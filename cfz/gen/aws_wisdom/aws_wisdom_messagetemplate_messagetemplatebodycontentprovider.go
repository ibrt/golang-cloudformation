// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wisdom

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider__Type is the CloudFormation type for AWS::Wisdom::MessageTemplate.MessageTemplateBodyContentProvider.
	AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider__Type = "AWS::Wisdom::MessageTemplate.MessageTemplateBodyContentProvider"
)

var (
	// AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider__PropertiesMap reports all the CloudFormation properties for AWS::Wisdom::MessageTemplate.MessageTemplateBodyContentProvider.
	AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider__PropertiesMap = struct {
		Content string
	}{
		Content: "Content",
	}

	// AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider__PropertiesSlice reports all the CloudFormation properties for AWS::Wisdom::MessageTemplate.MessageTemplateBodyContentProvider.
	AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider__PropertiesSlice = []string{
		AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider__PropertiesMap.Content,
	}
)

// AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider is a binding for AWS::Wisdom::MessageTemplate.MessageTemplateBodyContentProvider.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplatebodycontentprovider.html
type AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider struct {
	// Content is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-messagetemplatebodycontentprovider.html#cfn-wisdom-messagetemplate-messagetemplatebodycontentprovider-content
	Content cfz.Expression[string] `json:"Content,omitempty"`
}

// New__AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider initializes a new AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider.
func New__AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider() AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider {
	return AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider{}
}

// GetType returns the CloudFormation type.
func (AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider) GetType() string {
	return AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider__Type
}

// Set__Content updates property "Content".
func (t AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider) Set__Content(v cfz.Expression[string]) AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider {
	t.Content = v
	return t
}

// SetV__Content updates property "Content".
func (t AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider) SetV__Content(v string) AWS_Wisdom_MessageTemplate_MessageTemplateBodyContentProvider {
	t.Content = cfz.V(v)
	return t
}
