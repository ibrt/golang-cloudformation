// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudfront

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig__Type is the CloudFormation type for AWS::CloudFront::ContinuousDeploymentPolicy.SingleHeaderConfig.
	AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig__Type = "AWS::CloudFront::ContinuousDeploymentPolicy.SingleHeaderConfig"
)

var (
	// AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig__PropertiesMap reports all the CloudFormation properties for AWS::CloudFront::ContinuousDeploymentPolicy.SingleHeaderConfig.
	AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig__PropertiesMap = struct {
		Header string
		Value  string
	}{
		Header: "Header",
		Value:  "Value",
	}

	// AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFront::ContinuousDeploymentPolicy.SingleHeaderConfig.
	AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig__PropertiesSlice = []string{
		AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig__PropertiesMap.Header,
		AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig__PropertiesMap.Value,
	}
)

// AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig is a binding for AWS::CloudFront::ContinuousDeploymentPolicy.SingleHeaderConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.html
type AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig struct {
	// Header is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleheaderconfig-header
	Header cfz.Expression[string] `json:"Header,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.html#cfn-cloudfront-continuousdeploymentpolicy-singleheaderconfig-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig initializes a new AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig.
func New__AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig() AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig {
	return AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig) GetType() string {
	return AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig__Type
}

// Set__Header updates property "Header".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig) Set__Header(v cfz.Expression[string]) AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig {
	t.Header = v
	return t
}

// SetV__Header updates property "Header".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig) SetV__Header(v string) AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig {
	t.Header = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig) Set__Value(v cfz.Expression[string]) AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig) SetV__Value(v string) AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderConfig {
	t.Value = cfz.V(v)
	return t
}
