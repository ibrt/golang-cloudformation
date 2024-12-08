// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudwatch

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudWatch_AnomalyDetector_MetricStat__Type is the CloudFormation type for AWS::CloudWatch::AnomalyDetector.MetricStat.
	AWS_CloudWatch_AnomalyDetector_MetricStat__Type = "AWS::CloudWatch::AnomalyDetector.MetricStat"
)

var (
	// AWS_CloudWatch_AnomalyDetector_MetricStat__PropertiesMap reports all the CloudFormation properties for AWS::CloudWatch::AnomalyDetector.MetricStat.
	AWS_CloudWatch_AnomalyDetector_MetricStat__PropertiesMap = struct {
		Metric string
		Period string
		Stat   string
		Unit   string
	}{
		Metric: "Metric",
		Period: "Period",
		Stat:   "Stat",
		Unit:   "Unit",
	}

	// AWS_CloudWatch_AnomalyDetector_MetricStat__PropertiesSlice reports all the CloudFormation properties for AWS::CloudWatch::AnomalyDetector.MetricStat.
	AWS_CloudWatch_AnomalyDetector_MetricStat__PropertiesSlice = []string{
		AWS_CloudWatch_AnomalyDetector_MetricStat__PropertiesMap.Metric,
		AWS_CloudWatch_AnomalyDetector_MetricStat__PropertiesMap.Period,
		AWS_CloudWatch_AnomalyDetector_MetricStat__PropertiesMap.Stat,
		AWS_CloudWatch_AnomalyDetector_MetricStat__PropertiesMap.Unit,
	}
)

// AWS_CloudWatch_AnomalyDetector_MetricStat is a binding for AWS::CloudWatch::AnomalyDetector.MetricStat.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html
type AWS_CloudWatch_AnomalyDetector_MetricStat struct {
	// Metric is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html#cfn-cloudwatch-anomalydetector-metricstat-metric
	Metric cfz.Expression[AWS_CloudWatch_AnomalyDetector_Metric] `json:"Metric,omitempty"`

	// Period is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html#cfn-cloudwatch-anomalydetector-metricstat-period
	Period cfz.Expression[int32] `json:"Period,omitempty"`

	// Stat is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html#cfn-cloudwatch-anomalydetector-metricstat-stat
	Stat cfz.Expression[string] `json:"Stat,omitempty"`

	// Unit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html#cfn-cloudwatch-anomalydetector-metricstat-unit
	Unit cfz.Expression[string] `json:"Unit,omitempty"`
}

// New__AWS_CloudWatch_AnomalyDetector_MetricStat initializes a new AWS_CloudWatch_AnomalyDetector_MetricStat.
func New__AWS_CloudWatch_AnomalyDetector_MetricStat() AWS_CloudWatch_AnomalyDetector_MetricStat {
	return AWS_CloudWatch_AnomalyDetector_MetricStat{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudWatch_AnomalyDetector_MetricStat) GetType() string {
	return AWS_CloudWatch_AnomalyDetector_MetricStat__Type
}

// Set__Metric updates property "Metric".
func (t AWS_CloudWatch_AnomalyDetector_MetricStat) Set__Metric(v cfz.Expression[AWS_CloudWatch_AnomalyDetector_Metric]) AWS_CloudWatch_AnomalyDetector_MetricStat {
	t.Metric = v
	return t
}

// SetV__Metric updates property "Metric".
func (t AWS_CloudWatch_AnomalyDetector_MetricStat) SetV__Metric(v AWS_CloudWatch_AnomalyDetector_Metric) AWS_CloudWatch_AnomalyDetector_MetricStat {
	t.Metric = cfz.V(v)
	return t
}

// Set__Period updates property "Period".
func (t AWS_CloudWatch_AnomalyDetector_MetricStat) Set__Period(v cfz.Expression[int32]) AWS_CloudWatch_AnomalyDetector_MetricStat {
	t.Period = v
	return t
}

// SetV__Period updates property "Period".
func (t AWS_CloudWatch_AnomalyDetector_MetricStat) SetV__Period(v int32) AWS_CloudWatch_AnomalyDetector_MetricStat {
	t.Period = cfz.V(v)
	return t
}

// Set__Stat updates property "Stat".
func (t AWS_CloudWatch_AnomalyDetector_MetricStat) Set__Stat(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_MetricStat {
	t.Stat = v
	return t
}

// SetV__Stat updates property "Stat".
func (t AWS_CloudWatch_AnomalyDetector_MetricStat) SetV__Stat(v string) AWS_CloudWatch_AnomalyDetector_MetricStat {
	t.Stat = cfz.V(v)
	return t
}

// Set__Unit updates property "Unit".
func (t AWS_CloudWatch_AnomalyDetector_MetricStat) Set__Unit(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_MetricStat {
	t.Unit = v
	return t
}

// SetV__Unit updates property "Unit".
func (t AWS_CloudWatch_AnomalyDetector_MetricStat) SetV__Unit(v string) AWS_CloudWatch_AnomalyDetector_MetricStat {
	t.Unit = cfz.V(v)
	return t
}
