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

type rawOutputImpl struct {
	Description string               `json:"Description,omitempty"`
	Value       any                  `json:"Value,omitempty"`
	Export      *rawOutputImplExport `json:"Export,omitempty"`
}

type rawOutputImplExport struct {
	Name Expression[string] `json:"Name,omitempty"`
}

// OutputImpl is a provided implementation for an output whose value is resolved by an Expression[T].
type OutputImpl[T any] struct {
	LogicalName string             `json:"-"`
	Description string             `json:"-"`
	Value       Expression[T]      `json:"-"`
	ExportName  Expression[string] `json:"-"`
}

// NewOutput initializes a new output whose value is resolved by an Expression[T].
func NewOutput[T any](logicalName string, value Expression[T]) *OutputImpl[T] {
	return &OutputImpl[T]{
		LogicalName: logicalName,
		Value:       value,
	}
}

// GetOutputLogicalName implements the Output interface.
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

// MarshalJSON implements the Output interface.
func (o *OutputImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(rawOutputImpl{
		Description: o.Description,
		Value:       o.Value,
		Export:      memz.Ternary(o.ExportName != nil, &rawOutputImplExport{Name: o.ExportName}, nil),
	})
}

// OutputSliceImpl is a provided implementation for an output whose value is resolved by an ExpressionSlice[T].
type OutputSliceImpl[T any] struct {
	LogicalName string             `json:"-"`
	Description string             `json:"-"`
	Value       ExpressionSlice[T] `json:"-"`
	ExportName  Expression[string] `json:"-"`
}

// NewOutputSlice initializes a new output whose value is resolved by an ExpressionSlice[T].
func NewOutputSlice[T any](logicalName string, value ExpressionSlice[T]) *OutputSliceImpl[T] {
	return &OutputSliceImpl[T]{
		LogicalName: logicalName,
		Value:       value,
	}
}

// GetOutputLogicalName returns the logical name of the output.
func (o *OutputSliceImpl[T]) GetOutputLogicalName() string {
	return o.LogicalName
}

// SetLogicalName sets the logical name.
func (o *OutputSliceImpl[T]) SetLogicalName(logicalName string) *OutputSliceImpl[T] {
	o.Description = logicalName
	return o
}

// SetDescription sets the description.
func (o *OutputSliceImpl[T]) SetDescription(description string) *OutputSliceImpl[T] {
	o.Description = description
	return o
}

// SetValue sets the value.
func (o *OutputSliceImpl[T]) SetValue(value ExpressionSlice[T]) *OutputSliceImpl[T] {
	o.Value = value
	return o
}

// SetExportName sets the export name.
func (o *OutputSliceImpl[T]) SetExportName(exportName Expression[string]) *OutputSliceImpl[T] {
	o.ExportName = exportName
	return o
}

// SetVExportName sets the export name.
func (o *OutputSliceImpl[T]) SetVExportName(exportName string) *OutputSliceImpl[T] {
	o.ExportName = V(exportName)
	return o
}

// SetConventionalExportName sets the export name to "${AWS::StackName}-<current value of LogicalName>".
func (o *OutputSliceImpl[T]) SetConventionalExportName() *OutputSliceImpl[T] {
	o.ExportName = Sub(fmt.Sprintf("${AWS::StackName}-%v", o.LogicalName))
	return o
}

// MarshalJSON marshals the output.
func (o *OutputSliceImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(rawOutputImpl{
		Description: o.Description,
		Value:       o.Value,
		Export:      memz.Ternary(o.ExportName != nil, &rawOutputImplExport{Name: o.ExportName}, nil),
	})
}

// OutputMapImpl is a provided implementation for an output whose value is resolved by an ExpressionMap[T].
type OutputMapImpl[T any] struct {
	LogicalName string             `json:"-"`
	Description string             `json:"-"`
	Value       ExpressionMap[T]   `json:"-"`
	ExportName  Expression[string] `json:"-"`
}

// NewOutputMap initializes a new output whose value is resolved by an ExpressionMap[T].
func NewOutputMap[T any](logicalName string, value ExpressionMap[T]) *OutputMapImpl[T] {
	return &OutputMapImpl[T]{
		LogicalName: logicalName,
		Value:       value,
	}
}

// GetOutputLogicalName returns the logical name of the output.
func (o *OutputMapImpl[T]) GetOutputLogicalName() string {
	return o.LogicalName
}

// SetLogicalName sets the logical name.
func (o *OutputMapImpl[T]) SetLogicalName(logicalName string) *OutputMapImpl[T] {
	o.Description = logicalName
	return o
}

// SetDescription sets the description.
func (o *OutputMapImpl[T]) SetDescription(description string) *OutputMapImpl[T] {
	o.Description = description
	return o
}

// SetValue sets the value.
func (o *OutputMapImpl[T]) SetValue(value ExpressionMap[T]) *OutputMapImpl[T] {
	o.Value = value
	return o
}

// SetExportName sets the export name.
func (o *OutputMapImpl[T]) SetExportName(exportName Expression[string]) *OutputMapImpl[T] {
	o.ExportName = exportName
	return o
}

// SetVExportName sets the export name.
func (o *OutputMapImpl[T]) SetVExportName(exportName string) *OutputMapImpl[T] {
	o.ExportName = V(exportName)
	return o
}

// SetConventionalExportName sets the export name to "${AWS::StackName}-<current value of LogicalName>".
func (o *OutputMapImpl[T]) SetConventionalExportName() *OutputMapImpl[T] {
	o.ExportName = Sub(fmt.Sprintf("${AWS::StackName}-%v", o.LogicalName))
	return o
}

// MarshalJSON marshals the output.
func (o *OutputMapImpl[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(rawOutputImpl{
		Description: o.Description,
		Value:       o.Value,
		Export:      memz.Ternary(o.ExportName != nil, &rawOutputImplExport{Name: o.ExportName}, nil),
	})
}
