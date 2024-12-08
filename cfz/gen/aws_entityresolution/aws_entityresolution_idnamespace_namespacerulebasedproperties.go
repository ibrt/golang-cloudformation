// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_entityresolution

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__Type is the CloudFormation type for AWS::EntityResolution::IdNamespace.NamespaceRuleBasedProperties.
	AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__Type = "AWS::EntityResolution::IdNamespace.NamespaceRuleBasedProperties"
)

var (
	// AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__PropertiesMap reports all the CloudFormation properties for AWS::EntityResolution::IdNamespace.NamespaceRuleBasedProperties.
	AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__PropertiesMap = struct {
		AttributeMatchingModel string
		RecordMatchingModels   string
		RuleDefinitionTypes    string
		Rules                  string
	}{
		AttributeMatchingModel: "AttributeMatchingModel",
		RecordMatchingModels:   "RecordMatchingModels",
		RuleDefinitionTypes:    "RuleDefinitionTypes",
		Rules:                  "Rules",
	}

	// AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__PropertiesSlice reports all the CloudFormation properties for AWS::EntityResolution::IdNamespace.NamespaceRuleBasedProperties.
	AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__PropertiesSlice = []string{
		AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__PropertiesMap.AttributeMatchingModel,
		AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__PropertiesMap.RecordMatchingModels,
		AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__PropertiesMap.RuleDefinitionTypes,
		AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__PropertiesMap.Rules,
	}
)

// AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties is a binding for AWS::EntityResolution::IdNamespace.NamespaceRuleBasedProperties.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html
type AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties struct {
	// AttributeMatchingModel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html#cfn-entityresolution-idnamespace-namespacerulebasedproperties-attributematchingmodel
	AttributeMatchingModel cfz.Expression[string] `json:"AttributeMatchingModel,omitempty"`

	// RecordMatchingModels is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html#cfn-entityresolution-idnamespace-namespacerulebasedproperties-recordmatchingmodels
	RecordMatchingModels cfz.ExpressionSlice[string] `json:"RecordMatchingModels,omitempty"`

	// RuleDefinitionTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html#cfn-entityresolution-idnamespace-namespacerulebasedproperties-ruledefinitiontypes
	RuleDefinitionTypes cfz.ExpressionSlice[string] `json:"RuleDefinitionTypes,omitempty"`

	// Rules is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.html#cfn-entityresolution-idnamespace-namespacerulebasedproperties-rules
	Rules cfz.ExpressionSlice[AWS_EntityResolution_IdNamespace_Rule] `json:"Rules,omitempty"`
}

// New__AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties initializes a new AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties.
func New__AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties() AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	return AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties{}
}

// GetType returns the CloudFormation type.
func (AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) GetType() string {
	return AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties__Type
}

// Set__AttributeMatchingModel updates property "AttributeMatchingModel".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) Set__AttributeMatchingModel(v cfz.Expression[string]) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.AttributeMatchingModel = v
	return t
}

// SetV__AttributeMatchingModel updates property "AttributeMatchingModel".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) SetV__AttributeMatchingModel(v string) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.AttributeMatchingModel = cfz.V(v)
	return t
}

// Set__RecordMatchingModels updates property "RecordMatchingModels".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) Set__RecordMatchingModels(v cfz.ExpressionSlice[string]) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.RecordMatchingModels = v
	return t
}

// SetS__RecordMatchingModels updates property "RecordMatchingModels".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) SetS__RecordMatchingModels(v ...cfz.Expression[string]) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.RecordMatchingModels = cfz.S(v...)
	return t
}

// SetSV__RecordMatchingModels updates property "RecordMatchingModels".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) SetSV__RecordMatchingModels(v ...string) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.RecordMatchingModels = cfz.SV(v...)
	return t
}

// Set__RuleDefinitionTypes updates property "RuleDefinitionTypes".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) Set__RuleDefinitionTypes(v cfz.ExpressionSlice[string]) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.RuleDefinitionTypes = v
	return t
}

// SetS__RuleDefinitionTypes updates property "RuleDefinitionTypes".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) SetS__RuleDefinitionTypes(v ...cfz.Expression[string]) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.RuleDefinitionTypes = cfz.S(v...)
	return t
}

// SetSV__RuleDefinitionTypes updates property "RuleDefinitionTypes".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) SetSV__RuleDefinitionTypes(v ...string) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.RuleDefinitionTypes = cfz.SV(v...)
	return t
}

// Set__Rules updates property "Rules".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) Set__Rules(v cfz.ExpressionSlice[AWS_EntityResolution_IdNamespace_Rule]) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.Rules = v
	return t
}

// SetS__Rules updates property "Rules".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) SetS__Rules(v ...cfz.Expression[AWS_EntityResolution_IdNamespace_Rule]) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.Rules = cfz.S(v...)
	return t
}

// SetSV__Rules updates property "Rules".
func (t AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties) SetSV__Rules(v ...AWS_EntityResolution_IdNamespace_Rule) AWS_EntityResolution_IdNamespace_NamespaceRuleBasedProperties {
	t.Rules = cfz.SV(v...)
	return t
}
