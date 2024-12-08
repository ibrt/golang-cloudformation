// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ses

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SES_Template_Template__Type is the CloudFormation type for AWS::SES::Template.Template.
	AWS_SES_Template_Template__Type = "AWS::SES::Template.Template"
)

var (
	// AWS_SES_Template_Template__PropertiesMap reports all the CloudFormation properties for AWS::SES::Template.Template.
	AWS_SES_Template_Template__PropertiesMap = struct {
		HtmlPart     string
		SubjectPart  string
		TemplateName string
		TextPart     string
	}{
		HtmlPart:     "HtmlPart",
		SubjectPart:  "SubjectPart",
		TemplateName: "TemplateName",
		TextPart:     "TextPart",
	}

	// AWS_SES_Template_Template__PropertiesSlice reports all the CloudFormation properties for AWS::SES::Template.Template.
	AWS_SES_Template_Template__PropertiesSlice = []string{
		AWS_SES_Template_Template__PropertiesMap.HtmlPart,
		AWS_SES_Template_Template__PropertiesMap.SubjectPart,
		AWS_SES_Template_Template__PropertiesMap.TemplateName,
		AWS_SES_Template_Template__PropertiesMap.TextPart,
	}
)

// AWS_SES_Template_Template is a binding for AWS::SES::Template.Template.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html
type AWS_SES_Template_Template struct {
	// HtmlPart is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-htmlpart
	HtmlPart cfz.Expression[string] `json:"HtmlPart,omitempty"`

	// SubjectPart is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-subjectpart
	SubjectPart cfz.Expression[string] `json:"SubjectPart,omitempty"`

	// TemplateName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-templatename
	TemplateName cfz.Expression[string] `json:"TemplateName,omitempty"`

	// TextPart is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-textpart
	TextPart cfz.Expression[string] `json:"TextPart,omitempty"`
}

// New__AWS_SES_Template_Template initializes a new AWS_SES_Template_Template.
func New__AWS_SES_Template_Template() AWS_SES_Template_Template {
	return AWS_SES_Template_Template{}
}

// GetType returns the CloudFormation type.
func (AWS_SES_Template_Template) GetType() string {
	return AWS_SES_Template_Template__Type
}

// Set__HtmlPart updates property "HtmlPart".
func (t AWS_SES_Template_Template) Set__HtmlPart(v cfz.Expression[string]) AWS_SES_Template_Template {
	t.HtmlPart = v
	return t
}

// SetV__HtmlPart updates property "HtmlPart".
func (t AWS_SES_Template_Template) SetV__HtmlPart(v string) AWS_SES_Template_Template {
	t.HtmlPart = cfz.V(v)
	return t
}

// Set__SubjectPart updates property "SubjectPart".
func (t AWS_SES_Template_Template) Set__SubjectPart(v cfz.Expression[string]) AWS_SES_Template_Template {
	t.SubjectPart = v
	return t
}

// SetV__SubjectPart updates property "SubjectPart".
func (t AWS_SES_Template_Template) SetV__SubjectPart(v string) AWS_SES_Template_Template {
	t.SubjectPart = cfz.V(v)
	return t
}

// Set__TemplateName updates property "TemplateName".
func (t AWS_SES_Template_Template) Set__TemplateName(v cfz.Expression[string]) AWS_SES_Template_Template {
	t.TemplateName = v
	return t
}

// SetV__TemplateName updates property "TemplateName".
func (t AWS_SES_Template_Template) SetV__TemplateName(v string) AWS_SES_Template_Template {
	t.TemplateName = cfz.V(v)
	return t
}

// Set__TextPart updates property "TextPart".
func (t AWS_SES_Template_Template) Set__TextPart(v cfz.Expression[string]) AWS_SES_Template_Template {
	t.TextPart = v
	return t
}

// SetV__TextPart updates property "TextPart".
func (t AWS_SES_Template_Template) SetV__TextPart(v string) AWS_SES_Template_Template {
	t.TextPart = cfz.V(v)
	return t
}
