// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appconfig

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppConfig_Deployment_Tags__Type is the CloudFormation type for AWS::AppConfig::Deployment.Tags.
	AWS_AppConfig_Deployment_Tags__Type = "AWS::AppConfig::Deployment.Tags"
)

var (
	// AWS_AppConfig_Deployment_Tags__PropertiesMap reports all the CloudFormation properties for AWS::AppConfig::Deployment.Tags.
	AWS_AppConfig_Deployment_Tags__PropertiesMap = struct {
		Key   string
		Value string
	}{
		Key:   "Key",
		Value: "Value",
	}

	// AWS_AppConfig_Deployment_Tags__PropertiesSlice reports all the CloudFormation properties for AWS::AppConfig::Deployment.Tags.
	AWS_AppConfig_Deployment_Tags__PropertiesSlice = []string{
		AWS_AppConfig_Deployment_Tags__PropertiesMap.Key,
		AWS_AppConfig_Deployment_Tags__PropertiesMap.Value,
	}
)

// AWS_AppConfig_Deployment_Tags is a binding for AWS::AppConfig::Deployment.Tags.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html
type AWS_AppConfig_Deployment_Tags struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html#cfn-appconfig-deployment-tags-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html#cfn-appconfig-deployment-tags-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_AppConfig_Deployment_Tags initializes a new AWS_AppConfig_Deployment_Tags.
func New__AWS_AppConfig_Deployment_Tags() AWS_AppConfig_Deployment_Tags {
	return AWS_AppConfig_Deployment_Tags{}
}

// GetType returns the CloudFormation type.
func (AWS_AppConfig_Deployment_Tags) GetType() string {
	return AWS_AppConfig_Deployment_Tags__Type
}

// Set__Key updates property "Key".
func (t AWS_AppConfig_Deployment_Tags) Set__Key(v cfz.Expression[string]) AWS_AppConfig_Deployment_Tags {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_AppConfig_Deployment_Tags) SetV__Key(v string) AWS_AppConfig_Deployment_Tags {
	t.Key = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_AppConfig_Deployment_Tags) Set__Value(v cfz.Expression[string]) AWS_AppConfig_Deployment_Tags {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_AppConfig_Deployment_Tags) SetV__Value(v string) AWS_AppConfig_Deployment_Tags {
	t.Value = cfz.V(v)
	return t
}
