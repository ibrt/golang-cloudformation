// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_secretsmanager

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SecretsManager_Secret_GenerateSecretString__Type is the CloudFormation type for AWS::SecretsManager::Secret.GenerateSecretString.
	AWS_SecretsManager_Secret_GenerateSecretString__Type = "AWS::SecretsManager::Secret.GenerateSecretString"
)

var (
	// AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap reports all the CloudFormation properties for AWS::SecretsManager::Secret.GenerateSecretString.
	AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap = struct {
		ExcludeCharacters       string
		ExcludeLowercase        string
		ExcludeNumbers          string
		ExcludePunctuation      string
		ExcludeUppercase        string
		GenerateStringKey       string
		IncludeSpace            string
		PasswordLength          string
		RequireEachIncludedType string
		SecretStringTemplate    string
	}{
		ExcludeCharacters:       "ExcludeCharacters",
		ExcludeLowercase:        "ExcludeLowercase",
		ExcludeNumbers:          "ExcludeNumbers",
		ExcludePunctuation:      "ExcludePunctuation",
		ExcludeUppercase:        "ExcludeUppercase",
		GenerateStringKey:       "GenerateStringKey",
		IncludeSpace:            "IncludeSpace",
		PasswordLength:          "PasswordLength",
		RequireEachIncludedType: "RequireEachIncludedType",
		SecretStringTemplate:    "SecretStringTemplate",
	}

	// AWS_SecretsManager_Secret_GenerateSecretString__PropertiesSlice reports all the CloudFormation properties for AWS::SecretsManager::Secret.GenerateSecretString.
	AWS_SecretsManager_Secret_GenerateSecretString__PropertiesSlice = []string{
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.ExcludeCharacters,
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.ExcludeLowercase,
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.ExcludeNumbers,
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.ExcludePunctuation,
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.ExcludeUppercase,
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.GenerateStringKey,
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.IncludeSpace,
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.PasswordLength,
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.RequireEachIncludedType,
		AWS_SecretsManager_Secret_GenerateSecretString__PropertiesMap.SecretStringTemplate,
	}
)

// AWS_SecretsManager_Secret_GenerateSecretString is a binding for AWS::SecretsManager::Secret.GenerateSecretString.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html
type AWS_SecretsManager_Secret_GenerateSecretString struct {
	// ExcludeCharacters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludecharacters
	ExcludeCharacters cfz.Expression[string] `json:"ExcludeCharacters,omitempty"`

	// ExcludeLowercase is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludelowercase
	ExcludeLowercase cfz.Expression[bool] `json:"ExcludeLowercase,omitempty"`

	// ExcludeNumbers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludenumbers
	ExcludeNumbers cfz.Expression[bool] `json:"ExcludeNumbers,omitempty"`

	// ExcludePunctuation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludepunctuation
	ExcludePunctuation cfz.Expression[bool] `json:"ExcludePunctuation,omitempty"`

	// ExcludeUppercase is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludeuppercase
	ExcludeUppercase cfz.Expression[bool] `json:"ExcludeUppercase,omitempty"`

	// GenerateStringKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-generatestringkey
	GenerateStringKey cfz.Expression[string] `json:"GenerateStringKey,omitempty"`

	// IncludeSpace is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-includespace
	IncludeSpace cfz.Expression[bool] `json:"IncludeSpace,omitempty"`

	// PasswordLength is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-passwordlength
	PasswordLength cfz.Expression[int32] `json:"PasswordLength,omitempty"`

	// RequireEachIncludedType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-requireeachincludedtype
	RequireEachIncludedType cfz.Expression[bool] `json:"RequireEachIncludedType,omitempty"`

	// SecretStringTemplate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-secretstringtemplate
	SecretStringTemplate cfz.Expression[string] `json:"SecretStringTemplate,omitempty"`
}

// New__AWS_SecretsManager_Secret_GenerateSecretString initializes a new AWS_SecretsManager_Secret_GenerateSecretString.
func New__AWS_SecretsManager_Secret_GenerateSecretString() AWS_SecretsManager_Secret_GenerateSecretString {
	return AWS_SecretsManager_Secret_GenerateSecretString{}
}

