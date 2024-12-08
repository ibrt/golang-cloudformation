// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite__Type is the CloudFormation type for AWS::AppMesh::GatewayRoute.GrpcGatewayRouteRewrite.
	AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite__Type = "AWS::AppMesh::GatewayRoute.GrpcGatewayRouteRewrite"
)

var (
	// AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::GatewayRoute.GrpcGatewayRouteRewrite.
	AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite__PropertiesMap = struct {
		Hostname string
	}{
		Hostname: "Hostname",
	}

	// AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::GatewayRoute.GrpcGatewayRouteRewrite.
	AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite__PropertiesSlice = []string{
		AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite__PropertiesMap.Hostname,
	}
)

// AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite is a binding for AWS::AppMesh::GatewayRoute.GrpcGatewayRouteRewrite.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayrouterewrite.html
type AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite struct {
	// Hostname is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayrouterewrite.html#cfn-appmesh-gatewayroute-grpcgatewayrouterewrite-hostname
	Hostname cfz.Expression[AWS_AppMesh_GatewayRoute_GatewayRouteHostnameRewrite] `json:"Hostname,omitempty"`
}

// New__AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite initializes a new AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite.
func New__AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite() AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite {
	return AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite) GetType() string {
	return AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite__Type
}

// Set__Hostname updates property "Hostname".
func (t AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite) Set__Hostname(v cfz.Expression[AWS_AppMesh_GatewayRoute_GatewayRouteHostnameRewrite]) AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite {
	t.Hostname = v
	return t
}

// SetV__Hostname updates property "Hostname".
func (t AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite) SetV__Hostname(v AWS_AppMesh_GatewayRoute_GatewayRouteHostnameRewrite) AWS_AppMesh_GatewayRoute_GrpcGatewayRouteRewrite {
	t.Hostname = cfz.V(v)
	return t
}
