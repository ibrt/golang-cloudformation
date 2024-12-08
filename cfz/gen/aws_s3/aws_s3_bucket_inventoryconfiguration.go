// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_Bucket_InventoryConfiguration__Type is the CloudFormation type for AWS::S3::Bucket.InventoryConfiguration.
	AWS_S3_Bucket_InventoryConfiguration__Type = "AWS::S3::Bucket.InventoryConfiguration"
)

var (
	// AWS_S3_Bucket_InventoryConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::S3::Bucket.InventoryConfiguration.
	AWS_S3_Bucket_InventoryConfiguration__PropertiesMap = struct {
		Destination            string
		Enabled                string
		Id                     string
		IncludedObjectVersions string
		OptionalFields         string
		Prefix                 string
		ScheduleFrequency      string
	}{
		Destination:            "Destination",
		Enabled:                "Enabled",
		Id:                     "Id",
		IncludedObjectVersions: "IncludedObjectVersions",
		OptionalFields:         "OptionalFields",
		Prefix:                 "Prefix",
		ScheduleFrequency:      "ScheduleFrequency",
	}

	// AWS_S3_Bucket_InventoryConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::S3::Bucket.InventoryConfiguration.
	AWS_S3_Bucket_InventoryConfiguration__PropertiesSlice = []string{
		AWS_S3_Bucket_InventoryConfiguration__PropertiesMap.Destination,
		AWS_S3_Bucket_InventoryConfiguration__PropertiesMap.Enabled,
		AWS_S3_Bucket_InventoryConfiguration__PropertiesMap.Id,
		AWS_S3_Bucket_InventoryConfiguration__PropertiesMap.IncludedObjectVersions,
		AWS_S3_Bucket_InventoryConfiguration__PropertiesMap.OptionalFields,
		AWS_S3_Bucket_InventoryConfiguration__PropertiesMap.Prefix,
		AWS_S3_Bucket_InventoryConfiguration__PropertiesMap.ScheduleFrequency,
	}
)

// AWS_S3_Bucket_InventoryConfiguration is a binding for AWS::S3::Bucket.InventoryConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html
type AWS_S3_Bucket_InventoryConfiguration struct {
	// Destination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-destination
	Destination cfz.Expression[AWS_S3_Bucket_Destination] `json:"Destination,omitempty"`

	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`

	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-id
	Id cfz.Expression[string] `json:"Id,omitempty"`

	// IncludedObjectVersions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-includedobjectversions
	IncludedObjectVersions cfz.Expression[string] `json:"IncludedObjectVersions,omitempty"`

	// OptionalFields is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-optionalfields
	OptionalFields cfz.ExpressionSlice[string] `json:"OptionalFields,omitempty"`

	// Prefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-prefix
	Prefix cfz.Expression[string] `json:"Prefix,omitempty"`

	// ScheduleFrequency is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-schedulefrequency
	ScheduleFrequency cfz.Expression[string] `json:"ScheduleFrequency,omitempty"`
}

// New__AWS_S3_Bucket_InventoryConfiguration initializes a new AWS_S3_Bucket_InventoryConfiguration.
func New__AWS_S3_Bucket_InventoryConfiguration() AWS_S3_Bucket_InventoryConfiguration {
	return AWS_S3_Bucket_InventoryConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_Bucket_InventoryConfiguration) GetType() string {
	return AWS_S3_Bucket_InventoryConfiguration__Type
}

// Set__Destination updates property "Destination".
func (t AWS_S3_Bucket_InventoryConfiguration) Set__Destination(v cfz.Expression[AWS_S3_Bucket_Destination]) AWS_S3_Bucket_InventoryConfiguration {
	t.Destination = v
	return t
}

// SetV__Destination updates property "Destination".
func (t AWS_S3_Bucket_InventoryConfiguration) SetV__Destination(v AWS_S3_Bucket_Destination) AWS_S3_Bucket_InventoryConfiguration {
	t.Destination = cfz.V(v)
	return t
}

// Set__Enabled updates property "Enabled".
func (t AWS_S3_Bucket_InventoryConfiguration) Set__Enabled(v cfz.Expression[bool]) AWS_S3_Bucket_InventoryConfiguration {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_S3_Bucket_InventoryConfiguration) SetV__Enabled(v bool) AWS_S3_Bucket_InventoryConfiguration {
	t.Enabled = cfz.V(v)
	return t
}

// Set__Id updates property "Id".
func (t AWS_S3_Bucket_InventoryConfiguration) Set__Id(v cfz.Expression[string]) AWS_S3_Bucket_InventoryConfiguration {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t AWS_S3_Bucket_InventoryConfiguration) SetV__Id(v string) AWS_S3_Bucket_InventoryConfiguration {
	t.Id = cfz.V(v)
	return t
}

// Set__IncludedObjectVersions updates property "IncludedObjectVersions".
func (t AWS_S3_Bucket_InventoryConfiguration) Set__IncludedObjectVersions(v cfz.Expression[string]) AWS_S3_Bucket_InventoryConfiguration {
	t.IncludedObjectVersions = v
	return t
}

// SetV__IncludedObjectVersions updates property "IncludedObjectVersions".
func (t AWS_S3_Bucket_InventoryConfiguration) SetV__IncludedObjectVersions(v string) AWS_S3_Bucket_InventoryConfiguration {
	t.IncludedObjectVersions = cfz.V(v)
	return t
}

// Set__OptionalFields updates property "OptionalFields".
func (t AWS_S3_Bucket_InventoryConfiguration) Set__OptionalFields(v cfz.ExpressionSlice[string]) AWS_S3_Bucket_InventoryConfiguration {
	t.OptionalFields = v
	return t
}

// SetS__OptionalFields updates property "OptionalFields".
func (t AWS_S3_Bucket_InventoryConfiguration) SetS__OptionalFields(v ...cfz.Expression[string]) AWS_S3_Bucket_InventoryConfiguration {
	t.OptionalFields = cfz.S(v...)
	return t
}

// SetSV__OptionalFields updates property "OptionalFields".
func (t AWS_S3_Bucket_InventoryConfiguration) SetSV__OptionalFields(v ...string) AWS_S3_Bucket_InventoryConfiguration {
	t.OptionalFields = cfz.SV(v...)
	return t
}

// Set__Prefix updates property "Prefix".
func (t AWS_S3_Bucket_InventoryConfiguration) Set__Prefix(v cfz.Expression[string]) AWS_S3_Bucket_InventoryConfiguration {
	t.Prefix = v
	return t
}

// SetV__Prefix updates property "Prefix".
func (t AWS_S3_Bucket_InventoryConfiguration) SetV__Prefix(v string) AWS_S3_Bucket_InventoryConfiguration {
	t.Prefix = cfz.V(v)
	return t
}

// Set__ScheduleFrequency updates property "ScheduleFrequency".
func (t AWS_S3_Bucket_InventoryConfiguration) Set__ScheduleFrequency(v cfz.Expression[string]) AWS_S3_Bucket_InventoryConfiguration {
	t.ScheduleFrequency = v
	return t
}

// SetV__ScheduleFrequency updates property "ScheduleFrequency".
func (t AWS_S3_Bucket_InventoryConfiguration) SetV__ScheduleFrequency(v string) AWS_S3_Bucket_InventoryConfiguration {
	t.ScheduleFrequency = cfz.V(v)
	return t
}
