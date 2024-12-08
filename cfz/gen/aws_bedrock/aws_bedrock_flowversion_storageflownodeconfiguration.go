// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration__Type is the CloudFormation type for AWS::Bedrock::FlowVersion.StorageFlowNodeConfiguration.
	AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration__Type = "AWS::Bedrock::FlowVersion.StorageFlowNodeConfiguration"
)

var (
	// AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::FlowVersion.StorageFlowNodeConfiguration.
	AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration__PropertiesMap = struct {
		ServiceConfiguration string
	}{
		ServiceConfiguration: "ServiceConfiguration",
	}

	// AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::FlowVersion.StorageFlowNodeConfiguration.
	AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration__PropertiesSlice = []string{
		AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration__PropertiesMap.ServiceConfiguration,
	}
)

// AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration is a binding for AWS::Bedrock::FlowVersion.StorageFlowNodeConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-storageflownodeconfiguration.html
type AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration struct {
	// ServiceConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-storageflownodeconfiguration.html#cfn-bedrock-flowversion-storageflownodeconfiguration-serviceconfiguration
	ServiceConfiguration cfz.Expression[AWS_Bedrock_FlowVersion_StorageFlowNodeServiceConfiguration] `json:"ServiceConfiguration,omitempty"`
}

// New__AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration initializes a new AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration.
func New__AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration() AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration {
	return AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration) GetType() string {
	return AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration__Type
}

// Set__ServiceConfiguration updates property "ServiceConfiguration".
func (t AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration) Set__ServiceConfiguration(v cfz.Expression[AWS_Bedrock_FlowVersion_StorageFlowNodeServiceConfiguration]) AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration {
	t.ServiceConfiguration = v
	return t
}

// SetV__ServiceConfiguration updates property "ServiceConfiguration".
func (t AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration) SetV__ServiceConfiguration(v AWS_Bedrock_FlowVersion_StorageFlowNodeServiceConfiguration) AWS_Bedrock_FlowVersion_StorageFlowNodeConfiguration {
	t.ServiceConfiguration = cfz.V(v)
	return t
}
