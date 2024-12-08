// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_connect

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Connect_InstanceStorageConfig)(nil)
	_ cfz.Resource                   = (*AWS_Connect_InstanceStorageConfig)(nil)
)

const (
	// AWS_Connect_InstanceStorageConfig__Type is the CloudFormation type for AWS::Connect::InstanceStorageConfig.
	AWS_Connect_InstanceStorageConfig__Type = "AWS::Connect::InstanceStorageConfig"
)

var (
	// AWS_Connect_InstanceStorageConfig__AttributesMap reports all the CloudFormation attributes for AWS::Connect::InstanceStorageConfig.
	AWS_Connect_InstanceStorageConfig__AttributesMap = struct {
		AssociationId string
	}{
		AssociationId: "AssociationId",
	}

	// AWS_Connect_InstanceStorageConfig__AttributesSlice reports all the CloudFormation attributes for AWS::Connect::InstanceStorageConfig.
	AWS_Connect_InstanceStorageConfig__AttributesSlice = []string{
		AWS_Connect_InstanceStorageConfig__AttributesMap.AssociationId,
	}
)

var (
	// AWS_Connect_InstanceStorageConfig__PropertiesMap reports all the CloudFormation properties for AWS::Connect::InstanceStorageConfig.
	AWS_Connect_InstanceStorageConfig__PropertiesMap = struct {
		InstanceArn              string
		KinesisFirehoseConfig    string
		KinesisStreamConfig      string
		KinesisVideoStreamConfig string
		ResourceType             string
		S3Config                 string
		StorageType              string
	}{
		InstanceArn:              "InstanceArn",
		KinesisFirehoseConfig:    "KinesisFirehoseConfig",
		KinesisStreamConfig:      "KinesisStreamConfig",
		KinesisVideoStreamConfig: "KinesisVideoStreamConfig",
		ResourceType:             "ResourceType",
		S3Config:                 "S3Config",
		StorageType:              "StorageType",
	}

	// AWS_Connect_InstanceStorageConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Connect::InstanceStorageConfig.
	AWS_Connect_InstanceStorageConfig__PropertiesSlice = []string{
		AWS_Connect_InstanceStorageConfig__PropertiesMap.InstanceArn,
		AWS_Connect_InstanceStorageConfig__PropertiesMap.KinesisFirehoseConfig,
		AWS_Connect_InstanceStorageConfig__PropertiesMap.KinesisStreamConfig,
		AWS_Connect_InstanceStorageConfig__PropertiesMap.KinesisVideoStreamConfig,
		AWS_Connect_InstanceStorageConfig__PropertiesMap.ResourceType,
		AWS_Connect_InstanceStorageConfig__PropertiesMap.S3Config,
		AWS_Connect_InstanceStorageConfig__PropertiesMap.StorageType,
	}
)

// AWS_Connect_InstanceStorageConfig is a binding for AWS::Connect::InstanceStorageConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html
type AWS_Connect_InstanceStorageConfig struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// InstanceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-instancearn
	InstanceArn cfz.Expression[string] `json:"InstanceArn,omitempty"`

	// KinesisFirehoseConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-kinesisfirehoseconfig
	KinesisFirehoseConfig cfz.Expression[AWS_Connect_InstanceStorageConfig_KinesisFirehoseConfig] `json:"KinesisFirehoseConfig,omitempty"`

	// KinesisStreamConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-kinesisstreamconfig
	KinesisStreamConfig cfz.Expression[AWS_Connect_InstanceStorageConfig_KinesisStreamConfig] `json:"KinesisStreamConfig,omitempty"`

	// KinesisVideoStreamConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-kinesisvideostreamconfig
	KinesisVideoStreamConfig cfz.Expression[AWS_Connect_InstanceStorageConfig_KinesisVideoStreamConfig] `json:"KinesisVideoStreamConfig,omitempty"`

	// ResourceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-resourcetype
	ResourceType cfz.Expression[string] `json:"ResourceType,omitempty"`

	// S3Config is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-s3config
	S3Config cfz.Expression[AWS_Connect_InstanceStorageConfig_S3Config] `json:"S3Config,omitempty"`

	// StorageType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-storagetype
	StorageType cfz.Expression[string] `json:"StorageType,omitempty"`
}

