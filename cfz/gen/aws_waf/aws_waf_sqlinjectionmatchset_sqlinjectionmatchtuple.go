// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_waf

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple__Type is the CloudFormation type for AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple.
	AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple__Type = "AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple"
)

var (
	// AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple__PropertiesMap reports all the CloudFormation properties for AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple.
	AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple__PropertiesMap = struct {
		FieldToMatch       string
		TextTransformation string
	}{
		FieldToMatch:       "FieldToMatch",
		TextTransformation: "TextTransformation",
	}

	// AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple__PropertiesSlice reports all the CloudFormation properties for AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple.
	AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple__PropertiesSlice = []string{
		AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple__PropertiesMap.FieldToMatch,
		AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple__PropertiesMap.TextTransformation,
	}
)

// AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple is a binding for AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html
type AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple struct {
	// FieldToMatch is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples-fieldtomatch
	FieldToMatch cfz.Expression[AWS_WAF_SqlInjectionMatchSet_FieldToMatch] `json:"FieldToMatch,omitempty"`

	// TextTransformation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples-texttransformation
	TextTransformation cfz.Expression[string] `json:"TextTransformation,omitempty"`
}

// New__AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple initializes a new AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple.
func New__AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple() AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple {
	return AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple{}
}

// GetType returns the CloudFormation type.
func (AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple) GetType() string {
	return AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple__Type
}

// Set__FieldToMatch updates property "FieldToMatch".
func (t AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple) Set__FieldToMatch(v cfz.Expression[AWS_WAF_SqlInjectionMatchSet_FieldToMatch]) AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple {
	t.FieldToMatch = v
	return t
}

// SetV__FieldToMatch updates property "FieldToMatch".
func (t AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple) SetV__FieldToMatch(v AWS_WAF_SqlInjectionMatchSet_FieldToMatch) AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple {
	t.FieldToMatch = cfz.V(v)
	return t
}

// Set__TextTransformation updates property "TextTransformation".
func (t AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple) Set__TextTransformation(v cfz.Expression[string]) AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple {
	t.TextTransformation = v
	return t
}

// SetV__TextTransformation updates property "TextTransformation".
func (t AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple) SetV__TextTransformation(v string) AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple {
	t.TextTransformation = cfz.V(v)
	return t
}
