// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datazone

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataZone_ProjectMembership_Member__Type is the CloudFormation type for AWS::DataZone::ProjectMembership.Member.
	AWS_DataZone_ProjectMembership_Member__Type = "AWS::DataZone::ProjectMembership.Member"
)

var (
	// AWS_DataZone_ProjectMembership_Member__PropertiesMap reports all the CloudFormation properties for AWS::DataZone::ProjectMembership.Member.
	AWS_DataZone_ProjectMembership_Member__PropertiesMap = struct {
		GroupIdentifier string
		UserIdentifier  string
	}{
		GroupIdentifier: "GroupIdentifier",
		UserIdentifier:  "UserIdentifier",
	}

	// AWS_DataZone_ProjectMembership_Member__PropertiesSlice reports all the CloudFormation properties for AWS::DataZone::ProjectMembership.Member.
	AWS_DataZone_ProjectMembership_Member__PropertiesSlice = []string{
		AWS_DataZone_ProjectMembership_Member__PropertiesMap.GroupIdentifier,
		AWS_DataZone_ProjectMembership_Member__PropertiesMap.UserIdentifier,
	}
)

// AWS_DataZone_ProjectMembership_Member is a binding for AWS::DataZone::ProjectMembership.Member.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectmembership-member.html
type AWS_DataZone_ProjectMembership_Member struct {
	// GroupIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectmembership-member.html#cfn-datazone-projectmembership-member-groupidentifier
	GroupIdentifier cfz.Expression[string] `json:"GroupIdentifier,omitempty"`

	// UserIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectmembership-member.html#cfn-datazone-projectmembership-member-useridentifier
	UserIdentifier cfz.Expression[string] `json:"UserIdentifier,omitempty"`
}

// New__AWS_DataZone_ProjectMembership_Member initializes a new AWS_DataZone_ProjectMembership_Member.
func New__AWS_DataZone_ProjectMembership_Member() AWS_DataZone_ProjectMembership_Member {
	return AWS_DataZone_ProjectMembership_Member{}
}

// GetType returns the CloudFormation type.
func (AWS_DataZone_ProjectMembership_Member) GetType() string {
	return AWS_DataZone_ProjectMembership_Member__Type
}

// Set__GroupIdentifier updates property "GroupIdentifier".
func (t AWS_DataZone_ProjectMembership_Member) Set__GroupIdentifier(v cfz.Expression[string]) AWS_DataZone_ProjectMembership_Member {
	t.GroupIdentifier = v
	return t
}

// SetV__GroupIdentifier updates property "GroupIdentifier".
func (t AWS_DataZone_ProjectMembership_Member) SetV__GroupIdentifier(v string) AWS_DataZone_ProjectMembership_Member {
	t.GroupIdentifier = cfz.V(v)
	return t
}

// Set__UserIdentifier updates property "UserIdentifier".
func (t AWS_DataZone_ProjectMembership_Member) Set__UserIdentifier(v cfz.Expression[string]) AWS_DataZone_ProjectMembership_Member {
	t.UserIdentifier = v
	return t
}

// SetV__UserIdentifier updates property "UserIdentifier".
func (t AWS_DataZone_ProjectMembership_Member) SetV__UserIdentifier(v string) AWS_DataZone_ProjectMembership_Member {
	t.UserIdentifier = cfz.V(v)
	return t
}
