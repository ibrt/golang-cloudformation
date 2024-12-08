// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emr

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__Type is the CloudFormation type for AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition.
	AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__Type = "AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition"
)

var (
	// AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap reports all the CloudFormation properties for AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition.
	AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap = struct {
		ComparisonOperator string
		Dimensions         string
		EvaluationPeriods  string
		MetricName         string
		Namespace          string
		Period             string
		Statistic          string
		Threshold          string
		Unit               string
	}{
		ComparisonOperator: "ComparisonOperator",
		Dimensions:         "Dimensions",
		EvaluationPeriods:  "EvaluationPeriods",
		MetricName:         "MetricName",
		Namespace:          "Namespace",
		Period:             "Period",
		Statistic:          "Statistic",
		Threshold:          "Threshold",
		Unit:               "Unit",
	}

	// AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesSlice reports all the CloudFormation properties for AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition.
	AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesSlice = []string{
		AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap.ComparisonOperator,
		AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap.Dimensions,
		AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap.EvaluationPeriods,
		AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap.MetricName,
		AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap.Namespace,
		AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap.Period,
		AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap.Statistic,
		AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap.Threshold,
		AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__PropertiesMap.Unit,
	}
)

// AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition is a binding for AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html
type AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition struct {
	// ComparisonOperator is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-comparisonoperator
	ComparisonOperator cfz.Expression[string] `json:"ComparisonOperator,omitempty"`

	// Dimensions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-dimensions
	Dimensions cfz.ExpressionSlice[AWS_EMR_InstanceGroupConfig_MetricDimension] `json:"Dimensions,omitempty"`

	// EvaluationPeriods is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-evaluationperiods
	EvaluationPeriods cfz.Expression[int32] `json:"EvaluationPeriods,omitempty"`

	// MetricName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-metricname
	MetricName cfz.Expression[string] `json:"MetricName,omitempty"`

	// Namespace is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-namespace
	Namespace cfz.Expression[string] `json:"Namespace,omitempty"`

	// Period is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-period
	Period cfz.Expression[int32] `json:"Period,omitempty"`

	// Statistic is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-statistic
	Statistic cfz.Expression[string] `json:"Statistic,omitempty"`

	// Threshold is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-threshold
	Threshold cfz.Expression[float64] `json:"Threshold,omitempty"`

	// Unit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-unit
	Unit cfz.Expression[string] `json:"Unit,omitempty"`
}

// New__AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition initializes a new AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition.
func New__AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition() AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	return AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition{}
}

// GetType returns the CloudFormation type.
func (AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) GetType() string {
	return AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition__Type
}

// Set__ComparisonOperator updates property "ComparisonOperator".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) Set__ComparisonOperator(v cfz.Expression[string]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.ComparisonOperator = v
	return t
}

// SetV__ComparisonOperator updates property "ComparisonOperator".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetV__ComparisonOperator(v string) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.ComparisonOperator = cfz.V(v)
	return t
}

// Set__Dimensions updates property "Dimensions".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) Set__Dimensions(v cfz.ExpressionSlice[AWS_EMR_InstanceGroupConfig_MetricDimension]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Dimensions = v
	return t
}

// SetS__Dimensions updates property "Dimensions".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetS__Dimensions(v ...cfz.Expression[AWS_EMR_InstanceGroupConfig_MetricDimension]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Dimensions = cfz.S(v...)
	return t
}

// SetSV__Dimensions updates property "Dimensions".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetSV__Dimensions(v ...AWS_EMR_InstanceGroupConfig_MetricDimension) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Dimensions = cfz.SV(v...)
	return t
}

// Set__EvaluationPeriods updates property "EvaluationPeriods".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) Set__EvaluationPeriods(v cfz.Expression[int32]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.EvaluationPeriods = v
	return t
}

// SetV__EvaluationPeriods updates property "EvaluationPeriods".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetV__EvaluationPeriods(v int32) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.EvaluationPeriods = cfz.V(v)
	return t
}

// Set__MetricName updates property "MetricName".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) Set__MetricName(v cfz.Expression[string]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.MetricName = v
	return t
}

// SetV__MetricName updates property "MetricName".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetV__MetricName(v string) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.MetricName = cfz.V(v)
	return t
}

// Set__Namespace updates property "Namespace".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) Set__Namespace(v cfz.Expression[string]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Namespace = v
	return t
}

// SetV__Namespace updates property "Namespace".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetV__Namespace(v string) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Namespace = cfz.V(v)
	return t
}

// Set__Period updates property "Period".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) Set__Period(v cfz.Expression[int32]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Period = v
	return t
}

// SetV__Period updates property "Period".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetV__Period(v int32) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Period = cfz.V(v)
	return t
}

// Set__Statistic updates property "Statistic".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) Set__Statistic(v cfz.Expression[string]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Statistic = v
	return t
}

// SetV__Statistic updates property "Statistic".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetV__Statistic(v string) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Statistic = cfz.V(v)
	return t
}

// Set__Threshold updates property "Threshold".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) Set__Threshold(v cfz.Expression[float64]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Threshold = v
	return t
}

// SetV__Threshold updates property "Threshold".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetV__Threshold(v float64) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Threshold = cfz.V(v)
	return t
}

// Set__Unit updates property "Unit".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) Set__Unit(v cfz.Expression[string]) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Unit = v
	return t
}

// SetV__Unit updates property "Unit".
func (t AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition) SetV__Unit(v string) AWS_EMR_InstanceGroupConfig_CloudWatchAlarmDefinition {
	t.Unit = cfz.V(v)
	return t
}
