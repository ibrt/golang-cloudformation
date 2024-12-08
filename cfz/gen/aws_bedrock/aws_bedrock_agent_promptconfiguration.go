// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Agent_PromptConfiguration__Type is the CloudFormation type for AWS::Bedrock::Agent.PromptConfiguration.
	AWS_Bedrock_Agent_PromptConfiguration__Type = "AWS::Bedrock::Agent.PromptConfiguration"
)

var (
	// AWS_Bedrock_Agent_PromptConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Agent.PromptConfiguration.
	AWS_Bedrock_Agent_PromptConfiguration__PropertiesMap = struct {
		BasePromptTemplate     string
		InferenceConfiguration string
		ParserMode             string
		PromptCreationMode     string
		PromptState            string
		PromptType             string
	}{
		BasePromptTemplate:     "BasePromptTemplate",
		InferenceConfiguration: "InferenceConfiguration",
		ParserMode:             "ParserMode",
		PromptCreationMode:     "PromptCreationMode",
		PromptState:            "PromptState",
		PromptType:             "PromptType",
	}

	// AWS_Bedrock_Agent_PromptConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Agent.PromptConfiguration.
	AWS_Bedrock_Agent_PromptConfiguration__PropertiesSlice = []string{
		AWS_Bedrock_Agent_PromptConfiguration__PropertiesMap.BasePromptTemplate,
		AWS_Bedrock_Agent_PromptConfiguration__PropertiesMap.InferenceConfiguration,
		AWS_Bedrock_Agent_PromptConfiguration__PropertiesMap.ParserMode,
		AWS_Bedrock_Agent_PromptConfiguration__PropertiesMap.PromptCreationMode,
		AWS_Bedrock_Agent_PromptConfiguration__PropertiesMap.PromptState,
		AWS_Bedrock_Agent_PromptConfiguration__PropertiesMap.PromptType,
	}
)

// AWS_Bedrock_Agent_PromptConfiguration is a binding for AWS::Bedrock::Agent.PromptConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html
type AWS_Bedrock_Agent_PromptConfiguration struct {
	// BasePromptTemplate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-baseprompttemplate
	BasePromptTemplate cfz.Expression[string] `json:"BasePromptTemplate,omitempty"`

	// InferenceConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-inferenceconfiguration
	InferenceConfiguration cfz.Expression[AWS_Bedrock_Agent_InferenceConfiguration] `json:"InferenceConfiguration,omitempty"`

	// ParserMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-parsermode
	ParserMode cfz.Expression[string] `json:"ParserMode,omitempty"`

	// PromptCreationMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-promptcreationmode
	PromptCreationMode cfz.Expression[string] `json:"PromptCreationMode,omitempty"`

	// PromptState is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-promptstate
	PromptState cfz.Expression[string] `json:"PromptState,omitempty"`

	// PromptType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-prompttype
	PromptType cfz.Expression[string] `json:"PromptType,omitempty"`
}

// New__AWS_Bedrock_Agent_PromptConfiguration initializes a new AWS_Bedrock_Agent_PromptConfiguration.
func New__AWS_Bedrock_Agent_PromptConfiguration() AWS_Bedrock_Agent_PromptConfiguration {
	return AWS_Bedrock_Agent_PromptConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Agent_PromptConfiguration) GetType() string {
	return AWS_Bedrock_Agent_PromptConfiguration__Type
}

// Set__BasePromptTemplate updates property "BasePromptTemplate".
func (t AWS_Bedrock_Agent_PromptConfiguration) Set__BasePromptTemplate(v cfz.Expression[string]) AWS_Bedrock_Agent_PromptConfiguration {
	t.BasePromptTemplate = v
	return t
}

// SetV__BasePromptTemplate updates property "BasePromptTemplate".
func (t AWS_Bedrock_Agent_PromptConfiguration) SetV__BasePromptTemplate(v string) AWS_Bedrock_Agent_PromptConfiguration {
	t.BasePromptTemplate = cfz.V(v)
	return t
}

// Set__InferenceConfiguration updates property "InferenceConfiguration".
func (t AWS_Bedrock_Agent_PromptConfiguration) Set__InferenceConfiguration(v cfz.Expression[AWS_Bedrock_Agent_InferenceConfiguration]) AWS_Bedrock_Agent_PromptConfiguration {
	t.InferenceConfiguration = v
	return t
}

// SetV__InferenceConfiguration updates property "InferenceConfiguration".
func (t AWS_Bedrock_Agent_PromptConfiguration) SetV__InferenceConfiguration(v AWS_Bedrock_Agent_InferenceConfiguration) AWS_Bedrock_Agent_PromptConfiguration {
	t.InferenceConfiguration = cfz.V(v)
	return t
}

// Set__ParserMode updates property "ParserMode".
func (t AWS_Bedrock_Agent_PromptConfiguration) Set__ParserMode(v cfz.Expression[string]) AWS_Bedrock_Agent_PromptConfiguration {
	t.ParserMode = v
	return t
}

// SetV__ParserMode updates property "ParserMode".
func (t AWS_Bedrock_Agent_PromptConfiguration) SetV__ParserMode(v string) AWS_Bedrock_Agent_PromptConfiguration {
	t.ParserMode = cfz.V(v)
	return t
}

// Set__PromptCreationMode updates property "PromptCreationMode".
func (t AWS_Bedrock_Agent_PromptConfiguration) Set__PromptCreationMode(v cfz.Expression[string]) AWS_Bedrock_Agent_PromptConfiguration {
	t.PromptCreationMode = v
	return t
}

// SetV__PromptCreationMode updates property "PromptCreationMode".
func (t AWS_Bedrock_Agent_PromptConfiguration) SetV__PromptCreationMode(v string) AWS_Bedrock_Agent_PromptConfiguration {
	t.PromptCreationMode = cfz.V(v)
	return t
}

// Set__PromptState updates property "PromptState".
func (t AWS_Bedrock_Agent_PromptConfiguration) Set__PromptState(v cfz.Expression[string]) AWS_Bedrock_Agent_PromptConfiguration {
	t.PromptState = v
	return t
}

// SetV__PromptState updates property "PromptState".
func (t AWS_Bedrock_Agent_PromptConfiguration) SetV__PromptState(v string) AWS_Bedrock_Agent_PromptConfiguration {
	t.PromptState = cfz.V(v)
	return t
}

// Set__PromptType updates property "PromptType".
func (t AWS_Bedrock_Agent_PromptConfiguration) Set__PromptType(v cfz.Expression[string]) AWS_Bedrock_Agent_PromptConfiguration {
	t.PromptType = v
	return t
}

// SetV__PromptType updates property "PromptType".
func (t AWS_Bedrock_Agent_PromptConfiguration) SetV__PromptType(v string) AWS_Bedrock_Agent_PromptConfiguration {
	t.PromptType = cfz.V(v)
	return t
}
