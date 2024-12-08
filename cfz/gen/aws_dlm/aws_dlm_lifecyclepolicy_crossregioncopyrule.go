// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dlm

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__Type is the CloudFormation type for AWS::DLM::LifecyclePolicy.CrossRegionCopyRule.
	AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__Type = "AWS::DLM::LifecyclePolicy.CrossRegionCopyRule"
)

var (
	// AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesMap reports all the CloudFormation properties for AWS::DLM::LifecyclePolicy.CrossRegionCopyRule.
	AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesMap = struct {
		CmkArn        string
		CopyTags      string
		DeprecateRule string
		Encrypted     string
		RetainRule    string
		Target        string
		TargetRegion  string
	}{
		CmkArn:        "CmkArn",
		CopyTags:      "CopyTags",
		DeprecateRule: "DeprecateRule",
		Encrypted:     "Encrypted",
		RetainRule:    "RetainRule",
		Target:        "Target",
		TargetRegion:  "TargetRegion",
	}

	// AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesSlice reports all the CloudFormation properties for AWS::DLM::LifecyclePolicy.CrossRegionCopyRule.
	AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesSlice = []string{
		AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesMap.CmkArn,
		AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesMap.CopyTags,
		AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesMap.DeprecateRule,
		AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesMap.Encrypted,
		AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesMap.RetainRule,
		AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesMap.Target,
		AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__PropertiesMap.TargetRegion,
	}
)

// AWS_DLM_LifecyclePolicy_CrossRegionCopyRule is a binding for AWS::DLM::LifecyclePolicy.CrossRegionCopyRule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html
type AWS_DLM_LifecyclePolicy_CrossRegionCopyRule struct {
	// CmkArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-cmkarn
	CmkArn cfz.Expression[string] `json:"CmkArn,omitempty"`

	// CopyTags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-copytags
	CopyTags cfz.Expression[bool] `json:"CopyTags,omitempty"`

	// DeprecateRule is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-deprecaterule
	DeprecateRule cfz.Expression[AWS_DLM_LifecyclePolicy_CrossRegionCopyDeprecateRule] `json:"DeprecateRule,omitempty"`

	// Encrypted is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-encrypted
	Encrypted cfz.Expression[bool] `json:"Encrypted,omitempty"`

	// RetainRule is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-retainrule
	RetainRule cfz.Expression[AWS_DLM_LifecyclePolicy_CrossRegionCopyRetainRule] `json:"RetainRule,omitempty"`

	// Target is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-target
	Target cfz.Expression[string] `json:"Target,omitempty"`

	// TargetRegion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-targetregion
	TargetRegion cfz.Expression[string] `json:"TargetRegion,omitempty"`
}

// New__AWS_DLM_LifecyclePolicy_CrossRegionCopyRule initializes a new AWS_DLM_LifecyclePolicy_CrossRegionCopyRule.
func New__AWS_DLM_LifecyclePolicy_CrossRegionCopyRule() AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	return AWS_DLM_LifecyclePolicy_CrossRegionCopyRule{}
}

// GetType returns the CloudFormation type.
func (AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) GetType() string {
	return AWS_DLM_LifecyclePolicy_CrossRegionCopyRule__Type
}

// Set__CmkArn updates property "CmkArn".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) Set__CmkArn(v cfz.Expression[string]) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.CmkArn = v
	return t
}

// SetV__CmkArn updates property "CmkArn".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) SetV__CmkArn(v string) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.CmkArn = cfz.V(v)
	return t
}

// Set__CopyTags updates property "CopyTags".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) Set__CopyTags(v cfz.Expression[bool]) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.CopyTags = v
	return t
}

// SetV__CopyTags updates property "CopyTags".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) SetV__CopyTags(v bool) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.CopyTags = cfz.V(v)
	return t
}

// Set__DeprecateRule updates property "DeprecateRule".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) Set__DeprecateRule(v cfz.Expression[AWS_DLM_LifecyclePolicy_CrossRegionCopyDeprecateRule]) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.DeprecateRule = v
	return t
}

// SetV__DeprecateRule updates property "DeprecateRule".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) SetV__DeprecateRule(v AWS_DLM_LifecyclePolicy_CrossRegionCopyDeprecateRule) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.DeprecateRule = cfz.V(v)
	return t
}

// Set__Encrypted updates property "Encrypted".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) Set__Encrypted(v cfz.Expression[bool]) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.Encrypted = v
	return t
}

// SetV__Encrypted updates property "Encrypted".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) SetV__Encrypted(v bool) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.Encrypted = cfz.V(v)
	return t
}

// Set__RetainRule updates property "RetainRule".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) Set__RetainRule(v cfz.Expression[AWS_DLM_LifecyclePolicy_CrossRegionCopyRetainRule]) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.RetainRule = v
	return t
}

// SetV__RetainRule updates property "RetainRule".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) SetV__RetainRule(v AWS_DLM_LifecyclePolicy_CrossRegionCopyRetainRule) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.RetainRule = cfz.V(v)
	return t
}

// Set__Target updates property "Target".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) Set__Target(v cfz.Expression[string]) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.Target = v
	return t
}

// SetV__Target updates property "Target".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) SetV__Target(v string) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.Target = cfz.V(v)
	return t
}

// Set__TargetRegion updates property "TargetRegion".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) Set__TargetRegion(v cfz.Expression[string]) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.TargetRegion = v
	return t
}

// SetV__TargetRegion updates property "TargetRegion".
func (t AWS_DLM_LifecyclePolicy_CrossRegionCopyRule) SetV__TargetRegion(v string) AWS_DLM_LifecyclePolicy_CrossRegionCopyRule {
	t.TargetRegion = cfz.V(v)
	return t
}
