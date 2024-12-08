// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest__Type is the CloudFormation type for AWS::EC2::NetworkInsightsAccessScope.ResourceStatementRequest.
	AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest__Type = "AWS::EC2::NetworkInsightsAccessScope.ResourceStatementRequest"
)

var (
	// AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest__PropertiesMap reports all the CloudFormation properties for AWS::EC2::NetworkInsightsAccessScope.ResourceStatementRequest.
	AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest__PropertiesMap = struct {
		ResourceTypes string
		Resources     string
	}{
		ResourceTypes: "ResourceTypes",
		Resources:     "Resources",
	}

	// AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::NetworkInsightsAccessScope.ResourceStatementRequest.
	AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest__PropertiesSlice = []string{
		AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest__PropertiesMap.ResourceTypes,
		AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest__PropertiesMap.Resources,
	}
)

// AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest is a binding for AWS::EC2::NetworkInsightsAccessScope.ResourceStatementRequest.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-resourcestatementrequest.html
type AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest struct {
	// ResourceTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-resourcestatementrequest.html#cfn-ec2-networkinsightsaccessscope-resourcestatementrequest-resourcetypes
	ResourceTypes cfz.ExpressionSlice[string] `json:"ResourceTypes,omitempty"`

	// Resources is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-resourcestatementrequest.html#cfn-ec2-networkinsightsaccessscope-resourcestatementrequest-resources
	Resources cfz.ExpressionSlice[string] `json:"Resources,omitempty"`
}

// New__AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest initializes a new AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest.
func New__AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest() AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest {
	return AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest) GetType() string {
	return AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest__Type
}

// Set__ResourceTypes updates property "ResourceTypes".
func (t AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest) Set__ResourceTypes(v cfz.ExpressionSlice[string]) AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest {
	t.ResourceTypes = v
	return t
}

// SetS__ResourceTypes updates property "ResourceTypes".
func (t AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest) SetS__ResourceTypes(v ...cfz.Expression[string]) AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest {
	t.ResourceTypes = cfz.S(v...)
	return t
}

// SetSV__ResourceTypes updates property "ResourceTypes".
func (t AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest) SetSV__ResourceTypes(v ...string) AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest {
	t.ResourceTypes = cfz.SV(v...)
	return t
}

// Set__Resources updates property "Resources".
func (t AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest) Set__Resources(v cfz.ExpressionSlice[string]) AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest {
	t.Resources = v
	return t
}

// SetS__Resources updates property "Resources".
func (t AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest) SetS__Resources(v ...cfz.Expression[string]) AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest {
	t.Resources = cfz.S(v...)
	return t
}

// SetSV__Resources updates property "Resources".
func (t AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest) SetSV__Resources(v ...string) AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest {
	t.Resources = cfz.SV(v...)
	return t
}
