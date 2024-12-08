// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Lex_BotVersion)(nil)
	_ cfz.Resource                   = (*AWS_Lex_BotVersion)(nil)
)

const (
	// AWS_Lex_BotVersion__Type is the CloudFormation type for AWS::Lex::BotVersion.
	AWS_Lex_BotVersion__Type = "AWS::Lex::BotVersion"
)

var (
	// AWS_Lex_BotVersion__AttributesMap reports all the CloudFormation attributes for AWS::Lex::BotVersion.
	AWS_Lex_BotVersion__AttributesMap = struct {
		BotVersion string
	}{
		BotVersion: "BotVersion",
	}

	// AWS_Lex_BotVersion__AttributesSlice reports all the CloudFormation attributes for AWS::Lex::BotVersion.
	AWS_Lex_BotVersion__AttributesSlice = []string{
		AWS_Lex_BotVersion__AttributesMap.BotVersion,
	}
)

var (
	// AWS_Lex_BotVersion__PropertiesMap reports all the CloudFormation properties for AWS::Lex::BotVersion.
	AWS_Lex_BotVersion__PropertiesMap = struct {
		BotId                         string
		BotVersionLocaleSpecification string
		Description                   string
	}{
		BotId:                         "BotId",
		BotVersionLocaleSpecification: "BotVersionLocaleSpecification",
		Description:                   "Description",
	}

	// AWS_Lex_BotVersion__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::BotVersion.
	AWS_Lex_BotVersion__PropertiesSlice = []string{
		AWS_Lex_BotVersion__PropertiesMap.BotId,
		AWS_Lex_BotVersion__PropertiesMap.BotVersionLocaleSpecification,
		AWS_Lex_BotVersion__PropertiesMap.Description,
	}
)

// AWS_Lex_BotVersion is a binding for AWS::Lex::BotVersion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botversion.html
type AWS_Lex_BotVersion struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// BotId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botversion.html#cfn-lex-botversion-botid
	BotId cfz.Expression[string] `json:"BotId,omitempty"`

	// BotVersionLocaleSpecification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botversion.html#cfn-lex-botversion-botversionlocalespecification
	BotVersionLocaleSpecification cfz.ExpressionSlice[AWS_Lex_BotVersion_BotVersionLocaleSpecification] `json:"BotVersionLocaleSpecification,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-botversion.html#cfn-lex-botversion-description
	Description cfz.Expression[string] `json:"Description,omitempty"`
}

// New__AWS_Lex_BotVersion initializes a new *AWS_Lex_BotVersion.
func New__AWS_Lex_BotVersion(logicalName string) *AWS_Lex_BotVersion {
	return &AWS_Lex_BotVersion{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Lex_BotVersion) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Lex_BotVersion) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Lex_BotVersion) GetType() string {
	return AWS_Lex_BotVersion__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Lex_BotVersion) Set__LogicalName(v string) *AWS_Lex_BotVersion {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Lex_BotVersion) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Lex_BotVersion {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Lex_BotVersion) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Lex_BotVersion {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Lex_BotVersion) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Lex_BotVersion {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Lex_BotVersion) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Lex_BotVersion {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Lex_BotVersion) Set__RequestedOutputs(v []cfz.Output) *AWS_Lex_BotVersion {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Lex_BotVersion) Add__RequestedOutputs(v ...cfz.Output) *AWS_Lex_BotVersion {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__BotId updates property "BotId".
func (t *AWS_Lex_BotVersion) Set__BotId(v cfz.Expression[string]) *AWS_Lex_BotVersion {
	t.BotId = v
	return t
}

// SetV__BotId updates property "BotId".
func (t *AWS_Lex_BotVersion) SetV__BotId(v string) *AWS_Lex_BotVersion {
	t.BotId = cfz.V(v)
	return t
}

// Set__BotVersionLocaleSpecification updates property "BotVersionLocaleSpecification".
func (t *AWS_Lex_BotVersion) Set__BotVersionLocaleSpecification(v cfz.ExpressionSlice[AWS_Lex_BotVersion_BotVersionLocaleSpecification]) *AWS_Lex_BotVersion {
	t.BotVersionLocaleSpecification = v
	return t
}

// SetS__BotVersionLocaleSpecification updates property "BotVersionLocaleSpecification".
func (t *AWS_Lex_BotVersion) SetS__BotVersionLocaleSpecification(v ...cfz.Expression[AWS_Lex_BotVersion_BotVersionLocaleSpecification]) *AWS_Lex_BotVersion {
	t.BotVersionLocaleSpecification = cfz.S(v...)
	return t
}

// SetSV__BotVersionLocaleSpecification updates property "BotVersionLocaleSpecification".
func (t *AWS_Lex_BotVersion) SetSV__BotVersionLocaleSpecification(v ...AWS_Lex_BotVersion_BotVersionLocaleSpecification) *AWS_Lex_BotVersion {
	t.BotVersionLocaleSpecification = cfz.SV(v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Lex_BotVersion) Set__Description(v cfz.Expression[string]) *AWS_Lex_BotVersion {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Lex_BotVersion) SetV__Description(v string) *AWS_Lex_BotVersion {
	t.Description = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Lex_BotVersion) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__BotVersion returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: BotVersion
func (t *AWS_Lex_BotVersion) GetAtt__BotVersion() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Lex_BotVersion__AttributesMap.BotVersion))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Lex_BotVersion) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__BotVersion returns a conventionally configured output for an attribute of this resource.
// Attribute: BotVersion
func (t *AWS_Lex_BotVersion) GetConventionalOutputAtt__BotVersion(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttBotVersion", t.GetAtt__BotVersion())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Lex_BotVersion) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Lex_BotVersion

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

func (t *AWS_Lex_BotVersion) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
