package resource

import (
	"github.com/kuzxnia/eris/control-plane/pkg/agent"
)

type AgentResourceSelector struct {
	resourceClient ResourceClient
	repository     agent.AgentRepository
}

func ProvideAgentResourceSelector(resourceClient ResourceClient, repository agent.AgentRepository) *AgentResourceSelector {
	return &AgentResourceSelector{
		resourceClient: resourceClient,
		repository:     repository,
	}
}

func (s *AgentResourceSelector) Select(selector *Selector) ([]*Resource, error) {
	agents, err := s.repository.GetAgents()
	if err != nil {
		return nil, err
	}

	type ResourceResult struct {
		resources []*Resource
		err       error
	}

	resultChannel := make(chan *ResourceResult)
	for _, a := range agents {
		go func(agent *agent.Agent) {
			resources, err := s.resourceClient.GetApiV1Resource(agent.URL, selector)
			if err != nil {
				resultChannel <- &ResourceResult{
					resources: nil,
					err:       err,
				}
			}

			resultChannel <- &ResourceResult{
				resources: resources,
				err:       nil,
			}
		}(a)
	}

	resources := []*Resource{}
	for range agents {
		resourceFetchResult := <-resultChannel

		if resourceFetchResult.err != nil {
			return nil, resourceFetchResult.err
		}
		for _, resource := range resourceFetchResult.resources {
			resources = append(resources, resource)
		}
	}

  // todo: limit number of resource by
	// selector.PodAffectedPercentage

	return resources, nil
}

func (s *AgentResourceSelector) GetSelectorType() string {
	return "label_selector"
}
