package pod

// IpSwapAction is ...
type IpSwapAction struct{}

func ProvideIpSwapAction() *IpSwapAction {
	return &IpSwapAction{}
}

func (c *IpSwapAction) Perform() {
}

func (c *IpSwapAction) ActionType() string {
	return "ip-swap"
}
