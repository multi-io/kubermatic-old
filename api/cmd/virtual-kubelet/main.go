package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"reflect"
	"time"

	"github.com/kubermatic/kubermatic/api/pkg/virtualkubelet"

	"github.com/golang/glog"

	corev1 "k8s.io/api/core/v1"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	podAnswer = `
{
  "items": [
    {
      "metadata": {
        "namespace": "kube-system",
        "annotations": {
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
        "containers": [
          {
            "name": "k8s-auth-injector",
            "image": "alvaroaleman/k8s-auth-injector",
            "resources": {},
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeName": "kubermatic-virtual-kubelet",
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
	kubeconfig := flag.String("kubeconfig", "", "Path to the kubeconfig, not required if running in-cluster")
	masterURL := flag.String("master", "", "The address of the Kubernetes API server")
	flag.Parse()

	if *listenAddress == "" || *certFile == "" || *keyFile == "" || *kubeconfig == "" {
		glog.Fatalf(`All of "-listen-address", "-certfile", "-keyfile" and "-kubeconfig" must be non-empty!`)
	}

	cfg, err := clientcmd.BuildConfigFromFlags(*masterURL, *kubeconfig)
	if err != nil {
		glog.Fatalf("error building kubeconfig: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("error building kubernetes clientset for kubeClient: %v", err)
	}

	var podList corev1.PodList
	if err := json.Unmarshal([]byte(podAnswer), &podList); err != nil {
		glog.Fatalf("Failed to unmarshal static podlist: %v", err)
	}

	stopChannel := make(chan struct{})

	informerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*10)
	informerFactory.Start(stopChannel)
	glog.V(4).Infoln("Starting to sync...")
	for _, syncsMap := range []map[reflect.Type]bool{informerFactory.WaitForCacheSync(stopChannel)} {
		for key, synced := range syncsMap {
			if !synced {
				glog.Fatalf("unable to sync %s", key)
			}
		}
	}
	glog.V(4).Infoln("Finished syncing")

	glog.V(4).Infoln("Creating controller...")
	controller := virtualkubelet.New(kubeClient,
		informerFactory.Core().V1().Pods().Informer(),
		informerFactory.Core().V1().Pods().Lister(),
		informerFactory.Core().V1().Nodes().Lister())
	glog.V(4).Infoln("Successfully created controller")
	glog.V(4).Infoln("Starting controller...")
	go func() {
		if err := controller.Run(4, stopChannel); err != nil {
			glog.Fatalf("Failed to start controller: %v", err)
		}
	}()
	glog.V(4).Infoln("Successfully started controller")
	virtualKubeletPod := podList.Items[0]
	virtualKubeletPod.Name = "virtual-kubelet-testpod"
	virtualKubeletPod.Spec.Containers = []corev1.Container{{Name: "testcontainer", Image: "testimage"}}
	controller.SetPod(&virtualKubeletPod)

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
