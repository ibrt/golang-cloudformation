// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_MsSmoothGroupSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.MsSmoothGroupSettings.
	AWS_MediaLive_Channel_MsSmoothGroupSettings__Type = "AWS::MediaLive::Channel.MsSmoothGroupSettings"
)

var (
	// AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.MsSmoothGroupSettings.
	AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap = struct {
		AcquisitionPointId       string
		AudioOnlyTimecodeControl string
		CertificateMode          string
		ConnectionRetryInterval  string
		Destination              string
		EventId                  string
		EventIdMode              string
		EventStopBehavior        string
		FilecacheDuration        string
		FragmentLength           string
		InputLossAction          string
		NumRetries               string
		RestartDelay             string
		SegmentationMode         string
		SendDelayMs              string
		SparseTrackType          string
		StreamManifestBehavior   string
		TimestampOffset          string
		TimestampOffsetMode      string
	}{
		AcquisitionPointId:       "AcquisitionPointId",
		AudioOnlyTimecodeControl: "AudioOnlyTimecodeControl",
		CertificateMode:          "CertificateMode",
		ConnectionRetryInterval:  "ConnectionRetryInterval",
		Destination:              "Destination",
		EventId:                  "EventId",
		EventIdMode:              "EventIdMode",
		EventStopBehavior:        "EventStopBehavior",
		FilecacheDuration:        "FilecacheDuration",
		FragmentLength:           "FragmentLength",
		InputLossAction:          "InputLossAction",
		NumRetries:               "NumRetries",
		RestartDelay:             "RestartDelay",
		SegmentationMode:         "SegmentationMode",
		SendDelayMs:              "SendDelayMs",
		SparseTrackType:          "SparseTrackType",
		StreamManifestBehavior:   "StreamManifestBehavior",
		TimestampOffset:          "TimestampOffset",
		TimestampOffsetMode:      "TimestampOffsetMode",
	}

	// AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.MsSmoothGroupSettings.
	AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.AcquisitionPointId,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.AudioOnlyTimecodeControl,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.CertificateMode,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.ConnectionRetryInterval,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.Destination,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.EventId,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.EventIdMode,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.EventStopBehavior,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.FilecacheDuration,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.FragmentLength,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.InputLossAction,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.NumRetries,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.RestartDelay,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.SegmentationMode,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.SendDelayMs,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.SparseTrackType,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.StreamManifestBehavior,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.TimestampOffset,
		AWS_MediaLive_Channel_MsSmoothGroupSettings__PropertiesMap.TimestampOffsetMode,
	}
)

