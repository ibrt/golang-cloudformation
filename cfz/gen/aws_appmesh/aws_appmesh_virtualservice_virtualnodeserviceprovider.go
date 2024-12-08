// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualService_VirtualNodeServiceProvider__Type is the CloudFormation type for AWS::AppMesh::VirtualService.VirtualNodeServiceProvider.
	AWS_AppMesh_VirtualService_VirtualNodeServiceProvider__Type = "AWS::AppMesh::VirtualService.VirtualNodeServiceProvider"
)

var (
	// AWS_AppMesh_VirtualService_VirtualNodeServiceProvider__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualService.VirtualNodeServiceProvider.
	AWS_AppMesh_VirtualService_VirtualNodeServiceProvider__PropertiesMap = struct {
		VirtualNodeName string
	}{
		VirtualNodeName: "VirtualNodeName",
	}

	// AWS_AppMesh_VirtualService_VirtualNodeServiceProvider__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualService.VirtualNodeServiceProvider.
	AWS_AppMesh_VirtualService_VirtualNodeServiceProvider__PropertiesSlice = []string{
		AWS_AppMesh_VirtualService_VirtualNodeServiceProvider__PropertiesMap.VirtualNodeName,
	}
)

// AWS_AppMesh_VirtualService_VirtualNodeServiceProvider is a binding for AWS::AppMesh::VirtualService.VirtualNodeServiceProvider.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualnodeserviceprovider.html
type AWS_AppMesh_VirtualService_VirtualNodeServiceProvider struct {
	// VirtualNodeName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualservice-virtualnodeserviceprovider.html#cfn-appmesh-virtualservice-virtualnodeserviceprovider-virtualnodename
	VirtualNodeName cfz.Expression[string] `json:"VirtualNodeName,omitempty"`
}

// New__AWS_AppMesh_VirtualService_VirtualNodeServiceProvider initializes a new AWS_AppMesh_VirtualService_VirtualNodeServiceProvider.
func New__AWS_AppMesh_VirtualService_VirtualNodeServiceProvider() AWS_AppMesh_VirtualService_VirtualNodeServiceProvider {
	return AWS_AppMesh_VirtualService_VirtualNodeServiceProvider{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualService_VirtualNodeServiceProvider) GetType() string {
	return AWS_AppMesh_VirtualService_VirtualNodeServiceProvider__Type
}

// Set__VirtualNodeName updates property "VirtualNodeName".
func (t AWS_AppMesh_VirtualService_VirtualNodeServiceProvider) Set__VirtualNodeName(v cfz.Expression[string]) AWS_AppMesh_VirtualService_VirtualNodeServiceProvider {
	t.VirtualNodeName = v
	return t
}

// SetV__VirtualNodeName updates property "VirtualNodeName".
func (t AWS_AppMesh_VirtualService_VirtualNodeServiceProvider) SetV__VirtualNodeName(v string) AWS_AppMesh_VirtualService_VirtualNodeServiceProvider {
	t.VirtualNodeName = cfz.V(v)
	return t
}
