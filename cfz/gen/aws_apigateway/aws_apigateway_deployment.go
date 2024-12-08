// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apigateway

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ApiGateway_Deployment)(nil)
	_ cfz.Resource                   = (*AWS_ApiGateway_Deployment)(nil)
)

const (
	// AWS_ApiGateway_Deployment__Type is the CloudFormation type for AWS::ApiGateway::Deployment.
	AWS_ApiGateway_Deployment__Type = "AWS::ApiGateway::Deployment"
)

var (
	// AWS_ApiGateway_Deployment__AttributesMap reports all the CloudFormation attributes for AWS::ApiGateway::Deployment.
	AWS_ApiGateway_Deployment__AttributesMap = struct {
		DeploymentId string
	}{
		DeploymentId: "DeploymentId",
	}

	// AWS_ApiGateway_Deployment__AttributesSlice reports all the CloudFormation attributes for AWS::ApiGateway::Deployment.
	AWS_ApiGateway_Deployment__AttributesSlice = []string{
		AWS_ApiGateway_Deployment__AttributesMap.DeploymentId,
	}
)

var (
	// AWS_ApiGateway_Deployment__PropertiesMap reports all the CloudFormation properties for AWS::ApiGateway::Deployment.
	AWS_ApiGateway_Deployment__PropertiesMap = struct {
		DeploymentCanarySettings string
		Description              string
		RestApiId                string
		StageDescription         string
		StageName                string
	}{
		DeploymentCanarySettings: "DeploymentCanarySettings",
		Description:              "Description",
		RestApiId:                "RestApiId",
		StageDescription:         "StageDescription",
		StageName:                "StageName",
	}

	// AWS_ApiGateway_Deployment__PropertiesSlice reports all the CloudFormation properties for AWS::ApiGateway::Deployment.
	AWS_ApiGateway_Deployment__PropertiesSlice = []string{
		AWS_ApiGateway_Deployment__PropertiesMap.DeploymentCanarySettings,
		AWS_ApiGateway_Deployment__PropertiesMap.Description,
		AWS_ApiGateway_Deployment__PropertiesMap.RestApiId,
		AWS_ApiGateway_Deployment__PropertiesMap.StageDescription,
		AWS_ApiGateway_Deployment__PropertiesMap.StageName,
	}
)

// AWS_ApiGateway_Deployment is a binding for AWS::ApiGateway::Deployment.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html
type AWS_ApiGateway_Deployment struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DeploymentCanarySettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-deploymentcanarysettings
	DeploymentCanarySettings cfz.Expression[AWS_ApiGateway_Deployment_DeploymentCanarySettings] `json:"DeploymentCanarySettings,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// RestApiId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-restapiid
	RestApiId cfz.Expression[string] `json:"RestApiId,omitempty"`

	// StageDescription is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagedescription
	StageDescription cfz.Expression[AWS_ApiGateway_Deployment_StageDescription] `json:"StageDescription,omitempty"`

	// StageName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagename
	StageName cfz.Expression[string] `json:"StageName,omitempty"`
}

// New__AWS_ApiGateway_Deployment initializes a new *AWS_ApiGateway_Deployment.
func New__AWS_ApiGateway_Deployment(logicalName string) *AWS_ApiGateway_Deployment {
	return &AWS_ApiGateway_Deployment{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ApiGateway_Deployment) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ApiGateway_Deployment) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ApiGateway_Deployment) GetType() string {
	return AWS_ApiGateway_Deployment__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ApiGateway_Deployment) Set__LogicalName(v string) *AWS_ApiGateway_Deployment {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ApiGateway_Deployment) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ApiGateway_Deployment {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ApiGateway_Deployment) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ApiGateway_Deployment {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ApiGateway_Deployment) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ApiGateway_Deployment {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ApiGateway_Deployment) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ApiGateway_Deployment {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ApiGateway_Deployment) Set__RequestedOutputs(v []cfz.Output) *AWS_ApiGateway_Deployment {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ApiGateway_Deployment) Add__RequestedOutputs(v ...cfz.Output) *AWS_ApiGateway_Deployment {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DeploymentCanarySettings updates property "DeploymentCanarySettings".
func (t *AWS_ApiGateway_Deployment) Set__DeploymentCanarySettings(v cfz.Expression[AWS_ApiGateway_Deployment_DeploymentCanarySettings]) *AWS_ApiGateway_Deployment {
	t.DeploymentCanarySettings = v
	return t
}

// SetV__DeploymentCanarySettings updates property "DeploymentCanarySettings".
func (t *AWS_ApiGateway_Deployment) SetV__DeploymentCanarySettings(v AWS_ApiGateway_Deployment_DeploymentCanarySettings) *AWS_ApiGateway_Deployment {
	t.DeploymentCanarySettings = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_ApiGateway_Deployment) Set__Description(v cfz.Expression[string]) *AWS_ApiGateway_Deployment {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_ApiGateway_Deployment) SetV__Description(v string) *AWS_ApiGateway_Deployment {
	t.Description = cfz.V(v)
	return t
}

// Set__RestApiId updates property "RestApiId".
func (t *AWS_ApiGateway_Deployment) Set__RestApiId(v cfz.Expression[string]) *AWS_ApiGateway_Deployment {
	t.RestApiId = v
	return t
}

// SetV__RestApiId updates property "RestApiId".
func (t *AWS_ApiGateway_Deployment) SetV__RestApiId(v string) *AWS_ApiGateway_Deployment {
	t.RestApiId = cfz.V(v)
	return t
}

// Set__StageDescription updates property "StageDescription".
func (t *AWS_ApiGateway_Deployment) Set__StageDescription(v cfz.Expression[AWS_ApiGateway_Deployment_StageDescription]) *AWS_ApiGateway_Deployment {
	t.StageDescription = v
	return t
}

// SetV__StageDescription updates property "StageDescription".
func (t *AWS_ApiGateway_Deployment) SetV__StageDescription(v AWS_ApiGateway_Deployment_StageDescription) *AWS_ApiGateway_Deployment {
	t.StageDescription = cfz.V(v)
	return t
}

// Set__StageName updates property "StageName".
func (t *AWS_ApiGateway_Deployment) Set__StageName(v cfz.Expression[string]) *AWS_ApiGateway_Deployment {
	t.StageName = v
	return t
}

// SetV__StageName updates property "StageName".
func (t *AWS_ApiGateway_Deployment) SetV__StageName(v string) *AWS_ApiGateway_Deployment {
	t.StageName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ApiGateway_Deployment) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__DeploymentId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DeploymentId
func (t *AWS_ApiGateway_Deployment) GetAtt__DeploymentId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ApiGateway_Deployment__AttributesMap.DeploymentId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ApiGateway_Deployment) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DeploymentId returns a conventionally configured output for an attribute of this resource.
// Attribute: DeploymentId
func (t *AWS_ApiGateway_Deployment) GetConventionalOutputAtt__DeploymentId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDeploymentId", t.GetAtt__DeploymentId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ApiGateway_Deployment) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ApiGateway_Deployment

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

func (t *AWS_ApiGateway_Deployment) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
