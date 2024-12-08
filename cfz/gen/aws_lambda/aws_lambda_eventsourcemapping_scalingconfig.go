// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lambda

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lambda_EventSourceMapping_ScalingConfig__Type is the CloudFormation type for AWS::Lambda::EventSourceMapping.ScalingConfig.
	AWS_Lambda_EventSourceMapping_ScalingConfig__Type = "AWS::Lambda::EventSourceMapping.ScalingConfig"
)

var (
	// AWS_Lambda_EventSourceMapping_ScalingConfig__PropertiesMap reports all the CloudFormation properties for AWS::Lambda::EventSourceMapping.ScalingConfig.
	AWS_Lambda_EventSourceMapping_ScalingConfig__PropertiesMap = struct {
		MaximumConcurrency string
	}{
		MaximumConcurrency: "MaximumConcurrency",
	}

	// AWS_Lambda_EventSourceMapping_ScalingConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Lambda::EventSourceMapping.ScalingConfig.
	AWS_Lambda_EventSourceMapping_ScalingConfig__PropertiesSlice = []string{
		AWS_Lambda_EventSourceMapping_ScalingConfig__PropertiesMap.MaximumConcurrency,
	}
)

// AWS_Lambda_EventSourceMapping_ScalingConfig is a binding for AWS::Lambda::EventSourceMapping.ScalingConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-scalingconfig.html
type AWS_Lambda_EventSourceMapping_ScalingConfig struct {
	// MaximumConcurrency is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-scalingconfig.html#cfn-lambda-eventsourcemapping-scalingconfig-maximumconcurrency
	MaximumConcurrency cfz.Expression[int32] `json:"MaximumConcurrency,omitempty"`
}

// New__AWS_Lambda_EventSourceMapping_ScalingConfig initializes a new AWS_Lambda_EventSourceMapping_ScalingConfig.
func New__AWS_Lambda_EventSourceMapping_ScalingConfig() AWS_Lambda_EventSourceMapping_ScalingConfig {
	return AWS_Lambda_EventSourceMapping_ScalingConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Lambda_EventSourceMapping_ScalingConfig) GetType() string {
	return AWS_Lambda_EventSourceMapping_ScalingConfig__Type
}

// Set__MaximumConcurrency updates property "MaximumConcurrency".
func (t AWS_Lambda_EventSourceMapping_ScalingConfig) Set__MaximumConcurrency(v cfz.Expression[int32]) AWS_Lambda_EventSourceMapping_ScalingConfig {
	t.MaximumConcurrency = v
	return t
}

// SetV__MaximumConcurrency updates property "MaximumConcurrency".
func (t AWS_Lambda_EventSourceMapping_ScalingConfig) SetV__MaximumConcurrency(v int32) AWS_Lambda_EventSourceMapping_ScalingConfig {
	t.MaximumConcurrency = cfz.V(v)
	return t
}
