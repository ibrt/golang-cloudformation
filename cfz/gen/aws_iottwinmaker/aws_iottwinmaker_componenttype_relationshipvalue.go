// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iottwinmaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTTwinMaker_ComponentType_RelationshipValue__Type is the CloudFormation type for AWS::IoTTwinMaker::ComponentType.RelationshipValue.
	AWS_IoTTwinMaker_ComponentType_RelationshipValue__Type = "AWS::IoTTwinMaker::ComponentType.RelationshipValue"
)

var (
	// AWS_IoTTwinMaker_ComponentType_RelationshipValue__PropertiesMap reports all the CloudFormation properties for AWS::IoTTwinMaker::ComponentType.RelationshipValue.
	AWS_IoTTwinMaker_ComponentType_RelationshipValue__PropertiesMap = struct {
		TargetComponentName string
		TargetEntityId      string
	}{
		TargetComponentName: "TargetComponentName",
		TargetEntityId:      "TargetEntityId",
	}

	// AWS_IoTTwinMaker_ComponentType_RelationshipValue__PropertiesSlice reports all the CloudFormation properties for AWS::IoTTwinMaker::ComponentType.RelationshipValue.
	AWS_IoTTwinMaker_ComponentType_RelationshipValue__PropertiesSlice = []string{
		AWS_IoTTwinMaker_ComponentType_RelationshipValue__PropertiesMap.TargetComponentName,
		AWS_IoTTwinMaker_ComponentType_RelationshipValue__PropertiesMap.TargetEntityId,
	}
)

// AWS_IoTTwinMaker_ComponentType_RelationshipValue is a binding for AWS::IoTTwinMaker::ComponentType.RelationshipValue.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationshipvalue.html
type AWS_IoTTwinMaker_ComponentType_RelationshipValue struct {
	// TargetComponentName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationshipvalue.html#cfn-iottwinmaker-componenttype-relationshipvalue-targetcomponentname
	TargetComponentName cfz.Expression[string] `json:"TargetComponentName,omitempty"`

	// TargetEntityId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationshipvalue.html#cfn-iottwinmaker-componenttype-relationshipvalue-targetentityid
	TargetEntityId cfz.Expression[string] `json:"TargetEntityId,omitempty"`
}

// New__AWS_IoTTwinMaker_ComponentType_RelationshipValue initializes a new AWS_IoTTwinMaker_ComponentType_RelationshipValue.
func New__AWS_IoTTwinMaker_ComponentType_RelationshipValue() AWS_IoTTwinMaker_ComponentType_RelationshipValue {
	return AWS_IoTTwinMaker_ComponentType_RelationshipValue{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTTwinMaker_ComponentType_RelationshipValue) GetType() string {
	return AWS_IoTTwinMaker_ComponentType_RelationshipValue__Type
}

// Set__TargetComponentName updates property "TargetComponentName".
func (t AWS_IoTTwinMaker_ComponentType_RelationshipValue) Set__TargetComponentName(v cfz.Expression[string]) AWS_IoTTwinMaker_ComponentType_RelationshipValue {
	t.TargetComponentName = v
	return t
}

// SetV__TargetComponentName updates property "TargetComponentName".
func (t AWS_IoTTwinMaker_ComponentType_RelationshipValue) SetV__TargetComponentName(v string) AWS_IoTTwinMaker_ComponentType_RelationshipValue {
	t.TargetComponentName = cfz.V(v)
	return t
}

// Set__TargetEntityId updates property "TargetEntityId".
func (t AWS_IoTTwinMaker_ComponentType_RelationshipValue) Set__TargetEntityId(v cfz.Expression[string]) AWS_IoTTwinMaker_ComponentType_RelationshipValue {
	t.TargetEntityId = v
	return t
}

// SetV__TargetEntityId updates property "TargetEntityId".
func (t AWS_IoTTwinMaker_ComponentType_RelationshipValue) SetV__TargetEntityId(v string) AWS_IoTTwinMaker_ComponentType_RelationshipValue {
	t.TargetEntityId = cfz.V(v)
	return t
}
