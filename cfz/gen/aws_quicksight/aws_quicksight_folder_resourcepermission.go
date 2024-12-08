// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_Folder_ResourcePermission__Type is the CloudFormation type for AWS::QuickSight::Folder.ResourcePermission.
	AWS_QuickSight_Folder_ResourcePermission__Type = "AWS::QuickSight::Folder.ResourcePermission"
)

var (
	// AWS_QuickSight_Folder_ResourcePermission__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::Folder.ResourcePermission.
	AWS_QuickSight_Folder_ResourcePermission__PropertiesMap = struct {
		Actions   string
		Principal string
	}{
		Actions:   "Actions",
		Principal: "Principal",
	}

	// AWS_QuickSight_Folder_ResourcePermission__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::Folder.ResourcePermission.
	AWS_QuickSight_Folder_ResourcePermission__PropertiesSlice = []string{
		AWS_QuickSight_Folder_ResourcePermission__PropertiesMap.Actions,
		AWS_QuickSight_Folder_ResourcePermission__PropertiesMap.Principal,
	}
)

// AWS_QuickSight_Folder_ResourcePermission is a binding for AWS::QuickSight::Folder.ResourcePermission.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-folder-resourcepermission.html
type AWS_QuickSight_Folder_ResourcePermission struct {
	// Actions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-folder-resourcepermission.html#cfn-quicksight-folder-resourcepermission-actions
	Actions cfz.ExpressionSlice[string] `json:"Actions,omitempty"`

	// Principal is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-folder-resourcepermission.html#cfn-quicksight-folder-resourcepermission-principal
	Principal cfz.Expression[string] `json:"Principal,omitempty"`
}

// New__AWS_QuickSight_Folder_ResourcePermission initializes a new AWS_QuickSight_Folder_ResourcePermission.
func New__AWS_QuickSight_Folder_ResourcePermission() AWS_QuickSight_Folder_ResourcePermission {
	return AWS_QuickSight_Folder_ResourcePermission{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_Folder_ResourcePermission) GetType() string {
	return AWS_QuickSight_Folder_ResourcePermission__Type
}

// Set__Actions updates property "Actions".
func (t AWS_QuickSight_Folder_ResourcePermission) Set__Actions(v cfz.ExpressionSlice[string]) AWS_QuickSight_Folder_ResourcePermission {
	t.Actions = v
	return t
}

// SetS__Actions updates property "Actions".
func (t AWS_QuickSight_Folder_ResourcePermission) SetS__Actions(v ...cfz.Expression[string]) AWS_QuickSight_Folder_ResourcePermission {
	t.Actions = cfz.S(v...)
	return t
}

// SetSV__Actions updates property "Actions".
func (t AWS_QuickSight_Folder_ResourcePermission) SetSV__Actions(v ...string) AWS_QuickSight_Folder_ResourcePermission {
	t.Actions = cfz.SV(v...)
	return t
}

// Set__Principal updates property "Principal".
func (t AWS_QuickSight_Folder_ResourcePermission) Set__Principal(v cfz.Expression[string]) AWS_QuickSight_Folder_ResourcePermission {
	t.Principal = v
	return t
}

// SetV__Principal updates property "Principal".
func (t AWS_QuickSight_Folder_ResourcePermission) SetV__Principal(v string) AWS_QuickSight_Folder_ResourcePermission {
	t.Principal = cfz.V(v)
	return t
}
