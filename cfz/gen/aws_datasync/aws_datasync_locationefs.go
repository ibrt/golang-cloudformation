// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datasync

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_DataSync_LocationEFS)(nil)
	_ cfz.Resource                   = (*AWS_DataSync_LocationEFS)(nil)
)

const (
	// AWS_DataSync_LocationEFS__Type is the CloudFormation type for AWS::DataSync::LocationEFS.
	AWS_DataSync_LocationEFS__Type = "AWS::DataSync::LocationEFS"
)

var (
	// AWS_DataSync_LocationEFS__AttributesMap reports all the CloudFormation attributes for AWS::DataSync::LocationEFS.
	AWS_DataSync_LocationEFS__AttributesMap = struct {
		LocationArn string
		LocationUri string
	}{
		LocationArn: "LocationArn",
		LocationUri: "LocationUri",
	}

	// AWS_DataSync_LocationEFS__AttributesSlice reports all the CloudFormation attributes for AWS::DataSync::LocationEFS.
	AWS_DataSync_LocationEFS__AttributesSlice = []string{
		AWS_DataSync_LocationEFS__AttributesMap.LocationArn,
		AWS_DataSync_LocationEFS__AttributesMap.LocationUri,
	}
)

var (
	// AWS_DataSync_LocationEFS__PropertiesMap reports all the CloudFormation properties for AWS::DataSync::LocationEFS.
	AWS_DataSync_LocationEFS__PropertiesMap = struct {
		AccessPointArn          string
		Ec2Config               string
		EfsFilesystemArn        string
		FileSystemAccessRoleArn string
		InTransitEncryption     string
		Subdirectory            string
		Tags                    string
	}{
		AccessPointArn:          "AccessPointArn",
		Ec2Config:               "Ec2Config",
		EfsFilesystemArn:        "EfsFilesystemArn",
		FileSystemAccessRoleArn: "FileSystemAccessRoleArn",
		InTransitEncryption:     "InTransitEncryption",
		Subdirectory:            "Subdirectory",
		Tags:                    "Tags",
	}

	// AWS_DataSync_LocationEFS__PropertiesSlice reports all the CloudFormation properties for AWS::DataSync::LocationEFS.
	AWS_DataSync_LocationEFS__PropertiesSlice = []string{
		AWS_DataSync_LocationEFS__PropertiesMap.AccessPointArn,
		AWS_DataSync_LocationEFS__PropertiesMap.Ec2Config,
		AWS_DataSync_LocationEFS__PropertiesMap.EfsFilesystemArn,
		AWS_DataSync_LocationEFS__PropertiesMap.FileSystemAccessRoleArn,
		AWS_DataSync_LocationEFS__PropertiesMap.InTransitEncryption,
		AWS_DataSync_LocationEFS__PropertiesMap.Subdirectory,
		AWS_DataSync_LocationEFS__PropertiesMap.Tags,
	}
)

