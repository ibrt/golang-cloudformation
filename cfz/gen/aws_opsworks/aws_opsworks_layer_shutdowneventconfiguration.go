// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_opsworks

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_OpsWorks_Layer_ShutdownEventConfiguration__Type is the CloudFormation type for AWS::OpsWorks::Layer.ShutdownEventConfiguration.
	AWS_OpsWorks_Layer_ShutdownEventConfiguration__Type = "AWS::OpsWorks::Layer.ShutdownEventConfiguration"
)

var (
	// AWS_OpsWorks_Layer_ShutdownEventConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::OpsWorks::Layer.ShutdownEventConfiguration.
	AWS_OpsWorks_Layer_ShutdownEventConfiguration__PropertiesMap = struct {
		DelayUntilElbConnectionsDrained string
		ExecutionTimeout                string
	}{
		DelayUntilElbConnectionsDrained: "DelayUntilElbConnectionsDrained",
		ExecutionTimeout:                "ExecutionTimeout",
	}

	// AWS_OpsWorks_Layer_ShutdownEventConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::OpsWorks::Layer.ShutdownEventConfiguration.
	AWS_OpsWorks_Layer_ShutdownEventConfiguration__PropertiesSlice = []string{
		AWS_OpsWorks_Layer_ShutdownEventConfiguration__PropertiesMap.DelayUntilElbConnectionsDrained,
		AWS_OpsWorks_Layer_ShutdownEventConfiguration__PropertiesMap.ExecutionTimeout,
	}
)

// AWS_OpsWorks_Layer_ShutdownEventConfiguration is a binding for AWS::OpsWorks::Layer.ShutdownEventConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html
type AWS_OpsWorks_Layer_ShutdownEventConfiguration struct {
	// DelayUntilElbConnectionsDrained is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-delayuntilelbconnectionsdrained
	DelayUntilElbConnectionsDrained cfz.Expression[bool] `json:"DelayUntilElbConnectionsDrained,omitempty"`

	// ExecutionTimeout is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-executiontimeout
	ExecutionTimeout cfz.Expression[int32] `json:"ExecutionTimeout,omitempty"`
}

// New__AWS_OpsWorks_Layer_ShutdownEventConfiguration initializes a new AWS_OpsWorks_Layer_ShutdownEventConfiguration.
func New__AWS_OpsWorks_Layer_ShutdownEventConfiguration() AWS_OpsWorks_Layer_ShutdownEventConfiguration {
	return AWS_OpsWorks_Layer_ShutdownEventConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_OpsWorks_Layer_ShutdownEventConfiguration) GetType() string {
	return AWS_OpsWorks_Layer_ShutdownEventConfiguration__Type
}

// Set__DelayUntilElbConnectionsDrained updates property "DelayUntilElbConnectionsDrained".
func (t AWS_OpsWorks_Layer_ShutdownEventConfiguration) Set__DelayUntilElbConnectionsDrained(v cfz.Expression[bool]) AWS_OpsWorks_Layer_ShutdownEventConfiguration {
	t.DelayUntilElbConnectionsDrained = v
	return t
}

// SetV__DelayUntilElbConnectionsDrained updates property "DelayUntilElbConnectionsDrained".
func (t AWS_OpsWorks_Layer_ShutdownEventConfiguration) SetV__DelayUntilElbConnectionsDrained(v bool) AWS_OpsWorks_Layer_ShutdownEventConfiguration {
	t.DelayUntilElbConnectionsDrained = cfz.V(v)
	return t
}

// Set__ExecutionTimeout updates property "ExecutionTimeout".
func (t AWS_OpsWorks_Layer_ShutdownEventConfiguration) Set__ExecutionTimeout(v cfz.Expression[int32]) AWS_OpsWorks_Layer_ShutdownEventConfiguration {
	t.ExecutionTimeout = v
	return t
}

// SetV__ExecutionTimeout updates property "ExecutionTimeout".
func (t AWS_OpsWorks_Layer_ShutdownEventConfiguration) SetV__ExecutionTimeout(v int32) AWS_OpsWorks_Layer_ShutdownEventConfiguration {
	t.ExecutionTimeout = cfz.V(v)
	return t
}
