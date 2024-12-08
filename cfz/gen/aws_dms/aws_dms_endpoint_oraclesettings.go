// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dms

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DMS_Endpoint_OracleSettings__Type is the CloudFormation type for AWS::DMS::Endpoint.OracleSettings.
	AWS_DMS_Endpoint_OracleSettings__Type = "AWS::DMS::Endpoint.OracleSettings"
)

var (
	// AWS_DMS_Endpoint_OracleSettings__PropertiesMap reports all the CloudFormation properties for AWS::DMS::Endpoint.OracleSettings.
	AWS_DMS_Endpoint_OracleSettings__PropertiesMap = struct {
		AccessAlternateDirectly                string
		AddSupplementalLogging                 string
		AdditionalArchivedLogDestId            string
		AllowSelectNestedTables                string
		ArchivedLogDestId                      string
		ArchivedLogsOnly                       string
		AsmPassword                            string
		AsmServer                              string
		AsmUser                                string
		CharLengthSemantics                    string
		DirectPathNoLog                        string
		DirectPathParallelLoad                 string
		EnableHomogenousTablespace             string
		ExtraArchivedLogDestIds                string
		FailTasksOnLobTruncation               string
		NumberDatatypeScale                    string
		OraclePathPrefix                       string
		ParallelAsmReadThreads                 string
		ReadAheadBlocks                        string
		ReadTableSpaceName                     string
		ReplacePathPrefix                      string
		RetryInterval                          string
		SecretsManagerAccessRoleArn            string
		SecretsManagerOracleAsmAccessRoleArn   string
		SecretsManagerOracleAsmSecretId        string
		SecretsManagerSecretId                 string
		SecurityDbEncryption                   string
		SecurityDbEncryptionName               string
		SpatialDataOptionToGeoJsonFunctionName string
		StandbyDelayTime                       string
		UseAlternateFolderForOnline            string
		UseBFile                               string
		UseDirectPathFullLoad                  string
		UseLogminerReader                      string
		UsePathPrefix                          string
	}{
		AccessAlternateDirectly:                "AccessAlternateDirectly",
		AddSupplementalLogging:                 "AddSupplementalLogging",
		AdditionalArchivedLogDestId:            "AdditionalArchivedLogDestId",
		AllowSelectNestedTables:                "AllowSelectNestedTables",
		ArchivedLogDestId:                      "ArchivedLogDestId",
		ArchivedLogsOnly:                       "ArchivedLogsOnly",
		AsmPassword:                            "AsmPassword",
		AsmServer:                              "AsmServer",
		AsmUser:                                "AsmUser",
		CharLengthSemantics:                    "CharLengthSemantics",
		DirectPathNoLog:                        "DirectPathNoLog",
		DirectPathParallelLoad:                 "DirectPathParallelLoad",
		EnableHomogenousTablespace:             "EnableHomogenousTablespace",
		ExtraArchivedLogDestIds:                "ExtraArchivedLogDestIds",
		FailTasksOnLobTruncation:               "FailTasksOnLobTruncation",
		NumberDatatypeScale:                    "NumberDatatypeScale",
		OraclePathPrefix:                       "OraclePathPrefix",
		ParallelAsmReadThreads:                 "ParallelAsmReadThreads",
		ReadAheadBlocks:                        "ReadAheadBlocks",
		ReadTableSpaceName:                     "ReadTableSpaceName",
		ReplacePathPrefix:                      "ReplacePathPrefix",
		RetryInterval:                          "RetryInterval",
		SecretsManagerAccessRoleArn:            "SecretsManagerAccessRoleArn",
		SecretsManagerOracleAsmAccessRoleArn:   "SecretsManagerOracleAsmAccessRoleArn",
		SecretsManagerOracleAsmSecretId:        "SecretsManagerOracleAsmSecretId",
		SecretsManagerSecretId:                 "SecretsManagerSecretId",
		SecurityDbEncryption:                   "SecurityDbEncryption",
		SecurityDbEncryptionName:               "SecurityDbEncryptionName",
		SpatialDataOptionToGeoJsonFunctionName: "SpatialDataOptionToGeoJsonFunctionName",
		StandbyDelayTime:                       "StandbyDelayTime",
		UseAlternateFolderForOnline:            "UseAlternateFolderForOnline",
		UseBFile:                               "UseBFile",
		UseDirectPathFullLoad:                  "UseDirectPathFullLoad",
		UseLogminerReader:                      "UseLogminerReader",
		UsePathPrefix:                          "UsePathPrefix",
	}

	// AWS_DMS_Endpoint_OracleSettings__PropertiesSlice reports all the CloudFormation properties for AWS::DMS::Endpoint.OracleSettings.
	AWS_DMS_Endpoint_OracleSettings__PropertiesSlice = []string{
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.AccessAlternateDirectly,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.AddSupplementalLogging,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.AdditionalArchivedLogDestId,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.AllowSelectNestedTables,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.ArchivedLogDestId,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.ArchivedLogsOnly,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.AsmPassword,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.AsmServer,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.AsmUser,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.CharLengthSemantics,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.DirectPathNoLog,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.DirectPathParallelLoad,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.EnableHomogenousTablespace,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.ExtraArchivedLogDestIds,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.FailTasksOnLobTruncation,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.NumberDatatypeScale,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.OraclePathPrefix,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.ParallelAsmReadThreads,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.ReadAheadBlocks,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.ReadTableSpaceName,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.ReplacePathPrefix,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.RetryInterval,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.SecretsManagerAccessRoleArn,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.SecretsManagerOracleAsmAccessRoleArn,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.SecretsManagerOracleAsmSecretId,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.SecretsManagerSecretId,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.SecurityDbEncryption,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.SecurityDbEncryptionName,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.SpatialDataOptionToGeoJsonFunctionName,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.StandbyDelayTime,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.UseAlternateFolderForOnline,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.UseBFile,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.UseDirectPathFullLoad,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.UseLogminerReader,
		AWS_DMS_Endpoint_OracleSettings__PropertiesMap.UsePathPrefix,
	}
)

