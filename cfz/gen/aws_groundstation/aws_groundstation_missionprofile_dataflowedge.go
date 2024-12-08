// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_groundstation

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GroundStation_MissionProfile_DataflowEdge__Type is the CloudFormation type for AWS::GroundStation::MissionProfile.DataflowEdge.
	AWS_GroundStation_MissionProfile_DataflowEdge__Type = "AWS::GroundStation::MissionProfile.DataflowEdge"
)

var (
	// AWS_GroundStation_MissionProfile_DataflowEdge__PropertiesMap reports all the CloudFormation properties for AWS::GroundStation::MissionProfile.DataflowEdge.
	AWS_GroundStation_MissionProfile_DataflowEdge__PropertiesMap = struct {
		Destination string
		Source      string
	}{
		Destination: "Destination",
		Source:      "Source",
	}

	// AWS_GroundStation_MissionProfile_DataflowEdge__PropertiesSlice reports all the CloudFormation properties for AWS::GroundStation::MissionProfile.DataflowEdge.
	AWS_GroundStation_MissionProfile_DataflowEdge__PropertiesSlice = []string{
		AWS_GroundStation_MissionProfile_DataflowEdge__PropertiesMap.Destination,
		AWS_GroundStation_MissionProfile_DataflowEdge__PropertiesMap.Source,
	}
)

// AWS_GroundStation_MissionProfile_DataflowEdge is a binding for AWS::GroundStation::MissionProfile.DataflowEdge.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-dataflowedge.html
type AWS_GroundStation_MissionProfile_DataflowEdge struct {
	// Destination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-dataflowedge.html#cfn-groundstation-missionprofile-dataflowedge-destination
	Destination cfz.Expression[string] `json:"Destination,omitempty"`

	// Source is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-dataflowedge.html#cfn-groundstation-missionprofile-dataflowedge-source
	Source cfz.Expression[string] `json:"Source,omitempty"`
}

// New__AWS_GroundStation_MissionProfile_DataflowEdge initializes a new AWS_GroundStation_MissionProfile_DataflowEdge.
func New__AWS_GroundStation_MissionProfile_DataflowEdge() AWS_GroundStation_MissionProfile_DataflowEdge {
	return AWS_GroundStation_MissionProfile_DataflowEdge{}
}

// GetType returns the CloudFormation type.
func (AWS_GroundStation_MissionProfile_DataflowEdge) GetType() string {
	return AWS_GroundStation_MissionProfile_DataflowEdge__Type
}

// Set__Destination updates property "Destination".
func (t AWS_GroundStation_MissionProfile_DataflowEdge) Set__Destination(v cfz.Expression[string]) AWS_GroundStation_MissionProfile_DataflowEdge {
	t.Destination = v
	return t
}

// SetV__Destination updates property "Destination".
func (t AWS_GroundStation_MissionProfile_DataflowEdge) SetV__Destination(v string) AWS_GroundStation_MissionProfile_DataflowEdge {
	t.Destination = cfz.V(v)
	return t
}

// Set__Source updates property "Source".
func (t AWS_GroundStation_MissionProfile_DataflowEdge) Set__Source(v cfz.Expression[string]) AWS_GroundStation_MissionProfile_DataflowEdge {
	t.Source = v
	return t
}

// SetV__Source updates property "Source".
func (t AWS_GroundStation_MissionProfile_DataflowEdge) SetV__Source(v string) AWS_GroundStation_MissionProfile_DataflowEdge {
	t.Source = cfz.V(v)
	return t
}
