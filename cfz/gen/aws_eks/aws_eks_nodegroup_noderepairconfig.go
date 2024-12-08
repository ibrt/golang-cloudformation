// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_eks

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EKS_Nodegroup_NodeRepairConfig__Type is the CloudFormation type for AWS::EKS::Nodegroup.NodeRepairConfig.
	AWS_EKS_Nodegroup_NodeRepairConfig__Type = "AWS::EKS::Nodegroup.NodeRepairConfig"
)

var (
	// AWS_EKS_Nodegroup_NodeRepairConfig__PropertiesMap reports all the CloudFormation properties for AWS::EKS::Nodegroup.NodeRepairConfig.
	AWS_EKS_Nodegroup_NodeRepairConfig__PropertiesMap = struct {
		Enabled string
	}{
		Enabled: "Enabled",
	}

	// AWS_EKS_Nodegroup_NodeRepairConfig__PropertiesSlice reports all the CloudFormation properties for AWS::EKS::Nodegroup.NodeRepairConfig.
	AWS_EKS_Nodegroup_NodeRepairConfig__PropertiesSlice = []string{
		AWS_EKS_Nodegroup_NodeRepairConfig__PropertiesMap.Enabled,
	}
)

// AWS_EKS_Nodegroup_NodeRepairConfig is a binding for AWS::EKS::Nodegroup.NodeRepairConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html
type AWS_EKS_Nodegroup_NodeRepairConfig struct {
	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html#cfn-eks-nodegroup-noderepairconfig-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`
}

// New__AWS_EKS_Nodegroup_NodeRepairConfig initializes a new AWS_EKS_Nodegroup_NodeRepairConfig.
func New__AWS_EKS_Nodegroup_NodeRepairConfig() AWS_EKS_Nodegroup_NodeRepairConfig {
	return AWS_EKS_Nodegroup_NodeRepairConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_EKS_Nodegroup_NodeRepairConfig) GetType() string {
	return AWS_EKS_Nodegroup_NodeRepairConfig__Type
}

// Set__Enabled updates property "Enabled".
func (t AWS_EKS_Nodegroup_NodeRepairConfig) Set__Enabled(v cfz.Expression[bool]) AWS_EKS_Nodegroup_NodeRepairConfig {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_EKS_Nodegroup_NodeRepairConfig) SetV__Enabled(v bool) AWS_EKS_Nodegroup_NodeRepairConfig {
	t.Enabled = cfz.V(v)
	return t
}
