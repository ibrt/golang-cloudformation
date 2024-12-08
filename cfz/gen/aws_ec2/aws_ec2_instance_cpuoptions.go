// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_Instance_CpuOptions__Type is the CloudFormation type for AWS::EC2::Instance.CpuOptions.
	AWS_EC2_Instance_CpuOptions__Type = "AWS::EC2::Instance.CpuOptions"
)

var (
	// AWS_EC2_Instance_CpuOptions__PropertiesMap reports all the CloudFormation properties for AWS::EC2::Instance.CpuOptions.
	AWS_EC2_Instance_CpuOptions__PropertiesMap = struct {
		CoreCount      string
		ThreadsPerCore string
	}{
		CoreCount:      "CoreCount",
		ThreadsPerCore: "ThreadsPerCore",
	}

	// AWS_EC2_Instance_CpuOptions__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::Instance.CpuOptions.
	AWS_EC2_Instance_CpuOptions__PropertiesSlice = []string{
		AWS_EC2_Instance_CpuOptions__PropertiesMap.CoreCount,
		AWS_EC2_Instance_CpuOptions__PropertiesMap.ThreadsPerCore,
	}
)

// AWS_EC2_Instance_CpuOptions is a binding for AWS::EC2::Instance.CpuOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-cpuoptions.html
type AWS_EC2_Instance_CpuOptions struct {
	// CoreCount is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-cpuoptions.html#cfn-ec2-instance-cpuoptions-corecount
	CoreCount cfz.Expression[int32] `json:"CoreCount,omitempty"`

	// ThreadsPerCore is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-cpuoptions.html#cfn-ec2-instance-cpuoptions-threadspercore
	ThreadsPerCore cfz.Expression[int32] `json:"ThreadsPerCore,omitempty"`
}

// New__AWS_EC2_Instance_CpuOptions initializes a new AWS_EC2_Instance_CpuOptions.
func New__AWS_EC2_Instance_CpuOptions() AWS_EC2_Instance_CpuOptions {
	return AWS_EC2_Instance_CpuOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_Instance_CpuOptions) GetType() string {
	return AWS_EC2_Instance_CpuOptions__Type
}

// Set__CoreCount updates property "CoreCount".
func (t AWS_EC2_Instance_CpuOptions) Set__CoreCount(v cfz.Expression[int32]) AWS_EC2_Instance_CpuOptions {
	t.CoreCount = v
	return t
}

// SetV__CoreCount updates property "CoreCount".
func (t AWS_EC2_Instance_CpuOptions) SetV__CoreCount(v int32) AWS_EC2_Instance_CpuOptions {
	t.CoreCount = cfz.V(v)
	return t
}

// Set__ThreadsPerCore updates property "ThreadsPerCore".
func (t AWS_EC2_Instance_CpuOptions) Set__ThreadsPerCore(v cfz.Expression[int32]) AWS_EC2_Instance_CpuOptions {
	t.ThreadsPerCore = v
	return t
}

// SetV__ThreadsPerCore updates property "ThreadsPerCore".
func (t AWS_EC2_Instance_CpuOptions) SetV__ThreadsPerCore(v int32) AWS_EC2_Instance_CpuOptions {
	t.ThreadsPerCore = cfz.V(v)
	return t
}
