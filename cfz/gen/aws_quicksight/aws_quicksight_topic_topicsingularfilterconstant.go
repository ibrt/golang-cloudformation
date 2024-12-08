// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_Topic_TopicSingularFilterConstant__Type is the CloudFormation type for AWS::QuickSight::Topic.TopicSingularFilterConstant.
	AWS_QuickSight_Topic_TopicSingularFilterConstant__Type = "AWS::QuickSight::Topic.TopicSingularFilterConstant"
)

var (
	// AWS_QuickSight_Topic_TopicSingularFilterConstant__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::Topic.TopicSingularFilterConstant.
	AWS_QuickSight_Topic_TopicSingularFilterConstant__PropertiesMap = struct {
		ConstantType     string
		SingularConstant string
	}{
		ConstantType:     "ConstantType",
		SingularConstant: "SingularConstant",
	}

	// AWS_QuickSight_Topic_TopicSingularFilterConstant__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::Topic.TopicSingularFilterConstant.
	AWS_QuickSight_Topic_TopicSingularFilterConstant__PropertiesSlice = []string{
		AWS_QuickSight_Topic_TopicSingularFilterConstant__PropertiesMap.ConstantType,
		AWS_QuickSight_Topic_TopicSingularFilterConstant__PropertiesMap.SingularConstant,
	}
)

// AWS_QuickSight_Topic_TopicSingularFilterConstant is a binding for AWS::QuickSight::Topic.TopicSingularFilterConstant.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicsingularfilterconstant.html
type AWS_QuickSight_Topic_TopicSingularFilterConstant struct {
	// ConstantType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicsingularfilterconstant.html#cfn-quicksight-topic-topicsingularfilterconstant-constanttype
	ConstantType cfz.Expression[string] `json:"ConstantType,omitempty"`

	// SingularConstant is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicsingularfilterconstant.html#cfn-quicksight-topic-topicsingularfilterconstant-singularconstant
	SingularConstant cfz.Expression[string] `json:"SingularConstant,omitempty"`
}

// New__AWS_QuickSight_Topic_TopicSingularFilterConstant initializes a new AWS_QuickSight_Topic_TopicSingularFilterConstant.
func New__AWS_QuickSight_Topic_TopicSingularFilterConstant() AWS_QuickSight_Topic_TopicSingularFilterConstant {
	return AWS_QuickSight_Topic_TopicSingularFilterConstant{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_Topic_TopicSingularFilterConstant) GetType() string {
	return AWS_QuickSight_Topic_TopicSingularFilterConstant__Type
}

// Set__ConstantType updates property "ConstantType".
func (t AWS_QuickSight_Topic_TopicSingularFilterConstant) Set__ConstantType(v cfz.Expression[string]) AWS_QuickSight_Topic_TopicSingularFilterConstant {
	t.ConstantType = v
	return t
}

// SetV__ConstantType updates property "ConstantType".
func (t AWS_QuickSight_Topic_TopicSingularFilterConstant) SetV__ConstantType(v string) AWS_QuickSight_Topic_TopicSingularFilterConstant {
	t.ConstantType = cfz.V(v)
	return t
}

// Set__SingularConstant updates property "SingularConstant".
func (t AWS_QuickSight_Topic_TopicSingularFilterConstant) Set__SingularConstant(v cfz.Expression[string]) AWS_QuickSight_Topic_TopicSingularFilterConstant {
	t.SingularConstant = v
	return t
}

// SetV__SingularConstant updates property "SingularConstant".
func (t AWS_QuickSight_Topic_TopicSingularFilterConstant) SetV__SingularConstant(v string) AWS_QuickSight_Topic_TopicSingularFilterConstant {
	t.SingularConstant = cfz.V(v)
	return t
}
