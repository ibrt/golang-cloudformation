// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_entityresolution

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EntityResolution_MatchingWorkflow_OutputAttribute__Type is the CloudFormation type for AWS::EntityResolution::MatchingWorkflow.OutputAttribute.
	AWS_EntityResolution_MatchingWorkflow_OutputAttribute__Type = "AWS::EntityResolution::MatchingWorkflow.OutputAttribute"
)

var (
	// AWS_EntityResolution_MatchingWorkflow_OutputAttribute__PropertiesMap reports all the CloudFormation properties for AWS::EntityResolution::MatchingWorkflow.OutputAttribute.
	AWS_EntityResolution_MatchingWorkflow_OutputAttribute__PropertiesMap = struct {
		Hashed string
		Name   string
	}{
		Hashed: "Hashed",
		Name:   "Name",
	}

	// AWS_EntityResolution_MatchingWorkflow_OutputAttribute__PropertiesSlice reports all the CloudFormation properties for AWS::EntityResolution::MatchingWorkflow.OutputAttribute.
	AWS_EntityResolution_MatchingWorkflow_OutputAttribute__PropertiesSlice = []string{
		AWS_EntityResolution_MatchingWorkflow_OutputAttribute__PropertiesMap.Hashed,
		AWS_EntityResolution_MatchingWorkflow_OutputAttribute__PropertiesMap.Name,
	}
)

// AWS_EntityResolution_MatchingWorkflow_OutputAttribute is a binding for AWS::EntityResolution::MatchingWorkflow.OutputAttribute.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputattribute.html
type AWS_EntityResolution_MatchingWorkflow_OutputAttribute struct {
	// Hashed is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputattribute.html#cfn-entityresolution-matchingworkflow-outputattribute-hashed
	Hashed cfz.Expression[bool] `json:"Hashed,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputattribute.html#cfn-entityresolution-matchingworkflow-outputattribute-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_EntityResolution_MatchingWorkflow_OutputAttribute initializes a new AWS_EntityResolution_MatchingWorkflow_OutputAttribute.
func New__AWS_EntityResolution_MatchingWorkflow_OutputAttribute() AWS_EntityResolution_MatchingWorkflow_OutputAttribute {
	return AWS_EntityResolution_MatchingWorkflow_OutputAttribute{}
}

// GetType returns the CloudFormation type.
func (AWS_EntityResolution_MatchingWorkflow_OutputAttribute) GetType() string {
	return AWS_EntityResolution_MatchingWorkflow_OutputAttribute__Type
}

// Set__Hashed updates property "Hashed".
func (t AWS_EntityResolution_MatchingWorkflow_OutputAttribute) Set__Hashed(v cfz.Expression[bool]) AWS_EntityResolution_MatchingWorkflow_OutputAttribute {
	t.Hashed = v
	return t
}

// SetV__Hashed updates property "Hashed".
func (t AWS_EntityResolution_MatchingWorkflow_OutputAttribute) SetV__Hashed(v bool) AWS_EntityResolution_MatchingWorkflow_OutputAttribute {
	t.Hashed = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_EntityResolution_MatchingWorkflow_OutputAttribute) Set__Name(v cfz.Expression[string]) AWS_EntityResolution_MatchingWorkflow_OutputAttribute {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_EntityResolution_MatchingWorkflow_OutputAttribute) SetV__Name(v string) AWS_EntityResolution_MatchingWorkflow_OutputAttribute {
	t.Name = cfz.V(v)
	return t
}
