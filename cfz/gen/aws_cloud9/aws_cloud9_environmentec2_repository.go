// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloud9

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Cloud9_EnvironmentEC2_Repository__Type is the CloudFormation type for AWS::Cloud9::EnvironmentEC2.Repository.
	AWS_Cloud9_EnvironmentEC2_Repository__Type = "AWS::Cloud9::EnvironmentEC2.Repository"
)

var (
	// AWS_Cloud9_EnvironmentEC2_Repository__PropertiesMap reports all the CloudFormation properties for AWS::Cloud9::EnvironmentEC2.Repository.
	AWS_Cloud9_EnvironmentEC2_Repository__PropertiesMap = struct {
		PathComponent string
		RepositoryUrl string
	}{
		PathComponent: "PathComponent",
		RepositoryUrl: "RepositoryUrl",
	}

	// AWS_Cloud9_EnvironmentEC2_Repository__PropertiesSlice reports all the CloudFormation properties for AWS::Cloud9::EnvironmentEC2.Repository.
	AWS_Cloud9_EnvironmentEC2_Repository__PropertiesSlice = []string{
		AWS_Cloud9_EnvironmentEC2_Repository__PropertiesMap.PathComponent,
		AWS_Cloud9_EnvironmentEC2_Repository__PropertiesMap.RepositoryUrl,
	}
)

// AWS_Cloud9_EnvironmentEC2_Repository is a binding for AWS::Cloud9::EnvironmentEC2.Repository.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloud9-environmentec2-repository.html
type AWS_Cloud9_EnvironmentEC2_Repository struct {
	// PathComponent is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloud9-environmentec2-repository.html#cfn-cloud9-environmentec2-repository-pathcomponent
	PathComponent cfz.Expression[string] `json:"PathComponent,omitempty"`

	// RepositoryUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloud9-environmentec2-repository.html#cfn-cloud9-environmentec2-repository-repositoryurl
	RepositoryUrl cfz.Expression[string] `json:"RepositoryUrl,omitempty"`
}

// New__AWS_Cloud9_EnvironmentEC2_Repository initializes a new AWS_Cloud9_EnvironmentEC2_Repository.
func New__AWS_Cloud9_EnvironmentEC2_Repository() AWS_Cloud9_EnvironmentEC2_Repository {
	return AWS_Cloud9_EnvironmentEC2_Repository{}
}

// GetType returns the CloudFormation type.
func (AWS_Cloud9_EnvironmentEC2_Repository) GetType() string {
	return AWS_Cloud9_EnvironmentEC2_Repository__Type
}

// Set__PathComponent updates property "PathComponent".
func (t AWS_Cloud9_EnvironmentEC2_Repository) Set__PathComponent(v cfz.Expression[string]) AWS_Cloud9_EnvironmentEC2_Repository {
	t.PathComponent = v
	return t
}

// SetV__PathComponent updates property "PathComponent".
func (t AWS_Cloud9_EnvironmentEC2_Repository) SetV__PathComponent(v string) AWS_Cloud9_EnvironmentEC2_Repository {
	t.PathComponent = cfz.V(v)
	return t
}

// Set__RepositoryUrl updates property "RepositoryUrl".
func (t AWS_Cloud9_EnvironmentEC2_Repository) Set__RepositoryUrl(v cfz.Expression[string]) AWS_Cloud9_EnvironmentEC2_Repository {
	t.RepositoryUrl = v
	return t
}

// SetV__RepositoryUrl updates property "RepositoryUrl".
func (t AWS_Cloud9_EnvironmentEC2_Repository) SetV__RepositoryUrl(v string) AWS_Cloud9_EnvironmentEC2_Repository {
	t.RepositoryUrl = cfz.V(v)
	return t
}
