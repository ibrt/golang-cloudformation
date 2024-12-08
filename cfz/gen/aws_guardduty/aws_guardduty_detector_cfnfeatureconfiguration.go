// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_guardduty

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GuardDuty_Detector_CFNFeatureConfiguration__Type is the CloudFormation type for AWS::GuardDuty::Detector.CFNFeatureConfiguration.
	AWS_GuardDuty_Detector_CFNFeatureConfiguration__Type = "AWS::GuardDuty::Detector.CFNFeatureConfiguration"
)

var (
	// AWS_GuardDuty_Detector_CFNFeatureConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::GuardDuty::Detector.CFNFeatureConfiguration.
	AWS_GuardDuty_Detector_CFNFeatureConfiguration__PropertiesMap = struct {
		AdditionalConfiguration string
		Name                    string
		Status                  string
	}{
		AdditionalConfiguration: "AdditionalConfiguration",
		Name:                    "Name",
		Status:                  "Status",
	}

	// AWS_GuardDuty_Detector_CFNFeatureConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::GuardDuty::Detector.CFNFeatureConfiguration.
	AWS_GuardDuty_Detector_CFNFeatureConfiguration__PropertiesSlice = []string{
		AWS_GuardDuty_Detector_CFNFeatureConfiguration__PropertiesMap.AdditionalConfiguration,
		AWS_GuardDuty_Detector_CFNFeatureConfiguration__PropertiesMap.Name,
		AWS_GuardDuty_Detector_CFNFeatureConfiguration__PropertiesMap.Status,
	}
)

// AWS_GuardDuty_Detector_CFNFeatureConfiguration is a binding for AWS::GuardDuty::Detector.CFNFeatureConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html
type AWS_GuardDuty_Detector_CFNFeatureConfiguration struct {
	// AdditionalConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html#cfn-guardduty-detector-cfnfeatureconfiguration-additionalconfiguration
	AdditionalConfiguration cfz.ExpressionSlice[AWS_GuardDuty_Detector_CFNFeatureAdditionalConfiguration] `json:"AdditionalConfiguration,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html#cfn-guardduty-detector-cfnfeatureconfiguration-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Status is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnfeatureconfiguration.html#cfn-guardduty-detector-cfnfeatureconfiguration-status
	Status cfz.Expression[string] `json:"Status,omitempty"`
}

// New__AWS_GuardDuty_Detector_CFNFeatureConfiguration initializes a new AWS_GuardDuty_Detector_CFNFeatureConfiguration.
func New__AWS_GuardDuty_Detector_CFNFeatureConfiguration() AWS_GuardDuty_Detector_CFNFeatureConfiguration {
	return AWS_GuardDuty_Detector_CFNFeatureConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_GuardDuty_Detector_CFNFeatureConfiguration) GetType() string {
	return AWS_GuardDuty_Detector_CFNFeatureConfiguration__Type
}

// Set__AdditionalConfiguration updates property "AdditionalConfiguration".
func (t AWS_GuardDuty_Detector_CFNFeatureConfiguration) Set__AdditionalConfiguration(v cfz.ExpressionSlice[AWS_GuardDuty_Detector_CFNFeatureAdditionalConfiguration]) AWS_GuardDuty_Detector_CFNFeatureConfiguration {
	t.AdditionalConfiguration = v
	return t
}

// SetS__AdditionalConfiguration updates property "AdditionalConfiguration".
func (t AWS_GuardDuty_Detector_CFNFeatureConfiguration) SetS__AdditionalConfiguration(v ...cfz.Expression[AWS_GuardDuty_Detector_CFNFeatureAdditionalConfiguration]) AWS_GuardDuty_Detector_CFNFeatureConfiguration {
	t.AdditionalConfiguration = cfz.S(v...)
	return t
}

// SetSV__AdditionalConfiguration updates property "AdditionalConfiguration".
func (t AWS_GuardDuty_Detector_CFNFeatureConfiguration) SetSV__AdditionalConfiguration(v ...AWS_GuardDuty_Detector_CFNFeatureAdditionalConfiguration) AWS_GuardDuty_Detector_CFNFeatureConfiguration {
	t.AdditionalConfiguration = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t AWS_GuardDuty_Detector_CFNFeatureConfiguration) Set__Name(v cfz.Expression[string]) AWS_GuardDuty_Detector_CFNFeatureConfiguration {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_GuardDuty_Detector_CFNFeatureConfiguration) SetV__Name(v string) AWS_GuardDuty_Detector_CFNFeatureConfiguration {
	t.Name = cfz.V(v)
	return t
}

// Set__Status updates property "Status".
func (t AWS_GuardDuty_Detector_CFNFeatureConfiguration) Set__Status(v cfz.Expression[string]) AWS_GuardDuty_Detector_CFNFeatureConfiguration {
	t.Status = v
	return t
}

// SetV__Status updates property "Status".
func (t AWS_GuardDuty_Detector_CFNFeatureConfiguration) SetV__Status(v string) AWS_GuardDuty_Detector_CFNFeatureConfiguration {
	t.Status = cfz.V(v)
	return t
}
