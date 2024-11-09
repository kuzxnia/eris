package agent

type AgentRepository interface {
	// GetAgent(string) (*Agent, error)
	GetAgents() ([]*Agent, error)
	SaveAgent(*Agent) error
}
