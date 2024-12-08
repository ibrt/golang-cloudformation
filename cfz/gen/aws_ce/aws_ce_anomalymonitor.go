// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ce

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_CE_AnomalyMonitor)(nil)
	_ cfz.Resource                   = (*AWS_CE_AnomalyMonitor)(nil)
)

const (
	// AWS_CE_AnomalyMonitor__Type is the CloudFormation type for AWS::CE::AnomalyMonitor.
	AWS_CE_AnomalyMonitor__Type = "AWS::CE::AnomalyMonitor"
)

var (
	// AWS_CE_AnomalyMonitor__AttributesMap reports all the CloudFormation attributes for AWS::CE::AnomalyMonitor.
	AWS_CE_AnomalyMonitor__AttributesMap = struct {
		CreationDate          string
		DimensionalValueCount string
		LastEvaluatedDate     string
		LastUpdatedDate       string
		MonitorArn            string
	}{
		CreationDate:          "CreationDate",
		DimensionalValueCount: "DimensionalValueCount",
		LastEvaluatedDate:     "LastEvaluatedDate",
		LastUpdatedDate:       "LastUpdatedDate",
		MonitorArn:            "MonitorArn",
	}

	// AWS_CE_AnomalyMonitor__AttributesSlice reports all the CloudFormation attributes for AWS::CE::AnomalyMonitor.
	AWS_CE_AnomalyMonitor__AttributesSlice = []string{
		AWS_CE_AnomalyMonitor__AttributesMap.CreationDate,
		AWS_CE_AnomalyMonitor__AttributesMap.DimensionalValueCount,
		AWS_CE_AnomalyMonitor__AttributesMap.LastEvaluatedDate,
		AWS_CE_AnomalyMonitor__AttributesMap.LastUpdatedDate,
		AWS_CE_AnomalyMonitor__AttributesMap.MonitorArn,
	}
)

var (
	// AWS_CE_AnomalyMonitor__PropertiesMap reports all the CloudFormation properties for AWS::CE::AnomalyMonitor.
	AWS_CE_AnomalyMonitor__PropertiesMap = struct {
		MonitorDimension     string
		MonitorName          string
		MonitorSpecification string
		MonitorType          string
		ResourceTags         string
	}{
		MonitorDimension:     "MonitorDimension",
		MonitorName:          "MonitorName",
		MonitorSpecification: "MonitorSpecification",
		MonitorType:          "MonitorType",
		ResourceTags:         "ResourceTags",
	}

	// AWS_CE_AnomalyMonitor__PropertiesSlice reports all the CloudFormation properties for AWS::CE::AnomalyMonitor.
	AWS_CE_AnomalyMonitor__PropertiesSlice = []string{
		AWS_CE_AnomalyMonitor__PropertiesMap.MonitorDimension,
		AWS_CE_AnomalyMonitor__PropertiesMap.MonitorName,
		AWS_CE_AnomalyMonitor__PropertiesMap.MonitorSpecification,
		AWS_CE_AnomalyMonitor__PropertiesMap.MonitorType,
		AWS_CE_AnomalyMonitor__PropertiesMap.ResourceTags,
	}
)

// AWS_CE_AnomalyMonitor is a binding for AWS::CE::AnomalyMonitor.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html
type AWS_CE_AnomalyMonitor struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// MonitorDimension is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitordimension
	MonitorDimension cfz.Expression[string] `json:"MonitorDimension,omitempty"`

	// MonitorName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitorname
	MonitorName cfz.Expression[string] `json:"MonitorName,omitempty"`

	// MonitorSpecification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitorspecification
	MonitorSpecification cfz.Expression[string] `json:"MonitorSpecification,omitempty"`

	// MonitorType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitortype
	MonitorType cfz.Expression[string] `json:"MonitorType,omitempty"`

	// ResourceTags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-resourcetags
	ResourceTags cfz.ExpressionSlice[AWS_CE_AnomalyMonitor_ResourceTag] `json:"ResourceTags,omitempty"`
}

