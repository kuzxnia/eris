package repository

import (
	"errors"

	"github.com/kuzxnia/eris/control-plane/pkg/agent"
	"github.com/samber/lo"
)

type InMemoryAgentRepository struct {
	agents map[string]*agent.Agent
}

func ProvideInMemoryAgentRepository() *InMemoryAgentRepository {
	return &InMemoryAgentRepository{
		agents: map[string]*agent.Agent{},
	}
}

func (r *InMemoryAgentRepository) SaveAgent(agent *agent.Agent) error {
	_, isPresent := r.agents[agent.Id]
	if isPresent {
		return errors.New("Agent already exists")
	}

	r.agents[agent.Id] = agent
	return nil
}

func (r *InMemoryAgentRepository) GetAgents() ([]*agent.Agent, error) {
	return lo.Values(r.agents), nil
}
