// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appsync

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__Type is the CloudFormation type for AWS::AppSync::GraphQLApi.LambdaAuthorizerConfig.
	AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__Type = "AWS::AppSync::GraphQLApi.LambdaAuthorizerConfig"
)

var (
	// AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__PropertiesMap reports all the CloudFormation properties for AWS::AppSync::GraphQLApi.LambdaAuthorizerConfig.
	AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__PropertiesMap = struct {
		AuthorizerResultTtlInSeconds string
		AuthorizerUri                string
		IdentityValidationExpression string
	}{
		AuthorizerResultTtlInSeconds: "AuthorizerResultTtlInSeconds",
		AuthorizerUri:                "AuthorizerUri",
		IdentityValidationExpression: "IdentityValidationExpression",
	}

	// AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__PropertiesSlice reports all the CloudFormation properties for AWS::AppSync::GraphQLApi.LambdaAuthorizerConfig.
	AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__PropertiesSlice = []string{
		AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__PropertiesMap.AuthorizerResultTtlInSeconds,
		AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__PropertiesMap.AuthorizerUri,
		AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__PropertiesMap.IdentityValidationExpression,
	}
)

// AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig is a binding for AWS::AppSync::GraphQLApi.LambdaAuthorizerConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html
type AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig struct {
	// AuthorizerResultTtlInSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html#cfn-appsync-graphqlapi-lambdaauthorizerconfig-authorizerresultttlinseconds
	AuthorizerResultTtlInSeconds cfz.Expression[float64] `json:"AuthorizerResultTtlInSeconds,omitempty"`

	// AuthorizerUri is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html#cfn-appsync-graphqlapi-lambdaauthorizerconfig-authorizeruri
	AuthorizerUri cfz.Expression[string] `json:"AuthorizerUri,omitempty"`

	// IdentityValidationExpression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html#cfn-appsync-graphqlapi-lambdaauthorizerconfig-identityvalidationexpression
	IdentityValidationExpression cfz.Expression[string] `json:"IdentityValidationExpression,omitempty"`
}

// New__AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig initializes a new AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig.
func New__AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig() AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig {
	return AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig) GetType() string {
	return AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig__Type
}

// Set__AuthorizerResultTtlInSeconds updates property "AuthorizerResultTtlInSeconds".
func (t AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig) Set__AuthorizerResultTtlInSeconds(v cfz.Expression[float64]) AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig {
	t.AuthorizerResultTtlInSeconds = v
	return t
}

// SetV__AuthorizerResultTtlInSeconds updates property "AuthorizerResultTtlInSeconds".
func (t AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig) SetV__AuthorizerResultTtlInSeconds(v float64) AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig {
	t.AuthorizerResultTtlInSeconds = cfz.V(v)
	return t
}

// Set__AuthorizerUri updates property "AuthorizerUri".
func (t AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig) Set__AuthorizerUri(v cfz.Expression[string]) AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig {
	t.AuthorizerUri = v
	return t
}

// SetV__AuthorizerUri updates property "AuthorizerUri".
func (t AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig) SetV__AuthorizerUri(v string) AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig {
	t.AuthorizerUri = cfz.V(v)
	return t
}

// Set__IdentityValidationExpression updates property "IdentityValidationExpression".
func (t AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig) Set__IdentityValidationExpression(v cfz.Expression[string]) AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig {
	t.IdentityValidationExpression = v
	return t
}

// SetV__IdentityValidationExpression updates property "IdentityValidationExpression".
func (t AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig) SetV__IdentityValidationExpression(v string) AWS_AppSync_GraphQLApi_LambdaAuthorizerConfig {
	t.IdentityValidationExpression = cfz.V(v)
	return t
}
