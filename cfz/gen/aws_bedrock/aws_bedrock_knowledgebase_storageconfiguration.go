// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_KnowledgeBase_StorageConfiguration__Type is the CloudFormation type for AWS::Bedrock::KnowledgeBase.StorageConfiguration.
	AWS_Bedrock_KnowledgeBase_StorageConfiguration__Type = "AWS::Bedrock::KnowledgeBase.StorageConfiguration"
)

var (
	// AWS_Bedrock_KnowledgeBase_StorageConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::KnowledgeBase.StorageConfiguration.
	AWS_Bedrock_KnowledgeBase_StorageConfiguration__PropertiesMap = struct {
		MongoDbAtlasConfiguration         string
		OpensearchServerlessConfiguration string
		PineconeConfiguration             string
		RdsConfiguration                  string
		Type                              string
	}{
		MongoDbAtlasConfiguration:         "MongoDbAtlasConfiguration",
		OpensearchServerlessConfiguration: "OpensearchServerlessConfiguration",
		PineconeConfiguration:             "PineconeConfiguration",
		RdsConfiguration:                  "RdsConfiguration",
		Type:                              "Type",
	}

	// AWS_Bedrock_KnowledgeBase_StorageConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::KnowledgeBase.StorageConfiguration.
	AWS_Bedrock_KnowledgeBase_StorageConfiguration__PropertiesSlice = []string{
		AWS_Bedrock_KnowledgeBase_StorageConfiguration__PropertiesMap.MongoDbAtlasConfiguration,
		AWS_Bedrock_KnowledgeBase_StorageConfiguration__PropertiesMap.OpensearchServerlessConfiguration,
		AWS_Bedrock_KnowledgeBase_StorageConfiguration__PropertiesMap.PineconeConfiguration,
		AWS_Bedrock_KnowledgeBase_StorageConfiguration__PropertiesMap.RdsConfiguration,
		AWS_Bedrock_KnowledgeBase_StorageConfiguration__PropertiesMap.Type,
	}
)

// AWS_Bedrock_KnowledgeBase_StorageConfiguration is a binding for AWS::Bedrock::KnowledgeBase.StorageConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html
type AWS_Bedrock_KnowledgeBase_StorageConfiguration struct {
	// MongoDbAtlasConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-mongodbatlasconfiguration
	MongoDbAtlasConfiguration cfz.Expression[AWS_Bedrock_KnowledgeBase_MongoDbAtlasConfiguration] `json:"MongoDbAtlasConfiguration,omitempty"`

	// OpensearchServerlessConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-opensearchserverlessconfiguration
	OpensearchServerlessConfiguration cfz.Expression[AWS_Bedrock_KnowledgeBase_OpenSearchServerlessConfiguration] `json:"OpensearchServerlessConfiguration,omitempty"`

	// PineconeConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-pineconeconfiguration
	PineconeConfiguration cfz.Expression[AWS_Bedrock_KnowledgeBase_PineconeConfiguration] `json:"PineconeConfiguration,omitempty"`

	// RdsConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-rdsconfiguration
	RdsConfiguration cfz.Expression[AWS_Bedrock_KnowledgeBase_RdsConfiguration] `json:"RdsConfiguration,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-storageconfiguration.html#cfn-bedrock-knowledgebase-storageconfiguration-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_Bedrock_KnowledgeBase_StorageConfiguration initializes a new AWS_Bedrock_KnowledgeBase_StorageConfiguration.
func New__AWS_Bedrock_KnowledgeBase_StorageConfiguration() AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	return AWS_Bedrock_KnowledgeBase_StorageConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_KnowledgeBase_StorageConfiguration) GetType() string {
	return AWS_Bedrock_KnowledgeBase_StorageConfiguration__Type
}

// Set__MongoDbAtlasConfiguration updates property "MongoDbAtlasConfiguration".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) Set__MongoDbAtlasConfiguration(v cfz.Expression[AWS_Bedrock_KnowledgeBase_MongoDbAtlasConfiguration]) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.MongoDbAtlasConfiguration = v
	return t
}

// SetV__MongoDbAtlasConfiguration updates property "MongoDbAtlasConfiguration".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) SetV__MongoDbAtlasConfiguration(v AWS_Bedrock_KnowledgeBase_MongoDbAtlasConfiguration) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.MongoDbAtlasConfiguration = cfz.V(v)
	return t
}

// Set__OpensearchServerlessConfiguration updates property "OpensearchServerlessConfiguration".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) Set__OpensearchServerlessConfiguration(v cfz.Expression[AWS_Bedrock_KnowledgeBase_OpenSearchServerlessConfiguration]) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.OpensearchServerlessConfiguration = v
	return t
}

// SetV__OpensearchServerlessConfiguration updates property "OpensearchServerlessConfiguration".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) SetV__OpensearchServerlessConfiguration(v AWS_Bedrock_KnowledgeBase_OpenSearchServerlessConfiguration) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.OpensearchServerlessConfiguration = cfz.V(v)
	return t
}

// Set__PineconeConfiguration updates property "PineconeConfiguration".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) Set__PineconeConfiguration(v cfz.Expression[AWS_Bedrock_KnowledgeBase_PineconeConfiguration]) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.PineconeConfiguration = v
	return t
}

// SetV__PineconeConfiguration updates property "PineconeConfiguration".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) SetV__PineconeConfiguration(v AWS_Bedrock_KnowledgeBase_PineconeConfiguration) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.PineconeConfiguration = cfz.V(v)
	return t
}

// Set__RdsConfiguration updates property "RdsConfiguration".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) Set__RdsConfiguration(v cfz.Expression[AWS_Bedrock_KnowledgeBase_RdsConfiguration]) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.RdsConfiguration = v
	return t
}

// SetV__RdsConfiguration updates property "RdsConfiguration".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) SetV__RdsConfiguration(v AWS_Bedrock_KnowledgeBase_RdsConfiguration) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.RdsConfiguration = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) Set__Type(v cfz.Expression[string]) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_Bedrock_KnowledgeBase_StorageConfiguration) SetV__Type(v string) AWS_Bedrock_KnowledgeBase_StorageConfiguration {
	t.Type = cfz.V(v)
	return t
}
