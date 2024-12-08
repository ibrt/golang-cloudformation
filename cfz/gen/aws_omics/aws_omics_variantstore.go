// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_omics

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Omics_VariantStore)(nil)
	_ cfz.Resource                   = (*AWS_Omics_VariantStore)(nil)
)

const (
	// AWS_Omics_VariantStore__Type is the CloudFormation type for AWS::Omics::VariantStore.
	AWS_Omics_VariantStore__Type = "AWS::Omics::VariantStore"
)

var (
	// AWS_Omics_VariantStore__AttributesMap reports all the CloudFormation attributes for AWS::Omics::VariantStore.
	AWS_Omics_VariantStore__AttributesMap = struct {
		CreationTime   string
		Id             string
		Status         string
		StatusMessage  string
		StoreArn       string
		StoreSizeBytes string
		UpdateTime     string
	}{
		CreationTime:   "CreationTime",
		Id:             "Id",
		Status:         "Status",
		StatusMessage:  "StatusMessage",
		StoreArn:       "StoreArn",
		StoreSizeBytes: "StoreSizeBytes",
		UpdateTime:     "UpdateTime",
	}

	// AWS_Omics_VariantStore__AttributesSlice reports all the CloudFormation attributes for AWS::Omics::VariantStore.
	AWS_Omics_VariantStore__AttributesSlice = []string{
		AWS_Omics_VariantStore__AttributesMap.CreationTime,
		AWS_Omics_VariantStore__AttributesMap.Id,
		AWS_Omics_VariantStore__AttributesMap.Status,
		AWS_Omics_VariantStore__AttributesMap.StatusMessage,
		AWS_Omics_VariantStore__AttributesMap.StoreArn,
		AWS_Omics_VariantStore__AttributesMap.StoreSizeBytes,
		AWS_Omics_VariantStore__AttributesMap.UpdateTime,
	}
)

var (
	// AWS_Omics_VariantStore__PropertiesMap reports all the CloudFormation properties for AWS::Omics::VariantStore.
	AWS_Omics_VariantStore__PropertiesMap = struct {
		Description string
		Name        string
		Reference   string
		SseConfig   string
		Tags        string
	}{
		Description: "Description",
		Name:        "Name",
		Reference:   "Reference",
		SseConfig:   "SseConfig",
		Tags:        "Tags",
	}

	// AWS_Omics_VariantStore__PropertiesSlice reports all the CloudFormation properties for AWS::Omics::VariantStore.
	AWS_Omics_VariantStore__PropertiesSlice = []string{
		AWS_Omics_VariantStore__PropertiesMap.Description,
		AWS_Omics_VariantStore__PropertiesMap.Name,
		AWS_Omics_VariantStore__PropertiesMap.Reference,
		AWS_Omics_VariantStore__PropertiesMap.SseConfig,
		AWS_Omics_VariantStore__PropertiesMap.Tags,
	}
)

