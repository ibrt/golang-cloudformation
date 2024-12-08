// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_eventschemas

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EventSchemas_Schema)(nil)
	_ cfz.Resource                   = (*AWS_EventSchemas_Schema)(nil)
)

const (
	// AWS_EventSchemas_Schema__Type is the CloudFormation type for AWS::EventSchemas::Schema.
	AWS_EventSchemas_Schema__Type = "AWS::EventSchemas::Schema"
)

var (
	// AWS_EventSchemas_Schema__AttributesMap reports all the CloudFormation attributes for AWS::EventSchemas::Schema.
	AWS_EventSchemas_Schema__AttributesMap = struct {
		LastModified       string
		SchemaArn          string
		SchemaName         string
		SchemaVersion      string
		VersionCreatedDate string
	}{
		LastModified:       "LastModified",
		SchemaArn:          "SchemaArn",
		SchemaName:         "SchemaName",
		SchemaVersion:      "SchemaVersion",
		VersionCreatedDate: "VersionCreatedDate",
	}

	// AWS_EventSchemas_Schema__AttributesSlice reports all the CloudFormation attributes for AWS::EventSchemas::Schema.
	AWS_EventSchemas_Schema__AttributesSlice = []string{
		AWS_EventSchemas_Schema__AttributesMap.LastModified,
		AWS_EventSchemas_Schema__AttributesMap.SchemaArn,
		AWS_EventSchemas_Schema__AttributesMap.SchemaName,
		AWS_EventSchemas_Schema__AttributesMap.SchemaVersion,
		AWS_EventSchemas_Schema__AttributesMap.VersionCreatedDate,
	}
)

var (
	// AWS_EventSchemas_Schema__PropertiesMap reports all the CloudFormation properties for AWS::EventSchemas::Schema.
	AWS_EventSchemas_Schema__PropertiesMap = struct {
		Content      string
		Description  string
		RegistryName string
		SchemaName   string
		Tags         string
		Type         string
	}{
		Content:      "Content",
		Description:  "Description",
		RegistryName: "RegistryName",
		SchemaName:   "SchemaName",
		Tags:         "Tags",
		Type:         "Type",
	}

	// AWS_EventSchemas_Schema__PropertiesSlice reports all the CloudFormation properties for AWS::EventSchemas::Schema.
	AWS_EventSchemas_Schema__PropertiesSlice = []string{
		AWS_EventSchemas_Schema__PropertiesMap.Content,
		AWS_EventSchemas_Schema__PropertiesMap.Description,
		AWS_EventSchemas_Schema__PropertiesMap.RegistryName,
		AWS_EventSchemas_Schema__PropertiesMap.SchemaName,
		AWS_EventSchemas_Schema__PropertiesMap.Tags,
		AWS_EventSchemas_Schema__PropertiesMap.Type,
	}
)

