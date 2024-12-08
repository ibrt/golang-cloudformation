// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_imagebuilder

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ImageBuilder_Image_ImageScanningConfiguration__Type is the CloudFormation type for AWS::ImageBuilder::Image.ImageScanningConfiguration.
	AWS_ImageBuilder_Image_ImageScanningConfiguration__Type = "AWS::ImageBuilder::Image.ImageScanningConfiguration"
)

var (
	// AWS_ImageBuilder_Image_ImageScanningConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::ImageBuilder::Image.ImageScanningConfiguration.
	AWS_ImageBuilder_Image_ImageScanningConfiguration__PropertiesMap = struct {
		EcrConfiguration     string
		ImageScanningEnabled string
	}{
		EcrConfiguration:     "EcrConfiguration",
		ImageScanningEnabled: "ImageScanningEnabled",
	}

	// AWS_ImageBuilder_Image_ImageScanningConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::ImageBuilder::Image.ImageScanningConfiguration.
	AWS_ImageBuilder_Image_ImageScanningConfiguration__PropertiesSlice = []string{
		AWS_ImageBuilder_Image_ImageScanningConfiguration__PropertiesMap.EcrConfiguration,
		AWS_ImageBuilder_Image_ImageScanningConfiguration__PropertiesMap.ImageScanningEnabled,
	}
)

// AWS_ImageBuilder_Image_ImageScanningConfiguration is a binding for AWS::ImageBuilder::Image.ImageScanningConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagescanningconfiguration.html
type AWS_ImageBuilder_Image_ImageScanningConfiguration struct {
	// EcrConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagescanningconfiguration.html#cfn-imagebuilder-image-imagescanningconfiguration-ecrconfiguration
	EcrConfiguration cfz.Expression[AWS_ImageBuilder_Image_EcrConfiguration] `json:"EcrConfiguration,omitempty"`

	// ImageScanningEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagescanningconfiguration.html#cfn-imagebuilder-image-imagescanningconfiguration-imagescanningenabled
	ImageScanningEnabled cfz.Expression[bool] `json:"ImageScanningEnabled,omitempty"`
}

// New__AWS_ImageBuilder_Image_ImageScanningConfiguration initializes a new AWS_ImageBuilder_Image_ImageScanningConfiguration.
func New__AWS_ImageBuilder_Image_ImageScanningConfiguration() AWS_ImageBuilder_Image_ImageScanningConfiguration {
	return AWS_ImageBuilder_Image_ImageScanningConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_ImageBuilder_Image_ImageScanningConfiguration) GetType() string {
	return AWS_ImageBuilder_Image_ImageScanningConfiguration__Type
}

// Set__EcrConfiguration updates property "EcrConfiguration".
func (t AWS_ImageBuilder_Image_ImageScanningConfiguration) Set__EcrConfiguration(v cfz.Expression[AWS_ImageBuilder_Image_EcrConfiguration]) AWS_ImageBuilder_Image_ImageScanningConfiguration {
	t.EcrConfiguration = v
	return t
}

// SetV__EcrConfiguration updates property "EcrConfiguration".
func (t AWS_ImageBuilder_Image_ImageScanningConfiguration) SetV__EcrConfiguration(v AWS_ImageBuilder_Image_EcrConfiguration) AWS_ImageBuilder_Image_ImageScanningConfiguration {
	t.EcrConfiguration = cfz.V(v)
	return t
}

// Set__ImageScanningEnabled updates property "ImageScanningEnabled".
func (t AWS_ImageBuilder_Image_ImageScanningConfiguration) Set__ImageScanningEnabled(v cfz.Expression[bool]) AWS_ImageBuilder_Image_ImageScanningConfiguration {
	t.ImageScanningEnabled = v
	return t
}

// SetV__ImageScanningEnabled updates property "ImageScanningEnabled".
func (t AWS_ImageBuilder_Image_ImageScanningConfiguration) SetV__ImageScanningEnabled(v bool) AWS_ImageBuilder_Image_ImageScanningConfiguration {
	t.ImageScanningEnabled = cfz.V(v)
	return t
}
