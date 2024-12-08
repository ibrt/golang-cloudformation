// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datasync

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_DataSync_LocationAzureBlob)(nil)
	_ cfz.Resource                   = (*AWS_DataSync_LocationAzureBlob)(nil)
)

const (
	// AWS_DataSync_LocationAzureBlob__Type is the CloudFormation type for AWS::DataSync::LocationAzureBlob.
	AWS_DataSync_LocationAzureBlob__Type = "AWS::DataSync::LocationAzureBlob"
)

var (
	// AWS_DataSync_LocationAzureBlob__AttributesMap reports all the CloudFormation attributes for AWS::DataSync::LocationAzureBlob.
	AWS_DataSync_LocationAzureBlob__AttributesMap = struct {
		LocationArn string
		LocationUri string
	}{
		LocationArn: "LocationArn",
		LocationUri: "LocationUri",
	}

	// AWS_DataSync_LocationAzureBlob__AttributesSlice reports all the CloudFormation attributes for AWS::DataSync::LocationAzureBlob.
	AWS_DataSync_LocationAzureBlob__AttributesSlice = []string{
		AWS_DataSync_LocationAzureBlob__AttributesMap.LocationArn,
		AWS_DataSync_LocationAzureBlob__AttributesMap.LocationUri,
	}
)

var (
	// AWS_DataSync_LocationAzureBlob__PropertiesMap reports all the CloudFormation properties for AWS::DataSync::LocationAzureBlob.
	AWS_DataSync_LocationAzureBlob__PropertiesMap = struct {
		AgentArns                   string
		AzureAccessTier             string
		AzureBlobAuthenticationType string
		AzureBlobContainerUrl       string
		AzureBlobSasConfiguration   string
		AzureBlobType               string
		Subdirectory                string
		Tags                        string
	}{
		AgentArns:                   "AgentArns",
		AzureAccessTier:             "AzureAccessTier",
		AzureBlobAuthenticationType: "AzureBlobAuthenticationType",
		AzureBlobContainerUrl:       "AzureBlobContainerUrl",
		AzureBlobSasConfiguration:   "AzureBlobSasConfiguration",
		AzureBlobType:               "AzureBlobType",
		Subdirectory:                "Subdirectory",
		Tags:                        "Tags",
	}

	// AWS_DataSync_LocationAzureBlob__PropertiesSlice reports all the CloudFormation properties for AWS::DataSync::LocationAzureBlob.
	AWS_DataSync_LocationAzureBlob__PropertiesSlice = []string{
		AWS_DataSync_LocationAzureBlob__PropertiesMap.AgentArns,
		AWS_DataSync_LocationAzureBlob__PropertiesMap.AzureAccessTier,
		AWS_DataSync_LocationAzureBlob__PropertiesMap.AzureBlobAuthenticationType,
		AWS_DataSync_LocationAzureBlob__PropertiesMap.AzureBlobContainerUrl,
		AWS_DataSync_LocationAzureBlob__PropertiesMap.AzureBlobSasConfiguration,
		AWS_DataSync_LocationAzureBlob__PropertiesMap.AzureBlobType,
		AWS_DataSync_LocationAzureBlob__PropertiesMap.Subdirectory,
		AWS_DataSync_LocationAzureBlob__PropertiesMap.Tags,
	}
)

