// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_UserProfile_UserSettings__Type is the CloudFormation type for AWS::SageMaker::UserProfile.UserSettings.
	AWS_SageMaker_UserProfile_UserSettings__Type = "AWS::SageMaker::UserProfile.UserSettings"
)

var (
	// AWS_SageMaker_UserProfile_UserSettings__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::UserProfile.UserSettings.
	AWS_SageMaker_UserProfile_UserSettings__PropertiesMap = struct {
		CodeEditorAppSettings       string
		CustomFileSystemConfigs     string
		CustomPosixUserConfig       string
		DefaultLandingUri           string
		ExecutionRole               string
		JupyterLabAppSettings       string
		JupyterServerAppSettings    string
		KernelGatewayAppSettings    string
		RStudioServerProAppSettings string
		SecurityGroups              string
		SharingSettings             string
		SpaceStorageSettings        string
		StudioWebPortal             string
		StudioWebPortalSettings     string
	}{
		CodeEditorAppSettings:       "CodeEditorAppSettings",
		CustomFileSystemConfigs:     "CustomFileSystemConfigs",
		CustomPosixUserConfig:       "CustomPosixUserConfig",
		DefaultLandingUri:           "DefaultLandingUri",
		ExecutionRole:               "ExecutionRole",
		JupyterLabAppSettings:       "JupyterLabAppSettings",
		JupyterServerAppSettings:    "JupyterServerAppSettings",
		KernelGatewayAppSettings:    "KernelGatewayAppSettings",
		RStudioServerProAppSettings: "RStudioServerProAppSettings",
		SecurityGroups:              "SecurityGroups",
		SharingSettings:             "SharingSettings",
		SpaceStorageSettings:        "SpaceStorageSettings",
		StudioWebPortal:             "StudioWebPortal",
		StudioWebPortalSettings:     "StudioWebPortalSettings",
	}

	// AWS_SageMaker_UserProfile_UserSettings__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::UserProfile.UserSettings.
	AWS_SageMaker_UserProfile_UserSettings__PropertiesSlice = []string{
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.CodeEditorAppSettings,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.CustomFileSystemConfigs,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.CustomPosixUserConfig,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.DefaultLandingUri,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.ExecutionRole,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.JupyterLabAppSettings,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.JupyterServerAppSettings,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.KernelGatewayAppSettings,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.RStudioServerProAppSettings,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.SecurityGroups,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.SharingSettings,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.SpaceStorageSettings,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.StudioWebPortal,
		AWS_SageMaker_UserProfile_UserSettings__PropertiesMap.StudioWebPortalSettings,
	}
)

// AWS_SageMaker_UserProfile_UserSettings is a binding for AWS::SageMaker::UserProfile.UserSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html
type AWS_SageMaker_UserProfile_UserSettings struct {
	// CodeEditorAppSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-codeeditorappsettings
	CodeEditorAppSettings cfz.Expression[AWS_SageMaker_UserProfile_CodeEditorAppSettings] `json:"CodeEditorAppSettings,omitempty"`

	// CustomFileSystemConfigs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-customfilesystemconfigs
	CustomFileSystemConfigs cfz.ExpressionSlice[AWS_SageMaker_UserProfile_CustomFileSystemConfig] `json:"CustomFileSystemConfigs,omitempty"`

	// CustomPosixUserConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-customposixuserconfig
	CustomPosixUserConfig cfz.Expression[AWS_SageMaker_UserProfile_CustomPosixUserConfig] `json:"CustomPosixUserConfig,omitempty"`

	// DefaultLandingUri is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-defaultlandinguri
	DefaultLandingUri cfz.Expression[string] `json:"DefaultLandingUri,omitempty"`

	// ExecutionRole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-executionrole
	ExecutionRole cfz.Expression[string] `json:"ExecutionRole,omitempty"`

	// JupyterLabAppSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-jupyterlabappsettings
	JupyterLabAppSettings cfz.Expression[AWS_SageMaker_UserProfile_JupyterLabAppSettings] `json:"JupyterLabAppSettings,omitempty"`

	// JupyterServerAppSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-jupyterserverappsettings
	JupyterServerAppSettings cfz.Expression[AWS_SageMaker_UserProfile_JupyterServerAppSettings] `json:"JupyterServerAppSettings,omitempty"`

	// KernelGatewayAppSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-kernelgatewayappsettings
	KernelGatewayAppSettings cfz.Expression[AWS_SageMaker_UserProfile_KernelGatewayAppSettings] `json:"KernelGatewayAppSettings,omitempty"`

	// RStudioServerProAppSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-rstudioserverproappsettings
	RStudioServerProAppSettings cfz.Expression[AWS_SageMaker_UserProfile_RStudioServerProAppSettings] `json:"RStudioServerProAppSettings,omitempty"`

	// SecurityGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-securitygroups
	SecurityGroups cfz.ExpressionSlice[string] `json:"SecurityGroups,omitempty"`

	// SharingSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-sharingsettings
	SharingSettings cfz.Expression[AWS_SageMaker_UserProfile_SharingSettings] `json:"SharingSettings,omitempty"`

	// SpaceStorageSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-spacestoragesettings
	SpaceStorageSettings cfz.Expression[AWS_SageMaker_UserProfile_DefaultSpaceStorageSettings] `json:"SpaceStorageSettings,omitempty"`

	// StudioWebPortal is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-studiowebportal
	StudioWebPortal cfz.Expression[string] `json:"StudioWebPortal,omitempty"`

	// StudioWebPortalSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-studiowebportalsettings
	StudioWebPortalSettings cfz.Expression[AWS_SageMaker_UserProfile_StudioWebPortalSettings] `json:"StudioWebPortalSettings,omitempty"`
}

