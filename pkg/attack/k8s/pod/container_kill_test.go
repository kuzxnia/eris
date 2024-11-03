package pod

import (
	"testing"

	"github.com/kuzxnia/eris/internal/infrastructure/k8s"
	"github.com/stretchr/testify/mock"
)

type TestResourceProvider struct {
	mock.Mock
}

func (p *TestResourceProvider) GetResources(selector *k8s.ResourceSelector) ([]*k8s.Resource, error) {
	return nil, nil
}

func (p *TestResourceProvider) GetPods(selector *k8s.ResourceSelector) ([]k8s.Pod, error) {
	args := p.Called(selector)
	return args.Get(0).([]k8s.Pod), args.Error(1)
}

func (p *TestResourceProvider) RunOnContainer(pod k8s.Pod, container string, cmd string) error {
  p.Called(pod, container, cmd)
	return nil
}

func TestContainerKillAttack_Perform(t *testing.T) {
	// given
	selector := k8s.ResourceSelector{}
	resourceProvider := new(TestResourceProvider)
	pod := k8s.Pod{Context: "test"}
	resourceProvider.On("GetPods", mock.Anything).Return([]k8s.Pod{pod}, nil)
	resourceProvider.On("RunOnContainer", pod, "test", "/sbin/killall5").Return(nil)

	attackCommand := ContainerKillAttack{
		resourceProvider: resourceProvider,
	}
	// todo: add config

	// when
	attackCommand.Perform(&selector)

	// then
	resourceProvider.AssertCalled(t, "RunOnContainer", pod, "test", "/sbin/killall5")
	resourceProvider.AssertExpectations(t)
}
