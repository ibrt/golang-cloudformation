// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_customerprofiles

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CustomerProfiles_Integration_TriggerProperties__Type is the CloudFormation type for AWS::CustomerProfiles::Integration.TriggerProperties.
	AWS_CustomerProfiles_Integration_TriggerProperties__Type = "AWS::CustomerProfiles::Integration.TriggerProperties"
)

var (
	// AWS_CustomerProfiles_Integration_TriggerProperties__PropertiesMap reports all the CloudFormation properties for AWS::CustomerProfiles::Integration.TriggerProperties.
	AWS_CustomerProfiles_Integration_TriggerProperties__PropertiesMap = struct {
		Scheduled string
	}{
		Scheduled: "Scheduled",
	}

	// AWS_CustomerProfiles_Integration_TriggerProperties__PropertiesSlice reports all the CloudFormation properties for AWS::CustomerProfiles::Integration.TriggerProperties.
	AWS_CustomerProfiles_Integration_TriggerProperties__PropertiesSlice = []string{
		AWS_CustomerProfiles_Integration_TriggerProperties__PropertiesMap.Scheduled,
	}
)

// AWS_CustomerProfiles_Integration_TriggerProperties is a binding for AWS::CustomerProfiles::Integration.TriggerProperties.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-triggerproperties.html
type AWS_CustomerProfiles_Integration_TriggerProperties struct {
	// Scheduled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-triggerproperties.html#cfn-customerprofiles-integration-triggerproperties-scheduled
	Scheduled cfz.Expression[AWS_CustomerProfiles_Integration_ScheduledTriggerProperties] `json:"Scheduled,omitempty"`
}

// New__AWS_CustomerProfiles_Integration_TriggerProperties initializes a new AWS_CustomerProfiles_Integration_TriggerProperties.
func New__AWS_CustomerProfiles_Integration_TriggerProperties() AWS_CustomerProfiles_Integration_TriggerProperties {
	return AWS_CustomerProfiles_Integration_TriggerProperties{}
}

// GetType returns the CloudFormation type.
func (AWS_CustomerProfiles_Integration_TriggerProperties) GetType() string {
	return AWS_CustomerProfiles_Integration_TriggerProperties__Type
}

// Set__Scheduled updates property "Scheduled".
func (t AWS_CustomerProfiles_Integration_TriggerProperties) Set__Scheduled(v cfz.Expression[AWS_CustomerProfiles_Integration_ScheduledTriggerProperties]) AWS_CustomerProfiles_Integration_TriggerProperties {
	t.Scheduled = v
	return t
}

// SetV__Scheduled updates property "Scheduled".
func (t AWS_CustomerProfiles_Integration_TriggerProperties) SetV__Scheduled(v AWS_CustomerProfiles_Integration_ScheduledTriggerProperties) AWS_CustomerProfiles_Integration_TriggerProperties {
	t.Scheduled = cfz.V(v)
	return t
}
