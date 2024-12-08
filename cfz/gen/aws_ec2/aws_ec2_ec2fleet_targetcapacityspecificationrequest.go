// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__Type is the CloudFormation type for AWS::EC2::EC2Fleet.TargetCapacitySpecificationRequest.
	AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__Type = "AWS::EC2::EC2Fleet.TargetCapacitySpecificationRequest"
)

var (
	// AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__PropertiesMap reports all the CloudFormation properties for AWS::EC2::EC2Fleet.TargetCapacitySpecificationRequest.
	AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__PropertiesMap = struct {
		DefaultTargetCapacityType string
		OnDemandTargetCapacity    string
		SpotTargetCapacity        string
		TargetCapacityUnitType    string
		TotalTargetCapacity       string
	}{
		DefaultTargetCapacityType: "DefaultTargetCapacityType",
		OnDemandTargetCapacity:    "OnDemandTargetCapacity",
		SpotTargetCapacity:        "SpotTargetCapacity",
		TargetCapacityUnitType:    "TargetCapacityUnitType",
		TotalTargetCapacity:       "TotalTargetCapacity",
	}

	// AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::EC2Fleet.TargetCapacitySpecificationRequest.
	AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__PropertiesSlice = []string{
		AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__PropertiesMap.DefaultTargetCapacityType,
		AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__PropertiesMap.OnDemandTargetCapacity,
		AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__PropertiesMap.SpotTargetCapacity,
		AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__PropertiesMap.TargetCapacityUnitType,
		AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__PropertiesMap.TotalTargetCapacity,
	}
)

// AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest is a binding for AWS::EC2::EC2Fleet.TargetCapacitySpecificationRequest.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html
type AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest struct {
	// DefaultTargetCapacityType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-defaulttargetcapacitytype
	DefaultTargetCapacityType cfz.Expression[string] `json:"DefaultTargetCapacityType,omitempty"`

	// OnDemandTargetCapacity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-ondemandtargetcapacity
	OnDemandTargetCapacity cfz.Expression[int32] `json:"OnDemandTargetCapacity,omitempty"`

	// SpotTargetCapacity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-spottargetcapacity
	SpotTargetCapacity cfz.Expression[int32] `json:"SpotTargetCapacity,omitempty"`

	// TargetCapacityUnitType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-targetcapacityunittype
	TargetCapacityUnitType cfz.Expression[string] `json:"TargetCapacityUnitType,omitempty"`

	// TotalTargetCapacity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-totaltargetcapacity
	TotalTargetCapacity cfz.Expression[int32] `json:"TotalTargetCapacity,omitempty"`
}

// New__AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest initializes a new AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest.
func New__AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest() AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	return AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) GetType() string {
	return AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest__Type
}

// Set__DefaultTargetCapacityType updates property "DefaultTargetCapacityType".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) Set__DefaultTargetCapacityType(v cfz.Expression[string]) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.DefaultTargetCapacityType = v
	return t
}

// SetV__DefaultTargetCapacityType updates property "DefaultTargetCapacityType".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) SetV__DefaultTargetCapacityType(v string) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.DefaultTargetCapacityType = cfz.V(v)
	return t
}

// Set__OnDemandTargetCapacity updates property "OnDemandTargetCapacity".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) Set__OnDemandTargetCapacity(v cfz.Expression[int32]) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.OnDemandTargetCapacity = v
	return t
}

// SetV__OnDemandTargetCapacity updates property "OnDemandTargetCapacity".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) SetV__OnDemandTargetCapacity(v int32) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.OnDemandTargetCapacity = cfz.V(v)
	return t
}

// Set__SpotTargetCapacity updates property "SpotTargetCapacity".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) Set__SpotTargetCapacity(v cfz.Expression[int32]) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.SpotTargetCapacity = v
	return t
}

// SetV__SpotTargetCapacity updates property "SpotTargetCapacity".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) SetV__SpotTargetCapacity(v int32) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.SpotTargetCapacity = cfz.V(v)
	return t
}

// Set__TargetCapacityUnitType updates property "TargetCapacityUnitType".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) Set__TargetCapacityUnitType(v cfz.Expression[string]) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.TargetCapacityUnitType = v
	return t
}

// SetV__TargetCapacityUnitType updates property "TargetCapacityUnitType".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) SetV__TargetCapacityUnitType(v string) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.TargetCapacityUnitType = cfz.V(v)
	return t
}

// Set__TotalTargetCapacity updates property "TotalTargetCapacity".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) Set__TotalTargetCapacity(v cfz.Expression[int32]) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.TotalTargetCapacity = v
	return t
}

// SetV__TotalTargetCapacity updates property "TotalTargetCapacity".
func (t AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest) SetV__TotalTargetCapacity(v int32) AWS_EC2_EC2Fleet_TargetCapacitySpecificationRequest {
	t.TotalTargetCapacity = cfz.V(v)
	return t
}
