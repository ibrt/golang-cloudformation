// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_Bucket_CorsRule__Type is the CloudFormation type for AWS::S3::Bucket.CorsRule.
	AWS_S3_Bucket_CorsRule__Type = "AWS::S3::Bucket.CorsRule"
)

var (
	// AWS_S3_Bucket_CorsRule__PropertiesMap reports all the CloudFormation properties for AWS::S3::Bucket.CorsRule.
	AWS_S3_Bucket_CorsRule__PropertiesMap = struct {
		AllowedHeaders string
		AllowedMethods string
		AllowedOrigins string
		ExposedHeaders string
		Id             string
		MaxAge         string
	}{
		AllowedHeaders: "AllowedHeaders",
		AllowedMethods: "AllowedMethods",
		AllowedOrigins: "AllowedOrigins",
		ExposedHeaders: "ExposedHeaders",
		Id:             "Id",
		MaxAge:         "MaxAge",
	}

	// AWS_S3_Bucket_CorsRule__PropertiesSlice reports all the CloudFormation properties for AWS::S3::Bucket.CorsRule.
	AWS_S3_Bucket_CorsRule__PropertiesSlice = []string{
		AWS_S3_Bucket_CorsRule__PropertiesMap.AllowedHeaders,
		AWS_S3_Bucket_CorsRule__PropertiesMap.AllowedMethods,
		AWS_S3_Bucket_CorsRule__PropertiesMap.AllowedOrigins,
		AWS_S3_Bucket_CorsRule__PropertiesMap.ExposedHeaders,
		AWS_S3_Bucket_CorsRule__PropertiesMap.Id,
		AWS_S3_Bucket_CorsRule__PropertiesMap.MaxAge,
	}
)

// AWS_S3_Bucket_CorsRule is a binding for AWS::S3::Bucket.CorsRule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html
type AWS_S3_Bucket_CorsRule struct {
	// AllowedHeaders is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-allowedheaders
	AllowedHeaders cfz.ExpressionSlice[string] `json:"AllowedHeaders,omitempty"`

	// AllowedMethods is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-allowedmethods
	AllowedMethods cfz.ExpressionSlice[string] `json:"AllowedMethods,omitempty"`

	// AllowedOrigins is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-allowedorigins
	AllowedOrigins cfz.ExpressionSlice[string] `json:"AllowedOrigins,omitempty"`

	// ExposedHeaders is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-exposedheaders
	ExposedHeaders cfz.ExpressionSlice[string] `json:"ExposedHeaders,omitempty"`

	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-id
	Id cfz.Expression[string] `json:"Id,omitempty"`

	// MaxAge is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-maxage
	MaxAge cfz.Expression[int32] `json:"MaxAge,omitempty"`
}

// New__AWS_S3_Bucket_CorsRule initializes a new AWS_S3_Bucket_CorsRule.
func New__AWS_S3_Bucket_CorsRule() AWS_S3_Bucket_CorsRule {
	return AWS_S3_Bucket_CorsRule{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_Bucket_CorsRule) GetType() string {
	return AWS_S3_Bucket_CorsRule__Type
}

// Set__AllowedHeaders updates property "AllowedHeaders".
func (t AWS_S3_Bucket_CorsRule) Set__AllowedHeaders(v cfz.ExpressionSlice[string]) AWS_S3_Bucket_CorsRule {
	t.AllowedHeaders = v
	return t
}

// SetS__AllowedHeaders updates property "AllowedHeaders".
func (t AWS_S3_Bucket_CorsRule) SetS__AllowedHeaders(v ...cfz.Expression[string]) AWS_S3_Bucket_CorsRule {
	t.AllowedHeaders = cfz.S(v...)
	return t
}

// SetSV__AllowedHeaders updates property "AllowedHeaders".
func (t AWS_S3_Bucket_CorsRule) SetSV__AllowedHeaders(v ...string) AWS_S3_Bucket_CorsRule {
	t.AllowedHeaders = cfz.SV(v...)
	return t
}

// Set__AllowedMethods updates property "AllowedMethods".
func (t AWS_S3_Bucket_CorsRule) Set__AllowedMethods(v cfz.ExpressionSlice[string]) AWS_S3_Bucket_CorsRule {
	t.AllowedMethods = v
	return t
}

// SetS__AllowedMethods updates property "AllowedMethods".
func (t AWS_S3_Bucket_CorsRule) SetS__AllowedMethods(v ...cfz.Expression[string]) AWS_S3_Bucket_CorsRule {
	t.AllowedMethods = cfz.S(v...)
	return t
}

// SetSV__AllowedMethods updates property "AllowedMethods".
func (t AWS_S3_Bucket_CorsRule) SetSV__AllowedMethods(v ...string) AWS_S3_Bucket_CorsRule {
	t.AllowedMethods = cfz.SV(v...)
	return t
}

// Set__AllowedOrigins updates property "AllowedOrigins".
func (t AWS_S3_Bucket_CorsRule) Set__AllowedOrigins(v cfz.ExpressionSlice[string]) AWS_S3_Bucket_CorsRule {
	t.AllowedOrigins = v
	return t
}

// SetS__AllowedOrigins updates property "AllowedOrigins".
func (t AWS_S3_Bucket_CorsRule) SetS__AllowedOrigins(v ...cfz.Expression[string]) AWS_S3_Bucket_CorsRule {
	t.AllowedOrigins = cfz.S(v...)
	return t
}

// SetSV__AllowedOrigins updates property "AllowedOrigins".
func (t AWS_S3_Bucket_CorsRule) SetSV__AllowedOrigins(v ...string) AWS_S3_Bucket_CorsRule {
	t.AllowedOrigins = cfz.SV(v...)
	return t
}

// Set__ExposedHeaders updates property "ExposedHeaders".
func (t AWS_S3_Bucket_CorsRule) Set__ExposedHeaders(v cfz.ExpressionSlice[string]) AWS_S3_Bucket_CorsRule {
	t.ExposedHeaders = v
	return t
}

// SetS__ExposedHeaders updates property "ExposedHeaders".
func (t AWS_S3_Bucket_CorsRule) SetS__ExposedHeaders(v ...cfz.Expression[string]) AWS_S3_Bucket_CorsRule {
	t.ExposedHeaders = cfz.S(v...)
	return t
}

// SetSV__ExposedHeaders updates property "ExposedHeaders".
func (t AWS_S3_Bucket_CorsRule) SetSV__ExposedHeaders(v ...string) AWS_S3_Bucket_CorsRule {
	t.ExposedHeaders = cfz.SV(v...)
	return t
}

// Set__Id updates property "Id".
func (t AWS_S3_Bucket_CorsRule) Set__Id(v cfz.Expression[string]) AWS_S3_Bucket_CorsRule {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t AWS_S3_Bucket_CorsRule) SetV__Id(v string) AWS_S3_Bucket_CorsRule {
	t.Id = cfz.V(v)
	return t
}

// Set__MaxAge updates property "MaxAge".
func (t AWS_S3_Bucket_CorsRule) Set__MaxAge(v cfz.Expression[int32]) AWS_S3_Bucket_CorsRule {
	t.MaxAge = v
	return t
}

// SetV__MaxAge updates property "MaxAge".
func (t AWS_S3_Bucket_CorsRule) SetV__MaxAge(v int32) AWS_S3_Bucket_CorsRule {
	t.MaxAge = cfz.V(v)
	return t
}
