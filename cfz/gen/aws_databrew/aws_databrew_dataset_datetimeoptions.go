// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_databrew

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataBrew_Dataset_DatetimeOptions__Type is the CloudFormation type for AWS::DataBrew::Dataset.DatetimeOptions.
	AWS_DataBrew_Dataset_DatetimeOptions__Type = "AWS::DataBrew::Dataset.DatetimeOptions"
)

var (
	// AWS_DataBrew_Dataset_DatetimeOptions__PropertiesMap reports all the CloudFormation properties for AWS::DataBrew::Dataset.DatetimeOptions.
	AWS_DataBrew_Dataset_DatetimeOptions__PropertiesMap = struct {
		Format         string
		LocaleCode     string
		TimezoneOffset string
	}{
		Format:         "Format",
		LocaleCode:     "LocaleCode",
		TimezoneOffset: "TimezoneOffset",
	}

	// AWS_DataBrew_Dataset_DatetimeOptions__PropertiesSlice reports all the CloudFormation properties for AWS::DataBrew::Dataset.DatetimeOptions.
	AWS_DataBrew_Dataset_DatetimeOptions__PropertiesSlice = []string{
		AWS_DataBrew_Dataset_DatetimeOptions__PropertiesMap.Format,
		AWS_DataBrew_Dataset_DatetimeOptions__PropertiesMap.LocaleCode,
		AWS_DataBrew_Dataset_DatetimeOptions__PropertiesMap.TimezoneOffset,
	}
)

// AWS_DataBrew_Dataset_DatetimeOptions is a binding for AWS::DataBrew::Dataset.DatetimeOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datetimeoptions.html
type AWS_DataBrew_Dataset_DatetimeOptions struct {
	// Format is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datetimeoptions.html#cfn-databrew-dataset-datetimeoptions-format
	Format cfz.Expression[string] `json:"Format,omitempty"`

	// LocaleCode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datetimeoptions.html#cfn-databrew-dataset-datetimeoptions-localecode
	LocaleCode cfz.Expression[string] `json:"LocaleCode,omitempty"`

	// TimezoneOffset is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datetimeoptions.html#cfn-databrew-dataset-datetimeoptions-timezoneoffset
	TimezoneOffset cfz.Expression[string] `json:"TimezoneOffset,omitempty"`
}

// New__AWS_DataBrew_Dataset_DatetimeOptions initializes a new AWS_DataBrew_Dataset_DatetimeOptions.
func New__AWS_DataBrew_Dataset_DatetimeOptions() AWS_DataBrew_Dataset_DatetimeOptions {
	return AWS_DataBrew_Dataset_DatetimeOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_DataBrew_Dataset_DatetimeOptions) GetType() string {
	return AWS_DataBrew_Dataset_DatetimeOptions__Type
}

// Set__Format updates property "Format".
func (t AWS_DataBrew_Dataset_DatetimeOptions) Set__Format(v cfz.Expression[string]) AWS_DataBrew_Dataset_DatetimeOptions {
	t.Format = v
	return t
}

// SetV__Format updates property "Format".
func (t AWS_DataBrew_Dataset_DatetimeOptions) SetV__Format(v string) AWS_DataBrew_Dataset_DatetimeOptions {
	t.Format = cfz.V(v)
	return t
}

// Set__LocaleCode updates property "LocaleCode".
func (t AWS_DataBrew_Dataset_DatetimeOptions) Set__LocaleCode(v cfz.Expression[string]) AWS_DataBrew_Dataset_DatetimeOptions {
	t.LocaleCode = v
	return t
}

// SetV__LocaleCode updates property "LocaleCode".
func (t AWS_DataBrew_Dataset_DatetimeOptions) SetV__LocaleCode(v string) AWS_DataBrew_Dataset_DatetimeOptions {
	t.LocaleCode = cfz.V(v)
	return t
}

// Set__TimezoneOffset updates property "TimezoneOffset".
func (t AWS_DataBrew_Dataset_DatetimeOptions) Set__TimezoneOffset(v cfz.Expression[string]) AWS_DataBrew_Dataset_DatetimeOptions {
	t.TimezoneOffset = v
	return t
}

// SetV__TimezoneOffset updates property "TimezoneOffset".
func (t AWS_DataBrew_Dataset_DatetimeOptions) SetV__TimezoneOffset(v string) AWS_DataBrew_Dataset_DatetimeOptions {
	t.TimezoneOffset = cfz.V(v)
	return t
}
