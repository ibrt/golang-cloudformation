// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_opensearchservice

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_OpenSearchService_Application_IamIdentityCenterOptions__Type is the CloudFormation type for AWS::OpenSearchService::Application.IamIdentityCenterOptions.
	AWS_OpenSearchService_Application_IamIdentityCenterOptions__Type = "AWS::OpenSearchService::Application.IamIdentityCenterOptions"
)

var (
	// AWS_OpenSearchService_Application_IamIdentityCenterOptions__PropertiesMap reports all the CloudFormation properties for AWS::OpenSearchService::Application.IamIdentityCenterOptions.
	AWS_OpenSearchService_Application_IamIdentityCenterOptions__PropertiesMap = struct {
		Enabled                                string
		IamIdentityCenterInstanceArn           string
		IamRoleForIdentityCenterApplicationArn string
	}{
		Enabled:                                "Enabled",
		IamIdentityCenterInstanceArn:           "IamIdentityCenterInstanceArn",
		IamRoleForIdentityCenterApplicationArn: "IamRoleForIdentityCenterApplicationArn",
	}

	// AWS_OpenSearchService_Application_IamIdentityCenterOptions__PropertiesSlice reports all the CloudFormation properties for AWS::OpenSearchService::Application.IamIdentityCenterOptions.
	AWS_OpenSearchService_Application_IamIdentityCenterOptions__PropertiesSlice = []string{
		AWS_OpenSearchService_Application_IamIdentityCenterOptions__PropertiesMap.Enabled,
		AWS_OpenSearchService_Application_IamIdentityCenterOptions__PropertiesMap.IamIdentityCenterInstanceArn,
		AWS_OpenSearchService_Application_IamIdentityCenterOptions__PropertiesMap.IamRoleForIdentityCenterApplicationArn,
	}
)

// AWS_OpenSearchService_Application_IamIdentityCenterOptions is a binding for AWS::OpenSearchService::Application.IamIdentityCenterOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html
type AWS_OpenSearchService_Application_IamIdentityCenterOptions struct {
	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html#cfn-opensearchservice-application-iamidentitycenteroptions-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`

	// IamIdentityCenterInstanceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html#cfn-opensearchservice-application-iamidentitycenteroptions-iamidentitycenterinstancearn
	IamIdentityCenterInstanceArn cfz.Expression[string] `json:"IamIdentityCenterInstanceArn,omitempty"`

	// IamRoleForIdentityCenterApplicationArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-iamidentitycenteroptions.html#cfn-opensearchservice-application-iamidentitycenteroptions-iamroleforidentitycenterapplicationarn
	IamRoleForIdentityCenterApplicationArn cfz.Expression[string] `json:"IamRoleForIdentityCenterApplicationArn,omitempty"`
}

// New__AWS_OpenSearchService_Application_IamIdentityCenterOptions initializes a new AWS_OpenSearchService_Application_IamIdentityCenterOptions.
func New__AWS_OpenSearchService_Application_IamIdentityCenterOptions() AWS_OpenSearchService_Application_IamIdentityCenterOptions {
	return AWS_OpenSearchService_Application_IamIdentityCenterOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_OpenSearchService_Application_IamIdentityCenterOptions) GetType() string {
	return AWS_OpenSearchService_Application_IamIdentityCenterOptions__Type
}

// Set__Enabled updates property "Enabled".
func (t AWS_OpenSearchService_Application_IamIdentityCenterOptions) Set__Enabled(v cfz.Expression[bool]) AWS_OpenSearchService_Application_IamIdentityCenterOptions {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_OpenSearchService_Application_IamIdentityCenterOptions) SetV__Enabled(v bool) AWS_OpenSearchService_Application_IamIdentityCenterOptions {
	t.Enabled = cfz.V(v)
	return t
}

// Set__IamIdentityCenterInstanceArn updates property "IamIdentityCenterInstanceArn".
func (t AWS_OpenSearchService_Application_IamIdentityCenterOptions) Set__IamIdentityCenterInstanceArn(v cfz.Expression[string]) AWS_OpenSearchService_Application_IamIdentityCenterOptions {
	t.IamIdentityCenterInstanceArn = v
	return t
}

// SetV__IamIdentityCenterInstanceArn updates property "IamIdentityCenterInstanceArn".
func (t AWS_OpenSearchService_Application_IamIdentityCenterOptions) SetV__IamIdentityCenterInstanceArn(v string) AWS_OpenSearchService_Application_IamIdentityCenterOptions {
	t.IamIdentityCenterInstanceArn = cfz.V(v)
	return t
}

// Set__IamRoleForIdentityCenterApplicationArn updates property "IamRoleForIdentityCenterApplicationArn".
func (t AWS_OpenSearchService_Application_IamIdentityCenterOptions) Set__IamRoleForIdentityCenterApplicationArn(v cfz.Expression[string]) AWS_OpenSearchService_Application_IamIdentityCenterOptions {
	t.IamRoleForIdentityCenterApplicationArn = v
	return t
}

// SetV__IamRoleForIdentityCenterApplicationArn updates property "IamRoleForIdentityCenterApplicationArn".
func (t AWS_OpenSearchService_Application_IamIdentityCenterOptions) SetV__IamRoleForIdentityCenterApplicationArn(v string) AWS_OpenSearchService_Application_IamIdentityCenterOptions {
	t.IamRoleForIdentityCenterApplicationArn = cfz.V(v)
	return t
}
