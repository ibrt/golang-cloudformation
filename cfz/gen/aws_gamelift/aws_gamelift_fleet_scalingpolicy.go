// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_gamelift

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GameLift_Fleet_ScalingPolicy__Type is the CloudFormation type for AWS::GameLift::Fleet.ScalingPolicy.
	AWS_GameLift_Fleet_ScalingPolicy__Type = "AWS::GameLift::Fleet.ScalingPolicy"
)

var (
	// AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap reports all the CloudFormation properties for AWS::GameLift::Fleet.ScalingPolicy.
	AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap = struct {
		ComparisonOperator    string
		EvaluationPeriods     string
		Location              string
		MetricName            string
		Name                  string
		PolicyType            string
		ScalingAdjustment     string
		ScalingAdjustmentType string
		Status                string
		TargetConfiguration   string
		Threshold             string
		UpdateStatus          string
	}{
		ComparisonOperator:    "ComparisonOperator",
		EvaluationPeriods:     "EvaluationPeriods",
		Location:              "Location",
		MetricName:            "MetricName",
		Name:                  "Name",
		PolicyType:            "PolicyType",
		ScalingAdjustment:     "ScalingAdjustment",
		ScalingAdjustmentType: "ScalingAdjustmentType",
		Status:                "Status",
		TargetConfiguration:   "TargetConfiguration",
		Threshold:             "Threshold",
		UpdateStatus:          "UpdateStatus",
	}

	// AWS_GameLift_Fleet_ScalingPolicy__PropertiesSlice reports all the CloudFormation properties for AWS::GameLift::Fleet.ScalingPolicy.
	AWS_GameLift_Fleet_ScalingPolicy__PropertiesSlice = []string{
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.ComparisonOperator,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.EvaluationPeriods,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.Location,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.MetricName,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.Name,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.PolicyType,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.ScalingAdjustment,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.ScalingAdjustmentType,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.Status,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.TargetConfiguration,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.Threshold,
		AWS_GameLift_Fleet_ScalingPolicy__PropertiesMap.UpdateStatus,
	}
)

