// Code generated by schema-generate. DO NOT EDIT.

package promenade

// Images
type Images2 struct {
	Pause string `json:"pause"`
}

// Root
type KubeletSpec struct {
	Arguments []string `json:"arguments,omitempty"`
	Images    *Images2 `json:"images"`
}
