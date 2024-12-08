// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_groundstation

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GroundStation_Config_UplinkEchoConfig__Type is the CloudFormation type for AWS::GroundStation::Config.UplinkEchoConfig.
	AWS_GroundStation_Config_UplinkEchoConfig__Type = "AWS::GroundStation::Config.UplinkEchoConfig"
)

var (
	// AWS_GroundStation_Config_UplinkEchoConfig__PropertiesMap reports all the CloudFormation properties for AWS::GroundStation::Config.UplinkEchoConfig.
	AWS_GroundStation_Config_UplinkEchoConfig__PropertiesMap = struct {
		AntennaUplinkConfigArn string
		Enabled                string
	}{
		AntennaUplinkConfigArn: "AntennaUplinkConfigArn",
		Enabled:                "Enabled",
	}

	// AWS_GroundStation_Config_UplinkEchoConfig__PropertiesSlice reports all the CloudFormation properties for AWS::GroundStation::Config.UplinkEchoConfig.
	AWS_GroundStation_Config_UplinkEchoConfig__PropertiesSlice = []string{
		AWS_GroundStation_Config_UplinkEchoConfig__PropertiesMap.AntennaUplinkConfigArn,
		AWS_GroundStation_Config_UplinkEchoConfig__PropertiesMap.Enabled,
	}
)

// AWS_GroundStation_Config_UplinkEchoConfig is a binding for AWS::GroundStation::Config.UplinkEchoConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkechoconfig.html
type AWS_GroundStation_Config_UplinkEchoConfig struct {
	// AntennaUplinkConfigArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkechoconfig.html#cfn-groundstation-config-uplinkechoconfig-antennauplinkconfigarn
	AntennaUplinkConfigArn cfz.Expression[string] `json:"AntennaUplinkConfigArn,omitempty"`

	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkechoconfig.html#cfn-groundstation-config-uplinkechoconfig-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`
}

// New__AWS_GroundStation_Config_UplinkEchoConfig initializes a new AWS_GroundStation_Config_UplinkEchoConfig.
func New__AWS_GroundStation_Config_UplinkEchoConfig() AWS_GroundStation_Config_UplinkEchoConfig {
	return AWS_GroundStation_Config_UplinkEchoConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_GroundStation_Config_UplinkEchoConfig) GetType() string {
	return AWS_GroundStation_Config_UplinkEchoConfig__Type
}

// Set__AntennaUplinkConfigArn updates property "AntennaUplinkConfigArn".
func (t AWS_GroundStation_Config_UplinkEchoConfig) Set__AntennaUplinkConfigArn(v cfz.Expression[string]) AWS_GroundStation_Config_UplinkEchoConfig {
	t.AntennaUplinkConfigArn = v
	return t
}

// SetV__AntennaUplinkConfigArn updates property "AntennaUplinkConfigArn".
func (t AWS_GroundStation_Config_UplinkEchoConfig) SetV__AntennaUplinkConfigArn(v string) AWS_GroundStation_Config_UplinkEchoConfig {
	t.AntennaUplinkConfigArn = cfz.V(v)
	return t
}

// Set__Enabled updates property "Enabled".
func (t AWS_GroundStation_Config_UplinkEchoConfig) Set__Enabled(v cfz.Expression[bool]) AWS_GroundStation_Config_UplinkEchoConfig {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_GroundStation_Config_UplinkEchoConfig) SetV__Enabled(v bool) AWS_GroundStation_Config_UplinkEchoConfig {
	t.Enabled = cfz.V(v)
	return t
}
