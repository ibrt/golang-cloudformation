// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ecr

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ECR_ReplicationConfiguration_ReplicationDestination__Type is the CloudFormation type for AWS::ECR::ReplicationConfiguration.ReplicationDestination.
	AWS_ECR_ReplicationConfiguration_ReplicationDestination__Type = "AWS::ECR::ReplicationConfiguration.ReplicationDestination"
)

var (
	// AWS_ECR_ReplicationConfiguration_ReplicationDestination__PropertiesMap reports all the CloudFormation properties for AWS::ECR::ReplicationConfiguration.ReplicationDestination.
	AWS_ECR_ReplicationConfiguration_ReplicationDestination__PropertiesMap = struct {
		Region     string
		RegistryId string
	}{
		Region:     "Region",
		RegistryId: "RegistryId",
	}

	// AWS_ECR_ReplicationConfiguration_ReplicationDestination__PropertiesSlice reports all the CloudFormation properties for AWS::ECR::ReplicationConfiguration.ReplicationDestination.
	AWS_ECR_ReplicationConfiguration_ReplicationDestination__PropertiesSlice = []string{
		AWS_ECR_ReplicationConfiguration_ReplicationDestination__PropertiesMap.Region,
		AWS_ECR_ReplicationConfiguration_ReplicationDestination__PropertiesMap.RegistryId,
	}
)

// AWS_ECR_ReplicationConfiguration_ReplicationDestination is a binding for AWS::ECR::ReplicationConfiguration.ReplicationDestination.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-replicationdestination.html
type AWS_ECR_ReplicationConfiguration_ReplicationDestination struct {
	// Region is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-replicationdestination.html#cfn-ecr-replicationconfiguration-replicationdestination-region
	Region cfz.Expression[string] `json:"Region,omitempty"`

	// RegistryId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-replicationdestination.html#cfn-ecr-replicationconfiguration-replicationdestination-registryid
	RegistryId cfz.Expression[string] `json:"RegistryId,omitempty"`
}

// New__AWS_ECR_ReplicationConfiguration_ReplicationDestination initializes a new AWS_ECR_ReplicationConfiguration_ReplicationDestination.
func New__AWS_ECR_ReplicationConfiguration_ReplicationDestination() AWS_ECR_ReplicationConfiguration_ReplicationDestination {
	return AWS_ECR_ReplicationConfiguration_ReplicationDestination{}
}

// GetType returns the CloudFormation type.
func (AWS_ECR_ReplicationConfiguration_ReplicationDestination) GetType() string {
	return AWS_ECR_ReplicationConfiguration_ReplicationDestination__Type
}

// Set__Region updates property "Region".
func (t AWS_ECR_ReplicationConfiguration_ReplicationDestination) Set__Region(v cfz.Expression[string]) AWS_ECR_ReplicationConfiguration_ReplicationDestination {
	t.Region = v
	return t
}

// SetV__Region updates property "Region".
func (t AWS_ECR_ReplicationConfiguration_ReplicationDestination) SetV__Region(v string) AWS_ECR_ReplicationConfiguration_ReplicationDestination {
	t.Region = cfz.V(v)
	return t
}

// Set__RegistryId updates property "RegistryId".
func (t AWS_ECR_ReplicationConfiguration_ReplicationDestination) Set__RegistryId(v cfz.Expression[string]) AWS_ECR_ReplicationConfiguration_ReplicationDestination {
	t.RegistryId = v
	return t
}

// SetV__RegistryId updates property "RegistryId".
func (t AWS_ECR_ReplicationConfiguration_ReplicationDestination) SetV__RegistryId(v string) AWS_ECR_ReplicationConfiguration_ReplicationDestination {
	t.RegistryId = cfz.V(v)
	return t
}
