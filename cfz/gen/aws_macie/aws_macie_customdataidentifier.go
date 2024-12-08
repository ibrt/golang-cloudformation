// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_macie

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Macie_CustomDataIdentifier)(nil)
	_ cfz.Resource                   = (*AWS_Macie_CustomDataIdentifier)(nil)
)

const (
	// AWS_Macie_CustomDataIdentifier__Type is the CloudFormation type for AWS::Macie::CustomDataIdentifier.
	AWS_Macie_CustomDataIdentifier__Type = "AWS::Macie::CustomDataIdentifier"
)

var (
	// AWS_Macie_CustomDataIdentifier__AttributesMap reports all the CloudFormation attributes for AWS::Macie::CustomDataIdentifier.
	AWS_Macie_CustomDataIdentifier__AttributesMap = struct {
		Arn string
		Id  string
	}{
		Arn: "Arn",
		Id:  "Id",
	}

	// AWS_Macie_CustomDataIdentifier__AttributesSlice reports all the CloudFormation attributes for AWS::Macie::CustomDataIdentifier.
	AWS_Macie_CustomDataIdentifier__AttributesSlice = []string{
		AWS_Macie_CustomDataIdentifier__AttributesMap.Arn,
		AWS_Macie_CustomDataIdentifier__AttributesMap.Id,
	}
)

var (
	// AWS_Macie_CustomDataIdentifier__PropertiesMap reports all the CloudFormation properties for AWS::Macie::CustomDataIdentifier.
	AWS_Macie_CustomDataIdentifier__PropertiesMap = struct {
		Description          string
		IgnoreWords          string
		Keywords             string
		MaximumMatchDistance string
		Name                 string
		Regex                string
		Tags                 string
	}{
		Description:          "Description",
		IgnoreWords:          "IgnoreWords",
		Keywords:             "Keywords",
		MaximumMatchDistance: "MaximumMatchDistance",
		Name:                 "Name",
		Regex:                "Regex",
		Tags:                 "Tags",
	}

	// AWS_Macie_CustomDataIdentifier__PropertiesSlice reports all the CloudFormation properties for AWS::Macie::CustomDataIdentifier.
	AWS_Macie_CustomDataIdentifier__PropertiesSlice = []string{
		AWS_Macie_CustomDataIdentifier__PropertiesMap.Description,
		AWS_Macie_CustomDataIdentifier__PropertiesMap.IgnoreWords,
		AWS_Macie_CustomDataIdentifier__PropertiesMap.Keywords,
		AWS_Macie_CustomDataIdentifier__PropertiesMap.MaximumMatchDistance,
		AWS_Macie_CustomDataIdentifier__PropertiesMap.Name,
		AWS_Macie_CustomDataIdentifier__PropertiesMap.Regex,
		AWS_Macie_CustomDataIdentifier__PropertiesMap.Tags,
	}
)

// AWS_Macie_CustomDataIdentifier is a binding for AWS::Macie::CustomDataIdentifier.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html
type AWS_Macie_CustomDataIdentifier struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// IgnoreWords is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-ignorewords
	IgnoreWords cfz.ExpressionSlice[string] `json:"IgnoreWords,omitempty"`

	// Keywords is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-keywords
	Keywords cfz.ExpressionSlice[string] `json:"Keywords,omitempty"`

	// MaximumMatchDistance is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-maximummatchdistance
	MaximumMatchDistance cfz.Expression[int32] `json:"MaximumMatchDistance,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Regex is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-regex
	Regex cfz.Expression[string] `json:"Regex,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_Macie_CustomDataIdentifier initializes a new *AWS_Macie_CustomDataIdentifier.
func New__AWS_Macie_CustomDataIdentifier(logicalName string) *AWS_Macie_CustomDataIdentifier {
	return &AWS_Macie_CustomDataIdentifier{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Macie_CustomDataIdentifier) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Macie_CustomDataIdentifier) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Macie_CustomDataIdentifier) GetType() string {
	return AWS_Macie_CustomDataIdentifier__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Macie_CustomDataIdentifier) Set__LogicalName(v string) *AWS_Macie_CustomDataIdentifier {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Macie_CustomDataIdentifier) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Macie_CustomDataIdentifier {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Macie_CustomDataIdentifier) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Macie_CustomDataIdentifier {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Macie_CustomDataIdentifier) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Macie_CustomDataIdentifier {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Macie_CustomDataIdentifier) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Macie_CustomDataIdentifier {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Macie_CustomDataIdentifier) Set__RequestedOutputs(v []cfz.Output) *AWS_Macie_CustomDataIdentifier {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Macie_CustomDataIdentifier) Add__RequestedOutputs(v ...cfz.Output) *AWS_Macie_CustomDataIdentifier {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Macie_CustomDataIdentifier) Set__Description(v cfz.Expression[string]) *AWS_Macie_CustomDataIdentifier {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Macie_CustomDataIdentifier) SetV__Description(v string) *AWS_Macie_CustomDataIdentifier {
	t.Description = cfz.V(v)
	return t
}

