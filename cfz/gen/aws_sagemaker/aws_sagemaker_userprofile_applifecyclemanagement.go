// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_UserProfile_AppLifecycleManagement__Type is the CloudFormation type for AWS::SageMaker::UserProfile.AppLifecycleManagement.
	AWS_SageMaker_UserProfile_AppLifecycleManagement__Type = "AWS::SageMaker::UserProfile.AppLifecycleManagement"
)

var (
	// AWS_SageMaker_UserProfile_AppLifecycleManagement__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::UserProfile.AppLifecycleManagement.
	AWS_SageMaker_UserProfile_AppLifecycleManagement__PropertiesMap = struct {
		IdleSettings string
	}{
		IdleSettings: "IdleSettings",
	}

	// AWS_SageMaker_UserProfile_AppLifecycleManagement__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::UserProfile.AppLifecycleManagement.
	AWS_SageMaker_UserProfile_AppLifecycleManagement__PropertiesSlice = []string{
		AWS_SageMaker_UserProfile_AppLifecycleManagement__PropertiesMap.IdleSettings,
	}
)

// AWS_SageMaker_UserProfile_AppLifecycleManagement is a binding for AWS::SageMaker::UserProfile.AppLifecycleManagement.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-applifecyclemanagement.html
type AWS_SageMaker_UserProfile_AppLifecycleManagement struct {
	// IdleSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-applifecyclemanagement.html#cfn-sagemaker-userprofile-applifecyclemanagement-idlesettings
	IdleSettings cfz.Expression[AWS_SageMaker_UserProfile_IdleSettings] `json:"IdleSettings,omitempty"`
}

// New__AWS_SageMaker_UserProfile_AppLifecycleManagement initializes a new AWS_SageMaker_UserProfile_AppLifecycleManagement.
func New__AWS_SageMaker_UserProfile_AppLifecycleManagement() AWS_SageMaker_UserProfile_AppLifecycleManagement {
	return AWS_SageMaker_UserProfile_AppLifecycleManagement{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_UserProfile_AppLifecycleManagement) GetType() string {
	return AWS_SageMaker_UserProfile_AppLifecycleManagement__Type
}

// Set__IdleSettings updates property "IdleSettings".
func (t AWS_SageMaker_UserProfile_AppLifecycleManagement) Set__IdleSettings(v cfz.Expression[AWS_SageMaker_UserProfile_IdleSettings]) AWS_SageMaker_UserProfile_AppLifecycleManagement {
	t.IdleSettings = v
	return t
}

// SetV__IdleSettings updates property "IdleSettings".
func (t AWS_SageMaker_UserProfile_AppLifecycleManagement) SetV__IdleSettings(v AWS_SageMaker_UserProfile_IdleSettings) AWS_SageMaker_UserProfile_AppLifecycleManagement {
	t.IdleSettings = cfz.V(v)
	return t
}
