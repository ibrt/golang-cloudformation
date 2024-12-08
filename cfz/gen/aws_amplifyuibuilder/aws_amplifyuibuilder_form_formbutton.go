// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_amplifyuibuilder

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AmplifyUIBuilder_Form_FormButton__Type is the CloudFormation type for AWS::AmplifyUIBuilder::Form.FormButton.
	AWS_AmplifyUIBuilder_Form_FormButton__Type = "AWS::AmplifyUIBuilder::Form.FormButton"
)

var (
	// AWS_AmplifyUIBuilder_Form_FormButton__PropertiesMap reports all the CloudFormation properties for AWS::AmplifyUIBuilder::Form.FormButton.
	AWS_AmplifyUIBuilder_Form_FormButton__PropertiesMap = struct {
		Children string
		Excluded string
		Position string
	}{
		Children: "Children",
		Excluded: "Excluded",
		Position: "Position",
	}

	// AWS_AmplifyUIBuilder_Form_FormButton__PropertiesSlice reports all the CloudFormation properties for AWS::AmplifyUIBuilder::Form.FormButton.
	AWS_AmplifyUIBuilder_Form_FormButton__PropertiesSlice = []string{
		AWS_AmplifyUIBuilder_Form_FormButton__PropertiesMap.Children,
		AWS_AmplifyUIBuilder_Form_FormButton__PropertiesMap.Excluded,
		AWS_AmplifyUIBuilder_Form_FormButton__PropertiesMap.Position,
	}
)

// AWS_AmplifyUIBuilder_Form_FormButton is a binding for AWS::AmplifyUIBuilder::Form.FormButton.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formbutton.html
type AWS_AmplifyUIBuilder_Form_FormButton struct {
	// Children is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formbutton.html#cfn-amplifyuibuilder-form-formbutton-children
	Children cfz.Expression[string] `json:"Children,omitempty"`

	// Excluded is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formbutton.html#cfn-amplifyuibuilder-form-formbutton-excluded
	Excluded cfz.Expression[bool] `json:"Excluded,omitempty"`

	// Position is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formbutton.html#cfn-amplifyuibuilder-form-formbutton-position
	Position cfz.Expression[AWS_AmplifyUIBuilder_Form_FieldPosition] `json:"Position,omitempty"`
}

// New__AWS_AmplifyUIBuilder_Form_FormButton initializes a new AWS_AmplifyUIBuilder_Form_FormButton.
func New__AWS_AmplifyUIBuilder_Form_FormButton() AWS_AmplifyUIBuilder_Form_FormButton {
	return AWS_AmplifyUIBuilder_Form_FormButton{}
}

// GetType returns the CloudFormation type.
func (AWS_AmplifyUIBuilder_Form_FormButton) GetType() string {
	return AWS_AmplifyUIBuilder_Form_FormButton__Type
}

// Set__Children updates property "Children".
func (t AWS_AmplifyUIBuilder_Form_FormButton) Set__Children(v cfz.Expression[string]) AWS_AmplifyUIBuilder_Form_FormButton {
	t.Children = v
	return t
}

// SetV__Children updates property "Children".
func (t AWS_AmplifyUIBuilder_Form_FormButton) SetV__Children(v string) AWS_AmplifyUIBuilder_Form_FormButton {
	t.Children = cfz.V(v)
	return t
}

// Set__Excluded updates property "Excluded".
func (t AWS_AmplifyUIBuilder_Form_FormButton) Set__Excluded(v cfz.Expression[bool]) AWS_AmplifyUIBuilder_Form_FormButton {
	t.Excluded = v
	return t
}

// SetV__Excluded updates property "Excluded".
func (t AWS_AmplifyUIBuilder_Form_FormButton) SetV__Excluded(v bool) AWS_AmplifyUIBuilder_Form_FormButton {
	t.Excluded = cfz.V(v)
	return t
}

// Set__Position updates property "Position".
func (t AWS_AmplifyUIBuilder_Form_FormButton) Set__Position(v cfz.Expression[AWS_AmplifyUIBuilder_Form_FieldPosition]) AWS_AmplifyUIBuilder_Form_FormButton {
	t.Position = v
	return t
}

// SetV__Position updates property "Position".
func (t AWS_AmplifyUIBuilder_Form_FormButton) SetV__Position(v AWS_AmplifyUIBuilder_Form_FieldPosition) AWS_AmplifyUIBuilder_Form_FormButton {
	t.Position = cfz.V(v)
	return t
}
