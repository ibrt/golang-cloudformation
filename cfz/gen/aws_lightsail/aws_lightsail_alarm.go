// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lightsail

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Lightsail_Alarm)(nil)
	_ cfz.Resource                   = (*AWS_Lightsail_Alarm)(nil)
)

const (
	// AWS_Lightsail_Alarm__Type is the CloudFormation type for AWS::Lightsail::Alarm.
	AWS_Lightsail_Alarm__Type = "AWS::Lightsail::Alarm"
)

var (
	// AWS_Lightsail_Alarm__AttributesMap reports all the CloudFormation attributes for AWS::Lightsail::Alarm.
	AWS_Lightsail_Alarm__AttributesMap = struct {
		AlarmArn string
		State    string
	}{
		AlarmArn: "AlarmArn",
		State:    "State",
	}

	// AWS_Lightsail_Alarm__AttributesSlice reports all the CloudFormation attributes for AWS::Lightsail::Alarm.
	AWS_Lightsail_Alarm__AttributesSlice = []string{
		AWS_Lightsail_Alarm__AttributesMap.AlarmArn,
		AWS_Lightsail_Alarm__AttributesMap.State,
	}
)

var (
	// AWS_Lightsail_Alarm__PropertiesMap reports all the CloudFormation properties for AWS::Lightsail::Alarm.
	AWS_Lightsail_Alarm__PropertiesMap = struct {
		AlarmName             string
		ComparisonOperator    string
		ContactProtocols      string
		DatapointsToAlarm     string
		EvaluationPeriods     string
		MetricName            string
		MonitoredResourceName string
		NotificationEnabled   string
		NotificationTriggers  string
		Threshold             string
		TreatMissingData      string
	}{
		AlarmName:             "AlarmName",
		ComparisonOperator:    "ComparisonOperator",
		ContactProtocols:      "ContactProtocols",
		DatapointsToAlarm:     "DatapointsToAlarm",
		EvaluationPeriods:     "EvaluationPeriods",
		MetricName:            "MetricName",
		MonitoredResourceName: "MonitoredResourceName",
		NotificationEnabled:   "NotificationEnabled",
		NotificationTriggers:  "NotificationTriggers",
		Threshold:             "Threshold",
		TreatMissingData:      "TreatMissingData",
	}

	// AWS_Lightsail_Alarm__PropertiesSlice reports all the CloudFormation properties for AWS::Lightsail::Alarm.
	AWS_Lightsail_Alarm__PropertiesSlice = []string{
		AWS_Lightsail_Alarm__PropertiesMap.AlarmName,
		AWS_Lightsail_Alarm__PropertiesMap.ComparisonOperator,
		AWS_Lightsail_Alarm__PropertiesMap.ContactProtocols,
		AWS_Lightsail_Alarm__PropertiesMap.DatapointsToAlarm,
		AWS_Lightsail_Alarm__PropertiesMap.EvaluationPeriods,
		AWS_Lightsail_Alarm__PropertiesMap.MetricName,
		AWS_Lightsail_Alarm__PropertiesMap.MonitoredResourceName,
		AWS_Lightsail_Alarm__PropertiesMap.NotificationEnabled,
		AWS_Lightsail_Alarm__PropertiesMap.NotificationTriggers,
		AWS_Lightsail_Alarm__PropertiesMap.Threshold,
		AWS_Lightsail_Alarm__PropertiesMap.TreatMissingData,
	}
)

// AWS_Lightsail_Alarm is a binding for AWS::Lightsail::Alarm.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html
type AWS_Lightsail_Alarm struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AlarmName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-alarmname
	AlarmName cfz.Expression[string] `json:"AlarmName,omitempty"`

	// ComparisonOperator is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-comparisonoperator
	ComparisonOperator cfz.Expression[string] `json:"ComparisonOperator,omitempty"`

	// ContactProtocols is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-contactprotocols
	ContactProtocols cfz.ExpressionSlice[string] `json:"ContactProtocols,omitempty"`

	// DatapointsToAlarm is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-datapointstoalarm
	DatapointsToAlarm cfz.Expression[int32] `json:"DatapointsToAlarm,omitempty"`

	// EvaluationPeriods is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-evaluationperiods
	EvaluationPeriods cfz.Expression[int32] `json:"EvaluationPeriods,omitempty"`

	// MetricName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-metricname
	MetricName cfz.Expression[string] `json:"MetricName,omitempty"`

	// MonitoredResourceName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-monitoredresourcename
	MonitoredResourceName cfz.Expression[string] `json:"MonitoredResourceName,omitempty"`

	// NotificationEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-notificationenabled
	NotificationEnabled cfz.Expression[bool] `json:"NotificationEnabled,omitempty"`

	// NotificationTriggers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-notificationtriggers
	NotificationTriggers cfz.ExpressionSlice[string] `json:"NotificationTriggers,omitempty"`

	// Threshold is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-threshold
	Threshold cfz.Expression[float64] `json:"Threshold,omitempty"`

	// TreatMissingData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html#cfn-lightsail-alarm-treatmissingdata
	TreatMissingData cfz.Expression[string] `json:"TreatMissingData,omitempty"`
}

