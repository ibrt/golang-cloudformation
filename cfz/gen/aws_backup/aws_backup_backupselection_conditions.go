// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_backup

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Backup_BackupSelection_Conditions__Type is the CloudFormation type for AWS::Backup::BackupSelection.Conditions.
	AWS_Backup_BackupSelection_Conditions__Type = "AWS::Backup::BackupSelection.Conditions"
)

var (
	// AWS_Backup_BackupSelection_Conditions__PropertiesMap reports all the CloudFormation properties for AWS::Backup::BackupSelection.Conditions.
	AWS_Backup_BackupSelection_Conditions__PropertiesMap = struct {
		StringEquals    string
		StringLike      string
		StringNotEquals string
		StringNotLike   string
	}{
		StringEquals:    "StringEquals",
		StringLike:      "StringLike",
		StringNotEquals: "StringNotEquals",
		StringNotLike:   "StringNotLike",
	}

	// AWS_Backup_BackupSelection_Conditions__PropertiesSlice reports all the CloudFormation properties for AWS::Backup::BackupSelection.Conditions.
	AWS_Backup_BackupSelection_Conditions__PropertiesSlice = []string{
		AWS_Backup_BackupSelection_Conditions__PropertiesMap.StringEquals,
		AWS_Backup_BackupSelection_Conditions__PropertiesMap.StringLike,
		AWS_Backup_BackupSelection_Conditions__PropertiesMap.StringNotEquals,
		AWS_Backup_BackupSelection_Conditions__PropertiesMap.StringNotLike,
	}
)

// AWS_Backup_BackupSelection_Conditions is a binding for AWS::Backup::BackupSelection.Conditions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html
type AWS_Backup_BackupSelection_Conditions struct {
	// StringEquals is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html#cfn-backup-backupselection-conditions-stringequals
	StringEquals cfz.ExpressionSlice[AWS_Backup_BackupSelection_ConditionParameter] `json:"StringEquals,omitempty"`

	// StringLike is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html#cfn-backup-backupselection-conditions-stringlike
	StringLike cfz.ExpressionSlice[AWS_Backup_BackupSelection_ConditionParameter] `json:"StringLike,omitempty"`

	// StringNotEquals is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html#cfn-backup-backupselection-conditions-stringnotequals
	StringNotEquals cfz.ExpressionSlice[AWS_Backup_BackupSelection_ConditionParameter] `json:"StringNotEquals,omitempty"`

	// StringNotLike is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html#cfn-backup-backupselection-conditions-stringnotlike
	StringNotLike cfz.ExpressionSlice[AWS_Backup_BackupSelection_ConditionParameter] `json:"StringNotLike,omitempty"`
}

// New__AWS_Backup_BackupSelection_Conditions initializes a new AWS_Backup_BackupSelection_Conditions.
func New__AWS_Backup_BackupSelection_Conditions() AWS_Backup_BackupSelection_Conditions {
	return AWS_Backup_BackupSelection_Conditions{}
}

// GetType returns the CloudFormation type.
func (AWS_Backup_BackupSelection_Conditions) GetType() string {
	return AWS_Backup_BackupSelection_Conditions__Type
}

// Set__StringEquals updates property "StringEquals".
func (t AWS_Backup_BackupSelection_Conditions) Set__StringEquals(v cfz.ExpressionSlice[AWS_Backup_BackupSelection_ConditionParameter]) AWS_Backup_BackupSelection_Conditions {
	t.StringEquals = v
	return t
}

// SetS__StringEquals updates property "StringEquals".
func (t AWS_Backup_BackupSelection_Conditions) SetS__StringEquals(v ...cfz.Expression[AWS_Backup_BackupSelection_ConditionParameter]) AWS_Backup_BackupSelection_Conditions {
	t.StringEquals = cfz.S(v...)
	return t
}

// SetSV__StringEquals updates property "StringEquals".
func (t AWS_Backup_BackupSelection_Conditions) SetSV__StringEquals(v ...AWS_Backup_BackupSelection_ConditionParameter) AWS_Backup_BackupSelection_Conditions {
	t.StringEquals = cfz.SV(v...)
	return t
}

// Set__StringLike updates property "StringLike".
func (t AWS_Backup_BackupSelection_Conditions) Set__StringLike(v cfz.ExpressionSlice[AWS_Backup_BackupSelection_ConditionParameter]) AWS_Backup_BackupSelection_Conditions {
	t.StringLike = v
	return t
}

// SetS__StringLike updates property "StringLike".
func (t AWS_Backup_BackupSelection_Conditions) SetS__StringLike(v ...cfz.Expression[AWS_Backup_BackupSelection_ConditionParameter]) AWS_Backup_BackupSelection_Conditions {
	t.StringLike = cfz.S(v...)
	return t
}

// SetSV__StringLike updates property "StringLike".
func (t AWS_Backup_BackupSelection_Conditions) SetSV__StringLike(v ...AWS_Backup_BackupSelection_ConditionParameter) AWS_Backup_BackupSelection_Conditions {
	t.StringLike = cfz.SV(v...)
	return t
}

// Set__StringNotEquals updates property "StringNotEquals".
func (t AWS_Backup_BackupSelection_Conditions) Set__StringNotEquals(v cfz.ExpressionSlice[AWS_Backup_BackupSelection_ConditionParameter]) AWS_Backup_BackupSelection_Conditions {
	t.StringNotEquals = v
	return t
}

// SetS__StringNotEquals updates property "StringNotEquals".
func (t AWS_Backup_BackupSelection_Conditions) SetS__StringNotEquals(v ...cfz.Expression[AWS_Backup_BackupSelection_ConditionParameter]) AWS_Backup_BackupSelection_Conditions {
	t.StringNotEquals = cfz.S(v...)
	return t
}

// SetSV__StringNotEquals updates property "StringNotEquals".
func (t AWS_Backup_BackupSelection_Conditions) SetSV__StringNotEquals(v ...AWS_Backup_BackupSelection_ConditionParameter) AWS_Backup_BackupSelection_Conditions {
	t.StringNotEquals = cfz.SV(v...)
	return t
}

// Set__StringNotLike updates property "StringNotLike".
func (t AWS_Backup_BackupSelection_Conditions) Set__StringNotLike(v cfz.ExpressionSlice[AWS_Backup_BackupSelection_ConditionParameter]) AWS_Backup_BackupSelection_Conditions {
	t.StringNotLike = v
	return t
}

// SetS__StringNotLike updates property "StringNotLike".
func (t AWS_Backup_BackupSelection_Conditions) SetS__StringNotLike(v ...cfz.Expression[AWS_Backup_BackupSelection_ConditionParameter]) AWS_Backup_BackupSelection_Conditions {
	t.StringNotLike = cfz.S(v...)
	return t
}

// SetSV__StringNotLike updates property "StringNotLike".
func (t AWS_Backup_BackupSelection_Conditions) SetSV__StringNotLike(v ...AWS_Backup_BackupSelection_ConditionParameter) AWS_Backup_BackupSelection_Conditions {
	t.StringNotLike = cfz.SV(v...)
	return t
}
