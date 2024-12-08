// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emrserverless

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EMRServerless_Application_ConfigurationObject__Type is the CloudFormation type for AWS::EMRServerless::Application.ConfigurationObject.
	AWS_EMRServerless_Application_ConfigurationObject__Type = "AWS::EMRServerless::Application.ConfigurationObject"
)

var (
	// AWS_EMRServerless_Application_ConfigurationObject__PropertiesMap reports all the CloudFormation properties for AWS::EMRServerless::Application.ConfigurationObject.
	AWS_EMRServerless_Application_ConfigurationObject__PropertiesMap = struct {
		Classification string
		Configurations string
		Properties     string
	}{
		Classification: "Classification",
		Configurations: "Configurations",
		Properties:     "Properties",
	}

	// AWS_EMRServerless_Application_ConfigurationObject__PropertiesSlice reports all the CloudFormation properties for AWS::EMRServerless::Application.ConfigurationObject.
	AWS_EMRServerless_Application_ConfigurationObject__PropertiesSlice = []string{
		AWS_EMRServerless_Application_ConfigurationObject__PropertiesMap.Classification,
		AWS_EMRServerless_Application_ConfigurationObject__PropertiesMap.Configurations,
		AWS_EMRServerless_Application_ConfigurationObject__PropertiesMap.Properties,
	}
)

// AWS_EMRServerless_Application_ConfigurationObject is a binding for AWS::EMRServerless::Application.ConfigurationObject.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html
type AWS_EMRServerless_Application_ConfigurationObject struct {
	// Classification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html#cfn-emrserverless-application-configurationobject-classification
	Classification cfz.Expression[string] `json:"Classification,omitempty"`

	// Configurations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html#cfn-emrserverless-application-configurationobject-configurations
	Configurations cfz.ExpressionSlice[AWS_EMRServerless_Application_ConfigurationObject] `json:"Configurations,omitempty"`

	// Properties is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-configurationobject.html#cfn-emrserverless-application-configurationobject-properties
	Properties cfz.ExpressionMap[string] `json:"Properties,omitempty"`
}

// New__AWS_EMRServerless_Application_ConfigurationObject initializes a new AWS_EMRServerless_Application_ConfigurationObject.
func New__AWS_EMRServerless_Application_ConfigurationObject() AWS_EMRServerless_Application_ConfigurationObject {
	return AWS_EMRServerless_Application_ConfigurationObject{}
}

// GetType returns the CloudFormation type.
func (AWS_EMRServerless_Application_ConfigurationObject) GetType() string {
	return AWS_EMRServerless_Application_ConfigurationObject__Type
}

// Set__Classification updates property "Classification".
func (t AWS_EMRServerless_Application_ConfigurationObject) Set__Classification(v cfz.Expression[string]) AWS_EMRServerless_Application_ConfigurationObject {
	t.Classification = v
	return t
}

// SetV__Classification updates property "Classification".
func (t AWS_EMRServerless_Application_ConfigurationObject) SetV__Classification(v string) AWS_EMRServerless_Application_ConfigurationObject {
	t.Classification = cfz.V(v)
	return t
}

// Set__Configurations updates property "Configurations".
func (t AWS_EMRServerless_Application_ConfigurationObject) Set__Configurations(v cfz.ExpressionSlice[AWS_EMRServerless_Application_ConfigurationObject]) AWS_EMRServerless_Application_ConfigurationObject {
	t.Configurations = v
	return t
}

// SetS__Configurations updates property "Configurations".
func (t AWS_EMRServerless_Application_ConfigurationObject) SetS__Configurations(v ...cfz.Expression[AWS_EMRServerless_Application_ConfigurationObject]) AWS_EMRServerless_Application_ConfigurationObject {
	t.Configurations = cfz.S(v...)
	return t
}

// SetSV__Configurations updates property "Configurations".
func (t AWS_EMRServerless_Application_ConfigurationObject) SetSV__Configurations(v ...AWS_EMRServerless_Application_ConfigurationObject) AWS_EMRServerless_Application_ConfigurationObject {
	t.Configurations = cfz.SV(v...)
	return t
}

// Set__Properties updates property "Properties".
func (t AWS_EMRServerless_Application_ConfigurationObject) Set__Properties(v cfz.ExpressionMap[string]) AWS_EMRServerless_Application_ConfigurationObject {
	t.Properties = v
	return t
}

// SetM__Properties updates property "Properties".
func (t AWS_EMRServerless_Application_ConfigurationObject) SetM__Properties(v ...map[string]cfz.Expression[string]) AWS_EMRServerless_Application_ConfigurationObject {
	t.Properties = cfz.M(v...)
	return t
}

// SetMV__Properties updates property "Properties".
func (t AWS_EMRServerless_Application_ConfigurationObject) SetMV__Properties(v ...map[string]string) AWS_EMRServerless_Application_ConfigurationObject {
	t.Properties = cfz.MV(v...)
	return t
}
