// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_batch

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Batch_JobDefinition_ResourceRequirement__Type is the CloudFormation type for AWS::Batch::JobDefinition.ResourceRequirement.
	AWS_Batch_JobDefinition_ResourceRequirement__Type = "AWS::Batch::JobDefinition.ResourceRequirement"
)

var (
	// AWS_Batch_JobDefinition_ResourceRequirement__PropertiesMap reports all the CloudFormation properties for AWS::Batch::JobDefinition.ResourceRequirement.
	AWS_Batch_JobDefinition_ResourceRequirement__PropertiesMap = struct {
		Type  string
		Value string
	}{
		Type:  "Type",
		Value: "Value",
	}

	// AWS_Batch_JobDefinition_ResourceRequirement__PropertiesSlice reports all the CloudFormation properties for AWS::Batch::JobDefinition.ResourceRequirement.
	AWS_Batch_JobDefinition_ResourceRequirement__PropertiesSlice = []string{
		AWS_Batch_JobDefinition_ResourceRequirement__PropertiesMap.Type,
		AWS_Batch_JobDefinition_ResourceRequirement__PropertiesMap.Value,
	}
)

// AWS_Batch_JobDefinition_ResourceRequirement is a binding for AWS::Batch::JobDefinition.ResourceRequirement.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resourcerequirement.html
type AWS_Batch_JobDefinition_ResourceRequirement struct {
	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resourcerequirement.html#cfn-batch-jobdefinition-resourcerequirement-type
	Type cfz.Expression[string] `json:"Type,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resourcerequirement.html#cfn-batch-jobdefinition-resourcerequirement-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_Batch_JobDefinition_ResourceRequirement initializes a new AWS_Batch_JobDefinition_ResourceRequirement.
func New__AWS_Batch_JobDefinition_ResourceRequirement() AWS_Batch_JobDefinition_ResourceRequirement {
	return AWS_Batch_JobDefinition_ResourceRequirement{}
}

// GetType returns the CloudFormation type.
func (AWS_Batch_JobDefinition_ResourceRequirement) GetType() string {
	return AWS_Batch_JobDefinition_ResourceRequirement__Type
}

// Set__Type updates property "Type".
func (t AWS_Batch_JobDefinition_ResourceRequirement) Set__Type(v cfz.Expression[string]) AWS_Batch_JobDefinition_ResourceRequirement {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_Batch_JobDefinition_ResourceRequirement) SetV__Type(v string) AWS_Batch_JobDefinition_ResourceRequirement {
	t.Type = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_Batch_JobDefinition_ResourceRequirement) Set__Value(v cfz.Expression[string]) AWS_Batch_JobDefinition_ResourceRequirement {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_Batch_JobDefinition_ResourceRequirement) SetV__Value(v string) AWS_Batch_JobDefinition_ResourceRequirement {
	t.Value = cfz.V(v)
	return t
}
