// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_SageMaker_ImageVersion)(nil)
	_ cfz.Resource                   = (*AWS_SageMaker_ImageVersion)(nil)
)

const (
	// AWS_SageMaker_ImageVersion__Type is the CloudFormation type for AWS::SageMaker::ImageVersion.
	AWS_SageMaker_ImageVersion__Type = "AWS::SageMaker::ImageVersion"
)

var (
	// AWS_SageMaker_ImageVersion__AttributesMap reports all the CloudFormation attributes for AWS::SageMaker::ImageVersion.
	AWS_SageMaker_ImageVersion__AttributesMap = struct {
		ContainerImage  string
		ImageArn        string
		ImageVersionArn string
		Version         string
	}{
		ContainerImage:  "ContainerImage",
		ImageArn:        "ImageArn",
		ImageVersionArn: "ImageVersionArn",
		Version:         "Version",
	}

	// AWS_SageMaker_ImageVersion__AttributesSlice reports all the CloudFormation attributes for AWS::SageMaker::ImageVersion.
	AWS_SageMaker_ImageVersion__AttributesSlice = []string{
		AWS_SageMaker_ImageVersion__AttributesMap.ContainerImage,
		AWS_SageMaker_ImageVersion__AttributesMap.ImageArn,
		AWS_SageMaker_ImageVersion__AttributesMap.ImageVersionArn,
		AWS_SageMaker_ImageVersion__AttributesMap.Version,
	}
)

var (
	// AWS_SageMaker_ImageVersion__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ImageVersion.
	AWS_SageMaker_ImageVersion__PropertiesMap = struct {
		Alias           string
		Aliases         string
		BaseImage       string
		Horovod         string
		ImageName       string
		JobType         string
		MLFramework     string
		Processor       string
		ProgrammingLang string
		ReleaseNotes    string
		VendorGuidance  string
	}{
		Alias:           "Alias",
		Aliases:         "Aliases",
		BaseImage:       "BaseImage",
		Horovod:         "Horovod",
		ImageName:       "ImageName",
		JobType:         "JobType",
		MLFramework:     "MLFramework",
		Processor:       "Processor",
		ProgrammingLang: "ProgrammingLang",
		ReleaseNotes:    "ReleaseNotes",
		VendorGuidance:  "VendorGuidance",
	}

	// AWS_SageMaker_ImageVersion__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ImageVersion.
	AWS_SageMaker_ImageVersion__PropertiesSlice = []string{
		AWS_SageMaker_ImageVersion__PropertiesMap.Alias,
		AWS_SageMaker_ImageVersion__PropertiesMap.Aliases,
		AWS_SageMaker_ImageVersion__PropertiesMap.BaseImage,
		AWS_SageMaker_ImageVersion__PropertiesMap.Horovod,
		AWS_SageMaker_ImageVersion__PropertiesMap.ImageName,
		AWS_SageMaker_ImageVersion__PropertiesMap.JobType,
		AWS_SageMaker_ImageVersion__PropertiesMap.MLFramework,
		AWS_SageMaker_ImageVersion__PropertiesMap.Processor,
		AWS_SageMaker_ImageVersion__PropertiesMap.ProgrammingLang,
		AWS_SageMaker_ImageVersion__PropertiesMap.ReleaseNotes,
		AWS_SageMaker_ImageVersion__PropertiesMap.VendorGuidance,
	}
)

// AWS_SageMaker_ImageVersion is a binding for AWS::SageMaker::ImageVersion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html
type AWS_SageMaker_ImageVersion struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Alias is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-alias
	Alias cfz.Expression[string] `json:"Alias,omitempty"`

	// Aliases is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-aliases
	Aliases cfz.ExpressionSlice[string] `json:"Aliases,omitempty"`

	// BaseImage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-baseimage
	BaseImage cfz.Expression[string] `json:"BaseImage,omitempty"`

	// Horovod is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-horovod
	Horovod cfz.Expression[bool] `json:"Horovod,omitempty"`

	// ImageName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-imagename
	ImageName cfz.Expression[string] `json:"ImageName,omitempty"`

	// JobType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-jobtype
	JobType cfz.Expression[string] `json:"JobType,omitempty"`

	// MLFramework is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-mlframework
	MLFramework cfz.Expression[string] `json:"MLFramework,omitempty"`

	// Processor is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-processor
	Processor cfz.Expression[string] `json:"Processor,omitempty"`

	// ProgrammingLang is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-programminglang
	ProgrammingLang cfz.Expression[string] `json:"ProgrammingLang,omitempty"`

	// ReleaseNotes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-releasenotes
	ReleaseNotes cfz.Expression[string] `json:"ReleaseNotes,omitempty"`

	// VendorGuidance is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-vendorguidance
	VendorGuidance cfz.Expression[string] `json:"VendorGuidance,omitempty"`
}

