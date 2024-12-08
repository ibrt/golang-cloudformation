// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest__Type is the CloudFormation type for AWS::EC2::ClientVpnEndpoint.CertificateAuthenticationRequest.
	AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest__Type = "AWS::EC2::ClientVpnEndpoint.CertificateAuthenticationRequest"
)

var (
	// AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest__PropertiesMap reports all the CloudFormation properties for AWS::EC2::ClientVpnEndpoint.CertificateAuthenticationRequest.
	AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest__PropertiesMap = struct {
		ClientRootCertificateChainArn string
	}{
		ClientRootCertificateChainArn: "ClientRootCertificateChainArn",
	}

	// AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::ClientVpnEndpoint.CertificateAuthenticationRequest.
	AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest__PropertiesSlice = []string{
		AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest__PropertiesMap.ClientRootCertificateChainArn,
	}
)

// AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest is a binding for AWS::EC2::ClientVpnEndpoint.CertificateAuthenticationRequest.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-certificateauthenticationrequest.html
type AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest struct {
	// ClientRootCertificateChainArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-certificateauthenticationrequest.html#cfn-ec2-clientvpnendpoint-certificateauthenticationrequest-clientrootcertificatechainarn
	ClientRootCertificateChainArn cfz.Expression[string] `json:"ClientRootCertificateChainArn,omitempty"`
}

// New__AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest initializes a new AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest.
func New__AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest() AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest {
	return AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest) GetType() string {
	return AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest__Type
}

// Set__ClientRootCertificateChainArn updates property "ClientRootCertificateChainArn".
func (t AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest) Set__ClientRootCertificateChainArn(v cfz.Expression[string]) AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest {
	t.ClientRootCertificateChainArn = v
	return t
}

// SetV__ClientRootCertificateChainArn updates property "ClientRootCertificateChainArn".
func (t AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest) SetV__ClientRootCertificateChainArn(v string) AWS_EC2_ClientVpnEndpoint_CertificateAuthenticationRequest {
	t.ClientRootCertificateChainArn = cfz.V(v)
	return t
}
