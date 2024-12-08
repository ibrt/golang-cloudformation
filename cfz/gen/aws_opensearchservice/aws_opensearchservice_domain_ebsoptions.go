// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_opensearchservice

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_OpenSearchService_Domain_EBSOptions__Type is the CloudFormation type for AWS::OpenSearchService::Domain.EBSOptions.
	AWS_OpenSearchService_Domain_EBSOptions__Type = "AWS::OpenSearchService::Domain.EBSOptions"
)

var (
	// AWS_OpenSearchService_Domain_EBSOptions__PropertiesMap reports all the CloudFormation properties for AWS::OpenSearchService::Domain.EBSOptions.
	AWS_OpenSearchService_Domain_EBSOptions__PropertiesMap = struct {
		EBSEnabled string
		Iops       string
		Throughput string
		VolumeSize string
		VolumeType string
	}{
		EBSEnabled: "EBSEnabled",
		Iops:       "Iops",
		Throughput: "Throughput",
		VolumeSize: "VolumeSize",
		VolumeType: "VolumeType",
	}

	// AWS_OpenSearchService_Domain_EBSOptions__PropertiesSlice reports all the CloudFormation properties for AWS::OpenSearchService::Domain.EBSOptions.
	AWS_OpenSearchService_Domain_EBSOptions__PropertiesSlice = []string{
		AWS_OpenSearchService_Domain_EBSOptions__PropertiesMap.EBSEnabled,
		AWS_OpenSearchService_Domain_EBSOptions__PropertiesMap.Iops,
		AWS_OpenSearchService_Domain_EBSOptions__PropertiesMap.Throughput,
		AWS_OpenSearchService_Domain_EBSOptions__PropertiesMap.VolumeSize,
		AWS_OpenSearchService_Domain_EBSOptions__PropertiesMap.VolumeType,
	}
)

// AWS_OpenSearchService_Domain_EBSOptions is a binding for AWS::OpenSearchService::Domain.EBSOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html
type AWS_OpenSearchService_Domain_EBSOptions struct {
	// EBSEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-ebsenabled
	EBSEnabled cfz.Expression[bool] `json:"EBSEnabled,omitempty"`

	// Iops is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-iops
	Iops cfz.Expression[int32] `json:"Iops,omitempty"`

	// Throughput is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-throughput
	Throughput cfz.Expression[int32] `json:"Throughput,omitempty"`

	// VolumeSize is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-volumesize
	VolumeSize cfz.Expression[int32] `json:"VolumeSize,omitempty"`

	// VolumeType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-ebsoptions.html#cfn-opensearchservice-domain-ebsoptions-volumetype
	VolumeType cfz.Expression[string] `json:"VolumeType,omitempty"`
}

// New__AWS_OpenSearchService_Domain_EBSOptions initializes a new AWS_OpenSearchService_Domain_EBSOptions.
func New__AWS_OpenSearchService_Domain_EBSOptions() AWS_OpenSearchService_Domain_EBSOptions {
	return AWS_OpenSearchService_Domain_EBSOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_OpenSearchService_Domain_EBSOptions) GetType() string {
	return AWS_OpenSearchService_Domain_EBSOptions__Type
}

// Set__EBSEnabled updates property "EBSEnabled".
func (t AWS_OpenSearchService_Domain_EBSOptions) Set__EBSEnabled(v cfz.Expression[bool]) AWS_OpenSearchService_Domain_EBSOptions {
	t.EBSEnabled = v
	return t
}

// SetV__EBSEnabled updates property "EBSEnabled".
func (t AWS_OpenSearchService_Domain_EBSOptions) SetV__EBSEnabled(v bool) AWS_OpenSearchService_Domain_EBSOptions {
	t.EBSEnabled = cfz.V(v)
	return t
}

// Set__Iops updates property "Iops".
func (t AWS_OpenSearchService_Domain_EBSOptions) Set__Iops(v cfz.Expression[int32]) AWS_OpenSearchService_Domain_EBSOptions {
	t.Iops = v
	return t
}

// SetV__Iops updates property "Iops".
func (t AWS_OpenSearchService_Domain_EBSOptions) SetV__Iops(v int32) AWS_OpenSearchService_Domain_EBSOptions {
	t.Iops = cfz.V(v)
	return t
}

// Set__Throughput updates property "Throughput".
func (t AWS_OpenSearchService_Domain_EBSOptions) Set__Throughput(v cfz.Expression[int32]) AWS_OpenSearchService_Domain_EBSOptions {
	t.Throughput = v
	return t
}

// SetV__Throughput updates property "Throughput".
func (t AWS_OpenSearchService_Domain_EBSOptions) SetV__Throughput(v int32) AWS_OpenSearchService_Domain_EBSOptions {
	t.Throughput = cfz.V(v)
	return t
}

// Set__VolumeSize updates property "VolumeSize".
func (t AWS_OpenSearchService_Domain_EBSOptions) Set__VolumeSize(v cfz.Expression[int32]) AWS_OpenSearchService_Domain_EBSOptions {
	t.VolumeSize = v
	return t
}

// SetV__VolumeSize updates property "VolumeSize".
func (t AWS_OpenSearchService_Domain_EBSOptions) SetV__VolumeSize(v int32) AWS_OpenSearchService_Domain_EBSOptions {
	t.VolumeSize = cfz.V(v)
	return t
}

// Set__VolumeType updates property "VolumeType".
func (t AWS_OpenSearchService_Domain_EBSOptions) Set__VolumeType(v cfz.Expression[string]) AWS_OpenSearchService_Domain_EBSOptions {
	t.VolumeType = v
	return t
}

// SetV__VolumeType updates property "VolumeType".
func (t AWS_OpenSearchService_Domain_EBSOptions) SetV__VolumeType(v string) AWS_OpenSearchService_Domain_EBSOptions {
	t.VolumeType = cfz.V(v)
	return t
}
