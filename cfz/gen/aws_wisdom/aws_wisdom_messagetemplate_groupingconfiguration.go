// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wisdom

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Wisdom_MessageTemplate_GroupingConfiguration__Type is the CloudFormation type for AWS::Wisdom::MessageTemplate.GroupingConfiguration.
	AWS_Wisdom_MessageTemplate_GroupingConfiguration__Type = "AWS::Wisdom::MessageTemplate.GroupingConfiguration"
)

var (
	// AWS_Wisdom_MessageTemplate_GroupingConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Wisdom::MessageTemplate.GroupingConfiguration.
	AWS_Wisdom_MessageTemplate_GroupingConfiguration__PropertiesMap = struct {
		Criteria string
		Values   string
	}{
		Criteria: "Criteria",
		Values:   "Values",
	}

	// AWS_Wisdom_MessageTemplate_GroupingConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Wisdom::MessageTemplate.GroupingConfiguration.
	AWS_Wisdom_MessageTemplate_GroupingConfiguration__PropertiesSlice = []string{
		AWS_Wisdom_MessageTemplate_GroupingConfiguration__PropertiesMap.Criteria,
		AWS_Wisdom_MessageTemplate_GroupingConfiguration__PropertiesMap.Values,
	}
)

// AWS_Wisdom_MessageTemplate_GroupingConfiguration is a binding for AWS::Wisdom::MessageTemplate.GroupingConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-groupingconfiguration.html
type AWS_Wisdom_MessageTemplate_GroupingConfiguration struct {
	// Criteria is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-groupingconfiguration.html#cfn-wisdom-messagetemplate-groupingconfiguration-criteria
	Criteria cfz.Expression[string] `json:"Criteria,omitempty"`

	// Values is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-groupingconfiguration.html#cfn-wisdom-messagetemplate-groupingconfiguration-values
	Values cfz.ExpressionSlice[string] `json:"Values,omitempty"`
}

// New__AWS_Wisdom_MessageTemplate_GroupingConfiguration initializes a new AWS_Wisdom_MessageTemplate_GroupingConfiguration.
func New__AWS_Wisdom_MessageTemplate_GroupingConfiguration() AWS_Wisdom_MessageTemplate_GroupingConfiguration {
	return AWS_Wisdom_MessageTemplate_GroupingConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Wisdom_MessageTemplate_GroupingConfiguration) GetType() string {
	return AWS_Wisdom_MessageTemplate_GroupingConfiguration__Type
}

// Set__Criteria updates property "Criteria".
func (t AWS_Wisdom_MessageTemplate_GroupingConfiguration) Set__Criteria(v cfz.Expression[string]) AWS_Wisdom_MessageTemplate_GroupingConfiguration {
	t.Criteria = v
	return t
}

// SetV__Criteria updates property "Criteria".
func (t AWS_Wisdom_MessageTemplate_GroupingConfiguration) SetV__Criteria(v string) AWS_Wisdom_MessageTemplate_GroupingConfiguration {
	t.Criteria = cfz.V(v)
	return t
}

// Set__Values updates property "Values".
func (t AWS_Wisdom_MessageTemplate_GroupingConfiguration) Set__Values(v cfz.ExpressionSlice[string]) AWS_Wisdom_MessageTemplate_GroupingConfiguration {
	t.Values = v
	return t
}

// SetS__Values updates property "Values".
func (t AWS_Wisdom_MessageTemplate_GroupingConfiguration) SetS__Values(v ...cfz.Expression[string]) AWS_Wisdom_MessageTemplate_GroupingConfiguration {
	t.Values = cfz.S(v...)
	return t
}

// SetSV__Values updates property "Values".
func (t AWS_Wisdom_MessageTemplate_GroupingConfiguration) SetSV__Values(v ...string) AWS_Wisdom_MessageTemplate_GroupingConfiguration {
	t.Values = cfz.SV(v...)
	return t
}
