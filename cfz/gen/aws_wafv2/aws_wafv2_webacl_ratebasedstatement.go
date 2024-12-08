// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_RateBasedStatement__Type is the CloudFormation type for AWS::WAFv2::WebACL.RateBasedStatement.
	AWS_WAFv2_WebACL_RateBasedStatement__Type = "AWS::WAFv2::WebACL.RateBasedStatement"
)

var (
	// AWS_WAFv2_WebACL_RateBasedStatement__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.RateBasedStatement.
	AWS_WAFv2_WebACL_RateBasedStatement__PropertiesMap = struct {
		AggregateKeyType    string
		CustomKeys          string
		EvaluationWindowSec string
		ForwardedIPConfig   string
		Limit               string
		ScopeDownStatement  string
	}{
		AggregateKeyType:    "AggregateKeyType",
		CustomKeys:          "CustomKeys",
		EvaluationWindowSec: "EvaluationWindowSec",
		ForwardedIPConfig:   "ForwardedIPConfig",
		Limit:               "Limit",
		ScopeDownStatement:  "ScopeDownStatement",
	}

	// AWS_WAFv2_WebACL_RateBasedStatement__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.RateBasedStatement.
	AWS_WAFv2_WebACL_RateBasedStatement__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_RateBasedStatement__PropertiesMap.AggregateKeyType,
		AWS_WAFv2_WebACL_RateBasedStatement__PropertiesMap.CustomKeys,
		AWS_WAFv2_WebACL_RateBasedStatement__PropertiesMap.EvaluationWindowSec,
		AWS_WAFv2_WebACL_RateBasedStatement__PropertiesMap.ForwardedIPConfig,
		AWS_WAFv2_WebACL_RateBasedStatement__PropertiesMap.Limit,
		AWS_WAFv2_WebACL_RateBasedStatement__PropertiesMap.ScopeDownStatement,
	}
)

// AWS_WAFv2_WebACL_RateBasedStatement is a binding for AWS::WAFv2::WebACL.RateBasedStatement.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html
type AWS_WAFv2_WebACL_RateBasedStatement struct {
	// AggregateKeyType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-aggregatekeytype
	AggregateKeyType cfz.Expression[string] `json:"AggregateKeyType,omitempty"`

	// CustomKeys is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-customkeys
	CustomKeys cfz.ExpressionSlice[AWS_WAFv2_WebACL_RateBasedStatementCustomKey] `json:"CustomKeys,omitempty"`

	// EvaluationWindowSec is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-evaluationwindowsec
	EvaluationWindowSec cfz.Expression[int32] `json:"EvaluationWindowSec,omitempty"`

	// ForwardedIPConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-forwardedipconfig
	ForwardedIPConfig cfz.Expression[AWS_WAFv2_WebACL_ForwardedIPConfiguration] `json:"ForwardedIPConfig,omitempty"`

	// Limit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-limit
	Limit cfz.Expression[int32] `json:"Limit,omitempty"`

	// ScopeDownStatement is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-scopedownstatement
	ScopeDownStatement cfz.Expression[AWS_WAFv2_WebACL_Statement] `json:"ScopeDownStatement,omitempty"`
}

