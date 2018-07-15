package virtualkubelet

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"

	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	listerscorev1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	NodeName = "kubermatic-virtual-kubelet"
)

type Controller struct {
	kubeClient  kubernetes.Interface
	podsLister  listerscorev1.PodLister
	nodesLister listerscorev1.NodeLister
	workqueue   workqueue.RateLimitingInterface
	managedPods map[string]*corev1.Pod
	mapLock     sync.Mutex
}

func New(
	kubeClient kubernetes.Interface,
	podInformer cache.SharedIndexInformer,
	podLister listerscorev1.PodLister,
	nodeLister listerscorev1.NodeLister) *Controller {

	controller := &Controller{
		kubeClient:  kubeClient,
		workqueue:   workqueue.NewRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(1*time.Second, 5*time.Minute)),
		podsLister:  podLister,
		nodesLister: nodeLister,
		managedPods: map[string]*corev1.Pod{},
	}

	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueuePod,
		UpdateFunc: func(_, new interface{}) {
			controller.enqueuePod(new)
		},
	})

	return controller
}

func (c *Controller) SetPod(pod *corev1.Pod) {
	c.managedPods[pod.Name] = pod
	c.workqueue.AddRateLimited(pod.Name)
}

func (c *Controller) handlePod(pod *corev1.Pod) {
	if pod.Namespace != metav1.NamespaceSystem {
		return
	}

	for managedPod := range c.managedPods {
		if managedPod == pod.Name {
			c.workqueue.AddRateLimited(pod.Name)
			return
		}
	}
}

// Run starts the control loop
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	go c.keepNodeReady()
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	return nil
}

func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	key, quit := c.workqueue.Get()
	if quit {
		return false
	}

	defer c.workqueue.Done(key)

	err := c.syncHandler(key.(string))
	if err == nil {
		c.workqueue.Forget(key)
		return true
	}

	utilruntime.HandleError(fmt.Errorf("%v failed with: %v", key, err))
	c.workqueue.AddRateLimited(key)

	return true
}

func (c *Controller) enqueuePod(obj interface{}) {
	if pod, isPod := obj.(*corev1.Pod); isPod {
		c.workqueue.AddRateLimited(pod.Name)
	}
	utilruntime.HandleError(fmt.Errorf("Got an object that was not a pod: %+v", obj))
}

func (c *Controller) syncHandler(key string) error {
	_, err := c.podsLister.Pods(metav1.NamespaceSystem).Get(key)
	if err != nil {
		if kerrors.IsNotFound(err) {
			return c.createPod(key)
		}
		return err
	}
	return nil
}

func (c *Controller) createPod(key string) error {

	// Ensure pod is in the kube-system namespace even if the client
	// did not explicitly set that
	c.managedPods[key].Namespace = metav1.NamespaceSystem
	// Ensure pod is scheduled on the virtual kubelet even if the client
	// did not explicitly set that
	c.managedPods[key].Spec.NodeName = NodeName

	// Pod must have a "kubernetes.io/config.mirror" annotation
	if c.managedPods[key].Annotations == nil {
		c.managedPods[key].Annotations = map[string]string{}
	}
	c.managedPods[key].Annotations["kubernetes.io/config.mirror"] = key

	createdPod, err := c.kubeClient.CoreV1().Pods(metav1.NamespaceSystem).Create(c.managedPods[key])
	if err != nil {
		glog.V(4).Infof("Failed to create pod %s: %v", c.managedPods[key].Name, err)
		return err
	}
	glog.V(4).Infof("Successfully created pod %s", createdPod.Name)
	c.mapLock.Lock()
	defer c.mapLock.Unlock()
	c.managedPods[key] = createdPod

	return nil
}