// AWS_DMS_Endpoint_OracleSettings is a binding for AWS::DMS::Endpoint.OracleSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html
type AWS_DMS_Endpoint_OracleSettings struct {
	// AccessAlternateDirectly is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-accessalternatedirectly
	AccessAlternateDirectly cfz.Expression[bool] `json:"AccessAlternateDirectly,omitempty"`

	// AddSupplementalLogging is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-addsupplementallogging
	AddSupplementalLogging cfz.Expression[bool] `json:"AddSupplementalLogging,omitempty"`

	// AdditionalArchivedLogDestId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-additionalarchivedlogdestid
	AdditionalArchivedLogDestId cfz.Expression[int32] `json:"AdditionalArchivedLogDestId,omitempty"`

	// AllowSelectNestedTables is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-allowselectnestedtables
	AllowSelectNestedTables cfz.Expression[bool] `json:"AllowSelectNestedTables,omitempty"`

	// ArchivedLogDestId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-archivedlogdestid
	ArchivedLogDestId cfz.Expression[int32] `json:"ArchivedLogDestId,omitempty"`

	// ArchivedLogsOnly is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-archivedlogsonly
	ArchivedLogsOnly cfz.Expression[bool] `json:"ArchivedLogsOnly,omitempty"`

	// AsmPassword is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-asmpassword
	AsmPassword cfz.Expression[string] `json:"AsmPassword,omitempty"`

	// AsmServer is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-asmserver
	AsmServer cfz.Expression[string] `json:"AsmServer,omitempty"`

	// AsmUser is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-asmuser
	AsmUser cfz.Expression[string] `json:"AsmUser,omitempty"`

	// CharLengthSemantics is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-charlengthsemantics
	CharLengthSemantics cfz.Expression[string] `json:"CharLengthSemantics,omitempty"`

	// DirectPathNoLog is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-directpathnolog
	DirectPathNoLog cfz.Expression[bool] `json:"DirectPathNoLog,omitempty"`

	// DirectPathParallelLoad is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-directpathparallelload
	DirectPathParallelLoad cfz.Expression[bool] `json:"DirectPathParallelLoad,omitempty"`

	// EnableHomogenousTablespace is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-enablehomogenoustablespace
	EnableHomogenousTablespace cfz.Expression[bool] `json:"EnableHomogenousTablespace,omitempty"`

	// ExtraArchivedLogDestIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-extraarchivedlogdestids
	ExtraArchivedLogDestIds cfz.ExpressionSlice[int32] `json:"ExtraArchivedLogDestIds,omitempty"`

	// FailTasksOnLobTruncation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-failtasksonlobtruncation
	FailTasksOnLobTruncation cfz.Expression[bool] `json:"FailTasksOnLobTruncation,omitempty"`

	// NumberDatatypeScale is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-numberdatatypescale
	NumberDatatypeScale cfz.Expression[int32] `json:"NumberDatatypeScale,omitempty"`

	// OraclePathPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-oraclepathprefix
	OraclePathPrefix cfz.Expression[string] `json:"OraclePathPrefix,omitempty"`

	// ParallelAsmReadThreads is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-parallelasmreadthreads
	ParallelAsmReadThreads cfz.Expression[int32] `json:"ParallelAsmReadThreads,omitempty"`

	// ReadAheadBlocks is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-readaheadblocks
	ReadAheadBlocks cfz.Expression[int32] `json:"ReadAheadBlocks,omitempty"`

	// ReadTableSpaceName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-readtablespacename
	ReadTableSpaceName cfz.Expression[bool] `json:"ReadTableSpaceName,omitempty"`

	// ReplacePathPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-replacepathprefix
	ReplacePathPrefix cfz.Expression[bool] `json:"ReplacePathPrefix,omitempty"`

	// RetryInterval is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-retryinterval
	RetryInterval cfz.Expression[int32] `json:"RetryInterval,omitempty"`

	// SecretsManagerAccessRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanageraccessrolearn
	SecretsManagerAccessRoleArn cfz.Expression[string] `json:"SecretsManagerAccessRoleArn,omitempty"`

	// SecretsManagerOracleAsmAccessRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanageroracleasmaccessrolearn
	SecretsManagerOracleAsmAccessRoleArn cfz.Expression[string] `json:"SecretsManagerOracleAsmAccessRoleArn,omitempty"`

	// SecretsManagerOracleAsmSecretId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanageroracleasmsecretid
	SecretsManagerOracleAsmSecretId cfz.Expression[string] `json:"SecretsManagerOracleAsmSecretId,omitempty"`

	// SecretsManagerSecretId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanagersecretid
	SecretsManagerSecretId cfz.Expression[string] `json:"SecretsManagerSecretId,omitempty"`

	// SecurityDbEncryption is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-securitydbencryption
	SecurityDbEncryption cfz.Expression[string] `json:"SecurityDbEncryption,omitempty"`

	// SecurityDbEncryptionName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-securitydbencryptionname
	SecurityDbEncryptionName cfz.Expression[string] `json:"SecurityDbEncryptionName,omitempty"`

	// SpatialDataOptionToGeoJsonFunctionName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-spatialdataoptiontogeojsonfunctionname
	SpatialDataOptionToGeoJsonFunctionName cfz.Expression[string] `json:"SpatialDataOptionToGeoJsonFunctionName,omitempty"`

	// StandbyDelayTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-standbydelaytime
	StandbyDelayTime cfz.Expression[int32] `json:"StandbyDelayTime,omitempty"`

	// UseAlternateFolderForOnline is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usealternatefolderforonline
	UseAlternateFolderForOnline cfz.Expression[bool] `json:"UseAlternateFolderForOnline,omitempty"`

	// UseBFile is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usebfile
	UseBFile cfz.Expression[bool] `json:"UseBFile,omitempty"`

	// UseDirectPathFullLoad is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usedirectpathfullload
	UseDirectPathFullLoad cfz.Expression[bool] `json:"UseDirectPathFullLoad,omitempty"`

	// UseLogminerReader is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-uselogminerreader
	UseLogminerReader cfz.Expression[bool] `json:"UseLogminerReader,omitempty"`

	// UsePathPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usepathprefix
	UsePathPrefix cfz.Expression[string] `json:"UsePathPrefix,omitempty"`
}

