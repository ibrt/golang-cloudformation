// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lookoutmetrics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_LookoutMetrics_AnomalyDetector_MetricSource__Type is the CloudFormation type for AWS::LookoutMetrics::AnomalyDetector.MetricSource.
	AWS_LookoutMetrics_AnomalyDetector_MetricSource__Type = "AWS::LookoutMetrics::AnomalyDetector.MetricSource"
)

var (
	// AWS_LookoutMetrics_AnomalyDetector_MetricSource__PropertiesMap reports all the CloudFormation properties for AWS::LookoutMetrics::AnomalyDetector.MetricSource.
	AWS_LookoutMetrics_AnomalyDetector_MetricSource__PropertiesMap = struct {
		AppFlowConfig        string
		CloudwatchConfig     string
		RDSSourceConfig      string
		RedshiftSourceConfig string
		S3SourceConfig       string
	}{
		AppFlowConfig:        "AppFlowConfig",
		CloudwatchConfig:     "CloudwatchConfig",
		RDSSourceConfig:      "RDSSourceConfig",
		RedshiftSourceConfig: "RedshiftSourceConfig",
		S3SourceConfig:       "S3SourceConfig",
	}

	// AWS_LookoutMetrics_AnomalyDetector_MetricSource__PropertiesSlice reports all the CloudFormation properties for AWS::LookoutMetrics::AnomalyDetector.MetricSource.
	AWS_LookoutMetrics_AnomalyDetector_MetricSource__PropertiesSlice = []string{
		AWS_LookoutMetrics_AnomalyDetector_MetricSource__PropertiesMap.AppFlowConfig,
		AWS_LookoutMetrics_AnomalyDetector_MetricSource__PropertiesMap.CloudwatchConfig,
		AWS_LookoutMetrics_AnomalyDetector_MetricSource__PropertiesMap.RDSSourceConfig,
		AWS_LookoutMetrics_AnomalyDetector_MetricSource__PropertiesMap.RedshiftSourceConfig,
		AWS_LookoutMetrics_AnomalyDetector_MetricSource__PropertiesMap.S3SourceConfig,
	}
)

// AWS_LookoutMetrics_AnomalyDetector_MetricSource is a binding for AWS::LookoutMetrics::AnomalyDetector.MetricSource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricsource.html
type AWS_LookoutMetrics_AnomalyDetector_MetricSource struct {
	// AppFlowConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricsource.html#cfn-lookoutmetrics-anomalydetector-metricsource-appflowconfig
	AppFlowConfig cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_AppFlowConfig] `json:"AppFlowConfig,omitempty"`

	// CloudwatchConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricsource.html#cfn-lookoutmetrics-anomalydetector-metricsource-cloudwatchconfig
	CloudwatchConfig cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_CloudwatchConfig] `json:"CloudwatchConfig,omitempty"`

	// RDSSourceConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricsource.html#cfn-lookoutmetrics-anomalydetector-metricsource-rdssourceconfig
	RDSSourceConfig cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_RDSSourceConfig] `json:"RDSSourceConfig,omitempty"`

	// RedshiftSourceConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricsource.html#cfn-lookoutmetrics-anomalydetector-metricsource-redshiftsourceconfig
	RedshiftSourceConfig cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_RedshiftSourceConfig] `json:"RedshiftSourceConfig,omitempty"`

	// S3SourceConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-metricsource.html#cfn-lookoutmetrics-anomalydetector-metricsource-s3sourceconfig
	S3SourceConfig cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_S3SourceConfig] `json:"S3SourceConfig,omitempty"`
}

// New__AWS_LookoutMetrics_AnomalyDetector_MetricSource initializes a new AWS_LookoutMetrics_AnomalyDetector_MetricSource.
func New__AWS_LookoutMetrics_AnomalyDetector_MetricSource() AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	return AWS_LookoutMetrics_AnomalyDetector_MetricSource{}
}

// GetType returns the CloudFormation type.
func (AWS_LookoutMetrics_AnomalyDetector_MetricSource) GetType() string {
	return AWS_LookoutMetrics_AnomalyDetector_MetricSource__Type
}

// Set__AppFlowConfig updates property "AppFlowConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) Set__AppFlowConfig(v cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_AppFlowConfig]) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.AppFlowConfig = v
	return t
}

// SetV__AppFlowConfig updates property "AppFlowConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) SetV__AppFlowConfig(v AWS_LookoutMetrics_AnomalyDetector_AppFlowConfig) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.AppFlowConfig = cfz.V(v)
	return t
}

// Set__CloudwatchConfig updates property "CloudwatchConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) Set__CloudwatchConfig(v cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_CloudwatchConfig]) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.CloudwatchConfig = v
	return t
}

// SetV__CloudwatchConfig updates property "CloudwatchConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) SetV__CloudwatchConfig(v AWS_LookoutMetrics_AnomalyDetector_CloudwatchConfig) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.CloudwatchConfig = cfz.V(v)
	return t
}

// Set__RDSSourceConfig updates property "RDSSourceConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) Set__RDSSourceConfig(v cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_RDSSourceConfig]) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.RDSSourceConfig = v
	return t
}

// SetV__RDSSourceConfig updates property "RDSSourceConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) SetV__RDSSourceConfig(v AWS_LookoutMetrics_AnomalyDetector_RDSSourceConfig) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.RDSSourceConfig = cfz.V(v)
	return t
}

// Set__RedshiftSourceConfig updates property "RedshiftSourceConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) Set__RedshiftSourceConfig(v cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_RedshiftSourceConfig]) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.RedshiftSourceConfig = v
	return t
}

// SetV__RedshiftSourceConfig updates property "RedshiftSourceConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) SetV__RedshiftSourceConfig(v AWS_LookoutMetrics_AnomalyDetector_RedshiftSourceConfig) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.RedshiftSourceConfig = cfz.V(v)
	return t
}

// Set__S3SourceConfig updates property "S3SourceConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) Set__S3SourceConfig(v cfz.Expression[AWS_LookoutMetrics_AnomalyDetector_S3SourceConfig]) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.S3SourceConfig = v
	return t
}

// SetV__S3SourceConfig updates property "S3SourceConfig".
func (t AWS_LookoutMetrics_AnomalyDetector_MetricSource) SetV__S3SourceConfig(v AWS_LookoutMetrics_AnomalyDetector_S3SourceConfig) AWS_LookoutMetrics_AnomalyDetector_MetricSource {
	t.S3SourceConfig = cfz.V(v)
	return t
}
