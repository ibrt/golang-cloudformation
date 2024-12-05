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

// ResourceDeletionPolicy describes a CloudFormation resource deletion policy.
type ResourceDeletionPolicy interface {
	GetResourceDeletionPolicy() string
	json.Marshaler
}

type resourceDeletionPolicy string

// GetResourceDeletionPolicy implements the ResourceDeletionPolicy interface.
func (p resourceDeletionPolicy) GetResourceDeletionPolicy() string {
	return string(p)
}

// MarshalJSON implements the ResourceDeletionPolicy interface.
func (p resourceDeletionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.GetResourceDeletionPolicy())
}

var (
	// ResourceDeletionPolicies provides quick access to the supported CloudFormation resource deletion policies.
	ResourceDeletionPolicies = struct {
		Delete               ResourceDeletionPolicy
		Snapshot             ResourceDeletionPolicy
		RetainExceptOnCreate ResourceDeletionPolicy
		Retain               ResourceDeletionPolicy
	}{
		Delete:               resourceDeletionPolicy("Delete"),
		Snapshot:             resourceDeletionPolicy("Snapshot"),
		RetainExceptOnCreate: resourceDeletionPolicy("RetainExceptOnCreate"),
		Retain:               resourceDeletionPolicy("Retain"),
	}
)

// ResourceUpdateReplacePolicy describes a CloudFormation resource update replace policy.
type ResourceUpdateReplacePolicy interface {
	GetResourceUpdateReplacePolicy() string
	json.Marshaler
}

type resourceUpdateReplacePolicy string

// GetResourceUpdateReplacePolicy implements the ResourceUpdateReplacePolicy interface.
func (p resourceUpdateReplacePolicy) GetResourceUpdateReplacePolicy() string {
	return string(p)
}

// MarshalJSON implements the ResourceUpdateReplacePolicy interface.
func (p resourceUpdateReplacePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.GetResourceUpdateReplacePolicy())
}

var (
	// ResourceUpdateReplacePolicies provides quick access to the supported update replace policies.
	ResourceUpdateReplacePolicies = struct {
		Delete   ResourceUpdateReplacePolicy
		Snapshot ResourceUpdateReplacePolicy
		Retain   ResourceUpdateReplacePolicy
	}{
		Delete:   resourceUpdateReplacePolicy("Delete"),
		Snapshot: resourceUpdateReplacePolicy("Snapshot"),
		Retain:   resourceUpdateReplacePolicy("Retain"),
	}
)
