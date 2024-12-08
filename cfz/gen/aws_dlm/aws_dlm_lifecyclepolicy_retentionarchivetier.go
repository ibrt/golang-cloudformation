// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dlm

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DLM_LifecyclePolicy_RetentionArchiveTier__Type is the CloudFormation type for AWS::DLM::LifecyclePolicy.RetentionArchiveTier.
	AWS_DLM_LifecyclePolicy_RetentionArchiveTier__Type = "AWS::DLM::LifecyclePolicy.RetentionArchiveTier"
)

var (
	// AWS_DLM_LifecyclePolicy_RetentionArchiveTier__PropertiesMap reports all the CloudFormation properties for AWS::DLM::LifecyclePolicy.RetentionArchiveTier.
	AWS_DLM_LifecyclePolicy_RetentionArchiveTier__PropertiesMap = struct {
		Count        string
		Interval     string
		IntervalUnit string
	}{
		Count:        "Count",
		Interval:     "Interval",
		IntervalUnit: "IntervalUnit",
	}

	// AWS_DLM_LifecyclePolicy_RetentionArchiveTier__PropertiesSlice reports all the CloudFormation properties for AWS::DLM::LifecyclePolicy.RetentionArchiveTier.
	AWS_DLM_LifecyclePolicy_RetentionArchiveTier__PropertiesSlice = []string{
		AWS_DLM_LifecyclePolicy_RetentionArchiveTier__PropertiesMap.Count,
		AWS_DLM_LifecyclePolicy_RetentionArchiveTier__PropertiesMap.Interval,
		AWS_DLM_LifecyclePolicy_RetentionArchiveTier__PropertiesMap.IntervalUnit,
	}
)

// AWS_DLM_LifecyclePolicy_RetentionArchiveTier is a binding for AWS::DLM::LifecyclePolicy.RetentionArchiveTier.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retentionarchivetier.html
type AWS_DLM_LifecyclePolicy_RetentionArchiveTier struct {
	// Count is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retentionarchivetier.html#cfn-dlm-lifecyclepolicy-retentionarchivetier-count
	Count cfz.Expression[int32] `json:"Count,omitempty"`

	// Interval is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retentionarchivetier.html#cfn-dlm-lifecyclepolicy-retentionarchivetier-interval
	Interval cfz.Expression[int32] `json:"Interval,omitempty"`

	// IntervalUnit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retentionarchivetier.html#cfn-dlm-lifecyclepolicy-retentionarchivetier-intervalunit
	IntervalUnit cfz.Expression[string] `json:"IntervalUnit,omitempty"`
}

// New__AWS_DLM_LifecyclePolicy_RetentionArchiveTier initializes a new AWS_DLM_LifecyclePolicy_RetentionArchiveTier.
func New__AWS_DLM_LifecyclePolicy_RetentionArchiveTier() AWS_DLM_LifecyclePolicy_RetentionArchiveTier {
	return AWS_DLM_LifecyclePolicy_RetentionArchiveTier{}
}

// GetType returns the CloudFormation type.
func (AWS_DLM_LifecyclePolicy_RetentionArchiveTier) GetType() string {
	return AWS_DLM_LifecyclePolicy_RetentionArchiveTier__Type
}

// Set__Count updates property "Count".
func (t AWS_DLM_LifecyclePolicy_RetentionArchiveTier) Set__Count(v cfz.Expression[int32]) AWS_DLM_LifecyclePolicy_RetentionArchiveTier {
	t.Count = v
	return t
}

// SetV__Count updates property "Count".
func (t AWS_DLM_LifecyclePolicy_RetentionArchiveTier) SetV__Count(v int32) AWS_DLM_LifecyclePolicy_RetentionArchiveTier {
	t.Count = cfz.V(v)
	return t
}

// Set__Interval updates property "Interval".
func (t AWS_DLM_LifecyclePolicy_RetentionArchiveTier) Set__Interval(v cfz.Expression[int32]) AWS_DLM_LifecyclePolicy_RetentionArchiveTier {
	t.Interval = v
	return t
}

// SetV__Interval updates property "Interval".
func (t AWS_DLM_LifecyclePolicy_RetentionArchiveTier) SetV__Interval(v int32) AWS_DLM_LifecyclePolicy_RetentionArchiveTier {
	t.Interval = cfz.V(v)
	return t
}

// Set__IntervalUnit updates property "IntervalUnit".
func (t AWS_DLM_LifecyclePolicy_RetentionArchiveTier) Set__IntervalUnit(v cfz.Expression[string]) AWS_DLM_LifecyclePolicy_RetentionArchiveTier {
	t.IntervalUnit = v
	return t
}

// SetV__IntervalUnit updates property "IntervalUnit".
func (t AWS_DLM_LifecyclePolicy_RetentionArchiveTier) SetV__IntervalUnit(v string) AWS_DLM_LifecyclePolicy_RetentionArchiveTier {
	t.IntervalUnit = cfz.V(v)
	return t
}
