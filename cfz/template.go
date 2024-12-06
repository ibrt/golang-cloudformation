package cfz

import (
	"encoding/json"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/memz"
)

// Template describes a template.
type Template struct {
	resources map[string]Resource
	outputs   map[string]Output
}

// NewTemplate initializes a new template.
func NewTemplate() *Template {
	return &Template{
		resources: make(map[string]Resource),
	}
}

// AddResource adds a resource to the template.
func (t *Template) AddResource(r Resource) error {
	t.resources[r.GetResourceLogicalName()] = r
	return nil
}

// MustAddResource is like AddResource, but panics on error.
func (t *Template) MustAddResource(r Resource) *Template {
	errorz.MaybeMustWrap(t.AddResource(r))
	return t
}

// AddOutput adds an output to the template.
func (t *Template) AddOutput(o Output) error {
	t.outputs[o.GetOutputLogicalName()] = o
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (t *Template) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		AWSTemplateFormatVersion string              `json:"AWSTemplateFormatVersion,omitempty"`
		Resources                map[string]Resource `json:"Resources,omitempty"`
	}{
		AWSTemplateFormatVersion: "2010-09-09",
		Resources:                memz.ShallowCopyMap(t.resources),
	})
}
