// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_elasticloadbalancingv2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ElasticLoadBalancingV2_TargetGroup)(nil)
	_ cfz.Resource                   = (*AWS_ElasticLoadBalancingV2_TargetGroup)(nil)
)

const (
	// AWS_ElasticLoadBalancingV2_TargetGroup__Type is the CloudFormation type for AWS::ElasticLoadBalancingV2::TargetGroup.
	AWS_ElasticLoadBalancingV2_TargetGroup__Type = "AWS::ElasticLoadBalancingV2::TargetGroup"
)

var (
	// AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap reports all the CloudFormation attributes for AWS::ElasticLoadBalancingV2::TargetGroup.
	AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap = struct {
		LoadBalancerArns    string
		TargetGroupArn      string
		TargetGroupFullName string
		TargetGroupName     string
	}{
		LoadBalancerArns:    "LoadBalancerArns",
		TargetGroupArn:      "TargetGroupArn",
		TargetGroupFullName: "TargetGroupFullName",
		TargetGroupName:     "TargetGroupName",
	}

	// AWS_ElasticLoadBalancingV2_TargetGroup__AttributesSlice reports all the CloudFormation attributes for AWS::ElasticLoadBalancingV2::TargetGroup.
	AWS_ElasticLoadBalancingV2_TargetGroup__AttributesSlice = []string{
		AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap.LoadBalancerArns,
		AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap.TargetGroupArn,
		AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap.TargetGroupFullName,
		AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap.TargetGroupName,
	}
)

var (
	// AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::TargetGroup.
	AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap = struct {
		HealthCheckEnabled         string
		HealthCheckIntervalSeconds string
		HealthCheckPath            string
		HealthCheckPort            string
		HealthCheckProtocol        string
		HealthCheckTimeoutSeconds  string
		HealthyThresholdCount      string
		IpAddressType              string
		Matcher                    string
		Name                       string
		Port                       string
		Protocol                   string
		ProtocolVersion            string
		Tags                       string
		TargetGroupAttributes      string
		TargetType                 string
		Targets                    string
		UnhealthyThresholdCount    string
		VpcId                      string
	}{
		HealthCheckEnabled:         "HealthCheckEnabled",
		HealthCheckIntervalSeconds: "HealthCheckIntervalSeconds",
		HealthCheckPath:            "HealthCheckPath",
		HealthCheckPort:            "HealthCheckPort",
		HealthCheckProtocol:        "HealthCheckProtocol",
		HealthCheckTimeoutSeconds:  "HealthCheckTimeoutSeconds",
		HealthyThresholdCount:      "HealthyThresholdCount",
		IpAddressType:              "IpAddressType",
		Matcher:                    "Matcher",
		Name:                       "Name",
		Port:                       "Port",
		Protocol:                   "Protocol",
		ProtocolVersion:            "ProtocolVersion",
		Tags:                       "Tags",
		TargetGroupAttributes:      "TargetGroupAttributes",
		TargetType:                 "TargetType",
		Targets:                    "Targets",
		UnhealthyThresholdCount:    "UnhealthyThresholdCount",
		VpcId:                      "VpcId",
	}

	// AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesSlice reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::TargetGroup.
	AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesSlice = []string{
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.HealthCheckEnabled,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.HealthCheckIntervalSeconds,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.HealthCheckPath,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.HealthCheckPort,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.HealthCheckProtocol,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.HealthCheckTimeoutSeconds,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.HealthyThresholdCount,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.IpAddressType,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.Matcher,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.Name,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.Port,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.Protocol,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.ProtocolVersion,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.Tags,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.TargetGroupAttributes,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.TargetType,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.Targets,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.UnhealthyThresholdCount,
		AWS_ElasticLoadBalancingV2_TargetGroup__PropertiesMap.VpcId,
	}
)

