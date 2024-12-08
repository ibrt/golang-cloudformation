// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_managedblockchain

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ManagedBlockchain_Node_NodeConfiguration__Type is the CloudFormation type for AWS::ManagedBlockchain::Node.NodeConfiguration.
	AWS_ManagedBlockchain_Node_NodeConfiguration__Type = "AWS::ManagedBlockchain::Node.NodeConfiguration"
)

var (
	// AWS_ManagedBlockchain_Node_NodeConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::ManagedBlockchain::Node.NodeConfiguration.
	AWS_ManagedBlockchain_Node_NodeConfiguration__PropertiesMap = struct {
		AvailabilityZone string
		InstanceType     string
	}{
		AvailabilityZone: "AvailabilityZone",
		InstanceType:     "InstanceType",
	}

	// AWS_ManagedBlockchain_Node_NodeConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::ManagedBlockchain::Node.NodeConfiguration.
	AWS_ManagedBlockchain_Node_NodeConfiguration__PropertiesSlice = []string{
		AWS_ManagedBlockchain_Node_NodeConfiguration__PropertiesMap.AvailabilityZone,
		AWS_ManagedBlockchain_Node_NodeConfiguration__PropertiesMap.InstanceType,
	}
)

// AWS_ManagedBlockchain_Node_NodeConfiguration is a binding for AWS::ManagedBlockchain::Node.NodeConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-node-nodeconfiguration.html
type AWS_ManagedBlockchain_Node_NodeConfiguration struct {
	// AvailabilityZone is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-node-nodeconfiguration.html#cfn-managedblockchain-node-nodeconfiguration-availabilityzone
	AvailabilityZone cfz.Expression[string] `json:"AvailabilityZone,omitempty"`

	// InstanceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-node-nodeconfiguration.html#cfn-managedblockchain-node-nodeconfiguration-instancetype
	InstanceType cfz.Expression[string] `json:"InstanceType,omitempty"`
}

// New__AWS_ManagedBlockchain_Node_NodeConfiguration initializes a new AWS_ManagedBlockchain_Node_NodeConfiguration.
func New__AWS_ManagedBlockchain_Node_NodeConfiguration() AWS_ManagedBlockchain_Node_NodeConfiguration {
	return AWS_ManagedBlockchain_Node_NodeConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_ManagedBlockchain_Node_NodeConfiguration) GetType() string {
	return AWS_ManagedBlockchain_Node_NodeConfiguration__Type
}

// Set__AvailabilityZone updates property "AvailabilityZone".
func (t AWS_ManagedBlockchain_Node_NodeConfiguration) Set__AvailabilityZone(v cfz.Expression[string]) AWS_ManagedBlockchain_Node_NodeConfiguration {
	t.AvailabilityZone = v
	return t
}

// SetV__AvailabilityZone updates property "AvailabilityZone".
func (t AWS_ManagedBlockchain_Node_NodeConfiguration) SetV__AvailabilityZone(v string) AWS_ManagedBlockchain_Node_NodeConfiguration {
	t.AvailabilityZone = cfz.V(v)
	return t
}

// Set__InstanceType updates property "InstanceType".
func (t AWS_ManagedBlockchain_Node_NodeConfiguration) Set__InstanceType(v cfz.Expression[string]) AWS_ManagedBlockchain_Node_NodeConfiguration {
	t.InstanceType = v
	return t
}

// SetV__InstanceType updates property "InstanceType".
func (t AWS_ManagedBlockchain_Node_NodeConfiguration) SetV__InstanceType(v string) AWS_ManagedBlockchain_Node_NodeConfiguration {
	t.InstanceType = cfz.V(v)
	return t
}
