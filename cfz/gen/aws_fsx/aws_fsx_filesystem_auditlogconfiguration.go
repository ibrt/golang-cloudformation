// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_fsx

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FSx_FileSystem_AuditLogConfiguration__Type is the CloudFormation type for AWS::FSx::FileSystem.AuditLogConfiguration.
	AWS_FSx_FileSystem_AuditLogConfiguration__Type = "AWS::FSx::FileSystem.AuditLogConfiguration"
)

var (
	// AWS_FSx_FileSystem_AuditLogConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::FSx::FileSystem.AuditLogConfiguration.
	AWS_FSx_FileSystem_AuditLogConfiguration__PropertiesMap = struct {
		AuditLogDestination          string
		FileAccessAuditLogLevel      string
		FileShareAccessAuditLogLevel string
	}{
		AuditLogDestination:          "AuditLogDestination",
		FileAccessAuditLogLevel:      "FileAccessAuditLogLevel",
		FileShareAccessAuditLogLevel: "FileShareAccessAuditLogLevel",
	}

	// AWS_FSx_FileSystem_AuditLogConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::FSx::FileSystem.AuditLogConfiguration.
	AWS_FSx_FileSystem_AuditLogConfiguration__PropertiesSlice = []string{
		AWS_FSx_FileSystem_AuditLogConfiguration__PropertiesMap.AuditLogDestination,
		AWS_FSx_FileSystem_AuditLogConfiguration__PropertiesMap.FileAccessAuditLogLevel,
		AWS_FSx_FileSystem_AuditLogConfiguration__PropertiesMap.FileShareAccessAuditLogLevel,
	}
)

// AWS_FSx_FileSystem_AuditLogConfiguration is a binding for AWS::FSx::FileSystem.AuditLogConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration-auditlogconfiguration.html
type AWS_FSx_FileSystem_AuditLogConfiguration struct {
	// AuditLogDestination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration-auditlogconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-auditlogconfiguration-auditlogdestination
	AuditLogDestination cfz.Expression[string] `json:"AuditLogDestination,omitempty"`

	// FileAccessAuditLogLevel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration-auditlogconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-auditlogconfiguration-fileaccessauditloglevel
	FileAccessAuditLogLevel cfz.Expression[string] `json:"FileAccessAuditLogLevel,omitempty"`

	// FileShareAccessAuditLogLevel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration-auditlogconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-auditlogconfiguration-fileshareaccessauditloglevel
	FileShareAccessAuditLogLevel cfz.Expression[string] `json:"FileShareAccessAuditLogLevel,omitempty"`
}

// New__AWS_FSx_FileSystem_AuditLogConfiguration initializes a new AWS_FSx_FileSystem_AuditLogConfiguration.
func New__AWS_FSx_FileSystem_AuditLogConfiguration() AWS_FSx_FileSystem_AuditLogConfiguration {
	return AWS_FSx_FileSystem_AuditLogConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_FSx_FileSystem_AuditLogConfiguration) GetType() string {
	return AWS_FSx_FileSystem_AuditLogConfiguration__Type
}

// Set__AuditLogDestination updates property "AuditLogDestination".
func (t AWS_FSx_FileSystem_AuditLogConfiguration) Set__AuditLogDestination(v cfz.Expression[string]) AWS_FSx_FileSystem_AuditLogConfiguration {
	t.AuditLogDestination = v
	return t
}

// SetV__AuditLogDestination updates property "AuditLogDestination".
func (t AWS_FSx_FileSystem_AuditLogConfiguration) SetV__AuditLogDestination(v string) AWS_FSx_FileSystem_AuditLogConfiguration {
	t.AuditLogDestination = cfz.V(v)
	return t
}

// Set__FileAccessAuditLogLevel updates property "FileAccessAuditLogLevel".
func (t AWS_FSx_FileSystem_AuditLogConfiguration) Set__FileAccessAuditLogLevel(v cfz.Expression[string]) AWS_FSx_FileSystem_AuditLogConfiguration {
	t.FileAccessAuditLogLevel = v
	return t
}

// SetV__FileAccessAuditLogLevel updates property "FileAccessAuditLogLevel".
func (t AWS_FSx_FileSystem_AuditLogConfiguration) SetV__FileAccessAuditLogLevel(v string) AWS_FSx_FileSystem_AuditLogConfiguration {
	t.FileAccessAuditLogLevel = cfz.V(v)
	return t
}

// Set__FileShareAccessAuditLogLevel updates property "FileShareAccessAuditLogLevel".
func (t AWS_FSx_FileSystem_AuditLogConfiguration) Set__FileShareAccessAuditLogLevel(v cfz.Expression[string]) AWS_FSx_FileSystem_AuditLogConfiguration {
	t.FileShareAccessAuditLogLevel = v
	return t
}

// SetV__FileShareAccessAuditLogLevel updates property "FileShareAccessAuditLogLevel".
func (t AWS_FSx_FileSystem_AuditLogConfiguration) SetV__FileShareAccessAuditLogLevel(v string) AWS_FSx_FileSystem_AuditLogConfiguration {
	t.FileShareAccessAuditLogLevel = cfz.V(v)
	return t
}
