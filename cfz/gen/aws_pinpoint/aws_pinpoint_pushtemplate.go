// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pinpoint

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Pinpoint_PushTemplate)(nil)
	_ cfz.Resource                   = (*AWS_Pinpoint_PushTemplate)(nil)
)

const (
	// AWS_Pinpoint_PushTemplate__Type is the CloudFormation type for AWS::Pinpoint::PushTemplate.
	AWS_Pinpoint_PushTemplate__Type = "AWS::Pinpoint::PushTemplate"
)

var (
	// AWS_Pinpoint_PushTemplate__AttributesMap reports all the CloudFormation attributes for AWS::Pinpoint::PushTemplate.
	AWS_Pinpoint_PushTemplate__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Pinpoint_PushTemplate__AttributesSlice reports all the CloudFormation attributes for AWS::Pinpoint::PushTemplate.
	AWS_Pinpoint_PushTemplate__AttributesSlice = []string{
		AWS_Pinpoint_PushTemplate__AttributesMap.Arn,
	}
)

var (
	// AWS_Pinpoint_PushTemplate__PropertiesMap reports all the CloudFormation properties for AWS::Pinpoint::PushTemplate.
	AWS_Pinpoint_PushTemplate__PropertiesMap = struct {
		ADM                  string
		APNS                 string
		Baidu                string
		Default              string
		DefaultSubstitutions string
		GCM                  string
		Tags                 string
		TemplateDescription  string
		TemplateName         string
	}{
		ADM:                  "ADM",
		APNS:                 "APNS",
		Baidu:                "Baidu",
		Default:              "Default",
		DefaultSubstitutions: "DefaultSubstitutions",
		GCM:                  "GCM",
		Tags:                 "Tags",
		TemplateDescription:  "TemplateDescription",
		TemplateName:         "TemplateName",
	}

	// AWS_Pinpoint_PushTemplate__PropertiesSlice reports all the CloudFormation properties for AWS::Pinpoint::PushTemplate.
	AWS_Pinpoint_PushTemplate__PropertiesSlice = []string{
		AWS_Pinpoint_PushTemplate__PropertiesMap.ADM,
		AWS_Pinpoint_PushTemplate__PropertiesMap.APNS,
		AWS_Pinpoint_PushTemplate__PropertiesMap.Baidu,
		AWS_Pinpoint_PushTemplate__PropertiesMap.Default,
		AWS_Pinpoint_PushTemplate__PropertiesMap.DefaultSubstitutions,
		AWS_Pinpoint_PushTemplate__PropertiesMap.GCM,
		AWS_Pinpoint_PushTemplate__PropertiesMap.Tags,
		AWS_Pinpoint_PushTemplate__PropertiesMap.TemplateDescription,
		AWS_Pinpoint_PushTemplate__PropertiesMap.TemplateName,
	}
)

