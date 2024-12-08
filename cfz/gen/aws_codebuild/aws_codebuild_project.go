// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codebuild

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_CodeBuild_Project)(nil)
	_ cfz.Resource                   = (*AWS_CodeBuild_Project)(nil)
)

const (
	// AWS_CodeBuild_Project__Type is the CloudFormation type for AWS::CodeBuild::Project.
	AWS_CodeBuild_Project__Type = "AWS::CodeBuild::Project"
)

var (
	// AWS_CodeBuild_Project__AttributesMap reports all the CloudFormation attributes for AWS::CodeBuild::Project.
	AWS_CodeBuild_Project__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_CodeBuild_Project__AttributesSlice reports all the CloudFormation attributes for AWS::CodeBuild::Project.
	AWS_CodeBuild_Project__AttributesSlice = []string{
		AWS_CodeBuild_Project__AttributesMap.Arn,
	}
)

var (
	// AWS_CodeBuild_Project__PropertiesMap reports all the CloudFormation properties for AWS::CodeBuild::Project.
	AWS_CodeBuild_Project__PropertiesMap = struct {
		Artifacts               string
		AutoRetryLimit          string
		BadgeEnabled            string
		BuildBatchConfig        string
		Cache                   string
		ConcurrentBuildLimit    string
		Description             string
		EncryptionKey           string
		Environment             string
		FileSystemLocations     string
		LogsConfig              string
		Name                    string
		QueuedTimeoutInMinutes  string
		ResourceAccessRole      string
		SecondaryArtifacts      string
		SecondarySourceVersions string
		SecondarySources        string
		ServiceRole             string
		Source                  string
		SourceVersion           string
		Tags                    string
		TimeoutInMinutes        string
		Triggers                string
		Visibility              string
		VpcConfig               string
	}{
		Artifacts:               "Artifacts",
		AutoRetryLimit:          "AutoRetryLimit",
		BadgeEnabled:            "BadgeEnabled",
		BuildBatchConfig:        "BuildBatchConfig",
		Cache:                   "Cache",
		ConcurrentBuildLimit:    "ConcurrentBuildLimit",
		Description:             "Description",
		EncryptionKey:           "EncryptionKey",
		Environment:             "Environment",
		FileSystemLocations:     "FileSystemLocations",
		LogsConfig:              "LogsConfig",
		Name:                    "Name",
		QueuedTimeoutInMinutes:  "QueuedTimeoutInMinutes",
		ResourceAccessRole:      "ResourceAccessRole",
		SecondaryArtifacts:      "SecondaryArtifacts",
		SecondarySourceVersions: "SecondarySourceVersions",
		SecondarySources:        "SecondarySources",
		ServiceRole:             "ServiceRole",
		Source:                  "Source",
		SourceVersion:           "SourceVersion",
		Tags:                    "Tags",
		TimeoutInMinutes:        "TimeoutInMinutes",
		Triggers:                "Triggers",
		Visibility:              "Visibility",
		VpcConfig:               "VpcConfig",
	}

	// AWS_CodeBuild_Project__PropertiesSlice reports all the CloudFormation properties for AWS::CodeBuild::Project.
	AWS_CodeBuild_Project__PropertiesSlice = []string{
		AWS_CodeBuild_Project__PropertiesMap.Artifacts,
		AWS_CodeBuild_Project__PropertiesMap.AutoRetryLimit,
		AWS_CodeBuild_Project__PropertiesMap.BadgeEnabled,
		AWS_CodeBuild_Project__PropertiesMap.BuildBatchConfig,
		AWS_CodeBuild_Project__PropertiesMap.Cache,
		AWS_CodeBuild_Project__PropertiesMap.ConcurrentBuildLimit,
		AWS_CodeBuild_Project__PropertiesMap.Description,
		AWS_CodeBuild_Project__PropertiesMap.EncryptionKey,
		AWS_CodeBuild_Project__PropertiesMap.Environment,
		AWS_CodeBuild_Project__PropertiesMap.FileSystemLocations,
		AWS_CodeBuild_Project__PropertiesMap.LogsConfig,
		AWS_CodeBuild_Project__PropertiesMap.Name,
		AWS_CodeBuild_Project__PropertiesMap.QueuedTimeoutInMinutes,
		AWS_CodeBuild_Project__PropertiesMap.ResourceAccessRole,
		AWS_CodeBuild_Project__PropertiesMap.SecondaryArtifacts,
		AWS_CodeBuild_Project__PropertiesMap.SecondarySourceVersions,
		AWS_CodeBuild_Project__PropertiesMap.SecondarySources,
		AWS_CodeBuild_Project__PropertiesMap.ServiceRole,
		AWS_CodeBuild_Project__PropertiesMap.Source,
		AWS_CodeBuild_Project__PropertiesMap.SourceVersion,
		AWS_CodeBuild_Project__PropertiesMap.Tags,
		AWS_CodeBuild_Project__PropertiesMap.TimeoutInMinutes,
		AWS_CodeBuild_Project__PropertiesMap.Triggers,
		AWS_CodeBuild_Project__PropertiesMap.Visibility,
		AWS_CodeBuild_Project__PropertiesMap.VpcConfig,
	}
)

