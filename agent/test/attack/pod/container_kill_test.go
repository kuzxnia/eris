package pod

import (
	"testing"

	"github.com/kuzxnia/eris/agent/internal/infrastructure/k8s"
	k8sPod "github.com/kuzxnia/eris/agent/pkg/action/k8s/pod"
	"github.com/kuzxnia/eris/agent/test/mocks"
	"github.com/stretchr/testify/mock"
)

func TestContainerKillAttack_Perform(t *testing.T) {
	// given
	// selector := k8s.ResourceSelector{}
	resourceProvider := new(mocks.TestResourceProvider)
	pod := k8s.Pod{Context: "test"}
	resourceProvider.On("GetPods", mock.Anything).Return([]k8s.Pod{pod}, nil)
	resourceProvider.On("RunOnContainer", pod, "test", "/sbin/killall5").Return(nil)

	attackCommand := k8sPod.ProvideContainerKillAction(resourceProvider)
	// todo: add config

	// when
	attackCommand.Perform()

	// then
	resourceProvider.AssertCalled(t, "RunOnContainer", pod, "test", "/sbin/killall5")
	resourceProvider.AssertExpectations(t)
}
