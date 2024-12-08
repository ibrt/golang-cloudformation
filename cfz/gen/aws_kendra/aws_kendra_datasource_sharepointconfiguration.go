// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kendra

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Kendra_DataSource_SharePointConfiguration__Type is the CloudFormation type for AWS::Kendra::DataSource.SharePointConfiguration.
	AWS_Kendra_DataSource_SharePointConfiguration__Type = "AWS::Kendra::DataSource.SharePointConfiguration"
)

var (
	// AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Kendra::DataSource.SharePointConfiguration.
	AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap = struct {
		CrawlAttachments       string
		DisableLocalGroups     string
		DocumentTitleFieldName string
		ExclusionPatterns      string
		FieldMappings          string
		InclusionPatterns      string
		SecretArn              string
		SharePointVersion      string
		SslCertificateS3Path   string
		Urls                   string
		UseChangeLog           string
		VpcConfiguration       string
	}{
		CrawlAttachments:       "CrawlAttachments",
		DisableLocalGroups:     "DisableLocalGroups",
		DocumentTitleFieldName: "DocumentTitleFieldName",
		ExclusionPatterns:      "ExclusionPatterns",
		FieldMappings:          "FieldMappings",
		InclusionPatterns:      "InclusionPatterns",
		SecretArn:              "SecretArn",
		SharePointVersion:      "SharePointVersion",
		SslCertificateS3Path:   "SslCertificateS3Path",
		Urls:                   "Urls",
		UseChangeLog:           "UseChangeLog",
		VpcConfiguration:       "VpcConfiguration",
	}

	// AWS_Kendra_DataSource_SharePointConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Kendra::DataSource.SharePointConfiguration.
	AWS_Kendra_DataSource_SharePointConfiguration__PropertiesSlice = []string{
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.CrawlAttachments,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.DisableLocalGroups,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.DocumentTitleFieldName,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.ExclusionPatterns,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.FieldMappings,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.InclusionPatterns,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.SecretArn,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.SharePointVersion,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.SslCertificateS3Path,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.Urls,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.UseChangeLog,
		AWS_Kendra_DataSource_SharePointConfiguration__PropertiesMap.VpcConfiguration,
	}
)

// AWS_Kendra_DataSource_SharePointConfiguration is a binding for AWS::Kendra::DataSource.SharePointConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html
type AWS_Kendra_DataSource_SharePointConfiguration struct {
	// CrawlAttachments is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-crawlattachments
	CrawlAttachments cfz.Expression[bool] `json:"CrawlAttachments,omitempty"`

	// DisableLocalGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-disablelocalgroups
	DisableLocalGroups cfz.Expression[bool] `json:"DisableLocalGroups,omitempty"`

	// DocumentTitleFieldName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-documenttitlefieldname
	DocumentTitleFieldName cfz.Expression[string] `json:"DocumentTitleFieldName,omitempty"`

	// ExclusionPatterns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-exclusionpatterns
	ExclusionPatterns cfz.ExpressionSlice[string] `json:"ExclusionPatterns,omitempty"`

	// FieldMappings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-fieldmappings
	FieldMappings cfz.ExpressionSlice[AWS_Kendra_DataSource_DataSourceToIndexFieldMapping] `json:"FieldMappings,omitempty"`

	// InclusionPatterns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-inclusionpatterns
	InclusionPatterns cfz.ExpressionSlice[string] `json:"InclusionPatterns,omitempty"`

	// SecretArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-secretarn
	SecretArn cfz.Expression[string] `json:"SecretArn,omitempty"`

	// SharePointVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-sharepointversion
	SharePointVersion cfz.Expression[string] `json:"SharePointVersion,omitempty"`

	// SslCertificateS3Path is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-sslcertificates3path
	SslCertificateS3Path cfz.Expression[AWS_Kendra_DataSource_S3Path] `json:"SslCertificateS3Path,omitempty"`

	// Urls is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-urls
	Urls cfz.ExpressionSlice[string] `json:"Urls,omitempty"`

	// UseChangeLog is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-usechangelog
	UseChangeLog cfz.Expression[bool] `json:"UseChangeLog,omitempty"`

	// VpcConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-vpcconfiguration
	VpcConfiguration cfz.Expression[AWS_Kendra_DataSource_DataSourceVpcConfiguration] `json:"VpcConfiguration,omitempty"`
}

// New__AWS_Kendra_DataSource_SharePointConfiguration initializes a new AWS_Kendra_DataSource_SharePointConfiguration.
func New__AWS_Kendra_DataSource_SharePointConfiguration() AWS_Kendra_DataSource_SharePointConfiguration {
	return AWS_Kendra_DataSource_SharePointConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Kendra_DataSource_SharePointConfiguration) GetType() string {
	return AWS_Kendra_DataSource_SharePointConfiguration__Type
}

// Set__CrawlAttachments updates property "CrawlAttachments".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__CrawlAttachments(v cfz.Expression[bool]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.CrawlAttachments = v
	return t
}

// SetV__CrawlAttachments updates property "CrawlAttachments".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetV__CrawlAttachments(v bool) AWS_Kendra_DataSource_SharePointConfiguration {
	t.CrawlAttachments = cfz.V(v)
	return t
}

// Set__DisableLocalGroups updates property "DisableLocalGroups".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__DisableLocalGroups(v cfz.Expression[bool]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.DisableLocalGroups = v
	return t
}