// AWS_EventSchemas_Schema is a binding for AWS::EventSchemas::Schema.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html
type AWS_EventSchemas_Schema struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Content is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-content
	Content cfz.Expression[string] `json:"Content,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// RegistryName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-registryname
	RegistryName cfz.Expression[string] `json:"RegistryName,omitempty"`

	// SchemaName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-schemaname
	SchemaName cfz.Expression[string] `json:"SchemaName,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-tags
	Tags cfz.ExpressionSlice[AWS_EventSchemas_Schema_TagsEntry] `json:"Tags,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-schema.html#cfn-eventschemas-schema-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_EventSchemas_Schema initializes a new *AWS_EventSchemas_Schema.
func New__AWS_EventSchemas_Schema(logicalName string) *AWS_EventSchemas_Schema {
	return &AWS_EventSchemas_Schema{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EventSchemas_Schema) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EventSchemas_Schema) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EventSchemas_Schema) GetType() string {
	return AWS_EventSchemas_Schema__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EventSchemas_Schema) Set__LogicalName(v string) *AWS_EventSchemas_Schema {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EventSchemas_Schema) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EventSchemas_Schema {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EventSchemas_Schema) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EventSchemas_Schema {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EventSchemas_Schema) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EventSchemas_Schema {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EventSchemas_Schema) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EventSchemas_Schema {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EventSchemas_Schema) Set__RequestedOutputs(v []cfz.Output) *AWS_EventSchemas_Schema {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EventSchemas_Schema) Add__RequestedOutputs(v ...cfz.Output) *AWS_EventSchemas_Schema {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Content updates property "Content".
func (t *AWS_EventSchemas_Schema) Set__Content(v cfz.Expression[string]) *AWS_EventSchemas_Schema {
	t.Content = v
	return t
}

// SetV__Content updates property "Content".
func (t *AWS_EventSchemas_Schema) SetV__Content(v string) *AWS_EventSchemas_Schema {
	t.Content = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_EventSchemas_Schema) Set__Description(v cfz.Expression[string]) *AWS_EventSchemas_Schema {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_EventSchemas_Schema) SetV__Description(v string) *AWS_EventSchemas_Schema {
	t.Description = cfz.V(v)
	return t
}

// Set__RegistryName updates property "RegistryName".
func (t *AWS_EventSchemas_Schema) Set__RegistryName(v cfz.Expression[string]) *AWS_EventSchemas_Schema {
	t.RegistryName = v
	return t
}

// SetV__RegistryName updates property "RegistryName".
func (t *AWS_EventSchemas_Schema) SetV__RegistryName(v string) *AWS_EventSchemas_Schema {
	t.RegistryName = cfz.V(v)
	return t
}

// Set__SchemaName updates property "SchemaName".
func (t *AWS_EventSchemas_Schema) Set__SchemaName(v cfz.Expression[string]) *AWS_EventSchemas_Schema {
	t.SchemaName = v
	return t
}

// SetV__SchemaName updates property "SchemaName".
func (t *AWS_EventSchemas_Schema) SetV__SchemaName(v string) *AWS_EventSchemas_Schema {
	t.SchemaName = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EventSchemas_Schema) Set__Tags(v cfz.ExpressionSlice[AWS_EventSchemas_Schema_TagsEntry]) *AWS_EventSchemas_Schema {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_EventSchemas_Schema) SetS__Tags(v ...cfz.Expression[AWS_EventSchemas_Schema_TagsEntry]) *AWS_EventSchemas_Schema {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_EventSchemas_Schema) SetSV__Tags(v ...AWS_EventSchemas_Schema_TagsEntry) *AWS_EventSchemas_Schema {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__Type updates property "Type".
func (t *AWS_EventSchemas_Schema) Set__Type(v cfz.Expression[string]) *AWS_EventSchemas_Schema {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t *AWS_EventSchemas_Schema) SetV__Type(v string) *AWS_EventSchemas_Schema {
	t.Type = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EventSchemas_Schema) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__LastModified returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LastModified
func (t *AWS_EventSchemas_Schema) GetAtt__LastModified() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EventSchemas_Schema__AttributesMap.LastModified))
}

// GetAtt__SchemaArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: SchemaArn
func (t *AWS_EventSchemas_Schema) GetAtt__SchemaArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EventSchemas_Schema__AttributesMap.SchemaArn))
}

// GetAtt__SchemaName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: SchemaName
func (t *AWS_EventSchemas_Schema) GetAtt__SchemaName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EventSchemas_Schema__AttributesMap.SchemaName))
}

// GetAtt__SchemaVersion returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: SchemaVersion
func (t *AWS_EventSchemas_Schema) GetAtt__SchemaVersion() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EventSchemas_Schema__AttributesMap.SchemaVersion))
}

// GetAtt__VersionCreatedDate returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: VersionCreatedDate
func (t *AWS_EventSchemas_Schema) GetAtt__VersionCreatedDate() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EventSchemas_Schema__AttributesMap.VersionCreatedDate))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EventSchemas_Schema) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LastModified returns a conventionally configured output for an attribute of this resource.
// Attribute: LastModified
func (t *AWS_EventSchemas_Schema) GetConventionalOutputAtt__LastModified(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLastModified", t.GetAtt__LastModified())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__SchemaArn returns a conventionally configured output for an attribute of this resource.
// Attribute: SchemaArn
func (t *AWS_EventSchemas_Schema) GetConventionalOutputAtt__SchemaArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttSchemaArn", t.GetAtt__SchemaArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__SchemaName returns a conventionally configured output for an attribute of this resource.
// Attribute: SchemaName
func (t *AWS_EventSchemas_Schema) GetConventionalOutputAtt__SchemaName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttSchemaName", t.GetAtt__SchemaName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__SchemaVersion returns a conventionally configured output for an attribute of this resource.
// Attribute: SchemaVersion
func (t *AWS_EventSchemas_Schema) GetConventionalOutputAtt__SchemaVersion(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttSchemaVersion", t.GetAtt__SchemaVersion())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__VersionCreatedDate returns a conventionally configured output for an attribute of this resource.
// Attribute: VersionCreatedDate
func (t *AWS_EventSchemas_Schema) GetConventionalOutputAtt__VersionCreatedDate(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttVersionCreatedDate", t.GetAtt__VersionCreatedDate())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EventSchemas_Schema) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EventSchemas_Schema

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

func (t *AWS_EventSchemas_Schema) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
