package k8s

import (
	"context" // ✅ 添加 context 包
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

// 获取 Kubernetes Client
func getKubernetesClient() (*kubernetes.Clientset, dynamic.Interface, error) {
	var config *rest.Config
	var err error

	// 判断是否有 KUBERNETES_SERVICE_HOST 环境变量，以此判断是否在 Kubernetes 集群内部
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		// 集群内使用 inClusterConfig
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, nil, fmt.Errorf("error creating in-cluster config: %v", err)
		}
	} else {
		// 本地开发使用 kubeconfig
		kubeconfigPath := "D:\\\\desktop\\\\config.txt" // 替换为你的 kubeconfig 路径
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			return nil, nil, fmt.Errorf("error building config from kubeconfig: %v", err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Kubernetes client: %v", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating dynamic client: %v", err)
	}

	return clientset, dynamicClient, nil
}

func Checkclientstatus() {
	// 获取 Kubernetes 客户端
	clientset, dynamicClient, err := getKubernetesClient()
	if err != nil {
		log.Fatalf("❌ 连接 Kubernetes 失败: %v", err)
	}

	// 测试 1: 获取 Nodes
	fmt.Println("🔎 正在获取 Kubernetes 集群中的 Nodes...")
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{}) // ✅ 添加 context.TODO()
	if err != nil {
		log.Fatalf("❌ 获取 Nodes 失败: %v", err)
	}
	fmt.Printf("✅ 成功连接集群！共找到 %d 个节点。\n", len(nodes.Items))
	for _, node := range nodes.Items {
		fmt.Printf(" - 节点名称: %s\n", node.Name)
	}

	// 测试 2: 获取 default 命名空间中的 Pods
	fmt.Println("\n🔎 正在获取 'monitor" +
		"' 命名空间中的 Pods...")
	pods, err := clientset.CoreV1().Pods("monitor").List(context.TODO(), v1.ListOptions{}) // ✅ 添加 context.TODO()
	if err != nil {
		log.Fatalf("❌ 获取 Pods 失败: %v", err)
	}
	fmt.Printf("✅ 找到 %d 个 Pods。\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf(" - Pod: %s (状态: %s)\n", pod.Name, pod.Status.Phase)
	}

	// 测试 3: 使用 dynamicClient 查询 Deployments
	fmt.Println("\n🔎 正在获取 'monitor' 命名空间中的 Deployments...")
	deploymentsResource := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}                   // ✅ 修正类型
	deployments, err := dynamicClient.Resource(deploymentsResource).Namespace("monitor").List(context.TODO(), v1.ListOptions{}) // ✅ 添加 context.TODO()
	if err != nil {
		log.Fatalf("❌ 获取 Deployments 失败: %v", err)
	}
	fmt.Printf("✅ 找到 %d 个 Deployments。\n", len(deployments.Items))
	for _, deploy := range deployments.Items {
		fmt.Printf(" - Deployment: %s\n", deploy.GetName())
	}
}
