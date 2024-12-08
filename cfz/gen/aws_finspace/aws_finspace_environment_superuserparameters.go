// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_finspace

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FinSpace_Environment_SuperuserParameters__Type is the CloudFormation type for AWS::FinSpace::Environment.SuperuserParameters.
	AWS_FinSpace_Environment_SuperuserParameters__Type = "AWS::FinSpace::Environment.SuperuserParameters"
)

var (
	// AWS_FinSpace_Environment_SuperuserParameters__PropertiesMap reports all the CloudFormation properties for AWS::FinSpace::Environment.SuperuserParameters.
	AWS_FinSpace_Environment_SuperuserParameters__PropertiesMap = struct {
		EmailAddress string
		FirstName    string
		LastName     string
	}{
		EmailAddress: "EmailAddress",
		FirstName:    "FirstName",
		LastName:     "LastName",
	}

	// AWS_FinSpace_Environment_SuperuserParameters__PropertiesSlice reports all the CloudFormation properties for AWS::FinSpace::Environment.SuperuserParameters.
	AWS_FinSpace_Environment_SuperuserParameters__PropertiesSlice = []string{
		AWS_FinSpace_Environment_SuperuserParameters__PropertiesMap.EmailAddress,
		AWS_FinSpace_Environment_SuperuserParameters__PropertiesMap.FirstName,
		AWS_FinSpace_Environment_SuperuserParameters__PropertiesMap.LastName,
	}
)

// AWS_FinSpace_Environment_SuperuserParameters is a binding for AWS::FinSpace::Environment.SuperuserParameters.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html
type AWS_FinSpace_Environment_SuperuserParameters struct {
	// EmailAddress is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html#cfn-finspace-environment-superuserparameters-emailaddress
	EmailAddress cfz.Expression[string] `json:"EmailAddress,omitempty"`

	// FirstName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html#cfn-finspace-environment-superuserparameters-firstname
	FirstName cfz.Expression[string] `json:"FirstName,omitempty"`

	// LastName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-superuserparameters.html#cfn-finspace-environment-superuserparameters-lastname
	LastName cfz.Expression[string] `json:"LastName,omitempty"`
}

// New__AWS_FinSpace_Environment_SuperuserParameters initializes a new AWS_FinSpace_Environment_SuperuserParameters.
func New__AWS_FinSpace_Environment_SuperuserParameters() AWS_FinSpace_Environment_SuperuserParameters {
	return AWS_FinSpace_Environment_SuperuserParameters{}
}

// GetType returns the CloudFormation type.
func (AWS_FinSpace_Environment_SuperuserParameters) GetType() string {
	return AWS_FinSpace_Environment_SuperuserParameters__Type
}

// Set__EmailAddress updates property "EmailAddress".
func (t AWS_FinSpace_Environment_SuperuserParameters) Set__EmailAddress(v cfz.Expression[string]) AWS_FinSpace_Environment_SuperuserParameters {
	t.EmailAddress = v
	return t
}

// SetV__EmailAddress updates property "EmailAddress".
func (t AWS_FinSpace_Environment_SuperuserParameters) SetV__EmailAddress(v string) AWS_FinSpace_Environment_SuperuserParameters {
	t.EmailAddress = cfz.V(v)
	return t
}

// Set__FirstName updates property "FirstName".
func (t AWS_FinSpace_Environment_SuperuserParameters) Set__FirstName(v cfz.Expression[string]) AWS_FinSpace_Environment_SuperuserParameters {
	t.FirstName = v
	return t
}

// SetV__FirstName updates property "FirstName".
func (t AWS_FinSpace_Environment_SuperuserParameters) SetV__FirstName(v string) AWS_FinSpace_Environment_SuperuserParameters {
	t.FirstName = cfz.V(v)
	return t
}

// Set__LastName updates property "LastName".
func (t AWS_FinSpace_Environment_SuperuserParameters) Set__LastName(v cfz.Expression[string]) AWS_FinSpace_Environment_SuperuserParameters {
	t.LastName = v
	return t
}

// SetV__LastName updates property "LastName".
func (t AWS_FinSpace_Environment_SuperuserParameters) SetV__LastName(v string) AWS_FinSpace_Environment_SuperuserParameters {
	t.LastName = cfz.V(v)
	return t
}
