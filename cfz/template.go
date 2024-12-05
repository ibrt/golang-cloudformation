package cfz

import (
	"encoding/json"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/memz"
)

// Template describes a template.
type Template struct {
	resourceMap   map[string]Resource
	resourceSlice []Resource
}

// NewTemplate initializes a new template.
func NewTemplate() *Template {
	return &Template{
		resourceMap:   make(map[string]Resource),
		resourceSlice: make([]Resource, 0),
	}
}

// AddResource adds a resource to the template.
func (t *Template) AddResource(resource Resource) error {
	t.resourceMap[resource.GetResourceLogicalName()] = resource
	t.resourceSlice = append(t.resourceSlice, resource)
	return nil
}

// MustAddResource is like AddResource, but panics on error.
func (t *Template) MustAddResource(resource Resource) *Template {
	errorz.MaybeMustWrap(t.AddResource(resource))
	return t
}

// MarshalJSON implements the json.Marshaler interface.
func (t *Template) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		AWSTemplateFormatVersion string              `json:"AWSTemplateFormatVersion,omitempty"`
		Resources                map[string]Resource `json:"Resources,omitempty"`
	}{
		AWSTemplateFormatVersion: "2010-09-09",
		Resources:                memz.ShallowCopyMap(t.resourceMap),
	})
}
