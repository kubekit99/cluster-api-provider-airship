// Code generated by schema-generate. DO NOT EDIT.

package shipyard

// GroupsItems
type GroupsItems struct {
	Critical        bool              `json:"critical"`
	DependsOn       []string          `json:"depends_on"`
	Name            string            `json:"name"`
	Selectors       []*SelectorsItems `json:"selectors"`
	SuccessCriteria *SuccessCriteria  `json:"success_criteria,omitempty"`
}

// Root
type DeploymentStrategySpec struct {
	Groups []*GroupsItems `json:"groups"`
}

// SelectorsItems
type SelectorsItems struct {
	NodeLabels []string `json:"node_labels,omitempty"`
	NodeNames  []string `json:"node_names,omitempty"`
	NodeTags   []string `json:"node_tags,omitempty"`
	RackNames  []string `json:"rack_names,omitempty"`
}

// SuccessCriteria
type SuccessCriteria struct {
	MaximumFailedNodes     int `json:"maximum_failed_nodes,omitempty"`
	MinimumSuccessfulNodes int `json:"minimum_successful_nodes,omitempty"`
	PercentSuccessfulNodes int `json:"percent_successful_nodes,omitempty"`
}
