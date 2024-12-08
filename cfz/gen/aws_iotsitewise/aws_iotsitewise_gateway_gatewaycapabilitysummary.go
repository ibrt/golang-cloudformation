// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotsitewise

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary__Type is the CloudFormation type for AWS::IoTSiteWise::Gateway.GatewayCapabilitySummary.
	AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary__Type = "AWS::IoTSiteWise::Gateway.GatewayCapabilitySummary"
)

var (
	// AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary__PropertiesMap reports all the CloudFormation properties for AWS::IoTSiteWise::Gateway.GatewayCapabilitySummary.
	AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary__PropertiesMap = struct {
		CapabilityConfiguration string
		CapabilityNamespace     string
	}{
		CapabilityConfiguration: "CapabilityConfiguration",
		CapabilityNamespace:     "CapabilityNamespace",
	}

	// AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary__PropertiesSlice reports all the CloudFormation properties for AWS::IoTSiteWise::Gateway.GatewayCapabilitySummary.
	AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary__PropertiesSlice = []string{
		AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary__PropertiesMap.CapabilityConfiguration,
		AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary__PropertiesMap.CapabilityNamespace,
	}
)

// AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary is a binding for AWS::IoTSiteWise::Gateway.GatewayCapabilitySummary.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewaycapabilitysummary.html
type AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary struct {
	// CapabilityConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewaycapabilitysummary.html#cfn-iotsitewise-gateway-gatewaycapabilitysummary-capabilityconfiguration
	CapabilityConfiguration cfz.Expression[string] `json:"CapabilityConfiguration,omitempty"`

	// CapabilityNamespace is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewaycapabilitysummary.html#cfn-iotsitewise-gateway-gatewaycapabilitysummary-capabilitynamespace
	CapabilityNamespace cfz.Expression[string] `json:"CapabilityNamespace,omitempty"`
}

// New__AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary initializes a new AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary.
func New__AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary() AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary {
	return AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary) GetType() string {
	return AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary__Type
}

// Set__CapabilityConfiguration updates property "CapabilityConfiguration".
func (t AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary) Set__CapabilityConfiguration(v cfz.Expression[string]) AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary {
	t.CapabilityConfiguration = v
	return t
}

// SetV__CapabilityConfiguration updates property "CapabilityConfiguration".
func (t AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary) SetV__CapabilityConfiguration(v string) AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary {
	t.CapabilityConfiguration = cfz.V(v)
	return t
}

// Set__CapabilityNamespace updates property "CapabilityNamespace".
func (t AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary) Set__CapabilityNamespace(v cfz.Expression[string]) AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary {
	t.CapabilityNamespace = v
	return t
}

// SetV__CapabilityNamespace updates property "CapabilityNamespace".
func (t AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary) SetV__CapabilityNamespace(v string) AWS_IoTSiteWise_Gateway_GatewayCapabilitySummary {
	t.CapabilityNamespace = cfz.V(v)
	return t
}
