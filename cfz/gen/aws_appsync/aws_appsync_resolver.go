// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appsync

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_AppSync_Resolver)(nil)
	_ cfz.Resource                   = (*AWS_AppSync_Resolver)(nil)
)

const (
	// AWS_AppSync_Resolver__Type is the CloudFormation type for AWS::AppSync::Resolver.
	AWS_AppSync_Resolver__Type = "AWS::AppSync::Resolver"
)

var (
	// AWS_AppSync_Resolver__AttributesMap reports all the CloudFormation attributes for AWS::AppSync::Resolver.
	AWS_AppSync_Resolver__AttributesMap = struct {
		FieldName   string
		ResolverArn string
		TypeName    string
	}{
		FieldName:   "FieldName",
		ResolverArn: "ResolverArn",
		TypeName:    "TypeName",
	}

	// AWS_AppSync_Resolver__AttributesSlice reports all the CloudFormation attributes for AWS::AppSync::Resolver.
	AWS_AppSync_Resolver__AttributesSlice = []string{
		AWS_AppSync_Resolver__AttributesMap.FieldName,
		AWS_AppSync_Resolver__AttributesMap.ResolverArn,
		AWS_AppSync_Resolver__AttributesMap.TypeName,
	}
)

var (
	// AWS_AppSync_Resolver__PropertiesMap reports all the CloudFormation properties for AWS::AppSync::Resolver.
	AWS_AppSync_Resolver__PropertiesMap = struct {
		ApiId                             string
		CachingConfig                     string
		Code                              string
		CodeS3Location                    string
		DataSourceName                    string
		FieldName                         string
		Kind                              string
		MaxBatchSize                      string
		MetricsConfig                     string
		PipelineConfig                    string
		RequestMappingTemplate            string
		RequestMappingTemplateS3Location  string
		ResponseMappingTemplate           string
		ResponseMappingTemplateS3Location string
		Runtime                           string
		SyncConfig                        string
		TypeName                          string
	}{
		ApiId:                             "ApiId",
		CachingConfig:                     "CachingConfig",
		Code:                              "Code",
		CodeS3Location:                    "CodeS3Location",
		DataSourceName:                    "DataSourceName",
		FieldName:                         "FieldName",
		Kind:                              "Kind",
		MaxBatchSize:                      "MaxBatchSize",
		MetricsConfig:                     "MetricsConfig",
		PipelineConfig:                    "PipelineConfig",
		RequestMappingTemplate:            "RequestMappingTemplate",
		RequestMappingTemplateS3Location:  "RequestMappingTemplateS3Location",
		ResponseMappingTemplate:           "ResponseMappingTemplate",
		ResponseMappingTemplateS3Location: "ResponseMappingTemplateS3Location",
		Runtime:                           "Runtime",
		SyncConfig:                        "SyncConfig",
		TypeName:                          "TypeName",
	}

	// AWS_AppSync_Resolver__PropertiesSlice reports all the CloudFormation properties for AWS::AppSync::Resolver.
	AWS_AppSync_Resolver__PropertiesSlice = []string{
		AWS_AppSync_Resolver__PropertiesMap.ApiId,
		AWS_AppSync_Resolver__PropertiesMap.CachingConfig,
		AWS_AppSync_Resolver__PropertiesMap.Code,
		AWS_AppSync_Resolver__PropertiesMap.CodeS3Location,
		AWS_AppSync_Resolver__PropertiesMap.DataSourceName,
		AWS_AppSync_Resolver__PropertiesMap.FieldName,
		AWS_AppSync_Resolver__PropertiesMap.Kind,
		AWS_AppSync_Resolver__PropertiesMap.MaxBatchSize,
		AWS_AppSync_Resolver__PropertiesMap.MetricsConfig,
		AWS_AppSync_Resolver__PropertiesMap.PipelineConfig,
		AWS_AppSync_Resolver__PropertiesMap.RequestMappingTemplate,
		AWS_AppSync_Resolver__PropertiesMap.RequestMappingTemplateS3Location,
		AWS_AppSync_Resolver__PropertiesMap.ResponseMappingTemplate,
		AWS_AppSync_Resolver__PropertiesMap.ResponseMappingTemplateS3Location,
		AWS_AppSync_Resolver__PropertiesMap.Runtime,
		AWS_AppSync_Resolver__PropertiesMap.SyncConfig,
		AWS_AppSync_Resolver__PropertiesMap.TypeName,
	}
)

