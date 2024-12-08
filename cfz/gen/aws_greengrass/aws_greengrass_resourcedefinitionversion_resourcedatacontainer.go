// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_greengrass

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__Type is the CloudFormation type for AWS::Greengrass::ResourceDefinitionVersion.ResourceDataContainer.
	AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__Type = "AWS::Greengrass::ResourceDefinitionVersion.ResourceDataContainer"
)

var (
	// AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__PropertiesMap reports all the CloudFormation properties for AWS::Greengrass::ResourceDefinitionVersion.ResourceDataContainer.
	AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__PropertiesMap = struct {
		LocalDeviceResourceData                   string
		LocalVolumeResourceData                   string
		S3MachineLearningModelResourceData        string
		SageMakerMachineLearningModelResourceData string
		SecretsManagerSecretResourceData          string
	}{
		LocalDeviceResourceData:                   "LocalDeviceResourceData",
		LocalVolumeResourceData:                   "LocalVolumeResourceData",
		S3MachineLearningModelResourceData:        "S3MachineLearningModelResourceData",
		SageMakerMachineLearningModelResourceData: "SageMakerMachineLearningModelResourceData",
		SecretsManagerSecretResourceData:          "SecretsManagerSecretResourceData",
	}

	// AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__PropertiesSlice reports all the CloudFormation properties for AWS::Greengrass::ResourceDefinitionVersion.ResourceDataContainer.
	AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__PropertiesSlice = []string{
		AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__PropertiesMap.LocalDeviceResourceData,
		AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__PropertiesMap.LocalVolumeResourceData,
		AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__PropertiesMap.S3MachineLearningModelResourceData,
		AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__PropertiesMap.SageMakerMachineLearningModelResourceData,
		AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__PropertiesMap.SecretsManagerSecretResourceData,
	}
)

// AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer is a binding for AWS::Greengrass::ResourceDefinitionVersion.ResourceDataContainer.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html
type AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer struct {
	// LocalDeviceResourceData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-localdeviceresourcedata
	LocalDeviceResourceData cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_LocalDeviceResourceData] `json:"LocalDeviceResourceData,omitempty"`

	// LocalVolumeResourceData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-localvolumeresourcedata
	LocalVolumeResourceData cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData] `json:"LocalVolumeResourceData,omitempty"`

	// S3MachineLearningModelResourceData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-s3machinelearningmodelresourcedata
	S3MachineLearningModelResourceData cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_S3MachineLearningModelResourceData] `json:"S3MachineLearningModelResourceData,omitempty"`

	// SageMakerMachineLearningModelResourceData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-sagemakermachinelearningmodelresourcedata
	SageMakerMachineLearningModelResourceData cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_SageMakerMachineLearningModelResourceData] `json:"SageMakerMachineLearningModelResourceData,omitempty"`

	// SecretsManagerSecretResourceData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-secretsmanagersecretresourcedata
	SecretsManagerSecretResourceData cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_SecretsManagerSecretResourceData] `json:"SecretsManagerSecretResourceData,omitempty"`
}

// New__AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer initializes a new AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer.
func New__AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer() AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	return AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer{}
}

// GetType returns the CloudFormation type.
func (AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) GetType() string {
	return AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer__Type
}

// Set__LocalDeviceResourceData updates property "LocalDeviceResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) Set__LocalDeviceResourceData(v cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_LocalDeviceResourceData]) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.LocalDeviceResourceData = v
	return t
}

// SetV__LocalDeviceResourceData updates property "LocalDeviceResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) SetV__LocalDeviceResourceData(v AWS_Greengrass_ResourceDefinitionVersion_LocalDeviceResourceData) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.LocalDeviceResourceData = cfz.V(v)
	return t
}

// Set__LocalVolumeResourceData updates property "LocalVolumeResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) Set__LocalVolumeResourceData(v cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData]) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.LocalVolumeResourceData = v
	return t
}

// SetV__LocalVolumeResourceData updates property "LocalVolumeResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) SetV__LocalVolumeResourceData(v AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.LocalVolumeResourceData = cfz.V(v)
	return t
}

// Set__S3MachineLearningModelResourceData updates property "S3MachineLearningModelResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) Set__S3MachineLearningModelResourceData(v cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_S3MachineLearningModelResourceData]) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.S3MachineLearningModelResourceData = v
	return t
}

// SetV__S3MachineLearningModelResourceData updates property "S3MachineLearningModelResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) SetV__S3MachineLearningModelResourceData(v AWS_Greengrass_ResourceDefinitionVersion_S3MachineLearningModelResourceData) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.S3MachineLearningModelResourceData = cfz.V(v)
	return t
}

// Set__SageMakerMachineLearningModelResourceData updates property "SageMakerMachineLearningModelResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) Set__SageMakerMachineLearningModelResourceData(v cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_SageMakerMachineLearningModelResourceData]) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.SageMakerMachineLearningModelResourceData = v
	return t
}

// SetV__SageMakerMachineLearningModelResourceData updates property "SageMakerMachineLearningModelResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) SetV__SageMakerMachineLearningModelResourceData(v AWS_Greengrass_ResourceDefinitionVersion_SageMakerMachineLearningModelResourceData) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.SageMakerMachineLearningModelResourceData = cfz.V(v)
	return t
}

// Set__SecretsManagerSecretResourceData updates property "SecretsManagerSecretResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) Set__SecretsManagerSecretResourceData(v cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_SecretsManagerSecretResourceData]) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.SecretsManagerSecretResourceData = v
	return t
}

// SetV__SecretsManagerSecretResourceData updates property "SecretsManagerSecretResourceData".
func (t AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer) SetV__SecretsManagerSecretResourceData(v AWS_Greengrass_ResourceDefinitionVersion_SecretsManagerSecretResourceData) AWS_Greengrass_ResourceDefinitionVersion_ResourceDataContainer {
	t.SecretsManagerSecretResourceData = cfz.V(v)
	return t
}
