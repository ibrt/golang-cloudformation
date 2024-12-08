// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_SageMaker_FeatureGroup)(nil)
	_ cfz.Resource                   = (*AWS_SageMaker_FeatureGroup)(nil)
)

const (
	// AWS_SageMaker_FeatureGroup__Type is the CloudFormation type for AWS::SageMaker::FeatureGroup.
	AWS_SageMaker_FeatureGroup__Type = "AWS::SageMaker::FeatureGroup"
)

var (
	// AWS_SageMaker_FeatureGroup__AttributesMap reports all the CloudFormation attributes for AWS::SageMaker::FeatureGroup.
	AWS_SageMaker_FeatureGroup__AttributesMap = struct {
		CreationTime       string
		FeatureGroupStatus string
	}{
		CreationTime:       "CreationTime",
		FeatureGroupStatus: "FeatureGroupStatus",
	}

	// AWS_SageMaker_FeatureGroup__AttributesSlice reports all the CloudFormation attributes for AWS::SageMaker::FeatureGroup.
	AWS_SageMaker_FeatureGroup__AttributesSlice = []string{
		AWS_SageMaker_FeatureGroup__AttributesMap.CreationTime,
		AWS_SageMaker_FeatureGroup__AttributesMap.FeatureGroupStatus,
	}
)

var (
	// AWS_SageMaker_FeatureGroup__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::FeatureGroup.
	AWS_SageMaker_FeatureGroup__PropertiesMap = struct {
		Description                 string
		EventTimeFeatureName        string
		FeatureDefinitions          string
		FeatureGroupName            string
		OfflineStoreConfig          string
		OnlineStoreConfig           string
		RecordIdentifierFeatureName string
		RoleArn                     string
		Tags                        string
		ThroughputConfig            string
	}{
		Description:                 "Description",
		EventTimeFeatureName:        "EventTimeFeatureName",
		FeatureDefinitions:          "FeatureDefinitions",
		FeatureGroupName:            "FeatureGroupName",
		OfflineStoreConfig:          "OfflineStoreConfig",
		OnlineStoreConfig:           "OnlineStoreConfig",
		RecordIdentifierFeatureName: "RecordIdentifierFeatureName",
		RoleArn:                     "RoleArn",
		Tags:                        "Tags",
		ThroughputConfig:            "ThroughputConfig",
	}

	// AWS_SageMaker_FeatureGroup__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::FeatureGroup.
	AWS_SageMaker_FeatureGroup__PropertiesSlice = []string{
		AWS_SageMaker_FeatureGroup__PropertiesMap.Description,
		AWS_SageMaker_FeatureGroup__PropertiesMap.EventTimeFeatureName,
		AWS_SageMaker_FeatureGroup__PropertiesMap.FeatureDefinitions,
		AWS_SageMaker_FeatureGroup__PropertiesMap.FeatureGroupName,
		AWS_SageMaker_FeatureGroup__PropertiesMap.OfflineStoreConfig,
		AWS_SageMaker_FeatureGroup__PropertiesMap.OnlineStoreConfig,
		AWS_SageMaker_FeatureGroup__PropertiesMap.RecordIdentifierFeatureName,
		AWS_SageMaker_FeatureGroup__PropertiesMap.RoleArn,
		AWS_SageMaker_FeatureGroup__PropertiesMap.Tags,
		AWS_SageMaker_FeatureGroup__PropertiesMap.ThroughputConfig,
	}
)

