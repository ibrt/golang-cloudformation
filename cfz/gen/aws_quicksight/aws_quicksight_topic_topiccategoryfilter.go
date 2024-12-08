// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_Topic_TopicCategoryFilter__Type is the CloudFormation type for AWS::QuickSight::Topic.TopicCategoryFilter.
	AWS_QuickSight_Topic_TopicCategoryFilter__Type = "AWS::QuickSight::Topic.TopicCategoryFilter"
)

var (
	// AWS_QuickSight_Topic_TopicCategoryFilter__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::Topic.TopicCategoryFilter.
	AWS_QuickSight_Topic_TopicCategoryFilter__PropertiesMap = struct {
		CategoryFilterFunction string
		CategoryFilterType     string
		Constant               string
		Inverse                string
	}{
		CategoryFilterFunction: "CategoryFilterFunction",
		CategoryFilterType:     "CategoryFilterType",
		Constant:               "Constant",
		Inverse:                "Inverse",
	}

	// AWS_QuickSight_Topic_TopicCategoryFilter__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::Topic.TopicCategoryFilter.
	AWS_QuickSight_Topic_TopicCategoryFilter__PropertiesSlice = []string{
		AWS_QuickSight_Topic_TopicCategoryFilter__PropertiesMap.CategoryFilterFunction,
		AWS_QuickSight_Topic_TopicCategoryFilter__PropertiesMap.CategoryFilterType,
		AWS_QuickSight_Topic_TopicCategoryFilter__PropertiesMap.Constant,
		AWS_QuickSight_Topic_TopicCategoryFilter__PropertiesMap.Inverse,
	}
)

// AWS_QuickSight_Topic_TopicCategoryFilter is a binding for AWS::QuickSight::Topic.TopicCategoryFilter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html
type AWS_QuickSight_Topic_TopicCategoryFilter struct {
	// CategoryFilterFunction is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html#cfn-quicksight-topic-topiccategoryfilter-categoryfilterfunction
	CategoryFilterFunction cfz.Expression[string] `json:"CategoryFilterFunction,omitempty"`

	// CategoryFilterType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html#cfn-quicksight-topic-topiccategoryfilter-categoryfiltertype
	CategoryFilterType cfz.Expression[string] `json:"CategoryFilterType,omitempty"`

	// Constant is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html#cfn-quicksight-topic-topiccategoryfilter-constant
	Constant cfz.Expression[AWS_QuickSight_Topic_TopicCategoryFilterConstant] `json:"Constant,omitempty"`

	// Inverse is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html#cfn-quicksight-topic-topiccategoryfilter-inverse
	Inverse cfz.Expression[bool] `json:"Inverse,omitempty"`
}

// New__AWS_QuickSight_Topic_TopicCategoryFilter initializes a new AWS_QuickSight_Topic_TopicCategoryFilter.
func New__AWS_QuickSight_Topic_TopicCategoryFilter() AWS_QuickSight_Topic_TopicCategoryFilter {
	return AWS_QuickSight_Topic_TopicCategoryFilter{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_Topic_TopicCategoryFilter) GetType() string {
	return AWS_QuickSight_Topic_TopicCategoryFilter__Type
}

// Set__CategoryFilterFunction updates property "CategoryFilterFunction".
func (t AWS_QuickSight_Topic_TopicCategoryFilter) Set__CategoryFilterFunction(v cfz.Expression[string]) AWS_QuickSight_Topic_TopicCategoryFilter {
	t.CategoryFilterFunction = v
	return t
}

// SetV__CategoryFilterFunction updates property "CategoryFilterFunction".
func (t AWS_QuickSight_Topic_TopicCategoryFilter) SetV__CategoryFilterFunction(v string) AWS_QuickSight_Topic_TopicCategoryFilter {
	t.CategoryFilterFunction = cfz.V(v)
	return t
}

// Set__CategoryFilterType updates property "CategoryFilterType".
func (t AWS_QuickSight_Topic_TopicCategoryFilter) Set__CategoryFilterType(v cfz.Expression[string]) AWS_QuickSight_Topic_TopicCategoryFilter {
	t.CategoryFilterType = v
	return t
}

// SetV__CategoryFilterType updates property "CategoryFilterType".
func (t AWS_QuickSight_Topic_TopicCategoryFilter) SetV__CategoryFilterType(v string) AWS_QuickSight_Topic_TopicCategoryFilter {
	t.CategoryFilterType = cfz.V(v)
	return t
}

// Set__Constant updates property "Constant".
func (t AWS_QuickSight_Topic_TopicCategoryFilter) Set__Constant(v cfz.Expression[AWS_QuickSight_Topic_TopicCategoryFilterConstant]) AWS_QuickSight_Topic_TopicCategoryFilter {
	t.Constant = v
	return t
}

// SetV__Constant updates property "Constant".
func (t AWS_QuickSight_Topic_TopicCategoryFilter) SetV__Constant(v AWS_QuickSight_Topic_TopicCategoryFilterConstant) AWS_QuickSight_Topic_TopicCategoryFilter {
	t.Constant = cfz.V(v)
	return t
}

// Set__Inverse updates property "Inverse".
func (t AWS_QuickSight_Topic_TopicCategoryFilter) Set__Inverse(v cfz.Expression[bool]) AWS_QuickSight_Topic_TopicCategoryFilter {
	t.Inverse = v
	return t
}

// SetV__Inverse updates property "Inverse".
func (t AWS_QuickSight_Topic_TopicCategoryFilter) SetV__Inverse(v bool) AWS_QuickSight_Topic_TopicCategoryFilter {
	t.Inverse = cfz.V(v)
	return t
}
