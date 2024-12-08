// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dms

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_DMS_DataMigration)(nil)
	_ cfz.Resource                   = (*AWS_DMS_DataMigration)(nil)
)

const (
	// AWS_DMS_DataMigration__Type is the CloudFormation type for AWS::DMS::DataMigration.
	AWS_DMS_DataMigration__Type = "AWS::DMS::DataMigration"
)

var (
	// AWS_DMS_DataMigration__AttributesMap reports all the CloudFormation attributes for AWS::DMS::DataMigration.
	AWS_DMS_DataMigration__AttributesMap = struct {
		DataMigrationArn        string
		DataMigrationCreateTime string
	}{
		DataMigrationArn:        "DataMigrationArn",
		DataMigrationCreateTime: "DataMigrationCreateTime",
	}

	// AWS_DMS_DataMigration__AttributesSlice reports all the CloudFormation attributes for AWS::DMS::DataMigration.
	AWS_DMS_DataMigration__AttributesSlice = []string{
		AWS_DMS_DataMigration__AttributesMap.DataMigrationArn,
		AWS_DMS_DataMigration__AttributesMap.DataMigrationCreateTime,
	}
)

var (
	// AWS_DMS_DataMigration__PropertiesMap reports all the CloudFormation properties for AWS::DMS::DataMigration.
	AWS_DMS_DataMigration__PropertiesMap = struct {
		DataMigrationIdentifier    string
		DataMigrationName          string
		DataMigrationSettings      string
		DataMigrationType          string
		MigrationProjectIdentifier string
		ServiceAccessRoleArn       string
		SourceDataSettings         string
		Tags                       string
	}{
		DataMigrationIdentifier:    "DataMigrationIdentifier",
		DataMigrationName:          "DataMigrationName",
		DataMigrationSettings:      "DataMigrationSettings",
		DataMigrationType:          "DataMigrationType",
		MigrationProjectIdentifier: "MigrationProjectIdentifier",
		ServiceAccessRoleArn:       "ServiceAccessRoleArn",
		SourceDataSettings:         "SourceDataSettings",
		Tags:                       "Tags",
	}

	// AWS_DMS_DataMigration__PropertiesSlice reports all the CloudFormation properties for AWS::DMS::DataMigration.
	AWS_DMS_DataMigration__PropertiesSlice = []string{
		AWS_DMS_DataMigration__PropertiesMap.DataMigrationIdentifier,
		AWS_DMS_DataMigration__PropertiesMap.DataMigrationName,
		AWS_DMS_DataMigration__PropertiesMap.DataMigrationSettings,
		AWS_DMS_DataMigration__PropertiesMap.DataMigrationType,
		AWS_DMS_DataMigration__PropertiesMap.MigrationProjectIdentifier,
		AWS_DMS_DataMigration__PropertiesMap.ServiceAccessRoleArn,
		AWS_DMS_DataMigration__PropertiesMap.SourceDataSettings,
		AWS_DMS_DataMigration__PropertiesMap.Tags,
	}
)

