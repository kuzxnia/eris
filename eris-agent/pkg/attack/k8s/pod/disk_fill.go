package pod

// DiskFillAttack is ...
type DiskFillAttack struct{}

func ProvideDiskFillAttack() *DiskFillAttack {
	return &DiskFillAttack{}
}

func (c *DiskFillAttack) Perform() {
}

func (c *DiskFillAttack) AttackType() string {
	return "disk-fill"
}
