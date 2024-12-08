// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_elasticloadbalancing

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ElasticLoadBalancing_LoadBalancer_Listeners__Type is the CloudFormation type for AWS::ElasticLoadBalancing::LoadBalancer.Listeners.
	AWS_ElasticLoadBalancing_LoadBalancer_Listeners__Type = "AWS::ElasticLoadBalancing::LoadBalancer.Listeners"
)

var (
	// AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesMap reports all the CloudFormation properties for AWS::ElasticLoadBalancing::LoadBalancer.Listeners.
	AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesMap = struct {
		InstancePort     string
		InstanceProtocol string
		LoadBalancerPort string
		PolicyNames      string
		Protocol         string
		SSLCertificateId string
	}{
		InstancePort:     "InstancePort",
		InstanceProtocol: "InstanceProtocol",
		LoadBalancerPort: "LoadBalancerPort",
		PolicyNames:      "PolicyNames",
		Protocol:         "Protocol",
		SSLCertificateId: "SSLCertificateId",
	}

	// AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesSlice reports all the CloudFormation properties for AWS::ElasticLoadBalancing::LoadBalancer.Listeners.
	AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesSlice = []string{
		AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesMap.InstancePort,
		AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesMap.InstanceProtocol,
		AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesMap.LoadBalancerPort,
		AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesMap.PolicyNames,
		AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesMap.Protocol,
		AWS_ElasticLoadBalancing_LoadBalancer_Listeners__PropertiesMap.SSLCertificateId,
	}
)

// AWS_ElasticLoadBalancing_LoadBalancer_Listeners is a binding for AWS::ElasticLoadBalancing::LoadBalancer.Listeners.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html
type AWS_ElasticLoadBalancing_LoadBalancer_Listeners struct {
	// InstancePort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-instanceport
	InstancePort cfz.Expression[string] `json:"InstancePort,omitempty"`

	// InstanceProtocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-instanceprotocol
	InstanceProtocol cfz.Expression[string] `json:"InstanceProtocol,omitempty"`

	// LoadBalancerPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-loadbalancerport
	LoadBalancerPort cfz.Expression[string] `json:"LoadBalancerPort,omitempty"`

	// PolicyNames is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames
	PolicyNames cfz.ExpressionSlice[string] `json:"PolicyNames,omitempty"`

	// Protocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-protocol
	Protocol cfz.Expression[string] `json:"Protocol,omitempty"`

	// SSLCertificateId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-sslcertificateid
	SSLCertificateId cfz.Expression[string] `json:"SSLCertificateId,omitempty"`
}

// New__AWS_ElasticLoadBalancing_LoadBalancer_Listeners initializes a new AWS_ElasticLoadBalancing_LoadBalancer_Listeners.
func New__AWS_ElasticLoadBalancing_LoadBalancer_Listeners() AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	return AWS_ElasticLoadBalancing_LoadBalancer_Listeners{}
}

// GetType returns the CloudFormation type.
func (AWS_ElasticLoadBalancing_LoadBalancer_Listeners) GetType() string {
	return AWS_ElasticLoadBalancing_LoadBalancer_Listeners__Type
}

// Set__InstancePort updates property "InstancePort".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) Set__InstancePort(v cfz.Expression[string]) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.InstancePort = v
	return t
}

// SetV__InstancePort updates property "InstancePort".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) SetV__InstancePort(v string) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.InstancePort = cfz.V(v)
	return t
}

// Set__InstanceProtocol updates property "InstanceProtocol".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) Set__InstanceProtocol(v cfz.Expression[string]) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.InstanceProtocol = v
	return t
}

// SetV__InstanceProtocol updates property "InstanceProtocol".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) SetV__InstanceProtocol(v string) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.InstanceProtocol = cfz.V(v)
	return t
}

// Set__LoadBalancerPort updates property "LoadBalancerPort".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) Set__LoadBalancerPort(v cfz.Expression[string]) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.LoadBalancerPort = v
	return t
}

// SetV__LoadBalancerPort updates property "LoadBalancerPort".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) SetV__LoadBalancerPort(v string) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.LoadBalancerPort = cfz.V(v)
	return t
}

// Set__PolicyNames updates property "PolicyNames".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) Set__PolicyNames(v cfz.ExpressionSlice[string]) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.PolicyNames = v
	return t
}

// SetS__PolicyNames updates property "PolicyNames".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) SetS__PolicyNames(v ...cfz.Expression[string]) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.PolicyNames = cfz.S(v...)
	return t
}

// SetSV__PolicyNames updates property "PolicyNames".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) SetSV__PolicyNames(v ...string) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.PolicyNames = cfz.SV(v...)
	return t
}

// Set__Protocol updates property "Protocol".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) Set__Protocol(v cfz.Expression[string]) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.Protocol = v
	return t
}

// SetV__Protocol updates property "Protocol".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) SetV__Protocol(v string) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.Protocol = cfz.V(v)
	return t
}

// Set__SSLCertificateId updates property "SSLCertificateId".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) Set__SSLCertificateId(v cfz.Expression[string]) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.SSLCertificateId = v
	return t
}

// SetV__SSLCertificateId updates property "SSLCertificateId".
func (t AWS_ElasticLoadBalancing_LoadBalancer_Listeners) SetV__SSLCertificateId(v string) AWS_ElasticLoadBalancing_LoadBalancer_Listeners {
	t.SSLCertificateId = cfz.V(v)
	return t
}