// Set__IgnoreWords updates property "IgnoreWords".
func (t *AWS_Macie_CustomDataIdentifier) Set__IgnoreWords(v cfz.ExpressionSlice[string]) *AWS_Macie_CustomDataIdentifier {
	t.IgnoreWords = v
	return t
}

// SetS__IgnoreWords updates property "IgnoreWords".
func (t *AWS_Macie_CustomDataIdentifier) SetS__IgnoreWords(v ...cfz.Expression[string]) *AWS_Macie_CustomDataIdentifier {
	t.IgnoreWords = cfz.S(v...)
	return t
}

// SetSV__IgnoreWords updates property "IgnoreWords".
func (t *AWS_Macie_CustomDataIdentifier) SetSV__IgnoreWords(v ...string) *AWS_Macie_CustomDataIdentifier {
	t.IgnoreWords = cfz.SV(v...)
	return t
}

// Set__Keywords updates property "Keywords".
func (t *AWS_Macie_CustomDataIdentifier) Set__Keywords(v cfz.ExpressionSlice[string]) *AWS_Macie_CustomDataIdentifier {
	t.Keywords = v
	return t
}

// SetS__Keywords updates property "Keywords".
func (t *AWS_Macie_CustomDataIdentifier) SetS__Keywords(v ...cfz.Expression[string]) *AWS_Macie_CustomDataIdentifier {
	t.Keywords = cfz.S(v...)
	return t
}

// SetSV__Keywords updates property "Keywords".
func (t *AWS_Macie_CustomDataIdentifier) SetSV__Keywords(v ...string) *AWS_Macie_CustomDataIdentifier {
	t.Keywords = cfz.SV(v...)
	return t
}

// Set__MaximumMatchDistance updates property "MaximumMatchDistance".
func (t *AWS_Macie_CustomDataIdentifier) Set__MaximumMatchDistance(v cfz.Expression[int32]) *AWS_Macie_CustomDataIdentifier {
	t.MaximumMatchDistance = v
	return t
}

// SetV__MaximumMatchDistance updates property "MaximumMatchDistance".
func (t *AWS_Macie_CustomDataIdentifier) SetV__MaximumMatchDistance(v int32) *AWS_Macie_CustomDataIdentifier {
	t.MaximumMatchDistance = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Macie_CustomDataIdentifier) Set__Name(v cfz.Expression[string]) *AWS_Macie_CustomDataIdentifier {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Macie_CustomDataIdentifier) SetV__Name(v string) *AWS_Macie_CustomDataIdentifier {
	t.Name = cfz.V(v)
	return t
}

// Set__Regex updates property "Regex".
func (t *AWS_Macie_CustomDataIdentifier) Set__Regex(v cfz.Expression[string]) *AWS_Macie_CustomDataIdentifier {
	t.Regex = v
	return t
}

// SetV__Regex updates property "Regex".
func (t *AWS_Macie_CustomDataIdentifier) SetV__Regex(v string) *AWS_Macie_CustomDataIdentifier {
	t.Regex = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Macie_CustomDataIdentifier) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Macie_CustomDataIdentifier {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Macie_CustomDataIdentifier) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Macie_CustomDataIdentifier {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Macie_CustomDataIdentifier) SetSV__Tags(v ...cfz.Tag) *AWS_Macie_CustomDataIdentifier {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Macie_CustomDataIdentifier) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Macie_CustomDataIdentifier) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Macie_CustomDataIdentifier__AttributesMap.Arn))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_Macie_CustomDataIdentifier) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Macie_CustomDataIdentifier__AttributesMap.Id))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Macie_CustomDataIdentifier) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Macie_CustomDataIdentifier) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_Macie_CustomDataIdentifier) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Macie_CustomDataIdentifier) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Macie_CustomDataIdentifier

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

func (t *AWS_Macie_CustomDataIdentifier) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
