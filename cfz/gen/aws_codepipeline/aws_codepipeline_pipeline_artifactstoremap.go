// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codepipeline

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodePipeline_Pipeline_ArtifactStoreMap__Type is the CloudFormation type for AWS::CodePipeline::Pipeline.ArtifactStoreMap.
	AWS_CodePipeline_Pipeline_ArtifactStoreMap__Type = "AWS::CodePipeline::Pipeline.ArtifactStoreMap"
)

var (
	// AWS_CodePipeline_Pipeline_ArtifactStoreMap__PropertiesMap reports all the CloudFormation properties for AWS::CodePipeline::Pipeline.ArtifactStoreMap.
	AWS_CodePipeline_Pipeline_ArtifactStoreMap__PropertiesMap = struct {
		ArtifactStore string
		Region        string
	}{
		ArtifactStore: "ArtifactStore",
		Region:        "Region",
	}

	// AWS_CodePipeline_Pipeline_ArtifactStoreMap__PropertiesSlice reports all the CloudFormation properties for AWS::CodePipeline::Pipeline.ArtifactStoreMap.
	AWS_CodePipeline_Pipeline_ArtifactStoreMap__PropertiesSlice = []string{
		AWS_CodePipeline_Pipeline_ArtifactStoreMap__PropertiesMap.ArtifactStore,
		AWS_CodePipeline_Pipeline_ArtifactStoreMap__PropertiesMap.Region,
	}
)

// AWS_CodePipeline_Pipeline_ArtifactStoreMap is a binding for AWS::CodePipeline::Pipeline.ArtifactStoreMap.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html
type AWS_CodePipeline_Pipeline_ArtifactStoreMap struct {
	// ArtifactStore is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html#cfn-codepipeline-pipeline-artifactstoremap-artifactstore
	ArtifactStore cfz.Expression[AWS_CodePipeline_Pipeline_ArtifactStore] `json:"ArtifactStore,omitempty"`

	// Region is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html#cfn-codepipeline-pipeline-artifactstoremap-region
	Region cfz.Expression[string] `json:"Region,omitempty"`
}

// New__AWS_CodePipeline_Pipeline_ArtifactStoreMap initializes a new AWS_CodePipeline_Pipeline_ArtifactStoreMap.
func New__AWS_CodePipeline_Pipeline_ArtifactStoreMap() AWS_CodePipeline_Pipeline_ArtifactStoreMap {
	return AWS_CodePipeline_Pipeline_ArtifactStoreMap{}
}

// GetType returns the CloudFormation type.
func (AWS_CodePipeline_Pipeline_ArtifactStoreMap) GetType() string {
	return AWS_CodePipeline_Pipeline_ArtifactStoreMap__Type
}

// Set__ArtifactStore updates property "ArtifactStore".
func (t AWS_CodePipeline_Pipeline_ArtifactStoreMap) Set__ArtifactStore(v cfz.Expression[AWS_CodePipeline_Pipeline_ArtifactStore]) AWS_CodePipeline_Pipeline_ArtifactStoreMap {
	t.ArtifactStore = v
	return t
}

// SetV__ArtifactStore updates property "ArtifactStore".
func (t AWS_CodePipeline_Pipeline_ArtifactStoreMap) SetV__ArtifactStore(v AWS_CodePipeline_Pipeline_ArtifactStore) AWS_CodePipeline_Pipeline_ArtifactStoreMap {
	t.ArtifactStore = cfz.V(v)
	return t
}

// Set__Region updates property "Region".
func (t AWS_CodePipeline_Pipeline_ArtifactStoreMap) Set__Region(v cfz.Expression[string]) AWS_CodePipeline_Pipeline_ArtifactStoreMap {
	t.Region = v
	return t
}

// SetV__Region updates property "Region".
func (t AWS_CodePipeline_Pipeline_ArtifactStoreMap) SetV__Region(v string) AWS_CodePipeline_Pipeline_ArtifactStoreMap {
	t.Region = cfz.V(v)
	return t
}
