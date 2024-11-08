package k8s

type CompositeResourceProvider struct {
	k8sClient []*KubernetesClient
}

func ProvideCompositeResourceProvider() *CompositeResourceProvider {
	return &CompositeResourceProvider{}
}

// todo: multiple clients
// from multiple contexts
// combine results
func (p *CompositeResourceProvider) GetResources(selector *ResourceSelector) ([]*Resource, error) {
	return nil, nil
}

func (p *CompositeResourceProvider) GetPods(selector *ResourceSelector) ([]Pod, error) {
	return nil, nil
}

func (p *CompositeResourceProvider) RunOnContainer(pod Pod, container string, cmd string) error {
	return nil
}

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
