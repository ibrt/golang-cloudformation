// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_DataSource_CustomTransformationConfiguration__Type is the CloudFormation type for AWS::Bedrock::DataSource.CustomTransformationConfiguration.
	AWS_Bedrock_DataSource_CustomTransformationConfiguration__Type = "AWS::Bedrock::DataSource.CustomTransformationConfiguration"
)

var (
	// AWS_Bedrock_DataSource_CustomTransformationConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::DataSource.CustomTransformationConfiguration.
	AWS_Bedrock_DataSource_CustomTransformationConfiguration__PropertiesMap = struct {
		IntermediateStorage string
		Transformations     string
	}{
		IntermediateStorage: "IntermediateStorage",
		Transformations:     "Transformations",
	}

	// AWS_Bedrock_DataSource_CustomTransformationConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::DataSource.CustomTransformationConfiguration.
	AWS_Bedrock_DataSource_CustomTransformationConfiguration__PropertiesSlice = []string{
		AWS_Bedrock_DataSource_CustomTransformationConfiguration__PropertiesMap.IntermediateStorage,
		AWS_Bedrock_DataSource_CustomTransformationConfiguration__PropertiesMap.Transformations,
	}
)

// AWS_Bedrock_DataSource_CustomTransformationConfiguration is a binding for AWS::Bedrock::DataSource.CustomTransformationConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-customtransformationconfiguration.html
type AWS_Bedrock_DataSource_CustomTransformationConfiguration struct {
	// IntermediateStorage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-customtransformationconfiguration.html#cfn-bedrock-datasource-customtransformationconfiguration-intermediatestorage
	IntermediateStorage cfz.Expression[AWS_Bedrock_DataSource_IntermediateStorage] `json:"IntermediateStorage,omitempty"`

	// Transformations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-customtransformationconfiguration.html#cfn-bedrock-datasource-customtransformationconfiguration-transformations
	Transformations cfz.ExpressionSlice[AWS_Bedrock_DataSource_Transformation] `json:"Transformations,omitempty"`
}

// New__AWS_Bedrock_DataSource_CustomTransformationConfiguration initializes a new AWS_Bedrock_DataSource_CustomTransformationConfiguration.
func New__AWS_Bedrock_DataSource_CustomTransformationConfiguration() AWS_Bedrock_DataSource_CustomTransformationConfiguration {
	return AWS_Bedrock_DataSource_CustomTransformationConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_DataSource_CustomTransformationConfiguration) GetType() string {
	return AWS_Bedrock_DataSource_CustomTransformationConfiguration__Type
}

// Set__IntermediateStorage updates property "IntermediateStorage".
func (t AWS_Bedrock_DataSource_CustomTransformationConfiguration) Set__IntermediateStorage(v cfz.Expression[AWS_Bedrock_DataSource_IntermediateStorage]) AWS_Bedrock_DataSource_CustomTransformationConfiguration {
	t.IntermediateStorage = v
	return t
}

// SetV__IntermediateStorage updates property "IntermediateStorage".
func (t AWS_Bedrock_DataSource_CustomTransformationConfiguration) SetV__IntermediateStorage(v AWS_Bedrock_DataSource_IntermediateStorage) AWS_Bedrock_DataSource_CustomTransformationConfiguration {
	t.IntermediateStorage = cfz.V(v)
	return t
}

// Set__Transformations updates property "Transformations".
func (t AWS_Bedrock_DataSource_CustomTransformationConfiguration) Set__Transformations(v cfz.ExpressionSlice[AWS_Bedrock_DataSource_Transformation]) AWS_Bedrock_DataSource_CustomTransformationConfiguration {
	t.Transformations = v
	return t
}

// SetS__Transformations updates property "Transformations".
func (t AWS_Bedrock_DataSource_CustomTransformationConfiguration) SetS__Transformations(v ...cfz.Expression[AWS_Bedrock_DataSource_Transformation]) AWS_Bedrock_DataSource_CustomTransformationConfiguration {
	t.Transformations = cfz.S(v...)
	return t
}

// SetSV__Transformations updates property "Transformations".
func (t AWS_Bedrock_DataSource_CustomTransformationConfiguration) SetSV__Transformations(v ...AWS_Bedrock_DataSource_Transformation) AWS_Bedrock_DataSource_CustomTransformationConfiguration {
	t.Transformations = cfz.SV(v...)
	return t
}
