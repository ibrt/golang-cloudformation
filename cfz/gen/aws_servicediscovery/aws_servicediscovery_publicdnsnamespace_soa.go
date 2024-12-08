// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_servicediscovery

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ServiceDiscovery_PublicDnsNamespace_SOA__Type is the CloudFormation type for AWS::ServiceDiscovery::PublicDnsNamespace.SOA.
	AWS_ServiceDiscovery_PublicDnsNamespace_SOA__Type = "AWS::ServiceDiscovery::PublicDnsNamespace.SOA"
)

var (
	// AWS_ServiceDiscovery_PublicDnsNamespace_SOA__PropertiesMap reports all the CloudFormation properties for AWS::ServiceDiscovery::PublicDnsNamespace.SOA.
	AWS_ServiceDiscovery_PublicDnsNamespace_SOA__PropertiesMap = struct {
		TTL string
	}{
		TTL: "TTL",
	}

	// AWS_ServiceDiscovery_PublicDnsNamespace_SOA__PropertiesSlice reports all the CloudFormation properties for AWS::ServiceDiscovery::PublicDnsNamespace.SOA.
	AWS_ServiceDiscovery_PublicDnsNamespace_SOA__PropertiesSlice = []string{
		AWS_ServiceDiscovery_PublicDnsNamespace_SOA__PropertiesMap.TTL,
	}
)

// AWS_ServiceDiscovery_PublicDnsNamespace_SOA is a binding for AWS::ServiceDiscovery::PublicDnsNamespace.SOA.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-publicdnsnamespace-soa.html
type AWS_ServiceDiscovery_PublicDnsNamespace_SOA struct {
	// TTL is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-publicdnsnamespace-soa.html#cfn-servicediscovery-publicdnsnamespace-soa-ttl
	TTL cfz.Expression[float64] `json:"TTL,omitempty"`
}

// New__AWS_ServiceDiscovery_PublicDnsNamespace_SOA initializes a new AWS_ServiceDiscovery_PublicDnsNamespace_SOA.
func New__AWS_ServiceDiscovery_PublicDnsNamespace_SOA() AWS_ServiceDiscovery_PublicDnsNamespace_SOA {
	return AWS_ServiceDiscovery_PublicDnsNamespace_SOA{}
}

// GetType returns the CloudFormation type.
func (AWS_ServiceDiscovery_PublicDnsNamespace_SOA) GetType() string {
	return AWS_ServiceDiscovery_PublicDnsNamespace_SOA__Type
}

// Set__TTL updates property "TTL".
func (t AWS_ServiceDiscovery_PublicDnsNamespace_SOA) Set__TTL(v cfz.Expression[float64]) AWS_ServiceDiscovery_PublicDnsNamespace_SOA {
	t.TTL = v
	return t
}

// SetV__TTL updates property "TTL".
func (t AWS_ServiceDiscovery_PublicDnsNamespace_SOA) SetV__TTL(v float64) AWS_ServiceDiscovery_PublicDnsNamespace_SOA {
	t.TTL = cfz.V(v)
	return t
}
