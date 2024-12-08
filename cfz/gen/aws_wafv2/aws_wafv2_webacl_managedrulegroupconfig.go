// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_ManagedRuleGroupConfig__Type is the CloudFormation type for AWS::WAFv2::WebACL.ManagedRuleGroupConfig.
	AWS_WAFv2_WebACL_ManagedRuleGroupConfig__Type = "AWS::WAFv2::WebACL.ManagedRuleGroupConfig"
)

var (
	// AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.ManagedRuleGroupConfig.
	AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesMap = struct {
		AWSManagedRulesACFPRuleSet       string
		AWSManagedRulesATPRuleSet        string
		AWSManagedRulesBotControlRuleSet string
		LoginPath                        string
		PasswordField                    string
		PayloadType                      string
		UsernameField                    string
	}{
		AWSManagedRulesACFPRuleSet:       "AWSManagedRulesACFPRuleSet",
		AWSManagedRulesATPRuleSet:        "AWSManagedRulesATPRuleSet",
		AWSManagedRulesBotControlRuleSet: "AWSManagedRulesBotControlRuleSet",
		LoginPath:                        "LoginPath",
		PasswordField:                    "PasswordField",
		PayloadType:                      "PayloadType",
		UsernameField:                    "UsernameField",
	}

	// AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.ManagedRuleGroupConfig.
	AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesMap.AWSManagedRulesACFPRuleSet,
		AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesMap.AWSManagedRulesATPRuleSet,
		AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesMap.AWSManagedRulesBotControlRuleSet,
		AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesMap.LoginPath,
		AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesMap.PasswordField,
		AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesMap.PayloadType,
		AWS_WAFv2_WebACL_ManagedRuleGroupConfig__PropertiesMap.UsernameField,
	}
)

// AWS_WAFv2_WebACL_ManagedRuleGroupConfig is a binding for AWS::WAFv2::WebACL.ManagedRuleGroupConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html
type AWS_WAFv2_WebACL_ManagedRuleGroupConfig struct {
	// AWSManagedRulesACFPRuleSet is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-awsmanagedrulesacfpruleset
	AWSManagedRulesACFPRuleSet cfz.Expression[AWS_WAFv2_WebACL_AWSManagedRulesACFPRuleSet] `json:"AWSManagedRulesACFPRuleSet,omitempty"`

	// AWSManagedRulesATPRuleSet is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-awsmanagedrulesatpruleset
	AWSManagedRulesATPRuleSet cfz.Expression[AWS_WAFv2_WebACL_AWSManagedRulesATPRuleSet] `json:"AWSManagedRulesATPRuleSet,omitempty"`

	// AWSManagedRulesBotControlRuleSet is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-awsmanagedrulesbotcontrolruleset
	AWSManagedRulesBotControlRuleSet cfz.Expression[AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet] `json:"AWSManagedRulesBotControlRuleSet,omitempty"`

	// LoginPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-loginpath
	LoginPath cfz.Expression[string] `json:"LoginPath,omitempty"`

	// PasswordField is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-passwordfield
	PasswordField cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier] `json:"PasswordField,omitempty"`

	// PayloadType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-payloadtype
	PayloadType cfz.Expression[string] `json:"PayloadType,omitempty"`

	// UsernameField is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-managedrulegroupconfig.html#cfn-wafv2-webacl-managedrulegroupconfig-usernamefield
	UsernameField cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier] `json:"UsernameField,omitempty"`
}

// New__AWS_WAFv2_WebACL_ManagedRuleGroupConfig initializes a new AWS_WAFv2_WebACL_ManagedRuleGroupConfig.
func New__AWS_WAFv2_WebACL_ManagedRuleGroupConfig() AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	return AWS_WAFv2_WebACL_ManagedRuleGroupConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_ManagedRuleGroupConfig) GetType() string {
	return AWS_WAFv2_WebACL_ManagedRuleGroupConfig__Type
}

// Set__AWSManagedRulesACFPRuleSet updates property "AWSManagedRulesACFPRuleSet".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) Set__AWSManagedRulesACFPRuleSet(v cfz.Expression[AWS_WAFv2_WebACL_AWSManagedRulesACFPRuleSet]) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.AWSManagedRulesACFPRuleSet = v
	return t
}

// SetV__AWSManagedRulesACFPRuleSet updates property "AWSManagedRulesACFPRuleSet".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) SetV__AWSManagedRulesACFPRuleSet(v AWS_WAFv2_WebACL_AWSManagedRulesACFPRuleSet) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.AWSManagedRulesACFPRuleSet = cfz.V(v)
	return t
}

// Set__AWSManagedRulesATPRuleSet updates property "AWSManagedRulesATPRuleSet".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) Set__AWSManagedRulesATPRuleSet(v cfz.Expression[AWS_WAFv2_WebACL_AWSManagedRulesATPRuleSet]) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.AWSManagedRulesATPRuleSet = v
	return t
}

// SetV__AWSManagedRulesATPRuleSet updates property "AWSManagedRulesATPRuleSet".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) SetV__AWSManagedRulesATPRuleSet(v AWS_WAFv2_WebACL_AWSManagedRulesATPRuleSet) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.AWSManagedRulesATPRuleSet = cfz.V(v)
	return t
}

// Set__AWSManagedRulesBotControlRuleSet updates property "AWSManagedRulesBotControlRuleSet".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) Set__AWSManagedRulesBotControlRuleSet(v cfz.Expression[AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet]) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.AWSManagedRulesBotControlRuleSet = v
	return t
}

// SetV__AWSManagedRulesBotControlRuleSet updates property "AWSManagedRulesBotControlRuleSet".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) SetV__AWSManagedRulesBotControlRuleSet(v AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.AWSManagedRulesBotControlRuleSet = cfz.V(v)
	return t
}

// Set__LoginPath updates property "LoginPath".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) Set__LoginPath(v cfz.Expression[string]) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.LoginPath = v
	return t
}

// SetV__LoginPath updates property "LoginPath".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) SetV__LoginPath(v string) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.LoginPath = cfz.V(v)
	return t
}

// Set__PasswordField updates property "PasswordField".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) Set__PasswordField(v cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier]) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.PasswordField = v
	return t
}

// SetV__PasswordField updates property "PasswordField".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) SetV__PasswordField(v AWS_WAFv2_WebACL_FieldIdentifier) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.PasswordField = cfz.V(v)
	return t
}

// Set__PayloadType updates property "PayloadType".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) Set__PayloadType(v cfz.Expression[string]) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.PayloadType = v
	return t
}

// SetV__PayloadType updates property "PayloadType".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) SetV__PayloadType(v string) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.PayloadType = cfz.V(v)
	return t
}

// Set__UsernameField updates property "UsernameField".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) Set__UsernameField(v cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier]) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.UsernameField = v
	return t
}

// SetV__UsernameField updates property "UsernameField".
func (t AWS_WAFv2_WebACL_ManagedRuleGroupConfig) SetV__UsernameField(v AWS_WAFv2_WebACL_FieldIdentifier) AWS_WAFv2_WebACL_ManagedRuleGroupConfig {
	t.UsernameField = cfz.V(v)
	return t
}
