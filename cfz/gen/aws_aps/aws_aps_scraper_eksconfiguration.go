// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_aps

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_APS_Scraper_EksConfiguration__Type is the CloudFormation type for AWS::APS::Scraper.EksConfiguration.
	AWS_APS_Scraper_EksConfiguration__Type = "AWS::APS::Scraper.EksConfiguration"
)

var (
	// AWS_APS_Scraper_EksConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::APS::Scraper.EksConfiguration.
	AWS_APS_Scraper_EksConfiguration__PropertiesMap = struct {
		ClusterArn       string
		SecurityGroupIds string
		SubnetIds        string
	}{
		ClusterArn:       "ClusterArn",
		SecurityGroupIds: "SecurityGroupIds",
		SubnetIds:        "SubnetIds",
	}

	// AWS_APS_Scraper_EksConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::APS::Scraper.EksConfiguration.
	AWS_APS_Scraper_EksConfiguration__PropertiesSlice = []string{
		AWS_APS_Scraper_EksConfiguration__PropertiesMap.ClusterArn,
		AWS_APS_Scraper_EksConfiguration__PropertiesMap.SecurityGroupIds,
		AWS_APS_Scraper_EksConfiguration__PropertiesMap.SubnetIds,
	}
)

// AWS_APS_Scraper_EksConfiguration is a binding for AWS::APS::Scraper.EksConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-eksconfiguration.html
type AWS_APS_Scraper_EksConfiguration struct {
	// ClusterArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-eksconfiguration.html#cfn-aps-scraper-eksconfiguration-clusterarn
	ClusterArn cfz.Expression[string] `json:"ClusterArn,omitempty"`

	// SecurityGroupIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-eksconfiguration.html#cfn-aps-scraper-eksconfiguration-securitygroupids
	SecurityGroupIds cfz.ExpressionSlice[string] `json:"SecurityGroupIds,omitempty"`

	// SubnetIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-eksconfiguration.html#cfn-aps-scraper-eksconfiguration-subnetids
	SubnetIds cfz.ExpressionSlice[string] `json:"SubnetIds,omitempty"`
}

// New__AWS_APS_Scraper_EksConfiguration initializes a new AWS_APS_Scraper_EksConfiguration.
func New__AWS_APS_Scraper_EksConfiguration() AWS_APS_Scraper_EksConfiguration {
	return AWS_APS_Scraper_EksConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_APS_Scraper_EksConfiguration) GetType() string {
	return AWS_APS_Scraper_EksConfiguration__Type
}

// Set__ClusterArn updates property "ClusterArn".
func (t AWS_APS_Scraper_EksConfiguration) Set__ClusterArn(v cfz.Expression[string]) AWS_APS_Scraper_EksConfiguration {
	t.ClusterArn = v
	return t
}

// SetV__ClusterArn updates property "ClusterArn".
func (t AWS_APS_Scraper_EksConfiguration) SetV__ClusterArn(v string) AWS_APS_Scraper_EksConfiguration {
	t.ClusterArn = cfz.V(v)
	return t
}

// Set__SecurityGroupIds updates property "SecurityGroupIds".
func (t AWS_APS_Scraper_EksConfiguration) Set__SecurityGroupIds(v cfz.ExpressionSlice[string]) AWS_APS_Scraper_EksConfiguration {
	t.SecurityGroupIds = v
	return t
}

// SetS__SecurityGroupIds updates property "SecurityGroupIds".
func (t AWS_APS_Scraper_EksConfiguration) SetS__SecurityGroupIds(v ...cfz.Expression[string]) AWS_APS_Scraper_EksConfiguration {
	t.SecurityGroupIds = cfz.S(v...)
	return t
}

// SetSV__SecurityGroupIds updates property "SecurityGroupIds".
func (t AWS_APS_Scraper_EksConfiguration) SetSV__SecurityGroupIds(v ...string) AWS_APS_Scraper_EksConfiguration {
	t.SecurityGroupIds = cfz.SV(v...)
	return t
}

// Set__SubnetIds updates property "SubnetIds".
func (t AWS_APS_Scraper_EksConfiguration) Set__SubnetIds(v cfz.ExpressionSlice[string]) AWS_APS_Scraper_EksConfiguration {
	t.SubnetIds = v
	return t
}

// SetS__SubnetIds updates property "SubnetIds".
func (t AWS_APS_Scraper_EksConfiguration) SetS__SubnetIds(v ...cfz.Expression[string]) AWS_APS_Scraper_EksConfiguration {
	t.SubnetIds = cfz.S(v...)
	return t
}

// SetSV__SubnetIds updates property "SubnetIds".
func (t AWS_APS_Scraper_EksConfiguration) SetSV__SubnetIds(v ...string) AWS_APS_Scraper_EksConfiguration {
	t.SubnetIds = cfz.SV(v...)
	return t
}