// New__AWS_CE_AnomalyMonitor initializes a new *AWS_CE_AnomalyMonitor.
func New__AWS_CE_AnomalyMonitor(logicalName string) *AWS_CE_AnomalyMonitor {
	return &AWS_CE_AnomalyMonitor{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_CE_AnomalyMonitor) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_CE_AnomalyMonitor) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_CE_AnomalyMonitor) GetType() string {
	return AWS_CE_AnomalyMonitor__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_CE_AnomalyMonitor) Set__LogicalName(v string) *AWS_CE_AnomalyMonitor {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_CE_AnomalyMonitor) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_CE_AnomalyMonitor {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_CE_AnomalyMonitor) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_CE_AnomalyMonitor {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_CE_AnomalyMonitor) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_CE_AnomalyMonitor {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_CE_AnomalyMonitor) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_CE_AnomalyMonitor {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_CE_AnomalyMonitor) Set__RequestedOutputs(v []cfz.Output) *AWS_CE_AnomalyMonitor {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_CE_AnomalyMonitor) Add__RequestedOutputs(v ...cfz.Output) *AWS_CE_AnomalyMonitor {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__MonitorDimension updates property "MonitorDimension".
func (t *AWS_CE_AnomalyMonitor) Set__MonitorDimension(v cfz.Expression[string]) *AWS_CE_AnomalyMonitor {
	t.MonitorDimension = v
	return t
}

// SetV__MonitorDimension updates property "MonitorDimension".
func (t *AWS_CE_AnomalyMonitor) SetV__MonitorDimension(v string) *AWS_CE_AnomalyMonitor {
	t.MonitorDimension = cfz.V(v)
	return t
}

// Set__MonitorName updates property "MonitorName".
func (t *AWS_CE_AnomalyMonitor) Set__MonitorName(v cfz.Expression[string]) *AWS_CE_AnomalyMonitor {
	t.MonitorName = v
	return t
}

// SetV__MonitorName updates property "MonitorName".
func (t *AWS_CE_AnomalyMonitor) SetV__MonitorName(v string) *AWS_CE_AnomalyMonitor {
	t.MonitorName = cfz.V(v)
	return t
}

// Set__MonitorSpecification updates property "MonitorSpecification".
func (t *AWS_CE_AnomalyMonitor) Set__MonitorSpecification(v cfz.Expression[string]) *AWS_CE_AnomalyMonitor {
	t.MonitorSpecification = v
	return t
}

// SetV__MonitorSpecification updates property "MonitorSpecification".
func (t *AWS_CE_AnomalyMonitor) SetV__MonitorSpecification(v string) *AWS_CE_AnomalyMonitor {
	t.MonitorSpecification = cfz.V(v)
	return t
}

// Set__MonitorType updates property "MonitorType".
func (t *AWS_CE_AnomalyMonitor) Set__MonitorType(v cfz.Expression[string]) *AWS_CE_AnomalyMonitor {
	t.MonitorType = v
	return t
}

// SetV__MonitorType updates property "MonitorType".
func (t *AWS_CE_AnomalyMonitor) SetV__MonitorType(v string) *AWS_CE_AnomalyMonitor {
	t.MonitorType = cfz.V(v)
	return t
}

// Set__ResourceTags updates property "ResourceTags".
func (t *AWS_CE_AnomalyMonitor) Set__ResourceTags(v cfz.ExpressionSlice[AWS_CE_AnomalyMonitor_ResourceTag]) *AWS_CE_AnomalyMonitor {
	t.ResourceTags = v
	return t
}

// SetS__ResourceTags updates property "ResourceTags".
func (t *AWS_CE_AnomalyMonitor) SetS__ResourceTags(v ...cfz.Expression[AWS_CE_AnomalyMonitor_ResourceTag]) *AWS_CE_AnomalyMonitor {
	t.ResourceTags = cfz.S(v...)
	return t
}

// SetSV__ResourceTags updates property "ResourceTags".
func (t *AWS_CE_AnomalyMonitor) SetSV__ResourceTags(v ...AWS_CE_AnomalyMonitor_ResourceTag) *AWS_CE_AnomalyMonitor {
	t.ResourceTags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_CE_AnomalyMonitor) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CreationDate returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreationDate
func (t *AWS_CE_AnomalyMonitor) GetAtt__CreationDate() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CE_AnomalyMonitor__AttributesMap.CreationDate))
}

// GetAtt__DimensionalValueCount returns a $cfz.Expression[int32] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DimensionalValueCount
func (t *AWS_CE_AnomalyMonitor) GetAtt__DimensionalValueCount() cfz.Expression[int32] {
	return cfz.GetAtt[int32](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CE_AnomalyMonitor__AttributesMap.DimensionalValueCount))
}

// GetAtt__LastEvaluatedDate returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LastEvaluatedDate
func (t *AWS_CE_AnomalyMonitor) GetAtt__LastEvaluatedDate() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CE_AnomalyMonitor__AttributesMap.LastEvaluatedDate))
}

// GetAtt__LastUpdatedDate returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LastUpdatedDate
func (t *AWS_CE_AnomalyMonitor) GetAtt__LastUpdatedDate() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CE_AnomalyMonitor__AttributesMap.LastUpdatedDate))
}

// GetAtt__MonitorArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: MonitorArn
func (t *AWS_CE_AnomalyMonitor) GetAtt__MonitorArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CE_AnomalyMonitor__AttributesMap.MonitorArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_CE_AnomalyMonitor) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreationDate returns a conventionally configured output for an attribute of this resource.
// Attribute: CreationDate
func (t *AWS_CE_AnomalyMonitor) GetConventionalOutputAtt__CreationDate(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreationDate", t.GetAtt__CreationDate())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DimensionalValueCount returns a conventionally configured output for an attribute of this resource.
// Attribute: DimensionalValueCount
func (t *AWS_CE_AnomalyMonitor) GetConventionalOutputAtt__DimensionalValueCount(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDimensionalValueCount", t.GetAtt__DimensionalValueCount())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LastEvaluatedDate returns a conventionally configured output for an attribute of this resource.
// Attribute: LastEvaluatedDate
func (t *AWS_CE_AnomalyMonitor) GetConventionalOutputAtt__LastEvaluatedDate(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLastEvaluatedDate", t.GetAtt__LastEvaluatedDate())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LastUpdatedDate returns a conventionally configured output for an attribute of this resource.
// Attribute: LastUpdatedDate
func (t *AWS_CE_AnomalyMonitor) GetConventionalOutputAtt__LastUpdatedDate(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLastUpdatedDate", t.GetAtt__LastUpdatedDate())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__MonitorArn returns a conventionally configured output for an attribute of this resource.
// Attribute: MonitorArn
func (t *AWS_CE_AnomalyMonitor) GetConventionalOutputAtt__MonitorArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttMonitorArn", t.GetAtt__MonitorArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_CE_AnomalyMonitor) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_CE_AnomalyMonitor

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

func (t *AWS_CE_AnomalyMonitor) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
