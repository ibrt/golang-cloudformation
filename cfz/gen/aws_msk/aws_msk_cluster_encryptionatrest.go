// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_msk

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MSK_Cluster_EncryptionAtRest__Type is the CloudFormation type for AWS::MSK::Cluster.EncryptionAtRest.
	AWS_MSK_Cluster_EncryptionAtRest__Type = "AWS::MSK::Cluster.EncryptionAtRest"
)

var (
	// AWS_MSK_Cluster_EncryptionAtRest__PropertiesMap reports all the CloudFormation properties for AWS::MSK::Cluster.EncryptionAtRest.
	AWS_MSK_Cluster_EncryptionAtRest__PropertiesMap = struct {
		DataVolumeKMSKeyId string
	}{
		DataVolumeKMSKeyId: "DataVolumeKMSKeyId",
	}

	// AWS_MSK_Cluster_EncryptionAtRest__PropertiesSlice reports all the CloudFormation properties for AWS::MSK::Cluster.EncryptionAtRest.
	AWS_MSK_Cluster_EncryptionAtRest__PropertiesSlice = []string{
		AWS_MSK_Cluster_EncryptionAtRest__PropertiesMap.DataVolumeKMSKeyId,
	}
)

// AWS_MSK_Cluster_EncryptionAtRest is a binding for AWS::MSK::Cluster.EncryptionAtRest.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionatrest.html
type AWS_MSK_Cluster_EncryptionAtRest struct {
	// DataVolumeKMSKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionatrest.html#cfn-msk-cluster-encryptionatrest-datavolumekmskeyid
	DataVolumeKMSKeyId cfz.Expression[string] `json:"DataVolumeKMSKeyId,omitempty"`
}

// New__AWS_MSK_Cluster_EncryptionAtRest initializes a new AWS_MSK_Cluster_EncryptionAtRest.
func New__AWS_MSK_Cluster_EncryptionAtRest() AWS_MSK_Cluster_EncryptionAtRest {
	return AWS_MSK_Cluster_EncryptionAtRest{}
}

// GetType returns the CloudFormation type.
func (AWS_MSK_Cluster_EncryptionAtRest) GetType() string {
	return AWS_MSK_Cluster_EncryptionAtRest__Type
}

// Set__DataVolumeKMSKeyId updates property "DataVolumeKMSKeyId".
func (t AWS_MSK_Cluster_EncryptionAtRest) Set__DataVolumeKMSKeyId(v cfz.Expression[string]) AWS_MSK_Cluster_EncryptionAtRest {
	t.DataVolumeKMSKeyId = v
	return t
}

// SetV__DataVolumeKMSKeyId updates property "DataVolumeKMSKeyId".
func (t AWS_MSK_Cluster_EncryptionAtRest) SetV__DataVolumeKMSKeyId(v string) AWS_MSK_Cluster_EncryptionAtRest {
	t.DataVolumeKMSKeyId = cfz.V(v)
	return t
}
