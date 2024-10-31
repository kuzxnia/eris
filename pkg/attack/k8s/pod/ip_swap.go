package pod

// IpSwapAttack is ...
type IpSwapAttack struct{}

func ProvideIpSwapAttack() *IpSwapAttack {
	return &IpSwapAttack{}
}

func (c *IpSwapAttack) Perform() {
}

func (c *IpSwapAttack) AttackType() string {
	return "ip-swap"
}
