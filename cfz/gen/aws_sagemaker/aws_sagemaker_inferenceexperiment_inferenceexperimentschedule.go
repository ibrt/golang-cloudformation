// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule__Type is the CloudFormation type for AWS::SageMaker::InferenceExperiment.InferenceExperimentSchedule.
	AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule__Type = "AWS::SageMaker::InferenceExperiment.InferenceExperimentSchedule"
)

var (
	// AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::InferenceExperiment.InferenceExperimentSchedule.
	AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule__PropertiesMap = struct {
		EndTime   string
		StartTime string
	}{
		EndTime:   "EndTime",
		StartTime: "StartTime",
	}

	// AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::InferenceExperiment.InferenceExperimentSchedule.
	AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule__PropertiesSlice = []string{
		AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule__PropertiesMap.EndTime,
		AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule__PropertiesMap.StartTime,
	}
)

// AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule is a binding for AWS::SageMaker::InferenceExperiment.InferenceExperimentSchedule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-inferenceexperimentschedule.html
type AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule struct {
	// EndTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-inferenceexperimentschedule.html#cfn-sagemaker-inferenceexperiment-inferenceexperimentschedule-endtime
	EndTime cfz.Expression[string] `json:"EndTime,omitempty"`

	// StartTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-inferenceexperimentschedule.html#cfn-sagemaker-inferenceexperiment-inferenceexperimentschedule-starttime
	StartTime cfz.Expression[string] `json:"StartTime,omitempty"`
}

// New__AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule initializes a new AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule.
func New__AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule() AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule {
	return AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule) GetType() string {
	return AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule__Type
}

// Set__EndTime updates property "EndTime".
func (t AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule) Set__EndTime(v cfz.Expression[string]) AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule {
	t.EndTime = v
	return t
}

// SetV__EndTime updates property "EndTime".
func (t AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule) SetV__EndTime(v string) AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule {
	t.EndTime = cfz.V(v)
	return t
}

// Set__StartTime updates property "StartTime".
func (t AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule) Set__StartTime(v cfz.Expression[string]) AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule {
	t.StartTime = v
	return t
}

// SetV__StartTime updates property "StartTime".
func (t AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule) SetV__StartTime(v string) AWS_SageMaker_InferenceExperiment_InferenceExperimentSchedule {
	t.StartTime = cfz.V(v)
	return t
}
