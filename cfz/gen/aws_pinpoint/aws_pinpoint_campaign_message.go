// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pinpoint

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pinpoint_Campaign_Message__Type is the CloudFormation type for AWS::Pinpoint::Campaign.Message.
	AWS_Pinpoint_Campaign_Message__Type = "AWS::Pinpoint::Campaign.Message"
)

var (
	// AWS_Pinpoint_Campaign_Message__PropertiesMap reports all the CloudFormation properties for AWS::Pinpoint::Campaign.Message.
	AWS_Pinpoint_Campaign_Message__PropertiesMap = struct {
		Action            string
		Body              string
		ImageIconUrl      string
		ImageSmallIconUrl string
		ImageUrl          string
		JsonBody          string
		MediaUrl          string
		RawContent        string
		SilentPush        string
		TimeToLive        string
		Title             string
		Url               string
	}{
		Action:            "Action",
		Body:              "Body",
		ImageIconUrl:      "ImageIconUrl",
		ImageSmallIconUrl: "ImageSmallIconUrl",
		ImageUrl:          "ImageUrl",
		JsonBody:          "JsonBody",
		MediaUrl:          "MediaUrl",
		RawContent:        "RawContent",
		SilentPush:        "SilentPush",
		TimeToLive:        "TimeToLive",
		Title:             "Title",
		Url:               "Url",
	}

	// AWS_Pinpoint_Campaign_Message__PropertiesSlice reports all the CloudFormation properties for AWS::Pinpoint::Campaign.Message.
	AWS_Pinpoint_Campaign_Message__PropertiesSlice = []string{
		AWS_Pinpoint_Campaign_Message__PropertiesMap.Action,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.Body,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.ImageIconUrl,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.ImageSmallIconUrl,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.ImageUrl,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.JsonBody,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.MediaUrl,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.RawContent,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.SilentPush,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.TimeToLive,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.Title,
		AWS_Pinpoint_Campaign_Message__PropertiesMap.Url,
	}
)

// AWS_Pinpoint_Campaign_Message is a binding for AWS::Pinpoint::Campaign.Message.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html
type AWS_Pinpoint_Campaign_Message struct {
	// Action is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-action
	Action cfz.Expression[string] `json:"Action,omitempty"`

	// Body is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-body
	Body cfz.Expression[string] `json:"Body,omitempty"`

	// ImageIconUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-imageiconurl
	ImageIconUrl cfz.Expression[string] `json:"ImageIconUrl,omitempty"`

	// ImageSmallIconUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-imagesmalliconurl
	ImageSmallIconUrl cfz.Expression[string] `json:"ImageSmallIconUrl,omitempty"`

	// ImageUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-imageurl
	ImageUrl cfz.Expression[string] `json:"ImageUrl,omitempty"`

	// JsonBody is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-jsonbody
	JsonBody cfz.Expression[string] `json:"JsonBody,omitempty"`

	// MediaUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-mediaurl
	MediaUrl cfz.Expression[string] `json:"MediaUrl,omitempty"`

	// RawContent is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-rawcontent
	RawContent cfz.Expression[string] `json:"RawContent,omitempty"`

	// SilentPush is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-silentpush
	SilentPush cfz.Expression[bool] `json:"SilentPush,omitempty"`

	// TimeToLive is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-timetolive
	TimeToLive cfz.Expression[int32] `json:"TimeToLive,omitempty"`

	// Title is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-title
	Title cfz.Expression[string] `json:"Title,omitempty"`

	// Url is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html#cfn-pinpoint-campaign-message-url
	Url cfz.Expression[string] `json:"Url,omitempty"`
}

// New__AWS_Pinpoint_Campaign_Message initializes a new AWS_Pinpoint_Campaign_Message.
func New__AWS_Pinpoint_Campaign_Message() AWS_Pinpoint_Campaign_Message {
	return AWS_Pinpoint_Campaign_Message{}
}

// GetType returns the CloudFormation type.
func (AWS_Pinpoint_Campaign_Message) GetType() string {
	return AWS_Pinpoint_Campaign_Message__Type
}

// Set__Action updates property "Action".
func (t AWS_Pinpoint_Campaign_Message) Set__Action(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.Action = v
	return t
}

// SetV__Action updates property "Action".
func (t AWS_Pinpoint_Campaign_Message) SetV__Action(v string) AWS_Pinpoint_Campaign_Message {
	t.Action = cfz.V(v)
	return t
}

// Set__Body updates property "Body".
func (t AWS_Pinpoint_Campaign_Message) Set__Body(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.Body = v
	return t
}