// New__AWS_SageMaker_ImageVersion initializes a new *AWS_SageMaker_ImageVersion.
func New__AWS_SageMaker_ImageVersion(logicalName string) *AWS_SageMaker_ImageVersion {
	return &AWS_SageMaker_ImageVersion{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_SageMaker_ImageVersion) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_SageMaker_ImageVersion) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_SageMaker_ImageVersion) GetType() string {
	return AWS_SageMaker_ImageVersion__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_SageMaker_ImageVersion) Set__LogicalName(v string) *AWS_SageMaker_ImageVersion {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_SageMaker_ImageVersion) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_SageMaker_ImageVersion {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_SageMaker_ImageVersion) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_SageMaker_ImageVersion {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_SageMaker_ImageVersion) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_SageMaker_ImageVersion {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_SageMaker_ImageVersion) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_SageMaker_ImageVersion {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_SageMaker_ImageVersion) Set__RequestedOutputs(v []cfz.Output) *AWS_SageMaker_ImageVersion {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_SageMaker_ImageVersion) Add__RequestedOutputs(v ...cfz.Output) *AWS_SageMaker_ImageVersion {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Alias updates property "Alias".
func (t *AWS_SageMaker_ImageVersion) Set__Alias(v cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.Alias = v
	return t
}

// SetV__Alias updates property "Alias".
func (t *AWS_SageMaker_ImageVersion) SetV__Alias(v string) *AWS_SageMaker_ImageVersion {
	t.Alias = cfz.V(v)
	return t
}

// Set__Aliases updates property "Aliases".
func (t *AWS_SageMaker_ImageVersion) Set__Aliases(v cfz.ExpressionSlice[string]) *AWS_SageMaker_ImageVersion {
	t.Aliases = v
	return t
}

// SetS__Aliases updates property "Aliases".
func (t *AWS_SageMaker_ImageVersion) SetS__Aliases(v ...cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.Aliases = cfz.S(v...)
	return t
}

// SetSV__Aliases updates property "Aliases".
func (t *AWS_SageMaker_ImageVersion) SetSV__Aliases(v ...string) *AWS_SageMaker_ImageVersion {
	t.Aliases = cfz.SV(v...)
	return t
}

// Set__BaseImage updates property "BaseImage".
func (t *AWS_SageMaker_ImageVersion) Set__BaseImage(v cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.BaseImage = v
	return t
}

// SetV__BaseImage updates property "BaseImage".
func (t *AWS_SageMaker_ImageVersion) SetV__BaseImage(v string) *AWS_SageMaker_ImageVersion {
	t.BaseImage = cfz.V(v)
	return t
}

// Set__Horovod updates property "Horovod".
func (t *AWS_SageMaker_ImageVersion) Set__Horovod(v cfz.Expression[bool]) *AWS_SageMaker_ImageVersion {
	t.Horovod = v
	return t
}

// SetV__Horovod updates property "Horovod".
func (t *AWS_SageMaker_ImageVersion) SetV__Horovod(v bool) *AWS_SageMaker_ImageVersion {
	t.Horovod = cfz.V(v)
	return t
}

// Set__ImageName updates property "ImageName".
func (t *AWS_SageMaker_ImageVersion) Set__ImageName(v cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.ImageName = v
	return t
}

// SetV__ImageName updates property "ImageName".
func (t *AWS_SageMaker_ImageVersion) SetV__ImageName(v string) *AWS_SageMaker_ImageVersion {
	t.ImageName = cfz.V(v)
	return t
}

// Set__JobType updates property "JobType".
func (t *AWS_SageMaker_ImageVersion) Set__JobType(v cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.JobType = v
	return t
}

// SetV__JobType updates property "JobType".
func (t *AWS_SageMaker_ImageVersion) SetV__JobType(v string) *AWS_SageMaker_ImageVersion {
	t.JobType = cfz.V(v)
	return t
}

// Set__MLFramework updates property "MLFramework".
func (t *AWS_SageMaker_ImageVersion) Set__MLFramework(v cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.MLFramework = v
	return t
}

// SetV__MLFramework updates property "MLFramework".
func (t *AWS_SageMaker_ImageVersion) SetV__MLFramework(v string) *AWS_SageMaker_ImageVersion {
	t.MLFramework = cfz.V(v)
	return t
}

// Set__Processor updates property "Processor".
func (t *AWS_SageMaker_ImageVersion) Set__Processor(v cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.Processor = v
	return t
}

// SetV__Processor updates property "Processor".
func (t *AWS_SageMaker_ImageVersion) SetV__Processor(v string) *AWS_SageMaker_ImageVersion {
	t.Processor = cfz.V(v)
	return t
}

// Set__ProgrammingLang updates property "ProgrammingLang".
func (t *AWS_SageMaker_ImageVersion) Set__ProgrammingLang(v cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.ProgrammingLang = v
	return t
}

// SetV__ProgrammingLang updates property "ProgrammingLang".
func (t *AWS_SageMaker_ImageVersion) SetV__ProgrammingLang(v string) *AWS_SageMaker_ImageVersion {
	t.ProgrammingLang = cfz.V(v)
	return t
}

// Set__ReleaseNotes updates property "ReleaseNotes".
func (t *AWS_SageMaker_ImageVersion) Set__ReleaseNotes(v cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.ReleaseNotes = v
	return t
}

// SetV__ReleaseNotes updates property "ReleaseNotes".
func (t *AWS_SageMaker_ImageVersion) SetV__ReleaseNotes(v string) *AWS_SageMaker_ImageVersion {
	t.ReleaseNotes = cfz.V(v)
	return t
}

// Set__VendorGuidance updates property "VendorGuidance".
func (t *AWS_SageMaker_ImageVersion) Set__VendorGuidance(v cfz.Expression[string]) *AWS_SageMaker_ImageVersion {
	t.VendorGuidance = v
	return t
}

// SetV__VendorGuidance updates property "VendorGuidance".
func (t *AWS_SageMaker_ImageVersion) SetV__VendorGuidance(v string) *AWS_SageMaker_ImageVersion {
	t.VendorGuidance = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_SageMaker_ImageVersion) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__ContainerImage returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ContainerImage
func (t *AWS_SageMaker_ImageVersion) GetAtt__ContainerImage() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SageMaker_ImageVersion__AttributesMap.ContainerImage))
}

