// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_workspaces

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WorkSpaces_WorkspacesPool_ApplicationSettings__Type is the CloudFormation type for AWS::WorkSpaces::WorkspacesPool.ApplicationSettings.
	AWS_WorkSpaces_WorkspacesPool_ApplicationSettings__Type = "AWS::WorkSpaces::WorkspacesPool.ApplicationSettings"
)

var (
	// AWS_WorkSpaces_WorkspacesPool_ApplicationSettings__PropertiesMap reports all the CloudFormation properties for AWS::WorkSpaces::WorkspacesPool.ApplicationSettings.
	AWS_WorkSpaces_WorkspacesPool_ApplicationSettings__PropertiesMap = struct {
		SettingsGroup string
		Status        string
	}{
		SettingsGroup: "SettingsGroup",
		Status:        "Status",
	}

	// AWS_WorkSpaces_WorkspacesPool_ApplicationSettings__PropertiesSlice reports all the CloudFormation properties for AWS::WorkSpaces::WorkspacesPool.ApplicationSettings.
	AWS_WorkSpaces_WorkspacesPool_ApplicationSettings__PropertiesSlice = []string{
		AWS_WorkSpaces_WorkspacesPool_ApplicationSettings__PropertiesMap.SettingsGroup,
		AWS_WorkSpaces_WorkspacesPool_ApplicationSettings__PropertiesMap.Status,
	}
)

// AWS_WorkSpaces_WorkspacesPool_ApplicationSettings is a binding for AWS::WorkSpaces::WorkspacesPool.ApplicationSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-applicationsettings.html
type AWS_WorkSpaces_WorkspacesPool_ApplicationSettings struct {
	// SettingsGroup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-applicationsettings.html#cfn-workspaces-workspacespool-applicationsettings-settingsgroup
	SettingsGroup cfz.Expression[string] `json:"SettingsGroup,omitempty"`

	// Status is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-applicationsettings.html#cfn-workspaces-workspacespool-applicationsettings-status
	Status cfz.Expression[string] `json:"Status,omitempty"`
}

// New__AWS_WorkSpaces_WorkspacesPool_ApplicationSettings initializes a new AWS_WorkSpaces_WorkspacesPool_ApplicationSettings.
func New__AWS_WorkSpaces_WorkspacesPool_ApplicationSettings() AWS_WorkSpaces_WorkspacesPool_ApplicationSettings {
	return AWS_WorkSpaces_WorkspacesPool_ApplicationSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_WorkSpaces_WorkspacesPool_ApplicationSettings) GetType() string {
	return AWS_WorkSpaces_WorkspacesPool_ApplicationSettings__Type
}

// Set__SettingsGroup updates property "SettingsGroup".
func (t AWS_WorkSpaces_WorkspacesPool_ApplicationSettings) Set__SettingsGroup(v cfz.Expression[string]) AWS_WorkSpaces_WorkspacesPool_ApplicationSettings {
	t.SettingsGroup = v
	return t
}

// SetV__SettingsGroup updates property "SettingsGroup".
func (t AWS_WorkSpaces_WorkspacesPool_ApplicationSettings) SetV__SettingsGroup(v string) AWS_WorkSpaces_WorkspacesPool_ApplicationSettings {
	t.SettingsGroup = cfz.V(v)
	return t
}

// Set__Status updates property "Status".
func (t AWS_WorkSpaces_WorkspacesPool_ApplicationSettings) Set__Status(v cfz.Expression[string]) AWS_WorkSpaces_WorkspacesPool_ApplicationSettings {
	t.Status = v
	return t
}

// SetV__Status updates property "Status".
func (t AWS_WorkSpaces_WorkspacesPool_ApplicationSettings) SetV__Status(v string) AWS_WorkSpaces_WorkspacesPool_ApplicationSettings {
	t.Status = cfz.V(v)
	return t
}