// AWS_CodeBuild_Project is a binding for AWS::CodeBuild::Project.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html
type AWS_CodeBuild_Project struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Artifacts is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-artifacts
	Artifacts cfz.Expression[AWS_CodeBuild_Project_Artifacts] `json:"Artifacts,omitempty"`

	// AutoRetryLimit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-autoretrylimit
	AutoRetryLimit cfz.Expression[int32] `json:"AutoRetryLimit,omitempty"`

	// BadgeEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-badgeenabled
	BadgeEnabled cfz.Expression[bool] `json:"BadgeEnabled,omitempty"`

	// BuildBatchConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-buildbatchconfig
	BuildBatchConfig cfz.Expression[AWS_CodeBuild_Project_ProjectBuildBatchConfig] `json:"BuildBatchConfig,omitempty"`

	// Cache is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-cache
	Cache cfz.Expression[AWS_CodeBuild_Project_ProjectCache] `json:"Cache,omitempty"`

	// ConcurrentBuildLimit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-concurrentbuildlimit
	ConcurrentBuildLimit cfz.Expression[int32] `json:"ConcurrentBuildLimit,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// EncryptionKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-encryptionkey
	EncryptionKey cfz.Expression[string] `json:"EncryptionKey,omitempty"`

	// Environment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-environment
	Environment cfz.Expression[AWS_CodeBuild_Project_Environment] `json:"Environment,omitempty"`

	// FileSystemLocations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-filesystemlocations
	FileSystemLocations cfz.ExpressionSlice[AWS_CodeBuild_Project_ProjectFileSystemLocation] `json:"FileSystemLocations,omitempty"`

	// LogsConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-logsconfig
	LogsConfig cfz.Expression[AWS_CodeBuild_Project_LogsConfig] `json:"LogsConfig,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// QueuedTimeoutInMinutes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-queuedtimeoutinminutes
	QueuedTimeoutInMinutes cfz.Expression[int32] `json:"QueuedTimeoutInMinutes,omitempty"`

	// ResourceAccessRole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-resourceaccessrole
	ResourceAccessRole cfz.Expression[string] `json:"ResourceAccessRole,omitempty"`

	// SecondaryArtifacts is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-secondaryartifacts
	SecondaryArtifacts cfz.ExpressionSlice[AWS_CodeBuild_Project_Artifacts] `json:"SecondaryArtifacts,omitempty"`

	// SecondarySourceVersions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-secondarysourceversions
	SecondarySourceVersions cfz.ExpressionSlice[AWS_CodeBuild_Project_ProjectSourceVersion] `json:"SecondarySourceVersions,omitempty"`

	// SecondarySources is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-secondarysources
	SecondarySources cfz.ExpressionSlice[AWS_CodeBuild_Project_Source] `json:"SecondarySources,omitempty"`

	// ServiceRole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-servicerole
	ServiceRole cfz.Expression[string] `json:"ServiceRole,omitempty"`

	// Source is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-source
	Source cfz.Expression[AWS_CodeBuild_Project_Source] `json:"Source,omitempty"`

	// SourceVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-sourceversion
	SourceVersion cfz.Expression[string] `json:"SourceVersion,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// TimeoutInMinutes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-timeoutinminutes
	TimeoutInMinutes cfz.Expression[int32] `json:"TimeoutInMinutes,omitempty"`

	// Triggers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-triggers
	Triggers cfz.Expression[AWS_CodeBuild_Project_ProjectTriggers] `json:"Triggers,omitempty"`

	// Visibility is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-visibility
	Visibility cfz.Expression[string] `json:"Visibility,omitempty"`

	// VpcConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-vpcconfig
	VpcConfig cfz.Expression[AWS_CodeBuild_Project_VpcConfig] `json:"VpcConfig,omitempty"`
}

// New__AWS_CodeBuild_Project initializes a new *AWS_CodeBuild_Project.
func New__AWS_CodeBuild_Project(logicalName string) *AWS_CodeBuild_Project {
	return &AWS_CodeBuild_Project{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_CodeBuild_Project) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_CodeBuild_Project) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_CodeBuild_Project) GetType() string {
	return AWS_CodeBuild_Project__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_CodeBuild_Project) Set__LogicalName(v string) *AWS_CodeBuild_Project {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_CodeBuild_Project) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_CodeBuild_Project {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_CodeBuild_Project) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_CodeBuild_Project {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_CodeBuild_Project) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_CodeBuild_Project {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_CodeBuild_Project) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_CodeBuild_Project {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_CodeBuild_Project) Set__RequestedOutputs(v []cfz.Output) *AWS_CodeBuild_Project {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_CodeBuild_Project) Add__RequestedOutputs(v ...cfz.Output) *AWS_CodeBuild_Project {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Artifacts updates property "Artifacts".
func (t *AWS_CodeBuild_Project) Set__Artifacts(v cfz.Expression[AWS_CodeBuild_Project_Artifacts]) *AWS_CodeBuild_Project {
	t.Artifacts = v
	return t
}

// SetV__Artifacts updates property "Artifacts".
func (t *AWS_CodeBuild_Project) SetV__Artifacts(v AWS_CodeBuild_Project_Artifacts) *AWS_CodeBuild_Project {
	t.Artifacts = cfz.V(v)
	return t
}

// Set__AutoRetryLimit updates property "AutoRetryLimit".
func (t *AWS_CodeBuild_Project) Set__AutoRetryLimit(v cfz.Expression[int32]) *AWS_CodeBuild_Project {
	t.AutoRetryLimit = v
	return t
}

// SetV__AutoRetryLimit updates property "AutoRetryLimit".
func (t *AWS_CodeBuild_Project) SetV__AutoRetryLimit(v int32) *AWS_CodeBuild_Project {
	t.AutoRetryLimit = cfz.V(v)
	return t
}

// Set__BadgeEnabled updates property "BadgeEnabled".
func (t *AWS_CodeBuild_Project) Set__BadgeEnabled(v cfz.Expression[bool]) *AWS_CodeBuild_Project {
	t.BadgeEnabled = v
	return t
}

// SetV__BadgeEnabled updates property "BadgeEnabled".
func (t *AWS_CodeBuild_Project) SetV__BadgeEnabled(v bool) *AWS_CodeBuild_Project {
	t.BadgeEnabled = cfz.V(v)
	return t
}

// Set__BuildBatchConfig updates property "BuildBatchConfig".
func (t *AWS_CodeBuild_Project) Set__BuildBatchConfig(v cfz.Expression[AWS_CodeBuild_Project_ProjectBuildBatchConfig]) *AWS_CodeBuild_Project {
	t.BuildBatchConfig = v
	return t
}

// SetV__BuildBatchConfig updates property "BuildBatchConfig".
func (t *AWS_CodeBuild_Project) SetV__BuildBatchConfig(v AWS_CodeBuild_Project_ProjectBuildBatchConfig) *AWS_CodeBuild_Project {
	t.BuildBatchConfig = cfz.V(v)
	return t
}

// Set__Cache updates property "Cache".
func (t *AWS_CodeBuild_Project) Set__Cache(v cfz.Expression[AWS_CodeBuild_Project_ProjectCache]) *AWS_CodeBuild_Project {
	t.Cache = v
	return t
}

// SetV__Cache updates property "Cache".
func (t *AWS_CodeBuild_Project) SetV__Cache(v AWS_CodeBuild_Project_ProjectCache) *AWS_CodeBuild_Project {
	t.Cache = cfz.V(v)
	return t
}

// Set__ConcurrentBuildLimit updates property "ConcurrentBuildLimit".
func (t *AWS_CodeBuild_Project) Set__ConcurrentBuildLimit(v cfz.Expression[int32]) *AWS_CodeBuild_Project {
	t.ConcurrentBuildLimit = v
	return t
}

// SetV__ConcurrentBuildLimit updates property "ConcurrentBuildLimit".
func (t *AWS_CodeBuild_Project) SetV__ConcurrentBuildLimit(v int32) *AWS_CodeBuild_Project {
	t.ConcurrentBuildLimit = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_CodeBuild_Project) Set__Description(v cfz.Expression[string]) *AWS_CodeBuild_Project {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_CodeBuild_Project) SetV__Description(v string) *AWS_CodeBuild_Project {
	t.Description = cfz.V(v)
	return t
}

// Set__EncryptionKey updates property "EncryptionKey".
func (t *AWS_CodeBuild_Project) Set__EncryptionKey(v cfz.Expression[string]) *AWS_CodeBuild_Project {
	t.EncryptionKey = v
	return t
}

// SetV__EncryptionKey updates property "EncryptionKey".
func (t *AWS_CodeBuild_Project) SetV__EncryptionKey(v string) *AWS_CodeBuild_Project {
	t.EncryptionKey = cfz.V(v)
	return t
}

// Set__Environment updates property "Environment".
func (t *AWS_CodeBuild_Project) Set__Environment(v cfz.Expression[AWS_CodeBuild_Project_Environment]) *AWS_CodeBuild_Project {
	t.Environment = v
	return t
}

// SetV__Environment updates property "Environment".
func (t *AWS_CodeBuild_Project) SetV__Environment(v AWS_CodeBuild_Project_Environment) *AWS_CodeBuild_Project {
	t.Environment = cfz.V(v)
	return t
}

// Set__FileSystemLocations updates property "FileSystemLocations".
func (t *AWS_CodeBuild_Project) Set__FileSystemLocations(v cfz.ExpressionSlice[AWS_CodeBuild_Project_ProjectFileSystemLocation]) *AWS_CodeBuild_Project {
	t.FileSystemLocations = v
	return t
}

// SetS__FileSystemLocations updates property "FileSystemLocations".
func (t *AWS_CodeBuild_Project) SetS__FileSystemLocations(v ...cfz.Expression[AWS_CodeBuild_Project_ProjectFileSystemLocation]) *AWS_CodeBuild_Project {
	t.FileSystemLocations = cfz.S(v...)
	return t
}

// SetSV__FileSystemLocations updates property "FileSystemLocations".
func (t *AWS_CodeBuild_Project) SetSV__FileSystemLocations(v ...AWS_CodeBuild_Project_ProjectFileSystemLocation) *AWS_CodeBuild_Project {
	t.FileSystemLocations = cfz.SV(v...)
	return t
}

// Set__LogsConfig updates property "LogsConfig".
func (t *AWS_CodeBuild_Project) Set__LogsConfig(v cfz.Expression[AWS_CodeBuild_Project_LogsConfig]) *AWS_CodeBuild_Project {
	t.LogsConfig = v
	return t
}

// SetV__LogsConfig updates property "LogsConfig".
func (t *AWS_CodeBuild_Project) SetV__LogsConfig(v AWS_CodeBuild_Project_LogsConfig) *AWS_CodeBuild_Project {
	t.LogsConfig = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_CodeBuild_Project) Set__Name(v cfz.Expression[string]) *AWS_CodeBuild_Project {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_CodeBuild_Project) SetV__Name(v string) *AWS_CodeBuild_Project {
	t.Name = cfz.V(v)
	return t
}

// Set__QueuedTimeoutInMinutes updates property "QueuedTimeoutInMinutes".
func (t *AWS_CodeBuild_Project) Set__QueuedTimeoutInMinutes(v cfz.Expression[int32]) *AWS_CodeBuild_Project {
	t.QueuedTimeoutInMinutes = v
	return t
}

// SetV__QueuedTimeoutInMinutes updates property "QueuedTimeoutInMinutes".
func (t *AWS_CodeBuild_Project) SetV__QueuedTimeoutInMinutes(v int32) *AWS_CodeBuild_Project {
	t.QueuedTimeoutInMinutes = cfz.V(v)
	return t
}

// Set__ResourceAccessRole updates property "ResourceAccessRole".
func (t *AWS_CodeBuild_Project) Set__ResourceAccessRole(v cfz.Expression[string]) *AWS_CodeBuild_Project {
	t.ResourceAccessRole = v
	return t
}

// SetV__ResourceAccessRole updates property "ResourceAccessRole".
func (t *AWS_CodeBuild_Project) SetV__ResourceAccessRole(v string) *AWS_CodeBuild_Project {
	t.ResourceAccessRole = cfz.V(v)
	return t
}

// Set__SecondaryArtifacts updates property "SecondaryArtifacts".
func (t *AWS_CodeBuild_Project) Set__SecondaryArtifacts(v cfz.ExpressionSlice[AWS_CodeBuild_Project_Artifacts]) *AWS_CodeBuild_Project {
	t.SecondaryArtifacts = v
	return t
}

// SetS__SecondaryArtifacts updates property "SecondaryArtifacts".
func (t *AWS_CodeBuild_Project) SetS__SecondaryArtifacts(v ...cfz.Expression[AWS_CodeBuild_Project_Artifacts]) *AWS_CodeBuild_Project {
	t.SecondaryArtifacts = cfz.S(v...)
	return t
}

// SetSV__SecondaryArtifacts updates property "SecondaryArtifacts".
func (t *AWS_CodeBuild_Project) SetSV__SecondaryArtifacts(v ...AWS_CodeBuild_Project_Artifacts) *AWS_CodeBuild_Project {
	t.SecondaryArtifacts = cfz.SV(v...)
	return t
}

// Set__SecondarySourceVersions updates property "SecondarySourceVersions".
func (t *AWS_CodeBuild_Project) Set__SecondarySourceVersions(v cfz.ExpressionSlice[AWS_CodeBuild_Project_ProjectSourceVersion]) *AWS_CodeBuild_Project {
	t.SecondarySourceVersions = v
	return t
}

// SetS__SecondarySourceVersions updates property "SecondarySourceVersions".
func (t *AWS_CodeBuild_Project) SetS__SecondarySourceVersions(v ...cfz.Expression[AWS_CodeBuild_Project_ProjectSourceVersion]) *AWS_CodeBuild_Project {
	t.SecondarySourceVersions = cfz.S(v...)
	return t
}

// SetSV__SecondarySourceVersions updates property "SecondarySourceVersions".
func (t *AWS_CodeBuild_Project) SetSV__SecondarySourceVersions(v ...AWS_CodeBuild_Project_ProjectSourceVersion) *AWS_CodeBuild_Project {
	t.SecondarySourceVersions = cfz.SV(v...)
	return t
}

// Set__SecondarySources updates property "SecondarySources".
func (t *AWS_CodeBuild_Project) Set__SecondarySources(v cfz.ExpressionSlice[AWS_CodeBuild_Project_Source]) *AWS_CodeBuild_Project {
	t.SecondarySources = v
	return t
}

// SetS__SecondarySources updates property "SecondarySources".
func (t *AWS_CodeBuild_Project) SetS__SecondarySources(v ...cfz.Expression[AWS_CodeBuild_Project_Source]) *AWS_CodeBuild_Project {
	t.SecondarySources = cfz.S(v...)
	return t
}

// SetSV__SecondarySources updates property "SecondarySources".
func (t *AWS_CodeBuild_Project) SetSV__SecondarySources(v ...AWS_CodeBuild_Project_Source) *AWS_CodeBuild_Project {
	t.SecondarySources = cfz.SV(v...)
	return t
}

// Set__ServiceRole updates property "ServiceRole".
func (t *AWS_CodeBuild_Project) Set__ServiceRole(v cfz.Expression[string]) *AWS_CodeBuild_Project {
	t.ServiceRole = v
	return t
}

// SetV__ServiceRole updates property "ServiceRole".
func (t *AWS_CodeBuild_Project) SetV__ServiceRole(v string) *AWS_CodeBuild_Project {
	t.ServiceRole = cfz.V(v)
	return t
}

// Set__Source updates property "Source".
func (t *AWS_CodeBuild_Project) Set__Source(v cfz.Expression[AWS_CodeBuild_Project_Source]) *AWS_CodeBuild_Project {
	t.Source = v
	return t
}

// SetV__Source updates property "Source".
func (t *AWS_CodeBuild_Project) SetV__Source(v AWS_CodeBuild_Project_Source) *AWS_CodeBuild_Project {
	t.Source = cfz.V(v)
	return t
}

// Set__SourceVersion updates property "SourceVersion".
func (t *AWS_CodeBuild_Project) Set__SourceVersion(v cfz.Expression[string]) *AWS_CodeBuild_Project {
	t.SourceVersion = v
	return t
}

// SetV__SourceVersion updates property "SourceVersion".
func (t *AWS_CodeBuild_Project) SetV__SourceVersion(v string) *AWS_CodeBuild_Project {
	t.SourceVersion = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_CodeBuild_Project) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_CodeBuild_Project {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_CodeBuild_Project) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_CodeBuild_Project {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_CodeBuild_Project) SetSV__Tags(v ...cfz.Tag) *AWS_CodeBuild_Project {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TimeoutInMinutes updates property "TimeoutInMinutes".
func (t *AWS_CodeBuild_Project) Set__TimeoutInMinutes(v cfz.Expression[int32]) *AWS_CodeBuild_Project {
	t.TimeoutInMinutes = v
	return t
}

// SetV__TimeoutInMinutes updates property "TimeoutInMinutes".
func (t *AWS_CodeBuild_Project) SetV__TimeoutInMinutes(v int32) *AWS_CodeBuild_Project {
	t.TimeoutInMinutes = cfz.V(v)
	return t
}

// Set__Triggers updates property "Triggers".
func (t *AWS_CodeBuild_Project) Set__Triggers(v cfz.Expression[AWS_CodeBuild_Project_ProjectTriggers]) *AWS_CodeBuild_Project {
	t.Triggers = v
	return t
}

// SetV__Triggers updates property "Triggers".
func (t *AWS_CodeBuild_Project) SetV__Triggers(v AWS_CodeBuild_Project_ProjectTriggers) *AWS_CodeBuild_Project {
	t.Triggers = cfz.V(v)
	return t
}

// Set__Visibility updates property "Visibility".
func (t *AWS_CodeBuild_Project) Set__Visibility(v cfz.Expression[string]) *AWS_CodeBuild_Project {
	t.Visibility = v
	return t
}

// SetV__Visibility updates property "Visibility".
func (t *AWS_CodeBuild_Project) SetV__Visibility(v string) *AWS_CodeBuild_Project {
	t.Visibility = cfz.V(v)
	return t
}

// Set__VpcConfig updates property "VpcConfig".
func (t *AWS_CodeBuild_Project) Set__VpcConfig(v cfz.Expression[AWS_CodeBuild_Project_VpcConfig]) *AWS_CodeBuild_Project {
	t.VpcConfig = v
	return t
}

// SetV__VpcConfig updates property "VpcConfig".
func (t *AWS_CodeBuild_Project) SetV__VpcConfig(v AWS_CodeBuild_Project_VpcConfig) *AWS_CodeBuild_Project {
	t.VpcConfig = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_CodeBuild_Project) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_CodeBuild_Project) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CodeBuild_Project__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_CodeBuild_Project) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_CodeBuild_Project) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_CodeBuild_Project) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_CodeBuild_Project

	return json.Marshal(struct {
		Type                string                          `json:"Type"`
		DependsOn           []string                        `json:"DependsOn,omitempty"`
		DeletionPolicy      cfz.ResourceDeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Properties          *Properties                     `json:"Properties,omitempty"`
	}{
		Type:       t.GetType(),
		DependsOn:  t.getDependsOn(),
		Properties: (*Properties)(t),
	})
}

func (t *AWS_CodeBuild_Project) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
