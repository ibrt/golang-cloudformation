// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_Workteam_OidcMemberDefinition__Type is the CloudFormation type for AWS::SageMaker::Workteam.OidcMemberDefinition.
	AWS_SageMaker_Workteam_OidcMemberDefinition__Type = "AWS::SageMaker::Workteam.OidcMemberDefinition"
)

var (
	// AWS_SageMaker_Workteam_OidcMemberDefinition__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::Workteam.OidcMemberDefinition.
	AWS_SageMaker_Workteam_OidcMemberDefinition__PropertiesMap = struct {
		OidcGroups string
	}{
		OidcGroups: "OidcGroups",
	}

	// AWS_SageMaker_Workteam_OidcMemberDefinition__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::Workteam.OidcMemberDefinition.
	AWS_SageMaker_Workteam_OidcMemberDefinition__PropertiesSlice = []string{
		AWS_SageMaker_Workteam_OidcMemberDefinition__PropertiesMap.OidcGroups,
	}
)

// AWS_SageMaker_Workteam_OidcMemberDefinition is a binding for AWS::SageMaker::Workteam.OidcMemberDefinition.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-oidcmemberdefinition.html
type AWS_SageMaker_Workteam_OidcMemberDefinition struct {
	// OidcGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-oidcmemberdefinition.html#cfn-sagemaker-workteam-oidcmemberdefinition-oidcgroups
	OidcGroups cfz.ExpressionSlice[string] `json:"OidcGroups,omitempty"`
}

// New__AWS_SageMaker_Workteam_OidcMemberDefinition initializes a new AWS_SageMaker_Workteam_OidcMemberDefinition.
func New__AWS_SageMaker_Workteam_OidcMemberDefinition() AWS_SageMaker_Workteam_OidcMemberDefinition {
	return AWS_SageMaker_Workteam_OidcMemberDefinition{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_Workteam_OidcMemberDefinition) GetType() string {
	return AWS_SageMaker_Workteam_OidcMemberDefinition__Type
}

// Set__OidcGroups updates property "OidcGroups".
func (t AWS_SageMaker_Workteam_OidcMemberDefinition) Set__OidcGroups(v cfz.ExpressionSlice[string]) AWS_SageMaker_Workteam_OidcMemberDefinition {
	t.OidcGroups = v
	return t
}

// SetS__OidcGroups updates property "OidcGroups".
func (t AWS_SageMaker_Workteam_OidcMemberDefinition) SetS__OidcGroups(v ...cfz.Expression[string]) AWS_SageMaker_Workteam_OidcMemberDefinition {
	t.OidcGroups = cfz.S(v...)
	return t
}

// SetSV__OidcGroups updates property "OidcGroups".
func (t AWS_SageMaker_Workteam_OidcMemberDefinition) SetSV__OidcGroups(v ...string) AWS_SageMaker_Workteam_OidcMemberDefinition {
	t.OidcGroups = cfz.SV(v...)
	return t
}
