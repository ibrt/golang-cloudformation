// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_inspector

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Inspector_AssessmentTarget)(nil)
	_ cfz.Resource                   = (*AWS_Inspector_AssessmentTarget)(nil)
)

const (
	// AWS_Inspector_AssessmentTarget__Type is the CloudFormation type for AWS::Inspector::AssessmentTarget.
	AWS_Inspector_AssessmentTarget__Type = "AWS::Inspector::AssessmentTarget"
)

var (
	// AWS_Inspector_AssessmentTarget__AttributesMap reports all the CloudFormation attributes for AWS::Inspector::AssessmentTarget.
	AWS_Inspector_AssessmentTarget__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Inspector_AssessmentTarget__AttributesSlice reports all the CloudFormation attributes for AWS::Inspector::AssessmentTarget.
	AWS_Inspector_AssessmentTarget__AttributesSlice = []string{
		AWS_Inspector_AssessmentTarget__AttributesMap.Arn,
	}
)

var (
	// AWS_Inspector_AssessmentTarget__PropertiesMap reports all the CloudFormation properties for AWS::Inspector::AssessmentTarget.
	AWS_Inspector_AssessmentTarget__PropertiesMap = struct {
		AssessmentTargetName string
		ResourceGroupArn     string
	}{
		AssessmentTargetName: "AssessmentTargetName",
		ResourceGroupArn:     "ResourceGroupArn",
	}

	// AWS_Inspector_AssessmentTarget__PropertiesSlice reports all the CloudFormation properties for AWS::Inspector::AssessmentTarget.
	AWS_Inspector_AssessmentTarget__PropertiesSlice = []string{
		AWS_Inspector_AssessmentTarget__PropertiesMap.AssessmentTargetName,
		AWS_Inspector_AssessmentTarget__PropertiesMap.ResourceGroupArn,
	}
)

// AWS_Inspector_AssessmentTarget is a binding for AWS::Inspector::AssessmentTarget.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttarget.html
type AWS_Inspector_AssessmentTarget struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AssessmentTargetName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttarget.html#cfn-inspector-assessmenttarget-assessmenttargetname
	AssessmentTargetName cfz.Expression[string] `json:"AssessmentTargetName,omitempty"`

	// ResourceGroupArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttarget.html#cfn-inspector-assessmenttarget-resourcegrouparn
	ResourceGroupArn cfz.Expression[string] `json:"ResourceGroupArn,omitempty"`
}

// New__AWS_Inspector_AssessmentTarget initializes a new *AWS_Inspector_AssessmentTarget.
func New__AWS_Inspector_AssessmentTarget(logicalName string) *AWS_Inspector_AssessmentTarget {
	return &AWS_Inspector_AssessmentTarget{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Inspector_AssessmentTarget) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Inspector_AssessmentTarget) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Inspector_AssessmentTarget) GetType() string {
	return AWS_Inspector_AssessmentTarget__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Inspector_AssessmentTarget) Set__LogicalName(v string) *AWS_Inspector_AssessmentTarget {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Inspector_AssessmentTarget) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Inspector_AssessmentTarget {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Inspector_AssessmentTarget) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Inspector_AssessmentTarget {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Inspector_AssessmentTarget) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Inspector_AssessmentTarget {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Inspector_AssessmentTarget) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Inspector_AssessmentTarget {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Inspector_AssessmentTarget) Set__RequestedOutputs(v []cfz.Output) *AWS_Inspector_AssessmentTarget {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Inspector_AssessmentTarget) Add__RequestedOutputs(v ...cfz.Output) *AWS_Inspector_AssessmentTarget {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AssessmentTargetName updates property "AssessmentTargetName".
func (t *AWS_Inspector_AssessmentTarget) Set__AssessmentTargetName(v cfz.Expression[string]) *AWS_Inspector_AssessmentTarget {
	t.AssessmentTargetName = v
	return t
}

// SetV__AssessmentTargetName updates property "AssessmentTargetName".
func (t *AWS_Inspector_AssessmentTarget) SetV__AssessmentTargetName(v string) *AWS_Inspector_AssessmentTarget {
	t.AssessmentTargetName = cfz.V(v)
	return t
}

// Set__ResourceGroupArn updates property "ResourceGroupArn".
func (t *AWS_Inspector_AssessmentTarget) Set__ResourceGroupArn(v cfz.Expression[string]) *AWS_Inspector_AssessmentTarget {
	t.ResourceGroupArn = v
	return t
}

// SetV__ResourceGroupArn updates property "ResourceGroupArn".
func (t *AWS_Inspector_AssessmentTarget) SetV__ResourceGroupArn(v string) *AWS_Inspector_AssessmentTarget {
	t.ResourceGroupArn = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Inspector_AssessmentTarget) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Inspector_AssessmentTarget) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Inspector_AssessmentTarget__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Inspector_AssessmentTarget) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Inspector_AssessmentTarget) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Inspector_AssessmentTarget) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Inspector_AssessmentTarget

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

func (t *AWS_Inspector_AssessmentTarget) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
