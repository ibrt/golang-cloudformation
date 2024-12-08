// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dms

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DMS_Endpoint_S3Settings__Type is the CloudFormation type for AWS::DMS::Endpoint.S3Settings.
	AWS_DMS_Endpoint_S3Settings__Type = "AWS::DMS::Endpoint.S3Settings"
)

var (
	// AWS_DMS_Endpoint_S3Settings__PropertiesMap reports all the CloudFormation properties for AWS::DMS::Endpoint.S3Settings.
	AWS_DMS_Endpoint_S3Settings__PropertiesMap = struct {
		AddColumnName                        string
		AddTrailingPaddingCharacter          string
		BucketFolder                         string
		BucketName                           string
		CannedAclForObjects                  string
		CdcInsertsAndUpdates                 string
		CdcInsertsOnly                       string
		CdcMaxBatchInterval                  string
		CdcMinFileSize                       string
		CdcPath                              string
		CompressionType                      string
		CsvDelimiter                         string
		CsvNoSupValue                        string
		CsvNullValue                         string
		CsvRowDelimiter                      string
		DataFormat                           string
		DataPageSize                         string
		DatePartitionDelimiter               string
		DatePartitionEnabled                 string
		DatePartitionSequence                string
		DatePartitionTimezone                string
		DictPageSizeLimit                    string
		EnableStatistics                     string
		EncodingType                         string
		EncryptionMode                       string
		ExpectedBucketOwner                  string
		ExternalTableDefinition              string
		GlueCatalogGeneration                string
		IgnoreHeaderRows                     string
		IncludeOpForFullLoad                 string
		MaxFileSize                          string
		ParquetTimestampInMillisecond        string
		ParquetVersion                       string
		PreserveTransactions                 string
		Rfc4180                              string
		RowGroupLength                       string
		ServerSideEncryptionKmsKeyId         string
		ServiceAccessRoleArn                 string
		TimestampColumnName                  string
		UseCsvNoSupValue                     string
		UseTaskStartTimeForFullLoadTimestamp string
	}{
		AddColumnName:                        "AddColumnName",
		AddTrailingPaddingCharacter:          "AddTrailingPaddingCharacter",
		BucketFolder:                         "BucketFolder",
		BucketName:                           "BucketName",
		CannedAclForObjects:                  "CannedAclForObjects",
		CdcInsertsAndUpdates:                 "CdcInsertsAndUpdates",
		CdcInsertsOnly:                       "CdcInsertsOnly",
		CdcMaxBatchInterval:                  "CdcMaxBatchInterval",
		CdcMinFileSize:                       "CdcMinFileSize",
		CdcPath:                              "CdcPath",
		CompressionType:                      "CompressionType",
		CsvDelimiter:                         "CsvDelimiter",
		CsvNoSupValue:                        "CsvNoSupValue",
		CsvNullValue:                         "CsvNullValue",
		CsvRowDelimiter:                      "CsvRowDelimiter",
		DataFormat:                           "DataFormat",
		DataPageSize:                         "DataPageSize",
		DatePartitionDelimiter:               "DatePartitionDelimiter",
		DatePartitionEnabled:                 "DatePartitionEnabled",
		DatePartitionSequence:                "DatePartitionSequence",
		DatePartitionTimezone:                "DatePartitionTimezone",
		DictPageSizeLimit:                    "DictPageSizeLimit",
		EnableStatistics:                     "EnableStatistics",
		EncodingType:                         "EncodingType",
		EncryptionMode:                       "EncryptionMode",
		ExpectedBucketOwner:                  "ExpectedBucketOwner",
		ExternalTableDefinition:              "ExternalTableDefinition",
		GlueCatalogGeneration:                "GlueCatalogGeneration",
		IgnoreHeaderRows:                     "IgnoreHeaderRows",
		IncludeOpForFullLoad:                 "IncludeOpForFullLoad",
		MaxFileSize:                          "MaxFileSize",
		ParquetTimestampInMillisecond:        "ParquetTimestampInMillisecond",
		ParquetVersion:                       "ParquetVersion",
		PreserveTransactions:                 "PreserveTransactions",
		Rfc4180:                              "Rfc4180",
		RowGroupLength:                       "RowGroupLength",
		ServerSideEncryptionKmsKeyId:         "ServerSideEncryptionKmsKeyId",
		ServiceAccessRoleArn:                 "ServiceAccessRoleArn",
		TimestampColumnName:                  "TimestampColumnName",
		UseCsvNoSupValue:                     "UseCsvNoSupValue",
		UseTaskStartTimeForFullLoadTimestamp: "UseTaskStartTimeForFullLoadTimestamp",
	}

	// AWS_DMS_Endpoint_S3Settings__PropertiesSlice reports all the CloudFormation properties for AWS::DMS::Endpoint.S3Settings.
	AWS_DMS_Endpoint_S3Settings__PropertiesSlice = []string{
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.AddColumnName,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.AddTrailingPaddingCharacter,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.BucketFolder,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.BucketName,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CannedAclForObjects,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CdcInsertsAndUpdates,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CdcInsertsOnly,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CdcMaxBatchInterval,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CdcMinFileSize,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CdcPath,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CompressionType,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CsvDelimiter,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CsvNoSupValue,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CsvNullValue,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.CsvRowDelimiter,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.DataFormat,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.DataPageSize,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.DatePartitionDelimiter,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.DatePartitionEnabled,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.DatePartitionSequence,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.DatePartitionTimezone,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.DictPageSizeLimit,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.EnableStatistics,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.EncodingType,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.EncryptionMode,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.ExpectedBucketOwner,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.ExternalTableDefinition,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.GlueCatalogGeneration,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.IgnoreHeaderRows,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.IncludeOpForFullLoad,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.MaxFileSize,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.ParquetTimestampInMillisecond,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.ParquetVersion,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.PreserveTransactions,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.Rfc4180,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.RowGroupLength,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.ServerSideEncryptionKmsKeyId,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.ServiceAccessRoleArn,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.TimestampColumnName,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.UseCsvNoSupValue,
		AWS_DMS_Endpoint_S3Settings__PropertiesMap.UseTaskStartTimeForFullLoadTimestamp,
	}
)