// AWS_MediaLive_Channel_MsSmoothGroupSettings is a binding for AWS::MediaLive::Channel.MsSmoothGroupSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html
type AWS_MediaLive_Channel_MsSmoothGroupSettings struct {
	// AcquisitionPointId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-acquisitionpointid
	AcquisitionPointId cfz.Expression[string] `json:"AcquisitionPointId,omitempty"`

	// AudioOnlyTimecodeControl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-audioonlytimecodecontrol
	AudioOnlyTimecodeControl cfz.Expression[string] `json:"AudioOnlyTimecodeControl,omitempty"`

	// CertificateMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-certificatemode
	CertificateMode cfz.Expression[string] `json:"CertificateMode,omitempty"`

	// ConnectionRetryInterval is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-connectionretryinterval
	ConnectionRetryInterval cfz.Expression[int32] `json:"ConnectionRetryInterval,omitempty"`

	// Destination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-destination
	Destination cfz.Expression[AWS_MediaLive_Channel_OutputLocationRef] `json:"Destination,omitempty"`

	// EventId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-eventid
	EventId cfz.Expression[string] `json:"EventId,omitempty"`

	// EventIdMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-eventidmode
	EventIdMode cfz.Expression[string] `json:"EventIdMode,omitempty"`

	// EventStopBehavior is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-eventstopbehavior
	EventStopBehavior cfz.Expression[string] `json:"EventStopBehavior,omitempty"`

	// FilecacheDuration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-filecacheduration
	FilecacheDuration cfz.Expression[int32] `json:"FilecacheDuration,omitempty"`

	// FragmentLength is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-fragmentlength
	FragmentLength cfz.Expression[int32] `json:"FragmentLength,omitempty"`

	// InputLossAction is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-inputlossaction
	InputLossAction cfz.Expression[string] `json:"InputLossAction,omitempty"`

	// NumRetries is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-numretries
	NumRetries cfz.Expression[int32] `json:"NumRetries,omitempty"`

	// RestartDelay is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-restartdelay
	RestartDelay cfz.Expression[int32] `json:"RestartDelay,omitempty"`

	// SegmentationMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-segmentationmode
	SegmentationMode cfz.Expression[string] `json:"SegmentationMode,omitempty"`

	// SendDelayMs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-senddelayms
	SendDelayMs cfz.Expression[int32] `json:"SendDelayMs,omitempty"`

	// SparseTrackType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-sparsetracktype
	SparseTrackType cfz.Expression[string] `json:"SparseTrackType,omitempty"`

	// StreamManifestBehavior is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-streammanifestbehavior
	StreamManifestBehavior cfz.Expression[string] `json:"StreamManifestBehavior,omitempty"`

	// TimestampOffset is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-timestampoffset
	TimestampOffset cfz.Expression[string] `json:"TimestampOffset,omitempty"`

	// TimestampOffsetMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothgroupsettings.html#cfn-medialive-channel-mssmoothgroupsettings-timestampoffsetmode
	TimestampOffsetMode cfz.Expression[string] `json:"TimestampOffsetMode,omitempty"`
}

// New__AWS_MediaLive_Channel_MsSmoothGroupSettings initializes a new AWS_MediaLive_Channel_MsSmoothGroupSettings.
func New__AWS_MediaLive_Channel_MsSmoothGroupSettings() AWS_MediaLive_Channel_MsSmoothGroupSettings {
	return AWS_MediaLive_Channel_MsSmoothGroupSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_MsSmoothGroupSettings) GetType() string {
	return AWS_MediaLive_Channel_MsSmoothGroupSettings__Type
}

// Set__AcquisitionPointId updates property "AcquisitionPointId".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__AcquisitionPointId(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.AcquisitionPointId = v
	return t
}

// SetV__AcquisitionPointId updates property "AcquisitionPointId".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__AcquisitionPointId(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.AcquisitionPointId = cfz.V(v)
	return t
}

// Set__AudioOnlyTimecodeControl updates property "AudioOnlyTimecodeControl".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__AudioOnlyTimecodeControl(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.AudioOnlyTimecodeControl = v
	return t
}

// SetV__AudioOnlyTimecodeControl updates property "AudioOnlyTimecodeControl".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__AudioOnlyTimecodeControl(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.AudioOnlyTimecodeControl = cfz.V(v)
	return t
}

// Set__CertificateMode updates property "CertificateMode".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__CertificateMode(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.CertificateMode = v
	return t
}

// SetV__CertificateMode updates property "CertificateMode".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__CertificateMode(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.CertificateMode = cfz.V(v)
	return t
}

// Set__ConnectionRetryInterval updates property "ConnectionRetryInterval".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__ConnectionRetryInterval(v cfz.Expression[int32]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.ConnectionRetryInterval = v
	return t
}

// SetV__ConnectionRetryInterval updates property "ConnectionRetryInterval".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__ConnectionRetryInterval(v int32) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.ConnectionRetryInterval = cfz.V(v)
	return t
}

// Set__Destination updates property "Destination".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__Destination(v cfz.Expression[AWS_MediaLive_Channel_OutputLocationRef]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.Destination = v
	return t
}

// SetV__Destination updates property "Destination".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__Destination(v AWS_MediaLive_Channel_OutputLocationRef) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.Destination = cfz.V(v)
	return t
}

// Set__EventId updates property "EventId".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__EventId(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.EventId = v
	return t
}

// SetV__EventId updates property "EventId".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__EventId(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.EventId = cfz.V(v)
	return t
}

