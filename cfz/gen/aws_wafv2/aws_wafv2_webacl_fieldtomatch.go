// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_FieldToMatch__Type is the CloudFormation type for AWS::WAFv2::WebACL.FieldToMatch.
	AWS_WAFv2_WebACL_FieldToMatch__Type = "AWS::WAFv2::WebACL.FieldToMatch"
)

var (
	// AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.FieldToMatch.
	AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap = struct {
		AllQueryArguments   string
		Body                string
		Cookies             string
		Headers             string
		JA3Fingerprint      string
		JsonBody            string
		Method              string
		QueryString         string
		SingleHeader        string
		SingleQueryArgument string
		UriPath             string
	}{
		AllQueryArguments:   "AllQueryArguments",
		Body:                "Body",
		Cookies:             "Cookies",
		Headers:             "Headers",
		JA3Fingerprint:      "JA3Fingerprint",
		JsonBody:            "JsonBody",
		Method:              "Method",
		QueryString:         "QueryString",
		SingleHeader:        "SingleHeader",
		SingleQueryArgument: "SingleQueryArgument",
		UriPath:             "UriPath",
	}

	// AWS_WAFv2_WebACL_FieldToMatch__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.FieldToMatch.
	AWS_WAFv2_WebACL_FieldToMatch__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.AllQueryArguments,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.Body,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.Cookies,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.Headers,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.JA3Fingerprint,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.JsonBody,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.Method,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.QueryString,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.SingleHeader,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.SingleQueryArgument,
		AWS_WAFv2_WebACL_FieldToMatch__PropertiesMap.UriPath,
	}
)

// AWS_WAFv2_WebACL_FieldToMatch is a binding for AWS::WAFv2::WebACL.FieldToMatch.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html
type AWS_WAFv2_WebACL_FieldToMatch struct {
	// AllQueryArguments is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-allqueryarguments
	AllQueryArguments cfz.Expression[json.RawMessage] `json:"AllQueryArguments,omitempty"`

	// Body is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-body
	Body cfz.Expression[AWS_WAFv2_WebACL_Body] `json:"Body,omitempty"`

	// Cookies is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-cookies
	Cookies cfz.Expression[AWS_WAFv2_WebACL_Cookies] `json:"Cookies,omitempty"`

	// Headers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-headers
	Headers cfz.Expression[AWS_WAFv2_WebACL_Headers] `json:"Headers,omitempty"`

	// JA3Fingerprint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-ja3fingerprint
	JA3Fingerprint cfz.Expression[AWS_WAFv2_WebACL_JA3Fingerprint] `json:"JA3Fingerprint,omitempty"`

	// JsonBody is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-jsonbody
	JsonBody cfz.Expression[AWS_WAFv2_WebACL_JsonBody] `json:"JsonBody,omitempty"`

	// Method is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-method
	Method cfz.Expression[json.RawMessage] `json:"Method,omitempty"`

	// QueryString is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-querystring
	QueryString cfz.Expression[json.RawMessage] `json:"QueryString,omitempty"`

	// SingleHeader is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-singleheader
	SingleHeader cfz.Expression[AWS_WAFv2_WebACL_SingleHeader] `json:"SingleHeader,omitempty"`

	// SingleQueryArgument is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-singlequeryargument
	SingleQueryArgument cfz.Expression[AWS_WAFv2_WebACL_SingleQueryArgument] `json:"SingleQueryArgument,omitempty"`

	// UriPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-uripath
	UriPath cfz.Expression[json.RawMessage] `json:"UriPath,omitempty"`
}

// New__AWS_WAFv2_WebACL_FieldToMatch initializes a new AWS_WAFv2_WebACL_FieldToMatch.
func New__AWS_WAFv2_WebACL_FieldToMatch() AWS_WAFv2_WebACL_FieldToMatch {
	return AWS_WAFv2_WebACL_FieldToMatch{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_FieldToMatch) GetType() string {
	return AWS_WAFv2_WebACL_FieldToMatch__Type
}

// Set__AllQueryArguments updates property "AllQueryArguments".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__AllQueryArguments(v cfz.Expression[json.RawMessage]) AWS_WAFv2_WebACL_FieldToMatch {
	t.AllQueryArguments = v
	return t
}

// SetV__AllQueryArguments updates property "AllQueryArguments".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__AllQueryArguments(v json.RawMessage) AWS_WAFv2_WebACL_FieldToMatch {
	t.AllQueryArguments = cfz.V(v)
	return t
}

// Set__Body updates property "Body".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__Body(v cfz.Expression[AWS_WAFv2_WebACL_Body]) AWS_WAFv2_WebACL_FieldToMatch {
	t.Body = v
	return t
}

