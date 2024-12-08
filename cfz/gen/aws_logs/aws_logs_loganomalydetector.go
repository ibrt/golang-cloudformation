// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_logs

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Logs_LogAnomalyDetector)(nil)
	_ cfz.Resource                   = (*AWS_Logs_LogAnomalyDetector)(nil)
)

const (
	// AWS_Logs_LogAnomalyDetector__Type is the CloudFormation type for AWS::Logs::LogAnomalyDetector.
	AWS_Logs_LogAnomalyDetector__Type = "AWS::Logs::LogAnomalyDetector"
)

var (
	// AWS_Logs_LogAnomalyDetector__AttributesMap reports all the CloudFormation attributes for AWS::Logs::LogAnomalyDetector.
	AWS_Logs_LogAnomalyDetector__AttributesMap = struct {
		AnomalyDetectorArn    string
		AnomalyDetectorStatus string
		CreationTimeStamp     string
		LastModifiedTimeStamp string
	}{
		AnomalyDetectorArn:    "AnomalyDetectorArn",
		AnomalyDetectorStatus: "AnomalyDetectorStatus",
		CreationTimeStamp:     "CreationTimeStamp",
		LastModifiedTimeStamp: "LastModifiedTimeStamp",
	}

	// AWS_Logs_LogAnomalyDetector__AttributesSlice reports all the CloudFormation attributes for AWS::Logs::LogAnomalyDetector.
	AWS_Logs_LogAnomalyDetector__AttributesSlice = []string{
		AWS_Logs_LogAnomalyDetector__AttributesMap.AnomalyDetectorArn,
		AWS_Logs_LogAnomalyDetector__AttributesMap.AnomalyDetectorStatus,
		AWS_Logs_LogAnomalyDetector__AttributesMap.CreationTimeStamp,
		AWS_Logs_LogAnomalyDetector__AttributesMap.LastModifiedTimeStamp,
	}
)

var (
	// AWS_Logs_LogAnomalyDetector__PropertiesMap reports all the CloudFormation properties for AWS::Logs::LogAnomalyDetector.
	AWS_Logs_LogAnomalyDetector__PropertiesMap = struct {
		AccountId             string
		AnomalyVisibilityTime string
		DetectorName          string
		EvaluationFrequency   string
		FilterPattern         string
		KmsKeyId              string
		LogGroupArnList       string
	}{
		AccountId:             "AccountId",
		AnomalyVisibilityTime: "AnomalyVisibilityTime",
		DetectorName:          "DetectorName",
		EvaluationFrequency:   "EvaluationFrequency",
		FilterPattern:         "FilterPattern",
		KmsKeyId:              "KmsKeyId",
		LogGroupArnList:       "LogGroupArnList",
	}

	// AWS_Logs_LogAnomalyDetector__PropertiesSlice reports all the CloudFormation properties for AWS::Logs::LogAnomalyDetector.
	AWS_Logs_LogAnomalyDetector__PropertiesSlice = []string{
		AWS_Logs_LogAnomalyDetector__PropertiesMap.AccountId,
		AWS_Logs_LogAnomalyDetector__PropertiesMap.AnomalyVisibilityTime,
		AWS_Logs_LogAnomalyDetector__PropertiesMap.DetectorName,
		AWS_Logs_LogAnomalyDetector__PropertiesMap.EvaluationFrequency,
		AWS_Logs_LogAnomalyDetector__PropertiesMap.FilterPattern,
		AWS_Logs_LogAnomalyDetector__PropertiesMap.KmsKeyId,
		AWS_Logs_LogAnomalyDetector__PropertiesMap.LogGroupArnList,
	}
)

// AWS_Logs_LogAnomalyDetector is a binding for AWS::Logs::LogAnomalyDetector.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html
type AWS_Logs_LogAnomalyDetector struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AccountId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-accountid
	AccountId cfz.Expression[string] `json:"AccountId,omitempty"`

	// AnomalyVisibilityTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-anomalyvisibilitytime
	AnomalyVisibilityTime cfz.Expression[float64] `json:"AnomalyVisibilityTime,omitempty"`

	// DetectorName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-detectorname
	DetectorName cfz.Expression[string] `json:"DetectorName,omitempty"`

	// EvaluationFrequency is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-evaluationfrequency
	EvaluationFrequency cfz.Expression[string] `json:"EvaluationFrequency,omitempty"`

	// FilterPattern is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-filterpattern
	FilterPattern cfz.Expression[string] `json:"FilterPattern,omitempty"`

	// KmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-kmskeyid
	KmsKeyId cfz.Expression[string] `json:"KmsKeyId,omitempty"`

	// LogGroupArnList is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-loggrouparnlist
	LogGroupArnList cfz.ExpressionSlice[string] `json:"LogGroupArnList,omitempty"`
}

