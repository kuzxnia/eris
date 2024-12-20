package repository

import (
	"github.com/kuzxnia/eris/control-plane/pkg/exception"
	"github.com/kuzxnia/eris/control-plane/pkg/workflow"
)

type InMemoryWorkflowConfigRepository struct {
	configs map[string]*workflow.WorkflowConfig
}

func ProvideInMemoryWorkflowConfigRepository() *InMemoryWorkflowConfigRepository {
	return &InMemoryWorkflowConfigRepository{
		configs: map[string]*workflow.WorkflowConfig{},
	}
}

func (r *InMemoryWorkflowConfigRepository) CreateConfig(cfg *workflow.WorkflowConfig) error {
	_, isPresent := r.configs[cfg.Name]
	if isPresent {
		return exception.ResourceAlreadyExistsError{}
	}

	cfg.Version = 1
	r.configs[cfg.Name] = cfg
	return nil
}

func (r *InMemoryWorkflowConfigRepository) UpdateWorkflow(cfg *workflow.WorkflowConfig) error {
	dbCfg, isPresent := r.configs[cfg.Name]
	if isPresent {
		cfg.Version = dbCfg.Version + 1
	} else {
		return exception.NotFoundError{}
	}
	r.configs[cfg.Name] = cfg
	return nil
}

func (r *InMemoryWorkflowConfigRepository) GetConfig(cfgName string) (*workflow.WorkflowConfig, error) {
	cfg, isPresent := r.configs[cfgName]
	if !isPresent {
		return nil, exception.NotFoundError{}
	}

	return cfg, nil
}
