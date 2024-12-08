// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_customerprofiles

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CustomerProfiles_Integration_ServiceNowSourceProperties__Type is the CloudFormation type for AWS::CustomerProfiles::Integration.ServiceNowSourceProperties.
	AWS_CustomerProfiles_Integration_ServiceNowSourceProperties__Type = "AWS::CustomerProfiles::Integration.ServiceNowSourceProperties"
)

var (
	// AWS_CustomerProfiles_Integration_ServiceNowSourceProperties__PropertiesMap reports all the CloudFormation properties for AWS::CustomerProfiles::Integration.ServiceNowSourceProperties.
	AWS_CustomerProfiles_Integration_ServiceNowSourceProperties__PropertiesMap = struct {
		Object string
	}{
		Object: "Object",
	}

	// AWS_CustomerProfiles_Integration_ServiceNowSourceProperties__PropertiesSlice reports all the CloudFormation properties for AWS::CustomerProfiles::Integration.ServiceNowSourceProperties.
	AWS_CustomerProfiles_Integration_ServiceNowSourceProperties__PropertiesSlice = []string{
		AWS_CustomerProfiles_Integration_ServiceNowSourceProperties__PropertiesMap.Object,
	}
)

// AWS_CustomerProfiles_Integration_ServiceNowSourceProperties is a binding for AWS::CustomerProfiles::Integration.ServiceNowSourceProperties.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-servicenowsourceproperties.html
type AWS_CustomerProfiles_Integration_ServiceNowSourceProperties struct {
	// Object is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-servicenowsourceproperties.html#cfn-customerprofiles-integration-servicenowsourceproperties-object
	Object cfz.Expression[string] `json:"Object,omitempty"`
}

// New__AWS_CustomerProfiles_Integration_ServiceNowSourceProperties initializes a new AWS_CustomerProfiles_Integration_ServiceNowSourceProperties.
func New__AWS_CustomerProfiles_Integration_ServiceNowSourceProperties() AWS_CustomerProfiles_Integration_ServiceNowSourceProperties {
	return AWS_CustomerProfiles_Integration_ServiceNowSourceProperties{}
}

// GetType returns the CloudFormation type.
func (AWS_CustomerProfiles_Integration_ServiceNowSourceProperties) GetType() string {
	return AWS_CustomerProfiles_Integration_ServiceNowSourceProperties__Type
}

// Set__Object updates property "Object".
func (t AWS_CustomerProfiles_Integration_ServiceNowSourceProperties) Set__Object(v cfz.Expression[string]) AWS_CustomerProfiles_Integration_ServiceNowSourceProperties {
	t.Object = v
	return t
}

// SetV__Object updates property "Object".
func (t AWS_CustomerProfiles_Integration_ServiceNowSourceProperties) SetV__Object(v string) AWS_CustomerProfiles_Integration_ServiceNowSourceProperties {
	t.Object = cfz.V(v)
	return t
}
