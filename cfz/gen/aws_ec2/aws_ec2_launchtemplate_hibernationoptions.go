// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_LaunchTemplate_HibernationOptions__Type is the CloudFormation type for AWS::EC2::LaunchTemplate.HibernationOptions.
	AWS_EC2_LaunchTemplate_HibernationOptions__Type = "AWS::EC2::LaunchTemplate.HibernationOptions"
)

var (
	// AWS_EC2_LaunchTemplate_HibernationOptions__PropertiesMap reports all the CloudFormation properties for AWS::EC2::LaunchTemplate.HibernationOptions.
	AWS_EC2_LaunchTemplate_HibernationOptions__PropertiesMap = struct {
		Configured string
	}{
		Configured: "Configured",
	}

	// AWS_EC2_LaunchTemplate_HibernationOptions__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::LaunchTemplate.HibernationOptions.
	AWS_EC2_LaunchTemplate_HibernationOptions__PropertiesSlice = []string{
		AWS_EC2_LaunchTemplate_HibernationOptions__PropertiesMap.Configured,
	}
)

// AWS_EC2_LaunchTemplate_HibernationOptions is a binding for AWS::EC2::LaunchTemplate.HibernationOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-hibernationoptions.html
type AWS_EC2_LaunchTemplate_HibernationOptions struct {
	// Configured is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-hibernationoptions.html#cfn-ec2-launchtemplate-hibernationoptions-configured
	Configured cfz.Expression[bool] `json:"Configured,omitempty"`
}

// New__AWS_EC2_LaunchTemplate_HibernationOptions initializes a new AWS_EC2_LaunchTemplate_HibernationOptions.
func New__AWS_EC2_LaunchTemplate_HibernationOptions() AWS_EC2_LaunchTemplate_HibernationOptions {
	return AWS_EC2_LaunchTemplate_HibernationOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_LaunchTemplate_HibernationOptions) GetType() string {
	return AWS_EC2_LaunchTemplate_HibernationOptions__Type
}

// Set__Configured updates property "Configured".
func (t AWS_EC2_LaunchTemplate_HibernationOptions) Set__Configured(v cfz.Expression[bool]) AWS_EC2_LaunchTemplate_HibernationOptions {
	t.Configured = v
	return t
}

// SetV__Configured updates property "Configured".
func (t AWS_EC2_LaunchTemplate_HibernationOptions) SetV__Configured(v bool) AWS_EC2_LaunchTemplate_HibernationOptions {
	t.Configured = cfz.V(v)
	return t
}
