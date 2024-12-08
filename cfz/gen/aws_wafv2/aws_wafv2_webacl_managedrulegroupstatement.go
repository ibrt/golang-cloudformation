// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_ManagedRuleGroupStatement__Type is the CloudFormation type for AWS::WAFv2::WebACL.ManagedRuleGroupStatement.
	AWS_WAFv2_WebACL_ManagedRuleGroupStatement__Type = "AWS::WAFv2::WebACL.ManagedRuleGroupStatement"
)

var (
	// AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.ManagedRuleGroupStatement.
	AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesMap = struct {
		ExcludedRules           string
		ManagedRuleGroupConfigs string
		Name                    string
		RuleActionOverrides     string
		ScopeDownStatement      string
		VendorName              string
		Version                 string
	}{
		ExcludedRules:           "ExcludedRules",
		ManagedRuleGroupConfigs: "ManagedRuleGroupConfigs",
		Name:                    "Name",
		RuleActionOverrides:     "RuleActionOverrides",
		ScopeDownStatement:      "ScopeDownStatement",
		VendorName:              "VendorName",
		Version:                 "Version",
	}

	// AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.ManagedRuleGroupStatement.
	AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesMap.ExcludedRules,
		AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesMap.ManagedRuleGroupConfigs,
		AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesMap.Name,
		AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesMap.RuleActionOverrides,
		AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesMap.ScopeDownStatement,
		AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesMap.VendorName,
		AWS_WAFv2_WebACL_ManagedRuleGroupStatement__PropertiesMap.Version,
	}
)

// AWS_WAFv2_WebACL_ManagedRuleGroupStatement is a binding for AWS::WAFv2::WebACL.ManagedRuleGroupStatement.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html
type AWS_WAFv2_WebACL_ManagedRuleGroupStatement struct {
	// ExcludedRules is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-excludedrules
	ExcludedRules cfz.ExpressionSlice[AWS_WAFv2_WebACL_ExcludedRule] `json:"ExcludedRules,omitempty"`

	// ManagedRuleGroupConfigs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-managedrulegroupconfigs
	ManagedRuleGroupConfigs cfz.ExpressionSlice[AWS_WAFv2_WebACL_ManagedRuleGroupConfig] `json:"ManagedRuleGroupConfigs,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// RuleActionOverrides is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-ruleactionoverrides
	RuleActionOverrides cfz.ExpressionSlice[AWS_WAFv2_WebACL_RuleActionOverride] `json:"RuleActionOverrides,omitempty"`

	// ScopeDownStatement is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-scopedownstatement
	ScopeDownStatement cfz.Expression[AWS_WAFv2_WebACL_Statement] `json:"ScopeDownStatement,omitempty"`

	// VendorName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-vendorname
	VendorName cfz.Expression[string] `json:"VendorName,omitempty"`

	// Version is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupstatement.html#cfn-wafv2-webacl-managedrulegroupstatement-version
	Version cfz.Expression[string] `json:"Version,omitempty"`
}

// New__AWS_WAFv2_WebACL_ManagedRuleGroupStatement initializes a new AWS_WAFv2_WebACL_ManagedRuleGroupStatement.
func New__AWS_WAFv2_WebACL_ManagedRuleGroupStatement() AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	return AWS_WAFv2_WebACL_ManagedRuleGroupStatement{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_ManagedRuleGroupStatement) GetType() string {
	return AWS_WAFv2_WebACL_ManagedRuleGroupStatement__Type
}

// Set__ExcludedRules updates property "ExcludedRules".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) Set__ExcludedRules(v cfz.ExpressionSlice[AWS_WAFv2_WebACL_ExcludedRule]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.ExcludedRules = v
	return t
}

// SetS__ExcludedRules updates property "ExcludedRules".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetS__ExcludedRules(v ...cfz.Expression[AWS_WAFv2_WebACL_ExcludedRule]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.ExcludedRules = cfz.S(v...)
	return t
}

// SetSV__ExcludedRules updates property "ExcludedRules".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetSV__ExcludedRules(v ...AWS_WAFv2_WebACL_ExcludedRule) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.ExcludedRules = cfz.SV(v...)
	return t
}

// Set__ManagedRuleGroupConfigs updates property "ManagedRuleGroupConfigs".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) Set__ManagedRuleGroupConfigs(v cfz.ExpressionSlice[AWS_WAFv2_WebACL_ManagedRuleGroupConfig]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.ManagedRuleGroupConfigs = v
	return t
}

// SetS__ManagedRuleGroupConfigs updates property "ManagedRuleGroupConfigs".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetS__ManagedRuleGroupConfigs(v ...cfz.Expression[AWS_WAFv2_WebACL_ManagedRuleGroupConfig]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.ManagedRuleGroupConfigs = cfz.S(v...)
	return t
}

// SetSV__ManagedRuleGroupConfigs updates property "ManagedRuleGroupConfigs".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetSV__ManagedRuleGroupConfigs(v ...AWS_WAFv2_WebACL_ManagedRuleGroupConfig) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.ManagedRuleGroupConfigs = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) Set__Name(v cfz.Expression[string]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetV__Name(v string) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.Name = cfz.V(v)
	return t
}

// Set__RuleActionOverrides updates property "RuleActionOverrides".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) Set__RuleActionOverrides(v cfz.ExpressionSlice[AWS_WAFv2_WebACL_RuleActionOverride]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.RuleActionOverrides = v
	return t
}

// SetS__RuleActionOverrides updates property "RuleActionOverrides".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetS__RuleActionOverrides(v ...cfz.Expression[AWS_WAFv2_WebACL_RuleActionOverride]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.RuleActionOverrides = cfz.S(v...)
	return t
}

// SetSV__RuleActionOverrides updates property "RuleActionOverrides".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetSV__RuleActionOverrides(v ...AWS_WAFv2_WebACL_RuleActionOverride) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.RuleActionOverrides = cfz.SV(v...)
	return t
}

// Set__ScopeDownStatement updates property "ScopeDownStatement".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) Set__ScopeDownStatement(v cfz.Expression[AWS_WAFv2_WebACL_Statement]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.ScopeDownStatement = v
	return t
}

// SetV__ScopeDownStatement updates property "ScopeDownStatement".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetV__ScopeDownStatement(v AWS_WAFv2_WebACL_Statement) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.ScopeDownStatement = cfz.V(v)
	return t
}

// Set__VendorName updates property "VendorName".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) Set__VendorName(v cfz.Expression[string]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.VendorName = v
	return t
}

// SetV__VendorName updates property "VendorName".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetV__VendorName(v string) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.VendorName = cfz.V(v)
	return t
}

// Set__Version updates property "Version".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) Set__Version(v cfz.Expression[string]) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.Version = v
	return t
}

// SetV__Version updates property "Version".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupStatement) SetV__Version(v string) AWS_WAFv2_WebACL_ManagedRuleGroupStatement {
	t.Version = cfz.V(v)
	return t
}