// Set__EventIdMode updates property "EventIdMode".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__EventIdMode(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.EventIdMode = v
	return t
}

// SetV__EventIdMode updates property "EventIdMode".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__EventIdMode(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.EventIdMode = cfz.V(v)
	return t
}

// Set__EventStopBehavior updates property "EventStopBehavior".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__EventStopBehavior(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.EventStopBehavior = v
	return t
}

// SetV__EventStopBehavior updates property "EventStopBehavior".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__EventStopBehavior(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.EventStopBehavior = cfz.V(v)
	return t
}

// Set__FilecacheDuration updates property "FilecacheDuration".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__FilecacheDuration(v cfz.Expression[int32]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.FilecacheDuration = v
	return t
}

// SetV__FilecacheDuration updates property "FilecacheDuration".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__FilecacheDuration(v int32) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.FilecacheDuration = cfz.V(v)
	return t
}

// Set__FragmentLength updates property "FragmentLength".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__FragmentLength(v cfz.Expression[int32]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.FragmentLength = v
	return t
}

// SetV__FragmentLength updates property "FragmentLength".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__FragmentLength(v int32) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.FragmentLength = cfz.V(v)
	return t
}

// Set__InputLossAction updates property "InputLossAction".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__InputLossAction(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.InputLossAction = v
	return t
}

// SetV__InputLossAction updates property "InputLossAction".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__InputLossAction(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.InputLossAction = cfz.V(v)
	return t
}

// Set__NumRetries updates property "NumRetries".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__NumRetries(v cfz.Expression[int32]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.NumRetries = v
	return t
}

// SetV__NumRetries updates property "NumRetries".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__NumRetries(v int32) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.NumRetries = cfz.V(v)
	return t
}

// Set__RestartDelay updates property "RestartDelay".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__RestartDelay(v cfz.Expression[int32]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.RestartDelay = v
	return t
}

// SetV__RestartDelay updates property "RestartDelay".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__RestartDelay(v int32) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.RestartDelay = cfz.V(v)
	return t
}

// Set__SegmentationMode updates property "SegmentationMode".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__SegmentationMode(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.SegmentationMode = v
	return t
}

// SetV__SegmentationMode updates property "SegmentationMode".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__SegmentationMode(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.SegmentationMode = cfz.V(v)
	return t
}

// Set__SendDelayMs updates property "SendDelayMs".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__SendDelayMs(v cfz.Expression[int32]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.SendDelayMs = v
	return t
}

// SetV__SendDelayMs updates property "SendDelayMs".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__SendDelayMs(v int32) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.SendDelayMs = cfz.V(v)
	return t
}

// Set__SparseTrackType updates property "SparseTrackType".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__SparseTrackType(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.SparseTrackType = v
	return t
}

// SetV__SparseTrackType updates property "SparseTrackType".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__SparseTrackType(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.SparseTrackType = cfz.V(v)
	return t
}

// Set__StreamManifestBehavior updates property "StreamManifestBehavior".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__StreamManifestBehavior(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.StreamManifestBehavior = v
	return t
}

// SetV__StreamManifestBehavior updates property "StreamManifestBehavior".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__StreamManifestBehavior(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.StreamManifestBehavior = cfz.V(v)
	return t
}

// Set__TimestampOffset updates property "TimestampOffset".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__TimestampOffset(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.TimestampOffset = v
	return t
}

// SetV__TimestampOffset updates property "TimestampOffset".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__TimestampOffset(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.TimestampOffset = cfz.V(v)
	return t
}

// Set__TimestampOffsetMode updates property "TimestampOffsetMode".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) Set__TimestampOffsetMode(v cfz.Expression[string]) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.TimestampOffsetMode = v
	return t
}

// SetV__TimestampOffsetMode updates property "TimestampOffsetMode".
func (t AWS_MediaLive_Channel_MsSmoothGroupSettings) SetV__TimestampOffsetMode(v string) AWS_MediaLive_Channel_MsSmoothGroupSettings {
	t.TimestampOffsetMode = cfz.V(v)
	return t
}
