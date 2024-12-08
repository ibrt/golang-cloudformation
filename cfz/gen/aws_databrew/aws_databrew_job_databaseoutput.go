// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_databrew

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataBrew_Job_DatabaseOutput__Type is the CloudFormation type for AWS::DataBrew::Job.DatabaseOutput.
	AWS_DataBrew_Job_DatabaseOutput__Type = "AWS::DataBrew::Job.DatabaseOutput"
)

var (
	// AWS_DataBrew_Job_DatabaseOutput__PropertiesMap reports all the CloudFormation properties for AWS::DataBrew::Job.DatabaseOutput.
	AWS_DataBrew_Job_DatabaseOutput__PropertiesMap = struct {
		DatabaseOptions    string
		DatabaseOutputMode string
		GlueConnectionName string
	}{
		DatabaseOptions:    "DatabaseOptions",
		DatabaseOutputMode: "DatabaseOutputMode",
		GlueConnectionName: "GlueConnectionName",
	}

	// AWS_DataBrew_Job_DatabaseOutput__PropertiesSlice reports all the CloudFormation properties for AWS::DataBrew::Job.DatabaseOutput.
	AWS_DataBrew_Job_DatabaseOutput__PropertiesSlice = []string{
		AWS_DataBrew_Job_DatabaseOutput__PropertiesMap.DatabaseOptions,
		AWS_DataBrew_Job_DatabaseOutput__PropertiesMap.DatabaseOutputMode,
		AWS_DataBrew_Job_DatabaseOutput__PropertiesMap.GlueConnectionName,
	}
)

// AWS_DataBrew_Job_DatabaseOutput is a binding for AWS::DataBrew::Job.DatabaseOutput.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databaseoutput.html
type AWS_DataBrew_Job_DatabaseOutput struct {
	// DatabaseOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databaseoutput.html#cfn-databrew-job-databaseoutput-databaseoptions
	DatabaseOptions cfz.Expression[AWS_DataBrew_Job_DatabaseTableOutputOptions] `json:"DatabaseOptions,omitempty"`

	// DatabaseOutputMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databaseoutput.html#cfn-databrew-job-databaseoutput-databaseoutputmode
	DatabaseOutputMode cfz.Expression[string] `json:"DatabaseOutputMode,omitempty"`

	// GlueConnectionName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databaseoutput.html#cfn-databrew-job-databaseoutput-glueconnectionname
	GlueConnectionName cfz.Expression[string] `json:"GlueConnectionName,omitempty"`
}

// New__AWS_DataBrew_Job_DatabaseOutput initializes a new AWS_DataBrew_Job_DatabaseOutput.
func New__AWS_DataBrew_Job_DatabaseOutput() AWS_DataBrew_Job_DatabaseOutput {
	return AWS_DataBrew_Job_DatabaseOutput{}
}

// GetType returns the CloudFormation type.
func (AWS_DataBrew_Job_DatabaseOutput) GetType() string {
	return AWS_DataBrew_Job_DatabaseOutput__Type
}

// Set__DatabaseOptions updates property "DatabaseOptions".
func (t AWS_DataBrew_Job_DatabaseOutput) Set__DatabaseOptions(v cfz.Expression[AWS_DataBrew_Job_DatabaseTableOutputOptions]) AWS_DataBrew_Job_DatabaseOutput {
	t.DatabaseOptions = v
	return t
}

// SetV__DatabaseOptions updates property "DatabaseOptions".
func (t AWS_DataBrew_Job_DatabaseOutput) SetV__DatabaseOptions(v AWS_DataBrew_Job_DatabaseTableOutputOptions) AWS_DataBrew_Job_DatabaseOutput {
	t.DatabaseOptions = cfz.V(v)
	return t
}

// Set__DatabaseOutputMode updates property "DatabaseOutputMode".
func (t AWS_DataBrew_Job_DatabaseOutput) Set__DatabaseOutputMode(v cfz.Expression[string]) AWS_DataBrew_Job_DatabaseOutput {
	t.DatabaseOutputMode = v
	return t
}

// SetV__DatabaseOutputMode updates property "DatabaseOutputMode".
func (t AWS_DataBrew_Job_DatabaseOutput) SetV__DatabaseOutputMode(v string) AWS_DataBrew_Job_DatabaseOutput {
	t.DatabaseOutputMode = cfz.V(v)
	return t
}

// Set__GlueConnectionName updates property "GlueConnectionName".
func (t AWS_DataBrew_Job_DatabaseOutput) Set__GlueConnectionName(v cfz.Expression[string]) AWS_DataBrew_Job_DatabaseOutput {
	t.GlueConnectionName = v
	return t
}

// SetV__GlueConnectionName updates property "GlueConnectionName".
func (t AWS_DataBrew_Job_DatabaseOutput) SetV__GlueConnectionName(v string) AWS_DataBrew_Job_DatabaseOutput {
	t.GlueConnectionName = cfz.V(v)
	return t
}
