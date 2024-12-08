// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediaconnect

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaConnect_Bridge_BridgeFlowSource__Type is the CloudFormation type for AWS::MediaConnect::Bridge.BridgeFlowSource.
	AWS_MediaConnect_Bridge_BridgeFlowSource__Type = "AWS::MediaConnect::Bridge.BridgeFlowSource"
)

var (
	// AWS_MediaConnect_Bridge_BridgeFlowSource__PropertiesMap reports all the CloudFormation properties for AWS::MediaConnect::Bridge.BridgeFlowSource.
	AWS_MediaConnect_Bridge_BridgeFlowSource__PropertiesMap = struct {
		FlowArn                    string
		FlowVpcInterfaceAttachment string
		Name                       string
	}{
		FlowArn:                    "FlowArn",
		FlowVpcInterfaceAttachment: "FlowVpcInterfaceAttachment",
		Name:                       "Name",
	}

	// AWS_MediaConnect_Bridge_BridgeFlowSource__PropertiesSlice reports all the CloudFormation properties for AWS::MediaConnect::Bridge.BridgeFlowSource.
	AWS_MediaConnect_Bridge_BridgeFlowSource__PropertiesSlice = []string{
		AWS_MediaConnect_Bridge_BridgeFlowSource__PropertiesMap.FlowArn,
		AWS_MediaConnect_Bridge_BridgeFlowSource__PropertiesMap.FlowVpcInterfaceAttachment,
		AWS_MediaConnect_Bridge_BridgeFlowSource__PropertiesMap.Name,
	}
)

// AWS_MediaConnect_Bridge_BridgeFlowSource is a binding for AWS::MediaConnect::Bridge.BridgeFlowSource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgeflowsource.html
type AWS_MediaConnect_Bridge_BridgeFlowSource struct {
	// FlowArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgeflowsource.html#cfn-mediaconnect-bridge-bridgeflowsource-flowarn
	FlowArn cfz.Expression[string] `json:"FlowArn,omitempty"`

	// FlowVpcInterfaceAttachment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgeflowsource.html#cfn-mediaconnect-bridge-bridgeflowsource-flowvpcinterfaceattachment
	FlowVpcInterfaceAttachment cfz.Expression[AWS_MediaConnect_Bridge_VpcInterfaceAttachment] `json:"FlowVpcInterfaceAttachment,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgeflowsource.html#cfn-mediaconnect-bridge-bridgeflowsource-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_MediaConnect_Bridge_BridgeFlowSource initializes a new AWS_MediaConnect_Bridge_BridgeFlowSource.
func New__AWS_MediaConnect_Bridge_BridgeFlowSource() AWS_MediaConnect_Bridge_BridgeFlowSource {
	return AWS_MediaConnect_Bridge_BridgeFlowSource{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaConnect_Bridge_BridgeFlowSource) GetType() string {
	return AWS_MediaConnect_Bridge_BridgeFlowSource__Type
}

// Set__FlowArn updates property "FlowArn".
func (t AWS_MediaConnect_Bridge_BridgeFlowSource) Set__FlowArn(v cfz.Expression[string]) AWS_MediaConnect_Bridge_BridgeFlowSource {
	t.FlowArn = v
	return t
}

// SetV__FlowArn updates property "FlowArn".
func (t AWS_MediaConnect_Bridge_BridgeFlowSource) SetV__FlowArn(v string) AWS_MediaConnect_Bridge_BridgeFlowSource {
	t.FlowArn = cfz.V(v)
	return t
}

// Set__FlowVpcInterfaceAttachment updates property "FlowVpcInterfaceAttachment".
func (t AWS_MediaConnect_Bridge_BridgeFlowSource) Set__FlowVpcInterfaceAttachment(v cfz.Expression[AWS_MediaConnect_Bridge_VpcInterfaceAttachment]) AWS_MediaConnect_Bridge_BridgeFlowSource {
	t.FlowVpcInterfaceAttachment = v
	return t
}

// SetV__FlowVpcInterfaceAttachment updates property "FlowVpcInterfaceAttachment".
func (t AWS_MediaConnect_Bridge_BridgeFlowSource) SetV__FlowVpcInterfaceAttachment(v AWS_MediaConnect_Bridge_VpcInterfaceAttachment) AWS_MediaConnect_Bridge_BridgeFlowSource {
	t.FlowVpcInterfaceAttachment = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_MediaConnect_Bridge_BridgeFlowSource) Set__Name(v cfz.Expression[string]) AWS_MediaConnect_Bridge_BridgeFlowSource {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_MediaConnect_Bridge_BridgeFlowSource) SetV__Name(v string) AWS_MediaConnect_Bridge_BridgeFlowSource {
	t.Name = cfz.V(v)
	return t
}
