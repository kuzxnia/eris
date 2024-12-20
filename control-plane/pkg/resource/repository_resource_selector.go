package resource

type RepositoryResourceSelector struct{}

// todo:
// this is not related to workflow but to execution 
// repository for job execution,
// workflow is just scaffold how workflow will be executed
// like facilyty and with history and soo
// like backup execution
// or rename workflow as workflow config

func ProvideRepositoryResourceSelector() *RepositoryResourceSelector {
	return &RepositoryResourceSelector{}
}

func (s *RepositoryResourceSelector) Select(selector *Selector) ([]*Resource, error) {
	return nil, nil
}

func (s *RepositoryResourceSelector) GetSelectorType() string {
	return "previous_action"
}
