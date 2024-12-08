package cfspecz

import (
	"github.com/ibrt/golang-utils/memz"
)

// NewDefaultPatchManager initializes a new patch manager and configures it with default, known patchers.
func NewDefaultPatchManager() *PatchManager {
	return NewPatchManager().
		// Special treatment for "CustomResource".
		// Clients can handle it by implementing the cfz.Resource interface.
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::CloudFormation::CustomResource",
		}).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::CloudWatch::AnomalyDetector.MetricDataQueries",
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::CloudWatch::InsightRule.Tags",
		}).
		// References invalid, deleted type "AWS::CloudWatch::InsightRule.Tags",
		// fallback to List(Tag) (as documented)..
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::CloudWatch::InsightRule",
			PropertyName: "Tags",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "Tags",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "",
				ItemType:          "Tag",
			},
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::CodeBuild::Project.FilterGroup",
		}).
		// References invalid, deleted type "AWS::CodeBuild::Project.FilterGroup",
		// fallback to JSON.
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::CodeBuild::Project.ProjectTriggers",
			PropertyName: "FilterGroups",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "",
				ItemType:          "FilterGroup",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "Json",
				Type:              "",
				PrimitiveItemType: "",
				ItemType:          "",
			},
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::DLM::LifecyclePolicy.CrossRegionCopyTargets",
		}).
		// References invalid, deleted type "AWS::DLM::LifecyclePolicy.CrossRegionCopyTargets",
		// fallback to JSON.
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::DLM::LifecyclePolicy",
			PropertyName: "CrossRegionCopyTargets",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "CrossRegionCopyTargets",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "Json",
				Type:              "",
				PrimitiveItemType: "",
				ItemType:          "",
			},
		}).
		// References invalid, deleted type "AWS::DLM::LifecyclePolicy.CrossRegionCopyTargets",
		// fallback to JSON.
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::DLM::LifecyclePolicy.PolicyDetails",
			PropertyName: "CrossRegionCopyTargets",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "CrossRegionCopyTargets",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "Json",
				Type:              "",
				PrimitiveItemType: "",
				ItemType:          "",
			},
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::DLM::LifecyclePolicy.ExcludeTags",
		}).
		// References invalid, deleted type "AWS::DLM::LifecyclePolicy.ExcludeTags",
		// fallback to JSON.
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::DLM::LifecyclePolicy.Exclusions",
			PropertyName: "ExcludeTags",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "ExcludeTags",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "Json",
				Type:              "",
				PrimitiveItemType: "",
				ItemType:          "",
			},
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::DLM::LifecyclePolicy.ExcludeVolumeTypesList",
		}).
		// References invalid, deleted type "AWS::DLM::LifecyclePolicy.ExcludeVolumeTypesList",
		// fallback to JSON.
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::DLM::LifecyclePolicy.Exclusions",
			PropertyName: "ExcludeVolumeTypes",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "ExcludeVolumeTypesList",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "Json",
				Type:              "",
				PrimitiveItemType: "",
				ItemType:          "",
			},
		}).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::DLM::LifecyclePolicy.VolumeTypeValues",
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::Glue::Table.MetadataOperation",
		}).
		// References invalid, deleted type "AWS::Glue::Table.MetadataOperation",
		// fallback to String (as documented).
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::Glue::Table.IcebergInput",
			PropertyName: "MetadataOperation",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "MetadataOperation",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "String",
				Type:              "",
				PrimitiveItemType: "",
				ItemType:          "",
			},
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::Glue::SecurityConfiguration.S3Encryptions",
		}).
		// References invalid, deleted type "AWS::Glue::SecurityConfiguration.S3Encryptions",
		// fallback to JSON.
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::Glue::SecurityConfiguration.EncryptionConfiguration",
			PropertyName: "S3Encryptions",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "S3Encryptions",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "Json",
				Type:              "",
				PrimitiveItemType: "",
				ItemType:          "",
			},
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::LakeFormation::DataLakeSettings.Admins",
		}).
		// References invalid, deleted type "AWS::LakeFormation::DataLakeSettings.Admins",
		//fallback to List(DataLakePrincipal) (as documented).
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::LakeFormation::DataLakeSettings",
			PropertyName: "Admins",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "Admins",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "",
				ItemType:          "DataLakePrincipal",
			},
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::LakeFormation::DataLakeSettings.CreateDatabaseDefaultPermissions",
		}).
		// References invalid, deleted type "AWS::LakeFormation::DataLakeSettings.CreateDatabaseDefaultPermissions",
		// fallback to List(PrincipalPermissions) (as documented).
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::LakeFormation::DataLakeSettings",
			PropertyName: "CreateDatabaseDefaultPermissions",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "CreateDatabaseDefaultPermissions",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "",
				ItemType:          "PrincipalPermissions",
			},
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::LakeFormation::DataLakeSettings.CreateTableDefaultPermissions",
		}).
		// References invalid, deleted type "AWS::LakeFormation::DataLakeSettings.CreateTableDefaultPermissions",
		// fallback to List(PrincipalPermissions) (as documented).
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::LakeFormation::DataLakeSettings",
			PropertyName: "CreateTableDefaultPermissions",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "CreateTableDefaultPermissions",
				PrimitiveItemType: "",
				ItemType:          "",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "",
				ItemType:          "PrincipalPermissions",
			},
		}).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::LakeFormation::DataLakeSettings.ExternalDataFilteringAllowList",
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::SageMaker::EndpointConfig.ClarifyFeatureType",
		}).
		// References invalid, deleted type "AWS::SageMaker::EndpointConfig.ClarifyHeader", fallback to List(String) (as documented).
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::SageMaker::EndpointConfig.ClarifyInferenceConfig",
			PropertyName: "FeatureTypes",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "",
				ItemType:          "ClarifyFeatureType",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "String",
				ItemType:          "",
			},
		}).
		// Invalid type (referenced).
		RegisterRawPatch(&RawPatchDeleteType{
			TypeName: "AWS::SageMaker::EndpointConfig.ClarifyHeader",
		}).
		// References invalid, deleted type "AWS::SageMaker::EndpointConfig.ClarifyHeader", fallback to List(String) (as documented).
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::SageMaker::EndpointConfig.ClarifyInferenceConfig",
			PropertyName: "FeatureHeaders",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "",
				ItemType:          "ClarifyHeader",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "String",
				ItemType:          "",
			},
		}).
		// References invalid, deleted type "AWS::SageMaker::EndpointConfig.ClarifyHeader", fallback to List(String) (as documented).
		RegisterRawPatch(&RawPatchFixPropertyType{
			TypeName:     "AWS::SageMaker::EndpointConfig.ClarifyInferenceConfig",
			PropertyName: "LabelHeaders",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "",
				ItemType:          "ClarifyHeader",
			},
			FixedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "String",
				ItemType:          "",
			},
		}).
		// Special treatment for "Tag" because it is the only structured type shared across resources.
		// It is also the only structured type that does not have a "." in its name.
		RegisterSpecPatch(&SpecPatchDeleteType{
			TypeName:              "Tag",
			ForceIsStructuredType: memz.Ptr(true),
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::CleanRooms::IdNamespaceAssociation", &TypePatchDeleteAttribute{
			AttributeName: "InputReferenceProperties.IdMappingWorkflowsSupported",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "List",
				PrimitiveItemType: "Json",
				ItemType:          "",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::CloudFormation::WaitCondition", &TypePatchDeleteAttribute{
			AttributeName: "Data",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "Json",
				Type:              "",
				PrimitiveItemType: "",
				ItemType:          "",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::IoTTwinMaker::Scene", &TypePatchDeleteAttribute{
			AttributeName: "GeneratedSceneMetadata",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "String",
				ItemType:          "",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::MediaLive::SignalMap", &TypePatchDeleteAttribute{
			AttributeName: "FailedMediaResourceMap",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "",
				ItemType:          "MediaResource",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::MediaLive::SignalMap", &TypePatchDeleteAttribute{
			AttributeName: "MediaResourceMap",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "",
				ItemType:          "MediaResource",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::OpenSearchService::Domain", &TypePatchDeleteAttribute{
			AttributeName: "DomainEndpoints",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "String",
				ItemType:          "",
			},
		}).
		// Undocumented attribute with invalid type.
		RegisterTypePatch("AWS::ServiceCatalog::CloudFormationProvisionedProduct", &TypePatchDeleteAttribute{
			AttributeName: "Outputs",
			ExpectedFields: &PropertyOrAttributeTypeFields{
				PrimitiveType:     "",
				Type:              "Map",
				PrimitiveItemType: "String",
				ItemType:          "",
			},
		})
}
