// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apigateway

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ApiGateway_RequestValidator)(nil)
	_ cfz.Resource                   = (*AWS_ApiGateway_RequestValidator)(nil)
)

const (
	// AWS_ApiGateway_RequestValidator__Type is the CloudFormation type for AWS::ApiGateway::RequestValidator.
	AWS_ApiGateway_RequestValidator__Type = "AWS::ApiGateway::RequestValidator"
)

var (
	// AWS_ApiGateway_RequestValidator__AttributesMap reports all the CloudFormation attributes for AWS::ApiGateway::RequestValidator.
	AWS_ApiGateway_RequestValidator__AttributesMap = struct {
		RequestValidatorId string
	}{
		RequestValidatorId: "RequestValidatorId",
	}

	// AWS_ApiGateway_RequestValidator__AttributesSlice reports all the CloudFormation attributes for AWS::ApiGateway::RequestValidator.
	AWS_ApiGateway_RequestValidator__AttributesSlice = []string{
		AWS_ApiGateway_RequestValidator__AttributesMap.RequestValidatorId,
	}
)

var (
	// AWS_ApiGateway_RequestValidator__PropertiesMap reports all the CloudFormation properties for AWS::ApiGateway::RequestValidator.
	AWS_ApiGateway_RequestValidator__PropertiesMap = struct {
		Name                      string
		RestApiId                 string
		ValidateRequestBody       string
		ValidateRequestParameters string
	}{
		Name:                      "Name",
		RestApiId:                 "RestApiId",
		ValidateRequestBody:       "ValidateRequestBody",
		ValidateRequestParameters: "ValidateRequestParameters",
	}

	// AWS_ApiGateway_RequestValidator__PropertiesSlice reports all the CloudFormation properties for AWS::ApiGateway::RequestValidator.
	AWS_ApiGateway_RequestValidator__PropertiesSlice = []string{
		AWS_ApiGateway_RequestValidator__PropertiesMap.Name,
		AWS_ApiGateway_RequestValidator__PropertiesMap.RestApiId,
		AWS_ApiGateway_RequestValidator__PropertiesMap.ValidateRequestBody,
		AWS_ApiGateway_RequestValidator__PropertiesMap.ValidateRequestParameters,
	}
)

// AWS_ApiGateway_RequestValidator is a binding for AWS::ApiGateway::RequestValidator.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html
type AWS_ApiGateway_RequestValidator struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// RestApiId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-restapiid
	RestApiId cfz.Expression[string] `json:"RestApiId,omitempty"`

	// ValidateRequestBody is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-validaterequestbody
	ValidateRequestBody cfz.Expression[bool] `json:"ValidateRequestBody,omitempty"`

	// ValidateRequestParameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-validaterequestparameters
	ValidateRequestParameters cfz.Expression[bool] `json:"ValidateRequestParameters,omitempty"`
}

// New__AWS_ApiGateway_RequestValidator initializes a new *AWS_ApiGateway_RequestValidator.
func New__AWS_ApiGateway_RequestValidator(logicalName string) *AWS_ApiGateway_RequestValidator {
	return &AWS_ApiGateway_RequestValidator{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ApiGateway_RequestValidator) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ApiGateway_RequestValidator) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ApiGateway_RequestValidator) GetType() string {
	return AWS_ApiGateway_RequestValidator__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ApiGateway_RequestValidator) Set__LogicalName(v string) *AWS_ApiGateway_RequestValidator {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ApiGateway_RequestValidator) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ApiGateway_RequestValidator {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ApiGateway_RequestValidator) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ApiGateway_RequestValidator {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ApiGateway_RequestValidator) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ApiGateway_RequestValidator {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ApiGateway_RequestValidator) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ApiGateway_RequestValidator {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ApiGateway_RequestValidator) Set__RequestedOutputs(v []cfz.Output) *AWS_ApiGateway_RequestValidator {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ApiGateway_RequestValidator) Add__RequestedOutputs(v ...cfz.Output) *AWS_ApiGateway_RequestValidator {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_ApiGateway_RequestValidator) Set__Name(v cfz.Expression[string]) *AWS_ApiGateway_RequestValidator {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_ApiGateway_RequestValidator) SetV__Name(v string) *AWS_ApiGateway_RequestValidator {
	t.Name = cfz.V(v)
	return t
}

// Set__RestApiId updates property "RestApiId".
func (t *AWS_ApiGateway_RequestValidator) Set__RestApiId(v cfz.Expression[string]) *AWS_ApiGateway_RequestValidator {
	t.RestApiId = v
	return t
}

// SetV__RestApiId updates property "RestApiId".
func (t *AWS_ApiGateway_RequestValidator) SetV__RestApiId(v string) *AWS_ApiGateway_RequestValidator {
	t.RestApiId = cfz.V(v)
	return t
}

// Set__ValidateRequestBody updates property "ValidateRequestBody".
func (t *AWS_ApiGateway_RequestValidator) Set__ValidateRequestBody(v cfz.Expression[bool]) *AWS_ApiGateway_RequestValidator {
	t.ValidateRequestBody = v
	return t
}

// SetV__ValidateRequestBody updates property "ValidateRequestBody".
func (t *AWS_ApiGateway_RequestValidator) SetV__ValidateRequestBody(v bool) *AWS_ApiGateway_RequestValidator {
	t.ValidateRequestBody = cfz.V(v)
	return t
}

// Set__ValidateRequestParameters updates property "ValidateRequestParameters".
func (t *AWS_ApiGateway_RequestValidator) Set__ValidateRequestParameters(v cfz.Expression[bool]) *AWS_ApiGateway_RequestValidator {
	t.ValidateRequestParameters = v
	return t
}

// SetV__ValidateRequestParameters updates property "ValidateRequestParameters".
func (t *AWS_ApiGateway_RequestValidator) SetV__ValidateRequestParameters(v bool) *AWS_ApiGateway_RequestValidator {
	t.ValidateRequestParameters = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ApiGateway_RequestValidator) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__RequestValidatorId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: RequestValidatorId
func (t *AWS_ApiGateway_RequestValidator) GetAtt__RequestValidatorId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ApiGateway_RequestValidator__AttributesMap.RequestValidatorId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ApiGateway_RequestValidator) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__RequestValidatorId returns a conventionally configured output for an attribute of this resource.
// Attribute: RequestValidatorId
func (t *AWS_ApiGateway_RequestValidator) GetConventionalOutputAtt__RequestValidatorId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttRequestValidatorId", t.GetAtt__RequestValidatorId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ApiGateway_RequestValidator) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ApiGateway_RequestValidator

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

func (t *AWS_ApiGateway_RequestValidator) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
