// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet__Type is the CloudFormation type for AWS::WAFv2::WebACL.AWSManagedRulesBotControlRuleSet.
	AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet__Type = "AWS::WAFv2::WebACL.AWSManagedRulesBotControlRuleSet"
)

var (
	// AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.AWSManagedRulesBotControlRuleSet.
	AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet__PropertiesMap = struct {
		EnableMachineLearning string
		InspectionLevel       string
	}{
		EnableMachineLearning: "EnableMachineLearning",
		InspectionLevel:       "InspectionLevel",
	}

	// AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.AWSManagedRulesBotControlRuleSet.
	AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet__PropertiesMap.EnableMachineLearning,
		AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet__PropertiesMap.InspectionLevel,
	}
)

// AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet is a binding for AWS::WAFv2::WebACL.AWSManagedRulesBotControlRuleSet.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesbotcontrolruleset.html
type AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet struct {
	// EnableMachineLearning is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesbotcontrolruleset.html#cfn-wafv2-webacl-awsmanagedrulesbotcontrolruleset-enablemachinelearning
	EnableMachineLearning cfz.Expression[bool] `json:"EnableMachineLearning,omitempty"`

	// InspectionLevel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-awsmanagedrulesbotcontrolruleset.html#cfn-wafv2-webacl-awsmanagedrulesbotcontrolruleset-inspectionlevel
	InspectionLevel cfz.Expression[string] `json:"InspectionLevel,omitempty"`
}

// New__AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet initializes a new AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet.
func New__AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet() AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet {
	return AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet) GetType() string {
	return AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet__Type
}

// Set__EnableMachineLearning updates property "EnableMachineLearning".
func (t AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet) Set__EnableMachineLearning(v cfz.Expression[bool]) AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet {
	t.EnableMachineLearning = v
	return t
}

// SetV__EnableMachineLearning updates property "EnableMachineLearning".
func (t AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet) SetV__EnableMachineLearning(v bool) AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet {
	t.EnableMachineLearning = cfz.V(v)
	return t
}

// Set__InspectionLevel updates property "InspectionLevel".
func (t AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet) Set__InspectionLevel(v cfz.Expression[string]) AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet {
	t.InspectionLevel = v
	return t
}

// SetV__InspectionLevel updates property "InspectionLevel".
func (t AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet) SetV__InspectionLevel(v string) AWS_WAFv2_WebACL_AWSManagedRulesBotControlRuleSet {
	t.InspectionLevel = cfz.V(v)
	return t
}
