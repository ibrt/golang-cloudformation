// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_Topic_TopicCategoryFilterConstant__Type is the CloudFormation type for AWS::QuickSight::Topic.TopicCategoryFilterConstant.
	AWS_QuickSight_Topic_TopicCategoryFilterConstant__Type = "AWS::QuickSight::Topic.TopicCategoryFilterConstant"
)

var (
	// AWS_QuickSight_Topic_TopicCategoryFilterConstant__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::Topic.TopicCategoryFilterConstant.
	AWS_QuickSight_Topic_TopicCategoryFilterConstant__PropertiesMap = struct {
		CollectiveConstant string
		ConstantType       string
		SingularConstant   string
	}{
		CollectiveConstant: "CollectiveConstant",
		ConstantType:       "ConstantType",
		SingularConstant:   "SingularConstant",
	}

	// AWS_QuickSight_Topic_TopicCategoryFilterConstant__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::Topic.TopicCategoryFilterConstant.
	AWS_QuickSight_Topic_TopicCategoryFilterConstant__PropertiesSlice = []string{
		AWS_QuickSight_Topic_TopicCategoryFilterConstant__PropertiesMap.CollectiveConstant,
		AWS_QuickSight_Topic_TopicCategoryFilterConstant__PropertiesMap.ConstantType,
		AWS_QuickSight_Topic_TopicCategoryFilterConstant__PropertiesMap.SingularConstant,
	}
)

// AWS_QuickSight_Topic_TopicCategoryFilterConstant is a binding for AWS::QuickSight::Topic.TopicCategoryFilterConstant.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html
type AWS_QuickSight_Topic_TopicCategoryFilterConstant struct {
	// CollectiveConstant is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html#cfn-quicksight-topic-topiccategoryfilterconstant-collectiveconstant
	CollectiveConstant cfz.Expression[AWS_QuickSight_Topic_CollectiveConstant] `json:"CollectiveConstant,omitempty"`

	// ConstantType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html#cfn-quicksight-topic-topiccategoryfilterconstant-constanttype
	ConstantType cfz.Expression[string] `json:"ConstantType,omitempty"`

	// SingularConstant is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilterconstant.html#cfn-quicksight-topic-topiccategoryfilterconstant-singularconstant
	SingularConstant cfz.Expression[string] `json:"SingularConstant,omitempty"`
}

// New__AWS_QuickSight_Topic_TopicCategoryFilterConstant initializes a new AWS_QuickSight_Topic_TopicCategoryFilterConstant.
func New__AWS_QuickSight_Topic_TopicCategoryFilterConstant() AWS_QuickSight_Topic_TopicCategoryFilterConstant {
	return AWS_QuickSight_Topic_TopicCategoryFilterConstant{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_Topic_TopicCategoryFilterConstant) GetType() string {
	return AWS_QuickSight_Topic_TopicCategoryFilterConstant__Type
}

// Set__CollectiveConstant updates property "CollectiveConstant".
func (t AWS_QuickSight_Topic_TopicCategoryFilterConstant) Set__CollectiveConstant(v cfz.Expression[AWS_QuickSight_Topic_CollectiveConstant]) AWS_QuickSight_Topic_TopicCategoryFilterConstant {
	t.CollectiveConstant = v
	return t
}

// SetV__CollectiveConstant updates property "CollectiveConstant".
func (t AWS_QuickSight_Topic_TopicCategoryFilterConstant) SetV__CollectiveConstant(v AWS_QuickSight_Topic_CollectiveConstant) AWS_QuickSight_Topic_TopicCategoryFilterConstant {
	t.CollectiveConstant = cfz.V(v)
	return t
}

// Set__ConstantType updates property "ConstantType".
func (t AWS_QuickSight_Topic_TopicCategoryFilterConstant) Set__ConstantType(v cfz.Expression[string]) AWS_QuickSight_Topic_TopicCategoryFilterConstant {
	t.ConstantType = v
	return t
}

// SetV__ConstantType updates property "ConstantType".
func (t AWS_QuickSight_Topic_TopicCategoryFilterConstant) SetV__ConstantType(v string) AWS_QuickSight_Topic_TopicCategoryFilterConstant {
	t.ConstantType = cfz.V(v)
	return t
}

// Set__SingularConstant updates property "SingularConstant".
func (t AWS_QuickSight_Topic_TopicCategoryFilterConstant) Set__SingularConstant(v cfz.Expression[string]) AWS_QuickSight_Topic_TopicCategoryFilterConstant {
	t.SingularConstant = v
	return t
}

// SetV__SingularConstant updates property "SingularConstant".
func (t AWS_QuickSight_Topic_TopicCategoryFilterConstant) SetV__SingularConstant(v string) AWS_QuickSight_Topic_TopicCategoryFilterConstant {
	t.SingularConstant = cfz.V(v)
	return t
}
