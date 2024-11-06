package probe

type ProbeOption func(*Probe)

type Probe struct {
	Type ProbeType
	Mode ProbeMode


}

func (p *Probe) ApplyOptions(options ...ProbeOption) {
	for _, option := range options {
		option(p)
	}
}

// http
// cmd
// k8s
type ProbeType interface{}

// SoT: Executed at the Start of Test as a pre-chaos check
// EoT: Executed at the End of Test as a post-chaos check
// Edge: Executed both, before and after the chaos
// Continuous: The probe is executed continuously, with a specified polling interval during the chaos injection.
// OnChaos: The probe is executed continuously, with a specified polling interval strictly for chaos duration of chaos

// pre-action
// post-action
// edge
// during-action
// full
type ProbeMode interface{}

// attributes
// decoreator/
