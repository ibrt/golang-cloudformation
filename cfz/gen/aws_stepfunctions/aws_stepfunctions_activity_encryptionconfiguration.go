// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_stepfunctions

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_StepFunctions_Activity_EncryptionConfiguration__Type is the CloudFormation type for AWS::StepFunctions::Activity.EncryptionConfiguration.
	AWS_StepFunctions_Activity_EncryptionConfiguration__Type = "AWS::StepFunctions::Activity.EncryptionConfiguration"
)

var (
	// AWS_StepFunctions_Activity_EncryptionConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::StepFunctions::Activity.EncryptionConfiguration.
	AWS_StepFunctions_Activity_EncryptionConfiguration__PropertiesMap = struct {
		KmsDataKeyReusePeriodSeconds string
		KmsKeyId                     string
		Type                         string
	}{
		KmsDataKeyReusePeriodSeconds: "KmsDataKeyReusePeriodSeconds",
		KmsKeyId:                     "KmsKeyId",
		Type:                         "Type",
	}

	// AWS_StepFunctions_Activity_EncryptionConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::StepFunctions::Activity.EncryptionConfiguration.
	AWS_StepFunctions_Activity_EncryptionConfiguration__PropertiesSlice = []string{
		AWS_StepFunctions_Activity_EncryptionConfiguration__PropertiesMap.KmsDataKeyReusePeriodSeconds,
		AWS_StepFunctions_Activity_EncryptionConfiguration__PropertiesMap.KmsKeyId,
		AWS_StepFunctions_Activity_EncryptionConfiguration__PropertiesMap.Type,
	}
)

// AWS_StepFunctions_Activity_EncryptionConfiguration is a binding for AWS::StepFunctions::Activity.EncryptionConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-encryptionconfiguration.html
type AWS_StepFunctions_Activity_EncryptionConfiguration struct {
	// KmsDataKeyReusePeriodSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-encryptionconfiguration.html#cfn-stepfunctions-activity-encryptionconfiguration-kmsdatakeyreuseperiodseconds
	KmsDataKeyReusePeriodSeconds cfz.Expression[int32] `json:"KmsDataKeyReusePeriodSeconds,omitempty"`

	// KmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-encryptionconfiguration.html#cfn-stepfunctions-activity-encryptionconfiguration-kmskeyid
	KmsKeyId cfz.Expression[string] `json:"KmsKeyId,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-encryptionconfiguration.html#cfn-stepfunctions-activity-encryptionconfiguration-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_StepFunctions_Activity_EncryptionConfiguration initializes a new AWS_StepFunctions_Activity_EncryptionConfiguration.
func New__AWS_StepFunctions_Activity_EncryptionConfiguration() AWS_StepFunctions_Activity_EncryptionConfiguration {
	return AWS_StepFunctions_Activity_EncryptionConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_StepFunctions_Activity_EncryptionConfiguration) GetType() string {
	return AWS_StepFunctions_Activity_EncryptionConfiguration__Type
}

// Set__KmsDataKeyReusePeriodSeconds updates property "KmsDataKeyReusePeriodSeconds".
func (t AWS_StepFunctions_Activity_EncryptionConfiguration) Set__KmsDataKeyReusePeriodSeconds(v cfz.Expression[int32]) AWS_StepFunctions_Activity_EncryptionConfiguration {
	t.KmsDataKeyReusePeriodSeconds = v
	return t
}

// SetV__KmsDataKeyReusePeriodSeconds updates property "KmsDataKeyReusePeriodSeconds".
func (t AWS_StepFunctions_Activity_EncryptionConfiguration) SetV__KmsDataKeyReusePeriodSeconds(v int32) AWS_StepFunctions_Activity_EncryptionConfiguration {
	t.KmsDataKeyReusePeriodSeconds = cfz.V(v)
	return t
}

// Set__KmsKeyId updates property "KmsKeyId".
func (t AWS_StepFunctions_Activity_EncryptionConfiguration) Set__KmsKeyId(v cfz.Expression[string]) AWS_StepFunctions_Activity_EncryptionConfiguration {
	t.KmsKeyId = v
	return t
}

// SetV__KmsKeyId updates property "KmsKeyId".
func (t AWS_StepFunctions_Activity_EncryptionConfiguration) SetV__KmsKeyId(v string) AWS_StepFunctions_Activity_EncryptionConfiguration {
	t.KmsKeyId = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_StepFunctions_Activity_EncryptionConfiguration) Set__Type(v cfz.Expression[string]) AWS_StepFunctions_Activity_EncryptionConfiguration {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_StepFunctions_Activity_EncryptionConfiguration) SetV__Type(v string) AWS_StepFunctions_Activity_EncryptionConfiguration {
	t.Type = cfz.V(v)
	return t
}
