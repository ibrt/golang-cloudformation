// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wisdom

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Wisdom_KnowledgeBase)(nil)
	_ cfz.Resource                   = (*AWS_Wisdom_KnowledgeBase)(nil)
)

const (
	// AWS_Wisdom_KnowledgeBase__Type is the CloudFormation type for AWS::Wisdom::KnowledgeBase.
	AWS_Wisdom_KnowledgeBase__Type = "AWS::Wisdom::KnowledgeBase"
)

var (
	// AWS_Wisdom_KnowledgeBase__AttributesMap reports all the CloudFormation attributes for AWS::Wisdom::KnowledgeBase.
	AWS_Wisdom_KnowledgeBase__AttributesMap = struct {
		KnowledgeBaseArn string
		KnowledgeBaseId  string
	}{
		KnowledgeBaseArn: "KnowledgeBaseArn",
		KnowledgeBaseId:  "KnowledgeBaseId",
	}

	// AWS_Wisdom_KnowledgeBase__AttributesSlice reports all the CloudFormation attributes for AWS::Wisdom::KnowledgeBase.
	AWS_Wisdom_KnowledgeBase__AttributesSlice = []string{
		AWS_Wisdom_KnowledgeBase__AttributesMap.KnowledgeBaseArn,
		AWS_Wisdom_KnowledgeBase__AttributesMap.KnowledgeBaseId,
	}
)

var (
	// AWS_Wisdom_KnowledgeBase__PropertiesMap reports all the CloudFormation properties for AWS::Wisdom::KnowledgeBase.
	AWS_Wisdom_KnowledgeBase__PropertiesMap = struct {
		Description                       string
		KnowledgeBaseType                 string
		Name                              string
		RenderingConfiguration            string
		ServerSideEncryptionConfiguration string
		SourceConfiguration               string
		Tags                              string
		VectorIngestionConfiguration      string
	}{
		Description:                       "Description",
		KnowledgeBaseType:                 "KnowledgeBaseType",
		Name:                              "Name",
		RenderingConfiguration:            "RenderingConfiguration",
		ServerSideEncryptionConfiguration: "ServerSideEncryptionConfiguration",
		SourceConfiguration:               "SourceConfiguration",
		Tags:                              "Tags",
		VectorIngestionConfiguration:      "VectorIngestionConfiguration",
	}

	// AWS_Wisdom_KnowledgeBase__PropertiesSlice reports all the CloudFormation properties for AWS::Wisdom::KnowledgeBase.
	AWS_Wisdom_KnowledgeBase__PropertiesSlice = []string{
		AWS_Wisdom_KnowledgeBase__PropertiesMap.Description,
		AWS_Wisdom_KnowledgeBase__PropertiesMap.KnowledgeBaseType,
		AWS_Wisdom_KnowledgeBase__PropertiesMap.Name,
		AWS_Wisdom_KnowledgeBase__PropertiesMap.RenderingConfiguration,
		AWS_Wisdom_KnowledgeBase__PropertiesMap.ServerSideEncryptionConfiguration,
		AWS_Wisdom_KnowledgeBase__PropertiesMap.SourceConfiguration,
		AWS_Wisdom_KnowledgeBase__PropertiesMap.Tags,
		AWS_Wisdom_KnowledgeBase__PropertiesMap.VectorIngestionConfiguration,
	}
)