// AWS_ElasticLoadBalancingV2_TargetGroup is a binding for AWS::ElasticLoadBalancingV2::TargetGroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html
type AWS_ElasticLoadBalancingV2_TargetGroup struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// HealthCheckEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckenabled
	HealthCheckEnabled cfz.Expression[bool] `json:"HealthCheckEnabled,omitempty"`

	// HealthCheckIntervalSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckintervalseconds
	HealthCheckIntervalSeconds cfz.Expression[int32] `json:"HealthCheckIntervalSeconds,omitempty"`

	// HealthCheckPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckpath
	HealthCheckPath cfz.Expression[string] `json:"HealthCheckPath,omitempty"`

	// HealthCheckPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckport
	HealthCheckPort cfz.Expression[string] `json:"HealthCheckPort,omitempty"`

	// HealthCheckProtocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckprotocol
	HealthCheckProtocol cfz.Expression[string] `json:"HealthCheckProtocol,omitempty"`

	// HealthCheckTimeoutSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthchecktimeoutseconds
	HealthCheckTimeoutSeconds cfz.Expression[int32] `json:"HealthCheckTimeoutSeconds,omitempty"`

	// HealthyThresholdCount is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthythresholdcount
	HealthyThresholdCount cfz.Expression[int32] `json:"HealthyThresholdCount,omitempty"`

	// IpAddressType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-ipaddresstype
	IpAddressType cfz.Expression[string] `json:"IpAddressType,omitempty"`

	// Matcher is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-matcher
	Matcher cfz.Expression[AWS_ElasticLoadBalancingV2_TargetGroup_Matcher] `json:"Matcher,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Port is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-port
	Port cfz.Expression[int32] `json:"Port,omitempty"`

	// Protocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-protocol
	Protocol cfz.Expression[string] `json:"Protocol,omitempty"`

	// ProtocolVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-protocolversion
	ProtocolVersion cfz.Expression[string] `json:"ProtocolVersion,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// TargetGroupAttributes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-targetgroupattributes
	TargetGroupAttributes cfz.ExpressionSlice[AWS_ElasticLoadBalancingV2_TargetGroup_TargetGroupAttribute] `json:"TargetGroupAttributes,omitempty"`

	// TargetType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-targettype
	TargetType cfz.Expression[string] `json:"TargetType,omitempty"`

	// Targets is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-targets
	Targets cfz.ExpressionSlice[AWS_ElasticLoadBalancingV2_TargetGroup_TargetDescription] `json:"Targets,omitempty"`

	// UnhealthyThresholdCount is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-unhealthythresholdcount
	UnhealthyThresholdCount cfz.Expression[int32] `json:"UnhealthyThresholdCount,omitempty"`

	// VpcId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-vpcid
	VpcId cfz.Expression[string] `json:"VpcId,omitempty"`
}

// New__AWS_ElasticLoadBalancingV2_TargetGroup initializes a new *AWS_ElasticLoadBalancingV2_TargetGroup.
func New__AWS_ElasticLoadBalancingV2_TargetGroup(logicalName string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	return &AWS_ElasticLoadBalancingV2_TargetGroup{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ElasticLoadBalancingV2_TargetGroup) GetType() string {
	return AWS_ElasticLoadBalancingV2_TargetGroup__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__LogicalName(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__RequestedOutputs(v []cfz.Output) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Add__RequestedOutputs(v ...cfz.Output) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__HealthCheckEnabled updates property "HealthCheckEnabled".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__HealthCheckEnabled(v cfz.Expression[bool]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckEnabled = v
	return t
}

// SetV__HealthCheckEnabled updates property "HealthCheckEnabled".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__HealthCheckEnabled(v bool) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckEnabled = cfz.V(v)
	return t
}

// Set__HealthCheckIntervalSeconds updates property "HealthCheckIntervalSeconds".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__HealthCheckIntervalSeconds(v cfz.Expression[int32]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckIntervalSeconds = v
	return t
}

// SetV__HealthCheckIntervalSeconds updates property "HealthCheckIntervalSeconds".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__HealthCheckIntervalSeconds(v int32) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckIntervalSeconds = cfz.V(v)
	return t
}

