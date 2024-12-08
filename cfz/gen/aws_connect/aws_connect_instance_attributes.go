// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_connect

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Connect_Instance_Attributes__Type is the CloudFormation type for AWS::Connect::Instance.Attributes.
	AWS_Connect_Instance_Attributes__Type = "AWS::Connect::Instance.Attributes"
)

var (
	// AWS_Connect_Instance_Attributes__PropertiesMap reports all the CloudFormation properties for AWS::Connect::Instance.Attributes.
	AWS_Connect_Instance_Attributes__PropertiesMap = struct {
		AutoResolveBestVoices string
		ContactLens           string
		ContactflowLogs       string
		EarlyMedia            string
		InboundCalls          string
		OutboundCalls         string
		UseCustomTTSVoices    string
	}{
		AutoResolveBestVoices: "AutoResolveBestVoices",
		ContactLens:           "ContactLens",
		ContactflowLogs:       "ContactflowLogs",
		EarlyMedia:            "EarlyMedia",
		InboundCalls:          "InboundCalls",
		OutboundCalls:         "OutboundCalls",
		UseCustomTTSVoices:    "UseCustomTTSVoices",
	}

	// AWS_Connect_Instance_Attributes__PropertiesSlice reports all the CloudFormation properties for AWS::Connect::Instance.Attributes.
	AWS_Connect_Instance_Attributes__PropertiesSlice = []string{
		AWS_Connect_Instance_Attributes__PropertiesMap.AutoResolveBestVoices,
		AWS_Connect_Instance_Attributes__PropertiesMap.ContactLens,
		AWS_Connect_Instance_Attributes__PropertiesMap.ContactflowLogs,
		AWS_Connect_Instance_Attributes__PropertiesMap.EarlyMedia,
		AWS_Connect_Instance_Attributes__PropertiesMap.InboundCalls,
		AWS_Connect_Instance_Attributes__PropertiesMap.OutboundCalls,
		AWS_Connect_Instance_Attributes__PropertiesMap.UseCustomTTSVoices,
	}
)

// AWS_Connect_Instance_Attributes is a binding for AWS::Connect::Instance.Attributes.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html
type AWS_Connect_Instance_Attributes struct {
	// AutoResolveBestVoices is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-autoresolvebestvoices
	AutoResolveBestVoices cfz.Expression[bool] `json:"AutoResolveBestVoices,omitempty"`

	// ContactLens is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-contactlens
	ContactLens cfz.Expression[bool] `json:"ContactLens,omitempty"`

	// ContactflowLogs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-contactflowlogs
	ContactflowLogs cfz.Expression[bool] `json:"ContactflowLogs,omitempty"`

	// EarlyMedia is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-earlymedia
	EarlyMedia cfz.Expression[bool] `json:"EarlyMedia,omitempty"`

	// InboundCalls is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-inboundcalls
	InboundCalls cfz.Expression[bool] `json:"InboundCalls,omitempty"`

	// OutboundCalls is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-outboundcalls
	OutboundCalls cfz.Expression[bool] `json:"OutboundCalls,omitempty"`

	// UseCustomTTSVoices is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-usecustomttsvoices
	UseCustomTTSVoices cfz.Expression[bool] `json:"UseCustomTTSVoices,omitempty"`
}

// New__AWS_Connect_Instance_Attributes initializes a new AWS_Connect_Instance_Attributes.
func New__AWS_Connect_Instance_Attributes() AWS_Connect_Instance_Attributes {
	return AWS_Connect_Instance_Attributes{}
}

// GetType returns the CloudFormation type.
func (AWS_Connect_Instance_Attributes) GetType() string {
	return AWS_Connect_Instance_Attributes__Type
}

// Set__AutoResolveBestVoices updates property "AutoResolveBestVoices".
func (t AWS_Connect_Instance_Attributes) Set__AutoResolveBestVoices(v cfz.Expression[bool]) AWS_Connect_Instance_Attributes {
	t.AutoResolveBestVoices = v
	return t
}

// SetV__AutoResolveBestVoices updates property "AutoResolveBestVoices".
func (t AWS_Connect_Instance_Attributes) SetV__AutoResolveBestVoices(v bool) AWS_Connect_Instance_Attributes {
	t.AutoResolveBestVoices = cfz.V(v)
	return t
}

// Set__ContactLens updates property "ContactLens".
func (t AWS_Connect_Instance_Attributes) Set__ContactLens(v cfz.Expression[bool]) AWS_Connect_Instance_Attributes {
	t.ContactLens = v
	return t
}

// SetV__ContactLens updates property "ContactLens".
func (t AWS_Connect_Instance_Attributes) SetV__ContactLens(v bool) AWS_Connect_Instance_Attributes {
	t.ContactLens = cfz.V(v)
	return t
}

// Set__ContactflowLogs updates property "ContactflowLogs".
func (t AWS_Connect_Instance_Attributes) Set__ContactflowLogs(v cfz.Expression[bool]) AWS_Connect_Instance_Attributes {
	t.ContactflowLogs = v
	return t
}

// SetV__ContactflowLogs updates property "ContactflowLogs".
func (t AWS_Connect_Instance_Attributes) SetV__ContactflowLogs(v bool) AWS_Connect_Instance_Attributes {
	t.ContactflowLogs = cfz.V(v)
	return t
}

// Set__EarlyMedia updates property "EarlyMedia".
func (t AWS_Connect_Instance_Attributes) Set__EarlyMedia(v cfz.Expression[bool]) AWS_Connect_Instance_Attributes {
	t.EarlyMedia = v
	return t
}

// SetV__EarlyMedia updates property "EarlyMedia".
func (t AWS_Connect_Instance_Attributes) SetV__EarlyMedia(v bool) AWS_Connect_Instance_Attributes {
	t.EarlyMedia = cfz.V(v)
	return t
}

// Set__InboundCalls updates property "InboundCalls".
func (t AWS_Connect_Instance_Attributes) Set__InboundCalls(v cfz.Expression[bool]) AWS_Connect_Instance_Attributes {
	t.InboundCalls = v
	return t
}

// SetV__InboundCalls updates property "InboundCalls".
func (t AWS_Connect_Instance_Attributes) SetV__InboundCalls(v bool) AWS_Connect_Instance_Attributes {
	t.InboundCalls = cfz.V(v)
	return t
}

// Set__OutboundCalls updates property "OutboundCalls".
func (t AWS_Connect_Instance_Attributes) Set__OutboundCalls(v cfz.Expression[bool]) AWS_Connect_Instance_Attributes {
	t.OutboundCalls = v
	return t
}

// SetV__OutboundCalls updates property "OutboundCalls".
func (t AWS_Connect_Instance_Attributes) SetV__OutboundCalls(v bool) AWS_Connect_Instance_Attributes {
	t.OutboundCalls = cfz.V(v)
	return t
}

// Set__UseCustomTTSVoices updates property "UseCustomTTSVoices".
func (t AWS_Connect_Instance_Attributes) Set__UseCustomTTSVoices(v cfz.Expression[bool]) AWS_Connect_Instance_Attributes {
	t.UseCustomTTSVoices = v
	return t
}

// SetV__UseCustomTTSVoices updates property "UseCustomTTSVoices".
func (t AWS_Connect_Instance_Attributes) SetV__UseCustomTTSVoices(v bool) AWS_Connect_Instance_Attributes {
	t.UseCustomTTSVoices = cfz.V(v)
	return t
}
