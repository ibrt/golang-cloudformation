// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig__Type is the CloudFormation type for AWS::Bedrock::Guardrail.SensitiveInformationPolicyConfig.
	AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig__Type = "AWS::Bedrock::Guardrail.SensitiveInformationPolicyConfig"
)

var (
	// AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Guardrail.SensitiveInformationPolicyConfig.
	AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig__PropertiesMap = struct {
		PiiEntitiesConfig string
		RegexesConfig     string
	}{
		PiiEntitiesConfig: "PiiEntitiesConfig",
		RegexesConfig:     "RegexesConfig",
	}

	// AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Guardrail.SensitiveInformationPolicyConfig.
	AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig__PropertiesSlice = []string{
		AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig__PropertiesMap.PiiEntitiesConfig,
		AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig__PropertiesMap.RegexesConfig,
	}
)

// AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig is a binding for AWS::Bedrock::Guardrail.SensitiveInformationPolicyConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-sensitiveinformationpolicyconfig.html
type AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig struct {
	// PiiEntitiesConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-sensitiveinformationpolicyconfig.html#cfn-bedrock-guardrail-sensitiveinformationpolicyconfig-piientitiesconfig
	PiiEntitiesConfig cfz.ExpressionSlice[AWS_Bedrock_Guardrail_PiiEntityConfig] `json:"PiiEntitiesConfig,omitempty"`

	// RegexesConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-sensitiveinformationpolicyconfig.html#cfn-bedrock-guardrail-sensitiveinformationpolicyconfig-regexesconfig
	RegexesConfig cfz.ExpressionSlice[AWS_Bedrock_Guardrail_RegexConfig] `json:"RegexesConfig,omitempty"`
}

// New__AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig initializes a new AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig.
func New__AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig() AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig {
	return AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig) GetType() string {
	return AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig__Type
}

// Set__PiiEntitiesConfig updates property "PiiEntitiesConfig".
func (t AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig) Set__PiiEntitiesConfig(v cfz.ExpressionSlice[AWS_Bedrock_Guardrail_PiiEntityConfig]) AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig {
	t.PiiEntitiesConfig = v
	return t
}

// SetS__PiiEntitiesConfig updates property "PiiEntitiesConfig".
func (t AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig) SetS__PiiEntitiesConfig(v ...cfz.Expression[AWS_Bedrock_Guardrail_PiiEntityConfig]) AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig {
	t.PiiEntitiesConfig = cfz.S(v...)
	return t
}

// SetSV__PiiEntitiesConfig updates property "PiiEntitiesConfig".
func (t AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig) SetSV__PiiEntitiesConfig(v ...AWS_Bedrock_Guardrail_PiiEntityConfig) AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig {
	t.PiiEntitiesConfig = cfz.SV(v...)
	return t
}

// Set__RegexesConfig updates property "RegexesConfig".
func (t AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig) Set__RegexesConfig(v cfz.ExpressionSlice[AWS_Bedrock_Guardrail_RegexConfig]) AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig {
	t.RegexesConfig = v
	return t
}

// SetS__RegexesConfig updates property "RegexesConfig".
func (t AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig) SetS__RegexesConfig(v ...cfz.Expression[AWS_Bedrock_Guardrail_RegexConfig]) AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig {
	t.RegexesConfig = cfz.S(v...)
	return t
}

// SetSV__RegexesConfig updates property "RegexesConfig".
func (t AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig) SetSV__RegexesConfig(v ...AWS_Bedrock_Guardrail_RegexConfig) AWS_Bedrock_Guardrail_SensitiveInformationPolicyConfig {
	t.RegexesConfig = cfz.SV(v...)
	return t
}
