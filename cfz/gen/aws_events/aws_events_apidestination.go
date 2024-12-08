// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_events

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Events_ApiDestination)(nil)
	_ cfz.Resource                   = (*AWS_Events_ApiDestination)(nil)
)

const (
	// AWS_Events_ApiDestination__Type is the CloudFormation type for AWS::Events::ApiDestination.
	AWS_Events_ApiDestination__Type = "AWS::Events::ApiDestination"
)

var (
	// AWS_Events_ApiDestination__AttributesMap reports all the CloudFormation attributes for AWS::Events::ApiDestination.
	AWS_Events_ApiDestination__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Events_ApiDestination__AttributesSlice reports all the CloudFormation attributes for AWS::Events::ApiDestination.
	AWS_Events_ApiDestination__AttributesSlice = []string{
		AWS_Events_ApiDestination__AttributesMap.Arn,
	}
)

var (
	// AWS_Events_ApiDestination__PropertiesMap reports all the CloudFormation properties for AWS::Events::ApiDestination.
	AWS_Events_ApiDestination__PropertiesMap = struct {
		ConnectionArn                string
		Description                  string
		HttpMethod                   string
		InvocationEndpoint           string
		InvocationRateLimitPerSecond string
		Name                         string
	}{
		ConnectionArn:                "ConnectionArn",
		Description:                  "Description",
		HttpMethod:                   "HttpMethod",
		InvocationEndpoint:           "InvocationEndpoint",
		InvocationRateLimitPerSecond: "InvocationRateLimitPerSecond",
		Name:                         "Name",
	}

	// AWS_Events_ApiDestination__PropertiesSlice reports all the CloudFormation properties for AWS::Events::ApiDestination.
	AWS_Events_ApiDestination__PropertiesSlice = []string{
		AWS_Events_ApiDestination__PropertiesMap.ConnectionArn,
		AWS_Events_ApiDestination__PropertiesMap.Description,
		AWS_Events_ApiDestination__PropertiesMap.HttpMethod,
		AWS_Events_ApiDestination__PropertiesMap.InvocationEndpoint,
		AWS_Events_ApiDestination__PropertiesMap.InvocationRateLimitPerSecond,
		AWS_Events_ApiDestination__PropertiesMap.Name,
	}
)

// AWS_Events_ApiDestination is a binding for AWS::Events::ApiDestination.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html
type AWS_Events_ApiDestination struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ConnectionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-connectionarn
	ConnectionArn cfz.Expression[string] `json:"ConnectionArn,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// HttpMethod is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-httpmethod
	HttpMethod cfz.Expression[string] `json:"HttpMethod,omitempty"`

	// InvocationEndpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-invocationendpoint
	InvocationEndpoint cfz.Expression[string] `json:"InvocationEndpoint,omitempty"`

	// InvocationRateLimitPerSecond is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-invocationratelimitpersecond
	InvocationRateLimitPerSecond cfz.Expression[int32] `json:"InvocationRateLimitPerSecond,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-apidestination.html#cfn-events-apidestination-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_Events_ApiDestination initializes a new *AWS_Events_ApiDestination.
func New__AWS_Events_ApiDestination(logicalName string) *AWS_Events_ApiDestination {
	return &AWS_Events_ApiDestination{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Events_ApiDestination) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Events_ApiDestination) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Events_ApiDestination) GetType() string {
	return AWS_Events_ApiDestination__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Events_ApiDestination) Set__LogicalName(v string) *AWS_Events_ApiDestination {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Events_ApiDestination) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Events_ApiDestination {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Events_ApiDestination) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Events_ApiDestination {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Events_ApiDestination) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Events_ApiDestination {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Events_ApiDestination) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Events_ApiDestination {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Events_ApiDestination) Set__RequestedOutputs(v []cfz.Output) *AWS_Events_ApiDestination {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Events_ApiDestination) Add__RequestedOutputs(v ...cfz.Output) *AWS_Events_ApiDestination {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ConnectionArn updates property "ConnectionArn".
func (t *AWS_Events_ApiDestination) Set__ConnectionArn(v cfz.Expression[string]) *AWS_Events_ApiDestination {
	t.ConnectionArn = v
	return t
}

// SetV__ConnectionArn updates property "ConnectionArn".
func (t *AWS_Events_ApiDestination) SetV__ConnectionArn(v string) *AWS_Events_ApiDestination {
	t.ConnectionArn = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Events_ApiDestination) Set__Description(v cfz.Expression[string]) *AWS_Events_ApiDestination {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Events_ApiDestination) SetV__Description(v string) *AWS_Events_ApiDestination {
	t.Description = cfz.V(v)
	return t
}

// Set__HttpMethod updates property "HttpMethod".
func (t *AWS_Events_ApiDestination) Set__HttpMethod(v cfz.Expression[string]) *AWS_Events_ApiDestination {
	t.HttpMethod = v
	return t
}

// SetV__HttpMethod updates property "HttpMethod".
func (t *AWS_Events_ApiDestination) SetV__HttpMethod(v string) *AWS_Events_ApiDestination {
	t.HttpMethod = cfz.V(v)
	return t
}

// Set__InvocationEndpoint updates property "InvocationEndpoint".
func (t *AWS_Events_ApiDestination) Set__InvocationEndpoint(v cfz.Expression[string]) *AWS_Events_ApiDestination {
	t.InvocationEndpoint = v
	return t
}

// SetV__InvocationEndpoint updates property "InvocationEndpoint".
func (t *AWS_Events_ApiDestination) SetV__InvocationEndpoint(v string) *AWS_Events_ApiDestination {
	t.InvocationEndpoint = cfz.V(v)
	return t
}

// Set__InvocationRateLimitPerSecond updates property "InvocationRateLimitPerSecond".
func (t *AWS_Events_ApiDestination) Set__InvocationRateLimitPerSecond(v cfz.Expression[int32]) *AWS_Events_ApiDestination {
	t.InvocationRateLimitPerSecond = v
	return t
}

// SetV__InvocationRateLimitPerSecond updates property "InvocationRateLimitPerSecond".
func (t *AWS_Events_ApiDestination) SetV__InvocationRateLimitPerSecond(v int32) *AWS_Events_ApiDestination {
	t.InvocationRateLimitPerSecond = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Events_ApiDestination) Set__Name(v cfz.Expression[string]) *AWS_Events_ApiDestination {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Events_ApiDestination) SetV__Name(v string) *AWS_Events_ApiDestination {
	t.Name = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Events_ApiDestination) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Events_ApiDestination) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Events_ApiDestination__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Events_ApiDestination) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Events_ApiDestination) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Events_ApiDestination) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Events_ApiDestination

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

func (t *AWS_Events_ApiDestination) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
