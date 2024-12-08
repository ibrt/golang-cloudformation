// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_groundstation

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GroundStation_Config_UplinkSpectrumConfig__Type is the CloudFormation type for AWS::GroundStation::Config.UplinkSpectrumConfig.
	AWS_GroundStation_Config_UplinkSpectrumConfig__Type = "AWS::GroundStation::Config.UplinkSpectrumConfig"
)

var (
	// AWS_GroundStation_Config_UplinkSpectrumConfig__PropertiesMap reports all the CloudFormation properties for AWS::GroundStation::Config.UplinkSpectrumConfig.
	AWS_GroundStation_Config_UplinkSpectrumConfig__PropertiesMap = struct {
		CenterFrequency string
		Polarization    string
	}{
		CenterFrequency: "CenterFrequency",
		Polarization:    "Polarization",
	}

	// AWS_GroundStation_Config_UplinkSpectrumConfig__PropertiesSlice reports all the CloudFormation properties for AWS::GroundStation::Config.UplinkSpectrumConfig.
	AWS_GroundStation_Config_UplinkSpectrumConfig__PropertiesSlice = []string{
		AWS_GroundStation_Config_UplinkSpectrumConfig__PropertiesMap.CenterFrequency,
		AWS_GroundStation_Config_UplinkSpectrumConfig__PropertiesMap.Polarization,
	}
)

// AWS_GroundStation_Config_UplinkSpectrumConfig is a binding for AWS::GroundStation::Config.UplinkSpectrumConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkspectrumconfig.html
type AWS_GroundStation_Config_UplinkSpectrumConfig struct {
	// CenterFrequency is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkspectrumconfig.html#cfn-groundstation-config-uplinkspectrumconfig-centerfrequency
	CenterFrequency cfz.Expression[AWS_GroundStation_Config_Frequency] `json:"CenterFrequency,omitempty"`

	// Polarization is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkspectrumconfig.html#cfn-groundstation-config-uplinkspectrumconfig-polarization
	Polarization cfz.Expression[string] `json:"Polarization,omitempty"`
}

// New__AWS_GroundStation_Config_UplinkSpectrumConfig initializes a new AWS_GroundStation_Config_UplinkSpectrumConfig.
func New__AWS_GroundStation_Config_UplinkSpectrumConfig() AWS_GroundStation_Config_UplinkSpectrumConfig {
	return AWS_GroundStation_Config_UplinkSpectrumConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_GroundStation_Config_UplinkSpectrumConfig) GetType() string {
	return AWS_GroundStation_Config_UplinkSpectrumConfig__Type
}

// Set__CenterFrequency updates property "CenterFrequency".
func (t AWS_GroundStation_Config_UplinkSpectrumConfig) Set__CenterFrequency(v cfz.Expression[AWS_GroundStation_Config_Frequency]) AWS_GroundStation_Config_UplinkSpectrumConfig {
	t.CenterFrequency = v
	return t
}

// SetV__CenterFrequency updates property "CenterFrequency".
func (t AWS_GroundStation_Config_UplinkSpectrumConfig) SetV__CenterFrequency(v AWS_GroundStation_Config_Frequency) AWS_GroundStation_Config_UplinkSpectrumConfig {
	t.CenterFrequency = cfz.V(v)
	return t
}

// Set__Polarization updates property "Polarization".
func (t AWS_GroundStation_Config_UplinkSpectrumConfig) Set__Polarization(v cfz.Expression[string]) AWS_GroundStation_Config_UplinkSpectrumConfig {
	t.Polarization = v
	return t
}

// SetV__Polarization updates property "Polarization".
func (t AWS_GroundStation_Config_UplinkSpectrumConfig) SetV__Polarization(v string) AWS_GroundStation_Config_UplinkSpectrumConfig {
	t.Polarization = cfz.V(v)
	return t
}