// New__AWS_DMS_Endpoint_OracleSettings initializes a new AWS_DMS_Endpoint_OracleSettings.
func New__AWS_DMS_Endpoint_OracleSettings() AWS_DMS_Endpoint_OracleSettings {
	return AWS_DMS_Endpoint_OracleSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_DMS_Endpoint_OracleSettings) GetType() string {
	return AWS_DMS_Endpoint_OracleSettings__Type
}

// Set__AccessAlternateDirectly updates property "AccessAlternateDirectly".
func (t AWS_DMS_Endpoint_OracleSettings) Set__AccessAlternateDirectly(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.AccessAlternateDirectly = v
	return t
}

// SetV__AccessAlternateDirectly updates property "AccessAlternateDirectly".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__AccessAlternateDirectly(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.AccessAlternateDirectly = cfz.V(v)
	return t
}

// Set__AddSupplementalLogging updates property "AddSupplementalLogging".
func (t AWS_DMS_Endpoint_OracleSettings) Set__AddSupplementalLogging(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.AddSupplementalLogging = v
	return t
}

// SetV__AddSupplementalLogging updates property "AddSupplementalLogging".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__AddSupplementalLogging(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.AddSupplementalLogging = cfz.V(v)
	return t
}

// Set__AdditionalArchivedLogDestId updates property "AdditionalArchivedLogDestId".
func (t AWS_DMS_Endpoint_OracleSettings) Set__AdditionalArchivedLogDestId(v cfz.Expression[int32]) AWS_DMS_Endpoint_OracleSettings {
	t.AdditionalArchivedLogDestId = v
	return t
}

