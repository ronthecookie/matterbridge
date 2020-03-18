// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ExternalConnection undocumented
type ExternalConnection struct {
	// Entity is the base model of ExternalConnection
	Entity
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Configuration undocumented
	Configuration *Configuration `json:"configuration,omitempty"`
	// Schema undocumented
	Schema *Schema `json:"schema,omitempty"`
	// Items undocumented
	Items []ExternalItem `json:"items,omitempty"`
	// Operations undocumented
	Operations []ConnectionOperation `json:"operations,omitempty"`
}