// New__AWS_Logs_LogAnomalyDetector initializes a new *AWS_Logs_LogAnomalyDetector.
func New__AWS_Logs_LogAnomalyDetector(logicalName string) *AWS_Logs_LogAnomalyDetector {
	return &AWS_Logs_LogAnomalyDetector{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Logs_LogAnomalyDetector) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Logs_LogAnomalyDetector) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Logs_LogAnomalyDetector) GetType() string {
	return AWS_Logs_LogAnomalyDetector__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Logs_LogAnomalyDetector) Set__LogicalName(v string) *AWS_Logs_LogAnomalyDetector {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Logs_LogAnomalyDetector) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Logs_LogAnomalyDetector {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Logs_LogAnomalyDetector) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Logs_LogAnomalyDetector {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Logs_LogAnomalyDetector) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Logs_LogAnomalyDetector {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Logs_LogAnomalyDetector) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Logs_LogAnomalyDetector {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Logs_LogAnomalyDetector) Set__RequestedOutputs(v []cfz.Output) *AWS_Logs_LogAnomalyDetector {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Logs_LogAnomalyDetector) Add__RequestedOutputs(v ...cfz.Output) *AWS_Logs_LogAnomalyDetector {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AccountId updates property "AccountId".
func (t *AWS_Logs_LogAnomalyDetector) Set__AccountId(v cfz.Expression[string]) *AWS_Logs_LogAnomalyDetector {
	t.AccountId = v
	return t
}

// SetV__AccountId updates property "AccountId".
func (t *AWS_Logs_LogAnomalyDetector) SetV__AccountId(v string) *AWS_Logs_LogAnomalyDetector {
	t.AccountId = cfz.V(v)
	return t
}

// Set__AnomalyVisibilityTime updates property "AnomalyVisibilityTime".
func (t *AWS_Logs_LogAnomalyDetector) Set__AnomalyVisibilityTime(v cfz.Expression[float64]) *AWS_Logs_LogAnomalyDetector {
	t.AnomalyVisibilityTime = v
	return t
}

// SetV__AnomalyVisibilityTime updates property "AnomalyVisibilityTime".
func (t *AWS_Logs_LogAnomalyDetector) SetV__AnomalyVisibilityTime(v float64) *AWS_Logs_LogAnomalyDetector {
	t.AnomalyVisibilityTime = cfz.V(v)
	return t
}

// Set__DetectorName updates property "DetectorName".
func (t *AWS_Logs_LogAnomalyDetector) Set__DetectorName(v cfz.Expression[string]) *AWS_Logs_LogAnomalyDetector {
	t.DetectorName = v
	return t
}

// SetV__DetectorName updates property "DetectorName".
func (t *AWS_Logs_LogAnomalyDetector) SetV__DetectorName(v string) *AWS_Logs_LogAnomalyDetector {
	t.DetectorName = cfz.V(v)
	return t
}

// Set__EvaluationFrequency updates property "EvaluationFrequency".
func (t *AWS_Logs_LogAnomalyDetector) Set__EvaluationFrequency(v cfz.Expression[string]) *AWS_Logs_LogAnomalyDetector {
	t.EvaluationFrequency = v
	return t
}

// SetV__EvaluationFrequency updates property "EvaluationFrequency".
func (t *AWS_Logs_LogAnomalyDetector) SetV__EvaluationFrequency(v string) *AWS_Logs_LogAnomalyDetector {
	t.EvaluationFrequency = cfz.V(v)
	return t
}

// Set__FilterPattern updates property "FilterPattern".
func (t *AWS_Logs_LogAnomalyDetector) Set__FilterPattern(v cfz.Expression[string]) *AWS_Logs_LogAnomalyDetector {
	t.FilterPattern = v
	return t
}

// SetV__FilterPattern updates property "FilterPattern".
func (t *AWS_Logs_LogAnomalyDetector) SetV__FilterPattern(v string) *AWS_Logs_LogAnomalyDetector {
	t.FilterPattern = cfz.V(v)
	return t
}

// Set__KmsKeyId updates property "KmsKeyId".
func (t *AWS_Logs_LogAnomalyDetector) Set__KmsKeyId(v cfz.Expression[string]) *AWS_Logs_LogAnomalyDetector {
	t.KmsKeyId = v
	return t
}

// SetV__KmsKeyId updates property "KmsKeyId".
func (t *AWS_Logs_LogAnomalyDetector) SetV__KmsKeyId(v string) *AWS_Logs_LogAnomalyDetector {
	t.KmsKeyId = cfz.V(v)
	return t
}

// Set__LogGroupArnList updates property "LogGroupArnList".
func (t *AWS_Logs_LogAnomalyDetector) Set__LogGroupArnList(v cfz.ExpressionSlice[string]) *AWS_Logs_LogAnomalyDetector {
	t.LogGroupArnList = v
	return t
}

// SetS__LogGroupArnList updates property "LogGroupArnList".
func (t *AWS_Logs_LogAnomalyDetector) SetS__LogGroupArnList(v ...cfz.Expression[string]) *AWS_Logs_LogAnomalyDetector {
	t.LogGroupArnList = cfz.S(v...)
	return t
}

// SetSV__LogGroupArnList updates property "LogGroupArnList".
func (t *AWS_Logs_LogAnomalyDetector) SetSV__LogGroupArnList(v ...string) *AWS_Logs_LogAnomalyDetector {
	t.LogGroupArnList = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Logs_LogAnomalyDetector) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__AnomalyDetectorArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AnomalyDetectorArn
func (t *AWS_Logs_LogAnomalyDetector) GetAtt__AnomalyDetectorArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Logs_LogAnomalyDetector__AttributesMap.AnomalyDetectorArn))
}