// GetType returns the CloudFormation type.
func (AWS_SecretsManager_Secret_GenerateSecretString) GetType() string {
	return AWS_SecretsManager_Secret_GenerateSecretString__Type
}

// Set__ExcludeCharacters updates property "ExcludeCharacters".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__ExcludeCharacters(v cfz.Expression[string]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludeCharacters = v
	return t
}

// SetV__ExcludeCharacters updates property "ExcludeCharacters".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__ExcludeCharacters(v string) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludeCharacters = cfz.V(v)
	return t
}

// Set__ExcludeLowercase updates property "ExcludeLowercase".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__ExcludeLowercase(v cfz.Expression[bool]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludeLowercase = v
	return t
}

// SetV__ExcludeLowercase updates property "ExcludeLowercase".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__ExcludeLowercase(v bool) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludeLowercase = cfz.V(v)
	return t
}

// Set__ExcludeNumbers updates property "ExcludeNumbers".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__ExcludeNumbers(v cfz.Expression[bool]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludeNumbers = v
	return t
}

// SetV__ExcludeNumbers updates property "ExcludeNumbers".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__ExcludeNumbers(v bool) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludeNumbers = cfz.V(v)
	return t
}

// Set__ExcludePunctuation updates property "ExcludePunctuation".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__ExcludePunctuation(v cfz.Expression[bool]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludePunctuation = v
	return t
}

// SetV__ExcludePunctuation updates property "ExcludePunctuation".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__ExcludePunctuation(v bool) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludePunctuation = cfz.V(v)
	return t
}

// Set__ExcludeUppercase updates property "ExcludeUppercase".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__ExcludeUppercase(v cfz.Expression[bool]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludeUppercase = v
	return t
}

// SetV__ExcludeUppercase updates property "ExcludeUppercase".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__ExcludeUppercase(v bool) AWS_SecretsManager_Secret_GenerateSecretString {
	t.ExcludeUppercase = cfz.V(v)
	return t
}

// Set__GenerateStringKey updates property "GenerateStringKey".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__GenerateStringKey(v cfz.Expression[string]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.GenerateStringKey = v
	return t
}

// SetV__GenerateStringKey updates property "GenerateStringKey".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__GenerateStringKey(v string) AWS_SecretsManager_Secret_GenerateSecretString {
	t.GenerateStringKey = cfz.V(v)
	return t
}

// Set__IncludeSpace updates property "IncludeSpace".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__IncludeSpace(v cfz.Expression[bool]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.IncludeSpace = v
	return t
}

// SetV__IncludeSpace updates property "IncludeSpace".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__IncludeSpace(v bool) AWS_SecretsManager_Secret_GenerateSecretString {
	t.IncludeSpace = cfz.V(v)
	return t
}

// Set__PasswordLength updates property "PasswordLength".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__PasswordLength(v cfz.Expression[int32]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.PasswordLength = v
	return t
}

// SetV__PasswordLength updates property "PasswordLength".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__PasswordLength(v int32) AWS_SecretsManager_Secret_GenerateSecretString {
	t.PasswordLength = cfz.V(v)
	return t
}

// Set__RequireEachIncludedType updates property "RequireEachIncludedType".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__RequireEachIncludedType(v cfz.Expression[bool]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.RequireEachIncludedType = v
	return t
}

// SetV__RequireEachIncludedType updates property "RequireEachIncludedType".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__RequireEachIncludedType(v bool) AWS_SecretsManager_Secret_GenerateSecretString {
	t.RequireEachIncludedType = cfz.V(v)
	return t
}

// Set__SecretStringTemplate updates property "SecretStringTemplate".
func (t AWS_SecretsManager_Secret_GenerateSecretString) Set__SecretStringTemplate(v cfz.Expression[string]) AWS_SecretsManager_Secret_GenerateSecretString {
	t.SecretStringTemplate = v
	return t
}

// SetV__SecretStringTemplate updates property "SecretStringTemplate".
func (t AWS_SecretsManager_Secret_GenerateSecretString) SetV__SecretStringTemplate(v string) AWS_SecretsManager_Secret_GenerateSecretString {
	t.SecretStringTemplate = cfz.V(v)
	return t
}