// AWS_DMS_Endpoint_S3Settings is a binding for AWS::DMS::Endpoint.S3Settings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html
type AWS_DMS_Endpoint_S3Settings struct {
	// AddColumnName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-addcolumnname
	AddColumnName cfz.Expression[bool] `json:"AddColumnName,omitempty"`

	// AddTrailingPaddingCharacter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-addtrailingpaddingcharacter
	AddTrailingPaddingCharacter cfz.Expression[bool] `json:"AddTrailingPaddingCharacter,omitempty"`

	// BucketFolder is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-bucketfolder
	BucketFolder cfz.Expression[string] `json:"BucketFolder,omitempty"`

	// BucketName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-bucketname
	BucketName cfz.Expression[string] `json:"BucketName,omitempty"`

	// CannedAclForObjects is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-cannedaclforobjects
	CannedAclForObjects cfz.Expression[string] `json:"CannedAclForObjects,omitempty"`

	// CdcInsertsAndUpdates is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-cdcinsertsandupdates
	CdcInsertsAndUpdates cfz.Expression[bool] `json:"CdcInsertsAndUpdates,omitempty"`

	// CdcInsertsOnly is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-cdcinsertsonly
	CdcInsertsOnly cfz.Expression[bool] `json:"CdcInsertsOnly,omitempty"`

	// CdcMaxBatchInterval is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-cdcmaxbatchinterval
	CdcMaxBatchInterval cfz.Expression[int32] `json:"CdcMaxBatchInterval,omitempty"`

	// CdcMinFileSize is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-cdcminfilesize
	CdcMinFileSize cfz.Expression[int32] `json:"CdcMinFileSize,omitempty"`

	// CdcPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-cdcpath
	CdcPath cfz.Expression[string] `json:"CdcPath,omitempty"`

	// CompressionType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-compressiontype
	CompressionType cfz.Expression[string] `json:"CompressionType,omitempty"`

	// CsvDelimiter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-csvdelimiter
	CsvDelimiter cfz.Expression[string] `json:"CsvDelimiter,omitempty"`

	// CsvNoSupValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-csvnosupvalue
	CsvNoSupValue cfz.Expression[string] `json:"CsvNoSupValue,omitempty"`

	// CsvNullValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-csvnullvalue
	CsvNullValue cfz.Expression[string] `json:"CsvNullValue,omitempty"`

	// CsvRowDelimiter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-csvrowdelimiter
	CsvRowDelimiter cfz.Expression[string] `json:"CsvRowDelimiter,omitempty"`

	// DataFormat is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-dataformat
	DataFormat cfz.Expression[string] `json:"DataFormat,omitempty"`

	// DataPageSize is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-datapagesize
	DataPageSize cfz.Expression[int32] `json:"DataPageSize,omitempty"`

	// DatePartitionDelimiter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-datepartitiondelimiter
	DatePartitionDelimiter cfz.Expression[string] `json:"DatePartitionDelimiter,omitempty"`

	// DatePartitionEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-datepartitionenabled
	DatePartitionEnabled cfz.Expression[bool] `json:"DatePartitionEnabled,omitempty"`

	// DatePartitionSequence is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-datepartitionsequence
	DatePartitionSequence cfz.Expression[string] `json:"DatePartitionSequence,omitempty"`

	// DatePartitionTimezone is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-datepartitiontimezone
	DatePartitionTimezone cfz.Expression[string] `json:"DatePartitionTimezone,omitempty"`

	// DictPageSizeLimit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-dictpagesizelimit
	DictPageSizeLimit cfz.Expression[int32] `json:"DictPageSizeLimit,omitempty"`

	// EnableStatistics is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-enablestatistics
	EnableStatistics cfz.Expression[bool] `json:"EnableStatistics,omitempty"`

	// EncodingType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-encodingtype
	EncodingType cfz.Expression[string] `json:"EncodingType,omitempty"`

	// EncryptionMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-encryptionmode
	EncryptionMode cfz.Expression[string] `json:"EncryptionMode,omitempty"`

	// ExpectedBucketOwner is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-expectedbucketowner
	ExpectedBucketOwner cfz.Expression[string] `json:"ExpectedBucketOwner,omitempty"`

	// ExternalTableDefinition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-externaltabledefinition
	ExternalTableDefinition cfz.Expression[string] `json:"ExternalTableDefinition,omitempty"`

	// GlueCatalogGeneration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-gluecataloggeneration
	GlueCatalogGeneration cfz.Expression[bool] `json:"GlueCatalogGeneration,omitempty"`

	// IgnoreHeaderRows is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-ignoreheaderrows
	IgnoreHeaderRows cfz.Expression[int32] `json:"IgnoreHeaderRows,omitempty"`

	// IncludeOpForFullLoad is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-includeopforfullload
	IncludeOpForFullLoad cfz.Expression[bool] `json:"IncludeOpForFullLoad,omitempty"`

	// MaxFileSize is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-maxfilesize
	MaxFileSize cfz.Expression[int32] `json:"MaxFileSize,omitempty"`

	// ParquetTimestampInMillisecond is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-parquettimestampinmillisecond
	ParquetTimestampInMillisecond cfz.Expression[bool] `json:"ParquetTimestampInMillisecond,omitempty"`

	// ParquetVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-parquetversion
	ParquetVersion cfz.Expression[string] `json:"ParquetVersion,omitempty"`

	// PreserveTransactions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-preservetransactions
	PreserveTransactions cfz.Expression[bool] `json:"PreserveTransactions,omitempty"`

	// Rfc4180 is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-rfc4180
	Rfc4180 cfz.Expression[bool] `json:"Rfc4180,omitempty"`

	// RowGroupLength is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-rowgrouplength
	RowGroupLength cfz.Expression[int32] `json:"RowGroupLength,omitempty"`

	// ServerSideEncryptionKmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-serversideencryptionkmskeyid
	ServerSideEncryptionKmsKeyId cfz.Expression[string] `json:"ServerSideEncryptionKmsKeyId,omitempty"`

	// ServiceAccessRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-serviceaccessrolearn
	ServiceAccessRoleArn cfz.Expression[string] `json:"ServiceAccessRoleArn,omitempty"`

	// TimestampColumnName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-timestampcolumnname
	TimestampColumnName cfz.Expression[string] `json:"TimestampColumnName,omitempty"`

	// UseCsvNoSupValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-usecsvnosupvalue
	UseCsvNoSupValue cfz.Expression[bool] `json:"UseCsvNoSupValue,omitempty"`

	// UseTaskStartTimeForFullLoadTimestamp is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-usetaskstarttimeforfullloadtimestamp
	UseTaskStartTimeForFullLoadTimestamp cfz.Expression[bool] `json:"UseTaskStartTimeForFullLoadTimestamp,omitempty"`
}

