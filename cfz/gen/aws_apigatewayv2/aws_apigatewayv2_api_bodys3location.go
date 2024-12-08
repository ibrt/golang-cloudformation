// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apigatewayv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ApiGatewayV2_Api_BodyS3Location__Type is the CloudFormation type for AWS::ApiGatewayV2::Api.BodyS3Location.
	AWS_ApiGatewayV2_Api_BodyS3Location__Type = "AWS::ApiGatewayV2::Api.BodyS3Location"
)

var (
	// AWS_ApiGatewayV2_Api_BodyS3Location__PropertiesMap reports all the CloudFormation properties for AWS::ApiGatewayV2::Api.BodyS3Location.
	AWS_ApiGatewayV2_Api_BodyS3Location__PropertiesMap = struct {
		Bucket  string
		Etag    string
		Key     string
		Version string
	}{
		Bucket:  "Bucket",
		Etag:    "Etag",
		Key:     "Key",
		Version: "Version",
	}

	// AWS_ApiGatewayV2_Api_BodyS3Location__PropertiesSlice reports all the CloudFormation properties for AWS::ApiGatewayV2::Api.BodyS3Location.
	AWS_ApiGatewayV2_Api_BodyS3Location__PropertiesSlice = []string{
		AWS_ApiGatewayV2_Api_BodyS3Location__PropertiesMap.Bucket,
		AWS_ApiGatewayV2_Api_BodyS3Location__PropertiesMap.Etag,
		AWS_ApiGatewayV2_Api_BodyS3Location__PropertiesMap.Key,
		AWS_ApiGatewayV2_Api_BodyS3Location__PropertiesMap.Version,
	}
)

// AWS_ApiGatewayV2_Api_BodyS3Location is a binding for AWS::ApiGatewayV2::Api.BodyS3Location.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html
type AWS_ApiGatewayV2_Api_BodyS3Location struct {
	// Bucket is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-bucket
	Bucket cfz.Expression[string] `json:"Bucket,omitempty"`

	// Etag is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-etag
	Etag cfz.Expression[string] `json:"Etag,omitempty"`

	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Version is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-bodys3location.html#cfn-apigatewayv2-api-bodys3location-version
	Version cfz.Expression[string] `json:"Version,omitempty"`
}

// New__AWS_ApiGatewayV2_Api_BodyS3Location initializes a new AWS_ApiGatewayV2_Api_BodyS3Location.
func New__AWS_ApiGatewayV2_Api_BodyS3Location() AWS_ApiGatewayV2_Api_BodyS3Location {
	return AWS_ApiGatewayV2_Api_BodyS3Location{}
}

// GetType returns the CloudFormation type.
func (AWS_ApiGatewayV2_Api_BodyS3Location) GetType() string {
	return AWS_ApiGatewayV2_Api_BodyS3Location__Type
}

// Set__Bucket updates property "Bucket".
func (t AWS_ApiGatewayV2_Api_BodyS3Location) Set__Bucket(v cfz.Expression[string]) AWS_ApiGatewayV2_Api_BodyS3Location {
	t.Bucket = v
	return t
}

// SetV__Bucket updates property "Bucket".
func (t AWS_ApiGatewayV2_Api_BodyS3Location) SetV__Bucket(v string) AWS_ApiGatewayV2_Api_BodyS3Location {
	t.Bucket = cfz.V(v)
	return t
}

// Set__Etag updates property "Etag".
func (t AWS_ApiGatewayV2_Api_BodyS3Location) Set__Etag(v cfz.Expression[string]) AWS_ApiGatewayV2_Api_BodyS3Location {
	t.Etag = v
	return t
}

// SetV__Etag updates property "Etag".
func (t AWS_ApiGatewayV2_Api_BodyS3Location) SetV__Etag(v string) AWS_ApiGatewayV2_Api_BodyS3Location {
	t.Etag = cfz.V(v)
	return t
}

// Set__Key updates property "Key".
func (t AWS_ApiGatewayV2_Api_BodyS3Location) Set__Key(v cfz.Expression[string]) AWS_ApiGatewayV2_Api_BodyS3Location {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_ApiGatewayV2_Api_BodyS3Location) SetV__Key(v string) AWS_ApiGatewayV2_Api_BodyS3Location {
	t.Key = cfz.V(v)
	return t
}

// Set__Version updates property "Version".
func (t AWS_ApiGatewayV2_Api_BodyS3Location) Set__Version(v cfz.Expression[string]) AWS_ApiGatewayV2_Api_BodyS3Location {
	t.Version = v
	return t
}

// SetV__Version updates property "Version".
func (t AWS_ApiGatewayV2_Api_BodyS3Location) SetV__Version(v string) AWS_ApiGatewayV2_Api_BodyS3Location {
	t.Version = cfz.V(v)
	return t
}
