// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_NetworkInsightsPath)(nil)
	_ cfz.Resource                   = (*AWS_EC2_NetworkInsightsPath)(nil)
)

const (
	// AWS_EC2_NetworkInsightsPath__Type is the CloudFormation type for AWS::EC2::NetworkInsightsPath.
	AWS_EC2_NetworkInsightsPath__Type = "AWS::EC2::NetworkInsightsPath"
)

var (
	// AWS_EC2_NetworkInsightsPath__AttributesMap reports all the CloudFormation attributes for AWS::EC2::NetworkInsightsPath.
	AWS_EC2_NetworkInsightsPath__AttributesMap = struct {
		CreatedDate            string
		DestinationArn         string
		NetworkInsightsPathArn string
		NetworkInsightsPathId  string
		SourceArn              string
	}{
		CreatedDate:            "CreatedDate",
		DestinationArn:         "DestinationArn",
		NetworkInsightsPathArn: "NetworkInsightsPathArn",
		NetworkInsightsPathId:  "NetworkInsightsPathId",
		SourceArn:              "SourceArn",
	}

	// AWS_EC2_NetworkInsightsPath__AttributesSlice reports all the CloudFormation attributes for AWS::EC2::NetworkInsightsPath.
	AWS_EC2_NetworkInsightsPath__AttributesSlice = []string{
		AWS_EC2_NetworkInsightsPath__AttributesMap.CreatedDate,
		AWS_EC2_NetworkInsightsPath__AttributesMap.DestinationArn,
		AWS_EC2_NetworkInsightsPath__AttributesMap.NetworkInsightsPathArn,
		AWS_EC2_NetworkInsightsPath__AttributesMap.NetworkInsightsPathId,
		AWS_EC2_NetworkInsightsPath__AttributesMap.SourceArn,
	}
)

var (
	// AWS_EC2_NetworkInsightsPath__PropertiesMap reports all the CloudFormation properties for AWS::EC2::NetworkInsightsPath.
	AWS_EC2_NetworkInsightsPath__PropertiesMap = struct {
		Destination         string
		DestinationIp       string
		DestinationPort     string
		FilterAtDestination string
		FilterAtSource      string
		Protocol            string
		Source              string
		SourceIp            string
		Tags                string
	}{
		Destination:         "Destination",
		DestinationIp:       "DestinationIp",
		DestinationPort:     "DestinationPort",
		FilterAtDestination: "FilterAtDestination",
		FilterAtSource:      "FilterAtSource",
		Protocol:            "Protocol",
		Source:              "Source",
		SourceIp:            "SourceIp",
		Tags:                "Tags",
	}

	// AWS_EC2_NetworkInsightsPath__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::NetworkInsightsPath.
	AWS_EC2_NetworkInsightsPath__PropertiesSlice = []string{
		AWS_EC2_NetworkInsightsPath__PropertiesMap.Destination,
		AWS_EC2_NetworkInsightsPath__PropertiesMap.DestinationIp,
		AWS_EC2_NetworkInsightsPath__PropertiesMap.DestinationPort,
		AWS_EC2_NetworkInsightsPath__PropertiesMap.FilterAtDestination,
		AWS_EC2_NetworkInsightsPath__PropertiesMap.FilterAtSource,
		AWS_EC2_NetworkInsightsPath__PropertiesMap.Protocol,
		AWS_EC2_NetworkInsightsPath__PropertiesMap.Source,
		AWS_EC2_NetworkInsightsPath__PropertiesMap.SourceIp,
		AWS_EC2_NetworkInsightsPath__PropertiesMap.Tags,
	}
)

