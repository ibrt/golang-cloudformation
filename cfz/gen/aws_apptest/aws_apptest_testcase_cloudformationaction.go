// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apptest

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppTest_TestCase_CloudFormationAction__Type is the CloudFormation type for AWS::AppTest::TestCase.CloudFormationAction.
	AWS_AppTest_TestCase_CloudFormationAction__Type = "AWS::AppTest::TestCase.CloudFormationAction"
)

var (
	// AWS_AppTest_TestCase_CloudFormationAction__PropertiesMap reports all the CloudFormation properties for AWS::AppTest::TestCase.CloudFormationAction.
	AWS_AppTest_TestCase_CloudFormationAction__PropertiesMap = struct {
		ActionType string
		Resource   string
	}{
		ActionType: "ActionType",
		Resource:   "Resource",
	}

	// AWS_AppTest_TestCase_CloudFormationAction__PropertiesSlice reports all the CloudFormation properties for AWS::AppTest::TestCase.CloudFormationAction.
	AWS_AppTest_TestCase_CloudFormationAction__PropertiesSlice = []string{
		AWS_AppTest_TestCase_CloudFormationAction__PropertiesMap.ActionType,
		AWS_AppTest_TestCase_CloudFormationAction__PropertiesMap.Resource,
	}
)

// AWS_AppTest_TestCase_CloudFormationAction is a binding for AWS::AppTest::TestCase.CloudFormationAction.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-cloudformationaction.html
type AWS_AppTest_TestCase_CloudFormationAction struct {
	// ActionType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-cloudformationaction.html#cfn-apptest-testcase-cloudformationaction-actiontype
	ActionType cfz.Expression[string] `json:"ActionType,omitempty"`

	// Resource is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-cloudformationaction.html#cfn-apptest-testcase-cloudformationaction-resource
	Resource cfz.Expression[string] `json:"Resource,omitempty"`
}

// New__AWS_AppTest_TestCase_CloudFormationAction initializes a new AWS_AppTest_TestCase_CloudFormationAction.
func New__AWS_AppTest_TestCase_CloudFormationAction() AWS_AppTest_TestCase_CloudFormationAction {
	return AWS_AppTest_TestCase_CloudFormationAction{}
}

// GetType returns the CloudFormation type.
func (AWS_AppTest_TestCase_CloudFormationAction) GetType() string {
	return AWS_AppTest_TestCase_CloudFormationAction__Type
}

// Set__ActionType updates property "ActionType".
func (t AWS_AppTest_TestCase_CloudFormationAction) Set__ActionType(v cfz.Expression[string]) AWS_AppTest_TestCase_CloudFormationAction {
	t.ActionType = v
	return t
}

// SetV__ActionType updates property "ActionType".
func (t AWS_AppTest_TestCase_CloudFormationAction) SetV__ActionType(v string) AWS_AppTest_TestCase_CloudFormationAction {
	t.ActionType = cfz.V(v)
	return t
}

// Set__Resource updates property "Resource".
func (t AWS_AppTest_TestCase_CloudFormationAction) Set__Resource(v cfz.Expression[string]) AWS_AppTest_TestCase_CloudFormationAction {
	t.Resource = v
	return t
}

// SetV__Resource updates property "Resource".
func (t AWS_AppTest_TestCase_CloudFormationAction) SetV__Resource(v string) AWS_AppTest_TestCase_CloudFormationAction {
	t.Resource = cfz.V(v)
	return t
}
