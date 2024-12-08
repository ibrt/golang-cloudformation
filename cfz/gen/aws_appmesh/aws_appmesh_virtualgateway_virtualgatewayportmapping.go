// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping__Type is the CloudFormation type for AWS::AppMesh::VirtualGateway.VirtualGatewayPortMapping.
	AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping__Type = "AWS::AppMesh::VirtualGateway.VirtualGatewayPortMapping"
)

var (
	// AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewayPortMapping.
	AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping__PropertiesMap = struct {
		Port     string
		Protocol string
	}{
		Port:     "Port",
		Protocol: "Protocol",
	}

	// AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewayPortMapping.
	AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping__PropertiesSlice = []string{
		AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping__PropertiesMap.Port,
		AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping__PropertiesMap.Protocol,
	}
)

// AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping is a binding for AWS::AppMesh::VirtualGateway.VirtualGatewayPortMapping.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayportmapping.html
type AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping struct {
	// Port is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayportmapping.html#cfn-appmesh-virtualgateway-virtualgatewayportmapping-port
	Port cfz.Expression[int32] `json:"Port,omitempty"`

	// Protocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayportmapping.html#cfn-appmesh-virtualgateway-virtualgatewayportmapping-protocol
	Protocol cfz.Expression[string] `json:"Protocol,omitempty"`
}

// New__AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping initializes a new AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping.
func New__AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping() AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping {
	return AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping) GetType() string {
	return AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping__Type
}

// Set__Port updates property "Port".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping) Set__Port(v cfz.Expression[int32]) AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping {
	t.Port = v
	return t
}

// SetV__Port updates property "Port".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping) SetV__Port(v int32) AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping {
	t.Port = cfz.V(v)
	return t
}

// Set__Protocol updates property "Protocol".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping) Set__Protocol(v cfz.Expression[string]) AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping {
	t.Protocol = v
	return t
}

// SetV__Protocol updates property "Protocol".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping) SetV__Protocol(v string) AWS_AppMesh_VirtualGateway_VirtualGatewayPortMapping {
	t.Protocol = cfz.V(v)
	return t
}
