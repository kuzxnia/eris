package node

// NodeSchedulingDisabledAttack is ...
type NodeSchedulingDisabledAttack struct{}

func ProvideNodeSchedulingDisabledAttack() *NodeSchedulingDisabledAttack {
	return &NodeSchedulingDisabledAttack{}
}

func (c *NodeSchedulingDisabledAttack) Perform() {
}

func (c *NodeSchedulingDisabledAttack) AttackType() string {
	return "node-scheduling-disabled"
}
