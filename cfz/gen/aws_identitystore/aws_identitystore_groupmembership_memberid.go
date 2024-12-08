// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_identitystore

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IdentityStore_GroupMembership_MemberId__Type is the CloudFormation type for AWS::IdentityStore::GroupMembership.MemberId.
	AWS_IdentityStore_GroupMembership_MemberId__Type = "AWS::IdentityStore::GroupMembership.MemberId"
)

var (
	// AWS_IdentityStore_GroupMembership_MemberId__PropertiesMap reports all the CloudFormation properties for AWS::IdentityStore::GroupMembership.MemberId.
	AWS_IdentityStore_GroupMembership_MemberId__PropertiesMap = struct {
		UserId string
	}{
		UserId: "UserId",
	}

	// AWS_IdentityStore_GroupMembership_MemberId__PropertiesSlice reports all the CloudFormation properties for AWS::IdentityStore::GroupMembership.MemberId.
	AWS_IdentityStore_GroupMembership_MemberId__PropertiesSlice = []string{
		AWS_IdentityStore_GroupMembership_MemberId__PropertiesMap.UserId,
	}
)

// AWS_IdentityStore_GroupMembership_MemberId is a binding for AWS::IdentityStore::GroupMembership.MemberId.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-groupmembership-memberid.html
type AWS_IdentityStore_GroupMembership_MemberId struct {
	// UserId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-groupmembership-memberid.html#cfn-identitystore-groupmembership-memberid-userid
	UserId cfz.Expression[string] `json:"UserId,omitempty"`
}

// New__AWS_IdentityStore_GroupMembership_MemberId initializes a new AWS_IdentityStore_GroupMembership_MemberId.
func New__AWS_IdentityStore_GroupMembership_MemberId() AWS_IdentityStore_GroupMembership_MemberId {
	return AWS_IdentityStore_GroupMembership_MemberId{}
}

// GetType returns the CloudFormation type.
func (AWS_IdentityStore_GroupMembership_MemberId) GetType() string {
	return AWS_IdentityStore_GroupMembership_MemberId__Type
}

// Set__UserId updates property "UserId".
func (t AWS_IdentityStore_GroupMembership_MemberId) Set__UserId(v cfz.Expression[string]) AWS_IdentityStore_GroupMembership_MemberId {
	t.UserId = v
	return t
}

// SetV__UserId updates property "UserId".
func (t AWS_IdentityStore_GroupMembership_MemberId) SetV__UserId(v string) AWS_IdentityStore_GroupMembership_MemberId {
	t.UserId = cfz.V(v)
	return t
}
