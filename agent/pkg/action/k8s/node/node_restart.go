package node

// NodeRestartAction is ...
type NodeRestartAction struct{}

func ProvideNodeRestartAction() *NodeRestartAction {
	return &NodeRestartAction{}
}

func (c *NodeRestartAction) Perform() {
}

func (c *NodeRestartAction) ActionType() string {
	return "node-restart"
}
