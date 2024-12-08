// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codedeploy

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodeDeploy_DeploymentGroup_DeploymentStyle__Type is the CloudFormation type for AWS::CodeDeploy::DeploymentGroup.DeploymentStyle.
	AWS_CodeDeploy_DeploymentGroup_DeploymentStyle__Type = "AWS::CodeDeploy::DeploymentGroup.DeploymentStyle"
)

var (
	// AWS_CodeDeploy_DeploymentGroup_DeploymentStyle__PropertiesMap reports all the CloudFormation properties for AWS::CodeDeploy::DeploymentGroup.DeploymentStyle.
	AWS_CodeDeploy_DeploymentGroup_DeploymentStyle__PropertiesMap = struct {
		DeploymentOption string
		DeploymentType   string
	}{
		DeploymentOption: "DeploymentOption",
		DeploymentType:   "DeploymentType",
	}

	// AWS_CodeDeploy_DeploymentGroup_DeploymentStyle__PropertiesSlice reports all the CloudFormation properties for AWS::CodeDeploy::DeploymentGroup.DeploymentStyle.
	AWS_CodeDeploy_DeploymentGroup_DeploymentStyle__PropertiesSlice = []string{
		AWS_CodeDeploy_DeploymentGroup_DeploymentStyle__PropertiesMap.DeploymentOption,
		AWS_CodeDeploy_DeploymentGroup_DeploymentStyle__PropertiesMap.DeploymentType,
	}
)

// AWS_CodeDeploy_DeploymentGroup_DeploymentStyle is a binding for AWS::CodeDeploy::DeploymentGroup.DeploymentStyle.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deploymentstyle.html
type AWS_CodeDeploy_DeploymentGroup_DeploymentStyle struct {
	// DeploymentOption is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deploymentstyle.html#cfn-codedeploy-deploymentgroup-deploymentstyle-deploymentoption
	DeploymentOption cfz.Expression[string] `json:"DeploymentOption,omitempty"`

	// DeploymentType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deploymentstyle.html#cfn-codedeploy-deploymentgroup-deploymentstyle-deploymenttype
	DeploymentType cfz.Expression[string] `json:"DeploymentType,omitempty"`
}

// New__AWS_CodeDeploy_DeploymentGroup_DeploymentStyle initializes a new AWS_CodeDeploy_DeploymentGroup_DeploymentStyle.
func New__AWS_CodeDeploy_DeploymentGroup_DeploymentStyle() AWS_CodeDeploy_DeploymentGroup_DeploymentStyle {
	return AWS_CodeDeploy_DeploymentGroup_DeploymentStyle{}
}

// GetType returns the CloudFormation type.
func (AWS_CodeDeploy_DeploymentGroup_DeploymentStyle) GetType() string {
	return AWS_CodeDeploy_DeploymentGroup_DeploymentStyle__Type
}

// Set__DeploymentOption updates property "DeploymentOption".
func (t AWS_CodeDeploy_DeploymentGroup_DeploymentStyle) Set__DeploymentOption(v cfz.Expression[string]) AWS_CodeDeploy_DeploymentGroup_DeploymentStyle {
	t.DeploymentOption = v
	return t
}

// SetV__DeploymentOption updates property "DeploymentOption".
func (t AWS_CodeDeploy_DeploymentGroup_DeploymentStyle) SetV__DeploymentOption(v string) AWS_CodeDeploy_DeploymentGroup_DeploymentStyle {
	t.DeploymentOption = cfz.V(v)
	return t
}

// Set__DeploymentType updates property "DeploymentType".
func (t AWS_CodeDeploy_DeploymentGroup_DeploymentStyle) Set__DeploymentType(v cfz.Expression[string]) AWS_CodeDeploy_DeploymentGroup_DeploymentStyle {
	t.DeploymentType = v
	return t
}

// SetV__DeploymentType updates property "DeploymentType".
func (t AWS_CodeDeploy_DeploymentGroup_DeploymentStyle) SetV__DeploymentType(v string) AWS_CodeDeploy_DeploymentGroup_DeploymentStyle {
	t.DeploymentType = cfz.V(v)
	return t
}
