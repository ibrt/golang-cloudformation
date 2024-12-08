// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_Theme_DataColorPalette__Type is the CloudFormation type for AWS::QuickSight::Theme.DataColorPalette.
	AWS_QuickSight_Theme_DataColorPalette__Type = "AWS::QuickSight::Theme.DataColorPalette"
)

var (
	// AWS_QuickSight_Theme_DataColorPalette__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::Theme.DataColorPalette.
	AWS_QuickSight_Theme_DataColorPalette__PropertiesMap = struct {
		Colors         string
		EmptyFillColor string
		MinMaxGradient string
	}{
		Colors:         "Colors",
		EmptyFillColor: "EmptyFillColor",
		MinMaxGradient: "MinMaxGradient",
	}

	// AWS_QuickSight_Theme_DataColorPalette__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::Theme.DataColorPalette.
	AWS_QuickSight_Theme_DataColorPalette__PropertiesSlice = []string{
		AWS_QuickSight_Theme_DataColorPalette__PropertiesMap.Colors,
		AWS_QuickSight_Theme_DataColorPalette__PropertiesMap.EmptyFillColor,
		AWS_QuickSight_Theme_DataColorPalette__PropertiesMap.MinMaxGradient,
	}
)

// AWS_QuickSight_Theme_DataColorPalette is a binding for AWS::QuickSight::Theme.DataColorPalette.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html
type AWS_QuickSight_Theme_DataColorPalette struct {
	// Colors is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html#cfn-quicksight-theme-datacolorpalette-colors
	Colors cfz.ExpressionSlice[string] `json:"Colors,omitempty"`

	// EmptyFillColor is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html#cfn-quicksight-theme-datacolorpalette-emptyfillcolor
	EmptyFillColor cfz.Expression[string] `json:"EmptyFillColor,omitempty"`

	// MinMaxGradient is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-datacolorpalette.html#cfn-quicksight-theme-datacolorpalette-minmaxgradient
	MinMaxGradient cfz.ExpressionSlice[string] `json:"MinMaxGradient,omitempty"`
}

// New__AWS_QuickSight_Theme_DataColorPalette initializes a new AWS_QuickSight_Theme_DataColorPalette.
func New__AWS_QuickSight_Theme_DataColorPalette() AWS_QuickSight_Theme_DataColorPalette {
	return AWS_QuickSight_Theme_DataColorPalette{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_Theme_DataColorPalette) GetType() string {
	return AWS_QuickSight_Theme_DataColorPalette__Type
}

// Set__Colors updates property "Colors".
func (t AWS_QuickSight_Theme_DataColorPalette) Set__Colors(v cfz.ExpressionSlice[string]) AWS_QuickSight_Theme_DataColorPalette {
	t.Colors = v
	return t
}

// SetS__Colors updates property "Colors".
func (t AWS_QuickSight_Theme_DataColorPalette) SetS__Colors(v ...cfz.Expression[string]) AWS_QuickSight_Theme_DataColorPalette {
	t.Colors = cfz.S(v...)
	return t
}

// SetSV__Colors updates property "Colors".
func (t AWS_QuickSight_Theme_DataColorPalette) SetSV__Colors(v ...string) AWS_QuickSight_Theme_DataColorPalette {
	t.Colors = cfz.SV(v...)
	return t
}

// Set__EmptyFillColor updates property "EmptyFillColor".
func (t AWS_QuickSight_Theme_DataColorPalette) Set__EmptyFillColor(v cfz.Expression[string]) AWS_QuickSight_Theme_DataColorPalette {
	t.EmptyFillColor = v
	return t
}

// SetV__EmptyFillColor updates property "EmptyFillColor".
func (t AWS_QuickSight_Theme_DataColorPalette) SetV__EmptyFillColor(v string) AWS_QuickSight_Theme_DataColorPalette {
	t.EmptyFillColor = cfz.V(v)
	return t
}

// Set__MinMaxGradient updates property "MinMaxGradient".
func (t AWS_QuickSight_Theme_DataColorPalette) Set__MinMaxGradient(v cfz.ExpressionSlice[string]) AWS_QuickSight_Theme_DataColorPalette {
	t.MinMaxGradient = v
	return t
}

// SetS__MinMaxGradient updates property "MinMaxGradient".
func (t AWS_QuickSight_Theme_DataColorPalette) SetS__MinMaxGradient(v ...cfz.Expression[string]) AWS_QuickSight_Theme_DataColorPalette {
	t.MinMaxGradient = cfz.S(v...)
	return t
}

// SetSV__MinMaxGradient updates property "MinMaxGradient".
func (t AWS_QuickSight_Theme_DataColorPalette) SetSV__MinMaxGradient(v ...string) AWS_QuickSight_Theme_DataColorPalette {
	t.MinMaxGradient = cfz.SV(v...)
	return t
}
