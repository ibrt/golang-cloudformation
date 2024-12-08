// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_config

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Config_ConfigurationRecorder_RecordingModeOverride__Type is the CloudFormation type for AWS::Config::ConfigurationRecorder.RecordingModeOverride.
	AWS_Config_ConfigurationRecorder_RecordingModeOverride__Type = "AWS::Config::ConfigurationRecorder.RecordingModeOverride"
)

var (
	// AWS_Config_ConfigurationRecorder_RecordingModeOverride__PropertiesMap reports all the CloudFormation properties for AWS::Config::ConfigurationRecorder.RecordingModeOverride.
	AWS_Config_ConfigurationRecorder_RecordingModeOverride__PropertiesMap = struct {
		Description        string
		RecordingFrequency string
		ResourceTypes      string
	}{
		Description:        "Description",
		RecordingFrequency: "RecordingFrequency",
		ResourceTypes:      "ResourceTypes",
	}

	// AWS_Config_ConfigurationRecorder_RecordingModeOverride__PropertiesSlice reports all the CloudFormation properties for AWS::Config::ConfigurationRecorder.RecordingModeOverride.
	AWS_Config_ConfigurationRecorder_RecordingModeOverride__PropertiesSlice = []string{
		AWS_Config_ConfigurationRecorder_RecordingModeOverride__PropertiesMap.Description,
		AWS_Config_ConfigurationRecorder_RecordingModeOverride__PropertiesMap.RecordingFrequency,
		AWS_Config_ConfigurationRecorder_RecordingModeOverride__PropertiesMap.ResourceTypes,
	}
)

// AWS_Config_ConfigurationRecorder_RecordingModeOverride is a binding for AWS::Config::ConfigurationRecorder.RecordingModeOverride.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmodeoverride.html
type AWS_Config_ConfigurationRecorder_RecordingModeOverride struct {
	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmodeoverride.html#cfn-config-configurationrecorder-recordingmodeoverride-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// RecordingFrequency is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmodeoverride.html#cfn-config-configurationrecorder-recordingmodeoverride-recordingfrequency
	RecordingFrequency cfz.Expression[string] `json:"RecordingFrequency,omitempty"`

	// ResourceTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmodeoverride.html#cfn-config-configurationrecorder-recordingmodeoverride-resourcetypes
	ResourceTypes cfz.ExpressionSlice[string] `json:"ResourceTypes,omitempty"`
}

// New__AWS_Config_ConfigurationRecorder_RecordingModeOverride initializes a new AWS_Config_ConfigurationRecorder_RecordingModeOverride.
func New__AWS_Config_ConfigurationRecorder_RecordingModeOverride() AWS_Config_ConfigurationRecorder_RecordingModeOverride {
	return AWS_Config_ConfigurationRecorder_RecordingModeOverride{}
}

// GetType returns the CloudFormation type.
func (AWS_Config_ConfigurationRecorder_RecordingModeOverride) GetType() string {
	return AWS_Config_ConfigurationRecorder_RecordingModeOverride__Type
}

// Set__Description updates property "Description".
func (t AWS_Config_ConfigurationRecorder_RecordingModeOverride) Set__Description(v cfz.Expression[string]) AWS_Config_ConfigurationRecorder_RecordingModeOverride {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_Config_ConfigurationRecorder_RecordingModeOverride) SetV__Description(v string) AWS_Config_ConfigurationRecorder_RecordingModeOverride {
	t.Description = cfz.V(v)
	return t
}

// Set__RecordingFrequency updates property "RecordingFrequency".
func (t AWS_Config_ConfigurationRecorder_RecordingModeOverride) Set__RecordingFrequency(v cfz.Expression[string]) AWS_Config_ConfigurationRecorder_RecordingModeOverride {
	t.RecordingFrequency = v
	return t
}

// SetV__RecordingFrequency updates property "RecordingFrequency".
func (t AWS_Config_ConfigurationRecorder_RecordingModeOverride) SetV__RecordingFrequency(v string) AWS_Config_ConfigurationRecorder_RecordingModeOverride {
	t.RecordingFrequency = cfz.V(v)
	return t
}

// Set__ResourceTypes updates property "ResourceTypes".
func (t AWS_Config_ConfigurationRecorder_RecordingModeOverride) Set__ResourceTypes(v cfz.ExpressionSlice[string]) AWS_Config_ConfigurationRecorder_RecordingModeOverride {
	t.ResourceTypes = v
	return t
}

// SetS__ResourceTypes updates property "ResourceTypes".
func (t AWS_Config_ConfigurationRecorder_RecordingModeOverride) SetS__ResourceTypes(v ...cfz.Expression[string]) AWS_Config_ConfigurationRecorder_RecordingModeOverride {
	t.ResourceTypes = cfz.S(v...)
	return t
}

// SetSV__ResourceTypes updates property "ResourceTypes".
func (t AWS_Config_ConfigurationRecorder_RecordingModeOverride) SetSV__ResourceTypes(v ...string) AWS_Config_ConfigurationRecorder_RecordingModeOverride {
	t.ResourceTypes = cfz.SV(v...)
	return t
}
