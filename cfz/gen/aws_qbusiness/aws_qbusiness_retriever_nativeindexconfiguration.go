// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_qbusiness

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QBusiness_Retriever_NativeIndexConfiguration__Type is the CloudFormation type for AWS::QBusiness::Retriever.NativeIndexConfiguration.
	AWS_QBusiness_Retriever_NativeIndexConfiguration__Type = "AWS::QBusiness::Retriever.NativeIndexConfiguration"
)

var (
	// AWS_QBusiness_Retriever_NativeIndexConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::QBusiness::Retriever.NativeIndexConfiguration.
	AWS_QBusiness_Retriever_NativeIndexConfiguration__PropertiesMap = struct {
		IndexId string
	}{
		IndexId: "IndexId",
	}

	// AWS_QBusiness_Retriever_NativeIndexConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::QBusiness::Retriever.NativeIndexConfiguration.
	AWS_QBusiness_Retriever_NativeIndexConfiguration__PropertiesSlice = []string{
		AWS_QBusiness_Retriever_NativeIndexConfiguration__PropertiesMap.IndexId,
	}
)

// AWS_QBusiness_Retriever_NativeIndexConfiguration is a binding for AWS::QBusiness::Retriever.NativeIndexConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-nativeindexconfiguration.html
type AWS_QBusiness_Retriever_NativeIndexConfiguration struct {
	// IndexId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-nativeindexconfiguration.html#cfn-qbusiness-retriever-nativeindexconfiguration-indexid
	IndexId cfz.Expression[string] `json:"IndexId,omitempty"`
}

// New__AWS_QBusiness_Retriever_NativeIndexConfiguration initializes a new AWS_QBusiness_Retriever_NativeIndexConfiguration.
func New__AWS_QBusiness_Retriever_NativeIndexConfiguration() AWS_QBusiness_Retriever_NativeIndexConfiguration {
	return AWS_QBusiness_Retriever_NativeIndexConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_QBusiness_Retriever_NativeIndexConfiguration) GetType() string {
	return AWS_QBusiness_Retriever_NativeIndexConfiguration__Type
}

// Set__IndexId updates property "IndexId".
func (t AWS_QBusiness_Retriever_NativeIndexConfiguration) Set__IndexId(v cfz.Expression[string]) AWS_QBusiness_Retriever_NativeIndexConfiguration {
	t.IndexId = v
	return t
}

// SetV__IndexId updates property "IndexId".
func (t AWS_QBusiness_Retriever_NativeIndexConfiguration) SetV__IndexId(v string) AWS_QBusiness_Retriever_NativeIndexConfiguration {
	t.IndexId = cfz.V(v)
	return t
}
