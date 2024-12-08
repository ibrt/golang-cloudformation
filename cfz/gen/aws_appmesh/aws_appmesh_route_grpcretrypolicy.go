// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appmesh

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppMesh_Route_GrpcRetryPolicy__Type is the CloudFormation type for AWS::AppMesh::Route.GrpcRetryPolicy.
	AWS_AppMesh_Route_GrpcRetryPolicy__Type = "AWS::AppMesh::Route.GrpcRetryPolicy"
)

var (
	// AWS_AppMesh_Route_GrpcRetryPolicy__PropertiesMap reports all the CloudFormation properties for AWS::AppMesh::Route.GrpcRetryPolicy.
	AWS_AppMesh_Route_GrpcRetryPolicy__PropertiesMap = struct {
		GrpcRetryEvents string
		HttpRetryEvents string
		MaxRetries      string
		PerRetryTimeout string
		TcpRetryEvents  string
	}{
		GrpcRetryEvents: "GrpcRetryEvents",
		HttpRetryEvents: "HttpRetryEvents",
		MaxRetries:      "MaxRetries",
		PerRetryTimeout: "PerRetryTimeout",
		TcpRetryEvents:  "TcpRetryEvents",
	}

	// AWS_AppMesh_Route_GrpcRetryPolicy__PropertiesSlice reports all the CloudFormation properties for AWS::AppMesh::Route.GrpcRetryPolicy.
	AWS_AppMesh_Route_GrpcRetryPolicy__PropertiesSlice = []string{
		AWS_AppMesh_Route_GrpcRetryPolicy__PropertiesMap.GrpcRetryEvents,
		AWS_AppMesh_Route_GrpcRetryPolicy__PropertiesMap.HttpRetryEvents,
		AWS_AppMesh_Route_GrpcRetryPolicy__PropertiesMap.MaxRetries,
		AWS_AppMesh_Route_GrpcRetryPolicy__PropertiesMap.PerRetryTimeout,
		AWS_AppMesh_Route_GrpcRetryPolicy__PropertiesMap.TcpRetryEvents,
	}
)

// AWS_AppMesh_Route_GrpcRetryPolicy is a binding for AWS::AppMesh::Route.GrpcRetryPolicy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcretrypolicy.html
type AWS_AppMesh_Route_GrpcRetryPolicy struct {
	// GrpcRetryEvents is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcretrypolicy.html#cfn-appmesh-route-grpcretrypolicy-grpcretryevents
	GrpcRetryEvents cfz.ExpressionSlice[string] `json:"GrpcRetryEvents,omitempty"`

	// HttpRetryEvents is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcretrypolicy.html#cfn-appmesh-route-grpcretrypolicy-httpretryevents
	HttpRetryEvents cfz.ExpressionSlice[string] `json:"HttpRetryEvents,omitempty"`

	// MaxRetries is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcretrypolicy.html#cfn-appmesh-route-grpcretrypolicy-maxretries
	MaxRetries cfz.Expression[int32] `json:"MaxRetries,omitempty"`

	// PerRetryTimeout is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcretrypolicy.html#cfn-appmesh-route-grpcretrypolicy-perretrytimeout
	PerRetryTimeout cfz.Expression[AWS_AppMesh_Route_Duration] `json:"PerRetryTimeout,omitempty"`

	// TcpRetryEvents is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcretrypolicy.html#cfn-appmesh-route-grpcretrypolicy-tcpretryevents
	TcpRetryEvents cfz.ExpressionSlice[string] `json:"TcpRetryEvents,omitempty"`
}

// New__AWS_AppMesh_Route_GrpcRetryPolicy initializes a new AWS_AppMesh_Route_GrpcRetryPolicy.
func New__AWS_AppMesh_Route_GrpcRetryPolicy() AWS_AppMesh_Route_GrpcRetryPolicy {
	return AWS_AppMesh_Route_GrpcRetryPolicy{}
}

// GetType returns the CloudFormation type.
func (AWS_AppMesh_Route_GrpcRetryPolicy) GetType() string {
	return AWS_AppMesh_Route_GrpcRetryPolicy__Type
}

// Set__GrpcRetryEvents updates property "GrpcRetryEvents".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) Set__GrpcRetryEvents(v cfz.ExpressionSlice[string]) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.GrpcRetryEvents = v
	return t
}

// SetS__GrpcRetryEvents updates property "GrpcRetryEvents".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) SetS__GrpcRetryEvents(v ...cfz.Expression[string]) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.GrpcRetryEvents = cfz.S(v...)
	return t
}

// SetSV__GrpcRetryEvents updates property "GrpcRetryEvents".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) SetSV__GrpcRetryEvents(v ...string) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.GrpcRetryEvents = cfz.SV(v...)
	return t
}

// Set__HttpRetryEvents updates property "HttpRetryEvents".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) Set__HttpRetryEvents(v cfz.ExpressionSlice[string]) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.HttpRetryEvents = v
	return t
}

// SetS__HttpRetryEvents updates property "HttpRetryEvents".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) SetS__HttpRetryEvents(v ...cfz.Expression[string]) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.HttpRetryEvents = cfz.S(v...)
	return t
}

// SetSV__HttpRetryEvents updates property "HttpRetryEvents".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) SetSV__HttpRetryEvents(v ...string) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.HttpRetryEvents = cfz.SV(v...)
	return t
}

// Set__MaxRetries updates property "MaxRetries".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) Set__MaxRetries(v cfz.Expression[int32]) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.MaxRetries = v
	return t
}

// SetV__MaxRetries updates property "MaxRetries".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) SetV__MaxRetries(v int32) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.MaxRetries = cfz.V(v)
	return t
}

// Set__PerRetryTimeout updates property "PerRetryTimeout".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) Set__PerRetryTimeout(v cfz.Expression[AWS_AppMesh_Route_Duration]) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.PerRetryTimeout = v
	return t
}

// SetV__PerRetryTimeout updates property "PerRetryTimeout".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) SetV__PerRetryTimeout(v AWS_AppMesh_Route_Duration) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.PerRetryTimeout = cfz.V(v)
	return t
}

// Set__TcpRetryEvents updates property "TcpRetryEvents".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) Set__TcpRetryEvents(v cfz.ExpressionSlice[string]) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.TcpRetryEvents = v
	return t
}

// SetS__TcpRetryEvents updates property "TcpRetryEvents".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) SetS__TcpRetryEvents(v ...cfz.Expression[string]) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.TcpRetryEvents = cfz.S(v...)
	return t
}

// SetSV__TcpRetryEvents updates property "TcpRetryEvents".
func (t AWS_AppMesh_Route_GrpcRetryPolicy) SetSV__TcpRetryEvents(v ...string) AWS_AppMesh_Route_GrpcRetryPolicy {
	t.TcpRetryEvents = cfz.SV(v...)
	return t
}
