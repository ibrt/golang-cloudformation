// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_elasticache

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails__Type is the CloudFormation type for AWS::ElastiCache::ReplicationGroup.KinesisFirehoseDestinationDetails.
	AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails__Type = "AWS::ElastiCache::ReplicationGroup.KinesisFirehoseDestinationDetails"
)

var (
	// AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails__PropertiesMap reports all the CloudFormation properties for AWS::ElastiCache::ReplicationGroup.KinesisFirehoseDestinationDetails.
	AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails__PropertiesMap = struct {
		DeliveryStream string
	}{
		DeliveryStream: "DeliveryStream",
	}

	// AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails__PropertiesSlice reports all the CloudFormation properties for AWS::ElastiCache::ReplicationGroup.KinesisFirehoseDestinationDetails.
	AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails__PropertiesSlice = []string{
		AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails__PropertiesMap.DeliveryStream,
	}
)

// AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails is a binding for AWS::ElastiCache::ReplicationGroup.KinesisFirehoseDestinationDetails.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-kinesisfirehosedestinationdetails.html
type AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails struct {
	// DeliveryStream is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-kinesisfirehosedestinationdetails.html#cfn-elasticache-replicationgroup-kinesisfirehosedestinationdetails-deliverystream
	DeliveryStream cfz.Expression[string] `json:"DeliveryStream,omitempty"`
}

// New__AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails initializes a new AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails.
func New__AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails() AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails {
	return AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails{}
}

// GetType returns the CloudFormation type.
func (AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails) GetType() string {
	return AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails__Type
}

// Set__DeliveryStream updates property "DeliveryStream".
func (t AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails) Set__DeliveryStream(v cfz.Expression[string]) AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails {
	t.DeliveryStream = v
	return t
}

// SetV__DeliveryStream updates property "DeliveryStream".
func (t AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails) SetV__DeliveryStream(v string) AWS_ElastiCache_ReplicationGroup_KinesisFirehoseDestinationDetails {
	t.DeliveryStream = cfz.V(v)
	return t
}
