// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pipes

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__Type is the CloudFormation type for AWS::Pipes::Pipe.PipeSourceManagedStreamingKafkaParameters.
	AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__Type = "AWS::Pipes::Pipe.PipeSourceManagedStreamingKafkaParameters"
)

var (
	// AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesMap reports all the CloudFormation properties for AWS::Pipes::Pipe.PipeSourceManagedStreamingKafkaParameters.
	AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesMap = struct {
		BatchSize                      string
		ConsumerGroupID                string
		Credentials                    string
		MaximumBatchingWindowInSeconds string
		StartingPosition               string
		TopicName                      string
	}{
		BatchSize:                      "BatchSize",
		ConsumerGroupID:                "ConsumerGroupID",
		Credentials:                    "Credentials",
		MaximumBatchingWindowInSeconds: "MaximumBatchingWindowInSeconds",
		StartingPosition:               "StartingPosition",
		TopicName:                      "TopicName",
	}

	// AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesSlice reports all the CloudFormation properties for AWS::Pipes::Pipe.PipeSourceManagedStreamingKafkaParameters.
	AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesSlice = []string{
		AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesMap.BatchSize,
		AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesMap.ConsumerGroupID,
		AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesMap.Credentials,
		AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesMap.MaximumBatchingWindowInSeconds,
		AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesMap.StartingPosition,
		AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__PropertiesMap.TopicName,
	}
)

// AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters is a binding for AWS::Pipes::Pipe.PipeSourceManagedStreamingKafkaParameters.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html
type AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters struct {
	// BatchSize is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-batchsize
	BatchSize cfz.Expression[int32] `json:"BatchSize,omitempty"`

	// ConsumerGroupID is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-consumergroupid
	ConsumerGroupID cfz.Expression[string] `json:"ConsumerGroupID,omitempty"`

	// Credentials is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-credentials
	Credentials cfz.Expression[AWS_Pipes_Pipe_MSKAccessCredentials] `json:"Credentials,omitempty"`

	// MaximumBatchingWindowInSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-maximumbatchingwindowinseconds
	MaximumBatchingWindowInSeconds cfz.Expression[int32] `json:"MaximumBatchingWindowInSeconds,omitempty"`

	// StartingPosition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-startingposition
	StartingPosition cfz.Expression[string] `json:"StartingPosition,omitempty"`

	// TopicName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-topicname
	TopicName cfz.Expression[string] `json:"TopicName,omitempty"`
}

// New__AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters initializes a new AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters.
func New__AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters() AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	return AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters{}
}

// GetType returns the CloudFormation type.
func (AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) GetType() string {
	return AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters__Type
}

// Set__BatchSize updates property "BatchSize".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) Set__BatchSize(v cfz.Expression[int32]) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.BatchSize = v
	return t
}

// SetV__BatchSize updates property "BatchSize".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) SetV__BatchSize(v int32) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.BatchSize = cfz.V(v)
	return t
}

// Set__ConsumerGroupID updates property "ConsumerGroupID".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) Set__ConsumerGroupID(v cfz.Expression[string]) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.ConsumerGroupID = v
	return t
}

// SetV__ConsumerGroupID updates property "ConsumerGroupID".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) SetV__ConsumerGroupID(v string) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.ConsumerGroupID = cfz.V(v)
	return t
}

// Set__Credentials updates property "Credentials".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) Set__Credentials(v cfz.Expression[AWS_Pipes_Pipe_MSKAccessCredentials]) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.Credentials = v
	return t
}

// SetV__Credentials updates property "Credentials".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) SetV__Credentials(v AWS_Pipes_Pipe_MSKAccessCredentials) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.Credentials = cfz.V(v)
	return t
}

// Set__MaximumBatchingWindowInSeconds updates property "MaximumBatchingWindowInSeconds".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) Set__MaximumBatchingWindowInSeconds(v cfz.Expression[int32]) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.MaximumBatchingWindowInSeconds = v
	return t
}

// SetV__MaximumBatchingWindowInSeconds updates property "MaximumBatchingWindowInSeconds".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) SetV__MaximumBatchingWindowInSeconds(v int32) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.MaximumBatchingWindowInSeconds = cfz.V(v)
	return t
}

// Set__StartingPosition updates property "StartingPosition".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) Set__StartingPosition(v cfz.Expression[string]) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.StartingPosition = v
	return t
}

// SetV__StartingPosition updates property "StartingPosition".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) SetV__StartingPosition(v string) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.StartingPosition = cfz.V(v)
	return t
}

// Set__TopicName updates property "TopicName".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) Set__TopicName(v cfz.Expression[string]) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.TopicName = v
	return t
}

// SetV__TopicName updates property "TopicName".
func (t AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters) SetV__TopicName(v string) AWS_Pipes_Pipe_PipeSourceManagedStreamingKafkaParameters {
	t.TopicName = cfz.V(v)
	return t
}
