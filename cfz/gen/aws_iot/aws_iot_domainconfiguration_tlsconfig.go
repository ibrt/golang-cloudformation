// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoT_DomainConfiguration_TlsConfig__Type is the CloudFormation type for AWS::IoT::DomainConfiguration.TlsConfig.
	AWS_IoT_DomainConfiguration_TlsConfig__Type = "AWS::IoT::DomainConfiguration.TlsConfig"
)

var (
	// AWS_IoT_DomainConfiguration_TlsConfig__PropertiesMap reports all the CloudFormation properties for AWS::IoT::DomainConfiguration.TlsConfig.
	AWS_IoT_DomainConfiguration_TlsConfig__PropertiesMap = struct {
		SecurityPolicy string
	}{
		SecurityPolicy: "SecurityPolicy",
	}

	// AWS_IoT_DomainConfiguration_TlsConfig__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::DomainConfiguration.TlsConfig.
	AWS_IoT_DomainConfiguration_TlsConfig__PropertiesSlice = []string{
		AWS_IoT_DomainConfiguration_TlsConfig__PropertiesMap.SecurityPolicy,
	}
)

// AWS_IoT_DomainConfiguration_TlsConfig is a binding for AWS::IoT::DomainConfiguration.TlsConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-tlsconfig.html
type AWS_IoT_DomainConfiguration_TlsConfig struct {
	// SecurityPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-tlsconfig.html#cfn-iot-domainconfiguration-tlsconfig-securitypolicy
	SecurityPolicy cfz.Expression[string] `json:"SecurityPolicy,omitempty"`
}

// New__AWS_IoT_DomainConfiguration_TlsConfig initializes a new AWS_IoT_DomainConfiguration_TlsConfig.
func New__AWS_IoT_DomainConfiguration_TlsConfig() AWS_IoT_DomainConfiguration_TlsConfig {
	return AWS_IoT_DomainConfiguration_TlsConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_IoT_DomainConfiguration_TlsConfig) GetType() string {
	return AWS_IoT_DomainConfiguration_TlsConfig__Type
}

// Set__SecurityPolicy updates property "SecurityPolicy".
func (t AWS_IoT_DomainConfiguration_TlsConfig) Set__SecurityPolicy(v cfz.Expression[string]) AWS_IoT_DomainConfiguration_TlsConfig {
	t.SecurityPolicy = v
	return t
}

// SetV__SecurityPolicy updates property "SecurityPolicy".
func (t AWS_IoT_DomainConfiguration_TlsConfig) SetV__SecurityPolicy(v string) AWS_IoT_DomainConfiguration_TlsConfig {
	t.SecurityPolicy = cfz.V(v)
	return t
}
