// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_glue

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Glue_SecurityConfiguration_CloudWatchEncryption__Type is the CloudFormation type for AWS::Glue::SecurityConfiguration.CloudWatchEncryption.
	AWS_Glue_SecurityConfiguration_CloudWatchEncryption__Type = "AWS::Glue::SecurityConfiguration.CloudWatchEncryption"
)

var (
	// AWS_Glue_SecurityConfiguration_CloudWatchEncryption__PropertiesMap reports all the CloudFormation properties for AWS::Glue::SecurityConfiguration.CloudWatchEncryption.
	AWS_Glue_SecurityConfiguration_CloudWatchEncryption__PropertiesMap = struct {
		CloudWatchEncryptionMode string
		KmsKeyArn                string
	}{
		CloudWatchEncryptionMode: "CloudWatchEncryptionMode",
		KmsKeyArn:                "KmsKeyArn",
	}

	// AWS_Glue_SecurityConfiguration_CloudWatchEncryption__PropertiesSlice reports all the CloudFormation properties for AWS::Glue::SecurityConfiguration.CloudWatchEncryption.
	AWS_Glue_SecurityConfiguration_CloudWatchEncryption__PropertiesSlice = []string{
		AWS_Glue_SecurityConfiguration_CloudWatchEncryption__PropertiesMap.CloudWatchEncryptionMode,
		AWS_Glue_SecurityConfiguration_CloudWatchEncryption__PropertiesMap.KmsKeyArn,
	}
)

// AWS_Glue_SecurityConfiguration_CloudWatchEncryption is a binding for AWS::Glue::SecurityConfiguration.CloudWatchEncryption.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-cloudwatchencryption.html
type AWS_Glue_SecurityConfiguration_CloudWatchEncryption struct {
	// CloudWatchEncryptionMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-cloudwatchencryption.html#cfn-glue-securityconfiguration-cloudwatchencryption-cloudwatchencryptionmode
	CloudWatchEncryptionMode cfz.Expression[string] `json:"CloudWatchEncryptionMode,omitempty"`

	// KmsKeyArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-cloudwatchencryption.html#cfn-glue-securityconfiguration-cloudwatchencryption-kmskeyarn
	KmsKeyArn cfz.Expression[string] `json:"KmsKeyArn,omitempty"`
}

// New__AWS_Glue_SecurityConfiguration_CloudWatchEncryption initializes a new AWS_Glue_SecurityConfiguration_CloudWatchEncryption.
func New__AWS_Glue_SecurityConfiguration_CloudWatchEncryption() AWS_Glue_SecurityConfiguration_CloudWatchEncryption {
	return AWS_Glue_SecurityConfiguration_CloudWatchEncryption{}
}

// GetType returns the CloudFormation type.
func (AWS_Glue_SecurityConfiguration_CloudWatchEncryption) GetType() string {
	return AWS_Glue_SecurityConfiguration_CloudWatchEncryption__Type
}

// Set__CloudWatchEncryptionMode updates property "CloudWatchEncryptionMode".
func (t AWS_Glue_SecurityConfiguration_CloudWatchEncryption) Set__CloudWatchEncryptionMode(v cfz.Expression[string]) AWS_Glue_SecurityConfiguration_CloudWatchEncryption {
	t.CloudWatchEncryptionMode = v
	return t
}

// SetV__CloudWatchEncryptionMode updates property "CloudWatchEncryptionMode".
func (t AWS_Glue_SecurityConfiguration_CloudWatchEncryption) SetV__CloudWatchEncryptionMode(v string) AWS_Glue_SecurityConfiguration_CloudWatchEncryption {
	t.CloudWatchEncryptionMode = cfz.V(v)
	return t
}

// Set__KmsKeyArn updates property "KmsKeyArn".
func (t AWS_Glue_SecurityConfiguration_CloudWatchEncryption) Set__KmsKeyArn(v cfz.Expression[string]) AWS_Glue_SecurityConfiguration_CloudWatchEncryption {
	t.KmsKeyArn = v
	return t
}

// SetV__KmsKeyArn updates property "KmsKeyArn".
func (t AWS_Glue_SecurityConfiguration_CloudWatchEncryption) SetV__KmsKeyArn(v string) AWS_Glue_SecurityConfiguration_CloudWatchEncryption {
	t.KmsKeyArn = cfz.V(v)
	return t
}
