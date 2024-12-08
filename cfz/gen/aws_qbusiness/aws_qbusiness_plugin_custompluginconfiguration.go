// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_qbusiness

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QBusiness_Plugin_CustomPluginConfiguration__Type is the CloudFormation type for AWS::QBusiness::Plugin.CustomPluginConfiguration.
	AWS_QBusiness_Plugin_CustomPluginConfiguration__Type = "AWS::QBusiness::Plugin.CustomPluginConfiguration"
)

var (
	// AWS_QBusiness_Plugin_CustomPluginConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::QBusiness::Plugin.CustomPluginConfiguration.
	AWS_QBusiness_Plugin_CustomPluginConfiguration__PropertiesMap = struct {
		ApiSchema     string
		ApiSchemaType string
		Description   string
	}{
		ApiSchema:     "ApiSchema",
		ApiSchemaType: "ApiSchemaType",
		Description:   "Description",
	}

	// AWS_QBusiness_Plugin_CustomPluginConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::QBusiness::Plugin.CustomPluginConfiguration.
	AWS_QBusiness_Plugin_CustomPluginConfiguration__PropertiesSlice = []string{
		AWS_QBusiness_Plugin_CustomPluginConfiguration__PropertiesMap.ApiSchema,
		AWS_QBusiness_Plugin_CustomPluginConfiguration__PropertiesMap.ApiSchemaType,
		AWS_QBusiness_Plugin_CustomPluginConfiguration__PropertiesMap.Description,
	}
)

// AWS_QBusiness_Plugin_CustomPluginConfiguration is a binding for AWS::QBusiness::Plugin.CustomPluginConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-custompluginconfiguration.html
type AWS_QBusiness_Plugin_CustomPluginConfiguration struct {
	// ApiSchema is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-custompluginconfiguration.html#cfn-qbusiness-plugin-custompluginconfiguration-apischema
	ApiSchema cfz.Expression[AWS_QBusiness_Plugin_APISchema] `json:"ApiSchema,omitempty"`

	// ApiSchemaType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-custompluginconfiguration.html#cfn-qbusiness-plugin-custompluginconfiguration-apischematype
	ApiSchemaType cfz.Expression[string] `json:"ApiSchemaType,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-custompluginconfiguration.html#cfn-qbusiness-plugin-custompluginconfiguration-description
	Description cfz.Expression[string] `json:"Description,omitempty"`
}

// New__AWS_QBusiness_Plugin_CustomPluginConfiguration initializes a new AWS_QBusiness_Plugin_CustomPluginConfiguration.
func New__AWS_QBusiness_Plugin_CustomPluginConfiguration() AWS_QBusiness_Plugin_CustomPluginConfiguration {
	return AWS_QBusiness_Plugin_CustomPluginConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_QBusiness_Plugin_CustomPluginConfiguration) GetType() string {
	return AWS_QBusiness_Plugin_CustomPluginConfiguration__Type
}

// Set__ApiSchema updates property "ApiSchema".
func (t AWS_QBusiness_Plugin_CustomPluginConfiguration) Set__ApiSchema(v cfz.Expression[AWS_QBusiness_Plugin_APISchema]) AWS_QBusiness_Plugin_CustomPluginConfiguration {
	t.ApiSchema = v
	return t
}

// SetV__ApiSchema updates property "ApiSchema".
func (t AWS_QBusiness_Plugin_CustomPluginConfiguration) SetV__ApiSchema(v AWS_QBusiness_Plugin_APISchema) AWS_QBusiness_Plugin_CustomPluginConfiguration {
	t.ApiSchema = cfz.V(v)
	return t
}

// Set__ApiSchemaType updates property "ApiSchemaType".
func (t AWS_QBusiness_Plugin_CustomPluginConfiguration) Set__ApiSchemaType(v cfz.Expression[string]) AWS_QBusiness_Plugin_CustomPluginConfiguration {
	t.ApiSchemaType = v
	return t
}

// SetV__ApiSchemaType updates property "ApiSchemaType".
func (t AWS_QBusiness_Plugin_CustomPluginConfiguration) SetV__ApiSchemaType(v string) AWS_QBusiness_Plugin_CustomPluginConfiguration {
	t.ApiSchemaType = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_QBusiness_Plugin_CustomPluginConfiguration) Set__Description(v cfz.Expression[string]) AWS_QBusiness_Plugin_CustomPluginConfiguration {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_QBusiness_Plugin_CustomPluginConfiguration) SetV__Description(v string) AWS_QBusiness_Plugin_CustomPluginConfiguration {
	t.Description = cfz.V(v)
	return t
}