// New__AWS_SageMaker_UserProfile_UserSettings initializes a new AWS_SageMaker_UserProfile_UserSettings.
func New__AWS_SageMaker_UserProfile_UserSettings() AWS_SageMaker_UserProfile_UserSettings {
	return AWS_SageMaker_UserProfile_UserSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_UserProfile_UserSettings) GetType() string {
	return AWS_SageMaker_UserProfile_UserSettings__Type
}

// Set__CodeEditorAppSettings updates property "CodeEditorAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__CodeEditorAppSettings(v cfz.Expression[AWS_SageMaker_UserProfile_CodeEditorAppSettings]) AWS_SageMaker_UserProfile_UserSettings {
	t.CodeEditorAppSettings = v
	return t
}

// SetV__CodeEditorAppSettings updates property "CodeEditorAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__CodeEditorAppSettings(v AWS_SageMaker_UserProfile_CodeEditorAppSettings) AWS_SageMaker_UserProfile_UserSettings {
	t.CodeEditorAppSettings = cfz.V(v)
	return t
}

// Set__CustomFileSystemConfigs updates property "CustomFileSystemConfigs".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__CustomFileSystemConfigs(v cfz.ExpressionSlice[AWS_SageMaker_UserProfile_CustomFileSystemConfig]) AWS_SageMaker_UserProfile_UserSettings {
	t.CustomFileSystemConfigs = v
	return t
}

// SetS__CustomFileSystemConfigs updates property "CustomFileSystemConfigs".
func (t AWS_SageMaker_UserProfile_UserSettings) SetS__CustomFileSystemConfigs(v ...cfz.Expression[AWS_SageMaker_UserProfile_CustomFileSystemConfig]) AWS_SageMaker_UserProfile_UserSettings {
	t.CustomFileSystemConfigs = cfz.S(v...)
	return t
}

// SetSV__CustomFileSystemConfigs updates property "CustomFileSystemConfigs".
func (t AWS_SageMaker_UserProfile_UserSettings) SetSV__CustomFileSystemConfigs(v ...AWS_SageMaker_UserProfile_CustomFileSystemConfig) AWS_SageMaker_UserProfile_UserSettings {
	t.CustomFileSystemConfigs = cfz.SV(v...)
	return t
}

// Set__CustomPosixUserConfig updates property "CustomPosixUserConfig".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__CustomPosixUserConfig(v cfz.Expression[AWS_SageMaker_UserProfile_CustomPosixUserConfig]) AWS_SageMaker_UserProfile_UserSettings {
	t.CustomPosixUserConfig = v
	return t
}

// SetV__CustomPosixUserConfig updates property "CustomPosixUserConfig".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__CustomPosixUserConfig(v AWS_SageMaker_UserProfile_CustomPosixUserConfig) AWS_SageMaker_UserProfile_UserSettings {
	t.CustomPosixUserConfig = cfz.V(v)
	return t
}

// Set__DefaultLandingUri updates property "DefaultLandingUri".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__DefaultLandingUri(v cfz.Expression[string]) AWS_SageMaker_UserProfile_UserSettings {
	t.DefaultLandingUri = v
	return t
}

// SetV__DefaultLandingUri updates property "DefaultLandingUri".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__DefaultLandingUri(v string) AWS_SageMaker_UserProfile_UserSettings {
	t.DefaultLandingUri = cfz.V(v)
	return t
}

// Set__ExecutionRole updates property "ExecutionRole".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__ExecutionRole(v cfz.Expression[string]) AWS_SageMaker_UserProfile_UserSettings {
	t.ExecutionRole = v
	return t
}

// SetV__ExecutionRole updates property "ExecutionRole".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__ExecutionRole(v string) AWS_SageMaker_UserProfile_UserSettings {
	t.ExecutionRole = cfz.V(v)
	return t
}

// Set__JupyterLabAppSettings updates property "JupyterLabAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__JupyterLabAppSettings(v cfz.Expression[AWS_SageMaker_UserProfile_JupyterLabAppSettings]) AWS_SageMaker_UserProfile_UserSettings {
	t.JupyterLabAppSettings = v
	return t
}

// SetV__JupyterLabAppSettings updates property "JupyterLabAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__JupyterLabAppSettings(v AWS_SageMaker_UserProfile_JupyterLabAppSettings) AWS_SageMaker_UserProfile_UserSettings {
	t.JupyterLabAppSettings = cfz.V(v)
	return t
}

