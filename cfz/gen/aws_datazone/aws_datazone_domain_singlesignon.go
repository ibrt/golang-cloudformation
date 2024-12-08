// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datazone

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataZone_Domain_SingleSignOn__Type is the CloudFormation type for AWS::DataZone::Domain.SingleSignOn.
	AWS_DataZone_Domain_SingleSignOn__Type = "AWS::DataZone::Domain.SingleSignOn"
)

var (
	// AWS_DataZone_Domain_SingleSignOn__PropertiesMap reports all the CloudFormation properties for AWS::DataZone::Domain.SingleSignOn.
	AWS_DataZone_Domain_SingleSignOn__PropertiesMap = struct {
		Type           string
		UserAssignment string
	}{
		Type:           "Type",
		UserAssignment: "UserAssignment",
	}

	// AWS_DataZone_Domain_SingleSignOn__PropertiesSlice reports all the CloudFormation properties for AWS::DataZone::Domain.SingleSignOn.
	AWS_DataZone_Domain_SingleSignOn__PropertiesSlice = []string{
		AWS_DataZone_Domain_SingleSignOn__PropertiesMap.Type,
		AWS_DataZone_Domain_SingleSignOn__PropertiesMap.UserAssignment,
	}
)

// AWS_DataZone_Domain_SingleSignOn is a binding for AWS::DataZone::Domain.SingleSignOn.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-domain-singlesignon.html
type AWS_DataZone_Domain_SingleSignOn struct {
	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-domain-singlesignon.html#cfn-datazone-domain-singlesignon-type
	Type cfz.Expression[string] `json:"Type,omitempty"`

	// UserAssignment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-domain-singlesignon.html#cfn-datazone-domain-singlesignon-userassignment
	UserAssignment cfz.Expression[string] `json:"UserAssignment,omitempty"`
}

// New__AWS_DataZone_Domain_SingleSignOn initializes a new AWS_DataZone_Domain_SingleSignOn.
func New__AWS_DataZone_Domain_SingleSignOn() AWS_DataZone_Domain_SingleSignOn {
	return AWS_DataZone_Domain_SingleSignOn{}
}

// GetType returns the CloudFormation type.
func (AWS_DataZone_Domain_SingleSignOn) GetType() string {
	return AWS_DataZone_Domain_SingleSignOn__Type
}

// Set__Type updates property "Type".
func (t AWS_DataZone_Domain_SingleSignOn) Set__Type(v cfz.Expression[string]) AWS_DataZone_Domain_SingleSignOn {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_DataZone_Domain_SingleSignOn) SetV__Type(v string) AWS_DataZone_Domain_SingleSignOn {
	t.Type = cfz.V(v)
	return t
}

// Set__UserAssignment updates property "UserAssignment".
func (t AWS_DataZone_Domain_SingleSignOn) Set__UserAssignment(v cfz.Expression[string]) AWS_DataZone_Domain_SingleSignOn {
	t.UserAssignment = v
	return t
}

// SetV__UserAssignment updates property "UserAssignment".
func (t AWS_DataZone_Domain_SingleSignOn) SetV__UserAssignment(v string) AWS_DataZone_Domain_SingleSignOn {
	t.UserAssignment = cfz.V(v)
	return t
}
