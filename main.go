package main

import (
	"goK8sClient/pod"
)

func main() {
	// conn.Init()

	// namespace.GetNs("test")
	// namespace.GetNsList()
	// namespace.CreateNs("hello")
	// namespace.DeleteNs("hello")

	pod.GetPod("test", "xxx-xx-xxx-api-b456cf57b-lg9kp")
	// pod.GetPodList("test")
	// pod.CreatePod("hello")
	// pod.DeletePod("hello")
}
