// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_omics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Omics_SequenceStore_SseConfig__Type is the CloudFormation type for AWS::Omics::SequenceStore.SseConfig.
	AWS_Omics_SequenceStore_SseConfig__Type = "AWS::Omics::SequenceStore.SseConfig"
)

var (
	// AWS_Omics_SequenceStore_SseConfig__PropertiesMap reports all the CloudFormation properties for AWS::Omics::SequenceStore.SseConfig.
	AWS_Omics_SequenceStore_SseConfig__PropertiesMap = struct {
		KeyArn string
		Type   string
	}{
		KeyArn: "KeyArn",
		Type:   "Type",
	}

	// AWS_Omics_SequenceStore_SseConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Omics::SequenceStore.SseConfig.
	AWS_Omics_SequenceStore_SseConfig__PropertiesSlice = []string{
		AWS_Omics_SequenceStore_SseConfig__PropertiesMap.KeyArn,
		AWS_Omics_SequenceStore_SseConfig__PropertiesMap.Type,
	}
)

// AWS_Omics_SequenceStore_SseConfig is a binding for AWS::Omics::SequenceStore.SseConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-sequencestore-sseconfig.html
type AWS_Omics_SequenceStore_SseConfig struct {
	// KeyArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-sequencestore-sseconfig.html#cfn-omics-sequencestore-sseconfig-keyarn
	KeyArn cfz.Expression[string] `json:"KeyArn,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-sequencestore-sseconfig.html#cfn-omics-sequencestore-sseconfig-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_Omics_SequenceStore_SseConfig initializes a new AWS_Omics_SequenceStore_SseConfig.
func New__AWS_Omics_SequenceStore_SseConfig() AWS_Omics_SequenceStore_SseConfig {
	return AWS_Omics_SequenceStore_SseConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Omics_SequenceStore_SseConfig) GetType() string {
	return AWS_Omics_SequenceStore_SseConfig__Type
}

// Set__KeyArn updates property "KeyArn".
func (t AWS_Omics_SequenceStore_SseConfig) Set__KeyArn(v cfz.Expression[string]) AWS_Omics_SequenceStore_SseConfig {
	t.KeyArn = v
	return t
}

// SetV__KeyArn updates property "KeyArn".
func (t AWS_Omics_SequenceStore_SseConfig) SetV__KeyArn(v string) AWS_Omics_SequenceStore_SseConfig {
	t.KeyArn = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_Omics_SequenceStore_SseConfig) Set__Type(v cfz.Expression[string]) AWS_Omics_SequenceStore_SseConfig {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_Omics_SequenceStore_SseConfig) SetV__Type(v string) AWS_Omics_SequenceStore_SseConfig {
	t.Type = cfz.V(v)
	return t
}
