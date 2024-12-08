// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appstream

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_AppStream_StackFleetAssociation)(nil)
	_ cfz.Resource                   = (*AWS_AppStream_StackFleetAssociation)(nil)
)

const (
	// AWS_AppStream_StackFleetAssociation__Type is the CloudFormation type for AWS::AppStream::StackFleetAssociation.
	AWS_AppStream_StackFleetAssociation__Type = "AWS::AppStream::StackFleetAssociation"
)

var (
	// AWS_AppStream_StackFleetAssociation__PropertiesMap reports all the CloudFormation properties for AWS::AppStream::StackFleetAssociation.
	AWS_AppStream_StackFleetAssociation__PropertiesMap = struct {
		FleetName string
		StackName string
	}{
		FleetName: "FleetName",
		StackName: "StackName",
	}

	// AWS_AppStream_StackFleetAssociation__PropertiesSlice reports all the CloudFormation properties for AWS::AppStream::StackFleetAssociation.
	AWS_AppStream_StackFleetAssociation__PropertiesSlice = []string{
		AWS_AppStream_StackFleetAssociation__PropertiesMap.FleetName,
		AWS_AppStream_StackFleetAssociation__PropertiesMap.StackName,
	}
)

// AWS_AppStream_StackFleetAssociation is a binding for AWS::AppStream::StackFleetAssociation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html
type AWS_AppStream_StackFleetAssociation struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// FleetName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-fleetname
	FleetName cfz.Expression[string] `json:"FleetName,omitempty"`

	// StackName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-stackname
	StackName cfz.Expression[string] `json:"StackName,omitempty"`
}

// New__AWS_AppStream_StackFleetAssociation initializes a new *AWS_AppStream_StackFleetAssociation.
func New__AWS_AppStream_StackFleetAssociation(logicalName string) *AWS_AppStream_StackFleetAssociation {
	return &AWS_AppStream_StackFleetAssociation{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_AppStream_StackFleetAssociation) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_AppStream_StackFleetAssociation) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_AppStream_StackFleetAssociation) GetType() string {
	return AWS_AppStream_StackFleetAssociation__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_AppStream_StackFleetAssociation) Set__LogicalName(v string) *AWS_AppStream_StackFleetAssociation {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_AppStream_StackFleetAssociation) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_AppStream_StackFleetAssociation {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_AppStream_StackFleetAssociation) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_AppStream_StackFleetAssociation {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_AppStream_StackFleetAssociation) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_AppStream_StackFleetAssociation {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_AppStream_StackFleetAssociation) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_AppStream_StackFleetAssociation {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_AppStream_StackFleetAssociation) Set__RequestedOutputs(v []cfz.Output) *AWS_AppStream_StackFleetAssociation {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_AppStream_StackFleetAssociation) Add__RequestedOutputs(v ...cfz.Output) *AWS_AppStream_StackFleetAssociation {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__FleetName updates property "FleetName".
func (t *AWS_AppStream_StackFleetAssociation) Set__FleetName(v cfz.Expression[string]) *AWS_AppStream_StackFleetAssociation {
	t.FleetName = v
	return t
}

// SetV__FleetName updates property "FleetName".
func (t *AWS_AppStream_StackFleetAssociation) SetV__FleetName(v string) *AWS_AppStream_StackFleetAssociation {
	t.FleetName = cfz.V(v)
	return t
}

// Set__StackName updates property "StackName".
func (t *AWS_AppStream_StackFleetAssociation) Set__StackName(v cfz.Expression[string]) *AWS_AppStream_StackFleetAssociation {
	t.StackName = v
	return t
}

// SetV__StackName updates property "StackName".
func (t *AWS_AppStream_StackFleetAssociation) SetV__StackName(v string) *AWS_AppStream_StackFleetAssociation {
	t.StackName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_AppStream_StackFleetAssociation) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_AppStream_StackFleetAssociation) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_AppStream_StackFleetAssociation) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_AppStream_StackFleetAssociation

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

func (t *AWS_AppStream_StackFleetAssociation) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
