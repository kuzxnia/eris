package node

// NodeSchedulingDisabledAction is ...
type NodeSchedulingDisabledAction struct{}

func ProvideNodeSchedulingDisabledAction() *NodeSchedulingDisabledAction {
	return &NodeSchedulingDisabledAction{}
}

func (c *NodeSchedulingDisabledAction) Perform() {
}

func (c *NodeSchedulingDisabledAction) ActionType() string {
	return "node-scheduling-disabled"
}
