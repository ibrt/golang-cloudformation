// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_applicationsignals

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ApplicationSignals_ServiceLevelObjective_Goal__Type is the CloudFormation type for AWS::ApplicationSignals::ServiceLevelObjective.Goal.
	AWS_ApplicationSignals_ServiceLevelObjective_Goal__Type = "AWS::ApplicationSignals::ServiceLevelObjective.Goal"
)

var (
	// AWS_ApplicationSignals_ServiceLevelObjective_Goal__PropertiesMap reports all the CloudFormation properties for AWS::ApplicationSignals::ServiceLevelObjective.Goal.
	AWS_ApplicationSignals_ServiceLevelObjective_Goal__PropertiesMap = struct {
		AttainmentGoal   string
		Interval         string
		WarningThreshold string
	}{
		AttainmentGoal:   "AttainmentGoal",
		Interval:         "Interval",
		WarningThreshold: "WarningThreshold",
	}

	// AWS_ApplicationSignals_ServiceLevelObjective_Goal__PropertiesSlice reports all the CloudFormation properties for AWS::ApplicationSignals::ServiceLevelObjective.Goal.
	AWS_ApplicationSignals_ServiceLevelObjective_Goal__PropertiesSlice = []string{
		AWS_ApplicationSignals_ServiceLevelObjective_Goal__PropertiesMap.AttainmentGoal,
		AWS_ApplicationSignals_ServiceLevelObjective_Goal__PropertiesMap.Interval,
		AWS_ApplicationSignals_ServiceLevelObjective_Goal__PropertiesMap.WarningThreshold,
	}
)

// AWS_ApplicationSignals_ServiceLevelObjective_Goal is a binding for AWS::ApplicationSignals::ServiceLevelObjective.Goal.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-goal.html
type AWS_ApplicationSignals_ServiceLevelObjective_Goal struct {
	// AttainmentGoal is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-goal.html#cfn-applicationsignals-servicelevelobjective-goal-attainmentgoal
	AttainmentGoal cfz.Expression[float64] `json:"AttainmentGoal,omitempty"`

	// Interval is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-goal.html#cfn-applicationsignals-servicelevelobjective-goal-interval
	Interval cfz.Expression[AWS_ApplicationSignals_ServiceLevelObjective_Interval] `json:"Interval,omitempty"`

	// WarningThreshold is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-goal.html#cfn-applicationsignals-servicelevelobjective-goal-warningthreshold
	WarningThreshold cfz.Expression[float64] `json:"WarningThreshold,omitempty"`
}

// New__AWS_ApplicationSignals_ServiceLevelObjective_Goal initializes a new AWS_ApplicationSignals_ServiceLevelObjective_Goal.
func New__AWS_ApplicationSignals_ServiceLevelObjective_Goal() AWS_ApplicationSignals_ServiceLevelObjective_Goal {
	return AWS_ApplicationSignals_ServiceLevelObjective_Goal{}
}

// GetType returns the CloudFormation type.
func (AWS_ApplicationSignals_ServiceLevelObjective_Goal) GetType() string {
	return AWS_ApplicationSignals_ServiceLevelObjective_Goal__Type
}

// Set__AttainmentGoal updates property "AttainmentGoal".
func (t AWS_ApplicationSignals_ServiceLevelObjective_Goal) Set__AttainmentGoal(v cfz.Expression[float64]) AWS_ApplicationSignals_ServiceLevelObjective_Goal {
	t.AttainmentGoal = v
	return t
}

// SetV__AttainmentGoal updates property "AttainmentGoal".
func (t AWS_ApplicationSignals_ServiceLevelObjective_Goal) SetV__AttainmentGoal(v float64) AWS_ApplicationSignals_ServiceLevelObjective_Goal {
	t.AttainmentGoal = cfz.V(v)
	return t
}

// Set__Interval updates property "Interval".
func (t AWS_ApplicationSignals_ServiceLevelObjective_Goal) Set__Interval(v cfz.Expression[AWS_ApplicationSignals_ServiceLevelObjective_Interval]) AWS_ApplicationSignals_ServiceLevelObjective_Goal {
	t.Interval = v
	return t
}

// SetV__Interval updates property "Interval".
func (t AWS_ApplicationSignals_ServiceLevelObjective_Goal) SetV__Interval(v AWS_ApplicationSignals_ServiceLevelObjective_Interval) AWS_ApplicationSignals_ServiceLevelObjective_Goal {
	t.Interval = cfz.V(v)
	return t
}

// Set__WarningThreshold updates property "WarningThreshold".
func (t AWS_ApplicationSignals_ServiceLevelObjective_Goal) Set__WarningThreshold(v cfz.Expression[float64]) AWS_ApplicationSignals_ServiceLevelObjective_Goal {
	t.WarningThreshold = v
	return t
}

// SetV__WarningThreshold updates property "WarningThreshold".
func (t AWS_ApplicationSignals_ServiceLevelObjective_Goal) SetV__WarningThreshold(v float64) AWS_ApplicationSignals_ServiceLevelObjective_Goal {
	t.WarningThreshold = cfz.V(v)
	return t
}
