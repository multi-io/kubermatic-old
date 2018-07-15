package virtualkubelet

import (
	"time"

	"github.com/golang/glog"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Controller) keepNodeReady() {
	for {
		time.Sleep(2 * time.Second)

		currentTime := metav1.Now()
		newNodeReadyCondition := corev1.NodeCondition{
			Type:              corev1.NodeReady,
			Status:            corev1.ConditionTrue,
			Reason:            "KubeletReady",
			Message:           "kubelet is posting ready status",
			LastHeartbeatTime: currentTime,
		}
		//node, err := c.nodesLister.Get(NodeName)
		node, err := c.kubeClient.CoreV1().Nodes().Get(NodeName, metav1.GetOptions{})
		if err != nil {
			glog.Errorf("Failed to get node: %v", err)
			continue
		}
		node.Status.Conditions = []corev1.NodeCondition{newNodeReadyCondition}
		if _, err := c.kubeClient.CoreV1().Nodes().UpdateStatus(node); err != nil {
			glog.Errorf("Failed to set node ready condition: %v", err)
		}
	}
}
