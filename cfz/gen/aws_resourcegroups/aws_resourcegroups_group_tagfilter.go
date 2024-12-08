// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_resourcegroups

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ResourceGroups_Group_TagFilter__Type is the CloudFormation type for AWS::ResourceGroups::Group.TagFilter.
	AWS_ResourceGroups_Group_TagFilter__Type = "AWS::ResourceGroups::Group.TagFilter"
)

var (
	// AWS_ResourceGroups_Group_TagFilter__PropertiesMap reports all the CloudFormation properties for AWS::ResourceGroups::Group.TagFilter.
	AWS_ResourceGroups_Group_TagFilter__PropertiesMap = struct {
		Key    string
		Values string
	}{
		Key:    "Key",
		Values: "Values",
	}

	// AWS_ResourceGroups_Group_TagFilter__PropertiesSlice reports all the CloudFormation properties for AWS::ResourceGroups::Group.TagFilter.
	AWS_ResourceGroups_Group_TagFilter__PropertiesSlice = []string{
		AWS_ResourceGroups_Group_TagFilter__PropertiesMap.Key,
		AWS_ResourceGroups_Group_TagFilter__PropertiesMap.Values,
	}
)

// AWS_ResourceGroups_Group_TagFilter is a binding for AWS::ResourceGroups::Group.TagFilter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-tagfilter.html
type AWS_ResourceGroups_Group_TagFilter struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-tagfilter.html#cfn-resourcegroups-group-tagfilter-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Values is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-tagfilter.html#cfn-resourcegroups-group-tagfilter-values
	Values cfz.ExpressionSlice[string] `json:"Values,omitempty"`
}

// New__AWS_ResourceGroups_Group_TagFilter initializes a new AWS_ResourceGroups_Group_TagFilter.
func New__AWS_ResourceGroups_Group_TagFilter() AWS_ResourceGroups_Group_TagFilter {
	return AWS_ResourceGroups_Group_TagFilter{}
}

// GetType returns the CloudFormation type.
func (AWS_ResourceGroups_Group_TagFilter) GetType() string {
	return AWS_ResourceGroups_Group_TagFilter__Type
}

// Set__Key updates property "Key".
func (t AWS_ResourceGroups_Group_TagFilter) Set__Key(v cfz.Expression[string]) AWS_ResourceGroups_Group_TagFilter {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_ResourceGroups_Group_TagFilter) SetV__Key(v string) AWS_ResourceGroups_Group_TagFilter {
	t.Key = cfz.V(v)
	return t
}

// Set__Values updates property "Values".
func (t AWS_ResourceGroups_Group_TagFilter) Set__Values(v cfz.ExpressionSlice[string]) AWS_ResourceGroups_Group_TagFilter {
	t.Values = v
	return t
}

// SetS__Values updates property "Values".
func (t AWS_ResourceGroups_Group_TagFilter) SetS__Values(v ...cfz.Expression[string]) AWS_ResourceGroups_Group_TagFilter {
	t.Values = cfz.S(v...)
	return t
}

// SetSV__Values updates property "Values".
func (t AWS_ResourceGroups_Group_TagFilter) SetSV__Values(v ...string) AWS_ResourceGroups_Group_TagFilter {
	t.Values = cfz.SV(v...)
	return t
}
