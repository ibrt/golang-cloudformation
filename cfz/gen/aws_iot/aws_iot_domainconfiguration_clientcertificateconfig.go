// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoT_DomainConfiguration_ClientCertificateConfig__Type is the CloudFormation type for AWS::IoT::DomainConfiguration.ClientCertificateConfig.
	AWS_IoT_DomainConfiguration_ClientCertificateConfig__Type = "AWS::IoT::DomainConfiguration.ClientCertificateConfig"
)

var (
	// AWS_IoT_DomainConfiguration_ClientCertificateConfig__PropertiesMap reports all the CloudFormation properties for AWS::IoT::DomainConfiguration.ClientCertificateConfig.
	AWS_IoT_DomainConfiguration_ClientCertificateConfig__PropertiesMap = struct {
		ClientCertificateCallbackArn string
	}{
		ClientCertificateCallbackArn: "ClientCertificateCallbackArn",
	}

	// AWS_IoT_DomainConfiguration_ClientCertificateConfig__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::DomainConfiguration.ClientCertificateConfig.
	AWS_IoT_DomainConfiguration_ClientCertificateConfig__PropertiesSlice = []string{
		AWS_IoT_DomainConfiguration_ClientCertificateConfig__PropertiesMap.ClientCertificateCallbackArn,
	}
)

// AWS_IoT_DomainConfiguration_ClientCertificateConfig is a binding for AWS::IoT::DomainConfiguration.ClientCertificateConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-clientcertificateconfig.html
type AWS_IoT_DomainConfiguration_ClientCertificateConfig struct {
	// ClientCertificateCallbackArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-clientcertificateconfig.html#cfn-iot-domainconfiguration-clientcertificateconfig-clientcertificatecallbackarn
	ClientCertificateCallbackArn cfz.Expression[string] `json:"ClientCertificateCallbackArn,omitempty"`
}

// New__AWS_IoT_DomainConfiguration_ClientCertificateConfig initializes a new AWS_IoT_DomainConfiguration_ClientCertificateConfig.
func New__AWS_IoT_DomainConfiguration_ClientCertificateConfig() AWS_IoT_DomainConfiguration_ClientCertificateConfig {
	return AWS_IoT_DomainConfiguration_ClientCertificateConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_IoT_DomainConfiguration_ClientCertificateConfig) GetType() string {
	return AWS_IoT_DomainConfiguration_ClientCertificateConfig__Type
}

// Set__ClientCertificateCallbackArn updates property "ClientCertificateCallbackArn".
func (t AWS_IoT_DomainConfiguration_ClientCertificateConfig) Set__ClientCertificateCallbackArn(v cfz.Expression[string]) AWS_IoT_DomainConfiguration_ClientCertificateConfig {
	t.ClientCertificateCallbackArn = v
	return t
}

// SetV__ClientCertificateCallbackArn updates property "ClientCertificateCallbackArn".
func (t AWS_IoT_DomainConfiguration_ClientCertificateConfig) SetV__ClientCertificateCallbackArn(v string) AWS_IoT_DomainConfiguration_ClientCertificateConfig {
	t.ClientCertificateCallbackArn = cfz.V(v)
	return t
}
