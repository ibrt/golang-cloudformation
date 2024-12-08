// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__Type is the CloudFormation type for AWS::AppMesh::VirtualGateway.VirtualGatewaySpec.
	AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__Type = "AWS::AppMesh::VirtualGateway.VirtualGatewaySpec"
)

var (
	// AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewaySpec.
	AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__PropertiesMap = struct {
		BackendDefaults string
		Listeners       string
		Logging         string
	}{
		BackendDefaults: "BackendDefaults",
		Listeners:       "Listeners",
		Logging:         "Logging",
	}

	// AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::VirtualGateway.VirtualGatewaySpec.
	AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__PropertiesSlice = []string{
		AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__PropertiesMap.BackendDefaults,
		AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__PropertiesMap.Listeners,
		AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__PropertiesMap.Logging,
	}
)

// AWS_AppMesh_VirtualGateway_VirtualGatewaySpec is a binding for AWS::AppMesh::VirtualGateway.VirtualGatewaySpec.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayspec.html
type AWS_AppMesh_VirtualGateway_VirtualGatewaySpec struct {
	// BackendDefaults is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayspec.html#cfn-appmesh-virtualgateway-virtualgatewayspec-backenddefaults
	BackendDefaults cfz.Expression[AWS_AppMesh_VirtualGateway_VirtualGatewayBackendDefaults] `json:"BackendDefaults,omitempty"`

	// Listeners is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayspec.html#cfn-appmesh-virtualgateway-virtualgatewayspec-listeners
	Listeners cfz.ExpressionSlice[AWS_AppMesh_VirtualGateway_VirtualGatewayListener] `json:"Listeners,omitempty"`

	// Logging is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayspec.html#cfn-appmesh-virtualgateway-virtualgatewayspec-logging
	Logging cfz.Expression[AWS_AppMesh_VirtualGateway_VirtualGatewayLogging] `json:"Logging,omitempty"`
}

// New__AWS_AppMesh_VirtualGateway_VirtualGatewaySpec initializes a new AWS_AppMesh_VirtualGateway_VirtualGatewaySpec.
func New__AWS_AppMesh_VirtualGateway_VirtualGatewaySpec() AWS_AppMesh_VirtualGateway_VirtualGatewaySpec {
	return AWS_AppMesh_VirtualGateway_VirtualGatewaySpec{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_VirtualGateway_VirtualGatewaySpec) GetType() string {
	return AWS_AppMesh_VirtualGateway_VirtualGatewaySpec__Type
}

// Set__BackendDefaults updates property "BackendDefaults".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewaySpec) Set__BackendDefaults(v cfz.Expression[AWS_AppMesh_VirtualGateway_VirtualGatewayBackendDefaults]) AWS_AppMesh_VirtualGateway_VirtualGatewaySpec {
	t.BackendDefaults = v
	return t
}

// SetV__BackendDefaults updates property "BackendDefaults".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewaySpec) SetV__BackendDefaults(v AWS_AppMesh_VirtualGateway_VirtualGatewayBackendDefaults) AWS_AppMesh_VirtualGateway_VirtualGatewaySpec {
	t.BackendDefaults = cfz.V(v)
	return t
}

// Set__Listeners updates property "Listeners".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewaySpec) Set__Listeners(v cfz.ExpressionSlice[AWS_AppMesh_VirtualGateway_VirtualGatewayListener]) AWS_AppMesh_VirtualGateway_VirtualGatewaySpec {
	t.Listeners = v
	return t
}

// SetS__Listeners updates property "Listeners".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewaySpec) SetS__Listeners(v ...cfz.Expression[AWS_AppMesh_VirtualGateway_VirtualGatewayListener]) AWS_AppMesh_VirtualGateway_VirtualGatewaySpec {
	t.Listeners = cfz.S(v...)
	return t
}

// SetSV__Listeners updates property "Listeners".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewaySpec) SetSV__Listeners(v ...AWS_AppMesh_VirtualGateway_VirtualGatewayListener) AWS_AppMesh_VirtualGateway_VirtualGatewaySpec {
	t.Listeners = cfz.SV(v...)
	return t
}

// Set__Logging updates property "Logging".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewaySpec) Set__Logging(v cfz.Expression[AWS_AppMesh_VirtualGateway_VirtualGatewayLogging]) AWS_AppMesh_VirtualGateway_VirtualGatewaySpec {
	t.Logging = v
	return t
}

// SetV__Logging updates property "Logging".
func (t AWS_AppMesh_VirtualGateway_VirtualGatewaySpec) SetV__Logging(v AWS_AppMesh_VirtualGateway_VirtualGatewayLogging) AWS_AppMesh_VirtualGateway_VirtualGatewaySpec {
	t.Logging = cfz.V(v)
	return t
}
