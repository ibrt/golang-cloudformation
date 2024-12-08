// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_transfer

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Transfer_Connector_As2Config__Type is the CloudFormation type for AWS::Transfer::Connector.As2Config.
	AWS_Transfer_Connector_As2Config__Type = "AWS::Transfer::Connector.As2Config"
)

var (
	// AWS_Transfer_Connector_As2Config__PropertiesMap reports all the CloudFormation properties for AWS::Transfer::Connector.As2Config.
	AWS_Transfer_Connector_As2Config__PropertiesMap = struct {
		BasicAuthSecretId   string
		Compression         string
		EncryptionAlgorithm string
		LocalProfileId      string
		MdnResponse         string
		MdnSigningAlgorithm string
		MessageSubject      string
		PartnerProfileId    string
		SigningAlgorithm    string
	}{
		BasicAuthSecretId:   "BasicAuthSecretId",
		Compression:         "Compression",
		EncryptionAlgorithm: "EncryptionAlgorithm",
		LocalProfileId:      "LocalProfileId",
		MdnResponse:         "MdnResponse",
		MdnSigningAlgorithm: "MdnSigningAlgorithm",
		MessageSubject:      "MessageSubject",
		PartnerProfileId:    "PartnerProfileId",
		SigningAlgorithm:    "SigningAlgorithm",
	}

	// AWS_Transfer_Connector_As2Config__PropertiesSlice reports all the CloudFormation properties for AWS::Transfer::Connector.As2Config.
	AWS_Transfer_Connector_As2Config__PropertiesSlice = []string{
		AWS_Transfer_Connector_As2Config__PropertiesMap.BasicAuthSecretId,
		AWS_Transfer_Connector_As2Config__PropertiesMap.Compression,
		AWS_Transfer_Connector_As2Config__PropertiesMap.EncryptionAlgorithm,
		AWS_Transfer_Connector_As2Config__PropertiesMap.LocalProfileId,
		AWS_Transfer_Connector_As2Config__PropertiesMap.MdnResponse,
		AWS_Transfer_Connector_As2Config__PropertiesMap.MdnSigningAlgorithm,
		AWS_Transfer_Connector_As2Config__PropertiesMap.MessageSubject,
		AWS_Transfer_Connector_As2Config__PropertiesMap.PartnerProfileId,
		AWS_Transfer_Connector_As2Config__PropertiesMap.SigningAlgorithm,
	}
)

// AWS_Transfer_Connector_As2Config is a binding for AWS::Transfer::Connector.As2Config.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html
type AWS_Transfer_Connector_As2Config struct {
	// BasicAuthSecretId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-basicauthsecretid
	BasicAuthSecretId cfz.Expression[string] `json:"BasicAuthSecretId,omitempty"`

	// Compression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-compression
	Compression cfz.Expression[string] `json:"Compression,omitempty"`

	// EncryptionAlgorithm is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-encryptionalgorithm
	EncryptionAlgorithm cfz.Expression[string] `json:"EncryptionAlgorithm,omitempty"`

	// LocalProfileId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-localprofileid
	LocalProfileId cfz.Expression[string] `json:"LocalProfileId,omitempty"`

	// MdnResponse is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-mdnresponse
	MdnResponse cfz.Expression[string] `json:"MdnResponse,omitempty"`

	// MdnSigningAlgorithm is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-mdnsigningalgorithm
	MdnSigningAlgorithm cfz.Expression[string] `json:"MdnSigningAlgorithm,omitempty"`

	// MessageSubject is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-messagesubject
	MessageSubject cfz.Expression[string] `json:"MessageSubject,omitempty"`

	// PartnerProfileId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-partnerprofileid
	PartnerProfileId cfz.Expression[string] `json:"PartnerProfileId,omitempty"`

	// SigningAlgorithm is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-as2config.html#cfn-transfer-connector-as2config-signingalgorithm
	SigningAlgorithm cfz.Expression[string] `json:"SigningAlgorithm,omitempty"`
}

// New__AWS_Transfer_Connector_As2Config initializes a new AWS_Transfer_Connector_As2Config.
func New__AWS_Transfer_Connector_As2Config() AWS_Transfer_Connector_As2Config {
	return AWS_Transfer_Connector_As2Config{}
}

// GetType returns the CloudFormation type.
func (AWS_Transfer_Connector_As2Config) GetType() string {
	return AWS_Transfer_Connector_As2Config__Type
}

// Set__BasicAuthSecretId updates property "BasicAuthSecretId".
func (t AWS_Transfer_Connector_As2Config) Set__BasicAuthSecretId(v cfz.Expression[string]) AWS_Transfer_Connector_As2Config {
	t.BasicAuthSecretId = v
	return t
}

