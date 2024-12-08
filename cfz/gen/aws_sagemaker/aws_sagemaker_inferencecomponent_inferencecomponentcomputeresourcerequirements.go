// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__Type is the CloudFormation type for AWS::SageMaker::InferenceComponent.InferenceComponentComputeResourceRequirements.
	AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__Type = "AWS::SageMaker::InferenceComponent.InferenceComponentComputeResourceRequirements"
)

var (
	// AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::InferenceComponent.InferenceComponentComputeResourceRequirements.
	AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__PropertiesMap = struct {
		MaxMemoryRequiredInMb              string
		MinMemoryRequiredInMb              string
		NumberOfAcceleratorDevicesRequired string
		NumberOfCpuCoresRequired           string
	}{
		MaxMemoryRequiredInMb:              "MaxMemoryRequiredInMb",
		MinMemoryRequiredInMb:              "MinMemoryRequiredInMb",
		NumberOfAcceleratorDevicesRequired: "NumberOfAcceleratorDevicesRequired",
		NumberOfCpuCoresRequired:           "NumberOfCpuCoresRequired",
	}

	// AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::InferenceComponent.InferenceComponentComputeResourceRequirements.
	AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__PropertiesSlice = []string{
		AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__PropertiesMap.MaxMemoryRequiredInMb,
		AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__PropertiesMap.MinMemoryRequiredInMb,
		AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__PropertiesMap.NumberOfAcceleratorDevicesRequired,
		AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__PropertiesMap.NumberOfCpuCoresRequired,
	}
)

// AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements is a binding for AWS::SageMaker::InferenceComponent.InferenceComponentComputeResourceRequirements.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html
type AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements struct {
	// MaxMemoryRequiredInMb is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html#cfn-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements-maxmemoryrequiredinmb
	MaxMemoryRequiredInMb cfz.Expression[int32] `json:"MaxMemoryRequiredInMb,omitempty"`

	// MinMemoryRequiredInMb is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html#cfn-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements-minmemoryrequiredinmb
	MinMemoryRequiredInMb cfz.Expression[int32] `json:"MinMemoryRequiredInMb,omitempty"`

	// NumberOfAcceleratorDevicesRequired is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html#cfn-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements-numberofacceleratordevicesrequired
	NumberOfAcceleratorDevicesRequired cfz.Expression[float64] `json:"NumberOfAcceleratorDevicesRequired,omitempty"`

	// NumberOfCpuCoresRequired is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html#cfn-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements-numberofcpucoresrequired
	NumberOfCpuCoresRequired cfz.Expression[float64] `json:"NumberOfCpuCoresRequired,omitempty"`
}

// New__AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements initializes a new AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements.
func New__AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements() AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements {
	return AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements) GetType() string {
	return AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements__Type
}

// Set__MaxMemoryRequiredInMb updates property "MaxMemoryRequiredInMb".
func (t AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements) Set__MaxMemoryRequiredInMb(v cfz.Expression[int32]) AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements {
	t.MaxMemoryRequiredInMb = v
	return t
}

// SetV__MaxMemoryRequiredInMb updates property "MaxMemoryRequiredInMb".
func (t AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements) SetV__MaxMemoryRequiredInMb(v int32) AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements {
	t.MaxMemoryRequiredInMb = cfz.V(v)
	return t
}

// Set__MinMemoryRequiredInMb updates property "MinMemoryRequiredInMb".
func (t AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements) Set__MinMemoryRequiredInMb(v cfz.Expression[int32]) AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements {
	t.MinMemoryRequiredInMb = v
	return t
}

// SetV__MinMemoryRequiredInMb updates property "MinMemoryRequiredInMb".
func (t AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements) SetV__MinMemoryRequiredInMb(v int32) AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements {
	t.MinMemoryRequiredInMb = cfz.V(v)
	return t
}

// Set__NumberOfAcceleratorDevicesRequired updates property "NumberOfAcceleratorDevicesRequired".
func (t AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements) Set__NumberOfAcceleratorDevicesRequired(v cfz.Expression[float64]) AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements {
	t.NumberOfAcceleratorDevicesRequired = v
	return t
}

// SetV__NumberOfAcceleratorDevicesRequired updates property "NumberOfAcceleratorDevicesRequired".
func (t AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements) SetV__NumberOfAcceleratorDevicesRequired(v float64) AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements {
	t.NumberOfAcceleratorDevicesRequired = cfz.V(v)
	return t
}

// Set__NumberOfCpuCoresRequired updates property "NumberOfCpuCoresRequired".
func (t AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements) Set__NumberOfCpuCoresRequired(v cfz.Expression[float64]) AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements {
	t.NumberOfCpuCoresRequired = v
	return t
}

// SetV__NumberOfCpuCoresRequired updates property "NumberOfCpuCoresRequired".
func (t AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements) SetV__NumberOfCpuCoresRequired(v float64) AWS_SageMaker_InferenceComponent_InferenceComponentComputeResourceRequirements {
	t.NumberOfCpuCoresRequired = cfz.V(v)
	return t
}
