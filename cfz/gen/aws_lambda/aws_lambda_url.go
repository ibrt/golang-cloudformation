// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lambda

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Lambda_Url)(nil)
	_ cfz.Resource                   = (*AWS_Lambda_Url)(nil)
)

const (
	// AWS_Lambda_Url__Type is the CloudFormation type for AWS::Lambda::Url.
	AWS_Lambda_Url__Type = "AWS::Lambda::Url"
)

var (
	// AWS_Lambda_Url__AttributesMap reports all the CloudFormation attributes for AWS::Lambda::Url.
	AWS_Lambda_Url__AttributesMap = struct {
		FunctionArn string
		FunctionUrl string
	}{
		FunctionArn: "FunctionArn",
		FunctionUrl: "FunctionUrl",
	}

	// AWS_Lambda_Url__AttributesSlice reports all the CloudFormation attributes for AWS::Lambda::Url.
	AWS_Lambda_Url__AttributesSlice = []string{
		AWS_Lambda_Url__AttributesMap.FunctionArn,
		AWS_Lambda_Url__AttributesMap.FunctionUrl,
	}
)

var (
	// AWS_Lambda_Url__PropertiesMap reports all the CloudFormation properties for AWS::Lambda::Url.
	AWS_Lambda_Url__PropertiesMap = struct {
		AuthType          string
		Cors              string
		InvokeMode        string
		Qualifier         string
		TargetFunctionArn string
	}{
		AuthType:          "AuthType",
		Cors:              "Cors",
		InvokeMode:        "InvokeMode",
		Qualifier:         "Qualifier",
		TargetFunctionArn: "TargetFunctionArn",
	}

	// AWS_Lambda_Url__PropertiesSlice reports all the CloudFormation properties for AWS::Lambda::Url.
	AWS_Lambda_Url__PropertiesSlice = []string{
		AWS_Lambda_Url__PropertiesMap.AuthType,
		AWS_Lambda_Url__PropertiesMap.Cors,
		AWS_Lambda_Url__PropertiesMap.InvokeMode,
		AWS_Lambda_Url__PropertiesMap.Qualifier,
		AWS_Lambda_Url__PropertiesMap.TargetFunctionArn,
	}
)

// AWS_Lambda_Url is a binding for AWS::Lambda::Url.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html
type AWS_Lambda_Url struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AuthType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-authtype
	AuthType cfz.Expression[string] `json:"AuthType,omitempty"`

	// Cors is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-cors
	Cors cfz.Expression[AWS_Lambda_Url_Cors] `json:"Cors,omitempty"`

	// InvokeMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-invokemode
	InvokeMode cfz.Expression[string] `json:"InvokeMode,omitempty"`

	// Qualifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-qualifier
	Qualifier cfz.Expression[string] `json:"Qualifier,omitempty"`

	// TargetFunctionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-targetfunctionarn
	TargetFunctionArn cfz.Expression[string] `json:"TargetFunctionArn,omitempty"`
}

// New__AWS_Lambda_Url initializes a new *AWS_Lambda_Url.
func New__AWS_Lambda_Url(logicalName string) *AWS_Lambda_Url {
	return &AWS_Lambda_Url{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Lambda_Url) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Lambda_Url) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Lambda_Url) GetType() string {
	return AWS_Lambda_Url__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Lambda_Url) Set__LogicalName(v string) *AWS_Lambda_Url {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Lambda_Url) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Lambda_Url {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Lambda_Url) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Lambda_Url {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Lambda_Url) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Lambda_Url {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Lambda_Url) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Lambda_Url {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Lambda_Url) Set__RequestedOutputs(v []cfz.Output) *AWS_Lambda_Url {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Lambda_Url) Add__RequestedOutputs(v ...cfz.Output) *AWS_Lambda_Url {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AuthType updates property "AuthType".
func (t *AWS_Lambda_Url) Set__AuthType(v cfz.Expression[string]) *AWS_Lambda_Url {
	t.AuthType = v
	return t
}

// SetV__AuthType updates property "AuthType".
func (t *AWS_Lambda_Url) SetV__AuthType(v string) *AWS_Lambda_Url {
	t.AuthType = cfz.V(v)
	return t
}

// Set__Cors updates property "Cors".
func (t *AWS_Lambda_Url) Set__Cors(v cfz.Expression[AWS_Lambda_Url_Cors]) *AWS_Lambda_Url {
	t.Cors = v
	return t
}

// SetV__Cors updates property "Cors".
func (t *AWS_Lambda_Url) SetV__Cors(v AWS_Lambda_Url_Cors) *AWS_Lambda_Url {
	t.Cors = cfz.V(v)
	return t
}

// Set__InvokeMode updates property "InvokeMode".
func (t *AWS_Lambda_Url) Set__InvokeMode(v cfz.Expression[string]) *AWS_Lambda_Url {
	t.InvokeMode = v
	return t
}

// SetV__InvokeMode updates property "InvokeMode".
func (t *AWS_Lambda_Url) SetV__InvokeMode(v string) *AWS_Lambda_Url {
	t.InvokeMode = cfz.V(v)
	return t
}

// Set__Qualifier updates property "Qualifier".
func (t *AWS_Lambda_Url) Set__Qualifier(v cfz.Expression[string]) *AWS_Lambda_Url {
	t.Qualifier = v
	return t
}

// SetV__Qualifier updates property "Qualifier".
func (t *AWS_Lambda_Url) SetV__Qualifier(v string) *AWS_Lambda_Url {
	t.Qualifier = cfz.V(v)
	return t
}

// Set__TargetFunctionArn updates property "TargetFunctionArn".
func (t *AWS_Lambda_Url) Set__TargetFunctionArn(v cfz.Expression[string]) *AWS_Lambda_Url {
	t.TargetFunctionArn = v
	return t
}

// SetV__TargetFunctionArn updates property "TargetFunctionArn".
func (t *AWS_Lambda_Url) SetV__TargetFunctionArn(v string) *AWS_Lambda_Url {
	t.TargetFunctionArn = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Lambda_Url) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__FunctionArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: FunctionArn
func (t *AWS_Lambda_Url) GetAtt__FunctionArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Lambda_Url__AttributesMap.FunctionArn))
}

// GetAtt__FunctionUrl returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: FunctionUrl
func (t *AWS_Lambda_Url) GetAtt__FunctionUrl() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Lambda_Url__AttributesMap.FunctionUrl))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Lambda_Url) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__FunctionArn returns a conventionally configured output for an attribute of this resource.
// Attribute: FunctionArn
func (t *AWS_Lambda_Url) GetConventionalOutputAtt__FunctionArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFunctionArn", t.GetAtt__FunctionArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__FunctionUrl returns a conventionally configured output for an attribute of this resource.
// Attribute: FunctionUrl
func (t *AWS_Lambda_Url) GetConventionalOutputAtt__FunctionUrl(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFunctionUrl", t.GetAtt__FunctionUrl())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Lambda_Url) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Lambda_Url

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

func (t *AWS_Lambda_Url) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