// New__AWS_WAFv2_WebACL_RateBasedStatement initializes a new AWS_WAFv2_WebACL_RateBasedStatement.
func New__AWS_WAFv2_WebACL_RateBasedStatement() AWS_WAFv2_WebACL_RateBasedStatement {
	return AWS_WAFv2_WebACL_RateBasedStatement{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_RateBasedStatement) GetType() string {
	return AWS_WAFv2_WebACL_RateBasedStatement__Type
}

// Set__AggregateKeyType updates property "AggregateKeyType".
func (t AWS_WAFv2_WebACL_RateBasedStatement) Set__AggregateKeyType(v cfz.Expression[string]) AWS_WAFv2_WebACL_RateBasedStatement {
	t.AggregateKeyType = v
	return t
}

// SetV__AggregateKeyType updates property "AggregateKeyType".
func (t AWS_WAFv2_WebACL_RateBasedStatement) SetV__AggregateKeyType(v string) AWS_WAFv2_WebACL_RateBasedStatement {
	t.AggregateKeyType = cfz.V(v)
	return t
}

// Set__CustomKeys updates property "CustomKeys".
func (t AWS_WAFv2_WebACL_RateBasedStatement) Set__CustomKeys(v cfz.ExpressionSlice[AWS_WAFv2_WebACL_RateBasedStatementCustomKey]) AWS_WAFv2_WebACL_RateBasedStatement {
	t.CustomKeys = v
	return t
}

// SetS__CustomKeys updates property "CustomKeys".
func (t AWS_WAFv2_WebACL_RateBasedStatement) SetS__CustomKeys(v ...cfz.Expression[AWS_WAFv2_WebACL_RateBasedStatementCustomKey]) AWS_WAFv2_WebACL_RateBasedStatement {
	t.CustomKeys = cfz.S(v...)
	return t
}

// SetSV__CustomKeys updates property "CustomKeys".
func (t AWS_WAFv2_WebACL_RateBasedStatement) SetSV__CustomKeys(v ...AWS_WAFv2_WebACL_RateBasedStatementCustomKey) AWS_WAFv2_WebACL_RateBasedStatement {
	t.CustomKeys = cfz.SV(v...)
	return t
}

// Set__EvaluationWindowSec updates property "EvaluationWindowSec".
func (t AWS_WAFv2_WebACL_RateBasedStatement) Set__EvaluationWindowSec(v cfz.Expression[int32]) AWS_WAFv2_WebACL_RateBasedStatement {
	t.EvaluationWindowSec = v
	return t
}

// SetV__EvaluationWindowSec updates property "EvaluationWindowSec".
func (t AWS_WAFv2_WebACL_RateBasedStatement) SetV__EvaluationWindowSec(v int32) AWS_WAFv2_WebACL_RateBasedStatement {
	t.EvaluationWindowSec = cfz.V(v)
	return t
}

// Set__ForwardedIPConfig updates property "ForwardedIPConfig".
func (t AWS_WAFv2_WebACL_RateBasedStatement) Set__ForwardedIPConfig(v cfz.Expression[AWS_WAFv2_WebACL_ForwardedIPConfiguration]) AWS_WAFv2_WebACL_RateBasedStatement {
	t.ForwardedIPConfig = v
	return t
}

// SetV__ForwardedIPConfig updates property "ForwardedIPConfig".
func (t AWS_WAFv2_WebACL_RateBasedStatement) SetV__ForwardedIPConfig(v AWS_WAFv2_WebACL_ForwardedIPConfiguration) AWS_WAFv2_WebACL_RateBasedStatement {
	t.ForwardedIPConfig = cfz.V(v)
	return t
}

// Set__Limit updates property "Limit".
func (t AWS_WAFv2_WebACL_RateBasedStatement) Set__Limit(v cfz.Expression[int32]) AWS_WAFv2_WebACL_RateBasedStatement {
	t.Limit = v
	return t
}

// SetV__Limit updates property "Limit".
func (t AWS_WAFv2_WebACL_RateBasedStatement) SetV__Limit(v int32) AWS_WAFv2_WebACL_RateBasedStatement {
	t.Limit = cfz.V(v)
	return t
}

// Set__ScopeDownStatement updates property "ScopeDownStatement".
func (t AWS_WAFv2_WebACL_RateBasedStatement) Set__ScopeDownStatement(v cfz.Expression[AWS_WAFv2_WebACL_Statement]) AWS_WAFv2_WebACL_RateBasedStatement {
	t.ScopeDownStatement = v
	return t
}

// SetV__ScopeDownStatement updates property "ScopeDownStatement".
func (t AWS_WAFv2_WebACL_RateBasedStatement) SetV__ScopeDownStatement(v AWS_WAFv2_WebACL_Statement) AWS_WAFv2_WebACL_RateBasedStatement {
	t.ScopeDownStatement = cfz.V(v)
	return t
}
