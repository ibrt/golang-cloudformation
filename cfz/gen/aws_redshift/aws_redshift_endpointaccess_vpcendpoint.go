// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_redshift

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Redshift_EndpointAccess_VpcEndpoint__Type is the CloudFormation type for AWS::Redshift::EndpointAccess.VpcEndpoint.
	AWS_Redshift_EndpointAccess_VpcEndpoint__Type = "AWS::Redshift::EndpointAccess.VpcEndpoint"
)

var (
	// AWS_Redshift_EndpointAccess_VpcEndpoint__PropertiesMap reports all the CloudFormation properties for AWS::Redshift::EndpointAccess.VpcEndpoint.
	AWS_Redshift_EndpointAccess_VpcEndpoint__PropertiesMap = struct {
		NetworkInterfaces string
		VpcEndpointId     string
		VpcId             string
	}{
		NetworkInterfaces: "NetworkInterfaces",
		VpcEndpointId:     "VpcEndpointId",
		VpcId:             "VpcId",
	}

	// AWS_Redshift_EndpointAccess_VpcEndpoint__PropertiesSlice reports all the CloudFormation properties for AWS::Redshift::EndpointAccess.VpcEndpoint.
	AWS_Redshift_EndpointAccess_VpcEndpoint__PropertiesSlice = []string{
		AWS_Redshift_EndpointAccess_VpcEndpoint__PropertiesMap.NetworkInterfaces,
		AWS_Redshift_EndpointAccess_VpcEndpoint__PropertiesMap.VpcEndpointId,
		AWS_Redshift_EndpointAccess_VpcEndpoint__PropertiesMap.VpcId,
	}
)

// AWS_Redshift_EndpointAccess_VpcEndpoint is a binding for AWS::Redshift::EndpointAccess.VpcEndpoint.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html
type AWS_Redshift_EndpointAccess_VpcEndpoint struct {
	// NetworkInterfaces is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html#cfn-redshift-endpointaccess-vpcendpoint-networkinterfaces
	NetworkInterfaces cfz.ExpressionSlice[AWS_Redshift_EndpointAccess_NetworkInterface] `json:"NetworkInterfaces,omitempty"`

	// VpcEndpointId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html#cfn-redshift-endpointaccess-vpcendpoint-vpcendpointid
	VpcEndpointId cfz.Expression[string] `json:"VpcEndpointId,omitempty"`

	// VpcId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html#cfn-redshift-endpointaccess-vpcendpoint-vpcid
	VpcId cfz.Expression[string] `json:"VpcId,omitempty"`
}

// New__AWS_Redshift_EndpointAccess_VpcEndpoint initializes a new AWS_Redshift_EndpointAccess_VpcEndpoint.
func New__AWS_Redshift_EndpointAccess_VpcEndpoint() AWS_Redshift_EndpointAccess_VpcEndpoint {
	return AWS_Redshift_EndpointAccess_VpcEndpoint{}
}

// GetType returns the CloudFormation type.
func (AWS_Redshift_EndpointAccess_VpcEndpoint) GetType() string {
	return AWS_Redshift_EndpointAccess_VpcEndpoint__Type
}

// Set__NetworkInterfaces updates property "NetworkInterfaces".
func (t AWS_Redshift_EndpointAccess_VpcEndpoint) Set__NetworkInterfaces(v cfz.ExpressionSlice[AWS_Redshift_EndpointAccess_NetworkInterface]) AWS_Redshift_EndpointAccess_VpcEndpoint {
	t.NetworkInterfaces = v
	return t
}

// SetS__NetworkInterfaces updates property "NetworkInterfaces".
func (t AWS_Redshift_EndpointAccess_VpcEndpoint) SetS__NetworkInterfaces(v ...cfz.Expression[AWS_Redshift_EndpointAccess_NetworkInterface]) AWS_Redshift_EndpointAccess_VpcEndpoint {
	t.NetworkInterfaces = cfz.S(v...)
	return t
}

// SetSV__NetworkInterfaces updates property "NetworkInterfaces".
func (t AWS_Redshift_EndpointAccess_VpcEndpoint) SetSV__NetworkInterfaces(v ...AWS_Redshift_EndpointAccess_NetworkInterface) AWS_Redshift_EndpointAccess_VpcEndpoint {
	t.NetworkInterfaces = cfz.SV(v...)
	return t
}

// Set__VpcEndpointId updates property "VpcEndpointId".
func (t AWS_Redshift_EndpointAccess_VpcEndpoint) Set__VpcEndpointId(v cfz.Expression[string]) AWS_Redshift_EndpointAccess_VpcEndpoint {
	t.VpcEndpointId = v
	return t
}

// SetV__VpcEndpointId updates property "VpcEndpointId".
func (t AWS_Redshift_EndpointAccess_VpcEndpoint) SetV__VpcEndpointId(v string) AWS_Redshift_EndpointAccess_VpcEndpoint {
	t.VpcEndpointId = cfz.V(v)
	return t
}

// Set__VpcId updates property "VpcId".
func (t AWS_Redshift_EndpointAccess_VpcEndpoint) Set__VpcId(v cfz.Expression[string]) AWS_Redshift_EndpointAccess_VpcEndpoint {
	t.VpcId = v
	return t
}

// SetV__VpcId updates property "VpcId".
func (t AWS_Redshift_EndpointAccess_VpcEndpoint) SetV__VpcId(v string) AWS_Redshift_EndpointAccess_VpcEndpoint {
	t.VpcId = cfz.V(v)
	return t
}
