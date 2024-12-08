// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dynamodb

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DynamoDB_Table_InputFormatOptions__Type is the CloudFormation type for AWS::DynamoDB::Table.InputFormatOptions.
	AWS_DynamoDB_Table_InputFormatOptions__Type = "AWS::DynamoDB::Table.InputFormatOptions"
)

var (
	// AWS_DynamoDB_Table_InputFormatOptions__PropertiesMap reports all the CloudFormation properties for AWS::DynamoDB::Table.InputFormatOptions.
	AWS_DynamoDB_Table_InputFormatOptions__PropertiesMap = struct {
		Csv string
	}{
		Csv: "Csv",
	}

	// AWS_DynamoDB_Table_InputFormatOptions__PropertiesSlice reports all the CloudFormation properties for AWS::DynamoDB::Table.InputFormatOptions.
	AWS_DynamoDB_Table_InputFormatOptions__PropertiesSlice = []string{
		AWS_DynamoDB_Table_InputFormatOptions__PropertiesMap.Csv,
	}
)

// AWS_DynamoDB_Table_InputFormatOptions is a binding for AWS::DynamoDB::Table.InputFormatOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-inputformatoptions.html
type AWS_DynamoDB_Table_InputFormatOptions struct {
	// Csv is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-inputformatoptions.html#cfn-dynamodb-table-inputformatoptions-csv
	Csv cfz.Expression[AWS_DynamoDB_Table_Csv] `json:"Csv,omitempty"`
}

// New__AWS_DynamoDB_Table_InputFormatOptions initializes a new AWS_DynamoDB_Table_InputFormatOptions.
func New__AWS_DynamoDB_Table_InputFormatOptions() AWS_DynamoDB_Table_InputFormatOptions {
	return AWS_DynamoDB_Table_InputFormatOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_DynamoDB_Table_InputFormatOptions) GetType() string {
	return AWS_DynamoDB_Table_InputFormatOptions__Type
}

// Set__Csv updates property "Csv".
func (t AWS_DynamoDB_Table_InputFormatOptions) Set__Csv(v cfz.Expression[AWS_DynamoDB_Table_Csv]) AWS_DynamoDB_Table_InputFormatOptions {
	t.Csv = v
	return t
}

// SetV__Csv updates property "Csv".
func (t AWS_DynamoDB_Table_InputFormatOptions) SetV__Csv(v AWS_DynamoDB_Table_Csv) AWS_DynamoDB_Table_InputFormatOptions {
	t.Csv = cfz.V(v)
	return t
}
