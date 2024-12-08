// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iam

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_IAM_UserPolicy)(nil)
	_ cfz.Resource                   = (*AWS_IAM_UserPolicy)(nil)
)

const (
	// AWS_IAM_UserPolicy__Type is the CloudFormation type for AWS::IAM::UserPolicy.
	AWS_IAM_UserPolicy__Type = "AWS::IAM::UserPolicy"
)

var (
	// AWS_IAM_UserPolicy__PropertiesMap reports all the CloudFormation properties for AWS::IAM::UserPolicy.
	AWS_IAM_UserPolicy__PropertiesMap = struct {
		PolicyDocument string
		PolicyName     string
		UserName       string
	}{
		PolicyDocument: "PolicyDocument",
		PolicyName:     "PolicyName",
		UserName:       "UserName",
	}

	// AWS_IAM_UserPolicy__PropertiesSlice reports all the CloudFormation properties for AWS::IAM::UserPolicy.
	AWS_IAM_UserPolicy__PropertiesSlice = []string{
		AWS_IAM_UserPolicy__PropertiesMap.PolicyDocument,
		AWS_IAM_UserPolicy__PropertiesMap.PolicyName,
		AWS_IAM_UserPolicy__PropertiesMap.UserName,
	}
)

// AWS_IAM_UserPolicy is a binding for AWS::IAM::UserPolicy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-userpolicy.html
type AWS_IAM_UserPolicy struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// PolicyDocument is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-userpolicy.html#cfn-iam-userpolicy-policydocument
	PolicyDocument cfz.Expression[json.RawMessage] `json:"PolicyDocument,omitempty"`

	// PolicyName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-userpolicy.html#cfn-iam-userpolicy-policyname
	PolicyName cfz.Expression[string] `json:"PolicyName,omitempty"`

	// UserName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-userpolicy.html#cfn-iam-userpolicy-username
	UserName cfz.Expression[string] `json:"UserName,omitempty"`
}

// New__AWS_IAM_UserPolicy initializes a new *AWS_IAM_UserPolicy.
func New__AWS_IAM_UserPolicy(logicalName string) *AWS_IAM_UserPolicy {
	return &AWS_IAM_UserPolicy{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_IAM_UserPolicy) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_IAM_UserPolicy) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_IAM_UserPolicy) GetType() string {
	return AWS_IAM_UserPolicy__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_IAM_UserPolicy) Set__LogicalName(v string) *AWS_IAM_UserPolicy {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_IAM_UserPolicy) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_IAM_UserPolicy {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_IAM_UserPolicy) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_IAM_UserPolicy {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_IAM_UserPolicy) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_IAM_UserPolicy {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_IAM_UserPolicy) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_IAM_UserPolicy {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_IAM_UserPolicy) Set__RequestedOutputs(v []cfz.Output) *AWS_IAM_UserPolicy {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_IAM_UserPolicy) Add__RequestedOutputs(v ...cfz.Output) *AWS_IAM_UserPolicy {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__PolicyDocument updates property "PolicyDocument".
func (t *AWS_IAM_UserPolicy) Set__PolicyDocument(v cfz.Expression[json.RawMessage]) *AWS_IAM_UserPolicy {
	t.PolicyDocument = v
	return t
}

// SetV__PolicyDocument updates property "PolicyDocument".
func (t *AWS_IAM_UserPolicy) SetV__PolicyDocument(v json.RawMessage) *AWS_IAM_UserPolicy {
	t.PolicyDocument = cfz.V(v)
	return t
}

// Set__PolicyName updates property "PolicyName".
func (t *AWS_IAM_UserPolicy) Set__PolicyName(v cfz.Expression[string]) *AWS_IAM_UserPolicy {
	t.PolicyName = v
	return t
}

// SetV__PolicyName updates property "PolicyName".
func (t *AWS_IAM_UserPolicy) SetV__PolicyName(v string) *AWS_IAM_UserPolicy {
	t.PolicyName = cfz.V(v)
	return t
}

// Set__UserName updates property "UserName".
func (t *AWS_IAM_UserPolicy) Set__UserName(v cfz.Expression[string]) *AWS_IAM_UserPolicy {
	t.UserName = v
	return t
}

// SetV__UserName updates property "UserName".
func (t *AWS_IAM_UserPolicy) SetV__UserName(v string) *AWS_IAM_UserPolicy {
	t.UserName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_IAM_UserPolicy) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_IAM_UserPolicy) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_IAM_UserPolicy) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_IAM_UserPolicy

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

func (t *AWS_IAM_UserPolicy) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
