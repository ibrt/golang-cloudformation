// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafregional

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFRegional_RateBasedRule_Predicate__Type is the CloudFormation type for AWS::WAFRegional::RateBasedRule.Predicate.
	AWS_WAFRegional_RateBasedRule_Predicate__Type = "AWS::WAFRegional::RateBasedRule.Predicate"
)

var (
	// AWS_WAFRegional_RateBasedRule_Predicate__PropertiesMap reports all the CloudFormation properties for AWS::WAFRegional::RateBasedRule.Predicate.
	AWS_WAFRegional_RateBasedRule_Predicate__PropertiesMap = struct {
		DataId  string
		Negated string
		Type    string
	}{
		DataId:  "DataId",
		Negated: "Negated",
		Type:    "Type",
	}

	// AWS_WAFRegional_RateBasedRule_Predicate__PropertiesSlice reports all the CloudFormation properties for AWS::WAFRegional::RateBasedRule.Predicate.
	AWS_WAFRegional_RateBasedRule_Predicate__PropertiesSlice = []string{
		AWS_WAFRegional_RateBasedRule_Predicate__PropertiesMap.DataId,
		AWS_WAFRegional_RateBasedRule_Predicate__PropertiesMap.Negated,
		AWS_WAFRegional_RateBasedRule_Predicate__PropertiesMap.Type,
	}
)

// AWS_WAFRegional_RateBasedRule_Predicate is a binding for AWS::WAFRegional::RateBasedRule.Predicate.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ratebasedrule-predicate.html
type AWS_WAFRegional_RateBasedRule_Predicate struct {
	// DataId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ratebasedrule-predicate.html#cfn-wafregional-ratebasedrule-predicate-dataid
	DataId cfz.Expression[string] `json:"DataId,omitempty"`

	// Negated is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ratebasedrule-predicate.html#cfn-wafregional-ratebasedrule-predicate-negated
	Negated cfz.Expression[bool] `json:"Negated,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ratebasedrule-predicate.html#cfn-wafregional-ratebasedrule-predicate-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_WAFRegional_RateBasedRule_Predicate initializes a new AWS_WAFRegional_RateBasedRule_Predicate.
func New__AWS_WAFRegional_RateBasedRule_Predicate() AWS_WAFRegional_RateBasedRule_Predicate {
	return AWS_WAFRegional_RateBasedRule_Predicate{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFRegional_RateBasedRule_Predicate) GetType() string {
	return AWS_WAFRegional_RateBasedRule_Predicate__Type
}

// Set__DataId updates property "DataId".
func (t AWS_WAFRegional_RateBasedRule_Predicate) Set__DataId(v cfz.Expression[string]) AWS_WAFRegional_RateBasedRule_Predicate {
	t.DataId = v
	return t
}

// SetV__DataId updates property "DataId".
func (t AWS_WAFRegional_RateBasedRule_Predicate) SetV__DataId(v string) AWS_WAFRegional_RateBasedRule_Predicate {
	t.DataId = cfz.V(v)
	return t
}

// Set__Negated updates property "Negated".
func (t AWS_WAFRegional_RateBasedRule_Predicate) Set__Negated(v cfz.Expression[bool]) AWS_WAFRegional_RateBasedRule_Predicate {
	t.Negated = v
	return t
}

// SetV__Negated updates property "Negated".
func (t AWS_WAFRegional_RateBasedRule_Predicate) SetV__Negated(v bool) AWS_WAFRegional_RateBasedRule_Predicate {
	t.Negated = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_WAFRegional_RateBasedRule_Predicate) Set__Type(v cfz.Expression[string]) AWS_WAFRegional_RateBasedRule_Predicate {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_WAFRegional_RateBasedRule_Predicate) SetV__Type(v string) AWS_WAFRegional_RateBasedRule_Predicate {
	t.Type = cfz.V(v)
	return t
}
