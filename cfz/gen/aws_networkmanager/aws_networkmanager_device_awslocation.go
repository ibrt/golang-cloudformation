// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_networkmanager

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_NetworkManager_Device_AWSLocation__Type is the CloudFormation type for AWS::NetworkManager::Device.AWSLocation.
	AWS_NetworkManager_Device_AWSLocation__Type = "AWS::NetworkManager::Device.AWSLocation"
)

var (
	// AWS_NetworkManager_Device_AWSLocation__PropertiesMap reports all the CloudFormation properties for AWS::NetworkManager::Device.AWSLocation.
	AWS_NetworkManager_Device_AWSLocation__PropertiesMap = struct {
		SubnetArn string
		Zone      string
	}{
		SubnetArn: "SubnetArn",
		Zone:      "Zone",
	}

	// AWS_NetworkManager_Device_AWSLocation__PropertiesSlice reports all the CloudFormation properties for AWS::NetworkManager::Device.AWSLocation.
	AWS_NetworkManager_Device_AWSLocation__PropertiesSlice = []string{
		AWS_NetworkManager_Device_AWSLocation__PropertiesMap.SubnetArn,
		AWS_NetworkManager_Device_AWSLocation__PropertiesMap.Zone,
	}
)

// AWS_NetworkManager_Device_AWSLocation is a binding for AWS::NetworkManager::Device.AWSLocation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-awslocation.html
type AWS_NetworkManager_Device_AWSLocation struct {
	// SubnetArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-awslocation.html#cfn-networkmanager-device-awslocation-subnetarn
	SubnetArn cfz.Expression[string] `json:"SubnetArn,omitempty"`

	// Zone is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-awslocation.html#cfn-networkmanager-device-awslocation-zone
	Zone cfz.Expression[string] `json:"Zone,omitempty"`
}

// New__AWS_NetworkManager_Device_AWSLocation initializes a new AWS_NetworkManager_Device_AWSLocation.
func New__AWS_NetworkManager_Device_AWSLocation() AWS_NetworkManager_Device_AWSLocation {
	return AWS_NetworkManager_Device_AWSLocation{}
}

// GetType returns the CloudFormation type.
func (AWS_NetworkManager_Device_AWSLocation) GetType() string {
	return AWS_NetworkManager_Device_AWSLocation__Type
}

// Set__SubnetArn updates property "SubnetArn".
func (t AWS_NetworkManager_Device_AWSLocation) Set__SubnetArn(v cfz.Expression[string]) AWS_NetworkManager_Device_AWSLocation {
	t.SubnetArn = v
	return t
}

// SetV__SubnetArn updates property "SubnetArn".
func (t AWS_NetworkManager_Device_AWSLocation) SetV__SubnetArn(v string) AWS_NetworkManager_Device_AWSLocation {
	t.SubnetArn = cfz.V(v)
	return t
}

// Set__Zone updates property "Zone".
func (t AWS_NetworkManager_Device_AWSLocation) Set__Zone(v cfz.Expression[string]) AWS_NetworkManager_Device_AWSLocation {
	t.Zone = v
	return t
}

// SetV__Zone updates property "Zone".
func (t AWS_NetworkManager_Device_AWSLocation) SetV__Zone(v string) AWS_NetworkManager_Device_AWSLocation {
	t.Zone = cfz.V(v)
	return t
}
