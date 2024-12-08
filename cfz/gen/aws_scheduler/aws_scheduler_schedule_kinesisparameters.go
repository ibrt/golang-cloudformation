// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_scheduler

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Scheduler_Schedule_KinesisParameters__Type is the CloudFormation type for AWS::Scheduler::Schedule.KinesisParameters.
	AWS_Scheduler_Schedule_KinesisParameters__Type = "AWS::Scheduler::Schedule.KinesisParameters"
)

var (
	// AWS_Scheduler_Schedule_KinesisParameters__PropertiesMap reports all the CloudFormation properties for AWS::Scheduler::Schedule.KinesisParameters.
	AWS_Scheduler_Schedule_KinesisParameters__PropertiesMap = struct {
		PartitionKey string
	}{
		PartitionKey: "PartitionKey",
	}

	// AWS_Scheduler_Schedule_KinesisParameters__PropertiesSlice reports all the CloudFormation properties for AWS::Scheduler::Schedule.KinesisParameters.
	AWS_Scheduler_Schedule_KinesisParameters__PropertiesSlice = []string{
		AWS_Scheduler_Schedule_KinesisParameters__PropertiesMap.PartitionKey,
	}
)

// AWS_Scheduler_Schedule_KinesisParameters is a binding for AWS::Scheduler::Schedule.KinesisParameters.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-kinesisparameters.html
type AWS_Scheduler_Schedule_KinesisParameters struct {
	// PartitionKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-kinesisparameters.html#cfn-scheduler-schedule-kinesisparameters-partitionkey
	PartitionKey cfz.Expression[string] `json:"PartitionKey,omitempty"`
}

// New__AWS_Scheduler_Schedule_KinesisParameters initializes a new AWS_Scheduler_Schedule_KinesisParameters.
func New__AWS_Scheduler_Schedule_KinesisParameters() AWS_Scheduler_Schedule_KinesisParameters {
	return AWS_Scheduler_Schedule_KinesisParameters{}
}

// GetType returns the CloudFormation type.
func (AWS_Scheduler_Schedule_KinesisParameters) GetType() string {
	return AWS_Scheduler_Schedule_KinesisParameters__Type
}

// Set__PartitionKey updates property "PartitionKey".
func (t AWS_Scheduler_Schedule_KinesisParameters) Set__PartitionKey(v cfz.Expression[string]) AWS_Scheduler_Schedule_KinesisParameters {
	t.PartitionKey = v
	return t
}

// SetV__PartitionKey updates property "PartitionKey".
func (t AWS_Scheduler_Schedule_KinesisParameters) SetV__PartitionKey(v string) AWS_Scheduler_Schedule_KinesisParameters {
	t.PartitionKey = cfz.V(v)
	return t
}
