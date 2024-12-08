// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediaconnect

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaConnect_Flow_Fmtp__Type is the CloudFormation type for AWS::MediaConnect::Flow.Fmtp.
	AWS_MediaConnect_Flow_Fmtp__Type = "AWS::MediaConnect::Flow.Fmtp"
)

var (
	// AWS_MediaConnect_Flow_Fmtp__PropertiesMap reports all the CloudFormation properties for AWS::MediaConnect::Flow.Fmtp.
	AWS_MediaConnect_Flow_Fmtp__PropertiesMap = struct {
		ChannelOrder   string
		Colorimetry    string
		ExactFramerate string
		Par            string
		Range          string
		ScanMode       string
		Tcs            string
	}{
		ChannelOrder:   "ChannelOrder",
		Colorimetry:    "Colorimetry",
		ExactFramerate: "ExactFramerate",
		Par:            "Par",
		Range:          "Range",
		ScanMode:       "ScanMode",
		Tcs:            "Tcs",
	}

	// AWS_MediaConnect_Flow_Fmtp__PropertiesSlice reports all the CloudFormation properties for AWS::MediaConnect::Flow.Fmtp.
	AWS_MediaConnect_Flow_Fmtp__PropertiesSlice = []string{
		AWS_MediaConnect_Flow_Fmtp__PropertiesMap.ChannelOrder,
		AWS_MediaConnect_Flow_Fmtp__PropertiesMap.Colorimetry,
		AWS_MediaConnect_Flow_Fmtp__PropertiesMap.ExactFramerate,
		AWS_MediaConnect_Flow_Fmtp__PropertiesMap.Par,
		AWS_MediaConnect_Flow_Fmtp__PropertiesMap.Range,
		AWS_MediaConnect_Flow_Fmtp__PropertiesMap.ScanMode,
		AWS_MediaConnect_Flow_Fmtp__PropertiesMap.Tcs,
	}
)

// AWS_MediaConnect_Flow_Fmtp is a binding for AWS::MediaConnect::Flow.Fmtp.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html
type AWS_MediaConnect_Flow_Fmtp struct {
	// ChannelOrder is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-channelorder
	ChannelOrder cfz.Expression[string] `json:"ChannelOrder,omitempty"`

	// Colorimetry is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-colorimetry
	Colorimetry cfz.Expression[string] `json:"Colorimetry,omitempty"`

	// ExactFramerate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-exactframerate
	ExactFramerate cfz.Expression[string] `json:"ExactFramerate,omitempty"`

	// Par is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-par
	Par cfz.Expression[string] `json:"Par,omitempty"`

	// Range is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-range
	Range cfz.Expression[string] `json:"Range,omitempty"`

	// ScanMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-scanmode
	ScanMode cfz.Expression[string] `json:"ScanMode,omitempty"`

	// Tcs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-fmtp.html#cfn-mediaconnect-flow-fmtp-tcs
	Tcs cfz.Expression[string] `json:"Tcs,omitempty"`
}

// New__AWS_MediaConnect_Flow_Fmtp initializes a new AWS_MediaConnect_Flow_Fmtp.
func New__AWS_MediaConnect_Flow_Fmtp() AWS_MediaConnect_Flow_Fmtp {
	return AWS_MediaConnect_Flow_Fmtp{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaConnect_Flow_Fmtp) GetType() string {
	return AWS_MediaConnect_Flow_Fmtp__Type
}

// Set__ChannelOrder updates property "ChannelOrder".
func (t AWS_MediaConnect_Flow_Fmtp) Set__ChannelOrder(v cfz.Expression[string]) AWS_MediaConnect_Flow_Fmtp {
	t.ChannelOrder = v
	return t
}

// SetV__ChannelOrder updates property "ChannelOrder".
func (t AWS_MediaConnect_Flow_Fmtp) SetV__ChannelOrder(v string) AWS_MediaConnect_Flow_Fmtp {
	t.ChannelOrder = cfz.V(v)
	return t
}

// Set__Colorimetry updates property "Colorimetry".
func (t AWS_MediaConnect_Flow_Fmtp) Set__Colorimetry(v cfz.Expression[string]) AWS_MediaConnect_Flow_Fmtp {
	t.Colorimetry = v
	return t
}

// SetV__Colorimetry updates property "Colorimetry".
func (t AWS_MediaConnect_Flow_Fmtp) SetV__Colorimetry(v string) AWS_MediaConnect_Flow_Fmtp {
	t.Colorimetry = cfz.V(v)
	return t
}

// Set__ExactFramerate updates property "ExactFramerate".
func (t AWS_MediaConnect_Flow_Fmtp) Set__ExactFramerate(v cfz.Expression[string]) AWS_MediaConnect_Flow_Fmtp {
	t.ExactFramerate = v
	return t
}

// SetV__ExactFramerate updates property "ExactFramerate".
func (t AWS_MediaConnect_Flow_Fmtp) SetV__ExactFramerate(v string) AWS_MediaConnect_Flow_Fmtp {
	t.ExactFramerate = cfz.V(v)
	return t
}

// Set__Par updates property "Par".
func (t AWS_MediaConnect_Flow_Fmtp) Set__Par(v cfz.Expression[string]) AWS_MediaConnect_Flow_Fmtp {
	t.Par = v
	return t
}

// SetV__Par updates property "Par".
func (t AWS_MediaConnect_Flow_Fmtp) SetV__Par(v string) AWS_MediaConnect_Flow_Fmtp {
	t.Par = cfz.V(v)
	return t
}

// Set__Range updates property "Range".
func (t AWS_MediaConnect_Flow_Fmtp) Set__Range(v cfz.Expression[string]) AWS_MediaConnect_Flow_Fmtp {
	t.Range = v
	return t
}

// SetV__Range updates property "Range".
func (t AWS_MediaConnect_Flow_Fmtp) SetV__Range(v string) AWS_MediaConnect_Flow_Fmtp {
	t.Range = cfz.V(v)
	return t
}

// Set__ScanMode updates property "ScanMode".
func (t AWS_MediaConnect_Flow_Fmtp) Set__ScanMode(v cfz.Expression[string]) AWS_MediaConnect_Flow_Fmtp {
	t.ScanMode = v
	return t
}

// SetV__ScanMode updates property "ScanMode".
func (t AWS_MediaConnect_Flow_Fmtp) SetV__ScanMode(v string) AWS_MediaConnect_Flow_Fmtp {
	t.ScanMode = cfz.V(v)
	return t
}

// Set__Tcs updates property "Tcs".
func (t AWS_MediaConnect_Flow_Fmtp) Set__Tcs(v cfz.Expression[string]) AWS_MediaConnect_Flow_Fmtp {
	t.Tcs = v
	return t
}

// SetV__Tcs updates property "Tcs".
func (t AWS_MediaConnect_Flow_Fmtp) SetV__Tcs(v string) AWS_MediaConnect_Flow_Fmtp {
	t.Tcs = cfz.V(v)
	return t
}
