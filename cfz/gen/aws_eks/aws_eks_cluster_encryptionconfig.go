// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_eks

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EKS_Cluster_EncryptionConfig__Type is the CloudFormation type for AWS::EKS::Cluster.EncryptionConfig.
	AWS_EKS_Cluster_EncryptionConfig__Type = "AWS::EKS::Cluster.EncryptionConfig"
)

var (
	// AWS_EKS_Cluster_EncryptionConfig__PropertiesMap reports all the CloudFormation properties for AWS::EKS::Cluster.EncryptionConfig.
	AWS_EKS_Cluster_EncryptionConfig__PropertiesMap = struct {
		Provider  string
		Resources string
	}{
		Provider:  "Provider",
		Resources: "Resources",
	}

	// AWS_EKS_Cluster_EncryptionConfig__PropertiesSlice reports all the CloudFormation properties for AWS::EKS::Cluster.EncryptionConfig.
	AWS_EKS_Cluster_EncryptionConfig__PropertiesSlice = []string{
		AWS_EKS_Cluster_EncryptionConfig__PropertiesMap.Provider,
		AWS_EKS_Cluster_EncryptionConfig__PropertiesMap.Resources,
	}
)

// AWS_EKS_Cluster_EncryptionConfig is a binding for AWS::EKS::Cluster.EncryptionConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-encryptionconfig.html
type AWS_EKS_Cluster_EncryptionConfig struct {
	// Provider is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-encryptionconfig.html#cfn-eks-cluster-encryptionconfig-provider
	Provider cfz.Expression[AWS_EKS_Cluster_Provider] `json:"Provider,omitempty"`

	// Resources is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-encryptionconfig.html#cfn-eks-cluster-encryptionconfig-resources
	Resources cfz.ExpressionSlice[string] `json:"Resources,omitempty"`
}

// New__AWS_EKS_Cluster_EncryptionConfig initializes a new AWS_EKS_Cluster_EncryptionConfig.
func New__AWS_EKS_Cluster_EncryptionConfig() AWS_EKS_Cluster_EncryptionConfig {
	return AWS_EKS_Cluster_EncryptionConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_EKS_Cluster_EncryptionConfig) GetType() string {
	return AWS_EKS_Cluster_EncryptionConfig__Type
}

// Set__Provider updates property "Provider".
func (t AWS_EKS_Cluster_EncryptionConfig) Set__Provider(v cfz.Expression[AWS_EKS_Cluster_Provider]) AWS_EKS_Cluster_EncryptionConfig {
	t.Provider = v
	return t
}

// SetV__Provider updates property "Provider".
func (t AWS_EKS_Cluster_EncryptionConfig) SetV__Provider(v AWS_EKS_Cluster_Provider) AWS_EKS_Cluster_EncryptionConfig {
	t.Provider = cfz.V(v)
	return t
}

// Set__Resources updates property "Resources".
func (t AWS_EKS_Cluster_EncryptionConfig) Set__Resources(v cfz.ExpressionSlice[string]) AWS_EKS_Cluster_EncryptionConfig {
	t.Resources = v
	return t
}

// SetS__Resources updates property "Resources".
func (t AWS_EKS_Cluster_EncryptionConfig) SetS__Resources(v ...cfz.Expression[string]) AWS_EKS_Cluster_EncryptionConfig {
	t.Resources = cfz.S(v...)
	return t
}

// SetSV__Resources updates property "Resources".
func (t AWS_EKS_Cluster_EncryptionConfig) SetSV__Resources(v ...string) AWS_EKS_Cluster_EncryptionConfig {
	t.Resources = cfz.SV(v...)
	return t
}
