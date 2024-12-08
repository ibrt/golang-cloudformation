// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apptest

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppTest_TestCase_Input__Type is the CloudFormation type for AWS::AppTest::TestCase.Input.
	AWS_AppTest_TestCase_Input__Type = "AWS::AppTest::TestCase.Input"
)

var (
	// AWS_AppTest_TestCase_Input__PropertiesMap reports all the CloudFormation properties for AWS::AppTest::TestCase.Input.
	AWS_AppTest_TestCase_Input__PropertiesMap = struct {
		File string
	}{
		File: "File",
	}

	// AWS_AppTest_TestCase_Input__PropertiesSlice reports all the CloudFormation properties for AWS::AppTest::TestCase.Input.
	AWS_AppTest_TestCase_Input__PropertiesSlice = []string{
		AWS_AppTest_TestCase_Input__PropertiesMap.File,
	}
)

// AWS_AppTest_TestCase_Input is a binding for AWS::AppTest::TestCase.Input.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-input.html
type AWS_AppTest_TestCase_Input struct {
	// File is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-input.html#cfn-apptest-testcase-input-file
	File cfz.Expression[AWS_AppTest_TestCase_InputFile] `json:"File,omitempty"`
}

// New__AWS_AppTest_TestCase_Input initializes a new AWS_AppTest_TestCase_Input.
func New__AWS_AppTest_TestCase_Input() AWS_AppTest_TestCase_Input {
	return AWS_AppTest_TestCase_Input{}
}

// GetType returns the CloudFormation type.
func (AWS_AppTest_TestCase_Input) GetType() string {
	return AWS_AppTest_TestCase_Input__Type
}

// Set__File updates property "File".
func (t AWS_AppTest_TestCase_Input) Set__File(v cfz.Expression[AWS_AppTest_TestCase_InputFile]) AWS_AppTest_TestCase_Input {
	t.File = v
	return t
}

// SetV__File updates property "File".
func (t AWS_AppTest_TestCase_Input) SetV__File(v AWS_AppTest_TestCase_InputFile) AWS_AppTest_TestCase_Input {
	t.File = cfz.V(v)
	return t
}
