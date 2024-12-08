// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_MonitoringSchedule_NetworkConfig__Type is the CloudFormation type for AWS::SageMaker::MonitoringSchedule.NetworkConfig.
	AWS_SageMaker_MonitoringSchedule_NetworkConfig__Type = "AWS::SageMaker::MonitoringSchedule.NetworkConfig"
)

var (
	// AWS_SageMaker_MonitoringSchedule_NetworkConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::MonitoringSchedule.NetworkConfig.
	AWS_SageMaker_MonitoringSchedule_NetworkConfig__PropertiesMap = struct {
		EnableInterContainerTrafficEncryption string
		EnableNetworkIsolation                string
		VpcConfig                             string
	}{
		EnableInterContainerTrafficEncryption: "EnableInterContainerTrafficEncryption",
		EnableNetworkIsolation:                "EnableNetworkIsolation",
		VpcConfig:                             "VpcConfig",
	}

	// AWS_SageMaker_MonitoringSchedule_NetworkConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::MonitoringSchedule.NetworkConfig.
	AWS_SageMaker_MonitoringSchedule_NetworkConfig__PropertiesSlice = []string{
		AWS_SageMaker_MonitoringSchedule_NetworkConfig__PropertiesMap.EnableInterContainerTrafficEncryption,
		AWS_SageMaker_MonitoringSchedule_NetworkConfig__PropertiesMap.EnableNetworkIsolation,
		AWS_SageMaker_MonitoringSchedule_NetworkConfig__PropertiesMap.VpcConfig,
	}
)

// AWS_SageMaker_MonitoringSchedule_NetworkConfig is a binding for AWS::SageMaker::MonitoringSchedule.NetworkConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-networkconfig.html
type AWS_SageMaker_MonitoringSchedule_NetworkConfig struct {
	// EnableInterContainerTrafficEncryption is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-networkconfig.html#cfn-sagemaker-monitoringschedule-networkconfig-enableintercontainertrafficencryption
	EnableInterContainerTrafficEncryption cfz.Expression[bool] `json:"EnableInterContainerTrafficEncryption,omitempty"`

	// EnableNetworkIsolation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-networkconfig.html#cfn-sagemaker-monitoringschedule-networkconfig-enablenetworkisolation
	EnableNetworkIsolation cfz.Expression[bool] `json:"EnableNetworkIsolation,omitempty"`

	// VpcConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-networkconfig.html#cfn-sagemaker-monitoringschedule-networkconfig-vpcconfig
	VpcConfig cfz.Expression[AWS_SageMaker_MonitoringSchedule_VpcConfig] `json:"VpcConfig,omitempty"`
}

// New__AWS_SageMaker_MonitoringSchedule_NetworkConfig initializes a new AWS_SageMaker_MonitoringSchedule_NetworkConfig.
func New__AWS_SageMaker_MonitoringSchedule_NetworkConfig() AWS_SageMaker_MonitoringSchedule_NetworkConfig {
	return AWS_SageMaker_MonitoringSchedule_NetworkConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_MonitoringSchedule_NetworkConfig) GetType() string {
	return AWS_SageMaker_MonitoringSchedule_NetworkConfig__Type
}

// Set__EnableInterContainerTrafficEncryption updates property "EnableInterContainerTrafficEncryption".
func (t AWS_SageMaker_MonitoringSchedule_NetworkConfig) Set__EnableInterContainerTrafficEncryption(v cfz.Expression[bool]) AWS_SageMaker_MonitoringSchedule_NetworkConfig {
	t.EnableInterContainerTrafficEncryption = v
	return t
}

// SetV__EnableInterContainerTrafficEncryption updates property "EnableInterContainerTrafficEncryption".
func (t AWS_SageMaker_MonitoringSchedule_NetworkConfig) SetV__EnableInterContainerTrafficEncryption(v bool) AWS_SageMaker_MonitoringSchedule_NetworkConfig {
	t.EnableInterContainerTrafficEncryption = cfz.V(v)
	return t
}

// Set__EnableNetworkIsolation updates property "EnableNetworkIsolation".
func (t AWS_SageMaker_MonitoringSchedule_NetworkConfig) Set__EnableNetworkIsolation(v cfz.Expression[bool]) AWS_SageMaker_MonitoringSchedule_NetworkConfig {
	t.EnableNetworkIsolation = v
	return t
}

// SetV__EnableNetworkIsolation updates property "EnableNetworkIsolation".
func (t AWS_SageMaker_MonitoringSchedule_NetworkConfig) SetV__EnableNetworkIsolation(v bool) AWS_SageMaker_MonitoringSchedule_NetworkConfig {
	t.EnableNetworkIsolation = cfz.V(v)
	return t
}

// Set__VpcConfig updates property "VpcConfig".
func (t AWS_SageMaker_MonitoringSchedule_NetworkConfig) Set__VpcConfig(v cfz.Expression[AWS_SageMaker_MonitoringSchedule_VpcConfig]) AWS_SageMaker_MonitoringSchedule_NetworkConfig {
	t.VpcConfig = v
	return t
}

// SetV__VpcConfig updates property "VpcConfig".
func (t AWS_SageMaker_MonitoringSchedule_NetworkConfig) SetV__VpcConfig(v AWS_SageMaker_MonitoringSchedule_VpcConfig) AWS_SageMaker_MonitoringSchedule_NetworkConfig {
	t.VpcConfig = cfz.V(v)
	return t
}
