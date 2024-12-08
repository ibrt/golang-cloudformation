// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoT_JobTemplate_PresignedUrlConfig__Type is the CloudFormation type for AWS::IoT::JobTemplate.PresignedUrlConfig.
	AWS_IoT_JobTemplate_PresignedUrlConfig__Type = "AWS::IoT::JobTemplate.PresignedUrlConfig"
)

var (
	// AWS_IoT_JobTemplate_PresignedUrlConfig__PropertiesMap reports all the CloudFormation properties for AWS::IoT::JobTemplate.PresignedUrlConfig.
	AWS_IoT_JobTemplate_PresignedUrlConfig__PropertiesMap = struct {
		ExpiresInSec string
		RoleArn      string
	}{
		ExpiresInSec: "ExpiresInSec",
		RoleArn:      "RoleArn",
	}

	// AWS_IoT_JobTemplate_PresignedUrlConfig__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::JobTemplate.PresignedUrlConfig.
	AWS_IoT_JobTemplate_PresignedUrlConfig__PropertiesSlice = []string{
		AWS_IoT_JobTemplate_PresignedUrlConfig__PropertiesMap.ExpiresInSec,
		AWS_IoT_JobTemplate_PresignedUrlConfig__PropertiesMap.RoleArn,
	}
)

// AWS_IoT_JobTemplate_PresignedUrlConfig is a binding for AWS::IoT::JobTemplate.PresignedUrlConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-presignedurlconfig.html
type AWS_IoT_JobTemplate_PresignedUrlConfig struct {
	// ExpiresInSec is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-presignedurlconfig.html#cfn-iot-jobtemplate-presignedurlconfig-expiresinsec
	ExpiresInSec cfz.Expression[int32] `json:"ExpiresInSec,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-presignedurlconfig.html#cfn-iot-jobtemplate-presignedurlconfig-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`
}

// New__AWS_IoT_JobTemplate_PresignedUrlConfig initializes a new AWS_IoT_JobTemplate_PresignedUrlConfig.
func New__AWS_IoT_JobTemplate_PresignedUrlConfig() AWS_IoT_JobTemplate_PresignedUrlConfig {
	return AWS_IoT_JobTemplate_PresignedUrlConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_IoT_JobTemplate_PresignedUrlConfig) GetType() string {
	return AWS_IoT_JobTemplate_PresignedUrlConfig__Type
}

// Set__ExpiresInSec updates property "ExpiresInSec".
func (t AWS_IoT_JobTemplate_PresignedUrlConfig) Set__ExpiresInSec(v cfz.Expression[int32]) AWS_IoT_JobTemplate_PresignedUrlConfig {
	t.ExpiresInSec = v
	return t
}

// SetV__ExpiresInSec updates property "ExpiresInSec".
func (t AWS_IoT_JobTemplate_PresignedUrlConfig) SetV__ExpiresInSec(v int32) AWS_IoT_JobTemplate_PresignedUrlConfig {
	t.ExpiresInSec = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t AWS_IoT_JobTemplate_PresignedUrlConfig) Set__RoleArn(v cfz.Expression[string]) AWS_IoT_JobTemplate_PresignedUrlConfig {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t AWS_IoT_JobTemplate_PresignedUrlConfig) SetV__RoleArn(v string) AWS_IoT_JobTemplate_PresignedUrlConfig {
	t.RoleArn = cfz.V(v)
	return t
}