// GetAtt__AnomalyDetectorStatus returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AnomalyDetectorStatus
func (t *AWS_Logs_LogAnomalyDetector) GetAtt__AnomalyDetectorStatus() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Logs_LogAnomalyDetector__AttributesMap.AnomalyDetectorStatus))
}

// GetAtt__CreationTimeStamp returns a $cfz.Expression[float64] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreationTimeStamp
func (t *AWS_Logs_LogAnomalyDetector) GetAtt__CreationTimeStamp() cfz.Expression[float64] {
	return cfz.GetAtt[float64](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Logs_LogAnomalyDetector__AttributesMap.CreationTimeStamp))
}

// GetAtt__LastModifiedTimeStamp returns a $cfz.Expression[float64] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LastModifiedTimeStamp
func (t *AWS_Logs_LogAnomalyDetector) GetAtt__LastModifiedTimeStamp() cfz.Expression[float64] {
	return cfz.GetAtt[float64](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Logs_LogAnomalyDetector__AttributesMap.LastModifiedTimeStamp))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Logs_LogAnomalyDetector) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AnomalyDetectorArn returns a conventionally configured output for an attribute of this resource.
// Attribute: AnomalyDetectorArn
func (t *AWS_Logs_LogAnomalyDetector) GetConventionalOutputAtt__AnomalyDetectorArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAnomalyDetectorArn", t.GetAtt__AnomalyDetectorArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AnomalyDetectorStatus returns a conventionally configured output for an attribute of this resource.
// Attribute: AnomalyDetectorStatus
func (t *AWS_Logs_LogAnomalyDetector) GetConventionalOutputAtt__AnomalyDetectorStatus(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAnomalyDetectorStatus", t.GetAtt__AnomalyDetectorStatus())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreationTimeStamp returns a conventionally configured output for an attribute of this resource.
// Attribute: CreationTimeStamp
func (t *AWS_Logs_LogAnomalyDetector) GetConventionalOutputAtt__CreationTimeStamp(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreationTimeStamp", t.GetAtt__CreationTimeStamp())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LastModifiedTimeStamp returns a conventionally configured output for an attribute of this resource.
// Attribute: LastModifiedTimeStamp
func (t *AWS_Logs_LogAnomalyDetector) GetConventionalOutputAtt__LastModifiedTimeStamp(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLastModifiedTimeStamp", t.GetAtt__LastModifiedTimeStamp())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Logs_LogAnomalyDetector) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Logs_LogAnomalyDetector

	return json.Marshal(struct {
		Type                string                          `json:"Type"`
		DependsOn           []string                        `json:"DependsOn,omitempty"`
		DeletionPolicy      cfz.ResourceDeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Properties          *Properties                     `json:"Properties,omitempty"`
	}{
		Type:       t.GetType(),
		DependsOn:  t.getDependsOn(),
		Properties: (*Properties)(t),
	})
}

func (t *AWS_Logs_LogAnomalyDetector) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
