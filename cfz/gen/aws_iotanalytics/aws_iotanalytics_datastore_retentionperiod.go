// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotanalytics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTAnalytics_Datastore_RetentionPeriod__Type is the CloudFormation type for AWS::IoTAnalytics::Datastore.RetentionPeriod.
	AWS_IoTAnalytics_Datastore_RetentionPeriod__Type = "AWS::IoTAnalytics::Datastore.RetentionPeriod"
)

var (
	// AWS_IoTAnalytics_Datastore_RetentionPeriod__PropertiesMap reports all the CloudFormation properties for AWS::IoTAnalytics::Datastore.RetentionPeriod.
	AWS_IoTAnalytics_Datastore_RetentionPeriod__PropertiesMap = struct {
		NumberOfDays string
		Unlimited    string
	}{
		NumberOfDays: "NumberOfDays",
		Unlimited:    "Unlimited",
	}

	// AWS_IoTAnalytics_Datastore_RetentionPeriod__PropertiesSlice reports all the CloudFormation properties for AWS::IoTAnalytics::Datastore.RetentionPeriod.
	AWS_IoTAnalytics_Datastore_RetentionPeriod__PropertiesSlice = []string{
		AWS_IoTAnalytics_Datastore_RetentionPeriod__PropertiesMap.NumberOfDays,
		AWS_IoTAnalytics_Datastore_RetentionPeriod__PropertiesMap.Unlimited,
	}
)

// AWS_IoTAnalytics_Datastore_RetentionPeriod is a binding for AWS::IoTAnalytics::Datastore.RetentionPeriod.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-retentionperiod.html
type AWS_IoTAnalytics_Datastore_RetentionPeriod struct {
	// NumberOfDays is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-retentionperiod.html#cfn-iotanalytics-datastore-retentionperiod-numberofdays
	NumberOfDays cfz.Expression[int32] `json:"NumberOfDays,omitempty"`

	// Unlimited is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-retentionperiod.html#cfn-iotanalytics-datastore-retentionperiod-unlimited
	Unlimited cfz.Expression[bool] `json:"Unlimited,omitempty"`
}

// New__AWS_IoTAnalytics_Datastore_RetentionPeriod initializes a new AWS_IoTAnalytics_Datastore_RetentionPeriod.
func New__AWS_IoTAnalytics_Datastore_RetentionPeriod() AWS_IoTAnalytics_Datastore_RetentionPeriod {
	return AWS_IoTAnalytics_Datastore_RetentionPeriod{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTAnalytics_Datastore_RetentionPeriod) GetType() string {
	return AWS_IoTAnalytics_Datastore_RetentionPeriod__Type
}

// Set__NumberOfDays updates property "NumberOfDays".
func (t AWS_IoTAnalytics_Datastore_RetentionPeriod) Set__NumberOfDays(v cfz.Expression[int32]) AWS_IoTAnalytics_Datastore_RetentionPeriod {
	t.NumberOfDays = v
	return t
}

// SetV__NumberOfDays updates property "NumberOfDays".
func (t AWS_IoTAnalytics_Datastore_RetentionPeriod) SetV__NumberOfDays(v int32) AWS_IoTAnalytics_Datastore_RetentionPeriod {
	t.NumberOfDays = cfz.V(v)
	return t
}

// Set__Unlimited updates property "Unlimited".
func (t AWS_IoTAnalytics_Datastore_RetentionPeriod) Set__Unlimited(v cfz.Expression[bool]) AWS_IoTAnalytics_Datastore_RetentionPeriod {
	t.Unlimited = v
	return t
}

// SetV__Unlimited updates property "Unlimited".
func (t AWS_IoTAnalytics_Datastore_RetentionPeriod) SetV__Unlimited(v bool) AWS_IoTAnalytics_Datastore_RetentionPeriod {
	t.Unlimited = cfz.V(v)
	return t
}
