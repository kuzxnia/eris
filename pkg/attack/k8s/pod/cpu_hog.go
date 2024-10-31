package pod

// CpuHogAttack is ...
type CpuHogAttack struct{}

func ProvideCpuHogAttack() *CpuHogAttack {
	return &CpuHogAttack{}
}

func (c *CpuHogAttack) Perform() {
}

func (c *CpuHogAttack) AttackType() string {
	return "cpu-hog"
}
