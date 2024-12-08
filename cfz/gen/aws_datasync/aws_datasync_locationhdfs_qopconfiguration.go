// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datasync

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataSync_LocationHDFS_QopConfiguration__Type is the CloudFormation type for AWS::DataSync::LocationHDFS.QopConfiguration.
	AWS_DataSync_LocationHDFS_QopConfiguration__Type = "AWS::DataSync::LocationHDFS.QopConfiguration"
)

var (
	// AWS_DataSync_LocationHDFS_QopConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::DataSync::LocationHDFS.QopConfiguration.
	AWS_DataSync_LocationHDFS_QopConfiguration__PropertiesMap = struct {
		DataTransferProtection string
		RpcProtection          string
	}{
		DataTransferProtection: "DataTransferProtection",
		RpcProtection:          "RpcProtection",
	}

	// AWS_DataSync_LocationHDFS_QopConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::DataSync::LocationHDFS.QopConfiguration.
	AWS_DataSync_LocationHDFS_QopConfiguration__PropertiesSlice = []string{
		AWS_DataSync_LocationHDFS_QopConfiguration__PropertiesMap.DataTransferProtection,
		AWS_DataSync_LocationHDFS_QopConfiguration__PropertiesMap.RpcProtection,
	}
)

// AWS_DataSync_LocationHDFS_QopConfiguration is a binding for AWS::DataSync::LocationHDFS.QopConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationhdfs-qopconfiguration.html
type AWS_DataSync_LocationHDFS_QopConfiguration struct {
	// DataTransferProtection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationhdfs-qopconfiguration.html#cfn-datasync-locationhdfs-qopconfiguration-datatransferprotection
	DataTransferProtection cfz.Expression[string] `json:"DataTransferProtection,omitempty"`

	// RpcProtection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationhdfs-qopconfiguration.html#cfn-datasync-locationhdfs-qopconfiguration-rpcprotection
	RpcProtection cfz.Expression[string] `json:"RpcProtection,omitempty"`
}

// New__AWS_DataSync_LocationHDFS_QopConfiguration initializes a new AWS_DataSync_LocationHDFS_QopConfiguration.
func New__AWS_DataSync_LocationHDFS_QopConfiguration() AWS_DataSync_LocationHDFS_QopConfiguration {
	return AWS_DataSync_LocationHDFS_QopConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_DataSync_LocationHDFS_QopConfiguration) GetType() string {
	return AWS_DataSync_LocationHDFS_QopConfiguration__Type
}

// Set__DataTransferProtection updates property "DataTransferProtection".
func (t AWS_DataSync_LocationHDFS_QopConfiguration) Set__DataTransferProtection(v cfz.Expression[string]) AWS_DataSync_LocationHDFS_QopConfiguration {
	t.DataTransferProtection = v
	return t
}

// SetV__DataTransferProtection updates property "DataTransferProtection".
func (t AWS_DataSync_LocationHDFS_QopConfiguration) SetV__DataTransferProtection(v string) AWS_DataSync_LocationHDFS_QopConfiguration {
	t.DataTransferProtection = cfz.V(v)
	return t
}

// Set__RpcProtection updates property "RpcProtection".
func (t AWS_DataSync_LocationHDFS_QopConfiguration) Set__RpcProtection(v cfz.Expression[string]) AWS_DataSync_LocationHDFS_QopConfiguration {
	t.RpcProtection = v
	return t
}

// SetV__RpcProtection updates property "RpcProtection".
func (t AWS_DataSync_LocationHDFS_QopConfiguration) SetV__RpcProtection(v string) AWS_DataSync_LocationHDFS_QopConfiguration {
	t.RpcProtection = cfz.V(v)
	return t
}
