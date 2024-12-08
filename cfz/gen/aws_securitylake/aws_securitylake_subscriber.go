// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_securitylake

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_SecurityLake_Subscriber)(nil)
	_ cfz.Resource                   = (*AWS_SecurityLake_Subscriber)(nil)
)

const (
	// AWS_SecurityLake_Subscriber__Type is the CloudFormation type for AWS::SecurityLake::Subscriber.
	AWS_SecurityLake_Subscriber__Type = "AWS::SecurityLake::Subscriber"
)

var (
	// AWS_SecurityLake_Subscriber__AttributesMap reports all the CloudFormation attributes for AWS::SecurityLake::Subscriber.
	AWS_SecurityLake_Subscriber__AttributesMap = struct {
		ResourceShareArn  string
		ResourceShareName string
		S3BucketArn       string
		SubscriberArn     string
		SubscriberRoleArn string
	}{
		ResourceShareArn:  "ResourceShareArn",
		ResourceShareName: "ResourceShareName",
		S3BucketArn:       "S3BucketArn",
		SubscriberArn:     "SubscriberArn",
		SubscriberRoleArn: "SubscriberRoleArn",
	}

	// AWS_SecurityLake_Subscriber__AttributesSlice reports all the CloudFormation attributes for AWS::SecurityLake::Subscriber.
	AWS_SecurityLake_Subscriber__AttributesSlice = []string{
		AWS_SecurityLake_Subscriber__AttributesMap.ResourceShareArn,
		AWS_SecurityLake_Subscriber__AttributesMap.ResourceShareName,
		AWS_SecurityLake_Subscriber__AttributesMap.S3BucketArn,
		AWS_SecurityLake_Subscriber__AttributesMap.SubscriberArn,
		AWS_SecurityLake_Subscriber__AttributesMap.SubscriberRoleArn,
	}
)

var (
	// AWS_SecurityLake_Subscriber__PropertiesMap reports all the CloudFormation properties for AWS::SecurityLake::Subscriber.
	AWS_SecurityLake_Subscriber__PropertiesMap = struct {
		AccessTypes           string
		DataLakeArn           string
		Sources               string
		SubscriberDescription string
		SubscriberIdentity    string
		SubscriberName        string
		Tags                  string
	}{
		AccessTypes:           "AccessTypes",
		DataLakeArn:           "DataLakeArn",
		Sources:               "Sources",
		SubscriberDescription: "SubscriberDescription",
		SubscriberIdentity:    "SubscriberIdentity",
		SubscriberName:        "SubscriberName",
		Tags:                  "Tags",
	}

	// AWS_SecurityLake_Subscriber__PropertiesSlice reports all the CloudFormation properties for AWS::SecurityLake::Subscriber.
	AWS_SecurityLake_Subscriber__PropertiesSlice = []string{
		AWS_SecurityLake_Subscriber__PropertiesMap.AccessTypes,
		AWS_SecurityLake_Subscriber__PropertiesMap.DataLakeArn,
		AWS_SecurityLake_Subscriber__PropertiesMap.Sources,
		AWS_SecurityLake_Subscriber__PropertiesMap.SubscriberDescription,
		AWS_SecurityLake_Subscriber__PropertiesMap.SubscriberIdentity,
		AWS_SecurityLake_Subscriber__PropertiesMap.SubscriberName,
		AWS_SecurityLake_Subscriber__PropertiesMap.Tags,
	}
)

