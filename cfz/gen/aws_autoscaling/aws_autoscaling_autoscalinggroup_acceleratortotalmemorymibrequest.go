// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_autoscaling

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest__Type is the CloudFormation type for AWS::AutoScaling::AutoScalingGroup.AcceleratorTotalMemoryMiBRequest.
	AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest__Type = "AWS::AutoScaling::AutoScalingGroup.AcceleratorTotalMemoryMiBRequest"
)

var (
	// AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest__PropertiesMap reports all the CloudFormation properties for AWS::AutoScaling::AutoScalingGroup.AcceleratorTotalMemoryMiBRequest.
	AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest__PropertiesMap = struct {
		Max string
		Min string
	}{
		Max: "Max",
		Min: "Min",
	}

	// AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest__PropertiesSlice reports all the CloudFormation properties for AWS::AutoScaling::AutoScalingGroup.AcceleratorTotalMemoryMiBRequest.
	AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest__PropertiesSlice = []string{
		AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest__PropertiesMap.Max,
		AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest__PropertiesMap.Min,
	}
)

// AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest is a binding for AWS::AutoScaling::AutoScalingGroup.AcceleratorTotalMemoryMiBRequest.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-acceleratortotalmemorymibrequest.html
type AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest struct {
	// Max is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-acceleratortotalmemorymibrequest.html#cfn-autoscaling-autoscalinggroup-acceleratortotalmemorymibrequest-max
	Max cfz.Expression[int32] `json:"Max,omitempty"`

	// Min is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-acceleratortotalmemorymibrequest.html#cfn-autoscaling-autoscalinggroup-acceleratortotalmemorymibrequest-min
	Min cfz.Expression[int32] `json:"Min,omitempty"`
}

// New__AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest initializes a new AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest.
func New__AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest() AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest {
	return AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest{}
}

// GetType returns the CloudFormation type.
func (AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest) GetType() string {
	return AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest__Type
}

// Set__Max updates property "Max".
func (t AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest) Set__Max(v cfz.Expression[int32]) AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest {
	t.Max = v
	return t
}

// SetV__Max updates property "Max".
func (t AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest) SetV__Max(v int32) AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest {
	t.Max = cfz.V(v)
	return t
}

// Set__Min updates property "Min".
func (t AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest) Set__Min(v cfz.Expression[int32]) AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest {
	t.Min = v
	return t
}

// SetV__Min updates property "Min".
func (t AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest) SetV__Min(v int32) AWS_AutoScaling_AutoScalingGroup_AcceleratorTotalMemoryMiBRequest {
	t.Min = cfz.V(v)
	return t
}
