// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig__Type is the CloudFormation type for AWS::Bedrock::Guardrail.ContextualGroundingFilterConfig.
	AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig__Type = "AWS::Bedrock::Guardrail.ContextualGroundingFilterConfig"
)

var (
	// AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Guardrail.ContextualGroundingFilterConfig.
	AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig__PropertiesMap = struct {
		Threshold string
		Type      string
	}{
		Threshold: "Threshold",
		Type:      "Type",
	}

	// AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Guardrail.ContextualGroundingFilterConfig.
	AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig__PropertiesSlice = []string{
		AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig__PropertiesMap.Threshold,
		AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig__PropertiesMap.Type,
	}
)

// AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig is a binding for AWS::Bedrock::Guardrail.ContextualGroundingFilterConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html
type AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig struct {
	// Threshold is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-threshold
	Threshold cfz.Expression[float64] `json:"Threshold,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.html#cfn-bedrock-guardrail-contextualgroundingfilterconfig-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig initializes a new AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig.
func New__AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig() AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig {
	return AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig) GetType() string {
	return AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig__Type
}

// Set__Threshold updates property "Threshold".
func (t AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig) Set__Threshold(v cfz.Expression[float64]) AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig {
	t.Threshold = v
	return t
}

// SetV__Threshold updates property "Threshold".
func (t AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig) SetV__Threshold(v float64) AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig {
	t.Threshold = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig) Set__Type(v cfz.Expression[string]) AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig) SetV__Type(v string) AWS_Bedrock_Guardrail_ContextualGroundingFilterConfig {
	t.Type = cfz.V(v)
	return t
}
