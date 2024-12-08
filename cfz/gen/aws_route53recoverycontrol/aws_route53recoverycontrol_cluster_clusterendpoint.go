// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_route53recoverycontrol

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Route53RecoveryControl_Cluster_ClusterEndpoint__Type is the CloudFormation type for AWS::Route53RecoveryControl::Cluster.ClusterEndpoint.
	AWS_Route53RecoveryControl_Cluster_ClusterEndpoint__Type = "AWS::Route53RecoveryControl::Cluster.ClusterEndpoint"
)

var (
	// AWS_Route53RecoveryControl_Cluster_ClusterEndpoint__PropertiesMap reports all the CloudFormation properties for AWS::Route53RecoveryControl::Cluster.ClusterEndpoint.
	AWS_Route53RecoveryControl_Cluster_ClusterEndpoint__PropertiesMap = struct {
		Endpoint string
		Region   string
	}{
		Endpoint: "Endpoint",
		Region:   "Region",
	}

	// AWS_Route53RecoveryControl_Cluster_ClusterEndpoint__PropertiesSlice reports all the CloudFormation properties for AWS::Route53RecoveryControl::Cluster.ClusterEndpoint.
	AWS_Route53RecoveryControl_Cluster_ClusterEndpoint__PropertiesSlice = []string{
		AWS_Route53RecoveryControl_Cluster_ClusterEndpoint__PropertiesMap.Endpoint,
		AWS_Route53RecoveryControl_Cluster_ClusterEndpoint__PropertiesMap.Region,
	}
)

// AWS_Route53RecoveryControl_Cluster_ClusterEndpoint is a binding for AWS::Route53RecoveryControl::Cluster.ClusterEndpoint.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoverycontrol-cluster-clusterendpoint.html
type AWS_Route53RecoveryControl_Cluster_ClusterEndpoint struct {
	// Endpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoverycontrol-cluster-clusterendpoint.html#cfn-route53recoverycontrol-cluster-clusterendpoint-endpoint
	Endpoint cfz.Expression[string] `json:"Endpoint,omitempty"`

	// Region is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoverycontrol-cluster-clusterendpoint.html#cfn-route53recoverycontrol-cluster-clusterendpoint-region
	Region cfz.Expression[string] `json:"Region,omitempty"`
}

// New__AWS_Route53RecoveryControl_Cluster_ClusterEndpoint initializes a new AWS_Route53RecoveryControl_Cluster_ClusterEndpoint.
func New__AWS_Route53RecoveryControl_Cluster_ClusterEndpoint() AWS_Route53RecoveryControl_Cluster_ClusterEndpoint {
	return AWS_Route53RecoveryControl_Cluster_ClusterEndpoint{}
}

// GetType returns the CloudFormation type.
func (AWS_Route53RecoveryControl_Cluster_ClusterEndpoint) GetType() string {
	return AWS_Route53RecoveryControl_Cluster_ClusterEndpoint__Type
}

// Set__Endpoint updates property "Endpoint".
func (t AWS_Route53RecoveryControl_Cluster_ClusterEndpoint) Set__Endpoint(v cfz.Expression[string]) AWS_Route53RecoveryControl_Cluster_ClusterEndpoint {
	t.Endpoint = v
	return t
}

// SetV__Endpoint updates property "Endpoint".
func (t AWS_Route53RecoveryControl_Cluster_ClusterEndpoint) SetV__Endpoint(v string) AWS_Route53RecoveryControl_Cluster_ClusterEndpoint {
	t.Endpoint = cfz.V(v)
	return t
}

// Set__Region updates property "Region".
func (t AWS_Route53RecoveryControl_Cluster_ClusterEndpoint) Set__Region(v cfz.Expression[string]) AWS_Route53RecoveryControl_Cluster_ClusterEndpoint {
	t.Region = v
	return t
}

// SetV__Region updates property "Region".
func (t AWS_Route53RecoveryControl_Cluster_ClusterEndpoint) SetV__Region(v string) AWS_Route53RecoveryControl_Cluster_ClusterEndpoint {
	t.Region = cfz.V(v)
	return t
}