// SetV__DisableLocalGroups updates property "DisableLocalGroups".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetV__DisableLocalGroups(v bool) AWS_Kendra_DataSource_SharePointConfiguration {
	t.DisableLocalGroups = cfz.V(v)
	return t
}

// Set__DocumentTitleFieldName updates property "DocumentTitleFieldName".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__DocumentTitleFieldName(v cfz.Expression[string]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.DocumentTitleFieldName = v
	return t
}

// SetV__DocumentTitleFieldName updates property "DocumentTitleFieldName".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetV__DocumentTitleFieldName(v string) AWS_Kendra_DataSource_SharePointConfiguration {
	t.DocumentTitleFieldName = cfz.V(v)
	return t
}

// Set__ExclusionPatterns updates property "ExclusionPatterns".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__ExclusionPatterns(v cfz.ExpressionSlice[string]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.ExclusionPatterns = v
	return t
}

// SetS__ExclusionPatterns updates property "ExclusionPatterns".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetS__ExclusionPatterns(v ...cfz.Expression[string]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.ExclusionPatterns = cfz.S(v...)
	return t
}

// SetSV__ExclusionPatterns updates property "ExclusionPatterns".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetSV__ExclusionPatterns(v ...string) AWS_Kendra_DataSource_SharePointConfiguration {
	t.ExclusionPatterns = cfz.SV(v...)
	return t
}

// Set__FieldMappings updates property "FieldMappings".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__FieldMappings(v cfz.ExpressionSlice[AWS_Kendra_DataSource_DataSourceToIndexFieldMapping]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.FieldMappings = v
	return t
}

// SetS__FieldMappings updates property "FieldMappings".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetS__FieldMappings(v ...cfz.Expression[AWS_Kendra_DataSource_DataSourceToIndexFieldMapping]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.FieldMappings = cfz.S(v...)
	return t
}

// SetSV__FieldMappings updates property "FieldMappings".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetSV__FieldMappings(v ...AWS_Kendra_DataSource_DataSourceToIndexFieldMapping) AWS_Kendra_DataSource_SharePointConfiguration {
	t.FieldMappings = cfz.SV(v...)
	return t
}

// Set__InclusionPatterns updates property "InclusionPatterns".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__InclusionPatterns(v cfz.ExpressionSlice[string]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.InclusionPatterns = v
	return t
}

// SetS__InclusionPatterns updates property "InclusionPatterns".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetS__InclusionPatterns(v ...cfz.Expression[string]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.InclusionPatterns = cfz.S(v...)
	return t
}

// SetSV__InclusionPatterns updates property "InclusionPatterns".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetSV__InclusionPatterns(v ...string) AWS_Kendra_DataSource_SharePointConfiguration {
	t.InclusionPatterns = cfz.SV(v...)
	return t
}

// Set__SecretArn updates property "SecretArn".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__SecretArn(v cfz.Expression[string]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.SecretArn = v
	return t
}

// SetV__SecretArn updates property "SecretArn".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetV__SecretArn(v string) AWS_Kendra_DataSource_SharePointConfiguration {
	t.SecretArn = cfz.V(v)
	return t
}

// Set__SharePointVersion updates property "SharePointVersion".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__SharePointVersion(v cfz.Expression[string]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.SharePointVersion = v
	return t
}

// SetV__SharePointVersion updates property "SharePointVersion".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetV__SharePointVersion(v string) AWS_Kendra_DataSource_SharePointConfiguration {
	t.SharePointVersion = cfz.V(v)
	return t
}

// Set__SslCertificateS3Path updates property "SslCertificateS3Path".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__SslCertificateS3Path(v cfz.Expression[AWS_Kendra_DataSource_S3Path]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.SslCertificateS3Path = v
	return t
}

// SetV__SslCertificateS3Path updates property "SslCertificateS3Path".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetV__SslCertificateS3Path(v AWS_Kendra_DataSource_S3Path) AWS_Kendra_DataSource_SharePointConfiguration {
	t.SslCertificateS3Path = cfz.V(v)
	return t
}

// Set__Urls updates property "Urls".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__Urls(v cfz.ExpressionSlice[string]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.Urls = v
	return t
}

// SetS__Urls updates property "Urls".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetS__Urls(v ...cfz.Expression[string]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.Urls = cfz.S(v...)
	return t
}

// SetSV__Urls updates property "Urls".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetSV__Urls(v ...string) AWS_Kendra_DataSource_SharePointConfiguration {
	t.Urls = cfz.SV(v...)
	return t
}

// Set__UseChangeLog updates property "UseChangeLog".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__UseChangeLog(v cfz.Expression[bool]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.UseChangeLog = v
	return t
}

// SetV__UseChangeLog updates property "UseChangeLog".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetV__UseChangeLog(v bool) AWS_Kendra_DataSource_SharePointConfiguration {
	t.UseChangeLog = cfz.V(v)
	return t
}

// Set__VpcConfiguration updates property "VpcConfiguration".
func (t AWS_Kendra_DataSource_SharePointConfiguration) Set__VpcConfiguration(v cfz.Expression[AWS_Kendra_DataSource_DataSourceVpcConfiguration]) AWS_Kendra_DataSource_SharePointConfiguration {
	t.VpcConfiguration = v
	return t
}

// SetV__VpcConfiguration updates property "VpcConfiguration".
func (t AWS_Kendra_DataSource_SharePointConfiguration) SetV__VpcConfiguration(v AWS_Kendra_DataSource_DataSourceVpcConfiguration) AWS_Kendra_DataSource_SharePointConfiguration {
	t.VpcConfiguration = cfz.V(v)
	return t
}
