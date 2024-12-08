// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediatailor

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_MediaTailor_VodSource)(nil)
	_ cfz.Resource                   = (*AWS_MediaTailor_VodSource)(nil)
)

const (
	// AWS_MediaTailor_VodSource__Type is the CloudFormation type for AWS::MediaTailor::VodSource.
	AWS_MediaTailor_VodSource__Type = "AWS::MediaTailor::VodSource"
)

var (
	// AWS_MediaTailor_VodSource__AttributesMap reports all the CloudFormation attributes for AWS::MediaTailor::VodSource.
	AWS_MediaTailor_VodSource__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_MediaTailor_VodSource__AttributesSlice reports all the CloudFormation attributes for AWS::MediaTailor::VodSource.
	AWS_MediaTailor_VodSource__AttributesSlice = []string{
		AWS_MediaTailor_VodSource__AttributesMap.Arn,
	}
)

var (
	// AWS_MediaTailor_VodSource__PropertiesMap reports all the CloudFormation properties for AWS::MediaTailor::VodSource.
	AWS_MediaTailor_VodSource__PropertiesMap = struct {
		HttpPackageConfigurations string
		SourceLocationName        string
		Tags                      string
		VodSourceName             string
	}{
		HttpPackageConfigurations: "HttpPackageConfigurations",
		SourceLocationName:        "SourceLocationName",
		Tags:                      "Tags",
		VodSourceName:             "VodSourceName",
	}

	// AWS_MediaTailor_VodSource__PropertiesSlice reports all the CloudFormation properties for AWS::MediaTailor::VodSource.
	AWS_MediaTailor_VodSource__PropertiesSlice = []string{
		AWS_MediaTailor_VodSource__PropertiesMap.HttpPackageConfigurations,
		AWS_MediaTailor_VodSource__PropertiesMap.SourceLocationName,
		AWS_MediaTailor_VodSource__PropertiesMap.Tags,
		AWS_MediaTailor_VodSource__PropertiesMap.VodSourceName,
	}
)

// AWS_MediaTailor_VodSource is a binding for AWS::MediaTailor::VodSource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html
type AWS_MediaTailor_VodSource struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// HttpPackageConfigurations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html#cfn-mediatailor-vodsource-httppackageconfigurations
	HttpPackageConfigurations cfz.ExpressionSlice[AWS_MediaTailor_VodSource_HttpPackageConfiguration] `json:"HttpPackageConfigurations,omitempty"`

	// SourceLocationName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html#cfn-mediatailor-vodsource-sourcelocationname
	SourceLocationName cfz.Expression[string] `json:"SourceLocationName,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html#cfn-mediatailor-vodsource-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// VodSourceName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html#cfn-mediatailor-vodsource-vodsourcename
	VodSourceName cfz.Expression[string] `json:"VodSourceName,omitempty"`
}

// New__AWS_MediaTailor_VodSource initializes a new *AWS_MediaTailor_VodSource.
func New__AWS_MediaTailor_VodSource(logicalName string) *AWS_MediaTailor_VodSource {
	return &AWS_MediaTailor_VodSource{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_MediaTailor_VodSource) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_MediaTailor_VodSource) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_MediaTailor_VodSource) GetType() string {
	return AWS_MediaTailor_VodSource__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_MediaTailor_VodSource) Set__LogicalName(v string) *AWS_MediaTailor_VodSource {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_MediaTailor_VodSource) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_MediaTailor_VodSource {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_MediaTailor_VodSource) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_MediaTailor_VodSource {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_MediaTailor_VodSource) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_MediaTailor_VodSource {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_MediaTailor_VodSource) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_MediaTailor_VodSource {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_MediaTailor_VodSource) Set__RequestedOutputs(v []cfz.Output) *AWS_MediaTailor_VodSource {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_MediaTailor_VodSource) Add__RequestedOutputs(v ...cfz.Output) *AWS_MediaTailor_VodSource {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__HttpPackageConfigurations updates property "HttpPackageConfigurations".
func (t *AWS_MediaTailor_VodSource) Set__HttpPackageConfigurations(v cfz.ExpressionSlice[AWS_MediaTailor_VodSource_HttpPackageConfiguration]) *AWS_MediaTailor_VodSource {
	t.HttpPackageConfigurations = v
	return t
}

// SetS__HttpPackageConfigurations updates property "HttpPackageConfigurations".
func (t *AWS_MediaTailor_VodSource) SetS__HttpPackageConfigurations(v ...cfz.Expression[AWS_MediaTailor_VodSource_HttpPackageConfiguration]) *AWS_MediaTailor_VodSource {
	t.HttpPackageConfigurations = cfz.S(v...)
	return t
}

// SetSV__HttpPackageConfigurations updates property "HttpPackageConfigurations".
func (t *AWS_MediaTailor_VodSource) SetSV__HttpPackageConfigurations(v ...AWS_MediaTailor_VodSource_HttpPackageConfiguration) *AWS_MediaTailor_VodSource {
	t.HttpPackageConfigurations = cfz.SV(v...)
	return t
}

// Set__SourceLocationName updates property "SourceLocationName".
func (t *AWS_MediaTailor_VodSource) Set__SourceLocationName(v cfz.Expression[string]) *AWS_MediaTailor_VodSource {
	t.SourceLocationName = v
	return t
}

// SetV__SourceLocationName updates property "SourceLocationName".
func (t *AWS_MediaTailor_VodSource) SetV__SourceLocationName(v string) *AWS_MediaTailor_VodSource {
	t.SourceLocationName = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_MediaTailor_VodSource) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_MediaTailor_VodSource {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_MediaTailor_VodSource) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_MediaTailor_VodSource {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_MediaTailor_VodSource) SetSV__Tags(v ...cfz.Tag) *AWS_MediaTailor_VodSource {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__VodSourceName updates property "VodSourceName".
func (t *AWS_MediaTailor_VodSource) Set__VodSourceName(v cfz.Expression[string]) *AWS_MediaTailor_VodSource {
	t.VodSourceName = v
	return t
}

// SetV__VodSourceName updates property "VodSourceName".
func (t *AWS_MediaTailor_VodSource) SetV__VodSourceName(v string) *AWS_MediaTailor_VodSource {
	t.VodSourceName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_MediaTailor_VodSource) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_MediaTailor_VodSource) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MediaTailor_VodSource__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_MediaTailor_VodSource) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_MediaTailor_VodSource) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_MediaTailor_VodSource) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_MediaTailor_VodSource

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

func (t *AWS_MediaTailor_VodSource) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
