// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_greengrass

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__Type is the CloudFormation type for AWS::Greengrass::ResourceDefinitionVersion.LocalVolumeResourceData.
	AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__Type = "AWS::Greengrass::ResourceDefinitionVersion.LocalVolumeResourceData"
)

var (
	// AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__PropertiesMap reports all the CloudFormation properties for AWS::Greengrass::ResourceDefinitionVersion.LocalVolumeResourceData.
	AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__PropertiesMap = struct {
		DestinationPath   string
		GroupOwnerSetting string
		SourcePath        string
	}{
		DestinationPath:   "DestinationPath",
		GroupOwnerSetting: "GroupOwnerSetting",
		SourcePath:        "SourcePath",
	}

	// AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__PropertiesSlice reports all the CloudFormation properties for AWS::Greengrass::ResourceDefinitionVersion.LocalVolumeResourceData.
	AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__PropertiesSlice = []string{
		AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__PropertiesMap.DestinationPath,
		AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__PropertiesMap.GroupOwnerSetting,
		AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__PropertiesMap.SourcePath,
	}
)

// AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData is a binding for AWS::Greengrass::ResourceDefinitionVersion.LocalVolumeResourceData.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.html
type AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData struct {
	// DestinationPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.html#cfn-greengrass-resourcedefinitionversion-localvolumeresourcedata-destinationpath
	DestinationPath cfz.Expression[string] `json:"DestinationPath,omitempty"`

	// GroupOwnerSetting is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.html#cfn-greengrass-resourcedefinitionversion-localvolumeresourcedata-groupownersetting
	GroupOwnerSetting cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_GroupOwnerSetting] `json:"GroupOwnerSetting,omitempty"`

	// SourcePath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.html#cfn-greengrass-resourcedefinitionversion-localvolumeresourcedata-sourcepath
	SourcePath cfz.Expression[string] `json:"SourcePath,omitempty"`
}

// New__AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData initializes a new AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData.
func New__AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData() AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData {
	return AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData{}
}

// GetType returns the CloudFormation type.
func (AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData) GetType() string {
	return AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData__Type
}

// Set__DestinationPath updates property "DestinationPath".
func (t AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData) Set__DestinationPath(v cfz.Expression[string]) AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData {
	t.DestinationPath = v
	return t
}

// SetV__DestinationPath updates property "DestinationPath".
func (t AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData) SetV__DestinationPath(v string) AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData {
	t.DestinationPath = cfz.V(v)
	return t
}

// Set__GroupOwnerSetting updates property "GroupOwnerSetting".
func (t AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData) Set__GroupOwnerSetting(v cfz.Expression[AWS_Greengrass_ResourceDefinitionVersion_GroupOwnerSetting]) AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData {
	t.GroupOwnerSetting = v
	return t
}

// SetV__GroupOwnerSetting updates property "GroupOwnerSetting".
func (t AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData) SetV__GroupOwnerSetting(v AWS_Greengrass_ResourceDefinitionVersion_GroupOwnerSetting) AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData {
	t.GroupOwnerSetting = cfz.V(v)
	return t
}

// Set__SourcePath updates property "SourcePath".
func (t AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData) Set__SourcePath(v cfz.Expression[string]) AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData {
	t.SourcePath = v
	return t
}

// SetV__SourcePath updates property "SourcePath".
func (t AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData) SetV__SourcePath(v string) AWS_Greengrass_ResourceDefinitionVersion_LocalVolumeResourceData {
	t.SourcePath = cfz.V(v)
	return t
}
