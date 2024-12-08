// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_EndpointConfig_DataCaptureConfig__Type is the CloudFormation type for AWS::SageMaker::EndpointConfig.DataCaptureConfig.
	AWS_SageMaker_EndpointConfig_DataCaptureConfig__Type = "AWS::SageMaker::EndpointConfig.DataCaptureConfig"
)

var (
	// AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::EndpointConfig.DataCaptureConfig.
	AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesMap = struct {
		CaptureContentTypeHeader  string
		CaptureOptions            string
		DestinationS3Uri          string
		EnableCapture             string
		InitialSamplingPercentage string
		KmsKeyId                  string
	}{
		CaptureContentTypeHeader:  "CaptureContentTypeHeader",
		CaptureOptions:            "CaptureOptions",
		DestinationS3Uri:          "DestinationS3Uri",
		EnableCapture:             "EnableCapture",
		InitialSamplingPercentage: "InitialSamplingPercentage",
		KmsKeyId:                  "KmsKeyId",
	}

	// AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::EndpointConfig.DataCaptureConfig.
	AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesSlice = []string{
		AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesMap.CaptureContentTypeHeader,
		AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesMap.CaptureOptions,
		AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesMap.DestinationS3Uri,
		AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesMap.EnableCapture,
		AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesMap.InitialSamplingPercentage,
		AWS_SageMaker_EndpointConfig_DataCaptureConfig__PropertiesMap.KmsKeyId,
	}
)

// AWS_SageMaker_EndpointConfig_DataCaptureConfig is a binding for AWS::SageMaker::EndpointConfig.DataCaptureConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html
type AWS_SageMaker_EndpointConfig_DataCaptureConfig struct {
	// CaptureContentTypeHeader is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-capturecontenttypeheader
	CaptureContentTypeHeader cfz.Expression[AWS_SageMaker_EndpointConfig_CaptureContentTypeHeader] `json:"CaptureContentTypeHeader,omitempty"`

	// CaptureOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-captureoptions
	CaptureOptions cfz.ExpressionSlice[AWS_SageMaker_EndpointConfig_CaptureOption] `json:"CaptureOptions,omitempty"`

	// DestinationS3Uri is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-destinations3uri
	DestinationS3Uri cfz.Expression[string] `json:"DestinationS3Uri,omitempty"`

	// EnableCapture is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-enablecapture
	EnableCapture cfz.Expression[bool] `json:"EnableCapture,omitempty"`

	// InitialSamplingPercentage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-initialsamplingpercentage
	InitialSamplingPercentage cfz.Expression[int32] `json:"InitialSamplingPercentage,omitempty"`

	// KmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-datacaptureconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig-kmskeyid
	KmsKeyId cfz.Expression[string] `json:"KmsKeyId,omitempty"`
}

// New__AWS_SageMaker_EndpointConfig_DataCaptureConfig initializes a new AWS_SageMaker_EndpointConfig_DataCaptureConfig.
func New__AWS_SageMaker_EndpointConfig_DataCaptureConfig() AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	return AWS_SageMaker_EndpointConfig_DataCaptureConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_EndpointConfig_DataCaptureConfig) GetType() string {
	return AWS_SageMaker_EndpointConfig_DataCaptureConfig__Type
}

// Set__CaptureContentTypeHeader updates property "CaptureContentTypeHeader".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) Set__CaptureContentTypeHeader(v cfz.Expression[AWS_SageMaker_EndpointConfig_CaptureContentTypeHeader]) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.CaptureContentTypeHeader = v
	return t
}

// SetV__CaptureContentTypeHeader updates property "CaptureContentTypeHeader".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) SetV__CaptureContentTypeHeader(v AWS_SageMaker_EndpointConfig_CaptureContentTypeHeader) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.CaptureContentTypeHeader = cfz.V(v)
	return t
}

// Set__CaptureOptions updates property "CaptureOptions".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) Set__CaptureOptions(v cfz.ExpressionSlice[AWS_SageMaker_EndpointConfig_CaptureOption]) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.CaptureOptions = v
	return t
}

// SetS__CaptureOptions updates property "CaptureOptions".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) SetS__CaptureOptions(v ...cfz.Expression[AWS_SageMaker_EndpointConfig_CaptureOption]) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.CaptureOptions = cfz.S(v...)
	return t
}

// SetSV__CaptureOptions updates property "CaptureOptions".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) SetSV__CaptureOptions(v ...AWS_SageMaker_EndpointConfig_CaptureOption) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.CaptureOptions = cfz.SV(v...)
	return t
}

// Set__DestinationS3Uri updates property "DestinationS3Uri".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) Set__DestinationS3Uri(v cfz.Expression[string]) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.DestinationS3Uri = v
	return t
}

// SetV__DestinationS3Uri updates property "DestinationS3Uri".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) SetV__DestinationS3Uri(v string) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.DestinationS3Uri = cfz.V(v)
	return t
}

// Set__EnableCapture updates property "EnableCapture".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) Set__EnableCapture(v cfz.Expression[bool]) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.EnableCapture = v
	return t
}

// SetV__EnableCapture updates property "EnableCapture".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) SetV__EnableCapture(v bool) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.EnableCapture = cfz.V(v)
	return t
}

// Set__InitialSamplingPercentage updates property "InitialSamplingPercentage".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) Set__InitialSamplingPercentage(v cfz.Expression[int32]) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.InitialSamplingPercentage = v
	return t
}

// SetV__InitialSamplingPercentage updates property "InitialSamplingPercentage".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) SetV__InitialSamplingPercentage(v int32) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.InitialSamplingPercentage = cfz.V(v)
	return t
}

// Set__KmsKeyId updates property "KmsKeyId".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) Set__KmsKeyId(v cfz.Expression[string]) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.KmsKeyId = v
	return t
}

// SetV__KmsKeyId updates property "KmsKeyId".
func (t AWS_SageMaker_EndpointConfig_DataCaptureConfig) SetV__KmsKeyId(v string) AWS_SageMaker_EndpointConfig_DataCaptureConfig {
	t.KmsKeyId = cfz.V(v)
	return t
}
