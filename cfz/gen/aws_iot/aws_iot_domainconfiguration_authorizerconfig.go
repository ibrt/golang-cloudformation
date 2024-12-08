// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoT_DomainConfiguration_AuthorizerConfig__Type is the CloudFormation type for AWS::IoT::DomainConfiguration.AuthorizerConfig.
	AWS_IoT_DomainConfiguration_AuthorizerConfig__Type = "AWS::IoT::DomainConfiguration.AuthorizerConfig"
)

var (
	// AWS_IoT_DomainConfiguration_AuthorizerConfig__PropertiesMap reports all the CloudFormation properties for AWS::IoT::DomainConfiguration.AuthorizerConfig.
	AWS_IoT_DomainConfiguration_AuthorizerConfig__PropertiesMap = struct {
		AllowAuthorizerOverride string
		DefaultAuthorizerName   string
	}{
		AllowAuthorizerOverride: "AllowAuthorizerOverride",
		DefaultAuthorizerName:   "DefaultAuthorizerName",
	}

	// AWS_IoT_DomainConfiguration_AuthorizerConfig__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::DomainConfiguration.AuthorizerConfig.
	AWS_IoT_DomainConfiguration_AuthorizerConfig__PropertiesSlice = []string{
		AWS_IoT_DomainConfiguration_AuthorizerConfig__PropertiesMap.AllowAuthorizerOverride,
		AWS_IoT_DomainConfiguration_AuthorizerConfig__PropertiesMap.DefaultAuthorizerName,
	}
)

// AWS_IoT_DomainConfiguration_AuthorizerConfig is a binding for AWS::IoT::DomainConfiguration.AuthorizerConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-authorizerconfig.html
type AWS_IoT_DomainConfiguration_AuthorizerConfig struct {
	// AllowAuthorizerOverride is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-authorizerconfig.html#cfn-iot-domainconfiguration-authorizerconfig-allowauthorizeroverride
	AllowAuthorizerOverride cfz.Expression[bool] `json:"AllowAuthorizerOverride,omitempty"`

	// DefaultAuthorizerName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-authorizerconfig.html#cfn-iot-domainconfiguration-authorizerconfig-defaultauthorizername
	DefaultAuthorizerName cfz.Expression[string] `json:"DefaultAuthorizerName,omitempty"`
}

// New__AWS_IoT_DomainConfiguration_AuthorizerConfig initializes a new AWS_IoT_DomainConfiguration_AuthorizerConfig.
func New__AWS_IoT_DomainConfiguration_AuthorizerConfig() AWS_IoT_DomainConfiguration_AuthorizerConfig {
	return AWS_IoT_DomainConfiguration_AuthorizerConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_IoT_DomainConfiguration_AuthorizerConfig) GetType() string {
	return AWS_IoT_DomainConfiguration_AuthorizerConfig__Type
}

// Set__AllowAuthorizerOverride updates property "AllowAuthorizerOverride".
func (t AWS_IoT_DomainConfiguration_AuthorizerConfig) Set__AllowAuthorizerOverride(v cfz.Expression[bool]) AWS_IoT_DomainConfiguration_AuthorizerConfig {
	t.AllowAuthorizerOverride = v
	return t
}

// SetV__AllowAuthorizerOverride updates property "AllowAuthorizerOverride".
func (t AWS_IoT_DomainConfiguration_AuthorizerConfig) SetV__AllowAuthorizerOverride(v bool) AWS_IoT_DomainConfiguration_AuthorizerConfig {
	t.AllowAuthorizerOverride = cfz.V(v)
	return t
}

// Set__DefaultAuthorizerName updates property "DefaultAuthorizerName".
func (t AWS_IoT_DomainConfiguration_AuthorizerConfig) Set__DefaultAuthorizerName(v cfz.Expression[string]) AWS_IoT_DomainConfiguration_AuthorizerConfig {
	t.DefaultAuthorizerName = v
	return t
}

// SetV__DefaultAuthorizerName updates property "DefaultAuthorizerName".
func (t AWS_IoT_DomainConfiguration_AuthorizerConfig) SetV__DefaultAuthorizerName(v string) AWS_IoT_DomainConfiguration_AuthorizerConfig {
	t.DefaultAuthorizerName = cfz.V(v)
	return t
}
