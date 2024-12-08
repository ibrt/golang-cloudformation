// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisanalyticsv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput__Type is the CloudFormation type for AWS::KinesisAnalyticsV2::ApplicationOutput.KinesisFirehoseOutput.
	AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput__Type = "AWS::KinesisAnalyticsV2::ApplicationOutput.KinesisFirehoseOutput"
)

var (
	// AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput__PropertiesMap reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::ApplicationOutput.KinesisFirehoseOutput.
	AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput__PropertiesMap = struct {
		ResourceARN string
	}{
		ResourceARN: "ResourceARN",
	}

	// AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::ApplicationOutput.KinesisFirehoseOutput.
	AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput__PropertiesSlice = []string{
		AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput__PropertiesMap.ResourceARN,
	}
)

// AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput is a binding for AWS::KinesisAnalyticsV2::ApplicationOutput.KinesisFirehoseOutput.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-kinesisfirehoseoutput.html
type AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput struct {
	// ResourceARN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-kinesisfirehoseoutput.html#cfn-kinesisanalyticsv2-applicationoutput-kinesisfirehoseoutput-resourcearn
	ResourceARN cfz.Expression[string] `json:"ResourceARN,omitempty"`
}

// New__AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput initializes a new AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput.
func New__AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput() AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput {
	return AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput) GetType() string {
	return AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput__Type
}

// Set__ResourceARN updates property "ResourceARN".
func (t AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput) Set__ResourceARN(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput {
	t.ResourceARN = v
	return t
}

// SetV__ResourceARN updates property "ResourceARN".
func (t AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput) SetV__ResourceARN(v string) AWS_KinesisAnalyticsV2_ApplicationOutput_KinesisFirehoseOutput {
	t.ResourceARN = cfz.V(v)
	return t
}