// SetV__Body updates property "Body".
func (t AWS_Pinpoint_Campaign_Message) SetV__Body(v string) AWS_Pinpoint_Campaign_Message {
	t.Body = cfz.V(v)
	return t
}

// Set__ImageIconUrl updates property "ImageIconUrl".
func (t AWS_Pinpoint_Campaign_Message) Set__ImageIconUrl(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.ImageIconUrl = v
	return t
}

// SetV__ImageIconUrl updates property "ImageIconUrl".
func (t AWS_Pinpoint_Campaign_Message) SetV__ImageIconUrl(v string) AWS_Pinpoint_Campaign_Message {
	t.ImageIconUrl = cfz.V(v)
	return t
}

// Set__ImageSmallIconUrl updates property "ImageSmallIconUrl".
func (t AWS_Pinpoint_Campaign_Message) Set__ImageSmallIconUrl(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.ImageSmallIconUrl = v
	return t
}

// SetV__ImageSmallIconUrl updates property "ImageSmallIconUrl".
func (t AWS_Pinpoint_Campaign_Message) SetV__ImageSmallIconUrl(v string) AWS_Pinpoint_Campaign_Message {
	t.ImageSmallIconUrl = cfz.V(v)
	return t
}

// Set__ImageUrl updates property "ImageUrl".
func (t AWS_Pinpoint_Campaign_Message) Set__ImageUrl(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.ImageUrl = v
	return t
}

// SetV__ImageUrl updates property "ImageUrl".
func (t AWS_Pinpoint_Campaign_Message) SetV__ImageUrl(v string) AWS_Pinpoint_Campaign_Message {
	t.ImageUrl = cfz.V(v)
	return t
}

// Set__JsonBody updates property "JsonBody".
func (t AWS_Pinpoint_Campaign_Message) Set__JsonBody(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.JsonBody = v
	return t
}

// SetV__JsonBody updates property "JsonBody".
func (t AWS_Pinpoint_Campaign_Message) SetV__JsonBody(v string) AWS_Pinpoint_Campaign_Message {
	t.JsonBody = cfz.V(v)
	return t
}

// Set__MediaUrl updates property "MediaUrl".
func (t AWS_Pinpoint_Campaign_Message) Set__MediaUrl(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.MediaUrl = v
	return t
}

// SetV__MediaUrl updates property "MediaUrl".
func (t AWS_Pinpoint_Campaign_Message) SetV__MediaUrl(v string) AWS_Pinpoint_Campaign_Message {
	t.MediaUrl = cfz.V(v)
	return t
}

// Set__RawContent updates property "RawContent".
func (t AWS_Pinpoint_Campaign_Message) Set__RawContent(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.RawContent = v
	return t
}

// SetV__RawContent updates property "RawContent".
func (t AWS_Pinpoint_Campaign_Message) SetV__RawContent(v string) AWS_Pinpoint_Campaign_Message {
	t.RawContent = cfz.V(v)
	return t
}

// Set__SilentPush updates property "SilentPush".
func (t AWS_Pinpoint_Campaign_Message) Set__SilentPush(v cfz.Expression[bool]) AWS_Pinpoint_Campaign_Message {
	t.SilentPush = v
	return t
}

// SetV__SilentPush updates property "SilentPush".
func (t AWS_Pinpoint_Campaign_Message) SetV__SilentPush(v bool) AWS_Pinpoint_Campaign_Message {
	t.SilentPush = cfz.V(v)
	return t
}

// Set__TimeToLive updates property "TimeToLive".
func (t AWS_Pinpoint_Campaign_Message) Set__TimeToLive(v cfz.Expression[int32]) AWS_Pinpoint_Campaign_Message {
	t.TimeToLive = v
	return t
}

// SetV__TimeToLive updates property "TimeToLive".
func (t AWS_Pinpoint_Campaign_Message) SetV__TimeToLive(v int32) AWS_Pinpoint_Campaign_Message {
	t.TimeToLive = cfz.V(v)
	return t
}

// Set__Title updates property "Title".
func (t AWS_Pinpoint_Campaign_Message) Set__Title(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.Title = v
	return t
}

// SetV__Title updates property "Title".
func (t AWS_Pinpoint_Campaign_Message) SetV__Title(v string) AWS_Pinpoint_Campaign_Message {
	t.Title = cfz.V(v)
	return t
}

// Set__Url updates property "Url".
func (t AWS_Pinpoint_Campaign_Message) Set__Url(v cfz.Expression[string]) AWS_Pinpoint_Campaign_Message {
	t.Url = v
	return t
}

// SetV__Url updates property "Url".
func (t AWS_Pinpoint_Campaign_Message) SetV__Url(v string) AWS_Pinpoint_Campaign_Message {
	t.Url = cfz.V(v)
	return t
}
