// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_UserProfile_CustomImage__Type is the CloudFormation type for AWS::SageMaker::UserProfile.CustomImage.
	AWS_SageMaker_UserProfile_CustomImage__Type = "AWS::SageMaker::UserProfile.CustomImage"
)

var (
	// AWS_SageMaker_UserProfile_CustomImage__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::UserProfile.CustomImage.
	AWS_SageMaker_UserProfile_CustomImage__PropertiesMap = struct {
		AppImageConfigName string
		ImageName          string
		ImageVersionNumber string
	}{
		AppImageConfigName: "AppImageConfigName",
		ImageName:          "ImageName",
		ImageVersionNumber: "ImageVersionNumber",
	}

	// AWS_SageMaker_UserProfile_CustomImage__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::UserProfile.CustomImage.
	AWS_SageMaker_UserProfile_CustomImage__PropertiesSlice = []string{
		AWS_SageMaker_UserProfile_CustomImage__PropertiesMap.AppImageConfigName,
		AWS_SageMaker_UserProfile_CustomImage__PropertiesMap.ImageName,
		AWS_SageMaker_UserProfile_CustomImage__PropertiesMap.ImageVersionNumber,
	}
)

// AWS_SageMaker_UserProfile_CustomImage is a binding for AWS::SageMaker::UserProfile.CustomImage.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-customimage.html
type AWS_SageMaker_UserProfile_CustomImage struct {
	// AppImageConfigName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-customimage.html#cfn-sagemaker-userprofile-customimage-appimageconfigname
	AppImageConfigName cfz.Expression[string] `json:"AppImageConfigName,omitempty"`

	// ImageName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-customimage.html#cfn-sagemaker-userprofile-customimage-imagename
	ImageName cfz.Expression[string] `json:"ImageName,omitempty"`

	// ImageVersionNumber is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-customimage.html#cfn-sagemaker-userprofile-customimage-imageversionnumber
	ImageVersionNumber cfz.Expression[int32] `json:"ImageVersionNumber,omitempty"`
}

// New__AWS_SageMaker_UserProfile_CustomImage initializes a new AWS_SageMaker_UserProfile_CustomImage.
func New__AWS_SageMaker_UserProfile_CustomImage() AWS_SageMaker_UserProfile_CustomImage {
	return AWS_SageMaker_UserProfile_CustomImage{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_UserProfile_CustomImage) GetType() string {
	return AWS_SageMaker_UserProfile_CustomImage__Type
}

// Set__AppImageConfigName updates property "AppImageConfigName".
func (t AWS_SageMaker_UserProfile_CustomImage) Set__AppImageConfigName(v cfz.Expression[string]) AWS_SageMaker_UserProfile_CustomImage {
	t.AppImageConfigName = v
	return t
}

// SetV__AppImageConfigName updates property "AppImageConfigName".
func (t AWS_SageMaker_UserProfile_CustomImage) SetV__AppImageConfigName(v string) AWS_SageMaker_UserProfile_CustomImage {
	t.AppImageConfigName = cfz.V(v)
	return t
}

// Set__ImageName updates property "ImageName".
func (t AWS_SageMaker_UserProfile_CustomImage) Set__ImageName(v cfz.Expression[string]) AWS_SageMaker_UserProfile_CustomImage {
	t.ImageName = v
	return t
}

// SetV__ImageName updates property "ImageName".
func (t AWS_SageMaker_UserProfile_CustomImage) SetV__ImageName(v string) AWS_SageMaker_UserProfile_CustomImage {
	t.ImageName = cfz.V(v)
	return t
}

// Set__ImageVersionNumber updates property "ImageVersionNumber".
func (t AWS_SageMaker_UserProfile_CustomImage) Set__ImageVersionNumber(v cfz.Expression[int32]) AWS_SageMaker_UserProfile_CustomImage {
	t.ImageVersionNumber = v
	return t
}

// SetV__ImageVersionNumber updates property "ImageVersionNumber".
func (t AWS_SageMaker_UserProfile_CustomImage) SetV__ImageVersionNumber(v int32) AWS_SageMaker_UserProfile_CustomImage {
	t.ImageVersionNumber = cfz.V(v)
	return t
}
