// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_databrew

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataBrew_Job_ValidationConfiguration__Type is the CloudFormation type for AWS::DataBrew::Job.ValidationConfiguration.
	AWS_DataBrew_Job_ValidationConfiguration__Type = "AWS::DataBrew::Job.ValidationConfiguration"
)

var (
	// AWS_DataBrew_Job_ValidationConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::DataBrew::Job.ValidationConfiguration.
	AWS_DataBrew_Job_ValidationConfiguration__PropertiesMap = struct {
		RulesetArn     string
		ValidationMode string
	}{
		RulesetArn:     "RulesetArn",
		ValidationMode: "ValidationMode",
	}

	// AWS_DataBrew_Job_ValidationConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::DataBrew::Job.ValidationConfiguration.
	AWS_DataBrew_Job_ValidationConfiguration__PropertiesSlice = []string{
		AWS_DataBrew_Job_ValidationConfiguration__PropertiesMap.RulesetArn,
		AWS_DataBrew_Job_ValidationConfiguration__PropertiesMap.ValidationMode,
	}
)

// AWS_DataBrew_Job_ValidationConfiguration is a binding for AWS::DataBrew::Job.ValidationConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-validationconfiguration.html
type AWS_DataBrew_Job_ValidationConfiguration struct {
	// RulesetArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-validationconfiguration.html#cfn-databrew-job-validationconfiguration-rulesetarn
	RulesetArn cfz.Expression[string] `json:"RulesetArn,omitempty"`

	// ValidationMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-validationconfiguration.html#cfn-databrew-job-validationconfiguration-validationmode
	ValidationMode cfz.Expression[string] `json:"ValidationMode,omitempty"`
}

// New__AWS_DataBrew_Job_ValidationConfiguration initializes a new AWS_DataBrew_Job_ValidationConfiguration.
func New__AWS_DataBrew_Job_ValidationConfiguration() AWS_DataBrew_Job_ValidationConfiguration {
	return AWS_DataBrew_Job_ValidationConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_DataBrew_Job_ValidationConfiguration) GetType() string {
	return AWS_DataBrew_Job_ValidationConfiguration__Type
}

// Set__RulesetArn updates property "RulesetArn".
func (t AWS_DataBrew_Job_ValidationConfiguration) Set__RulesetArn(v cfz.Expression[string]) AWS_DataBrew_Job_ValidationConfiguration {
	t.RulesetArn = v
	return t
}

// SetV__RulesetArn updates property "RulesetArn".
func (t AWS_DataBrew_Job_ValidationConfiguration) SetV__RulesetArn(v string) AWS_DataBrew_Job_ValidationConfiguration {
	t.RulesetArn = cfz.V(v)
	return t
}

// Set__ValidationMode updates property "ValidationMode".
func (t AWS_DataBrew_Job_ValidationConfiguration) Set__ValidationMode(v cfz.Expression[string]) AWS_DataBrew_Job_ValidationConfiguration {
	t.ValidationMode = v
	return t
}

// SetV__ValidationMode updates property "ValidationMode".
func (t AWS_DataBrew_Job_ValidationConfiguration) SetV__ValidationMode(v string) AWS_DataBrew_Job_ValidationConfiguration {
	t.ValidationMode = cfz.V(v)
	return t
}
