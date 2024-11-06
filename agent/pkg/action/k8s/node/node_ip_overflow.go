package node

// NodeIpOverflowAction is ...
type NodeIpOverflowAction struct{}

func ProvideNodeIpOverflowAction() *NodeIpOverflowAction {
	return &NodeIpOverflowAction{}
}

func (c *NodeIpOverflowAction) Perform() {
}

func (c *NodeIpOverflowAction) ActionType() string {
	return "node-ip-overflow"
}
