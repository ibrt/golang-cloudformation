// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apigateway

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ApiGateway_DomainName_EndpointConfiguration__Type is the CloudFormation type for AWS::ApiGateway::DomainName.EndpointConfiguration.
	AWS_ApiGateway_DomainName_EndpointConfiguration__Type = "AWS::ApiGateway::DomainName.EndpointConfiguration"
)

var (
	// AWS_ApiGateway_DomainName_EndpointConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::ApiGateway::DomainName.EndpointConfiguration.
	AWS_ApiGateway_DomainName_EndpointConfiguration__PropertiesMap = struct {
		Types string
	}{
		Types: "Types",
	}

	// AWS_ApiGateway_DomainName_EndpointConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::ApiGateway::DomainName.EndpointConfiguration.
	AWS_ApiGateway_DomainName_EndpointConfiguration__PropertiesSlice = []string{
		AWS_ApiGateway_DomainName_EndpointConfiguration__PropertiesMap.Types,
	}
)

// AWS_ApiGateway_DomainName_EndpointConfiguration is a binding for AWS::ApiGateway::DomainName.EndpointConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html
type AWS_ApiGateway_DomainName_EndpointConfiguration struct {
	// Types is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html#cfn-apigateway-domainname-endpointconfiguration-types
	Types cfz.ExpressionSlice[string] `json:"Types,omitempty"`
}

// New__AWS_ApiGateway_DomainName_EndpointConfiguration initializes a new AWS_ApiGateway_DomainName_EndpointConfiguration.
func New__AWS_ApiGateway_DomainName_EndpointConfiguration() AWS_ApiGateway_DomainName_EndpointConfiguration {
	return AWS_ApiGateway_DomainName_EndpointConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_ApiGateway_DomainName_EndpointConfiguration) GetType() string {
	return AWS_ApiGateway_DomainName_EndpointConfiguration__Type
}

// Set__Types updates property "Types".
func (t AWS_ApiGateway_DomainName_EndpointConfiguration) Set__Types(v cfz.ExpressionSlice[string]) AWS_ApiGateway_DomainName_EndpointConfiguration {
	t.Types = v
	return t
}

// SetS__Types updates property "Types".
func (t AWS_ApiGateway_DomainName_EndpointConfiguration) SetS__Types(v ...cfz.Expression[string]) AWS_ApiGateway_DomainName_EndpointConfiguration {
	t.Types = cfz.S(v...)
	return t
}

// SetSV__Types updates property "Types".
func (t AWS_ApiGateway_DomainName_EndpointConfiguration) SetSV__Types(v ...string) AWS_ApiGateway_DomainName_EndpointConfiguration {
	t.Types = cfz.SV(v...)
	return t
}
