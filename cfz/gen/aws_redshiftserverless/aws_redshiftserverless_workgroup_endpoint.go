// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_redshiftserverless

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_RedshiftServerless_Workgroup_Endpoint__Type is the CloudFormation type for AWS::RedshiftServerless::Workgroup.Endpoint.
	AWS_RedshiftServerless_Workgroup_Endpoint__Type = "AWS::RedshiftServerless::Workgroup.Endpoint"
)

var (
	// AWS_RedshiftServerless_Workgroup_Endpoint__PropertiesMap reports all the CloudFormation properties for AWS::RedshiftServerless::Workgroup.Endpoint.
	AWS_RedshiftServerless_Workgroup_Endpoint__PropertiesMap = struct {
		Address      string
		Port         string
		VpcEndpoints string
	}{
		Address:      "Address",
		Port:         "Port",
		VpcEndpoints: "VpcEndpoints",
	}

	// AWS_RedshiftServerless_Workgroup_Endpoint__PropertiesSlice reports all the CloudFormation properties for AWS::RedshiftServerless::Workgroup.Endpoint.
	AWS_RedshiftServerless_Workgroup_Endpoint__PropertiesSlice = []string{
		AWS_RedshiftServerless_Workgroup_Endpoint__PropertiesMap.Address,
		AWS_RedshiftServerless_Workgroup_Endpoint__PropertiesMap.Port,
		AWS_RedshiftServerless_Workgroup_Endpoint__PropertiesMap.VpcEndpoints,
	}
)

// AWS_RedshiftServerless_Workgroup_Endpoint is a binding for AWS::RedshiftServerless::Workgroup.Endpoint.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-endpoint.html
type AWS_RedshiftServerless_Workgroup_Endpoint struct {
	// Address is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-endpoint.html#cfn-redshiftserverless-workgroup-endpoint-address
	Address cfz.Expression[string] `json:"Address,omitempty"`

	// Port is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-endpoint.html#cfn-redshiftserverless-workgroup-endpoint-port
	Port cfz.Expression[int32] `json:"Port,omitempty"`

	// VpcEndpoints is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-endpoint.html#cfn-redshiftserverless-workgroup-endpoint-vpcendpoints
	VpcEndpoints cfz.ExpressionSlice[AWS_RedshiftServerless_Workgroup_VpcEndpoint] `json:"VpcEndpoints,omitempty"`
}

// New__AWS_RedshiftServerless_Workgroup_Endpoint initializes a new AWS_RedshiftServerless_Workgroup_Endpoint.
func New__AWS_RedshiftServerless_Workgroup_Endpoint() AWS_RedshiftServerless_Workgroup_Endpoint {
	return AWS_RedshiftServerless_Workgroup_Endpoint{}
}

// GetType returns the CloudFormation type.
func (AWS_RedshiftServerless_Workgroup_Endpoint) GetType() string {
	return AWS_RedshiftServerless_Workgroup_Endpoint__Type
}

// Set__Address updates property "Address".
func (t AWS_RedshiftServerless_Workgroup_Endpoint) Set__Address(v cfz.Expression[string]) AWS_RedshiftServerless_Workgroup_Endpoint {
	t.Address = v
	return t
}

// SetV__Address updates property "Address".
func (t AWS_RedshiftServerless_Workgroup_Endpoint) SetV__Address(v string) AWS_RedshiftServerless_Workgroup_Endpoint {
	t.Address = cfz.V(v)
	return t
}

// Set__Port updates property "Port".
func (t AWS_RedshiftServerless_Workgroup_Endpoint) Set__Port(v cfz.Expression[int32]) AWS_RedshiftServerless_Workgroup_Endpoint {
	t.Port = v
	return t
}

// SetV__Port updates property "Port".
func (t AWS_RedshiftServerless_Workgroup_Endpoint) SetV__Port(v int32) AWS_RedshiftServerless_Workgroup_Endpoint {
	t.Port = cfz.V(v)
	return t
}

// Set__VpcEndpoints updates property "VpcEndpoints".
func (t AWS_RedshiftServerless_Workgroup_Endpoint) Set__VpcEndpoints(v cfz.ExpressionSlice[AWS_RedshiftServerless_Workgroup_VpcEndpoint]) AWS_RedshiftServerless_Workgroup_Endpoint {
	t.VpcEndpoints = v
	return t
}

// SetS__VpcEndpoints updates property "VpcEndpoints".
func (t AWS_RedshiftServerless_Workgroup_Endpoint) SetS__VpcEndpoints(v ...cfz.Expression[AWS_RedshiftServerless_Workgroup_VpcEndpoint]) AWS_RedshiftServerless_Workgroup_Endpoint {
	t.VpcEndpoints = cfz.S(v...)
	return t
}

// SetSV__VpcEndpoints updates property "VpcEndpoints".
func (t AWS_RedshiftServerless_Workgroup_Endpoint) SetSV__VpcEndpoints(v ...AWS_RedshiftServerless_Workgroup_VpcEndpoint) AWS_RedshiftServerless_Workgroup_Endpoint {
	t.VpcEndpoints = cfz.SV(v...)
	return t
}