// New__AWS_DMS_Endpoint_S3Settings initializes a new AWS_DMS_Endpoint_S3Settings.
func New__AWS_DMS_Endpoint_S3Settings() AWS_DMS_Endpoint_S3Settings {
	return AWS_DMS_Endpoint_S3Settings{}
}

// GetType returns the CloudFormation type.
func (AWS_DMS_Endpoint_S3Settings) GetType() string {
	return AWS_DMS_Endpoint_S3Settings__Type
}

// Set__AddColumnName updates property "AddColumnName".
func (t AWS_DMS_Endpoint_S3Settings) Set__AddColumnName(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.AddColumnName = v
	return t
}

// SetV__AddColumnName updates property "AddColumnName".
func (t AWS_DMS_Endpoint_S3Settings) SetV__AddColumnName(v bool) AWS_DMS_Endpoint_S3Settings {
	t.AddColumnName = cfz.V(v)
	return t
}

// Set__AddTrailingPaddingCharacter updates property "AddTrailingPaddingCharacter".
func (t AWS_DMS_Endpoint_S3Settings) Set__AddTrailingPaddingCharacter(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.AddTrailingPaddingCharacter = v
	return t
}

// SetV__AddTrailingPaddingCharacter updates property "AddTrailingPaddingCharacter".
func (t AWS_DMS_Endpoint_S3Settings) SetV__AddTrailingPaddingCharacter(v bool) AWS_DMS_Endpoint_S3Settings {
	t.AddTrailingPaddingCharacter = cfz.V(v)
	return t
}

