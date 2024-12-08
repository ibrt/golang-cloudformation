// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kendra

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Kendra_DataSource_OneDriveConfiguration__Type is the CloudFormation type for AWS::Kendra::DataSource.OneDriveConfiguration.
	AWS_Kendra_DataSource_OneDriveConfiguration__Type = "AWS::Kendra::DataSource.OneDriveConfiguration"
)

var (
	// AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Kendra::DataSource.OneDriveConfiguration.
	AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesMap = struct {
		DisableLocalGroups string
		ExclusionPatterns  string
		FieldMappings      string
		InclusionPatterns  string
		OneDriveUsers      string
		SecretArn          string
		TenantDomain       string
	}{
		DisableLocalGroups: "DisableLocalGroups",
		ExclusionPatterns:  "ExclusionPatterns",
		FieldMappings:      "FieldMappings",
		InclusionPatterns:  "InclusionPatterns",
		OneDriveUsers:      "OneDriveUsers",
		SecretArn:          "SecretArn",
		TenantDomain:       "TenantDomain",
	}

	// AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Kendra::DataSource.OneDriveConfiguration.
	AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesSlice = []string{
		AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesMap.DisableLocalGroups,
		AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesMap.ExclusionPatterns,
		AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesMap.FieldMappings,
		AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesMap.InclusionPatterns,
		AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesMap.OneDriveUsers,
		AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesMap.SecretArn,
		AWS_Kendra_DataSource_OneDriveConfiguration__PropertiesMap.TenantDomain,
	}
)

// AWS_Kendra_DataSource_OneDriveConfiguration is a binding for AWS::Kendra::DataSource.OneDriveConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html
type AWS_Kendra_DataSource_OneDriveConfiguration struct {
	// DisableLocalGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-disablelocalgroups
	DisableLocalGroups cfz.Expression[bool] `json:"DisableLocalGroups,omitempty"`

	// ExclusionPatterns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-exclusionpatterns
	ExclusionPatterns cfz.ExpressionSlice[string] `json:"ExclusionPatterns,omitempty"`

	// FieldMappings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-fieldmappings
	FieldMappings cfz.ExpressionSlice[AWS_Kendra_DataSource_DataSourceToIndexFieldMapping] `json:"FieldMappings,omitempty"`

	// InclusionPatterns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-inclusionpatterns
	InclusionPatterns cfz.ExpressionSlice[string] `json:"InclusionPatterns,omitempty"`

	// OneDriveUsers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-onedriveusers
	OneDriveUsers cfz.Expression[AWS_Kendra_DataSource_OneDriveUsers] `json:"OneDriveUsers,omitempty"`

	// SecretArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-secretarn
	SecretArn cfz.Expression[string] `json:"SecretArn,omitempty"`

	// TenantDomain is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-tenantdomain
	TenantDomain cfz.Expression[string] `json:"TenantDomain,omitempty"`
}

// New__AWS_Kendra_DataSource_OneDriveConfiguration initializes a new AWS_Kendra_DataSource_OneDriveConfiguration.
func New__AWS_Kendra_DataSource_OneDriveConfiguration() AWS_Kendra_DataSource_OneDriveConfiguration {
	return AWS_Kendra_DataSource_OneDriveConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Kendra_DataSource_OneDriveConfiguration) GetType() string {
	return AWS_Kendra_DataSource_OneDriveConfiguration__Type
}

// Set__DisableLocalGroups updates property "DisableLocalGroups".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) Set__DisableLocalGroups(v cfz.Expression[bool]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.DisableLocalGroups = v
	return t
}

// SetV__DisableLocalGroups updates property "DisableLocalGroups".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetV__DisableLocalGroups(v bool) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.DisableLocalGroups = cfz.V(v)
	return t
}

// Set__ExclusionPatterns updates property "ExclusionPatterns".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) Set__ExclusionPatterns(v cfz.ExpressionSlice[string]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.ExclusionPatterns = v
	return t
}

// SetS__ExclusionPatterns updates property "ExclusionPatterns".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetS__ExclusionPatterns(v ...cfz.Expression[string]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.ExclusionPatterns = cfz.S(v...)
	return t
}

// SetSV__ExclusionPatterns updates property "ExclusionPatterns".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetSV__ExclusionPatterns(v ...string) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.ExclusionPatterns = cfz.SV(v...)
	return t
}

// Set__FieldMappings updates property "FieldMappings".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) Set__FieldMappings(v cfz.ExpressionSlice[AWS_Kendra_DataSource_DataSourceToIndexFieldMapping]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.FieldMappings = v
	return t
}

// SetS__FieldMappings updates property "FieldMappings".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetS__FieldMappings(v ...cfz.Expression[AWS_Kendra_DataSource_DataSourceToIndexFieldMapping]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.FieldMappings = cfz.S(v...)
	return t
}

// SetSV__FieldMappings updates property "FieldMappings".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetSV__FieldMappings(v ...AWS_Kendra_DataSource_DataSourceToIndexFieldMapping) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.FieldMappings = cfz.SV(v...)
	return t
}

// Set__InclusionPatterns updates property "InclusionPatterns".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) Set__InclusionPatterns(v cfz.ExpressionSlice[string]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.InclusionPatterns = v
	return t
}

// SetS__InclusionPatterns updates property "InclusionPatterns".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetS__InclusionPatterns(v ...cfz.Expression[string]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.InclusionPatterns = cfz.S(v...)
	return t
}

// SetSV__InclusionPatterns updates property "InclusionPatterns".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetSV__InclusionPatterns(v ...string) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.InclusionPatterns = cfz.SV(v...)
	return t
}

// Set__OneDriveUsers updates property "OneDriveUsers".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) Set__OneDriveUsers(v cfz.Expression[AWS_Kendra_DataSource_OneDriveUsers]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.OneDriveUsers = v
	return t
}

// SetV__OneDriveUsers updates property "OneDriveUsers".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetV__OneDriveUsers(v AWS_Kendra_DataSource_OneDriveUsers) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.OneDriveUsers = cfz.V(v)
	return t
}

// Set__SecretArn updates property "SecretArn".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) Set__SecretArn(v cfz.Expression[string]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.SecretArn = v
	return t
}

// SetV__SecretArn updates property "SecretArn".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetV__SecretArn(v string) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.SecretArn = cfz.V(v)
	return t
}

// Set__TenantDomain updates property "TenantDomain".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) Set__TenantDomain(v cfz.Expression[string]) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.TenantDomain = v
	return t
}

// SetV__TenantDomain updates property "TenantDomain".
func (t AWS_Kendra_DataSource_OneDriveConfiguration) SetV__TenantDomain(v string) AWS_Kendra_DataSource_OneDriveConfiguration {
	t.TenantDomain = cfz.V(v)
	return t
}
