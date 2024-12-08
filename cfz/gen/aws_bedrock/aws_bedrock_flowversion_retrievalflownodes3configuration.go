// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration__Type is the CloudFormation type for AWS::Bedrock::FlowVersion.RetrievalFlowNodeS3Configuration.
	AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration__Type = "AWS::Bedrock::FlowVersion.RetrievalFlowNodeS3Configuration"
)

var (
	// AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::FlowVersion.RetrievalFlowNodeS3Configuration.
	AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration__PropertiesMap = struct {
		BucketName string
	}{
		BucketName: "BucketName",
	}

	// AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::FlowVersion.RetrievalFlowNodeS3Configuration.
	AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration__PropertiesSlice = []string{
		AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration__PropertiesMap.BucketName,
	}
)

// AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration is a binding for AWS::Bedrock::FlowVersion.RetrievalFlowNodeS3Configuration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-retrievalflownodes3configuration.html
type AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration struct {
	// BucketName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-retrievalflownodes3configuration.html#cfn-bedrock-flowversion-retrievalflownodes3configuration-bucketname
	BucketName cfz.Expression[string] `json:"BucketName,omitempty"`
}

// New__AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration initializes a new AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration.
func New__AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration() AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration {
	return AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration) GetType() string {
	return AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration__Type
}

// Set__BucketName updates property "BucketName".
func (t AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration) Set__BucketName(v cfz.Expression[string]) AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration {
	t.BucketName = v
	return t
}

// SetV__BucketName updates property "BucketName".
func (t AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration) SetV__BucketName(v string) AWS_Bedrock_FlowVersion_RetrievalFlowNodeS3Configuration {
	t.BucketName = cfz.V(v)
	return t
}
