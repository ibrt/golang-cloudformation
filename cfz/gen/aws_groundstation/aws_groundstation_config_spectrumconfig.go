// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_groundstation

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GroundStation_Config_SpectrumConfig__Type is the CloudFormation type for AWS::GroundStation::Config.SpectrumConfig.
	AWS_GroundStation_Config_SpectrumConfig__Type = "AWS::GroundStation::Config.SpectrumConfig"
)

var (
	// AWS_GroundStation_Config_SpectrumConfig__PropertiesMap reports all the CloudFormation properties for AWS::GroundStation::Config.SpectrumConfig.
	AWS_GroundStation_Config_SpectrumConfig__PropertiesMap = struct {
		Bandwidth       string
		CenterFrequency string
		Polarization    string
	}{
		Bandwidth:       "Bandwidth",
		CenterFrequency: "CenterFrequency",
		Polarization:    "Polarization",
	}

	// AWS_GroundStation_Config_SpectrumConfig__PropertiesSlice reports all the CloudFormation properties for AWS::GroundStation::Config.SpectrumConfig.
	AWS_GroundStation_Config_SpectrumConfig__PropertiesSlice = []string{
		AWS_GroundStation_Config_SpectrumConfig__PropertiesMap.Bandwidth,
		AWS_GroundStation_Config_SpectrumConfig__PropertiesMap.CenterFrequency,
		AWS_GroundStation_Config_SpectrumConfig__PropertiesMap.Polarization,
	}
)

// AWS_GroundStation_Config_SpectrumConfig is a binding for AWS::GroundStation::Config.SpectrumConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-spectrumconfig.html
type AWS_GroundStation_Config_SpectrumConfig struct {
	// Bandwidth is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-spectrumconfig.html#cfn-groundstation-config-spectrumconfig-bandwidth
	Bandwidth cfz.Expression[AWS_GroundStation_Config_FrequencyBandwidth] `json:"Bandwidth,omitempty"`

	// CenterFrequency is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-spectrumconfig.html#cfn-groundstation-config-spectrumconfig-centerfrequency
	CenterFrequency cfz.Expression[AWS_GroundStation_Config_Frequency] `json:"CenterFrequency,omitempty"`

	// Polarization is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-spectrumconfig.html#cfn-groundstation-config-spectrumconfig-polarization
	Polarization cfz.Expression[string] `json:"Polarization,omitempty"`
}

// New__AWS_GroundStation_Config_SpectrumConfig initializes a new AWS_GroundStation_Config_SpectrumConfig.
func New__AWS_GroundStation_Config_SpectrumConfig() AWS_GroundStation_Config_SpectrumConfig {
	return AWS_GroundStation_Config_SpectrumConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_GroundStation_Config_SpectrumConfig) GetType() string {
	return AWS_GroundStation_Config_SpectrumConfig__Type
}

// Set__Bandwidth updates property "Bandwidth".
func (t AWS_GroundStation_Config_SpectrumConfig) Set__Bandwidth(v cfz.Expression[AWS_GroundStation_Config_FrequencyBandwidth]) AWS_GroundStation_Config_SpectrumConfig {
	t.Bandwidth = v
	return t
}

// SetV__Bandwidth updates property "Bandwidth".
func (t AWS_GroundStation_Config_SpectrumConfig) SetV__Bandwidth(v AWS_GroundStation_Config_FrequencyBandwidth) AWS_GroundStation_Config_SpectrumConfig {
	t.Bandwidth = cfz.V(v)
	return t
}

// Set__CenterFrequency updates property "CenterFrequency".
func (t AWS_GroundStation_Config_SpectrumConfig) Set__CenterFrequency(v cfz.Expression[AWS_GroundStation_Config_Frequency]) AWS_GroundStation_Config_SpectrumConfig {
	t.CenterFrequency = v
	return t
}

// SetV__CenterFrequency updates property "CenterFrequency".
func (t AWS_GroundStation_Config_SpectrumConfig) SetV__CenterFrequency(v AWS_GroundStation_Config_Frequency) AWS_GroundStation_Config_SpectrumConfig {
	t.CenterFrequency = cfz.V(v)
	return t
}

// Set__Polarization updates property "Polarization".
func (t AWS_GroundStation_Config_SpectrumConfig) Set__Polarization(v cfz.Expression[string]) AWS_GroundStation_Config_SpectrumConfig {
	t.Polarization = v
	return t
}

// SetV__Polarization updates property "Polarization".
func (t AWS_GroundStation_Config_SpectrumConfig) SetV__Polarization(v string) AWS_GroundStation_Config_SpectrumConfig {
	t.Polarization = cfz.V(v)
	return t
}
