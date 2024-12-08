// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codeartifact

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodeArtifact_PackageGroup_RestrictionType__Type is the CloudFormation type for AWS::CodeArtifact::PackageGroup.RestrictionType.
	AWS_CodeArtifact_PackageGroup_RestrictionType__Type = "AWS::CodeArtifact::PackageGroup.RestrictionType"
)

var (
	// AWS_CodeArtifact_PackageGroup_RestrictionType__PropertiesMap reports all the CloudFormation properties for AWS::CodeArtifact::PackageGroup.RestrictionType.
	AWS_CodeArtifact_PackageGroup_RestrictionType__PropertiesMap = struct {
		Repositories    string
		RestrictionMode string
	}{
		Repositories:    "Repositories",
		RestrictionMode: "RestrictionMode",
	}

	// AWS_CodeArtifact_PackageGroup_RestrictionType__PropertiesSlice reports all the CloudFormation properties for AWS::CodeArtifact::PackageGroup.RestrictionType.
	AWS_CodeArtifact_PackageGroup_RestrictionType__PropertiesSlice = []string{
		AWS_CodeArtifact_PackageGroup_RestrictionType__PropertiesMap.Repositories,
		AWS_CodeArtifact_PackageGroup_RestrictionType__PropertiesMap.RestrictionMode,
	}
)

// AWS_CodeArtifact_PackageGroup_RestrictionType is a binding for AWS::CodeArtifact::PackageGroup.RestrictionType.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html
type AWS_CodeArtifact_PackageGroup_RestrictionType struct {
	// Repositories is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html#cfn-codeartifact-packagegroup-restrictiontype-repositories
	Repositories cfz.ExpressionSlice[string] `json:"Repositories,omitempty"`

	// RestrictionMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-packagegroup-restrictiontype.html#cfn-codeartifact-packagegroup-restrictiontype-restrictionmode
	RestrictionMode cfz.Expression[string] `json:"RestrictionMode,omitempty"`
}

// New__AWS_CodeArtifact_PackageGroup_RestrictionType initializes a new AWS_CodeArtifact_PackageGroup_RestrictionType.
func New__AWS_CodeArtifact_PackageGroup_RestrictionType() AWS_CodeArtifact_PackageGroup_RestrictionType {
	return AWS_CodeArtifact_PackageGroup_RestrictionType{}
}

// GetType returns the CloudFormation type.
func (AWS_CodeArtifact_PackageGroup_RestrictionType) GetType() string {
	return AWS_CodeArtifact_PackageGroup_RestrictionType__Type
}

// Set__Repositories updates property "Repositories".
func (t AWS_CodeArtifact_PackageGroup_RestrictionType) Set__Repositories(v cfz.ExpressionSlice[string]) AWS_CodeArtifact_PackageGroup_RestrictionType {
	t.Repositories = v
	return t
}

// SetS__Repositories updates property "Repositories".
func (t AWS_CodeArtifact_PackageGroup_RestrictionType) SetS__Repositories(v ...cfz.Expression[string]) AWS_CodeArtifact_PackageGroup_RestrictionType {
	t.Repositories = cfz.S(v...)
	return t
}

// SetSV__Repositories updates property "Repositories".
func (t AWS_CodeArtifact_PackageGroup_RestrictionType) SetSV__Repositories(v ...string) AWS_CodeArtifact_PackageGroup_RestrictionType {
	t.Repositories = cfz.SV(v...)
	return t
}

// Set__RestrictionMode updates property "RestrictionMode".
func (t AWS_CodeArtifact_PackageGroup_RestrictionType) Set__RestrictionMode(v cfz.Expression[string]) AWS_CodeArtifact_PackageGroup_RestrictionType {
	t.RestrictionMode = v
	return t
}

// SetV__RestrictionMode updates property "RestrictionMode".
func (t AWS_CodeArtifact_PackageGroup_RestrictionType) SetV__RestrictionMode(v string) AWS_CodeArtifact_PackageGroup_RestrictionType {
	t.RestrictionMode = cfz.V(v)
	return t
}