// Set__BucketFolder updates property "BucketFolder".
func (t AWS_DMS_Endpoint_S3Settings) Set__BucketFolder(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.BucketFolder = v
	return t
}

// SetV__BucketFolder updates property "BucketFolder".
func (t AWS_DMS_Endpoint_S3Settings) SetV__BucketFolder(v string) AWS_DMS_Endpoint_S3Settings {
	t.BucketFolder = cfz.V(v)
	return t
}

// Set__BucketName updates property "BucketName".
func (t AWS_DMS_Endpoint_S3Settings) Set__BucketName(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.BucketName = v
	return t
}

// SetV__BucketName updates property "BucketName".
func (t AWS_DMS_Endpoint_S3Settings) SetV__BucketName(v string) AWS_DMS_Endpoint_S3Settings {
	t.BucketName = cfz.V(v)
	return t
}

// Set__CannedAclForObjects updates property "CannedAclForObjects".
func (t AWS_DMS_Endpoint_S3Settings) Set__CannedAclForObjects(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.CannedAclForObjects = v
	return t
}

// SetV__CannedAclForObjects updates property "CannedAclForObjects".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CannedAclForObjects(v string) AWS_DMS_Endpoint_S3Settings {
	t.CannedAclForObjects = cfz.V(v)
	return t
}

// Set__CdcInsertsAndUpdates updates property "CdcInsertsAndUpdates".
func (t AWS_DMS_Endpoint_S3Settings) Set__CdcInsertsAndUpdates(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.CdcInsertsAndUpdates = v
	return t
}

// SetV__CdcInsertsAndUpdates updates property "CdcInsertsAndUpdates".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CdcInsertsAndUpdates(v bool) AWS_DMS_Endpoint_S3Settings {
	t.CdcInsertsAndUpdates = cfz.V(v)
	return t
}

// Set__CdcInsertsOnly updates property "CdcInsertsOnly".
func (t AWS_DMS_Endpoint_S3Settings) Set__CdcInsertsOnly(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.CdcInsertsOnly = v
	return t
}

// SetV__CdcInsertsOnly updates property "CdcInsertsOnly".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CdcInsertsOnly(v bool) AWS_DMS_Endpoint_S3Settings {
	t.CdcInsertsOnly = cfz.V(v)
	return t
}

// Set__CdcMaxBatchInterval updates property "CdcMaxBatchInterval".
func (t AWS_DMS_Endpoint_S3Settings) Set__CdcMaxBatchInterval(v cfz.Expression[int32]) AWS_DMS_Endpoint_S3Settings {
	t.CdcMaxBatchInterval = v
	return t
}

// SetV__CdcMaxBatchInterval updates property "CdcMaxBatchInterval".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CdcMaxBatchInterval(v int32) AWS_DMS_Endpoint_S3Settings {
	t.CdcMaxBatchInterval = cfz.V(v)
	return t
}

