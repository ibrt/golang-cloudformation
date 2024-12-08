// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_gamelift

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GameLift_ContainerFleet_ScalingPolicy__Type is the CloudFormation type for AWS::GameLift::ContainerFleet.ScalingPolicy.
	AWS_GameLift_ContainerFleet_ScalingPolicy__Type = "AWS::GameLift::ContainerFleet.ScalingPolicy"
)

var (
	// AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap reports all the CloudFormation properties for AWS::GameLift::ContainerFleet.ScalingPolicy.
	AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap = struct {
		ComparisonOperator    string
		EvaluationPeriods     string
		MetricName            string
		Name                  string
		PolicyType            string
		ScalingAdjustment     string
		ScalingAdjustmentType string
		TargetConfiguration   string
		Threshold             string
	}{
		ComparisonOperator:    "ComparisonOperator",
		EvaluationPeriods:     "EvaluationPeriods",
		MetricName:            "MetricName",
		Name:                  "Name",
		PolicyType:            "PolicyType",
		ScalingAdjustment:     "ScalingAdjustment",
		ScalingAdjustmentType: "ScalingAdjustmentType",
		TargetConfiguration:   "TargetConfiguration",
		Threshold:             "Threshold",
	}

	// AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesSlice reports all the CloudFormation properties for AWS::GameLift::ContainerFleet.ScalingPolicy.
	AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesSlice = []string{
		AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap.ComparisonOperator,
		AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap.EvaluationPeriods,
		AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap.MetricName,
		AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap.Name,
		AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap.PolicyType,
		AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap.ScalingAdjustment,
		AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap.ScalingAdjustmentType,
		AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap.TargetConfiguration,
		AWS_GameLift_ContainerFleet_ScalingPolicy__PropertiesMap.Threshold,
	}
)

// AWS_GameLift_ContainerFleet_ScalingPolicy is a binding for AWS::GameLift::ContainerFleet.ScalingPolicy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html
type AWS_GameLift_ContainerFleet_ScalingPolicy struct {
	// ComparisonOperator is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html#cfn-gamelift-containerfleet-scalingpolicy-comparisonoperator
	ComparisonOperator cfz.Expression[string] `json:"ComparisonOperator,omitempty"`

	// EvaluationPeriods is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html#cfn-gamelift-containerfleet-scalingpolicy-evaluationperiods
	EvaluationPeriods cfz.Expression[int32] `json:"EvaluationPeriods,omitempty"`

	// MetricName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html#cfn-gamelift-containerfleet-scalingpolicy-metricname
	MetricName cfz.Expression[string] `json:"MetricName,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html#cfn-gamelift-containerfleet-scalingpolicy-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// PolicyType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html#cfn-gamelift-containerfleet-scalingpolicy-policytype
	PolicyType cfz.Expression[string] `json:"PolicyType,omitempty"`

	// ScalingAdjustment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html#cfn-gamelift-containerfleet-scalingpolicy-scalingadjustment
	ScalingAdjustment cfz.Expression[int32] `json:"ScalingAdjustment,omitempty"`

	// ScalingAdjustmentType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html#cfn-gamelift-containerfleet-scalingpolicy-scalingadjustmenttype
	ScalingAdjustmentType cfz.Expression[string] `json:"ScalingAdjustmentType,omitempty"`

	// TargetConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html#cfn-gamelift-containerfleet-scalingpolicy-targetconfiguration
	TargetConfiguration cfz.Expression[AWS_GameLift_ContainerFleet_TargetConfiguration] `json:"TargetConfiguration,omitempty"`

	// Threshold is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-scalingpolicy.html#cfn-gamelift-containerfleet-scalingpolicy-threshold
	Threshold cfz.Expression[float64] `json:"Threshold,omitempty"`
}

// New__AWS_GameLift_ContainerFleet_ScalingPolicy initializes a new AWS_GameLift_ContainerFleet_ScalingPolicy.
func New__AWS_GameLift_ContainerFleet_ScalingPolicy() AWS_GameLift_ContainerFleet_ScalingPolicy {
	return AWS_GameLift_ContainerFleet_ScalingPolicy{}
}

// GetType returns the CloudFormation type.
func (AWS_GameLift_ContainerFleet_ScalingPolicy) GetType() string {
	return AWS_GameLift_ContainerFleet_ScalingPolicy__Type
}

// Set__ComparisonOperator updates property "ComparisonOperator".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) Set__ComparisonOperator(v cfz.Expression[string]) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.ComparisonOperator = v
	return t
}

// SetV__ComparisonOperator updates property "ComparisonOperator".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) SetV__ComparisonOperator(v string) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.ComparisonOperator = cfz.V(v)
	return t
}

// Set__EvaluationPeriods updates property "EvaluationPeriods".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) Set__EvaluationPeriods(v cfz.Expression[int32]) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.EvaluationPeriods = v
	return t
}

// SetV__EvaluationPeriods updates property "EvaluationPeriods".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) SetV__EvaluationPeriods(v int32) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.EvaluationPeriods = cfz.V(v)
	return t
}

// Set__MetricName updates property "MetricName".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) Set__MetricName(v cfz.Expression[string]) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.MetricName = v
	return t
}

// SetV__MetricName updates property "MetricName".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) SetV__MetricName(v string) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.MetricName = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) Set__Name(v cfz.Expression[string]) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) SetV__Name(v string) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.Name = cfz.V(v)
	return t
}

// Set__PolicyType updates property "PolicyType".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) Set__PolicyType(v cfz.Expression[string]) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.PolicyType = v
	return t
}

// SetV__PolicyType updates property "PolicyType".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) SetV__PolicyType(v string) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.PolicyType = cfz.V(v)
	return t
}

// Set__ScalingAdjustment updates property "ScalingAdjustment".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) Set__ScalingAdjustment(v cfz.Expression[int32]) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.ScalingAdjustment = v
	return t
}

// SetV__ScalingAdjustment updates property "ScalingAdjustment".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) SetV__ScalingAdjustment(v int32) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.ScalingAdjustment = cfz.V(v)
	return t
}

// Set__ScalingAdjustmentType updates property "ScalingAdjustmentType".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) Set__ScalingAdjustmentType(v cfz.Expression[string]) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.ScalingAdjustmentType = v
	return t
}

// SetV__ScalingAdjustmentType updates property "ScalingAdjustmentType".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) SetV__ScalingAdjustmentType(v string) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.ScalingAdjustmentType = cfz.V(v)
	return t
}

// Set__TargetConfiguration updates property "TargetConfiguration".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) Set__TargetConfiguration(v cfz.Expression[AWS_GameLift_ContainerFleet_TargetConfiguration]) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.TargetConfiguration = v
	return t
}

// SetV__TargetConfiguration updates property "TargetConfiguration".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) SetV__TargetConfiguration(v AWS_GameLift_ContainerFleet_TargetConfiguration) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.TargetConfiguration = cfz.V(v)
	return t
}

// Set__Threshold updates property "Threshold".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) Set__Threshold(v cfz.Expression[float64]) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.Threshold = v
	return t
}

// SetV__Threshold updates property "Threshold".
func (t AWS_GameLift_ContainerFleet_ScalingPolicy) SetV__Threshold(v float64) AWS_GameLift_ContainerFleet_ScalingPolicy {
	t.Threshold = cfz.V(v)
	return t
}
