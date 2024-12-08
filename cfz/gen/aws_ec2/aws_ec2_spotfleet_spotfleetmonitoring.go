// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_SpotFleet_SpotFleetMonitoring__Type is the CloudFormation type for AWS::EC2::SpotFleet.SpotFleetMonitoring.
	AWS_EC2_SpotFleet_SpotFleetMonitoring__Type = "AWS::EC2::SpotFleet.SpotFleetMonitoring"
)

var (
	// AWS_EC2_SpotFleet_SpotFleetMonitoring__PropertiesMap reports all the CloudFormation properties for AWS::EC2::SpotFleet.SpotFleetMonitoring.
	AWS_EC2_SpotFleet_SpotFleetMonitoring__PropertiesMap = struct {
		Enabled string
	}{
		Enabled: "Enabled",
	}

	// AWS_EC2_SpotFleet_SpotFleetMonitoring__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::SpotFleet.SpotFleetMonitoring.
	AWS_EC2_SpotFleet_SpotFleetMonitoring__PropertiesSlice = []string{
		AWS_EC2_SpotFleet_SpotFleetMonitoring__PropertiesMap.Enabled,
	}
)

// AWS_EC2_SpotFleet_SpotFleetMonitoring is a binding for AWS::EC2::SpotFleet.SpotFleetMonitoring.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetmonitoring.html
type AWS_EC2_SpotFleet_SpotFleetMonitoring struct {
	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetmonitoring.html#cfn-ec2-spotfleet-spotfleetmonitoring-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`
}

// New__AWS_EC2_SpotFleet_SpotFleetMonitoring initializes a new AWS_EC2_SpotFleet_SpotFleetMonitoring.
func New__AWS_EC2_SpotFleet_SpotFleetMonitoring() AWS_EC2_SpotFleet_SpotFleetMonitoring {
	return AWS_EC2_SpotFleet_SpotFleetMonitoring{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_SpotFleet_SpotFleetMonitoring) GetType() string {
	return AWS_EC2_SpotFleet_SpotFleetMonitoring__Type
}

// Set__Enabled updates property "Enabled".
func (t AWS_EC2_SpotFleet_SpotFleetMonitoring) Set__Enabled(v cfz.Expression[bool]) AWS_EC2_SpotFleet_SpotFleetMonitoring {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_EC2_SpotFleet_SpotFleetMonitoring) SetV__Enabled(v bool) AWS_EC2_SpotFleet_SpotFleetMonitoring {
	t.Enabled = cfz.V(v)
	return t
}
