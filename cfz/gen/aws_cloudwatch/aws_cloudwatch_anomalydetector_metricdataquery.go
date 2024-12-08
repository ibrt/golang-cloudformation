// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudwatch

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudWatch_AnomalyDetector_MetricDataQuery__Type is the CloudFormation type for AWS::CloudWatch::AnomalyDetector.MetricDataQuery.
	AWS_CloudWatch_AnomalyDetector_MetricDataQuery__Type = "AWS::CloudWatch::AnomalyDetector.MetricDataQuery"
)

var (
	// AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesMap reports all the CloudFormation properties for AWS::CloudWatch::AnomalyDetector.MetricDataQuery.
	AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesMap = struct {
		AccountId  string
		Expression string
		Id         string
		Label      string
		MetricStat string
		Period     string
		ReturnData string
	}{
		AccountId:  "AccountId",
		Expression: "Expression",
		Id:         "Id",
		Label:      "Label",
		MetricStat: "MetricStat",
		Period:     "Period",
		ReturnData: "ReturnData",
	}

	// AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesSlice reports all the CloudFormation properties for AWS::CloudWatch::AnomalyDetector.MetricDataQuery.
	AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesSlice = []string{
		AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesMap.AccountId,
		AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesMap.Expression,
		AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesMap.Id,
		AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesMap.Label,
		AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesMap.MetricStat,
		AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesMap.Period,
		AWS_CloudWatch_AnomalyDetector_MetricDataQuery__PropertiesMap.ReturnData,
	}
)

// AWS_CloudWatch_AnomalyDetector_MetricDataQuery is a binding for AWS::CloudWatch::AnomalyDetector.MetricDataQuery.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html
type AWS_CloudWatch_AnomalyDetector_MetricDataQuery struct {
	// AccountId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-accountid
	AccountId cfz.Expression[string] `json:"AccountId,omitempty"`

	// Expression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-expression
	Expression cfz.Expression[string] `json:"Expression,omitempty"`

	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-id
	Id cfz.Expression[string] `json:"Id,omitempty"`

	// Label is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-label
	Label cfz.Expression[string] `json:"Label,omitempty"`

	// MetricStat is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-metricstat
	MetricStat cfz.Expression[AWS_CloudWatch_AnomalyDetector_MetricStat] `json:"MetricStat,omitempty"`

	// Period is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-period
	Period cfz.Expression[int32] `json:"Period,omitempty"`

	// ReturnData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-returndata
	ReturnData cfz.Expression[bool] `json:"ReturnData,omitempty"`
}

// New__AWS_CloudWatch_AnomalyDetector_MetricDataQuery initializes a new AWS_CloudWatch_AnomalyDetector_MetricDataQuery.
func New__AWS_CloudWatch_AnomalyDetector_MetricDataQuery() AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	return AWS_CloudWatch_AnomalyDetector_MetricDataQuery{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudWatch_AnomalyDetector_MetricDataQuery) GetType() string {
	return AWS_CloudWatch_AnomalyDetector_MetricDataQuery__Type
}

// Set__AccountId updates property "AccountId".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) Set__AccountId(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.AccountId = v
	return t
}

// SetV__AccountId updates property "AccountId".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) SetV__AccountId(v string) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.AccountId = cfz.V(v)
	return t
}

// Set__Expression updates property "Expression".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) Set__Expression(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.Expression = v
	return t
}

// SetV__Expression updates property "Expression".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) SetV__Expression(v string) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.Expression = cfz.V(v)
	return t
}

// Set__Id updates property "Id".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) Set__Id(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) SetV__Id(v string) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.Id = cfz.V(v)
	return t
}

// Set__Label updates property "Label".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) Set__Label(v cfz.Expression[string]) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.Label = v
	return t
}

// SetV__Label updates property "Label".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) SetV__Label(v string) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.Label = cfz.V(v)
	return t
}

// Set__MetricStat updates property "MetricStat".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) Set__MetricStat(v cfz.Expression[AWS_CloudWatch_AnomalyDetector_MetricStat]) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.MetricStat = v
	return t
}

// SetV__MetricStat updates property "MetricStat".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) SetV__MetricStat(v AWS_CloudWatch_AnomalyDetector_MetricStat) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.MetricStat = cfz.V(v)
	return t
}

// Set__Period updates property "Period".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) Set__Period(v cfz.Expression[int32]) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.Period = v
	return t
}

// SetV__Period updates property "Period".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) SetV__Period(v int32) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.Period = cfz.V(v)
	return t
}

// Set__ReturnData updates property "ReturnData".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) Set__ReturnData(v cfz.Expression[bool]) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.ReturnData = v
	return t
}

// SetV__ReturnData updates property "ReturnData".
func (t AWS_CloudWatch_AnomalyDetector_MetricDataQuery) SetV__ReturnData(v bool) AWS_CloudWatch_AnomalyDetector_MetricDataQuery {
	t.ReturnData = cfz.V(v)
	return t
}
