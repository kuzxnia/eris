package dto

// workflowConfigRequest
// workflowConfigRequest
type WorkflowRequest struct{}

type WorkflowConfigRequest struct {
	Name     string           `json:"name,omitempty"`
	Actions  []*ActionRequest `json:"actions,omitempty"`
	Contexts ContextsRequest  `json:"contexts"`
	Sources  SourcesRequest   `json:"sources"`
}

type ContextsRequest map[string]any

type SourcesRequest map[string]any

type ActionRequest struct {
	Name          string           `json:"name,omitempty"`
	Type          string           `json:"type,omitempty"`
	Timeout       string           `json:"timeout,omitempty"`
	Selector      *SelectorRequest `json:"selector,omitempty"`
	Probe         *ProbeRequest    `json:"probe,omitempty"`
	ProcessName   *string          `json:"process_name,omitempty"`
	ContainerName *string          `json:"container_name,omitempty"`
}

type ProbeRequest struct {
	Type                   string   `json:"type,omitempty"`
	Mode                   string   `json:"mode,omitempty"`
	Sources                []string `json:"sources,omitempty"`
	Check                  string   `json:"check,omitempty"`
	Frequency              string   `json:"frequency,omitempty"`
	Timeout                string   `json:"timeout,omitempty"`
	BlockAfterUntilSuccess bool     `json:"block_after_until_success,omitempty"`
	BlockTimeout           string   `json:"block_timeout,omitempty"`
}

type SelectorRequest struct {
	Type                  string   `json:"type,omitempty"`
	PodAffectedPercentage string   `json:"pod_affected_percentage"`
	Label                 string   `json:"label,omitempty"`
	Namespaces            []string `json:"namespaces,omitempty"`
	ActionName            string   `json:"action_name,omitempty"`
	TargetPod             string   `json:"target_pod,omitempty"`
}
