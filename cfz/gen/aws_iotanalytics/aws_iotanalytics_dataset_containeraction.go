// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotanalytics

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTAnalytics_Dataset_ContainerAction__Type is the CloudFormation type for AWS::IoTAnalytics::Dataset.ContainerAction.
	AWS_IoTAnalytics_Dataset_ContainerAction__Type = "AWS::IoTAnalytics::Dataset.ContainerAction"
)

var (
	// AWS_IoTAnalytics_Dataset_ContainerAction__PropertiesMap reports all the CloudFormation properties for AWS::IoTAnalytics::Dataset.ContainerAction.
	AWS_IoTAnalytics_Dataset_ContainerAction__PropertiesMap = struct {
		ExecutionRoleArn      string
		Image                 string
		ResourceConfiguration string
		Variables             string
	}{
		ExecutionRoleArn:      "ExecutionRoleArn",
		Image:                 "Image",
		ResourceConfiguration: "ResourceConfiguration",
		Variables:             "Variables",
	}

	// AWS_IoTAnalytics_Dataset_ContainerAction__PropertiesSlice reports all the CloudFormation properties for AWS::IoTAnalytics::Dataset.ContainerAction.
	AWS_IoTAnalytics_Dataset_ContainerAction__PropertiesSlice = []string{
		AWS_IoTAnalytics_Dataset_ContainerAction__PropertiesMap.ExecutionRoleArn,
		AWS_IoTAnalytics_Dataset_ContainerAction__PropertiesMap.Image,
		AWS_IoTAnalytics_Dataset_ContainerAction__PropertiesMap.ResourceConfiguration,
		AWS_IoTAnalytics_Dataset_ContainerAction__PropertiesMap.Variables,
	}
)

// AWS_IoTAnalytics_Dataset_ContainerAction is a binding for AWS::IoTAnalytics::Dataset.ContainerAction.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html
type AWS_IoTAnalytics_Dataset_ContainerAction struct {
	// ExecutionRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html#cfn-iotanalytics-dataset-containeraction-executionrolearn
	ExecutionRoleArn cfz.Expression[string] `json:"ExecutionRoleArn,omitempty"`

	// Image is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html#cfn-iotanalytics-dataset-containeraction-image
	Image cfz.Expression[string] `json:"Image,omitempty"`

	// ResourceConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html#cfn-iotanalytics-dataset-containeraction-resourceconfiguration
	ResourceConfiguration cfz.Expression[AWS_IoTAnalytics_Dataset_ResourceConfiguration] `json:"ResourceConfiguration,omitempty"`

	// Variables is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html#cfn-iotanalytics-dataset-containeraction-variables
	Variables cfz.ExpressionSlice[AWS_IoTAnalytics_Dataset_Variable] `json:"Variables,omitempty"`
}

// New__AWS_IoTAnalytics_Dataset_ContainerAction initializes a new AWS_IoTAnalytics_Dataset_ContainerAction.
func New__AWS_IoTAnalytics_Dataset_ContainerAction() AWS_IoTAnalytics_Dataset_ContainerAction {
	return AWS_IoTAnalytics_Dataset_ContainerAction{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTAnalytics_Dataset_ContainerAction) GetType() string {
	return AWS_IoTAnalytics_Dataset_ContainerAction__Type
}

// Set__ExecutionRoleArn updates property "ExecutionRoleArn".
func (t AWS_IoTAnalytics_Dataset_ContainerAction) Set__ExecutionRoleArn(v cfz.Expression[string]) AWS_IoTAnalytics_Dataset_ContainerAction {
	t.ExecutionRoleArn = v
	return t
}

// SetV__ExecutionRoleArn updates property "ExecutionRoleArn".
func (t AWS_IoTAnalytics_Dataset_ContainerAction) SetV__ExecutionRoleArn(v string) AWS_IoTAnalytics_Dataset_ContainerAction {
	t.ExecutionRoleArn = cfz.V(v)
	return t
}

// Set__Image updates property "Image".
func (t AWS_IoTAnalytics_Dataset_ContainerAction) Set__Image(v cfz.Expression[string]) AWS_IoTAnalytics_Dataset_ContainerAction {
	t.Image = v
	return t
}

// SetV__Image updates property "Image".
func (t AWS_IoTAnalytics_Dataset_ContainerAction) SetV__Image(v string) AWS_IoTAnalytics_Dataset_ContainerAction {
	t.Image = cfz.V(v)
	return t
}

// Set__ResourceConfiguration updates property "ResourceConfiguration".
func (t AWS_IoTAnalytics_Dataset_ContainerAction) Set__ResourceConfiguration(v cfz.Expression[AWS_IoTAnalytics_Dataset_ResourceConfiguration]) AWS_IoTAnalytics_Dataset_ContainerAction {
	t.ResourceConfiguration = v
	return t
}

// SetV__ResourceConfiguration updates property "ResourceConfiguration".
func (t AWS_IoTAnalytics_Dataset_ContainerAction) SetV__ResourceConfiguration(v AWS_IoTAnalytics_Dataset_ResourceConfiguration) AWS_IoTAnalytics_Dataset_ContainerAction {
	t.ResourceConfiguration = cfz.V(v)
	return t
}

// Set__Variables updates property "Variables".
func (t AWS_IoTAnalytics_Dataset_ContainerAction) Set__Variables(v cfz.ExpressionSlice[AWS_IoTAnalytics_Dataset_Variable]) AWS_IoTAnalytics_Dataset_ContainerAction {
	t.Variables = v
	return t
}

// SetS__Variables updates property "Variables".
func (t AWS_IoTAnalytics_Dataset_ContainerAction) SetS__Variables(v ...cfz.Expression[AWS_IoTAnalytics_Dataset_Variable]) AWS_IoTAnalytics_Dataset_ContainerAction {
	t.Variables = cfz.S(v...)
	return t
}

// SetSV__Variables updates property "Variables".
func (t AWS_IoTAnalytics_Dataset_ContainerAction) SetSV__Variables(v ...AWS_IoTAnalytics_Dataset_Variable) AWS_IoTAnalytics_Dataset_ContainerAction {
	t.Variables = cfz.SV(v...)
	return t
}
