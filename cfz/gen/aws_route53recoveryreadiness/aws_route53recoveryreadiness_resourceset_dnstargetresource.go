// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_route53recoveryreadiness

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__Type is the CloudFormation type for AWS::Route53RecoveryReadiness::ResourceSet.DNSTargetResource.
	AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__Type = "AWS::Route53RecoveryReadiness::ResourceSet.DNSTargetResource"
)

var (
	// AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__PropertiesMap reports all the CloudFormation properties for AWS::Route53RecoveryReadiness::ResourceSet.DNSTargetResource.
	AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__PropertiesMap = struct {
		DomainName     string
		HostedZoneArn  string
		RecordSetId    string
		RecordType     string
		TargetResource string
	}{
		DomainName:     "DomainName",
		HostedZoneArn:  "HostedZoneArn",
		RecordSetId:    "RecordSetId",
		RecordType:     "RecordType",
		TargetResource: "TargetResource",
	}

	// AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__PropertiesSlice reports all the CloudFormation properties for AWS::Route53RecoveryReadiness::ResourceSet.DNSTargetResource.
	AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__PropertiesSlice = []string{
		AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__PropertiesMap.DomainName,
		AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__PropertiesMap.HostedZoneArn,
		AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__PropertiesMap.RecordSetId,
		AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__PropertiesMap.RecordType,
		AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__PropertiesMap.TargetResource,
	}
)

// AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource is a binding for AWS::Route53RecoveryReadiness::ResourceSet.DNSTargetResource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html
type AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource struct {
	// DomainName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-domainname
	DomainName cfz.Expression[string] `json:"DomainName,omitempty"`

	// HostedZoneArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-hostedzonearn
	HostedZoneArn cfz.Expression[string] `json:"HostedZoneArn,omitempty"`

	// RecordSetId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-recordsetid
	RecordSetId cfz.Expression[string] `json:"RecordSetId,omitempty"`

	// RecordType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-recordtype
	RecordType cfz.Expression[string] `json:"RecordType,omitempty"`

	// TargetResource is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-dnstargetresource.html#cfn-route53recoveryreadiness-resourceset-dnstargetresource-targetresource
	TargetResource cfz.Expression[AWS_Route53RecoveryReadiness_ResourceSet_TargetResource] `json:"TargetResource,omitempty"`
}

// New__AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource initializes a new AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource.
func New__AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource() AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	return AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource{}
}

// GetType returns the CloudFormation type.
func (AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) GetType() string {
	return AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource__Type
}

// Set__DomainName updates property "DomainName".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) Set__DomainName(v cfz.Expression[string]) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.DomainName = v
	return t
}

// SetV__DomainName updates property "DomainName".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) SetV__DomainName(v string) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.DomainName = cfz.V(v)
	return t
}

// Set__HostedZoneArn updates property "HostedZoneArn".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) Set__HostedZoneArn(v cfz.Expression[string]) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.HostedZoneArn = v
	return t
}

// SetV__HostedZoneArn updates property "HostedZoneArn".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) SetV__HostedZoneArn(v string) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.HostedZoneArn = cfz.V(v)
	return t
}

// Set__RecordSetId updates property "RecordSetId".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) Set__RecordSetId(v cfz.Expression[string]) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.RecordSetId = v
	return t
}

// SetV__RecordSetId updates property "RecordSetId".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) SetV__RecordSetId(v string) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.RecordSetId = cfz.V(v)
	return t
}

// Set__RecordType updates property "RecordType".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) Set__RecordType(v cfz.Expression[string]) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.RecordType = v
	return t
}

// SetV__RecordType updates property "RecordType".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) SetV__RecordType(v string) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.RecordType = cfz.V(v)
	return t
}

// Set__TargetResource updates property "TargetResource".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) Set__TargetResource(v cfz.Expression[AWS_Route53RecoveryReadiness_ResourceSet_TargetResource]) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.TargetResource = v
	return t
}

// SetV__TargetResource updates property "TargetResource".
func (t AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource) SetV__TargetResource(v AWS_Route53RecoveryReadiness_ResourceSet_TargetResource) AWS_Route53RecoveryReadiness_ResourceSet_DNSTargetResource {
	t.TargetResource = cfz.V(v)
	return t
}