// SetV__AdditionalArchivedLogDestId updates property "AdditionalArchivedLogDestId".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__AdditionalArchivedLogDestId(v int32) AWS_DMS_Endpoint_OracleSettings {
	t.AdditionalArchivedLogDestId = cfz.V(v)
	return t
}

// Set__AllowSelectNestedTables updates property "AllowSelectNestedTables".
func (t AWS_DMS_Endpoint_OracleSettings) Set__AllowSelectNestedTables(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.AllowSelectNestedTables = v
	return t
}

// SetV__AllowSelectNestedTables updates property "AllowSelectNestedTables".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__AllowSelectNestedTables(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.AllowSelectNestedTables = cfz.V(v)
	return t
}

// Set__ArchivedLogDestId updates property "ArchivedLogDestId".
func (t AWS_DMS_Endpoint_OracleSettings) Set__ArchivedLogDestId(v cfz.Expression[int32]) AWS_DMS_Endpoint_OracleSettings {
	t.ArchivedLogDestId = v
	return t
}

// SetV__ArchivedLogDestId updates property "ArchivedLogDestId".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__ArchivedLogDestId(v int32) AWS_DMS_Endpoint_OracleSettings {
	t.ArchivedLogDestId = cfz.V(v)
	return t
}

// Set__ArchivedLogsOnly updates property "ArchivedLogsOnly".
func (t AWS_DMS_Endpoint_OracleSettings) Set__ArchivedLogsOnly(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.ArchivedLogsOnly = v
	return t
}

// SetV__ArchivedLogsOnly updates property "ArchivedLogsOnly".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__ArchivedLogsOnly(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.ArchivedLogsOnly = cfz.V(v)
	return t
}

// Set__AsmPassword updates property "AsmPassword".
func (t AWS_DMS_Endpoint_OracleSettings) Set__AsmPassword(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.AsmPassword = v
	return t
}

// SetV__AsmPassword updates property "AsmPassword".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__AsmPassword(v string) AWS_DMS_Endpoint_OracleSettings {
	t.AsmPassword = cfz.V(v)
	return t
}

// Set__AsmServer updates property "AsmServer".
func (t AWS_DMS_Endpoint_OracleSettings) Set__AsmServer(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.AsmServer = v
	return t
}

// SetV__AsmServer updates property "AsmServer".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__AsmServer(v string) AWS_DMS_Endpoint_OracleSettings {
	t.AsmServer = cfz.V(v)
	return t
}

// Set__AsmUser updates property "AsmUser".
func (t AWS_DMS_Endpoint_OracleSettings) Set__AsmUser(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.AsmUser = v
	return t
}

// SetV__AsmUser updates property "AsmUser".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__AsmUser(v string) AWS_DMS_Endpoint_OracleSettings {
	t.AsmUser = cfz.V(v)
	return t
}

// Set__CharLengthSemantics updates property "CharLengthSemantics".
func (t AWS_DMS_Endpoint_OracleSettings) Set__CharLengthSemantics(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.CharLengthSemantics = v
	return t
}

// SetV__CharLengthSemantics updates property "CharLengthSemantics".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__CharLengthSemantics(v string) AWS_DMS_Endpoint_OracleSettings {
	t.CharLengthSemantics = cfz.V(v)
	return t
}

// Set__DirectPathNoLog updates property "DirectPathNoLog".
func (t AWS_DMS_Endpoint_OracleSettings) Set__DirectPathNoLog(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.DirectPathNoLog = v
	return t
}

