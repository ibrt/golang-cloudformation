// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codebuild

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodeBuild_Fleet_ComputeConfiguration__Type is the CloudFormation type for AWS::CodeBuild::Fleet.ComputeConfiguration.
	AWS_CodeBuild_Fleet_ComputeConfiguration__Type = "AWS::CodeBuild::Fleet.ComputeConfiguration"
)

var (
	// AWS_CodeBuild_Fleet_ComputeConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::CodeBuild::Fleet.ComputeConfiguration.
	AWS_CodeBuild_Fleet_ComputeConfiguration__PropertiesMap = struct {
		Disk        string
		MachineType string
		Memory      string
		VCpu        string
	}{
		Disk:        "disk",
		MachineType: "machineType",
		Memory:      "memory",
		VCpu:        "vCpu",
	}

	// AWS_CodeBuild_Fleet_ComputeConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::CodeBuild::Fleet.ComputeConfiguration.
	AWS_CodeBuild_Fleet_ComputeConfiguration__PropertiesSlice = []string{
		AWS_CodeBuild_Fleet_ComputeConfiguration__PropertiesMap.Disk,
		AWS_CodeBuild_Fleet_ComputeConfiguration__PropertiesMap.MachineType,
		AWS_CodeBuild_Fleet_ComputeConfiguration__PropertiesMap.Memory,
		AWS_CodeBuild_Fleet_ComputeConfiguration__PropertiesMap.VCpu,
	}
)

// AWS_CodeBuild_Fleet_ComputeConfiguration is a binding for AWS::CodeBuild::Fleet.ComputeConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html
type AWS_CodeBuild_Fleet_ComputeConfiguration struct {
	// disk is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html#cfn-codebuild-fleet-computeconfiguration-disk
	Disk cfz.Expression[int32] `json:"disk,omitempty"`

	// machineType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html#cfn-codebuild-fleet-computeconfiguration-machinetype
	MachineType cfz.Expression[string] `json:"machineType,omitempty"`

	// memory is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html#cfn-codebuild-fleet-computeconfiguration-memory
	Memory cfz.Expression[int32] `json:"memory,omitempty"`

	// vCpu is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html#cfn-codebuild-fleet-computeconfiguration-vcpu
	VCpu cfz.Expression[int32] `json:"vCpu,omitempty"`
}

// New__AWS_CodeBuild_Fleet_ComputeConfiguration initializes a new AWS_CodeBuild_Fleet_ComputeConfiguration.
func New__AWS_CodeBuild_Fleet_ComputeConfiguration() AWS_CodeBuild_Fleet_ComputeConfiguration {
	return AWS_CodeBuild_Fleet_ComputeConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_CodeBuild_Fleet_ComputeConfiguration) GetType() string {
	return AWS_CodeBuild_Fleet_ComputeConfiguration__Type
}

// Set__Disk updates property "disk".
func (t AWS_CodeBuild_Fleet_ComputeConfiguration) Set__Disk(v cfz.Expression[int32]) AWS_CodeBuild_Fleet_ComputeConfiguration {
	t.Disk = v
	return t
}

// SetV__Disk updates property "disk".
func (t AWS_CodeBuild_Fleet_ComputeConfiguration) SetV__Disk(v int32) AWS_CodeBuild_Fleet_ComputeConfiguration {
	t.Disk = cfz.V(v)
	return t
}

// Set__MachineType updates property "machineType".
func (t AWS_CodeBuild_Fleet_ComputeConfiguration) Set__MachineType(v cfz.Expression[string]) AWS_CodeBuild_Fleet_ComputeConfiguration {
	t.MachineType = v
	return t
}

// SetV__MachineType updates property "machineType".
func (t AWS_CodeBuild_Fleet_ComputeConfiguration) SetV__MachineType(v string) AWS_CodeBuild_Fleet_ComputeConfiguration {
	t.MachineType = cfz.V(v)
	return t
}

// Set__Memory updates property "memory".
func (t AWS_CodeBuild_Fleet_ComputeConfiguration) Set__Memory(v cfz.Expression[int32]) AWS_CodeBuild_Fleet_ComputeConfiguration {
	t.Memory = v
	return t
}

// SetV__Memory updates property "memory".
func (t AWS_CodeBuild_Fleet_ComputeConfiguration) SetV__Memory(v int32) AWS_CodeBuild_Fleet_ComputeConfiguration {
	t.Memory = cfz.V(v)
	return t
}

// Set__VCpu updates property "vCpu".
func (t AWS_CodeBuild_Fleet_ComputeConfiguration) Set__VCpu(v cfz.Expression[int32]) AWS_CodeBuild_Fleet_ComputeConfiguration {
	t.VCpu = v
	return t
}

// SetV__VCpu updates property "vCpu".
func (t AWS_CodeBuild_Fleet_ComputeConfiguration) SetV__VCpu(v int32) AWS_CodeBuild_Fleet_ComputeConfiguration {
	t.VCpu = cfz.V(v)
	return t
}
