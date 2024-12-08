// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_imagebuilder

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ImageBuilder_LifecyclePolicy_ExclusionRules__Type is the CloudFormation type for AWS::ImageBuilder::LifecyclePolicy.ExclusionRules.
	AWS_ImageBuilder_LifecyclePolicy_ExclusionRules__Type = "AWS::ImageBuilder::LifecyclePolicy.ExclusionRules"
)

var (
	// AWS_ImageBuilder_LifecyclePolicy_ExclusionRules__PropertiesMap reports all the CloudFormation properties for AWS::ImageBuilder::LifecyclePolicy.ExclusionRules.
	AWS_ImageBuilder_LifecyclePolicy_ExclusionRules__PropertiesMap = struct {
		Amis   string
		TagMap string
	}{
		Amis:   "Amis",
		TagMap: "TagMap",
	}

	// AWS_ImageBuilder_LifecyclePolicy_ExclusionRules__PropertiesSlice reports all the CloudFormation properties for AWS::ImageBuilder::LifecyclePolicy.ExclusionRules.
	AWS_ImageBuilder_LifecyclePolicy_ExclusionRules__PropertiesSlice = []string{
		AWS_ImageBuilder_LifecyclePolicy_ExclusionRules__PropertiesMap.Amis,
		AWS_ImageBuilder_LifecyclePolicy_ExclusionRules__PropertiesMap.TagMap,
	}
)

// AWS_ImageBuilder_LifecyclePolicy_ExclusionRules is a binding for AWS::ImageBuilder::LifecyclePolicy.ExclusionRules.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.html
type AWS_ImageBuilder_LifecyclePolicy_ExclusionRules struct {
	// Amis is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.html#cfn-imagebuilder-lifecyclepolicy-exclusionrules-amis
	Amis cfz.Expression[AWS_ImageBuilder_LifecyclePolicy_AmiExclusionRules] `json:"Amis,omitempty"`

	// TagMap is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.html#cfn-imagebuilder-lifecyclepolicy-exclusionrules-tagmap
	TagMap cfz.ExpressionMap[string] `json:"TagMap,omitempty"`
}

// New__AWS_ImageBuilder_LifecyclePolicy_ExclusionRules initializes a new AWS_ImageBuilder_LifecyclePolicy_ExclusionRules.
func New__AWS_ImageBuilder_LifecyclePolicy_ExclusionRules() AWS_ImageBuilder_LifecyclePolicy_ExclusionRules {
	return AWS_ImageBuilder_LifecyclePolicy_ExclusionRules{}
}

// GetType returns the CloudFormation type.
func (AWS_ImageBuilder_LifecyclePolicy_ExclusionRules) GetType() string {
	return AWS_ImageBuilder_LifecyclePolicy_ExclusionRules__Type
}

// Set__Amis updates property "Amis".
func (t AWS_ImageBuilder_LifecyclePolicy_ExclusionRules) Set__Amis(v cfz.Expression[AWS_ImageBuilder_LifecyclePolicy_AmiExclusionRules]) AWS_ImageBuilder_LifecyclePolicy_ExclusionRules {
	t.Amis = v
	return t
}

// SetV__Amis updates property "Amis".
func (t AWS_ImageBuilder_LifecyclePolicy_ExclusionRules) SetV__Amis(v AWS_ImageBuilder_LifecyclePolicy_AmiExclusionRules) AWS_ImageBuilder_LifecyclePolicy_ExclusionRules {
	t.Amis = cfz.V(v)
	return t
}

// Set__TagMap updates property "TagMap".
func (t AWS_ImageBuilder_LifecyclePolicy_ExclusionRules) Set__TagMap(v cfz.ExpressionMap[string]) AWS_ImageBuilder_LifecyclePolicy_ExclusionRules {
	t.TagMap = v
	return t
}

// SetM__TagMap updates property "TagMap".
func (t AWS_ImageBuilder_LifecyclePolicy_ExclusionRules) SetM__TagMap(v ...map[string]cfz.Expression[string]) AWS_ImageBuilder_LifecyclePolicy_ExclusionRules {
	t.TagMap = cfz.M(v...)
	return t
}

// SetMV__TagMap updates property "TagMap".
func (t AWS_ImageBuilder_LifecyclePolicy_ExclusionRules) SetMV__TagMap(v ...map[string]string) AWS_ImageBuilder_LifecyclePolicy_ExclusionRules {
	t.TagMap = cfz.MV(v...)
	return t
}
