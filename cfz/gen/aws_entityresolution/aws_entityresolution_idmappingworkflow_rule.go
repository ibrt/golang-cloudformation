// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_entityresolution

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EntityResolution_IdMappingWorkflow_Rule__Type is the CloudFormation type for AWS::EntityResolution::IdMappingWorkflow.Rule.
	AWS_EntityResolution_IdMappingWorkflow_Rule__Type = "AWS::EntityResolution::IdMappingWorkflow.Rule"
)

var (
	// AWS_EntityResolution_IdMappingWorkflow_Rule__PropertiesMap reports all the CloudFormation properties for AWS::EntityResolution::IdMappingWorkflow.Rule.
	AWS_EntityResolution_IdMappingWorkflow_Rule__PropertiesMap = struct {
		MatchingKeys string
		RuleName     string
	}{
		MatchingKeys: "MatchingKeys",
		RuleName:     "RuleName",
	}

	// AWS_EntityResolution_IdMappingWorkflow_Rule__PropertiesSlice reports all the CloudFormation properties for AWS::EntityResolution::IdMappingWorkflow.Rule.
	AWS_EntityResolution_IdMappingWorkflow_Rule__PropertiesSlice = []string{
		AWS_EntityResolution_IdMappingWorkflow_Rule__PropertiesMap.MatchingKeys,
		AWS_EntityResolution_IdMappingWorkflow_Rule__PropertiesMap.RuleName,
	}
)

// AWS_EntityResolution_IdMappingWorkflow_Rule is a binding for AWS::EntityResolution::IdMappingWorkflow.Rule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-rule.html
type AWS_EntityResolution_IdMappingWorkflow_Rule struct {
	// MatchingKeys is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-rule.html#cfn-entityresolution-idmappingworkflow-rule-matchingkeys
	MatchingKeys cfz.ExpressionSlice[string] `json:"MatchingKeys,omitempty"`

	// RuleName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-rule.html#cfn-entityresolution-idmappingworkflow-rule-rulename
	RuleName cfz.Expression[string] `json:"RuleName,omitempty"`
}

// New__AWS_EntityResolution_IdMappingWorkflow_Rule initializes a new AWS_EntityResolution_IdMappingWorkflow_Rule.
func New__AWS_EntityResolution_IdMappingWorkflow_Rule() AWS_EntityResolution_IdMappingWorkflow_Rule {
	return AWS_EntityResolution_IdMappingWorkflow_Rule{}
}

// GetType returns the CloudFormation type.
func (AWS_EntityResolution_IdMappingWorkflow_Rule) GetType() string {
	return AWS_EntityResolution_IdMappingWorkflow_Rule__Type
}

// Set__MatchingKeys updates property "MatchingKeys".
func (t AWS_EntityResolution_IdMappingWorkflow_Rule) Set__MatchingKeys(v cfz.ExpressionSlice[string]) AWS_EntityResolution_IdMappingWorkflow_Rule {
	t.MatchingKeys = v
	return t
}

// SetS__MatchingKeys updates property "MatchingKeys".
func (t AWS_EntityResolution_IdMappingWorkflow_Rule) SetS__MatchingKeys(v ...cfz.Expression[string]) AWS_EntityResolution_IdMappingWorkflow_Rule {
	t.MatchingKeys = cfz.S(v...)
	return t
}

// SetSV__MatchingKeys updates property "MatchingKeys".
func (t AWS_EntityResolution_IdMappingWorkflow_Rule) SetSV__MatchingKeys(v ...string) AWS_EntityResolution_IdMappingWorkflow_Rule {
	t.MatchingKeys = cfz.SV(v...)
	return t
}

// Set__RuleName updates property "RuleName".
func (t AWS_EntityResolution_IdMappingWorkflow_Rule) Set__RuleName(v cfz.Expression[string]) AWS_EntityResolution_IdMappingWorkflow_Rule {
	t.RuleName = v
	return t
}

// SetV__RuleName updates property "RuleName".
func (t AWS_EntityResolution_IdMappingWorkflow_Rule) SetV__RuleName(v string) AWS_EntityResolution_IdMappingWorkflow_Rule {
	t.RuleName = cfz.V(v)
	return t
}
