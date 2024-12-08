// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_healthlake

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_HealthLake_FHIRDatastore)(nil)
	_ cfz.Resource                   = (*AWS_HealthLake_FHIRDatastore)(nil)
)

const (
	// AWS_HealthLake_FHIRDatastore__Type is the CloudFormation type for AWS::HealthLake::FHIRDatastore.
	AWS_HealthLake_FHIRDatastore__Type = "AWS::HealthLake::FHIRDatastore"
)

var (
	// AWS_HealthLake_FHIRDatastore__AttributesMap reports all the CloudFormation attributes for AWS::HealthLake::FHIRDatastore.
	AWS_HealthLake_FHIRDatastore__AttributesMap = struct {
		CreatedAt         string
		CreatedAt_Nanos   string
		CreatedAt_Seconds string
		DatastoreArn      string
		DatastoreEndpoint string
		DatastoreId       string
		DatastoreStatus   string
	}{
		CreatedAt:         "CreatedAt",
		CreatedAt_Nanos:   "CreatedAt.Nanos",
		CreatedAt_Seconds: "CreatedAt.Seconds",
		DatastoreArn:      "DatastoreArn",
		DatastoreEndpoint: "DatastoreEndpoint",
		DatastoreId:       "DatastoreId",
		DatastoreStatus:   "DatastoreStatus",
	}

	// AWS_HealthLake_FHIRDatastore__AttributesSlice reports all the CloudFormation attributes for AWS::HealthLake::FHIRDatastore.
	AWS_HealthLake_FHIRDatastore__AttributesSlice = []string{
		AWS_HealthLake_FHIRDatastore__AttributesMap.CreatedAt,
		AWS_HealthLake_FHIRDatastore__AttributesMap.CreatedAt_Nanos,
		AWS_HealthLake_FHIRDatastore__AttributesMap.CreatedAt_Seconds,
		AWS_HealthLake_FHIRDatastore__AttributesMap.DatastoreArn,
		AWS_HealthLake_FHIRDatastore__AttributesMap.DatastoreEndpoint,
		AWS_HealthLake_FHIRDatastore__AttributesMap.DatastoreId,
		AWS_HealthLake_FHIRDatastore__AttributesMap.DatastoreStatus,
	}
)

var (
	// AWS_HealthLake_FHIRDatastore__PropertiesMap reports all the CloudFormation properties for AWS::HealthLake::FHIRDatastore.
	AWS_HealthLake_FHIRDatastore__PropertiesMap = struct {
		DatastoreName                 string
		DatastoreTypeVersion          string
		IdentityProviderConfiguration string
		PreloadDataConfig             string
		SseConfiguration              string
		Tags                          string
	}{
		DatastoreName:                 "DatastoreName",
		DatastoreTypeVersion:          "DatastoreTypeVersion",
		IdentityProviderConfiguration: "IdentityProviderConfiguration",
		PreloadDataConfig:             "PreloadDataConfig",
		SseConfiguration:              "SseConfiguration",
		Tags:                          "Tags",
	}

	// AWS_HealthLake_FHIRDatastore__PropertiesSlice reports all the CloudFormation properties for AWS::HealthLake::FHIRDatastore.
	AWS_HealthLake_FHIRDatastore__PropertiesSlice = []string{
		AWS_HealthLake_FHIRDatastore__PropertiesMap.DatastoreName,
		AWS_HealthLake_FHIRDatastore__PropertiesMap.DatastoreTypeVersion,
		AWS_HealthLake_FHIRDatastore__PropertiesMap.IdentityProviderConfiguration,
		AWS_HealthLake_FHIRDatastore__PropertiesMap.PreloadDataConfig,
		AWS_HealthLake_FHIRDatastore__PropertiesMap.SseConfiguration,
		AWS_HealthLake_FHIRDatastore__PropertiesMap.Tags,
	}
)

