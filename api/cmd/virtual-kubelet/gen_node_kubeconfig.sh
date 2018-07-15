#!/usr/bin/env bash

set -eu

die() {
	echo $@
	exit 1
}

[[ $1 == "" ]] && die "must pass nodename!"
[[ -e ca.crt ]] || die "ca.crt file must be present in current dir!"
[[ -e ca.key ]] || die "ca.key file must be present in current dir!"
[[ -e ../kubelet.conf ]] || die "file ../kubelet.conf must be present!"
export NODENAME=$1

exit() {
  rm -f node-${NODENAME}.csr node-${NODENAME}.key node-${NODENAME}-openssl.cfg node-${NODENAME}.crt
}
trap exit EXIT

cat <<EOF > node-${1}-openssl.cfg
[req]
prompt = no
req_extensions = v3_req
distinguished_name = req_distinguished_name

[req_distinguished_name]
C=US
ST=CA
O=system:nodes
CN=system:node:${1}
emailAddress = root@kubernetes.io

[v3_req]
basicConstraints = CA:FALSE
keyUsage         = digitalSignature,keyEncipherment
extendedKeyUsage = clientAuth
EOF


openssl genrsa -out node-${1}.key 1024
openssl req -new -key node-${1}.key -out node-${1}.csr -config node-${1}-openssl.cfg
openssl x509 -req -days 3650 -in node-${1}.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out node-${1}.crt -extfile node-${1}-openssl.cfg

echo "Successfully created certificates!"

cat <<EOF >node-${1}.kubeconfig
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: $(cat ca.crt|base64 -w0)
    server: $(cat ../kubelet.conf |grep server|awk '{ print $2 }')
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: system:node:$1
  name: system:node:$1@kubernetes
current-context: system:node:$1@kubernetes
kind: Config
preferences: {}
users:
- name: system:node:$1
  user:
    client-certificate-data:  $(cat node-${1}.crt |base64 -w0)
    client-key-data: $(cat node-${1}.key|base64 -w0)
EOF

echo "Successfully finished!"
