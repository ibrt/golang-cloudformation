// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualNode_Logging__Type is the CloudFormation type for AWS::AppMesh::VirtualNode.Logging.
	AWS_AppMesh_VirtualNode_Logging__Type = "AWS::AppMesh::VirtualNode.Logging"
)

var (
	// AWS_AppMesh_VirtualNode_Logging__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualNode.Logging.
	AWS_AppMesh_VirtualNode_Logging__PropertiesMap = struct {
		AccessLog string
	}{
		AccessLog: "AccessLog",
	}

	// AWS_AppMesh_VirtualNode_Logging__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualNode.Logging.
	AWS_AppMesh_VirtualNode_Logging__PropertiesSlice = []string{
		AWS_AppMesh_VirtualNode_Logging__PropertiesMap.AccessLog,
	}
)

// AWS_AppMesh_VirtualNode_Logging is a binding for AWS::AppMesh::VirtualNode.Logging.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-logging.html
type AWS_AppMesh_VirtualNode_Logging struct {
	// AccessLog is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-logging.html#cfn-appmesh-virtualnode-logging-accesslog
	AccessLog cfz.Expression[AWS_AppMesh_VirtualNode_AccessLog] `json:"AccessLog,omitempty"`
}

// New__AWS_AppMesh_VirtualNode_Logging initializes a new AWS_AppMesh_VirtualNode_Logging.
func New__AWS_AppMesh_VirtualNode_Logging() AWS_AppMesh_VirtualNode_Logging {
	return AWS_AppMesh_VirtualNode_Logging{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualNode_Logging) GetType() string {
	return AWS_AppMesh_VirtualNode_Logging__Type
}

// Set__AccessLog updates property "AccessLog".
func (t AWS_AppMesh_VirtualNode_Logging) Set__AccessLog(v cfz.Expression[AWS_AppMesh_VirtualNode_AccessLog]) AWS_AppMesh_VirtualNode_Logging {
	t.AccessLog = v
	return t
}

// SetV__AccessLog updates property "AccessLog".
func (t AWS_AppMesh_VirtualNode_Logging) SetV__AccessLog(v AWS_AppMesh_VirtualNode_AccessLog) AWS_AppMesh_VirtualNode_Logging {
	t.AccessLog = cfz.V(v)
	return t
}