// AWS_GameLift_Fleet_ScalingPolicy is a binding for AWS::GameLift::Fleet.ScalingPolicy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html
type AWS_GameLift_Fleet_ScalingPolicy struct {
	// ComparisonOperator is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-comparisonoperator
	ComparisonOperator cfz.Expression[string] `json:"ComparisonOperator,omitempty"`

	// EvaluationPeriods is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-evaluationperiods
	EvaluationPeriods cfz.Expression[int32] `json:"EvaluationPeriods,omitempty"`

	// Location is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-location
	Location cfz.Expression[string] `json:"Location,omitempty"`

	// MetricName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-metricname
	MetricName cfz.Expression[string] `json:"MetricName,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// PolicyType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-policytype
	PolicyType cfz.Expression[string] `json:"PolicyType,omitempty"`

	// ScalingAdjustment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-scalingadjustment
	ScalingAdjustment cfz.Expression[int32] `json:"ScalingAdjustment,omitempty"`

	// ScalingAdjustmentType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-scalingadjustmenttype
	ScalingAdjustmentType cfz.Expression[string] `json:"ScalingAdjustmentType,omitempty"`

	// Status is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-status
	Status cfz.Expression[string] `json:"Status,omitempty"`

	// TargetConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-targetconfiguration
	TargetConfiguration cfz.Expression[AWS_GameLift_Fleet_TargetConfiguration] `json:"TargetConfiguration,omitempty"`

	// Threshold is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-threshold
	Threshold cfz.Expression[float64] `json:"Threshold,omitempty"`

	// UpdateStatus is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-updatestatus
	UpdateStatus cfz.Expression[string] `json:"UpdateStatus,omitempty"`
}

// New__AWS_GameLift_Fleet_ScalingPolicy initializes a new AWS_GameLift_Fleet_ScalingPolicy.
func New__AWS_GameLift_Fleet_ScalingPolicy() AWS_GameLift_Fleet_ScalingPolicy {
	return AWS_GameLift_Fleet_ScalingPolicy{}
}

// GetType returns the CloudFormation type.
func (AWS_GameLift_Fleet_ScalingPolicy) GetType() string {
	return AWS_GameLift_Fleet_ScalingPolicy__Type
}

// Set__ComparisonOperator updates property "ComparisonOperator".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__ComparisonOperator(v cfz.Expression[string]) AWS_GameLift_Fleet_ScalingPolicy {
	t.ComparisonOperator = v
	return t
}

// SetV__ComparisonOperator updates property "ComparisonOperator".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__ComparisonOperator(v string) AWS_GameLift_Fleet_ScalingPolicy {
	t.ComparisonOperator = cfz.V(v)
	return t
}

// Set__EvaluationPeriods updates property "EvaluationPeriods".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__EvaluationPeriods(v cfz.Expression[int32]) AWS_GameLift_Fleet_ScalingPolicy {
	t.EvaluationPeriods = v
	return t
}

// SetV__EvaluationPeriods updates property "EvaluationPeriods".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__EvaluationPeriods(v int32) AWS_GameLift_Fleet_ScalingPolicy {
	t.EvaluationPeriods = cfz.V(v)
	return t
}

// Set__Location updates property "Location".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__Location(v cfz.Expression[string]) AWS_GameLift_Fleet_ScalingPolicy {
	t.Location = v
	return t
}

// SetV__Location updates property "Location".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__Location(v string) AWS_GameLift_Fleet_ScalingPolicy {
	t.Location = cfz.V(v)
	return t
}

// Set__MetricName updates property "MetricName".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__MetricName(v cfz.Expression[string]) AWS_GameLift_Fleet_ScalingPolicy {
	t.MetricName = v
	return t
}

// SetV__MetricName updates property "MetricName".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__MetricName(v string) AWS_GameLift_Fleet_ScalingPolicy {
	t.MetricName = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__Name(v cfz.Expression[string]) AWS_GameLift_Fleet_ScalingPolicy {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__Name(v string) AWS_GameLift_Fleet_ScalingPolicy {
	t.Name = cfz.V(v)
	return t
}

// Set__PolicyType updates property "PolicyType".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__PolicyType(v cfz.Expression[string]) AWS_GameLift_Fleet_ScalingPolicy {
	t.PolicyType = v
	return t
}

// SetV__PolicyType updates property "PolicyType".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__PolicyType(v string) AWS_GameLift_Fleet_ScalingPolicy {
	t.PolicyType = cfz.V(v)
	return t
}

// Set__ScalingAdjustment updates property "ScalingAdjustment".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__ScalingAdjustment(v cfz.Expression[int32]) AWS_GameLift_Fleet_ScalingPolicy {
	t.ScalingAdjustment = v
	return t
}

// SetV__ScalingAdjustment updates property "ScalingAdjustment".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__ScalingAdjustment(v int32) AWS_GameLift_Fleet_ScalingPolicy {
	t.ScalingAdjustment = cfz.V(v)
	return t
}

// Set__ScalingAdjustmentType updates property "ScalingAdjustmentType".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__ScalingAdjustmentType(v cfz.Expression[string]) AWS_GameLift_Fleet_ScalingPolicy {
	t.ScalingAdjustmentType = v
	return t
}

// SetV__ScalingAdjustmentType updates property "ScalingAdjustmentType".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__ScalingAdjustmentType(v string) AWS_GameLift_Fleet_ScalingPolicy {
	t.ScalingAdjustmentType = cfz.V(v)
	return t
}

// Set__Status updates property "Status".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__Status(v cfz.Expression[string]) AWS_GameLift_Fleet_ScalingPolicy {
	t.Status = v
	return t
}

// SetV__Status updates property "Status".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__Status(v string) AWS_GameLift_Fleet_ScalingPolicy {
	t.Status = cfz.V(v)
	return t
}

// Set__TargetConfiguration updates property "TargetConfiguration".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__TargetConfiguration(v cfz.Expression[AWS_GameLift_Fleet_TargetConfiguration]) AWS_GameLift_Fleet_ScalingPolicy {
	t.TargetConfiguration = v
	return t
}

// SetV__TargetConfiguration updates property "TargetConfiguration".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__TargetConfiguration(v AWS_GameLift_Fleet_TargetConfiguration) AWS_GameLift_Fleet_ScalingPolicy {
	t.TargetConfiguration = cfz.V(v)
	return t
}

// Set__Threshold updates property "Threshold".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__Threshold(v cfz.Expression[float64]) AWS_GameLift_Fleet_ScalingPolicy {
	t.Threshold = v
	return t
}

// SetV__Threshold updates property "Threshold".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__Threshold(v float64) AWS_GameLift_Fleet_ScalingPolicy {
	t.Threshold = cfz.V(v)
	return t
}

// Set__UpdateStatus updates property "UpdateStatus".
func (t AWS_GameLift_Fleet_ScalingPolicy) Set__UpdateStatus(v cfz.Expression[string]) AWS_GameLift_Fleet_ScalingPolicy {
	t.UpdateStatus = v
	return t
}

// SetV__UpdateStatus updates property "UpdateStatus".
func (t AWS_GameLift_Fleet_ScalingPolicy) SetV__UpdateStatus(v string) AWS_GameLift_Fleet_ScalingPolicy {
	t.UpdateStatus = cfz.V(v)
	return t
}
