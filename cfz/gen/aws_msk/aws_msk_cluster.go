// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_msk

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_MSK_Cluster)(nil)
	_ cfz.Resource                   = (*AWS_MSK_Cluster)(nil)
)

const (
	// AWS_MSK_Cluster__Type is the CloudFormation type for AWS::MSK::Cluster.
	AWS_MSK_Cluster__Type = "AWS::MSK::Cluster"
)

var (
	// AWS_MSK_Cluster__AttributesMap reports all the CloudFormation attributes for AWS::MSK::Cluster.
	AWS_MSK_Cluster__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_MSK_Cluster__AttributesSlice reports all the CloudFormation attributes for AWS::MSK::Cluster.
	AWS_MSK_Cluster__AttributesSlice = []string{
		AWS_MSK_Cluster__AttributesMap.Arn,
	}
)

var (
	// AWS_MSK_Cluster__PropertiesMap reports all the CloudFormation properties for AWS::MSK::Cluster.
	AWS_MSK_Cluster__PropertiesMap = struct {
		BrokerNodeGroupInfo  string
		ClientAuthentication string
		ClusterName          string
		ConfigurationInfo    string
		CurrentVersion       string
		EncryptionInfo       string
		EnhancedMonitoring   string
		KafkaVersion         string
		LoggingInfo          string
		NumberOfBrokerNodes  string
		OpenMonitoring       string
		StorageMode          string
		Tags                 string
	}{
		BrokerNodeGroupInfo:  "BrokerNodeGroupInfo",
		ClientAuthentication: "ClientAuthentication",
		ClusterName:          "ClusterName",
		ConfigurationInfo:    "ConfigurationInfo",
		CurrentVersion:       "CurrentVersion",
		EncryptionInfo:       "EncryptionInfo",
		EnhancedMonitoring:   "EnhancedMonitoring",
		KafkaVersion:         "KafkaVersion",
		LoggingInfo:          "LoggingInfo",
		NumberOfBrokerNodes:  "NumberOfBrokerNodes",
		OpenMonitoring:       "OpenMonitoring",
		StorageMode:          "StorageMode",
		Tags:                 "Tags",
	}

	// AWS_MSK_Cluster__PropertiesSlice reports all the CloudFormation properties for AWS::MSK::Cluster.
	AWS_MSK_Cluster__PropertiesSlice = []string{
		AWS_MSK_Cluster__PropertiesMap.BrokerNodeGroupInfo,
		AWS_MSK_Cluster__PropertiesMap.ClientAuthentication,
		AWS_MSK_Cluster__PropertiesMap.ClusterName,
		AWS_MSK_Cluster__PropertiesMap.ConfigurationInfo,
		AWS_MSK_Cluster__PropertiesMap.CurrentVersion,
		AWS_MSK_Cluster__PropertiesMap.EncryptionInfo,
		AWS_MSK_Cluster__PropertiesMap.EnhancedMonitoring,
		AWS_MSK_Cluster__PropertiesMap.KafkaVersion,
		AWS_MSK_Cluster__PropertiesMap.LoggingInfo,
		AWS_MSK_Cluster__PropertiesMap.NumberOfBrokerNodes,
		AWS_MSK_Cluster__PropertiesMap.OpenMonitoring,
		AWS_MSK_Cluster__PropertiesMap.StorageMode,
		AWS_MSK_Cluster__PropertiesMap.Tags,
	}
)

