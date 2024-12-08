// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_fis

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__Type is the CloudFormation type for AWS::FIS::ExperimentTemplate.ExperimentTemplateAction.
	AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__Type = "AWS::FIS::ExperimentTemplate.ExperimentTemplateAction"
)

var (
	// AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__PropertiesMap reports all the CloudFormation properties for AWS::FIS::ExperimentTemplate.ExperimentTemplateAction.
	AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__PropertiesMap = struct {
		ActionId    string
		Description string
		Parameters  string
		StartAfter  string
		Targets     string
	}{
		ActionId:    "ActionId",
		Description: "Description",
		Parameters:  "Parameters",
		StartAfter:  "StartAfter",
		Targets:     "Targets",
	}

	// AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__PropertiesSlice reports all the CloudFormation properties for AWS::FIS::ExperimentTemplate.ExperimentTemplateAction.
	AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__PropertiesSlice = []string{
		AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__PropertiesMap.ActionId,
		AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__PropertiesMap.Description,
		AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__PropertiesMap.Parameters,
		AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__PropertiesMap.StartAfter,
		AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__PropertiesMap.Targets,
	}
)

// AWS_FIS_ExperimentTemplate_ExperimentTemplateAction is a binding for AWS::FIS::ExperimentTemplate.ExperimentTemplateAction.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html
type AWS_FIS_ExperimentTemplate_ExperimentTemplateAction struct {
	// ActionId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-actionid
	ActionId cfz.Expression[string] `json:"ActionId,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Parameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-parameters
	Parameters cfz.ExpressionMap[string] `json:"Parameters,omitempty"`

	// StartAfter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-startafter
	StartAfter cfz.ExpressionSlice[string] `json:"StartAfter,omitempty"`

	// Targets is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateaction.html#cfn-fis-experimenttemplate-experimenttemplateaction-targets
	Targets cfz.ExpressionMap[string] `json:"Targets,omitempty"`
}

// New__AWS_FIS_ExperimentTemplate_ExperimentTemplateAction initializes a new AWS_FIS_ExperimentTemplate_ExperimentTemplateAction.
func New__AWS_FIS_ExperimentTemplate_ExperimentTemplateAction() AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	return AWS_FIS_ExperimentTemplate_ExperimentTemplateAction{}
}

// GetType returns the CloudFormation type.
func (AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) GetType() string {
	return AWS_FIS_ExperimentTemplate_ExperimentTemplateAction__Type
}

// Set__ActionId updates property "ActionId".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) Set__ActionId(v cfz.Expression[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.ActionId = v
	return t
}

// SetV__ActionId updates property "ActionId".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) SetV__ActionId(v string) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.ActionId = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) Set__Description(v cfz.Expression[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) SetV__Description(v string) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.Description = cfz.V(v)
	return t
}

// Set__Parameters updates property "Parameters".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) Set__Parameters(v cfz.ExpressionMap[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.Parameters = v
	return t
}

// SetM__Parameters updates property "Parameters".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) SetM__Parameters(v ...map[string]cfz.Expression[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.Parameters = cfz.M(v...)
	return t
}

// SetMV__Parameters updates property "Parameters".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) SetMV__Parameters(v ...map[string]string) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.Parameters = cfz.MV(v...)
	return t
}

// Set__StartAfter updates property "StartAfter".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) Set__StartAfter(v cfz.ExpressionSlice[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.StartAfter = v
	return t
}

// SetS__StartAfter updates property "StartAfter".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) SetS__StartAfter(v ...cfz.Expression[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.StartAfter = cfz.S(v...)
	return t
}

// SetSV__StartAfter updates property "StartAfter".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) SetSV__StartAfter(v ...string) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.StartAfter = cfz.SV(v...)
	return t
}

// Set__Targets updates property "Targets".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) Set__Targets(v cfz.ExpressionMap[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.Targets = v
	return t
}

// SetM__Targets updates property "Targets".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) SetM__Targets(v ...map[string]cfz.Expression[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.Targets = cfz.M(v...)
	return t
}

// SetMV__Targets updates property "Targets".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateAction) SetMV__Targets(v ...map[string]string) AWS_FIS_ExperimentTemplate_ExperimentTemplateAction {
	t.Targets = cfz.MV(v...)
	return t
}
