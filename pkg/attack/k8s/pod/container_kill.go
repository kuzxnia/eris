package pod

import (
	"fmt"

	"github.com/kuzxnia/eris/internal/infrastructure/k8s"
)

// ContainerKillAttack is ...
type ContainerKillAttack struct {
	resourceProvider k8s.ResourceProvider
}

func ProvideContainerKillAttack(
	resourceProvider k8s.ResourceProvider,
) *ContainerKillAttack {
	return &ContainerKillAttack{
		resourceProvider: resourceProvider,
	}
}

// beform should have context as param or build inside ???
func (c *ContainerKillAttack) Perform(selector *k8s.ResourceSelector) error {
  // todo: build context or get from config
	// selector := k8s.ResourceSelector{}

	pods, err := c.resourceProvider.GetPods(selector)
	if err != nil {
		return err
	}
	// todo: perform logic related to custom % of
	// - pod affected percentage
	//     (how many resources will be affected from selected app-resource)
	// - target pod
	//     - with regex
	// - target container
	//     - with regex (multiple)

	// or raise that no container found
  fmt.Println("pods", pods)
	pod := pods[0]

	// todo: use some filters to determine target pod or raise that no container found
	containers := pod.GetContainersNames()

  fmt.Println(pod, containers[0], "/sbin/killall5")
	return c.resourceProvider.RunOnContainer(pod, containers[0], "/sbin/killall5")

	// return nil
}

func (c *ContainerKillAttack) AttackType() string {
	return "container-kill"
}
