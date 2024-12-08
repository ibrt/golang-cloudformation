// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apprunner

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration__Type is the CloudFormation type for AWS::AppRunner::VpcIngressConnection.IngressVpcConfiguration.
	AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration__Type = "AWS::AppRunner::VpcIngressConnection.IngressVpcConfiguration"
)

var (
	// AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::AppRunner::VpcIngressConnection.IngressVpcConfiguration.
	AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration__PropertiesMap = struct {
		VpcEndpointId string
		VpcId         string
	}{
		VpcEndpointId: "VpcEndpointId",
		VpcId:         "VpcId",
	}

	// AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::AppRunner::VpcIngressConnection.IngressVpcConfiguration.
	AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration__PropertiesSlice = []string{
		AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration__PropertiesMap.VpcEndpointId,
		AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration__PropertiesMap.VpcId,
	}
)

// AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration is a binding for AWS::AppRunner::VpcIngressConnection.IngressVpcConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-vpcingressconnection-ingressvpcconfiguration.html
type AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration struct {
	// VpcEndpointId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-vpcingressconnection-ingressvpcconfiguration.html#cfn-apprunner-vpcingressconnection-ingressvpcconfiguration-vpcendpointid
	VpcEndpointId cfz.Expression[string] `json:"VpcEndpointId,omitempty"`

	// VpcId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-vpcingressconnection-ingressvpcconfiguration.html#cfn-apprunner-vpcingressconnection-ingressvpcconfiguration-vpcid
	VpcId cfz.Expression[string] `json:"VpcId,omitempty"`
}

// New__AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration initializes a new AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration.
func New__AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration() AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration {
	return AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration) GetType() string {
	return AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration__Type
}

// Set__VpcEndpointId updates property "VpcEndpointId".
func (t AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration) Set__VpcEndpointId(v cfz.Expression[string]) AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration {
	t.VpcEndpointId = v
	return t
}

// SetV__VpcEndpointId updates property "VpcEndpointId".
func (t AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration) SetV__VpcEndpointId(v string) AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration {
	t.VpcEndpointId = cfz.V(v)
	return t
}

// Set__VpcId updates property "VpcId".
func (t AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration) Set__VpcId(v cfz.Expression[string]) AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration {
	t.VpcId = v
	return t
}

// SetV__VpcId updates property "VpcId".
func (t AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration) SetV__VpcId(v string) AWS_AppRunner_VpcIngressConnection_IngressVpcConfiguration {
	t.VpcId = cfz.V(v)
	return t
}
