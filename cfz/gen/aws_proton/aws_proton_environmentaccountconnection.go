// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_proton

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Proton_EnvironmentAccountConnection)(nil)
	_ cfz.Resource                   = (*AWS_Proton_EnvironmentAccountConnection)(nil)
)

const (
	// AWS_Proton_EnvironmentAccountConnection__Type is the CloudFormation type for AWS::Proton::EnvironmentAccountConnection.
	AWS_Proton_EnvironmentAccountConnection__Type = "AWS::Proton::EnvironmentAccountConnection"
)

var (
	// AWS_Proton_EnvironmentAccountConnection__AttributesMap reports all the CloudFormation attributes for AWS::Proton::EnvironmentAccountConnection.
	AWS_Proton_EnvironmentAccountConnection__AttributesMap = struct {
		Arn    string
		Id     string
		Status string
	}{
		Arn:    "Arn",
		Id:     "Id",
		Status: "Status",
	}

	// AWS_Proton_EnvironmentAccountConnection__AttributesSlice reports all the CloudFormation attributes for AWS::Proton::EnvironmentAccountConnection.
	AWS_Proton_EnvironmentAccountConnection__AttributesSlice = []string{
		AWS_Proton_EnvironmentAccountConnection__AttributesMap.Arn,
		AWS_Proton_EnvironmentAccountConnection__AttributesMap.Id,
		AWS_Proton_EnvironmentAccountConnection__AttributesMap.Status,
	}
)

var (
	// AWS_Proton_EnvironmentAccountConnection__PropertiesMap reports all the CloudFormation properties for AWS::Proton::EnvironmentAccountConnection.
	AWS_Proton_EnvironmentAccountConnection__PropertiesMap = struct {
		CodebuildRoleArn     string
		ComponentRoleArn     string
		EnvironmentAccountId string
		EnvironmentName      string
		ManagementAccountId  string
		RoleArn              string
		Tags                 string
	}{
		CodebuildRoleArn:     "CodebuildRoleArn",
		ComponentRoleArn:     "ComponentRoleArn",
		EnvironmentAccountId: "EnvironmentAccountId",
		EnvironmentName:      "EnvironmentName",
		ManagementAccountId:  "ManagementAccountId",
		RoleArn:              "RoleArn",
		Tags:                 "Tags",
	}

	// AWS_Proton_EnvironmentAccountConnection__PropertiesSlice reports all the CloudFormation properties for AWS::Proton::EnvironmentAccountConnection.
	AWS_Proton_EnvironmentAccountConnection__PropertiesSlice = []string{
		AWS_Proton_EnvironmentAccountConnection__PropertiesMap.CodebuildRoleArn,
		AWS_Proton_EnvironmentAccountConnection__PropertiesMap.ComponentRoleArn,
		AWS_Proton_EnvironmentAccountConnection__PropertiesMap.EnvironmentAccountId,
		AWS_Proton_EnvironmentAccountConnection__PropertiesMap.EnvironmentName,
		AWS_Proton_EnvironmentAccountConnection__PropertiesMap.ManagementAccountId,
		AWS_Proton_EnvironmentAccountConnection__PropertiesMap.RoleArn,
		AWS_Proton_EnvironmentAccountConnection__PropertiesMap.Tags,
	}
)

