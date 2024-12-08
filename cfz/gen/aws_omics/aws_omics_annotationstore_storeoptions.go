// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_omics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Omics_AnnotationStore_StoreOptions__Type is the CloudFormation type for AWS::Omics::AnnotationStore.StoreOptions.
	AWS_Omics_AnnotationStore_StoreOptions__Type = "AWS::Omics::AnnotationStore.StoreOptions"
)

var (
	// AWS_Omics_AnnotationStore_StoreOptions__PropertiesMap reports all the CloudFormation properties for AWS::Omics::AnnotationStore.StoreOptions.
	AWS_Omics_AnnotationStore_StoreOptions__PropertiesMap = struct {
		TsvStoreOptions string
	}{
		TsvStoreOptions: "TsvStoreOptions",
	}

	// AWS_Omics_AnnotationStore_StoreOptions__PropertiesSlice reports all the CloudFormation properties for AWS::Omics::AnnotationStore.StoreOptions.
	AWS_Omics_AnnotationStore_StoreOptions__PropertiesSlice = []string{
		AWS_Omics_AnnotationStore_StoreOptions__PropertiesMap.TsvStoreOptions,
	}
)

// AWS_Omics_AnnotationStore_StoreOptions is a binding for AWS::Omics::AnnotationStore.StoreOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-storeoptions.html
type AWS_Omics_AnnotationStore_StoreOptions struct {
	// TsvStoreOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-storeoptions.html#cfn-omics-annotationstore-storeoptions-tsvstoreoptions
	TsvStoreOptions cfz.Expression[AWS_Omics_AnnotationStore_TsvStoreOptions] `json:"TsvStoreOptions,omitempty"`
}

// New__AWS_Omics_AnnotationStore_StoreOptions initializes a new AWS_Omics_AnnotationStore_StoreOptions.
func New__AWS_Omics_AnnotationStore_StoreOptions() AWS_Omics_AnnotationStore_StoreOptions {
	return AWS_Omics_AnnotationStore_StoreOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_Omics_AnnotationStore_StoreOptions) GetType() string {
	return AWS_Omics_AnnotationStore_StoreOptions__Type
}

// Set__TsvStoreOptions updates property "TsvStoreOptions".
func (t AWS_Omics_AnnotationStore_StoreOptions) Set__TsvStoreOptions(v cfz.Expression[AWS_Omics_AnnotationStore_TsvStoreOptions]) AWS_Omics_AnnotationStore_StoreOptions {
	t.TsvStoreOptions = v
	return t
}

// SetV__TsvStoreOptions updates property "TsvStoreOptions".
func (t AWS_Omics_AnnotationStore_StoreOptions) SetV__TsvStoreOptions(v AWS_Omics_AnnotationStore_TsvStoreOptions) AWS_Omics_AnnotationStore_StoreOptions {
	t.TsvStoreOptions = cfz.V(v)
	return t
}
