// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudfront

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFront_Distribution_OriginGroups__Type is the CloudFormation type for AWS::CloudFront::Distribution.OriginGroups.
	AWS_CloudFront_Distribution_OriginGroups__Type = "AWS::CloudFront::Distribution.OriginGroups"
)

var (
	// AWS_CloudFront_Distribution_OriginGroups__PropertiesMap reports all the CloudFormation properties for AWS::CloudFront::Distribution.OriginGroups.
	AWS_CloudFront_Distribution_OriginGroups__PropertiesMap = struct {
		Items    string
		Quantity string
	}{
		Items:    "Items",
		Quantity: "Quantity",
	}

	// AWS_CloudFront_Distribution_OriginGroups__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFront::Distribution.OriginGroups.
	AWS_CloudFront_Distribution_OriginGroups__PropertiesSlice = []string{
		AWS_CloudFront_Distribution_OriginGroups__PropertiesMap.Items,
		AWS_CloudFront_Distribution_OriginGroups__PropertiesMap.Quantity,
	}
)

// AWS_CloudFront_Distribution_OriginGroups is a binding for AWS::CloudFront::Distribution.OriginGroups.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroups.html
type AWS_CloudFront_Distribution_OriginGroups struct {
	// Items is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroups.html#cfn-cloudfront-distribution-origingroups-items
	Items cfz.ExpressionSlice[AWS_CloudFront_Distribution_OriginGroup] `json:"Items,omitempty"`

	// Quantity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroups.html#cfn-cloudfront-distribution-origingroups-quantity
	Quantity cfz.Expression[int32] `json:"Quantity,omitempty"`
}

// New__AWS_CloudFront_Distribution_OriginGroups initializes a new AWS_CloudFront_Distribution_OriginGroups.
func New__AWS_CloudFront_Distribution_OriginGroups() AWS_CloudFront_Distribution_OriginGroups {
	return AWS_CloudFront_Distribution_OriginGroups{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFront_Distribution_OriginGroups) GetType() string {
	return AWS_CloudFront_Distribution_OriginGroups__Type
}

// Set__Items updates property "Items".
func (t AWS_CloudFront_Distribution_OriginGroups) Set__Items(v cfz.ExpressionSlice[AWS_CloudFront_Distribution_OriginGroup]) AWS_CloudFront_Distribution_OriginGroups {
	t.Items = v
	return t
}

// SetS__Items updates property "Items".
func (t AWS_CloudFront_Distribution_OriginGroups) SetS__Items(v ...cfz.Expression[AWS_CloudFront_Distribution_OriginGroup]) AWS_CloudFront_Distribution_OriginGroups {
	t.Items = cfz.S(v...)
	return t
}

// SetSV__Items updates property "Items".
func (t AWS_CloudFront_Distribution_OriginGroups) SetSV__Items(v ...AWS_CloudFront_Distribution_OriginGroup) AWS_CloudFront_Distribution_OriginGroups {
	t.Items = cfz.SV(v...)
	return t
}

// Set__Quantity updates property "Quantity".
func (t AWS_CloudFront_Distribution_OriginGroups) Set__Quantity(v cfz.Expression[int32]) AWS_CloudFront_Distribution_OriginGroups {
	t.Quantity = v
	return t
}

// SetV__Quantity updates property "Quantity".
func (t AWS_CloudFront_Distribution_OriginGroups) SetV__Quantity(v int32) AWS_CloudFront_Distribution_OriginGroups {
	t.Quantity = cfz.V(v)
	return t
}
