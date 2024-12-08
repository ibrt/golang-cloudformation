// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_backup

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Backup_RestoreTestingSelection)(nil)
	_ cfz.Resource                   = (*AWS_Backup_RestoreTestingSelection)(nil)
)

const (
	// AWS_Backup_RestoreTestingSelection__Type is the CloudFormation type for AWS::Backup::RestoreTestingSelection.
	AWS_Backup_RestoreTestingSelection__Type = "AWS::Backup::RestoreTestingSelection"
)

var (
	// AWS_Backup_RestoreTestingSelection__PropertiesMap reports all the CloudFormation properties for AWS::Backup::RestoreTestingSelection.
	AWS_Backup_RestoreTestingSelection__PropertiesMap = struct {
		IamRoleArn                  string
		ProtectedResourceArns       string
		ProtectedResourceConditions string
		ProtectedResourceType       string
		RestoreMetadataOverrides    string
		RestoreTestingPlanName      string
		RestoreTestingSelectionName string
		ValidationWindowHours       string
	}{
		IamRoleArn:                  "IamRoleArn",
		ProtectedResourceArns:       "ProtectedResourceArns",
		ProtectedResourceConditions: "ProtectedResourceConditions",
		ProtectedResourceType:       "ProtectedResourceType",
		RestoreMetadataOverrides:    "RestoreMetadataOverrides",
		RestoreTestingPlanName:      "RestoreTestingPlanName",
		RestoreTestingSelectionName: "RestoreTestingSelectionName",
		ValidationWindowHours:       "ValidationWindowHours",
	}

	// AWS_Backup_RestoreTestingSelection__PropertiesSlice reports all the CloudFormation properties for AWS::Backup::RestoreTestingSelection.
	AWS_Backup_RestoreTestingSelection__PropertiesSlice = []string{
		AWS_Backup_RestoreTestingSelection__PropertiesMap.IamRoleArn,
		AWS_Backup_RestoreTestingSelection__PropertiesMap.ProtectedResourceArns,
		AWS_Backup_RestoreTestingSelection__PropertiesMap.ProtectedResourceConditions,
		AWS_Backup_RestoreTestingSelection__PropertiesMap.ProtectedResourceType,
		AWS_Backup_RestoreTestingSelection__PropertiesMap.RestoreMetadataOverrides,
		AWS_Backup_RestoreTestingSelection__PropertiesMap.RestoreTestingPlanName,
		AWS_Backup_RestoreTestingSelection__PropertiesMap.RestoreTestingSelectionName,
		AWS_Backup_RestoreTestingSelection__PropertiesMap.ValidationWindowHours,
	}
)

// AWS_Backup_RestoreTestingSelection is a binding for AWS::Backup::RestoreTestingSelection.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html
type AWS_Backup_RestoreTestingSelection struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// IamRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-iamrolearn
	IamRoleArn cfz.Expression[string] `json:"IamRoleArn,omitempty"`

	// ProtectedResourceArns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-protectedresourcearns
	ProtectedResourceArns cfz.ExpressionSlice[string] `json:"ProtectedResourceArns,omitempty"`

	// ProtectedResourceConditions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-protectedresourceconditions
	ProtectedResourceConditions cfz.Expression[AWS_Backup_RestoreTestingSelection_ProtectedResourceConditions] `json:"ProtectedResourceConditions,omitempty"`

	// ProtectedResourceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-protectedresourcetype
	ProtectedResourceType cfz.Expression[string] `json:"ProtectedResourceType,omitempty"`

	// RestoreMetadataOverrides is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-restoremetadataoverrides
	RestoreMetadataOverrides cfz.ExpressionMap[string] `json:"RestoreMetadataOverrides,omitempty"`

	// RestoreTestingPlanName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-restoretestingplanname
	RestoreTestingPlanName cfz.Expression[string] `json:"RestoreTestingPlanName,omitempty"`

	// RestoreTestingSelectionName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-restoretestingselectionname
	RestoreTestingSelectionName cfz.Expression[string] `json:"RestoreTestingSelectionName,omitempty"`

	// ValidationWindowHours is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingselection.html#cfn-backup-restoretestingselection-validationwindowhours
	ValidationWindowHours cfz.Expression[int32] `json:"ValidationWindowHours,omitempty"`
}