// SetV__DirectPathNoLog updates property "DirectPathNoLog".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__DirectPathNoLog(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.DirectPathNoLog = cfz.V(v)
	return t
}

// Set__DirectPathParallelLoad updates property "DirectPathParallelLoad".
func (t AWS_DMS_Endpoint_OracleSettings) Set__DirectPathParallelLoad(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.DirectPathParallelLoad = v
	return t
}

// SetV__DirectPathParallelLoad updates property "DirectPathParallelLoad".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__DirectPathParallelLoad(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.DirectPathParallelLoad = cfz.V(v)
	return t
}

// Set__EnableHomogenousTablespace updates property "EnableHomogenousTablespace".
func (t AWS_DMS_Endpoint_OracleSettings) Set__EnableHomogenousTablespace(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.EnableHomogenousTablespace = v
	return t
}

// SetV__EnableHomogenousTablespace updates property "EnableHomogenousTablespace".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__EnableHomogenousTablespace(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.EnableHomogenousTablespace = cfz.V(v)
	return t
}

// Set__ExtraArchivedLogDestIds updates property "ExtraArchivedLogDestIds".
func (t AWS_DMS_Endpoint_OracleSettings) Set__ExtraArchivedLogDestIds(v cfz.ExpressionSlice[int32]) AWS_DMS_Endpoint_OracleSettings {
	t.ExtraArchivedLogDestIds = v
	return t
}

// SetS__ExtraArchivedLogDestIds updates property "ExtraArchivedLogDestIds".
func (t AWS_DMS_Endpoint_OracleSettings) SetS__ExtraArchivedLogDestIds(v ...cfz.Expression[int32]) AWS_DMS_Endpoint_OracleSettings {
	t.ExtraArchivedLogDestIds = cfz.S(v...)
	return t
}

// SetSV__ExtraArchivedLogDestIds updates property "ExtraArchivedLogDestIds".
func (t AWS_DMS_Endpoint_OracleSettings) SetSV__ExtraArchivedLogDestIds(v ...int32) AWS_DMS_Endpoint_OracleSettings {
	t.ExtraArchivedLogDestIds = cfz.SV(v...)
	return t
}

// Set__FailTasksOnLobTruncation updates property "FailTasksOnLobTruncation".
func (t AWS_DMS_Endpoint_OracleSettings) Set__FailTasksOnLobTruncation(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.FailTasksOnLobTruncation = v
	return t
}

// SetV__FailTasksOnLobTruncation updates property "FailTasksOnLobTruncation".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__FailTasksOnLobTruncation(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.FailTasksOnLobTruncation = cfz.V(v)
	return t
}

// Set__NumberDatatypeScale updates property "NumberDatatypeScale".
func (t AWS_DMS_Endpoint_OracleSettings) Set__NumberDatatypeScale(v cfz.Expression[int32]) AWS_DMS_Endpoint_OracleSettings {
	t.NumberDatatypeScale = v
	return t
}

// SetV__NumberDatatypeScale updates property "NumberDatatypeScale".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__NumberDatatypeScale(v int32) AWS_DMS_Endpoint_OracleSettings {
	t.NumberDatatypeScale = cfz.V(v)
	return t
}

// Set__OraclePathPrefix updates property "OraclePathPrefix".
func (t AWS_DMS_Endpoint_OracleSettings) Set__OraclePathPrefix(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.OraclePathPrefix = v
	return t
}

// SetV__OraclePathPrefix updates property "OraclePathPrefix".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__OraclePathPrefix(v string) AWS_DMS_Endpoint_OracleSettings {
	t.OraclePathPrefix = cfz.V(v)
	return t
}

// Set__ParallelAsmReadThreads updates property "ParallelAsmReadThreads".
func (t AWS_DMS_Endpoint_OracleSettings) Set__ParallelAsmReadThreads(v cfz.Expression[int32]) AWS_DMS_Endpoint_OracleSettings {
	t.ParallelAsmReadThreads = v
	return t
}

// SetV__ParallelAsmReadThreads updates property "ParallelAsmReadThreads".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__ParallelAsmReadThreads(v int32) AWS_DMS_Endpoint_OracleSettings {
	t.ParallelAsmReadThreads = cfz.V(v)
	return t
}

