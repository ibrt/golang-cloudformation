// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_opsworks

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_OpsWorks_App_Source__Type is the CloudFormation type for AWS::OpsWorks::App.Source.
	AWS_OpsWorks_App_Source__Type = "AWS::OpsWorks::App.Source"
)

var (
	// AWS_OpsWorks_App_Source__PropertiesMap reports all the CloudFormation properties for AWS::OpsWorks::App.Source.
	AWS_OpsWorks_App_Source__PropertiesMap = struct {
		Password string
		Revision string
		SshKey   string
		Type     string
		Url      string
		Username string
	}{
		Password: "Password",
		Revision: "Revision",
		SshKey:   "SshKey",
		Type:     "Type",
		Url:      "Url",
		Username: "Username",
	}

	// AWS_OpsWorks_App_Source__PropertiesSlice reports all the CloudFormation properties for AWS::OpsWorks::App.Source.
	AWS_OpsWorks_App_Source__PropertiesSlice = []string{
		AWS_OpsWorks_App_Source__PropertiesMap.Password,
		AWS_OpsWorks_App_Source__PropertiesMap.Revision,
		AWS_OpsWorks_App_Source__PropertiesMap.SshKey,
		AWS_OpsWorks_App_Source__PropertiesMap.Type,
		AWS_OpsWorks_App_Source__PropertiesMap.Url,
		AWS_OpsWorks_App_Source__PropertiesMap.Username,
	}
)

// AWS_OpsWorks_App_Source is a binding for AWS::OpsWorks::App.Source.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html
type AWS_OpsWorks_App_Source struct {
	// Password is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-pw
	Password cfz.Expression[string] `json:"Password,omitempty"`

	// Revision is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-revision
	Revision cfz.Expression[string] `json:"Revision,omitempty"`

	// SshKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-sshkey
	SshKey cfz.Expression[string] `json:"SshKey,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-type
	Type cfz.Expression[string] `json:"Type,omitempty"`

	// Url is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-url
	Url cfz.Expression[string] `json:"Url,omitempty"`

	// Username is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-username
	Username cfz.Expression[string] `json:"Username,omitempty"`
}

// New__AWS_OpsWorks_App_Source initializes a new AWS_OpsWorks_App_Source.
func New__AWS_OpsWorks_App_Source() AWS_OpsWorks_App_Source {
	return AWS_OpsWorks_App_Source{}
}

// GetType returns the CloudFormation type.
func (AWS_OpsWorks_App_Source) GetType() string {
	return AWS_OpsWorks_App_Source__Type
}

// Set__Password updates property "Password".
func (t AWS_OpsWorks_App_Source) Set__Password(v cfz.Expression[string]) AWS_OpsWorks_App_Source {
	t.Password = v
	return t
}

// SetV__Password updates property "Password".
func (t AWS_OpsWorks_App_Source) SetV__Password(v string) AWS_OpsWorks_App_Source {
	t.Password = cfz.V(v)
	return t
}

// Set__Revision updates property "Revision".
func (t AWS_OpsWorks_App_Source) Set__Revision(v cfz.Expression[string]) AWS_OpsWorks_App_Source {
	t.Revision = v
	return t
}

// SetV__Revision updates property "Revision".
func (t AWS_OpsWorks_App_Source) SetV__Revision(v string) AWS_OpsWorks_App_Source {
	t.Revision = cfz.V(v)
	return t
}

// Set__SshKey updates property "SshKey".
func (t AWS_OpsWorks_App_Source) Set__SshKey(v cfz.Expression[string]) AWS_OpsWorks_App_Source {
	t.SshKey = v
	return t
}

// SetV__SshKey updates property "SshKey".
func (t AWS_OpsWorks_App_Source) SetV__SshKey(v string) AWS_OpsWorks_App_Source {
	t.SshKey = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_OpsWorks_App_Source) Set__Type(v cfz.Expression[string]) AWS_OpsWorks_App_Source {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_OpsWorks_App_Source) SetV__Type(v string) AWS_OpsWorks_App_Source {
	t.Type = cfz.V(v)
	return t
}

// Set__Url updates property "Url".
func (t AWS_OpsWorks_App_Source) Set__Url(v cfz.Expression[string]) AWS_OpsWorks_App_Source {
	t.Url = v
	return t
}

// SetV__Url updates property "Url".
func (t AWS_OpsWorks_App_Source) SetV__Url(v string) AWS_OpsWorks_App_Source {
	t.Url = cfz.V(v)
	return t
}

// Set__Username updates property "Username".
func (t AWS_OpsWorks_App_Source) Set__Username(v cfz.Expression[string]) AWS_OpsWorks_App_Source {
	t.Username = v
	return t
}

// SetV__Username updates property "Username".
func (t AWS_OpsWorks_App_Source) SetV__Username(v string) AWS_OpsWorks_App_Source {
	t.Username = cfz.V(v)
	return t
}
