// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pinpointemail

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_PinpointEmail_ConfigurationSet_SendingOptions__Type is the CloudFormation type for AWS::PinpointEmail::ConfigurationSet.SendingOptions.
	AWS_PinpointEmail_ConfigurationSet_SendingOptions__Type = "AWS::PinpointEmail::ConfigurationSet.SendingOptions"
)

var (
	// AWS_PinpointEmail_ConfigurationSet_SendingOptions__PropertiesMap reports all the CloudFormation properties for AWS::PinpointEmail::ConfigurationSet.SendingOptions.
	AWS_PinpointEmail_ConfigurationSet_SendingOptions__PropertiesMap = struct {
		SendingEnabled string
	}{
		SendingEnabled: "SendingEnabled",
	}

	// AWS_PinpointEmail_ConfigurationSet_SendingOptions__PropertiesSlice reports all the CloudFormation properties for AWS::PinpointEmail::ConfigurationSet.SendingOptions.
	AWS_PinpointEmail_ConfigurationSet_SendingOptions__PropertiesSlice = []string{
		AWS_PinpointEmail_ConfigurationSet_SendingOptions__PropertiesMap.SendingEnabled,
	}
)

// AWS_PinpointEmail_ConfigurationSet_SendingOptions is a binding for AWS::PinpointEmail::ConfigurationSet.SendingOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-sendingoptions.html
type AWS_PinpointEmail_ConfigurationSet_SendingOptions struct {
	// SendingEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-sendingoptions.html#cfn-pinpointemail-configurationset-sendingoptions-sendingenabled
	SendingEnabled cfz.Expression[bool] `json:"SendingEnabled,omitempty"`
}

// New__AWS_PinpointEmail_ConfigurationSet_SendingOptions initializes a new AWS_PinpointEmail_ConfigurationSet_SendingOptions.
func New__AWS_PinpointEmail_ConfigurationSet_SendingOptions() AWS_PinpointEmail_ConfigurationSet_SendingOptions {
	return AWS_PinpointEmail_ConfigurationSet_SendingOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_PinpointEmail_ConfigurationSet_SendingOptions) GetType() string {
	return AWS_PinpointEmail_ConfigurationSet_SendingOptions__Type
}

// Set__SendingEnabled updates property "SendingEnabled".
func (t AWS_PinpointEmail_ConfigurationSet_SendingOptions) Set__SendingEnabled(v cfz.Expression[bool]) AWS_PinpointEmail_ConfigurationSet_SendingOptions {
	t.SendingEnabled = v
	return t
}

// SetV__SendingEnabled updates property "SendingEnabled".
func (t AWS_PinpointEmail_ConfigurationSet_SendingOptions) SetV__SendingEnabled(v bool) AWS_PinpointEmail_ConfigurationSet_SendingOptions {
	t.SendingEnabled = cfz.V(v)
	return t
}
