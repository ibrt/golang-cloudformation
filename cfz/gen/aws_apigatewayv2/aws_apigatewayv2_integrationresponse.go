// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apigatewayv2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ApiGatewayV2_IntegrationResponse)(nil)
	_ cfz.Resource                   = (*AWS_ApiGatewayV2_IntegrationResponse)(nil)
)

const (
	// AWS_ApiGatewayV2_IntegrationResponse__Type is the CloudFormation type for AWS::ApiGatewayV2::IntegrationResponse.
	AWS_ApiGatewayV2_IntegrationResponse__Type = "AWS::ApiGatewayV2::IntegrationResponse"
)

var (
	// AWS_ApiGatewayV2_IntegrationResponse__AttributesMap reports all the CloudFormation attributes for AWS::ApiGatewayV2::IntegrationResponse.
	AWS_ApiGatewayV2_IntegrationResponse__AttributesMap = struct {
		IntegrationResponseId string
	}{
		IntegrationResponseId: "IntegrationResponseId",
	}

	// AWS_ApiGatewayV2_IntegrationResponse__AttributesSlice reports all the CloudFormation attributes for AWS::ApiGatewayV2::IntegrationResponse.
	AWS_ApiGatewayV2_IntegrationResponse__AttributesSlice = []string{
		AWS_ApiGatewayV2_IntegrationResponse__AttributesMap.IntegrationResponseId,
	}
)

var (
	// AWS_ApiGatewayV2_IntegrationResponse__PropertiesMap reports all the CloudFormation properties for AWS::ApiGatewayV2::IntegrationResponse.
	AWS_ApiGatewayV2_IntegrationResponse__PropertiesMap = struct {
		ApiId                       string
		ContentHandlingStrategy     string
		IntegrationId               string
		IntegrationResponseKey      string
		ResponseParameters          string
		ResponseTemplates           string
		TemplateSelectionExpression string
	}{
		ApiId:                       "ApiId",
		ContentHandlingStrategy:     "ContentHandlingStrategy",
		IntegrationId:               "IntegrationId",
		IntegrationResponseKey:      "IntegrationResponseKey",
		ResponseParameters:          "ResponseParameters",
		ResponseTemplates:           "ResponseTemplates",
		TemplateSelectionExpression: "TemplateSelectionExpression",
	}

	// AWS_ApiGatewayV2_IntegrationResponse__PropertiesSlice reports all the CloudFormation properties for AWS::ApiGatewayV2::IntegrationResponse.
	AWS_ApiGatewayV2_IntegrationResponse__PropertiesSlice = []string{
		AWS_ApiGatewayV2_IntegrationResponse__PropertiesMap.ApiId,
		AWS_ApiGatewayV2_IntegrationResponse__PropertiesMap.ContentHandlingStrategy,
		AWS_ApiGatewayV2_IntegrationResponse__PropertiesMap.IntegrationId,
		AWS_ApiGatewayV2_IntegrationResponse__PropertiesMap.IntegrationResponseKey,
		AWS_ApiGatewayV2_IntegrationResponse__PropertiesMap.ResponseParameters,
		AWS_ApiGatewayV2_IntegrationResponse__PropertiesMap.ResponseTemplates,
		AWS_ApiGatewayV2_IntegrationResponse__PropertiesMap.TemplateSelectionExpression,
	}
)

// AWS_ApiGatewayV2_IntegrationResponse is a binding for AWS::ApiGatewayV2::IntegrationResponse.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html
type AWS_ApiGatewayV2_IntegrationResponse struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-apiid
	ApiId cfz.Expression[string] `json:"ApiId,omitempty"`

	// ContentHandlingStrategy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-contenthandlingstrategy
	ContentHandlingStrategy cfz.Expression[string] `json:"ContentHandlingStrategy,omitempty"`

	// IntegrationId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-integrationid
	IntegrationId cfz.Expression[string] `json:"IntegrationId,omitempty"`

	// IntegrationResponseKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-integrationresponsekey
	IntegrationResponseKey cfz.Expression[string] `json:"IntegrationResponseKey,omitempty"`

	// ResponseParameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-responseparameters
	ResponseParameters cfz.Expression[json.RawMessage] `json:"ResponseParameters,omitempty"`

	// ResponseTemplates is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-responsetemplates
	ResponseTemplates cfz.Expression[json.RawMessage] `json:"ResponseTemplates,omitempty"`

	// TemplateSelectionExpression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html#cfn-apigatewayv2-integrationresponse-templateselectionexpression
	TemplateSelectionExpression cfz.Expression[string] `json:"TemplateSelectionExpression,omitempty"`
}

