package cfz

import (
	"encoding/json"
	"fmt"

	"github.com/ibrt/golang-utils/memz"
)

var (
	_ json.Marshaler = (*Output)(nil)
)

// Output describes an output.
type Output struct {
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

// NewOutput initializes a new output.
func NewOutput(logicalName string) *Output {
	return &Output{
		LogicalName: logicalName,
	}
}

// GetOutputLogicalName returns the logical name of the output.
func (o *Output) GetOutputLogicalName() string {
	return o.LogicalName
}

// SetDescription sets the description.
func (o *Output) SetDescription(description string) *Output {
	o.Description = description
	return o
}

// SetValue sets the value.
func (o *Output) SetValue(value Expression[string]) *Output {
	o.Value = value
	return o
}

// SetVValue sets the value.
func (o *Output) SetVValue(value string) *Output {
	o.Value = V(value)
	return o
}

// SetExportName sets the export name.
func (o *Output) SetExportName(exportName Expression[string]) *Output {
	o.ExportName = exportName
	return o
}

// SetVExportName sets the export name.
func (o *Output) SetVExportName(exportName string) *Output {
	o.ExportName = V(exportName)
	return o
}

// SetConventionalExportName sets the export name to "${AWS::StackName}-<current value of LogicalName>".
func (o *Output) SetConventionalExportName() *Output {
	o.ExportName = Sub(fmt.Sprintf("${AWS::StackName}-%v", o.LogicalName))
	return o
}

// MarshalJSON marshals the output.
func (o *Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(outputImpl{
		Description: o.Description,
		Value:       o.Value,
		Export:      memz.Ternary(o.ExportName != nil, &outputImplExport{Name: o.ExportName}, nil),
	})
}
