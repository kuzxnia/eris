package pod

// ContainerKillAttack is ...
type ContainerKillAttack struct{}

func ProvideContainerKillAttack() *ContainerKillAttack {
	return &ContainerKillAttack{}
}

func (c *ContainerKillAttack) Perform() {
}

func (c *ContainerKillAttack) AttackType() string {
	return "container-kill"
}
