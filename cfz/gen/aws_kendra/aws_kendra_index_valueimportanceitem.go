// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kendra

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Kendra_Index_ValueImportanceItem__Type is the CloudFormation type for AWS::Kendra::Index.ValueImportanceItem.
	AWS_Kendra_Index_ValueImportanceItem__Type = "AWS::Kendra::Index.ValueImportanceItem"
)

var (
	// AWS_Kendra_Index_ValueImportanceItem__PropertiesMap reports all the CloudFormation properties for AWS::Kendra::Index.ValueImportanceItem.
	AWS_Kendra_Index_ValueImportanceItem__PropertiesMap = struct {
		Key   string
		Value string
	}{
		Key:   "Key",
		Value: "Value",
	}

	// AWS_Kendra_Index_ValueImportanceItem__PropertiesSlice reports all the CloudFormation properties for AWS::Kendra::Index.ValueImportanceItem.
	AWS_Kendra_Index_ValueImportanceItem__PropertiesSlice = []string{
		AWS_Kendra_Index_ValueImportanceItem__PropertiesMap.Key,
		AWS_Kendra_Index_ValueImportanceItem__PropertiesMap.Value,
	}
)

// AWS_Kendra_Index_ValueImportanceItem is a binding for AWS::Kendra::Index.ValueImportanceItem.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-valueimportanceitem.html
type AWS_Kendra_Index_ValueImportanceItem struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-valueimportanceitem.html#cfn-kendra-index-valueimportanceitem-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-valueimportanceitem.html#cfn-kendra-index-valueimportanceitem-value
	Value cfz.Expression[int32] `json:"Value,omitempty"`
}

// New__AWS_Kendra_Index_ValueImportanceItem initializes a new AWS_Kendra_Index_ValueImportanceItem.
func New__AWS_Kendra_Index_ValueImportanceItem() AWS_Kendra_Index_ValueImportanceItem {
	return AWS_Kendra_Index_ValueImportanceItem{}
}

// GetType returns the CloudFormation type.
func (AWS_Kendra_Index_ValueImportanceItem) GetType() string {
	return AWS_Kendra_Index_ValueImportanceItem__Type
}

// Set__Key updates property "Key".
func (t AWS_Kendra_Index_ValueImportanceItem) Set__Key(v cfz.Expression[string]) AWS_Kendra_Index_ValueImportanceItem {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_Kendra_Index_ValueImportanceItem) SetV__Key(v string) AWS_Kendra_Index_ValueImportanceItem {
	t.Key = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_Kendra_Index_ValueImportanceItem) Set__Value(v cfz.Expression[int32]) AWS_Kendra_Index_ValueImportanceItem {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_Kendra_Index_ValueImportanceItem) SetV__Value(v int32) AWS_Kendra_Index_ValueImportanceItem {
	t.Value = cfz.V(v)
	return t
}
