// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dms

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DMS_DataProvider_Settings__Type is the CloudFormation type for AWS::DMS::DataProvider.Settings.
	AWS_DMS_DataProvider_Settings__Type = "AWS::DMS::DataProvider.Settings"
)

var (
	// AWS_DMS_DataProvider_Settings__PropertiesMap reports all the CloudFormation properties for AWS::DMS::DataProvider.Settings.
	AWS_DMS_DataProvider_Settings__PropertiesMap = struct {
		MicrosoftSqlServerSettings string
		MySqlSettings              string
		OracleSettings             string
		PostgreSqlSettings         string
	}{
		MicrosoftSqlServerSettings: "MicrosoftSqlServerSettings",
		MySqlSettings:              "MySqlSettings",
		OracleSettings:             "OracleSettings",
		PostgreSqlSettings:         "PostgreSqlSettings",
	}

	// AWS_DMS_DataProvider_Settings__PropertiesSlice reports all the CloudFormation properties for AWS::DMS::DataProvider.Settings.
	AWS_DMS_DataProvider_Settings__PropertiesSlice = []string{
		AWS_DMS_DataProvider_Settings__PropertiesMap.MicrosoftSqlServerSettings,
		AWS_DMS_DataProvider_Settings__PropertiesMap.MySqlSettings,
		AWS_DMS_DataProvider_Settings__PropertiesMap.OracleSettings,
		AWS_DMS_DataProvider_Settings__PropertiesMap.PostgreSqlSettings,
	}
)

// AWS_DMS_DataProvider_Settings is a binding for AWS::DMS::DataProvider.Settings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html
type AWS_DMS_DataProvider_Settings struct {
	// MicrosoftSqlServerSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-microsoftsqlserversettings
	MicrosoftSqlServerSettings cfz.Expression[AWS_DMS_DataProvider_MicrosoftSqlServerSettings] `json:"MicrosoftSqlServerSettings,omitempty"`

	// MySqlSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-mysqlsettings
	MySqlSettings cfz.Expression[AWS_DMS_DataProvider_MySqlSettings] `json:"MySqlSettings,omitempty"`

	// OracleSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-oraclesettings
	OracleSettings cfz.Expression[AWS_DMS_DataProvider_OracleSettings] `json:"OracleSettings,omitempty"`

	// PostgreSqlSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-settings.html#cfn-dms-dataprovider-settings-postgresqlsettings
	PostgreSqlSettings cfz.Expression[AWS_DMS_DataProvider_PostgreSqlSettings] `json:"PostgreSqlSettings,omitempty"`
}

// New__AWS_DMS_DataProvider_Settings initializes a new AWS_DMS_DataProvider_Settings.
func New__AWS_DMS_DataProvider_Settings() AWS_DMS_DataProvider_Settings {
	return AWS_DMS_DataProvider_Settings{}
}

// GetType returns the CloudFormation type.
func (AWS_DMS_DataProvider_Settings) GetType() string {
	return AWS_DMS_DataProvider_Settings__Type
}

// Set__MicrosoftSqlServerSettings updates property "MicrosoftSqlServerSettings".
func (t AWS_DMS_DataProvider_Settings) Set__MicrosoftSqlServerSettings(v cfz.Expression[AWS_DMS_DataProvider_MicrosoftSqlServerSettings]) AWS_DMS_DataProvider_Settings {
	t.MicrosoftSqlServerSettings = v
	return t
}

// SetV__MicrosoftSqlServerSettings updates property "MicrosoftSqlServerSettings".
func (t AWS_DMS_DataProvider_Settings) SetV__MicrosoftSqlServerSettings(v AWS_DMS_DataProvider_MicrosoftSqlServerSettings) AWS_DMS_DataProvider_Settings {
	t.MicrosoftSqlServerSettings = cfz.V(v)
	return t
}

// Set__MySqlSettings updates property "MySqlSettings".
func (t AWS_DMS_DataProvider_Settings) Set__MySqlSettings(v cfz.Expression[AWS_DMS_DataProvider_MySqlSettings]) AWS_DMS_DataProvider_Settings {
	t.MySqlSettings = v
	return t
}

// SetV__MySqlSettings updates property "MySqlSettings".
func (t AWS_DMS_DataProvider_Settings) SetV__MySqlSettings(v AWS_DMS_DataProvider_MySqlSettings) AWS_DMS_DataProvider_Settings {
	t.MySqlSettings = cfz.V(v)
	return t
}

// Set__OracleSettings updates property "OracleSettings".
func (t AWS_DMS_DataProvider_Settings) Set__OracleSettings(v cfz.Expression[AWS_DMS_DataProvider_OracleSettings]) AWS_DMS_DataProvider_Settings {
	t.OracleSettings = v
	return t
}

// SetV__OracleSettings updates property "OracleSettings".
func (t AWS_DMS_DataProvider_Settings) SetV__OracleSettings(v AWS_DMS_DataProvider_OracleSettings) AWS_DMS_DataProvider_Settings {
	t.OracleSettings = cfz.V(v)
	return t
}

// Set__PostgreSqlSettings updates property "PostgreSqlSettings".
func (t AWS_DMS_DataProvider_Settings) Set__PostgreSqlSettings(v cfz.Expression[AWS_DMS_DataProvider_PostgreSqlSettings]) AWS_DMS_DataProvider_Settings {
	t.PostgreSqlSettings = v
	return t
}

// SetV__PostgreSqlSettings updates property "PostgreSqlSettings".
func (t AWS_DMS_DataProvider_Settings) SetV__PostgreSqlSettings(v AWS_DMS_DataProvider_PostgreSqlSettings) AWS_DMS_DataProvider_Settings {
	t.PostgreSqlSettings = cfz.V(v)
	return t
}
