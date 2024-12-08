// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_DataSet_DateTimeDatasetParameter__Type is the CloudFormation type for AWS::QuickSight::DataSet.DateTimeDatasetParameter.
	AWS_QuickSight_DataSet_DateTimeDatasetParameter__Type = "AWS::QuickSight::DataSet.DateTimeDatasetParameter"
)

var (
	// AWS_QuickSight_DataSet_DateTimeDatasetParameter__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::DataSet.DateTimeDatasetParameter.
	AWS_QuickSight_DataSet_DateTimeDatasetParameter__PropertiesMap = struct {
		DefaultValues   string
		Id              string
		Name            string
		TimeGranularity string
		ValueType       string
	}{
		DefaultValues:   "DefaultValues",
		Id:              "Id",
		Name:            "Name",
		TimeGranularity: "TimeGranularity",
		ValueType:       "ValueType",
	}

	// AWS_QuickSight_DataSet_DateTimeDatasetParameter__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::DataSet.DateTimeDatasetParameter.
	AWS_QuickSight_DataSet_DateTimeDatasetParameter__PropertiesSlice = []string{
		AWS_QuickSight_DataSet_DateTimeDatasetParameter__PropertiesMap.DefaultValues,
		AWS_QuickSight_DataSet_DateTimeDatasetParameter__PropertiesMap.Id,
		AWS_QuickSight_DataSet_DateTimeDatasetParameter__PropertiesMap.Name,
		AWS_QuickSight_DataSet_DateTimeDatasetParameter__PropertiesMap.TimeGranularity,
		AWS_QuickSight_DataSet_DateTimeDatasetParameter__PropertiesMap.ValueType,
	}
)

// AWS_QuickSight_DataSet_DateTimeDatasetParameter is a binding for AWS::QuickSight::DataSet.DateTimeDatasetParameter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html
type AWS_QuickSight_DataSet_DateTimeDatasetParameter struct {
	// DefaultValues is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-defaultvalues
	DefaultValues cfz.Expression[AWS_QuickSight_DataSet_DateTimeDatasetParameterDefaultValues] `json:"DefaultValues,omitempty"`

	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-id
	Id cfz.Expression[string] `json:"Id,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// TimeGranularity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-timegranularity
	TimeGranularity cfz.Expression[string] `json:"TimeGranularity,omitempty"`

	// ValueType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-valuetype
	ValueType cfz.Expression[string] `json:"ValueType,omitempty"`
}

// New__AWS_QuickSight_DataSet_DateTimeDatasetParameter initializes a new AWS_QuickSight_DataSet_DateTimeDatasetParameter.
func New__AWS_QuickSight_DataSet_DateTimeDatasetParameter() AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	return AWS_QuickSight_DataSet_DateTimeDatasetParameter{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_DataSet_DateTimeDatasetParameter) GetType() string {
	return AWS_QuickSight_DataSet_DateTimeDatasetParameter__Type
}

// Set__DefaultValues updates property "DefaultValues".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) Set__DefaultValues(v cfz.Expression[AWS_QuickSight_DataSet_DateTimeDatasetParameterDefaultValues]) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.DefaultValues = v
	return t
}

// SetV__DefaultValues updates property "DefaultValues".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) SetV__DefaultValues(v AWS_QuickSight_DataSet_DateTimeDatasetParameterDefaultValues) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.DefaultValues = cfz.V(v)
	return t
}

// Set__Id updates property "Id".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) Set__Id(v cfz.Expression[string]) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) SetV__Id(v string) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.Id = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) Set__Name(v cfz.Expression[string]) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) SetV__Name(v string) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.Name = cfz.V(v)
	return t
}

// Set__TimeGranularity updates property "TimeGranularity".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) Set__TimeGranularity(v cfz.Expression[string]) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.TimeGranularity = v
	return t
}

// SetV__TimeGranularity updates property "TimeGranularity".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) SetV__TimeGranularity(v string) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.TimeGranularity = cfz.V(v)
	return t
}

// Set__ValueType updates property "ValueType".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) Set__ValueType(v cfz.Expression[string]) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.ValueType = v
	return t
}

// SetV__ValueType updates property "ValueType".
func (t AWS_QuickSight_DataSet_DateTimeDatasetParameter) SetV__ValueType(v string) AWS_QuickSight_DataSet_DateTimeDatasetParameter {
	t.ValueType = cfz.V(v)
	return t
}
