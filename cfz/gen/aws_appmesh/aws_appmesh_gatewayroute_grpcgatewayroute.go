// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_GatewayRoute_GrpcGatewayRoute__Type is the CloudFormation type for AWS::AppMesh::GatewayRoute.GrpcGatewayRoute.
	AWS_AppMesh_GatewayRoute_GrpcGatewayRoute__Type = "AWS::AppMesh::GatewayRoute.GrpcGatewayRoute"
)

var (
	// AWS_AppMesh_GatewayRoute_GrpcGatewayRoute__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::GatewayRoute.GrpcGatewayRoute.
	AWS_AppMesh_GatewayRoute_GrpcGatewayRoute__PropertiesMap = struct {
		Action string
		Match  string
	}{
		Action: "Action",
		Match:  "Match",
	}

	// AWS_AppMesh_GatewayRoute_GrpcGatewayRoute__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::GatewayRoute.GrpcGatewayRoute.
	AWS_AppMesh_GatewayRoute_GrpcGatewayRoute__PropertiesSlice = []string{
		AWS_AppMesh_GatewayRoute_GrpcGatewayRoute__PropertiesMap.Action,
		AWS_AppMesh_GatewayRoute_GrpcGatewayRoute__PropertiesMap.Match,
	}
)

// AWS_AppMesh_GatewayRoute_GrpcGatewayRoute is a binding for AWS::AppMesh::GatewayRoute.GrpcGatewayRoute.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroute.html
type AWS_AppMesh_GatewayRoute_GrpcGatewayRoute struct {
	// Action is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroute.html#cfn-appmesh-gatewayroute-grpcgatewayroute-action
	Action cfz.Expression[AWS_AppMesh_GatewayRoute_GrpcGatewayRouteAction] `json:"Action,omitempty"`

	// Match is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroute.html#cfn-appmesh-gatewayroute-grpcgatewayroute-match
	Match cfz.Expression[AWS_AppMesh_GatewayRoute_GrpcGatewayRouteMatch] `json:"Match,omitempty"`
}

// New__AWS_AppMesh_GatewayRoute_GrpcGatewayRoute initializes a new AWS_AppMesh_GatewayRoute_GrpcGatewayRoute.
func New__AWS_AppMesh_GatewayRoute_GrpcGatewayRoute() AWS_AppMesh_GatewayRoute_GrpcGatewayRoute {
	return AWS_AppMesh_GatewayRoute_GrpcGatewayRoute{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_GatewayRoute_GrpcGatewayRoute) GetType() string {
	return AWS_AppMesh_GatewayRoute_GrpcGatewayRoute__Type
}

// Set__Action updates property "Action".
func (t AWS_AppMesh_GatewayRoute_GrpcGatewayRoute) Set__Action(v cfz.Expression[AWS_AppMesh_GatewayRoute_GrpcGatewayRouteAction]) AWS_AppMesh_GatewayRoute_GrpcGatewayRoute {
	t.Action = v
	return t
}

// SetV__Action updates property "Action".
func (t AWS_AppMesh_GatewayRoute_GrpcGatewayRoute) SetV__Action(v AWS_AppMesh_GatewayRoute_GrpcGatewayRouteAction) AWS_AppMesh_GatewayRoute_GrpcGatewayRoute {
	t.Action = cfz.V(v)
	return t
}

// Set__Match updates property "Match".
func (t AWS_AppMesh_GatewayRoute_GrpcGatewayRoute) Set__Match(v cfz.Expression[AWS_AppMesh_GatewayRoute_GrpcGatewayRouteMatch]) AWS_AppMesh_GatewayRoute_GrpcGatewayRoute {
	t.Match = v
	return t
}

// SetV__Match updates property "Match".
func (t AWS_AppMesh_GatewayRoute_GrpcGatewayRoute) SetV__Match(v AWS_AppMesh_GatewayRoute_GrpcGatewayRouteMatch) AWS_AppMesh_GatewayRoute_GrpcGatewayRoute {
	t.Match = cfz.V(v)
	return t
}
