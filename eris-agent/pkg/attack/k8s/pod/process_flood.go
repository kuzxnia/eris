package pod

// ProcessFloodAttack is ...
type ProcessFloodAttack struct{}

func ProvideProcessFloodAttack() *ProcessFloodAttack {
	return &ProcessFloodAttack{}
}

func (c *ProcessFloodAttack) Perform() {
}

func (c *ProcessFloodAttack) AttackType() string {
	return "process-flood"
}
