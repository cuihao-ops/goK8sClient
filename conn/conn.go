package conn

import (
	"flag"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 入参为空，出参是配置文件的指针
func Init() *kubernetes.Clientset {
	var kubeconfig *string

	/*
		运行脚本格式如下：
		go run .\conn.go config D:\goK8sClient\config

		flag.String说明：第一个是传入的参数，第二个是命令行分隔符，第三个是传入参数的具体值
		flag.String("kubeconfig", "config", "D:\\goK8sClient\\config")
	*/
	kubeconfig = flag.String("kubeconfig", "config", "D:\\goK8sClient\\config")

	flag.Parse()

	// 绑定上面config配置文件
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 连接集群
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("成功连接k8s集群...")
	}

	return clientset

}
