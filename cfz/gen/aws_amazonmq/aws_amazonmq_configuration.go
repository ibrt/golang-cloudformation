// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_amazonmq

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_AmazonMQ_Configuration)(nil)
	_ cfz.Resource                   = (*AWS_AmazonMQ_Configuration)(nil)
)

const (
	// AWS_AmazonMQ_Configuration__Type is the CloudFormation type for AWS::AmazonMQ::Configuration.
	AWS_AmazonMQ_Configuration__Type = "AWS::AmazonMQ::Configuration"
)

var (
	// AWS_AmazonMQ_Configuration__AttributesMap reports all the CloudFormation attributes for AWS::AmazonMQ::Configuration.
	AWS_AmazonMQ_Configuration__AttributesMap = struct {
		Arn      string
		Id       string
		Revision string
	}{
		Arn:      "Arn",
		Id:       "Id",
		Revision: "Revision",
	}

	// AWS_AmazonMQ_Configuration__AttributesSlice reports all the CloudFormation attributes for AWS::AmazonMQ::Configuration.
	AWS_AmazonMQ_Configuration__AttributesSlice = []string{
		AWS_AmazonMQ_Configuration__AttributesMap.Arn,
		AWS_AmazonMQ_Configuration__AttributesMap.Id,
		AWS_AmazonMQ_Configuration__AttributesMap.Revision,
	}
)

var (
	// AWS_AmazonMQ_Configuration__PropertiesMap reports all the CloudFormation properties for AWS::AmazonMQ::Configuration.
	AWS_AmazonMQ_Configuration__PropertiesMap = struct {
		AuthenticationStrategy string
		Data                   string
		Description            string
		EngineType             string
		EngineVersion          string
		Name                   string
		Tags                   string
	}{
		AuthenticationStrategy: "AuthenticationStrategy",
		Data:                   "Data",
		Description:            "Description",
		EngineType:             "EngineType",
		EngineVersion:          "EngineVersion",
		Name:                   "Name",
		Tags:                   "Tags",
	}

	// AWS_AmazonMQ_Configuration__PropertiesSlice reports all the CloudFormation properties for AWS::AmazonMQ::Configuration.
	AWS_AmazonMQ_Configuration__PropertiesSlice = []string{
		AWS_AmazonMQ_Configuration__PropertiesMap.AuthenticationStrategy,
		AWS_AmazonMQ_Configuration__PropertiesMap.Data,
		AWS_AmazonMQ_Configuration__PropertiesMap.Description,
		AWS_AmazonMQ_Configuration__PropertiesMap.EngineType,
		AWS_AmazonMQ_Configuration__PropertiesMap.EngineVersion,
		AWS_AmazonMQ_Configuration__PropertiesMap.Name,
		AWS_AmazonMQ_Configuration__PropertiesMap.Tags,
	}
)

// AWS_AmazonMQ_Configuration is a binding for AWS::AmazonMQ::Configuration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html
type AWS_AmazonMQ_Configuration struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AuthenticationStrategy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-authenticationstrategy
	AuthenticationStrategy cfz.Expression[string] `json:"AuthenticationStrategy,omitempty"`

	// Data is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-data
	Data cfz.Expression[string] `json:"Data,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// EngineType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-enginetype
	EngineType cfz.Expression[string] `json:"EngineType,omitempty"`

	// EngineVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-engineversion
	EngineVersion cfz.Expression[string] `json:"EngineVersion,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html#cfn-amazonmq-configuration-tags
	Tags cfz.ExpressionSlice[AWS_AmazonMQ_Configuration_TagsEntry] `json:"Tags,omitempty"`
}

