// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_fis

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__Type is the CloudFormation type for AWS::FIS::ExperimentTemplate.ExperimentTemplateExperimentReportConfiguration.
	AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__Type = "AWS::FIS::ExperimentTemplate.ExperimentTemplateExperimentReportConfiguration"
)

var (
	// AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::FIS::ExperimentTemplate.ExperimentTemplateExperimentReportConfiguration.
	AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__PropertiesMap = struct {
		DataSources            string
		Outputs                string
		PostExperimentDuration string
		PreExperimentDuration  string
	}{
		DataSources:            "DataSources",
		Outputs:                "Outputs",
		PostExperimentDuration: "PostExperimentDuration",
		PreExperimentDuration:  "PreExperimentDuration",
	}

	// AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::FIS::ExperimentTemplate.ExperimentTemplateExperimentReportConfiguration.
	AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__PropertiesSlice = []string{
		AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__PropertiesMap.DataSources,
		AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__PropertiesMap.Outputs,
		AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__PropertiesMap.PostExperimentDuration,
		AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__PropertiesMap.PreExperimentDuration,
	}
)

// AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration is a binding for AWS::FIS::ExperimentTemplate.ExperimentTemplateExperimentReportConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html
type AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration struct {
	// DataSources is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html#cfn-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration-datasources
	DataSources cfz.Expression[AWS_FIS_ExperimentTemplate_DataSources] `json:"DataSources,omitempty"`

	// Outputs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html#cfn-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration-outputs
	Outputs cfz.Expression[AWS_FIS_ExperimentTemplate_Outputs] `json:"Outputs,omitempty"`

	// PostExperimentDuration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html#cfn-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration-postexperimentduration
	PostExperimentDuration cfz.Expression[string] `json:"PostExperimentDuration,omitempty"`

	// PreExperimentDuration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html#cfn-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration-preexperimentduration
	PreExperimentDuration cfz.Expression[string] `json:"PreExperimentDuration,omitempty"`
}

// New__AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration initializes a new AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration.
func New__AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration() AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration {
	return AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration) GetType() string {
	return AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration__Type
}

// Set__DataSources updates property "DataSources".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration) Set__DataSources(v cfz.Expression[AWS_FIS_ExperimentTemplate_DataSources]) AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration {
	t.DataSources = v
	return t
}

// SetV__DataSources updates property "DataSources".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration) SetV__DataSources(v AWS_FIS_ExperimentTemplate_DataSources) AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration {
	t.DataSources = cfz.V(v)
	return t
}

// Set__Outputs updates property "Outputs".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration) Set__Outputs(v cfz.Expression[AWS_FIS_ExperimentTemplate_Outputs]) AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration {
	t.Outputs = v
	return t
}

// SetV__Outputs updates property "Outputs".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration) SetV__Outputs(v AWS_FIS_ExperimentTemplate_Outputs) AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration {
	t.Outputs = cfz.V(v)
	return t
}

// Set__PostExperimentDuration updates property "PostExperimentDuration".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration) Set__PostExperimentDuration(v cfz.Expression[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration {
	t.PostExperimentDuration = v
	return t
}

// SetV__PostExperimentDuration updates property "PostExperimentDuration".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration) SetV__PostExperimentDuration(v string) AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration {
	t.PostExperimentDuration = cfz.V(v)
	return t
}

// Set__PreExperimentDuration updates property "PreExperimentDuration".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration) Set__PreExperimentDuration(v cfz.Expression[string]) AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration {
	t.PreExperimentDuration = v
	return t
}

// SetV__PreExperimentDuration updates property "PreExperimentDuration".
func (t AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration) SetV__PreExperimentDuration(v string) AWS_FIS_ExperimentTemplate_ExperimentTemplateExperimentReportConfiguration {
	t.PreExperimentDuration = cfz.V(v)
	return t
}
