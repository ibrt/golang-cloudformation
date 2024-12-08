// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_athena

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Athena_WorkGroup_AclConfiguration__Type is the CloudFormation type for AWS::Athena::WorkGroup.AclConfiguration.
	AWS_Athena_WorkGroup_AclConfiguration__Type = "AWS::Athena::WorkGroup.AclConfiguration"
)

var (
	// AWS_Athena_WorkGroup_AclConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Athena::WorkGroup.AclConfiguration.
	AWS_Athena_WorkGroup_AclConfiguration__PropertiesMap = struct {
		S3AclOption string
	}{
		S3AclOption: "S3AclOption",
	}

	// AWS_Athena_WorkGroup_AclConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Athena::WorkGroup.AclConfiguration.
	AWS_Athena_WorkGroup_AclConfiguration__PropertiesSlice = []string{
		AWS_Athena_WorkGroup_AclConfiguration__PropertiesMap.S3AclOption,
	}
)

// AWS_Athena_WorkGroup_AclConfiguration is a binding for AWS::Athena::WorkGroup.AclConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-aclconfiguration.html
type AWS_Athena_WorkGroup_AclConfiguration struct {
	// S3AclOption is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-aclconfiguration.html#cfn-athena-workgroup-aclconfiguration-s3acloption
	S3AclOption cfz.Expression[string] `json:"S3AclOption,omitempty"`
}

// New__AWS_Athena_WorkGroup_AclConfiguration initializes a new AWS_Athena_WorkGroup_AclConfiguration.
func New__AWS_Athena_WorkGroup_AclConfiguration() AWS_Athena_WorkGroup_AclConfiguration {
	return AWS_Athena_WorkGroup_AclConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Athena_WorkGroup_AclConfiguration) GetType() string {
	return AWS_Athena_WorkGroup_AclConfiguration__Type
}

// Set__S3AclOption updates property "S3AclOption".
func (t AWS_Athena_WorkGroup_AclConfiguration) Set__S3AclOption(v cfz.Expression[string]) AWS_Athena_WorkGroup_AclConfiguration {
	t.S3AclOption = v
	return t
}

// SetV__S3AclOption updates property "S3AclOption".
func (t AWS_Athena_WorkGroup_AclConfiguration) SetV__S3AclOption(v string) AWS_Athena_WorkGroup_AclConfiguration {
	t.S3AclOption = cfz.V(v)
	return t
}
