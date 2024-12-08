// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_databrew

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataBrew_Job_ProfileConfiguration__Type is the CloudFormation type for AWS::DataBrew::Job.ProfileConfiguration.
	AWS_DataBrew_Job_ProfileConfiguration__Type = "AWS::DataBrew::Job.ProfileConfiguration"
)

var (
	// AWS_DataBrew_Job_ProfileConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::DataBrew::Job.ProfileConfiguration.
	AWS_DataBrew_Job_ProfileConfiguration__PropertiesMap = struct {
		ColumnStatisticsConfigurations string
		DatasetStatisticsConfiguration string
		EntityDetectorConfiguration    string
		ProfileColumns                 string
	}{
		ColumnStatisticsConfigurations: "ColumnStatisticsConfigurations",
		DatasetStatisticsConfiguration: "DatasetStatisticsConfiguration",
		EntityDetectorConfiguration:    "EntityDetectorConfiguration",
		ProfileColumns:                 "ProfileColumns",
	}

	// AWS_DataBrew_Job_ProfileConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::DataBrew::Job.ProfileConfiguration.
	AWS_DataBrew_Job_ProfileConfiguration__PropertiesSlice = []string{
		AWS_DataBrew_Job_ProfileConfiguration__PropertiesMap.ColumnStatisticsConfigurations,
		AWS_DataBrew_Job_ProfileConfiguration__PropertiesMap.DatasetStatisticsConfiguration,
		AWS_DataBrew_Job_ProfileConfiguration__PropertiesMap.EntityDetectorConfiguration,
		AWS_DataBrew_Job_ProfileConfiguration__PropertiesMap.ProfileColumns,
	}
)

// AWS_DataBrew_Job_ProfileConfiguration is a binding for AWS::DataBrew::Job.ProfileConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html
type AWS_DataBrew_Job_ProfileConfiguration struct {
	// ColumnStatisticsConfigurations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-columnstatisticsconfigurations
	ColumnStatisticsConfigurations cfz.ExpressionSlice[AWS_DataBrew_Job_ColumnStatisticsConfiguration] `json:"ColumnStatisticsConfigurations,omitempty"`

	// DatasetStatisticsConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-datasetstatisticsconfiguration
	DatasetStatisticsConfiguration cfz.Expression[AWS_DataBrew_Job_StatisticsConfiguration] `json:"DatasetStatisticsConfiguration,omitempty"`

	// EntityDetectorConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-entitydetectorconfiguration
	EntityDetectorConfiguration cfz.Expression[AWS_DataBrew_Job_EntityDetectorConfiguration] `json:"EntityDetectorConfiguration,omitempty"`

	// ProfileColumns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-profilecolumns
	ProfileColumns cfz.ExpressionSlice[AWS_DataBrew_Job_ColumnSelector] `json:"ProfileColumns,omitempty"`
}

// New__AWS_DataBrew_Job_ProfileConfiguration initializes a new AWS_DataBrew_Job_ProfileConfiguration.
func New__AWS_DataBrew_Job_ProfileConfiguration() AWS_DataBrew_Job_ProfileConfiguration {
	return AWS_DataBrew_Job_ProfileConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_DataBrew_Job_ProfileConfiguration) GetType() string {
	return AWS_DataBrew_Job_ProfileConfiguration__Type
}

// Set__ColumnStatisticsConfigurations updates property "ColumnStatisticsConfigurations".
func (t AWS_DataBrew_Job_ProfileConfiguration) Set__ColumnStatisticsConfigurations(v cfz.ExpressionSlice[AWS_DataBrew_Job_ColumnStatisticsConfiguration]) AWS_DataBrew_Job_ProfileConfiguration {
	t.ColumnStatisticsConfigurations = v
	return t
}

// SetS__ColumnStatisticsConfigurations updates property "ColumnStatisticsConfigurations".
func (t AWS_DataBrew_Job_ProfileConfiguration) SetS__ColumnStatisticsConfigurations(v ...cfz.Expression[AWS_DataBrew_Job_ColumnStatisticsConfiguration]) AWS_DataBrew_Job_ProfileConfiguration {
	t.ColumnStatisticsConfigurations = cfz.S(v...)
	return t
}

// SetSV__ColumnStatisticsConfigurations updates property "ColumnStatisticsConfigurations".
func (t AWS_DataBrew_Job_ProfileConfiguration) SetSV__ColumnStatisticsConfigurations(v ...AWS_DataBrew_Job_ColumnStatisticsConfiguration) AWS_DataBrew_Job_ProfileConfiguration {
	t.ColumnStatisticsConfigurations = cfz.SV(v...)
	return t
}

// Set__DatasetStatisticsConfiguration updates property "DatasetStatisticsConfiguration".
func (t AWS_DataBrew_Job_ProfileConfiguration) Set__DatasetStatisticsConfiguration(v cfz.Expression[AWS_DataBrew_Job_StatisticsConfiguration]) AWS_DataBrew_Job_ProfileConfiguration {
	t.DatasetStatisticsConfiguration = v
	return t
}

// SetV__DatasetStatisticsConfiguration updates property "DatasetStatisticsConfiguration".
func (t AWS_DataBrew_Job_ProfileConfiguration) SetV__DatasetStatisticsConfiguration(v AWS_DataBrew_Job_StatisticsConfiguration) AWS_DataBrew_Job_ProfileConfiguration {
	t.DatasetStatisticsConfiguration = cfz.V(v)
	return t
}

// Set__EntityDetectorConfiguration updates property "EntityDetectorConfiguration".
func (t AWS_DataBrew_Job_ProfileConfiguration) Set__EntityDetectorConfiguration(v cfz.Expression[AWS_DataBrew_Job_EntityDetectorConfiguration]) AWS_DataBrew_Job_ProfileConfiguration {
	t.EntityDetectorConfiguration = v
	return t
}

// SetV__EntityDetectorConfiguration updates property "EntityDetectorConfiguration".
func (t AWS_DataBrew_Job_ProfileConfiguration) SetV__EntityDetectorConfiguration(v AWS_DataBrew_Job_EntityDetectorConfiguration) AWS_DataBrew_Job_ProfileConfiguration {
	t.EntityDetectorConfiguration = cfz.V(v)
	return t
}

// Set__ProfileColumns updates property "ProfileColumns".
func (t AWS_DataBrew_Job_ProfileConfiguration) Set__ProfileColumns(v cfz.ExpressionSlice[AWS_DataBrew_Job_ColumnSelector]) AWS_DataBrew_Job_ProfileConfiguration {
	t.ProfileColumns = v
	return t
}

// SetS__ProfileColumns updates property "ProfileColumns".
func (t AWS_DataBrew_Job_ProfileConfiguration) SetS__ProfileColumns(v ...cfz.Expression[AWS_DataBrew_Job_ColumnSelector]) AWS_DataBrew_Job_ProfileConfiguration {
	t.ProfileColumns = cfz.S(v...)
	return t
}

// SetSV__ProfileColumns updates property "ProfileColumns".
func (t AWS_DataBrew_Job_ProfileConfiguration) SetSV__ProfileColumns(v ...AWS_DataBrew_Job_ColumnSelector) AWS_DataBrew_Job_ProfileConfiguration {
	t.ProfileColumns = cfz.SV(v...)
	return t
}
