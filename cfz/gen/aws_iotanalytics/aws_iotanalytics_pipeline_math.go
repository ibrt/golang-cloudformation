// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotanalytics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTAnalytics_Pipeline_Math__Type is the CloudFormation type for AWS::IoTAnalytics::Pipeline.Math.
	AWS_IoTAnalytics_Pipeline_Math__Type = "AWS::IoTAnalytics::Pipeline.Math"
)

var (
	// AWS_IoTAnalytics_Pipeline_Math__PropertiesMap reports all the CloudFormation properties for AWS::IoTAnalytics::Pipeline.Math.
	AWS_IoTAnalytics_Pipeline_Math__PropertiesMap = struct {
		Attribute string
		Math      string
		Name      string
		Next      string
	}{
		Attribute: "Attribute",
		Math:      "Math",
		Name:      "Name",
		Next:      "Next",
	}

	// AWS_IoTAnalytics_Pipeline_Math__PropertiesSlice reports all the CloudFormation properties for AWS::IoTAnalytics::Pipeline.Math.
	AWS_IoTAnalytics_Pipeline_Math__PropertiesSlice = []string{
		AWS_IoTAnalytics_Pipeline_Math__PropertiesMap.Attribute,
		AWS_IoTAnalytics_Pipeline_Math__PropertiesMap.Math,
		AWS_IoTAnalytics_Pipeline_Math__PropertiesMap.Name,
		AWS_IoTAnalytics_Pipeline_Math__PropertiesMap.Next,
	}
)

// AWS_IoTAnalytics_Pipeline_Math is a binding for AWS::IoTAnalytics::Pipeline.Math.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html
type AWS_IoTAnalytics_Pipeline_Math struct {
	// Attribute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html#cfn-iotanalytics-pipeline-math-attribute
	Attribute cfz.Expression[string] `json:"Attribute,omitempty"`

	// Math is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html#cfn-iotanalytics-pipeline-math-math
	Math cfz.Expression[string] `json:"Math,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html#cfn-iotanalytics-pipeline-math-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Next is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html#cfn-iotanalytics-pipeline-math-next
	Next cfz.Expression[string] `json:"Next,omitempty"`
}

// New__AWS_IoTAnalytics_Pipeline_Math initializes a new AWS_IoTAnalytics_Pipeline_Math.
func New__AWS_IoTAnalytics_Pipeline_Math() AWS_IoTAnalytics_Pipeline_Math {
	return AWS_IoTAnalytics_Pipeline_Math{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTAnalytics_Pipeline_Math) GetType() string {
	return AWS_IoTAnalytics_Pipeline_Math__Type
}

// Set__Attribute updates property "Attribute".
func (t AWS_IoTAnalytics_Pipeline_Math) Set__Attribute(v cfz.Expression[string]) AWS_IoTAnalytics_Pipeline_Math {
	t.Attribute = v
	return t
}

// SetV__Attribute updates property "Attribute".
func (t AWS_IoTAnalytics_Pipeline_Math) SetV__Attribute(v string) AWS_IoTAnalytics_Pipeline_Math {
	t.Attribute = cfz.V(v)
	return t
}

// Set__Math updates property "Math".
func (t AWS_IoTAnalytics_Pipeline_Math) Set__Math(v cfz.Expression[string]) AWS_IoTAnalytics_Pipeline_Math {
	t.Math = v
	return t
}

// SetV__Math updates property "Math".
func (t AWS_IoTAnalytics_Pipeline_Math) SetV__Math(v string) AWS_IoTAnalytics_Pipeline_Math {
	t.Math = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_IoTAnalytics_Pipeline_Math) Set__Name(v cfz.Expression[string]) AWS_IoTAnalytics_Pipeline_Math {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_IoTAnalytics_Pipeline_Math) SetV__Name(v string) AWS_IoTAnalytics_Pipeline_Math {
	t.Name = cfz.V(v)
	return t
}

// Set__Next updates property "Next".
func (t AWS_IoTAnalytics_Pipeline_Math) Set__Next(v cfz.Expression[string]) AWS_IoTAnalytics_Pipeline_Math {
	t.Next = v
	return t
}

// SetV__Next updates property "Next".
func (t AWS_IoTAnalytics_Pipeline_Math) SetV__Next(v string) AWS_IoTAnalytics_Pipeline_Math {
	t.Next = cfz.V(v)
	return t
}
