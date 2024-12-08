// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_qbusiness

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QBusiness_DataSource_DocumentAttributeCondition__Type is the CloudFormation type for AWS::QBusiness::DataSource.DocumentAttributeCondition.
	AWS_QBusiness_DataSource_DocumentAttributeCondition__Type = "AWS::QBusiness::DataSource.DocumentAttributeCondition"
)

var (
	// AWS_QBusiness_DataSource_DocumentAttributeCondition__PropertiesMap reports all the CloudFormation properties for AWS::QBusiness::DataSource.DocumentAttributeCondition.
	AWS_QBusiness_DataSource_DocumentAttributeCondition__PropertiesMap = struct {
		Key      string
		Operator string
		Value    string
	}{
		Key:      "Key",
		Operator: "Operator",
		Value:    "Value",
	}

	// AWS_QBusiness_DataSource_DocumentAttributeCondition__PropertiesSlice reports all the CloudFormation properties for AWS::QBusiness::DataSource.DocumentAttributeCondition.
	AWS_QBusiness_DataSource_DocumentAttributeCondition__PropertiesSlice = []string{
		AWS_QBusiness_DataSource_DocumentAttributeCondition__PropertiesMap.Key,
		AWS_QBusiness_DataSource_DocumentAttributeCondition__PropertiesMap.Operator,
		AWS_QBusiness_DataSource_DocumentAttributeCondition__PropertiesMap.Value,
	}
)

// AWS_QBusiness_DataSource_DocumentAttributeCondition is a binding for AWS::QBusiness::DataSource.DocumentAttributeCondition.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributecondition.html
type AWS_QBusiness_DataSource_DocumentAttributeCondition struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributecondition.html#cfn-qbusiness-datasource-documentattributecondition-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Operator is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributecondition.html#cfn-qbusiness-datasource-documentattributecondition-operator
	Operator cfz.Expression[string] `json:"Operator,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributecondition.html#cfn-qbusiness-datasource-documentattributecondition-value
	Value cfz.Expression[AWS_QBusiness_DataSource_DocumentAttributeValue] `json:"Value,omitempty"`
}

// New__AWS_QBusiness_DataSource_DocumentAttributeCondition initializes a new AWS_QBusiness_DataSource_DocumentAttributeCondition.
func New__AWS_QBusiness_DataSource_DocumentAttributeCondition() AWS_QBusiness_DataSource_DocumentAttributeCondition {
	return AWS_QBusiness_DataSource_DocumentAttributeCondition{}
}

// GetType returns the CloudFormation type.
func (AWS_QBusiness_DataSource_DocumentAttributeCondition) GetType() string {
	return AWS_QBusiness_DataSource_DocumentAttributeCondition__Type
}

// Set__Key updates property "Key".
func (t AWS_QBusiness_DataSource_DocumentAttributeCondition) Set__Key(v cfz.Expression[string]) AWS_QBusiness_DataSource_DocumentAttributeCondition {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_QBusiness_DataSource_DocumentAttributeCondition) SetV__Key(v string) AWS_QBusiness_DataSource_DocumentAttributeCondition {
	t.Key = cfz.V(v)
	return t
}

// Set__Operator updates property "Operator".
func (t AWS_QBusiness_DataSource_DocumentAttributeCondition) Set__Operator(v cfz.Expression[string]) AWS_QBusiness_DataSource_DocumentAttributeCondition {
	t.Operator = v
	return t
}

// SetV__Operator updates property "Operator".
func (t AWS_QBusiness_DataSource_DocumentAttributeCondition) SetV__Operator(v string) AWS_QBusiness_DataSource_DocumentAttributeCondition {
	t.Operator = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_QBusiness_DataSource_DocumentAttributeCondition) Set__Value(v cfz.Expression[AWS_QBusiness_DataSource_DocumentAttributeValue]) AWS_QBusiness_DataSource_DocumentAttributeCondition {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_QBusiness_DataSource_DocumentAttributeCondition) SetV__Value(v AWS_QBusiness_DataSource_DocumentAttributeValue) AWS_QBusiness_DataSource_DocumentAttributeCondition {
	t.Value = cfz.V(v)
	return t
}
