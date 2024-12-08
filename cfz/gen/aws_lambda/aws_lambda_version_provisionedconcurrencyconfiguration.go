// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lambda

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lambda_Version_ProvisionedConcurrencyConfiguration__Type is the CloudFormation type for AWS::Lambda::Version.ProvisionedConcurrencyConfiguration.
	AWS_Lambda_Version_ProvisionedConcurrencyConfiguration__Type = "AWS::Lambda::Version.ProvisionedConcurrencyConfiguration"
)

var (
	// AWS_Lambda_Version_ProvisionedConcurrencyConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Lambda::Version.ProvisionedConcurrencyConfiguration.
	AWS_Lambda_Version_ProvisionedConcurrencyConfiguration__PropertiesMap = struct {
		ProvisionedConcurrentExecutions string
	}{
		ProvisionedConcurrentExecutions: "ProvisionedConcurrentExecutions",
	}

	// AWS_Lambda_Version_ProvisionedConcurrencyConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Lambda::Version.ProvisionedConcurrencyConfiguration.
	AWS_Lambda_Version_ProvisionedConcurrencyConfiguration__PropertiesSlice = []string{
		AWS_Lambda_Version_ProvisionedConcurrencyConfiguration__PropertiesMap.ProvisionedConcurrentExecutions,
	}
)

// AWS_Lambda_Version_ProvisionedConcurrencyConfiguration is a binding for AWS::Lambda::Version.ProvisionedConcurrencyConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-version-provisionedconcurrencyconfiguration.html
type AWS_Lambda_Version_ProvisionedConcurrencyConfiguration struct {
	// ProvisionedConcurrentExecutions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-version-provisionedconcurrencyconfiguration.html#cfn-lambda-version-provisionedconcurrencyconfiguration-provisionedconcurrentexecutions
	ProvisionedConcurrentExecutions cfz.Expression[int32] `json:"ProvisionedConcurrentExecutions,omitempty"`
}

// New__AWS_Lambda_Version_ProvisionedConcurrencyConfiguration initializes a new AWS_Lambda_Version_ProvisionedConcurrencyConfiguration.
func New__AWS_Lambda_Version_ProvisionedConcurrencyConfiguration() AWS_Lambda_Version_ProvisionedConcurrencyConfiguration {
	return AWS_Lambda_Version_ProvisionedConcurrencyConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Lambda_Version_ProvisionedConcurrencyConfiguration) GetType() string {
	return AWS_Lambda_Version_ProvisionedConcurrencyConfiguration__Type
}

// Set__ProvisionedConcurrentExecutions updates property "ProvisionedConcurrentExecutions".
func (t AWS_Lambda_Version_ProvisionedConcurrencyConfiguration) Set__ProvisionedConcurrentExecutions(v cfz.Expression[int32]) AWS_Lambda_Version_ProvisionedConcurrencyConfiguration {
	t.ProvisionedConcurrentExecutions = v
	return t
}

// SetV__ProvisionedConcurrentExecutions updates property "ProvisionedConcurrentExecutions".
func (t AWS_Lambda_Version_ProvisionedConcurrencyConfiguration) SetV__ProvisionedConcurrentExecutions(v int32) AWS_Lambda_Version_ProvisionedConcurrencyConfiguration {
	t.ProvisionedConcurrentExecutions = cfz.V(v)
	return t
}
