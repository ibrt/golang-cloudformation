package cfspecz

/*
// NewDefaultSpecPatchManager initializes a new spec patch manager and configures it with default, known patchers.
func NewDefaultSpecPatchManager() *PatchManager {
	return NewPatchManager().
		RegisterRawPatch(&RawPatchMalformedType{
			TypeName: "AWS::CodeBuild::Project.FilterGroup",
			PropertyFixes: []*RawPatchMalformedTypePropertyFix{
				{
					TypeName:     "AWS::CodeBuild::Project.ProjectTriggers",
					PropertyName: "FilterGroups",
					ExpectedFields: &RawPatchMalformedTypeFields{
						PrimitiveType:     "",
						Type:              "List",
						PrimitiveItemType: "",
						ItemType:          "FilterGroup",
					},
				},
			},
		}).
		// Spurious structured type, property is List(List(WebhookFilter)).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::CodeBuild::Project.FilterGroup"}).
		RegisterRawPatchSetPath([]string{"PropertyTypes", "AWS::CodeBuild::Project.ProjectTriggers", "Properties", "FilterGroups", "Type"}, "List", nil).
		RegisterRawPatchSetPath([]string{"PropertyTypes", "AWS::CodeBuild::Project.ProjectTriggers", "Properties", "FilterGroups", "ItemType"}, "FilterGroup", nil).
		RegisterRawPatchSetPath([]string{"PropertyTypes", "AWS::CodeBuild::Project.ProjectTriggers", "Properties", "FilterGroups", "PrimitiveType"}, nil, "Json").

		// Spurious structured type, property is List(Tag).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::DLM::LifecyclePolicy.ExcludeTags"}).

		// Spurious structured type, property is List(VolumeTypeValues)
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::DLM::LifecyclePolicy.ExcludeVolumeTypesList"}).

		// Note: all these structured types are malformed and don't appear in the CloudFormation documentation.
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::CloudWatch::AnomalyDetector.MetricDataQueries"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::CloudWatch::InsightRule.Tags"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::DLM::LifecyclePolicy.CrossRegionCopyTargets"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::DLM::LifecyclePolicy.VolumeTypeValues"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::Glue::SecurityConfiguration.S3Encryptions"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::Glue::Table.MetadataOperation"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::LakeFormation::DataLakeSettings.Admins"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::LakeFormation::DataLakeSettings.CreateDatabaseDefaultPermissions"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::LakeFormation::DataLakeSettings.CreateTableDefaultPermissions"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::LakeFormation::DataLakeSettings.ExternalDataFilteringAllowList"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::SageMaker::EndpointConfig.ClarifyFeatureType"}).
		RegisterRawPatchDelPath([]string{"PropertyTypes", "AWS::SageMaker::EndpointConfig.ClarifyHeader"}).
		// Note: we currently do not natively support custom resources, although clients can create their own by implementing cfz.Resource.
		// We remove it here because it has an unexpected "AdditionalProperties" field.
		RegisterRawPatchDelPath([]string{"ResourceTypes", "AWS::CloudFormation::CustomResource"}).
		RegisterSpecPatch(
			func(ic *SpecIssueCollector, s *Spec) {
				if _, ok := s.PropertyTypes["Tag"]; ok {
					delete(s.PropertyTypes, "Tag")
					return
				}

				ic.CollectIssue(s, "outdated patch: missing property type: 'Tag'")
			}).
		RegisterTypePatchDeleteAttribute(
			"AWS::CleanRooms::IdNamespaceAssociation",
			"InputReferenceProperties.IdMappingWorkflowsSupported",
			func(a *Attribute) bool { return a.Type == "List" && a.PrimitiveItemType == "Json" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::CloudFormation::WaitCondition",
			"Data",
			func(a *Attribute) bool { return a.PrimitiveType == "Json" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::IoTTwinMaker::Scene",
			"GeneratedSceneMetadata",
			func(a *Attribute) bool { return a.Type == "Map" && a.PrimitiveItemType == "String" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::MediaLive::SignalMap",
			"MediaResourceMap",
			func(a *Attribute) bool { return a.Type == "Map" && a.ItemType == "MediaResource" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::MediaLive::SignalMap",
			"FailedMediaResourceMap",
			func(a *Attribute) bool { return a.Type == "Map" && a.ItemType == "MediaResource" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::OpenSearchService::Domain",
			"DomainEndpoints",
			func(a *Attribute) bool { return a.Type == "Map" && a.PrimitiveItemType == "String" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::ServiceCatalog::CloudFormationProvisionedProduct",
			"Outputs",
			func(a *Attribute) bool { return a.Type == "Map" && a.PrimitiveItemType == "String" })
}
*/