// AWS_DataSync_LocationAzureBlob is a binding for AWS::DataSync::LocationAzureBlob.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html
type AWS_DataSync_LocationAzureBlob struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AgentArns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-agentarns
	AgentArns cfz.ExpressionSlice[string] `json:"AgentArns,omitempty"`

	// AzureAccessTier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureaccesstier
	AzureAccessTier cfz.Expression[string] `json:"AzureAccessTier,omitempty"`

	// AzureBlobAuthenticationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobauthenticationtype
	AzureBlobAuthenticationType cfz.Expression[string] `json:"AzureBlobAuthenticationType,omitempty"`

	// AzureBlobContainerUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobcontainerurl
	AzureBlobContainerUrl cfz.Expression[string] `json:"AzureBlobContainerUrl,omitempty"`

	// AzureBlobSasConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobsasconfiguration
	AzureBlobSasConfiguration cfz.Expression[AWS_DataSync_LocationAzureBlob_AzureBlobSasConfiguration] `json:"AzureBlobSasConfiguration,omitempty"`

	// AzureBlobType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-azureblobtype
	AzureBlobType cfz.Expression[string] `json:"AzureBlobType,omitempty"`

	// Subdirectory is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-subdirectory
	Subdirectory cfz.Expression[string] `json:"Subdirectory,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationazureblob.html#cfn-datasync-locationazureblob-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_DataSync_LocationAzureBlob initializes a new *AWS_DataSync_LocationAzureBlob.
func New__AWS_DataSync_LocationAzureBlob(logicalName string) *AWS_DataSync_LocationAzureBlob {
	return &AWS_DataSync_LocationAzureBlob{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_DataSync_LocationAzureBlob) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_DataSync_LocationAzureBlob) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_DataSync_LocationAzureBlob) GetType() string {
	return AWS_DataSync_LocationAzureBlob__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_DataSync_LocationAzureBlob) Set__LogicalName(v string) *AWS_DataSync_LocationAzureBlob {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_DataSync_LocationAzureBlob) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_DataSync_LocationAzureBlob {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_DataSync_LocationAzureBlob) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_DataSync_LocationAzureBlob {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_DataSync_LocationAzureBlob) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_DataSync_LocationAzureBlob {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_DataSync_LocationAzureBlob) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_DataSync_LocationAzureBlob {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_DataSync_LocationAzureBlob) Set__RequestedOutputs(v []cfz.Output) *AWS_DataSync_LocationAzureBlob {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_DataSync_LocationAzureBlob) Add__RequestedOutputs(v ...cfz.Output) *AWS_DataSync_LocationAzureBlob {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AgentArns updates property "AgentArns".
func (t *AWS_DataSync_LocationAzureBlob) Set__AgentArns(v cfz.ExpressionSlice[string]) *AWS_DataSync_LocationAzureBlob {
	t.AgentArns = v
	return t
}

// SetS__AgentArns updates property "AgentArns".
func (t *AWS_DataSync_LocationAzureBlob) SetS__AgentArns(v ...cfz.Expression[string]) *AWS_DataSync_LocationAzureBlob {
	t.AgentArns = cfz.S(v...)
	return t
}

// SetSV__AgentArns updates property "AgentArns".
func (t *AWS_DataSync_LocationAzureBlob) SetSV__AgentArns(v ...string) *AWS_DataSync_LocationAzureBlob {
	t.AgentArns = cfz.SV(v...)
	return t
}

// Set__AzureAccessTier updates property "AzureAccessTier".
func (t *AWS_DataSync_LocationAzureBlob) Set__AzureAccessTier(v cfz.Expression[string]) *AWS_DataSync_LocationAzureBlob {
	t.AzureAccessTier = v
	return t
}

// SetV__AzureAccessTier updates property "AzureAccessTier".
func (t *AWS_DataSync_LocationAzureBlob) SetV__AzureAccessTier(v string) *AWS_DataSync_LocationAzureBlob {
	t.AzureAccessTier = cfz.V(v)
	return t
}

// Set__AzureBlobAuthenticationType updates property "AzureBlobAuthenticationType".
func (t *AWS_DataSync_LocationAzureBlob) Set__AzureBlobAuthenticationType(v cfz.Expression[string]) *AWS_DataSync_LocationAzureBlob {
	t.AzureBlobAuthenticationType = v
	return t
}

// SetV__AzureBlobAuthenticationType updates property "AzureBlobAuthenticationType".
func (t *AWS_DataSync_LocationAzureBlob) SetV__AzureBlobAuthenticationType(v string) *AWS_DataSync_LocationAzureBlob {
	t.AzureBlobAuthenticationType = cfz.V(v)
	return t
}

// Set__AzureBlobContainerUrl updates property "AzureBlobContainerUrl".
func (t *AWS_DataSync_LocationAzureBlob) Set__AzureBlobContainerUrl(v cfz.Expression[string]) *AWS_DataSync_LocationAzureBlob {
	t.AzureBlobContainerUrl = v
	return t
}

// SetV__AzureBlobContainerUrl updates property "AzureBlobContainerUrl".
func (t *AWS_DataSync_LocationAzureBlob) SetV__AzureBlobContainerUrl(v string) *AWS_DataSync_LocationAzureBlob {
	t.AzureBlobContainerUrl = cfz.V(v)
	return t
}

// Set__AzureBlobSasConfiguration updates property "AzureBlobSasConfiguration".
func (t *AWS_DataSync_LocationAzureBlob) Set__AzureBlobSasConfiguration(v cfz.Expression[AWS_DataSync_LocationAzureBlob_AzureBlobSasConfiguration]) *AWS_DataSync_LocationAzureBlob {
	t.AzureBlobSasConfiguration = v
	return t
}

// SetV__AzureBlobSasConfiguration updates property "AzureBlobSasConfiguration".
func (t *AWS_DataSync_LocationAzureBlob) SetV__AzureBlobSasConfiguration(v AWS_DataSync_LocationAzureBlob_AzureBlobSasConfiguration) *AWS_DataSync_LocationAzureBlob {
	t.AzureBlobSasConfiguration = cfz.V(v)
	return t
}

// Set__AzureBlobType updates property "AzureBlobType".
func (t *AWS_DataSync_LocationAzureBlob) Set__AzureBlobType(v cfz.Expression[string]) *AWS_DataSync_LocationAzureBlob {
	t.AzureBlobType = v
	return t
}

// SetV__AzureBlobType updates property "AzureBlobType".
func (t *AWS_DataSync_LocationAzureBlob) SetV__AzureBlobType(v string) *AWS_DataSync_LocationAzureBlob {
	t.AzureBlobType = cfz.V(v)
	return t
}

// Set__Subdirectory updates property "Subdirectory".
func (t *AWS_DataSync_LocationAzureBlob) Set__Subdirectory(v cfz.Expression[string]) *AWS_DataSync_LocationAzureBlob {
	t.Subdirectory = v
	return t
}

// SetV__Subdirectory updates property "Subdirectory".
func (t *AWS_DataSync_LocationAzureBlob) SetV__Subdirectory(v string) *AWS_DataSync_LocationAzureBlob {
	t.Subdirectory = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_DataSync_LocationAzureBlob) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_DataSync_LocationAzureBlob {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_DataSync_LocationAzureBlob) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_DataSync_LocationAzureBlob {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_DataSync_LocationAzureBlob) SetSV__Tags(v ...cfz.Tag) *AWS_DataSync_LocationAzureBlob {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_DataSync_LocationAzureBlob) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__LocationArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LocationArn
func (t *AWS_DataSync_LocationAzureBlob) GetAtt__LocationArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataSync_LocationAzureBlob__AttributesMap.LocationArn))
}

// GetAtt__LocationUri returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LocationUri
func (t *AWS_DataSync_LocationAzureBlob) GetAtt__LocationUri() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataSync_LocationAzureBlob__AttributesMap.LocationUri))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_DataSync_LocationAzureBlob) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LocationArn returns a conventionally configured output for an attribute of this resource.
// Attribute: LocationArn
func (t *AWS_DataSync_LocationAzureBlob) GetConventionalOutputAtt__LocationArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLocationArn", t.GetAtt__LocationArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LocationUri returns a conventionally configured output for an attribute of this resource.
// Attribute: LocationUri
func (t *AWS_DataSync_LocationAzureBlob) GetConventionalOutputAtt__LocationUri(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLocationUri", t.GetAtt__LocationUri())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_DataSync_LocationAzureBlob) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_DataSync_LocationAzureBlob

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

func (t *AWS_DataSync_LocationAzureBlob) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
