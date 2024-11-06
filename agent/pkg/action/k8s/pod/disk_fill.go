package pod

// DiskFillAction is ...
type DiskFillAction struct{}

func ProvideDiskFillAction() *DiskFillAction {
	return &DiskFillAction{}
}

func (c *DiskFillAction) Perform() {
}

func (c *DiskFillAction) ActionType() string {
	return "disk-fill"
}
