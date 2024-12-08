// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_Bucket_EventBridgeConfiguration__Type is the CloudFormation type for AWS::S3::Bucket.EventBridgeConfiguration.
	AWS_S3_Bucket_EventBridgeConfiguration__Type = "AWS::S3::Bucket.EventBridgeConfiguration"
)

var (
	// AWS_S3_Bucket_EventBridgeConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::S3::Bucket.EventBridgeConfiguration.
	AWS_S3_Bucket_EventBridgeConfiguration__PropertiesMap = struct {
		EventBridgeEnabled string
	}{
		EventBridgeEnabled: "EventBridgeEnabled",
	}

	// AWS_S3_Bucket_EventBridgeConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::S3::Bucket.EventBridgeConfiguration.
	AWS_S3_Bucket_EventBridgeConfiguration__PropertiesSlice = []string{
		AWS_S3_Bucket_EventBridgeConfiguration__PropertiesMap.EventBridgeEnabled,
	}
)

// AWS_S3_Bucket_EventBridgeConfiguration is a binding for AWS::S3::Bucket.EventBridgeConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-eventbridgeconfiguration.html
type AWS_S3_Bucket_EventBridgeConfiguration struct {
	// EventBridgeEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-eventbridgeconfiguration.html#cfn-s3-bucket-eventbridgeconfiguration-eventbridgeenabled
	EventBridgeEnabled cfz.Expression[bool] `json:"EventBridgeEnabled,omitempty"`
}

// New__AWS_S3_Bucket_EventBridgeConfiguration initializes a new AWS_S3_Bucket_EventBridgeConfiguration.
func New__AWS_S3_Bucket_EventBridgeConfiguration() AWS_S3_Bucket_EventBridgeConfiguration {
	return AWS_S3_Bucket_EventBridgeConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_Bucket_EventBridgeConfiguration) GetType() string {
	return AWS_S3_Bucket_EventBridgeConfiguration__Type
}

// Set__EventBridgeEnabled updates property "EventBridgeEnabled".
func (t AWS_S3_Bucket_EventBridgeConfiguration) Set__EventBridgeEnabled(v cfz.Expression[bool]) AWS_S3_Bucket_EventBridgeConfiguration {
	t.EventBridgeEnabled = v
	return t
}

// SetV__EventBridgeEnabled updates property "EventBridgeEnabled".
func (t AWS_S3_Bucket_EventBridgeConfiguration) SetV__EventBridgeEnabled(v bool) AWS_S3_Bucket_EventBridgeConfiguration {
	t.EventBridgeEnabled = cfz.V(v)
	return t
}
