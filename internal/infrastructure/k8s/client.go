package k8s

type KubernetesClient struct {
	context   string
	namespace string
}

func ProvideKubernetesClient() *KubernetesClient {
	return &KubernetesClient{}
}
