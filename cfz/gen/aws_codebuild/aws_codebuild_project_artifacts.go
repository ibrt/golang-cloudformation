// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codebuild

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodeBuild_Project_Artifacts__Type is the CloudFormation type for AWS::CodeBuild::Project.Artifacts.
	AWS_CodeBuild_Project_Artifacts__Type = "AWS::CodeBuild::Project.Artifacts"
)

var (
	// AWS_CodeBuild_Project_Artifacts__PropertiesMap reports all the CloudFormation properties for AWS::CodeBuild::Project.Artifacts.
	AWS_CodeBuild_Project_Artifacts__PropertiesMap = struct {
		ArtifactIdentifier   string
		EncryptionDisabled   string
		Location             string
		Name                 string
		NamespaceType        string
		OverrideArtifactName string
		Packaging            string
		Path                 string
		Type                 string
	}{
		ArtifactIdentifier:   "ArtifactIdentifier",
		EncryptionDisabled:   "EncryptionDisabled",
		Location:             "Location",
		Name:                 "Name",
		NamespaceType:        "NamespaceType",
		OverrideArtifactName: "OverrideArtifactName",
		Packaging:            "Packaging",
		Path:                 "Path",
		Type:                 "Type",
	}

	// AWS_CodeBuild_Project_Artifacts__PropertiesSlice reports all the CloudFormation properties for AWS::CodeBuild::Project.Artifacts.
	AWS_CodeBuild_Project_Artifacts__PropertiesSlice = []string{
		AWS_CodeBuild_Project_Artifacts__PropertiesMap.ArtifactIdentifier,
		AWS_CodeBuild_Project_Artifacts__PropertiesMap.EncryptionDisabled,
		AWS_CodeBuild_Project_Artifacts__PropertiesMap.Location,
		AWS_CodeBuild_Project_Artifacts__PropertiesMap.Name,
		AWS_CodeBuild_Project_Artifacts__PropertiesMap.NamespaceType,
		AWS_CodeBuild_Project_Artifacts__PropertiesMap.OverrideArtifactName,
		AWS_CodeBuild_Project_Artifacts__PropertiesMap.Packaging,
		AWS_CodeBuild_Project_Artifacts__PropertiesMap.Path,
		AWS_CodeBuild_Project_Artifacts__PropertiesMap.Type,
	}
)

