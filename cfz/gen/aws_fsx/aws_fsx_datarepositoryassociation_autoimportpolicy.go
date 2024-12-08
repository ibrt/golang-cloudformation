// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_fsx

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FSx_DataRepositoryAssociation_AutoImportPolicy__Type is the CloudFormation type for AWS::FSx::DataRepositoryAssociation.AutoImportPolicy.
	AWS_FSx_DataRepositoryAssociation_AutoImportPolicy__Type = "AWS::FSx::DataRepositoryAssociation.AutoImportPolicy"
)

var (
	// AWS_FSx_DataRepositoryAssociation_AutoImportPolicy__PropertiesMap reports all the CloudFormation properties for AWS::FSx::DataRepositoryAssociation.AutoImportPolicy.
	AWS_FSx_DataRepositoryAssociation_AutoImportPolicy__PropertiesMap = struct {
		Events string
	}{
		Events: "Events",
	}

	// AWS_FSx_DataRepositoryAssociation_AutoImportPolicy__PropertiesSlice reports all the CloudFormation properties for AWS::FSx::DataRepositoryAssociation.AutoImportPolicy.
	AWS_FSx_DataRepositoryAssociation_AutoImportPolicy__PropertiesSlice = []string{
		AWS_FSx_DataRepositoryAssociation_AutoImportPolicy__PropertiesMap.Events,
	}
)

// AWS_FSx_DataRepositoryAssociation_AutoImportPolicy is a binding for AWS::FSx::DataRepositoryAssociation.AutoImportPolicy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-datarepositoryassociation-autoimportpolicy.html
type AWS_FSx_DataRepositoryAssociation_AutoImportPolicy struct {
	// Events is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-datarepositoryassociation-autoimportpolicy.html#cfn-fsx-datarepositoryassociation-autoimportpolicy-events
	Events cfz.ExpressionSlice[string] `json:"Events,omitempty"`
}

// New__AWS_FSx_DataRepositoryAssociation_AutoImportPolicy initializes a new AWS_FSx_DataRepositoryAssociation_AutoImportPolicy.
func New__AWS_FSx_DataRepositoryAssociation_AutoImportPolicy() AWS_FSx_DataRepositoryAssociation_AutoImportPolicy {
	return AWS_FSx_DataRepositoryAssociation_AutoImportPolicy{}
}

// GetType returns the CloudFormation type.
func (AWS_FSx_DataRepositoryAssociation_AutoImportPolicy) GetType() string {
	return AWS_FSx_DataRepositoryAssociation_AutoImportPolicy__Type
}

// Set__Events updates property "Events".
func (t AWS_FSx_DataRepositoryAssociation_AutoImportPolicy) Set__Events(v cfz.ExpressionSlice[string]) AWS_FSx_DataRepositoryAssociation_AutoImportPolicy {
	t.Events = v
	return t
}

// SetS__Events updates property "Events".
func (t AWS_FSx_DataRepositoryAssociation_AutoImportPolicy) SetS__Events(v ...cfz.Expression[string]) AWS_FSx_DataRepositoryAssociation_AutoImportPolicy {
	t.Events = cfz.S(v...)
	return t
}

// SetSV__Events updates property "Events".
func (t AWS_FSx_DataRepositoryAssociation_AutoImportPolicy) SetSV__Events(v ...string) AWS_FSx_DataRepositoryAssociation_AutoImportPolicy {
	t.Events = cfz.SV(v...)
	return t
}
