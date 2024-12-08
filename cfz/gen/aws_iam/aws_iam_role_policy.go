// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iam

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IAM_Role_Policy__Type is the CloudFormation type for AWS::IAM::Role.Policy.
	AWS_IAM_Role_Policy__Type = "AWS::IAM::Role.Policy"
)

var (
	// AWS_IAM_Role_Policy__PropertiesMap reports all the CloudFormation properties for AWS::IAM::Role.Policy.
	AWS_IAM_Role_Policy__PropertiesMap = struct {
		PolicyDocument string
		PolicyName     string
	}{
		PolicyDocument: "PolicyDocument",
		PolicyName:     "PolicyName",
	}

	// AWS_IAM_Role_Policy__PropertiesSlice reports all the CloudFormation properties for AWS::IAM::Role.Policy.
	AWS_IAM_Role_Policy__PropertiesSlice = []string{
		AWS_IAM_Role_Policy__PropertiesMap.PolicyDocument,
		AWS_IAM_Role_Policy__PropertiesMap.PolicyName,
	}
)

// AWS_IAM_Role_Policy is a binding for AWS::IAM::Role.Policy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html
type AWS_IAM_Role_Policy struct {
	// PolicyDocument is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html#cfn-iam-role-policy-policydocument
	PolicyDocument cfz.Expression[json.RawMessage] `json:"PolicyDocument,omitempty"`

	// PolicyName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html#cfn-iam-role-policy-policyname
	PolicyName cfz.Expression[string] `json:"PolicyName,omitempty"`
}

// New__AWS_IAM_Role_Policy initializes a new AWS_IAM_Role_Policy.
func New__AWS_IAM_Role_Policy() AWS_IAM_Role_Policy {
	return AWS_IAM_Role_Policy{}
}

// GetType returns the CloudFormation type.
func (AWS_IAM_Role_Policy) GetType() string {
	return AWS_IAM_Role_Policy__Type
}

// Set__PolicyDocument updates property "PolicyDocument".
func (t AWS_IAM_Role_Policy) Set__PolicyDocument(v cfz.Expression[json.RawMessage]) AWS_IAM_Role_Policy {
	t.PolicyDocument = v
	return t
}

// SetV__PolicyDocument updates property "PolicyDocument".
func (t AWS_IAM_Role_Policy) SetV__PolicyDocument(v json.RawMessage) AWS_IAM_Role_Policy {
	t.PolicyDocument = cfz.V(v)
	return t
}

// Set__PolicyName updates property "PolicyName".
func (t AWS_IAM_Role_Policy) Set__PolicyName(v cfz.Expression[string]) AWS_IAM_Role_Policy {
	t.PolicyName = v
	return t
}

// SetV__PolicyName updates property "PolicyName".
func (t AWS_IAM_Role_Policy) SetV__PolicyName(v string) AWS_IAM_Role_Policy {
	t.PolicyName = cfz.V(v)
	return t
}