// AWS_AppSync_Resolver is a binding for AWS::AppSync::Resolver.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html
type AWS_AppSync_Resolver struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ApiId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-apiid
	ApiId cfz.Expression[string] `json:"ApiId,omitempty"`

	// CachingConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-cachingconfig
	CachingConfig cfz.Expression[AWS_AppSync_Resolver_CachingConfig] `json:"CachingConfig,omitempty"`

	// Code is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-code
	Code cfz.Expression[string] `json:"Code,omitempty"`

	// CodeS3Location is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-codes3location
	CodeS3Location cfz.Expression[string] `json:"CodeS3Location,omitempty"`

	// DataSourceName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-datasourcename
	DataSourceName cfz.Expression[string] `json:"DataSourceName,omitempty"`

	// FieldName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-fieldname
	FieldName cfz.Expression[string] `json:"FieldName,omitempty"`

	// Kind is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-kind
	Kind cfz.Expression[string] `json:"Kind,omitempty"`

	// MaxBatchSize is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-maxbatchsize
	MaxBatchSize cfz.Expression[int32] `json:"MaxBatchSize,omitempty"`

	// MetricsConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-metricsconfig
	MetricsConfig cfz.Expression[string] `json:"MetricsConfig,omitempty"`

	// PipelineConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-pipelineconfig
	PipelineConfig cfz.Expression[AWS_AppSync_Resolver_PipelineConfig] `json:"PipelineConfig,omitempty"`

	// RequestMappingTemplate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplate
	RequestMappingTemplate cfz.Expression[string] `json:"RequestMappingTemplate,omitempty"`

	// RequestMappingTemplateS3Location is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplates3location
	RequestMappingTemplateS3Location cfz.Expression[string] `json:"RequestMappingTemplateS3Location,omitempty"`

	// ResponseMappingTemplate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplate
	ResponseMappingTemplate cfz.Expression[string] `json:"ResponseMappingTemplate,omitempty"`

	// ResponseMappingTemplateS3Location is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplates3location
	ResponseMappingTemplateS3Location cfz.Expression[string] `json:"ResponseMappingTemplateS3Location,omitempty"`

	// Runtime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-runtime
	Runtime cfz.Expression[AWS_AppSync_Resolver_AppSyncRuntime] `json:"Runtime,omitempty"`

	// SyncConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-syncconfig
	SyncConfig cfz.Expression[AWS_AppSync_Resolver_SyncConfig] `json:"SyncConfig,omitempty"`

	// TypeName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-typename
	TypeName cfz.Expression[string] `json:"TypeName,omitempty"`
}

