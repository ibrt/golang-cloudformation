// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_config

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__Type is the CloudFormation type for AWS::Config::OrganizationConfigRule.OrganizationManagedRuleMetadata.
	AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__Type = "AWS::Config::OrganizationConfigRule.OrganizationManagedRuleMetadata"
)

var (
	// AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap reports all the CloudFormation properties for AWS::Config::OrganizationConfigRule.OrganizationManagedRuleMetadata.
	AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap = struct {
		Description               string
		InputParameters           string
		MaximumExecutionFrequency string
		ResourceIdScope           string
		ResourceTypesScope        string
		RuleIdentifier            string
		TagKeyScope               string
		TagValueScope             string
	}{
		Description:               "Description",
		InputParameters:           "InputParameters",
		MaximumExecutionFrequency: "MaximumExecutionFrequency",
		ResourceIdScope:           "ResourceIdScope",
		ResourceTypesScope:        "ResourceTypesScope",
		RuleIdentifier:            "RuleIdentifier",
		TagKeyScope:               "TagKeyScope",
		TagValueScope:             "TagValueScope",
	}

	// AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesSlice reports all the CloudFormation properties for AWS::Config::OrganizationConfigRule.OrganizationManagedRuleMetadata.
	AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesSlice = []string{
		AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap.Description,
		AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap.InputParameters,
		AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap.MaximumExecutionFrequency,
		AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap.ResourceIdScope,
		AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap.ResourceTypesScope,
		AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap.RuleIdentifier,
		AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap.TagKeyScope,
		AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__PropertiesMap.TagValueScope,
	}
)

// AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata is a binding for AWS::Config::OrganizationConfigRule.OrganizationManagedRuleMetadata.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html
type AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata struct {
	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// InputParameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-inputparameters
	InputParameters cfz.Expression[string] `json:"InputParameters,omitempty"`

	// MaximumExecutionFrequency is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-maximumexecutionfrequency
	MaximumExecutionFrequency cfz.Expression[string] `json:"MaximumExecutionFrequency,omitempty"`

	// ResourceIdScope is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-resourceidscope
	ResourceIdScope cfz.Expression[string] `json:"ResourceIdScope,omitempty"`

	// ResourceTypesScope is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-resourcetypesscope
	ResourceTypesScope cfz.ExpressionSlice[string] `json:"ResourceTypesScope,omitempty"`

	// RuleIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-ruleidentifier
	RuleIdentifier cfz.Expression[string] `json:"RuleIdentifier,omitempty"`

	// TagKeyScope is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-tagkeyscope
	TagKeyScope cfz.Expression[string] `json:"TagKeyScope,omitempty"`

	// TagValueScope is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-tagvaluescope
	TagValueScope cfz.Expression[string] `json:"TagValueScope,omitempty"`
}

// New__AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata initializes a new AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata.
func New__AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata() AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	return AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata{}
}

// GetType returns the CloudFormation type.
func (AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) GetType() string {
	return AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata__Type
}

// Set__Description updates property "Description".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) Set__Description(v cfz.Expression[string]) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) SetV__Description(v string) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.Description = cfz.V(v)
	return t
}

// Set__InputParameters updates property "InputParameters".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) Set__InputParameters(v cfz.Expression[string]) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.InputParameters = v
	return t
}

// SetV__InputParameters updates property "InputParameters".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) SetV__InputParameters(v string) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.InputParameters = cfz.V(v)
	return t
}

// Set__MaximumExecutionFrequency updates property "MaximumExecutionFrequency".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) Set__MaximumExecutionFrequency(v cfz.Expression[string]) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.MaximumExecutionFrequency = v
	return t
}

// SetV__MaximumExecutionFrequency updates property "MaximumExecutionFrequency".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) SetV__MaximumExecutionFrequency(v string) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.MaximumExecutionFrequency = cfz.V(v)
	return t
}

// Set__ResourceIdScope updates property "ResourceIdScope".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) Set__ResourceIdScope(v cfz.Expression[string]) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.ResourceIdScope = v
	return t
}

// SetV__ResourceIdScope updates property "ResourceIdScope".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) SetV__ResourceIdScope(v string) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.ResourceIdScope = cfz.V(v)
	return t
}

// Set__ResourceTypesScope updates property "ResourceTypesScope".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) Set__ResourceTypesScope(v cfz.ExpressionSlice[string]) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.ResourceTypesScope = v
	return t
}

// SetS__ResourceTypesScope updates property "ResourceTypesScope".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) SetS__ResourceTypesScope(v ...cfz.Expression[string]) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.ResourceTypesScope = cfz.S(v...)
	return t
}

// SetSV__ResourceTypesScope updates property "ResourceTypesScope".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) SetSV__ResourceTypesScope(v ...string) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.ResourceTypesScope = cfz.SV(v...)
	return t
}

// Set__RuleIdentifier updates property "RuleIdentifier".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) Set__RuleIdentifier(v cfz.Expression[string]) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.RuleIdentifier = v
	return t
}

// SetV__RuleIdentifier updates property "RuleIdentifier".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) SetV__RuleIdentifier(v string) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.RuleIdentifier = cfz.V(v)
	return t
}

// Set__TagKeyScope updates property "TagKeyScope".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) Set__TagKeyScope(v cfz.Expression[string]) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.TagKeyScope = v
	return t
}

// SetV__TagKeyScope updates property "TagKeyScope".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) SetV__TagKeyScope(v string) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.TagKeyScope = cfz.V(v)
	return t
}

// Set__TagValueScope updates property "TagValueScope".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) Set__TagValueScope(v cfz.Expression[string]) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.TagValueScope = v
	return t
}

// SetV__TagValueScope updates property "TagValueScope".
func (t AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata) SetV__TagValueScope(v string) AWS_Config_OrganizationConfigRule_OrganizationManagedRuleMetadata {
	t.TagValueScope = cfz.V(v)
	return t
}
