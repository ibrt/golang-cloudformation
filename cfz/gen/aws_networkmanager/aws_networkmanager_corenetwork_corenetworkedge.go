// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_networkmanager

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__Type is the CloudFormation type for AWS::NetworkManager::CoreNetwork.CoreNetworkEdge.
	AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__Type = "AWS::NetworkManager::CoreNetwork.CoreNetworkEdge"
)

var (
	// AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__PropertiesMap reports all the CloudFormation properties for AWS::NetworkManager::CoreNetwork.CoreNetworkEdge.
	AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__PropertiesMap = struct {
		Asn              string
		EdgeLocation     string
		InsideCidrBlocks string
	}{
		Asn:              "Asn",
		EdgeLocation:     "EdgeLocation",
		InsideCidrBlocks: "InsideCidrBlocks",
	}

	// AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__PropertiesSlice reports all the CloudFormation properties for AWS::NetworkManager::CoreNetwork.CoreNetworkEdge.
	AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__PropertiesSlice = []string{
		AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__PropertiesMap.Asn,
		AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__PropertiesMap.EdgeLocation,
		AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__PropertiesMap.InsideCidrBlocks,
	}
)

// AWS_NetworkManager_CoreNetwork_CoreNetworkEdge is a binding for AWS::NetworkManager::CoreNetwork.CoreNetworkEdge.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html
type AWS_NetworkManager_CoreNetwork_CoreNetworkEdge struct {
	// Asn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html#cfn-networkmanager-corenetwork-corenetworkedge-asn
	Asn cfz.Expression[float64] `json:"Asn,omitempty"`

	// EdgeLocation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html#cfn-networkmanager-corenetwork-corenetworkedge-edgelocation
	EdgeLocation cfz.Expression[string] `json:"EdgeLocation,omitempty"`

	// InsideCidrBlocks is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html#cfn-networkmanager-corenetwork-corenetworkedge-insidecidrblocks
	InsideCidrBlocks cfz.ExpressionSlice[string] `json:"InsideCidrBlocks,omitempty"`
}

// New__AWS_NetworkManager_CoreNetwork_CoreNetworkEdge initializes a new AWS_NetworkManager_CoreNetwork_CoreNetworkEdge.
func New__AWS_NetworkManager_CoreNetwork_CoreNetworkEdge() AWS_NetworkManager_CoreNetwork_CoreNetworkEdge {
	return AWS_NetworkManager_CoreNetwork_CoreNetworkEdge{}
}

// GetType returns the CloudFormation type.
func (AWS_NetworkManager_CoreNetwork_CoreNetworkEdge) GetType() string {
	return AWS_NetworkManager_CoreNetwork_CoreNetworkEdge__Type
}

// Set__Asn updates property "Asn".
func (t AWS_NetworkManager_CoreNetwork_CoreNetworkEdge) Set__Asn(v cfz.Expression[float64]) AWS_NetworkManager_CoreNetwork_CoreNetworkEdge {
	t.Asn = v
	return t
}

// SetV__Asn updates property "Asn".
func (t AWS_NetworkManager_CoreNetwork_CoreNetworkEdge) SetV__Asn(v float64) AWS_NetworkManager_CoreNetwork_CoreNetworkEdge {
	t.Asn = cfz.V(v)
	return t
}

// Set__EdgeLocation updates property "EdgeLocation".
func (t AWS_NetworkManager_CoreNetwork_CoreNetworkEdge) Set__EdgeLocation(v cfz.Expression[string]) AWS_NetworkManager_CoreNetwork_CoreNetworkEdge {
	t.EdgeLocation = v
	return t
}

// SetV__EdgeLocation updates property "EdgeLocation".
func (t AWS_NetworkManager_CoreNetwork_CoreNetworkEdge) SetV__EdgeLocation(v string) AWS_NetworkManager_CoreNetwork_CoreNetworkEdge {
	t.EdgeLocation = cfz.V(v)
	return t
}

// Set__InsideCidrBlocks updates property "InsideCidrBlocks".
func (t AWS_NetworkManager_CoreNetwork_CoreNetworkEdge) Set__InsideCidrBlocks(v cfz.ExpressionSlice[string]) AWS_NetworkManager_CoreNetwork_CoreNetworkEdge {
	t.InsideCidrBlocks = v
	return t
}

// SetS__InsideCidrBlocks updates property "InsideCidrBlocks".
func (t AWS_NetworkManager_CoreNetwork_CoreNetworkEdge) SetS__InsideCidrBlocks(v ...cfz.Expression[string]) AWS_NetworkManager_CoreNetwork_CoreNetworkEdge {
	t.InsideCidrBlocks = cfz.S(v...)
	return t
}

// SetSV__InsideCidrBlocks updates property "InsideCidrBlocks".
func (t AWS_NetworkManager_CoreNetwork_CoreNetworkEdge) SetSV__InsideCidrBlocks(v ...string) AWS_NetworkManager_CoreNetwork_CoreNetworkEdge {
	t.InsideCidrBlocks = cfz.SV(v...)
	return t
}
