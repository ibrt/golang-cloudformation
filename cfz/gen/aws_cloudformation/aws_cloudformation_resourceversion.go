// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudformation

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_CloudFormation_ResourceVersion)(nil)
	_ cfz.Resource                   = (*AWS_CloudFormation_ResourceVersion)(nil)
)

const (
	// AWS_CloudFormation_ResourceVersion__Type is the CloudFormation type for AWS::CloudFormation::ResourceVersion.
	AWS_CloudFormation_ResourceVersion__Type = "AWS::CloudFormation::ResourceVersion"
)

var (
	// AWS_CloudFormation_ResourceVersion__AttributesMap reports all the CloudFormation attributes for AWS::CloudFormation::ResourceVersion.
	AWS_CloudFormation_ResourceVersion__AttributesMap = struct {
		Arn              string
		IsDefaultVersion string
		ProvisioningType string
		TypeArn          string
		VersionId        string
		Visibility       string
	}{
		Arn:              "Arn",
		IsDefaultVersion: "IsDefaultVersion",
		ProvisioningType: "ProvisioningType",
		TypeArn:          "TypeArn",
		VersionId:        "VersionId",
		Visibility:       "Visibility",
	}

	// AWS_CloudFormation_ResourceVersion__AttributesSlice reports all the CloudFormation attributes for AWS::CloudFormation::ResourceVersion.
	AWS_CloudFormation_ResourceVersion__AttributesSlice = []string{
		AWS_CloudFormation_ResourceVersion__AttributesMap.Arn,
		AWS_CloudFormation_ResourceVersion__AttributesMap.IsDefaultVersion,
		AWS_CloudFormation_ResourceVersion__AttributesMap.ProvisioningType,
		AWS_CloudFormation_ResourceVersion__AttributesMap.TypeArn,
		AWS_CloudFormation_ResourceVersion__AttributesMap.VersionId,
		AWS_CloudFormation_ResourceVersion__AttributesMap.Visibility,
	}
)

var (
	// AWS_CloudFormation_ResourceVersion__PropertiesMap reports all the CloudFormation properties for AWS::CloudFormation::ResourceVersion.
	AWS_CloudFormation_ResourceVersion__PropertiesMap = struct {
		ExecutionRoleArn     string
		LoggingConfig        string
		SchemaHandlerPackage string
		TypeName             string
	}{
		ExecutionRoleArn:     "ExecutionRoleArn",
		LoggingConfig:        "LoggingConfig",
		SchemaHandlerPackage: "SchemaHandlerPackage",
		TypeName:             "TypeName",
	}

	// AWS_CloudFormation_ResourceVersion__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFormation::ResourceVersion.
	AWS_CloudFormation_ResourceVersion__PropertiesSlice = []string{
		AWS_CloudFormation_ResourceVersion__PropertiesMap.ExecutionRoleArn,
		AWS_CloudFormation_ResourceVersion__PropertiesMap.LoggingConfig,
		AWS_CloudFormation_ResourceVersion__PropertiesMap.SchemaHandlerPackage,
		AWS_CloudFormation_ResourceVersion__PropertiesMap.TypeName,
	}
)

// AWS_CloudFormation_ResourceVersion is a binding for AWS::CloudFormation::ResourceVersion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html
type AWS_CloudFormation_ResourceVersion struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ExecutionRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-executionrolearn
	ExecutionRoleArn cfz.Expression[string] `json:"ExecutionRoleArn,omitempty"`

	// LoggingConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-loggingconfig
	LoggingConfig cfz.Expression[AWS_CloudFormation_ResourceVersion_LoggingConfig] `json:"LoggingConfig,omitempty"`

	// SchemaHandlerPackage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-schemahandlerpackage
	SchemaHandlerPackage cfz.Expression[string] `json:"SchemaHandlerPackage,omitempty"`

	// TypeName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-typename
	TypeName cfz.Expression[string] `json:"TypeName,omitempty"`
}

