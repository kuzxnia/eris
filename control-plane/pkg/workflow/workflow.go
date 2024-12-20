package workflow

import "time"

type Workflow struct {
	Id               string // todo: uuid
	Status           string
	Actions          []*ActionExecutionContext // different actoin with resources
	StartTs          time.Time
	LastTransitionTs time.Time
	Logs             []*struct {
		Timestamp time.Time
		Message   string
	}
	Config *WorkflowConfig
}

type ActionExecutionContext struct {
	Action           *Action
	Resources        []*Resource
	Status           string
	StartTs          time.Time
	LastTransitionTs time.Time
}

type ActionResult struct {
	Message string
	Err     error
}

type Resource struct {
	Name string
}

type WorkflowConfig struct {
	Name     string
	Contexts Contexts
	Sources  Sources
	Actions  []*Action
	Version  int
	// todo: handle
	IsDeleted bool
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
	Type                  string
	PodAffectedPercentage string
	Label                 string
	Namespaces            []string
	ActionName            string
	TargetPod             string
}
