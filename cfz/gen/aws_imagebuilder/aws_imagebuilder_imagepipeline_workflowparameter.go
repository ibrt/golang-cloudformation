// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_imagebuilder

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ImageBuilder_ImagePipeline_WorkflowParameter__Type is the CloudFormation type for AWS::ImageBuilder::ImagePipeline.WorkflowParameter.
	AWS_ImageBuilder_ImagePipeline_WorkflowParameter__Type = "AWS::ImageBuilder::ImagePipeline.WorkflowParameter"
)

var (
	// AWS_ImageBuilder_ImagePipeline_WorkflowParameter__PropertiesMap reports all the CloudFormation properties for AWS::ImageBuilder::ImagePipeline.WorkflowParameter.
	AWS_ImageBuilder_ImagePipeline_WorkflowParameter__PropertiesMap = struct {
		Name  string
		Value string
	}{
		Name:  "Name",
		Value: "Value",
	}

	// AWS_ImageBuilder_ImagePipeline_WorkflowParameter__PropertiesSlice reports all the CloudFormation properties for AWS::ImageBuilder::ImagePipeline.WorkflowParameter.
	AWS_ImageBuilder_ImagePipeline_WorkflowParameter__PropertiesSlice = []string{
		AWS_ImageBuilder_ImagePipeline_WorkflowParameter__PropertiesMap.Name,
		AWS_ImageBuilder_ImagePipeline_WorkflowParameter__PropertiesMap.Value,
	}
)

// AWS_ImageBuilder_ImagePipeline_WorkflowParameter is a binding for AWS::ImageBuilder::ImagePipeline.WorkflowParameter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowparameter.html
type AWS_ImageBuilder_ImagePipeline_WorkflowParameter struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowparameter.html#cfn-imagebuilder-imagepipeline-workflowparameter-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-workflowparameter.html#cfn-imagebuilder-imagepipeline-workflowparameter-value
	Value cfz.ExpressionSlice[string] `json:"Value,omitempty"`
}

// New__AWS_ImageBuilder_ImagePipeline_WorkflowParameter initializes a new AWS_ImageBuilder_ImagePipeline_WorkflowParameter.
func New__AWS_ImageBuilder_ImagePipeline_WorkflowParameter() AWS_ImageBuilder_ImagePipeline_WorkflowParameter {
	return AWS_ImageBuilder_ImagePipeline_WorkflowParameter{}
}

// GetType returns the CloudFormation type.
func (AWS_ImageBuilder_ImagePipeline_WorkflowParameter) GetType() string {
	return AWS_ImageBuilder_ImagePipeline_WorkflowParameter__Type
}

// Set__Name updates property "Name".
func (t AWS_ImageBuilder_ImagePipeline_WorkflowParameter) Set__Name(v cfz.Expression[string]) AWS_ImageBuilder_ImagePipeline_WorkflowParameter {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_ImageBuilder_ImagePipeline_WorkflowParameter) SetV__Name(v string) AWS_ImageBuilder_ImagePipeline_WorkflowParameter {
	t.Name = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_ImageBuilder_ImagePipeline_WorkflowParameter) Set__Value(v cfz.ExpressionSlice[string]) AWS_ImageBuilder_ImagePipeline_WorkflowParameter {
	t.Value = v
	return t
}

// SetS__Value updates property "Value".
func (t AWS_ImageBuilder_ImagePipeline_WorkflowParameter) SetS__Value(v ...cfz.Expression[string]) AWS_ImageBuilder_ImagePipeline_WorkflowParameter {
	t.Value = cfz.S(v...)
	return t
}

// SetSV__Value updates property "Value".
func (t AWS_ImageBuilder_ImagePipeline_WorkflowParameter) SetSV__Value(v ...string) AWS_ImageBuilder_ImagePipeline_WorkflowParameter {
	t.Value = cfz.SV(v...)
	return t
}
