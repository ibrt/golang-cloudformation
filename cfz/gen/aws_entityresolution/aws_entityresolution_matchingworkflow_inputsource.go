// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_entityresolution

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EntityResolution_MatchingWorkflow_InputSource__Type is the CloudFormation type for AWS::EntityResolution::MatchingWorkflow.InputSource.
	AWS_EntityResolution_MatchingWorkflow_InputSource__Type = "AWS::EntityResolution::MatchingWorkflow.InputSource"
)

var (
	// AWS_EntityResolution_MatchingWorkflow_InputSource__PropertiesMap reports all the CloudFormation properties for AWS::EntityResolution::MatchingWorkflow.InputSource.
	AWS_EntityResolution_MatchingWorkflow_InputSource__PropertiesMap = struct {
		ApplyNormalization string
		InputSourceARN     string
		SchemaArn          string
	}{
		ApplyNormalization: "ApplyNormalization",
		InputSourceARN:     "InputSourceARN",
		SchemaArn:          "SchemaArn",
	}

	// AWS_EntityResolution_MatchingWorkflow_InputSource__PropertiesSlice reports all the CloudFormation properties for AWS::EntityResolution::MatchingWorkflow.InputSource.
	AWS_EntityResolution_MatchingWorkflow_InputSource__PropertiesSlice = []string{
		AWS_EntityResolution_MatchingWorkflow_InputSource__PropertiesMap.ApplyNormalization,
		AWS_EntityResolution_MatchingWorkflow_InputSource__PropertiesMap.InputSourceARN,
		AWS_EntityResolution_MatchingWorkflow_InputSource__PropertiesMap.SchemaArn,
	}
)

// AWS_EntityResolution_MatchingWorkflow_InputSource is a binding for AWS::EntityResolution::MatchingWorkflow.InputSource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html
type AWS_EntityResolution_MatchingWorkflow_InputSource struct {
	// ApplyNormalization is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-applynormalization
	ApplyNormalization cfz.Expression[bool] `json:"ApplyNormalization,omitempty"`

	// InputSourceARN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-inputsourcearn
	InputSourceARN cfz.Expression[string] `json:"InputSourceARN,omitempty"`

	// SchemaArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-schemaarn
	SchemaArn cfz.Expression[string] `json:"SchemaArn,omitempty"`
}

// New__AWS_EntityResolution_MatchingWorkflow_InputSource initializes a new AWS_EntityResolution_MatchingWorkflow_InputSource.
func New__AWS_EntityResolution_MatchingWorkflow_InputSource() AWS_EntityResolution_MatchingWorkflow_InputSource {
	return AWS_EntityResolution_MatchingWorkflow_InputSource{}
}

// GetType returns the CloudFormation type.
func (AWS_EntityResolution_MatchingWorkflow_InputSource) GetType() string {
	return AWS_EntityResolution_MatchingWorkflow_InputSource__Type
}

// Set__ApplyNormalization updates property "ApplyNormalization".
func (t AWS_EntityResolution_MatchingWorkflow_InputSource) Set__ApplyNormalization(v cfz.Expression[bool]) AWS_EntityResolution_MatchingWorkflow_InputSource {
	t.ApplyNormalization = v
	return t
}

// SetV__ApplyNormalization updates property "ApplyNormalization".
func (t AWS_EntityResolution_MatchingWorkflow_InputSource) SetV__ApplyNormalization(v bool) AWS_EntityResolution_MatchingWorkflow_InputSource {
	t.ApplyNormalization = cfz.V(v)
	return t
}

// Set__InputSourceARN updates property "InputSourceARN".
func (t AWS_EntityResolution_MatchingWorkflow_InputSource) Set__InputSourceARN(v cfz.Expression[string]) AWS_EntityResolution_MatchingWorkflow_InputSource {
	t.InputSourceARN = v
	return t
}

// SetV__InputSourceARN updates property "InputSourceARN".
func (t AWS_EntityResolution_MatchingWorkflow_InputSource) SetV__InputSourceARN(v string) AWS_EntityResolution_MatchingWorkflow_InputSource {
	t.InputSourceARN = cfz.V(v)
	return t
}

// Set__SchemaArn updates property "SchemaArn".
func (t AWS_EntityResolution_MatchingWorkflow_InputSource) Set__SchemaArn(v cfz.Expression[string]) AWS_EntityResolution_MatchingWorkflow_InputSource {
	t.SchemaArn = v
	return t
}

// SetV__SchemaArn updates property "SchemaArn".
func (t AWS_EntityResolution_MatchingWorkflow_InputSource) SetV__SchemaArn(v string) AWS_EntityResolution_MatchingWorkflow_InputSource {
	t.SchemaArn = cfz.V(v)
	return t
}
