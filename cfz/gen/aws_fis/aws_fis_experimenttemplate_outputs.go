// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_fis

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FIS_ExperimentTemplate_Outputs__Type is the CloudFormation type for AWS::FIS::ExperimentTemplate.Outputs.
	AWS_FIS_ExperimentTemplate_Outputs__Type = "AWS::FIS::ExperimentTemplate.Outputs"
)

var (
	// AWS_FIS_ExperimentTemplate_Outputs__PropertiesMap reports all the CloudFormation properties for AWS::FIS::ExperimentTemplate.Outputs.
	AWS_FIS_ExperimentTemplate_Outputs__PropertiesMap = struct {
		ExperimentReportS3Configuration string
	}{
		ExperimentReportS3Configuration: "ExperimentReportS3Configuration",
	}

	// AWS_FIS_ExperimentTemplate_Outputs__PropertiesSlice reports all the CloudFormation properties for AWS::FIS::ExperimentTemplate.Outputs.
	AWS_FIS_ExperimentTemplate_Outputs__PropertiesSlice = []string{
		AWS_FIS_ExperimentTemplate_Outputs__PropertiesMap.ExperimentReportS3Configuration,
	}
)

// AWS_FIS_ExperimentTemplate_Outputs is a binding for AWS::FIS::ExperimentTemplate.Outputs.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-outputs.html
type AWS_FIS_ExperimentTemplate_Outputs struct {
	// ExperimentReportS3Configuration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-outputs.html#cfn-fis-experimenttemplate-outputs-experimentreports3configuration
	ExperimentReportS3Configuration cfz.Expression[AWS_FIS_ExperimentTemplate_ExperimentReportS3Configuration] `json:"ExperimentReportS3Configuration,omitempty"`
}

// New__AWS_FIS_ExperimentTemplate_Outputs initializes a new AWS_FIS_ExperimentTemplate_Outputs.
func New__AWS_FIS_ExperimentTemplate_Outputs() AWS_FIS_ExperimentTemplate_Outputs {
	return AWS_FIS_ExperimentTemplate_Outputs{}
}

// GetType returns the CloudFormation type.
func (AWS_FIS_ExperimentTemplate_Outputs) GetType() string {
	return AWS_FIS_ExperimentTemplate_Outputs__Type
}

// Set__ExperimentReportS3Configuration updates property "ExperimentReportS3Configuration".
func (t AWS_FIS_ExperimentTemplate_Outputs) Set__ExperimentReportS3Configuration(v cfz.Expression[AWS_FIS_ExperimentTemplate_ExperimentReportS3Configuration]) AWS_FIS_ExperimentTemplate_Outputs {
	t.ExperimentReportS3Configuration = v
	return t
}

// SetV__ExperimentReportS3Configuration updates property "ExperimentReportS3Configuration".
func (t AWS_FIS_ExperimentTemplate_Outputs) SetV__ExperimentReportS3Configuration(v AWS_FIS_ExperimentTemplate_ExperimentReportS3Configuration) AWS_FIS_ExperimentTemplate_Outputs {
	t.ExperimentReportS3Configuration = cfz.V(v)
	return t
}
