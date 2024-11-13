package agent

import (
	"fmt"

	"github.com/kuzxnia/eris/control-plane/pkg"
	"github.com/kuzxnia/eris/control-plane/pkg/clients"
)

type AgentResourceSelector struct {
	resourceClient clients.ResourceClient
	repository     AgentRepository
}

func ProvideAgentResourceSelector(resourceClient clients.ResourceClient, repository AgentRepository) *AgentResourceSelector {
	return &AgentResourceSelector{
		resourceClient: resourceClient,
		repository:     repository,
	}
}

func (s *AgentResourceSelector) Select(selector *pkg.Selector) ([]*pkg.Resource, error) {
	agents, err := s.repository.GetAgents()
	if err != nil {
		return nil, err
	}

	type ResourceResult struct {
		resources []*pkg.Resource
		err       error
	}

	resultChannel := make(chan *ResourceResult)
	for _, agent := range agents {
		go func(agent *Agent) {
			resources, err := s.resourceClient.GetApiV1Resource(agent.URL, selector)
			// resp, err := s.httpClient.R().
			// 	SetResult(&pkg.Resource{}).
			// 	SetQueryParams(map[string]string{}).
			// 	Get(agent.URL + selectResourcesEndpoint)
			if err != nil {
				resultChannel <- &ResourceResult{
					resources: nil,
					err:       err,
				}
			}

			// if resp.IsSuccess() {
			resultChannel <- &ResourceResult{
				// resources: resp.Result().(*pkg.Resource),
				resources: resources,
				err:       nil,
			}
			// }
		}(agent)
	}

	resources := []*pkg.Resource{}
	for range agents {
		resourceFetchResult := <-resultChannel

		if resourceFetchResult.err != nil {
			err = resourceFetchResult.err
			break
		}

		fmt.Println("in for")
		for _, resource := range resourceFetchResult.resources {
			resources = append(resources, resource)
		}
	}
	fmt.Println("after")

	return resources, nil
}
