// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_QuickSight_Folder)(nil)
	_ cfz.Resource                   = (*AWS_QuickSight_Folder)(nil)
)

const (
	// AWS_QuickSight_Folder__Type is the CloudFormation type for AWS::QuickSight::Folder.
	AWS_QuickSight_Folder__Type = "AWS::QuickSight::Folder"
)

var (
	// AWS_QuickSight_Folder__AttributesMap reports all the CloudFormation attributes for AWS::QuickSight::Folder.
	AWS_QuickSight_Folder__AttributesMap = struct {
		Arn             string
		CreatedTime     string
		LastUpdatedTime string
	}{
		Arn:             "Arn",
		CreatedTime:     "CreatedTime",
		LastUpdatedTime: "LastUpdatedTime",
	}

	// AWS_QuickSight_Folder__AttributesSlice reports all the CloudFormation attributes for AWS::QuickSight::Folder.
	AWS_QuickSight_Folder__AttributesSlice = []string{
		AWS_QuickSight_Folder__AttributesMap.Arn,
		AWS_QuickSight_Folder__AttributesMap.CreatedTime,
		AWS_QuickSight_Folder__AttributesMap.LastUpdatedTime,
	}
)

var (
	// AWS_QuickSight_Folder__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::Folder.
	AWS_QuickSight_Folder__PropertiesMap = struct {
		AwsAccountId    string
		FolderId        string
		FolderType      string
		Name            string
		ParentFolderArn string
		Permissions     string
		SharingModel    string
		Tags            string
	}{
		AwsAccountId:    "AwsAccountId",
		FolderId:        "FolderId",
		FolderType:      "FolderType",
		Name:            "Name",
		ParentFolderArn: "ParentFolderArn",
		Permissions:     "Permissions",
		SharingModel:    "SharingModel",
		Tags:            "Tags",
	}

	// AWS_QuickSight_Folder__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::Folder.
	AWS_QuickSight_Folder__PropertiesSlice = []string{
		AWS_QuickSight_Folder__PropertiesMap.AwsAccountId,
		AWS_QuickSight_Folder__PropertiesMap.FolderId,
		AWS_QuickSight_Folder__PropertiesMap.FolderType,
		AWS_QuickSight_Folder__PropertiesMap.Name,
		AWS_QuickSight_Folder__PropertiesMap.ParentFolderArn,
		AWS_QuickSight_Folder__PropertiesMap.Permissions,
		AWS_QuickSight_Folder__PropertiesMap.SharingModel,
		AWS_QuickSight_Folder__PropertiesMap.Tags,
	}
)

// AWS_QuickSight_Folder is a binding for AWS::QuickSight::Folder.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html
type AWS_QuickSight_Folder struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AwsAccountId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-awsaccountid
	AwsAccountId cfz.Expression[string] `json:"AwsAccountId,omitempty"`

	// FolderId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-folderid
	FolderId cfz.Expression[string] `json:"FolderId,omitempty"`

	// FolderType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-foldertype
	FolderType cfz.Expression[string] `json:"FolderType,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// ParentFolderArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-parentfolderarn
	ParentFolderArn cfz.Expression[string] `json:"ParentFolderArn,omitempty"`

	// Permissions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-permissions
	Permissions cfz.ExpressionSlice[AWS_QuickSight_Folder_ResourcePermission] `json:"Permissions,omitempty"`

	// SharingModel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-sharingmodel
	SharingModel cfz.Expression[string] `json:"SharingModel,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_QuickSight_Folder initializes a new *AWS_QuickSight_Folder.
