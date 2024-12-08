// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_timestream

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Timestream_Table_S3Configuration__Type is the CloudFormation type for AWS::Timestream::Table.S3Configuration.
	AWS_Timestream_Table_S3Configuration__Type = "AWS::Timestream::Table.S3Configuration"
)

var (
	// AWS_Timestream_Table_S3Configuration__PropertiesMap reports all the CloudFormation properties for AWS::Timestream::Table.S3Configuration.
	AWS_Timestream_Table_S3Configuration__PropertiesMap = struct {
		BucketName       string
		EncryptionOption string
		KmsKeyId         string
		ObjectKeyPrefix  string
	}{
		BucketName:       "BucketName",
		EncryptionOption: "EncryptionOption",
		KmsKeyId:         "KmsKeyId",
		ObjectKeyPrefix:  "ObjectKeyPrefix",
	}

	// AWS_Timestream_Table_S3Configuration__PropertiesSlice reports all the CloudFormation properties for AWS::Timestream::Table.S3Configuration.
	AWS_Timestream_Table_S3Configuration__PropertiesSlice = []string{
		AWS_Timestream_Table_S3Configuration__PropertiesMap.BucketName,
		AWS_Timestream_Table_S3Configuration__PropertiesMap.EncryptionOption,
		AWS_Timestream_Table_S3Configuration__PropertiesMap.KmsKeyId,
		AWS_Timestream_Table_S3Configuration__PropertiesMap.ObjectKeyPrefix,
	}
)

// AWS_Timestream_Table_S3Configuration is a binding for AWS::Timestream::Table.S3Configuration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-s3configuration.html
type AWS_Timestream_Table_S3Configuration struct {
	// BucketName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-s3configuration.html#cfn-timestream-table-s3configuration-bucketname
	BucketName cfz.Expression[string] `json:"BucketName,omitempty"`

	// EncryptionOption is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-s3configuration.html#cfn-timestream-table-s3configuration-encryptionoption
	EncryptionOption cfz.Expression[string] `json:"EncryptionOption,omitempty"`

	// KmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-s3configuration.html#cfn-timestream-table-s3configuration-kmskeyid
	KmsKeyId cfz.Expression[string] `json:"KmsKeyId,omitempty"`

	// ObjectKeyPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-s3configuration.html#cfn-timestream-table-s3configuration-objectkeyprefix
	ObjectKeyPrefix cfz.Expression[string] `json:"ObjectKeyPrefix,omitempty"`
}

// New__AWS_Timestream_Table_S3Configuration initializes a new AWS_Timestream_Table_S3Configuration.
func New__AWS_Timestream_Table_S3Configuration() AWS_Timestream_Table_S3Configuration {
	return AWS_Timestream_Table_S3Configuration{}
}

// GetType returns the CloudFormation type.
func (AWS_Timestream_Table_S3Configuration) GetType() string {
	return AWS_Timestream_Table_S3Configuration__Type
}

// Set__BucketName updates property "BucketName".
func (t AWS_Timestream_Table_S3Configuration) Set__BucketName(v cfz.Expression[string]) AWS_Timestream_Table_S3Configuration {
	t.BucketName = v
	return t
}

// SetV__BucketName updates property "BucketName".
func (t AWS_Timestream_Table_S3Configuration) SetV__BucketName(v string) AWS_Timestream_Table_S3Configuration {
	t.BucketName = cfz.V(v)
	return t
}

// Set__EncryptionOption updates property "EncryptionOption".
func (t AWS_Timestream_Table_S3Configuration) Set__EncryptionOption(v cfz.Expression[string]) AWS_Timestream_Table_S3Configuration {
	t.EncryptionOption = v
	return t
}

// SetV__EncryptionOption updates property "EncryptionOption".
func (t AWS_Timestream_Table_S3Configuration) SetV__EncryptionOption(v string) AWS_Timestream_Table_S3Configuration {
	t.EncryptionOption = cfz.V(v)
	return t
}

// Set__KmsKeyId updates property "KmsKeyId".
func (t AWS_Timestream_Table_S3Configuration) Set__KmsKeyId(v cfz.Expression[string]) AWS_Timestream_Table_S3Configuration {
	t.KmsKeyId = v
	return t
}

// SetV__KmsKeyId updates property "KmsKeyId".
func (t AWS_Timestream_Table_S3Configuration) SetV__KmsKeyId(v string) AWS_Timestream_Table_S3Configuration {
	t.KmsKeyId = cfz.V(v)
	return t
}

// Set__ObjectKeyPrefix updates property "ObjectKeyPrefix".
func (t AWS_Timestream_Table_S3Configuration) Set__ObjectKeyPrefix(v cfz.Expression[string]) AWS_Timestream_Table_S3Configuration {
	t.ObjectKeyPrefix = v
	return t
}

// SetV__ObjectKeyPrefix updates property "ObjectKeyPrefix".
func (t AWS_Timestream_Table_S3Configuration) SetV__ObjectKeyPrefix(v string) AWS_Timestream_Table_S3Configuration {
	t.ObjectKeyPrefix = cfz.V(v)
	return t
}