// AWS_Pinpoint_PushTemplate is a binding for AWS::Pinpoint::PushTemplate.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html
type AWS_Pinpoint_PushTemplate struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ADM is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-adm
	ADM cfz.Expression[AWS_Pinpoint_PushTemplate_AndroidPushNotificationTemplate] `json:"ADM,omitempty"`

	// APNS is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-apns
	APNS cfz.Expression[AWS_Pinpoint_PushTemplate_APNSPushNotificationTemplate] `json:"APNS,omitempty"`

	// Baidu is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-baidu
	Baidu cfz.Expression[AWS_Pinpoint_PushTemplate_AndroidPushNotificationTemplate] `json:"Baidu,omitempty"`

	// Default is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-default
	Default cfz.Expression[AWS_Pinpoint_PushTemplate_DefaultPushNotificationTemplate] `json:"Default,omitempty"`

	// DefaultSubstitutions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-defaultsubstitutions
	DefaultSubstitutions cfz.Expression[string] `json:"DefaultSubstitutions,omitempty"`

	// GCM is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-gcm
	GCM cfz.Expression[AWS_Pinpoint_PushTemplate_AndroidPushNotificationTemplate] `json:"GCM,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-tags
	Tags cfz.Expression[json.RawMessage] `json:"Tags,omitempty"`

	// TemplateDescription is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatedescription
	TemplateDescription cfz.Expression[string] `json:"TemplateDescription,omitempty"`

	// TemplateName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-pushtemplate.html#cfn-pinpoint-pushtemplate-templatename
	TemplateName cfz.Expression[string] `json:"TemplateName,omitempty"`
}

// New__AWS_Pinpoint_PushTemplate initializes a new *AWS_Pinpoint_PushTemplate.
func New__AWS_Pinpoint_PushTemplate(logicalName string) *AWS_Pinpoint_PushTemplate {
	return &AWS_Pinpoint_PushTemplate{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Pinpoint_PushTemplate) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Pinpoint_PushTemplate) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Pinpoint_PushTemplate) GetType() string {
	return AWS_Pinpoint_PushTemplate__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Pinpoint_PushTemplate) Set__LogicalName(v string) *AWS_Pinpoint_PushTemplate {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Pinpoint_PushTemplate) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Pinpoint_PushTemplate {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Pinpoint_PushTemplate) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Pinpoint_PushTemplate {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Pinpoint_PushTemplate) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Pinpoint_PushTemplate {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Pinpoint_PushTemplate) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Pinpoint_PushTemplate {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Pinpoint_PushTemplate) Set__RequestedOutputs(v []cfz.Output) *AWS_Pinpoint_PushTemplate {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Pinpoint_PushTemplate) Add__RequestedOutputs(v ...cfz.Output) *AWS_Pinpoint_PushTemplate {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ADM updates property "ADM".
func (t *AWS_Pinpoint_PushTemplate) Set__ADM(v cfz.Expression[AWS_Pinpoint_PushTemplate_AndroidPushNotificationTemplate]) *AWS_Pinpoint_PushTemplate {
	t.ADM = v
	return t
}

// SetV__ADM updates property "ADM".
func (t *AWS_Pinpoint_PushTemplate) SetV__ADM(v AWS_Pinpoint_PushTemplate_AndroidPushNotificationTemplate) *AWS_Pinpoint_PushTemplate {
	t.ADM = cfz.V(v)
	return t
}

// Set__APNS updates property "APNS".
func (t *AWS_Pinpoint_PushTemplate) Set__APNS(v cfz.Expression[AWS_Pinpoint_PushTemplate_APNSPushNotificationTemplate]) *AWS_Pinpoint_PushTemplate {
	t.APNS = v
	return t
}

// SetV__APNS updates property "APNS".
func (t *AWS_Pinpoint_PushTemplate) SetV__APNS(v AWS_Pinpoint_PushTemplate_APNSPushNotificationTemplate) *AWS_Pinpoint_PushTemplate {
	t.APNS = cfz.V(v)
	return t
}

// Set__Baidu updates property "Baidu".
func (t *AWS_Pinpoint_PushTemplate) Set__Baidu(v cfz.Expression[AWS_Pinpoint_PushTemplate_AndroidPushNotificationTemplate]) *AWS_Pinpoint_PushTemplate {
	t.Baidu = v
	return t
}

// SetV__Baidu updates property "Baidu".
func (t *AWS_Pinpoint_PushTemplate) SetV__Baidu(v AWS_Pinpoint_PushTemplate_AndroidPushNotificationTemplate) *AWS_Pinpoint_PushTemplate {
	t.Baidu = cfz.V(v)
	return t
}

// Set__Default updates property "Default".
func (t *AWS_Pinpoint_PushTemplate) Set__Default(v cfz.Expression[AWS_Pinpoint_PushTemplate_DefaultPushNotificationTemplate]) *AWS_Pinpoint_PushTemplate {
	t.Default = v
	return t
}

// SetV__Default updates property "Default".
func (t *AWS_Pinpoint_PushTemplate) SetV__Default(v AWS_Pinpoint_PushTemplate_DefaultPushNotificationTemplate) *AWS_Pinpoint_PushTemplate {
	t.Default = cfz.V(v)
	return t
}

// Set__DefaultSubstitutions updates property "DefaultSubstitutions".
func (t *AWS_Pinpoint_PushTemplate) Set__DefaultSubstitutions(v cfz.Expression[string]) *AWS_Pinpoint_PushTemplate {
	t.DefaultSubstitutions = v
	return t
}

// SetV__DefaultSubstitutions updates property "DefaultSubstitutions".
func (t *AWS_Pinpoint_PushTemplate) SetV__DefaultSubstitutions(v string) *AWS_Pinpoint_PushTemplate {
	t.DefaultSubstitutions = cfz.V(v)
	return t
}

// Set__GCM updates property "GCM".
func (t *AWS_Pinpoint_PushTemplate) Set__GCM(v cfz.Expression[AWS_Pinpoint_PushTemplate_AndroidPushNotificationTemplate]) *AWS_Pinpoint_PushTemplate {
	t.GCM = v
	return t
}

// SetV__GCM updates property "GCM".
func (t *AWS_Pinpoint_PushTemplate) SetV__GCM(v AWS_Pinpoint_PushTemplate_AndroidPushNotificationTemplate) *AWS_Pinpoint_PushTemplate {
	t.GCM = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Pinpoint_PushTemplate) Set__Tags(v cfz.Expression[json.RawMessage]) *AWS_Pinpoint_PushTemplate {
	t.Tags = v
	return t
}

// SetV__Tags updates property "Tags".
func (t *AWS_Pinpoint_PushTemplate) SetV__Tags(v json.RawMessage) *AWS_Pinpoint_PushTemplate {
	t.Tags = cfz.V(v)
	return t
}

// Set__TemplateDescription updates property "TemplateDescription".
func (t *AWS_Pinpoint_PushTemplate) Set__TemplateDescription(v cfz.Expression[string]) *AWS_Pinpoint_PushTemplate {
	t.TemplateDescription = v
	return t
}

// SetV__TemplateDescription updates property "TemplateDescription".
func (t *AWS_Pinpoint_PushTemplate) SetV__TemplateDescription(v string) *AWS_Pinpoint_PushTemplate {
	t.TemplateDescription = cfz.V(v)
	return t
}

// Set__TemplateName updates property "TemplateName".
func (t *AWS_Pinpoint_PushTemplate) Set__TemplateName(v cfz.Expression[string]) *AWS_Pinpoint_PushTemplate {
	t.TemplateName = v
	return t
}

// SetV__TemplateName updates property "TemplateName".
func (t *AWS_Pinpoint_PushTemplate) SetV__TemplateName(v string) *AWS_Pinpoint_PushTemplate {
	t.TemplateName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Pinpoint_PushTemplate) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Pinpoint_PushTemplate) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Pinpoint_PushTemplate__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Pinpoint_PushTemplate) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Pinpoint_PushTemplate) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Pinpoint_PushTemplate) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Pinpoint_PushTemplate

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

func (t *AWS_Pinpoint_PushTemplate) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