// Set__JupyterServerAppSettings updates property "JupyterServerAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__JupyterServerAppSettings(v cfz.Expression[AWS_SageMaker_UserProfile_JupyterServerAppSettings]) AWS_SageMaker_UserProfile_UserSettings {
	t.JupyterServerAppSettings = v
	return t
}

// SetV__JupyterServerAppSettings updates property "JupyterServerAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__JupyterServerAppSettings(v AWS_SageMaker_UserProfile_JupyterServerAppSettings) AWS_SageMaker_UserProfile_UserSettings {
	t.JupyterServerAppSettings = cfz.V(v)
	return t
}

// Set__KernelGatewayAppSettings updates property "KernelGatewayAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__KernelGatewayAppSettings(v cfz.Expression[AWS_SageMaker_UserProfile_KernelGatewayAppSettings]) AWS_SageMaker_UserProfile_UserSettings {
	t.KernelGatewayAppSettings = v
	return t
}

// SetV__KernelGatewayAppSettings updates property "KernelGatewayAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__KernelGatewayAppSettings(v AWS_SageMaker_UserProfile_KernelGatewayAppSettings) AWS_SageMaker_UserProfile_UserSettings {
	t.KernelGatewayAppSettings = cfz.V(v)
	return t
}

// Set__RStudioServerProAppSettings updates property "RStudioServerProAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__RStudioServerProAppSettings(v cfz.Expression[AWS_SageMaker_UserProfile_RStudioServerProAppSettings]) AWS_SageMaker_UserProfile_UserSettings {
	t.RStudioServerProAppSettings = v
	return t
}

// SetV__RStudioServerProAppSettings updates property "RStudioServerProAppSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__RStudioServerProAppSettings(v AWS_SageMaker_UserProfile_RStudioServerProAppSettings) AWS_SageMaker_UserProfile_UserSettings {
	t.RStudioServerProAppSettings = cfz.V(v)
	return t
}

// Set__SecurityGroups updates property "SecurityGroups".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__SecurityGroups(v cfz.ExpressionSlice[string]) AWS_SageMaker_UserProfile_UserSettings {
	t.SecurityGroups = v
	return t
}

// SetS__SecurityGroups updates property "SecurityGroups".
func (t AWS_SageMaker_UserProfile_UserSettings) SetS__SecurityGroups(v ...cfz.Expression[string]) AWS_SageMaker_UserProfile_UserSettings {
	t.SecurityGroups = cfz.S(v...)
	return t
}

// SetSV__SecurityGroups updates property "SecurityGroups".
func (t AWS_SageMaker_UserProfile_UserSettings) SetSV__SecurityGroups(v ...string) AWS_SageMaker_UserProfile_UserSettings {
	t.SecurityGroups = cfz.SV(v...)
	return t
}

// Set__SharingSettings updates property "SharingSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__SharingSettings(v cfz.Expression[AWS_SageMaker_UserProfile_SharingSettings]) AWS_SageMaker_UserProfile_UserSettings {
	t.SharingSettings = v
	return t
}

// SetV__SharingSettings updates property "SharingSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__SharingSettings(v AWS_SageMaker_UserProfile_SharingSettings) AWS_SageMaker_UserProfile_UserSettings {
	t.SharingSettings = cfz.V(v)
	return t
}

// Set__SpaceStorageSettings updates property "SpaceStorageSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__SpaceStorageSettings(v cfz.Expression[AWS_SageMaker_UserProfile_DefaultSpaceStorageSettings]) AWS_SageMaker_UserProfile_UserSettings {
	t.SpaceStorageSettings = v
	return t
}

// SetV__SpaceStorageSettings updates property "SpaceStorageSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__SpaceStorageSettings(v AWS_SageMaker_UserProfile_DefaultSpaceStorageSettings) AWS_SageMaker_UserProfile_UserSettings {
	t.SpaceStorageSettings = cfz.V(v)
	return t
}

// Set__StudioWebPortal updates property "StudioWebPortal".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__StudioWebPortal(v cfz.Expression[string]) AWS_SageMaker_UserProfile_UserSettings {
	t.StudioWebPortal = v
	return t
}

// SetV__StudioWebPortal updates property "StudioWebPortal".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__StudioWebPortal(v string) AWS_SageMaker_UserProfile_UserSettings {
	t.StudioWebPortal = cfz.V(v)
	return t
}

// Set__StudioWebPortalSettings updates property "StudioWebPortalSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) Set__StudioWebPortalSettings(v cfz.Expression[AWS_SageMaker_UserProfile_StudioWebPortalSettings]) AWS_SageMaker_UserProfile_UserSettings {
	t.StudioWebPortalSettings = v
	return t
}

// SetV__StudioWebPortalSettings updates property "StudioWebPortalSettings".
func (t AWS_SageMaker_UserProfile_UserSettings) SetV__StudioWebPortalSettings(v AWS_SageMaker_UserProfile_StudioWebPortalSettings) AWS_SageMaker_UserProfile_UserSettings {
	t.StudioWebPortalSettings = cfz.V(v)
	return t
}
