// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotanalytics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration__Type is the CloudFormation type for AWS::IoTAnalytics::Dataset.DeltaTimeSessionWindowConfiguration.
	AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration__Type = "AWS::IoTAnalytics::Dataset.DeltaTimeSessionWindowConfiguration"
)

var (
	// AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::IoTAnalytics::Dataset.DeltaTimeSessionWindowConfiguration.
	AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration__PropertiesMap = struct {
		TimeoutInMinutes string
	}{
		TimeoutInMinutes: "TimeoutInMinutes",
	}

	// AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::IoTAnalytics::Dataset.DeltaTimeSessionWindowConfiguration.
	AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration__PropertiesSlice = []string{
		AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration__PropertiesMap.TimeoutInMinutes,
	}
)

// AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration is a binding for AWS::IoTAnalytics::Dataset.DeltaTimeSessionWindowConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatimesessionwindowconfiguration.html
type AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration struct {
	// TimeoutInMinutes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatimesessionwindowconfiguration.html#cfn-iotanalytics-dataset-deltatimesessionwindowconfiguration-timeoutinminutes
	TimeoutInMinutes cfz.Expression[int32] `json:"TimeoutInMinutes,omitempty"`
}

// New__AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration initializes a new AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration.
func New__AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration() AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration {
	return AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration) GetType() string {
	return AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration__Type
}

// Set__TimeoutInMinutes updates property "TimeoutInMinutes".
func (t AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration) Set__TimeoutInMinutes(v cfz.Expression[int32]) AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration {
	t.TimeoutInMinutes = v
	return t
}

// SetV__TimeoutInMinutes updates property "TimeoutInMinutes".
func (t AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration) SetV__TimeoutInMinutes(v int32) AWS_IoTAnalytics_Dataset_DeltaTimeSessionWindowConfiguration {
	t.TimeoutInMinutes = cfz.V(v)
	return t
}
