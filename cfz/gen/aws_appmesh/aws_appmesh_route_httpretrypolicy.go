// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_Route_HttpRetryPolicy__Type is the CloudFormation type for AWS::AppMesh::Route.HttpRetryPolicy.
	AWS_AppMesh_Route_HttpRetryPolicy__Type = "AWS::AppMesh::Route.HttpRetryPolicy"
)

var (
	// AWS_AppMesh_Route_HttpRetryPolicy__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::Route.HttpRetryPolicy.
	AWS_AppMesh_Route_HttpRetryPolicy__PropertiesMap = struct {
		HttpRetryEvents string
		MaxRetries      string
		PerRetryTimeout string
		TcpRetryEvents  string
	}{
		HttpRetryEvents: "HttpRetryEvents",
		MaxRetries:      "MaxRetries",
		PerRetryTimeout: "PerRetryTimeout",
		TcpRetryEvents:  "TcpRetryEvents",
	}

	// AWS_AppMesh_Route_HttpRetryPolicy__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::Route.HttpRetryPolicy.
	AWS_AppMesh_Route_HttpRetryPolicy__PropertiesSlice = []string{
		AWS_AppMesh_Route_HttpRetryPolicy__PropertiesMap.HttpRetryEvents,
		AWS_AppMesh_Route_HttpRetryPolicy__PropertiesMap.MaxRetries,
		AWS_AppMesh_Route_HttpRetryPolicy__PropertiesMap.PerRetryTimeout,
		AWS_AppMesh_Route_HttpRetryPolicy__PropertiesMap.TcpRetryEvents,
	}
)

// AWS_AppMesh_Route_HttpRetryPolicy is a binding for AWS::AppMesh::Route.HttpRetryPolicy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html
type AWS_AppMesh_Route_HttpRetryPolicy struct {
	// HttpRetryEvents is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-httpretryevents
	HttpRetryEvents cfz.ExpressionSlice[string] `json:"HttpRetryEvents,omitempty"`

	// MaxRetries is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-maxretries
	MaxRetries cfz.Expression[int32] `json:"MaxRetries,omitempty"`

	// PerRetryTimeout is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-perretrytimeout
	PerRetryTimeout cfz.Expression[AWS_AppMesh_Route_Duration] `json:"PerRetryTimeout,omitempty"`

	// TcpRetryEvents is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-tcpretryevents
	TcpRetryEvents cfz.ExpressionSlice[string] `json:"TcpRetryEvents,omitempty"`
}

// New__AWS_AppMesh_Route_HttpRetryPolicy initializes a new AWS_AppMesh_Route_HttpRetryPolicy.
func New__AWS_AppMesh_Route_HttpRetryPolicy() AWS_AppMesh_Route_HttpRetryPolicy {
	return AWS_AppMesh_Route_HttpRetryPolicy{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_Route_HttpRetryPolicy) GetType() string {
	return AWS_AppMesh_Route_HttpRetryPolicy__Type
}

// Set__HttpRetryEvents updates property "HttpRetryEvents".
func (t AWS_AppMesh_Route_HttpRetryPolicy) Set__HttpRetryEvents(v cfz.ExpressionSlice[string]) AWS_AppMesh_Route_HttpRetryPolicy {
	t.HttpRetryEvents = v
	return t
}

// SetS__HttpRetryEvents updates property "HttpRetryEvents".
func (t AWS_AppMesh_Route_HttpRetryPolicy) SetS__HttpRetryEvents(v ...cfz.Expression[string]) AWS_AppMesh_Route_HttpRetryPolicy {
	t.HttpRetryEvents = cfz.S(v...)
	return t
}

// SetSV__HttpRetryEvents updates property "HttpRetryEvents".
func (t AWS_AppMesh_Route_HttpRetryPolicy) SetSV__HttpRetryEvents(v ...string) AWS_AppMesh_Route_HttpRetryPolicy {
	t.HttpRetryEvents = cfz.SV(v...)
	return t
}

// Set__MaxRetries updates property "MaxRetries".
func (t AWS_AppMesh_Route_HttpRetryPolicy) Set__MaxRetries(v cfz.Expression[int32]) AWS_AppMesh_Route_HttpRetryPolicy {
	t.MaxRetries = v
	return t
}

// SetV__MaxRetries updates property "MaxRetries".
func (t AWS_AppMesh_Route_HttpRetryPolicy) SetV__MaxRetries(v int32) AWS_AppMesh_Route_HttpRetryPolicy {
	t.MaxRetries = cfz.V(v)
	return t
}

// Set__PerRetryTimeout updates property "PerRetryTimeout".
func (t AWS_AppMesh_Route_HttpRetryPolicy) Set__PerRetryTimeout(v cfz.Expression[AWS_AppMesh_Route_Duration]) AWS_AppMesh_Route_HttpRetryPolicy {
	t.PerRetryTimeout = v
	return t
}

// SetV__PerRetryTimeout updates property "PerRetryTimeout".
func (t AWS_AppMesh_Route_HttpRetryPolicy) SetV__PerRetryTimeout(v AWS_AppMesh_Route_Duration) AWS_AppMesh_Route_HttpRetryPolicy {
	t.PerRetryTimeout = cfz.V(v)
	return t
}

// Set__TcpRetryEvents updates property "TcpRetryEvents".
func (t AWS_AppMesh_Route_HttpRetryPolicy) Set__TcpRetryEvents(v cfz.ExpressionSlice[string]) AWS_AppMesh_Route_HttpRetryPolicy {
	t.TcpRetryEvents = v
	return t
}

// SetS__TcpRetryEvents updates property "TcpRetryEvents".
func (t AWS_AppMesh_Route_HttpRetryPolicy) SetS__TcpRetryEvents(v ...cfz.Expression[string]) AWS_AppMesh_Route_HttpRetryPolicy {
	t.TcpRetryEvents = cfz.S(v...)
	return t
}

// SetSV__TcpRetryEvents updates property "TcpRetryEvents".
func (t AWS_AppMesh_Route_HttpRetryPolicy) SetSV__TcpRetryEvents(v ...string) AWS_AppMesh_Route_HttpRetryPolicy {
	t.TcpRetryEvents = cfz.SV(v...)
	return t
}
