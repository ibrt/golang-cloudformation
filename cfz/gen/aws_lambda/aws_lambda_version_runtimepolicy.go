// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lambda

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lambda_Version_RuntimePolicy__Type is the CloudFormation type for AWS::Lambda::Version.RuntimePolicy.
	AWS_Lambda_Version_RuntimePolicy__Type = "AWS::Lambda::Version.RuntimePolicy"
)

var (
	// AWS_Lambda_Version_RuntimePolicy__PropertiesMap reports all the CloudFormation properties for AWS::Lambda::Version.RuntimePolicy.
	AWS_Lambda_Version_RuntimePolicy__PropertiesMap = struct {
		RuntimeVersionArn string
		UpdateRuntimeOn   string
	}{
		RuntimeVersionArn: "RuntimeVersionArn",
		UpdateRuntimeOn:   "UpdateRuntimeOn",
	}

	// AWS_Lambda_Version_RuntimePolicy__PropertiesSlice reports all the CloudFormation properties for AWS::Lambda::Version.RuntimePolicy.
	AWS_Lambda_Version_RuntimePolicy__PropertiesSlice = []string{
		AWS_Lambda_Version_RuntimePolicy__PropertiesMap.RuntimeVersionArn,
		AWS_Lambda_Version_RuntimePolicy__PropertiesMap.UpdateRuntimeOn,
	}
)

// AWS_Lambda_Version_RuntimePolicy is a binding for AWS::Lambda::Version.RuntimePolicy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-version-runtimepolicy.html
type AWS_Lambda_Version_RuntimePolicy struct {
	// RuntimeVersionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-version-runtimepolicy.html#cfn-lambda-version-runtimepolicy-runtimeversionarn
	RuntimeVersionArn cfz.Expression[string] `json:"RuntimeVersionArn,omitempty"`

	// UpdateRuntimeOn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-version-runtimepolicy.html#cfn-lambda-version-runtimepolicy-updateruntimeon
	UpdateRuntimeOn cfz.Expression[string] `json:"UpdateRuntimeOn,omitempty"`
}

// New__AWS_Lambda_Version_RuntimePolicy initializes a new AWS_Lambda_Version_RuntimePolicy.
func New__AWS_Lambda_Version_RuntimePolicy() AWS_Lambda_Version_RuntimePolicy {
	return AWS_Lambda_Version_RuntimePolicy{}
}

// GetType returns the CloudFormation type.
func (AWS_Lambda_Version_RuntimePolicy) GetType() string {
	return AWS_Lambda_Version_RuntimePolicy__Type
}

// Set__RuntimeVersionArn updates property "RuntimeVersionArn".
func (t AWS_Lambda_Version_RuntimePolicy) Set__RuntimeVersionArn(v cfz.Expression[string]) AWS_Lambda_Version_RuntimePolicy {
	t.RuntimeVersionArn = v
	return t
}

// SetV__RuntimeVersionArn updates property "RuntimeVersionArn".
func (t AWS_Lambda_Version_RuntimePolicy) SetV__RuntimeVersionArn(v string) AWS_Lambda_Version_RuntimePolicy {
	t.RuntimeVersionArn = cfz.V(v)
	return t
}

// Set__UpdateRuntimeOn updates property "UpdateRuntimeOn".
func (t AWS_Lambda_Version_RuntimePolicy) Set__UpdateRuntimeOn(v cfz.Expression[string]) AWS_Lambda_Version_RuntimePolicy {
	t.UpdateRuntimeOn = v
	return t
}

// SetV__UpdateRuntimeOn updates property "UpdateRuntimeOn".
func (t AWS_Lambda_Version_RuntimePolicy) SetV__UpdateRuntimeOn(v string) AWS_Lambda_Version_RuntimePolicy {
	t.UpdateRuntimeOn = cfz.V(v)
	return t
}
