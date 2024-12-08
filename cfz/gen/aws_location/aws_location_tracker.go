// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_location

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Location_Tracker)(nil)
	_ cfz.Resource                   = (*AWS_Location_Tracker)(nil)
)

const (
	// AWS_Location_Tracker__Type is the CloudFormation type for AWS::Location::Tracker.
	AWS_Location_Tracker__Type = "AWS::Location::Tracker"
)

var (
	// AWS_Location_Tracker__AttributesMap reports all the CloudFormation attributes for AWS::Location::Tracker.
	AWS_Location_Tracker__AttributesMap = struct {
		Arn        string
		CreateTime string
		TrackerArn string
		UpdateTime string
	}{
		Arn:        "Arn",
		CreateTime: "CreateTime",
		TrackerArn: "TrackerArn",
		UpdateTime: "UpdateTime",
	}

	// AWS_Location_Tracker__AttributesSlice reports all the CloudFormation attributes for AWS::Location::Tracker.
	AWS_Location_Tracker__AttributesSlice = []string{
		AWS_Location_Tracker__AttributesMap.Arn,
		AWS_Location_Tracker__AttributesMap.CreateTime,
		AWS_Location_Tracker__AttributesMap.TrackerArn,
		AWS_Location_Tracker__AttributesMap.UpdateTime,
	}
)

var (
	// AWS_Location_Tracker__PropertiesMap reports all the CloudFormation properties for AWS::Location::Tracker.
	AWS_Location_Tracker__PropertiesMap = struct {
		Description                   string
		EventBridgeEnabled            string
		KmsKeyEnableGeospatialQueries string
		KmsKeyId                      string
		PositionFiltering             string
		Tags                          string
		TrackerName                   string
	}{
		Description:                   "Description",
		EventBridgeEnabled:            "EventBridgeEnabled",
		KmsKeyEnableGeospatialQueries: "KmsKeyEnableGeospatialQueries",
		KmsKeyId:                      "KmsKeyId",
		PositionFiltering:             "PositionFiltering",
		Tags:                          "Tags",
		TrackerName:                   "TrackerName",
	}

	// AWS_Location_Tracker__PropertiesSlice reports all the CloudFormation properties for AWS::Location::Tracker.
	AWS_Location_Tracker__PropertiesSlice = []string{
		AWS_Location_Tracker__PropertiesMap.Description,
		AWS_Location_Tracker__PropertiesMap.EventBridgeEnabled,
		AWS_Location_Tracker__PropertiesMap.KmsKeyEnableGeospatialQueries,
		AWS_Location_Tracker__PropertiesMap.KmsKeyId,
		AWS_Location_Tracker__PropertiesMap.PositionFiltering,
		AWS_Location_Tracker__PropertiesMap.Tags,
		AWS_Location_Tracker__PropertiesMap.TrackerName,
	}
)

