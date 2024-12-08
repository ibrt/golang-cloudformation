// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_servicediscovery

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ServiceDiscovery_PrivateDnsNamespace_Properties__Type is the CloudFormation type for AWS::ServiceDiscovery::PrivateDnsNamespace.Properties.
	AWS_ServiceDiscovery_PrivateDnsNamespace_Properties__Type = "AWS::ServiceDiscovery::PrivateDnsNamespace.Properties"
)

var (
	// AWS_ServiceDiscovery_PrivateDnsNamespace_Properties__PropertiesMap reports all the CloudFormation properties for AWS::ServiceDiscovery::PrivateDnsNamespace.Properties.
	AWS_ServiceDiscovery_PrivateDnsNamespace_Properties__PropertiesMap = struct {
		DnsProperties string
	}{
		DnsProperties: "DnsProperties",
	}

	// AWS_ServiceDiscovery_PrivateDnsNamespace_Properties__PropertiesSlice reports all the CloudFormation properties for AWS::ServiceDiscovery::PrivateDnsNamespace.Properties.
	AWS_ServiceDiscovery_PrivateDnsNamespace_Properties__PropertiesSlice = []string{
		AWS_ServiceDiscovery_PrivateDnsNamespace_Properties__PropertiesMap.DnsProperties,
	}
)

// AWS_ServiceDiscovery_PrivateDnsNamespace_Properties is a binding for AWS::ServiceDiscovery::PrivateDnsNamespace.Properties.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-privatednsnamespace-properties.html
type AWS_ServiceDiscovery_PrivateDnsNamespace_Properties struct {
	// DnsProperties is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-privatednsnamespace-properties.html#cfn-servicediscovery-privatednsnamespace-properties-dnsproperties
	DnsProperties cfz.Expression[AWS_ServiceDiscovery_PrivateDnsNamespace_PrivateDnsPropertiesMutable] `json:"DnsProperties,omitempty"`
}

// New__AWS_ServiceDiscovery_PrivateDnsNamespace_Properties initializes a new AWS_ServiceDiscovery_PrivateDnsNamespace_Properties.
func New__AWS_ServiceDiscovery_PrivateDnsNamespace_Properties() AWS_ServiceDiscovery_PrivateDnsNamespace_Properties {
	return AWS_ServiceDiscovery_PrivateDnsNamespace_Properties{}
}

// GetType returns the CloudFormation type.
func (AWS_ServiceDiscovery_PrivateDnsNamespace_Properties) GetType() string {
	return AWS_ServiceDiscovery_PrivateDnsNamespace_Properties__Type
}

// Set__DnsProperties updates property "DnsProperties".
func (t AWS_ServiceDiscovery_PrivateDnsNamespace_Properties) Set__DnsProperties(v cfz.Expression[AWS_ServiceDiscovery_PrivateDnsNamespace_PrivateDnsPropertiesMutable]) AWS_ServiceDiscovery_PrivateDnsNamespace_Properties {
	t.DnsProperties = v
	return t
}

// SetV__DnsProperties updates property "DnsProperties".
func (t AWS_ServiceDiscovery_PrivateDnsNamespace_Properties) SetV__DnsProperties(v AWS_ServiceDiscovery_PrivateDnsNamespace_PrivateDnsPropertiesMutable) AWS_ServiceDiscovery_PrivateDnsNamespace_Properties {
	t.DnsProperties = cfz.V(v)
	return t
}
