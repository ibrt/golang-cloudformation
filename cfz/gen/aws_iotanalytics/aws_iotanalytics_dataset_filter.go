// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotanalytics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTAnalytics_Dataset_Filter__Type is the CloudFormation type for AWS::IoTAnalytics::Dataset.Filter.
	AWS_IoTAnalytics_Dataset_Filter__Type = "AWS::IoTAnalytics::Dataset.Filter"
)

var (
	// AWS_IoTAnalytics_Dataset_Filter__PropertiesMap reports all the CloudFormation properties for AWS::IoTAnalytics::Dataset.Filter.
	AWS_IoTAnalytics_Dataset_Filter__PropertiesMap = struct {
		DeltaTime string
	}{
		DeltaTime: "DeltaTime",
	}

	// AWS_IoTAnalytics_Dataset_Filter__PropertiesSlice reports all the CloudFormation properties for AWS::IoTAnalytics::Dataset.Filter.
	AWS_IoTAnalytics_Dataset_Filter__PropertiesSlice = []string{
		AWS_IoTAnalytics_Dataset_Filter__PropertiesMap.DeltaTime,
	}
)

// AWS_IoTAnalytics_Dataset_Filter is a binding for AWS::IoTAnalytics::Dataset.Filter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-filter.html
type AWS_IoTAnalytics_Dataset_Filter struct {
	// DeltaTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-filter.html#cfn-iotanalytics-dataset-filter-deltatime
	DeltaTime cfz.Expression[AWS_IoTAnalytics_Dataset_DeltaTime] `json:"DeltaTime,omitempty"`
}

// New__AWS_IoTAnalytics_Dataset_Filter initializes a new AWS_IoTAnalytics_Dataset_Filter.
func New__AWS_IoTAnalytics_Dataset_Filter() AWS_IoTAnalytics_Dataset_Filter {
	return AWS_IoTAnalytics_Dataset_Filter{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTAnalytics_Dataset_Filter) GetType() string {
	return AWS_IoTAnalytics_Dataset_Filter__Type
}

// Set__DeltaTime updates property "DeltaTime".
func (t AWS_IoTAnalytics_Dataset_Filter) Set__DeltaTime(v cfz.Expression[AWS_IoTAnalytics_Dataset_DeltaTime]) AWS_IoTAnalytics_Dataset_Filter {
	t.DeltaTime = v
	return t
}

// SetV__DeltaTime updates property "DeltaTime".
func (t AWS_IoTAnalytics_Dataset_Filter) SetV__DeltaTime(v AWS_IoTAnalytics_Dataset_DeltaTime) AWS_IoTAnalytics_Dataset_Filter {
	t.DeltaTime = cfz.V(v)
	return t
}