// GetAtt__ImageArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ImageArn
func (t *AWS_SageMaker_ImageVersion) GetAtt__ImageArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SageMaker_ImageVersion__AttributesMap.ImageArn))
}

// GetAtt__ImageVersionArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ImageVersionArn
func (t *AWS_SageMaker_ImageVersion) GetAtt__ImageVersionArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SageMaker_ImageVersion__AttributesMap.ImageVersionArn))
}

// GetAtt__Version returns a $cfz.Expression[int32] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Version
func (t *AWS_SageMaker_ImageVersion) GetAtt__Version() cfz.Expression[int32] {
	return cfz.GetAtt[int32](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SageMaker_ImageVersion__AttributesMap.Version))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_SageMaker_ImageVersion) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ContainerImage returns a conventionally configured output for an attribute of this resource.
// Attribute: ContainerImage
func (t *AWS_SageMaker_ImageVersion) GetConventionalOutputAtt__ContainerImage(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttContainerImage", t.GetAtt__ContainerImage())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ImageArn returns a conventionally configured output for an attribute of this resource.
// Attribute: ImageArn
func (t *AWS_SageMaker_ImageVersion) GetConventionalOutputAtt__ImageArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttImageArn", t.GetAtt__ImageArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ImageVersionArn returns a conventionally configured output for an attribute of this resource.
// Attribute: ImageVersionArn
func (t *AWS_SageMaker_ImageVersion) GetConventionalOutputAtt__ImageVersionArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttImageVersionArn", t.GetAtt__ImageVersionArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Version returns a conventionally configured output for an attribute of this resource.
// Attribute: Version
func (t *AWS_SageMaker_ImageVersion) GetConventionalOutputAtt__Version(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttVersion", t.GetAtt__Version())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_SageMaker_ImageVersion) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_SageMaker_ImageVersion

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

func (t *AWS_SageMaker_ImageVersion) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
