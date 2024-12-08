// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pinpoint

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pinpoint_Campaign_InAppMessageContent__Type is the CloudFormation type for AWS::Pinpoint::Campaign.InAppMessageContent.
	AWS_Pinpoint_Campaign_InAppMessageContent__Type = "AWS::Pinpoint::Campaign.InAppMessageContent"
)

var (
	// AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesMap reports all the CloudFormation properties for AWS::Pinpoint::Campaign.InAppMessageContent.
	AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesMap = struct {
		BackgroundColor string
		BodyConfig      string
		HeaderConfig    string
		ImageUrl        string
		PrimaryBtn      string
		SecondaryBtn    string
	}{
		BackgroundColor: "BackgroundColor",
		BodyConfig:      "BodyConfig",
		HeaderConfig:    "HeaderConfig",
		ImageUrl:        "ImageUrl",
		PrimaryBtn:      "PrimaryBtn",
		SecondaryBtn:    "SecondaryBtn",
	}

	// AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesSlice reports all the CloudFormation properties for AWS::Pinpoint::Campaign.InAppMessageContent.
	AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesSlice = []string{
		AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesMap.BackgroundColor,
		AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesMap.BodyConfig,
		AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesMap.HeaderConfig,
		AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesMap.ImageUrl,
		AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesMap.PrimaryBtn,
		AWS_Pinpoint_Campaign_InAppMessageContent__PropertiesMap.SecondaryBtn,
	}
)

// AWS_Pinpoint_Campaign_InAppMessageContent is a binding for AWS::Pinpoint::Campaign.InAppMessageContent.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html
type AWS_Pinpoint_Campaign_InAppMessageContent struct {
	// BackgroundColor is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-backgroundcolor
	BackgroundColor cfz.Expression[string] `json:"BackgroundColor,omitempty"`

	// BodyConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-bodyconfig
	BodyConfig cfz.Expression[AWS_Pinpoint_Campaign_InAppMessageBodyConfig] `json:"BodyConfig,omitempty"`

	// HeaderConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-headerconfig
	HeaderConfig cfz.Expression[AWS_Pinpoint_Campaign_InAppMessageHeaderConfig] `json:"HeaderConfig,omitempty"`

	// ImageUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-imageurl
	ImageUrl cfz.Expression[string] `json:"ImageUrl,omitempty"`

	// PrimaryBtn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-primarybtn
	PrimaryBtn cfz.Expression[AWS_Pinpoint_Campaign_InAppMessageButton] `json:"PrimaryBtn,omitempty"`

	// SecondaryBtn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-secondarybtn
	SecondaryBtn cfz.Expression[AWS_Pinpoint_Campaign_InAppMessageButton] `json:"SecondaryBtn,omitempty"`
}

// New__AWS_Pinpoint_Campaign_InAppMessageContent initializes a new AWS_Pinpoint_Campaign_InAppMessageContent.
func New__AWS_Pinpoint_Campaign_InAppMessageContent() AWS_Pinpoint_Campaign_InAppMessageContent {
	return AWS_Pinpoint_Campaign_InAppMessageContent{}
}

// GetType returns the CloudFormation type.
func (AWS_Pinpoint_Campaign_InAppMessageContent) GetType() string {
	return AWS_Pinpoint_Campaign_InAppMessageContent__Type
}

// Set__BackgroundColor updates property "BackgroundColor".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) Set__BackgroundColor(v cfz.Expression[string]) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.BackgroundColor = v
	return t
}

// SetV__BackgroundColor updates property "BackgroundColor".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) SetV__BackgroundColor(v string) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.BackgroundColor = cfz.V(v)
	return t
}

// Set__BodyConfig updates property "BodyConfig".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) Set__BodyConfig(v cfz.Expression[AWS_Pinpoint_Campaign_InAppMessageBodyConfig]) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.BodyConfig = v
	return t
}

// SetV__BodyConfig updates property "BodyConfig".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) SetV__BodyConfig(v AWS_Pinpoint_Campaign_InAppMessageBodyConfig) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.BodyConfig = cfz.V(v)
	return t
}

// Set__HeaderConfig updates property "HeaderConfig".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) Set__HeaderConfig(v cfz.Expression[AWS_Pinpoint_Campaign_InAppMessageHeaderConfig]) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.HeaderConfig = v
	return t
}

// SetV__HeaderConfig updates property "HeaderConfig".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) SetV__HeaderConfig(v AWS_Pinpoint_Campaign_InAppMessageHeaderConfig) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.HeaderConfig = cfz.V(v)
	return t
}

// Set__ImageUrl updates property "ImageUrl".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) Set__ImageUrl(v cfz.Expression[string]) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.ImageUrl = v
	return t
}

// SetV__ImageUrl updates property "ImageUrl".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) SetV__ImageUrl(v string) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.ImageUrl = cfz.V(v)
	return t
}

// Set__PrimaryBtn updates property "PrimaryBtn".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) Set__PrimaryBtn(v cfz.Expression[AWS_Pinpoint_Campaign_InAppMessageButton]) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.PrimaryBtn = v
	return t
}

// SetV__PrimaryBtn updates property "PrimaryBtn".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) SetV__PrimaryBtn(v AWS_Pinpoint_Campaign_InAppMessageButton) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.PrimaryBtn = cfz.V(v)
	return t
}

// Set__SecondaryBtn updates property "SecondaryBtn".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) Set__SecondaryBtn(v cfz.Expression[AWS_Pinpoint_Campaign_InAppMessageButton]) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.SecondaryBtn = v
	return t
}

// SetV__SecondaryBtn updates property "SecondaryBtn".
func (t AWS_Pinpoint_Campaign_InAppMessageContent) SetV__SecondaryBtn(v AWS_Pinpoint_Campaign_InAppMessageButton) AWS_Pinpoint_Campaign_InAppMessageContent {
	t.SecondaryBtn = cfz.V(v)
	return t
}
