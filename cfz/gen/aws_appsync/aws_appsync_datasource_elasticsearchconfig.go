// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appsync

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppSync_DataSource_ElasticsearchConfig__Type is the CloudFormation type for AWS::AppSync::DataSource.ElasticsearchConfig.
	AWS_AppSync_DataSource_ElasticsearchConfig__Type = "AWS::AppSync::DataSource.ElasticsearchConfig"
)

var (
	// AWS_AppSync_DataSource_ElasticsearchConfig__PropertiesMap reports all the CloudFormation properties for AWS::AppSync::DataSource.ElasticsearchConfig.
	AWS_AppSync_DataSource_ElasticsearchConfig__PropertiesMap = struct {
		AwsRegion string
		Endpoint  string
	}{
		AwsRegion: "AwsRegion",
		Endpoint:  "Endpoint",
	}

	// AWS_AppSync_DataSource_ElasticsearchConfig__PropertiesSlice reports all the CloudFormation properties for AWS::AppSync::DataSource.ElasticsearchConfig.
	AWS_AppSync_DataSource_ElasticsearchConfig__PropertiesSlice = []string{
		AWS_AppSync_DataSource_ElasticsearchConfig__PropertiesMap.AwsRegion,
		AWS_AppSync_DataSource_ElasticsearchConfig__PropertiesMap.Endpoint,
	}
)

// AWS_AppSync_DataSource_ElasticsearchConfig is a binding for AWS::AppSync::DataSource.ElasticsearchConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-elasticsearchconfig.html
type AWS_AppSync_DataSource_ElasticsearchConfig struct {
	// AwsRegion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-elasticsearchconfig.html#cfn-appsync-datasource-elasticsearchconfig-awsregion
	AwsRegion cfz.Expression[string] `json:"AwsRegion,omitempty"`

	// Endpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-elasticsearchconfig.html#cfn-appsync-datasource-elasticsearchconfig-endpoint
	Endpoint cfz.Expression[string] `json:"Endpoint,omitempty"`
}

// New__AWS_AppSync_DataSource_ElasticsearchConfig initializes a new AWS_AppSync_DataSource_ElasticsearchConfig.
func New__AWS_AppSync_DataSource_ElasticsearchConfig() AWS_AppSync_DataSource_ElasticsearchConfig {
	return AWS_AppSync_DataSource_ElasticsearchConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_AppSync_DataSource_ElasticsearchConfig) GetType() string {
	return AWS_AppSync_DataSource_ElasticsearchConfig__Type
}

// Set__AwsRegion updates property "AwsRegion".
func (t AWS_AppSync_DataSource_ElasticsearchConfig) Set__AwsRegion(v cfz.Expression[string]) AWS_AppSync_DataSource_ElasticsearchConfig {
	t.AwsRegion = v
	return t
}

// SetV__AwsRegion updates property "AwsRegion".
func (t AWS_AppSync_DataSource_ElasticsearchConfig) SetV__AwsRegion(v string) AWS_AppSync_DataSource_ElasticsearchConfig {
	t.AwsRegion = cfz.V(v)
	return t
}

// Set__Endpoint updates property "Endpoint".
func (t AWS_AppSync_DataSource_ElasticsearchConfig) Set__Endpoint(v cfz.Expression[string]) AWS_AppSync_DataSource_ElasticsearchConfig {
	t.Endpoint = v
	return t
}

// SetV__Endpoint updates property "Endpoint".
func (t AWS_AppSync_DataSource_ElasticsearchConfig) SetV__Endpoint(v string) AWS_AppSync_DataSource_ElasticsearchConfig {
	t.Endpoint = cfz.V(v)
	return t
}
