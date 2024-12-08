// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate__Type is the CloudFormation type for AWS::AppMesh::VirtualGateway.VirtualGatewayListenerTlsFileCertificate.
	AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate__Type = "AWS::AppMesh::VirtualGateway.VirtualGatewayListenerTlsFileCertificate"
)

var (
	// AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewayListenerTlsFileCertificate.
	AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate__PropertiesMap = struct {
		CertificateChain string
		PrivateKey       string
	}{
		CertificateChain: "CertificateChain",
		PrivateKey:       "PrivateKey",
	}

	// AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewayListenerTlsFileCertificate.
	AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate__PropertiesSlice = []string{
		AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate__PropertiesMap.CertificateChain,
		AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate__PropertiesMap.PrivateKey,
	}
)

// AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate is a binding for AWS::AppMesh::VirtualGateway.VirtualGatewayListenerTlsFileCertificate.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.html
type AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate struct {
	// CertificateChain is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate-certificatechain
	CertificateChain cfz.Expression[string] `json:"CertificateChain,omitempty"`

	// PrivateKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate-privatekey
	PrivateKey cfz.Expression[string] `json:"PrivateKey,omitempty"`
}

// New__AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate initializes a new AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate.
func New__AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate() AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate {
	return AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate) GetType() string {
	return AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate__Type
}

// Set__CertificateChain updates property "CertificateChain".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate) Set__CertificateChain(v cfz.Expression[string]) AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate {
	t.CertificateChain = v
	return t
}

// SetV__CertificateChain updates property "CertificateChain".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate) SetV__CertificateChain(v string) AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate {
	t.CertificateChain = cfz.V(v)
	return t
}

// Set__PrivateKey updates property "PrivateKey".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate) Set__PrivateKey(v cfz.Expression[string]) AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate {
	t.PrivateKey = v
	return t
}

// SetV__PrivateKey updates property "PrivateKey".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate) SetV__PrivateKey(v string) AWS_AppMesh_VirtualGateway_VirtualGatewayListenerTlsFileCertificate {
	t.PrivateKey = cfz.V(v)
	return t
}
