package pod

import (
	"context"
	"fmt"
	"goK8sClient/conn"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodList(ns string) {
	nsList, err := conn.Init().CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("这里有【%d】个pod在【%s】namespace 中...\n", len(nsList.Items), ns)
	for _, v := range nsList.Items {
		fmt.Println(v.ObjectMeta.Name)
	}
}

func GetPod(ns string, name string) {
	podGet, err := conn.Init().CoreV1().Pods(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("【%s】pod在【%s】namespace没有找到...\n", name, ns)
	} else {
		fmt.Printf("Pod 【%s】的UID为【%s】状态为【%s】创建时间为【%s】...\n", podGet.Name, podGet.UID, podGet.Status.Phase, podGet.CreationTimestamp)
	}
}
