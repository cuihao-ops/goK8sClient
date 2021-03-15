package getnamespace

import (
	"context"
	"fmt"
	"goK8sClient/conn"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNaList() {
	nsList, err := conn.Init().CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("这里有【%d】个namespace在集群中...\n", len(nsList.Items))
}

func GetNa() {
	ns1 := "test05"
	nsGet, err := conn.Init().CoreV1().Namespaces().Get(context.TODO(), ns1, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("【%s】在namespace没有找到...\n", ns1)
	} else {
		fmt.Printf("Namespace【%s】的UID为【%s】状态为【%s】创建时间为【%s】...\n", nsGet.Name, nsGet.UID, nsGet.Status.Phase, nsGet.CreationTimestamp)
	}
}
