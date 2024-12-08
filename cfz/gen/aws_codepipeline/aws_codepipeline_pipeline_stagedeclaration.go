// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codepipeline

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodePipeline_Pipeline_StageDeclaration__Type is the CloudFormation type for AWS::CodePipeline::Pipeline.StageDeclaration.
	AWS_CodePipeline_Pipeline_StageDeclaration__Type = "AWS::CodePipeline::Pipeline.StageDeclaration"
)

var (
	// AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesMap reports all the CloudFormation properties for AWS::CodePipeline::Pipeline.StageDeclaration.
	AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesMap = struct {
		Actions     string
		BeforeEntry string
		Blockers    string
		Name        string
		OnFailure   string
		OnSuccess   string
	}{
		Actions:     "Actions",
		BeforeEntry: "BeforeEntry",
		Blockers:    "Blockers",
		Name:        "Name",
		OnFailure:   "OnFailure",
		OnSuccess:   "OnSuccess",
	}

	// AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesSlice reports all the CloudFormation properties for AWS::CodePipeline::Pipeline.StageDeclaration.
	AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesSlice = []string{
		AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesMap.Actions,
		AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesMap.BeforeEntry,
		AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesMap.Blockers,
		AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesMap.Name,
		AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesMap.OnFailure,
		AWS_CodePipeline_Pipeline_StageDeclaration__PropertiesMap.OnSuccess,
	}
)

// AWS_CodePipeline_Pipeline_StageDeclaration is a binding for AWS::CodePipeline::Pipeline.StageDeclaration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html
type AWS_CodePipeline_Pipeline_StageDeclaration struct {
	// Actions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-actions
	Actions cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_ActionDeclaration] `json:"Actions,omitempty"`

	// BeforeEntry is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-beforeentry
	BeforeEntry cfz.Expression[AWS_CodePipeline_Pipeline_BeforeEntryConditions] `json:"BeforeEntry,omitempty"`

	// Blockers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-blockers
	Blockers cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_BlockerDeclaration] `json:"Blockers,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// OnFailure is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-onfailure
	OnFailure cfz.Expression[AWS_CodePipeline_Pipeline_FailureConditions] `json:"OnFailure,omitempty"`

	// OnSuccess is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-onsuccess
	OnSuccess cfz.Expression[AWS_CodePipeline_Pipeline_SuccessConditions] `json:"OnSuccess,omitempty"`
}

// New__AWS_CodePipeline_Pipeline_StageDeclaration initializes a new AWS_CodePipeline_Pipeline_StageDeclaration.
func New__AWS_CodePipeline_Pipeline_StageDeclaration() AWS_CodePipeline_Pipeline_StageDeclaration {
	return AWS_CodePipeline_Pipeline_StageDeclaration{}
}

// GetType returns the CloudFormation type.
func (AWS_CodePipeline_Pipeline_StageDeclaration) GetType() string {
	return AWS_CodePipeline_Pipeline_StageDeclaration__Type
}

// Set__Actions updates property "Actions".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) Set__Actions(v cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_ActionDeclaration]) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.Actions = v
	return t
}

// SetS__Actions updates property "Actions".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) SetS__Actions(v ...cfz.Expression[AWS_CodePipeline_Pipeline_ActionDeclaration]) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.Actions = cfz.S(v...)
	return t
}

// SetSV__Actions updates property "Actions".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) SetSV__Actions(v ...AWS_CodePipeline_Pipeline_ActionDeclaration) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.Actions = cfz.SV(v...)
	return t
}

// Set__BeforeEntry updates property "BeforeEntry".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) Set__BeforeEntry(v cfz.Expression[AWS_CodePipeline_Pipeline_BeforeEntryConditions]) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.BeforeEntry = v
	return t
}

// SetV__BeforeEntry updates property "BeforeEntry".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) SetV__BeforeEntry(v AWS_CodePipeline_Pipeline_BeforeEntryConditions) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.BeforeEntry = cfz.V(v)
	return t
}

// Set__Blockers updates property "Blockers".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) Set__Blockers(v cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_BlockerDeclaration]) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.Blockers = v
	return t
}

// SetS__Blockers updates property "Blockers".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) SetS__Blockers(v ...cfz.Expression[AWS_CodePipeline_Pipeline_BlockerDeclaration]) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.Blockers = cfz.S(v...)
	return t
}

// SetSV__Blockers updates property "Blockers".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) SetSV__Blockers(v ...AWS_CodePipeline_Pipeline_BlockerDeclaration) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.Blockers = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) Set__Name(v cfz.Expression[string]) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) SetV__Name(v string) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.Name = cfz.V(v)
	return t
}

// Set__OnFailure updates property "OnFailure".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) Set__OnFailure(v cfz.Expression[AWS_CodePipeline_Pipeline_FailureConditions]) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.OnFailure = v
	return t
}

// SetV__OnFailure updates property "OnFailure".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) SetV__OnFailure(v AWS_CodePipeline_Pipeline_FailureConditions) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.OnFailure = cfz.V(v)
	return t
}

// Set__OnSuccess updates property "OnSuccess".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) Set__OnSuccess(v cfz.Expression[AWS_CodePipeline_Pipeline_SuccessConditions]) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.OnSuccess = v
	return t
}

// SetV__OnSuccess updates property "OnSuccess".
func (t AWS_CodePipeline_Pipeline_StageDeclaration) SetV__OnSuccess(v AWS_CodePipeline_Pipeline_SuccessConditions) AWS_CodePipeline_Pipeline_StageDeclaration {
	t.OnSuccess = cfz.V(v)
	return t
}