// SetV__Body updates property "Body".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__Body(v AWS_WAFv2_WebACL_Body) AWS_WAFv2_WebACL_FieldToMatch {
	t.Body = cfz.V(v)
	return t
}

// Set__Cookies updates property "Cookies".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__Cookies(v cfz.Expression[AWS_WAFv2_WebACL_Cookies]) AWS_WAFv2_WebACL_FieldToMatch {
	t.Cookies = v
	return t
}

// SetV__Cookies updates property "Cookies".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__Cookies(v AWS_WAFv2_WebACL_Cookies) AWS_WAFv2_WebACL_FieldToMatch {
	t.Cookies = cfz.V(v)
	return t
}

// Set__Headers updates property "Headers".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__Headers(v cfz.Expression[AWS_WAFv2_WebACL_Headers]) AWS_WAFv2_WebACL_FieldToMatch {
	t.Headers = v
	return t
}

// SetV__Headers updates property "Headers".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__Headers(v AWS_WAFv2_WebACL_Headers) AWS_WAFv2_WebACL_FieldToMatch {
	t.Headers = cfz.V(v)
	return t
}

// Set__JA3Fingerprint updates property "JA3Fingerprint".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__JA3Fingerprint(v cfz.Expression[AWS_WAFv2_WebACL_JA3Fingerprint]) AWS_WAFv2_WebACL_FieldToMatch {
	t.JA3Fingerprint = v
	return t
}

// SetV__JA3Fingerprint updates property "JA3Fingerprint".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__JA3Fingerprint(v AWS_WAFv2_WebACL_JA3Fingerprint) AWS_WAFv2_WebACL_FieldToMatch {
	t.JA3Fingerprint = cfz.V(v)
	return t
}

// Set__JsonBody updates property "JsonBody".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__JsonBody(v cfz.Expression[AWS_WAFv2_WebACL_JsonBody]) AWS_WAFv2_WebACL_FieldToMatch {
	t.JsonBody = v
	return t
}

// SetV__JsonBody updates property "JsonBody".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__JsonBody(v AWS_WAFv2_WebACL_JsonBody) AWS_WAFv2_WebACL_FieldToMatch {
	t.JsonBody = cfz.V(v)
	return t
}

// Set__Method updates property "Method".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__Method(v cfz.Expression[json.RawMessage]) AWS_WAFv2_WebACL_FieldToMatch {
	t.Method = v
	return t
}

// SetV__Method updates property "Method".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__Method(v json.RawMessage) AWS_WAFv2_WebACL_FieldToMatch {
	t.Method = cfz.V(v)
	return t
}

// Set__QueryString updates property "QueryString".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__QueryString(v cfz.Expression[json.RawMessage]) AWS_WAFv2_WebACL_FieldToMatch {
	t.QueryString = v
	return t
}

// SetV__QueryString updates property "QueryString".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__QueryString(v json.RawMessage) AWS_WAFv2_WebACL_FieldToMatch {
	t.QueryString = cfz.V(v)
	return t
}

// Set__SingleHeader updates property "SingleHeader".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__SingleHeader(v cfz.Expression[AWS_WAFv2_WebACL_SingleHeader]) AWS_WAFv2_WebACL_FieldToMatch {
	t.SingleHeader = v
	return t
}

// SetV__SingleHeader updates property "SingleHeader".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__SingleHeader(v AWS_WAFv2_WebACL_SingleHeader) AWS_WAFv2_WebACL_FieldToMatch {
	t.SingleHeader = cfz.V(v)
	return t
}

// Set__SingleQueryArgument updates property "SingleQueryArgument".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__SingleQueryArgument(v cfz.Expression[AWS_WAFv2_WebACL_SingleQueryArgument]) AWS_WAFv2_WebACL_FieldToMatch {
	t.SingleQueryArgument = v
	return t
}

// SetV__SingleQueryArgument updates property "SingleQueryArgument".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__SingleQueryArgument(v AWS_WAFv2_WebACL_SingleQueryArgument) AWS_WAFv2_WebACL_FieldToMatch {
	t.SingleQueryArgument = cfz.V(v)
	return t
}

// Set__UriPath updates property "UriPath".
func (t AWS_WAFv2_WebACL_FieldToMatch) Set__UriPath(v cfz.Expression[json.RawMessage]) AWS_WAFv2_WebACL_FieldToMatch {
	t.UriPath = v
	return t
}

// SetV__UriPath updates property "UriPath".
func (t AWS_WAFv2_WebACL_FieldToMatch) SetV__UriPath(v json.RawMessage) AWS_WAFv2_WebACL_FieldToMatch {
	t.UriPath = cfz.V(v)
	return t
}
