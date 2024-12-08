// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualNode_TlsValidationContextTrust__Type is the CloudFormation type for AWS::AppMesh::VirtualNode.TlsValidationContextTrust.
	AWS_AppMesh_VirtualNode_TlsValidationContextTrust__Type = "AWS::AppMesh::VirtualNode.TlsValidationContextTrust"
)

var (
	// AWS_AppMesh_VirtualNode_TlsValidationContextTrust__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualNode.TlsValidationContextTrust.
	AWS_AppMesh_VirtualNode_TlsValidationContextTrust__PropertiesMap = struct {
		ACM  string
		File string
		SDS  string
	}{
		ACM:  "ACM",
		File: "File",
		SDS:  "SDS",
	}

	// AWS_AppMesh_VirtualNode_TlsValidationContextTrust__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualNode.TlsValidationContextTrust.
	AWS_AppMesh_VirtualNode_TlsValidationContextTrust__PropertiesSlice = []string{
		AWS_AppMesh_VirtualNode_TlsValidationContextTrust__PropertiesMap.ACM,
		AWS_AppMesh_VirtualNode_TlsValidationContextTrust__PropertiesMap.File,
		AWS_AppMesh_VirtualNode_TlsValidationContextTrust__PropertiesMap.SDS,
	}
)

// AWS_AppMesh_VirtualNode_TlsValidationContextTrust is a binding for AWS::AppMesh::VirtualNode.TlsValidationContextTrust.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontexttrust.html
type AWS_AppMesh_VirtualNode_TlsValidationContextTrust struct {
	// ACM is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontexttrust.html#cfn-appmesh-virtualnode-tlsvalidationcontexttrust-acm
	ACM cfz.Expression[AWS_AppMesh_VirtualNode_TlsValidationContextAcmTrust] `json:"ACM,omitempty"`

	// File is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontexttrust.html#cfn-appmesh-virtualnode-tlsvalidationcontexttrust-file
	File cfz.Expression[AWS_AppMesh_VirtualNode_TlsValidationContextFileTrust] `json:"File,omitempty"`

	// SDS is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-tlsvalidationcontexttrust.html#cfn-appmesh-virtualnode-tlsvalidationcontexttrust-sds
	SDS cfz.Expression[AWS_AppMesh_VirtualNode_TlsValidationContextSdsTrust] `json:"SDS,omitempty"`
}

// New__AWS_AppMesh_VirtualNode_TlsValidationContextTrust initializes a new AWS_AppMesh_VirtualNode_TlsValidationContextTrust.
func New__AWS_AppMesh_VirtualNode_TlsValidationContextTrust() AWS_AppMesh_VirtualNode_TlsValidationContextTrust {
	return AWS_AppMesh_VirtualNode_TlsValidationContextTrust{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualNode_TlsValidationContextTrust) GetType() string {
	return AWS_AppMesh_VirtualNode_TlsValidationContextTrust__Type
}

// Set__ACM updates property "ACM".
func (t AWS_AppMesh_VirtualNode_TlsValidationContextTrust) Set__ACM(v cfz.Expression[AWS_AppMesh_VirtualNode_TlsValidationContextAcmTrust]) AWS_AppMesh_VirtualNode_TlsValidationContextTrust {
	t.ACM = v
	return t
}

// SetV__ACM updates property "ACM".
func (t AWS_AppMesh_VirtualNode_TlsValidationContextTrust) SetV__ACM(v AWS_AppMesh_VirtualNode_TlsValidationContextAcmTrust) AWS_AppMesh_VirtualNode_TlsValidationContextTrust {
	t.ACM = cfz.V(v)
	return t
}

// Set__File updates property "File".
func (t AWS_AppMesh_VirtualNode_TlsValidationContextTrust) Set__File(v cfz.Expression[AWS_AppMesh_VirtualNode_TlsValidationContextFileTrust]) AWS_AppMesh_VirtualNode_TlsValidationContextTrust {
	t.File = v
	return t
}

// SetV__File updates property "File".
func (t AWS_AppMesh_VirtualNode_TlsValidationContextTrust) SetV__File(v AWS_AppMesh_VirtualNode_TlsValidationContextFileTrust) AWS_AppMesh_VirtualNode_TlsValidationContextTrust {
	t.File = cfz.V(v)
	return t
}

// Set__SDS updates property "SDS".
func (t AWS_AppMesh_VirtualNode_TlsValidationContextTrust) Set__SDS(v cfz.Expression[AWS_AppMesh_VirtualNode_TlsValidationContextSdsTrust]) AWS_AppMesh_VirtualNode_TlsValidationContextTrust {
	t.SDS = v
	return t
}

// SetV__SDS updates property "SDS".
func (t AWS_AppMesh_VirtualNode_TlsValidationContextTrust) SetV__SDS(v AWS_AppMesh_VirtualNode_TlsValidationContextSdsTrust) AWS_AppMesh_VirtualNode_TlsValidationContextTrust {
	t.SDS = cfz.V(v)
	return t
}
