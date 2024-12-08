// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_synthetics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Synthetics_Canary_RunConfig__Type is the CloudFormation type for AWS::Synthetics::Canary.RunConfig.
	AWS_Synthetics_Canary_RunConfig__Type = "AWS::Synthetics::Canary.RunConfig"
)

var (
	// AWS_Synthetics_Canary_RunConfig__PropertiesMap reports all the CloudFormation properties for AWS::Synthetics::Canary.RunConfig.
	AWS_Synthetics_Canary_RunConfig__PropertiesMap = struct {
		ActiveTracing        string
		EnvironmentVariables string
		MemoryInMB           string
		TimeoutInSeconds     string
	}{
		ActiveTracing:        "ActiveTracing",
		EnvironmentVariables: "EnvironmentVariables",
		MemoryInMB:           "MemoryInMB",
		TimeoutInSeconds:     "TimeoutInSeconds",
	}

	// AWS_Synthetics_Canary_RunConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Synthetics::Canary.RunConfig.
	AWS_Synthetics_Canary_RunConfig__PropertiesSlice = []string{
		AWS_Synthetics_Canary_RunConfig__PropertiesMap.ActiveTracing,
		AWS_Synthetics_Canary_RunConfig__PropertiesMap.EnvironmentVariables,
		AWS_Synthetics_Canary_RunConfig__PropertiesMap.MemoryInMB,
		AWS_Synthetics_Canary_RunConfig__PropertiesMap.TimeoutInSeconds,
	}
)

// AWS_Synthetics_Canary_RunConfig is a binding for AWS::Synthetics::Canary.RunConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html
type AWS_Synthetics_Canary_RunConfig struct {
	// ActiveTracing is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-activetracing
	ActiveTracing cfz.Expression[bool] `json:"ActiveTracing,omitempty"`

	// EnvironmentVariables is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-environmentvariables
	EnvironmentVariables cfz.ExpressionMap[string] `json:"EnvironmentVariables,omitempty"`

	// MemoryInMB is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-memoryinmb
	MemoryInMB cfz.Expression[int32] `json:"MemoryInMB,omitempty"`

	// TimeoutInSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-timeoutinseconds
	TimeoutInSeconds cfz.Expression[int32] `json:"TimeoutInSeconds,omitempty"`
}

// New__AWS_Synthetics_Canary_RunConfig initializes a new AWS_Synthetics_Canary_RunConfig.
func New__AWS_Synthetics_Canary_RunConfig() AWS_Synthetics_Canary_RunConfig {
	return AWS_Synthetics_Canary_RunConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Synthetics_Canary_RunConfig) GetType() string {
	return AWS_Synthetics_Canary_RunConfig__Type
}

// Set__ActiveTracing updates property "ActiveTracing".
func (t AWS_Synthetics_Canary_RunConfig) Set__ActiveTracing(v cfz.Expression[bool]) AWS_Synthetics_Canary_RunConfig {
	t.ActiveTracing = v
	return t
}

// SetV__ActiveTracing updates property "ActiveTracing".
func (t AWS_Synthetics_Canary_RunConfig) SetV__ActiveTracing(v bool) AWS_Synthetics_Canary_RunConfig {
	t.ActiveTracing = cfz.V(v)
	return t
}

// Set__EnvironmentVariables updates property "EnvironmentVariables".
func (t AWS_Synthetics_Canary_RunConfig) Set__EnvironmentVariables(v cfz.ExpressionMap[string]) AWS_Synthetics_Canary_RunConfig {
	t.EnvironmentVariables = v
	return t
}

// SetM__EnvironmentVariables updates property "EnvironmentVariables".
func (t AWS_Synthetics_Canary_RunConfig) SetM__EnvironmentVariables(v ...map[string]cfz.Expression[string]) AWS_Synthetics_Canary_RunConfig {
	t.EnvironmentVariables = cfz.M(v...)
	return t
}

// SetMV__EnvironmentVariables updates property "EnvironmentVariables".
func (t AWS_Synthetics_Canary_RunConfig) SetMV__EnvironmentVariables(v ...map[string]string) AWS_Synthetics_Canary_RunConfig {
	t.EnvironmentVariables = cfz.MV(v...)
	return t
}

// Set__MemoryInMB updates property "MemoryInMB".
func (t AWS_Synthetics_Canary_RunConfig) Set__MemoryInMB(v cfz.Expression[int32]) AWS_Synthetics_Canary_RunConfig {
	t.MemoryInMB = v
	return t
}

// SetV__MemoryInMB updates property "MemoryInMB".
func (t AWS_Synthetics_Canary_RunConfig) SetV__MemoryInMB(v int32) AWS_Synthetics_Canary_RunConfig {
	t.MemoryInMB = cfz.V(v)
	return t
}

// Set__TimeoutInSeconds updates property "TimeoutInSeconds".
func (t AWS_Synthetics_Canary_RunConfig) Set__TimeoutInSeconds(v cfz.Expression[int32]) AWS_Synthetics_Canary_RunConfig {
	t.TimeoutInSeconds = v
	return t
}

// SetV__TimeoutInSeconds updates property "TimeoutInSeconds".
func (t AWS_Synthetics_Canary_RunConfig) SetV__TimeoutInSeconds(v int32) AWS_Synthetics_Canary_RunConfig {
	t.TimeoutInSeconds = cfz.V(v)
	return t
}
