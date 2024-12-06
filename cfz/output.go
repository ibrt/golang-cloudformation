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
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html
type Output interface {
	OutputPartialLogicalName
	json.Marshaler
}

// OutputImpl is a provided implementation of the Output interface, which can be directly used as a struct pointer.
type OutputImpl struct {
	LogicalName string             `json:"-"`
	Description string             `json:"-"`
	Value       Expression[string] `json:"-"`
	ExportName  Expression[string] `json:"-"`
}

type outputImpl struct {
	Description string             `json:"Description,omitempty"`
	Value       Expression[string] `json:"Value,omitempty"`
	Export      *outputImplExport  `json:"Export,omitempty"`
}

type outputImplExport struct {
	Name Expression[string] `json:"Name,omitempty"`
}

// GetOutputLogicalName implements the Output interface.
func (o *OutputImpl) GetOutputLogicalName() string {
	return o.LogicalName
}

// MarshalJSON implements the Output interface.
func (o *OutputImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(outputImpl{
		Description: o.Description,
		Value:       o.Value,
		Export:      memz.Ternary(o.ExportName != nil, &outputImplExport{Name: o.ExportName}, nil),
	})
}

// OutputBuilder is a provided utility for creating outputs with a fluent interface.
type OutputBuilder struct {
	i *OutputImpl
}

// NewOutputBuilder initializes a new output builder.
func NewOutputBuilder(logicalName string) *OutputBuilder {
	return &OutputBuilder{
		i: &OutputImpl{
			LogicalName: logicalName,
		},
	}
}

// SetDescription sets the description.
func (b *OutputBuilder) SetDescription(description string) *OutputBuilder {
	b.i.Description = description
	return b
}

// SetValue sets the value from an Expression[string]
func (b *OutputBuilder) SetValue(value Expression[string]) *OutputBuilder {
	b.i.Value = value
	return b
}

// SetValueString sets the value from a constant string value.
func (b *OutputBuilder) SetValueString(value string) *OutputBuilder {
	b.i.Value = V(value)
	return b
}

// SetExportName sets the export name from an Expression[string]
func (b *OutputBuilder) SetExportName(exportName Expression[string]) *OutputBuilder {
	b.i.ExportName = exportName
	return b
}

// SetExportNameString sets the export name from a constant string value.
func (b *OutputBuilder) SetExportNameString(exportName Expression[string]) *OutputBuilder {
	b.i.ExportName = exportName
	return b
}

// SetConventionalExportName generates and sets an export name like "${AWS::StackName}-${OutputLogicalName}".
func (b *OutputBuilder) SetConventionalExportName() *OutputBuilder {
	b.i.ExportName = Sub(fmt.Sprintf("${AWS::StackName}-%v", b.i.LogicalName))
	return b
}

// Build the output.
func (b *OutputBuilder) Build() Output {
	return &OutputImpl{
		LogicalName: b.i.LogicalName,
		Description: b.i.Description,
		Value:       b.i.Value,
		ExportName:  b.i.ExportName,
	}
}
