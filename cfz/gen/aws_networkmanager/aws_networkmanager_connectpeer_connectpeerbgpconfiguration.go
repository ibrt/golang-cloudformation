// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_networkmanager

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__Type is the CloudFormation type for AWS::NetworkManager::ConnectPeer.ConnectPeerBgpConfiguration.
	AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__Type = "AWS::NetworkManager::ConnectPeer.ConnectPeerBgpConfiguration"
)

var (
	// AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::NetworkManager::ConnectPeer.ConnectPeerBgpConfiguration.
	AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__PropertiesMap = struct {
		CoreNetworkAddress string
		CoreNetworkAsn     string
		PeerAddress        string
		PeerAsn            string
	}{
		CoreNetworkAddress: "CoreNetworkAddress",
		CoreNetworkAsn:     "CoreNetworkAsn",
		PeerAddress:        "PeerAddress",
		PeerAsn:            "PeerAsn",
	}

	// AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::NetworkManager::ConnectPeer.ConnectPeerBgpConfiguration.
	AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__PropertiesSlice = []string{
		AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__PropertiesMap.CoreNetworkAddress,
		AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__PropertiesMap.CoreNetworkAsn,
		AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__PropertiesMap.PeerAddress,
		AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__PropertiesMap.PeerAsn,
	}
)

// AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration is a binding for AWS::NetworkManager::ConnectPeer.ConnectPeerBgpConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html
type AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration struct {
	// CoreNetworkAddress is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html#cfn-networkmanager-connectpeer-connectpeerbgpconfiguration-corenetworkaddress
	CoreNetworkAddress cfz.Expression[string] `json:"CoreNetworkAddress,omitempty"`

	// CoreNetworkAsn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html#cfn-networkmanager-connectpeer-connectpeerbgpconfiguration-corenetworkasn
	CoreNetworkAsn cfz.Expression[float64] `json:"CoreNetworkAsn,omitempty"`

	// PeerAddress is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html#cfn-networkmanager-connectpeer-connectpeerbgpconfiguration-peeraddress
	PeerAddress cfz.Expression[string] `json:"PeerAddress,omitempty"`

	// PeerAsn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectpeer-connectpeerbgpconfiguration.html#cfn-networkmanager-connectpeer-connectpeerbgpconfiguration-peerasn
	PeerAsn cfz.Expression[float64] `json:"PeerAsn,omitempty"`
}

// New__AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration initializes a new AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration.
func New__AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration() AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration {
	return AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration) GetType() string {
	return AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration__Type
}

// Set__CoreNetworkAddress updates property "CoreNetworkAddress".
func (t AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration) Set__CoreNetworkAddress(v cfz.Expression[string]) AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration {
	t.CoreNetworkAddress = v
	return t
}

// SetV__CoreNetworkAddress updates property "CoreNetworkAddress".
func (t AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration) SetV__CoreNetworkAddress(v string) AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration {
	t.CoreNetworkAddress = cfz.V(v)
	return t
}

// Set__CoreNetworkAsn updates property "CoreNetworkAsn".
func (t AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration) Set__CoreNetworkAsn(v cfz.Expression[float64]) AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration {
	t.CoreNetworkAsn = v
	return t
}

// SetV__CoreNetworkAsn updates property "CoreNetworkAsn".
func (t AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration) SetV__CoreNetworkAsn(v float64) AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration {
	t.CoreNetworkAsn = cfz.V(v)
	return t
}

// Set__PeerAddress updates property "PeerAddress".
func (t AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration) Set__PeerAddress(v cfz.Expression[string]) AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration {
	t.PeerAddress = v
	return t
}

// SetV__PeerAddress updates property "PeerAddress".
func (t AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration) SetV__PeerAddress(v string) AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration {
	t.PeerAddress = cfz.V(v)
	return t
}

// Set__PeerAsn updates property "PeerAsn".
func (t AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration) Set__PeerAsn(v cfz.Expression[float64]) AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration {
	t.PeerAsn = v
	return t
}

// SetV__PeerAsn updates property "PeerAsn".
func (t AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration) SetV__PeerAsn(v float64) AWS_NetworkManager_ConnectPeer_ConnectPeerBgpConfiguration {
	t.PeerAsn = cfz.V(v)
	return t
}
