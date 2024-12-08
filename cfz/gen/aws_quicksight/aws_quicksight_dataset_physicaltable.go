// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_DataSet_PhysicalTable__Type is the CloudFormation type for AWS::QuickSight::DataSet.PhysicalTable.
	AWS_QuickSight_DataSet_PhysicalTable__Type = "AWS::QuickSight::DataSet.PhysicalTable"
)

var (
	// AWS_QuickSight_DataSet_PhysicalTable__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::DataSet.PhysicalTable.
	AWS_QuickSight_DataSet_PhysicalTable__PropertiesMap = struct {
		CustomSql       string
		RelationalTable string
		S3Source        string
	}{
		CustomSql:       "CustomSql",
		RelationalTable: "RelationalTable",
		S3Source:        "S3Source",
	}

	// AWS_QuickSight_DataSet_PhysicalTable__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::DataSet.PhysicalTable.
	AWS_QuickSight_DataSet_PhysicalTable__PropertiesSlice = []string{
		AWS_QuickSight_DataSet_PhysicalTable__PropertiesMap.CustomSql,
		AWS_QuickSight_DataSet_PhysicalTable__PropertiesMap.RelationalTable,
		AWS_QuickSight_DataSet_PhysicalTable__PropertiesMap.S3Source,
	}
)

// AWS_QuickSight_DataSet_PhysicalTable is a binding for AWS::QuickSight::DataSet.PhysicalTable.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-physicaltable.html
type AWS_QuickSight_DataSet_PhysicalTable struct {
	// CustomSql is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-physicaltable.html#cfn-quicksight-dataset-physicaltable-customsql
	CustomSql cfz.Expression[AWS_QuickSight_DataSet_CustomSql] `json:"CustomSql,omitempty"`

	// RelationalTable is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-physicaltable.html#cfn-quicksight-dataset-physicaltable-relationaltable
	RelationalTable cfz.Expression[AWS_QuickSight_DataSet_RelationalTable] `json:"RelationalTable,omitempty"`

	// S3Source is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-physicaltable.html#cfn-quicksight-dataset-physicaltable-s3source
	S3Source cfz.Expression[AWS_QuickSight_DataSet_S3Source] `json:"S3Source,omitempty"`
}

// New__AWS_QuickSight_DataSet_PhysicalTable initializes a new AWS_QuickSight_DataSet_PhysicalTable.
func New__AWS_QuickSight_DataSet_PhysicalTable() AWS_QuickSight_DataSet_PhysicalTable {
	return AWS_QuickSight_DataSet_PhysicalTable{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_DataSet_PhysicalTable) GetType() string {
	return AWS_QuickSight_DataSet_PhysicalTable__Type
}

// Set__CustomSql updates property "CustomSql".
func (t AWS_QuickSight_DataSet_PhysicalTable) Set__CustomSql(v cfz.Expression[AWS_QuickSight_DataSet_CustomSql]) AWS_QuickSight_DataSet_PhysicalTable {
	t.CustomSql = v
	return t
}

// SetV__CustomSql updates property "CustomSql".
func (t AWS_QuickSight_DataSet_PhysicalTable) SetV__CustomSql(v AWS_QuickSight_DataSet_CustomSql) AWS_QuickSight_DataSet_PhysicalTable {
	t.CustomSql = cfz.V(v)
	return t
}

// Set__RelationalTable updates property "RelationalTable".
func (t AWS_QuickSight_DataSet_PhysicalTable) Set__RelationalTable(v cfz.Expression[AWS_QuickSight_DataSet_RelationalTable]) AWS_QuickSight_DataSet_PhysicalTable {
	t.RelationalTable = v
	return t
}

// SetV__RelationalTable updates property "RelationalTable".
func (t AWS_QuickSight_DataSet_PhysicalTable) SetV__RelationalTable(v AWS_QuickSight_DataSet_RelationalTable) AWS_QuickSight_DataSet_PhysicalTable {
	t.RelationalTable = cfz.V(v)
	return t
}

// Set__S3Source updates property "S3Source".
func (t AWS_QuickSight_DataSet_PhysicalTable) Set__S3Source(v cfz.Expression[AWS_QuickSight_DataSet_S3Source]) AWS_QuickSight_DataSet_PhysicalTable {
	t.S3Source = v
	return t
}

// SetV__S3Source updates property "S3Source".
func (t AWS_QuickSight_DataSet_PhysicalTable) SetV__S3Source(v AWS_QuickSight_DataSet_S3Source) AWS_QuickSight_DataSet_PhysicalTable {
	t.S3Source = cfz.V(v)
	return t
}