// Set__CdcMinFileSize updates property "CdcMinFileSize".
func (t AWS_DMS_Endpoint_S3Settings) Set__CdcMinFileSize(v cfz.Expression[int32]) AWS_DMS_Endpoint_S3Settings {
	t.CdcMinFileSize = v
	return t
}

// SetV__CdcMinFileSize updates property "CdcMinFileSize".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CdcMinFileSize(v int32) AWS_DMS_Endpoint_S3Settings {
	t.CdcMinFileSize = cfz.V(v)
	return t
}

// Set__CdcPath updates property "CdcPath".
func (t AWS_DMS_Endpoint_S3Settings) Set__CdcPath(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.CdcPath = v
	return t
}

// SetV__CdcPath updates property "CdcPath".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CdcPath(v string) AWS_DMS_Endpoint_S3Settings {
	t.CdcPath = cfz.V(v)
	return t
}

// Set__CompressionType updates property "CompressionType".
func (t AWS_DMS_Endpoint_S3Settings) Set__CompressionType(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.CompressionType = v
	return t
}

// SetV__CompressionType updates property "CompressionType".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CompressionType(v string) AWS_DMS_Endpoint_S3Settings {
	t.CompressionType = cfz.V(v)
	return t
}

// Set__CsvDelimiter updates property "CsvDelimiter".
func (t AWS_DMS_Endpoint_S3Settings) Set__CsvDelimiter(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.CsvDelimiter = v
	return t
}

// SetV__CsvDelimiter updates property "CsvDelimiter".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CsvDelimiter(v string) AWS_DMS_Endpoint_S3Settings {
	t.CsvDelimiter = cfz.V(v)
	return t
}

// Set__CsvNoSupValue updates property "CsvNoSupValue".
func (t AWS_DMS_Endpoint_S3Settings) Set__CsvNoSupValue(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.CsvNoSupValue = v
	return t
}

// SetV__CsvNoSupValue updates property "CsvNoSupValue".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CsvNoSupValue(v string) AWS_DMS_Endpoint_S3Settings {
	t.CsvNoSupValue = cfz.V(v)
	return t
}

// Set__CsvNullValue updates property "CsvNullValue".
func (t AWS_DMS_Endpoint_S3Settings) Set__CsvNullValue(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.CsvNullValue = v
	return t
}

// SetV__CsvNullValue updates property "CsvNullValue".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CsvNullValue(v string) AWS_DMS_Endpoint_S3Settings {
	t.CsvNullValue = cfz.V(v)
	return t
}

// Set__CsvRowDelimiter updates property "CsvRowDelimiter".
func (t AWS_DMS_Endpoint_S3Settings) Set__CsvRowDelimiter(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.CsvRowDelimiter = v
	return t
}

// SetV__CsvRowDelimiter updates property "CsvRowDelimiter".
func (t AWS_DMS_Endpoint_S3Settings) SetV__CsvRowDelimiter(v string) AWS_DMS_Endpoint_S3Settings {
	t.CsvRowDelimiter = cfz.V(v)
	return t
}

// Set__DataFormat updates property "DataFormat".
func (t AWS_DMS_Endpoint_S3Settings) Set__DataFormat(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.DataFormat = v
	return t
}

// SetV__DataFormat updates property "DataFormat".
func (t AWS_DMS_Endpoint_S3Settings) SetV__DataFormat(v string) AWS_DMS_Endpoint_S3Settings {
	t.DataFormat = cfz.V(v)
	return t
}

// Set__DataPageSize updates property "DataPageSize".
func (t AWS_DMS_Endpoint_S3Settings) Set__DataPageSize(v cfz.Expression[int32]) AWS_DMS_Endpoint_S3Settings {
	t.DataPageSize = v
	return t
}

// SetV__DataPageSize updates property "DataPageSize".
func (t AWS_DMS_Endpoint_S3Settings) SetV__DataPageSize(v int32) AWS_DMS_Endpoint_S3Settings {
	t.DataPageSize = cfz.V(v)
	return t
}

// Set__DatePartitionDelimiter updates property "DatePartitionDelimiter".
func (t AWS_DMS_Endpoint_S3Settings) Set__DatePartitionDelimiter(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.DatePartitionDelimiter = v
	return t
}

