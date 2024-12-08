// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_amazonmq

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AmazonMQ_Configuration_TagsEntry__Type is the CloudFormation type for AWS::AmazonMQ::Configuration.TagsEntry.
	AWS_AmazonMQ_Configuration_TagsEntry__Type = "AWS::AmazonMQ::Configuration.TagsEntry"
)

var (
	// AWS_AmazonMQ_Configuration_TagsEntry__PropertiesMap reports all the CloudFormation properties for AWS::AmazonMQ::Configuration.TagsEntry.
	AWS_AmazonMQ_Configuration_TagsEntry__PropertiesMap = struct {
		Key   string
		Value string
	}{
		Key:   "Key",
		Value: "Value",
	}

	// AWS_AmazonMQ_Configuration_TagsEntry__PropertiesSlice reports all the CloudFormation properties for AWS::AmazonMQ::Configuration.TagsEntry.
	AWS_AmazonMQ_Configuration_TagsEntry__PropertiesSlice = []string{
		AWS_AmazonMQ_Configuration_TagsEntry__PropertiesMap.Key,
		AWS_AmazonMQ_Configuration_TagsEntry__PropertiesMap.Value,
	}
)

// AWS_AmazonMQ_Configuration_TagsEntry is a binding for AWS::AmazonMQ::Configuration.TagsEntry.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configuration-tagsentry.html
type AWS_AmazonMQ_Configuration_TagsEntry struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configuration-tagsentry.html#cfn-amazonmq-configuration-tagsentry-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configuration-tagsentry.html#cfn-amazonmq-configuration-tagsentry-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_AmazonMQ_Configuration_TagsEntry initializes a new AWS_AmazonMQ_Configuration_TagsEntry.
func New__AWS_AmazonMQ_Configuration_TagsEntry() AWS_AmazonMQ_Configuration_TagsEntry {
	return AWS_AmazonMQ_Configuration_TagsEntry{}
}

// GetType returns the CloudFormation type.
func (AWS_AmazonMQ_Configuration_TagsEntry) GetType() string {
	return AWS_AmazonMQ_Configuration_TagsEntry__Type
}

// Set__Key updates property "Key".
func (t AWS_AmazonMQ_Configuration_TagsEntry) Set__Key(v cfz.Expression[string]) AWS_AmazonMQ_Configuration_TagsEntry {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_AmazonMQ_Configuration_TagsEntry) SetV__Key(v string) AWS_AmazonMQ_Configuration_TagsEntry {
	t.Key = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_AmazonMQ_Configuration_TagsEntry) Set__Value(v cfz.Expression[string]) AWS_AmazonMQ_Configuration_TagsEntry {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_AmazonMQ_Configuration_TagsEntry) SetV__Value(v string) AWS_AmazonMQ_Configuration_TagsEntry {
	t.Value = cfz.V(v)
	return t
}
