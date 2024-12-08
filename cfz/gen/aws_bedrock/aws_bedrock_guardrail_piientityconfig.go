// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Guardrail_PiiEntityConfig__Type is the CloudFormation type for AWS::Bedrock::Guardrail.PiiEntityConfig.
	AWS_Bedrock_Guardrail_PiiEntityConfig__Type = "AWS::Bedrock::Guardrail.PiiEntityConfig"
)

var (
	// AWS_Bedrock_Guardrail_PiiEntityConfig__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Guardrail.PiiEntityConfig.
	AWS_Bedrock_Guardrail_PiiEntityConfig__PropertiesMap = struct {
		Action string
		Type   string
	}{
		Action: "Action",
		Type:   "Type",
	}

	// AWS_Bedrock_Guardrail_PiiEntityConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Guardrail.PiiEntityConfig.
	AWS_Bedrock_Guardrail_PiiEntityConfig__PropertiesSlice = []string{
		AWS_Bedrock_Guardrail_PiiEntityConfig__PropertiesMap.Action,
		AWS_Bedrock_Guardrail_PiiEntityConfig__PropertiesMap.Type,
	}
)

// AWS_Bedrock_Guardrail_PiiEntityConfig is a binding for AWS::Bedrock::Guardrail.PiiEntityConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-piientityconfig.html
type AWS_Bedrock_Guardrail_PiiEntityConfig struct {
	// Action is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-piientityconfig.html#cfn-bedrock-guardrail-piientityconfig-action
	Action cfz.Expression[string] `json:"Action,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-piientityconfig.html#cfn-bedrock-guardrail-piientityconfig-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_Bedrock_Guardrail_PiiEntityConfig initializes a new AWS_Bedrock_Guardrail_PiiEntityConfig.
func New__AWS_Bedrock_Guardrail_PiiEntityConfig() AWS_Bedrock_Guardrail_PiiEntityConfig {
	return AWS_Bedrock_Guardrail_PiiEntityConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Guardrail_PiiEntityConfig) GetType() string {
	return AWS_Bedrock_Guardrail_PiiEntityConfig__Type
}

// Set__Action updates property "Action".
func (t AWS_Bedrock_Guardrail_PiiEntityConfig) Set__Action(v cfz.Expression[string]) AWS_Bedrock_Guardrail_PiiEntityConfig {
	t.Action = v
	return t
}

// SetV__Action updates property "Action".
func (t AWS_Bedrock_Guardrail_PiiEntityConfig) SetV__Action(v string) AWS_Bedrock_Guardrail_PiiEntityConfig {
	t.Action = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_Bedrock_Guardrail_PiiEntityConfig) Set__Type(v cfz.Expression[string]) AWS_Bedrock_Guardrail_PiiEntityConfig {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_Bedrock_Guardrail_PiiEntityConfig) SetV__Type(v string) AWS_Bedrock_Guardrail_PiiEntityConfig {
	t.Type = cfz.V(v)
	return t
}
