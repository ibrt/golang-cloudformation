// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_glue

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Glue_Partition_SerdeInfo__Type is the CloudFormation type for AWS::Glue::Partition.SerdeInfo.
	AWS_Glue_Partition_SerdeInfo__Type = "AWS::Glue::Partition.SerdeInfo"
)

var (
	// AWS_Glue_Partition_SerdeInfo__PropertiesMap reports all the CloudFormation properties for AWS::Glue::Partition.SerdeInfo.
	AWS_Glue_Partition_SerdeInfo__PropertiesMap = struct {
		Name                 string
		Parameters           string
		SerializationLibrary string
	}{
		Name:                 "Name",
		Parameters:           "Parameters",
		SerializationLibrary: "SerializationLibrary",
	}

	// AWS_Glue_Partition_SerdeInfo__PropertiesSlice reports all the CloudFormation properties for AWS::Glue::Partition.SerdeInfo.
	AWS_Glue_Partition_SerdeInfo__PropertiesSlice = []string{
		AWS_Glue_Partition_SerdeInfo__PropertiesMap.Name,
		AWS_Glue_Partition_SerdeInfo__PropertiesMap.Parameters,
		AWS_Glue_Partition_SerdeInfo__PropertiesMap.SerializationLibrary,
	}
)

// AWS_Glue_Partition_SerdeInfo is a binding for AWS::Glue::Partition.SerdeInfo.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-serdeinfo.html
type AWS_Glue_Partition_SerdeInfo struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-serdeinfo.html#cfn-glue-partition-serdeinfo-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Parameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-serdeinfo.html#cfn-glue-partition-serdeinfo-parameters
	Parameters cfz.Expression[json.RawMessage] `json:"Parameters,omitempty"`

	// SerializationLibrary is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-serdeinfo.html#cfn-glue-partition-serdeinfo-serializationlibrary
	SerializationLibrary cfz.Expression[string] `json:"SerializationLibrary,omitempty"`
}

// New__AWS_Glue_Partition_SerdeInfo initializes a new AWS_Glue_Partition_SerdeInfo.
func New__AWS_Glue_Partition_SerdeInfo() AWS_Glue_Partition_SerdeInfo {
	return AWS_Glue_Partition_SerdeInfo{}
}

// GetType returns the CloudFormation type.
func (AWS_Glue_Partition_SerdeInfo) GetType() string {
	return AWS_Glue_Partition_SerdeInfo__Type
}

// Set__Name updates property "Name".
func (t AWS_Glue_Partition_SerdeInfo) Set__Name(v cfz.Expression[string]) AWS_Glue_Partition_SerdeInfo {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_Glue_Partition_SerdeInfo) SetV__Name(v string) AWS_Glue_Partition_SerdeInfo {
	t.Name = cfz.V(v)
	return t
}

// Set__Parameters updates property "Parameters".
func (t AWS_Glue_Partition_SerdeInfo) Set__Parameters(v cfz.Expression[json.RawMessage]) AWS_Glue_Partition_SerdeInfo {
	t.Parameters = v
	return t
}

// SetV__Parameters updates property "Parameters".
func (t AWS_Glue_Partition_SerdeInfo) SetV__Parameters(v json.RawMessage) AWS_Glue_Partition_SerdeInfo {
	t.Parameters = cfz.V(v)
	return t
}

// Set__SerializationLibrary updates property "SerializationLibrary".
func (t AWS_Glue_Partition_SerdeInfo) Set__SerializationLibrary(v cfz.Expression[string]) AWS_Glue_Partition_SerdeInfo {
	t.SerializationLibrary = v
	return t
}

// SetV__SerializationLibrary updates property "SerializationLibrary".
func (t AWS_Glue_Partition_SerdeInfo) SetV__SerializationLibrary(v string) AWS_Glue_Partition_SerdeInfo {
	t.SerializationLibrary = cfz.V(v)
	return t
}
