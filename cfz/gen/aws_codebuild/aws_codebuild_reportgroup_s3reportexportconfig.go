// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codebuild

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodeBuild_ReportGroup_S3ReportExportConfig__Type is the CloudFormation type for AWS::CodeBuild::ReportGroup.S3ReportExportConfig.
	AWS_CodeBuild_ReportGroup_S3ReportExportConfig__Type = "AWS::CodeBuild::ReportGroup.S3ReportExportConfig"
)

var (
	// AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesMap reports all the CloudFormation properties for AWS::CodeBuild::ReportGroup.S3ReportExportConfig.
	AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesMap = struct {
		Bucket             string
		BucketOwner        string
		EncryptionDisabled string
		EncryptionKey      string
		Packaging          string
		Path               string
	}{
		Bucket:             "Bucket",
		BucketOwner:        "BucketOwner",
		EncryptionDisabled: "EncryptionDisabled",
		EncryptionKey:      "EncryptionKey",
		Packaging:          "Packaging",
		Path:               "Path",
	}

	// AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesSlice reports all the CloudFormation properties for AWS::CodeBuild::ReportGroup.S3ReportExportConfig.
	AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesSlice = []string{
		AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesMap.Bucket,
		AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesMap.BucketOwner,
		AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesMap.EncryptionDisabled,
		AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesMap.EncryptionKey,
		AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesMap.Packaging,
		AWS_CodeBuild_ReportGroup_S3ReportExportConfig__PropertiesMap.Path,
	}
)

// AWS_CodeBuild_ReportGroup_S3ReportExportConfig is a binding for AWS::CodeBuild::ReportGroup.S3ReportExportConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html
type AWS_CodeBuild_ReportGroup_S3ReportExportConfig struct {
	// Bucket is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-bucket
	Bucket cfz.Expression[string] `json:"Bucket,omitempty"`

	// BucketOwner is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-bucketowner
	BucketOwner cfz.Expression[string] `json:"BucketOwner,omitempty"`

	// EncryptionDisabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-encryptiondisabled
	EncryptionDisabled cfz.Expression[bool] `json:"EncryptionDisabled,omitempty"`

	// EncryptionKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-encryptionkey
	EncryptionKey cfz.Expression[string] `json:"EncryptionKey,omitempty"`

	// Packaging is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-packaging
	Packaging cfz.Expression[string] `json:"Packaging,omitempty"`

	// Path is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-path
	Path cfz.Expression[string] `json:"Path,omitempty"`
}

// New__AWS_CodeBuild_ReportGroup_S3ReportExportConfig initializes a new AWS_CodeBuild_ReportGroup_S3ReportExportConfig.
func New__AWS_CodeBuild_ReportGroup_S3ReportExportConfig() AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	return AWS_CodeBuild_ReportGroup_S3ReportExportConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_CodeBuild_ReportGroup_S3ReportExportConfig) GetType() string {
	return AWS_CodeBuild_ReportGroup_S3ReportExportConfig__Type
}

// Set__Bucket updates property "Bucket".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) Set__Bucket(v cfz.Expression[string]) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.Bucket = v
	return t
}

// SetV__Bucket updates property "Bucket".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) SetV__Bucket(v string) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.Bucket = cfz.V(v)
	return t
}

// Set__BucketOwner updates property "BucketOwner".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) Set__BucketOwner(v cfz.Expression[string]) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.BucketOwner = v
	return t
}

// SetV__BucketOwner updates property "BucketOwner".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) SetV__BucketOwner(v string) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.BucketOwner = cfz.V(v)
	return t
}

// Set__EncryptionDisabled updates property "EncryptionDisabled".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) Set__EncryptionDisabled(v cfz.Expression[bool]) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.EncryptionDisabled = v
	return t
}

// SetV__EncryptionDisabled updates property "EncryptionDisabled".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) SetV__EncryptionDisabled(v bool) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.EncryptionDisabled = cfz.V(v)
	return t
}

// Set__EncryptionKey updates property "EncryptionKey".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) Set__EncryptionKey(v cfz.Expression[string]) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.EncryptionKey = v
	return t
}

// SetV__EncryptionKey updates property "EncryptionKey".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) SetV__EncryptionKey(v string) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.EncryptionKey = cfz.V(v)
	return t
}

// Set__Packaging updates property "Packaging".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) Set__Packaging(v cfz.Expression[string]) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.Packaging = v
	return t
}

// SetV__Packaging updates property "Packaging".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) SetV__Packaging(v string) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.Packaging = cfz.V(v)
	return t
}

// Set__Path updates property "Path".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) Set__Path(v cfz.Expression[string]) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.Path = v
	return t
}

// SetV__Path updates property "Path".
func (t AWS_CodeBuild_ReportGroup_S3ReportExportConfig) SetV__Path(v string) AWS_CodeBuild_ReportGroup_S3ReportExportConfig {
	t.Path = cfz.V(v)
	return t
}
