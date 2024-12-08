// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotanalytics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTAnalytics_Pipeline_SelectAttributes__Type is the CloudFormation type for AWS::IoTAnalytics::Pipeline.SelectAttributes.
	AWS_IoTAnalytics_Pipeline_SelectAttributes__Type = "AWS::IoTAnalytics::Pipeline.SelectAttributes"
)

var (
	// AWS_IoTAnalytics_Pipeline_SelectAttributes__PropertiesMap reports all the CloudFormation properties for AWS::IoTAnalytics::Pipeline.SelectAttributes.
	AWS_IoTAnalytics_Pipeline_SelectAttributes__PropertiesMap = struct {
		Attributes string
		Name       string
		Next       string
	}{
		Attributes: "Attributes",
		Name:       "Name",
		Next:       "Next",
	}

	// AWS_IoTAnalytics_Pipeline_SelectAttributes__PropertiesSlice reports all the CloudFormation properties for AWS::IoTAnalytics::Pipeline.SelectAttributes.
	AWS_IoTAnalytics_Pipeline_SelectAttributes__PropertiesSlice = []string{
		AWS_IoTAnalytics_Pipeline_SelectAttributes__PropertiesMap.Attributes,
		AWS_IoTAnalytics_Pipeline_SelectAttributes__PropertiesMap.Name,
		AWS_IoTAnalytics_Pipeline_SelectAttributes__PropertiesMap.Next,
	}
)

// AWS_IoTAnalytics_Pipeline_SelectAttributes is a binding for AWS::IoTAnalytics::Pipeline.SelectAttributes.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-selectattributes.html
type AWS_IoTAnalytics_Pipeline_SelectAttributes struct {
	// Attributes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-selectattributes.html#cfn-iotanalytics-pipeline-selectattributes-attributes
	Attributes cfz.ExpressionSlice[string] `json:"Attributes,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-selectattributes.html#cfn-iotanalytics-pipeline-selectattributes-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Next is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-selectattributes.html#cfn-iotanalytics-pipeline-selectattributes-next
	Next cfz.Expression[string] `json:"Next,omitempty"`
}

// New__AWS_IoTAnalytics_Pipeline_SelectAttributes initializes a new AWS_IoTAnalytics_Pipeline_SelectAttributes.
func New__AWS_IoTAnalytics_Pipeline_SelectAttributes() AWS_IoTAnalytics_Pipeline_SelectAttributes {
	return AWS_IoTAnalytics_Pipeline_SelectAttributes{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTAnalytics_Pipeline_SelectAttributes) GetType() string {
	return AWS_IoTAnalytics_Pipeline_SelectAttributes__Type
}

// Set__Attributes updates property "Attributes".
func (t AWS_IoTAnalytics_Pipeline_SelectAttributes) Set__Attributes(v cfz.ExpressionSlice[string]) AWS_IoTAnalytics_Pipeline_SelectAttributes {
	t.Attributes = v
	return t
}

// SetS__Attributes updates property "Attributes".
func (t AWS_IoTAnalytics_Pipeline_SelectAttributes) SetS__Attributes(v ...cfz.Expression[string]) AWS_IoTAnalytics_Pipeline_SelectAttributes {
	t.Attributes = cfz.S(v...)
	return t
}

// SetSV__Attributes updates property "Attributes".
func (t AWS_IoTAnalytics_Pipeline_SelectAttributes) SetSV__Attributes(v ...string) AWS_IoTAnalytics_Pipeline_SelectAttributes {
	t.Attributes = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t AWS_IoTAnalytics_Pipeline_SelectAttributes) Set__Name(v cfz.Expression[string]) AWS_IoTAnalytics_Pipeline_SelectAttributes {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_IoTAnalytics_Pipeline_SelectAttributes) SetV__Name(v string) AWS_IoTAnalytics_Pipeline_SelectAttributes {
	t.Name = cfz.V(v)
	return t
}

// Set__Next updates property "Next".
func (t AWS_IoTAnalytics_Pipeline_SelectAttributes) Set__Next(v cfz.Expression[string]) AWS_IoTAnalytics_Pipeline_SelectAttributes {
	t.Next = v
	return t
}

// SetV__Next updates property "Next".
func (t AWS_IoTAnalytics_Pipeline_SelectAttributes) SetV__Next(v string) AWS_IoTAnalytics_Pipeline_SelectAttributes {
	t.Next = cfz.V(v)
	return t
}