// AWS_MSK_Cluster is a binding for AWS::MSK::Cluster.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html
type AWS_MSK_Cluster struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// BrokerNodeGroupInfo is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-brokernodegroupinfo
	BrokerNodeGroupInfo cfz.Expression[AWS_MSK_Cluster_BrokerNodeGroupInfo] `json:"BrokerNodeGroupInfo,omitempty"`

	// ClientAuthentication is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-clientauthentication
	ClientAuthentication cfz.Expression[AWS_MSK_Cluster_ClientAuthentication] `json:"ClientAuthentication,omitempty"`

	// ClusterName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-clustername
	ClusterName cfz.Expression[string] `json:"ClusterName,omitempty"`

	// ConfigurationInfo is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-configurationinfo
	ConfigurationInfo cfz.Expression[AWS_MSK_Cluster_ConfigurationInfo] `json:"ConfigurationInfo,omitempty"`

	// CurrentVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-currentversion
	CurrentVersion cfz.Expression[string] `json:"CurrentVersion,omitempty"`

	// EncryptionInfo is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-encryptioninfo
	EncryptionInfo cfz.Expression[AWS_MSK_Cluster_EncryptionInfo] `json:"EncryptionInfo,omitempty"`

	// EnhancedMonitoring is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-enhancedmonitoring
	EnhancedMonitoring cfz.Expression[string] `json:"EnhancedMonitoring,omitempty"`

	// KafkaVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-kafkaversion
	KafkaVersion cfz.Expression[string] `json:"KafkaVersion,omitempty"`

	// LoggingInfo is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-logginginfo
	LoggingInfo cfz.Expression[AWS_MSK_Cluster_LoggingInfo] `json:"LoggingInfo,omitempty"`

	// NumberOfBrokerNodes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-numberofbrokernodes
	NumberOfBrokerNodes cfz.Expression[int32] `json:"NumberOfBrokerNodes,omitempty"`

	// OpenMonitoring is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-openmonitoring
	OpenMonitoring cfz.Expression[AWS_MSK_Cluster_OpenMonitoring] `json:"OpenMonitoring,omitempty"`

	// StorageMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-storagemode
	StorageMode cfz.Expression[string] `json:"StorageMode,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-cluster.html#cfn-msk-cluster-tags
	Tags cfz.ExpressionMap[string] `json:"Tags,omitempty"`
}

// New__AWS_MSK_Cluster initializes a new *AWS_MSK_Cluster.
func New__AWS_MSK_Cluster(logicalName string) *AWS_MSK_Cluster {
	return &AWS_MSK_Cluster{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_MSK_Cluster) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_MSK_Cluster) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_MSK_Cluster) GetType() string {
	return AWS_MSK_Cluster__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_MSK_Cluster) Set__LogicalName(v string) *AWS_MSK_Cluster {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_MSK_Cluster) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_MSK_Cluster {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_MSK_Cluster) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_MSK_Cluster {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_MSK_Cluster) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_MSK_Cluster {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_MSK_Cluster) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_MSK_Cluster {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_MSK_Cluster) Set__RequestedOutputs(v []cfz.Output) *AWS_MSK_Cluster {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_MSK_Cluster) Add__RequestedOutputs(v ...cfz.Output) *AWS_MSK_Cluster {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__BrokerNodeGroupInfo updates property "BrokerNodeGroupInfo".
func (t *AWS_MSK_Cluster) Set__BrokerNodeGroupInfo(v cfz.Expression[AWS_MSK_Cluster_BrokerNodeGroupInfo]) *AWS_MSK_Cluster {
	t.BrokerNodeGroupInfo = v
	return t
}

// SetV__BrokerNodeGroupInfo updates property "BrokerNodeGroupInfo".
func (t *AWS_MSK_Cluster) SetV__BrokerNodeGroupInfo(v AWS_MSK_Cluster_BrokerNodeGroupInfo) *AWS_MSK_Cluster {
	t.BrokerNodeGroupInfo = cfz.V(v)
	return t
}

// Set__ClientAuthentication updates property "ClientAuthentication".
func (t *AWS_MSK_Cluster) Set__ClientAuthentication(v cfz.Expression[AWS_MSK_Cluster_ClientAuthentication]) *AWS_MSK_Cluster {
	t.ClientAuthentication = v
	return t
}

// SetV__ClientAuthentication updates property "ClientAuthentication".
func (t *AWS_MSK_Cluster) SetV__ClientAuthentication(v AWS_MSK_Cluster_ClientAuthentication) *AWS_MSK_Cluster {
	t.ClientAuthentication = cfz.V(v)
	return t
}

// Set__ClusterName updates property "ClusterName".
func (t *AWS_MSK_Cluster) Set__ClusterName(v cfz.Expression[string]) *AWS_MSK_Cluster {
	t.ClusterName = v
	return t
}

// SetV__ClusterName updates property "ClusterName".
func (t *AWS_MSK_Cluster) SetV__ClusterName(v string) *AWS_MSK_Cluster {
	t.ClusterName = cfz.V(v)
	return t
}

// Set__ConfigurationInfo updates property "ConfigurationInfo".
func (t *AWS_MSK_Cluster) Set__ConfigurationInfo(v cfz.Expression[AWS_MSK_Cluster_ConfigurationInfo]) *AWS_MSK_Cluster {
	t.ConfigurationInfo = v
	return t
}

// SetV__ConfigurationInfo updates property "ConfigurationInfo".
func (t *AWS_MSK_Cluster) SetV__ConfigurationInfo(v AWS_MSK_Cluster_ConfigurationInfo) *AWS_MSK_Cluster {
	t.ConfigurationInfo = cfz.V(v)
	return t
}

// Set__CurrentVersion updates property "CurrentVersion".
func (t *AWS_MSK_Cluster) Set__CurrentVersion(v cfz.Expression[string]) *AWS_MSK_Cluster {
	t.CurrentVersion = v
	return t
}

// SetV__CurrentVersion updates property "CurrentVersion".
func (t *AWS_MSK_Cluster) SetV__CurrentVersion(v string) *AWS_MSK_Cluster {
	t.CurrentVersion = cfz.V(v)
	return t
}

// Set__EncryptionInfo updates property "EncryptionInfo".
func (t *AWS_MSK_Cluster) Set__EncryptionInfo(v cfz.Expression[AWS_MSK_Cluster_EncryptionInfo]) *AWS_MSK_Cluster {
	t.EncryptionInfo = v
	return t
}

// SetV__EncryptionInfo updates property "EncryptionInfo".
func (t *AWS_MSK_Cluster) SetV__EncryptionInfo(v AWS_MSK_Cluster_EncryptionInfo) *AWS_MSK_Cluster {
	t.EncryptionInfo = cfz.V(v)
	return t
}

// Set__EnhancedMonitoring updates property "EnhancedMonitoring".
func (t *AWS_MSK_Cluster) Set__EnhancedMonitoring(v cfz.Expression[string]) *AWS_MSK_Cluster {
	t.EnhancedMonitoring = v
	return t
}

// SetV__EnhancedMonitoring updates property "EnhancedMonitoring".
func (t *AWS_MSK_Cluster) SetV__EnhancedMonitoring(v string) *AWS_MSK_Cluster {
	t.EnhancedMonitoring = cfz.V(v)
	return t
}

// Set__KafkaVersion updates property "KafkaVersion".
func (t *AWS_MSK_Cluster) Set__KafkaVersion(v cfz.Expression[string]) *AWS_MSK_Cluster {
	t.KafkaVersion = v
	return t
}

// SetV__KafkaVersion updates property "KafkaVersion".
func (t *AWS_MSK_Cluster) SetV__KafkaVersion(v string) *AWS_MSK_Cluster {
	t.KafkaVersion = cfz.V(v)
	return t
}

// Set__LoggingInfo updates property "LoggingInfo".
func (t *AWS_MSK_Cluster) Set__LoggingInfo(v cfz.Expression[AWS_MSK_Cluster_LoggingInfo]) *AWS_MSK_Cluster {
	t.LoggingInfo = v
	return t
}

// SetV__LoggingInfo updates property "LoggingInfo".
func (t *AWS_MSK_Cluster) SetV__LoggingInfo(v AWS_MSK_Cluster_LoggingInfo) *AWS_MSK_Cluster {
	t.LoggingInfo = cfz.V(v)
	return t
}

// Set__NumberOfBrokerNodes updates property "NumberOfBrokerNodes".
func (t *AWS_MSK_Cluster) Set__NumberOfBrokerNodes(v cfz.Expression[int32]) *AWS_MSK_Cluster {
	t.NumberOfBrokerNodes = v
	return t
}

// SetV__NumberOfBrokerNodes updates property "NumberOfBrokerNodes".
func (t *AWS_MSK_Cluster) SetV__NumberOfBrokerNodes(v int32) *AWS_MSK_Cluster {
	t.NumberOfBrokerNodes = cfz.V(v)
	return t
}

// Set__OpenMonitoring updates property "OpenMonitoring".
func (t *AWS_MSK_Cluster) Set__OpenMonitoring(v cfz.Expression[AWS_MSK_Cluster_OpenMonitoring]) *AWS_MSK_Cluster {
	t.OpenMonitoring = v
	return t
}

// SetV__OpenMonitoring updates property "OpenMonitoring".
func (t *AWS_MSK_Cluster) SetV__OpenMonitoring(v AWS_MSK_Cluster_OpenMonitoring) *AWS_MSK_Cluster {
	t.OpenMonitoring = cfz.V(v)
	return t
}

// Set__StorageMode updates property "StorageMode".
func (t *AWS_MSK_Cluster) Set__StorageMode(v cfz.Expression[string]) *AWS_MSK_Cluster {
	t.StorageMode = v
	return t
}

// SetV__StorageMode updates property "StorageMode".
func (t *AWS_MSK_Cluster) SetV__StorageMode(v string) *AWS_MSK_Cluster {
	t.StorageMode = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_MSK_Cluster) Set__Tags(v cfz.ExpressionMap[string]) *AWS_MSK_Cluster {
	t.Tags = v
	return t
}

// SetM__Tags updates property "Tags".
func (t *AWS_MSK_Cluster) SetM__Tags(v ...map[string]cfz.Expression[string]) *AWS_MSK_Cluster {
	t.Tags = cfz.M(v...)
	return t
}

// SetMV__Tags updates property "Tags".
func (t *AWS_MSK_Cluster) SetMV__Tags(v ...map[string]string) *AWS_MSK_Cluster {
	t.Tags = cfz.MV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_MSK_Cluster) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_MSK_Cluster) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MSK_Cluster__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_MSK_Cluster) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_MSK_Cluster) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_MSK_Cluster) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_MSK_Cluster

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

func (t *AWS_MSK_Cluster) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
