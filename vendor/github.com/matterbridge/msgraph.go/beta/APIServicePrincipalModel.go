// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// APIServicePrincipal undocumented
type APIServicePrincipal struct {
	// Object is the base model of APIServicePrincipal
	Object
	// ResourceSpecificApplicationPermissions undocumented
	ResourceSpecificApplicationPermissions []ResourceSpecificPermission `json:"resourceSpecificApplicationPermissions,omitempty"`
}