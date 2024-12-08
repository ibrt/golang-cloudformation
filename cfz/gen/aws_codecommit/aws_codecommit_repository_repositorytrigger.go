// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codecommit

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CodeCommit_Repository_RepositoryTrigger__Type is the CloudFormation type for AWS::CodeCommit::Repository.RepositoryTrigger.
	AWS_CodeCommit_Repository_RepositoryTrigger__Type = "AWS::CodeCommit::Repository.RepositoryTrigger"
)

var (
	// AWS_CodeCommit_Repository_RepositoryTrigger__PropertiesMap reports all the CloudFormation properties for AWS::CodeCommit::Repository.RepositoryTrigger.
	AWS_CodeCommit_Repository_RepositoryTrigger__PropertiesMap = struct {
		Branches       string
		CustomData     string
		DestinationArn string
		Events         string
		Name           string
	}{
		Branches:       "Branches",
		CustomData:     "CustomData",
		DestinationArn: "DestinationArn",
		Events:         "Events",
		Name:           "Name",
	}

	// AWS_CodeCommit_Repository_RepositoryTrigger__PropertiesSlice reports all the CloudFormation properties for AWS::CodeCommit::Repository.RepositoryTrigger.
	AWS_CodeCommit_Repository_RepositoryTrigger__PropertiesSlice = []string{
		AWS_CodeCommit_Repository_RepositoryTrigger__PropertiesMap.Branches,
		AWS_CodeCommit_Repository_RepositoryTrigger__PropertiesMap.CustomData,
		AWS_CodeCommit_Repository_RepositoryTrigger__PropertiesMap.DestinationArn,
		AWS_CodeCommit_Repository_RepositoryTrigger__PropertiesMap.Events,
		AWS_CodeCommit_Repository_RepositoryTrigger__PropertiesMap.Name,
	}
)

// AWS_CodeCommit_Repository_RepositoryTrigger is a binding for AWS::CodeCommit::Repository.RepositoryTrigger.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html
type AWS_CodeCommit_Repository_RepositoryTrigger struct {
	// Branches is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-branches
	Branches cfz.ExpressionSlice[string] `json:"Branches,omitempty"`

	// CustomData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-customdata
	CustomData cfz.Expression[string] `json:"CustomData,omitempty"`

	// DestinationArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-destinationarn
	DestinationArn cfz.Expression[string] `json:"DestinationArn,omitempty"`

	// Events is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-events
	Events cfz.ExpressionSlice[string] `json:"Events,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_CodeCommit_Repository_RepositoryTrigger initializes a new AWS_CodeCommit_Repository_RepositoryTrigger.
func New__AWS_CodeCommit_Repository_RepositoryTrigger() AWS_CodeCommit_Repository_RepositoryTrigger {
	return AWS_CodeCommit_Repository_RepositoryTrigger{}
}

// GetType returns the CloudFormation type.
func (AWS_CodeCommit_Repository_RepositoryTrigger) GetType() string {
	return AWS_CodeCommit_Repository_RepositoryTrigger__Type
}

// Set__Branches updates property "Branches".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) Set__Branches(v cfz.ExpressionSlice[string]) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.Branches = v
	return t
}

// SetS__Branches updates property "Branches".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) SetS__Branches(v ...cfz.Expression[string]) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.Branches = cfz.S(v...)
	return t
}

// SetSV__Branches updates property "Branches".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) SetSV__Branches(v ...string) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.Branches = cfz.SV(v...)
	return t
}

// Set__CustomData updates property "CustomData".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) Set__CustomData(v cfz.Expression[string]) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.CustomData = v
	return t
}

// SetV__CustomData updates property "CustomData".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) SetV__CustomData(v string) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.CustomData = cfz.V(v)
	return t
}

// Set__DestinationArn updates property "DestinationArn".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) Set__DestinationArn(v cfz.Expression[string]) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.DestinationArn = v
	return t
}

// SetV__DestinationArn updates property "DestinationArn".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) SetV__DestinationArn(v string) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.DestinationArn = cfz.V(v)
	return t
}

// Set__Events updates property "Events".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) Set__Events(v cfz.ExpressionSlice[string]) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.Events = v
	return t
}

// SetS__Events updates property "Events".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) SetS__Events(v ...cfz.Expression[string]) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.Events = cfz.S(v...)
	return t
}

// SetSV__Events updates property "Events".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) SetSV__Events(v ...string) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.Events = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) Set__Name(v cfz.Expression[string]) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_CodeCommit_Repository_RepositoryTrigger) SetV__Name(v string) AWS_CodeCommit_Repository_RepositoryTrigger {
	t.Name = cfz.V(v)
	return t
}
