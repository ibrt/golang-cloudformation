// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_athena

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Athena_WorkGroup_EngineVersion__Type is the CloudFormation type for AWS::Athena::WorkGroup.EngineVersion.
	AWS_Athena_WorkGroup_EngineVersion__Type = "AWS::Athena::WorkGroup.EngineVersion"
)

var (
	// AWS_Athena_WorkGroup_EngineVersion__PropertiesMap reports all the CloudFormation properties for AWS::Athena::WorkGroup.EngineVersion.
	AWS_Athena_WorkGroup_EngineVersion__PropertiesMap = struct {
		EffectiveEngineVersion string
		SelectedEngineVersion  string
	}{
		EffectiveEngineVersion: "EffectiveEngineVersion",
		SelectedEngineVersion:  "SelectedEngineVersion",
	}

	// AWS_Athena_WorkGroup_EngineVersion__PropertiesSlice reports all the CloudFormation properties for AWS::Athena::WorkGroup.EngineVersion.
	AWS_Athena_WorkGroup_EngineVersion__PropertiesSlice = []string{
		AWS_Athena_WorkGroup_EngineVersion__PropertiesMap.EffectiveEngineVersion,
		AWS_Athena_WorkGroup_EngineVersion__PropertiesMap.SelectedEngineVersion,
	}
)

// AWS_Athena_WorkGroup_EngineVersion is a binding for AWS::Athena::WorkGroup.EngineVersion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineversion.html
type AWS_Athena_WorkGroup_EngineVersion struct {
	// EffectiveEngineVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineversion.html#cfn-athena-workgroup-engineversion-effectiveengineversion
	EffectiveEngineVersion cfz.Expression[string] `json:"EffectiveEngineVersion,omitempty"`

	// SelectedEngineVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineversion.html#cfn-athena-workgroup-engineversion-selectedengineversion
	SelectedEngineVersion cfz.Expression[string] `json:"SelectedEngineVersion,omitempty"`
}

// New__AWS_Athena_WorkGroup_EngineVersion initializes a new AWS_Athena_WorkGroup_EngineVersion.
func New__AWS_Athena_WorkGroup_EngineVersion() AWS_Athena_WorkGroup_EngineVersion {
	return AWS_Athena_WorkGroup_EngineVersion{}
}

// GetType returns the CloudFormation type.
func (AWS_Athena_WorkGroup_EngineVersion) GetType() string {
	return AWS_Athena_WorkGroup_EngineVersion__Type
}

// Set__EffectiveEngineVersion updates property "EffectiveEngineVersion".
func (t AWS_Athena_WorkGroup_EngineVersion) Set__EffectiveEngineVersion(v cfz.Expression[string]) AWS_Athena_WorkGroup_EngineVersion {
	t.EffectiveEngineVersion = v
	return t
}

// SetV__EffectiveEngineVersion updates property "EffectiveEngineVersion".
func (t AWS_Athena_WorkGroup_EngineVersion) SetV__EffectiveEngineVersion(v string) AWS_Athena_WorkGroup_EngineVersion {
	t.EffectiveEngineVersion = cfz.V(v)
	return t
}

// Set__SelectedEngineVersion updates property "SelectedEngineVersion".
func (t AWS_Athena_WorkGroup_EngineVersion) Set__SelectedEngineVersion(v cfz.Expression[string]) AWS_Athena_WorkGroup_EngineVersion {
	t.SelectedEngineVersion = v
	return t
}

// SetV__SelectedEngineVersion updates property "SelectedEngineVersion".
func (t AWS_Athena_WorkGroup_EngineVersion) SetV__SelectedEngineVersion(v string) AWS_Athena_WorkGroup_EngineVersion {
	t.SelectedEngineVersion = cfz.V(v)
	return t
}
