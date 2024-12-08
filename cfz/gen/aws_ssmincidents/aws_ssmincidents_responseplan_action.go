// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ssmincidents

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SSMIncidents_ResponsePlan_Action__Type is the CloudFormation type for AWS::SSMIncidents::ResponsePlan.Action.
	AWS_SSMIncidents_ResponsePlan_Action__Type = "AWS::SSMIncidents::ResponsePlan.Action"
)

var (
	// AWS_SSMIncidents_ResponsePlan_Action__PropertiesMap reports all the CloudFormation properties for AWS::SSMIncidents::ResponsePlan.Action.
	AWS_SSMIncidents_ResponsePlan_Action__PropertiesMap = struct {
		SsmAutomation string
	}{
		SsmAutomation: "SsmAutomation",
	}

	// AWS_SSMIncidents_ResponsePlan_Action__PropertiesSlice reports all the CloudFormation properties for AWS::SSMIncidents::ResponsePlan.Action.
	AWS_SSMIncidents_ResponsePlan_Action__PropertiesSlice = []string{
		AWS_SSMIncidents_ResponsePlan_Action__PropertiesMap.SsmAutomation,
	}
)

// AWS_SSMIncidents_ResponsePlan_Action is a binding for AWS::SSMIncidents::ResponsePlan.Action.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-action.html
type AWS_SSMIncidents_ResponsePlan_Action struct {
	// SsmAutomation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-action.html#cfn-ssmincidents-responseplan-action-ssmautomation
	SsmAutomation cfz.Expression[AWS_SSMIncidents_ResponsePlan_SsmAutomation] `json:"SsmAutomation,omitempty"`
}

// New__AWS_SSMIncidents_ResponsePlan_Action initializes a new AWS_SSMIncidents_ResponsePlan_Action.
func New__AWS_SSMIncidents_ResponsePlan_Action() AWS_SSMIncidents_ResponsePlan_Action {
	return AWS_SSMIncidents_ResponsePlan_Action{}
}

// GetType returns the CloudFormation type.
func (AWS_SSMIncidents_ResponsePlan_Action) GetType() string {
	return AWS_SSMIncidents_ResponsePlan_Action__Type
}

// Set__SsmAutomation updates property "SsmAutomation".
func (t AWS_SSMIncidents_ResponsePlan_Action) Set__SsmAutomation(v cfz.Expression[AWS_SSMIncidents_ResponsePlan_SsmAutomation]) AWS_SSMIncidents_ResponsePlan_Action {
	t.SsmAutomation = v
	return t
}

// SetV__SsmAutomation updates property "SsmAutomation".
func (t AWS_SSMIncidents_ResponsePlan_Action) SetV__SsmAutomation(v AWS_SSMIncidents_ResponsePlan_SsmAutomation) AWS_SSMIncidents_ResponsePlan_Action {
	t.SsmAutomation = cfz.V(v)
	return t
}
