package getnamespace

import (
	"context"
	"fmt"
	"goK8sClient/conn"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNsList() {
	nsList, err := conn.Init().CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("这里有【%d】个namespace在集群中...\n", len(nsList.Items))
}

func GetNs() {
	ns1 := "test05"
	nsGet, err := conn.Init().CoreV1().Namespaces().Get(context.TODO(), ns1, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("【%s】在namespace没有找到...\n", ns1)
	} else {
		fmt.Printf("Namespace【%s】的UID为【%s】状态为【%s】创建时间为【%s】...\n", nsGet.Name, nsGet.UID, nsGet.Status.Phase, nsGet.CreationTimestamp)
	}
}

func CreateNs() {
	name := "hello"
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	_, err := conn.Init().CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("创建namespace失败...")
	} else {
		fmt.Println("创建namespace成功...")
	}
}
