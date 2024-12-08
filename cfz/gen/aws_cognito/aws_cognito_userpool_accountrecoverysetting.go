// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cognito

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Cognito_UserPool_AccountRecoverySetting__Type is the CloudFormation type for AWS::Cognito::UserPool.AccountRecoverySetting.
	AWS_Cognito_UserPool_AccountRecoverySetting__Type = "AWS::Cognito::UserPool.AccountRecoverySetting"
)

var (
	// AWS_Cognito_UserPool_AccountRecoverySetting__PropertiesMap reports all the CloudFormation properties for AWS::Cognito::UserPool.AccountRecoverySetting.
	AWS_Cognito_UserPool_AccountRecoverySetting__PropertiesMap = struct {
		RecoveryMechanisms string
	}{
		RecoveryMechanisms: "RecoveryMechanisms",
	}

	// AWS_Cognito_UserPool_AccountRecoverySetting__PropertiesSlice reports all the CloudFormation properties for AWS::Cognito::UserPool.AccountRecoverySetting.
	AWS_Cognito_UserPool_AccountRecoverySetting__PropertiesSlice = []string{
		AWS_Cognito_UserPool_AccountRecoverySetting__PropertiesMap.RecoveryMechanisms,
	}
)

// AWS_Cognito_UserPool_AccountRecoverySetting is a binding for AWS::Cognito::UserPool.AccountRecoverySetting.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-accountrecoverysetting.html
type AWS_Cognito_UserPool_AccountRecoverySetting struct {
	// RecoveryMechanisms is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-accountrecoverysetting.html#cfn-cognito-userpool-accountrecoverysetting-recoverymechanisms
	RecoveryMechanisms cfz.ExpressionSlice[AWS_Cognito_UserPool_RecoveryOption] `json:"RecoveryMechanisms,omitempty"`
}

// New__AWS_Cognito_UserPool_AccountRecoverySetting initializes a new AWS_Cognito_UserPool_AccountRecoverySetting.
func New__AWS_Cognito_UserPool_AccountRecoverySetting() AWS_Cognito_UserPool_AccountRecoverySetting {
	return AWS_Cognito_UserPool_AccountRecoverySetting{}
}

// GetType returns the CloudFormation type.
func (AWS_Cognito_UserPool_AccountRecoverySetting) GetType() string {
	return AWS_Cognito_UserPool_AccountRecoverySetting__Type
}

// Set__RecoveryMechanisms updates property "RecoveryMechanisms".
func (t AWS_Cognito_UserPool_AccountRecoverySetting) Set__RecoveryMechanisms(v cfz.ExpressionSlice[AWS_Cognito_UserPool_RecoveryOption]) AWS_Cognito_UserPool_AccountRecoverySetting {
	t.RecoveryMechanisms = v
	return t
}

// SetS__RecoveryMechanisms updates property "RecoveryMechanisms".
func (t AWS_Cognito_UserPool_AccountRecoverySetting) SetS__RecoveryMechanisms(v ...cfz.Expression[AWS_Cognito_UserPool_RecoveryOption]) AWS_Cognito_UserPool_AccountRecoverySetting {
	t.RecoveryMechanisms = cfz.S(v...)
	return t
}

// SetSV__RecoveryMechanisms updates property "RecoveryMechanisms".
func (t AWS_Cognito_UserPool_AccountRecoverySetting) SetSV__RecoveryMechanisms(v ...AWS_Cognito_UserPool_RecoveryOption) AWS_Cognito_UserPool_AccountRecoverySetting {
	t.RecoveryMechanisms = cfz.SV(v...)
	return t
}