// Set__HealthCheckPath updates property "HealthCheckPath".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__HealthCheckPath(v cfz.Expression[string]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckPath = v
	return t
}

// SetV__HealthCheckPath updates property "HealthCheckPath".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__HealthCheckPath(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckPath = cfz.V(v)
	return t
}

// Set__HealthCheckPort updates property "HealthCheckPort".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__HealthCheckPort(v cfz.Expression[string]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckPort = v
	return t
}

// SetV__HealthCheckPort updates property "HealthCheckPort".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__HealthCheckPort(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckPort = cfz.V(v)
	return t
}

// Set__HealthCheckProtocol updates property "HealthCheckProtocol".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__HealthCheckProtocol(v cfz.Expression[string]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckProtocol = v
	return t
}

// SetV__HealthCheckProtocol updates property "HealthCheckProtocol".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__HealthCheckProtocol(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckProtocol = cfz.V(v)
	return t
}

// Set__HealthCheckTimeoutSeconds updates property "HealthCheckTimeoutSeconds".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__HealthCheckTimeoutSeconds(v cfz.Expression[int32]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckTimeoutSeconds = v
	return t
}

// SetV__HealthCheckTimeoutSeconds updates property "HealthCheckTimeoutSeconds".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__HealthCheckTimeoutSeconds(v int32) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthCheckTimeoutSeconds = cfz.V(v)
	return t
}

// Set__HealthyThresholdCount updates property "HealthyThresholdCount".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__HealthyThresholdCount(v cfz.Expression[int32]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthyThresholdCount = v
	return t
}

// SetV__HealthyThresholdCount updates property "HealthyThresholdCount".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__HealthyThresholdCount(v int32) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.HealthyThresholdCount = cfz.V(v)
	return t
}

// Set__IpAddressType updates property "IpAddressType".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__IpAddressType(v cfz.Expression[string]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.IpAddressType = v
	return t
}

// SetV__IpAddressType updates property "IpAddressType".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__IpAddressType(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.IpAddressType = cfz.V(v)
	return t
}

// Set__Matcher updates property "Matcher".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__Matcher(v cfz.Expression[AWS_ElasticLoadBalancingV2_TargetGroup_Matcher]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Matcher = v
	return t
}

// SetV__Matcher updates property "Matcher".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__Matcher(v AWS_ElasticLoadBalancingV2_TargetGroup_Matcher) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Matcher = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__Name(v cfz.Expression[string]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__Name(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Name = cfz.V(v)
	return t
}

// Set__Port updates property "Port".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__Port(v cfz.Expression[int32]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Port = v
	return t
}

// SetV__Port updates property "Port".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__Port(v int32) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Port = cfz.V(v)
	return t
}

// Set__Protocol updates property "Protocol".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__Protocol(v cfz.Expression[string]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Protocol = v
	return t
}

// SetV__Protocol updates property "Protocol".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__Protocol(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Protocol = cfz.V(v)
	return t
}

// Set__ProtocolVersion updates property "ProtocolVersion".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__ProtocolVersion(v cfz.Expression[string]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.ProtocolVersion = v
	return t
}

// SetV__ProtocolVersion updates property "ProtocolVersion".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__ProtocolVersion(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.ProtocolVersion = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetSV__Tags(v ...cfz.Tag) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TargetGroupAttributes updates property "TargetGroupAttributes".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__TargetGroupAttributes(v cfz.ExpressionSlice[AWS_ElasticLoadBalancingV2_TargetGroup_TargetGroupAttribute]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.TargetGroupAttributes = v
	return t
}

// SetS__TargetGroupAttributes updates property "TargetGroupAttributes".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetS__TargetGroupAttributes(v ...cfz.Expression[AWS_ElasticLoadBalancingV2_TargetGroup_TargetGroupAttribute]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.TargetGroupAttributes = cfz.S(v...)
	return t
}

