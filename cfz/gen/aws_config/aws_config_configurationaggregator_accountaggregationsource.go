// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_config

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Config_ConfigurationAggregator_AccountAggregationSource__Type is the CloudFormation type for AWS::Config::ConfigurationAggregator.AccountAggregationSource.
	AWS_Config_ConfigurationAggregator_AccountAggregationSource__Type = "AWS::Config::ConfigurationAggregator.AccountAggregationSource"
)

var (
	// AWS_Config_ConfigurationAggregator_AccountAggregationSource__PropertiesMap reports all the CloudFormation properties for AWS::Config::ConfigurationAggregator.AccountAggregationSource.
	AWS_Config_ConfigurationAggregator_AccountAggregationSource__PropertiesMap = struct {
		AccountIds    string
		AllAwsRegions string
		AwsRegions    string
	}{
		AccountIds:    "AccountIds",
		AllAwsRegions: "AllAwsRegions",
		AwsRegions:    "AwsRegions",
	}

	// AWS_Config_ConfigurationAggregator_AccountAggregationSource__PropertiesSlice reports all the CloudFormation properties for AWS::Config::ConfigurationAggregator.AccountAggregationSource.
	AWS_Config_ConfigurationAggregator_AccountAggregationSource__PropertiesSlice = []string{
		AWS_Config_ConfigurationAggregator_AccountAggregationSource__PropertiesMap.AccountIds,
		AWS_Config_ConfigurationAggregator_AccountAggregationSource__PropertiesMap.AllAwsRegions,
		AWS_Config_ConfigurationAggregator_AccountAggregationSource__PropertiesMap.AwsRegions,
	}
)

// AWS_Config_ConfigurationAggregator_AccountAggregationSource is a binding for AWS::Config::ConfigurationAggregator.AccountAggregationSource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html
type AWS_Config_ConfigurationAggregator_AccountAggregationSource struct {
	// AccountIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html#cfn-config-configurationaggregator-accountaggregationsource-accountids
	AccountIds cfz.ExpressionSlice[string] `json:"AccountIds,omitempty"`

	// AllAwsRegions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html#cfn-config-configurationaggregator-accountaggregationsource-allawsregions
	AllAwsRegions cfz.Expression[bool] `json:"AllAwsRegions,omitempty"`

	// AwsRegions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html#cfn-config-configurationaggregator-accountaggregationsource-awsregions
	AwsRegions cfz.ExpressionSlice[string] `json:"AwsRegions,omitempty"`
}

// New__AWS_Config_ConfigurationAggregator_AccountAggregationSource initializes a new AWS_Config_ConfigurationAggregator_AccountAggregationSource.
func New__AWS_Config_ConfigurationAggregator_AccountAggregationSource() AWS_Config_ConfigurationAggregator_AccountAggregationSource {
	return AWS_Config_ConfigurationAggregator_AccountAggregationSource{}
}

// GetType returns the CloudFormation type.
func (AWS_Config_ConfigurationAggregator_AccountAggregationSource) GetType() string {
	return AWS_Config_ConfigurationAggregator_AccountAggregationSource__Type
}

// Set__AccountIds updates property "AccountIds".
func (t AWS_Config_ConfigurationAggregator_AccountAggregationSource) Set__AccountIds(v cfz.ExpressionSlice[string]) AWS_Config_ConfigurationAggregator_AccountAggregationSource {
	t.AccountIds = v
	return t
}

// SetS__AccountIds updates property "AccountIds".
func (t AWS_Config_ConfigurationAggregator_AccountAggregationSource) SetS__AccountIds(v ...cfz.Expression[string]) AWS_Config_ConfigurationAggregator_AccountAggregationSource {
	t.AccountIds = cfz.S(v...)
	return t
}

// SetSV__AccountIds updates property "AccountIds".
func (t AWS_Config_ConfigurationAggregator_AccountAggregationSource) SetSV__AccountIds(v ...string) AWS_Config_ConfigurationAggregator_AccountAggregationSource {
	t.AccountIds = cfz.SV(v...)
	return t
}

// Set__AllAwsRegions updates property "AllAwsRegions".
func (t AWS_Config_ConfigurationAggregator_AccountAggregationSource) Set__AllAwsRegions(v cfz.Expression[bool]) AWS_Config_ConfigurationAggregator_AccountAggregationSource {
	t.AllAwsRegions = v
	return t
}

// SetV__AllAwsRegions updates property "AllAwsRegions".
func (t AWS_Config_ConfigurationAggregator_AccountAggregationSource) SetV__AllAwsRegions(v bool) AWS_Config_ConfigurationAggregator_AccountAggregationSource {
	t.AllAwsRegions = cfz.V(v)
	return t
}

// Set__AwsRegions updates property "AwsRegions".
func (t AWS_Config_ConfigurationAggregator_AccountAggregationSource) Set__AwsRegions(v cfz.ExpressionSlice[string]) AWS_Config_ConfigurationAggregator_AccountAggregationSource {
	t.AwsRegions = v
	return t
}

// SetS__AwsRegions updates property "AwsRegions".
func (t AWS_Config_ConfigurationAggregator_AccountAggregationSource) SetS__AwsRegions(v ...cfz.Expression[string]) AWS_Config_ConfigurationAggregator_AccountAggregationSource {
	t.AwsRegions = cfz.S(v...)
	return t
}

// SetSV__AwsRegions updates property "AwsRegions".
func (t AWS_Config_ConfigurationAggregator_AccountAggregationSource) SetSV__AwsRegions(v ...string) AWS_Config_ConfigurationAggregator_AccountAggregationSource {
	t.AwsRegions = cfz.SV(v...)
	return t
}
