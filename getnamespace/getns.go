package getnamespace

import (
	"context"
	"fmt"
	"goK8sClient/conn"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNa() {
	pods, err := conn.Init().CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("这里有 %d 个pod在集群中...\n", len(pods.Items))

}
