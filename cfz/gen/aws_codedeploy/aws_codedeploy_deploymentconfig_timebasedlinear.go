// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codedeploy

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear__Type is the CloudFormation type for AWS::CodeDeploy::DeploymentConfig.TimeBasedLinear.
	AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear__Type = "AWS::CodeDeploy::DeploymentConfig.TimeBasedLinear"
)

var (
	// AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear__PropertiesMap reports all the CloudFormation properties for AWS::CodeDeploy::DeploymentConfig.TimeBasedLinear.
	AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear__PropertiesMap = struct {
		LinearInterval   string
		LinearPercentage string
	}{
		LinearInterval:   "LinearInterval",
		LinearPercentage: "LinearPercentage",
	}

	// AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear__PropertiesSlice reports all the CloudFormation properties for AWS::CodeDeploy::DeploymentConfig.TimeBasedLinear.
	AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear__PropertiesSlice = []string{
		AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear__PropertiesMap.LinearInterval,
		AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear__PropertiesMap.LinearPercentage,
	}
)

// AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear is a binding for AWS::CodeDeploy::DeploymentConfig.TimeBasedLinear.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-timebasedlinear.html
type AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear struct {
	// LinearInterval is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-timebasedlinear.html#cfn-codedeploy-deploymentconfig-timebasedlinear-linearinterval
	LinearInterval cfz.Expression[int32] `json:"LinearInterval,omitempty"`

	// LinearPercentage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-timebasedlinear.html#cfn-codedeploy-deploymentconfig-timebasedlinear-linearpercentage
	LinearPercentage cfz.Expression[int32] `json:"LinearPercentage,omitempty"`
}

// New__AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear initializes a new AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear.
func New__AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear() AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear {
	return AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear{}
}

// GetType returns the CloudFormation type.
func (AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear) GetType() string {
	return AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear__Type
}

// Set__LinearInterval updates property "LinearInterval".
func (t AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear) Set__LinearInterval(v cfz.Expression[int32]) AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear {
	t.LinearInterval = v
	return t
}

// SetV__LinearInterval updates property "LinearInterval".
func (t AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear) SetV__LinearInterval(v int32) AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear {
	t.LinearInterval = cfz.V(v)
	return t
}

// Set__LinearPercentage updates property "LinearPercentage".
func (t AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear) Set__LinearPercentage(v cfz.Expression[int32]) AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear {
	t.LinearPercentage = v
	return t
}

// SetV__LinearPercentage updates property "LinearPercentage".
func (t AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear) SetV__LinearPercentage(v int32) AWS_CodeDeploy_DeploymentConfig_TimeBasedLinear {
	t.LinearPercentage = cfz.V(v)
	return t
}
