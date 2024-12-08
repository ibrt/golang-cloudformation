// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lightsail

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Lightsail_Database)(nil)
	_ cfz.Resource                   = (*AWS_Lightsail_Database)(nil)
)

const (
	// AWS_Lightsail_Database__Type is the CloudFormation type for AWS::Lightsail::Database.
	AWS_Lightsail_Database__Type = "AWS::Lightsail::Database"
)

var (
	// AWS_Lightsail_Database__AttributesMap reports all the CloudFormation attributes for AWS::Lightsail::Database.
	AWS_Lightsail_Database__AttributesMap = struct {
		DatabaseArn string
	}{
		DatabaseArn: "DatabaseArn",
	}

	// AWS_Lightsail_Database__AttributesSlice reports all the CloudFormation attributes for AWS::Lightsail::Database.
	AWS_Lightsail_Database__AttributesSlice = []string{
		AWS_Lightsail_Database__AttributesMap.DatabaseArn,
	}
)

var (
	// AWS_Lightsail_Database__PropertiesMap reports all the CloudFormation properties for AWS::Lightsail::Database.
	AWS_Lightsail_Database__PropertiesMap = struct {
		AvailabilityZone              string
		BackupRetention               string
		CaCertificateIdentifier       string
		MasterDatabaseName            string
		MasterUserPassword            string
		MasterUsername                string
		PreferredBackupWindow         string
		PreferredMaintenanceWindow    string
		PubliclyAccessible            string
		RelationalDatabaseBlueprintId string
		RelationalDatabaseBundleId    string
		RelationalDatabaseName        string
		RelationalDatabaseParameters  string
		RotateMasterUserPassword      string
		Tags                          string
	}{
		AvailabilityZone:              "AvailabilityZone",
		BackupRetention:               "BackupRetention",
		CaCertificateIdentifier:       "CaCertificateIdentifier",
		MasterDatabaseName:            "MasterDatabaseName",
		MasterUserPassword:            "MasterUserPassword",
		MasterUsername:                "MasterUsername",
		PreferredBackupWindow:         "PreferredBackupWindow",
		PreferredMaintenanceWindow:    "PreferredMaintenanceWindow",
		PubliclyAccessible:            "PubliclyAccessible",
		RelationalDatabaseBlueprintId: "RelationalDatabaseBlueprintId",
		RelationalDatabaseBundleId:    "RelationalDatabaseBundleId",
		RelationalDatabaseName:        "RelationalDatabaseName",
		RelationalDatabaseParameters:  "RelationalDatabaseParameters",
		RotateMasterUserPassword:      "RotateMasterUserPassword",
		Tags:                          "Tags",
	}

	// AWS_Lightsail_Database__PropertiesSlice reports all the CloudFormation properties for AWS::Lightsail::Database.
	AWS_Lightsail_Database__PropertiesSlice = []string{
		AWS_Lightsail_Database__PropertiesMap.AvailabilityZone,
		AWS_Lightsail_Database__PropertiesMap.BackupRetention,
		AWS_Lightsail_Database__PropertiesMap.CaCertificateIdentifier,
		AWS_Lightsail_Database__PropertiesMap.MasterDatabaseName,
		AWS_Lightsail_Database__PropertiesMap.MasterUserPassword,
		AWS_Lightsail_Database__PropertiesMap.MasterUsername,
		AWS_Lightsail_Database__PropertiesMap.PreferredBackupWindow,
		AWS_Lightsail_Database__PropertiesMap.PreferredMaintenanceWindow,
		AWS_Lightsail_Database__PropertiesMap.PubliclyAccessible,
		AWS_Lightsail_Database__PropertiesMap.RelationalDatabaseBlueprintId,
		AWS_Lightsail_Database__PropertiesMap.RelationalDatabaseBundleId,
		AWS_Lightsail_Database__PropertiesMap.RelationalDatabaseName,
		AWS_Lightsail_Database__PropertiesMap.RelationalDatabaseParameters,
		AWS_Lightsail_Database__PropertiesMap.RotateMasterUserPassword,
		AWS_Lightsail_Database__PropertiesMap.Tags,
	}
)

