// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_securityhub

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SecurityHub_Insight_NumberFilter__Type is the CloudFormation type for AWS::SecurityHub::Insight.NumberFilter.
	AWS_SecurityHub_Insight_NumberFilter__Type = "AWS::SecurityHub::Insight.NumberFilter"
)

var (
	// AWS_SecurityHub_Insight_NumberFilter__PropertiesMap reports all the CloudFormation properties for AWS::SecurityHub::Insight.NumberFilter.
	AWS_SecurityHub_Insight_NumberFilter__PropertiesMap = struct {
		Eq  string
		Gte string
		Lte string
	}{
		Eq:  "Eq",
		Gte: "Gte",
		Lte: "Lte",
	}

	// AWS_SecurityHub_Insight_NumberFilter__PropertiesSlice reports all the CloudFormation properties for AWS::SecurityHub::Insight.NumberFilter.
	AWS_SecurityHub_Insight_NumberFilter__PropertiesSlice = []string{
		AWS_SecurityHub_Insight_NumberFilter__PropertiesMap.Eq,
		AWS_SecurityHub_Insight_NumberFilter__PropertiesMap.Gte,
		AWS_SecurityHub_Insight_NumberFilter__PropertiesMap.Lte,
	}
)

// AWS_SecurityHub_Insight_NumberFilter is a binding for AWS::SecurityHub::Insight.NumberFilter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-numberfilter.html
type AWS_SecurityHub_Insight_NumberFilter struct {
	// Eq is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-numberfilter.html#cfn-securityhub-insight-numberfilter-eq
	Eq cfz.Expression[float64] `json:"Eq,omitempty"`

	// Gte is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-numberfilter.html#cfn-securityhub-insight-numberfilter-gte
	Gte cfz.Expression[float64] `json:"Gte,omitempty"`

	// Lte is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-numberfilter.html#cfn-securityhub-insight-numberfilter-lte
	Lte cfz.Expression[float64] `json:"Lte,omitempty"`
}

// New__AWS_SecurityHub_Insight_NumberFilter initializes a new AWS_SecurityHub_Insight_NumberFilter.
func New__AWS_SecurityHub_Insight_NumberFilter() AWS_SecurityHub_Insight_NumberFilter {
	return AWS_SecurityHub_Insight_NumberFilter{}
}

// GetType returns the CloudFormation type.
func (AWS_SecurityHub_Insight_NumberFilter) GetType() string {
	return AWS_SecurityHub_Insight_NumberFilter__Type
}

// Set__Eq updates property "Eq".
func (t AWS_SecurityHub_Insight_NumberFilter) Set__Eq(v cfz.Expression[float64]) AWS_SecurityHub_Insight_NumberFilter {
	t.Eq = v
	return t
}

// SetV__Eq updates property "Eq".
func (t AWS_SecurityHub_Insight_NumberFilter) SetV__Eq(v float64) AWS_SecurityHub_Insight_NumberFilter {
	t.Eq = cfz.V(v)
	return t
}

// Set__Gte updates property "Gte".
func (t AWS_SecurityHub_Insight_NumberFilter) Set__Gte(v cfz.Expression[float64]) AWS_SecurityHub_Insight_NumberFilter {
	t.Gte = v
	return t
}

// SetV__Gte updates property "Gte".
func (t AWS_SecurityHub_Insight_NumberFilter) SetV__Gte(v float64) AWS_SecurityHub_Insight_NumberFilter {
	t.Gte = cfz.V(v)
	return t
}

// Set__Lte updates property "Lte".
func (t AWS_SecurityHub_Insight_NumberFilter) Set__Lte(v cfz.Expression[float64]) AWS_SecurityHub_Insight_NumberFilter {
	t.Lte = v
	return t
}

// SetV__Lte updates property "Lte".
func (t AWS_SecurityHub_Insight_NumberFilter) SetV__Lte(v float64) AWS_SecurityHub_Insight_NumberFilter {
	t.Lte = cfz.V(v)
	return t
}
