// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kafkaconnect

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KafkaConnect_Connector_ApacheKafkaCluster__Type is the CloudFormation type for AWS::KafkaConnect::Connector.ApacheKafkaCluster.
	AWS_KafkaConnect_Connector_ApacheKafkaCluster__Type = "AWS::KafkaConnect::Connector.ApacheKafkaCluster"
)

var (
	// AWS_KafkaConnect_Connector_ApacheKafkaCluster__PropertiesMap reports all the CloudFormation properties for AWS::KafkaConnect::Connector.ApacheKafkaCluster.
	AWS_KafkaConnect_Connector_ApacheKafkaCluster__PropertiesMap = struct {
		BootstrapServers string
		Vpc              string
	}{
		BootstrapServers: "BootstrapServers",
		Vpc:              "Vpc",
	}

	// AWS_KafkaConnect_Connector_ApacheKafkaCluster__PropertiesSlice reports all the CloudFormation properties for AWS::KafkaConnect::Connector.ApacheKafkaCluster.
	AWS_KafkaConnect_Connector_ApacheKafkaCluster__PropertiesSlice = []string{
		AWS_KafkaConnect_Connector_ApacheKafkaCluster__PropertiesMap.BootstrapServers,
		AWS_KafkaConnect_Connector_ApacheKafkaCluster__PropertiesMap.Vpc,
	}
)

// AWS_KafkaConnect_Connector_ApacheKafkaCluster is a binding for AWS::KafkaConnect::Connector.ApacheKafkaCluster.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-apachekafkacluster.html
type AWS_KafkaConnect_Connector_ApacheKafkaCluster struct {
	// BootstrapServers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-apachekafkacluster.html#cfn-kafkaconnect-connector-apachekafkacluster-bootstrapservers
	BootstrapServers cfz.Expression[string] `json:"BootstrapServers,omitempty"`

	// Vpc is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-apachekafkacluster.html#cfn-kafkaconnect-connector-apachekafkacluster-vpc
	Vpc cfz.Expression[AWS_KafkaConnect_Connector_Vpc] `json:"Vpc,omitempty"`
}

// New__AWS_KafkaConnect_Connector_ApacheKafkaCluster initializes a new AWS_KafkaConnect_Connector_ApacheKafkaCluster.
func New__AWS_KafkaConnect_Connector_ApacheKafkaCluster() AWS_KafkaConnect_Connector_ApacheKafkaCluster {
	return AWS_KafkaConnect_Connector_ApacheKafkaCluster{}
}

// GetType returns the CloudFormation type.
func (AWS_KafkaConnect_Connector_ApacheKafkaCluster) GetType() string {
	return AWS_KafkaConnect_Connector_ApacheKafkaCluster__Type
}

// Set__BootstrapServers updates property "BootstrapServers".
func (t AWS_KafkaConnect_Connector_ApacheKafkaCluster) Set__BootstrapServers(v cfz.Expression[string]) AWS_KafkaConnect_Connector_ApacheKafkaCluster {
	t.BootstrapServers = v
	return t
}

// SetV__BootstrapServers updates property "BootstrapServers".
func (t AWS_KafkaConnect_Connector_ApacheKafkaCluster) SetV__BootstrapServers(v string) AWS_KafkaConnect_Connector_ApacheKafkaCluster {
	t.BootstrapServers = cfz.V(v)
	return t
}

// Set__Vpc updates property "Vpc".
func (t AWS_KafkaConnect_Connector_ApacheKafkaCluster) Set__Vpc(v cfz.Expression[AWS_KafkaConnect_Connector_Vpc]) AWS_KafkaConnect_Connector_ApacheKafkaCluster {
	t.Vpc = v
	return t
}

// SetV__Vpc updates property "Vpc".
func (t AWS_KafkaConnect_Connector_ApacheKafkaCluster) SetV__Vpc(v AWS_KafkaConnect_Connector_Vpc) AWS_KafkaConnect_Connector_ApacheKafkaCluster {
	t.Vpc = cfz.V(v)
	return t
}
