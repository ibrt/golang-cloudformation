// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codeartifact

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodeArtifact_PackageGroup_OriginConfiguration__Type is the CloudFormation type for AWS::CodeArtifact::PackageGroup.OriginConfiguration.
	AWS_CodeArtifact_PackageGroup_OriginConfiguration__Type = "AWS::CodeArtifact::PackageGroup.OriginConfiguration"
)

var (
	// AWS_CodeArtifact_PackageGroup_OriginConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::CodeArtifact::PackageGroup.OriginConfiguration.
	AWS_CodeArtifact_PackageGroup_OriginConfiguration__PropertiesMap = struct {
		Restrictions string
	}{
		Restrictions: "Restrictions",
	}

	// AWS_CodeArtifact_PackageGroup_OriginConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::CodeArtifact::PackageGroup.OriginConfiguration.
	AWS_CodeArtifact_PackageGroup_OriginConfiguration__PropertiesSlice = []string{
		AWS_CodeArtifact_PackageGroup_OriginConfiguration__PropertiesMap.Restrictions,
	}
)

// AWS_CodeArtifact_PackageGroup_OriginConfiguration is a binding for AWS::CodeArtifact::PackageGroup.OriginConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-originconfiguration.html
type AWS_CodeArtifact_PackageGroup_OriginConfiguration struct {
	// Restrictions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-originconfiguration.html#cfn-codeartifact-packagegroup-originconfiguration-restrictions
	Restrictions cfz.Expression[AWS_CodeArtifact_PackageGroup_Restrictions] `json:"Restrictions,omitempty"`
}

// New__AWS_CodeArtifact_PackageGroup_OriginConfiguration initializes a new AWS_CodeArtifact_PackageGroup_OriginConfiguration.
func New__AWS_CodeArtifact_PackageGroup_OriginConfiguration() AWS_CodeArtifact_PackageGroup_OriginConfiguration {
	return AWS_CodeArtifact_PackageGroup_OriginConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_CodeArtifact_PackageGroup_OriginConfiguration) GetType() string {
	return AWS_CodeArtifact_PackageGroup_OriginConfiguration__Type
}

// Set__Restrictions updates property "Restrictions".
func (t AWS_CodeArtifact_PackageGroup_OriginConfiguration) Set__Restrictions(v cfz.Expression[AWS_CodeArtifact_PackageGroup_Restrictions]) AWS_CodeArtifact_PackageGroup_OriginConfiguration {
	t.Restrictions = v
	return t
}

// SetV__Restrictions updates property "Restrictions".
func (t AWS_CodeArtifact_PackageGroup_OriginConfiguration) SetV__Restrictions(v AWS_CodeArtifact_PackageGroup_Restrictions) AWS_CodeArtifact_PackageGroup_OriginConfiguration {
	t.Restrictions = cfz.V(v)
	return t
}
