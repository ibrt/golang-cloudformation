// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelPackage_InferenceSpecification__Type is the CloudFormation type for AWS::SageMaker::ModelPackage.InferenceSpecification.
	AWS_SageMaker_ModelPackage_InferenceSpecification__Type = "AWS::SageMaker::ModelPackage.InferenceSpecification"
)

var (
	// AWS_SageMaker_ModelPackage_InferenceSpecification__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelPackage.InferenceSpecification.
	AWS_SageMaker_ModelPackage_InferenceSpecification__PropertiesMap = struct {
		Containers                              string
		SupportedContentTypes                   string
		SupportedRealtimeInferenceInstanceTypes string
		SupportedResponseMIMETypes              string
		SupportedTransformInstanceTypes         string
	}{
		Containers:                              "Containers",
		SupportedContentTypes:                   "SupportedContentTypes",
		SupportedRealtimeInferenceInstanceTypes: "SupportedRealtimeInferenceInstanceTypes",
		SupportedResponseMIMETypes:              "SupportedResponseMIMETypes",
		SupportedTransformInstanceTypes:         "SupportedTransformInstanceTypes",
	}

	// AWS_SageMaker_ModelPackage_InferenceSpecification__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelPackage.InferenceSpecification.
	AWS_SageMaker_ModelPackage_InferenceSpecification__PropertiesSlice = []string{
		AWS_SageMaker_ModelPackage_InferenceSpecification__PropertiesMap.Containers,
		AWS_SageMaker_ModelPackage_InferenceSpecification__PropertiesMap.SupportedContentTypes,
		AWS_SageMaker_ModelPackage_InferenceSpecification__PropertiesMap.SupportedRealtimeInferenceInstanceTypes,
		AWS_SageMaker_ModelPackage_InferenceSpecification__PropertiesMap.SupportedResponseMIMETypes,
		AWS_SageMaker_ModelPackage_InferenceSpecification__PropertiesMap.SupportedTransformInstanceTypes,
	}
)

// AWS_SageMaker_ModelPackage_InferenceSpecification is a binding for AWS::SageMaker::ModelPackage.InferenceSpecification.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html
type AWS_SageMaker_ModelPackage_InferenceSpecification struct {
	// Containers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-containers
	Containers cfz.ExpressionSlice[AWS_SageMaker_ModelPackage_ModelPackageContainerDefinition] `json:"Containers,omitempty"`

	// SupportedContentTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-supportedcontenttypes
	SupportedContentTypes cfz.ExpressionSlice[string] `json:"SupportedContentTypes,omitempty"`

	// SupportedRealtimeInferenceInstanceTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-supportedrealtimeinferenceinstancetypes
	SupportedRealtimeInferenceInstanceTypes cfz.ExpressionSlice[string] `json:"SupportedRealtimeInferenceInstanceTypes,omitempty"`

	// SupportedResponseMIMETypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-supportedresponsemimetypes
	SupportedResponseMIMETypes cfz.ExpressionSlice[string] `json:"SupportedResponseMIMETypes,omitempty"`

	// SupportedTransformInstanceTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-supportedtransforminstancetypes
	SupportedTransformInstanceTypes cfz.ExpressionSlice[string] `json:"SupportedTransformInstanceTypes,omitempty"`
}

// New__AWS_SageMaker_ModelPackage_InferenceSpecification initializes a new AWS_SageMaker_ModelPackage_InferenceSpecification.
func New__AWS_SageMaker_ModelPackage_InferenceSpecification() AWS_SageMaker_ModelPackage_InferenceSpecification {
	return AWS_SageMaker_ModelPackage_InferenceSpecification{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelPackage_InferenceSpecification) GetType() string {
	return AWS_SageMaker_ModelPackage_InferenceSpecification__Type
}

// Set__Containers updates property "Containers".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) Set__Containers(v cfz.ExpressionSlice[AWS_SageMaker_ModelPackage_ModelPackageContainerDefinition]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.Containers = v
	return t
}

// SetS__Containers updates property "Containers".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetS__Containers(v ...cfz.Expression[AWS_SageMaker_ModelPackage_ModelPackageContainerDefinition]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.Containers = cfz.S(v...)
	return t
}

// SetSV__Containers updates property "Containers".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetSV__Containers(v ...AWS_SageMaker_ModelPackage_ModelPackageContainerDefinition) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.Containers = cfz.SV(v...)
	return t
}

// Set__SupportedContentTypes updates property "SupportedContentTypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) Set__SupportedContentTypes(v cfz.ExpressionSlice[string]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedContentTypes = v
	return t
}

// SetS__SupportedContentTypes updates property "SupportedContentTypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetS__SupportedContentTypes(v ...cfz.Expression[string]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedContentTypes = cfz.S(v...)
	return t
}

// SetSV__SupportedContentTypes updates property "SupportedContentTypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetSV__SupportedContentTypes(v ...string) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedContentTypes = cfz.SV(v...)
	return t
}

// Set__SupportedRealtimeInferenceInstanceTypes updates property "SupportedRealtimeInferenceInstanceTypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) Set__SupportedRealtimeInferenceInstanceTypes(v cfz.ExpressionSlice[string]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedRealtimeInferenceInstanceTypes = v
	return t
}

// SetS__SupportedRealtimeInferenceInstanceTypes updates property "SupportedRealtimeInferenceInstanceTypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetS__SupportedRealtimeInferenceInstanceTypes(v ...cfz.Expression[string]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedRealtimeInferenceInstanceTypes = cfz.S(v...)
	return t
}

// SetSV__SupportedRealtimeInferenceInstanceTypes updates property "SupportedRealtimeInferenceInstanceTypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetSV__SupportedRealtimeInferenceInstanceTypes(v ...string) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedRealtimeInferenceInstanceTypes = cfz.SV(v...)
	return t
}

// Set__SupportedResponseMIMETypes updates property "SupportedResponseMIMETypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) Set__SupportedResponseMIMETypes(v cfz.ExpressionSlice[string]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedResponseMIMETypes = v
	return t
}

// SetS__SupportedResponseMIMETypes updates property "SupportedResponseMIMETypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetS__SupportedResponseMIMETypes(v ...cfz.Expression[string]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedResponseMIMETypes = cfz.S(v...)
	return t
}

// SetSV__SupportedResponseMIMETypes updates property "SupportedResponseMIMETypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetSV__SupportedResponseMIMETypes(v ...string) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedResponseMIMETypes = cfz.SV(v...)
	return t
}

// Set__SupportedTransformInstanceTypes updates property "SupportedTransformInstanceTypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) Set__SupportedTransformInstanceTypes(v cfz.ExpressionSlice[string]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedTransformInstanceTypes = v
	return t
}

// SetS__SupportedTransformInstanceTypes updates property "SupportedTransformInstanceTypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetS__SupportedTransformInstanceTypes(v ...cfz.Expression[string]) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedTransformInstanceTypes = cfz.S(v...)
	return t
}

// SetSV__SupportedTransformInstanceTypes updates property "SupportedTransformInstanceTypes".
func (t AWS_SageMaker_ModelPackage_InferenceSpecification) SetSV__SupportedTransformInstanceTypes(v ...string) AWS_SageMaker_ModelPackage_InferenceSpecification {
	t.SupportedTransformInstanceTypes = cfz.SV(v...)
	return t
}
