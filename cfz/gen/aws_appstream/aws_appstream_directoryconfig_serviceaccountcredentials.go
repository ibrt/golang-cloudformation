// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appstream

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppStream_DirectoryConfig_ServiceAccountCredentials__Type is the CloudFormation type for AWS::AppStream::DirectoryConfig.ServiceAccountCredentials.
	AWS_AppStream_DirectoryConfig_ServiceAccountCredentials__Type = "AWS::AppStream::DirectoryConfig.ServiceAccountCredentials"
)

var (
	// AWS_AppStream_DirectoryConfig_ServiceAccountCredentials__PropertiesMap reports all the CloudFormation properties for AWS::AppStream::DirectoryConfig.ServiceAccountCredentials.
	AWS_AppStream_DirectoryConfig_ServiceAccountCredentials__PropertiesMap = struct {
		AccountName     string
		AccountPassword string
	}{
		AccountName:     "AccountName",
		AccountPassword: "AccountPassword",
	}

	// AWS_AppStream_DirectoryConfig_ServiceAccountCredentials__PropertiesSlice reports all the CloudFormation properties for AWS::AppStream::DirectoryConfig.ServiceAccountCredentials.
	AWS_AppStream_DirectoryConfig_ServiceAccountCredentials__PropertiesSlice = []string{
		AWS_AppStream_DirectoryConfig_ServiceAccountCredentials__PropertiesMap.AccountName,
		AWS_AppStream_DirectoryConfig_ServiceAccountCredentials__PropertiesMap.AccountPassword,
	}
)

// AWS_AppStream_DirectoryConfig_ServiceAccountCredentials is a binding for AWS::AppStream::DirectoryConfig.ServiceAccountCredentials.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html
type AWS_AppStream_DirectoryConfig_ServiceAccountCredentials struct {
	// AccountName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html#cfn-appstream-directoryconfig-serviceaccountcredentials-accountname
	AccountName cfz.Expression[string] `json:"AccountName,omitempty"`

	// AccountPassword is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html#cfn-appstream-directoryconfig-serviceaccountcredentials-accountpassword
	AccountPassword cfz.Expression[string] `json:"AccountPassword,omitempty"`
}

// New__AWS_AppStream_DirectoryConfig_ServiceAccountCredentials initializes a new AWS_AppStream_DirectoryConfig_ServiceAccountCredentials.
func New__AWS_AppStream_DirectoryConfig_ServiceAccountCredentials() AWS_AppStream_DirectoryConfig_ServiceAccountCredentials {
	return AWS_AppStream_DirectoryConfig_ServiceAccountCredentials{}
}

// GetType returns the CloudFormation type.
func (AWS_AppStream_DirectoryConfig_ServiceAccountCredentials) GetType() string {
	return AWS_AppStream_DirectoryConfig_ServiceAccountCredentials__Type
}

// Set__AccountName updates property "AccountName".
func (t AWS_AppStream_DirectoryConfig_ServiceAccountCredentials) Set__AccountName(v cfz.Expression[string]) AWS_AppStream_DirectoryConfig_ServiceAccountCredentials {
	t.AccountName = v
	return t
}

// SetV__AccountName updates property "AccountName".
func (t AWS_AppStream_DirectoryConfig_ServiceAccountCredentials) SetV__AccountName(v string) AWS_AppStream_DirectoryConfig_ServiceAccountCredentials {
	t.AccountName = cfz.V(v)
	return t
}

// Set__AccountPassword updates property "AccountPassword".
func (t AWS_AppStream_DirectoryConfig_ServiceAccountCredentials) Set__AccountPassword(v cfz.Expression[string]) AWS_AppStream_DirectoryConfig_ServiceAccountCredentials {
	t.AccountPassword = v
	return t
}

// SetV__AccountPassword updates property "AccountPassword".
func (t AWS_AppStream_DirectoryConfig_ServiceAccountCredentials) SetV__AccountPassword(v string) AWS_AppStream_DirectoryConfig_ServiceAccountCredentials {
	t.AccountPassword = cfz.V(v)
	return t
}