// Set__ReadAheadBlocks updates property "ReadAheadBlocks".
func (t AWS_DMS_Endpoint_OracleSettings) Set__ReadAheadBlocks(v cfz.Expression[int32]) AWS_DMS_Endpoint_OracleSettings {
	t.ReadAheadBlocks = v
	return t
}

// SetV__ReadAheadBlocks updates property "ReadAheadBlocks".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__ReadAheadBlocks(v int32) AWS_DMS_Endpoint_OracleSettings {
	t.ReadAheadBlocks = cfz.V(v)
	return t
}

// Set__ReadTableSpaceName updates property "ReadTableSpaceName".
func (t AWS_DMS_Endpoint_OracleSettings) Set__ReadTableSpaceName(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.ReadTableSpaceName = v
	return t
}

// SetV__ReadTableSpaceName updates property "ReadTableSpaceName".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__ReadTableSpaceName(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.ReadTableSpaceName = cfz.V(v)
	return t
}

// Set__ReplacePathPrefix updates property "ReplacePathPrefix".
func (t AWS_DMS_Endpoint_OracleSettings) Set__ReplacePathPrefix(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.ReplacePathPrefix = v
	return t
}

// SetV__ReplacePathPrefix updates property "ReplacePathPrefix".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__ReplacePathPrefix(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.ReplacePathPrefix = cfz.V(v)
	return t
}

// Set__RetryInterval updates property "RetryInterval".
func (t AWS_DMS_Endpoint_OracleSettings) Set__RetryInterval(v cfz.Expression[int32]) AWS_DMS_Endpoint_OracleSettings {
	t.RetryInterval = v
	return t
}

// SetV__RetryInterval updates property "RetryInterval".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__RetryInterval(v int32) AWS_DMS_Endpoint_OracleSettings {
	t.RetryInterval = cfz.V(v)
	return t
}

// Set__SecretsManagerAccessRoleArn updates property "SecretsManagerAccessRoleArn".
func (t AWS_DMS_Endpoint_OracleSettings) Set__SecretsManagerAccessRoleArn(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.SecretsManagerAccessRoleArn = v
	return t
}

// SetV__SecretsManagerAccessRoleArn updates property "SecretsManagerAccessRoleArn".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__SecretsManagerAccessRoleArn(v string) AWS_DMS_Endpoint_OracleSettings {
	t.SecretsManagerAccessRoleArn = cfz.V(v)
	return t
}

// Set__SecretsManagerOracleAsmAccessRoleArn updates property "SecretsManagerOracleAsmAccessRoleArn".
func (t AWS_DMS_Endpoint_OracleSettings) Set__SecretsManagerOracleAsmAccessRoleArn(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.SecretsManagerOracleAsmAccessRoleArn = v
	return t
}

// SetV__SecretsManagerOracleAsmAccessRoleArn updates property "SecretsManagerOracleAsmAccessRoleArn".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__SecretsManagerOracleAsmAccessRoleArn(v string) AWS_DMS_Endpoint_OracleSettings {
	t.SecretsManagerOracleAsmAccessRoleArn = cfz.V(v)
	return t
}

// Set__SecretsManagerOracleAsmSecretId updates property "SecretsManagerOracleAsmSecretId".
func (t AWS_DMS_Endpoint_OracleSettings) Set__SecretsManagerOracleAsmSecretId(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.SecretsManagerOracleAsmSecretId = v
	return t
}

// SetV__SecretsManagerOracleAsmSecretId updates property "SecretsManagerOracleAsmSecretId".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__SecretsManagerOracleAsmSecretId(v string) AWS_DMS_Endpoint_OracleSettings {
	t.SecretsManagerOracleAsmSecretId = cfz.V(v)
	return t
}

// Set__SecretsManagerSecretId updates property "SecretsManagerSecretId".
func (t AWS_DMS_Endpoint_OracleSettings) Set__SecretsManagerSecretId(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.SecretsManagerSecretId = v
	return t
}

// SetV__SecretsManagerSecretId updates property "SecretsManagerSecretId".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__SecretsManagerSecretId(v string) AWS_DMS_Endpoint_OracleSettings {
	t.SecretsManagerSecretId = cfz.V(v)
	return t
}

// Set__SecurityDbEncryption updates property "SecurityDbEncryption".
func (t AWS_DMS_Endpoint_OracleSettings) Set__SecurityDbEncryption(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.SecurityDbEncryption = v
	return t
}

