// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_glue

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Glue_Connection)(nil)
	_ cfz.Resource                   = (*AWS_Glue_Connection)(nil)
)

const (
	// AWS_Glue_Connection__Type is the CloudFormation type for AWS::Glue::Connection.
	AWS_Glue_Connection__Type = "AWS::Glue::Connection"
)

var (
	// AWS_Glue_Connection__PropertiesMap reports all the CloudFormation properties for AWS::Glue::Connection.
	AWS_Glue_Connection__PropertiesMap = struct {
		CatalogId       string
		ConnectionInput string
	}{
		CatalogId:       "CatalogId",
		ConnectionInput: "ConnectionInput",
	}

	// AWS_Glue_Connection__PropertiesSlice reports all the CloudFormation properties for AWS::Glue::Connection.
	AWS_Glue_Connection__PropertiesSlice = []string{
		AWS_Glue_Connection__PropertiesMap.CatalogId,
		AWS_Glue_Connection__PropertiesMap.ConnectionInput,
	}
)

// AWS_Glue_Connection is a binding for AWS::Glue::Connection.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html
type AWS_Glue_Connection struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CatalogId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html#cfn-glue-connection-catalogid
	CatalogId cfz.Expression[string] `json:"CatalogId,omitempty"`

	// ConnectionInput is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html#cfn-glue-connection-connectioninput
	ConnectionInput cfz.Expression[AWS_Glue_Connection_ConnectionInput] `json:"ConnectionInput,omitempty"`
}

// New__AWS_Glue_Connection initializes a new *AWS_Glue_Connection.
func New__AWS_Glue_Connection(logicalName string) *AWS_Glue_Connection {
	return &AWS_Glue_Connection{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Glue_Connection) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Glue_Connection) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Glue_Connection) GetType() string {
	return AWS_Glue_Connection__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Glue_Connection) Set__LogicalName(v string) *AWS_Glue_Connection {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Glue_Connection) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Glue_Connection {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Glue_Connection) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Glue_Connection {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Glue_Connection) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Glue_Connection {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Glue_Connection) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Glue_Connection {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Glue_Connection) Set__RequestedOutputs(v []cfz.Output) *AWS_Glue_Connection {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Glue_Connection) Add__RequestedOutputs(v ...cfz.Output) *AWS_Glue_Connection {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CatalogId updates property "CatalogId".
func (t *AWS_Glue_Connection) Set__CatalogId(v cfz.Expression[string]) *AWS_Glue_Connection {
	t.CatalogId = v
	return t
}

// SetV__CatalogId updates property "CatalogId".
func (t *AWS_Glue_Connection) SetV__CatalogId(v string) *AWS_Glue_Connection {
	t.CatalogId = cfz.V(v)
	return t
}

// Set__ConnectionInput updates property "ConnectionInput".
func (t *AWS_Glue_Connection) Set__ConnectionInput(v cfz.Expression[AWS_Glue_Connection_ConnectionInput]) *AWS_Glue_Connection {
	t.ConnectionInput = v
	return t
}

// SetV__ConnectionInput updates property "ConnectionInput".
func (t *AWS_Glue_Connection) SetV__ConnectionInput(v AWS_Glue_Connection_ConnectionInput) *AWS_Glue_Connection {
	t.ConnectionInput = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Glue_Connection) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Glue_Connection) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Glue_Connection) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Glue_Connection

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

func (t *AWS_Glue_Connection) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
