// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_StorageLensGroup_And__Type is the CloudFormation type for AWS::S3::StorageLensGroup.And.
	AWS_S3_StorageLensGroup_And__Type = "AWS::S3::StorageLensGroup.And"
)

var (
	// AWS_S3_StorageLensGroup_And__PropertiesMap reports all the CloudFormation properties for AWS::S3::StorageLensGroup.And.
	AWS_S3_StorageLensGroup_And__PropertiesMap = struct {
		MatchAnyPrefix  string
		MatchAnySuffix  string
		MatchAnyTag     string
		MatchObjectAge  string
		MatchObjectSize string
	}{
		MatchAnyPrefix:  "MatchAnyPrefix",
		MatchAnySuffix:  "MatchAnySuffix",
		MatchAnyTag:     "MatchAnyTag",
		MatchObjectAge:  "MatchObjectAge",
		MatchObjectSize: "MatchObjectSize",
	}

	// AWS_S3_StorageLensGroup_And__PropertiesSlice reports all the CloudFormation properties for AWS::S3::StorageLensGroup.And.
	AWS_S3_StorageLensGroup_And__PropertiesSlice = []string{
		AWS_S3_StorageLensGroup_And__PropertiesMap.MatchAnyPrefix,
		AWS_S3_StorageLensGroup_And__PropertiesMap.MatchAnySuffix,
		AWS_S3_StorageLensGroup_And__PropertiesMap.MatchAnyTag,
		AWS_S3_StorageLensGroup_And__PropertiesMap.MatchObjectAge,
		AWS_S3_StorageLensGroup_And__PropertiesMap.MatchObjectSize,
	}
)

// AWS_S3_StorageLensGroup_And is a binding for AWS::S3::StorageLensGroup.And.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html
type AWS_S3_StorageLensGroup_And struct {
	// MatchAnyPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchanyprefix
	MatchAnyPrefix cfz.ExpressionSlice[string] `json:"MatchAnyPrefix,omitempty"`

	// MatchAnySuffix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchanysuffix
	MatchAnySuffix cfz.ExpressionSlice[string] `json:"MatchAnySuffix,omitempty"`

	// MatchAnyTag is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchanytag
	MatchAnyTag cfz.ExpressionSlice[cfz.Tag] `json:"MatchAnyTag,omitempty"`

	// MatchObjectAge is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchobjectage
	MatchObjectAge cfz.Expression[AWS_S3_StorageLensGroup_MatchObjectAge] `json:"MatchObjectAge,omitempty"`

	// MatchObjectSize is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchobjectsize
	MatchObjectSize cfz.Expression[AWS_S3_StorageLensGroup_MatchObjectSize] `json:"MatchObjectSize,omitempty"`
}

// New__AWS_S3_StorageLensGroup_And initializes a new AWS_S3_StorageLensGroup_And.
func New__AWS_S3_StorageLensGroup_And() AWS_S3_StorageLensGroup_And {
	return AWS_S3_StorageLensGroup_And{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_StorageLensGroup_And) GetType() string {
	return AWS_S3_StorageLensGroup_And__Type
}

// Set__MatchAnyPrefix updates property "MatchAnyPrefix".
func (t AWS_S3_StorageLensGroup_And) Set__MatchAnyPrefix(v cfz.ExpressionSlice[string]) AWS_S3_StorageLensGroup_And {
	t.MatchAnyPrefix = v
	return t
}

// SetS__MatchAnyPrefix updates property "MatchAnyPrefix".
func (t AWS_S3_StorageLensGroup_And) SetS__MatchAnyPrefix(v ...cfz.Expression[string]) AWS_S3_StorageLensGroup_And {
	t.MatchAnyPrefix = cfz.S(v...)
	return t
}

// SetSV__MatchAnyPrefix updates property "MatchAnyPrefix".
func (t AWS_S3_StorageLensGroup_And) SetSV__MatchAnyPrefix(v ...string) AWS_S3_StorageLensGroup_And {
	t.MatchAnyPrefix = cfz.SV(v...)
	return t
}

// Set__MatchAnySuffix updates property "MatchAnySuffix".
func (t AWS_S3_StorageLensGroup_And) Set__MatchAnySuffix(v cfz.ExpressionSlice[string]) AWS_S3_StorageLensGroup_And {
	t.MatchAnySuffix = v
	return t
}

// SetS__MatchAnySuffix updates property "MatchAnySuffix".
func (t AWS_S3_StorageLensGroup_And) SetS__MatchAnySuffix(v ...cfz.Expression[string]) AWS_S3_StorageLensGroup_And {
	t.MatchAnySuffix = cfz.S(v...)
	return t
}

// SetSV__MatchAnySuffix updates property "MatchAnySuffix".
func (t AWS_S3_StorageLensGroup_And) SetSV__MatchAnySuffix(v ...string) AWS_S3_StorageLensGroup_And {
	t.MatchAnySuffix = cfz.SV(v...)
	return t
}

// Set__MatchAnyTag updates property "MatchAnyTag".
func (t AWS_S3_StorageLensGroup_And) Set__MatchAnyTag(v cfz.ExpressionSlice[cfz.Tag]) AWS_S3_StorageLensGroup_And {
	t.MatchAnyTag = v
	return t
}

// SetS__MatchAnyTag updates property "MatchAnyTag".
func (t AWS_S3_StorageLensGroup_And) SetS__MatchAnyTag(v ...cfz.Expression[cfz.Tag]) AWS_S3_StorageLensGroup_And {
	t.MatchAnyTag = cfz.S(v...)
	return t
}

// SetSV__MatchAnyTag updates property "MatchAnyTag".
func (t AWS_S3_StorageLensGroup_And) SetSV__MatchAnyTag(v ...cfz.Tag) AWS_S3_StorageLensGroup_And {
	t.MatchAnyTag = cfz.SV(v...)
	return t
}

// Set__MatchObjectAge updates property "MatchObjectAge".
func (t AWS_S3_StorageLensGroup_And) Set__MatchObjectAge(v cfz.Expression[AWS_S3_StorageLensGroup_MatchObjectAge]) AWS_S3_StorageLensGroup_And {
	t.MatchObjectAge = v
	return t
}

// SetV__MatchObjectAge updates property "MatchObjectAge".
func (t AWS_S3_StorageLensGroup_And) SetV__MatchObjectAge(v AWS_S3_StorageLensGroup_MatchObjectAge) AWS_S3_StorageLensGroup_And {
	t.MatchObjectAge = cfz.V(v)
	return t
}

// Set__MatchObjectSize updates property "MatchObjectSize".
func (t AWS_S3_StorageLensGroup_And) Set__MatchObjectSize(v cfz.Expression[AWS_S3_StorageLensGroup_MatchObjectSize]) AWS_S3_StorageLensGroup_And {
	t.MatchObjectSize = v
	return t
}

// SetV__MatchObjectSize updates property "MatchObjectSize".
func (t AWS_S3_StorageLensGroup_And) SetV__MatchObjectSize(v AWS_S3_StorageLensGroup_MatchObjectSize) AWS_S3_StorageLensGroup_And {
	t.MatchObjectSize = cfz.V(v)
	return t
}
