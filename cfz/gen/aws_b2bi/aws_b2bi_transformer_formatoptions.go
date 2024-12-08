// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_b2bi

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_B2BI_Transformer_FormatOptions__Type is the CloudFormation type for AWS::B2BI::Transformer.FormatOptions.
	AWS_B2BI_Transformer_FormatOptions__Type = "AWS::B2BI::Transformer.FormatOptions"
)

var (
	// AWS_B2BI_Transformer_FormatOptions__PropertiesMap reports all the CloudFormation properties for AWS::B2BI::Transformer.FormatOptions.
	AWS_B2BI_Transformer_FormatOptions__PropertiesMap = struct {
		X12 string
	}{
		X12: "X12",
	}

	// AWS_B2BI_Transformer_FormatOptions__PropertiesSlice reports all the CloudFormation properties for AWS::B2BI::Transformer.FormatOptions.
	AWS_B2BI_Transformer_FormatOptions__PropertiesSlice = []string{
		AWS_B2BI_Transformer_FormatOptions__PropertiesMap.X12,
	}
)

// AWS_B2BI_Transformer_FormatOptions is a binding for AWS::B2BI::Transformer.FormatOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-formatoptions.html
type AWS_B2BI_Transformer_FormatOptions struct {
	// X12 is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-formatoptions.html#cfn-b2bi-transformer-formatoptions-x12
	X12 cfz.Expression[AWS_B2BI_Transformer_X12Details] `json:"X12,omitempty"`
}

// New__AWS_B2BI_Transformer_FormatOptions initializes a new AWS_B2BI_Transformer_FormatOptions.
func New__AWS_B2BI_Transformer_FormatOptions() AWS_B2BI_Transformer_FormatOptions {
	return AWS_B2BI_Transformer_FormatOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_B2BI_Transformer_FormatOptions) GetType() string {
	return AWS_B2BI_Transformer_FormatOptions__Type
}

// Set__X12 updates property "X12".
func (t AWS_B2BI_Transformer_FormatOptions) Set__X12(v cfz.Expression[AWS_B2BI_Transformer_X12Details]) AWS_B2BI_Transformer_FormatOptions {
	t.X12 = v
	return t
}

// SetV__X12 updates property "X12".
func (t AWS_B2BI_Transformer_FormatOptions) SetV__X12(v AWS_B2BI_Transformer_X12Details) AWS_B2BI_Transformer_FormatOptions {
	t.X12 = cfz.V(v)
	return t
}
