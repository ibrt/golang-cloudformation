// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_Bucket_MetricsConfiguration__Type is the CloudFormation type for AWS::S3::Bucket.MetricsConfiguration.
	AWS_S3_Bucket_MetricsConfiguration__Type = "AWS::S3::Bucket.MetricsConfiguration"
)

var (
	// AWS_S3_Bucket_MetricsConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::S3::Bucket.MetricsConfiguration.
	AWS_S3_Bucket_MetricsConfiguration__PropertiesMap = struct {
		AccessPointArn string
		Id             string
		Prefix         string
		TagFilters     string
	}{
		AccessPointArn: "AccessPointArn",
		Id:             "Id",
		Prefix:         "Prefix",
		TagFilters:     "TagFilters",
	}

	// AWS_S3_Bucket_MetricsConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::S3::Bucket.MetricsConfiguration.
	AWS_S3_Bucket_MetricsConfiguration__PropertiesSlice = []string{
		AWS_S3_Bucket_MetricsConfiguration__PropertiesMap.AccessPointArn,
		AWS_S3_Bucket_MetricsConfiguration__PropertiesMap.Id,
		AWS_S3_Bucket_MetricsConfiguration__PropertiesMap.Prefix,
		AWS_S3_Bucket_MetricsConfiguration__PropertiesMap.TagFilters,
	}
)

// AWS_S3_Bucket_MetricsConfiguration is a binding for AWS::S3::Bucket.MetricsConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html
type AWS_S3_Bucket_MetricsConfiguration struct {
	// AccessPointArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html#cfn-s3-bucket-metricsconfiguration-accesspointarn
	AccessPointArn cfz.Expression[string] `json:"AccessPointArn,omitempty"`

	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html#cfn-s3-bucket-metricsconfiguration-id
	Id cfz.Expression[string] `json:"Id,omitempty"`

	// Prefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html#cfn-s3-bucket-metricsconfiguration-prefix
	Prefix cfz.Expression[string] `json:"Prefix,omitempty"`

	// TagFilters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html#cfn-s3-bucket-metricsconfiguration-tagfilters
	TagFilters cfz.ExpressionSlice[AWS_S3_Bucket_TagFilter] `json:"TagFilters,omitempty"`
}

// New__AWS_S3_Bucket_MetricsConfiguration initializes a new AWS_S3_Bucket_MetricsConfiguration.
func New__AWS_S3_Bucket_MetricsConfiguration() AWS_S3_Bucket_MetricsConfiguration {
	return AWS_S3_Bucket_MetricsConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_Bucket_MetricsConfiguration) GetType() string {
	return AWS_S3_Bucket_MetricsConfiguration__Type
}

// Set__AccessPointArn updates property "AccessPointArn".
func (t AWS_S3_Bucket_MetricsConfiguration) Set__AccessPointArn(v cfz.Expression[string]) AWS_S3_Bucket_MetricsConfiguration {
	t.AccessPointArn = v
	return t
}

// SetV__AccessPointArn updates property "AccessPointArn".
func (t AWS_S3_Bucket_MetricsConfiguration) SetV__AccessPointArn(v string) AWS_S3_Bucket_MetricsConfiguration {
	t.AccessPointArn = cfz.V(v)
	return t
}

// Set__Id updates property "Id".
func (t AWS_S3_Bucket_MetricsConfiguration) Set__Id(v cfz.Expression[string]) AWS_S3_Bucket_MetricsConfiguration {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t AWS_S3_Bucket_MetricsConfiguration) SetV__Id(v string) AWS_S3_Bucket_MetricsConfiguration {
	t.Id = cfz.V(v)
	return t
}

// Set__Prefix updates property "Prefix".
func (t AWS_S3_Bucket_MetricsConfiguration) Set__Prefix(v cfz.Expression[string]) AWS_S3_Bucket_MetricsConfiguration {
	t.Prefix = v
	return t
}

// SetV__Prefix updates property "Prefix".
func (t AWS_S3_Bucket_MetricsConfiguration) SetV__Prefix(v string) AWS_S3_Bucket_MetricsConfiguration {
	t.Prefix = cfz.V(v)
	return t
}

// Set__TagFilters updates property "TagFilters".
func (t AWS_S3_Bucket_MetricsConfiguration) Set__TagFilters(v cfz.ExpressionSlice[AWS_S3_Bucket_TagFilter]) AWS_S3_Bucket_MetricsConfiguration {
	t.TagFilters = v
	return t
}

// SetS__TagFilters updates property "TagFilters".
func (t AWS_S3_Bucket_MetricsConfiguration) SetS__TagFilters(v ...cfz.Expression[AWS_S3_Bucket_TagFilter]) AWS_S3_Bucket_MetricsConfiguration {
	t.TagFilters = cfz.S(v...)
	return t
}

// SetSV__TagFilters updates property "TagFilters".
func (t AWS_S3_Bucket_MetricsConfiguration) SetSV__TagFilters(v ...AWS_S3_Bucket_TagFilter) AWS_S3_Bucket_MetricsConfiguration {
	t.TagFilters = cfz.SV(v...)
	return t
}
