// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apigatewayv2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ApiGatewayV2_Deployment)(nil)
	_ cfz.Resource                   = (*AWS_ApiGatewayV2_Deployment)(nil)
)

const (
	// AWS_ApiGatewayV2_Deployment__Type is the CloudFormation type for AWS::ApiGatewayV2::Deployment.
	AWS_ApiGatewayV2_Deployment__Type = "AWS::ApiGatewayV2::Deployment"
)

var (
	// AWS_ApiGatewayV2_Deployment__AttributesMap reports all the CloudFormation attributes for AWS::ApiGatewayV2::Deployment.
	AWS_ApiGatewayV2_Deployment__AttributesMap = struct {
		DeploymentId string
	}{
		DeploymentId: "DeploymentId",
	}

	// AWS_ApiGatewayV2_Deployment__AttributesSlice reports all the CloudFormation attributes for AWS::ApiGatewayV2::Deployment.
	AWS_ApiGatewayV2_Deployment__AttributesSlice = []string{
		AWS_ApiGatewayV2_Deployment__AttributesMap.DeploymentId,
	}
)

var (
	// AWS_ApiGatewayV2_Deployment__PropertiesMap reports all the CloudFormation properties for AWS::ApiGatewayV2::Deployment.
	AWS_ApiGatewayV2_Deployment__PropertiesMap = struct {
		ApiId       string
		Description string
		StageName   string
	}{
		ApiId:       "ApiId",
		Description: "Description",
		StageName:   "StageName",
	}

	// AWS_ApiGatewayV2_Deployment__PropertiesSlice reports all the CloudFormation properties for AWS::ApiGatewayV2::Deployment.
	AWS_ApiGatewayV2_Deployment__PropertiesSlice = []string{
		AWS_ApiGatewayV2_Deployment__PropertiesMap.ApiId,
		AWS_ApiGatewayV2_Deployment__PropertiesMap.Description,
		AWS_ApiGatewayV2_Deployment__PropertiesMap.StageName,
	}
)

// AWS_ApiGatewayV2_Deployment is a binding for AWS::ApiGatewayV2::Deployment.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html
type AWS_ApiGatewayV2_Deployment struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ApiId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html#cfn-apigatewayv2-deployment-apiid
	ApiId cfz.Expression[string] `json:"ApiId,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html#cfn-apigatewayv2-deployment-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// StageName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html#cfn-apigatewayv2-deployment-stagename
	StageName cfz.Expression[string] `json:"StageName,omitempty"`
}

// New__AWS_ApiGatewayV2_Deployment initializes a new *AWS_ApiGatewayV2_Deployment.
func New__AWS_ApiGatewayV2_Deployment(logicalName string) *AWS_ApiGatewayV2_Deployment {
	return &AWS_ApiGatewayV2_Deployment{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ApiGatewayV2_Deployment) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ApiGatewayV2_Deployment) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ApiGatewayV2_Deployment) GetType() string {
	return AWS_ApiGatewayV2_Deployment__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ApiGatewayV2_Deployment) Set__LogicalName(v string) *AWS_ApiGatewayV2_Deployment {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ApiGatewayV2_Deployment) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ApiGatewayV2_Deployment {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ApiGatewayV2_Deployment) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ApiGatewayV2_Deployment {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ApiGatewayV2_Deployment) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ApiGatewayV2_Deployment {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ApiGatewayV2_Deployment) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ApiGatewayV2_Deployment {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ApiGatewayV2_Deployment) Set__RequestedOutputs(v []cfz.Output) *AWS_ApiGatewayV2_Deployment {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ApiGatewayV2_Deployment) Add__RequestedOutputs(v ...cfz.Output) *AWS_ApiGatewayV2_Deployment {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ApiId updates property "ApiId".
func (t *AWS_ApiGatewayV2_Deployment) Set__ApiId(v cfz.Expression[string]) *AWS_ApiGatewayV2_Deployment {
	t.ApiId = v
	return t
}

// SetV__ApiId updates property "ApiId".
func (t *AWS_ApiGatewayV2_Deployment) SetV__ApiId(v string) *AWS_ApiGatewayV2_Deployment {
	t.ApiId = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_ApiGatewayV2_Deployment) Set__Description(v cfz.Expression[string]) *AWS_ApiGatewayV2_Deployment {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_ApiGatewayV2_Deployment) SetV__Description(v string) *AWS_ApiGatewayV2_Deployment {
	t.Description = cfz.V(v)
	return t
}

// Set__StageName updates property "StageName".
func (t *AWS_ApiGatewayV2_Deployment) Set__StageName(v cfz.Expression[string]) *AWS_ApiGatewayV2_Deployment {
	t.StageName = v
	return t
}

// SetV__StageName updates property "StageName".
func (t *AWS_ApiGatewayV2_Deployment) SetV__StageName(v string) *AWS_ApiGatewayV2_Deployment {
	t.StageName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ApiGatewayV2_Deployment) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__DeploymentId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DeploymentId
func (t *AWS_ApiGatewayV2_Deployment) GetAtt__DeploymentId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ApiGatewayV2_Deployment__AttributesMap.DeploymentId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ApiGatewayV2_Deployment) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DeploymentId returns a conventionally configured output for an attribute of this resource.
// Attribute: DeploymentId
func (t *AWS_ApiGatewayV2_Deployment) GetConventionalOutputAtt__DeploymentId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDeploymentId", t.GetAtt__DeploymentId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ApiGatewayV2_Deployment) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ApiGatewayV2_Deployment

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

func (t *AWS_ApiGatewayV2_Deployment) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
