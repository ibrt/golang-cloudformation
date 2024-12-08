// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest__Type is the CloudFormation type for AWS::EC2::NetworkInsightsAccessScope.PathStatementRequest.
	AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest__Type = "AWS::EC2::NetworkInsightsAccessScope.PathStatementRequest"
)

var (
	// AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest__PropertiesMap reports all the CloudFormation properties for AWS::EC2::NetworkInsightsAccessScope.PathStatementRequest.
	AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest__PropertiesMap = struct {
		PacketHeaderStatement string
		ResourceStatement     string
	}{
		PacketHeaderStatement: "PacketHeaderStatement",
		ResourceStatement:     "ResourceStatement",
	}

	// AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::NetworkInsightsAccessScope.PathStatementRequest.
	AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest__PropertiesSlice = []string{
		AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest__PropertiesMap.PacketHeaderStatement,
		AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest__PropertiesMap.ResourceStatement,
	}
)

// AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest is a binding for AWS::EC2::NetworkInsightsAccessScope.PathStatementRequest.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-pathstatementrequest.html
type AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest struct {
	// PacketHeaderStatement is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-pathstatementrequest.html#cfn-ec2-networkinsightsaccessscope-pathstatementrequest-packetheaderstatement
	PacketHeaderStatement cfz.Expression[AWS_EC2_NetworkInsightsAccessScope_PacketHeaderStatementRequest] `json:"PacketHeaderStatement,omitempty"`

	// ResourceStatement is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-pathstatementrequest.html#cfn-ec2-networkinsightsaccessscope-pathstatementrequest-resourcestatement
	ResourceStatement cfz.Expression[AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest] `json:"ResourceStatement,omitempty"`
}

// New__AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest initializes a new AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest.
func New__AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest() AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest {
	return AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest) GetType() string {
	return AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest__Type
}

// Set__PacketHeaderStatement updates property "PacketHeaderStatement".
func (t AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest) Set__PacketHeaderStatement(v cfz.Expression[AWS_EC2_NetworkInsightsAccessScope_PacketHeaderStatementRequest]) AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest {
	t.PacketHeaderStatement = v
	return t
}

// SetV__PacketHeaderStatement updates property "PacketHeaderStatement".
func (t AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest) SetV__PacketHeaderStatement(v AWS_EC2_NetworkInsightsAccessScope_PacketHeaderStatementRequest) AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest {
	t.PacketHeaderStatement = cfz.V(v)
	return t
}

// Set__ResourceStatement updates property "ResourceStatement".
func (t AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest) Set__ResourceStatement(v cfz.Expression[AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest]) AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest {
	t.ResourceStatement = v
	return t
}

// SetV__ResourceStatement updates property "ResourceStatement".
func (t AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest) SetV__ResourceStatement(v AWS_EC2_NetworkInsightsAccessScope_ResourceStatementRequest) AWS_EC2_NetworkInsightsAccessScope_PathStatementRequest {
	t.ResourceStatement = cfz.V(v)
	return t
}
