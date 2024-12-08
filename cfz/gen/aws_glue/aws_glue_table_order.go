// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_glue

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Glue_Table_Order__Type is the CloudFormation type for AWS::Glue::Table.Order.
	AWS_Glue_Table_Order__Type = "AWS::Glue::Table.Order"
)

var (
	// AWS_Glue_Table_Order__PropertiesMap reports all the CloudFormation properties for AWS::Glue::Table.Order.
	AWS_Glue_Table_Order__PropertiesMap = struct {
		Column    string
		SortOrder string
	}{
		Column:    "Column",
		SortOrder: "SortOrder",
	}

	// AWS_Glue_Table_Order__PropertiesSlice reports all the CloudFormation properties for AWS::Glue::Table.Order.
	AWS_Glue_Table_Order__PropertiesSlice = []string{
		AWS_Glue_Table_Order__PropertiesMap.Column,
		AWS_Glue_Table_Order__PropertiesMap.SortOrder,
	}
)

// AWS_Glue_Table_Order is a binding for AWS::Glue::Table.Order.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-order.html
type AWS_Glue_Table_Order struct {
	// Column is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-order.html#cfn-glue-table-order-column
	Column cfz.Expression[string] `json:"Column,omitempty"`

	// SortOrder is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-order.html#cfn-glue-table-order-sortorder
	SortOrder cfz.Expression[int32] `json:"SortOrder,omitempty"`
}

// New__AWS_Glue_Table_Order initializes a new AWS_Glue_Table_Order.
func New__AWS_Glue_Table_Order() AWS_Glue_Table_Order {
	return AWS_Glue_Table_Order{}
}

// GetType returns the CloudFormation type.
func (AWS_Glue_Table_Order) GetType() string {
	return AWS_Glue_Table_Order__Type
}

// Set__Column updates property "Column".
func (t AWS_Glue_Table_Order) Set__Column(v cfz.Expression[string]) AWS_Glue_Table_Order {
	t.Column = v
	return t
}

// SetV__Column updates property "Column".
func (t AWS_Glue_Table_Order) SetV__Column(v string) AWS_Glue_Table_Order {
	t.Column = cfz.V(v)
	return t
}

// Set__SortOrder updates property "SortOrder".
func (t AWS_Glue_Table_Order) Set__SortOrder(v cfz.Expression[int32]) AWS_Glue_Table_Order {
	t.SortOrder = v
	return t
}

// SetV__SortOrder updates property "SortOrder".
func (t AWS_Glue_Table_Order) SetV__SortOrder(v int32) AWS_Glue_Table_Order {
	t.SortOrder = cfz.V(v)
	return t
}
