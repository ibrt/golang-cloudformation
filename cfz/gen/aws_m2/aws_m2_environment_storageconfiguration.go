// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_m2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_M2_Environment_StorageConfiguration__Type is the CloudFormation type for AWS::M2::Environment.StorageConfiguration.
	AWS_M2_Environment_StorageConfiguration__Type = "AWS::M2::Environment.StorageConfiguration"
)

var (
	// AWS_M2_Environment_StorageConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::M2::Environment.StorageConfiguration.
	AWS_M2_Environment_StorageConfiguration__PropertiesMap = struct {
		Efs string
		Fsx string
	}{
		Efs: "Efs",
		Fsx: "Fsx",
	}

	// AWS_M2_Environment_StorageConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::M2::Environment.StorageConfiguration.
	AWS_M2_Environment_StorageConfiguration__PropertiesSlice = []string{
		AWS_M2_Environment_StorageConfiguration__PropertiesMap.Efs,
		AWS_M2_Environment_StorageConfiguration__PropertiesMap.Fsx,
	}
)

// AWS_M2_Environment_StorageConfiguration is a binding for AWS::M2::Environment.StorageConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-storageconfiguration.html
type AWS_M2_Environment_StorageConfiguration struct {
	// Efs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-storageconfiguration.html#cfn-m2-environment-storageconfiguration-efs
	Efs cfz.Expression[AWS_M2_Environment_EfsStorageConfiguration] `json:"Efs,omitempty"`

	// Fsx is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-storageconfiguration.html#cfn-m2-environment-storageconfiguration-fsx
	Fsx cfz.Expression[AWS_M2_Environment_FsxStorageConfiguration] `json:"Fsx,omitempty"`
}

// New__AWS_M2_Environment_StorageConfiguration initializes a new AWS_M2_Environment_StorageConfiguration.
func New__AWS_M2_Environment_StorageConfiguration() AWS_M2_Environment_StorageConfiguration {
	return AWS_M2_Environment_StorageConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_M2_Environment_StorageConfiguration) GetType() string {
	return AWS_M2_Environment_StorageConfiguration__Type
}

// Set__Efs updates property "Efs".
func (t AWS_M2_Environment_StorageConfiguration) Set__Efs(v cfz.Expression[AWS_M2_Environment_EfsStorageConfiguration]) AWS_M2_Environment_StorageConfiguration {
	t.Efs = v
	return t
}

// SetV__Efs updates property "Efs".
func (t AWS_M2_Environment_StorageConfiguration) SetV__Efs(v AWS_M2_Environment_EfsStorageConfiguration) AWS_M2_Environment_StorageConfiguration {
	t.Efs = cfz.V(v)
	return t
}

// Set__Fsx updates property "Fsx".
func (t AWS_M2_Environment_StorageConfiguration) Set__Fsx(v cfz.Expression[AWS_M2_Environment_FsxStorageConfiguration]) AWS_M2_Environment_StorageConfiguration {
	t.Fsx = v
	return t
}

// SetV__Fsx updates property "Fsx".
func (t AWS_M2_Environment_StorageConfiguration) SetV__Fsx(v AWS_M2_Environment_FsxStorageConfiguration) AWS_M2_Environment_StorageConfiguration {
	t.Fsx = cfz.V(v)
	return t
}
