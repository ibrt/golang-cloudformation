// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dms

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_DMS_MigrationProject)(nil)
	_ cfz.Resource                   = (*AWS_DMS_MigrationProject)(nil)
)

const (
	// AWS_DMS_MigrationProject__Type is the CloudFormation type for AWS::DMS::MigrationProject.
	AWS_DMS_MigrationProject__Type = "AWS::DMS::MigrationProject"
)

var (
	// AWS_DMS_MigrationProject__AttributesMap reports all the CloudFormation attributes for AWS::DMS::MigrationProject.
	AWS_DMS_MigrationProject__AttributesMap = struct {
		MigrationProjectArn string
	}{
		MigrationProjectArn: "MigrationProjectArn",
	}

	// AWS_DMS_MigrationProject__AttributesSlice reports all the CloudFormation attributes for AWS::DMS::MigrationProject.
	AWS_DMS_MigrationProject__AttributesSlice = []string{
		AWS_DMS_MigrationProject__AttributesMap.MigrationProjectArn,
	}
)

var (
	// AWS_DMS_MigrationProject__PropertiesMap reports all the CloudFormation properties for AWS::DMS::MigrationProject.
	AWS_DMS_MigrationProject__PropertiesMap = struct {
		Description                           string
		InstanceProfileArn                    string
		InstanceProfileIdentifier             string
		InstanceProfileName                   string
		MigrationProjectIdentifier            string
		MigrationProjectName                  string
		SchemaConversionApplicationAttributes string
		SourceDataProviderDescriptors         string
		Tags                                  string
		TargetDataProviderDescriptors         string
		TransformationRules                   string
	}{
		Description:                           "Description",
		InstanceProfileArn:                    "InstanceProfileArn",
		InstanceProfileIdentifier:             "InstanceProfileIdentifier",
		InstanceProfileName:                   "InstanceProfileName",
		MigrationProjectIdentifier:            "MigrationProjectIdentifier",
		MigrationProjectName:                  "MigrationProjectName",
		SchemaConversionApplicationAttributes: "SchemaConversionApplicationAttributes",
		SourceDataProviderDescriptors:         "SourceDataProviderDescriptors",
		Tags:                                  "Tags",
		TargetDataProviderDescriptors:         "TargetDataProviderDescriptors",
		TransformationRules:                   "TransformationRules",
	}

	// AWS_DMS_MigrationProject__PropertiesSlice reports all the CloudFormation properties for AWS::DMS::MigrationProject.
	AWS_DMS_MigrationProject__PropertiesSlice = []string{
		AWS_DMS_MigrationProject__PropertiesMap.Description,
		AWS_DMS_MigrationProject__PropertiesMap.InstanceProfileArn,
		AWS_DMS_MigrationProject__PropertiesMap.InstanceProfileIdentifier,
		AWS_DMS_MigrationProject__PropertiesMap.InstanceProfileName,
		AWS_DMS_MigrationProject__PropertiesMap.MigrationProjectIdentifier,
		AWS_DMS_MigrationProject__PropertiesMap.MigrationProjectName,
		AWS_DMS_MigrationProject__PropertiesMap.SchemaConversionApplicationAttributes,
		AWS_DMS_MigrationProject__PropertiesMap.SourceDataProviderDescriptors,
		AWS_DMS_MigrationProject__PropertiesMap.Tags,
		AWS_DMS_MigrationProject__PropertiesMap.TargetDataProviderDescriptors,
		AWS_DMS_MigrationProject__PropertiesMap.TransformationRules,
	}
)