// SetV__DatePartitionDelimiter updates property "DatePartitionDelimiter".
func (t AWS_DMS_Endpoint_S3Settings) SetV__DatePartitionDelimiter(v string) AWS_DMS_Endpoint_S3Settings {
	t.DatePartitionDelimiter = cfz.V(v)
	return t
}

// Set__DatePartitionEnabled updates property "DatePartitionEnabled".
func (t AWS_DMS_Endpoint_S3Settings) Set__DatePartitionEnabled(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.DatePartitionEnabled = v
	return t
}

// SetV__DatePartitionEnabled updates property "DatePartitionEnabled".
func (t AWS_DMS_Endpoint_S3Settings) SetV__DatePartitionEnabled(v bool) AWS_DMS_Endpoint_S3Settings {
	t.DatePartitionEnabled = cfz.V(v)
	return t
}

// Set__DatePartitionSequence updates property "DatePartitionSequence".
func (t AWS_DMS_Endpoint_S3Settings) Set__DatePartitionSequence(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.DatePartitionSequence = v
	return t
}

// SetV__DatePartitionSequence updates property "DatePartitionSequence".
func (t AWS_DMS_Endpoint_S3Settings) SetV__DatePartitionSequence(v string) AWS_DMS_Endpoint_S3Settings {
	t.DatePartitionSequence = cfz.V(v)
	return t
}

// Set__DatePartitionTimezone updates property "DatePartitionTimezone".
func (t AWS_DMS_Endpoint_S3Settings) Set__DatePartitionTimezone(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.DatePartitionTimezone = v
	return t
}

// SetV__DatePartitionTimezone updates property "DatePartitionTimezone".
func (t AWS_DMS_Endpoint_S3Settings) SetV__DatePartitionTimezone(v string) AWS_DMS_Endpoint_S3Settings {
	t.DatePartitionTimezone = cfz.V(v)
	return t
}

// Set__DictPageSizeLimit updates property "DictPageSizeLimit".
func (t AWS_DMS_Endpoint_S3Settings) Set__DictPageSizeLimit(v cfz.Expression[int32]) AWS_DMS_Endpoint_S3Settings {
	t.DictPageSizeLimit = v
	return t
}

// SetV__DictPageSizeLimit updates property "DictPageSizeLimit".
func (t AWS_DMS_Endpoint_S3Settings) SetV__DictPageSizeLimit(v int32) AWS_DMS_Endpoint_S3Settings {
	t.DictPageSizeLimit = cfz.V(v)
	return t
}

// Set__EnableStatistics updates property "EnableStatistics".
func (t AWS_DMS_Endpoint_S3Settings) Set__EnableStatistics(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.EnableStatistics = v
	return t
}

// SetV__EnableStatistics updates property "EnableStatistics".
func (t AWS_DMS_Endpoint_S3Settings) SetV__EnableStatistics(v bool) AWS_DMS_Endpoint_S3Settings {
	t.EnableStatistics = cfz.V(v)
	return t
}

// Set__EncodingType updates property "EncodingType".
func (t AWS_DMS_Endpoint_S3Settings) Set__EncodingType(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.EncodingType = v
	return t
}

// SetV__EncodingType updates property "EncodingType".
func (t AWS_DMS_Endpoint_S3Settings) SetV__EncodingType(v string) AWS_DMS_Endpoint_S3Settings {
	t.EncodingType = cfz.V(v)
	return t
}

// Set__EncryptionMode updates property "EncryptionMode".
func (t AWS_DMS_Endpoint_S3Settings) Set__EncryptionMode(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.EncryptionMode = v
	return t
}

// SetV__EncryptionMode updates property "EncryptionMode".
func (t AWS_DMS_Endpoint_S3Settings) SetV__EncryptionMode(v string) AWS_DMS_Endpoint_S3Settings {
	t.EncryptionMode = cfz.V(v)
	return t
}

// Set__ExpectedBucketOwner updates property "ExpectedBucketOwner".
func (t AWS_DMS_Endpoint_S3Settings) Set__ExpectedBucketOwner(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.ExpectedBucketOwner = v
	return t
}

// SetV__ExpectedBucketOwner updates property "ExpectedBucketOwner".
func (t AWS_DMS_Endpoint_S3Settings) SetV__ExpectedBucketOwner(v string) AWS_DMS_Endpoint_S3Settings {
	t.ExpectedBucketOwner = cfz.V(v)
	return t
}

