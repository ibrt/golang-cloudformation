// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_Endpoint_VariantProperty__Type is the CloudFormation type for AWS::SageMaker::Endpoint.VariantProperty.
	AWS_SageMaker_Endpoint_VariantProperty__Type = "AWS::SageMaker::Endpoint.VariantProperty"
)

var (
	// AWS_SageMaker_Endpoint_VariantProperty__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::Endpoint.VariantProperty.
	AWS_SageMaker_Endpoint_VariantProperty__PropertiesMap = struct {
		VariantPropertyType string
	}{
		VariantPropertyType: "VariantPropertyType",
	}

	// AWS_SageMaker_Endpoint_VariantProperty__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::Endpoint.VariantProperty.
	AWS_SageMaker_Endpoint_VariantProperty__PropertiesSlice = []string{
		AWS_SageMaker_Endpoint_VariantProperty__PropertiesMap.VariantPropertyType,
	}
)

// AWS_SageMaker_Endpoint_VariantProperty is a binding for AWS::SageMaker::Endpoint.VariantProperty.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-variantproperty.html
type AWS_SageMaker_Endpoint_VariantProperty struct {
	// VariantPropertyType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-variantproperty.html#cfn-sagemaker-endpoint-variantproperty-variantpropertytype
	VariantPropertyType cfz.Expression[string] `json:"VariantPropertyType,omitempty"`
}

// New__AWS_SageMaker_Endpoint_VariantProperty initializes a new AWS_SageMaker_Endpoint_VariantProperty.
func New__AWS_SageMaker_Endpoint_VariantProperty() AWS_SageMaker_Endpoint_VariantProperty {
	return AWS_SageMaker_Endpoint_VariantProperty{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_Endpoint_VariantProperty) GetType() string {
	return AWS_SageMaker_Endpoint_VariantProperty__Type
}

// Set__VariantPropertyType updates property "VariantPropertyType".
func (t AWS_SageMaker_Endpoint_VariantProperty) Set__VariantPropertyType(v cfz.Expression[string]) AWS_SageMaker_Endpoint_VariantProperty {
	t.VariantPropertyType = v
	return t
}

// SetV__VariantPropertyType updates property "VariantPropertyType".
func (t AWS_SageMaker_Endpoint_VariantProperty) SetV__VariantPropertyType(v string) AWS_SageMaker_Endpoint_VariantProperty {
	t.VariantPropertyType = cfz.V(v)
	return t
}
