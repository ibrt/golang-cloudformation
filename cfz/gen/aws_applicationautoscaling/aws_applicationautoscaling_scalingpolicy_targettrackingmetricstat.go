// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_applicationautoscaling

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__Type is the CloudFormation type for AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingMetricStat.
	AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__Type = "AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingMetricStat"
)

var (
	// AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__PropertiesMap reports all the CloudFormation properties for AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingMetricStat.
	AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__PropertiesMap = struct {
		Metric string
		Stat   string
		Unit   string
	}{
		Metric: "Metric",
		Stat:   "Stat",
		Unit:   "Unit",
	}

	// AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__PropertiesSlice reports all the CloudFormation properties for AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingMetricStat.
	AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__PropertiesSlice = []string{
		AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__PropertiesMap.Metric,
		AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__PropertiesMap.Stat,
		AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__PropertiesMap.Unit,
	}
)

// AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat is a binding for AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingMetricStat.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricstat.html
type AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat struct {
	// Metric is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-targettrackingmetricstat-metric
	Metric cfz.Expression[AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetric] `json:"Metric,omitempty"`

	// Stat is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-targettrackingmetricstat-stat
	Stat cfz.Expression[string] `json:"Stat,omitempty"`

	// Unit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-targettrackingmetricstat-unit
	Unit cfz.Expression[string] `json:"Unit,omitempty"`
}

// New__AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat initializes a new AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat.
func New__AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat() AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat {
	return AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat{}
}

// GetType returns the CloudFormation type.
func (AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat) GetType() string {
	return AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat__Type
}

// Set__Metric updates property "Metric".
func (t AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat) Set__Metric(v cfz.Expression[AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetric]) AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat {
	t.Metric = v
	return t
}

// SetV__Metric updates property "Metric".
func (t AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat) SetV__Metric(v AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetric) AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat {
	t.Metric = cfz.V(v)
	return t
}

// Set__Stat updates property "Stat".
func (t AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat) Set__Stat(v cfz.Expression[string]) AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat {
	t.Stat = v
	return t
}

// SetV__Stat updates property "Stat".
func (t AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat) SetV__Stat(v string) AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat {
	t.Stat = cfz.V(v)
	return t
}

// Set__Unit updates property "Unit".
func (t AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat) Set__Unit(v cfz.Expression[string]) AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat {
	t.Unit = v
	return t
}

// SetV__Unit updates property "Unit".
func (t AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat) SetV__Unit(v string) AWS_ApplicationAutoScaling_ScalingPolicy_TargetTrackingMetricStat {
	t.Unit = cfz.V(v)
	return t
}
