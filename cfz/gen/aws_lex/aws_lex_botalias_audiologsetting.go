// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lex_BotAlias_AudioLogSetting__Type is the CloudFormation type for AWS::Lex::BotAlias.AudioLogSetting.
	AWS_Lex_BotAlias_AudioLogSetting__Type = "AWS::Lex::BotAlias.AudioLogSetting"
)

var (
	// AWS_Lex_BotAlias_AudioLogSetting__PropertiesMap reports all the CloudFormation properties for AWS::Lex::BotAlias.AudioLogSetting.
	AWS_Lex_BotAlias_AudioLogSetting__PropertiesMap = struct {
		Destination string
		Enabled     string
	}{
		Destination: "Destination",
		Enabled:     "Enabled",
	}

	// AWS_Lex_BotAlias_AudioLogSetting__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::BotAlias.AudioLogSetting.
	AWS_Lex_BotAlias_AudioLogSetting__PropertiesSlice = []string{
		AWS_Lex_BotAlias_AudioLogSetting__PropertiesMap.Destination,
		AWS_Lex_BotAlias_AudioLogSetting__PropertiesMap.Enabled,
	}
)

// AWS_Lex_BotAlias_AudioLogSetting is a binding for AWS::Lex::BotAlias.AudioLogSetting.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-audiologsetting.html
type AWS_Lex_BotAlias_AudioLogSetting struct {
	// Destination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-audiologsetting.html#cfn-lex-botalias-audiologsetting-destination
	Destination cfz.Expression[AWS_Lex_BotAlias_AudioLogDestination] `json:"Destination,omitempty"`

	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-audiologsetting.html#cfn-lex-botalias-audiologsetting-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`
}

// New__AWS_Lex_BotAlias_AudioLogSetting initializes a new AWS_Lex_BotAlias_AudioLogSetting.
func New__AWS_Lex_BotAlias_AudioLogSetting() AWS_Lex_BotAlias_AudioLogSetting {
	return AWS_Lex_BotAlias_AudioLogSetting{}
}

// GetType returns the CloudFormation type.
func (AWS_Lex_BotAlias_AudioLogSetting) GetType() string {
	return AWS_Lex_BotAlias_AudioLogSetting__Type
}

// Set__Destination updates property "Destination".
func (t AWS_Lex_BotAlias_AudioLogSetting) Set__Destination(v cfz.Expression[AWS_Lex_BotAlias_AudioLogDestination]) AWS_Lex_BotAlias_AudioLogSetting {
	t.Destination = v
	return t
}

// SetV__Destination updates property "Destination".
func (t AWS_Lex_BotAlias_AudioLogSetting) SetV__Destination(v AWS_Lex_BotAlias_AudioLogDestination) AWS_Lex_BotAlias_AudioLogSetting {
	t.Destination = cfz.V(v)
	return t
}

// Set__Enabled updates property "Enabled".
func (t AWS_Lex_BotAlias_AudioLogSetting) Set__Enabled(v cfz.Expression[bool]) AWS_Lex_BotAlias_AudioLogSetting {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_Lex_BotAlias_AudioLogSetting) SetV__Enabled(v bool) AWS_Lex_BotAlias_AudioLogSetting {
	t.Enabled = cfz.V(v)
	return t
}
