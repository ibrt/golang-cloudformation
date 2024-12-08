// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_AppMesh_Mesh)(nil)
	_ cfz.Resource                   = (*AWS_AppMesh_Mesh)(nil)
)

const (
	// AWS_AppMesh_Mesh__Type is the CloudFormation type for AWS::AppMesh::Mesh.
	AWS_AppMesh_Mesh__Type = "AWS::AppMesh::Mesh"
)

var (
	// AWS_AppMesh_Mesh__AttributesMap reports all the CloudFormation attributes for AWS::AppMesh::Mesh.
	AWS_AppMesh_Mesh__AttributesMap = struct {
		Arn           string
		MeshName      string
		MeshOwner     string
		ResourceOwner string
		Uid           string
	}{
		Arn:           "Arn",
		MeshName:      "MeshName",
		MeshOwner:     "MeshOwner",
		ResourceOwner: "ResourceOwner",
		Uid:           "Uid",
	}

	// AWS_AppMesh_Mesh__AttributesSlice reports all the CloudFormation attributes for AWS::AppMesh::Mesh.
	AWS_AppMesh_Mesh__AttributesSlice = []string{
		AWS_AppMesh_Mesh__AttributesMap.Arn,
		AWS_AppMesh_Mesh__AttributesMap.MeshName,
		AWS_AppMesh_Mesh__AttributesMap.MeshOwner,
		AWS_AppMesh_Mesh__AttributesMap.ResourceOwner,
		AWS_AppMesh_Mesh__AttributesMap.Uid,
	}
)

var (
	// AWS_AppMesh_Mesh__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::Mesh.
	AWS_AppMesh_Mesh__PropertiesMap = struct {
		MeshName string
		Spec     string
		Tags     string
	}{
		MeshName: "MeshName",
		Spec:     "Spec",
		Tags:     "Tags",
	}

	// AWS_AppMesh_Mesh__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::Mesh.
	AWS_AppMesh_Mesh__PropertiesSlice = []string{
		AWS_AppMesh_Mesh__PropertiesMap.MeshName,
		AWS_AppMesh_Mesh__PropertiesMap.Spec,
		AWS_AppMesh_Mesh__PropertiesMap.Tags,
	}
)

// AWS_AppMesh_Mesh is a binding for AWS::AppMesh::Mesh.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html
type AWS_AppMesh_Mesh struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// MeshName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html#cfn-appmesh-mesh-meshname
	MeshName cfz.Expression[string] `json:"MeshName,omitempty"`

	// Spec is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html#cfn-appmesh-mesh-spec
	Spec cfz.Expression[AWS_AppMesh_Mesh_MeshSpec] `json:"Spec,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html#cfn-appmesh-mesh-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_AppMesh_Mesh initializes a new *AWS_AppMesh_Mesh.
func New__AWS_AppMesh_Mesh(logicalName string) *AWS_AppMesh_Mesh {
	return &AWS_AppMesh_Mesh{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_AppMesh_Mesh) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_AppMesh_Mesh) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_AppMesh_Mesh) GetType() string {
	return AWS_AppMesh_Mesh__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_AppMesh_Mesh) Set__LogicalName(v string) *AWS_AppMesh_Mesh {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_AppMesh_Mesh) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_AppMesh_Mesh {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_AppMesh_Mesh) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_AppMesh_Mesh {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_AppMesh_Mesh) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_AppMesh_Mesh {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_AppMesh_Mesh) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_AppMesh_Mesh {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_AppMesh_Mesh) Set__RequestedOutputs(v []cfz.Output) *AWS_AppMesh_Mesh {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_AppMesh_Mesh) Add__RequestedOutputs(v ...cfz.Output) *AWS_AppMesh_Mesh {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__MeshName updates property "MeshName".
func (t *AWS_AppMesh_Mesh) Set__MeshName(v cfz.Expression[string]) *AWS_AppMesh_Mesh {
	t.MeshName = v
	return t
}

// SetV__MeshName updates property "MeshName".
func (t *AWS_AppMesh_Mesh) SetV__MeshName(v string) *AWS_AppMesh_Mesh {
	t.MeshName = cfz.V(v)
	return t
}

// Set__Spec updates property "Spec".
func (t *AWS_AppMesh_Mesh) Set__Spec(v cfz.Expression[AWS_AppMesh_Mesh_MeshSpec]) *AWS_AppMesh_Mesh {
	t.Spec = v
	return t
}

// SetV__Spec updates property "Spec".
func (t *AWS_AppMesh_Mesh) SetV__Spec(v AWS_AppMesh_Mesh_MeshSpec) *AWS_AppMesh_Mesh {
	t.Spec = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_AppMesh_Mesh) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_AppMesh_Mesh {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_AppMesh_Mesh) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_AppMesh_Mesh {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_AppMesh_Mesh) SetSV__Tags(v ...cfz.Tag) *AWS_AppMesh_Mesh {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_AppMesh_Mesh) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_AppMesh_Mesh) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppMesh_Mesh__AttributesMap.Arn))
}

// GetAtt__MeshName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: MeshName
func (t *AWS_AppMesh_Mesh) GetAtt__MeshName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppMesh_Mesh__AttributesMap.MeshName))
}

// GetAtt__MeshOwner returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: MeshOwner
func (t *AWS_AppMesh_Mesh) GetAtt__MeshOwner() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppMesh_Mesh__AttributesMap.MeshOwner))
}

// GetAtt__ResourceOwner returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ResourceOwner
func (t *AWS_AppMesh_Mesh) GetAtt__ResourceOwner() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppMesh_Mesh__AttributesMap.ResourceOwner))
}

// GetAtt__Uid returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Uid
func (t *AWS_AppMesh_Mesh) GetAtt__Uid() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppMesh_Mesh__AttributesMap.Uid))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_AppMesh_Mesh) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_AppMesh_Mesh) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__MeshName returns a conventionally configured output for an attribute of this resource.
// Attribute: MeshName
func (t *AWS_AppMesh_Mesh) GetConventionalOutputAtt__MeshName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttMeshName", t.GetAtt__MeshName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__MeshOwner returns a conventionally configured output for an attribute of this resource.
// Attribute: MeshOwner
func (t *AWS_AppMesh_Mesh) GetConventionalOutputAtt__MeshOwner(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttMeshOwner", t.GetAtt__MeshOwner())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ResourceOwner returns a conventionally configured output for an attribute of this resource.
// Attribute: ResourceOwner
func (t *AWS_AppMesh_Mesh) GetConventionalOutputAtt__ResourceOwner(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttResourceOwner", t.GetAtt__ResourceOwner())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Uid returns a conventionally configured output for an attribute of this resource.
// Attribute: Uid
func (t *AWS_AppMesh_Mesh) GetConventionalOutputAtt__Uid(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttUid", t.GetAtt__Uid())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_AppMesh_Mesh) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_AppMesh_Mesh

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

func (t *AWS_AppMesh_Mesh) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
