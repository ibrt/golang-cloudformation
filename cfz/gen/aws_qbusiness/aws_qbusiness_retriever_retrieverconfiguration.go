// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_qbusiness

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QBusiness_Retriever_RetrieverConfiguration__Type is the CloudFormation type for AWS::QBusiness::Retriever.RetrieverConfiguration.
	AWS_QBusiness_Retriever_RetrieverConfiguration__Type = "AWS::QBusiness::Retriever.RetrieverConfiguration"
)

var (
	// AWS_QBusiness_Retriever_RetrieverConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::QBusiness::Retriever.RetrieverConfiguration.
	AWS_QBusiness_Retriever_RetrieverConfiguration__PropertiesMap = struct {
		KendraIndexConfiguration string
		NativeIndexConfiguration string
	}{
		KendraIndexConfiguration: "KendraIndexConfiguration",
		NativeIndexConfiguration: "NativeIndexConfiguration",
	}

	// AWS_QBusiness_Retriever_RetrieverConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::QBusiness::Retriever.RetrieverConfiguration.
	AWS_QBusiness_Retriever_RetrieverConfiguration__PropertiesSlice = []string{
		AWS_QBusiness_Retriever_RetrieverConfiguration__PropertiesMap.KendraIndexConfiguration,
		AWS_QBusiness_Retriever_RetrieverConfiguration__PropertiesMap.NativeIndexConfiguration,
	}
)

// AWS_QBusiness_Retriever_RetrieverConfiguration is a binding for AWS::QBusiness::Retriever.RetrieverConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html
type AWS_QBusiness_Retriever_RetrieverConfiguration struct {
	// KendraIndexConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html#cfn-qbusiness-retriever-retrieverconfiguration-kendraindexconfiguration
	KendraIndexConfiguration cfz.Expression[AWS_QBusiness_Retriever_KendraIndexConfiguration] `json:"KendraIndexConfiguration,omitempty"`

	// NativeIndexConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html#cfn-qbusiness-retriever-retrieverconfiguration-nativeindexconfiguration
	NativeIndexConfiguration cfz.Expression[AWS_QBusiness_Retriever_NativeIndexConfiguration] `json:"NativeIndexConfiguration,omitempty"`
}

// New__AWS_QBusiness_Retriever_RetrieverConfiguration initializes a new AWS_QBusiness_Retriever_RetrieverConfiguration.
func New__AWS_QBusiness_Retriever_RetrieverConfiguration() AWS_QBusiness_Retriever_RetrieverConfiguration {
	return AWS_QBusiness_Retriever_RetrieverConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_QBusiness_Retriever_RetrieverConfiguration) GetType() string {
	return AWS_QBusiness_Retriever_RetrieverConfiguration__Type
}

// Set__KendraIndexConfiguration updates property "KendraIndexConfiguration".
func (t AWS_QBusiness_Retriever_RetrieverConfiguration) Set__KendraIndexConfiguration(v cfz.Expression[AWS_QBusiness_Retriever_KendraIndexConfiguration]) AWS_QBusiness_Retriever_RetrieverConfiguration {
	t.KendraIndexConfiguration = v
	return t
}

// SetV__KendraIndexConfiguration updates property "KendraIndexConfiguration".
func (t AWS_QBusiness_Retriever_RetrieverConfiguration) SetV__KendraIndexConfiguration(v AWS_QBusiness_Retriever_KendraIndexConfiguration) AWS_QBusiness_Retriever_RetrieverConfiguration {
	t.KendraIndexConfiguration = cfz.V(v)
	return t
}

// Set__NativeIndexConfiguration updates property "NativeIndexConfiguration".
func (t AWS_QBusiness_Retriever_RetrieverConfiguration) Set__NativeIndexConfiguration(v cfz.Expression[AWS_QBusiness_Retriever_NativeIndexConfiguration]) AWS_QBusiness_Retriever_RetrieverConfiguration {
	t.NativeIndexConfiguration = v
	return t
}

// SetV__NativeIndexConfiguration updates property "NativeIndexConfiguration".
func (t AWS_QBusiness_Retriever_RetrieverConfiguration) SetV__NativeIndexConfiguration(v AWS_QBusiness_Retriever_NativeIndexConfiguration) AWS_QBusiness_Retriever_RetrieverConfiguration {
	t.NativeIndexConfiguration = cfz.V(v)
	return t
}