// SetV__BasicAuthSecretId updates property "BasicAuthSecretId".
func (t AWS_Transfer_Connector_As2Config) SetV__BasicAuthSecretId(v string) AWS_Transfer_Connector_As2Config {
	t.BasicAuthSecretId = cfz.V(v)
	return t
}

// Set__Compression updates property "Compression".
func (t AWS_Transfer_Connector_As2Config) Set__Compression(v cfz.Expression[string]) AWS_Transfer_Connector_As2Config {
	t.Compression = v
	return t
}

// SetV__Compression updates property "Compression".
func (t AWS_Transfer_Connector_As2Config) SetV__Compression(v string) AWS_Transfer_Connector_As2Config {
	t.Compression = cfz.V(v)
	return t
}

// Set__EncryptionAlgorithm updates property "EncryptionAlgorithm".
func (t AWS_Transfer_Connector_As2Config) Set__EncryptionAlgorithm(v cfz.Expression[string]) AWS_Transfer_Connector_As2Config {
	t.EncryptionAlgorithm = v
	return t
}

// SetV__EncryptionAlgorithm updates property "EncryptionAlgorithm".
func (t AWS_Transfer_Connector_As2Config) SetV__EncryptionAlgorithm(v string) AWS_Transfer_Connector_As2Config {
	t.EncryptionAlgorithm = cfz.V(v)
	return t
}

// Set__LocalProfileId updates property "LocalProfileId".
func (t AWS_Transfer_Connector_As2Config) Set__LocalProfileId(v cfz.Expression[string]) AWS_Transfer_Connector_As2Config {
	t.LocalProfileId = v
	return t
}

// SetV__LocalProfileId updates property "LocalProfileId".
func (t AWS_Transfer_Connector_As2Config) SetV__LocalProfileId(v string) AWS_Transfer_Connector_As2Config {
	t.LocalProfileId = cfz.V(v)
	return t
}

// Set__MdnResponse updates property "MdnResponse".
func (t AWS_Transfer_Connector_As2Config) Set__MdnResponse(v cfz.Expression[string]) AWS_Transfer_Connector_As2Config {
	t.MdnResponse = v
	return t
}

// SetV__MdnResponse updates property "MdnResponse".
func (t AWS_Transfer_Connector_As2Config) SetV__MdnResponse(v string) AWS_Transfer_Connector_As2Config {
	t.MdnResponse = cfz.V(v)
	return t
}

// Set__MdnSigningAlgorithm updates property "MdnSigningAlgorithm".
func (t AWS_Transfer_Connector_As2Config) Set__MdnSigningAlgorithm(v cfz.Expression[string]) AWS_Transfer_Connector_As2Config {
	t.MdnSigningAlgorithm = v
	return t
}

// SetV__MdnSigningAlgorithm updates property "MdnSigningAlgorithm".
func (t AWS_Transfer_Connector_As2Config) SetV__MdnSigningAlgorithm(v string) AWS_Transfer_Connector_As2Config {
	t.MdnSigningAlgorithm = cfz.V(v)
	return t
}

// Set__MessageSubject updates property "MessageSubject".
func (t AWS_Transfer_Connector_As2Config) Set__MessageSubject(v cfz.Expression[string]) AWS_Transfer_Connector_As2Config {
	t.MessageSubject = v
	return t
}

// SetV__MessageSubject updates property "MessageSubject".
func (t AWS_Transfer_Connector_As2Config) SetV__MessageSubject(v string) AWS_Transfer_Connector_As2Config {
	t.MessageSubject = cfz.V(v)
	return t
}

// Set__PartnerProfileId updates property "PartnerProfileId".
func (t AWS_Transfer_Connector_As2Config) Set__PartnerProfileId(v cfz.Expression[string]) AWS_Transfer_Connector_As2Config {
	t.PartnerProfileId = v
	return t
}

// SetV__PartnerProfileId updates property "PartnerProfileId".
func (t AWS_Transfer_Connector_As2Config) SetV__PartnerProfileId(v string) AWS_Transfer_Connector_As2Config {
	t.PartnerProfileId = cfz.V(v)
	return t
}

// Set__SigningAlgorithm updates property "SigningAlgorithm".
func (t AWS_Transfer_Connector_As2Config) Set__SigningAlgorithm(v cfz.Expression[string]) AWS_Transfer_Connector_As2Config {
	t.SigningAlgorithm = v
	return t
}

// SetV__SigningAlgorithm updates property "SigningAlgorithm".
func (t AWS_Transfer_Connector_As2Config) SetV__SigningAlgorithm(v string) AWS_Transfer_Connector_As2Config {
	t.SigningAlgorithm = cfz.V(v)
	return t
}
