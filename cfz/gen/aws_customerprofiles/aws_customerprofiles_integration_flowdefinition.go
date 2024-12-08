// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_customerprofiles

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CustomerProfiles_Integration_FlowDefinition__Type is the CloudFormation type for AWS::CustomerProfiles::Integration.FlowDefinition.
	AWS_CustomerProfiles_Integration_FlowDefinition__Type = "AWS::CustomerProfiles::Integration.FlowDefinition"
)

var (
	// AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesMap reports all the CloudFormation properties for AWS::CustomerProfiles::Integration.FlowDefinition.
	AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesMap = struct {
		Description      string
		FlowName         string
		KmsArn           string
		SourceFlowConfig string
		Tasks            string
		TriggerConfig    string
	}{
		Description:      "Description",
		FlowName:         "FlowName",
		KmsArn:           "KmsArn",
		SourceFlowConfig: "SourceFlowConfig",
		Tasks:            "Tasks",
		TriggerConfig:    "TriggerConfig",
	}

	// AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesSlice reports all the CloudFormation properties for AWS::CustomerProfiles::Integration.FlowDefinition.
	AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesSlice = []string{
		AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesMap.Description,
		AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesMap.FlowName,
		AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesMap.KmsArn,
		AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesMap.SourceFlowConfig,
		AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesMap.Tasks,
		AWS_CustomerProfiles_Integration_FlowDefinition__PropertiesMap.TriggerConfig,
	}
)

// AWS_CustomerProfiles_Integration_FlowDefinition is a binding for AWS::CustomerProfiles::Integration.FlowDefinition.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html
type AWS_CustomerProfiles_Integration_FlowDefinition struct {
	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// FlowName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-flowname
	FlowName cfz.Expression[string] `json:"FlowName,omitempty"`

	// KmsArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-kmsarn
	KmsArn cfz.Expression[string] `json:"KmsArn,omitempty"`

	// SourceFlowConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-sourceflowconfig
	SourceFlowConfig cfz.Expression[AWS_CustomerProfiles_Integration_SourceFlowConfig] `json:"SourceFlowConfig,omitempty"`

	// Tasks is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-tasks
	Tasks cfz.ExpressionSlice[AWS_CustomerProfiles_Integration_Task] `json:"Tasks,omitempty"`

	// TriggerConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-flowdefinition.html#cfn-customerprofiles-integration-flowdefinition-triggerconfig
	TriggerConfig cfz.Expression[AWS_CustomerProfiles_Integration_TriggerConfig] `json:"TriggerConfig,omitempty"`
}

// New__AWS_CustomerProfiles_Integration_FlowDefinition initializes a new AWS_CustomerProfiles_Integration_FlowDefinition.
func New__AWS_CustomerProfiles_Integration_FlowDefinition() AWS_CustomerProfiles_Integration_FlowDefinition {
	return AWS_CustomerProfiles_Integration_FlowDefinition{}
}

// GetType returns the CloudFormation type.
func (AWS_CustomerProfiles_Integration_FlowDefinition) GetType() string {
	return AWS_CustomerProfiles_Integration_FlowDefinition__Type
}

// Set__Description updates property "Description".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) Set__Description(v cfz.Expression[string]) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) SetV__Description(v string) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.Description = cfz.V(v)
	return t
}

// Set__FlowName updates property "FlowName".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) Set__FlowName(v cfz.Expression[string]) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.FlowName = v
	return t
}

// SetV__FlowName updates property "FlowName".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) SetV__FlowName(v string) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.FlowName = cfz.V(v)
	return t
}

// Set__KmsArn updates property "KmsArn".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) Set__KmsArn(v cfz.Expression[string]) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.KmsArn = v
	return t
}

// SetV__KmsArn updates property "KmsArn".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) SetV__KmsArn(v string) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.KmsArn = cfz.V(v)
	return t
}

// Set__SourceFlowConfig updates property "SourceFlowConfig".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) Set__SourceFlowConfig(v cfz.Expression[AWS_CustomerProfiles_Integration_SourceFlowConfig]) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.SourceFlowConfig = v
	return t
}

// SetV__SourceFlowConfig updates property "SourceFlowConfig".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) SetV__SourceFlowConfig(v AWS_CustomerProfiles_Integration_SourceFlowConfig) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.SourceFlowConfig = cfz.V(v)
	return t
}

// Set__Tasks updates property "Tasks".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) Set__Tasks(v cfz.ExpressionSlice[AWS_CustomerProfiles_Integration_Task]) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.Tasks = v
	return t
}

// SetS__Tasks updates property "Tasks".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) SetS__Tasks(v ...cfz.Expression[AWS_CustomerProfiles_Integration_Task]) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.Tasks = cfz.S(v...)
	return t
}

// SetSV__Tasks updates property "Tasks".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) SetSV__Tasks(v ...AWS_CustomerProfiles_Integration_Task) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.Tasks = cfz.SV(v...)
	return t
}

// Set__TriggerConfig updates property "TriggerConfig".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) Set__TriggerConfig(v cfz.Expression[AWS_CustomerProfiles_Integration_TriggerConfig]) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.TriggerConfig = v
	return t
}

// SetV__TriggerConfig updates property "TriggerConfig".
func (t AWS_CustomerProfiles_Integration_FlowDefinition) SetV__TriggerConfig(v AWS_CustomerProfiles_Integration_TriggerConfig) AWS_CustomerProfiles_Integration_FlowDefinition {
	t.TriggerConfig = cfz.V(v)
	return t
}