// New__AWS_Backup_RestoreTestingSelection initializes a new *AWS_Backup_RestoreTestingSelection.
func New__AWS_Backup_RestoreTestingSelection(logicalName string) *AWS_Backup_RestoreTestingSelection {
	return &AWS_Backup_RestoreTestingSelection{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Backup_RestoreTestingSelection) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Backup_RestoreTestingSelection) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Backup_RestoreTestingSelection) GetType() string {
	return AWS_Backup_RestoreTestingSelection__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Backup_RestoreTestingSelection) Set__LogicalName(v string) *AWS_Backup_RestoreTestingSelection {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Backup_RestoreTestingSelection) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Backup_RestoreTestingSelection {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Backup_RestoreTestingSelection) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Backup_RestoreTestingSelection {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Backup_RestoreTestingSelection) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Backup_RestoreTestingSelection {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Backup_RestoreTestingSelection) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Backup_RestoreTestingSelection {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Backup_RestoreTestingSelection) Set__RequestedOutputs(v []cfz.Output) *AWS_Backup_RestoreTestingSelection {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Backup_RestoreTestingSelection) Add__RequestedOutputs(v ...cfz.Output) *AWS_Backup_RestoreTestingSelection {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__IamRoleArn updates property "IamRoleArn".
func (t *AWS_Backup_RestoreTestingSelection) Set__IamRoleArn(v cfz.Expression[string]) *AWS_Backup_RestoreTestingSelection {
	t.IamRoleArn = v
	return t
}

// SetV__IamRoleArn updates property "IamRoleArn".
func (t *AWS_Backup_RestoreTestingSelection) SetV__IamRoleArn(v string) *AWS_Backup_RestoreTestingSelection {
	t.IamRoleArn = cfz.V(v)
	return t
}

// Set__ProtectedResourceArns updates property "ProtectedResourceArns".
func (t *AWS_Backup_RestoreTestingSelection) Set__ProtectedResourceArns(v cfz.ExpressionSlice[string]) *AWS_Backup_RestoreTestingSelection {
	t.ProtectedResourceArns = v
	return t
}

// SetS__ProtectedResourceArns updates property "ProtectedResourceArns".
func (t *AWS_Backup_RestoreTestingSelection) SetS__ProtectedResourceArns(v ...cfz.Expression[string]) *AWS_Backup_RestoreTestingSelection {
	t.ProtectedResourceArns = cfz.S(v...)
	return t
}

// SetSV__ProtectedResourceArns updates property "ProtectedResourceArns".
func (t *AWS_Backup_RestoreTestingSelection) SetSV__ProtectedResourceArns(v ...string) *AWS_Backup_RestoreTestingSelection {
	t.ProtectedResourceArns = cfz.SV(v...)
	return t
}

// Set__ProtectedResourceConditions updates property "ProtectedResourceConditions".
func (t *AWS_Backup_RestoreTestingSelection) Set__ProtectedResourceConditions(v cfz.Expression[AWS_Backup_RestoreTestingSelection_ProtectedResourceConditions]) *AWS_Backup_RestoreTestingSelection {
	t.ProtectedResourceConditions = v
	return t
}

// SetV__ProtectedResourceConditions updates property "ProtectedResourceConditions".
func (t *AWS_Backup_RestoreTestingSelection) SetV__ProtectedResourceConditions(v AWS_Backup_RestoreTestingSelection_ProtectedResourceConditions) *AWS_Backup_RestoreTestingSelection {
	t.ProtectedResourceConditions = cfz.V(v)
	return t
}

// Set__ProtectedResourceType updates property "ProtectedResourceType".
func (t *AWS_Backup_RestoreTestingSelection) Set__ProtectedResourceType(v cfz.Expression[string]) *AWS_Backup_RestoreTestingSelection {
	t.ProtectedResourceType = v
	return t
}

// SetV__ProtectedResourceType updates property "ProtectedResourceType".
func (t *AWS_Backup_RestoreTestingSelection) SetV__ProtectedResourceType(v string) *AWS_Backup_RestoreTestingSelection {
	t.ProtectedResourceType = cfz.V(v)
	return t
}

// Set__RestoreMetadataOverrides updates property "RestoreMetadataOverrides".
func (t *AWS_Backup_RestoreTestingSelection) Set__RestoreMetadataOverrides(v cfz.ExpressionMap[string]) *AWS_Backup_RestoreTestingSelection {
	t.RestoreMetadataOverrides = v
	return t
}

// SetM__RestoreMetadataOverrides updates property "RestoreMetadataOverrides".
func (t *AWS_Backup_RestoreTestingSelection) SetM__RestoreMetadataOverrides(v ...map[string]cfz.Expression[string]) *AWS_Backup_RestoreTestingSelection {
	t.RestoreMetadataOverrides = cfz.M(v...)
	return t
}

// SetMV__RestoreMetadataOverrides updates property "RestoreMetadataOverrides".
func (t *AWS_Backup_RestoreTestingSelection) SetMV__RestoreMetadataOverrides(v ...map[string]string) *AWS_Backup_RestoreTestingSelection {
	t.RestoreMetadataOverrides = cfz.MV(v...)
	return t
}

// Set__RestoreTestingPlanName updates property "RestoreTestingPlanName".
func (t *AWS_Backup_RestoreTestingSelection) Set__RestoreTestingPlanName(v cfz.Expression[string]) *AWS_Backup_RestoreTestingSelection {
	t.RestoreTestingPlanName = v
	return t
}

// SetV__RestoreTestingPlanName updates property "RestoreTestingPlanName".
func (t *AWS_Backup_RestoreTestingSelection) SetV__RestoreTestingPlanName(v string) *AWS_Backup_RestoreTestingSelection {
	t.RestoreTestingPlanName = cfz.V(v)
	return t
}

// Set__RestoreTestingSelectionName updates property "RestoreTestingSelectionName".
func (t *AWS_Backup_RestoreTestingSelection) Set__RestoreTestingSelectionName(v cfz.Expression[string]) *AWS_Backup_RestoreTestingSelection {
	t.RestoreTestingSelectionName = v
	return t
}

// SetV__RestoreTestingSelectionName updates property "RestoreTestingSelectionName".
func (t *AWS_Backup_RestoreTestingSelection) SetV__RestoreTestingSelectionName(v string) *AWS_Backup_RestoreTestingSelection {
	t.RestoreTestingSelectionName = cfz.V(v)
	return t
}

// Set__ValidationWindowHours updates property "ValidationWindowHours".
func (t *AWS_Backup_RestoreTestingSelection) Set__ValidationWindowHours(v cfz.Expression[int32]) *AWS_Backup_RestoreTestingSelection {
	t.ValidationWindowHours = v
	return t
}

// SetV__ValidationWindowHours updates property "ValidationWindowHours".
func (t *AWS_Backup_RestoreTestingSelection) SetV__ValidationWindowHours(v int32) *AWS_Backup_RestoreTestingSelection {
	t.ValidationWindowHours = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Backup_RestoreTestingSelection) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Backup_RestoreTestingSelection) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Backup_RestoreTestingSelection) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Backup_RestoreTestingSelection

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

func (t *AWS_Backup_RestoreTestingSelection) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
