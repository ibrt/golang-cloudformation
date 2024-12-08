// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudfront

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFront_Distribution_OriginGroupMember__Type is the CloudFormation type for AWS::CloudFront::Distribution.OriginGroupMember.
	AWS_CloudFront_Distribution_OriginGroupMember__Type = "AWS::CloudFront::Distribution.OriginGroupMember"
)

var (
	// AWS_CloudFront_Distribution_OriginGroupMember__PropertiesMap reports all the CloudFormation properties for AWS::CloudFront::Distribution.OriginGroupMember.
	AWS_CloudFront_Distribution_OriginGroupMember__PropertiesMap = struct {
		OriginId string
	}{
		OriginId: "OriginId",
	}

	// AWS_CloudFront_Distribution_OriginGroupMember__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFront::Distribution.OriginGroupMember.
	AWS_CloudFront_Distribution_OriginGroupMember__PropertiesSlice = []string{
		AWS_CloudFront_Distribution_OriginGroupMember__PropertiesMap.OriginId,
	}
)

// AWS_CloudFront_Distribution_OriginGroupMember is a binding for AWS::CloudFront::Distribution.OriginGroupMember.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroupmember.html
type AWS_CloudFront_Distribution_OriginGroupMember struct {
	// OriginId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroupmember.html#cfn-cloudfront-distribution-origingroupmember-originid
	OriginId cfz.Expression[string] `json:"OriginId,omitempty"`
}

// New__AWS_CloudFront_Distribution_OriginGroupMember initializes a new AWS_CloudFront_Distribution_OriginGroupMember.
func New__AWS_CloudFront_Distribution_OriginGroupMember() AWS_CloudFront_Distribution_OriginGroupMember {
	return AWS_CloudFront_Distribution_OriginGroupMember{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFront_Distribution_OriginGroupMember) GetType() string {
	return AWS_CloudFront_Distribution_OriginGroupMember__Type
}

// Set__OriginId updates property "OriginId".
func (t AWS_CloudFront_Distribution_OriginGroupMember) Set__OriginId(v cfz.Expression[string]) AWS_CloudFront_Distribution_OriginGroupMember {
	t.OriginId = v
	return t
}

// SetV__OriginId updates property "OriginId".
func (t AWS_CloudFront_Distribution_OriginGroupMember) SetV__OriginId(v string) AWS_CloudFront_Distribution_OriginGroupMember {
	t.OriginId = cfz.V(v)
	return t
}
