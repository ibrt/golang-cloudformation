// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emr

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EMR_Cluster_EbsConfiguration__Type is the CloudFormation type for AWS::EMR::Cluster.EbsConfiguration.
	AWS_EMR_Cluster_EbsConfiguration__Type = "AWS::EMR::Cluster.EbsConfiguration"
)

var (
	// AWS_EMR_Cluster_EbsConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::EMR::Cluster.EbsConfiguration.
	AWS_EMR_Cluster_EbsConfiguration__PropertiesMap = struct {
		EbsBlockDeviceConfigs string
		EbsOptimized          string
	}{
		EbsBlockDeviceConfigs: "EbsBlockDeviceConfigs",
		EbsOptimized:          "EbsOptimized",
	}

	// AWS_EMR_Cluster_EbsConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::EMR::Cluster.EbsConfiguration.
	AWS_EMR_Cluster_EbsConfiguration__PropertiesSlice = []string{
		AWS_EMR_Cluster_EbsConfiguration__PropertiesMap.EbsBlockDeviceConfigs,
		AWS_EMR_Cluster_EbsConfiguration__PropertiesMap.EbsOptimized,
	}
)

// AWS_EMR_Cluster_EbsConfiguration is a binding for AWS::EMR::Cluster.EbsConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-ebsconfiguration.html
type AWS_EMR_Cluster_EbsConfiguration struct {
	// EbsBlockDeviceConfigs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-ebsconfiguration.html#cfn-elasticmapreduce-cluster-ebsconfiguration-ebsblockdeviceconfigs
	EbsBlockDeviceConfigs cfz.ExpressionSlice[AWS_EMR_Cluster_EbsBlockDeviceConfig] `json:"EbsBlockDeviceConfigs,omitempty"`

	// EbsOptimized is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-ebsconfiguration.html#cfn-elasticmapreduce-cluster-ebsconfiguration-ebsoptimized
	EbsOptimized cfz.Expression[bool] `json:"EbsOptimized,omitempty"`
}

// New__AWS_EMR_Cluster_EbsConfiguration initializes a new AWS_EMR_Cluster_EbsConfiguration.
func New__AWS_EMR_Cluster_EbsConfiguration() AWS_EMR_Cluster_EbsConfiguration {
	return AWS_EMR_Cluster_EbsConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_EMR_Cluster_EbsConfiguration) GetType() string {
	return AWS_EMR_Cluster_EbsConfiguration__Type
}

// Set__EbsBlockDeviceConfigs updates property "EbsBlockDeviceConfigs".
func (t AWS_EMR_Cluster_EbsConfiguration) Set__EbsBlockDeviceConfigs(v cfz.ExpressionSlice[AWS_EMR_Cluster_EbsBlockDeviceConfig]) AWS_EMR_Cluster_EbsConfiguration {
	t.EbsBlockDeviceConfigs = v
	return t
}

// SetS__EbsBlockDeviceConfigs updates property "EbsBlockDeviceConfigs".
func (t AWS_EMR_Cluster_EbsConfiguration) SetS__EbsBlockDeviceConfigs(v ...cfz.Expression[AWS_EMR_Cluster_EbsBlockDeviceConfig]) AWS_EMR_Cluster_EbsConfiguration {
	t.EbsBlockDeviceConfigs = cfz.S(v...)
	return t
}

// SetSV__EbsBlockDeviceConfigs updates property "EbsBlockDeviceConfigs".
func (t AWS_EMR_Cluster_EbsConfiguration) SetSV__EbsBlockDeviceConfigs(v ...AWS_EMR_Cluster_EbsBlockDeviceConfig) AWS_EMR_Cluster_EbsConfiguration {
	t.EbsBlockDeviceConfigs = cfz.SV(v...)
	return t
}

// Set__EbsOptimized updates property "EbsOptimized".
func (t AWS_EMR_Cluster_EbsConfiguration) Set__EbsOptimized(v cfz.Expression[bool]) AWS_EMR_Cluster_EbsConfiguration {
	t.EbsOptimized = v
	return t
}

// SetV__EbsOptimized updates property "EbsOptimized".
func (t AWS_EMR_Cluster_EbsConfiguration) SetV__EbsOptimized(v bool) AWS_EMR_Cluster_EbsConfiguration {
	t.EbsOptimized = cfz.V(v)
	return t
}
