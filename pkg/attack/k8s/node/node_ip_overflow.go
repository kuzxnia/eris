package node

// NodeIpOverflowAttack is ...
type NodeIpOverflowAttack struct{}

func ProvideNodeIpOverflowAttack() *NodeIpOverflowAttack {
	return &NodeIpOverflowAttack{}
}

func (c *NodeIpOverflowAttack) Perform() {
}

func (c *NodeIpOverflowAttack) AttackType() string {
	return "node-ip-overflow"
}
