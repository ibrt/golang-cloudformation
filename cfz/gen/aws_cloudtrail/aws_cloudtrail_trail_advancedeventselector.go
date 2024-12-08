// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudtrail

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudTrail_Trail_AdvancedEventSelector__Type is the CloudFormation type for AWS::CloudTrail::Trail.AdvancedEventSelector.
	AWS_CloudTrail_Trail_AdvancedEventSelector__Type = "AWS::CloudTrail::Trail.AdvancedEventSelector"
)

var (
	// AWS_CloudTrail_Trail_AdvancedEventSelector__PropertiesMap reports all the CloudFormation properties for AWS::CloudTrail::Trail.AdvancedEventSelector.
	AWS_CloudTrail_Trail_AdvancedEventSelector__PropertiesMap = struct {
		FieldSelectors string
		Name           string
	}{
		FieldSelectors: "FieldSelectors",
		Name:           "Name",
	}

	// AWS_CloudTrail_Trail_AdvancedEventSelector__PropertiesSlice reports all the CloudFormation properties for AWS::CloudTrail::Trail.AdvancedEventSelector.
	AWS_CloudTrail_Trail_AdvancedEventSelector__PropertiesSlice = []string{
		AWS_CloudTrail_Trail_AdvancedEventSelector__PropertiesMap.FieldSelectors,
		AWS_CloudTrail_Trail_AdvancedEventSelector__PropertiesMap.Name,
	}
)

// AWS_CloudTrail_Trail_AdvancedEventSelector is a binding for AWS::CloudTrail::Trail.AdvancedEventSelector.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedeventselector.html
type AWS_CloudTrail_Trail_AdvancedEventSelector struct {
	// FieldSelectors is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedeventselector.html#cfn-cloudtrail-trail-advancedeventselector-fieldselectors
	FieldSelectors cfz.ExpressionSlice[AWS_CloudTrail_Trail_AdvancedFieldSelector] `json:"FieldSelectors,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedeventselector.html#cfn-cloudtrail-trail-advancedeventselector-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_CloudTrail_Trail_AdvancedEventSelector initializes a new AWS_CloudTrail_Trail_AdvancedEventSelector.
func New__AWS_CloudTrail_Trail_AdvancedEventSelector() AWS_CloudTrail_Trail_AdvancedEventSelector {
	return AWS_CloudTrail_Trail_AdvancedEventSelector{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudTrail_Trail_AdvancedEventSelector) GetType() string {
	return AWS_CloudTrail_Trail_AdvancedEventSelector__Type
}

// Set__FieldSelectors updates property "FieldSelectors".
func (t AWS_CloudTrail_Trail_AdvancedEventSelector) Set__FieldSelectors(v cfz.ExpressionSlice[AWS_CloudTrail_Trail_AdvancedFieldSelector]) AWS_CloudTrail_Trail_AdvancedEventSelector {
	t.FieldSelectors = v
	return t
}

// SetS__FieldSelectors updates property "FieldSelectors".
func (t AWS_CloudTrail_Trail_AdvancedEventSelector) SetS__FieldSelectors(v ...cfz.Expression[AWS_CloudTrail_Trail_AdvancedFieldSelector]) AWS_CloudTrail_Trail_AdvancedEventSelector {
	t.FieldSelectors = cfz.S(v...)
	return t
}

// SetSV__FieldSelectors updates property "FieldSelectors".
func (t AWS_CloudTrail_Trail_AdvancedEventSelector) SetSV__FieldSelectors(v ...AWS_CloudTrail_Trail_AdvancedFieldSelector) AWS_CloudTrail_Trail_AdvancedEventSelector {
	t.FieldSelectors = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t AWS_CloudTrail_Trail_AdvancedEventSelector) Set__Name(v cfz.Expression[string]) AWS_CloudTrail_Trail_AdvancedEventSelector {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_CloudTrail_Trail_AdvancedEventSelector) SetV__Name(v string) AWS_CloudTrail_Trail_AdvancedEventSelector {
	t.Name = cfz.V(v)
	return t
}
