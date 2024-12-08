// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codedeploy

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone__Type is the CloudFormation type for AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHostsPerZone.
	AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone__Type = "AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHostsPerZone"
)

var (
	// AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone__PropertiesMap reports all the CloudFormation properties for AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHostsPerZone.
	AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone__PropertiesMap = struct {
		Type  string
		Value string
	}{
		Type:  "Type",
		Value: "Value",
	}

	// AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone__PropertiesSlice reports all the CloudFormation properties for AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHostsPerZone.
	AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone__PropertiesSlice = []string{
		AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone__PropertiesMap.Type,
		AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone__PropertiesMap.Value,
	}
)

// AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone is a binding for AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHostsPerZone.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhostsperzone.html
type AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone struct {
	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhostsperzone.html#cfn-codedeploy-deploymentconfig-minimumhealthyhostsperzone-type
	Type cfz.Expression[string] `json:"Type,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhostsperzone.html#cfn-codedeploy-deploymentconfig-minimumhealthyhostsperzone-value
	Value cfz.Expression[int32] `json:"Value,omitempty"`
}

// New__AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone initializes a new AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone.
func New__AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone() AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone {
	return AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone{}
}

// GetType returns the CloudFormation type.
func (AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone) GetType() string {
	return AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone__Type
}

// Set__Type updates property "Type".
func (t AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone) Set__Type(v cfz.Expression[string]) AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone) SetV__Type(v string) AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone {
	t.Type = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone) Set__Value(v cfz.Expression[int32]) AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone) SetV__Value(v int32) AWS_CodeDeploy_DeploymentConfig_MinimumHealthyHostsPerZone {
	t.Value = cfz.V(v)
	return t
}