// AWS_Lightsail_Database is a binding for AWS::Lightsail::Database.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html
type AWS_Lightsail_Database struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AvailabilityZone is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-availabilityzone
	AvailabilityZone cfz.Expression[string] `json:"AvailabilityZone,omitempty"`

	// BackupRetention is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-backupretention
	BackupRetention cfz.Expression[bool] `json:"BackupRetention,omitempty"`

	// CaCertificateIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-cacertificateidentifier
	CaCertificateIdentifier cfz.Expression[string] `json:"CaCertificateIdentifier,omitempty"`

	// MasterDatabaseName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-masterdatabasename
	MasterDatabaseName cfz.Expression[string] `json:"MasterDatabaseName,omitempty"`

	// MasterUserPassword is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-masteruserpassword
	MasterUserPassword cfz.Expression[string] `json:"MasterUserPassword,omitempty"`

	// MasterUsername is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-masterusername
	MasterUsername cfz.Expression[string] `json:"MasterUsername,omitempty"`

	// PreferredBackupWindow is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-preferredbackupwindow
	PreferredBackupWindow cfz.Expression[string] `json:"PreferredBackupWindow,omitempty"`

	// PreferredMaintenanceWindow is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-preferredmaintenancewindow
	PreferredMaintenanceWindow cfz.Expression[string] `json:"PreferredMaintenanceWindow,omitempty"`

	// PubliclyAccessible is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-publiclyaccessible
	PubliclyAccessible cfz.Expression[bool] `json:"PubliclyAccessible,omitempty"`

	// RelationalDatabaseBlueprintId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-relationaldatabaseblueprintid
	RelationalDatabaseBlueprintId cfz.Expression[string] `json:"RelationalDatabaseBlueprintId,omitempty"`

	// RelationalDatabaseBundleId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-relationaldatabasebundleid
	RelationalDatabaseBundleId cfz.Expression[string] `json:"RelationalDatabaseBundleId,omitempty"`

	// RelationalDatabaseName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-relationaldatabasename
	RelationalDatabaseName cfz.Expression[string] `json:"RelationalDatabaseName,omitempty"`

	// RelationalDatabaseParameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-relationaldatabaseparameters
	RelationalDatabaseParameters cfz.ExpressionSlice[AWS_Lightsail_Database_RelationalDatabaseParameter] `json:"RelationalDatabaseParameters,omitempty"`

	// RotateMasterUserPassword is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-rotatemasteruserpassword
	RotateMasterUserPassword cfz.Expression[bool] `json:"RotateMasterUserPassword,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_Lightsail_Database initializes a new *AWS_Lightsail_Database.
func New__AWS_Lightsail_Database(logicalName string) *AWS_Lightsail_Database {
	return &AWS_Lightsail_Database{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Lightsail_Database) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Lightsail_Database) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Lightsail_Database) GetType() string {
	return AWS_Lightsail_Database__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Lightsail_Database) Set__LogicalName(v string) *AWS_Lightsail_Database {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Lightsail_Database) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Lightsail_Database {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Lightsail_Database) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Lightsail_Database {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Lightsail_Database) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Lightsail_Database {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Lightsail_Database) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Lightsail_Database {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Lightsail_Database) Set__RequestedOutputs(v []cfz.Output) *AWS_Lightsail_Database {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Lightsail_Database) Add__RequestedOutputs(v ...cfz.Output) *AWS_Lightsail_Database {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AvailabilityZone updates property "AvailabilityZone".
func (t *AWS_Lightsail_Database) Set__AvailabilityZone(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.AvailabilityZone = v
	return t
}

// SetV__AvailabilityZone updates property "AvailabilityZone".
func (t *AWS_Lightsail_Database) SetV__AvailabilityZone(v string) *AWS_Lightsail_Database {
	t.AvailabilityZone = cfz.V(v)
	return t
}

// Set__BackupRetention updates property "BackupRetention".
func (t *AWS_Lightsail_Database) Set__BackupRetention(v cfz.Expression[bool]) *AWS_Lightsail_Database {
	t.BackupRetention = v
	return t
}

// SetV__BackupRetention updates property "BackupRetention".
func (t *AWS_Lightsail_Database) SetV__BackupRetention(v bool) *AWS_Lightsail_Database {
	t.BackupRetention = cfz.V(v)
	return t
}

// Set__CaCertificateIdentifier updates property "CaCertificateIdentifier".
func (t *AWS_Lightsail_Database) Set__CaCertificateIdentifier(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.CaCertificateIdentifier = v
	return t
}

// SetV__CaCertificateIdentifier updates property "CaCertificateIdentifier".
func (t *AWS_Lightsail_Database) SetV__CaCertificateIdentifier(v string) *AWS_Lightsail_Database {
	t.CaCertificateIdentifier = cfz.V(v)
	return t
}

// Set__MasterDatabaseName updates property "MasterDatabaseName".
func (t *AWS_Lightsail_Database) Set__MasterDatabaseName(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.MasterDatabaseName = v
	return t
}

// SetV__MasterDatabaseName updates property "MasterDatabaseName".
func (t *AWS_Lightsail_Database) SetV__MasterDatabaseName(v string) *AWS_Lightsail_Database {
	t.MasterDatabaseName = cfz.V(v)
	return t
}

// Set__MasterUserPassword updates property "MasterUserPassword".
func (t *AWS_Lightsail_Database) Set__MasterUserPassword(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.MasterUserPassword = v
	return t
}

// SetV__MasterUserPassword updates property "MasterUserPassword".
func (t *AWS_Lightsail_Database) SetV__MasterUserPassword(v string) *AWS_Lightsail_Database {
	t.MasterUserPassword = cfz.V(v)
	return t
}

// Set__MasterUsername updates property "MasterUsername".
func (t *AWS_Lightsail_Database) Set__MasterUsername(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.MasterUsername = v
	return t
}

// SetV__MasterUsername updates property "MasterUsername".
func (t *AWS_Lightsail_Database) SetV__MasterUsername(v string) *AWS_Lightsail_Database {
	t.MasterUsername = cfz.V(v)
	return t
}

// Set__PreferredBackupWindow updates property "PreferredBackupWindow".
func (t *AWS_Lightsail_Database) Set__PreferredBackupWindow(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.PreferredBackupWindow = v
	return t
}

// SetV__PreferredBackupWindow updates property "PreferredBackupWindow".
func (t *AWS_Lightsail_Database) SetV__PreferredBackupWindow(v string) *AWS_Lightsail_Database {
	t.PreferredBackupWindow = cfz.V(v)
	return t
}

// Set__PreferredMaintenanceWindow updates property "PreferredMaintenanceWindow".
func (t *AWS_Lightsail_Database) Set__PreferredMaintenanceWindow(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.PreferredMaintenanceWindow = v
	return t
}

// SetV__PreferredMaintenanceWindow updates property "PreferredMaintenanceWindow".
func (t *AWS_Lightsail_Database) SetV__PreferredMaintenanceWindow(v string) *AWS_Lightsail_Database {
	t.PreferredMaintenanceWindow = cfz.V(v)
	return t
}

// Set__PubliclyAccessible updates property "PubliclyAccessible".
func (t *AWS_Lightsail_Database) Set__PubliclyAccessible(v cfz.Expression[bool]) *AWS_Lightsail_Database {
	t.PubliclyAccessible = v
	return t
}

// SetV__PubliclyAccessible updates property "PubliclyAccessible".
func (t *AWS_Lightsail_Database) SetV__PubliclyAccessible(v bool) *AWS_Lightsail_Database {
	t.PubliclyAccessible = cfz.V(v)
	return t
}

// Set__RelationalDatabaseBlueprintId updates property "RelationalDatabaseBlueprintId".
func (t *AWS_Lightsail_Database) Set__RelationalDatabaseBlueprintId(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.RelationalDatabaseBlueprintId = v
	return t
}

// SetV__RelationalDatabaseBlueprintId updates property "RelationalDatabaseBlueprintId".
func (t *AWS_Lightsail_Database) SetV__RelationalDatabaseBlueprintId(v string) *AWS_Lightsail_Database {
	t.RelationalDatabaseBlueprintId = cfz.V(v)
	return t
}

// Set__RelationalDatabaseBundleId updates property "RelationalDatabaseBundleId".
func (t *AWS_Lightsail_Database) Set__RelationalDatabaseBundleId(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.RelationalDatabaseBundleId = v
	return t
}

// SetV__RelationalDatabaseBundleId updates property "RelationalDatabaseBundleId".
func (t *AWS_Lightsail_Database) SetV__RelationalDatabaseBundleId(v string) *AWS_Lightsail_Database {
	t.RelationalDatabaseBundleId = cfz.V(v)
	return t
}

// Set__RelationalDatabaseName updates property "RelationalDatabaseName".
func (t *AWS_Lightsail_Database) Set__RelationalDatabaseName(v cfz.Expression[string]) *AWS_Lightsail_Database {
	t.RelationalDatabaseName = v
	return t
}

// SetV__RelationalDatabaseName updates property "RelationalDatabaseName".
func (t *AWS_Lightsail_Database) SetV__RelationalDatabaseName(v string) *AWS_Lightsail_Database {
	t.RelationalDatabaseName = cfz.V(v)
	return t
}

// Set__RelationalDatabaseParameters updates property "RelationalDatabaseParameters".
func (t *AWS_Lightsail_Database) Set__RelationalDatabaseParameters(v cfz.ExpressionSlice[AWS_Lightsail_Database_RelationalDatabaseParameter]) *AWS_Lightsail_Database {
	t.RelationalDatabaseParameters = v
	return t
}

// SetS__RelationalDatabaseParameters updates property "RelationalDatabaseParameters".
func (t *AWS_Lightsail_Database) SetS__RelationalDatabaseParameters(v ...cfz.Expression[AWS_Lightsail_Database_RelationalDatabaseParameter]) *AWS_Lightsail_Database {
	t.RelationalDatabaseParameters = cfz.S(v...)
	return t
}

// SetSV__RelationalDatabaseParameters updates property "RelationalDatabaseParameters".
func (t *AWS_Lightsail_Database) SetSV__RelationalDatabaseParameters(v ...AWS_Lightsail_Database_RelationalDatabaseParameter) *AWS_Lightsail_Database {
	t.RelationalDatabaseParameters = cfz.SV(v...)
	return t
}

// Set__RotateMasterUserPassword updates property "RotateMasterUserPassword".
func (t *AWS_Lightsail_Database) Set__RotateMasterUserPassword(v cfz.Expression[bool]) *AWS_Lightsail_Database {
	t.RotateMasterUserPassword = v
	return t
}

// SetV__RotateMasterUserPassword updates property "RotateMasterUserPassword".
func (t *AWS_Lightsail_Database) SetV__RotateMasterUserPassword(v bool) *AWS_Lightsail_Database {
	t.RotateMasterUserPassword = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Lightsail_Database) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Lightsail_Database {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Lightsail_Database) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Lightsail_Database {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Lightsail_Database) SetSV__Tags(v ...cfz.Tag) *AWS_Lightsail_Database {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Lightsail_Database) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__DatabaseArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DatabaseArn
func (t *AWS_Lightsail_Database) GetAtt__DatabaseArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Lightsail_Database__AttributesMap.DatabaseArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Lightsail_Database) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DatabaseArn returns a conventionally configured output for an attribute of this resource.
// Attribute: DatabaseArn
func (t *AWS_Lightsail_Database) GetConventionalOutputAtt__DatabaseArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDatabaseArn", t.GetAtt__DatabaseArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Lightsail_Database) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Lightsail_Database

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

func (t *AWS_Lightsail_Database) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
