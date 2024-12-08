// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoT_TopicRule_AssetPropertyTimestamp__Type is the CloudFormation type for AWS::IoT::TopicRule.AssetPropertyTimestamp.
	AWS_IoT_TopicRule_AssetPropertyTimestamp__Type = "AWS::IoT::TopicRule.AssetPropertyTimestamp"
)

var (
	// AWS_IoT_TopicRule_AssetPropertyTimestamp__PropertiesMap reports all the CloudFormation properties for AWS::IoT::TopicRule.AssetPropertyTimestamp.
	AWS_IoT_TopicRule_AssetPropertyTimestamp__PropertiesMap = struct {
		OffsetInNanos string
		TimeInSeconds string
	}{
		OffsetInNanos: "OffsetInNanos",
		TimeInSeconds: "TimeInSeconds",
	}

	// AWS_IoT_TopicRule_AssetPropertyTimestamp__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::TopicRule.AssetPropertyTimestamp.
	AWS_IoT_TopicRule_AssetPropertyTimestamp__PropertiesSlice = []string{
		AWS_IoT_TopicRule_AssetPropertyTimestamp__PropertiesMap.OffsetInNanos,
		AWS_IoT_TopicRule_AssetPropertyTimestamp__PropertiesMap.TimeInSeconds,
	}
)

// AWS_IoT_TopicRule_AssetPropertyTimestamp is a binding for AWS::IoT::TopicRule.AssetPropertyTimestamp.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertytimestamp.html
type AWS_IoT_TopicRule_AssetPropertyTimestamp struct {
	// OffsetInNanos is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertytimestamp.html#cfn-iot-topicrule-assetpropertytimestamp-offsetinnanos
	OffsetInNanos cfz.Expression[string] `json:"OffsetInNanos,omitempty"`

	// TimeInSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertytimestamp.html#cfn-iot-topicrule-assetpropertytimestamp-timeinseconds
	TimeInSeconds cfz.Expression[string] `json:"TimeInSeconds,omitempty"`
}

// New__AWS_IoT_TopicRule_AssetPropertyTimestamp initializes a new AWS_IoT_TopicRule_AssetPropertyTimestamp.
func New__AWS_IoT_TopicRule_AssetPropertyTimestamp() AWS_IoT_TopicRule_AssetPropertyTimestamp {
	return AWS_IoT_TopicRule_AssetPropertyTimestamp{}
}

// GetType returns the CloudFormation type.
func (AWS_IoT_TopicRule_AssetPropertyTimestamp) GetType() string {
	return AWS_IoT_TopicRule_AssetPropertyTimestamp__Type
}

// Set__OffsetInNanos updates property "OffsetInNanos".
func (t AWS_IoT_TopicRule_AssetPropertyTimestamp) Set__OffsetInNanos(v cfz.Expression[string]) AWS_IoT_TopicRule_AssetPropertyTimestamp {
	t.OffsetInNanos = v
	return t
}

// SetV__OffsetInNanos updates property "OffsetInNanos".
func (t AWS_IoT_TopicRule_AssetPropertyTimestamp) SetV__OffsetInNanos(v string) AWS_IoT_TopicRule_AssetPropertyTimestamp {
	t.OffsetInNanos = cfz.V(v)
	return t
}

// Set__TimeInSeconds updates property "TimeInSeconds".
func (t AWS_IoT_TopicRule_AssetPropertyTimestamp) Set__TimeInSeconds(v cfz.Expression[string]) AWS_IoT_TopicRule_AssetPropertyTimestamp {
	t.TimeInSeconds = v
	return t
}

// SetV__TimeInSeconds updates property "TimeInSeconds".
func (t AWS_IoT_TopicRule_AssetPropertyTimestamp) SetV__TimeInSeconds(v string) AWS_IoT_TopicRule_AssetPropertyTimestamp {
	t.TimeInSeconds = cfz.V(v)
	return t
}