// New__AWS_AppSync_Resolver initializes a new *AWS_AppSync_Resolver.
func New__AWS_AppSync_Resolver(logicalName string) *AWS_AppSync_Resolver {
	return &AWS_AppSync_Resolver{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_AppSync_Resolver) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_AppSync_Resolver) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_AppSync_Resolver) GetType() string {
	return AWS_AppSync_Resolver__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_AppSync_Resolver) Set__LogicalName(v string) *AWS_AppSync_Resolver {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_AppSync_Resolver) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_AppSync_Resolver {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_AppSync_Resolver) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_AppSync_Resolver {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_AppSync_Resolver) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_AppSync_Resolver {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_AppSync_Resolver) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_AppSync_Resolver {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_AppSync_Resolver) Set__RequestedOutputs(v []cfz.Output) *AWS_AppSync_Resolver {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_AppSync_Resolver) Add__RequestedOutputs(v ...cfz.Output) *AWS_AppSync_Resolver {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ApiId updates property "ApiId".
func (t *AWS_AppSync_Resolver) Set__ApiId(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.ApiId = v
	return t
}

// SetV__ApiId updates property "ApiId".
func (t *AWS_AppSync_Resolver) SetV__ApiId(v string) *AWS_AppSync_Resolver {
	t.ApiId = cfz.V(v)
	return t
}

// Set__CachingConfig updates property "CachingConfig".
func (t *AWS_AppSync_Resolver) Set__CachingConfig(v cfz.Expression[AWS_AppSync_Resolver_CachingConfig]) *AWS_AppSync_Resolver {
	t.CachingConfig = v
	return t
}

// SetV__CachingConfig updates property "CachingConfig".
func (t *AWS_AppSync_Resolver) SetV__CachingConfig(v AWS_AppSync_Resolver_CachingConfig) *AWS_AppSync_Resolver {
	t.CachingConfig = cfz.V(v)
	return t
}

// Set__Code updates property "Code".
func (t *AWS_AppSync_Resolver) Set__Code(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.Code = v
	return t
}

// SetV__Code updates property "Code".
func (t *AWS_AppSync_Resolver) SetV__Code(v string) *AWS_AppSync_Resolver {
	t.Code = cfz.V(v)
	return t
}

// Set__CodeS3Location updates property "CodeS3Location".
func (t *AWS_AppSync_Resolver) Set__CodeS3Location(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.CodeS3Location = v
	return t
}

// SetV__CodeS3Location updates property "CodeS3Location".
func (t *AWS_AppSync_Resolver) SetV__CodeS3Location(v string) *AWS_AppSync_Resolver {
	t.CodeS3Location = cfz.V(v)
	return t
}

// Set__DataSourceName updates property "DataSourceName".
func (t *AWS_AppSync_Resolver) Set__DataSourceName(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.DataSourceName = v
	return t
}

// SetV__DataSourceName updates property "DataSourceName".
func (t *AWS_AppSync_Resolver) SetV__DataSourceName(v string) *AWS_AppSync_Resolver {
	t.DataSourceName = cfz.V(v)
	return t
}

// Set__FieldName updates property "FieldName".
func (t *AWS_AppSync_Resolver) Set__FieldName(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.FieldName = v
	return t
}

// SetV__FieldName updates property "FieldName".
func (t *AWS_AppSync_Resolver) SetV__FieldName(v string) *AWS_AppSync_Resolver {
	t.FieldName = cfz.V(v)
	return t
}

// Set__Kind updates property "Kind".
func (t *AWS_AppSync_Resolver) Set__Kind(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.Kind = v
	return t
}

// SetV__Kind updates property "Kind".
func (t *AWS_AppSync_Resolver) SetV__Kind(v string) *AWS_AppSync_Resolver {
	t.Kind = cfz.V(v)
	return t
}

// Set__MaxBatchSize updates property "MaxBatchSize".
func (t *AWS_AppSync_Resolver) Set__MaxBatchSize(v cfz.Expression[int32]) *AWS_AppSync_Resolver {
	t.MaxBatchSize = v
	return t
}

// SetV__MaxBatchSize updates property "MaxBatchSize".
func (t *AWS_AppSync_Resolver) SetV__MaxBatchSize(v int32) *AWS_AppSync_Resolver {
	t.MaxBatchSize = cfz.V(v)
	return t
}

// Set__MetricsConfig updates property "MetricsConfig".
func (t *AWS_AppSync_Resolver) Set__MetricsConfig(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.MetricsConfig = v
	return t
}

// SetV__MetricsConfig updates property "MetricsConfig".
func (t *AWS_AppSync_Resolver) SetV__MetricsConfig(v string) *AWS_AppSync_Resolver {
	t.MetricsConfig = cfz.V(v)
	return t
}

// Set__PipelineConfig updates property "PipelineConfig".
func (t *AWS_AppSync_Resolver) Set__PipelineConfig(v cfz.Expression[AWS_AppSync_Resolver_PipelineConfig]) *AWS_AppSync_Resolver {
	t.PipelineConfig = v
	return t
}

// SetV__PipelineConfig updates property "PipelineConfig".
func (t *AWS_AppSync_Resolver) SetV__PipelineConfig(v AWS_AppSync_Resolver_PipelineConfig) *AWS_AppSync_Resolver {
	t.PipelineConfig = cfz.V(v)
	return t
}

// Set__RequestMappingTemplate updates property "RequestMappingTemplate".
func (t *AWS_AppSync_Resolver) Set__RequestMappingTemplate(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.RequestMappingTemplate = v
	return t
}

// SetV__RequestMappingTemplate updates property "RequestMappingTemplate".
func (t *AWS_AppSync_Resolver) SetV__RequestMappingTemplate(v string) *AWS_AppSync_Resolver {
	t.RequestMappingTemplate = cfz.V(v)
	return t
}

// Set__RequestMappingTemplateS3Location updates property "RequestMappingTemplateS3Location".
func (t *AWS_AppSync_Resolver) Set__RequestMappingTemplateS3Location(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.RequestMappingTemplateS3Location = v
	return t
}

// SetV__RequestMappingTemplateS3Location updates property "RequestMappingTemplateS3Location".
func (t *AWS_AppSync_Resolver) SetV__RequestMappingTemplateS3Location(v string) *AWS_AppSync_Resolver {
	t.RequestMappingTemplateS3Location = cfz.V(v)
	return t
}

// Set__ResponseMappingTemplate updates property "ResponseMappingTemplate".
func (t *AWS_AppSync_Resolver) Set__ResponseMappingTemplate(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.ResponseMappingTemplate = v
	return t
}

// SetV__ResponseMappingTemplate updates property "ResponseMappingTemplate".
func (t *AWS_AppSync_Resolver) SetV__ResponseMappingTemplate(v string) *AWS_AppSync_Resolver {
	t.ResponseMappingTemplate = cfz.V(v)
	return t
}

// Set__ResponseMappingTemplateS3Location updates property "ResponseMappingTemplateS3Location".
func (t *AWS_AppSync_Resolver) Set__ResponseMappingTemplateS3Location(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.ResponseMappingTemplateS3Location = v
	return t
}

// SetV__ResponseMappingTemplateS3Location updates property "ResponseMappingTemplateS3Location".
func (t *AWS_AppSync_Resolver) SetV__ResponseMappingTemplateS3Location(v string) *AWS_AppSync_Resolver {
	t.ResponseMappingTemplateS3Location = cfz.V(v)
	return t
}

// Set__Runtime updates property "Runtime".
func (t *AWS_AppSync_Resolver) Set__Runtime(v cfz.Expression[AWS_AppSync_Resolver_AppSyncRuntime]) *AWS_AppSync_Resolver {
	t.Runtime = v
	return t
}

// SetV__Runtime updates property "Runtime".
func (t *AWS_AppSync_Resolver) SetV__Runtime(v AWS_AppSync_Resolver_AppSyncRuntime) *AWS_AppSync_Resolver {
	t.Runtime = cfz.V(v)
	return t
}

// Set__SyncConfig updates property "SyncConfig".
func (t *AWS_AppSync_Resolver) Set__SyncConfig(v cfz.Expression[AWS_AppSync_Resolver_SyncConfig]) *AWS_AppSync_Resolver {
	t.SyncConfig = v
	return t
}

// SetV__SyncConfig updates property "SyncConfig".
func (t *AWS_AppSync_Resolver) SetV__SyncConfig(v AWS_AppSync_Resolver_SyncConfig) *AWS_AppSync_Resolver {
	t.SyncConfig = cfz.V(v)
	return t
}

// Set__TypeName updates property "TypeName".
func (t *AWS_AppSync_Resolver) Set__TypeName(v cfz.Expression[string]) *AWS_AppSync_Resolver {
	t.TypeName = v
	return t
}

// SetV__TypeName updates property "TypeName".
func (t *AWS_AppSync_Resolver) SetV__TypeName(v string) *AWS_AppSync_Resolver {
	t.TypeName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_AppSync_Resolver) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__FieldName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: FieldName
func (t *AWS_AppSync_Resolver) GetAtt__FieldName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppSync_Resolver__AttributesMap.FieldName))
}

