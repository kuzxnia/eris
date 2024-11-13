package resourceselector

import "go.uber.org/fx"

type ResourceSelector interface {
	Select(*Selector) ([]*Resource, error)
	GetSelectorType() string
}

type Resource struct {
	Name string `json:"name"`
}

type Selector struct {
	// "type": "label_selector",
	// "type": "previous_action",
	Type                  string
	PodAffectedPercentage string
	Label                 string
	Namespaces            []string
	ActionName            string
	TargetPod             string
}

var Module = fx.Options(
	fx.Provide(ProvideHttpResourceClient),
	fx.Provide(ProvideAgentResourceSelector),
)