// New__AWS_CloudFormation_ResourceVersion initializes a new *AWS_CloudFormation_ResourceVersion.
func New__AWS_CloudFormation_ResourceVersion(logicalName string) *AWS_CloudFormation_ResourceVersion {
	return &AWS_CloudFormation_ResourceVersion{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_CloudFormation_ResourceVersion) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_CloudFormation_ResourceVersion) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_CloudFormation_ResourceVersion) GetType() string {
	return AWS_CloudFormation_ResourceVersion__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_CloudFormation_ResourceVersion) Set__LogicalName(v string) *AWS_CloudFormation_ResourceVersion {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_CloudFormation_ResourceVersion) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_CloudFormation_ResourceVersion {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_CloudFormation_ResourceVersion) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_CloudFormation_ResourceVersion {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_CloudFormation_ResourceVersion) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_CloudFormation_ResourceVersion {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_CloudFormation_ResourceVersion) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_CloudFormation_ResourceVersion {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_CloudFormation_ResourceVersion) Set__RequestedOutputs(v []cfz.Output) *AWS_CloudFormation_ResourceVersion {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_CloudFormation_ResourceVersion) Add__RequestedOutputs(v ...cfz.Output) *AWS_CloudFormation_ResourceVersion {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ExecutionRoleArn updates property "ExecutionRoleArn".
func (t *AWS_CloudFormation_ResourceVersion) Set__ExecutionRoleArn(v cfz.Expression[string]) *AWS_CloudFormation_ResourceVersion {
	t.ExecutionRoleArn = v
	return t
}

// SetV__ExecutionRoleArn updates property "ExecutionRoleArn".
func (t *AWS_CloudFormation_ResourceVersion) SetV__ExecutionRoleArn(v string) *AWS_CloudFormation_ResourceVersion {
	t.ExecutionRoleArn = cfz.V(v)
	return t
}

// Set__LoggingConfig updates property "LoggingConfig".
func (t *AWS_CloudFormation_ResourceVersion) Set__LoggingConfig(v cfz.Expression[AWS_CloudFormation_ResourceVersion_LoggingConfig]) *AWS_CloudFormation_ResourceVersion {
	t.LoggingConfig = v
	return t
}

// SetV__LoggingConfig updates property "LoggingConfig".
func (t *AWS_CloudFormation_ResourceVersion) SetV__LoggingConfig(v AWS_CloudFormation_ResourceVersion_LoggingConfig) *AWS_CloudFormation_ResourceVersion {
	t.LoggingConfig = cfz.V(v)
	return t
}

// Set__SchemaHandlerPackage updates property "SchemaHandlerPackage".
func (t *AWS_CloudFormation_ResourceVersion) Set__SchemaHandlerPackage(v cfz.Expression[string]) *AWS_CloudFormation_ResourceVersion {
	t.SchemaHandlerPackage = v
	return t
}

// SetV__SchemaHandlerPackage updates property "SchemaHandlerPackage".
func (t *AWS_CloudFormation_ResourceVersion) SetV__SchemaHandlerPackage(v string) *AWS_CloudFormation_ResourceVersion {
	t.SchemaHandlerPackage = cfz.V(v)
	return t
}

// Set__TypeName updates property "TypeName".
func (t *AWS_CloudFormation_ResourceVersion) Set__TypeName(v cfz.Expression[string]) *AWS_CloudFormation_ResourceVersion {
	t.TypeName = v
	return t
}

// SetV__TypeName updates property "TypeName".
func (t *AWS_CloudFormation_ResourceVersion) SetV__TypeName(v string) *AWS_CloudFormation_ResourceVersion {
	t.TypeName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_CloudFormation_ResourceVersion) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_CloudFormation_ResourceVersion) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CloudFormation_ResourceVersion__AttributesMap.Arn))
}

// GetAtt__IsDefaultVersion returns a $cfz.Expression[bool] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: IsDefaultVersion
func (t *AWS_CloudFormation_ResourceVersion) GetAtt__IsDefaultVersion() cfz.Expression[bool] {
	return cfz.GetAtt[bool](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CloudFormation_ResourceVersion__AttributesMap.IsDefaultVersion))
}

// GetAtt__ProvisioningType returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ProvisioningType
func (t *AWS_CloudFormation_ResourceVersion) GetAtt__ProvisioningType() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CloudFormation_ResourceVersion__AttributesMap.ProvisioningType))
}

// GetAtt__TypeArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TypeArn
func (t *AWS_CloudFormation_ResourceVersion) GetAtt__TypeArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CloudFormation_ResourceVersion__AttributesMap.TypeArn))
}

// GetAtt__VersionId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: VersionId
func (t *AWS_CloudFormation_ResourceVersion) GetAtt__VersionId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CloudFormation_ResourceVersion__AttributesMap.VersionId))
}

// GetAtt__Visibility returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Visibility
func (t *AWS_CloudFormation_ResourceVersion) GetAtt__Visibility() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CloudFormation_ResourceVersion__AttributesMap.Visibility))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_CloudFormation_ResourceVersion) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_CloudFormation_ResourceVersion) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__IsDefaultVersion returns a conventionally configured output for an attribute of this resource.
// Attribute: IsDefaultVersion
func (t *AWS_CloudFormation_ResourceVersion) GetConventionalOutputAtt__IsDefaultVersion(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttIsDefaultVersion", t.GetAtt__IsDefaultVersion())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ProvisioningType returns a conventionally configured output for an attribute of this resource.
// Attribute: ProvisioningType
func (t *AWS_CloudFormation_ResourceVersion) GetConventionalOutputAtt__ProvisioningType(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttProvisioningType", t.GetAtt__ProvisioningType())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TypeArn returns a conventionally configured output for an attribute of this resource.
// Attribute: TypeArn
func (t *AWS_CloudFormation_ResourceVersion) GetConventionalOutputAtt__TypeArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTypeArn", t.GetAtt__TypeArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__VersionId returns a conventionally configured output for an attribute of this resource.
// Attribute: VersionId
func (t *AWS_CloudFormation_ResourceVersion) GetConventionalOutputAtt__VersionId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttVersionId", t.GetAtt__VersionId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Visibility returns a conventionally configured output for an attribute of this resource.
// Attribute: Visibility
func (t *AWS_CloudFormation_ResourceVersion) GetConventionalOutputAtt__Visibility(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttVisibility", t.GetAtt__Visibility())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_CloudFormation_ResourceVersion) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_CloudFormation_ResourceVersion

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

func (t *AWS_CloudFormation_ResourceVersion) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
