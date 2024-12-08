package cfspecz

import (
	"github.com/ibrt/golang-utils/memz"
)

// NewDefaultPatchManager initializes a new patch manager and configures it with default, known patchers.
func NewDefaultPatchManager() *PatchManager {
	return NewPatchManager().
		// Special treatment for "Tag" because it is the only structured type shared across resources.
		// It is also the only structured type that does not have a "." in its name.
		RegisterSpecPatch(&SpecPatchDeleteType{
			TypeName:              "Tag",
			ForceIsStructuredType: memz.Ptr(true),
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::CleanRooms::IdNamespaceAssociation", &TypePatchDeleteAttribute{
			AttributeName: "InputReferenceProperties.IdMappingWorkflowsSupported",
			ExpectedFields: &TypeRelatedFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "Json",
				ItemType:          "",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::CloudFormation::WaitCondition", &TypePatchDeleteAttribute{
			AttributeName: "Data",
			ExpectedFields: &TypeRelatedFields{
				PrimitiveType:     "Json",
				Type:              "",
				PrimitiveItemType: "",
				ItemType:          "",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::IoTTwinMaker::Scene", &TypePatchDeleteAttribute{
			AttributeName: "GeneratedSceneMetadata",
			ExpectedFields: &TypeRelatedFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "String",
				ItemType:          "",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::MediaLive::SignalMap", &TypePatchDeleteAttribute{
			AttributeName: "FailedMediaResourceMap",
			ExpectedFields: &TypeRelatedFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "",
				ItemType:          "MediaResource",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::MediaLive::SignalMap", &TypePatchDeleteAttribute{
			AttributeName: "MediaResourceMap",
			ExpectedFields: &TypeRelatedFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "",
				ItemType:          "MediaResource",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::OpenSearchService::Domain", &TypePatchDeleteAttribute{
			AttributeName: "DomainEndpoints",
			ExpectedFields: &TypeRelatedFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "String",
				ItemType:          "",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::ServiceCatalog::CloudFormationProvisionedProduct", &TypePatchDeleteAttribute{
			AttributeName: "Outputs",
			ExpectedFields: &TypeRelatedFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "String",
				ItemType:          "",
			},
		})

}
