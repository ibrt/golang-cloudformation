// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emr

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EMR_InstanceGroupConfig_Configuration__Type is the CloudFormation type for AWS::EMR::InstanceGroupConfig.Configuration.
	AWS_EMR_InstanceGroupConfig_Configuration__Type = "AWS::EMR::InstanceGroupConfig.Configuration"
)

var (
	// AWS_EMR_InstanceGroupConfig_Configuration__PropertiesMap reports all the CloudFormation properties for AWS::EMR::InstanceGroupConfig.Configuration.
	AWS_EMR_InstanceGroupConfig_Configuration__PropertiesMap = struct {
		Classification          string
		ConfigurationProperties string
		Configurations          string
	}{
		Classification:          "Classification",
		ConfigurationProperties: "ConfigurationProperties",
		Configurations:          "Configurations",
	}

	// AWS_EMR_InstanceGroupConfig_Configuration__PropertiesSlice reports all the CloudFormation properties for AWS::EMR::InstanceGroupConfig.Configuration.
	AWS_EMR_InstanceGroupConfig_Configuration__PropertiesSlice = []string{
		AWS_EMR_InstanceGroupConfig_Configuration__PropertiesMap.Classification,
		AWS_EMR_InstanceGroupConfig_Configuration__PropertiesMap.ConfigurationProperties,
		AWS_EMR_InstanceGroupConfig_Configuration__PropertiesMap.Configurations,
	}
)

// AWS_EMR_InstanceGroupConfig_Configuration is a binding for AWS::EMR::InstanceGroupConfig.Configuration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html
type AWS_EMR_InstanceGroupConfig_Configuration struct {
	// Classification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-classification
	Classification cfz.Expression[string] `json:"Classification,omitempty"`

	// ConfigurationProperties is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurationproperties
	ConfigurationProperties cfz.ExpressionMap[string] `json:"ConfigurationProperties,omitempty"`

	// Configurations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurations
	Configurations cfz.ExpressionSlice[AWS_EMR_InstanceGroupConfig_Configuration] `json:"Configurations,omitempty"`
}

// New__AWS_EMR_InstanceGroupConfig_Configuration initializes a new AWS_EMR_InstanceGroupConfig_Configuration.
func New__AWS_EMR_InstanceGroupConfig_Configuration() AWS_EMR_InstanceGroupConfig_Configuration {
	return AWS_EMR_InstanceGroupConfig_Configuration{}
}

// GetType returns the CloudFormation type.
func (AWS_EMR_InstanceGroupConfig_Configuration) GetType() string {
	return AWS_EMR_InstanceGroupConfig_Configuration__Type
}

// Set__Classification updates property "Classification".
func (t AWS_EMR_InstanceGroupConfig_Configuration) Set__Classification(v cfz.Expression[string]) AWS_EMR_InstanceGroupConfig_Configuration {
	t.Classification = v
	return t
}

// SetV__Classification updates property "Classification".
func (t AWS_EMR_InstanceGroupConfig_Configuration) SetV__Classification(v string) AWS_EMR_InstanceGroupConfig_Configuration {
	t.Classification = cfz.V(v)
	return t
}

// Set__ConfigurationProperties updates property "ConfigurationProperties".
func (t AWS_EMR_InstanceGroupConfig_Configuration) Set__ConfigurationProperties(v cfz.ExpressionMap[string]) AWS_EMR_InstanceGroupConfig_Configuration {
	t.ConfigurationProperties = v
	return t
}

// SetM__ConfigurationProperties updates property "ConfigurationProperties".
func (t AWS_EMR_InstanceGroupConfig_Configuration) SetM__ConfigurationProperties(v ...map[string]cfz.Expression[string]) AWS_EMR_InstanceGroupConfig_Configuration {
	t.ConfigurationProperties = cfz.M(v...)
	return t
}

// SetMV__ConfigurationProperties updates property "ConfigurationProperties".
func (t AWS_EMR_InstanceGroupConfig_Configuration) SetMV__ConfigurationProperties(v ...map[string]string) AWS_EMR_InstanceGroupConfig_Configuration {
	t.ConfigurationProperties = cfz.MV(v...)
	return t
}

// Set__Configurations updates property "Configurations".
func (t AWS_EMR_InstanceGroupConfig_Configuration) Set__Configurations(v cfz.ExpressionSlice[AWS_EMR_InstanceGroupConfig_Configuration]) AWS_EMR_InstanceGroupConfig_Configuration {
	t.Configurations = v
	return t
}

// SetS__Configurations updates property "Configurations".
func (t AWS_EMR_InstanceGroupConfig_Configuration) SetS__Configurations(v ...cfz.Expression[AWS_EMR_InstanceGroupConfig_Configuration]) AWS_EMR_InstanceGroupConfig_Configuration {
	t.Configurations = cfz.S(v...)
	return t
}

// SetSV__Configurations updates property "Configurations".
func (t AWS_EMR_InstanceGroupConfig_Configuration) SetSV__Configurations(v ...AWS_EMR_InstanceGroupConfig_Configuration) AWS_EMR_InstanceGroupConfig_Configuration {
	t.Configurations = cfz.SV(v...)
	return t
}
