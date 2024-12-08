// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dlm

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DLM_LifecyclePolicy_Exclusions__Type is the CloudFormation type for AWS::DLM::LifecyclePolicy.Exclusions.
	AWS_DLM_LifecyclePolicy_Exclusions__Type = "AWS::DLM::LifecyclePolicy.Exclusions"
)

var (
	// AWS_DLM_LifecyclePolicy_Exclusions__PropertiesMap reports all the CloudFormation properties for AWS::DLM::LifecyclePolicy.Exclusions.
	AWS_DLM_LifecyclePolicy_Exclusions__PropertiesMap = struct {
		ExcludeBootVolumes string
		ExcludeTags        string
		ExcludeVolumeTypes string
	}{
		ExcludeBootVolumes: "ExcludeBootVolumes",
		ExcludeTags:        "ExcludeTags",
		ExcludeVolumeTypes: "ExcludeVolumeTypes",
	}

	// AWS_DLM_LifecyclePolicy_Exclusions__PropertiesSlice reports all the CloudFormation properties for AWS::DLM::LifecyclePolicy.Exclusions.
	AWS_DLM_LifecyclePolicy_Exclusions__PropertiesSlice = []string{
		AWS_DLM_LifecyclePolicy_Exclusions__PropertiesMap.ExcludeBootVolumes,
		AWS_DLM_LifecyclePolicy_Exclusions__PropertiesMap.ExcludeTags,
		AWS_DLM_LifecyclePolicy_Exclusions__PropertiesMap.ExcludeVolumeTypes,
	}
)

// AWS_DLM_LifecyclePolicy_Exclusions is a binding for AWS::DLM::LifecyclePolicy.Exclusions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-exclusions.html
type AWS_DLM_LifecyclePolicy_Exclusions struct {
	// ExcludeBootVolumes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-exclusions.html#cfn-dlm-lifecyclepolicy-exclusions-excludebootvolumes
	ExcludeBootVolumes cfz.Expression[bool] `json:"ExcludeBootVolumes,omitempty"`

	// ExcludeTags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-exclusions.html#cfn-dlm-lifecyclepolicy-exclusions-excludetags
	ExcludeTags cfz.Expression[json.RawMessage] `json:"ExcludeTags,omitempty"`

	// ExcludeVolumeTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-exclusions.html#cfn-dlm-lifecyclepolicy-exclusions-excludevolumetypes
	ExcludeVolumeTypes cfz.Expression[json.RawMessage] `json:"ExcludeVolumeTypes,omitempty"`
}

// New__AWS_DLM_LifecyclePolicy_Exclusions initializes a new AWS_DLM_LifecyclePolicy_Exclusions.
func New__AWS_DLM_LifecyclePolicy_Exclusions() AWS_DLM_LifecyclePolicy_Exclusions {
	return AWS_DLM_LifecyclePolicy_Exclusions{}
}

// GetType returns the CloudFormation type.
func (AWS_DLM_LifecyclePolicy_Exclusions) GetType() string {
	return AWS_DLM_LifecyclePolicy_Exclusions__Type
}

// Set__ExcludeBootVolumes updates property "ExcludeBootVolumes".
func (t AWS_DLM_LifecyclePolicy_Exclusions) Set__ExcludeBootVolumes(v cfz.Expression[bool]) AWS_DLM_LifecyclePolicy_Exclusions {
	t.ExcludeBootVolumes = v
	return t
}

// SetV__ExcludeBootVolumes updates property "ExcludeBootVolumes".
func (t AWS_DLM_LifecyclePolicy_Exclusions) SetV__ExcludeBootVolumes(v bool) AWS_DLM_LifecyclePolicy_Exclusions {
	t.ExcludeBootVolumes = cfz.V(v)
	return t
}

// Set__ExcludeTags updates property "ExcludeTags".
func (t AWS_DLM_LifecyclePolicy_Exclusions) Set__ExcludeTags(v cfz.Expression[json.RawMessage]) AWS_DLM_LifecyclePolicy_Exclusions {
	t.ExcludeTags = v
	return t
}

// SetV__ExcludeTags updates property "ExcludeTags".
func (t AWS_DLM_LifecyclePolicy_Exclusions) SetV__ExcludeTags(v json.RawMessage) AWS_DLM_LifecyclePolicy_Exclusions {
	t.ExcludeTags = cfz.V(v)
	return t
}

// Set__ExcludeVolumeTypes updates property "ExcludeVolumeTypes".
func (t AWS_DLM_LifecyclePolicy_Exclusions) Set__ExcludeVolumeTypes(v cfz.Expression[json.RawMessage]) AWS_DLM_LifecyclePolicy_Exclusions {
	t.ExcludeVolumeTypes = v
	return t
}

// SetV__ExcludeVolumeTypes updates property "ExcludeVolumeTypes".
func (t AWS_DLM_LifecyclePolicy_Exclusions) SetV__ExcludeVolumeTypes(v json.RawMessage) AWS_DLM_LifecyclePolicy_Exclusions {
	t.ExcludeVolumeTypes = cfz.V(v)
	return t
}
