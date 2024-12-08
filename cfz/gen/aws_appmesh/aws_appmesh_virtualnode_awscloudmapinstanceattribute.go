// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute__Type is the CloudFormation type for AWS::AppMesh::VirtualNode.AwsCloudMapInstanceAttribute.
	AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute__Type = "AWS::AppMesh::VirtualNode.AwsCloudMapInstanceAttribute"
)

var (
	// AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualNode.AwsCloudMapInstanceAttribute.
	AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute__PropertiesMap = struct {
		Key   string
		Value string
	}{
		Key:   "Key",
		Value: "Value",
	}

	// AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualNode.AwsCloudMapInstanceAttribute.
	AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute__PropertiesSlice = []string{
		AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute__PropertiesMap.Key,
		AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute__PropertiesMap.Value,
	}
)

// AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute is a binding for AWS::AppMesh::VirtualNode.AwsCloudMapInstanceAttribute.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapinstanceattribute.html
type AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapinstanceattribute.html#cfn-appmesh-virtualnode-awscloudmapinstanceattribute-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapinstanceattribute.html#cfn-appmesh-virtualnode-awscloudmapinstanceattribute-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute initializes a new AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute.
func New__AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute() AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute {
	return AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute) GetType() string {
	return AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute__Type
}

// Set__Key updates property "Key".
func (t AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute) Set__Key(v cfz.Expression[string]) AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute) SetV__Key(v string) AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute {
	t.Key = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute) Set__Value(v cfz.Expression[string]) AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute) SetV__Value(v string) AWS_AppMesh_VirtualNode_AwsCloudMapInstanceAttribute {
	t.Value = cfz.V(v)
	return t
}
