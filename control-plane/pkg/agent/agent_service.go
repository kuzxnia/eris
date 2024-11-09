package agent

import (
	"github.com/kuzxnia/eris/control-plane/pkg/interfaces/web/dto"
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
	return s.repository.GetAgents()
}