// AWS_DMS_MigrationProject is a binding for AWS::DMS::MigrationProject.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html
type AWS_DMS_MigrationProject struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// InstanceProfileArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-instanceprofilearn
	InstanceProfileArn cfz.Expression[string] `json:"InstanceProfileArn,omitempty"`

	// InstanceProfileIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-instanceprofileidentifier
	InstanceProfileIdentifier cfz.Expression[string] `json:"InstanceProfileIdentifier,omitempty"`

	// InstanceProfileName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-instanceprofilename
	InstanceProfileName cfz.Expression[string] `json:"InstanceProfileName,omitempty"`

	// MigrationProjectIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-migrationprojectidentifier
	MigrationProjectIdentifier cfz.Expression[string] `json:"MigrationProjectIdentifier,omitempty"`

	// MigrationProjectName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-migrationprojectname
	MigrationProjectName cfz.Expression[string] `json:"MigrationProjectName,omitempty"`

	// SchemaConversionApplicationAttributes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-schemaconversionapplicationattributes
	SchemaConversionApplicationAttributes cfz.Expression[AWS_DMS_MigrationProject_SchemaConversionApplicationAttributes] `json:"SchemaConversionApplicationAttributes,omitempty"`

	// SourceDataProviderDescriptors is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-sourcedataproviderdescriptors
	SourceDataProviderDescriptors cfz.ExpressionSlice[AWS_DMS_MigrationProject_DataProviderDescriptor] `json:"SourceDataProviderDescriptors,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// TargetDataProviderDescriptors is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-targetdataproviderdescriptors
	TargetDataProviderDescriptors cfz.ExpressionSlice[AWS_DMS_MigrationProject_DataProviderDescriptor] `json:"TargetDataProviderDescriptors,omitempty"`

	// TransformationRules is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-transformationrules
	TransformationRules cfz.Expression[string] `json:"TransformationRules,omitempty"`
}

// New__AWS_DMS_MigrationProject initializes a new *AWS_DMS_MigrationProject.
func New__AWS_DMS_MigrationProject(logicalName string) *AWS_DMS_MigrationProject {
	return &AWS_DMS_MigrationProject{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_DMS_MigrationProject) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_DMS_MigrationProject) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_DMS_MigrationProject) GetType() string {
	return AWS_DMS_MigrationProject__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_DMS_MigrationProject) Set__LogicalName(v string) *AWS_DMS_MigrationProject {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_DMS_MigrationProject) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_DMS_MigrationProject {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_DMS_MigrationProject) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_DMS_MigrationProject {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_DMS_MigrationProject) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_DMS_MigrationProject {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_DMS_MigrationProject) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_DMS_MigrationProject {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_DMS_MigrationProject) Set__RequestedOutputs(v []cfz.Output) *AWS_DMS_MigrationProject {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_DMS_MigrationProject) Add__RequestedOutputs(v ...cfz.Output) *AWS_DMS_MigrationProject {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_DMS_MigrationProject) Set__Description(v cfz.Expression[string]) *AWS_DMS_MigrationProject {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_DMS_MigrationProject) SetV__Description(v string) *AWS_DMS_MigrationProject {
	t.Description = cfz.V(v)
	return t
}

// Set__InstanceProfileArn updates property "InstanceProfileArn".
func (t *AWS_DMS_MigrationProject) Set__InstanceProfileArn(v cfz.Expression[string]) *AWS_DMS_MigrationProject {
	t.InstanceProfileArn = v
	return t
}

// SetV__InstanceProfileArn updates property "InstanceProfileArn".
func (t *AWS_DMS_MigrationProject) SetV__InstanceProfileArn(v string) *AWS_DMS_MigrationProject {
	t.InstanceProfileArn = cfz.V(v)
	return t
}

// Set__InstanceProfileIdentifier updates property "InstanceProfileIdentifier".
func (t *AWS_DMS_MigrationProject) Set__InstanceProfileIdentifier(v cfz.Expression[string]) *AWS_DMS_MigrationProject {
	t.InstanceProfileIdentifier = v
	return t
}

// SetV__InstanceProfileIdentifier updates property "InstanceProfileIdentifier".
func (t *AWS_DMS_MigrationProject) SetV__InstanceProfileIdentifier(v string) *AWS_DMS_MigrationProject {
	t.InstanceProfileIdentifier = cfz.V(v)
	return t
}

// Set__InstanceProfileName updates property "InstanceProfileName".
func (t *AWS_DMS_MigrationProject) Set__InstanceProfileName(v cfz.Expression[string]) *AWS_DMS_MigrationProject {
	t.InstanceProfileName = v
	return t
}

// SetV__InstanceProfileName updates property "InstanceProfileName".
func (t *AWS_DMS_MigrationProject) SetV__InstanceProfileName(v string) *AWS_DMS_MigrationProject {
	t.InstanceProfileName = cfz.V(v)
	return t
}

// Set__MigrationProjectIdentifier updates property "MigrationProjectIdentifier".
func (t *AWS_DMS_MigrationProject) Set__MigrationProjectIdentifier(v cfz.Expression[string]) *AWS_DMS_MigrationProject {
	t.MigrationProjectIdentifier = v
	return t
}

// SetV__MigrationProjectIdentifier updates property "MigrationProjectIdentifier".
func (t *AWS_DMS_MigrationProject) SetV__MigrationProjectIdentifier(v string) *AWS_DMS_MigrationProject {
	t.MigrationProjectIdentifier = cfz.V(v)
	return t
}

// Set__MigrationProjectName updates property "MigrationProjectName".
func (t *AWS_DMS_MigrationProject) Set__MigrationProjectName(v cfz.Expression[string]) *AWS_DMS_MigrationProject {
	t.MigrationProjectName = v
	return t
}

// SetV__MigrationProjectName updates property "MigrationProjectName".
func (t *AWS_DMS_MigrationProject) SetV__MigrationProjectName(v string) *AWS_DMS_MigrationProject {
	t.MigrationProjectName = cfz.V(v)
	return t
}

// Set__SchemaConversionApplicationAttributes updates property "SchemaConversionApplicationAttributes".
func (t *AWS_DMS_MigrationProject) Set__SchemaConversionApplicationAttributes(v cfz.Expression[AWS_DMS_MigrationProject_SchemaConversionApplicationAttributes]) *AWS_DMS_MigrationProject {
	t.SchemaConversionApplicationAttributes = v
	return t
}

// SetV__SchemaConversionApplicationAttributes updates property "SchemaConversionApplicationAttributes".
func (t *AWS_DMS_MigrationProject) SetV__SchemaConversionApplicationAttributes(v AWS_DMS_MigrationProject_SchemaConversionApplicationAttributes) *AWS_DMS_MigrationProject {
	t.SchemaConversionApplicationAttributes = cfz.V(v)
	return t
}

// Set__SourceDataProviderDescriptors updates property "SourceDataProviderDescriptors".
func (t *AWS_DMS_MigrationProject) Set__SourceDataProviderDescriptors(v cfz.ExpressionSlice[AWS_DMS_MigrationProject_DataProviderDescriptor]) *AWS_DMS_MigrationProject {
	t.SourceDataProviderDescriptors = v
	return t
}

// SetS__SourceDataProviderDescriptors updates property "SourceDataProviderDescriptors".
func (t *AWS_DMS_MigrationProject) SetS__SourceDataProviderDescriptors(v ...cfz.Expression[AWS_DMS_MigrationProject_DataProviderDescriptor]) *AWS_DMS_MigrationProject {
	t.SourceDataProviderDescriptors = cfz.S(v...)
	return t
}

// SetSV__SourceDataProviderDescriptors updates property "SourceDataProviderDescriptors".
func (t *AWS_DMS_MigrationProject) SetSV__SourceDataProviderDescriptors(v ...AWS_DMS_MigrationProject_DataProviderDescriptor) *AWS_DMS_MigrationProject {
	t.SourceDataProviderDescriptors = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_DMS_MigrationProject) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_DMS_MigrationProject {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_DMS_MigrationProject) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_DMS_MigrationProject {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_DMS_MigrationProject) SetSV__Tags(v ...cfz.Tag) *AWS_DMS_MigrationProject {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TargetDataProviderDescriptors updates property "TargetDataProviderDescriptors".
func (t *AWS_DMS_MigrationProject) Set__TargetDataProviderDescriptors(v cfz.ExpressionSlice[AWS_DMS_MigrationProject_DataProviderDescriptor]) *AWS_DMS_MigrationProject {
	t.TargetDataProviderDescriptors = v
	return t
}

// SetS__TargetDataProviderDescriptors updates property "TargetDataProviderDescriptors".
func (t *AWS_DMS_MigrationProject) SetS__TargetDataProviderDescriptors(v ...cfz.Expression[AWS_DMS_MigrationProject_DataProviderDescriptor]) *AWS_DMS_MigrationProject {
	t.TargetDataProviderDescriptors = cfz.S(v...)
	return t
}

// SetSV__TargetDataProviderDescriptors updates property "TargetDataProviderDescriptors".
func (t *AWS_DMS_MigrationProject) SetSV__TargetDataProviderDescriptors(v ...AWS_DMS_MigrationProject_DataProviderDescriptor) *AWS_DMS_MigrationProject {
	t.TargetDataProviderDescriptors = cfz.SV(v...)
	return t
}

// Set__TransformationRules updates property "TransformationRules".
func (t *AWS_DMS_MigrationProject) Set__TransformationRules(v cfz.Expression[string]) *AWS_DMS_MigrationProject {
	t.TransformationRules = v
	return t
}

// SetV__TransformationRules updates property "TransformationRules".
func (t *AWS_DMS_MigrationProject) SetV__TransformationRules(v string) *AWS_DMS_MigrationProject {
	t.TransformationRules = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_DMS_MigrationProject) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__MigrationProjectArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: MigrationProjectArn
func (t *AWS_DMS_MigrationProject) GetAtt__MigrationProjectArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DMS_MigrationProject__AttributesMap.MigrationProjectArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_DMS_MigrationProject) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__MigrationProjectArn returns a conventionally configured output for an attribute of this resource.
// Attribute: MigrationProjectArn
func (t *AWS_DMS_MigrationProject) GetConventionalOutputAtt__MigrationProjectArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttMigrationProjectArn", t.GetAtt__MigrationProjectArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_DMS_MigrationProject) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_DMS_MigrationProject

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

func (t *AWS_DMS_MigrationProject) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
