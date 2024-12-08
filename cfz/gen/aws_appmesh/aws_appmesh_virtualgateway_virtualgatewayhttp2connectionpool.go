// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool__Type is the CloudFormation type for AWS::AppMesh::VirtualGateway.VirtualGatewayHttp2ConnectionPool.
	AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool__Type = "AWS::AppMesh::VirtualGateway.VirtualGatewayHttp2ConnectionPool"
)

var (
	// AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewayHttp2ConnectionPool.
	AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool__PropertiesMap = struct {
		MaxRequests string
	}{
		MaxRequests: "MaxRequests",
	}

	// AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewayHttp2ConnectionPool.
	AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool__PropertiesSlice = []string{
		AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool__PropertiesMap.MaxRequests,
	}
)

// AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool is a binding for AWS::AppMesh::VirtualGateway.VirtualGatewayHttp2ConnectionPool.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhttp2connectionpool.html
type AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool struct {
	// MaxRequests is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhttp2connectionpool.html#cfn-appmesh-virtualgateway-virtualgatewayhttp2connectionpool-maxrequests
	MaxRequests cfz.Expression[int32] `json:"MaxRequests,omitempty"`
}

// New__AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool initializes a new AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool.
func New__AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool() AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool {
	return AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool) GetType() string {
	return AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool__Type
}

// Set__MaxRequests updates property "MaxRequests".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool) Set__MaxRequests(v cfz.Expression[int32]) AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool {
	t.MaxRequests = v
	return t
}

// SetV__MaxRequests updates property "MaxRequests".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool) SetV__MaxRequests(v int32) AWS_AppMesh_VirtualGateway_VirtualGatewayHttp2ConnectionPool {
	t.MaxRequests = cfz.V(v)
	return t
}
