// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_MonitoringSchedule_Json__Type is the CloudFormation type for AWS::SageMaker::MonitoringSchedule.Json.
	AWS_SageMaker_MonitoringSchedule_Json__Type = "AWS::SageMaker::MonitoringSchedule.Json"
)

var (
	// AWS_SageMaker_MonitoringSchedule_Json__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::MonitoringSchedule.Json.
	AWS_SageMaker_MonitoringSchedule_Json__PropertiesMap = struct {
		Line string
	}{
		Line: "Line",
	}

	// AWS_SageMaker_MonitoringSchedule_Json__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::MonitoringSchedule.Json.
	AWS_SageMaker_MonitoringSchedule_Json__PropertiesSlice = []string{
		AWS_SageMaker_MonitoringSchedule_Json__PropertiesMap.Line,
	}
)

// AWS_SageMaker_MonitoringSchedule_Json is a binding for AWS::SageMaker::MonitoringSchedule.Json.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-json.html
type AWS_SageMaker_MonitoringSchedule_Json struct {
	// Line is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-json.html#cfn-sagemaker-monitoringschedule-json-line
	Line cfz.Expression[bool] `json:"Line,omitempty"`
}

// New__AWS_SageMaker_MonitoringSchedule_Json initializes a new AWS_SageMaker_MonitoringSchedule_Json.
func New__AWS_SageMaker_MonitoringSchedule_Json() AWS_SageMaker_MonitoringSchedule_Json {
	return AWS_SageMaker_MonitoringSchedule_Json{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_MonitoringSchedule_Json) GetType() string {
	return AWS_SageMaker_MonitoringSchedule_Json__Type
}

// Set__Line updates property "Line".
func (t AWS_SageMaker_MonitoringSchedule_Json) Set__Line(v cfz.Expression[bool]) AWS_SageMaker_MonitoringSchedule_Json {
	t.Line = v
	return t
}

// SetV__Line updates property "Line".
func (t AWS_SageMaker_MonitoringSchedule_Json) SetV__Line(v bool) AWS_SageMaker_MonitoringSchedule_Json {
	t.Line = cfz.V(v)
	return t
}
