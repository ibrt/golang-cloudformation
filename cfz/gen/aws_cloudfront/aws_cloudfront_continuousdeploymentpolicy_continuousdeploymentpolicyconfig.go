// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudfront

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__Type is the CloudFormation type for AWS::CloudFront::ContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfig.
	AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__Type = "AWS::CloudFront::ContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfig"
)

var (
	// AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesMap reports all the CloudFormation properties for AWS::CloudFront::ContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfig.
	AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesMap = struct {
		Enabled                     string
		SingleHeaderPolicyConfig    string
		SingleWeightPolicyConfig    string
		StagingDistributionDnsNames string
		TrafficConfig               string
		Type                        string
	}{
		Enabled:                     "Enabled",
		SingleHeaderPolicyConfig:    "SingleHeaderPolicyConfig",
		SingleWeightPolicyConfig:    "SingleWeightPolicyConfig",
		StagingDistributionDnsNames: "StagingDistributionDnsNames",
		TrafficConfig:               "TrafficConfig",
		Type:                        "Type",
	}

	// AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFront::ContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfig.
	AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesSlice = []string{
		AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesMap.Enabled,
		AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesMap.SingleHeaderPolicyConfig,
		AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesMap.SingleWeightPolicyConfig,
		AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesMap.StagingDistributionDnsNames,
		AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesMap.TrafficConfig,
		AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__PropertiesMap.Type,
	}
)

// AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig is a binding for AWS::CloudFront::ContinuousDeploymentPolicy.ContinuousDeploymentPolicyConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html
type AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig struct {
	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`

	// SingleHeaderPolicyConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-singleheaderpolicyconfig
	SingleHeaderPolicyConfig cfz.Expression[AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderPolicyConfig] `json:"SingleHeaderPolicyConfig,omitempty"`

	// SingleWeightPolicyConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-singleweightpolicyconfig
	SingleWeightPolicyConfig cfz.Expression[AWS_CloudFront_ContinuousDeploymentPolicy_SingleWeightPolicyConfig] `json:"SingleWeightPolicyConfig,omitempty"`

	// StagingDistributionDnsNames is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-stagingdistributiondnsnames
	StagingDistributionDnsNames cfz.ExpressionSlice[string] `json:"StagingDistributionDnsNames,omitempty"`

	// TrafficConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-trafficconfig
	TrafficConfig cfz.Expression[AWS_CloudFront_ContinuousDeploymentPolicy_TrafficConfig] `json:"TrafficConfig,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig initializes a new AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig.
func New__AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig() AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	return AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) GetType() string {
	return AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig__Type
}

// Set__Enabled updates property "Enabled".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) Set__Enabled(v cfz.Expression[bool]) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) SetV__Enabled(v bool) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.Enabled = cfz.V(v)
	return t
}

// Set__SingleHeaderPolicyConfig updates property "SingleHeaderPolicyConfig".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) Set__SingleHeaderPolicyConfig(v cfz.Expression[AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderPolicyConfig]) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.SingleHeaderPolicyConfig = v
	return t
}

// SetV__SingleHeaderPolicyConfig updates property "SingleHeaderPolicyConfig".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) SetV__SingleHeaderPolicyConfig(v AWS_CloudFront_ContinuousDeploymentPolicy_SingleHeaderPolicyConfig) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.SingleHeaderPolicyConfig = cfz.V(v)
	return t
}

// Set__SingleWeightPolicyConfig updates property "SingleWeightPolicyConfig".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) Set__SingleWeightPolicyConfig(v cfz.Expression[AWS_CloudFront_ContinuousDeploymentPolicy_SingleWeightPolicyConfig]) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.SingleWeightPolicyConfig = v
	return t
}

// SetV__SingleWeightPolicyConfig updates property "SingleWeightPolicyConfig".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) SetV__SingleWeightPolicyConfig(v AWS_CloudFront_ContinuousDeploymentPolicy_SingleWeightPolicyConfig) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.SingleWeightPolicyConfig = cfz.V(v)
	return t
}

// Set__StagingDistributionDnsNames updates property "StagingDistributionDnsNames".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) Set__StagingDistributionDnsNames(v cfz.ExpressionSlice[string]) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.StagingDistributionDnsNames = v
	return t
}

// SetS__StagingDistributionDnsNames updates property "StagingDistributionDnsNames".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) SetS__StagingDistributionDnsNames(v ...cfz.Expression[string]) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.StagingDistributionDnsNames = cfz.S(v...)
	return t
}

// SetSV__StagingDistributionDnsNames updates property "StagingDistributionDnsNames".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) SetSV__StagingDistributionDnsNames(v ...string) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.StagingDistributionDnsNames = cfz.SV(v...)
	return t
}

// Set__TrafficConfig updates property "TrafficConfig".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) Set__TrafficConfig(v cfz.Expression[AWS_CloudFront_ContinuousDeploymentPolicy_TrafficConfig]) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.TrafficConfig = v
	return t
}

// SetV__TrafficConfig updates property "TrafficConfig".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) SetV__TrafficConfig(v AWS_CloudFront_ContinuousDeploymentPolicy_TrafficConfig) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.TrafficConfig = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) Set__Type(v cfz.Expression[string]) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig) SetV__Type(v string) AWS_CloudFront_ContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfig {
	t.Type = cfz.V(v)
	return t
}
