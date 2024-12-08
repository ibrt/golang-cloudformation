// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_auditmanager

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AuditManager_Assessment_AssessmentReportsDestination__Type is the CloudFormation type for AWS::AuditManager::Assessment.AssessmentReportsDestination.
	AWS_AuditManager_Assessment_AssessmentReportsDestination__Type = "AWS::AuditManager::Assessment.AssessmentReportsDestination"
)

var (
	// AWS_AuditManager_Assessment_AssessmentReportsDestination__PropertiesMap reports all the CloudFormation properties for AWS::AuditManager::Assessment.AssessmentReportsDestination.
	AWS_AuditManager_Assessment_AssessmentReportsDestination__PropertiesMap = struct {
		Destination     string
		DestinationType string
	}{
		Destination:     "Destination",
		DestinationType: "DestinationType",
	}

	// AWS_AuditManager_Assessment_AssessmentReportsDestination__PropertiesSlice reports all the CloudFormation properties for AWS::AuditManager::Assessment.AssessmentReportsDestination.
	AWS_AuditManager_Assessment_AssessmentReportsDestination__PropertiesSlice = []string{
		AWS_AuditManager_Assessment_AssessmentReportsDestination__PropertiesMap.Destination,
		AWS_AuditManager_Assessment_AssessmentReportsDestination__PropertiesMap.DestinationType,
	}
)

// AWS_AuditManager_Assessment_AssessmentReportsDestination is a binding for AWS::AuditManager::Assessment.AssessmentReportsDestination.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-assessmentreportsdestination.html
type AWS_AuditManager_Assessment_AssessmentReportsDestination struct {
	// Destination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-assessmentreportsdestination.html#cfn-auditmanager-assessment-assessmentreportsdestination-destination
	Destination cfz.Expression[string] `json:"Destination,omitempty"`

	// DestinationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-assessmentreportsdestination.html#cfn-auditmanager-assessment-assessmentreportsdestination-destinationtype
	DestinationType cfz.Expression[string] `json:"DestinationType,omitempty"`
}

// New__AWS_AuditManager_Assessment_AssessmentReportsDestination initializes a new AWS_AuditManager_Assessment_AssessmentReportsDestination.
func New__AWS_AuditManager_Assessment_AssessmentReportsDestination() AWS_AuditManager_Assessment_AssessmentReportsDestination {
	return AWS_AuditManager_Assessment_AssessmentReportsDestination{}
}

// GetType returns the CloudFormation type.
func (AWS_AuditManager_Assessment_AssessmentReportsDestination) GetType() string {
	return AWS_AuditManager_Assessment_AssessmentReportsDestination__Type
}

// Set__Destination updates property "Destination".
func (t AWS_AuditManager_Assessment_AssessmentReportsDestination) Set__Destination(v cfz.Expression[string]) AWS_AuditManager_Assessment_AssessmentReportsDestination {
	t.Destination = v
	return t
}

// SetV__Destination updates property "Destination".
func (t AWS_AuditManager_Assessment_AssessmentReportsDestination) SetV__Destination(v string) AWS_AuditManager_Assessment_AssessmentReportsDestination {
	t.Destination = cfz.V(v)
	return t
}

// Set__DestinationType updates property "DestinationType".
func (t AWS_AuditManager_Assessment_AssessmentReportsDestination) Set__DestinationType(v cfz.Expression[string]) AWS_AuditManager_Assessment_AssessmentReportsDestination {
	t.DestinationType = v
	return t
}

// SetV__DestinationType updates property "DestinationType".
func (t AWS_AuditManager_Assessment_AssessmentReportsDestination) SetV__DestinationType(v string) AWS_AuditManager_Assessment_AssessmentReportsDestination {
	t.DestinationType = cfz.V(v)
	return t
}