// AWS_Location_Tracker is a binding for AWS::Location::Tracker.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html
type AWS_Location_Tracker struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// EventBridgeEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-eventbridgeenabled
	EventBridgeEnabled cfz.Expression[bool] `json:"EventBridgeEnabled,omitempty"`

	// KmsKeyEnableGeospatialQueries is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-kmskeyenablegeospatialqueries
	KmsKeyEnableGeospatialQueries cfz.Expression[bool] `json:"KmsKeyEnableGeospatialQueries,omitempty"`

	// KmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-kmskeyid
	KmsKeyId cfz.Expression[string] `json:"KmsKeyId,omitempty"`

	// PositionFiltering is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-positionfiltering
	PositionFiltering cfz.Expression[string] `json:"PositionFiltering,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// TrackerName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-trackername
	TrackerName cfz.Expression[string] `json:"TrackerName,omitempty"`
}

// New__AWS_Location_Tracker initializes a new *AWS_Location_Tracker.
func New__AWS_Location_Tracker(logicalName string) *AWS_Location_Tracker {
	return &AWS_Location_Tracker{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Location_Tracker) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Location_Tracker) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Location_Tracker) GetType() string {
	return AWS_Location_Tracker__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Location_Tracker) Set__LogicalName(v string) *AWS_Location_Tracker {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Location_Tracker) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Location_Tracker {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Location_Tracker) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Location_Tracker {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Location_Tracker) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Location_Tracker {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Location_Tracker) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Location_Tracker {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Location_Tracker) Set__RequestedOutputs(v []cfz.Output) *AWS_Location_Tracker {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Location_Tracker) Add__RequestedOutputs(v ...cfz.Output) *AWS_Location_Tracker {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Location_Tracker) Set__Description(v cfz.Expression[string]) *AWS_Location_Tracker {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Location_Tracker) SetV__Description(v string) *AWS_Location_Tracker {
	t.Description = cfz.V(v)
	return t
}

// Set__EventBridgeEnabled updates property "EventBridgeEnabled".
func (t *AWS_Location_Tracker) Set__EventBridgeEnabled(v cfz.Expression[bool]) *AWS_Location_Tracker {
	t.EventBridgeEnabled = v
	return t
}

// SetV__EventBridgeEnabled updates property "EventBridgeEnabled".
func (t *AWS_Location_Tracker) SetV__EventBridgeEnabled(v bool) *AWS_Location_Tracker {
	t.EventBridgeEnabled = cfz.V(v)
	return t
}

// Set__KmsKeyEnableGeospatialQueries updates property "KmsKeyEnableGeospatialQueries".
func (t *AWS_Location_Tracker) Set__KmsKeyEnableGeospatialQueries(v cfz.Expression[bool]) *AWS_Location_Tracker {
	t.KmsKeyEnableGeospatialQueries = v
	return t
}

// SetV__KmsKeyEnableGeospatialQueries updates property "KmsKeyEnableGeospatialQueries".
func (t *AWS_Location_Tracker) SetV__KmsKeyEnableGeospatialQueries(v bool) *AWS_Location_Tracker {
	t.KmsKeyEnableGeospatialQueries = cfz.V(v)
	return t
}

// Set__KmsKeyId updates property "KmsKeyId".
func (t *AWS_Location_Tracker) Set__KmsKeyId(v cfz.Expression[string]) *AWS_Location_Tracker {
	t.KmsKeyId = v
	return t
}

// SetV__KmsKeyId updates property "KmsKeyId".
func (t *AWS_Location_Tracker) SetV__KmsKeyId(v string) *AWS_Location_Tracker {
	t.KmsKeyId = cfz.V(v)
	return t
}

// Set__PositionFiltering updates property "PositionFiltering".
func (t *AWS_Location_Tracker) Set__PositionFiltering(v cfz.Expression[string]) *AWS_Location_Tracker {
	t.PositionFiltering = v
	return t
}

// SetV__PositionFiltering updates property "PositionFiltering".
func (t *AWS_Location_Tracker) SetV__PositionFiltering(v string) *AWS_Location_Tracker {
	t.PositionFiltering = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Location_Tracker) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Location_Tracker {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Location_Tracker) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Location_Tracker {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Location_Tracker) SetSV__Tags(v ...cfz.Tag) *AWS_Location_Tracker {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TrackerName updates property "TrackerName".
func (t *AWS_Location_Tracker) Set__TrackerName(v cfz.Expression[string]) *AWS_Location_Tracker {
	t.TrackerName = v
	return t
}

// SetV__TrackerName updates property "TrackerName".
func (t *AWS_Location_Tracker) SetV__TrackerName(v string) *AWS_Location_Tracker {
	t.TrackerName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Location_Tracker) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Location_Tracker) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Location_Tracker__AttributesMap.Arn))
}

// GetAtt__CreateTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreateTime
func (t *AWS_Location_Tracker) GetAtt__CreateTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Location_Tracker__AttributesMap.CreateTime))
}

// GetAtt__TrackerArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TrackerArn
func (t *AWS_Location_Tracker) GetAtt__TrackerArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Location_Tracker__AttributesMap.TrackerArn))
}

// GetAtt__UpdateTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: UpdateTime
func (t *AWS_Location_Tracker) GetAtt__UpdateTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Location_Tracker__AttributesMap.UpdateTime))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Location_Tracker) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Location_Tracker) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreateTime returns a conventionally configured output for an attribute of this resource.
// Attribute: CreateTime
func (t *AWS_Location_Tracker) GetConventionalOutputAtt__CreateTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreateTime", t.GetAtt__CreateTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TrackerArn returns a conventionally configured output for an attribute of this resource.
// Attribute: TrackerArn
func (t *AWS_Location_Tracker) GetConventionalOutputAtt__TrackerArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTrackerArn", t.GetAtt__TrackerArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__UpdateTime returns a conventionally configured output for an attribute of this resource.
// Attribute: UpdateTime
func (t *AWS_Location_Tracker) GetConventionalOutputAtt__UpdateTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttUpdateTime", t.GetAtt__UpdateTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Location_Tracker) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Location_Tracker

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

func (t *AWS_Location_Tracker) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
