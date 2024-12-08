// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_healthlake

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_HealthLake_FHIRDatastore_SseConfiguration__Type is the CloudFormation type for AWS::HealthLake::FHIRDatastore.SseConfiguration.
	AWS_HealthLake_FHIRDatastore_SseConfiguration__Type = "AWS::HealthLake::FHIRDatastore.SseConfiguration"
)

var (
	// AWS_HealthLake_FHIRDatastore_SseConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::HealthLake::FHIRDatastore.SseConfiguration.
	AWS_HealthLake_FHIRDatastore_SseConfiguration__PropertiesMap = struct {
		KmsEncryptionConfig string
	}{
		KmsEncryptionConfig: "KmsEncryptionConfig",
	}

	// AWS_HealthLake_FHIRDatastore_SseConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::HealthLake::FHIRDatastore.SseConfiguration.
	AWS_HealthLake_FHIRDatastore_SseConfiguration__PropertiesSlice = []string{
		AWS_HealthLake_FHIRDatastore_SseConfiguration__PropertiesMap.KmsEncryptionConfig,
	}
)

// AWS_HealthLake_FHIRDatastore_SseConfiguration is a binding for AWS::HealthLake::FHIRDatastore.SseConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-sseconfiguration.html
type AWS_HealthLake_FHIRDatastore_SseConfiguration struct {
	// KmsEncryptionConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-sseconfiguration.html#cfn-healthlake-fhirdatastore-sseconfiguration-kmsencryptionconfig
	KmsEncryptionConfig cfz.Expression[AWS_HealthLake_FHIRDatastore_KmsEncryptionConfig] `json:"KmsEncryptionConfig,omitempty"`
}

// New__AWS_HealthLake_FHIRDatastore_SseConfiguration initializes a new AWS_HealthLake_FHIRDatastore_SseConfiguration.
func New__AWS_HealthLake_FHIRDatastore_SseConfiguration() AWS_HealthLake_FHIRDatastore_SseConfiguration {
	return AWS_HealthLake_FHIRDatastore_SseConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_HealthLake_FHIRDatastore_SseConfiguration) GetType() string {
	return AWS_HealthLake_FHIRDatastore_SseConfiguration__Type
}

// Set__KmsEncryptionConfig updates property "KmsEncryptionConfig".
func (t AWS_HealthLake_FHIRDatastore_SseConfiguration) Set__KmsEncryptionConfig(v cfz.Expression[AWS_HealthLake_FHIRDatastore_KmsEncryptionConfig]) AWS_HealthLake_FHIRDatastore_SseConfiguration {
	t.KmsEncryptionConfig = v
	return t
}

// SetV__KmsEncryptionConfig updates property "KmsEncryptionConfig".
func (t AWS_HealthLake_FHIRDatastore_SseConfiguration) SetV__KmsEncryptionConfig(v AWS_HealthLake_FHIRDatastore_KmsEncryptionConfig) AWS_HealthLake_FHIRDatastore_SseConfiguration {
	t.KmsEncryptionConfig = cfz.V(v)
	return t
}
