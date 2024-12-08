// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust__Type is the CloudFormation type for AWS::AppMesh::VirtualGateway.VirtualGatewayTlsValidationContextFileTrust.
	AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust__Type = "AWS::AppMesh::VirtualGateway.VirtualGatewayTlsValidationContextFileTrust"
)

var (
	// AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewayTlsValidationContextFileTrust.
	AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust__PropertiesMap = struct {
		CertificateChain string
	}{
		CertificateChain: "CertificateChain",
	}

	// AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewayTlsValidationContextFileTrust.
	AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust__PropertiesSlice = []string{
		AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust__PropertiesMap.CertificateChain,
	}
)

// AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust is a binding for AWS::AppMesh::VirtualGateway.VirtualGatewayTlsValidationContextFileTrust.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextfiletrust.html
type AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust struct {
	// CertificateChain is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextfiletrust.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontextfiletrust-certificatechain
	CertificateChain cfz.Expression[string] `json:"CertificateChain,omitempty"`
}

// New__AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust initializes a new AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust.
func New__AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust() AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust {
	return AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust) GetType() string {
	return AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust__Type
}

// Set__CertificateChain updates property "CertificateChain".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust) Set__CertificateChain(v cfz.Expression[string]) AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust {
	t.CertificateChain = v
	return t
}

// SetV__CertificateChain updates property "CertificateChain".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust) SetV__CertificateChain(v string) AWS_AppMesh_VirtualGateway_VirtualGatewayTlsValidationContextFileTrust {
	t.CertificateChain = cfz.V(v)
	return t
}
