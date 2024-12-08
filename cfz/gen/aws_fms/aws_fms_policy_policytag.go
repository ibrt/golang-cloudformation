// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_fms

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FMS_Policy_PolicyTag__Type is the CloudFormation type for AWS::FMS::Policy.PolicyTag.
	AWS_FMS_Policy_PolicyTag__Type = "AWS::FMS::Policy.PolicyTag"
)

var (
	// AWS_FMS_Policy_PolicyTag__PropertiesMap reports all the CloudFormation properties for AWS::FMS::Policy.PolicyTag.
	AWS_FMS_Policy_PolicyTag__PropertiesMap = struct {
		Key   string
		Value string
	}{
		Key:   "Key",
		Value: "Value",
	}

	// AWS_FMS_Policy_PolicyTag__PropertiesSlice reports all the CloudFormation properties for AWS::FMS::Policy.PolicyTag.
	AWS_FMS_Policy_PolicyTag__PropertiesSlice = []string{
		AWS_FMS_Policy_PolicyTag__PropertiesMap.Key,
		AWS_FMS_Policy_PolicyTag__PropertiesMap.Value,
	}
)

// AWS_FMS_Policy_PolicyTag is a binding for AWS::FMS::Policy.PolicyTag.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policytag.html
type AWS_FMS_Policy_PolicyTag struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policytag.html#cfn-fms-policy-policytag-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policytag.html#cfn-fms-policy-policytag-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_FMS_Policy_PolicyTag initializes a new AWS_FMS_Policy_PolicyTag.
func New__AWS_FMS_Policy_PolicyTag() AWS_FMS_Policy_PolicyTag {
	return AWS_FMS_Policy_PolicyTag{}
}

// GetType returns the CloudFormation type.
func (AWS_FMS_Policy_PolicyTag) GetType() string {
	return AWS_FMS_Policy_PolicyTag__Type
}

// Set__Key updates property "Key".
func (t AWS_FMS_Policy_PolicyTag) Set__Key(v cfz.Expression[string]) AWS_FMS_Policy_PolicyTag {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_FMS_Policy_PolicyTag) SetV__Key(v string) AWS_FMS_Policy_PolicyTag {
	t.Key = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_FMS_Policy_PolicyTag) Set__Value(v cfz.Expression[string]) AWS_FMS_Policy_PolicyTag {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_FMS_Policy_PolicyTag) SetV__Value(v string) AWS_FMS_Policy_PolicyTag {
	t.Value = cfz.V(v)
	return t
}
