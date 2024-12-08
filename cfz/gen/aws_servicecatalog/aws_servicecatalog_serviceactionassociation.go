// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_servicecatalog

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ServiceCatalog_ServiceActionAssociation)(nil)
	_ cfz.Resource                   = (*AWS_ServiceCatalog_ServiceActionAssociation)(nil)
)

const (
	// AWS_ServiceCatalog_ServiceActionAssociation__Type is the CloudFormation type for AWS::ServiceCatalog::ServiceActionAssociation.
	AWS_ServiceCatalog_ServiceActionAssociation__Type = "AWS::ServiceCatalog::ServiceActionAssociation"
)

var (
	// AWS_ServiceCatalog_ServiceActionAssociation__PropertiesMap reports all the CloudFormation properties for AWS::ServiceCatalog::ServiceActionAssociation.
	AWS_ServiceCatalog_ServiceActionAssociation__PropertiesMap = struct {
		ProductId              string
		ProvisioningArtifactId string
		ServiceActionId        string
	}{
		ProductId:              "ProductId",
		ProvisioningArtifactId: "ProvisioningArtifactId",
		ServiceActionId:        "ServiceActionId",
	}

	// AWS_ServiceCatalog_ServiceActionAssociation__PropertiesSlice reports all the CloudFormation properties for AWS::ServiceCatalog::ServiceActionAssociation.
	AWS_ServiceCatalog_ServiceActionAssociation__PropertiesSlice = []string{
		AWS_ServiceCatalog_ServiceActionAssociation__PropertiesMap.ProductId,
		AWS_ServiceCatalog_ServiceActionAssociation__PropertiesMap.ProvisioningArtifactId,
		AWS_ServiceCatalog_ServiceActionAssociation__PropertiesMap.ServiceActionId,
	}
)

// AWS_ServiceCatalog_ServiceActionAssociation is a binding for AWS::ServiceCatalog::ServiceActionAssociation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceactionassociation.html
type AWS_ServiceCatalog_ServiceActionAssociation struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ProductId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceactionassociation.html#cfn-servicecatalog-serviceactionassociation-productid
	ProductId cfz.Expression[string] `json:"ProductId,omitempty"`

	// ProvisioningArtifactId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceactionassociation.html#cfn-servicecatalog-serviceactionassociation-provisioningartifactid
	ProvisioningArtifactId cfz.Expression[string] `json:"ProvisioningArtifactId,omitempty"`

	// ServiceActionId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceactionassociation.html#cfn-servicecatalog-serviceactionassociation-serviceactionid
	ServiceActionId cfz.Expression[string] `json:"ServiceActionId,omitempty"`
}

// New__AWS_ServiceCatalog_ServiceActionAssociation initializes a new *AWS_ServiceCatalog_ServiceActionAssociation.
func New__AWS_ServiceCatalog_ServiceActionAssociation(logicalName string) *AWS_ServiceCatalog_ServiceActionAssociation {
	return &AWS_ServiceCatalog_ServiceActionAssociation{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ServiceCatalog_ServiceActionAssociation) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ServiceCatalog_ServiceActionAssociation) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ServiceCatalog_ServiceActionAssociation) GetType() string {
	return AWS_ServiceCatalog_ServiceActionAssociation__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Set__LogicalName(v string) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Set__RequestedOutputs(v []cfz.Output) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Add__RequestedOutputs(v ...cfz.Output) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ProductId updates property "ProductId".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Set__ProductId(v cfz.Expression[string]) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.ProductId = v
	return t
}

// SetV__ProductId updates property "ProductId".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) SetV__ProductId(v string) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.ProductId = cfz.V(v)
	return t
}

// Set__ProvisioningArtifactId updates property "ProvisioningArtifactId".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Set__ProvisioningArtifactId(v cfz.Expression[string]) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.ProvisioningArtifactId = v
	return t
}

// SetV__ProvisioningArtifactId updates property "ProvisioningArtifactId".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) SetV__ProvisioningArtifactId(v string) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.ProvisioningArtifactId = cfz.V(v)
	return t
}

// Set__ServiceActionId updates property "ServiceActionId".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Set__ServiceActionId(v cfz.Expression[string]) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.ServiceActionId = v
	return t
}

// SetV__ServiceActionId updates property "ServiceActionId".
func (t *AWS_ServiceCatalog_ServiceActionAssociation) SetV__ServiceActionId(v string) *AWS_ServiceCatalog_ServiceActionAssociation {
	t.ServiceActionId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ServiceCatalog_ServiceActionAssociation) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ServiceCatalog_ServiceActionAssociation) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ServiceCatalog_ServiceActionAssociation) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ServiceCatalog_ServiceActionAssociation

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

func (t *AWS_ServiceCatalog_ServiceActionAssociation) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
