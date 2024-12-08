// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pipes

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__Type is the CloudFormation type for AWS::Pipes::Pipe.PipeSourceRabbitMQBrokerParameters.
	AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__Type = "AWS::Pipes::Pipe.PipeSourceRabbitMQBrokerParameters"
)

var (
	// AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__PropertiesMap reports all the CloudFormation properties for AWS::Pipes::Pipe.PipeSourceRabbitMQBrokerParameters.
	AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__PropertiesMap = struct {
		BatchSize                      string
		Credentials                    string
		MaximumBatchingWindowInSeconds string
		QueueName                      string
		VirtualHost                    string
	}{
		BatchSize:                      "BatchSize",
		Credentials:                    "Credentials",
		MaximumBatchingWindowInSeconds: "MaximumBatchingWindowInSeconds",
		QueueName:                      "QueueName",
		VirtualHost:                    "VirtualHost",
	}

	// AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__PropertiesSlice reports all the CloudFormation properties for AWS::Pipes::Pipe.PipeSourceRabbitMQBrokerParameters.
	AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__PropertiesSlice = []string{
		AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__PropertiesMap.BatchSize,
		AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__PropertiesMap.Credentials,
		AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__PropertiesMap.MaximumBatchingWindowInSeconds,
		AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__PropertiesMap.QueueName,
		AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__PropertiesMap.VirtualHost,
	}
)

// AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters is a binding for AWS::Pipes::Pipe.PipeSourceRabbitMQBrokerParameters.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html
type AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters struct {
	// BatchSize is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-batchsize
	BatchSize cfz.Expression[int32] `json:"BatchSize,omitempty"`

	// Credentials is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-credentials
	Credentials cfz.Expression[AWS_Pipes_Pipe_MQBrokerAccessCredentials] `json:"Credentials,omitempty"`

	// MaximumBatchingWindowInSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-maximumbatchingwindowinseconds
	MaximumBatchingWindowInSeconds cfz.Expression[int32] `json:"MaximumBatchingWindowInSeconds,omitempty"`

	// QueueName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-queuename
	QueueName cfz.Expression[string] `json:"QueueName,omitempty"`

	// VirtualHost is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html#cfn-pipes-pipe-pipesourcerabbitmqbrokerparameters-virtualhost
	VirtualHost cfz.Expression[string] `json:"VirtualHost,omitempty"`
}

// New__AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters initializes a new AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters.
func New__AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters() AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	return AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters{}
}

// GetType returns the CloudFormation type.
func (AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) GetType() string {
	return AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters__Type
}

// Set__BatchSize updates property "BatchSize".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) Set__BatchSize(v cfz.Expression[int32]) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.BatchSize = v
	return t
}

// SetV__BatchSize updates property "BatchSize".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) SetV__BatchSize(v int32) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.BatchSize = cfz.V(v)
	return t
}

// Set__Credentials updates property "Credentials".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) Set__Credentials(v cfz.Expression[AWS_Pipes_Pipe_MQBrokerAccessCredentials]) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.Credentials = v
	return t
}

// SetV__Credentials updates property "Credentials".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) SetV__Credentials(v AWS_Pipes_Pipe_MQBrokerAccessCredentials) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.Credentials = cfz.V(v)
	return t
}

// Set__MaximumBatchingWindowInSeconds updates property "MaximumBatchingWindowInSeconds".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) Set__MaximumBatchingWindowInSeconds(v cfz.Expression[int32]) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.MaximumBatchingWindowInSeconds = v
	return t
}

// SetV__MaximumBatchingWindowInSeconds updates property "MaximumBatchingWindowInSeconds".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) SetV__MaximumBatchingWindowInSeconds(v int32) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.MaximumBatchingWindowInSeconds = cfz.V(v)
	return t
}

// Set__QueueName updates property "QueueName".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) Set__QueueName(v cfz.Expression[string]) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.QueueName = v
	return t
}

// SetV__QueueName updates property "QueueName".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) SetV__QueueName(v string) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.QueueName = cfz.V(v)
	return t
}

// Set__VirtualHost updates property "VirtualHost".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) Set__VirtualHost(v cfz.Expression[string]) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.VirtualHost = v
	return t
}

// SetV__VirtualHost updates property "VirtualHost".
func (t AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters) SetV__VirtualHost(v string) AWS_Pipes_Pipe_PipeSourceRabbitMQBrokerParameters {
	t.VirtualHost = cfz.V(v)
	return t
}
