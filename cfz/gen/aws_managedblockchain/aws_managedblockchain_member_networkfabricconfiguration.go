// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_managedblockchain

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ManagedBlockchain_Member_NetworkFabricConfiguration__Type is the CloudFormation type for AWS::ManagedBlockchain::Member.NetworkFabricConfiguration.
	AWS_ManagedBlockchain_Member_NetworkFabricConfiguration__Type = "AWS::ManagedBlockchain::Member.NetworkFabricConfiguration"
)

var (
	// AWS_ManagedBlockchain_Member_NetworkFabricConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::ManagedBlockchain::Member.NetworkFabricConfiguration.
	AWS_ManagedBlockchain_Member_NetworkFabricConfiguration__PropertiesMap = struct {
		Edition string
	}{
		Edition: "Edition",
	}

	// AWS_ManagedBlockchain_Member_NetworkFabricConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::ManagedBlockchain::Member.NetworkFabricConfiguration.
	AWS_ManagedBlockchain_Member_NetworkFabricConfiguration__PropertiesSlice = []string{
		AWS_ManagedBlockchain_Member_NetworkFabricConfiguration__PropertiesMap.Edition,
	}
)

// AWS_ManagedBlockchain_Member_NetworkFabricConfiguration is a binding for AWS::ManagedBlockchain::Member.NetworkFabricConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkfabricconfiguration.html
type AWS_ManagedBlockchain_Member_NetworkFabricConfiguration struct {
	// Edition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkfabricconfiguration.html#cfn-managedblockchain-member-networkfabricconfiguration-edition
	Edition cfz.Expression[string] `json:"Edition,omitempty"`
}

// New__AWS_ManagedBlockchain_Member_NetworkFabricConfiguration initializes a new AWS_ManagedBlockchain_Member_NetworkFabricConfiguration.
func New__AWS_ManagedBlockchain_Member_NetworkFabricConfiguration() AWS_ManagedBlockchain_Member_NetworkFabricConfiguration {
	return AWS_ManagedBlockchain_Member_NetworkFabricConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_ManagedBlockchain_Member_NetworkFabricConfiguration) GetType() string {
	return AWS_ManagedBlockchain_Member_NetworkFabricConfiguration__Type
}

// Set__Edition updates property "Edition".
func (t AWS_ManagedBlockchain_Member_NetworkFabricConfiguration) Set__Edition(v cfz.Expression[string]) AWS_ManagedBlockchain_Member_NetworkFabricConfiguration {
	t.Edition = v
	return t
}

// SetV__Edition updates property "Edition".
func (t AWS_ManagedBlockchain_Member_NetworkFabricConfiguration) SetV__Edition(v string) AWS_ManagedBlockchain_Member_NetworkFabricConfiguration {
	t.Edition = cfz.V(v)
	return t
}
