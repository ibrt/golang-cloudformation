// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lookoutmetrics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig__Type is the CloudFormation type for AWS::LookoutMetrics::AnomalyDetector.AnomalyDetectorConfig.
	AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig__Type = "AWS::LookoutMetrics::AnomalyDetector.AnomalyDetectorConfig"
)

var (
	// AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig__PropertiesMap reports all the CloudFormation properties for AWS::LookoutMetrics::AnomalyDetector.AnomalyDetectorConfig.
	AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig__PropertiesMap = struct {
		AnomalyDetectorFrequency string
	}{
		AnomalyDetectorFrequency: "AnomalyDetectorFrequency",
	}

	// AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig__PropertiesSlice reports all the CloudFormation properties for AWS::LookoutMetrics::AnomalyDetector.AnomalyDetectorConfig.
	AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig__PropertiesSlice = []string{
		AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig__PropertiesMap.AnomalyDetectorFrequency,
	}
)

// AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig is a binding for AWS::LookoutMetrics::AnomalyDetector.AnomalyDetectorConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-anomalydetectorconfig.html
type AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig struct {
	// AnomalyDetectorFrequency is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-anomalydetectorconfig.html#cfn-lookoutmetrics-anomalydetector-anomalydetectorconfig-anomalydetectorfrequency
	AnomalyDetectorFrequency cfz.Expression[string] `json:"AnomalyDetectorFrequency,omitempty"`
}

// New__AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig initializes a new AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig.
func New__AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig() AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig {
	return AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig) GetType() string {
	return AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig__Type
}

// Set__AnomalyDetectorFrequency updates property "AnomalyDetectorFrequency".
func (t AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig) Set__AnomalyDetectorFrequency(v cfz.Expression[string]) AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig {
	t.AnomalyDetectorFrequency = v
	return t
}

// SetV__AnomalyDetectorFrequency updates property "AnomalyDetectorFrequency".
func (t AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig) SetV__AnomalyDetectorFrequency(v string) AWS_LookoutMetrics_AnomalyDetector_AnomalyDetectorConfig {
	t.AnomalyDetectorFrequency = cfz.V(v)
	return t
}