// AWS_CodeBuild_Project_Artifacts is a binding for AWS::CodeBuild::Project.Artifacts.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html
type AWS_CodeBuild_Project_Artifacts struct {
	// ArtifactIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-artifactidentifier
	ArtifactIdentifier cfz.Expression[string] `json:"ArtifactIdentifier,omitempty"`

	// EncryptionDisabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-encryptiondisabled
	EncryptionDisabled cfz.Expression[bool] `json:"EncryptionDisabled,omitempty"`

	// Location is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-location
	Location cfz.Expression[string] `json:"Location,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// NamespaceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-namespacetype
	NamespaceType cfz.Expression[string] `json:"NamespaceType,omitempty"`

	// OverrideArtifactName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-overrideartifactname
	OverrideArtifactName cfz.Expression[bool] `json:"OverrideArtifactName,omitempty"`

	// Packaging is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-packaging
	Packaging cfz.Expression[string] `json:"Packaging,omitempty"`

	// Path is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-path
	Path cfz.Expression[string] `json:"Path,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_CodeBuild_Project_Artifacts initializes a new AWS_CodeBuild_Project_Artifacts.
func New__AWS_CodeBuild_Project_Artifacts() AWS_CodeBuild_Project_Artifacts {
	return AWS_CodeBuild_Project_Artifacts{}
}

// GetType returns the CloudFormation type.
func (AWS_CodeBuild_Project_Artifacts) GetType() string {
	return AWS_CodeBuild_Project_Artifacts__Type
}

// Set__ArtifactIdentifier updates property "ArtifactIdentifier".
func (t AWS_CodeBuild_Project_Artifacts) Set__ArtifactIdentifier(v cfz.Expression[string]) AWS_CodeBuild_Project_Artifacts {
	t.ArtifactIdentifier = v
	return t
}

// SetV__ArtifactIdentifier updates property "ArtifactIdentifier".
func (t AWS_CodeBuild_Project_Artifacts) SetV__ArtifactIdentifier(v string) AWS_CodeBuild_Project_Artifacts {
	t.ArtifactIdentifier = cfz.V(v)
	return t
}

// Set__EncryptionDisabled updates property "EncryptionDisabled".
func (t AWS_CodeBuild_Project_Artifacts) Set__EncryptionDisabled(v cfz.Expression[bool]) AWS_CodeBuild_Project_Artifacts {
	t.EncryptionDisabled = v
	return t
}

// SetV__EncryptionDisabled updates property "EncryptionDisabled".
func (t AWS_CodeBuild_Project_Artifacts) SetV__EncryptionDisabled(v bool) AWS_CodeBuild_Project_Artifacts {
	t.EncryptionDisabled = cfz.V(v)
	return t
}

// Set__Location updates property "Location".
func (t AWS_CodeBuild_Project_Artifacts) Set__Location(v cfz.Expression[string]) AWS_CodeBuild_Project_Artifacts {
	t.Location = v
	return t
}

// SetV__Location updates property "Location".
func (t AWS_CodeBuild_Project_Artifacts) SetV__Location(v string) AWS_CodeBuild_Project_Artifacts {
	t.Location = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_CodeBuild_Project_Artifacts) Set__Name(v cfz.Expression[string]) AWS_CodeBuild_Project_Artifacts {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_CodeBuild_Project_Artifacts) SetV__Name(v string) AWS_CodeBuild_Project_Artifacts {
	t.Name = cfz.V(v)
	return t
}

// Set__NamespaceType updates property "NamespaceType".
func (t AWS_CodeBuild_Project_Artifacts) Set__NamespaceType(v cfz.Expression[string]) AWS_CodeBuild_Project_Artifacts {
	t.NamespaceType = v
	return t
}

// SetV__NamespaceType updates property "NamespaceType".
func (t AWS_CodeBuild_Project_Artifacts) SetV__NamespaceType(v string) AWS_CodeBuild_Project_Artifacts {
	t.NamespaceType = cfz.V(v)
	return t
}

// Set__OverrideArtifactName updates property "OverrideArtifactName".
func (t AWS_CodeBuild_Project_Artifacts) Set__OverrideArtifactName(v cfz.Expression[bool]) AWS_CodeBuild_Project_Artifacts {
	t.OverrideArtifactName = v
	return t
}

// SetV__OverrideArtifactName updates property "OverrideArtifactName".
func (t AWS_CodeBuild_Project_Artifacts) SetV__OverrideArtifactName(v bool) AWS_CodeBuild_Project_Artifacts {
	t.OverrideArtifactName = cfz.V(v)
	return t
}

// Set__Packaging updates property "Packaging".
func (t AWS_CodeBuild_Project_Artifacts) Set__Packaging(v cfz.Expression[string]) AWS_CodeBuild_Project_Artifacts {
	t.Packaging = v
	return t
}

// SetV__Packaging updates property "Packaging".
func (t AWS_CodeBuild_Project_Artifacts) SetV__Packaging(v string) AWS_CodeBuild_Project_Artifacts {
	t.Packaging = cfz.V(v)
	return t
}

// Set__Path updates property "Path".
func (t AWS_CodeBuild_Project_Artifacts) Set__Path(v cfz.Expression[string]) AWS_CodeBuild_Project_Artifacts {
	t.Path = v
	return t
}

// SetV__Path updates property "Path".
func (t AWS_CodeBuild_Project_Artifacts) SetV__Path(v string) AWS_CodeBuild_Project_Artifacts {
	t.Path = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_CodeBuild_Project_Artifacts) Set__Type(v cfz.Expression[string]) AWS_CodeBuild_Project_Artifacts {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_CodeBuild_Project_Artifacts) SetV__Type(v string) AWS_CodeBuild_Project_Artifacts {
	t.Type = cfz.V(v)
	return t
}