// AWS_Wisdom_KnowledgeBase is a binding for AWS::Wisdom::KnowledgeBase.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html
type AWS_Wisdom_KnowledgeBase struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// KnowledgeBaseType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-knowledgebasetype
	KnowledgeBaseType cfz.Expression[string] `json:"KnowledgeBaseType,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// RenderingConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-renderingconfiguration
	RenderingConfiguration cfz.Expression[AWS_Wisdom_KnowledgeBase_RenderingConfiguration] `json:"RenderingConfiguration,omitempty"`

	// ServerSideEncryptionConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-serversideencryptionconfiguration
	ServerSideEncryptionConfiguration cfz.Expression[AWS_Wisdom_KnowledgeBase_ServerSideEncryptionConfiguration] `json:"ServerSideEncryptionConfiguration,omitempty"`

	// SourceConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-sourceconfiguration
	SourceConfiguration cfz.Expression[AWS_Wisdom_KnowledgeBase_SourceConfiguration] `json:"SourceConfiguration,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// VectorIngestionConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-vectoringestionconfiguration
	VectorIngestionConfiguration cfz.Expression[AWS_Wisdom_KnowledgeBase_VectorIngestionConfiguration] `json:"VectorIngestionConfiguration,omitempty"`
}

// New__AWS_Wisdom_KnowledgeBase initializes a new *AWS_Wisdom_KnowledgeBase.
func New__AWS_Wisdom_KnowledgeBase(logicalName string) *AWS_Wisdom_KnowledgeBase {
	return &AWS_Wisdom_KnowledgeBase{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Wisdom_KnowledgeBase) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Wisdom_KnowledgeBase) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Wisdom_KnowledgeBase) GetType() string {
	return AWS_Wisdom_KnowledgeBase__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Wisdom_KnowledgeBase) Set__LogicalName(v string) *AWS_Wisdom_KnowledgeBase {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Wisdom_KnowledgeBase) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Wisdom_KnowledgeBase {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Wisdom_KnowledgeBase) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Wisdom_KnowledgeBase {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Wisdom_KnowledgeBase) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Wisdom_KnowledgeBase {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Wisdom_KnowledgeBase) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Wisdom_KnowledgeBase {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Wisdom_KnowledgeBase) Set__RequestedOutputs(v []cfz.Output) *AWS_Wisdom_KnowledgeBase {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Wisdom_KnowledgeBase) Add__RequestedOutputs(v ...cfz.Output) *AWS_Wisdom_KnowledgeBase {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Wisdom_KnowledgeBase) Set__Description(v cfz.Expression[string]) *AWS_Wisdom_KnowledgeBase {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Wisdom_KnowledgeBase) SetV__Description(v string) *AWS_Wisdom_KnowledgeBase {
	t.Description = cfz.V(v)
	return t
}

// Set__KnowledgeBaseType updates property "KnowledgeBaseType".
func (t *AWS_Wisdom_KnowledgeBase) Set__KnowledgeBaseType(v cfz.Expression[string]) *AWS_Wisdom_KnowledgeBase {
	t.KnowledgeBaseType = v
	return t
}

// SetV__KnowledgeBaseType updates property "KnowledgeBaseType".
func (t *AWS_Wisdom_KnowledgeBase) SetV__KnowledgeBaseType(v string) *AWS_Wisdom_KnowledgeBase {
	t.KnowledgeBaseType = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Wisdom_KnowledgeBase) Set__Name(v cfz.Expression[string]) *AWS_Wisdom_KnowledgeBase {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Wisdom_KnowledgeBase) SetV__Name(v string) *AWS_Wisdom_KnowledgeBase {
	t.Name = cfz.V(v)
	return t
}

// Set__RenderingConfiguration updates property "RenderingConfiguration".
func (t *AWS_Wisdom_KnowledgeBase) Set__RenderingConfiguration(v cfz.Expression[AWS_Wisdom_KnowledgeBase_RenderingConfiguration]) *AWS_Wisdom_KnowledgeBase {
	t.RenderingConfiguration = v
	return t
}

// SetV__RenderingConfiguration updates property "RenderingConfiguration".
func (t *AWS_Wisdom_KnowledgeBase) SetV__RenderingConfiguration(v AWS_Wisdom_KnowledgeBase_RenderingConfiguration) *AWS_Wisdom_KnowledgeBase {
	t.RenderingConfiguration = cfz.V(v)
	return t
}