// AWS_HealthLake_FHIRDatastore is a binding for AWS::HealthLake::FHIRDatastore.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html
type AWS_HealthLake_FHIRDatastore struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DatastoreName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-datastorename
	DatastoreName cfz.Expression[string] `json:"DatastoreName,omitempty"`

	// DatastoreTypeVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-datastoretypeversion
	DatastoreTypeVersion cfz.Expression[string] `json:"DatastoreTypeVersion,omitempty"`

	// IdentityProviderConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration
	IdentityProviderConfiguration cfz.Expression[AWS_HealthLake_FHIRDatastore_IdentityProviderConfiguration] `json:"IdentityProviderConfiguration,omitempty"`

	// PreloadDataConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-preloaddataconfig
	PreloadDataConfig cfz.Expression[AWS_HealthLake_FHIRDatastore_PreloadDataConfig] `json:"PreloadDataConfig,omitempty"`

	// SseConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-sseconfiguration
	SseConfiguration cfz.Expression[AWS_HealthLake_FHIRDatastore_SseConfiguration] `json:"SseConfiguration,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_HealthLake_FHIRDatastore initializes a new *AWS_HealthLake_FHIRDatastore.
func New__AWS_HealthLake_FHIRDatastore(logicalName string) *AWS_HealthLake_FHIRDatastore {
	return &AWS_HealthLake_FHIRDatastore{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_HealthLake_FHIRDatastore) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_HealthLake_FHIRDatastore) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_HealthLake_FHIRDatastore) GetType() string {
	return AWS_HealthLake_FHIRDatastore__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_HealthLake_FHIRDatastore) Set__LogicalName(v string) *AWS_HealthLake_FHIRDatastore {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_HealthLake_FHIRDatastore) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_HealthLake_FHIRDatastore {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_HealthLake_FHIRDatastore) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_HealthLake_FHIRDatastore {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_HealthLake_FHIRDatastore) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_HealthLake_FHIRDatastore {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_HealthLake_FHIRDatastore) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_HealthLake_FHIRDatastore {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_HealthLake_FHIRDatastore) Set__RequestedOutputs(v []cfz.Output) *AWS_HealthLake_FHIRDatastore {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_HealthLake_FHIRDatastore) Add__RequestedOutputs(v ...cfz.Output) *AWS_HealthLake_FHIRDatastore {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DatastoreName updates property "DatastoreName".
func (t *AWS_HealthLake_FHIRDatastore) Set__DatastoreName(v cfz.Expression[string]) *AWS_HealthLake_FHIRDatastore {
	t.DatastoreName = v
	return t
}

// SetV__DatastoreName updates property "DatastoreName".
func (t *AWS_HealthLake_FHIRDatastore) SetV__DatastoreName(v string) *AWS_HealthLake_FHIRDatastore {
	t.DatastoreName = cfz.V(v)
	return t
}

// Set__DatastoreTypeVersion updates property "DatastoreTypeVersion".
func (t *AWS_HealthLake_FHIRDatastore) Set__DatastoreTypeVersion(v cfz.Expression[string]) *AWS_HealthLake_FHIRDatastore {
	t.DatastoreTypeVersion = v
	return t
}

// SetV__DatastoreTypeVersion updates property "DatastoreTypeVersion".
func (t *AWS_HealthLake_FHIRDatastore) SetV__DatastoreTypeVersion(v string) *AWS_HealthLake_FHIRDatastore {
	t.DatastoreTypeVersion = cfz.V(v)
	return t
}

// Set__IdentityProviderConfiguration updates property "IdentityProviderConfiguration".
func (t *AWS_HealthLake_FHIRDatastore) Set__IdentityProviderConfiguration(v cfz.Expression[AWS_HealthLake_FHIRDatastore_IdentityProviderConfiguration]) *AWS_HealthLake_FHIRDatastore {
	t.IdentityProviderConfiguration = v
	return t
}

// SetV__IdentityProviderConfiguration updates property "IdentityProviderConfiguration".
func (t *AWS_HealthLake_FHIRDatastore) SetV__IdentityProviderConfiguration(v AWS_HealthLake_FHIRDatastore_IdentityProviderConfiguration) *AWS_HealthLake_FHIRDatastore {
	t.IdentityProviderConfiguration = cfz.V(v)
	return t
}

// Set__PreloadDataConfig updates property "PreloadDataConfig".
func (t *AWS_HealthLake_FHIRDatastore) Set__PreloadDataConfig(v cfz.Expression[AWS_HealthLake_FHIRDatastore_PreloadDataConfig]) *AWS_HealthLake_FHIRDatastore {
	t.PreloadDataConfig = v
	return t
}

// SetV__PreloadDataConfig updates property "PreloadDataConfig".
func (t *AWS_HealthLake_FHIRDatastore) SetV__PreloadDataConfig(v AWS_HealthLake_FHIRDatastore_PreloadDataConfig) *AWS_HealthLake_FHIRDatastore {
	t.PreloadDataConfig = cfz.V(v)
	return t
}

// Set__SseConfiguration updates property "SseConfiguration".
func (t *AWS_HealthLake_FHIRDatastore) Set__SseConfiguration(v cfz.Expression[AWS_HealthLake_FHIRDatastore_SseConfiguration]) *AWS_HealthLake_FHIRDatastore {
	t.SseConfiguration = v
	return t
}

// SetV__SseConfiguration updates property "SseConfiguration".
func (t *AWS_HealthLake_FHIRDatastore) SetV__SseConfiguration(v AWS_HealthLake_FHIRDatastore_SseConfiguration) *AWS_HealthLake_FHIRDatastore {
	t.SseConfiguration = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_HealthLake_FHIRDatastore) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_HealthLake_FHIRDatastore {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_HealthLake_FHIRDatastore) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_HealthLake_FHIRDatastore {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_HealthLake_FHIRDatastore) SetSV__Tags(v ...cfz.Tag) *AWS_HealthLake_FHIRDatastore {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_HealthLake_FHIRDatastore) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CreatedAt returns a $cfz.Expression[AWS_HealthLake_FHIRDatastore_CreatedAt] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreatedAt
func (t *AWS_HealthLake_FHIRDatastore) GetAtt__CreatedAt() cfz.Expression[AWS_HealthLake_FHIRDatastore_CreatedAt] {
	return cfz.GetAtt[AWS_HealthLake_FHIRDatastore_CreatedAt](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_HealthLake_FHIRDatastore__AttributesMap.CreatedAt))
}

// GetAtt__CreatedAt_Nanos returns a $cfz.Expression[int32] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreatedAt.Nanos
func (t *AWS_HealthLake_FHIRDatastore) GetAtt__CreatedAt_Nanos() cfz.Expression[int32] {
	return cfz.GetAtt[int32](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_HealthLake_FHIRDatastore__AttributesMap.CreatedAt_Nanos))
}

// GetAtt__CreatedAt_Seconds returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreatedAt.Seconds
func (t *AWS_HealthLake_FHIRDatastore) GetAtt__CreatedAt_Seconds() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_HealthLake_FHIRDatastore__AttributesMap.CreatedAt_Seconds))
}

// GetAtt__DatastoreArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DatastoreArn
func (t *AWS_HealthLake_FHIRDatastore) GetAtt__DatastoreArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_HealthLake_FHIRDatastore__AttributesMap.DatastoreArn))
}

// GetAtt__DatastoreEndpoint returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DatastoreEndpoint
func (t *AWS_HealthLake_FHIRDatastore) GetAtt__DatastoreEndpoint() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_HealthLake_FHIRDatastore__AttributesMap.DatastoreEndpoint))
}

// GetAtt__DatastoreId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DatastoreId
func (t *AWS_HealthLake_FHIRDatastore) GetAtt__DatastoreId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_HealthLake_FHIRDatastore__AttributesMap.DatastoreId))
}

// GetAtt__DatastoreStatus returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DatastoreStatus
func (t *AWS_HealthLake_FHIRDatastore) GetAtt__DatastoreStatus() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_HealthLake_FHIRDatastore__AttributesMap.DatastoreStatus))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_HealthLake_FHIRDatastore) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreatedAt returns a conventionally configured output for an attribute of this resource.
// Attribute: CreatedAt
func (t *AWS_HealthLake_FHIRDatastore) GetConventionalOutputAtt__CreatedAt(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreatedAt", t.GetAtt__CreatedAt())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreatedAt_Nanos returns a conventionally configured output for an attribute of this resource.
// Attribute: CreatedAt.Nanos
func (t *AWS_HealthLake_FHIRDatastore) GetConventionalOutputAtt__CreatedAt_Nanos(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreatedAtNanos", t.GetAtt__CreatedAt_Nanos())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreatedAt_Seconds returns a conventionally configured output for an attribute of this resource.
// Attribute: CreatedAt.Seconds
func (t *AWS_HealthLake_FHIRDatastore) GetConventionalOutputAtt__CreatedAt_Seconds(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreatedAtSeconds", t.GetAtt__CreatedAt_Seconds())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DatastoreArn returns a conventionally configured output for an attribute of this resource.
// Attribute: DatastoreArn
func (t *AWS_HealthLake_FHIRDatastore) GetConventionalOutputAtt__DatastoreArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDatastoreArn", t.GetAtt__DatastoreArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DatastoreEndpoint returns a conventionally configured output for an attribute of this resource.
// Attribute: DatastoreEndpoint
func (t *AWS_HealthLake_FHIRDatastore) GetConventionalOutputAtt__DatastoreEndpoint(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDatastoreEndpoint", t.GetAtt__DatastoreEndpoint())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DatastoreId returns a conventionally configured output for an attribute of this resource.
// Attribute: DatastoreId
func (t *AWS_HealthLake_FHIRDatastore) GetConventionalOutputAtt__DatastoreId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDatastoreId", t.GetAtt__DatastoreId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DatastoreStatus returns a conventionally configured output for an attribute of this resource.
// Attribute: DatastoreStatus
func (t *AWS_HealthLake_FHIRDatastore) GetConventionalOutputAtt__DatastoreStatus(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDatastoreStatus", t.GetAtt__DatastoreStatus())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_HealthLake_FHIRDatastore) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_HealthLake_FHIRDatastore

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

func (t *AWS_HealthLake_FHIRDatastore) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
