// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_gamelift

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_GameLift_Alias)(nil)
	_ cfz.Resource                   = (*AWS_GameLift_Alias)(nil)
)

const (
	// AWS_GameLift_Alias__Type is the CloudFormation type for AWS::GameLift::Alias.
	AWS_GameLift_Alias__Type = "AWS::GameLift::Alias"
)

var (
	// AWS_GameLift_Alias__AttributesMap reports all the CloudFormation attributes for AWS::GameLift::Alias.
	AWS_GameLift_Alias__AttributesMap = struct {
		AliasId string
	}{
		AliasId: "AliasId",
	}

	// AWS_GameLift_Alias__AttributesSlice reports all the CloudFormation attributes for AWS::GameLift::Alias.
	AWS_GameLift_Alias__AttributesSlice = []string{
		AWS_GameLift_Alias__AttributesMap.AliasId,
	}
)

var (
	// AWS_GameLift_Alias__PropertiesMap reports all the CloudFormation properties for AWS::GameLift::Alias.
	AWS_GameLift_Alias__PropertiesMap = struct {
		Description     string
		Name            string
		RoutingStrategy string
	}{
		Description:     "Description",
		Name:            "Name",
		RoutingStrategy: "RoutingStrategy",
	}

	// AWS_GameLift_Alias__PropertiesSlice reports all the CloudFormation properties for AWS::GameLift::Alias.
	AWS_GameLift_Alias__PropertiesSlice = []string{
		AWS_GameLift_Alias__PropertiesMap.Description,
		AWS_GameLift_Alias__PropertiesMap.Name,
		AWS_GameLift_Alias__PropertiesMap.RoutingStrategy,
	}
)

// AWS_GameLift_Alias is a binding for AWS::GameLift::Alias.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html
type AWS_GameLift_Alias struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// RoutingStrategy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-routingstrategy
	RoutingStrategy cfz.Expression[AWS_GameLift_Alias_RoutingStrategy] `json:"RoutingStrategy,omitempty"`
}

// New__AWS_GameLift_Alias initializes a new *AWS_GameLift_Alias.
func New__AWS_GameLift_Alias(logicalName string) *AWS_GameLift_Alias {
	return &AWS_GameLift_Alias{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_GameLift_Alias) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_GameLift_Alias) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_GameLift_Alias) GetType() string {
	return AWS_GameLift_Alias__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_GameLift_Alias) Set__LogicalName(v string) *AWS_GameLift_Alias {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_GameLift_Alias) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_GameLift_Alias {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_GameLift_Alias) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_GameLift_Alias {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_GameLift_Alias) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_GameLift_Alias {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_GameLift_Alias) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_GameLift_Alias {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_GameLift_Alias) Set__RequestedOutputs(v []cfz.Output) *AWS_GameLift_Alias {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_GameLift_Alias) Add__RequestedOutputs(v ...cfz.Output) *AWS_GameLift_Alias {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_GameLift_Alias) Set__Description(v cfz.Expression[string]) *AWS_GameLift_Alias {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_GameLift_Alias) SetV__Description(v string) *AWS_GameLift_Alias {
	t.Description = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_GameLift_Alias) Set__Name(v cfz.Expression[string]) *AWS_GameLift_Alias {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_GameLift_Alias) SetV__Name(v string) *AWS_GameLift_Alias {
	t.Name = cfz.V(v)
	return t
}

// Set__RoutingStrategy updates property "RoutingStrategy".
func (t *AWS_GameLift_Alias) Set__RoutingStrategy(v cfz.Expression[AWS_GameLift_Alias_RoutingStrategy]) *AWS_GameLift_Alias {
	t.RoutingStrategy = v
	return t
}

// SetV__RoutingStrategy updates property "RoutingStrategy".
func (t *AWS_GameLift_Alias) SetV__RoutingStrategy(v AWS_GameLift_Alias_RoutingStrategy) *AWS_GameLift_Alias {
	t.RoutingStrategy = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_GameLift_Alias) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__AliasId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AliasId
func (t *AWS_GameLift_Alias) GetAtt__AliasId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_Alias__AttributesMap.AliasId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_GameLift_Alias) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AliasId returns a conventionally configured output for an attribute of this resource.
// Attribute: AliasId
func (t *AWS_GameLift_Alias) GetConventionalOutputAtt__AliasId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAliasId", t.GetAtt__AliasId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_GameLift_Alias) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_GameLift_Alias

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

func (t *AWS_GameLift_Alias) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