// Set__ServerSideEncryptionConfiguration updates property "ServerSideEncryptionConfiguration".
func (t *AWS_Wisdom_KnowledgeBase) Set__ServerSideEncryptionConfiguration(v cfz.Expression[AWS_Wisdom_KnowledgeBase_ServerSideEncryptionConfiguration]) *AWS_Wisdom_KnowledgeBase {
	t.ServerSideEncryptionConfiguration = v
	return t
}

// SetV__ServerSideEncryptionConfiguration updates property "ServerSideEncryptionConfiguration".
func (t *AWS_Wisdom_KnowledgeBase) SetV__ServerSideEncryptionConfiguration(v AWS_Wisdom_KnowledgeBase_ServerSideEncryptionConfiguration) *AWS_Wisdom_KnowledgeBase {
	t.ServerSideEncryptionConfiguration = cfz.V(v)
	return t
}

// Set__SourceConfiguration updates property "SourceConfiguration".
func (t *AWS_Wisdom_KnowledgeBase) Set__SourceConfiguration(v cfz.Expression[AWS_Wisdom_KnowledgeBase_SourceConfiguration]) *AWS_Wisdom_KnowledgeBase {
	t.SourceConfiguration = v
	return t
}

// SetV__SourceConfiguration updates property "SourceConfiguration".
func (t *AWS_Wisdom_KnowledgeBase) SetV__SourceConfiguration(v AWS_Wisdom_KnowledgeBase_SourceConfiguration) *AWS_Wisdom_KnowledgeBase {
	t.SourceConfiguration = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Wisdom_KnowledgeBase) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Wisdom_KnowledgeBase {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Wisdom_KnowledgeBase) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Wisdom_KnowledgeBase {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Wisdom_KnowledgeBase) SetSV__Tags(v ...cfz.Tag) *AWS_Wisdom_KnowledgeBase {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__VectorIngestionConfiguration updates property "VectorIngestionConfiguration".
func (t *AWS_Wisdom_KnowledgeBase) Set__VectorIngestionConfiguration(v cfz.Expression[AWS_Wisdom_KnowledgeBase_VectorIngestionConfiguration]) *AWS_Wisdom_KnowledgeBase {
	t.VectorIngestionConfiguration = v
	return t
}

// SetV__VectorIngestionConfiguration updates property "VectorIngestionConfiguration".
func (t *AWS_Wisdom_KnowledgeBase) SetV__VectorIngestionConfiguration(v AWS_Wisdom_KnowledgeBase_VectorIngestionConfiguration) *AWS_Wisdom_KnowledgeBase {
	t.VectorIngestionConfiguration = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Wisdom_KnowledgeBase) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__KnowledgeBaseArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: KnowledgeBaseArn
func (t *AWS_Wisdom_KnowledgeBase) GetAtt__KnowledgeBaseArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Wisdom_KnowledgeBase__AttributesMap.KnowledgeBaseArn))
}

// GetAtt__KnowledgeBaseId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: KnowledgeBaseId
func (t *AWS_Wisdom_KnowledgeBase) GetAtt__KnowledgeBaseId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Wisdom_KnowledgeBase__AttributesMap.KnowledgeBaseId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Wisdom_KnowledgeBase) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__KnowledgeBaseArn returns a conventionally configured output for an attribute of this resource.
// Attribute: KnowledgeBaseArn
func (t *AWS_Wisdom_KnowledgeBase) GetConventionalOutputAtt__KnowledgeBaseArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttKnowledgeBaseArn", t.GetAtt__KnowledgeBaseArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__KnowledgeBaseId returns a conventionally configured output for an attribute of this resource.
// Attribute: KnowledgeBaseId
func (t *AWS_Wisdom_KnowledgeBase) GetConventionalOutputAtt__KnowledgeBaseId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttKnowledgeBaseId", t.GetAtt__KnowledgeBaseId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Wisdom_KnowledgeBase) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Wisdom_KnowledgeBase

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

func (t *AWS_Wisdom_KnowledgeBase) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
