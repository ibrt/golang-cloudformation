// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_CodeRepository_GitConfig__Type is the CloudFormation type for AWS::SageMaker::CodeRepository.GitConfig.
	AWS_SageMaker_CodeRepository_GitConfig__Type = "AWS::SageMaker::CodeRepository.GitConfig"
)

var (
	// AWS_SageMaker_CodeRepository_GitConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::CodeRepository.GitConfig.
	AWS_SageMaker_CodeRepository_GitConfig__PropertiesMap = struct {
		Branch        string
		RepositoryUrl string
		SecretArn     string
	}{
		Branch:        "Branch",
		RepositoryUrl: "RepositoryUrl",
		SecretArn:     "SecretArn",
	}

	// AWS_SageMaker_CodeRepository_GitConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::CodeRepository.GitConfig.
	AWS_SageMaker_CodeRepository_GitConfig__PropertiesSlice = []string{
		AWS_SageMaker_CodeRepository_GitConfig__PropertiesMap.Branch,
		AWS_SageMaker_CodeRepository_GitConfig__PropertiesMap.RepositoryUrl,
		AWS_SageMaker_CodeRepository_GitConfig__PropertiesMap.SecretArn,
	}
)

// AWS_SageMaker_CodeRepository_GitConfig is a binding for AWS::SageMaker::CodeRepository.GitConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-coderepository-gitconfig.html
type AWS_SageMaker_CodeRepository_GitConfig struct {
	// Branch is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-coderepository-gitconfig.html#cfn-sagemaker-coderepository-gitconfig-branch
	Branch cfz.Expression[string] `json:"Branch,omitempty"`

	// RepositoryUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-coderepository-gitconfig.html#cfn-sagemaker-coderepository-gitconfig-repositoryurl
	RepositoryUrl cfz.Expression[string] `json:"RepositoryUrl,omitempty"`

	// SecretArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-coderepository-gitconfig.html#cfn-sagemaker-coderepository-gitconfig-secretarn
	SecretArn cfz.Expression[string] `json:"SecretArn,omitempty"`
}

// New__AWS_SageMaker_CodeRepository_GitConfig initializes a new AWS_SageMaker_CodeRepository_GitConfig.
func New__AWS_SageMaker_CodeRepository_GitConfig() AWS_SageMaker_CodeRepository_GitConfig {
	return AWS_SageMaker_CodeRepository_GitConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_CodeRepository_GitConfig) GetType() string {
	return AWS_SageMaker_CodeRepository_GitConfig__Type
}

// Set__Branch updates property "Branch".
func (t AWS_SageMaker_CodeRepository_GitConfig) Set__Branch(v cfz.Expression[string]) AWS_SageMaker_CodeRepository_GitConfig {
	t.Branch = v
	return t
}

// SetV__Branch updates property "Branch".
func (t AWS_SageMaker_CodeRepository_GitConfig) SetV__Branch(v string) AWS_SageMaker_CodeRepository_GitConfig {
	t.Branch = cfz.V(v)
	return t
}

// Set__RepositoryUrl updates property "RepositoryUrl".
func (t AWS_SageMaker_CodeRepository_GitConfig) Set__RepositoryUrl(v cfz.Expression[string]) AWS_SageMaker_CodeRepository_GitConfig {
	t.RepositoryUrl = v
	return t
}

// SetV__RepositoryUrl updates property "RepositoryUrl".
func (t AWS_SageMaker_CodeRepository_GitConfig) SetV__RepositoryUrl(v string) AWS_SageMaker_CodeRepository_GitConfig {
	t.RepositoryUrl = cfz.V(v)
	return t
}

// Set__SecretArn updates property "SecretArn".
func (t AWS_SageMaker_CodeRepository_GitConfig) Set__SecretArn(v cfz.Expression[string]) AWS_SageMaker_CodeRepository_GitConfig {
	t.SecretArn = v
	return t
}

// SetV__SecretArn updates property "SecretArn".
func (t AWS_SageMaker_CodeRepository_GitConfig) SetV__SecretArn(v string) AWS_SageMaker_CodeRepository_GitConfig {
	t.SecretArn = cfz.V(v)
	return t
}