// GetAtt__ResolverArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ResolverArn
func (t *AWS_AppSync_Resolver) GetAtt__ResolverArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppSync_Resolver__AttributesMap.ResolverArn))
}

// GetAtt__TypeName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TypeName
func (t *AWS_AppSync_Resolver) GetAtt__TypeName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppSync_Resolver__AttributesMap.TypeName))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_AppSync_Resolver) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__FieldName returns a conventionally configured output for an attribute of this resource.
// Attribute: FieldName
func (t *AWS_AppSync_Resolver) GetConventionalOutputAtt__FieldName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFieldName", t.GetAtt__FieldName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ResolverArn returns a conventionally configured output for an attribute of this resource.
// Attribute: ResolverArn
func (t *AWS_AppSync_Resolver) GetConventionalOutputAtt__ResolverArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttResolverArn", t.GetAtt__ResolverArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TypeName returns a conventionally configured output for an attribute of this resource.
// Attribute: TypeName
func (t *AWS_AppSync_Resolver) GetConventionalOutputAtt__TypeName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTypeName", t.GetAtt__TypeName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_AppSync_Resolver) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_AppSync_Resolver

	return json.Marshal(struct {
		Type                string                          `json:"Type"`
		DependsOn           []string                        `json:"DependsOn,omitempty"`
		DeletionPolicy      cfz.ResourceDeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Properties          *Properties                     `json:"Properties,omitempty"`
	}{
		Type:       t.GetType(),
		DependsOn:  t.getDependsOn(),
		Properties: (*Properties)(t),
	})
}

func (t *AWS_AppSync_Resolver) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
