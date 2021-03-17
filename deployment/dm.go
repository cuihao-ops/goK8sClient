package deployment

import (
	"context"
	"fmt"
	"goK8sClient/conn"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	err error
)

func GetDmList(ns string) {
	dmList, err := conn.Init().AppsV1().Deployments(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("这里有【%d】个deployment在【%s】namespace 中...\n", len(dmList.Items), ns)
	for _, v := range dmList.Items {
		fmt.Println(v.ObjectMeta.Name)
	}
}

func GetDm(ns string, name string) {
	dmGet, err := conn.Init().AppsV1().Deployments(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		fmt.Printf("【%s】deployment在【%s】namespace没有找到...\n", name, ns)
	} else {
		fmt.Printf("deployment 【%s】的UID为【%s】创建时间为【%s】可用副本【%d】...\n", dmGet.Name, dmGet.UID, dmGet.CreationTimestamp, dmGet.Status.AvailableReplicas)
	}
}