// New__AWS_AmazonMQ_Configuration initializes a new *AWS_AmazonMQ_Configuration.
func New__AWS_AmazonMQ_Configuration(logicalName string) *AWS_AmazonMQ_Configuration {
	return &AWS_AmazonMQ_Configuration{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_AmazonMQ_Configuration) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_AmazonMQ_Configuration) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_AmazonMQ_Configuration) GetType() string {
	return AWS_AmazonMQ_Configuration__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_AmazonMQ_Configuration) Set__LogicalName(v string) *AWS_AmazonMQ_Configuration {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_AmazonMQ_Configuration) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_AmazonMQ_Configuration {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_AmazonMQ_Configuration) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_AmazonMQ_Configuration {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_AmazonMQ_Configuration) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_AmazonMQ_Configuration {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_AmazonMQ_Configuration) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_AmazonMQ_Configuration {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_AmazonMQ_Configuration) Set__RequestedOutputs(v []cfz.Output) *AWS_AmazonMQ_Configuration {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_AmazonMQ_Configuration) Add__RequestedOutputs(v ...cfz.Output) *AWS_AmazonMQ_Configuration {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AuthenticationStrategy updates property "AuthenticationStrategy".
func (t *AWS_AmazonMQ_Configuration) Set__AuthenticationStrategy(v cfz.Expression[string]) *AWS_AmazonMQ_Configuration {
	t.AuthenticationStrategy = v
	return t
}

// SetV__AuthenticationStrategy updates property "AuthenticationStrategy".
func (t *AWS_AmazonMQ_Configuration) SetV__AuthenticationStrategy(v string) *AWS_AmazonMQ_Configuration {
	t.AuthenticationStrategy = cfz.V(v)
	return t
}

// Set__Data updates property "Data".
func (t *AWS_AmazonMQ_Configuration) Set__Data(v cfz.Expression[string]) *AWS_AmazonMQ_Configuration {
	t.Data = v
	return t
}

// SetV__Data updates property "Data".
func (t *AWS_AmazonMQ_Configuration) SetV__Data(v string) *AWS_AmazonMQ_Configuration {
	t.Data = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_AmazonMQ_Configuration) Set__Description(v cfz.Expression[string]) *AWS_AmazonMQ_Configuration {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_AmazonMQ_Configuration) SetV__Description(v string) *AWS_AmazonMQ_Configuration {
	t.Description = cfz.V(v)
	return t
}

// Set__EngineType updates property "EngineType".
func (t *AWS_AmazonMQ_Configuration) Set__EngineType(v cfz.Expression[string]) *AWS_AmazonMQ_Configuration {
	t.EngineType = v
	return t
}

// SetV__EngineType updates property "EngineType".
func (t *AWS_AmazonMQ_Configuration) SetV__EngineType(v string) *AWS_AmazonMQ_Configuration {
	t.EngineType = cfz.V(v)
	return t
}

// Set__EngineVersion updates property "EngineVersion".
func (t *AWS_AmazonMQ_Configuration) Set__EngineVersion(v cfz.Expression[string]) *AWS_AmazonMQ_Configuration {
	t.EngineVersion = v
	return t
}

// SetV__EngineVersion updates property "EngineVersion".
func (t *AWS_AmazonMQ_Configuration) SetV__EngineVersion(v string) *AWS_AmazonMQ_Configuration {
	t.EngineVersion = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_AmazonMQ_Configuration) Set__Name(v cfz.Expression[string]) *AWS_AmazonMQ_Configuration {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_AmazonMQ_Configuration) SetV__Name(v string) *AWS_AmazonMQ_Configuration {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_AmazonMQ_Configuration) Set__Tags(v cfz.ExpressionSlice[AWS_AmazonMQ_Configuration_TagsEntry]) *AWS_AmazonMQ_Configuration {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_AmazonMQ_Configuration) SetS__Tags(v ...cfz.Expression[AWS_AmazonMQ_Configuration_TagsEntry]) *AWS_AmazonMQ_Configuration {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_AmazonMQ_Configuration) SetSV__Tags(v ...AWS_AmazonMQ_Configuration_TagsEntry) *AWS_AmazonMQ_Configuration {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_AmazonMQ_Configuration) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_AmazonMQ_Configuration) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AmazonMQ_Configuration__AttributesMap.Arn))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_AmazonMQ_Configuration) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AmazonMQ_Configuration__AttributesMap.Id))
}

// GetAtt__Revision returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Revision
func (t *AWS_AmazonMQ_Configuration) GetAtt__Revision() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AmazonMQ_Configuration__AttributesMap.Revision))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_AmazonMQ_Configuration) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_AmazonMQ_Configuration) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_AmazonMQ_Configuration) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Revision returns a conventionally configured output for an attribute of this resource.
// Attribute: Revision
func (t *AWS_AmazonMQ_Configuration) GetConventionalOutputAtt__Revision(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttRevision", t.GetAtt__Revision())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_AmazonMQ_Configuration) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_AmazonMQ_Configuration

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

func (t *AWS_AmazonMQ_Configuration) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
