// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appflow

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppFlow_Flow_ZendeskSourceProperties__Type is the CloudFormation type for AWS::AppFlow::Flow.ZendeskSourceProperties.
	AWS_AppFlow_Flow_ZendeskSourceProperties__Type = "AWS::AppFlow::Flow.ZendeskSourceProperties"
)

var (
	// AWS_AppFlow_Flow_ZendeskSourceProperties__PropertiesMap reports all the CloudFormation properties for AWS::AppFlow::Flow.ZendeskSourceProperties.
	AWS_AppFlow_Flow_ZendeskSourceProperties__PropertiesMap = struct {
		Object string
	}{
		Object: "Object",
	}

	// AWS_AppFlow_Flow_ZendeskSourceProperties__PropertiesSlice reports all the CloudFormation properties for AWS::AppFlow::Flow.ZendeskSourceProperties.
	AWS_AppFlow_Flow_ZendeskSourceProperties__PropertiesSlice = []string{
		AWS_AppFlow_Flow_ZendeskSourceProperties__PropertiesMap.Object,
	}
)

// AWS_AppFlow_Flow_ZendeskSourceProperties is a binding for AWS::AppFlow::Flow.ZendeskSourceProperties.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-zendesksourceproperties.html
type AWS_AppFlow_Flow_ZendeskSourceProperties struct {
	// Object is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-zendesksourceproperties.html#cfn-appflow-flow-zendesksourceproperties-object
	Object cfz.Expression[string] `json:"Object,omitempty"`
}

// New__AWS_AppFlow_Flow_ZendeskSourceProperties initializes a new AWS_AppFlow_Flow_ZendeskSourceProperties.
func New__AWS_AppFlow_Flow_ZendeskSourceProperties() AWS_AppFlow_Flow_ZendeskSourceProperties {
	return AWS_AppFlow_Flow_ZendeskSourceProperties{}
}

// GetType returns the CloudFormation type.
func (AWS_AppFlow_Flow_ZendeskSourceProperties) GetType() string {
	return AWS_AppFlow_Flow_ZendeskSourceProperties__Type
}

// Set__Object updates property "Object".
func (t AWS_AppFlow_Flow_ZendeskSourceProperties) Set__Object(v cfz.Expression[string]) AWS_AppFlow_Flow_ZendeskSourceProperties {
	t.Object = v
	return t
}

// SetV__Object updates property "Object".
func (t AWS_AppFlow_Flow_ZendeskSourceProperties) SetV__Object(v string) AWS_AppFlow_Flow_ZendeskSourceProperties {
	t.Object = cfz.V(v)
	return t
}
