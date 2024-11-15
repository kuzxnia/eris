package agent

import (
	"github.com/kuzxnia/eris/control-plane/pkg/web/dto"
)

type AgentService struct {
	mapper     *AgentMapper
	repository AgentRepository
}

// where to refresh agents,
// heartbeat

func ProvideAgentService(mapper *AgentMapper, repository AgentRepository) *AgentService {
	return &AgentService{
		mapper:     mapper,
		repository: repository,
	}
}

func (s *AgentService) CreateAgent(agentRequest dto.AgentRequest) error {
	agent, _ := s.mapper.Map(&agentRequest)
	return s.repository.SaveAgent(agent)
}

func (s *AgentService) GetAgents() ([]*Agent, error) {
	// todo: from repository map
	// repository return entity
	return s.repository.GetAgents()
}
