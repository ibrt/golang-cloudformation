// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_rbin

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Rbin_Rule_UnlockDelay__Type is the CloudFormation type for AWS::Rbin::Rule.UnlockDelay.
	AWS_Rbin_Rule_UnlockDelay__Type = "AWS::Rbin::Rule.UnlockDelay"
)

var (
	// AWS_Rbin_Rule_UnlockDelay__PropertiesMap reports all the CloudFormation properties for AWS::Rbin::Rule.UnlockDelay.
	AWS_Rbin_Rule_UnlockDelay__PropertiesMap = struct {
		UnlockDelayUnit  string
		UnlockDelayValue string
	}{
		UnlockDelayUnit:  "UnlockDelayUnit",
		UnlockDelayValue: "UnlockDelayValue",
	}

	// AWS_Rbin_Rule_UnlockDelay__PropertiesSlice reports all the CloudFormation properties for AWS::Rbin::Rule.UnlockDelay.
	AWS_Rbin_Rule_UnlockDelay__PropertiesSlice = []string{
		AWS_Rbin_Rule_UnlockDelay__PropertiesMap.UnlockDelayUnit,
		AWS_Rbin_Rule_UnlockDelay__PropertiesMap.UnlockDelayValue,
	}
)

// AWS_Rbin_Rule_UnlockDelay is a binding for AWS::Rbin::Rule.UnlockDelay.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-unlockdelay.html
type AWS_Rbin_Rule_UnlockDelay struct {
	// UnlockDelayUnit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-unlockdelay.html#cfn-rbin-rule-unlockdelay-unlockdelayunit
	UnlockDelayUnit cfz.Expression[string] `json:"UnlockDelayUnit,omitempty"`

	// UnlockDelayValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-unlockdelay.html#cfn-rbin-rule-unlockdelay-unlockdelayvalue
	UnlockDelayValue cfz.Expression[int32] `json:"UnlockDelayValue,omitempty"`
}

// New__AWS_Rbin_Rule_UnlockDelay initializes a new AWS_Rbin_Rule_UnlockDelay.
func New__AWS_Rbin_Rule_UnlockDelay() AWS_Rbin_Rule_UnlockDelay {
	return AWS_Rbin_Rule_UnlockDelay{}
}

// GetType returns the CloudFormation type.
func (AWS_Rbin_Rule_UnlockDelay) GetType() string {
	return AWS_Rbin_Rule_UnlockDelay__Type
}

// Set__UnlockDelayUnit updates property "UnlockDelayUnit".
func (t AWS_Rbin_Rule_UnlockDelay) Set__UnlockDelayUnit(v cfz.Expression[string]) AWS_Rbin_Rule_UnlockDelay {
	t.UnlockDelayUnit = v
	return t
}

// SetV__UnlockDelayUnit updates property "UnlockDelayUnit".
func (t AWS_Rbin_Rule_UnlockDelay) SetV__UnlockDelayUnit(v string) AWS_Rbin_Rule_UnlockDelay {
	t.UnlockDelayUnit = cfz.V(v)
	return t
}

// Set__UnlockDelayValue updates property "UnlockDelayValue".
func (t AWS_Rbin_Rule_UnlockDelay) Set__UnlockDelayValue(v cfz.Expression[int32]) AWS_Rbin_Rule_UnlockDelay {
	t.UnlockDelayValue = v
	return t
}

// SetV__UnlockDelayValue updates property "UnlockDelayValue".
func (t AWS_Rbin_Rule_UnlockDelay) SetV__UnlockDelayValue(v int32) AWS_Rbin_Rule_UnlockDelay {
	t.UnlockDelayValue = cfz.V(v)
	return t
}
