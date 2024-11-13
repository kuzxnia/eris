package pkg

type ResourceSelector interface {
	Select(*Selector) ([]*Resource, error)
}

type Resource struct {
	Name string `json:"name"`
}

type Selector struct {
	Type                  string
	PodAffectedPercentage string
	Label                 string
	Namespaces            []string
	ActionName            string
	TargetPod             string
}
