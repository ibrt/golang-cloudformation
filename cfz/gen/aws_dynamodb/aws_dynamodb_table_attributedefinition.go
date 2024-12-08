// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dynamodb

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DynamoDB_Table_AttributeDefinition__Type is the CloudFormation type for AWS::DynamoDB::Table.AttributeDefinition.
	AWS_DynamoDB_Table_AttributeDefinition__Type = "AWS::DynamoDB::Table.AttributeDefinition"
)

var (
	// AWS_DynamoDB_Table_AttributeDefinition__PropertiesMap reports all the CloudFormation properties for AWS::DynamoDB::Table.AttributeDefinition.
	AWS_DynamoDB_Table_AttributeDefinition__PropertiesMap = struct {
		AttributeName string
		AttributeType string
	}{
		AttributeName: "AttributeName",
		AttributeType: "AttributeType",
	}

	// AWS_DynamoDB_Table_AttributeDefinition__PropertiesSlice reports all the CloudFormation properties for AWS::DynamoDB::Table.AttributeDefinition.
	AWS_DynamoDB_Table_AttributeDefinition__PropertiesSlice = []string{
		AWS_DynamoDB_Table_AttributeDefinition__PropertiesMap.AttributeName,
		AWS_DynamoDB_Table_AttributeDefinition__PropertiesMap.AttributeType,
	}
)

// AWS_DynamoDB_Table_AttributeDefinition is a binding for AWS::DynamoDB::Table.AttributeDefinition.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-attributedefinition.html
type AWS_DynamoDB_Table_AttributeDefinition struct {
	// AttributeName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-attributedefinition.html#cfn-dynamodb-table-attributedefinition-attributename
	AttributeName cfz.Expression[string] `json:"AttributeName,omitempty"`

	// AttributeType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-attributedefinition.html#cfn-dynamodb-table-attributedefinition-attributetype
	AttributeType cfz.Expression[string] `json:"AttributeType,omitempty"`
}

// New__AWS_DynamoDB_Table_AttributeDefinition initializes a new AWS_DynamoDB_Table_AttributeDefinition.
func New__AWS_DynamoDB_Table_AttributeDefinition() AWS_DynamoDB_Table_AttributeDefinition {
	return AWS_DynamoDB_Table_AttributeDefinition{}
}

// GetType returns the CloudFormation type.
func (AWS_DynamoDB_Table_AttributeDefinition) GetType() string {
	return AWS_DynamoDB_Table_AttributeDefinition__Type
}

// Set__AttributeName updates property "AttributeName".
func (t AWS_DynamoDB_Table_AttributeDefinition) Set__AttributeName(v cfz.Expression[string]) AWS_DynamoDB_Table_AttributeDefinition {
	t.AttributeName = v
	return t
}

// SetV__AttributeName updates property "AttributeName".
func (t AWS_DynamoDB_Table_AttributeDefinition) SetV__AttributeName(v string) AWS_DynamoDB_Table_AttributeDefinition {
	t.AttributeName = cfz.V(v)
	return t
}

// Set__AttributeType updates property "AttributeType".
func (t AWS_DynamoDB_Table_AttributeDefinition) Set__AttributeType(v cfz.Expression[string]) AWS_DynamoDB_Table_AttributeDefinition {
	t.AttributeType = v
	return t
}

// SetV__AttributeType updates property "AttributeType".
func (t AWS_DynamoDB_Table_AttributeDefinition) SetV__AttributeType(v string) AWS_DynamoDB_Table_AttributeDefinition {
	t.AttributeType = cfz.V(v)
	return t
}