func New__AWS_QuickSight_Folder(logicalName string) *AWS_QuickSight_Folder {
	return &AWS_QuickSight_Folder{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_QuickSight_Folder) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_QuickSight_Folder) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_QuickSight_Folder) GetType() string {
	return AWS_QuickSight_Folder__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_QuickSight_Folder) Set__LogicalName(v string) *AWS_QuickSight_Folder {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_QuickSight_Folder) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_QuickSight_Folder {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_QuickSight_Folder) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_QuickSight_Folder {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_QuickSight_Folder) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_QuickSight_Folder {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_QuickSight_Folder) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_QuickSight_Folder {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_QuickSight_Folder) Set__RequestedOutputs(v []cfz.Output) *AWS_QuickSight_Folder {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_QuickSight_Folder) Add__RequestedOutputs(v ...cfz.Output) *AWS_QuickSight_Folder {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AwsAccountId updates property "AwsAccountId".
func (t *AWS_QuickSight_Folder) Set__AwsAccountId(v cfz.Expression[string]) *AWS_QuickSight_Folder {
	t.AwsAccountId = v
	return t
}

// SetV__AwsAccountId updates property "AwsAccountId".
func (t *AWS_QuickSight_Folder) SetV__AwsAccountId(v string) *AWS_QuickSight_Folder {
	t.AwsAccountId = cfz.V(v)
	return t
}

// Set__FolderId updates property "FolderId".
func (t *AWS_QuickSight_Folder) Set__FolderId(v cfz.Expression[string]) *AWS_QuickSight_Folder {
	t.FolderId = v
	return t
}

// SetV__FolderId updates property "FolderId".
func (t *AWS_QuickSight_Folder) SetV__FolderId(v string) *AWS_QuickSight_Folder {
	t.FolderId = cfz.V(v)
	return t
}

// Set__FolderType updates property "FolderType".
func (t *AWS_QuickSight_Folder) Set__FolderType(v cfz.Expression[string]) *AWS_QuickSight_Folder {
	t.FolderType = v
	return t
}

// SetV__FolderType updates property "FolderType".
func (t *AWS_QuickSight_Folder) SetV__FolderType(v string) *AWS_QuickSight_Folder {
	t.FolderType = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_QuickSight_Folder) Set__Name(v cfz.Expression[string]) *AWS_QuickSight_Folder {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_QuickSight_Folder) SetV__Name(v string) *AWS_QuickSight_Folder {
	t.Name = cfz.V(v)
	return t
}

// Set__ParentFolderArn updates property "ParentFolderArn".
func (t *AWS_QuickSight_Folder) Set__ParentFolderArn(v cfz.Expression[string]) *AWS_QuickSight_Folder {
	t.ParentFolderArn = v
	return t
}

// SetV__ParentFolderArn updates property "ParentFolderArn".
func (t *AWS_QuickSight_Folder) SetV__ParentFolderArn(v string) *AWS_QuickSight_Folder {
	t.ParentFolderArn = cfz.V(v)
	return t
}

// Set__Permissions updates property "Permissions".
func (t *AWS_QuickSight_Folder) Set__Permissions(v cfz.ExpressionSlice[AWS_QuickSight_Folder_ResourcePermission]) *AWS_QuickSight_Folder {
	t.Permissions = v
	return t
}

// SetS__Permissions updates property "Permissions".
func (t *AWS_QuickSight_Folder) SetS__Permissions(v ...cfz.Expression[AWS_QuickSight_Folder_ResourcePermission]) *AWS_QuickSight_Folder {
	t.Permissions = cfz.S(v...)
	return t
}

// SetSV__Permissions updates property "Permissions".
func (t *AWS_QuickSight_Folder) SetSV__Permissions(v ...AWS_QuickSight_Folder_ResourcePermission) *AWS_QuickSight_Folder {
	t.Permissions = cfz.SV(v...)
	return t
}

// Set__SharingModel updates property "SharingModel".
func (t *AWS_QuickSight_Folder) Set__SharingModel(v cfz.Expression[string]) *AWS_QuickSight_Folder {
	t.SharingModel = v
	return t
}

// SetV__SharingModel updates property "SharingModel".
func (t *AWS_QuickSight_Folder) SetV__SharingModel(v string) *AWS_QuickSight_Folder {
	t.SharingModel = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_QuickSight_Folder) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_QuickSight_Folder {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_QuickSight_Folder) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_QuickSight_Folder {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_QuickSight_Folder) SetSV__Tags(v ...cfz.Tag) *AWS_QuickSight_Folder {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_QuickSight_Folder) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_QuickSight_Folder) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_QuickSight_Folder__AttributesMap.Arn))
}

// GetAtt__CreatedTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreatedTime
func (t *AWS_QuickSight_Folder) GetAtt__CreatedTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_QuickSight_Folder__AttributesMap.CreatedTime))
}

// GetAtt__LastUpdatedTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LastUpdatedTime
func (t *AWS_QuickSight_Folder) GetAtt__LastUpdatedTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_QuickSight_Folder__AttributesMap.LastUpdatedTime))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_QuickSight_Folder) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_QuickSight_Folder) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreatedTime returns a conventionally configured output for an attribute of this resource.
// Attribute: CreatedTime
func (t *AWS_QuickSight_Folder) GetConventionalOutputAtt__CreatedTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreatedTime", t.GetAtt__CreatedTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LastUpdatedTime returns a conventionally configured output for an attribute of this resource.
// Attribute: LastUpdatedTime
func (t *AWS_QuickSight_Folder) GetConventionalOutputAtt__LastUpdatedTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLastUpdatedTime", t.GetAtt__LastUpdatedTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_QuickSight_Folder) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_QuickSight_Folder

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

func (t *AWS_QuickSight_Folder) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
