package main

import (
	"goK8sClient/pod"

	v1 "k8s.io/api/core/v1"
)

func main() {
	// conn.Init()

	// namespace.GetNs("test")
	// namespace.GetNsList()
	// namespace.CreateNs("hello")
	// namespace.DeleteNs("hello")

	// pod.GetPod("test", "xxx-xx-xxx-api-b456cf57b-lg9kp")
	// pod.GetPodList("test")
	// pod.CreatePod("hello", "nginx-test")
	// pod.DeletePod("hello", "nginx-test")

	// convert.Init("nginx.yaml", &v1.Pod{})
	pod.CreatePodYaml("nginx.yaml", "D:\\goK8sClient\\yaml\\", &v1.Pod{})
}