// AWS_SageMaker_FeatureGroup is a binding for AWS::SageMaker::FeatureGroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html
type AWS_SageMaker_FeatureGroup struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// EventTimeFeatureName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-eventtimefeaturename
	EventTimeFeatureName cfz.Expression[string] `json:"EventTimeFeatureName,omitempty"`

	// FeatureDefinitions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-featuredefinitions
	FeatureDefinitions cfz.ExpressionSlice[AWS_SageMaker_FeatureGroup_FeatureDefinition] `json:"FeatureDefinitions,omitempty"`

	// FeatureGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-featuregroupname
	FeatureGroupName cfz.Expression[string] `json:"FeatureGroupName,omitempty"`

	// OfflineStoreConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-offlinestoreconfig
	OfflineStoreConfig cfz.Expression[AWS_SageMaker_FeatureGroup_OfflineStoreConfig] `json:"OfflineStoreConfig,omitempty"`

	// OnlineStoreConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-onlinestoreconfig
	OnlineStoreConfig cfz.Expression[AWS_SageMaker_FeatureGroup_OnlineStoreConfig] `json:"OnlineStoreConfig,omitempty"`

	// RecordIdentifierFeatureName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-recordidentifierfeaturename
	RecordIdentifierFeatureName cfz.Expression[string] `json:"RecordIdentifierFeatureName,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// ThroughputConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-throughputconfig
	ThroughputConfig cfz.Expression[AWS_SageMaker_FeatureGroup_ThroughputConfig] `json:"ThroughputConfig,omitempty"`
}

// New__AWS_SageMaker_FeatureGroup initializes a new *AWS_SageMaker_FeatureGroup.
func New__AWS_SageMaker_FeatureGroup(logicalName string) *AWS_SageMaker_FeatureGroup {
	return &AWS_SageMaker_FeatureGroup{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_SageMaker_FeatureGroup) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_SageMaker_FeatureGroup) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_SageMaker_FeatureGroup) GetType() string {
	return AWS_SageMaker_FeatureGroup__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_SageMaker_FeatureGroup) Set__LogicalName(v string) *AWS_SageMaker_FeatureGroup {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_SageMaker_FeatureGroup) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_SageMaker_FeatureGroup {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_SageMaker_FeatureGroup) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_SageMaker_FeatureGroup {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_SageMaker_FeatureGroup) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_SageMaker_FeatureGroup {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_SageMaker_FeatureGroup) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_SageMaker_FeatureGroup {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_SageMaker_FeatureGroup) Set__RequestedOutputs(v []cfz.Output) *AWS_SageMaker_FeatureGroup {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_SageMaker_FeatureGroup) Add__RequestedOutputs(v ...cfz.Output) *AWS_SageMaker_FeatureGroup {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_SageMaker_FeatureGroup) Set__Description(v cfz.Expression[string]) *AWS_SageMaker_FeatureGroup {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_SageMaker_FeatureGroup) SetV__Description(v string) *AWS_SageMaker_FeatureGroup {
	t.Description = cfz.V(v)
	return t
}

// Set__EventTimeFeatureName updates property "EventTimeFeatureName".
func (t *AWS_SageMaker_FeatureGroup) Set__EventTimeFeatureName(v cfz.Expression[string]) *AWS_SageMaker_FeatureGroup {
	t.EventTimeFeatureName = v
	return t
}

// SetV__EventTimeFeatureName updates property "EventTimeFeatureName".
func (t *AWS_SageMaker_FeatureGroup) SetV__EventTimeFeatureName(v string) *AWS_SageMaker_FeatureGroup {
	t.EventTimeFeatureName = cfz.V(v)
	return t
}

// Set__FeatureDefinitions updates property "FeatureDefinitions".
func (t *AWS_SageMaker_FeatureGroup) Set__FeatureDefinitions(v cfz.ExpressionSlice[AWS_SageMaker_FeatureGroup_FeatureDefinition]) *AWS_SageMaker_FeatureGroup {
	t.FeatureDefinitions = v
	return t
}

// SetS__FeatureDefinitions updates property "FeatureDefinitions".
func (t *AWS_SageMaker_FeatureGroup) SetS__FeatureDefinitions(v ...cfz.Expression[AWS_SageMaker_FeatureGroup_FeatureDefinition]) *AWS_SageMaker_FeatureGroup {
	t.FeatureDefinitions = cfz.S(v...)
	return t
}

// SetSV__FeatureDefinitions updates property "FeatureDefinitions".
func (t *AWS_SageMaker_FeatureGroup) SetSV__FeatureDefinitions(v ...AWS_SageMaker_FeatureGroup_FeatureDefinition) *AWS_SageMaker_FeatureGroup {
	t.FeatureDefinitions = cfz.SV(v...)
	return t
}

// Set__FeatureGroupName updates property "FeatureGroupName".
func (t *AWS_SageMaker_FeatureGroup) Set__FeatureGroupName(v cfz.Expression[string]) *AWS_SageMaker_FeatureGroup {
	t.FeatureGroupName = v
	return t
}

// SetV__FeatureGroupName updates property "FeatureGroupName".
func (t *AWS_SageMaker_FeatureGroup) SetV__FeatureGroupName(v string) *AWS_SageMaker_FeatureGroup {
	t.FeatureGroupName = cfz.V(v)
	return t
}

// Set__OfflineStoreConfig updates property "OfflineStoreConfig".
func (t *AWS_SageMaker_FeatureGroup) Set__OfflineStoreConfig(v cfz.Expression[AWS_SageMaker_FeatureGroup_OfflineStoreConfig]) *AWS_SageMaker_FeatureGroup {
	t.OfflineStoreConfig = v
	return t
}

// SetV__OfflineStoreConfig updates property "OfflineStoreConfig".
func (t *AWS_SageMaker_FeatureGroup) SetV__OfflineStoreConfig(v AWS_SageMaker_FeatureGroup_OfflineStoreConfig) *AWS_SageMaker_FeatureGroup {
	t.OfflineStoreConfig = cfz.V(v)
	return t
}

// Set__OnlineStoreConfig updates property "OnlineStoreConfig".
func (t *AWS_SageMaker_FeatureGroup) Set__OnlineStoreConfig(v cfz.Expression[AWS_SageMaker_FeatureGroup_OnlineStoreConfig]) *AWS_SageMaker_FeatureGroup {
	t.OnlineStoreConfig = v
	return t
}

// SetV__OnlineStoreConfig updates property "OnlineStoreConfig".
func (t *AWS_SageMaker_FeatureGroup) SetV__OnlineStoreConfig(v AWS_SageMaker_FeatureGroup_OnlineStoreConfig) *AWS_SageMaker_FeatureGroup {
	t.OnlineStoreConfig = cfz.V(v)
	return t
}

// Set__RecordIdentifierFeatureName updates property "RecordIdentifierFeatureName".
func (t *AWS_SageMaker_FeatureGroup) Set__RecordIdentifierFeatureName(v cfz.Expression[string]) *AWS_SageMaker_FeatureGroup {
	t.RecordIdentifierFeatureName = v
	return t
}

// SetV__RecordIdentifierFeatureName updates property "RecordIdentifierFeatureName".
func (t *AWS_SageMaker_FeatureGroup) SetV__RecordIdentifierFeatureName(v string) *AWS_SageMaker_FeatureGroup {
	t.RecordIdentifierFeatureName = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t *AWS_SageMaker_FeatureGroup) Set__RoleArn(v cfz.Expression[string]) *AWS_SageMaker_FeatureGroup {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t *AWS_SageMaker_FeatureGroup) SetV__RoleArn(v string) *AWS_SageMaker_FeatureGroup {
	t.RoleArn = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_SageMaker_FeatureGroup) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_SageMaker_FeatureGroup {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_SageMaker_FeatureGroup) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_SageMaker_FeatureGroup {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_SageMaker_FeatureGroup) SetSV__Tags(v ...cfz.Tag) *AWS_SageMaker_FeatureGroup {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__ThroughputConfig updates property "ThroughputConfig".
func (t *AWS_SageMaker_FeatureGroup) Set__ThroughputConfig(v cfz.Expression[AWS_SageMaker_FeatureGroup_ThroughputConfig]) *AWS_SageMaker_FeatureGroup {
	t.ThroughputConfig = v
	return t
}

// SetV__ThroughputConfig updates property "ThroughputConfig".
func (t *AWS_SageMaker_FeatureGroup) SetV__ThroughputConfig(v AWS_SageMaker_FeatureGroup_ThroughputConfig) *AWS_SageMaker_FeatureGroup {
	t.ThroughputConfig = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_SageMaker_FeatureGroup) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CreationTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreationTime
func (t *AWS_SageMaker_FeatureGroup) GetAtt__CreationTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SageMaker_FeatureGroup__AttributesMap.CreationTime))
}

// GetAtt__FeatureGroupStatus returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: FeatureGroupStatus
func (t *AWS_SageMaker_FeatureGroup) GetAtt__FeatureGroupStatus() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SageMaker_FeatureGroup__AttributesMap.FeatureGroupStatus))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_SageMaker_FeatureGroup) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreationTime returns a conventionally configured output for an attribute of this resource.
// Attribute: CreationTime
func (t *AWS_SageMaker_FeatureGroup) GetConventionalOutputAtt__CreationTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreationTime", t.GetAtt__CreationTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__FeatureGroupStatus returns a conventionally configured output for an attribute of this resource.
// Attribute: FeatureGroupStatus
func (t *AWS_SageMaker_FeatureGroup) GetConventionalOutputAtt__FeatureGroupStatus(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFeatureGroupStatus", t.GetAtt__FeatureGroupStatus())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_SageMaker_FeatureGroup) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_SageMaker_FeatureGroup

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

func (t *AWS_SageMaker_FeatureGroup) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
