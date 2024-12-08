// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kendra

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Kendra_DataSource_OneDriveUsers__Type is the CloudFormation type for AWS::Kendra::DataSource.OneDriveUsers.
	AWS_Kendra_DataSource_OneDriveUsers__Type = "AWS::Kendra::DataSource.OneDriveUsers"
)

var (
	// AWS_Kendra_DataSource_OneDriveUsers__PropertiesMap reports all the CloudFormation properties for AWS::Kendra::DataSource.OneDriveUsers.
	AWS_Kendra_DataSource_OneDriveUsers__PropertiesMap = struct {
		OneDriveUserList   string
		OneDriveUserS3Path string
	}{
		OneDriveUserList:   "OneDriveUserList",
		OneDriveUserS3Path: "OneDriveUserS3Path",
	}

	// AWS_Kendra_DataSource_OneDriveUsers__PropertiesSlice reports all the CloudFormation properties for AWS::Kendra::DataSource.OneDriveUsers.
	AWS_Kendra_DataSource_OneDriveUsers__PropertiesSlice = []string{
		AWS_Kendra_DataSource_OneDriveUsers__PropertiesMap.OneDriveUserList,
		AWS_Kendra_DataSource_OneDriveUsers__PropertiesMap.OneDriveUserS3Path,
	}
)

// AWS_Kendra_DataSource_OneDriveUsers is a binding for AWS::Kendra::DataSource.OneDriveUsers.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveusers.html
type AWS_Kendra_DataSource_OneDriveUsers struct {
	// OneDriveUserList is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveusers.html#cfn-kendra-datasource-onedriveusers-onedriveuserlist
	OneDriveUserList cfz.ExpressionSlice[string] `json:"OneDriveUserList,omitempty"`

	// OneDriveUserS3Path is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveusers.html#cfn-kendra-datasource-onedriveusers-onedriveusers3path
	OneDriveUserS3Path cfz.Expression[AWS_Kendra_DataSource_S3Path] `json:"OneDriveUserS3Path,omitempty"`
}

// New__AWS_Kendra_DataSource_OneDriveUsers initializes a new AWS_Kendra_DataSource_OneDriveUsers.
func New__AWS_Kendra_DataSource_OneDriveUsers() AWS_Kendra_DataSource_OneDriveUsers {
	return AWS_Kendra_DataSource_OneDriveUsers{}
}

// GetType returns the CloudFormation type.
func (AWS_Kendra_DataSource_OneDriveUsers) GetType() string {
	return AWS_Kendra_DataSource_OneDriveUsers__Type
}

// Set__OneDriveUserList updates property "OneDriveUserList".
func (t AWS_Kendra_DataSource_OneDriveUsers) Set__OneDriveUserList(v cfz.ExpressionSlice[string]) AWS_Kendra_DataSource_OneDriveUsers {
	t.OneDriveUserList = v
	return t
}

// SetS__OneDriveUserList updates property "OneDriveUserList".
func (t AWS_Kendra_DataSource_OneDriveUsers) SetS__OneDriveUserList(v ...cfz.Expression[string]) AWS_Kendra_DataSource_OneDriveUsers {
	t.OneDriveUserList = cfz.S(v...)
	return t
}

// SetSV__OneDriveUserList updates property "OneDriveUserList".
func (t AWS_Kendra_DataSource_OneDriveUsers) SetSV__OneDriveUserList(v ...string) AWS_Kendra_DataSource_OneDriveUsers {
	t.OneDriveUserList = cfz.SV(v...)
	return t
}

// Set__OneDriveUserS3Path updates property "OneDriveUserS3Path".
func (t AWS_Kendra_DataSource_OneDriveUsers) Set__OneDriveUserS3Path(v cfz.Expression[AWS_Kendra_DataSource_S3Path]) AWS_Kendra_DataSource_OneDriveUsers {
	t.OneDriveUserS3Path = v
	return t
}

// SetV__OneDriveUserS3Path updates property "OneDriveUserS3Path".
func (t AWS_Kendra_DataSource_OneDriveUsers) SetV__OneDriveUserS3Path(v AWS_Kendra_DataSource_S3Path) AWS_Kendra_DataSource_OneDriveUsers {
	t.OneDriveUserS3Path = cfz.V(v)
	return t
}
