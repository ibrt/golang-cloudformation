// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_glue

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Glue_Crawler_JdbcTarget__Type is the CloudFormation type for AWS::Glue::Crawler.JdbcTarget.
	AWS_Glue_Crawler_JdbcTarget__Type = "AWS::Glue::Crawler.JdbcTarget"
)

var (
	// AWS_Glue_Crawler_JdbcTarget__PropertiesMap reports all the CloudFormation properties for AWS::Glue::Crawler.JdbcTarget.
	AWS_Glue_Crawler_JdbcTarget__PropertiesMap = struct {
		ConnectionName           string
		EnableAdditionalMetadata string
		Exclusions               string
		Path                     string
	}{
		ConnectionName:           "ConnectionName",
		EnableAdditionalMetadata: "EnableAdditionalMetadata",
		Exclusions:               "Exclusions",
		Path:                     "Path",
	}

	// AWS_Glue_Crawler_JdbcTarget__PropertiesSlice reports all the CloudFormation properties for AWS::Glue::Crawler.JdbcTarget.
	AWS_Glue_Crawler_JdbcTarget__PropertiesSlice = []string{
		AWS_Glue_Crawler_JdbcTarget__PropertiesMap.ConnectionName,
		AWS_Glue_Crawler_JdbcTarget__PropertiesMap.EnableAdditionalMetadata,
		AWS_Glue_Crawler_JdbcTarget__PropertiesMap.Exclusions,
		AWS_Glue_Crawler_JdbcTarget__PropertiesMap.Path,
	}
)

// AWS_Glue_Crawler_JdbcTarget is a binding for AWS::Glue::Crawler.JdbcTarget.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html
type AWS_Glue_Crawler_JdbcTarget struct {
	// ConnectionName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-connectionname
	ConnectionName cfz.Expression[string] `json:"ConnectionName,omitempty"`

	// EnableAdditionalMetadata is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-enableadditionalmetadata
	EnableAdditionalMetadata cfz.ExpressionSlice[string] `json:"EnableAdditionalMetadata,omitempty"`

	// Exclusions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-exclusions
	Exclusions cfz.ExpressionSlice[string] `json:"Exclusions,omitempty"`

	// Path is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-path
	Path cfz.Expression[string] `json:"Path,omitempty"`
}

// New__AWS_Glue_Crawler_JdbcTarget initializes a new AWS_Glue_Crawler_JdbcTarget.
func New__AWS_Glue_Crawler_JdbcTarget() AWS_Glue_Crawler_JdbcTarget {
	return AWS_Glue_Crawler_JdbcTarget{}
}

// GetType returns the CloudFormation type.
func (AWS_Glue_Crawler_JdbcTarget) GetType() string {
	return AWS_Glue_Crawler_JdbcTarget__Type
}

// Set__ConnectionName updates property "ConnectionName".
func (t AWS_Glue_Crawler_JdbcTarget) Set__ConnectionName(v cfz.Expression[string]) AWS_Glue_Crawler_JdbcTarget {
	t.ConnectionName = v
	return t
}

// SetV__ConnectionName updates property "ConnectionName".
func (t AWS_Glue_Crawler_JdbcTarget) SetV__ConnectionName(v string) AWS_Glue_Crawler_JdbcTarget {
	t.ConnectionName = cfz.V(v)
	return t
}

// Set__EnableAdditionalMetadata updates property "EnableAdditionalMetadata".
func (t AWS_Glue_Crawler_JdbcTarget) Set__EnableAdditionalMetadata(v cfz.ExpressionSlice[string]) AWS_Glue_Crawler_JdbcTarget {
	t.EnableAdditionalMetadata = v
	return t
}

// SetS__EnableAdditionalMetadata updates property "EnableAdditionalMetadata".
func (t AWS_Glue_Crawler_JdbcTarget) SetS__EnableAdditionalMetadata(v ...cfz.Expression[string]) AWS_Glue_Crawler_JdbcTarget {
	t.EnableAdditionalMetadata = cfz.S(v...)
	return t
}

// SetSV__EnableAdditionalMetadata updates property "EnableAdditionalMetadata".
func (t AWS_Glue_Crawler_JdbcTarget) SetSV__EnableAdditionalMetadata(v ...string) AWS_Glue_Crawler_JdbcTarget {
	t.EnableAdditionalMetadata = cfz.SV(v...)
	return t
}

// Set__Exclusions updates property "Exclusions".
func (t AWS_Glue_Crawler_JdbcTarget) Set__Exclusions(v cfz.ExpressionSlice[string]) AWS_Glue_Crawler_JdbcTarget {
	t.Exclusions = v
	return t
}

// SetS__Exclusions updates property "Exclusions".
func (t AWS_Glue_Crawler_JdbcTarget) SetS__Exclusions(v ...cfz.Expression[string]) AWS_Glue_Crawler_JdbcTarget {
	t.Exclusions = cfz.S(v...)
	return t
}

// SetSV__Exclusions updates property "Exclusions".
func (t AWS_Glue_Crawler_JdbcTarget) SetSV__Exclusions(v ...string) AWS_Glue_Crawler_JdbcTarget {
	t.Exclusions = cfz.SV(v...)
	return t
}

// Set__Path updates property "Path".
func (t AWS_Glue_Crawler_JdbcTarget) Set__Path(v cfz.Expression[string]) AWS_Glue_Crawler_JdbcTarget {
	t.Path = v
	return t
}

// SetV__Path updates property "Path".
func (t AWS_Glue_Crawler_JdbcTarget) SetV__Path(v string) AWS_Glue_Crawler_JdbcTarget {
	t.Path = cfz.V(v)
	return t
}
