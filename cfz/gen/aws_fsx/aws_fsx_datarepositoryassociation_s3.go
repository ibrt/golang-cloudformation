// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_fsx

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FSx_DataRepositoryAssociation_S3__Type is the CloudFormation type for AWS::FSx::DataRepositoryAssociation.S3.
	AWS_FSx_DataRepositoryAssociation_S3__Type = "AWS::FSx::DataRepositoryAssociation.S3"
)

var (
	// AWS_FSx_DataRepositoryAssociation_S3__PropertiesMap reports all the CloudFormation properties for AWS::FSx::DataRepositoryAssociation.S3.
	AWS_FSx_DataRepositoryAssociation_S3__PropertiesMap = struct {
		AutoExportPolicy string
		AutoImportPolicy string
	}{
		AutoExportPolicy: "AutoExportPolicy",
		AutoImportPolicy: "AutoImportPolicy",
	}

	// AWS_FSx_DataRepositoryAssociation_S3__PropertiesSlice reports all the CloudFormation properties for AWS::FSx::DataRepositoryAssociation.S3.
	AWS_FSx_DataRepositoryAssociation_S3__PropertiesSlice = []string{
		AWS_FSx_DataRepositoryAssociation_S3__PropertiesMap.AutoExportPolicy,
		AWS_FSx_DataRepositoryAssociation_S3__PropertiesMap.AutoImportPolicy,
	}
)

// AWS_FSx_DataRepositoryAssociation_S3 is a binding for AWS::FSx::DataRepositoryAssociation.S3.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-datarepositoryassociation-s3.html
type AWS_FSx_DataRepositoryAssociation_S3 struct {
	// AutoExportPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-datarepositoryassociation-s3.html#cfn-fsx-datarepositoryassociation-s3-autoexportpolicy
	AutoExportPolicy cfz.Expression[AWS_FSx_DataRepositoryAssociation_AutoExportPolicy] `json:"AutoExportPolicy,omitempty"`

	// AutoImportPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-datarepositoryassociation-s3.html#cfn-fsx-datarepositoryassociation-s3-autoimportpolicy
	AutoImportPolicy cfz.Expression[AWS_FSx_DataRepositoryAssociation_AutoImportPolicy] `json:"AutoImportPolicy,omitempty"`
}

// New__AWS_FSx_DataRepositoryAssociation_S3 initializes a new AWS_FSx_DataRepositoryAssociation_S3.
func New__AWS_FSx_DataRepositoryAssociation_S3() AWS_FSx_DataRepositoryAssociation_S3 {
	return AWS_FSx_DataRepositoryAssociation_S3{}
}

// GetType returns the CloudFormation type.
func (AWS_FSx_DataRepositoryAssociation_S3) GetType() string {
	return AWS_FSx_DataRepositoryAssociation_S3__Type
}

// Set__AutoExportPolicy updates property "AutoExportPolicy".
func (t AWS_FSx_DataRepositoryAssociation_S3) Set__AutoExportPolicy(v cfz.Expression[AWS_FSx_DataRepositoryAssociation_AutoExportPolicy]) AWS_FSx_DataRepositoryAssociation_S3 {
	t.AutoExportPolicy = v
	return t
}

// SetV__AutoExportPolicy updates property "AutoExportPolicy".
func (t AWS_FSx_DataRepositoryAssociation_S3) SetV__AutoExportPolicy(v AWS_FSx_DataRepositoryAssociation_AutoExportPolicy) AWS_FSx_DataRepositoryAssociation_S3 {
	t.AutoExportPolicy = cfz.V(v)
	return t
}

// Set__AutoImportPolicy updates property "AutoImportPolicy".
func (t AWS_FSx_DataRepositoryAssociation_S3) Set__AutoImportPolicy(v cfz.Expression[AWS_FSx_DataRepositoryAssociation_AutoImportPolicy]) AWS_FSx_DataRepositoryAssociation_S3 {
	t.AutoImportPolicy = v
	return t
}

// SetV__AutoImportPolicy updates property "AutoImportPolicy".
func (t AWS_FSx_DataRepositoryAssociation_S3) SetV__AutoImportPolicy(v AWS_FSx_DataRepositoryAssociation_AutoImportPolicy) AWS_FSx_DataRepositoryAssociation_S3 {
	t.AutoImportPolicy = cfz.V(v)
	return t
}
