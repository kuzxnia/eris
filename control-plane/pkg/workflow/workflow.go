package workflow

type Workflows struct {
	Contexts  Contexts
	Sources   Sources
	Workflows []Workflow
}

type Workflow struct {
	Name    string
	Actions []*Action
}

type Contexts map[string]any

type Sources map[string]any

type Action struct {
	Name          string
	Type          string
	Timeout       string
	Selector      *Selector
	Probe         *Probe
	ProcessName   *string
	ContainerName *string
}

type Probe struct {
	Type                   string
	Mode                   string
	Sources                []string
	Check                  string
	Frequency              string
	Timeout                string
	BlockAfterUntilSuccess bool
	BlockTimeout           string
}

type Selector struct {
	Type       string
	Label      *string
	Namespaces []string
	ActionName *string
	TargetPod  *string
}
