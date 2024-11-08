package k8s

type ResourceProvider interface {
	GetResources(*ResourceSelector) ([]*Resource, error)
	GetPods(*ResourceSelector) ([]Pod, error)
	RunOnContainer(pod Pod, container string, cmd string) error
}


// todo: move this 
type Pod struct {
	Context string
	// todo: data or else
}

func (r *Pod) GetContainersNames() []string {
	return []string{"test"}
}

type Resource struct {
	context string
	// todo: data or else
}

func (r *Resource) GetName() string {
	return ""
}

type ResourceSelector struct{}