// New__AWS_ApiGatewayV2_IntegrationResponse initializes a new *AWS_ApiGatewayV2_IntegrationResponse.
func New__AWS_ApiGatewayV2_IntegrationResponse(logicalName string) *AWS_ApiGatewayV2_IntegrationResponse {
	return &AWS_ApiGatewayV2_IntegrationResponse{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ApiGatewayV2_IntegrationResponse) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ApiGatewayV2_IntegrationResponse) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ApiGatewayV2_IntegrationResponse) GetType() string {
	return AWS_ApiGatewayV2_IntegrationResponse__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__LogicalName(v string) *AWS_ApiGatewayV2_IntegrationResponse {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ApiGatewayV2_IntegrationResponse {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ApiGatewayV2_IntegrationResponse {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ApiGatewayV2_IntegrationResponse {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ApiGatewayV2_IntegrationResponse {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__RequestedOutputs(v []cfz.Output) *AWS_ApiGatewayV2_IntegrationResponse {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Add__RequestedOutputs(v ...cfz.Output) *AWS_ApiGatewayV2_IntegrationResponse {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ApiId updates property "ApiId".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__ApiId(v cfz.Expression[string]) *AWS_ApiGatewayV2_IntegrationResponse {
	t.ApiId = v
	return t
}

// SetV__ApiId updates property "ApiId".
func (t *AWS_ApiGatewayV2_IntegrationResponse) SetV__ApiId(v string) *AWS_ApiGatewayV2_IntegrationResponse {
	t.ApiId = cfz.V(v)
	return t
}

// Set__ContentHandlingStrategy updates property "ContentHandlingStrategy".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__ContentHandlingStrategy(v cfz.Expression[string]) *AWS_ApiGatewayV2_IntegrationResponse {
	t.ContentHandlingStrategy = v
	return t
}

// SetV__ContentHandlingStrategy updates property "ContentHandlingStrategy".
func (t *AWS_ApiGatewayV2_IntegrationResponse) SetV__ContentHandlingStrategy(v string) *AWS_ApiGatewayV2_IntegrationResponse {
	t.ContentHandlingStrategy = cfz.V(v)
	return t
}

// Set__IntegrationId updates property "IntegrationId".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__IntegrationId(v cfz.Expression[string]) *AWS_ApiGatewayV2_IntegrationResponse {
	t.IntegrationId = v
	return t
}

// SetV__IntegrationId updates property "IntegrationId".
func (t *AWS_ApiGatewayV2_IntegrationResponse) SetV__IntegrationId(v string) *AWS_ApiGatewayV2_IntegrationResponse {
	t.IntegrationId = cfz.V(v)
	return t
}

// Set__IntegrationResponseKey updates property "IntegrationResponseKey".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__IntegrationResponseKey(v cfz.Expression[string]) *AWS_ApiGatewayV2_IntegrationResponse {
	t.IntegrationResponseKey = v
	return t
}

// SetV__IntegrationResponseKey updates property "IntegrationResponseKey".
func (t *AWS_ApiGatewayV2_IntegrationResponse) SetV__IntegrationResponseKey(v string) *AWS_ApiGatewayV2_IntegrationResponse {
	t.IntegrationResponseKey = cfz.V(v)
	return t
}

// Set__ResponseParameters updates property "ResponseParameters".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__ResponseParameters(v cfz.Expression[json.RawMessage]) *AWS_ApiGatewayV2_IntegrationResponse {
	t.ResponseParameters = v
	return t
}

// SetV__ResponseParameters updates property "ResponseParameters".
func (t *AWS_ApiGatewayV2_IntegrationResponse) SetV__ResponseParameters(v json.RawMessage) *AWS_ApiGatewayV2_IntegrationResponse {
	t.ResponseParameters = cfz.V(v)
	return t
}

// Set__ResponseTemplates updates property "ResponseTemplates".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__ResponseTemplates(v cfz.Expression[json.RawMessage]) *AWS_ApiGatewayV2_IntegrationResponse {
	t.ResponseTemplates = v
	return t
}

// SetV__ResponseTemplates updates property "ResponseTemplates".
func (t *AWS_ApiGatewayV2_IntegrationResponse) SetV__ResponseTemplates(v json.RawMessage) *AWS_ApiGatewayV2_IntegrationResponse {
	t.ResponseTemplates = cfz.V(v)
	return t
}

// Set__TemplateSelectionExpression updates property "TemplateSelectionExpression".
func (t *AWS_ApiGatewayV2_IntegrationResponse) Set__TemplateSelectionExpression(v cfz.Expression[string]) *AWS_ApiGatewayV2_IntegrationResponse {
	t.TemplateSelectionExpression = v
	return t
}

// SetV__TemplateSelectionExpression updates property "TemplateSelectionExpression".
func (t *AWS_ApiGatewayV2_IntegrationResponse) SetV__TemplateSelectionExpression(v string) *AWS_ApiGatewayV2_IntegrationResponse {
	t.TemplateSelectionExpression = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ApiGatewayV2_IntegrationResponse) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__IntegrationResponseId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: IntegrationResponseId
func (t *AWS_ApiGatewayV2_IntegrationResponse) GetAtt__IntegrationResponseId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ApiGatewayV2_IntegrationResponse__AttributesMap.IntegrationResponseId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ApiGatewayV2_IntegrationResponse) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__IntegrationResponseId returns a conventionally configured output for an attribute of this resource.
// Attribute: IntegrationResponseId
func (t *AWS_ApiGatewayV2_IntegrationResponse) GetConventionalOutputAtt__IntegrationResponseId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttIntegrationResponseId", t.GetAtt__IntegrationResponseId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ApiGatewayV2_IntegrationResponse) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ApiGatewayV2_IntegrationResponse

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

func (t *AWS_ApiGatewayV2_IntegrationResponse) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
