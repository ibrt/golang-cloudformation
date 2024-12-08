// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apigateway

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ApiGateway_UsagePlan_QuotaSettings__Type is the CloudFormation type for AWS::ApiGateway::UsagePlan.QuotaSettings.
	AWS_ApiGateway_UsagePlan_QuotaSettings__Type = "AWS::ApiGateway::UsagePlan.QuotaSettings"
)

var (
	// AWS_ApiGateway_UsagePlan_QuotaSettings__PropertiesMap reports all the CloudFormation properties for AWS::ApiGateway::UsagePlan.QuotaSettings.
	AWS_ApiGateway_UsagePlan_QuotaSettings__PropertiesMap = struct {
		Limit  string
		Offset string
		Period string
	}{
		Limit:  "Limit",
		Offset: "Offset",
		Period: "Period",
	}

	// AWS_ApiGateway_UsagePlan_QuotaSettings__PropertiesSlice reports all the CloudFormation properties for AWS::ApiGateway::UsagePlan.QuotaSettings.
	AWS_ApiGateway_UsagePlan_QuotaSettings__PropertiesSlice = []string{
		AWS_ApiGateway_UsagePlan_QuotaSettings__PropertiesMap.Limit,
		AWS_ApiGateway_UsagePlan_QuotaSettings__PropertiesMap.Offset,
		AWS_ApiGateway_UsagePlan_QuotaSettings__PropertiesMap.Period,
	}
)

// AWS_ApiGateway_UsagePlan_QuotaSettings is a binding for AWS::ApiGateway::UsagePlan.QuotaSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html
type AWS_ApiGateway_UsagePlan_QuotaSettings struct {
	// Limit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-limit
	Limit cfz.Expression[int32] `json:"Limit,omitempty"`

	// Offset is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-offset
	Offset cfz.Expression[int32] `json:"Offset,omitempty"`

	// Period is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-period
	Period cfz.Expression[string] `json:"Period,omitempty"`
}

// New__AWS_ApiGateway_UsagePlan_QuotaSettings initializes a new AWS_ApiGateway_UsagePlan_QuotaSettings.
func New__AWS_ApiGateway_UsagePlan_QuotaSettings() AWS_ApiGateway_UsagePlan_QuotaSettings {
	return AWS_ApiGateway_UsagePlan_QuotaSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_ApiGateway_UsagePlan_QuotaSettings) GetType() string {
	return AWS_ApiGateway_UsagePlan_QuotaSettings__Type
}

// Set__Limit updates property "Limit".
func (t AWS_ApiGateway_UsagePlan_QuotaSettings) Set__Limit(v cfz.Expression[int32]) AWS_ApiGateway_UsagePlan_QuotaSettings {
	t.Limit = v
	return t
}

// SetV__Limit updates property "Limit".
func (t AWS_ApiGateway_UsagePlan_QuotaSettings) SetV__Limit(v int32) AWS_ApiGateway_UsagePlan_QuotaSettings {
	t.Limit = cfz.V(v)
	return t
}

// Set__Offset updates property "Offset".
func (t AWS_ApiGateway_UsagePlan_QuotaSettings) Set__Offset(v cfz.Expression[int32]) AWS_ApiGateway_UsagePlan_QuotaSettings {
	t.Offset = v
	return t
}

// SetV__Offset updates property "Offset".
func (t AWS_ApiGateway_UsagePlan_QuotaSettings) SetV__Offset(v int32) AWS_ApiGateway_UsagePlan_QuotaSettings {
	t.Offset = cfz.V(v)
	return t
}

// Set__Period updates property "Period".
func (t AWS_ApiGateway_UsagePlan_QuotaSettings) Set__Period(v cfz.Expression[string]) AWS_ApiGateway_UsagePlan_QuotaSettings {
	t.Period = v
	return t
}

// SetV__Period updates property "Period".
func (t AWS_ApiGateway_UsagePlan_QuotaSettings) SetV__Period(v string) AWS_ApiGateway_UsagePlan_QuotaSettings {
	t.Period = cfz.V(v)
	return t
}
