// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_imagebuilder

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ImageBuilder_ImageRecipe_ComponentParameter__Type is the CloudFormation type for AWS::ImageBuilder::ImageRecipe.ComponentParameter.
	AWS_ImageBuilder_ImageRecipe_ComponentParameter__Type = "AWS::ImageBuilder::ImageRecipe.ComponentParameter"
)

var (
	// AWS_ImageBuilder_ImageRecipe_ComponentParameter__PropertiesMap reports all the CloudFormation properties for AWS::ImageBuilder::ImageRecipe.ComponentParameter.
	AWS_ImageBuilder_ImageRecipe_ComponentParameter__PropertiesMap = struct {
		Name  string
		Value string
	}{
		Name:  "Name",
		Value: "Value",
	}

	// AWS_ImageBuilder_ImageRecipe_ComponentParameter__PropertiesSlice reports all the CloudFormation properties for AWS::ImageBuilder::ImageRecipe.ComponentParameter.
	AWS_ImageBuilder_ImageRecipe_ComponentParameter__PropertiesSlice = []string{
		AWS_ImageBuilder_ImageRecipe_ComponentParameter__PropertiesMap.Name,
		AWS_ImageBuilder_ImageRecipe_ComponentParameter__PropertiesMap.Value,
	}
)

// AWS_ImageBuilder_ImageRecipe_ComponentParameter is a binding for AWS::ImageBuilder::ImageRecipe.ComponentParameter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentparameter.html
type AWS_ImageBuilder_ImageRecipe_ComponentParameter struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentparameter.html#cfn-imagebuilder-imagerecipe-componentparameter-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentparameter.html#cfn-imagebuilder-imagerecipe-componentparameter-value
	Value cfz.ExpressionSlice[string] `json:"Value,omitempty"`
}

// New__AWS_ImageBuilder_ImageRecipe_ComponentParameter initializes a new AWS_ImageBuilder_ImageRecipe_ComponentParameter.
func New__AWS_ImageBuilder_ImageRecipe_ComponentParameter() AWS_ImageBuilder_ImageRecipe_ComponentParameter {
	return AWS_ImageBuilder_ImageRecipe_ComponentParameter{}
}

// GetType returns the CloudFormation type.
func (AWS_ImageBuilder_ImageRecipe_ComponentParameter) GetType() string {
	return AWS_ImageBuilder_ImageRecipe_ComponentParameter__Type
}

// Set__Name updates property "Name".
func (t AWS_ImageBuilder_ImageRecipe_ComponentParameter) Set__Name(v cfz.Expression[string]) AWS_ImageBuilder_ImageRecipe_ComponentParameter {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_ImageBuilder_ImageRecipe_ComponentParameter) SetV__Name(v string) AWS_ImageBuilder_ImageRecipe_ComponentParameter {
	t.Name = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_ImageBuilder_ImageRecipe_ComponentParameter) Set__Value(v cfz.ExpressionSlice[string]) AWS_ImageBuilder_ImageRecipe_ComponentParameter {
	t.Value = v
	return t
}

// SetS__Value updates property "Value".
func (t AWS_ImageBuilder_ImageRecipe_ComponentParameter) SetS__Value(v ...cfz.Expression[string]) AWS_ImageBuilder_ImageRecipe_ComponentParameter {
	t.Value = cfz.S(v...)
	return t
}

// SetSV__Value updates property "Value".
func (t AWS_ImageBuilder_ImageRecipe_ComponentParameter) SetSV__Value(v ...string) AWS_ImageBuilder_ImageRecipe_ComponentParameter {
	t.Value = cfz.SV(v...)
	return t
}
