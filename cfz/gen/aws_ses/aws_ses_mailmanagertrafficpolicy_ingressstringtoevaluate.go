// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ses

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate__Type is the CloudFormation type for AWS::SES::MailManagerTrafficPolicy.IngressStringToEvaluate.
	AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate__Type = "AWS::SES::MailManagerTrafficPolicy.IngressStringToEvaluate"
)

var (
	// AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate__PropertiesMap reports all the CloudFormation properties for AWS::SES::MailManagerTrafficPolicy.IngressStringToEvaluate.
	AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate__PropertiesMap = struct {
		Attribute string
	}{
		Attribute: "Attribute",
	}

	// AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate__PropertiesSlice reports all the CloudFormation properties for AWS::SES::MailManagerTrafficPolicy.IngressStringToEvaluate.
	AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate__PropertiesSlice = []string{
		AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate__PropertiesMap.Attribute,
	}
)

// AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate is a binding for AWS::SES::MailManagerTrafficPolicy.IngressStringToEvaluate.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressstringtoevaluate.html
type AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate struct {
	// Attribute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressstringtoevaluate.html#cfn-ses-mailmanagertrafficpolicy-ingressstringtoevaluate-attribute
	Attribute cfz.Expression[string] `json:"Attribute,omitempty"`
}

// New__AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate initializes a new AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate.
func New__AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate() AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate {
	return AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate{}
}

// GetType returns the CloudFormation type.
func (AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate) GetType() string {
	return AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate__Type
}

// Set__Attribute updates property "Attribute".
func (t AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate) Set__Attribute(v cfz.Expression[string]) AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate {
	t.Attribute = v
	return t
}

// SetV__Attribute updates property "Attribute".
func (t AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate) SetV__Attribute(v string) AWS_SES_MailManagerTrafficPolicy_IngressStringToEvaluate {
	t.Attribute = cfz.V(v)
	return t
}
