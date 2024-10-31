package node

// NodeRestartAttack is ...
type NodeRestartAttack struct{}

func ProvideNodeRestartAttack() *NodeRestartAttack {
	return &NodeRestartAttack{}
}

func (c *NodeRestartAttack) Perform() {
}

func (c *NodeRestartAttack) AttackType() string {
	return "node-restart"
}
