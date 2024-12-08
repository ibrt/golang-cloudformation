// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualGateway_JsonFormatRef__Type is the CloudFormation type for AWS::AppMesh::VirtualGateway.JsonFormatRef.
	AWS_AppMesh_VirtualGateway_JsonFormatRef__Type = "AWS::AppMesh::VirtualGateway.JsonFormatRef"
)

var (
	// AWS_AppMesh_VirtualGateway_JsonFormatRef__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.JsonFormatRef.
	AWS_AppMesh_VirtualGateway_JsonFormatRef__PropertiesMap = struct {
		Key   string
		Value string
	}{
		Key:   "Key",
		Value: "Value",
	}

	// AWS_AppMesh_VirtualGateway_JsonFormatRef__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.JsonFormatRef.
	AWS_AppMesh_VirtualGateway_JsonFormatRef__PropertiesSlice = []string{
		AWS_AppMesh_VirtualGateway_JsonFormatRef__PropertiesMap.Key,
		AWS_AppMesh_VirtualGateway_JsonFormatRef__PropertiesMap.Value,
	}
)

// AWS_AppMesh_VirtualGateway_JsonFormatRef is a binding for AWS::AppMesh::VirtualGateway.JsonFormatRef.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-jsonformatref.html
type AWS_AppMesh_VirtualGateway_JsonFormatRef struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-jsonformatref.html#cfn-appmesh-virtualgateway-jsonformatref-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-jsonformatref.html#cfn-appmesh-virtualgateway-jsonformatref-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_AppMesh_VirtualGateway_JsonFormatRef initializes a new AWS_AppMesh_VirtualGateway_JsonFormatRef.
func New__AWS_AppMesh_VirtualGateway_JsonFormatRef() AWS_AppMesh_VirtualGateway_JsonFormatRef {
	return AWS_AppMesh_VirtualGateway_JsonFormatRef{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualGateway_JsonFormatRef) GetType() string {
	return AWS_AppMesh_VirtualGateway_JsonFormatRef__Type
}

// Set__Key updates property "Key".
func (t AWS_AppMesh_VirtualGateway_JsonFormatRef) Set__Key(v cfz.Expression[string]) AWS_AppMesh_VirtualGateway_JsonFormatRef {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_AppMesh_VirtualGateway_JsonFormatRef) SetV__Key(v string) AWS_AppMesh_VirtualGateway_JsonFormatRef {
	t.Key = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_AppMesh_VirtualGateway_JsonFormatRef) Set__Value(v cfz.Expression[string]) AWS_AppMesh_VirtualGateway_JsonFormatRef {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_AppMesh_VirtualGateway_JsonFormatRef) SetV__Value(v string) AWS_AppMesh_VirtualGateway_JsonFormatRef {
	t.Value = cfz.V(v)
	return t
}
