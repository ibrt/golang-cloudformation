// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_vpclattice

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_VpcLattice_Rule_FixedResponse__Type is the CloudFormation type for AWS::VpcLattice::Rule.FixedResponse.
	AWS_VpcLattice_Rule_FixedResponse__Type = "AWS::VpcLattice::Rule.FixedResponse"
)

var (
	// AWS_VpcLattice_Rule_FixedResponse__PropertiesMap reports all the CloudFormation properties for AWS::VpcLattice::Rule.FixedResponse.
	AWS_VpcLattice_Rule_FixedResponse__PropertiesMap = struct {
		StatusCode string
	}{
		StatusCode: "StatusCode",
	}

	// AWS_VpcLattice_Rule_FixedResponse__PropertiesSlice reports all the CloudFormation properties for AWS::VpcLattice::Rule.FixedResponse.
	AWS_VpcLattice_Rule_FixedResponse__PropertiesSlice = []string{
		AWS_VpcLattice_Rule_FixedResponse__PropertiesMap.StatusCode,
	}
)

// AWS_VpcLattice_Rule_FixedResponse is a binding for AWS::VpcLattice::Rule.FixedResponse.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-fixedresponse.html
type AWS_VpcLattice_Rule_FixedResponse struct {
	// StatusCode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-fixedresponse.html#cfn-vpclattice-rule-fixedresponse-statuscode
	StatusCode cfz.Expression[int32] `json:"StatusCode,omitempty"`
}

// New__AWS_VpcLattice_Rule_FixedResponse initializes a new AWS_VpcLattice_Rule_FixedResponse.
func New__AWS_VpcLattice_Rule_FixedResponse() AWS_VpcLattice_Rule_FixedResponse {
	return AWS_VpcLattice_Rule_FixedResponse{}
}

// GetType returns the CloudFormation type.
func (AWS_VpcLattice_Rule_FixedResponse) GetType() string {
	return AWS_VpcLattice_Rule_FixedResponse__Type
}

// Set__StatusCode updates property "StatusCode".
func (t AWS_VpcLattice_Rule_FixedResponse) Set__StatusCode(v cfz.Expression[int32]) AWS_VpcLattice_Rule_FixedResponse {
	t.StatusCode = v
	return t
}

// SetV__StatusCode updates property "StatusCode".
func (t AWS_VpcLattice_Rule_FixedResponse) SetV__StatusCode(v int32) AWS_VpcLattice_Rule_FixedResponse {
	t.StatusCode = cfz.V(v)
	return t
}