// AWS_Proton_EnvironmentAccountConnection is a binding for AWS::Proton::EnvironmentAccountConnection.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmentaccountconnection.html
type AWS_Proton_EnvironmentAccountConnection struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CodebuildRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmentaccountconnection.html#cfn-proton-environmentaccountconnection-codebuildrolearn
	CodebuildRoleArn cfz.Expression[string] `json:"CodebuildRoleArn,omitempty"`

	// ComponentRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmentaccountconnection.html#cfn-proton-environmentaccountconnection-componentrolearn
	ComponentRoleArn cfz.Expression[string] `json:"ComponentRoleArn,omitempty"`

	// EnvironmentAccountId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmentaccountconnection.html#cfn-proton-environmentaccountconnection-environmentaccountid
	EnvironmentAccountId cfz.Expression[string] `json:"EnvironmentAccountId,omitempty"`

	// EnvironmentName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmentaccountconnection.html#cfn-proton-environmentaccountconnection-environmentname
	EnvironmentName cfz.Expression[string] `json:"EnvironmentName,omitempty"`

	// ManagementAccountId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmentaccountconnection.html#cfn-proton-environmentaccountconnection-managementaccountid
	ManagementAccountId cfz.Expression[string] `json:"ManagementAccountId,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmentaccountconnection.html#cfn-proton-environmentaccountconnection-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmentaccountconnection.html#cfn-proton-environmentaccountconnection-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_Proton_EnvironmentAccountConnection initializes a new *AWS_Proton_EnvironmentAccountConnection.
func New__AWS_Proton_EnvironmentAccountConnection(logicalName string) *AWS_Proton_EnvironmentAccountConnection {
	return &AWS_Proton_EnvironmentAccountConnection{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Proton_EnvironmentAccountConnection) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Proton_EnvironmentAccountConnection) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Proton_EnvironmentAccountConnection) GetType() string {
	return AWS_Proton_EnvironmentAccountConnection__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__LogicalName(v string) *AWS_Proton_EnvironmentAccountConnection {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Proton_EnvironmentAccountConnection {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Proton_EnvironmentAccountConnection) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Proton_EnvironmentAccountConnection {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Proton_EnvironmentAccountConnection {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Proton_EnvironmentAccountConnection {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__RequestedOutputs(v []cfz.Output) *AWS_Proton_EnvironmentAccountConnection {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Proton_EnvironmentAccountConnection) Add__RequestedOutputs(v ...cfz.Output) *AWS_Proton_EnvironmentAccountConnection {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CodebuildRoleArn updates property "CodebuildRoleArn".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__CodebuildRoleArn(v cfz.Expression[string]) *AWS_Proton_EnvironmentAccountConnection {
	t.CodebuildRoleArn = v
	return t
}

// SetV__CodebuildRoleArn updates property "CodebuildRoleArn".
func (t *AWS_Proton_EnvironmentAccountConnection) SetV__CodebuildRoleArn(v string) *AWS_Proton_EnvironmentAccountConnection {
	t.CodebuildRoleArn = cfz.V(v)
	return t
}

// Set__ComponentRoleArn updates property "ComponentRoleArn".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__ComponentRoleArn(v cfz.Expression[string]) *AWS_Proton_EnvironmentAccountConnection {
	t.ComponentRoleArn = v
	return t
}

// SetV__ComponentRoleArn updates property "ComponentRoleArn".
func (t *AWS_Proton_EnvironmentAccountConnection) SetV__ComponentRoleArn(v string) *AWS_Proton_EnvironmentAccountConnection {
	t.ComponentRoleArn = cfz.V(v)
	return t
}

// Set__EnvironmentAccountId updates property "EnvironmentAccountId".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__EnvironmentAccountId(v cfz.Expression[string]) *AWS_Proton_EnvironmentAccountConnection {
	t.EnvironmentAccountId = v
	return t
}

// SetV__EnvironmentAccountId updates property "EnvironmentAccountId".
func (t *AWS_Proton_EnvironmentAccountConnection) SetV__EnvironmentAccountId(v string) *AWS_Proton_EnvironmentAccountConnection {
	t.EnvironmentAccountId = cfz.V(v)
	return t
}

// Set__EnvironmentName updates property "EnvironmentName".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__EnvironmentName(v cfz.Expression[string]) *AWS_Proton_EnvironmentAccountConnection {
	t.EnvironmentName = v
	return t
}

// SetV__EnvironmentName updates property "EnvironmentName".
func (t *AWS_Proton_EnvironmentAccountConnection) SetV__EnvironmentName(v string) *AWS_Proton_EnvironmentAccountConnection {
	t.EnvironmentName = cfz.V(v)
	return t
}

// Set__ManagementAccountId updates property "ManagementAccountId".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__ManagementAccountId(v cfz.Expression[string]) *AWS_Proton_EnvironmentAccountConnection {
	t.ManagementAccountId = v
	return t
}

// SetV__ManagementAccountId updates property "ManagementAccountId".
func (t *AWS_Proton_EnvironmentAccountConnection) SetV__ManagementAccountId(v string) *AWS_Proton_EnvironmentAccountConnection {
	t.ManagementAccountId = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__RoleArn(v cfz.Expression[string]) *AWS_Proton_EnvironmentAccountConnection {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t *AWS_Proton_EnvironmentAccountConnection) SetV__RoleArn(v string) *AWS_Proton_EnvironmentAccountConnection {
	t.RoleArn = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Proton_EnvironmentAccountConnection) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Proton_EnvironmentAccountConnection {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Proton_EnvironmentAccountConnection) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Proton_EnvironmentAccountConnection {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Proton_EnvironmentAccountConnection) SetSV__Tags(v ...cfz.Tag) *AWS_Proton_EnvironmentAccountConnection {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Proton_EnvironmentAccountConnection) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Proton_EnvironmentAccountConnection) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Proton_EnvironmentAccountConnection__AttributesMap.Arn))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_Proton_EnvironmentAccountConnection) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Proton_EnvironmentAccountConnection__AttributesMap.Id))
}

// GetAtt__Status returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Status
func (t *AWS_Proton_EnvironmentAccountConnection) GetAtt__Status() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Proton_EnvironmentAccountConnection__AttributesMap.Status))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Proton_EnvironmentAccountConnection) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Proton_EnvironmentAccountConnection) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_Proton_EnvironmentAccountConnection) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Status returns a conventionally configured output for an attribute of this resource.
// Attribute: Status
func (t *AWS_Proton_EnvironmentAccountConnection) GetConventionalOutputAtt__Status(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatus", t.GetAtt__Status())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Proton_EnvironmentAccountConnection) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Proton_EnvironmentAccountConnection

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

func (t *AWS_Proton_EnvironmentAccountConnection) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
