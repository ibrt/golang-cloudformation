// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_Bucket_ReplicationRuleAndOperator__Type is the CloudFormation type for AWS::S3::Bucket.ReplicationRuleAndOperator.
	AWS_S3_Bucket_ReplicationRuleAndOperator__Type = "AWS::S3::Bucket.ReplicationRuleAndOperator"
)

var (
	// AWS_S3_Bucket_ReplicationRuleAndOperator__PropertiesMap reports all the CloudFormation properties for AWS::S3::Bucket.ReplicationRuleAndOperator.
	AWS_S3_Bucket_ReplicationRuleAndOperator__PropertiesMap = struct {
		Prefix     string
		TagFilters string
	}{
		Prefix:     "Prefix",
		TagFilters: "TagFilters",
	}

	// AWS_S3_Bucket_ReplicationRuleAndOperator__PropertiesSlice reports all the CloudFormation properties for AWS::S3::Bucket.ReplicationRuleAndOperator.
	AWS_S3_Bucket_ReplicationRuleAndOperator__PropertiesSlice = []string{
		AWS_S3_Bucket_ReplicationRuleAndOperator__PropertiesMap.Prefix,
		AWS_S3_Bucket_ReplicationRuleAndOperator__PropertiesMap.TagFilters,
	}
)

// AWS_S3_Bucket_ReplicationRuleAndOperator is a binding for AWS::S3::Bucket.ReplicationRuleAndOperator.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationruleandoperator.html
type AWS_S3_Bucket_ReplicationRuleAndOperator struct {
	// Prefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationruleandoperator.html#cfn-s3-bucket-replicationruleandoperator-prefix
	Prefix cfz.Expression[string] `json:"Prefix,omitempty"`

	// TagFilters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationruleandoperator.html#cfn-s3-bucket-replicationruleandoperator-tagfilters
	TagFilters cfz.ExpressionSlice[AWS_S3_Bucket_TagFilter] `json:"TagFilters,omitempty"`
}

// New__AWS_S3_Bucket_ReplicationRuleAndOperator initializes a new AWS_S3_Bucket_ReplicationRuleAndOperator.
func New__AWS_S3_Bucket_ReplicationRuleAndOperator() AWS_S3_Bucket_ReplicationRuleAndOperator {
	return AWS_S3_Bucket_ReplicationRuleAndOperator{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_Bucket_ReplicationRuleAndOperator) GetType() string {
	return AWS_S3_Bucket_ReplicationRuleAndOperator__Type
}

// Set__Prefix updates property "Prefix".
func (t AWS_S3_Bucket_ReplicationRuleAndOperator) Set__Prefix(v cfz.Expression[string]) AWS_S3_Bucket_ReplicationRuleAndOperator {
	t.Prefix = v
	return t
}

// SetV__Prefix updates property "Prefix".
func (t AWS_S3_Bucket_ReplicationRuleAndOperator) SetV__Prefix(v string) AWS_S3_Bucket_ReplicationRuleAndOperator {
	t.Prefix = cfz.V(v)
	return t
}

// Set__TagFilters updates property "TagFilters".
func (t AWS_S3_Bucket_ReplicationRuleAndOperator) Set__TagFilters(v cfz.ExpressionSlice[AWS_S3_Bucket_TagFilter]) AWS_S3_Bucket_ReplicationRuleAndOperator {
	t.TagFilters = v
	return t
}

// SetS__TagFilters updates property "TagFilters".
func (t AWS_S3_Bucket_ReplicationRuleAndOperator) SetS__TagFilters(v ...cfz.Expression[AWS_S3_Bucket_TagFilter]) AWS_S3_Bucket_ReplicationRuleAndOperator {
	t.TagFilters = cfz.S(v...)
	return t
}

// SetSV__TagFilters updates property "TagFilters".
func (t AWS_S3_Bucket_ReplicationRuleAndOperator) SetSV__TagFilters(v ...AWS_S3_Bucket_TagFilter) AWS_S3_Bucket_ReplicationRuleAndOperator {
	t.TagFilters = cfz.SV(v...)
	return t
}
