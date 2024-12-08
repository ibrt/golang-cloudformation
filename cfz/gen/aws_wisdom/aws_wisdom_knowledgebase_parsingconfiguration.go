// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wisdom

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Wisdom_KnowledgeBase_ParsingConfiguration__Type is the CloudFormation type for AWS::Wisdom::KnowledgeBase.ParsingConfiguration.
	AWS_Wisdom_KnowledgeBase_ParsingConfiguration__Type = "AWS::Wisdom::KnowledgeBase.ParsingConfiguration"
)

var (
	// AWS_Wisdom_KnowledgeBase_ParsingConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Wisdom::KnowledgeBase.ParsingConfiguration.
	AWS_Wisdom_KnowledgeBase_ParsingConfiguration__PropertiesMap = struct {
		BedrockFoundationModelConfiguration string
		ParsingStrategy                     string
	}{
		BedrockFoundationModelConfiguration: "BedrockFoundationModelConfiguration",
		ParsingStrategy:                     "ParsingStrategy",
	}

	// AWS_Wisdom_KnowledgeBase_ParsingConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Wisdom::KnowledgeBase.ParsingConfiguration.
	AWS_Wisdom_KnowledgeBase_ParsingConfiguration__PropertiesSlice = []string{
		AWS_Wisdom_KnowledgeBase_ParsingConfiguration__PropertiesMap.BedrockFoundationModelConfiguration,
		AWS_Wisdom_KnowledgeBase_ParsingConfiguration__PropertiesMap.ParsingStrategy,
	}
)

// AWS_Wisdom_KnowledgeBase_ParsingConfiguration is a binding for AWS::Wisdom::KnowledgeBase.ParsingConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingconfiguration.html
type AWS_Wisdom_KnowledgeBase_ParsingConfiguration struct {
	// BedrockFoundationModelConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingconfiguration.html#cfn-wisdom-knowledgebase-parsingconfiguration-bedrockfoundationmodelconfiguration
	BedrockFoundationModelConfiguration cfz.Expression[AWS_Wisdom_KnowledgeBase_BedrockFoundationModelConfiguration] `json:"BedrockFoundationModelConfiguration,omitempty"`

	// ParsingStrategy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingconfiguration.html#cfn-wisdom-knowledgebase-parsingconfiguration-parsingstrategy
	ParsingStrategy cfz.Expression[string] `json:"ParsingStrategy,omitempty"`
}

// New__AWS_Wisdom_KnowledgeBase_ParsingConfiguration initializes a new AWS_Wisdom_KnowledgeBase_ParsingConfiguration.
func New__AWS_Wisdom_KnowledgeBase_ParsingConfiguration() AWS_Wisdom_KnowledgeBase_ParsingConfiguration {
	return AWS_Wisdom_KnowledgeBase_ParsingConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Wisdom_KnowledgeBase_ParsingConfiguration) GetType() string {
	return AWS_Wisdom_KnowledgeBase_ParsingConfiguration__Type
}

// Set__BedrockFoundationModelConfiguration updates property "BedrockFoundationModelConfiguration".
func (t AWS_Wisdom_KnowledgeBase_ParsingConfiguration) Set__BedrockFoundationModelConfiguration(v cfz.Expression[AWS_Wisdom_KnowledgeBase_BedrockFoundationModelConfiguration]) AWS_Wisdom_KnowledgeBase_ParsingConfiguration {
	t.BedrockFoundationModelConfiguration = v
	return t
}

// SetV__BedrockFoundationModelConfiguration updates property "BedrockFoundationModelConfiguration".
func (t AWS_Wisdom_KnowledgeBase_ParsingConfiguration) SetV__BedrockFoundationModelConfiguration(v AWS_Wisdom_KnowledgeBase_BedrockFoundationModelConfiguration) AWS_Wisdom_KnowledgeBase_ParsingConfiguration {
	t.BedrockFoundationModelConfiguration = cfz.V(v)
	return t
}

// Set__ParsingStrategy updates property "ParsingStrategy".
func (t AWS_Wisdom_KnowledgeBase_ParsingConfiguration) Set__ParsingStrategy(v cfz.Expression[string]) AWS_Wisdom_KnowledgeBase_ParsingConfiguration {
	t.ParsingStrategy = v
	return t
}

// SetV__ParsingStrategy updates property "ParsingStrategy".
func (t AWS_Wisdom_KnowledgeBase_ParsingConfiguration) SetV__ParsingStrategy(v string) AWS_Wisdom_KnowledgeBase_ParsingConfiguration {
	t.ParsingStrategy = cfz.V(v)
	return t
}
