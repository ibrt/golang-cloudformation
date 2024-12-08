// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_qbusiness

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QBusiness_DataSource_DocumentAttributeTarget__Type is the CloudFormation type for AWS::QBusiness::DataSource.DocumentAttributeTarget.
	AWS_QBusiness_DataSource_DocumentAttributeTarget__Type = "AWS::QBusiness::DataSource.DocumentAttributeTarget"
)

var (
	// AWS_QBusiness_DataSource_DocumentAttributeTarget__PropertiesMap reports all the CloudFormation properties for AWS::QBusiness::DataSource.DocumentAttributeTarget.
	AWS_QBusiness_DataSource_DocumentAttributeTarget__PropertiesMap = struct {
		AttributeValueOperator string
		Key                    string
		Value                  string
	}{
		AttributeValueOperator: "AttributeValueOperator",
		Key:                    "Key",
		Value:                  "Value",
	}

	// AWS_QBusiness_DataSource_DocumentAttributeTarget__PropertiesSlice reports all the CloudFormation properties for AWS::QBusiness::DataSource.DocumentAttributeTarget.
	AWS_QBusiness_DataSource_DocumentAttributeTarget__PropertiesSlice = []string{
		AWS_QBusiness_DataSource_DocumentAttributeTarget__PropertiesMap.AttributeValueOperator,
		AWS_QBusiness_DataSource_DocumentAttributeTarget__PropertiesMap.Key,
		AWS_QBusiness_DataSource_DocumentAttributeTarget__PropertiesMap.Value,
	}
)

// AWS_QBusiness_DataSource_DocumentAttributeTarget is a binding for AWS::QBusiness::DataSource.DocumentAttributeTarget.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributetarget.html
type AWS_QBusiness_DataSource_DocumentAttributeTarget struct {
	// AttributeValueOperator is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributetarget.html#cfn-qbusiness-datasource-documentattributetarget-attributevalueoperator
	AttributeValueOperator cfz.Expression[string] `json:"AttributeValueOperator,omitempty"`

	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributetarget.html#cfn-qbusiness-datasource-documentattributetarget-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributetarget.html#cfn-qbusiness-datasource-documentattributetarget-value
	Value cfz.Expression[AWS_QBusiness_DataSource_DocumentAttributeValue] `json:"Value,omitempty"`
}

// New__AWS_QBusiness_DataSource_DocumentAttributeTarget initializes a new AWS_QBusiness_DataSource_DocumentAttributeTarget.
func New__AWS_QBusiness_DataSource_DocumentAttributeTarget() AWS_QBusiness_DataSource_DocumentAttributeTarget {
	return AWS_QBusiness_DataSource_DocumentAttributeTarget{}
}

// GetType returns the CloudFormation type.
func (AWS_QBusiness_DataSource_DocumentAttributeTarget) GetType() string {
	return AWS_QBusiness_DataSource_DocumentAttributeTarget__Type
}

// Set__AttributeValueOperator updates property "AttributeValueOperator".
func (t AWS_QBusiness_DataSource_DocumentAttributeTarget) Set__AttributeValueOperator(v cfz.Expression[string]) AWS_QBusiness_DataSource_DocumentAttributeTarget {
	t.AttributeValueOperator = v
	return t
}

// SetV__AttributeValueOperator updates property "AttributeValueOperator".
func (t AWS_QBusiness_DataSource_DocumentAttributeTarget) SetV__AttributeValueOperator(v string) AWS_QBusiness_DataSource_DocumentAttributeTarget {
	t.AttributeValueOperator = cfz.V(v)
	return t
}

// Set__Key updates property "Key".
func (t AWS_QBusiness_DataSource_DocumentAttributeTarget) Set__Key(v cfz.Expression[string]) AWS_QBusiness_DataSource_DocumentAttributeTarget {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_QBusiness_DataSource_DocumentAttributeTarget) SetV__Key(v string) AWS_QBusiness_DataSource_DocumentAttributeTarget {
	t.Key = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_QBusiness_DataSource_DocumentAttributeTarget) Set__Value(v cfz.Expression[AWS_QBusiness_DataSource_DocumentAttributeValue]) AWS_QBusiness_DataSource_DocumentAttributeTarget {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_QBusiness_DataSource_DocumentAttributeTarget) SetV__Value(v AWS_QBusiness_DataSource_DocumentAttributeValue) AWS_QBusiness_DataSource_DocumentAttributeTarget {
	t.Value = cfz.V(v)
	return t
}
