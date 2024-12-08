// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_databrew

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataBrew_Job_EntityDetectorConfiguration__Type is the CloudFormation type for AWS::DataBrew::Job.EntityDetectorConfiguration.
	AWS_DataBrew_Job_EntityDetectorConfiguration__Type = "AWS::DataBrew::Job.EntityDetectorConfiguration"
)

var (
	// AWS_DataBrew_Job_EntityDetectorConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::DataBrew::Job.EntityDetectorConfiguration.
	AWS_DataBrew_Job_EntityDetectorConfiguration__PropertiesMap = struct {
		AllowedStatistics string
		EntityTypes       string
	}{
		AllowedStatistics: "AllowedStatistics",
		EntityTypes:       "EntityTypes",
	}

	// AWS_DataBrew_Job_EntityDetectorConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::DataBrew::Job.EntityDetectorConfiguration.
	AWS_DataBrew_Job_EntityDetectorConfiguration__PropertiesSlice = []string{
		AWS_DataBrew_Job_EntityDetectorConfiguration__PropertiesMap.AllowedStatistics,
		AWS_DataBrew_Job_EntityDetectorConfiguration__PropertiesMap.EntityTypes,
	}
)

// AWS_DataBrew_Job_EntityDetectorConfiguration is a binding for AWS::DataBrew::Job.EntityDetectorConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-entitydetectorconfiguration.html
type AWS_DataBrew_Job_EntityDetectorConfiguration struct {
	// AllowedStatistics is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-entitydetectorconfiguration.html#cfn-databrew-job-entitydetectorconfiguration-allowedstatistics
	AllowedStatistics cfz.Expression[AWS_DataBrew_Job_AllowedStatistics] `json:"AllowedStatistics,omitempty"`

	// EntityTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-entitydetectorconfiguration.html#cfn-databrew-job-entitydetectorconfiguration-entitytypes
	EntityTypes cfz.ExpressionSlice[string] `json:"EntityTypes,omitempty"`
}

// New__AWS_DataBrew_Job_EntityDetectorConfiguration initializes a new AWS_DataBrew_Job_EntityDetectorConfiguration.
func New__AWS_DataBrew_Job_EntityDetectorConfiguration() AWS_DataBrew_Job_EntityDetectorConfiguration {
	return AWS_DataBrew_Job_EntityDetectorConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_DataBrew_Job_EntityDetectorConfiguration) GetType() string {
	return AWS_DataBrew_Job_EntityDetectorConfiguration__Type
}

// Set__AllowedStatistics updates property "AllowedStatistics".
func (t AWS_DataBrew_Job_EntityDetectorConfiguration) Set__AllowedStatistics(v cfz.Expression[AWS_DataBrew_Job_AllowedStatistics]) AWS_DataBrew_Job_EntityDetectorConfiguration {
	t.AllowedStatistics = v
	return t
}

// SetV__AllowedStatistics updates property "AllowedStatistics".
func (t AWS_DataBrew_Job_EntityDetectorConfiguration) SetV__AllowedStatistics(v AWS_DataBrew_Job_AllowedStatistics) AWS_DataBrew_Job_EntityDetectorConfiguration {
	t.AllowedStatistics = cfz.V(v)
	return t
}

// Set__EntityTypes updates property "EntityTypes".
func (t AWS_DataBrew_Job_EntityDetectorConfiguration) Set__EntityTypes(v cfz.ExpressionSlice[string]) AWS_DataBrew_Job_EntityDetectorConfiguration {
	t.EntityTypes = v
	return t
}

// SetS__EntityTypes updates property "EntityTypes".
func (t AWS_DataBrew_Job_EntityDetectorConfiguration) SetS__EntityTypes(v ...cfz.Expression[string]) AWS_DataBrew_Job_EntityDetectorConfiguration {
	t.EntityTypes = cfz.S(v...)
	return t
}

// SetSV__EntityTypes updates property "EntityTypes".
func (t AWS_DataBrew_Job_EntityDetectorConfiguration) SetSV__EntityTypes(v ...string) AWS_DataBrew_Job_EntityDetectorConfiguration {
	t.EntityTypes = cfz.SV(v...)
	return t
}