// Set__ExternalTableDefinition updates property "ExternalTableDefinition".
func (t AWS_DMS_Endpoint_S3Settings) Set__ExternalTableDefinition(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.ExternalTableDefinition = v
	return t
}

// SetV__ExternalTableDefinition updates property "ExternalTableDefinition".
func (t AWS_DMS_Endpoint_S3Settings) SetV__ExternalTableDefinition(v string) AWS_DMS_Endpoint_S3Settings {
	t.ExternalTableDefinition = cfz.V(v)
	return t
}

// Set__GlueCatalogGeneration updates property "GlueCatalogGeneration".
func (t AWS_DMS_Endpoint_S3Settings) Set__GlueCatalogGeneration(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.GlueCatalogGeneration = v
	return t
}

// SetV__GlueCatalogGeneration updates property "GlueCatalogGeneration".
func (t AWS_DMS_Endpoint_S3Settings) SetV__GlueCatalogGeneration(v bool) AWS_DMS_Endpoint_S3Settings {
	t.GlueCatalogGeneration = cfz.V(v)
	return t
}

// Set__IgnoreHeaderRows updates property "IgnoreHeaderRows".
func (t AWS_DMS_Endpoint_S3Settings) Set__IgnoreHeaderRows(v cfz.Expression[int32]) AWS_DMS_Endpoint_S3Settings {
	t.IgnoreHeaderRows = v
	return t
}

// SetV__IgnoreHeaderRows updates property "IgnoreHeaderRows".
func (t AWS_DMS_Endpoint_S3Settings) SetV__IgnoreHeaderRows(v int32) AWS_DMS_Endpoint_S3Settings {
	t.IgnoreHeaderRows = cfz.V(v)
	return t
}

// Set__IncludeOpForFullLoad updates property "IncludeOpForFullLoad".
func (t AWS_DMS_Endpoint_S3Settings) Set__IncludeOpForFullLoad(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.IncludeOpForFullLoad = v
	return t
}

// SetV__IncludeOpForFullLoad updates property "IncludeOpForFullLoad".
func (t AWS_DMS_Endpoint_S3Settings) SetV__IncludeOpForFullLoad(v bool) AWS_DMS_Endpoint_S3Settings {
	t.IncludeOpForFullLoad = cfz.V(v)
	return t
}

// Set__MaxFileSize updates property "MaxFileSize".
func (t AWS_DMS_Endpoint_S3Settings) Set__MaxFileSize(v cfz.Expression[int32]) AWS_DMS_Endpoint_S3Settings {
	t.MaxFileSize = v
	return t
}

// SetV__MaxFileSize updates property "MaxFileSize".
func (t AWS_DMS_Endpoint_S3Settings) SetV__MaxFileSize(v int32) AWS_DMS_Endpoint_S3Settings {
	t.MaxFileSize = cfz.V(v)
	return t
}

// Set__ParquetTimestampInMillisecond updates property "ParquetTimestampInMillisecond".
func (t AWS_DMS_Endpoint_S3Settings) Set__ParquetTimestampInMillisecond(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.ParquetTimestampInMillisecond = v
	return t
}

// SetV__ParquetTimestampInMillisecond updates property "ParquetTimestampInMillisecond".
func (t AWS_DMS_Endpoint_S3Settings) SetV__ParquetTimestampInMillisecond(v bool) AWS_DMS_Endpoint_S3Settings {
	t.ParquetTimestampInMillisecond = cfz.V(v)
	return t
}

// Set__ParquetVersion updates property "ParquetVersion".
func (t AWS_DMS_Endpoint_S3Settings) Set__ParquetVersion(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.ParquetVersion = v
	return t
}

// SetV__ParquetVersion updates property "ParquetVersion".
func (t AWS_DMS_Endpoint_S3Settings) SetV__ParquetVersion(v string) AWS_DMS_Endpoint_S3Settings {
	t.ParquetVersion = cfz.V(v)
	return t
}

// Set__PreserveTransactions updates property "PreserveTransactions".
func (t AWS_DMS_Endpoint_S3Settings) Set__PreserveTransactions(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.PreserveTransactions = v
	return t
}

// SetV__PreserveTransactions updates property "PreserveTransactions".
func (t AWS_DMS_Endpoint_S3Settings) SetV__PreserveTransactions(v bool) AWS_DMS_Endpoint_S3Settings {
	t.PreserveTransactions = cfz.V(v)
	return t
}

