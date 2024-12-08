// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codepipeline

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration__Type is the CloudFormation type for AWS::CodePipeline::Pipeline.PipelineTriggerDeclaration.
	AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration__Type = "AWS::CodePipeline::Pipeline.PipelineTriggerDeclaration"
)

var (
	// AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration__PropertiesMap reports all the CloudFormation properties for AWS::CodePipeline::Pipeline.PipelineTriggerDeclaration.
	AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration__PropertiesMap = struct {
		GitConfiguration string
		ProviderType     string
	}{
		GitConfiguration: "GitConfiguration",
		ProviderType:     "ProviderType",
	}

	// AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration__PropertiesSlice reports all the CloudFormation properties for AWS::CodePipeline::Pipeline.PipelineTriggerDeclaration.
	AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration__PropertiesSlice = []string{
		AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration__PropertiesMap.GitConfiguration,
		AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration__PropertiesMap.ProviderType,
	}
)

// AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration is a binding for AWS::CodePipeline::Pipeline.PipelineTriggerDeclaration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-pipelinetriggerdeclaration.html
type AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration struct {
	// GitConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-pipelinetriggerdeclaration.html#cfn-codepipeline-pipeline-pipelinetriggerdeclaration-gitconfiguration
	GitConfiguration cfz.Expression[AWS_CodePipeline_Pipeline_GitConfiguration] `json:"GitConfiguration,omitempty"`

	// ProviderType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-pipelinetriggerdeclaration.html#cfn-codepipeline-pipeline-pipelinetriggerdeclaration-providertype
	ProviderType cfz.Expression[string] `json:"ProviderType,omitempty"`
}

// New__AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration initializes a new AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration.
func New__AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration() AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration {
	return AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration{}
}

// GetType returns the CloudFormation type.
func (AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration) GetType() string {
	return AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration__Type
}

// Set__GitConfiguration updates property "GitConfiguration".
func (t AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration) Set__GitConfiguration(v cfz.Expression[AWS_CodePipeline_Pipeline_GitConfiguration]) AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration {
	t.GitConfiguration = v
	return t
}

// SetV__GitConfiguration updates property "GitConfiguration".
func (t AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration) SetV__GitConfiguration(v AWS_CodePipeline_Pipeline_GitConfiguration) AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration {
	t.GitConfiguration = cfz.V(v)
	return t
}

// Set__ProviderType updates property "ProviderType".
func (t AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration) Set__ProviderType(v cfz.Expression[string]) AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration {
	t.ProviderType = v
	return t
}

// SetV__ProviderType updates property "ProviderType".
func (t AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration) SetV__ProviderType(v string) AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration {
	t.ProviderType = cfz.V(v)
	return t
}
