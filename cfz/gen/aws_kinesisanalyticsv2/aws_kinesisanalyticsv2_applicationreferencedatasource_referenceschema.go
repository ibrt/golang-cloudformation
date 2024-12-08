// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisanalyticsv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__Type is the CloudFormation type for AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource.ReferenceSchema.
	AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__Type = "AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource.ReferenceSchema"
)

var (
	// AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__PropertiesMap reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource.ReferenceSchema.
	AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__PropertiesMap = struct {
		RecordColumns  string
		RecordEncoding string
		RecordFormat   string
	}{
		RecordColumns:  "RecordColumns",
		RecordEncoding: "RecordEncoding",
		RecordFormat:   "RecordFormat",
	}

	// AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource.ReferenceSchema.
	AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__PropertiesSlice = []string{
		AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__PropertiesMap.RecordColumns,
		AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__PropertiesMap.RecordEncoding,
		AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__PropertiesMap.RecordFormat,
	}
)

// AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema is a binding for AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource.ReferenceSchema.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referenceschema.html
type AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema struct {
	// RecordColumns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-referenceschema-recordcolumns
	RecordColumns cfz.ExpressionSlice[AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_RecordColumn] `json:"RecordColumns,omitempty"`

	// RecordEncoding is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-referenceschema-recordencoding
	RecordEncoding cfz.Expression[string] `json:"RecordEncoding,omitempty"`

	// RecordFormat is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-referenceschema-recordformat
	RecordFormat cfz.Expression[AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_RecordFormat] `json:"RecordFormat,omitempty"`
}

// New__AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema initializes a new AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema.
func New__AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema() AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema {
	return AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema) GetType() string {
	return AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema__Type
}

// Set__RecordColumns updates property "RecordColumns".
func (t AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema) Set__RecordColumns(v cfz.ExpressionSlice[AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_RecordColumn]) AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema {
	t.RecordColumns = v
	return t
}

// SetS__RecordColumns updates property "RecordColumns".
func (t AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema) SetS__RecordColumns(v ...cfz.Expression[AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_RecordColumn]) AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema {
	t.RecordColumns = cfz.S(v...)
	return t
}

// SetSV__RecordColumns updates property "RecordColumns".
func (t AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema) SetSV__RecordColumns(v ...AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_RecordColumn) AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema {
	t.RecordColumns = cfz.SV(v...)
	return t
}

// Set__RecordEncoding updates property "RecordEncoding".
func (t AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema) Set__RecordEncoding(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema {
	t.RecordEncoding = v
	return t
}

// SetV__RecordEncoding updates property "RecordEncoding".
func (t AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema) SetV__RecordEncoding(v string) AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema {
	t.RecordEncoding = cfz.V(v)
	return t
}

// Set__RecordFormat updates property "RecordFormat".
func (t AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema) Set__RecordFormat(v cfz.Expression[AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_RecordFormat]) AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema {
	t.RecordFormat = v
	return t
}

// SetV__RecordFormat updates property "RecordFormat".
func (t AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema) SetV__RecordFormat(v AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_RecordFormat) AWS_KinesisAnalyticsV2_ApplicationReferenceDataSource_ReferenceSchema {
	t.RecordFormat = cfz.V(v)
	return t
}
