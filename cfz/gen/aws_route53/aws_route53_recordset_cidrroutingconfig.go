// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_route53

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Route53_RecordSet_CidrRoutingConfig__Type is the CloudFormation type for AWS::Route53::RecordSet.CidrRoutingConfig.
	AWS_Route53_RecordSet_CidrRoutingConfig__Type = "AWS::Route53::RecordSet.CidrRoutingConfig"
)

var (
	// AWS_Route53_RecordSet_CidrRoutingConfig__PropertiesMap reports all the CloudFormation properties for AWS::Route53::RecordSet.CidrRoutingConfig.
	AWS_Route53_RecordSet_CidrRoutingConfig__PropertiesMap = struct {
		CollectionId string
		LocationName string
	}{
		CollectionId: "CollectionId",
		LocationName: "LocationName",
	}

	// AWS_Route53_RecordSet_CidrRoutingConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Route53::RecordSet.CidrRoutingConfig.
	AWS_Route53_RecordSet_CidrRoutingConfig__PropertiesSlice = []string{
		AWS_Route53_RecordSet_CidrRoutingConfig__PropertiesMap.CollectionId,
		AWS_Route53_RecordSet_CidrRoutingConfig__PropertiesMap.LocationName,
	}
)

// AWS_Route53_RecordSet_CidrRoutingConfig is a binding for AWS::Route53::RecordSet.CidrRoutingConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-cidrroutingconfig.html
type AWS_Route53_RecordSet_CidrRoutingConfig struct {
	// CollectionId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-cidrroutingconfig.html#cfn-route53-cidrroutingconfig-collectionid
	CollectionId cfz.Expression[string] `json:"CollectionId,omitempty"`

	// LocationName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-cidrroutingconfig.html#cfn-route53-cidrroutingconfig-locationname
	LocationName cfz.Expression[string] `json:"LocationName,omitempty"`
}

// New__AWS_Route53_RecordSet_CidrRoutingConfig initializes a new AWS_Route53_RecordSet_CidrRoutingConfig.
func New__AWS_Route53_RecordSet_CidrRoutingConfig() AWS_Route53_RecordSet_CidrRoutingConfig {
	return AWS_Route53_RecordSet_CidrRoutingConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Route53_RecordSet_CidrRoutingConfig) GetType() string {
	return AWS_Route53_RecordSet_CidrRoutingConfig__Type
}

// Set__CollectionId updates property "CollectionId".
func (t AWS_Route53_RecordSet_CidrRoutingConfig) Set__CollectionId(v cfz.Expression[string]) AWS_Route53_RecordSet_CidrRoutingConfig {
	t.CollectionId = v
	return t
}

// SetV__CollectionId updates property "CollectionId".
func (t AWS_Route53_RecordSet_CidrRoutingConfig) SetV__CollectionId(v string) AWS_Route53_RecordSet_CidrRoutingConfig {
	t.CollectionId = cfz.V(v)
	return t
}

// Set__LocationName updates property "LocationName".
func (t AWS_Route53_RecordSet_CidrRoutingConfig) Set__LocationName(v cfz.Expression[string]) AWS_Route53_RecordSet_CidrRoutingConfig {
	t.LocationName = v
	return t
}

// SetV__LocationName updates property "LocationName".
func (t AWS_Route53_RecordSet_CidrRoutingConfig) SetV__LocationName(v string) AWS_Route53_RecordSet_CidrRoutingConfig {
	t.LocationName = cfz.V(v)
	return t
}
