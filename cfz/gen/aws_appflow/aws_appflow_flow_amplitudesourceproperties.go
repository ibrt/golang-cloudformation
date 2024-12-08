// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appflow

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppFlow_Flow_AmplitudeSourceProperties__Type is the CloudFormation type for AWS::AppFlow::Flow.AmplitudeSourceProperties.
	AWS_AppFlow_Flow_AmplitudeSourceProperties__Type = "AWS::AppFlow::Flow.AmplitudeSourceProperties"
)

var (
	// AWS_AppFlow_Flow_AmplitudeSourceProperties__PropertiesMap reports all the CloudFormation properties for AWS::AppFlow::Flow.AmplitudeSourceProperties.
	AWS_AppFlow_Flow_AmplitudeSourceProperties__PropertiesMap = struct {
		Object string
	}{
		Object: "Object",
	}

	// AWS_AppFlow_Flow_AmplitudeSourceProperties__PropertiesSlice reports all the CloudFormation properties for AWS::AppFlow::Flow.AmplitudeSourceProperties.
	AWS_AppFlow_Flow_AmplitudeSourceProperties__PropertiesSlice = []string{
		AWS_AppFlow_Flow_AmplitudeSourceProperties__PropertiesMap.Object,
	}
)

// AWS_AppFlow_Flow_AmplitudeSourceProperties is a binding for AWS::AppFlow::Flow.AmplitudeSourceProperties.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-amplitudesourceproperties.html
type AWS_AppFlow_Flow_AmplitudeSourceProperties struct {
	// Object is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-amplitudesourceproperties.html#cfn-appflow-flow-amplitudesourceproperties-object
	Object cfz.Expression[string] `json:"Object,omitempty"`
}

// New__AWS_AppFlow_Flow_AmplitudeSourceProperties initializes a new AWS_AppFlow_Flow_AmplitudeSourceProperties.
func New__AWS_AppFlow_Flow_AmplitudeSourceProperties() AWS_AppFlow_Flow_AmplitudeSourceProperties {
	return AWS_AppFlow_Flow_AmplitudeSourceProperties{}
}

// GetType returns the CloudFormation type.
func (AWS_AppFlow_Flow_AmplitudeSourceProperties) GetType() string {
	return AWS_AppFlow_Flow_AmplitudeSourceProperties__Type
}

// Set__Object updates property "Object".
func (t AWS_AppFlow_Flow_AmplitudeSourceProperties) Set__Object(v cfz.Expression[string]) AWS_AppFlow_Flow_AmplitudeSourceProperties {
	t.Object = v
	return t
}

// SetV__Object updates property "Object".
func (t AWS_AppFlow_Flow_AmplitudeSourceProperties) SetV__Object(v string) AWS_AppFlow_Flow_AmplitudeSourceProperties {
	t.Object = cfz.V(v)
	return t
}
