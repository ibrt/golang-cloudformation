// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_DataSet_LookbackWindow__Type is the CloudFormation type for AWS::QuickSight::DataSet.LookbackWindow.
	AWS_QuickSight_DataSet_LookbackWindow__Type = "AWS::QuickSight::DataSet.LookbackWindow"
)

var (
	// AWS_QuickSight_DataSet_LookbackWindow__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::DataSet.LookbackWindow.
	AWS_QuickSight_DataSet_LookbackWindow__PropertiesMap = struct {
		ColumnName string
		Size       string
		SizeUnit   string
	}{
		ColumnName: "ColumnName",
		Size:       "Size",
		SizeUnit:   "SizeUnit",
	}

	// AWS_QuickSight_DataSet_LookbackWindow__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::DataSet.LookbackWindow.
	AWS_QuickSight_DataSet_LookbackWindow__PropertiesSlice = []string{
		AWS_QuickSight_DataSet_LookbackWindow__PropertiesMap.ColumnName,
		AWS_QuickSight_DataSet_LookbackWindow__PropertiesMap.Size,
		AWS_QuickSight_DataSet_LookbackWindow__PropertiesMap.SizeUnit,
	}
)

// AWS_QuickSight_DataSet_LookbackWindow is a binding for AWS::QuickSight::DataSet.LookbackWindow.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html
type AWS_QuickSight_DataSet_LookbackWindow struct {
	// ColumnName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html#cfn-quicksight-dataset-lookbackwindow-columnname
	ColumnName cfz.Expression[string] `json:"ColumnName,omitempty"`

	// Size is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html#cfn-quicksight-dataset-lookbackwindow-size
	Size cfz.Expression[float64] `json:"Size,omitempty"`

	// SizeUnit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html#cfn-quicksight-dataset-lookbackwindow-sizeunit
	SizeUnit cfz.Expression[string] `json:"SizeUnit,omitempty"`
}

// New__AWS_QuickSight_DataSet_LookbackWindow initializes a new AWS_QuickSight_DataSet_LookbackWindow.
func New__AWS_QuickSight_DataSet_LookbackWindow() AWS_QuickSight_DataSet_LookbackWindow {
	return AWS_QuickSight_DataSet_LookbackWindow{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_DataSet_LookbackWindow) GetType() string {
	return AWS_QuickSight_DataSet_LookbackWindow__Type
}

// Set__ColumnName updates property "ColumnName".
func (t AWS_QuickSight_DataSet_LookbackWindow) Set__ColumnName(v cfz.Expression[string]) AWS_QuickSight_DataSet_LookbackWindow {
	t.ColumnName = v
	return t
}

// SetV__ColumnName updates property "ColumnName".
func (t AWS_QuickSight_DataSet_LookbackWindow) SetV__ColumnName(v string) AWS_QuickSight_DataSet_LookbackWindow {
	t.ColumnName = cfz.V(v)
	return t
}

// Set__Size updates property "Size".
func (t AWS_QuickSight_DataSet_LookbackWindow) Set__Size(v cfz.Expression[float64]) AWS_QuickSight_DataSet_LookbackWindow {
	t.Size = v
	return t
}

// SetV__Size updates property "Size".
func (t AWS_QuickSight_DataSet_LookbackWindow) SetV__Size(v float64) AWS_QuickSight_DataSet_LookbackWindow {
	t.Size = cfz.V(v)
	return t
}

// Set__SizeUnit updates property "SizeUnit".
func (t AWS_QuickSight_DataSet_LookbackWindow) Set__SizeUnit(v cfz.Expression[string]) AWS_QuickSight_DataSet_LookbackWindow {
	t.SizeUnit = v
	return t
}

// SetV__SizeUnit updates property "SizeUnit".
func (t AWS_QuickSight_DataSet_LookbackWindow) SetV__SizeUnit(v string) AWS_QuickSight_DataSet_LookbackWindow {
	t.SizeUnit = cfz.V(v)
	return t
}
