// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Command undocumented
type Command struct {
	// Entity is the base model of Command
	Entity
	// Status undocumented
	Status *string `json:"Status,omitempty"`
	// Type undocumented
	Type *string `json:"Type,omitempty"`
	// AppServiceName undocumented
	AppServiceName *string `json:"AppServiceName,omitempty"`
	// PackageFamilyName undocumented
	PackageFamilyName *string `json:"PackageFamilyName,omitempty"`
	// Error undocumented
	Error *string `json:"Error,omitempty"`
	// Payload undocumented
	Payload *PayloadRequestObject `json:"Payload,omitempty"`
	// PermissionTicket undocumented
	PermissionTicket *string `json:"PermissionTicket,omitempty"`
	// PostBackURI undocumented
	PostBackURI *string `json:"PostBackUri,omitempty"`
	// Responsepayload undocumented
	Responsepayload *PayloadResponse `json:"responsepayload,omitempty"`
}