// AWS_DMS_DataMigration is a binding for AWS::DMS::DataMigration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html
type AWS_DMS_DataMigration struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DataMigrationIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-datamigrationidentifier
	DataMigrationIdentifier cfz.Expression[string] `json:"DataMigrationIdentifier,omitempty"`

	// DataMigrationName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-datamigrationname
	DataMigrationName cfz.Expression[string] `json:"DataMigrationName,omitempty"`

	// DataMigrationSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-datamigrationsettings
	DataMigrationSettings cfz.Expression[AWS_DMS_DataMigration_DataMigrationSettings] `json:"DataMigrationSettings,omitempty"`

	// DataMigrationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-datamigrationtype
	DataMigrationType cfz.Expression[string] `json:"DataMigrationType,omitempty"`

	// MigrationProjectIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-migrationprojectidentifier
	MigrationProjectIdentifier cfz.Expression[string] `json:"MigrationProjectIdentifier,omitempty"`

	// ServiceAccessRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-serviceaccessrolearn
	ServiceAccessRoleArn cfz.Expression[string] `json:"ServiceAccessRoleArn,omitempty"`

	// SourceDataSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-sourcedatasettings
	SourceDataSettings cfz.ExpressionSlice[AWS_DMS_DataMigration_SourceDataSettings] `json:"SourceDataSettings,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_DMS_DataMigration initializes a new *AWS_DMS_DataMigration.
func New__AWS_DMS_DataMigration(logicalName string) *AWS_DMS_DataMigration {
	return &AWS_DMS_DataMigration{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_DMS_DataMigration) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_DMS_DataMigration) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_DMS_DataMigration) GetType() string {
	return AWS_DMS_DataMigration__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_DMS_DataMigration) Set__LogicalName(v string) *AWS_DMS_DataMigration {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_DMS_DataMigration) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_DMS_DataMigration {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_DMS_DataMigration) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_DMS_DataMigration {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_DMS_DataMigration) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_DMS_DataMigration {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_DMS_DataMigration) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_DMS_DataMigration {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_DMS_DataMigration) Set__RequestedOutputs(v []cfz.Output) *AWS_DMS_DataMigration {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_DMS_DataMigration) Add__RequestedOutputs(v ...cfz.Output) *AWS_DMS_DataMigration {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DataMigrationIdentifier updates property "DataMigrationIdentifier".
func (t *AWS_DMS_DataMigration) Set__DataMigrationIdentifier(v cfz.Expression[string]) *AWS_DMS_DataMigration {
	t.DataMigrationIdentifier = v
	return t
}

// SetV__DataMigrationIdentifier updates property "DataMigrationIdentifier".
func (t *AWS_DMS_DataMigration) SetV__DataMigrationIdentifier(v string) *AWS_DMS_DataMigration {
	t.DataMigrationIdentifier = cfz.V(v)
	return t
}

// Set__DataMigrationName updates property "DataMigrationName".
func (t *AWS_DMS_DataMigration) Set__DataMigrationName(v cfz.Expression[string]) *AWS_DMS_DataMigration {
	t.DataMigrationName = v
	return t
}

// SetV__DataMigrationName updates property "DataMigrationName".
func (t *AWS_DMS_DataMigration) SetV__DataMigrationName(v string) *AWS_DMS_DataMigration {
	t.DataMigrationName = cfz.V(v)
	return t
}

// Set__DataMigrationSettings updates property "DataMigrationSettings".
func (t *AWS_DMS_DataMigration) Set__DataMigrationSettings(v cfz.Expression[AWS_DMS_DataMigration_DataMigrationSettings]) *AWS_DMS_DataMigration {
	t.DataMigrationSettings = v
	return t
}

// SetV__DataMigrationSettings updates property "DataMigrationSettings".
func (t *AWS_DMS_DataMigration) SetV__DataMigrationSettings(v AWS_DMS_DataMigration_DataMigrationSettings) *AWS_DMS_DataMigration {
	t.DataMigrationSettings = cfz.V(v)
	return t
}

// Set__DataMigrationType updates property "DataMigrationType".
func (t *AWS_DMS_DataMigration) Set__DataMigrationType(v cfz.Expression[string]) *AWS_DMS_DataMigration {
	t.DataMigrationType = v
	return t
}

// SetV__DataMigrationType updates property "DataMigrationType".
func (t *AWS_DMS_DataMigration) SetV__DataMigrationType(v string) *AWS_DMS_DataMigration {
	t.DataMigrationType = cfz.V(v)
	return t
}

// Set__MigrationProjectIdentifier updates property "MigrationProjectIdentifier".
func (t *AWS_DMS_DataMigration) Set__MigrationProjectIdentifier(v cfz.Expression[string]) *AWS_DMS_DataMigration {
	t.MigrationProjectIdentifier = v
	return t
}

// SetV__MigrationProjectIdentifier updates property "MigrationProjectIdentifier".
func (t *AWS_DMS_DataMigration) SetV__MigrationProjectIdentifier(v string) *AWS_DMS_DataMigration {
	t.MigrationProjectIdentifier = cfz.V(v)
	return t
}

// Set__ServiceAccessRoleArn updates property "ServiceAccessRoleArn".
func (t *AWS_DMS_DataMigration) Set__ServiceAccessRoleArn(v cfz.Expression[string]) *AWS_DMS_DataMigration {
	t.ServiceAccessRoleArn = v
	return t
}

// SetV__ServiceAccessRoleArn updates property "ServiceAccessRoleArn".
func (t *AWS_DMS_DataMigration) SetV__ServiceAccessRoleArn(v string) *AWS_DMS_DataMigration {
	t.ServiceAccessRoleArn = cfz.V(v)
	return t
}

// Set__SourceDataSettings updates property "SourceDataSettings".
func (t *AWS_DMS_DataMigration) Set__SourceDataSettings(v cfz.ExpressionSlice[AWS_DMS_DataMigration_SourceDataSettings]) *AWS_DMS_DataMigration {
	t.SourceDataSettings = v
	return t
}

// SetS__SourceDataSettings updates property "SourceDataSettings".
func (t *AWS_DMS_DataMigration) SetS__SourceDataSettings(v ...cfz.Expression[AWS_DMS_DataMigration_SourceDataSettings]) *AWS_DMS_DataMigration {
	t.SourceDataSettings = cfz.S(v...)
	return t
}

// SetSV__SourceDataSettings updates property "SourceDataSettings".
func (t *AWS_DMS_DataMigration) SetSV__SourceDataSettings(v ...AWS_DMS_DataMigration_SourceDataSettings) *AWS_DMS_DataMigration {
	t.SourceDataSettings = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_DMS_DataMigration) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_DMS_DataMigration {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_DMS_DataMigration) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_DMS_DataMigration {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_DMS_DataMigration) SetSV__Tags(v ...cfz.Tag) *AWS_DMS_DataMigration {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_DMS_DataMigration) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__DataMigrationArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DataMigrationArn
func (t *AWS_DMS_DataMigration) GetAtt__DataMigrationArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DMS_DataMigration__AttributesMap.DataMigrationArn))
}

// GetAtt__DataMigrationCreateTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DataMigrationCreateTime
func (t *AWS_DMS_DataMigration) GetAtt__DataMigrationCreateTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DMS_DataMigration__AttributesMap.DataMigrationCreateTime))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_DMS_DataMigration) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DataMigrationArn returns a conventionally configured output for an attribute of this resource.
// Attribute: DataMigrationArn
func (t *AWS_DMS_DataMigration) GetConventionalOutputAtt__DataMigrationArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDataMigrationArn", t.GetAtt__DataMigrationArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DataMigrationCreateTime returns a conventionally configured output for an attribute of this resource.
// Attribute: DataMigrationCreateTime
func (t *AWS_DMS_DataMigration) GetConventionalOutputAtt__DataMigrationCreateTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDataMigrationCreateTime", t.GetAtt__DataMigrationCreateTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_DMS_DataMigration) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_DMS_DataMigration

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

func (t *AWS_DMS_DataMigration) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
