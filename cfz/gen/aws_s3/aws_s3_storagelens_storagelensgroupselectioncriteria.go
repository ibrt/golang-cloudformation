// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_StorageLens_StorageLensGroupSelectionCriteria__Type is the CloudFormation type for AWS::S3::StorageLens.StorageLensGroupSelectionCriteria.
	AWS_S3_StorageLens_StorageLensGroupSelectionCriteria__Type = "AWS::S3::StorageLens.StorageLensGroupSelectionCriteria"
)

var (
	// AWS_S3_StorageLens_StorageLensGroupSelectionCriteria__PropertiesMap reports all the CloudFormation properties for AWS::S3::StorageLens.StorageLensGroupSelectionCriteria.
	AWS_S3_StorageLens_StorageLensGroupSelectionCriteria__PropertiesMap = struct {
		Exclude string
		Include string
	}{
		Exclude: "Exclude",
		Include: "Include",
	}

	// AWS_S3_StorageLens_StorageLensGroupSelectionCriteria__PropertiesSlice reports all the CloudFormation properties for AWS::S3::StorageLens.StorageLensGroupSelectionCriteria.
	AWS_S3_StorageLens_StorageLensGroupSelectionCriteria__PropertiesSlice = []string{
		AWS_S3_StorageLens_StorageLensGroupSelectionCriteria__PropertiesMap.Exclude,
		AWS_S3_StorageLens_StorageLensGroupSelectionCriteria__PropertiesMap.Include,
	}
)

// AWS_S3_StorageLens_StorageLensGroupSelectionCriteria is a binding for AWS::S3::StorageLens.StorageLensGroupSelectionCriteria.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgroupselectioncriteria.html
type AWS_S3_StorageLens_StorageLensGroupSelectionCriteria struct {
	// Exclude is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgroupselectioncriteria.html#cfn-s3-storagelens-storagelensgroupselectioncriteria-exclude
	Exclude cfz.ExpressionSlice[string] `json:"Exclude,omitempty"`

	// Include is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgroupselectioncriteria.html#cfn-s3-storagelens-storagelensgroupselectioncriteria-include
	Include cfz.ExpressionSlice[string] `json:"Include,omitempty"`
}

// New__AWS_S3_StorageLens_StorageLensGroupSelectionCriteria initializes a new AWS_S3_StorageLens_StorageLensGroupSelectionCriteria.
func New__AWS_S3_StorageLens_StorageLensGroupSelectionCriteria() AWS_S3_StorageLens_StorageLensGroupSelectionCriteria {
	return AWS_S3_StorageLens_StorageLensGroupSelectionCriteria{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_StorageLens_StorageLensGroupSelectionCriteria) GetType() string {
	return AWS_S3_StorageLens_StorageLensGroupSelectionCriteria__Type
}

// Set__Exclude updates property "Exclude".
func (t AWS_S3_StorageLens_StorageLensGroupSelectionCriteria) Set__Exclude(v cfz.ExpressionSlice[string]) AWS_S3_StorageLens_StorageLensGroupSelectionCriteria {
	t.Exclude = v
	return t
}

// SetS__Exclude updates property "Exclude".
func (t AWS_S3_StorageLens_StorageLensGroupSelectionCriteria) SetS__Exclude(v ...cfz.Expression[string]) AWS_S3_StorageLens_StorageLensGroupSelectionCriteria {
	t.Exclude = cfz.S(v...)
	return t
}

// SetSV__Exclude updates property "Exclude".
func (t AWS_S3_StorageLens_StorageLensGroupSelectionCriteria) SetSV__Exclude(v ...string) AWS_S3_StorageLens_StorageLensGroupSelectionCriteria {
	t.Exclude = cfz.SV(v...)
	return t
}

// Set__Include updates property "Include".
func (t AWS_S3_StorageLens_StorageLensGroupSelectionCriteria) Set__Include(v cfz.ExpressionSlice[string]) AWS_S3_StorageLens_StorageLensGroupSelectionCriteria {
	t.Include = v
	return t
}

// SetS__Include updates property "Include".
func (t AWS_S3_StorageLens_StorageLensGroupSelectionCriteria) SetS__Include(v ...cfz.Expression[string]) AWS_S3_StorageLens_StorageLensGroupSelectionCriteria {
	t.Include = cfz.S(v...)
	return t
}

// SetSV__Include updates property "Include".
func (t AWS_S3_StorageLens_StorageLensGroupSelectionCriteria) SetSV__Include(v ...string) AWS_S3_StorageLens_StorageLensGroupSelectionCriteria {
	t.Include = cfz.SV(v...)
	return t
}
