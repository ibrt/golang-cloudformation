// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_transfer

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Transfer_Workflow_CustomStepDetails__Type is the CloudFormation type for AWS::Transfer::Workflow.CustomStepDetails.
	AWS_Transfer_Workflow_CustomStepDetails__Type = "AWS::Transfer::Workflow.CustomStepDetails"
)

var (
	// AWS_Transfer_Workflow_CustomStepDetails__PropertiesMap reports all the CloudFormation properties for AWS::Transfer::Workflow.CustomStepDetails.
	AWS_Transfer_Workflow_CustomStepDetails__PropertiesMap = struct {
		Name               string
		SourceFileLocation string
		Target             string
		TimeoutSeconds     string
	}{
		Name:               "Name",
		SourceFileLocation: "SourceFileLocation",
		Target:             "Target",
		TimeoutSeconds:     "TimeoutSeconds",
	}

	// AWS_Transfer_Workflow_CustomStepDetails__PropertiesSlice reports all the CloudFormation properties for AWS::Transfer::Workflow.CustomStepDetails.
	AWS_Transfer_Workflow_CustomStepDetails__PropertiesSlice = []string{
		AWS_Transfer_Workflow_CustomStepDetails__PropertiesMap.Name,
		AWS_Transfer_Workflow_CustomStepDetails__PropertiesMap.SourceFileLocation,
		AWS_Transfer_Workflow_CustomStepDetails__PropertiesMap.Target,
		AWS_Transfer_Workflow_CustomStepDetails__PropertiesMap.TimeoutSeconds,
	}
)

// AWS_Transfer_Workflow_CustomStepDetails is a binding for AWS::Transfer::Workflow.CustomStepDetails.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html
type AWS_Transfer_Workflow_CustomStepDetails struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html#cfn-transfer-workflow-customstepdetails-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// SourceFileLocation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html#cfn-transfer-workflow-customstepdetails-sourcefilelocation
	SourceFileLocation cfz.Expression[string] `json:"SourceFileLocation,omitempty"`

	// Target is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html#cfn-transfer-workflow-customstepdetails-target
	Target cfz.Expression[string] `json:"Target,omitempty"`

	// TimeoutSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html#cfn-transfer-workflow-customstepdetails-timeoutseconds
	TimeoutSeconds cfz.Expression[int32] `json:"TimeoutSeconds,omitempty"`
}

// New__AWS_Transfer_Workflow_CustomStepDetails initializes a new AWS_Transfer_Workflow_CustomStepDetails.
func New__AWS_Transfer_Workflow_CustomStepDetails() AWS_Transfer_Workflow_CustomStepDetails {
	return AWS_Transfer_Workflow_CustomStepDetails{}
}

// GetType returns the CloudFormation type.
func (AWS_Transfer_Workflow_CustomStepDetails) GetType() string {
	return AWS_Transfer_Workflow_CustomStepDetails__Type
}

// Set__Name updates property "Name".
func (t AWS_Transfer_Workflow_CustomStepDetails) Set__Name(v cfz.Expression[string]) AWS_Transfer_Workflow_CustomStepDetails {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_Transfer_Workflow_CustomStepDetails) SetV__Name(v string) AWS_Transfer_Workflow_CustomStepDetails {
	t.Name = cfz.V(v)
	return t
}

// Set__SourceFileLocation updates property "SourceFileLocation".
func (t AWS_Transfer_Workflow_CustomStepDetails) Set__SourceFileLocation(v cfz.Expression[string]) AWS_Transfer_Workflow_CustomStepDetails {
	t.SourceFileLocation = v
	return t
}

// SetV__SourceFileLocation updates property "SourceFileLocation".
func (t AWS_Transfer_Workflow_CustomStepDetails) SetV__SourceFileLocation(v string) AWS_Transfer_Workflow_CustomStepDetails {
	t.SourceFileLocation = cfz.V(v)
	return t
}

// Set__Target updates property "Target".
func (t AWS_Transfer_Workflow_CustomStepDetails) Set__Target(v cfz.Expression[string]) AWS_Transfer_Workflow_CustomStepDetails {
	t.Target = v
	return t
}

// SetV__Target updates property "Target".
func (t AWS_Transfer_Workflow_CustomStepDetails) SetV__Target(v string) AWS_Transfer_Workflow_CustomStepDetails {
	t.Target = cfz.V(v)
	return t
}

// Set__TimeoutSeconds updates property "TimeoutSeconds".
func (t AWS_Transfer_Workflow_CustomStepDetails) Set__TimeoutSeconds(v cfz.Expression[int32]) AWS_Transfer_Workflow_CustomStepDetails {
	t.TimeoutSeconds = v
	return t
}

// SetV__TimeoutSeconds updates property "TimeoutSeconds".
func (t AWS_Transfer_Workflow_CustomStepDetails) SetV__TimeoutSeconds(v int32) AWS_Transfer_Workflow_CustomStepDetails {
	t.TimeoutSeconds = cfz.V(v)
	return t
}
