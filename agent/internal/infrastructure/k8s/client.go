package k8s

import (
	"context"

	metav1"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type KubernetesClient struct {
  // to musze wziąc jakoś z env, bo inclusterconfig nie rozumie koncepcji context
  // narazie działamy w obrębie jednego contextu wiec to nie problem, będzie z tym problem później

	context   string
	namespace string

	coreV1 v1.CoreV1Interface
}

func ProvideKubernetesClient() *KubernetesClient {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

  clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	return &KubernetesClient{
    coreV1: clientset.CoreV1(),
  }
}
