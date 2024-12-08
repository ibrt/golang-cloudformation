// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_DataSet_RelationalTable__Type is the CloudFormation type for AWS::QuickSight::DataSet.RelationalTable.
	AWS_QuickSight_DataSet_RelationalTable__Type = "AWS::QuickSight::DataSet.RelationalTable"
)

var (
	// AWS_QuickSight_DataSet_RelationalTable__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::DataSet.RelationalTable.
	AWS_QuickSight_DataSet_RelationalTable__PropertiesMap = struct {
		Catalog       string
		DataSourceArn string
		InputColumns  string
		Name          string
		Schema        string
	}{
		Catalog:       "Catalog",
		DataSourceArn: "DataSourceArn",
		InputColumns:  "InputColumns",
		Name:          "Name",
		Schema:        "Schema",
	}

	// AWS_QuickSight_DataSet_RelationalTable__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::DataSet.RelationalTable.
	AWS_QuickSight_DataSet_RelationalTable__PropertiesSlice = []string{
		AWS_QuickSight_DataSet_RelationalTable__PropertiesMap.Catalog,
		AWS_QuickSight_DataSet_RelationalTable__PropertiesMap.DataSourceArn,
		AWS_QuickSight_DataSet_RelationalTable__PropertiesMap.InputColumns,
		AWS_QuickSight_DataSet_RelationalTable__PropertiesMap.Name,
		AWS_QuickSight_DataSet_RelationalTable__PropertiesMap.Schema,
	}
)

// AWS_QuickSight_DataSet_RelationalTable is a binding for AWS::QuickSight::DataSet.RelationalTable.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html
type AWS_QuickSight_DataSet_RelationalTable struct {
	// Catalog is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-catalog
	Catalog cfz.Expression[string] `json:"Catalog,omitempty"`

	// DataSourceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-datasourcearn
	DataSourceArn cfz.Expression[string] `json:"DataSourceArn,omitempty"`

	// InputColumns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-inputcolumns
	InputColumns cfz.ExpressionSlice[AWS_QuickSight_DataSet_InputColumn] `json:"InputColumns,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Schema is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-schema
	Schema cfz.Expression[string] `json:"Schema,omitempty"`
}

// New__AWS_QuickSight_DataSet_RelationalTable initializes a new AWS_QuickSight_DataSet_RelationalTable.
func New__AWS_QuickSight_DataSet_RelationalTable() AWS_QuickSight_DataSet_RelationalTable {
	return AWS_QuickSight_DataSet_RelationalTable{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_DataSet_RelationalTable) GetType() string {
	return AWS_QuickSight_DataSet_RelationalTable__Type
}

// Set__Catalog updates property "Catalog".
func (t AWS_QuickSight_DataSet_RelationalTable) Set__Catalog(v cfz.Expression[string]) AWS_QuickSight_DataSet_RelationalTable {
	t.Catalog = v
	return t
}

// SetV__Catalog updates property "Catalog".
func (t AWS_QuickSight_DataSet_RelationalTable) SetV__Catalog(v string) AWS_QuickSight_DataSet_RelationalTable {
	t.Catalog = cfz.V(v)
	return t
}

// Set__DataSourceArn updates property "DataSourceArn".
func (t AWS_QuickSight_DataSet_RelationalTable) Set__DataSourceArn(v cfz.Expression[string]) AWS_QuickSight_DataSet_RelationalTable {
	t.DataSourceArn = v
	return t
}

// SetV__DataSourceArn updates property "DataSourceArn".
func (t AWS_QuickSight_DataSet_RelationalTable) SetV__DataSourceArn(v string) AWS_QuickSight_DataSet_RelationalTable {
	t.DataSourceArn = cfz.V(v)
	return t
}

// Set__InputColumns updates property "InputColumns".
func (t AWS_QuickSight_DataSet_RelationalTable) Set__InputColumns(v cfz.ExpressionSlice[AWS_QuickSight_DataSet_InputColumn]) AWS_QuickSight_DataSet_RelationalTable {
	t.InputColumns = v
	return t
}

// SetS__InputColumns updates property "InputColumns".
func (t AWS_QuickSight_DataSet_RelationalTable) SetS__InputColumns(v ...cfz.Expression[AWS_QuickSight_DataSet_InputColumn]) AWS_QuickSight_DataSet_RelationalTable {
	t.InputColumns = cfz.S(v...)
	return t
}

// SetSV__InputColumns updates property "InputColumns".
func (t AWS_QuickSight_DataSet_RelationalTable) SetSV__InputColumns(v ...AWS_QuickSight_DataSet_InputColumn) AWS_QuickSight_DataSet_RelationalTable {
	t.InputColumns = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t AWS_QuickSight_DataSet_RelationalTable) Set__Name(v cfz.Expression[string]) AWS_QuickSight_DataSet_RelationalTable {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_QuickSight_DataSet_RelationalTable) SetV__Name(v string) AWS_QuickSight_DataSet_RelationalTable {
	t.Name = cfz.V(v)
	return t
}

// Set__Schema updates property "Schema".
func (t AWS_QuickSight_DataSet_RelationalTable) Set__Schema(v cfz.Expression[string]) AWS_QuickSight_DataSet_RelationalTable {
	t.Schema = v
	return t
}

// SetV__Schema updates property "Schema".
func (t AWS_QuickSight_DataSet_RelationalTable) SetV__Schema(v string) AWS_QuickSight_DataSet_RelationalTable {
	t.Schema = cfz.V(v)
	return t
}
