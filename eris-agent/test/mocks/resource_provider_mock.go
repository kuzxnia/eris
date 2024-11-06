package mocks

import (
	"github.com/kuzxnia/eris/eris-agent/internal/infrastructure/k8s"
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
