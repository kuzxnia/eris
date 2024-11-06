package node

// NodeDrainAction is ...
type NodeDrainAction struct{}

func ProvideNodeDrainAction() *NodeDrainAction {
	return &NodeDrainAction{}
}

func (c *NodeDrainAction) Perform() {
}

func (c *NodeDrainAction) ActionType() string {
	return "node-drain"
}