// SetSV__TargetGroupAttributes updates property "TargetGroupAttributes".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetSV__TargetGroupAttributes(v ...AWS_ElasticLoadBalancingV2_TargetGroup_TargetGroupAttribute) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.TargetGroupAttributes = cfz.SV(v...)
	return t
}

// Set__TargetType updates property "TargetType".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__TargetType(v cfz.Expression[string]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.TargetType = v
	return t
}

// SetV__TargetType updates property "TargetType".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__TargetType(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.TargetType = cfz.V(v)
	return t
}

// Set__Targets updates property "Targets".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__Targets(v cfz.ExpressionSlice[AWS_ElasticLoadBalancingV2_TargetGroup_TargetDescription]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Targets = v
	return t
}

// SetS__Targets updates property "Targets".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetS__Targets(v ...cfz.Expression[AWS_ElasticLoadBalancingV2_TargetGroup_TargetDescription]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Targets = cfz.S(v...)
	return t
}

// SetSV__Targets updates property "Targets".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetSV__Targets(v ...AWS_ElasticLoadBalancingV2_TargetGroup_TargetDescription) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.Targets = cfz.SV(v...)
	return t
}

// Set__UnhealthyThresholdCount updates property "UnhealthyThresholdCount".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__UnhealthyThresholdCount(v cfz.Expression[int32]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.UnhealthyThresholdCount = v
	return t
}

// SetV__UnhealthyThresholdCount updates property "UnhealthyThresholdCount".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__UnhealthyThresholdCount(v int32) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.UnhealthyThresholdCount = cfz.V(v)
	return t
}

// Set__VpcId updates property "VpcId".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Set__VpcId(v cfz.Expression[string]) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.VpcId = v
	return t
}

// SetV__VpcId updates property "VpcId".
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) SetV__VpcId(v string) *AWS_ElasticLoadBalancingV2_TargetGroup {
	t.VpcId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAttSlice__LoadBalancerArns returns a $cfz.ExpressionSlice[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LoadBalancerArns
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetAttSlice__LoadBalancerArns() cfz.ExpressionSlice[string] {
	return cfz.GetAttSlice[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap.LoadBalancerArns))
}

// GetAtt__TargetGroupArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TargetGroupArn
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetAtt__TargetGroupArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap.TargetGroupArn))
}

// GetAtt__TargetGroupFullName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TargetGroupFullName
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetAtt__TargetGroupFullName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap.TargetGroupFullName))
}

// GetAtt__TargetGroupName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TargetGroupName
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetAtt__TargetGroupName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ElasticLoadBalancingV2_TargetGroup__AttributesMap.TargetGroupName))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LoadBalancerArns returns a conventionally configured output for an attribute of this resource.
// Attribute: LoadBalancerArns
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetConventionalOutputAtt__LoadBalancerArns(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLoadBalancerArns", t.GetAttSlice__LoadBalancerArns())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TargetGroupArn returns a conventionally configured output for an attribute of this resource.
// Attribute: TargetGroupArn
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetConventionalOutputAtt__TargetGroupArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTargetGroupArn", t.GetAtt__TargetGroupArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TargetGroupFullName returns a conventionally configured output for an attribute of this resource.
// Attribute: TargetGroupFullName
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetConventionalOutputAtt__TargetGroupFullName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTargetGroupFullName", t.GetAtt__TargetGroupFullName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TargetGroupName returns a conventionally configured output for an attribute of this resource.
// Attribute: TargetGroupName
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) GetConventionalOutputAtt__TargetGroupName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTargetGroupName", t.GetAtt__TargetGroupName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ElasticLoadBalancingV2_TargetGroup) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ElasticLoadBalancingV2_TargetGroup

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

func (t *AWS_ElasticLoadBalancingV2_TargetGroup) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
