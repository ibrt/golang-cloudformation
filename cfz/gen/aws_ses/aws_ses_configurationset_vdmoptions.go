// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ses

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SES_ConfigurationSet_VdmOptions__Type is the CloudFormation type for AWS::SES::ConfigurationSet.VdmOptions.
	AWS_SES_ConfigurationSet_VdmOptions__Type = "AWS::SES::ConfigurationSet.VdmOptions"
)

var (
	// AWS_SES_ConfigurationSet_VdmOptions__PropertiesMap reports all the CloudFormation properties for AWS::SES::ConfigurationSet.VdmOptions.
	AWS_SES_ConfigurationSet_VdmOptions__PropertiesMap = struct {
		DashboardOptions string
		GuardianOptions  string
	}{
		DashboardOptions: "DashboardOptions",
		GuardianOptions:  "GuardianOptions",
	}

	// AWS_SES_ConfigurationSet_VdmOptions__PropertiesSlice reports all the CloudFormation properties for AWS::SES::ConfigurationSet.VdmOptions.
	AWS_SES_ConfigurationSet_VdmOptions__PropertiesSlice = []string{
		AWS_SES_ConfigurationSet_VdmOptions__PropertiesMap.DashboardOptions,
		AWS_SES_ConfigurationSet_VdmOptions__PropertiesMap.GuardianOptions,
	}
)

// AWS_SES_ConfigurationSet_VdmOptions is a binding for AWS::SES::ConfigurationSet.VdmOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-vdmoptions.html
type AWS_SES_ConfigurationSet_VdmOptions struct {
	// DashboardOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-vdmoptions.html#cfn-ses-configurationset-vdmoptions-dashboardoptions
	DashboardOptions cfz.Expression[AWS_SES_ConfigurationSet_DashboardOptions] `json:"DashboardOptions,omitempty"`

	// GuardianOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationset-vdmoptions.html#cfn-ses-configurationset-vdmoptions-guardianoptions
	GuardianOptions cfz.Expression[AWS_SES_ConfigurationSet_GuardianOptions] `json:"GuardianOptions,omitempty"`
}

// New__AWS_SES_ConfigurationSet_VdmOptions initializes a new AWS_SES_ConfigurationSet_VdmOptions.
func New__AWS_SES_ConfigurationSet_VdmOptions() AWS_SES_ConfigurationSet_VdmOptions {
	return AWS_SES_ConfigurationSet_VdmOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_SES_ConfigurationSet_VdmOptions) GetType() string {
	return AWS_SES_ConfigurationSet_VdmOptions__Type
}

// Set__DashboardOptions updates property "DashboardOptions".
func (t AWS_SES_ConfigurationSet_VdmOptions) Set__DashboardOptions(v cfz.Expression[AWS_SES_ConfigurationSet_DashboardOptions]) AWS_SES_ConfigurationSet_VdmOptions {
	t.DashboardOptions = v
	return t
}

// SetV__DashboardOptions updates property "DashboardOptions".
func (t AWS_SES_ConfigurationSet_VdmOptions) SetV__DashboardOptions(v AWS_SES_ConfigurationSet_DashboardOptions) AWS_SES_ConfigurationSet_VdmOptions {
	t.DashboardOptions = cfz.V(v)
	return t
}

// Set__GuardianOptions updates property "GuardianOptions".
func (t AWS_SES_ConfigurationSet_VdmOptions) Set__GuardianOptions(v cfz.Expression[AWS_SES_ConfigurationSet_GuardianOptions]) AWS_SES_ConfigurationSet_VdmOptions {
	t.GuardianOptions = v
	return t
}

// SetV__GuardianOptions updates property "GuardianOptions".
func (t AWS_SES_ConfigurationSet_VdmOptions) SetV__GuardianOptions(v AWS_SES_ConfigurationSet_GuardianOptions) AWS_SES_ConfigurationSet_VdmOptions {
	t.GuardianOptions = cfz.V(v)
	return t
}