// AWS_Omics_VariantStore is a binding for AWS::Omics::VariantStore.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html
type AWS_Omics_VariantStore struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Reference is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-reference
	Reference cfz.Expression[AWS_Omics_VariantStore_ReferenceItem] `json:"Reference,omitempty"`

	// SseConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-sseconfig
	SseConfig cfz.Expression[AWS_Omics_VariantStore_SseConfig] `json:"SseConfig,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html#cfn-omics-variantstore-tags
	Tags cfz.ExpressionMap[string] `json:"Tags,omitempty"`
}

// New__AWS_Omics_VariantStore initializes a new *AWS_Omics_VariantStore.
func New__AWS_Omics_VariantStore(logicalName string) *AWS_Omics_VariantStore {
	return &AWS_Omics_VariantStore{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Omics_VariantStore) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Omics_VariantStore) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Omics_VariantStore) GetType() string {
	return AWS_Omics_VariantStore__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Omics_VariantStore) Set__LogicalName(v string) *AWS_Omics_VariantStore {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Omics_VariantStore) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Omics_VariantStore {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Omics_VariantStore) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Omics_VariantStore {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Omics_VariantStore) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Omics_VariantStore {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Omics_VariantStore) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Omics_VariantStore {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Omics_VariantStore) Set__RequestedOutputs(v []cfz.Output) *AWS_Omics_VariantStore {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Omics_VariantStore) Add__RequestedOutputs(v ...cfz.Output) *AWS_Omics_VariantStore {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Omics_VariantStore) Set__Description(v cfz.Expression[string]) *AWS_Omics_VariantStore {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Omics_VariantStore) SetV__Description(v string) *AWS_Omics_VariantStore {
	t.Description = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Omics_VariantStore) Set__Name(v cfz.Expression[string]) *AWS_Omics_VariantStore {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Omics_VariantStore) SetV__Name(v string) *AWS_Omics_VariantStore {
	t.Name = cfz.V(v)
	return t
}

// Set__Reference updates property "Reference".
func (t *AWS_Omics_VariantStore) Set__Reference(v cfz.Expression[AWS_Omics_VariantStore_ReferenceItem]) *AWS_Omics_VariantStore {
	t.Reference = v
	return t
}

// SetV__Reference updates property "Reference".
func (t *AWS_Omics_VariantStore) SetV__Reference(v AWS_Omics_VariantStore_ReferenceItem) *AWS_Omics_VariantStore {
	t.Reference = cfz.V(v)
	return t
}

// Set__SseConfig updates property "SseConfig".
func (t *AWS_Omics_VariantStore) Set__SseConfig(v cfz.Expression[AWS_Omics_VariantStore_SseConfig]) *AWS_Omics_VariantStore {
	t.SseConfig = v
	return t
}

// SetV__SseConfig updates property "SseConfig".
func (t *AWS_Omics_VariantStore) SetV__SseConfig(v AWS_Omics_VariantStore_SseConfig) *AWS_Omics_VariantStore {
	t.SseConfig = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Omics_VariantStore) Set__Tags(v cfz.ExpressionMap[string]) *AWS_Omics_VariantStore {
	t.Tags = v
	return t
}

// SetM__Tags updates property "Tags".
func (t *AWS_Omics_VariantStore) SetM__Tags(v ...map[string]cfz.Expression[string]) *AWS_Omics_VariantStore {
	t.Tags = cfz.M(v...)
	return t
}

// SetMV__Tags updates property "Tags".
func (t *AWS_Omics_VariantStore) SetMV__Tags(v ...map[string]string) *AWS_Omics_VariantStore {
	t.Tags = cfz.MV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Omics_VariantStore) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CreationTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreationTime
func (t *AWS_Omics_VariantStore) GetAtt__CreationTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Omics_VariantStore__AttributesMap.CreationTime))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_Omics_VariantStore) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Omics_VariantStore__AttributesMap.Id))
}

// GetAtt__Status returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Status
func (t *AWS_Omics_VariantStore) GetAtt__Status() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Omics_VariantStore__AttributesMap.Status))
}

// GetAtt__StatusMessage returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: StatusMessage
func (t *AWS_Omics_VariantStore) GetAtt__StatusMessage() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Omics_VariantStore__AttributesMap.StatusMessage))
}

// GetAtt__StoreArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: StoreArn
func (t *AWS_Omics_VariantStore) GetAtt__StoreArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Omics_VariantStore__AttributesMap.StoreArn))
}

// GetAtt__StoreSizeBytes returns a $cfz.Expression[float64] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: StoreSizeBytes
func (t *AWS_Omics_VariantStore) GetAtt__StoreSizeBytes() cfz.Expression[float64] {
	return cfz.GetAtt[float64](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Omics_VariantStore__AttributesMap.StoreSizeBytes))
}

// GetAtt__UpdateTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: UpdateTime
func (t *AWS_Omics_VariantStore) GetAtt__UpdateTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Omics_VariantStore__AttributesMap.UpdateTime))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Omics_VariantStore) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreationTime returns a conventionally configured output for an attribute of this resource.
// Attribute: CreationTime
func (t *AWS_Omics_VariantStore) GetConventionalOutputAtt__CreationTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreationTime", t.GetAtt__CreationTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_Omics_VariantStore) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Status returns a conventionally configured output for an attribute of this resource.
// Attribute: Status
func (t *AWS_Omics_VariantStore) GetConventionalOutputAtt__Status(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatus", t.GetAtt__Status())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__StatusMessage returns a conventionally configured output for an attribute of this resource.
// Attribute: StatusMessage
func (t *AWS_Omics_VariantStore) GetConventionalOutputAtt__StatusMessage(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatusMessage", t.GetAtt__StatusMessage())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__StoreArn returns a conventionally configured output for an attribute of this resource.
// Attribute: StoreArn
func (t *AWS_Omics_VariantStore) GetConventionalOutputAtt__StoreArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStoreArn", t.GetAtt__StoreArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__StoreSizeBytes returns a conventionally configured output for an attribute of this resource.
// Attribute: StoreSizeBytes
func (t *AWS_Omics_VariantStore) GetConventionalOutputAtt__StoreSizeBytes(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStoreSizeBytes", t.GetAtt__StoreSizeBytes())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__UpdateTime returns a conventionally configured output for an attribute of this resource.
// Attribute: UpdateTime
func (t *AWS_Omics_VariantStore) GetConventionalOutputAtt__UpdateTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttUpdateTime", t.GetAtt__UpdateTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Omics_VariantStore) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Omics_VariantStore

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

func (t *AWS_Omics_VariantStore) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
