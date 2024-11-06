package pod

// CpuHogAction is ...
type CpuHogAction struct{}

func ProvideCpuHogAction() *CpuHogAction {
	return &CpuHogAction{}
}

func (c *CpuHogAction) Perform() {
}

func (c *CpuHogAction) ActionType() string {
	return "cpu-hog"
}
