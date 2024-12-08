// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_Domain_AppLifecycleManagement__Type is the CloudFormation type for AWS::SageMaker::Domain.AppLifecycleManagement.
	AWS_SageMaker_Domain_AppLifecycleManagement__Type = "AWS::SageMaker::Domain.AppLifecycleManagement"
)

var (
	// AWS_SageMaker_Domain_AppLifecycleManagement__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::Domain.AppLifecycleManagement.
	AWS_SageMaker_Domain_AppLifecycleManagement__PropertiesMap = struct {
		IdleSettings string
	}{
		IdleSettings: "IdleSettings",
	}

	// AWS_SageMaker_Domain_AppLifecycleManagement__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::Domain.AppLifecycleManagement.
	AWS_SageMaker_Domain_AppLifecycleManagement__PropertiesSlice = []string{
		AWS_SageMaker_Domain_AppLifecycleManagement__PropertiesMap.IdleSettings,
	}
)

// AWS_SageMaker_Domain_AppLifecycleManagement is a binding for AWS::SageMaker::Domain.AppLifecycleManagement.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-applifecyclemanagement.html
type AWS_SageMaker_Domain_AppLifecycleManagement struct {
	// IdleSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-applifecyclemanagement.html#cfn-sagemaker-domain-applifecyclemanagement-idlesettings
	IdleSettings cfz.Expression[AWS_SageMaker_Domain_IdleSettings] `json:"IdleSettings,omitempty"`
}

// New__AWS_SageMaker_Domain_AppLifecycleManagement initializes a new AWS_SageMaker_Domain_AppLifecycleManagement.
func New__AWS_SageMaker_Domain_AppLifecycleManagement() AWS_SageMaker_Domain_AppLifecycleManagement {
	return AWS_SageMaker_Domain_AppLifecycleManagement{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_Domain_AppLifecycleManagement) GetType() string {
	return AWS_SageMaker_Domain_AppLifecycleManagement__Type
}

// Set__IdleSettings updates property "IdleSettings".
func (t AWS_SageMaker_Domain_AppLifecycleManagement) Set__IdleSettings(v cfz.Expression[AWS_SageMaker_Domain_IdleSettings]) AWS_SageMaker_Domain_AppLifecycleManagement {
	t.IdleSettings = v
	return t
}

// SetV__IdleSettings updates property "IdleSettings".
func (t AWS_SageMaker_Domain_AppLifecycleManagement) SetV__IdleSettings(v AWS_SageMaker_Domain_IdleSettings) AWS_SageMaker_Domain_AppLifecycleManagement {
	t.IdleSettings = cfz.V(v)
	return t
}
