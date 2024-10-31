package node

// NodeDrainAttack is ...
type NodeDrainAttack struct{}

func ProvideNodeDrainAttack() *NodeDrainAttack {
	return &NodeDrainAttack{}
}

func (c *NodeDrainAttack) Perform() {
}

func (c *NodeDrainAttack) AttackType() string {
	return "node-drain"
}