// Set__Rfc4180 updates property "Rfc4180".
func (t AWS_DMS_Endpoint_S3Settings) Set__Rfc4180(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.Rfc4180 = v
	return t
}

// SetV__Rfc4180 updates property "Rfc4180".
func (t AWS_DMS_Endpoint_S3Settings) SetV__Rfc4180(v bool) AWS_DMS_Endpoint_S3Settings {
	t.Rfc4180 = cfz.V(v)
	return t
}

// Set__RowGroupLength updates property "RowGroupLength".
func (t AWS_DMS_Endpoint_S3Settings) Set__RowGroupLength(v cfz.Expression[int32]) AWS_DMS_Endpoint_S3Settings {
	t.RowGroupLength = v
	return t
}

// SetV__RowGroupLength updates property "RowGroupLength".
func (t AWS_DMS_Endpoint_S3Settings) SetV__RowGroupLength(v int32) AWS_DMS_Endpoint_S3Settings {
	t.RowGroupLength = cfz.V(v)
	return t
}

// Set__ServerSideEncryptionKmsKeyId updates property "ServerSideEncryptionKmsKeyId".
func (t AWS_DMS_Endpoint_S3Settings) Set__ServerSideEncryptionKmsKeyId(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.ServerSideEncryptionKmsKeyId = v
	return t
}

// SetV__ServerSideEncryptionKmsKeyId updates property "ServerSideEncryptionKmsKeyId".
func (t AWS_DMS_Endpoint_S3Settings) SetV__ServerSideEncryptionKmsKeyId(v string) AWS_DMS_Endpoint_S3Settings {
	t.ServerSideEncryptionKmsKeyId = cfz.V(v)
	return t
}

// Set__ServiceAccessRoleArn updates property "ServiceAccessRoleArn".
func (t AWS_DMS_Endpoint_S3Settings) Set__ServiceAccessRoleArn(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.ServiceAccessRoleArn = v
	return t
}

// SetV__ServiceAccessRoleArn updates property "ServiceAccessRoleArn".
func (t AWS_DMS_Endpoint_S3Settings) SetV__ServiceAccessRoleArn(v string) AWS_DMS_Endpoint_S3Settings {
	t.ServiceAccessRoleArn = cfz.V(v)
	return t
}

// Set__TimestampColumnName updates property "TimestampColumnName".
func (t AWS_DMS_Endpoint_S3Settings) Set__TimestampColumnName(v cfz.Expression[string]) AWS_DMS_Endpoint_S3Settings {
	t.TimestampColumnName = v
	return t
}

// SetV__TimestampColumnName updates property "TimestampColumnName".
func (t AWS_DMS_Endpoint_S3Settings) SetV__TimestampColumnName(v string) AWS_DMS_Endpoint_S3Settings {
	t.TimestampColumnName = cfz.V(v)
	return t
}

// Set__UseCsvNoSupValue updates property "UseCsvNoSupValue".
func (t AWS_DMS_Endpoint_S3Settings) Set__UseCsvNoSupValue(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.UseCsvNoSupValue = v
	return t
}

// SetV__UseCsvNoSupValue updates property "UseCsvNoSupValue".
func (t AWS_DMS_Endpoint_S3Settings) SetV__UseCsvNoSupValue(v bool) AWS_DMS_Endpoint_S3Settings {
	t.UseCsvNoSupValue = cfz.V(v)
	return t
}

// Set__UseTaskStartTimeForFullLoadTimestamp updates property "UseTaskStartTimeForFullLoadTimestamp".
func (t AWS_DMS_Endpoint_S3Settings) Set__UseTaskStartTimeForFullLoadTimestamp(v cfz.Expression[bool]) AWS_DMS_Endpoint_S3Settings {
	t.UseTaskStartTimeForFullLoadTimestamp = v
	return t
}

// SetV__UseTaskStartTimeForFullLoadTimestamp updates property "UseTaskStartTimeForFullLoadTimestamp".
func (t AWS_DMS_Endpoint_S3Settings) SetV__UseTaskStartTimeForFullLoadTimestamp(v bool) AWS_DMS_Endpoint_S3Settings {
	t.UseTaskStartTimeForFullLoadTimestamp = cfz.V(v)
	return t
}
