package cfz

import (
	"encoding/json"
)

// ResourcePartialLogicalName describes a subset of a CloudFormation resource.
type ResourcePartialLogicalName interface {
	GetResourceLogicalName() string
}

// ResourceLogicalName is a string that implements the ResourcePartialLogicalName interface.
type ResourceLogicalName string

// GetResourceLogicalName implements the ResourcePartialLogicalName interface.
func (r ResourceLogicalName) GetResourceLogicalName() string {
	return string(r)
}

// Resource describes a CloudFormation resource.
type Resource interface {
	ResourcePartialLogicalName
	GetType() string
	json.Marshaler
}