// New__AWS_Connect_InstanceStorageConfig initializes a new *AWS_Connect_InstanceStorageConfig.
func New__AWS_Connect_InstanceStorageConfig(logicalName string) *AWS_Connect_InstanceStorageConfig {
	return &AWS_Connect_InstanceStorageConfig{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Connect_InstanceStorageConfig) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Connect_InstanceStorageConfig) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Connect_InstanceStorageConfig) GetType() string {
	return AWS_Connect_InstanceStorageConfig__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Connect_InstanceStorageConfig) Set__LogicalName(v string) *AWS_Connect_InstanceStorageConfig {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Connect_InstanceStorageConfig) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Connect_InstanceStorageConfig {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Connect_InstanceStorageConfig) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Connect_InstanceStorageConfig {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Connect_InstanceStorageConfig) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Connect_InstanceStorageConfig {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Connect_InstanceStorageConfig) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Connect_InstanceStorageConfig {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Connect_InstanceStorageConfig) Set__RequestedOutputs(v []cfz.Output) *AWS_Connect_InstanceStorageConfig {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Connect_InstanceStorageConfig) Add__RequestedOutputs(v ...cfz.Output) *AWS_Connect_InstanceStorageConfig {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__InstanceArn updates property "InstanceArn".
func (t *AWS_Connect_InstanceStorageConfig) Set__InstanceArn(v cfz.Expression[string]) *AWS_Connect_InstanceStorageConfig {
	t.InstanceArn = v
	return t
}

// SetV__InstanceArn updates property "InstanceArn".
func (t *AWS_Connect_InstanceStorageConfig) SetV__InstanceArn(v string) *AWS_Connect_InstanceStorageConfig {
	t.InstanceArn = cfz.V(v)
	return t
}

// Set__KinesisFirehoseConfig updates property "KinesisFirehoseConfig".
func (t *AWS_Connect_InstanceStorageConfig) Set__KinesisFirehoseConfig(v cfz.Expression[AWS_Connect_InstanceStorageConfig_KinesisFirehoseConfig]) *AWS_Connect_InstanceStorageConfig {
	t.KinesisFirehoseConfig = v
	return t
}

// SetV__KinesisFirehoseConfig updates property "KinesisFirehoseConfig".
func (t *AWS_Connect_InstanceStorageConfig) SetV__KinesisFirehoseConfig(v AWS_Connect_InstanceStorageConfig_KinesisFirehoseConfig) *AWS_Connect_InstanceStorageConfig {
	t.KinesisFirehoseConfig = cfz.V(v)
	return t
}

// Set__KinesisStreamConfig updates property "KinesisStreamConfig".
func (t *AWS_Connect_InstanceStorageConfig) Set__KinesisStreamConfig(v cfz.Expression[AWS_Connect_InstanceStorageConfig_KinesisStreamConfig]) *AWS_Connect_InstanceStorageConfig {
	t.KinesisStreamConfig = v
	return t
}

// SetV__KinesisStreamConfig updates property "KinesisStreamConfig".
func (t *AWS_Connect_InstanceStorageConfig) SetV__KinesisStreamConfig(v AWS_Connect_InstanceStorageConfig_KinesisStreamConfig) *AWS_Connect_InstanceStorageConfig {
	t.KinesisStreamConfig = cfz.V(v)
	return t
}

// Set__KinesisVideoStreamConfig updates property "KinesisVideoStreamConfig".
func (t *AWS_Connect_InstanceStorageConfig) Set__KinesisVideoStreamConfig(v cfz.Expression[AWS_Connect_InstanceStorageConfig_KinesisVideoStreamConfig]) *AWS_Connect_InstanceStorageConfig {
	t.KinesisVideoStreamConfig = v
	return t
}

// SetV__KinesisVideoStreamConfig updates property "KinesisVideoStreamConfig".
func (t *AWS_Connect_InstanceStorageConfig) SetV__KinesisVideoStreamConfig(v AWS_Connect_InstanceStorageConfig_KinesisVideoStreamConfig) *AWS_Connect_InstanceStorageConfig {
	t.KinesisVideoStreamConfig = cfz.V(v)
	return t
}

// Set__ResourceType updates property "ResourceType".
func (t *AWS_Connect_InstanceStorageConfig) Set__ResourceType(v cfz.Expression[string]) *AWS_Connect_InstanceStorageConfig {
	t.ResourceType = v
	return t
}

// SetV__ResourceType updates property "ResourceType".
func (t *AWS_Connect_InstanceStorageConfig) SetV__ResourceType(v string) *AWS_Connect_InstanceStorageConfig {
	t.ResourceType = cfz.V(v)
	return t
}

// Set__S3Config updates property "S3Config".
func (t *AWS_Connect_InstanceStorageConfig) Set__S3Config(v cfz.Expression[AWS_Connect_InstanceStorageConfig_S3Config]) *AWS_Connect_InstanceStorageConfig {
	t.S3Config = v
	return t
}

// SetV__S3Config updates property "S3Config".
func (t *AWS_Connect_InstanceStorageConfig) SetV__S3Config(v AWS_Connect_InstanceStorageConfig_S3Config) *AWS_Connect_InstanceStorageConfig {
	t.S3Config = cfz.V(v)
	return t
}

// Set__StorageType updates property "StorageType".
func (t *AWS_Connect_InstanceStorageConfig) Set__StorageType(v cfz.Expression[string]) *AWS_Connect_InstanceStorageConfig {
	t.StorageType = v
	return t
}

// SetV__StorageType updates property "StorageType".
func (t *AWS_Connect_InstanceStorageConfig) SetV__StorageType(v string) *AWS_Connect_InstanceStorageConfig {
	t.StorageType = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Connect_InstanceStorageConfig) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__AssociationId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AssociationId
func (t *AWS_Connect_InstanceStorageConfig) GetAtt__AssociationId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Connect_InstanceStorageConfig__AttributesMap.AssociationId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Connect_InstanceStorageConfig) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AssociationId returns a conventionally configured output for an attribute of this resource.
// Attribute: AssociationId
func (t *AWS_Connect_InstanceStorageConfig) GetConventionalOutputAtt__AssociationId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAssociationId", t.GetAtt__AssociationId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Connect_InstanceStorageConfig) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Connect_InstanceStorageConfig

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

func (t *AWS_Connect_InstanceStorageConfig) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