// AWS_SecurityLake_Subscriber is a binding for AWS::SecurityLake::Subscriber.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html
type AWS_SecurityLake_Subscriber struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AccessTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-accesstypes
	AccessTypes cfz.ExpressionSlice[string] `json:"AccessTypes,omitempty"`

	// DataLakeArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-datalakearn
	DataLakeArn cfz.Expression[string] `json:"DataLakeArn,omitempty"`

	// Sources is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-sources
	Sources cfz.ExpressionSlice[AWS_SecurityLake_Subscriber_Source] `json:"Sources,omitempty"`

	// SubscriberDescription is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-subscriberdescription
	SubscriberDescription cfz.Expression[string] `json:"SubscriberDescription,omitempty"`

	// SubscriberIdentity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-subscriberidentity
	SubscriberIdentity cfz.Expression[AWS_SecurityLake_Subscriber_SubscriberIdentity] `json:"SubscriberIdentity,omitempty"`

	// SubscriberName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-subscribername
	SubscriberName cfz.Expression[string] `json:"SubscriberName,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscriber.html#cfn-securitylake-subscriber-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_SecurityLake_Subscriber initializes a new *AWS_SecurityLake_Subscriber.
func New__AWS_SecurityLake_Subscriber(logicalName string) *AWS_SecurityLake_Subscriber {
	return &AWS_SecurityLake_Subscriber{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_SecurityLake_Subscriber) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_SecurityLake_Subscriber) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_SecurityLake_Subscriber) GetType() string {
	return AWS_SecurityLake_Subscriber__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_SecurityLake_Subscriber) Set__LogicalName(v string) *AWS_SecurityLake_Subscriber {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_SecurityLake_Subscriber) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_SecurityLake_Subscriber {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_SecurityLake_Subscriber) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_SecurityLake_Subscriber {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_SecurityLake_Subscriber) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_SecurityLake_Subscriber {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_SecurityLake_Subscriber) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_SecurityLake_Subscriber {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_SecurityLake_Subscriber) Set__RequestedOutputs(v []cfz.Output) *AWS_SecurityLake_Subscriber {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_SecurityLake_Subscriber) Add__RequestedOutputs(v ...cfz.Output) *AWS_SecurityLake_Subscriber {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AccessTypes updates property "AccessTypes".
func (t *AWS_SecurityLake_Subscriber) Set__AccessTypes(v cfz.ExpressionSlice[string]) *AWS_SecurityLake_Subscriber {
	t.AccessTypes = v
	return t
}

// SetS__AccessTypes updates property "AccessTypes".
func (t *AWS_SecurityLake_Subscriber) SetS__AccessTypes(v ...cfz.Expression[string]) *AWS_SecurityLake_Subscriber {
	t.AccessTypes = cfz.S(v...)
	return t
}

// SetSV__AccessTypes updates property "AccessTypes".
func (t *AWS_SecurityLake_Subscriber) SetSV__AccessTypes(v ...string) *AWS_SecurityLake_Subscriber {
	t.AccessTypes = cfz.SV(v...)
	return t
}

// Set__DataLakeArn updates property "DataLakeArn".
func (t *AWS_SecurityLake_Subscriber) Set__DataLakeArn(v cfz.Expression[string]) *AWS_SecurityLake_Subscriber {
	t.DataLakeArn = v
	return t
}

// SetV__DataLakeArn updates property "DataLakeArn".
func (t *AWS_SecurityLake_Subscriber) SetV__DataLakeArn(v string) *AWS_SecurityLake_Subscriber {
	t.DataLakeArn = cfz.V(v)
	return t
}

// Set__Sources updates property "Sources".
func (t *AWS_SecurityLake_Subscriber) Set__Sources(v cfz.ExpressionSlice[AWS_SecurityLake_Subscriber_Source]) *AWS_SecurityLake_Subscriber {
	t.Sources = v
	return t
}

// SetS__Sources updates property "Sources".
func (t *AWS_SecurityLake_Subscriber) SetS__Sources(v ...cfz.Expression[AWS_SecurityLake_Subscriber_Source]) *AWS_SecurityLake_Subscriber {
	t.Sources = cfz.S(v...)
	return t
}

// SetSV__Sources updates property "Sources".
func (t *AWS_SecurityLake_Subscriber) SetSV__Sources(v ...AWS_SecurityLake_Subscriber_Source) *AWS_SecurityLake_Subscriber {
	t.Sources = cfz.SV(v...)
	return t
}

// Set__SubscriberDescription updates property "SubscriberDescription".
func (t *AWS_SecurityLake_Subscriber) Set__SubscriberDescription(v cfz.Expression[string]) *AWS_SecurityLake_Subscriber {
	t.SubscriberDescription = v
	return t
}

// SetV__SubscriberDescription updates property "SubscriberDescription".
func (t *AWS_SecurityLake_Subscriber) SetV__SubscriberDescription(v string) *AWS_SecurityLake_Subscriber {
	t.SubscriberDescription = cfz.V(v)
	return t
}

// Set__SubscriberIdentity updates property "SubscriberIdentity".
func (t *AWS_SecurityLake_Subscriber) Set__SubscriberIdentity(v cfz.Expression[AWS_SecurityLake_Subscriber_SubscriberIdentity]) *AWS_SecurityLake_Subscriber {
	t.SubscriberIdentity = v
	return t
}

// SetV__SubscriberIdentity updates property "SubscriberIdentity".
func (t *AWS_SecurityLake_Subscriber) SetV__SubscriberIdentity(v AWS_SecurityLake_Subscriber_SubscriberIdentity) *AWS_SecurityLake_Subscriber {
	t.SubscriberIdentity = cfz.V(v)
	return t
}

// Set__SubscriberName updates property "SubscriberName".
func (t *AWS_SecurityLake_Subscriber) Set__SubscriberName(v cfz.Expression[string]) *AWS_SecurityLake_Subscriber {
	t.SubscriberName = v
	return t
}

// SetV__SubscriberName updates property "SubscriberName".
func (t *AWS_SecurityLake_Subscriber) SetV__SubscriberName(v string) *AWS_SecurityLake_Subscriber {
	t.SubscriberName = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_SecurityLake_Subscriber) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_SecurityLake_Subscriber {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_SecurityLake_Subscriber) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_SecurityLake_Subscriber {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_SecurityLake_Subscriber) SetSV__Tags(v ...cfz.Tag) *AWS_SecurityLake_Subscriber {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_SecurityLake_Subscriber) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__ResourceShareArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ResourceShareArn
func (t *AWS_SecurityLake_Subscriber) GetAtt__ResourceShareArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SecurityLake_Subscriber__AttributesMap.ResourceShareArn))
}

// GetAtt__ResourceShareName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ResourceShareName
func (t *AWS_SecurityLake_Subscriber) GetAtt__ResourceShareName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SecurityLake_Subscriber__AttributesMap.ResourceShareName))
}

// GetAtt__S3BucketArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: S3BucketArn
func (t *AWS_SecurityLake_Subscriber) GetAtt__S3BucketArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SecurityLake_Subscriber__AttributesMap.S3BucketArn))
}

// GetAtt__SubscriberArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: SubscriberArn
func (t *AWS_SecurityLake_Subscriber) GetAtt__SubscriberArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SecurityLake_Subscriber__AttributesMap.SubscriberArn))
}

// GetAtt__SubscriberRoleArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: SubscriberRoleArn
func (t *AWS_SecurityLake_Subscriber) GetAtt__SubscriberRoleArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SecurityLake_Subscriber__AttributesMap.SubscriberRoleArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_SecurityLake_Subscriber) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ResourceShareArn returns a conventionally configured output for an attribute of this resource.
// Attribute: ResourceShareArn
func (t *AWS_SecurityLake_Subscriber) GetConventionalOutputAtt__ResourceShareArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttResourceShareArn", t.GetAtt__ResourceShareArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ResourceShareName returns a conventionally configured output for an attribute of this resource.
// Attribute: ResourceShareName
func (t *AWS_SecurityLake_Subscriber) GetConventionalOutputAtt__ResourceShareName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttResourceShareName", t.GetAtt__ResourceShareName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__S3BucketArn returns a conventionally configured output for an attribute of this resource.
// Attribute: S3BucketArn
func (t *AWS_SecurityLake_Subscriber) GetConventionalOutputAtt__S3BucketArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttS3BucketArn", t.GetAtt__S3BucketArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__SubscriberArn returns a conventionally configured output for an attribute of this resource.
// Attribute: SubscriberArn
func (t *AWS_SecurityLake_Subscriber) GetConventionalOutputAtt__SubscriberArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttSubscriberArn", t.GetAtt__SubscriberArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__SubscriberRoleArn returns a conventionally configured output for an attribute of this resource.
// Attribute: SubscriberRoleArn
func (t *AWS_SecurityLake_Subscriber) GetConventionalOutputAtt__SubscriberRoleArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttSubscriberRoleArn", t.GetAtt__SubscriberRoleArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_SecurityLake_Subscriber) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_SecurityLake_Subscriber

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

func (t *AWS_SecurityLake_Subscriber) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
