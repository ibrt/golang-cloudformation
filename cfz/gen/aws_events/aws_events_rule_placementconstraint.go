// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_events

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Events_Rule_PlacementConstraint__Type is the CloudFormation type for AWS::Events::Rule.PlacementConstraint.
	AWS_Events_Rule_PlacementConstraint__Type = "AWS::Events::Rule.PlacementConstraint"
)

var (
	// AWS_Events_Rule_PlacementConstraint__PropertiesMap reports all the CloudFormation properties for AWS::Events::Rule.PlacementConstraint.
	AWS_Events_Rule_PlacementConstraint__PropertiesMap = struct {
		Expression string
		Type       string
	}{
		Expression: "Expression",
		Type:       "Type",
	}

	// AWS_Events_Rule_PlacementConstraint__PropertiesSlice reports all the CloudFormation properties for AWS::Events::Rule.PlacementConstraint.
	AWS_Events_Rule_PlacementConstraint__PropertiesSlice = []string{
		AWS_Events_Rule_PlacementConstraint__PropertiesMap.Expression,
		AWS_Events_Rule_PlacementConstraint__PropertiesMap.Type,
	}
)

// AWS_Events_Rule_PlacementConstraint is a binding for AWS::Events::Rule.PlacementConstraint.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-placementconstraint.html
type AWS_Events_Rule_PlacementConstraint struct {
	// Expression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-placementconstraint.html#cfn-events-rule-placementconstraint-expression
	Expression cfz.Expression[string] `json:"Expression,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-placementconstraint.html#cfn-events-rule-placementconstraint-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_Events_Rule_PlacementConstraint initializes a new AWS_Events_Rule_PlacementConstraint.
func New__AWS_Events_Rule_PlacementConstraint() AWS_Events_Rule_PlacementConstraint {
	return AWS_Events_Rule_PlacementConstraint{}
}

// GetType returns the CloudFormation type.
func (AWS_Events_Rule_PlacementConstraint) GetType() string {
	return AWS_Events_Rule_PlacementConstraint__Type
}

// Set__Expression updates property "Expression".
func (t AWS_Events_Rule_PlacementConstraint) Set__Expression(v cfz.Expression[string]) AWS_Events_Rule_PlacementConstraint {
	t.Expression = v
	return t
}

// SetV__Expression updates property "Expression".
func (t AWS_Events_Rule_PlacementConstraint) SetV__Expression(v string) AWS_Events_Rule_PlacementConstraint {
	t.Expression = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_Events_Rule_PlacementConstraint) Set__Type(v cfz.Expression[string]) AWS_Events_Rule_PlacementConstraint {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_Events_Rule_PlacementConstraint) SetV__Type(v string) AWS_Events_Rule_PlacementConstraint {
	t.Type = cfz.V(v)
	return t
}
