// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lex_Bot_ExternalSourceSetting__Type is the CloudFormation type for AWS::Lex::Bot.ExternalSourceSetting.
	AWS_Lex_Bot_ExternalSourceSetting__Type = "AWS::Lex::Bot.ExternalSourceSetting"
)

var (
	// AWS_Lex_Bot_ExternalSourceSetting__PropertiesMap reports all the CloudFormation properties for AWS::Lex::Bot.ExternalSourceSetting.
	AWS_Lex_Bot_ExternalSourceSetting__PropertiesMap = struct {
		GrammarSlotTypeSetting string
	}{
		GrammarSlotTypeSetting: "GrammarSlotTypeSetting",
	}

	// AWS_Lex_Bot_ExternalSourceSetting__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::Bot.ExternalSourceSetting.
	AWS_Lex_Bot_ExternalSourceSetting__PropertiesSlice = []string{
		AWS_Lex_Bot_ExternalSourceSetting__PropertiesMap.GrammarSlotTypeSetting,
	}
)

// AWS_Lex_Bot_ExternalSourceSetting is a binding for AWS::Lex::Bot.ExternalSourceSetting.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-externalsourcesetting.html
type AWS_Lex_Bot_ExternalSourceSetting struct {
	// GrammarSlotTypeSetting is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-externalsourcesetting.html#cfn-lex-bot-externalsourcesetting-grammarslottypesetting
	GrammarSlotTypeSetting cfz.Expression[AWS_Lex_Bot_GrammarSlotTypeSetting] `json:"GrammarSlotTypeSetting,omitempty"`
}

// New__AWS_Lex_Bot_ExternalSourceSetting initializes a new AWS_Lex_Bot_ExternalSourceSetting.
func New__AWS_Lex_Bot_ExternalSourceSetting() AWS_Lex_Bot_ExternalSourceSetting {
	return AWS_Lex_Bot_ExternalSourceSetting{}
}

// GetType returns the CloudFormation type.
func (AWS_Lex_Bot_ExternalSourceSetting) GetType() string {
	return AWS_Lex_Bot_ExternalSourceSetting__Type
}

// Set__GrammarSlotTypeSetting updates property "GrammarSlotTypeSetting".
func (t AWS_Lex_Bot_ExternalSourceSetting) Set__GrammarSlotTypeSetting(v cfz.Expression[AWS_Lex_Bot_GrammarSlotTypeSetting]) AWS_Lex_Bot_ExternalSourceSetting {
	t.GrammarSlotTypeSetting = v
	return t
}

// SetV__GrammarSlotTypeSetting updates property "GrammarSlotTypeSetting".
func (t AWS_Lex_Bot_ExternalSourceSetting) SetV__GrammarSlotTypeSetting(v AWS_Lex_Bot_GrammarSlotTypeSetting) AWS_Lex_Bot_ExternalSourceSetting {
	t.GrammarSlotTypeSetting = cfz.V(v)
	return t
}
