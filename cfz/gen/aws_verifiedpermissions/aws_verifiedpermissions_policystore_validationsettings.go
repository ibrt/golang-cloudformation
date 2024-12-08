// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_verifiedpermissions

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_VerifiedPermissions_PolicyStore_ValidationSettings__Type is the CloudFormation type for AWS::VerifiedPermissions::PolicyStore.ValidationSettings.
	AWS_VerifiedPermissions_PolicyStore_ValidationSettings__Type = "AWS::VerifiedPermissions::PolicyStore.ValidationSettings"
)

var (
	// AWS_VerifiedPermissions_PolicyStore_ValidationSettings__PropertiesMap reports all the CloudFormation properties for AWS::VerifiedPermissions::PolicyStore.ValidationSettings.
	AWS_VerifiedPermissions_PolicyStore_ValidationSettings__PropertiesMap = struct {
		Mode string
	}{
		Mode: "Mode",
	}

	// AWS_VerifiedPermissions_PolicyStore_ValidationSettings__PropertiesSlice reports all the CloudFormation properties for AWS::VerifiedPermissions::PolicyStore.ValidationSettings.
	AWS_VerifiedPermissions_PolicyStore_ValidationSettings__PropertiesSlice = []string{
		AWS_VerifiedPermissions_PolicyStore_ValidationSettings__PropertiesMap.Mode,
	}
)

// AWS_VerifiedPermissions_PolicyStore_ValidationSettings is a binding for AWS::VerifiedPermissions::PolicyStore.ValidationSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-validationsettings.html
type AWS_VerifiedPermissions_PolicyStore_ValidationSettings struct {
	// Mode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-validationsettings.html#cfn-verifiedpermissions-policystore-validationsettings-mode
	Mode cfz.Expression[string] `json:"Mode,omitempty"`
}

// New__AWS_VerifiedPermissions_PolicyStore_ValidationSettings initializes a new AWS_VerifiedPermissions_PolicyStore_ValidationSettings.
func New__AWS_VerifiedPermissions_PolicyStore_ValidationSettings() AWS_VerifiedPermissions_PolicyStore_ValidationSettings {
	return AWS_VerifiedPermissions_PolicyStore_ValidationSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_VerifiedPermissions_PolicyStore_ValidationSettings) GetType() string {
	return AWS_VerifiedPermissions_PolicyStore_ValidationSettings__Type
}

// Set__Mode updates property "Mode".
func (t AWS_VerifiedPermissions_PolicyStore_ValidationSettings) Set__Mode(v cfz.Expression[string]) AWS_VerifiedPermissions_PolicyStore_ValidationSettings {
	t.Mode = v
	return t
}

// SetV__Mode updates property "Mode".
func (t AWS_VerifiedPermissions_PolicyStore_ValidationSettings) SetV__Mode(v string) AWS_VerifiedPermissions_PolicyStore_ValidationSettings {
	t.Mode = cfz.V(v)
	return t
}
