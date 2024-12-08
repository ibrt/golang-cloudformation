// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iottwinmaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTTwinMaker_Entity_Relationship__Type is the CloudFormation type for AWS::IoTTwinMaker::Entity.Relationship.
	AWS_IoTTwinMaker_Entity_Relationship__Type = "AWS::IoTTwinMaker::Entity.Relationship"
)

var (
	// AWS_IoTTwinMaker_Entity_Relationship__PropertiesMap reports all the CloudFormation properties for AWS::IoTTwinMaker::Entity.Relationship.
	AWS_IoTTwinMaker_Entity_Relationship__PropertiesMap = struct {
		RelationshipType      string
		TargetComponentTypeId string
	}{
		RelationshipType:      "RelationshipType",
		TargetComponentTypeId: "TargetComponentTypeId",
	}

	// AWS_IoTTwinMaker_Entity_Relationship__PropertiesSlice reports all the CloudFormation properties for AWS::IoTTwinMaker::Entity.Relationship.
	AWS_IoTTwinMaker_Entity_Relationship__PropertiesSlice = []string{
		AWS_IoTTwinMaker_Entity_Relationship__PropertiesMap.RelationshipType,
		AWS_IoTTwinMaker_Entity_Relationship__PropertiesMap.TargetComponentTypeId,
	}
)

// AWS_IoTTwinMaker_Entity_Relationship is a binding for AWS::IoTTwinMaker::Entity.Relationship.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-relationship.html
type AWS_IoTTwinMaker_Entity_Relationship struct {
	// RelationshipType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-relationship.html#cfn-iottwinmaker-entity-relationship-relationshiptype
	RelationshipType cfz.Expression[string] `json:"RelationshipType,omitempty"`

	// TargetComponentTypeId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-relationship.html#cfn-iottwinmaker-entity-relationship-targetcomponenttypeid
	TargetComponentTypeId cfz.Expression[string] `json:"TargetComponentTypeId,omitempty"`
}

// New__AWS_IoTTwinMaker_Entity_Relationship initializes a new AWS_IoTTwinMaker_Entity_Relationship.
func New__AWS_IoTTwinMaker_Entity_Relationship() AWS_IoTTwinMaker_Entity_Relationship {
	return AWS_IoTTwinMaker_Entity_Relationship{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTTwinMaker_Entity_Relationship) GetType() string {
	return AWS_IoTTwinMaker_Entity_Relationship__Type
}

// Set__RelationshipType updates property "RelationshipType".
func (t AWS_IoTTwinMaker_Entity_Relationship) Set__RelationshipType(v cfz.Expression[string]) AWS_IoTTwinMaker_Entity_Relationship {
	t.RelationshipType = v
	return t
}

// SetV__RelationshipType updates property "RelationshipType".
func (t AWS_IoTTwinMaker_Entity_Relationship) SetV__RelationshipType(v string) AWS_IoTTwinMaker_Entity_Relationship {
	t.RelationshipType = cfz.V(v)
	return t
}

// Set__TargetComponentTypeId updates property "TargetComponentTypeId".
func (t AWS_IoTTwinMaker_Entity_Relationship) Set__TargetComponentTypeId(v cfz.Expression[string]) AWS_IoTTwinMaker_Entity_Relationship {
	t.TargetComponentTypeId = v
	return t
}

// SetV__TargetComponentTypeId updates property "TargetComponentTypeId".
func (t AWS_IoTTwinMaker_Entity_Relationship) SetV__TargetComponentTypeId(v string) AWS_IoTTwinMaker_Entity_Relationship {
	t.TargetComponentTypeId = cfz.V(v)
	return t
}
