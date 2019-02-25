package apiserver

import (
	"crypto/x509"
	"fmt"
	"net"

	"github.com/kubermatic/kubermatic/api/pkg/resources"

	corev1 "k8s.io/api/core/v1"
	certutil "k8s.io/client-go/util/cert"
)

// TLSServingCertificateCreator returns a function to create/update the secret with the apiserver tls certificate used to serve https
func TLSServingCertificateCreator(data resources.SecretDataProvider) resources.NamedSecretCreatorGetter {
	return func() (string, resources.SecretCreator) {
		return resources.ApiserverTLSSecretName, func(se *corev1.Secret) (*corev1.Secret, error) {
			if se.Data == nil {
				se.Data = map[string][]byte{}
			}

			ca, err := data.GetRootCA()
			if err != nil {
				return nil, fmt.Errorf("failed to get cluster ca: %v", err)
			}

			externalIP, err := data.ExternalIP()
			if err != nil {
				return nil, fmt.Errorf("failed to get external IP for cluster: %v", err)
			}

			inClusterIP, err := resources.InClusterApiserverIP(data.Cluster())
			if err != nil {
				return nil, fmt.Errorf("failed to get the in-cluster ClusterIP for the apiserver: %v", err)
			}

			altNames := certutil.AltNames{
				DNSNames: []string{
					// ExternalName
					data.Cluster().Address.ExternalName,
					// User facing
					"kubernetes",
					"kubernetes.default",
					"kubernetes.default.svc",
					fmt.Sprintf("kubernetes.default.svc.%s", data.Cluster().Spec.ClusterNetwork.DNSDomain),
					// Internal - apiserver-external
					resources.ApiserverExternalServiceName,
					fmt.Sprintf("%s.%s", resources.ApiserverExternalServiceName, data.Cluster().Status.NamespaceName),
					fmt.Sprintf("%s.%s.svc", resources.ApiserverExternalServiceName, data.Cluster().Status.NamespaceName),
					fmt.Sprintf("%s.%s.svc.cluster.local", resources.ApiserverExternalServiceName, data.Cluster().Status.NamespaceName),
					// Internal - apiserver
					resources.ApiserverInternalServiceName,
					fmt.Sprintf("%s.%s", resources.ApiserverInternalServiceName, data.Cluster().Status.NamespaceName),
					fmt.Sprintf("%s.%s.svc", resources.ApiserverInternalServiceName, data.Cluster().Status.NamespaceName),
					fmt.Sprintf("%s.%s.svc.cluster.local", resources.ApiserverInternalServiceName, data.Cluster().Status.NamespaceName),
				},
				IPs: []net.IP{
					*externalIP,
					*inClusterIP,
				},
			}

			if b, exists := se.Data[resources.ApiserverTLSCertSecretKey]; exists {
				certs, err := certutil.ParseCertsPEM(b)
				if err != nil {
					return nil, fmt.Errorf("failed to parse certificate (key=%s) from existing secret: %v", resources.ApiserverTLSCertSecretKey, err)
				}

				if resources.IsServerCertificateValidForAllOf(certs[0], "kube-apiserver", altNames, ca.Cert) {
					return se, nil
				}
			}

			key, err := certutil.NewPrivateKey()
			if err != nil {
				return nil, fmt.Errorf("unable to create a server private key: %v", err)
			}

			config := certutil.Config{
				CommonName: "kube-apiserver",
				AltNames:   altNames,
				Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
			}

			cert, err := certutil.NewSignedCert(config, key, ca.Cert, ca.Key)
			if err != nil {
				return nil, fmt.Errorf("unable to sign the server certificate: %v", err)
			}

			se.Data[resources.ApiserverTLSKeySecretKey] = certutil.EncodePrivateKeyPEM(key)
			se.Data[resources.ApiserverTLSCertSecretKey] = certutil.EncodeCertPEM(cert)

			return se, nil
		}
	}
}
