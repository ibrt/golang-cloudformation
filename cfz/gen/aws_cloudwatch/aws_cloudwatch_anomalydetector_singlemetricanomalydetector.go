// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudwatch

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__Type is the CloudFormation type for AWS::CloudWatch::AnomalyDetector.SingleMetricAnomalyDetector.
	AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__Type = "AWS::CloudWatch::AnomalyDetector.SingleMetricAnomalyDetector"
)

var (
	// AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__PropertiesMap reports all the CloudFormation properties for AWS::CloudWatch::AnomalyDetector.SingleMetricAnomalyDetector.
	AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__PropertiesMap = struct {
		AccountId  string
		Dimensions string
		MetricName string
		Namespace  string
		Stat       string
	}{
		AccountId:  "AccountId",
		Dimensions: "Dimensions",
		MetricName: "MetricName",
		Namespace:  "Namespace",
		Stat:       "Stat",
	}

	// AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__PropertiesSlice reports all the CloudFormation properties for AWS::CloudWatch::AnomalyDetector.SingleMetricAnomalyDetector.
	AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__PropertiesSlice = []string{
		AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__PropertiesMap.AccountId,
		AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__PropertiesMap.Dimensions,
		AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__PropertiesMap.MetricName,
		AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__PropertiesMap.Namespace,
		AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__PropertiesMap.Stat,
	}
)

// AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector is a binding for AWS::CloudWatch::AnomalyDetector.SingleMetricAnomalyDetector.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html
type AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector struct {
	// AccountId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector-accountid
	AccountId cfz.Expression[string] `json:"AccountId,omitempty"`

	// Dimensions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector-dimensions
	Dimensions cfz.ExpressionSlice[AWS_CloudWatch_AnomalyDetector_Dimension] `json:"Dimensions,omitempty"`

	// MetricName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector-metricname
	MetricName cfz.Expression[string] `json:"MetricName,omitempty"`

	// Namespace is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector-namespace
	Namespace cfz.Expression[string] `json:"Namespace,omitempty"`

	// Stat is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector-stat
	Stat cfz.Expression[string] `json:"Stat,omitempty"`
}

// New__AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector initializes a new AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector.
func New__AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector() AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	return AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) GetType() string {
	return AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector__Type
}

// Set__AccountId updates property "AccountId".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) Set__AccountId(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.AccountId = v
	return t
}

// SetV__AccountId updates property "AccountId".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) SetV__AccountId(v string) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.AccountId = cfz.V(v)
	return t
}

// Set__Dimensions updates property "Dimensions".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) Set__Dimensions(v cfz.ExpressionSlice[AWS_CloudWatch_AnomalyDetector_Dimension]) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.Dimensions = v
	return t
}

// SetS__Dimensions updates property "Dimensions".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) SetS__Dimensions(v ...cfz.Expression[AWS_CloudWatch_AnomalyDetector_Dimension]) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.Dimensions = cfz.S(v...)
	return t
}

// SetSV__Dimensions updates property "Dimensions".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) SetSV__Dimensions(v ...AWS_CloudWatch_AnomalyDetector_Dimension) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.Dimensions = cfz.SV(v...)
	return t
}

// Set__MetricName updates property "MetricName".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) Set__MetricName(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.MetricName = v
	return t
}

// SetV__MetricName updates property "MetricName".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) SetV__MetricName(v string) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.MetricName = cfz.V(v)
	return t
}

// Set__Namespace updates property "Namespace".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) Set__Namespace(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.Namespace = v
	return t
}

// SetV__Namespace updates property "Namespace".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) SetV__Namespace(v string) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.Namespace = cfz.V(v)
	return t
}

// Set__Stat updates property "Stat".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) Set__Stat(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.Stat = v
	return t
}

// SetV__Stat updates property "Stat".
func (t AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector) SetV__Stat(v string) AWS_CloudWatch_AnomalyDetector_SingleMetricAnomalyDetector {
	t.Stat = cfz.V(v)
	return t
}
