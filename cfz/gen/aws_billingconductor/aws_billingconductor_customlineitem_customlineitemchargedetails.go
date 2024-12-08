// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_billingconductor

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__Type is the CloudFormation type for AWS::BillingConductor::CustomLineItem.CustomLineItemChargeDetails.
	AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__Type = "AWS::BillingConductor::CustomLineItem.CustomLineItemChargeDetails"
)

var (
	// AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__PropertiesMap reports all the CloudFormation properties for AWS::BillingConductor::CustomLineItem.CustomLineItemChargeDetails.
	AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__PropertiesMap = struct {
		Flat            string
		LineItemFilters string
		Percentage      string
		Type            string
	}{
		Flat:            "Flat",
		LineItemFilters: "LineItemFilters",
		Percentage:      "Percentage",
		Type:            "Type",
	}

	// AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__PropertiesSlice reports all the CloudFormation properties for AWS::BillingConductor::CustomLineItem.CustomLineItemChargeDetails.
	AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__PropertiesSlice = []string{
		AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__PropertiesMap.Flat,
		AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__PropertiesMap.LineItemFilters,
		AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__PropertiesMap.Percentage,
		AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__PropertiesMap.Type,
	}
)

// AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails is a binding for AWS::BillingConductor::CustomLineItem.CustomLineItemChargeDetails.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html
type AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails struct {
	// Flat is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html#cfn-billingconductor-customlineitem-customlineitemchargedetails-flat
	Flat cfz.Expression[AWS_BillingConductor_CustomLineItem_CustomLineItemFlatChargeDetails] `json:"Flat,omitempty"`

	// LineItemFilters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html#cfn-billingconductor-customlineitem-customlineitemchargedetails-lineitemfilters
	LineItemFilters cfz.ExpressionSlice[AWS_BillingConductor_CustomLineItem_LineItemFilter] `json:"LineItemFilters,omitempty"`

	// Percentage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html#cfn-billingconductor-customlineitem-customlineitemchargedetails-percentage
	Percentage cfz.Expression[AWS_BillingConductor_CustomLineItem_CustomLineItemPercentageChargeDetails] `json:"Percentage,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html#cfn-billingconductor-customlineitem-customlineitemchargedetails-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails initializes a new AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails.
func New__AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails() AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	return AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails{}
}

// GetType returns the CloudFormation type.
func (AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) GetType() string {
	return AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails__Type
}

// Set__Flat updates property "Flat".
func (t AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) Set__Flat(v cfz.Expression[AWS_BillingConductor_CustomLineItem_CustomLineItemFlatChargeDetails]) AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	t.Flat = v
	return t
}

// SetV__Flat updates property "Flat".
func (t AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) SetV__Flat(v AWS_BillingConductor_CustomLineItem_CustomLineItemFlatChargeDetails) AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	t.Flat = cfz.V(v)
	return t
}

// Set__LineItemFilters updates property "LineItemFilters".
func (t AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) Set__LineItemFilters(v cfz.ExpressionSlice[AWS_BillingConductor_CustomLineItem_LineItemFilter]) AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	t.LineItemFilters = v
	return t
}

// SetS__LineItemFilters updates property "LineItemFilters".
func (t AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) SetS__LineItemFilters(v ...cfz.Expression[AWS_BillingConductor_CustomLineItem_LineItemFilter]) AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	t.LineItemFilters = cfz.S(v...)
	return t
}

// SetSV__LineItemFilters updates property "LineItemFilters".
func (t AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) SetSV__LineItemFilters(v ...AWS_BillingConductor_CustomLineItem_LineItemFilter) AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	t.LineItemFilters = cfz.SV(v...)
	return t
}

// Set__Percentage updates property "Percentage".
func (t AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) Set__Percentage(v cfz.Expression[AWS_BillingConductor_CustomLineItem_CustomLineItemPercentageChargeDetails]) AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	t.Percentage = v
	return t
}

// SetV__Percentage updates property "Percentage".
func (t AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) SetV__Percentage(v AWS_BillingConductor_CustomLineItem_CustomLineItemPercentageChargeDetails) AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	t.Percentage = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) Set__Type(v cfz.Expression[string]) AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails) SetV__Type(v string) AWS_BillingConductor_CustomLineItem_CustomLineItemChargeDetails {
	t.Type = cfz.V(v)
	return t
}