// New__AWS_Lightsail_Alarm initializes a new *AWS_Lightsail_Alarm.
func New__AWS_Lightsail_Alarm(logicalName string) *AWS_Lightsail_Alarm {
	return &AWS_Lightsail_Alarm{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Lightsail_Alarm) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Lightsail_Alarm) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Lightsail_Alarm) GetType() string {
	return AWS_Lightsail_Alarm__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Lightsail_Alarm) Set__LogicalName(v string) *AWS_Lightsail_Alarm {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Lightsail_Alarm) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Lightsail_Alarm {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Lightsail_Alarm) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Lightsail_Alarm {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Lightsail_Alarm) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Lightsail_Alarm {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Lightsail_Alarm) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Lightsail_Alarm {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Lightsail_Alarm) Set__RequestedOutputs(v []cfz.Output) *AWS_Lightsail_Alarm {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Lightsail_Alarm) Add__RequestedOutputs(v ...cfz.Output) *AWS_Lightsail_Alarm {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AlarmName updates property "AlarmName".
func (t *AWS_Lightsail_Alarm) Set__AlarmName(v cfz.Expression[string]) *AWS_Lightsail_Alarm {
	t.AlarmName = v
	return t
}

// SetV__AlarmName updates property "AlarmName".
func (t *AWS_Lightsail_Alarm) SetV__AlarmName(v string) *AWS_Lightsail_Alarm {
	t.AlarmName = cfz.V(v)
	return t
}

// Set__ComparisonOperator updates property "ComparisonOperator".
func (t *AWS_Lightsail_Alarm) Set__ComparisonOperator(v cfz.Expression[string]) *AWS_Lightsail_Alarm {
	t.ComparisonOperator = v
	return t
}

// SetV__ComparisonOperator updates property "ComparisonOperator".
func (t *AWS_Lightsail_Alarm) SetV__ComparisonOperator(v string) *AWS_Lightsail_Alarm {
	t.ComparisonOperator = cfz.V(v)
	return t
}

// Set__ContactProtocols updates property "ContactProtocols".
func (t *AWS_Lightsail_Alarm) Set__ContactProtocols(v cfz.ExpressionSlice[string]) *AWS_Lightsail_Alarm {
	t.ContactProtocols = v
	return t
}

// SetS__ContactProtocols updates property "ContactProtocols".
func (t *AWS_Lightsail_Alarm) SetS__ContactProtocols(v ...cfz.Expression[string]) *AWS_Lightsail_Alarm {
	t.ContactProtocols = cfz.S(v...)
	return t
}

// SetSV__ContactProtocols updates property "ContactProtocols".
func (t *AWS_Lightsail_Alarm) SetSV__ContactProtocols(v ...string) *AWS_Lightsail_Alarm {
	t.ContactProtocols = cfz.SV(v...)
	return t
}

// Set__DatapointsToAlarm updates property "DatapointsToAlarm".
func (t *AWS_Lightsail_Alarm) Set__DatapointsToAlarm(v cfz.Expression[int32]) *AWS_Lightsail_Alarm {
	t.DatapointsToAlarm = v
	return t
}

// SetV__DatapointsToAlarm updates property "DatapointsToAlarm".
func (t *AWS_Lightsail_Alarm) SetV__DatapointsToAlarm(v int32) *AWS_Lightsail_Alarm {
	t.DatapointsToAlarm = cfz.V(v)
	return t
}

// Set__EvaluationPeriods updates property "EvaluationPeriods".
func (t *AWS_Lightsail_Alarm) Set__EvaluationPeriods(v cfz.Expression[int32]) *AWS_Lightsail_Alarm {
	t.EvaluationPeriods = v
	return t
}

// SetV__EvaluationPeriods updates property "EvaluationPeriods".
func (t *AWS_Lightsail_Alarm) SetV__EvaluationPeriods(v int32) *AWS_Lightsail_Alarm {
	t.EvaluationPeriods = cfz.V(v)
	return t
}

// Set__MetricName updates property "MetricName".
func (t *AWS_Lightsail_Alarm) Set__MetricName(v cfz.Expression[string]) *AWS_Lightsail_Alarm {
	t.MetricName = v
	return t
}

// SetV__MetricName updates property "MetricName".
func (t *AWS_Lightsail_Alarm) SetV__MetricName(v string) *AWS_Lightsail_Alarm {
	t.MetricName = cfz.V(v)
	return t
}

// Set__MonitoredResourceName updates property "MonitoredResourceName".
func (t *AWS_Lightsail_Alarm) Set__MonitoredResourceName(v cfz.Expression[string]) *AWS_Lightsail_Alarm {
	t.MonitoredResourceName = v
	return t
}

// SetV__MonitoredResourceName updates property "MonitoredResourceName".
func (t *AWS_Lightsail_Alarm) SetV__MonitoredResourceName(v string) *AWS_Lightsail_Alarm {
	t.MonitoredResourceName = cfz.V(v)
	return t
}

// Set__NotificationEnabled updates property "NotificationEnabled".
func (t *AWS_Lightsail_Alarm) Set__NotificationEnabled(v cfz.Expression[bool]) *AWS_Lightsail_Alarm {
	t.NotificationEnabled = v
	return t
}

// SetV__NotificationEnabled updates property "NotificationEnabled".
func (t *AWS_Lightsail_Alarm) SetV__NotificationEnabled(v bool) *AWS_Lightsail_Alarm {
	t.NotificationEnabled = cfz.V(v)
	return t
}

// Set__NotificationTriggers updates property "NotificationTriggers".
func (t *AWS_Lightsail_Alarm) Set__NotificationTriggers(v cfz.ExpressionSlice[string]) *AWS_Lightsail_Alarm {
	t.NotificationTriggers = v
	return t
}

// SetS__NotificationTriggers updates property "NotificationTriggers".
func (t *AWS_Lightsail_Alarm) SetS__NotificationTriggers(v ...cfz.Expression[string]) *AWS_Lightsail_Alarm {
	t.NotificationTriggers = cfz.S(v...)
	return t
}

// SetSV__NotificationTriggers updates property "NotificationTriggers".
func (t *AWS_Lightsail_Alarm) SetSV__NotificationTriggers(v ...string) *AWS_Lightsail_Alarm {
	t.NotificationTriggers = cfz.SV(v...)
	return t
}

// Set__Threshold updates property "Threshold".
func (t *AWS_Lightsail_Alarm) Set__Threshold(v cfz.Expression[float64]) *AWS_Lightsail_Alarm {
	t.Threshold = v
	return t
}

// SetV__Threshold updates property "Threshold".
func (t *AWS_Lightsail_Alarm) SetV__Threshold(v float64) *AWS_Lightsail_Alarm {
	t.Threshold = cfz.V(v)
	return t
}

// Set__TreatMissingData updates property "TreatMissingData".
func (t *AWS_Lightsail_Alarm) Set__TreatMissingData(v cfz.Expression[string]) *AWS_Lightsail_Alarm {
	t.TreatMissingData = v
	return t
}

// SetV__TreatMissingData updates property "TreatMissingData".
func (t *AWS_Lightsail_Alarm) SetV__TreatMissingData(v string) *AWS_Lightsail_Alarm {
	t.TreatMissingData = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Lightsail_Alarm) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__AlarmArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AlarmArn
func (t *AWS_Lightsail_Alarm) GetAtt__AlarmArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Lightsail_Alarm__AttributesMap.AlarmArn))
}

// GetAtt__State returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: State
func (t *AWS_Lightsail_Alarm) GetAtt__State() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Lightsail_Alarm__AttributesMap.State))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Lightsail_Alarm) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AlarmArn returns a conventionally configured output for an attribute of this resource.
// Attribute: AlarmArn
func (t *AWS_Lightsail_Alarm) GetConventionalOutputAtt__AlarmArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAlarmArn", t.GetAtt__AlarmArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__State returns a conventionally configured output for an attribute of this resource.
// Attribute: State
func (t *AWS_Lightsail_Alarm) GetConventionalOutputAtt__State(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttState", t.GetAtt__State())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Lightsail_Alarm) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Lightsail_Alarm

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

func (t *AWS_Lightsail_Alarm) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
