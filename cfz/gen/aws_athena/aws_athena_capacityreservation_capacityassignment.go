// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_athena

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Athena_CapacityReservation_CapacityAssignment__Type is the CloudFormation type for AWS::Athena::CapacityReservation.CapacityAssignment.
	AWS_Athena_CapacityReservation_CapacityAssignment__Type = "AWS::Athena::CapacityReservation.CapacityAssignment"
)

var (
	// AWS_Athena_CapacityReservation_CapacityAssignment__PropertiesMap reports all the CloudFormation properties for AWS::Athena::CapacityReservation.CapacityAssignment.
	AWS_Athena_CapacityReservation_CapacityAssignment__PropertiesMap = struct {
		WorkgroupNames string
	}{
		WorkgroupNames: "WorkgroupNames",
	}

	// AWS_Athena_CapacityReservation_CapacityAssignment__PropertiesSlice reports all the CloudFormation properties for AWS::Athena::CapacityReservation.CapacityAssignment.
	AWS_Athena_CapacityReservation_CapacityAssignment__PropertiesSlice = []string{
		AWS_Athena_CapacityReservation_CapacityAssignment__PropertiesMap.WorkgroupNames,
	}
)

// AWS_Athena_CapacityReservation_CapacityAssignment is a binding for AWS::Athena::CapacityReservation.CapacityAssignment.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-capacityreservation-capacityassignment.html
type AWS_Athena_CapacityReservation_CapacityAssignment struct {
	// WorkgroupNames is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-capacityreservation-capacityassignment.html#cfn-athena-capacityreservation-capacityassignment-workgroupnames
	WorkgroupNames cfz.ExpressionSlice[string] `json:"WorkgroupNames,omitempty"`
}

// New__AWS_Athena_CapacityReservation_CapacityAssignment initializes a new AWS_Athena_CapacityReservation_CapacityAssignment.
func New__AWS_Athena_CapacityReservation_CapacityAssignment() AWS_Athena_CapacityReservation_CapacityAssignment {
	return AWS_Athena_CapacityReservation_CapacityAssignment{}
}

// GetType returns the CloudFormation type.
func (AWS_Athena_CapacityReservation_CapacityAssignment) GetType() string {
	return AWS_Athena_CapacityReservation_CapacityAssignment__Type
}

// Set__WorkgroupNames updates property "WorkgroupNames".
func (t AWS_Athena_CapacityReservation_CapacityAssignment) Set__WorkgroupNames(v cfz.ExpressionSlice[string]) AWS_Athena_CapacityReservation_CapacityAssignment {
	t.WorkgroupNames = v
	return t
}

// SetS__WorkgroupNames updates property "WorkgroupNames".
func (t AWS_Athena_CapacityReservation_CapacityAssignment) SetS__WorkgroupNames(v ...cfz.Expression[string]) AWS_Athena_CapacityReservation_CapacityAssignment {
	t.WorkgroupNames = cfz.S(v...)
	return t
}

// SetSV__WorkgroupNames updates property "WorkgroupNames".
func (t AWS_Athena_CapacityReservation_CapacityAssignment) SetSV__WorkgroupNames(v ...string) AWS_Athena_CapacityReservation_CapacityAssignment {
	t.WorkgroupNames = cfz.SV(v...)
	return t
}
