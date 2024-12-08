// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_directoryservice

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DirectoryService_SimpleAD_VpcSettings__Type is the CloudFormation type for AWS::DirectoryService::SimpleAD.VpcSettings.
	AWS_DirectoryService_SimpleAD_VpcSettings__Type = "AWS::DirectoryService::SimpleAD.VpcSettings"
)

var (
	// AWS_DirectoryService_SimpleAD_VpcSettings__PropertiesMap reports all the CloudFormation properties for AWS::DirectoryService::SimpleAD.VpcSettings.
	AWS_DirectoryService_SimpleAD_VpcSettings__PropertiesMap = struct {
		SubnetIds string
		VpcId     string
	}{
		SubnetIds: "SubnetIds",
		VpcId:     "VpcId",
	}

	// AWS_DirectoryService_SimpleAD_VpcSettings__PropertiesSlice reports all the CloudFormation properties for AWS::DirectoryService::SimpleAD.VpcSettings.
	AWS_DirectoryService_SimpleAD_VpcSettings__PropertiesSlice = []string{
		AWS_DirectoryService_SimpleAD_VpcSettings__PropertiesMap.SubnetIds,
		AWS_DirectoryService_SimpleAD_VpcSettings__PropertiesMap.VpcId,
	}
)

// AWS_DirectoryService_SimpleAD_VpcSettings is a binding for AWS::DirectoryService::SimpleAD.VpcSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html
type AWS_DirectoryService_SimpleAD_VpcSettings struct {
	// SubnetIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-subnetids
	SubnetIds cfz.ExpressionSlice[string] `json:"SubnetIds,omitempty"`

	// VpcId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-vpcid
	VpcId cfz.Expression[string] `json:"VpcId,omitempty"`
}

// New__AWS_DirectoryService_SimpleAD_VpcSettings initializes a new AWS_DirectoryService_SimpleAD_VpcSettings.
func New__AWS_DirectoryService_SimpleAD_VpcSettings() AWS_DirectoryService_SimpleAD_VpcSettings {
	return AWS_DirectoryService_SimpleAD_VpcSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_DirectoryService_SimpleAD_VpcSettings) GetType() string {
	return AWS_DirectoryService_SimpleAD_VpcSettings__Type
}

// Set__SubnetIds updates property "SubnetIds".
func (t AWS_DirectoryService_SimpleAD_VpcSettings) Set__SubnetIds(v cfz.ExpressionSlice[string]) AWS_DirectoryService_SimpleAD_VpcSettings {
	t.SubnetIds = v
	return t
}

// SetS__SubnetIds updates property "SubnetIds".
func (t AWS_DirectoryService_SimpleAD_VpcSettings) SetS__SubnetIds(v ...cfz.Expression[string]) AWS_DirectoryService_SimpleAD_VpcSettings {
	t.SubnetIds = cfz.S(v...)
	return t
}

// SetSV__SubnetIds updates property "SubnetIds".
func (t AWS_DirectoryService_SimpleAD_VpcSettings) SetSV__SubnetIds(v ...string) AWS_DirectoryService_SimpleAD_VpcSettings {
	t.SubnetIds = cfz.SV(v...)
	return t
}

// Set__VpcId updates property "VpcId".
func (t AWS_DirectoryService_SimpleAD_VpcSettings) Set__VpcId(v cfz.Expression[string]) AWS_DirectoryService_SimpleAD_VpcSettings {
	t.VpcId = v
	return t
}

// SetV__VpcId updates property "VpcId".
func (t AWS_DirectoryService_SimpleAD_VpcSettings) SetV__VpcId(v string) AWS_DirectoryService_SimpleAD_VpcSettings {
	t.VpcId = cfz.V(v)
	return t
}