// AWS_DataSync_LocationEFS is a binding for AWS::DataSync::LocationEFS.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html
type AWS_DataSync_LocationEFS struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AccessPointArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-accesspointarn
	AccessPointArn cfz.Expression[string] `json:"AccessPointArn,omitempty"`

	// Ec2Config is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-ec2config
	Ec2Config cfz.Expression[AWS_DataSync_LocationEFS_Ec2Config] `json:"Ec2Config,omitempty"`

	// EfsFilesystemArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-efsfilesystemarn
	EfsFilesystemArn cfz.Expression[string] `json:"EfsFilesystemArn,omitempty"`

	// FileSystemAccessRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-filesystemaccessrolearn
	FileSystemAccessRoleArn cfz.Expression[string] `json:"FileSystemAccessRoleArn,omitempty"`

	// InTransitEncryption is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-intransitencryption
	InTransitEncryption cfz.Expression[string] `json:"InTransitEncryption,omitempty"`

	// Subdirectory is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-subdirectory
	Subdirectory cfz.Expression[string] `json:"Subdirectory,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_DataSync_LocationEFS initializes a new *AWS_DataSync_LocationEFS.
func New__AWS_DataSync_LocationEFS(logicalName string) *AWS_DataSync_LocationEFS {
	return &AWS_DataSync_LocationEFS{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_DataSync_LocationEFS) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_DataSync_LocationEFS) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_DataSync_LocationEFS) GetType() string {
	return AWS_DataSync_LocationEFS__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_DataSync_LocationEFS) Set__LogicalName(v string) *AWS_DataSync_LocationEFS {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_DataSync_LocationEFS) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_DataSync_LocationEFS {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_DataSync_LocationEFS) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_DataSync_LocationEFS {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_DataSync_LocationEFS) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_DataSync_LocationEFS {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_DataSync_LocationEFS) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_DataSync_LocationEFS {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_DataSync_LocationEFS) Set__RequestedOutputs(v []cfz.Output) *AWS_DataSync_LocationEFS {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_DataSync_LocationEFS) Add__RequestedOutputs(v ...cfz.Output) *AWS_DataSync_LocationEFS {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AccessPointArn updates property "AccessPointArn".
func (t *AWS_DataSync_LocationEFS) Set__AccessPointArn(v cfz.Expression[string]) *AWS_DataSync_LocationEFS {
	t.AccessPointArn = v
	return t
}

// SetV__AccessPointArn updates property "AccessPointArn".
func (t *AWS_DataSync_LocationEFS) SetV__AccessPointArn(v string) *AWS_DataSync_LocationEFS {
	t.AccessPointArn = cfz.V(v)
	return t
}

// Set__Ec2Config updates property "Ec2Config".
func (t *AWS_DataSync_LocationEFS) Set__Ec2Config(v cfz.Expression[AWS_DataSync_LocationEFS_Ec2Config]) *AWS_DataSync_LocationEFS {
	t.Ec2Config = v
	return t
}

// SetV__Ec2Config updates property "Ec2Config".
func (t *AWS_DataSync_LocationEFS) SetV__Ec2Config(v AWS_DataSync_LocationEFS_Ec2Config) *AWS_DataSync_LocationEFS {
	t.Ec2Config = cfz.V(v)
	return t
}

// Set__EfsFilesystemArn updates property "EfsFilesystemArn".
func (t *AWS_DataSync_LocationEFS) Set__EfsFilesystemArn(v cfz.Expression[string]) *AWS_DataSync_LocationEFS {
	t.EfsFilesystemArn = v
	return t
}

// SetV__EfsFilesystemArn updates property "EfsFilesystemArn".
func (t *AWS_DataSync_LocationEFS) SetV__EfsFilesystemArn(v string) *AWS_DataSync_LocationEFS {
	t.EfsFilesystemArn = cfz.V(v)
	return t
}

// Set__FileSystemAccessRoleArn updates property "FileSystemAccessRoleArn".
func (t *AWS_DataSync_LocationEFS) Set__FileSystemAccessRoleArn(v cfz.Expression[string]) *AWS_DataSync_LocationEFS {
	t.FileSystemAccessRoleArn = v
	return t
}

// SetV__FileSystemAccessRoleArn updates property "FileSystemAccessRoleArn".
func (t *AWS_DataSync_LocationEFS) SetV__FileSystemAccessRoleArn(v string) *AWS_DataSync_LocationEFS {
	t.FileSystemAccessRoleArn = cfz.V(v)
	return t
}

// Set__InTransitEncryption updates property "InTransitEncryption".
func (t *AWS_DataSync_LocationEFS) Set__InTransitEncryption(v cfz.Expression[string]) *AWS_DataSync_LocationEFS {
	t.InTransitEncryption = v
	return t
}

// SetV__InTransitEncryption updates property "InTransitEncryption".
func (t *AWS_DataSync_LocationEFS) SetV__InTransitEncryption(v string) *AWS_DataSync_LocationEFS {
	t.InTransitEncryption = cfz.V(v)
	return t
}

// Set__Subdirectory updates property "Subdirectory".
func (t *AWS_DataSync_LocationEFS) Set__Subdirectory(v cfz.Expression[string]) *AWS_DataSync_LocationEFS {
	t.Subdirectory = v
	return t
}

// SetV__Subdirectory updates property "Subdirectory".
func (t *AWS_DataSync_LocationEFS) SetV__Subdirectory(v string) *AWS_DataSync_LocationEFS {
	t.Subdirectory = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_DataSync_LocationEFS) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_DataSync_LocationEFS {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_DataSync_LocationEFS) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_DataSync_LocationEFS {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_DataSync_LocationEFS) SetSV__Tags(v ...cfz.Tag) *AWS_DataSync_LocationEFS {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_DataSync_LocationEFS) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__LocationArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LocationArn
func (t *AWS_DataSync_LocationEFS) GetAtt__LocationArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataSync_LocationEFS__AttributesMap.LocationArn))
}

// GetAtt__LocationUri returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LocationUri
func (t *AWS_DataSync_LocationEFS) GetAtt__LocationUri() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataSync_LocationEFS__AttributesMap.LocationUri))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_DataSync_LocationEFS) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LocationArn returns a conventionally configured output for an attribute of this resource.
// Attribute: LocationArn
func (t *AWS_DataSync_LocationEFS) GetConventionalOutputAtt__LocationArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLocationArn", t.GetAtt__LocationArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LocationUri returns a conventionally configured output for an attribute of this resource.
// Attribute: LocationUri
func (t *AWS_DataSync_LocationEFS) GetConventionalOutputAtt__LocationUri(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLocationUri", t.GetAtt__LocationUri())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_DataSync_LocationEFS) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_DataSync_LocationEFS

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

func (t *AWS_DataSync_LocationEFS) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
