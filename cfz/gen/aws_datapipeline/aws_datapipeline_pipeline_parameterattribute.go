// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datapipeline

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataPipeline_Pipeline_ParameterAttribute__Type is the CloudFormation type for AWS::DataPipeline::Pipeline.ParameterAttribute.
	AWS_DataPipeline_Pipeline_ParameterAttribute__Type = "AWS::DataPipeline::Pipeline.ParameterAttribute"
)

var (
	// AWS_DataPipeline_Pipeline_ParameterAttribute__PropertiesMap reports all the CloudFormation properties for AWS::DataPipeline::Pipeline.ParameterAttribute.
	AWS_DataPipeline_Pipeline_ParameterAttribute__PropertiesMap = struct {
		Key         string
		StringValue string
	}{
		Key:         "Key",
		StringValue: "StringValue",
	}

	// AWS_DataPipeline_Pipeline_ParameterAttribute__PropertiesSlice reports all the CloudFormation properties for AWS::DataPipeline::Pipeline.ParameterAttribute.
	AWS_DataPipeline_Pipeline_ParameterAttribute__PropertiesSlice = []string{
		AWS_DataPipeline_Pipeline_ParameterAttribute__PropertiesMap.Key,
		AWS_DataPipeline_Pipeline_ParameterAttribute__PropertiesMap.StringValue,
	}
)

// AWS_DataPipeline_Pipeline_ParameterAttribute is a binding for AWS::DataPipeline::Pipeline.ParameterAttribute.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterattribute.html
type AWS_DataPipeline_Pipeline_ParameterAttribute struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterattribute.html#cfn-datapipeline-pipeline-parameterattribute-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// StringValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterattribute.html#cfn-datapipeline-pipeline-parameterattribute-stringvalue
	StringValue cfz.Expression[string] `json:"StringValue,omitempty"`
}

// New__AWS_DataPipeline_Pipeline_ParameterAttribute initializes a new AWS_DataPipeline_Pipeline_ParameterAttribute.
func New__AWS_DataPipeline_Pipeline_ParameterAttribute() AWS_DataPipeline_Pipeline_ParameterAttribute {
	return AWS_DataPipeline_Pipeline_ParameterAttribute{}
}

// GetType returns the CloudFormation type.
func (AWS_DataPipeline_Pipeline_ParameterAttribute) GetType() string {
	return AWS_DataPipeline_Pipeline_ParameterAttribute__Type
}

// Set__Key updates property "Key".
func (t AWS_DataPipeline_Pipeline_ParameterAttribute) Set__Key(v cfz.Expression[string]) AWS_DataPipeline_Pipeline_ParameterAttribute {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_DataPipeline_Pipeline_ParameterAttribute) SetV__Key(v string) AWS_DataPipeline_Pipeline_ParameterAttribute {
	t.Key = cfz.V(v)
	return t
}

// Set__StringValue updates property "StringValue".
func (t AWS_DataPipeline_Pipeline_ParameterAttribute) Set__StringValue(v cfz.Expression[string]) AWS_DataPipeline_Pipeline_ParameterAttribute {
	t.StringValue = v
	return t
}

// SetV__StringValue updates property "StringValue".
func (t AWS_DataPipeline_Pipeline_ParameterAttribute) SetV__StringValue(v string) AWS_DataPipeline_Pipeline_ParameterAttribute {
	t.StringValue = cfz.V(v)
	return t
}
