// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_amplifyuibuilder

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AmplifyUIBuilder_Component_ActionParameters__Type is the CloudFormation type for AWS::AmplifyUIBuilder::Component.ActionParameters.
	AWS_AmplifyUIBuilder_Component_ActionParameters__Type = "AWS::AmplifyUIBuilder::Component.ActionParameters"
)

var (
	// AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap reports all the CloudFormation properties for AWS::AmplifyUIBuilder::Component.ActionParameters.
	AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap = struct {
		Anchor string
		Fields string
		Global string
		Id     string
		Model  string
		State  string
		Target string
		Type   string
		Url    string
	}{
		Anchor: "Anchor",
		Fields: "Fields",
		Global: "Global",
		Id:     "Id",
		Model:  "Model",
		State:  "State",
		Target: "Target",
		Type:   "Type",
		Url:    "Url",
	}

	// AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesSlice reports all the CloudFormation properties for AWS::AmplifyUIBuilder::Component.ActionParameters.
	AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesSlice = []string{
		AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap.Anchor,
		AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap.Fields,
		AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap.Global,
		AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap.Id,
		AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap.Model,
		AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap.State,
		AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap.Target,
		AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap.Type,
		AWS_AmplifyUIBuilder_Component_ActionParameters__PropertiesMap.Url,
	}
)

// AWS_AmplifyUIBuilder_Component_ActionParameters is a binding for AWS::AmplifyUIBuilder::Component.ActionParameters.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html
type AWS_AmplifyUIBuilder_Component_ActionParameters struct {
	// Anchor is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-anchor
	Anchor cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty] `json:"Anchor,omitempty"`

	// Fields is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-fields
	Fields cfz.ExpressionMap[AWS_AmplifyUIBuilder_Component_ComponentProperty] `json:"Fields,omitempty"`

	// Global is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-global
	Global cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty] `json:"Global,omitempty"`

	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-id
	Id cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty] `json:"Id,omitempty"`

	// Model is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-model
	Model cfz.Expression[string] `json:"Model,omitempty"`

	// State is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-state
	State cfz.Expression[AWS_AmplifyUIBuilder_Component_MutationActionSetStateParameter] `json:"State,omitempty"`

	// Target is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-target
	Target cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty] `json:"Target,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-type
	Type cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty] `json:"Type,omitempty"`

	// Url is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-actionparameters.html#cfn-amplifyuibuilder-component-actionparameters-url
	Url cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty] `json:"Url,omitempty"`
}

// New__AWS_AmplifyUIBuilder_Component_ActionParameters initializes a new AWS_AmplifyUIBuilder_Component_ActionParameters.
func New__AWS_AmplifyUIBuilder_Component_ActionParameters() AWS_AmplifyUIBuilder_Component_ActionParameters {
	return AWS_AmplifyUIBuilder_Component_ActionParameters{}
}

// GetType returns the CloudFormation type.
func (AWS_AmplifyUIBuilder_Component_ActionParameters) GetType() string {
	return AWS_AmplifyUIBuilder_Component_ActionParameters__Type
}

// Set__Anchor updates property "Anchor".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) Set__Anchor(v cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Anchor = v
	return t
}

// SetV__Anchor updates property "Anchor".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetV__Anchor(v AWS_AmplifyUIBuilder_Component_ComponentProperty) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Anchor = cfz.V(v)
	return t
}

// Set__Fields updates property "Fields".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) Set__Fields(v cfz.ExpressionMap[AWS_AmplifyUIBuilder_Component_ComponentProperty]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Fields = v
	return t
}

// SetM__Fields updates property "Fields".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetM__Fields(v ...map[string]cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Fields = cfz.M(v...)
	return t
}

// SetMV__Fields updates property "Fields".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetMV__Fields(v ...map[string]AWS_AmplifyUIBuilder_Component_ComponentProperty) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Fields = cfz.MV(v...)
	return t
}

// Set__Global updates property "Global".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) Set__Global(v cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Global = v
	return t
}

// SetV__Global updates property "Global".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetV__Global(v AWS_AmplifyUIBuilder_Component_ComponentProperty) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Global = cfz.V(v)
	return t
}

// Set__Id updates property "Id".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) Set__Id(v cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetV__Id(v AWS_AmplifyUIBuilder_Component_ComponentProperty) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Id = cfz.V(v)
	return t
}

// Set__Model updates property "Model".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) Set__Model(v cfz.Expression[string]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Model = v
	return t
}

// SetV__Model updates property "Model".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetV__Model(v string) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Model = cfz.V(v)
	return t
}

// Set__State updates property "State".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) Set__State(v cfz.Expression[AWS_AmplifyUIBuilder_Component_MutationActionSetStateParameter]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.State = v
	return t
}

// SetV__State updates property "State".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetV__State(v AWS_AmplifyUIBuilder_Component_MutationActionSetStateParameter) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.State = cfz.V(v)
	return t
}

// Set__Target updates property "Target".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) Set__Target(v cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Target = v
	return t
}

// SetV__Target updates property "Target".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetV__Target(v AWS_AmplifyUIBuilder_Component_ComponentProperty) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Target = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) Set__Type(v cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetV__Type(v AWS_AmplifyUIBuilder_Component_ComponentProperty) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Type = cfz.V(v)
	return t
}

// Set__Url updates property "Url".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) Set__Url(v cfz.Expression[AWS_AmplifyUIBuilder_Component_ComponentProperty]) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Url = v
	return t
}

// SetV__Url updates property "Url".
func (t AWS_AmplifyUIBuilder_Component_ActionParameters) SetV__Url(v AWS_AmplifyUIBuilder_Component_ComponentProperty) AWS_AmplifyUIBuilder_Component_ActionParameters {
	t.Url = cfz.V(v)
	return t
}
