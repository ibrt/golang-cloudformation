// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emr

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EMR_InstanceFleetConfig_SpotResizingSpecification__Type is the CloudFormation type for AWS::EMR::InstanceFleetConfig.SpotResizingSpecification.
	AWS_EMR_InstanceFleetConfig_SpotResizingSpecification__Type = "AWS::EMR::InstanceFleetConfig.SpotResizingSpecification"
)

var (
	// AWS_EMR_InstanceFleetConfig_SpotResizingSpecification__PropertiesMap reports all the CloudFormation properties for AWS::EMR::InstanceFleetConfig.SpotResizingSpecification.
	AWS_EMR_InstanceFleetConfig_SpotResizingSpecification__PropertiesMap = struct {
		AllocationStrategy     string
		TimeoutDurationMinutes string
	}{
		AllocationStrategy:     "AllocationStrategy",
		TimeoutDurationMinutes: "TimeoutDurationMinutes",
	}

	// AWS_EMR_InstanceFleetConfig_SpotResizingSpecification__PropertiesSlice reports all the CloudFormation properties for AWS::EMR::InstanceFleetConfig.SpotResizingSpecification.
	AWS_EMR_InstanceFleetConfig_SpotResizingSpecification__PropertiesSlice = []string{
		AWS_EMR_InstanceFleetConfig_SpotResizingSpecification__PropertiesMap.AllocationStrategy,
		AWS_EMR_InstanceFleetConfig_SpotResizingSpecification__PropertiesMap.TimeoutDurationMinutes,
	}
)

// AWS_EMR_InstanceFleetConfig_SpotResizingSpecification is a binding for AWS::EMR::InstanceFleetConfig.SpotResizingSpecification.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-spotresizingspecification.html
type AWS_EMR_InstanceFleetConfig_SpotResizingSpecification struct {
	// AllocationStrategy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-spotresizingspecification.html#cfn-elasticmapreduce-instancefleetconfig-spotresizingspecification-allocationstrategy
	AllocationStrategy cfz.Expression[string] `json:"AllocationStrategy,omitempty"`

	// TimeoutDurationMinutes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-spotresizingspecification.html#cfn-elasticmapreduce-instancefleetconfig-spotresizingspecification-timeoutdurationminutes
	TimeoutDurationMinutes cfz.Expression[int32] `json:"TimeoutDurationMinutes,omitempty"`
}

// New__AWS_EMR_InstanceFleetConfig_SpotResizingSpecification initializes a new AWS_EMR_InstanceFleetConfig_SpotResizingSpecification.
func New__AWS_EMR_InstanceFleetConfig_SpotResizingSpecification() AWS_EMR_InstanceFleetConfig_SpotResizingSpecification {
	return AWS_EMR_InstanceFleetConfig_SpotResizingSpecification{}
}

// GetType returns the CloudFormation type.
func (AWS_EMR_InstanceFleetConfig_SpotResizingSpecification) GetType() string {
	return AWS_EMR_InstanceFleetConfig_SpotResizingSpecification__Type
}

// Set__AllocationStrategy updates property "AllocationStrategy".
func (t AWS_EMR_InstanceFleetConfig_SpotResizingSpecification) Set__AllocationStrategy(v cfz.Expression[string]) AWS_EMR_InstanceFleetConfig_SpotResizingSpecification {
	t.AllocationStrategy = v
	return t
}

// SetV__AllocationStrategy updates property "AllocationStrategy".
func (t AWS_EMR_InstanceFleetConfig_SpotResizingSpecification) SetV__AllocationStrategy(v string) AWS_EMR_InstanceFleetConfig_SpotResizingSpecification {
	t.AllocationStrategy = cfz.V(v)
	return t
}

// Set__TimeoutDurationMinutes updates property "TimeoutDurationMinutes".
func (t AWS_EMR_InstanceFleetConfig_SpotResizingSpecification) Set__TimeoutDurationMinutes(v cfz.Expression[int32]) AWS_EMR_InstanceFleetConfig_SpotResizingSpecification {
	t.TimeoutDurationMinutes = v
	return t
}

// SetV__TimeoutDurationMinutes updates property "TimeoutDurationMinutes".
func (t AWS_EMR_InstanceFleetConfig_SpotResizingSpecification) SetV__TimeoutDurationMinutes(v int32) AWS_EMR_InstanceFleetConfig_SpotResizingSpecification {
	t.TimeoutDurationMinutes = cfz.V(v)
	return t
}
