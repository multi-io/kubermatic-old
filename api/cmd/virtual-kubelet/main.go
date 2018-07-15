package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
)

const (
	podAnswer = `
{
  "items": [
    {
      "metadata": {
        "name": "k8s-auth-injector-6c67dbdb88-mltk4",
        "generateName": "k8s-auth-injector-6c67dbdb88-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/k8s-auth-injector-6c67dbdb88-mltk4",
        "uid": "2725c1db-7d37-11e8-b749-1c1b0d187168",
        "resourceVersion": "988497",
        "creationTimestamp": "2018-07-01T14:00:46Z",
        "labels": {
          "app": "k8s-auth-injector",
          "pod-template-hash": "2723868644"
        },
        "annotations": {
          "cni.projectcalico.org/podIP": "10.244.0.52/32",
          "kubernetes.io/config.seen": "2018-07-08T11:44:20.95467905Z",
          "kubernetes.io/config.source": "api"
        },
        "ownerReferences": [
          {
            "apiVersion": "extensions/v1beta1",
            "kind": "ReplicaSet",
            "name": "k8s-auth-injector-6c67dbdb88",
            "uid": "27222b1c-7d37-11e8-b749-1c1b0d187168",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "default-token-77clf",
            "secret": {
              "secretName": "default-token-77clf",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "k8s-auth-injector",
            "image": "alvaroaleman/k8s-auth-injector",
            "resources": {},
            "volumeMounts": [
              {
                "name": "default-token-77clf",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "j1900",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ]
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2018-07-01T14:00:46Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2018-07-01T14:00:50Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": null
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2018-07-01T14:00:46Z"
          }
        ],
        "hostIP": "192.168.0.39",
        "podIP": "10.244.0.52",
        "startTime": "2018-07-01T14:00:46Z",
        "containerStatuses": [
          {
            "name": "k8s-auth-injector",
            "state": {
              "running": {
                "startedAt": "2018-07-01T14:00:50Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "alvaroaleman/k8s-auth-injector:latest",
            "imageID": "docker-pullable://alvaroaleman/k8s-auth-injector@sha256:c631af4c68f589414c3711a7ff1858f8a7aea66588c5341040f3cab35d273694",
            "containerID": "docker://afc04556cadf512d04dc05a1fc245c1e54c0e75e1c1a10f27e0c585b88bdf44d"
          }
        ],
        "qosClass": "BestEffort"
      }
    }
  ]
}`
)

func main() {
	listenAddress := flag.String("listen-address", ":8443", "address to listen on")
	certFile := flag.String("certfile", "", "Path to the certfile")
	keyFile := flag.String("keyfile", "", "Path to the keyfile")
	flag.Parse()

	if *listenAddress == "" || *certFile == "" || *keyFile == "" {
		glog.Fatalf(`All of "-listen-address", "-certfile", and "-keyfile" must be non-empty!`)
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/pods", podHandler)
	glog.Fatalln(http.ListenAndServeTLS(*listenAddress, *certFile, *keyFile, nil))
}

func podHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("Got a request from %s")
	glog.Infof("Returning podAnswer to %s", r.RemoteAddr)
	w.Write([]byte(podAnswer))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("Got an unexpected request with uri=%s", r.URL.String())
	w.WriteHeader(http.StatusBadRequest)
}
