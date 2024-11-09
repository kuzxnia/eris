package agent

import "github.com/kuzxnia/eris/control-plane/pkg/interfaces/web/dto"

type AgentMapper struct{}

func ProvideAgentMapper() *AgentMapper {
	return &AgentMapper{}
}

func (m *AgentMapper) Map(request *dto.AgentRequest) (*Agent, error) {
	return &Agent{
		Id:      request.Id,
		Context: request.Context,
	}, nil
}
