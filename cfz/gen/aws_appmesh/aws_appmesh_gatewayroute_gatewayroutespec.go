// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_GatewayRoute_GatewayRouteSpec__Type is the CloudFormation type for AWS::AppMesh::GatewayRoute.GatewayRouteSpec.
	AWS_AppMesh_GatewayRoute_GatewayRouteSpec__Type = "AWS::AppMesh::GatewayRoute.GatewayRouteSpec"
)

var (
	// AWS_AppMesh_GatewayRoute_GatewayRouteSpec__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::GatewayRoute.GatewayRouteSpec.
	AWS_AppMesh_GatewayRoute_GatewayRouteSpec__PropertiesMap = struct {
		GrpcRoute  string
		Http2Route string
		HttpRoute  string
		Priority   string
	}{
		GrpcRoute:  "GrpcRoute",
		Http2Route: "Http2Route",
		HttpRoute:  "HttpRoute",
		Priority:   "Priority",
	}

	// AWS_AppMesh_GatewayRoute_GatewayRouteSpec__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::GatewayRoute.GatewayRouteSpec.
	AWS_AppMesh_GatewayRoute_GatewayRouteSpec__PropertiesSlice = []string{
		AWS_AppMesh_GatewayRoute_GatewayRouteSpec__PropertiesMap.GrpcRoute,
		AWS_AppMesh_GatewayRoute_GatewayRouteSpec__PropertiesMap.Http2Route,
		AWS_AppMesh_GatewayRoute_GatewayRouteSpec__PropertiesMap.HttpRoute,
		AWS_AppMesh_GatewayRoute_GatewayRouteSpec__PropertiesMap.Priority,
	}
)

// AWS_AppMesh_GatewayRoute_GatewayRouteSpec is a binding for AWS::AppMesh::GatewayRoute.GatewayRouteSpec.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html
type AWS_AppMesh_GatewayRoute_GatewayRouteSpec struct {
	// GrpcRoute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html#cfn-appmesh-gatewayroute-gatewayroutespec-grpcroute
	GrpcRoute cfz.Expression[AWS_AppMesh_GatewayRoute_GrpcGatewayRoute] `json:"GrpcRoute,omitempty"`

	// Http2Route is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html#cfn-appmesh-gatewayroute-gatewayroutespec-http2route
	Http2Route cfz.Expression[AWS_AppMesh_GatewayRoute_HttpGatewayRoute] `json:"Http2Route,omitempty"`

	// HttpRoute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html#cfn-appmesh-gatewayroute-gatewayroutespec-httproute
	HttpRoute cfz.Expression[AWS_AppMesh_GatewayRoute_HttpGatewayRoute] `json:"HttpRoute,omitempty"`

	// Priority is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html#cfn-appmesh-gatewayroute-gatewayroutespec-priority
	Priority cfz.Expression[int32] `json:"Priority,omitempty"`
}

// New__AWS_AppMesh_GatewayRoute_GatewayRouteSpec initializes a new AWS_AppMesh_GatewayRoute_GatewayRouteSpec.
func New__AWS_AppMesh_GatewayRoute_GatewayRouteSpec() AWS_AppMesh_GatewayRoute_GatewayRouteSpec {
	return AWS_AppMesh_GatewayRoute_GatewayRouteSpec{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_GatewayRoute_GatewayRouteSpec) GetType() string {
	return AWS_AppMesh_GatewayRoute_GatewayRouteSpec__Type
}

// Set__GrpcRoute updates property "GrpcRoute".
func (t AWS_AppMesh_GatewayRoute_GatewayRouteSpec) Set__GrpcRoute(v cfz.Expression[AWS_AppMesh_GatewayRoute_GrpcGatewayRoute]) AWS_AppMesh_GatewayRoute_GatewayRouteSpec {
	t.GrpcRoute = v
	return t
}

// SetV__GrpcRoute updates property "GrpcRoute".
func (t AWS_AppMesh_GatewayRoute_GatewayRouteSpec) SetV__GrpcRoute(v AWS_AppMesh_GatewayRoute_GrpcGatewayRoute) AWS_AppMesh_GatewayRoute_GatewayRouteSpec {
	t.GrpcRoute = cfz.V(v)
	return t
}

// Set__Http2Route updates property "Http2Route".
func (t AWS_AppMesh_GatewayRoute_GatewayRouteSpec) Set__Http2Route(v cfz.Expression[AWS_AppMesh_GatewayRoute_HttpGatewayRoute]) AWS_AppMesh_GatewayRoute_GatewayRouteSpec {
	t.Http2Route = v
	return t
}

// SetV__Http2Route updates property "Http2Route".
func (t AWS_AppMesh_GatewayRoute_GatewayRouteSpec) SetV__Http2Route(v AWS_AppMesh_GatewayRoute_HttpGatewayRoute) AWS_AppMesh_GatewayRoute_GatewayRouteSpec {
	t.Http2Route = cfz.V(v)
	return t
}

// Set__HttpRoute updates property "HttpRoute".
func (t AWS_AppMesh_GatewayRoute_GatewayRouteSpec) Set__HttpRoute(v cfz.Expression[AWS_AppMesh_GatewayRoute_HttpGatewayRoute]) AWS_AppMesh_GatewayRoute_GatewayRouteSpec {
	t.HttpRoute = v
	return t
}

// SetV__HttpRoute updates property "HttpRoute".
func (t AWS_AppMesh_GatewayRoute_GatewayRouteSpec) SetV__HttpRoute(v AWS_AppMesh_GatewayRoute_HttpGatewayRoute) AWS_AppMesh_GatewayRoute_GatewayRouteSpec {
	t.HttpRoute = cfz.V(v)
	return t
}

// Set__Priority updates property "Priority".
func (t AWS_AppMesh_GatewayRoute_GatewayRouteSpec) Set__Priority(v cfz.Expression[int32]) AWS_AppMesh_GatewayRoute_GatewayRouteSpec {
	t.Priority = v
	return t
}

// SetV__Priority updates property "Priority".
func (t AWS_AppMesh_GatewayRoute_GatewayRouteSpec) SetV__Priority(v int32) AWS_AppMesh_GatewayRoute_GatewayRouteSpec {
	t.Priority = cfz.V(v)
	return t
}
