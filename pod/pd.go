package pod

import (
	"context"
	"encoding/json"
	"fmt"
	"goK8sClient/conn"
	"goK8sClient/util/convert"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	err error
)

func GetPodList(ns string) {
	podList, err := conn.Init().CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("这里有【%d】个pod在【%s】namespace 中...\n", len(podList.Items), ns)
	for _, v := range podList.Items {
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

func CreatePod(ns string, name string) {

	podContent := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app": "nginx",
			},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				v1.Container{
					Name:            "nginx",
					Image:           "registry.loan.com/library/nginx:latest",
					ImagePullPolicy: "IfNotPresent",
				}},
		},
	}

	_, err := conn.Init().CoreV1().Pods(ns).Create(context.TODO(), podContent, metav1.CreateOptions{})

	if err != nil {
		fmt.Println("创建pod失败...")
	} else {
		fmt.Println("创建pod成功...")
	}
}

func CreatePodYaml(fileName string, dirPath string, kind *v1.Pod) {

	data := convert.Init(fileName, dirPath)

	err = json.Unmarshal(data, kind)
	if err != nil {
		panic(err.Error())
	}

	ns := kind.ObjectMeta.Namespace

	_, err = conn.Init().CoreV1().Pods(ns).Create(context.TODO(), kind, metav1.CreateOptions{})

	if err != nil {
		fmt.Println("创建pod失败...")
	} else {
		fmt.Println("创建pod成功...")
	}
}

func DeletePod(ns string, name string) {
	err := conn.Init().CoreV1().Pods(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println("删除pod失败...")
	} else {
		fmt.Println("删除pod成功...")
	}
}