// AWS_EC2_NetworkInsightsPath is a binding for AWS::EC2::NetworkInsightsPath.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html
type AWS_EC2_NetworkInsightsPath struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Destination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-destination
	Destination cfz.Expression[string] `json:"Destination,omitempty"`

	// DestinationIp is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-destinationip
	DestinationIp cfz.Expression[string] `json:"DestinationIp,omitempty"`

	// DestinationPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-destinationport
	DestinationPort cfz.Expression[int32] `json:"DestinationPort,omitempty"`

	// FilterAtDestination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-filteratdestination
	FilterAtDestination cfz.Expression[AWS_EC2_NetworkInsightsPath_PathFilter] `json:"FilterAtDestination,omitempty"`

	// FilterAtSource is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-filteratsource
	FilterAtSource cfz.Expression[AWS_EC2_NetworkInsightsPath_PathFilter] `json:"FilterAtSource,omitempty"`

	// Protocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-protocol
	Protocol cfz.Expression[string] `json:"Protocol,omitempty"`

	// Source is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-source
	Source cfz.Expression[string] `json:"Source,omitempty"`

	// SourceIp is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-sourceip
	SourceIp cfz.Expression[string] `json:"SourceIp,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_EC2_NetworkInsightsPath initializes a new *AWS_EC2_NetworkInsightsPath.
func New__AWS_EC2_NetworkInsightsPath(logicalName string) *AWS_EC2_NetworkInsightsPath {
	return &AWS_EC2_NetworkInsightsPath{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_NetworkInsightsPath) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_NetworkInsightsPath) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_NetworkInsightsPath) GetType() string {
	return AWS_EC2_NetworkInsightsPath__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_NetworkInsightsPath) Set__LogicalName(v string) *AWS_EC2_NetworkInsightsPath {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_NetworkInsightsPath) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_NetworkInsightsPath {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_NetworkInsightsPath) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_NetworkInsightsPath {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_NetworkInsightsPath) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_NetworkInsightsPath {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_NetworkInsightsPath) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_NetworkInsightsPath {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_NetworkInsightsPath) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_NetworkInsightsPath {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_NetworkInsightsPath) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_NetworkInsightsPath {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Destination updates property "Destination".
func (t *AWS_EC2_NetworkInsightsPath) Set__Destination(v cfz.Expression[string]) *AWS_EC2_NetworkInsightsPath {
	t.Destination = v
	return t
}

// SetV__Destination updates property "Destination".
func (t *AWS_EC2_NetworkInsightsPath) SetV__Destination(v string) *AWS_EC2_NetworkInsightsPath {
	t.Destination = cfz.V(v)
	return t
}

// Set__DestinationIp updates property "DestinationIp".
func (t *AWS_EC2_NetworkInsightsPath) Set__DestinationIp(v cfz.Expression[string]) *AWS_EC2_NetworkInsightsPath {
	t.DestinationIp = v
	return t
}

// SetV__DestinationIp updates property "DestinationIp".
func (t *AWS_EC2_NetworkInsightsPath) SetV__DestinationIp(v string) *AWS_EC2_NetworkInsightsPath {
	t.DestinationIp = cfz.V(v)
	return t
}

// Set__DestinationPort updates property "DestinationPort".
func (t *AWS_EC2_NetworkInsightsPath) Set__DestinationPort(v cfz.Expression[int32]) *AWS_EC2_NetworkInsightsPath {
	t.DestinationPort = v
	return t
}

// SetV__DestinationPort updates property "DestinationPort".
func (t *AWS_EC2_NetworkInsightsPath) SetV__DestinationPort(v int32) *AWS_EC2_NetworkInsightsPath {
	t.DestinationPort = cfz.V(v)
	return t
}

// Set__FilterAtDestination updates property "FilterAtDestination".
func (t *AWS_EC2_NetworkInsightsPath) Set__FilterAtDestination(v cfz.Expression[AWS_EC2_NetworkInsightsPath_PathFilter]) *AWS_EC2_NetworkInsightsPath {
	t.FilterAtDestination = v
	return t
}

// SetV__FilterAtDestination updates property "FilterAtDestination".
func (t *AWS_EC2_NetworkInsightsPath) SetV__FilterAtDestination(v AWS_EC2_NetworkInsightsPath_PathFilter) *AWS_EC2_NetworkInsightsPath {
	t.FilterAtDestination = cfz.V(v)
	return t
}

// Set__FilterAtSource updates property "FilterAtSource".
func (t *AWS_EC2_NetworkInsightsPath) Set__FilterAtSource(v cfz.Expression[AWS_EC2_NetworkInsightsPath_PathFilter]) *AWS_EC2_NetworkInsightsPath {
	t.FilterAtSource = v
	return t
}

// SetV__FilterAtSource updates property "FilterAtSource".
func (t *AWS_EC2_NetworkInsightsPath) SetV__FilterAtSource(v AWS_EC2_NetworkInsightsPath_PathFilter) *AWS_EC2_NetworkInsightsPath {
	t.FilterAtSource = cfz.V(v)
	return t
}

// Set__Protocol updates property "Protocol".
func (t *AWS_EC2_NetworkInsightsPath) Set__Protocol(v cfz.Expression[string]) *AWS_EC2_NetworkInsightsPath {
	t.Protocol = v
	return t
}

// SetV__Protocol updates property "Protocol".
func (t *AWS_EC2_NetworkInsightsPath) SetV__Protocol(v string) *AWS_EC2_NetworkInsightsPath {
	t.Protocol = cfz.V(v)
	return t
}

// Set__Source updates property "Source".
func (t *AWS_EC2_NetworkInsightsPath) Set__Source(v cfz.Expression[string]) *AWS_EC2_NetworkInsightsPath {
	t.Source = v
	return t
}

// SetV__Source updates property "Source".
func (t *AWS_EC2_NetworkInsightsPath) SetV__Source(v string) *AWS_EC2_NetworkInsightsPath {
	t.Source = cfz.V(v)
	return t
}

// Set__SourceIp updates property "SourceIp".
func (t *AWS_EC2_NetworkInsightsPath) Set__SourceIp(v cfz.Expression[string]) *AWS_EC2_NetworkInsightsPath {
	t.SourceIp = v
	return t
}

// SetV__SourceIp updates property "SourceIp".
func (t *AWS_EC2_NetworkInsightsPath) SetV__SourceIp(v string) *AWS_EC2_NetworkInsightsPath {
	t.SourceIp = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EC2_NetworkInsightsPath) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_EC2_NetworkInsightsPath {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_EC2_NetworkInsightsPath) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_EC2_NetworkInsightsPath {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_EC2_NetworkInsightsPath) SetSV__Tags(v ...cfz.Tag) *AWS_EC2_NetworkInsightsPath {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_NetworkInsightsPath) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CreatedDate returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreatedDate
func (t *AWS_EC2_NetworkInsightsPath) GetAtt__CreatedDate() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsPath__AttributesMap.CreatedDate))
}

// GetAtt__DestinationArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DestinationArn
func (t *AWS_EC2_NetworkInsightsPath) GetAtt__DestinationArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsPath__AttributesMap.DestinationArn))
}

// GetAtt__NetworkInsightsPathArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: NetworkInsightsPathArn
func (t *AWS_EC2_NetworkInsightsPath) GetAtt__NetworkInsightsPathArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsPath__AttributesMap.NetworkInsightsPathArn))
}

// GetAtt__NetworkInsightsPathId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: NetworkInsightsPathId
func (t *AWS_EC2_NetworkInsightsPath) GetAtt__NetworkInsightsPathId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsPath__AttributesMap.NetworkInsightsPathId))
}

// GetAtt__SourceArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: SourceArn
func (t *AWS_EC2_NetworkInsightsPath) GetAtt__SourceArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsPath__AttributesMap.SourceArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_NetworkInsightsPath) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreatedDate returns a conventionally configured output for an attribute of this resource.
// Attribute: CreatedDate
func (t *AWS_EC2_NetworkInsightsPath) GetConventionalOutputAtt__CreatedDate(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreatedDate", t.GetAtt__CreatedDate())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DestinationArn returns a conventionally configured output for an attribute of this resource.
// Attribute: DestinationArn
func (t *AWS_EC2_NetworkInsightsPath) GetConventionalOutputAtt__DestinationArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDestinationArn", t.GetAtt__DestinationArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__NetworkInsightsPathArn returns a conventionally configured output for an attribute of this resource.
// Attribute: NetworkInsightsPathArn
func (t *AWS_EC2_NetworkInsightsPath) GetConventionalOutputAtt__NetworkInsightsPathArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttNetworkInsightsPathArn", t.GetAtt__NetworkInsightsPathArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__NetworkInsightsPathId returns a conventionally configured output for an attribute of this resource.
// Attribute: NetworkInsightsPathId
func (t *AWS_EC2_NetworkInsightsPath) GetConventionalOutputAtt__NetworkInsightsPathId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttNetworkInsightsPathId", t.GetAtt__NetworkInsightsPathId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__SourceArn returns a conventionally configured output for an attribute of this resource.
// Attribute: SourceArn
func (t *AWS_EC2_NetworkInsightsPath) GetConventionalOutputAtt__SourceArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttSourceArn", t.GetAtt__SourceArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_NetworkInsightsPath) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_NetworkInsightsPath

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

func (t *AWS_EC2_NetworkInsightsPath) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
