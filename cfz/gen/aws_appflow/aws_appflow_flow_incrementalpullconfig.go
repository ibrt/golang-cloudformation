// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appflow

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppFlow_Flow_IncrementalPullConfig__Type is the CloudFormation type for AWS::AppFlow::Flow.IncrementalPullConfig.
	AWS_AppFlow_Flow_IncrementalPullConfig__Type = "AWS::AppFlow::Flow.IncrementalPullConfig"
)

var (
	// AWS_AppFlow_Flow_IncrementalPullConfig__PropertiesMap reports all the CloudFormation properties for AWS::AppFlow::Flow.IncrementalPullConfig.
	AWS_AppFlow_Flow_IncrementalPullConfig__PropertiesMap = struct {
		DatetimeTypeFieldName string
	}{
		DatetimeTypeFieldName: "DatetimeTypeFieldName",
	}

	// AWS_AppFlow_Flow_IncrementalPullConfig__PropertiesSlice reports all the CloudFormation properties for AWS::AppFlow::Flow.IncrementalPullConfig.
	AWS_AppFlow_Flow_IncrementalPullConfig__PropertiesSlice = []string{
		AWS_AppFlow_Flow_IncrementalPullConfig__PropertiesMap.DatetimeTypeFieldName,
	}
)

// AWS_AppFlow_Flow_IncrementalPullConfig is a binding for AWS::AppFlow::Flow.IncrementalPullConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-incrementalpullconfig.html
type AWS_AppFlow_Flow_IncrementalPullConfig struct {
	// DatetimeTypeFieldName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-incrementalpullconfig.html#cfn-appflow-flow-incrementalpullconfig-datetimetypefieldname
	DatetimeTypeFieldName cfz.Expression[string] `json:"DatetimeTypeFieldName,omitempty"`
}

// New__AWS_AppFlow_Flow_IncrementalPullConfig initializes a new AWS_AppFlow_Flow_IncrementalPullConfig.
func New__AWS_AppFlow_Flow_IncrementalPullConfig() AWS_AppFlow_Flow_IncrementalPullConfig {
	return AWS_AppFlow_Flow_IncrementalPullConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_AppFlow_Flow_IncrementalPullConfig) GetType() string {
	return AWS_AppFlow_Flow_IncrementalPullConfig__Type
}

// Set__DatetimeTypeFieldName updates property "DatetimeTypeFieldName".
func (t AWS_AppFlow_Flow_IncrementalPullConfig) Set__DatetimeTypeFieldName(v cfz.Expression[string]) AWS_AppFlow_Flow_IncrementalPullConfig {
	t.DatetimeTypeFieldName = v
	return t
}

// SetV__DatetimeTypeFieldName updates property "DatetimeTypeFieldName".
func (t AWS_AppFlow_Flow_IncrementalPullConfig) SetV__DatetimeTypeFieldName(v string) AWS_AppFlow_Flow_IncrementalPullConfig {
	t.DatetimeTypeFieldName = cfz.V(v)
	return t
}
