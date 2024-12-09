package cfz

import (
	"encoding/json"
	"fmt"

	"github.com/ibrt/golang-utils/memz"
)

// OutputPartialLogicalName describes a subset of a CloudFormation output.
type OutputPartialLogicalName interface {
	GetOutputLogicalName() string
}

// Output describes a CloudFormation output.
type Output interface {
	OutputPartialLogicalName
	json.Marshaler
}

var (
	_ OutputPartialLogicalName = (*OutputImpl[any])(nil)
	_ Output                   = (*OutputImpl[any])(nil)
)

type outputImpl[T any] struct {
	Description string            `json:"Description,omitempty"`
	Value       Expression[T]     `json:"Value,omitempty"`
	Export      *outputImplExport `json:"Export,omitempty"`
}

type outputImplExport struct {
	Name Expression[string] `json:"Name,omitempty"`
}

// OutputImpl is a provided implementation for an output.
type OutputImpl[T any] struct {
	LogicalName string             `json:"-"`
	Description string             `json:"-"`
	Value       Expression[T]      `json:"-"`
	ExportName  Expression[string] `json:"-"`
}

// NewOutput initializes a new output.
func NewOutput[T any](logicalName string, value Expression[T]) *OutputImpl[T] {
	return &OutputImpl[T]{
		LogicalName: logicalName,
		Value:       value,
	}
}

// GetOutputLogicalName returns the logical name of the output.
func (o *OutputImpl[T]) GetOutputLogicalName() string {
	return o.LogicalName
}

// SetLogicalName sets the logical name.
func (o *OutputImpl[T]) SetLogicalName(logicalName string) *OutputImpl[T] {
	o.Description = logicalName
	return o
}

// SetDescription sets the description.
func (o *OutputImpl[T]) SetDescription(description string) *OutputImpl[T] {
	o.Description = description
	return o
}

// SetValue sets the value.
func (o *OutputImpl[T]) SetValue(value Expression[T]) *OutputImpl[T] {
	o.Value = value
	return o
}

// SetExportName sets the export name.
func (o *OutputImpl[T]) SetExportName(exportName Expression[string]) *OutputImpl[T] {
	o.ExportName = exportName
	return o
}

// SetVExportName sets the export name.
func (o *OutputImpl[T]) SetVExportName(exportName string) *OutputImpl[T] {
	o.ExportName = V(exportName)
	return o
}

// SetConventionalExportName sets the export name to "${AWS::StackName}-<current value of LogicalName>".
func (o *OutputImpl[T]) SetConventionalExportName() *OutputImpl[T] {
	o.ExportName = Sub(fmt.Sprintf("${AWS::StackName}-%v", o.LogicalName))
	return o
}

// MarshalJSON marshals the output.
func (o *OutputImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(outputImpl[T]{
		Description: o.Description,
		Value:       o.Value,
		Export:      memz.Ternary(o.ExportName != nil, &outputImplExport{Name: o.ExportName}, nil),
	})
}
