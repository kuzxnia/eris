package pod

// ProcessFloodAction is ...
type ProcessFloodAction struct{}

func ProvideProcessFloodAction() *ProcessFloodAction {
	return &ProcessFloodAction{}
}

func (c *ProcessFloodAction) Perform() {
}

func (c *ProcessFloodAction) ActionType() string {
	return "process-flood"
}