// SetV__SecurityDbEncryption updates property "SecurityDbEncryption".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__SecurityDbEncryption(v string) AWS_DMS_Endpoint_OracleSettings {
	t.SecurityDbEncryption = cfz.V(v)
	return t
}

// Set__SecurityDbEncryptionName updates property "SecurityDbEncryptionName".
func (t AWS_DMS_Endpoint_OracleSettings) Set__SecurityDbEncryptionName(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.SecurityDbEncryptionName = v
	return t
}

// SetV__SecurityDbEncryptionName updates property "SecurityDbEncryptionName".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__SecurityDbEncryptionName(v string) AWS_DMS_Endpoint_OracleSettings {
	t.SecurityDbEncryptionName = cfz.V(v)
	return t
}

// Set__SpatialDataOptionToGeoJsonFunctionName updates property "SpatialDataOptionToGeoJsonFunctionName".
func (t AWS_DMS_Endpoint_OracleSettings) Set__SpatialDataOptionToGeoJsonFunctionName(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.SpatialDataOptionToGeoJsonFunctionName = v
	return t
}

// SetV__SpatialDataOptionToGeoJsonFunctionName updates property "SpatialDataOptionToGeoJsonFunctionName".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__SpatialDataOptionToGeoJsonFunctionName(v string) AWS_DMS_Endpoint_OracleSettings {
	t.SpatialDataOptionToGeoJsonFunctionName = cfz.V(v)
	return t
}

// Set__StandbyDelayTime updates property "StandbyDelayTime".
func (t AWS_DMS_Endpoint_OracleSettings) Set__StandbyDelayTime(v cfz.Expression[int32]) AWS_DMS_Endpoint_OracleSettings {
	t.StandbyDelayTime = v
	return t
}

// SetV__StandbyDelayTime updates property "StandbyDelayTime".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__StandbyDelayTime(v int32) AWS_DMS_Endpoint_OracleSettings {
	t.StandbyDelayTime = cfz.V(v)
	return t
}

// Set__UseAlternateFolderForOnline updates property "UseAlternateFolderForOnline".
func (t AWS_DMS_Endpoint_OracleSettings) Set__UseAlternateFolderForOnline(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.UseAlternateFolderForOnline = v
	return t
}

// SetV__UseAlternateFolderForOnline updates property "UseAlternateFolderForOnline".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__UseAlternateFolderForOnline(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.UseAlternateFolderForOnline = cfz.V(v)
	return t
}

// Set__UseBFile updates property "UseBFile".
func (t AWS_DMS_Endpoint_OracleSettings) Set__UseBFile(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.UseBFile = v
	return t
}

// SetV__UseBFile updates property "UseBFile".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__UseBFile(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.UseBFile = cfz.V(v)
	return t
}

// Set__UseDirectPathFullLoad updates property "UseDirectPathFullLoad".
func (t AWS_DMS_Endpoint_OracleSettings) Set__UseDirectPathFullLoad(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.UseDirectPathFullLoad = v
	return t
}

// SetV__UseDirectPathFullLoad updates property "UseDirectPathFullLoad".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__UseDirectPathFullLoad(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.UseDirectPathFullLoad = cfz.V(v)
	return t
}

// Set__UseLogminerReader updates property "UseLogminerReader".
func (t AWS_DMS_Endpoint_OracleSettings) Set__UseLogminerReader(v cfz.Expression[bool]) AWS_DMS_Endpoint_OracleSettings {
	t.UseLogminerReader = v
	return t
}

// SetV__UseLogminerReader updates property "UseLogminerReader".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__UseLogminerReader(v bool) AWS_DMS_Endpoint_OracleSettings {
	t.UseLogminerReader = cfz.V(v)
	return t
}

// Set__UsePathPrefix updates property "UsePathPrefix".
func (t AWS_DMS_Endpoint_OracleSettings) Set__UsePathPrefix(v cfz.Expression[string]) AWS_DMS_Endpoint_OracleSettings {
	t.UsePathPrefix = v
	return t
}

// SetV__UsePathPrefix updates property "UsePathPrefix".
func (t AWS_DMS_Endpoint_OracleSettings) SetV__UsePathPrefix(v string) AWS_DMS_Endpoint_OracleSettings {
	t.UsePathPrefix = cfz.V(v)
	return t
}
