// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediapackage

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaPackage_Channel_IngestEndpoint__Type is the CloudFormation type for AWS::MediaPackage::Channel.IngestEndpoint.
	AWS_MediaPackage_Channel_IngestEndpoint__Type = "AWS::MediaPackage::Channel.IngestEndpoint"
)

var (
	// AWS_MediaPackage_Channel_IngestEndpoint__PropertiesMap reports all the CloudFormation properties for AWS::MediaPackage::Channel.IngestEndpoint.
	AWS_MediaPackage_Channel_IngestEndpoint__PropertiesMap = struct {
		Id       string
		Password string
		Url      string
		Username string
	}{
		Id:       "Id",
		Password: "Password",
		Url:      "Url",
		Username: "Username",
	}

	// AWS_MediaPackage_Channel_IngestEndpoint__PropertiesSlice reports all the CloudFormation properties for AWS::MediaPackage::Channel.IngestEndpoint.
	AWS_MediaPackage_Channel_IngestEndpoint__PropertiesSlice = []string{
		AWS_MediaPackage_Channel_IngestEndpoint__PropertiesMap.Id,
		AWS_MediaPackage_Channel_IngestEndpoint__PropertiesMap.Password,
		AWS_MediaPackage_Channel_IngestEndpoint__PropertiesMap.Url,
		AWS_MediaPackage_Channel_IngestEndpoint__PropertiesMap.Username,
	}
)

// AWS_MediaPackage_Channel_IngestEndpoint is a binding for AWS::MediaPackage::Channel.IngestEndpoint.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html
type AWS_MediaPackage_Channel_IngestEndpoint struct {
	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html#cfn-mediapackage-channel-ingestendpoint-id
	Id cfz.Expression[string] `json:"Id,omitempty"`

	// Password is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html#cfn-mediapackage-channel-ingestendpoint-password
	Password cfz.Expression[string] `json:"Password,omitempty"`

	// Url is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html#cfn-mediapackage-channel-ingestendpoint-url
	Url cfz.Expression[string] `json:"Url,omitempty"`

	// Username is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html#cfn-mediapackage-channel-ingestendpoint-username
	Username cfz.Expression[string] `json:"Username,omitempty"`
}

// New__AWS_MediaPackage_Channel_IngestEndpoint initializes a new AWS_MediaPackage_Channel_IngestEndpoint.
func New__AWS_MediaPackage_Channel_IngestEndpoint() AWS_MediaPackage_Channel_IngestEndpoint {
	return AWS_MediaPackage_Channel_IngestEndpoint{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaPackage_Channel_IngestEndpoint) GetType() string {
	return AWS_MediaPackage_Channel_IngestEndpoint__Type
}

// Set__Id updates property "Id".
func (t AWS_MediaPackage_Channel_IngestEndpoint) Set__Id(v cfz.Expression[string]) AWS_MediaPackage_Channel_IngestEndpoint {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t AWS_MediaPackage_Channel_IngestEndpoint) SetV__Id(v string) AWS_MediaPackage_Channel_IngestEndpoint {
	t.Id = cfz.V(v)
	return t
}

// Set__Password updates property "Password".
func (t AWS_MediaPackage_Channel_IngestEndpoint) Set__Password(v cfz.Expression[string]) AWS_MediaPackage_Channel_IngestEndpoint {
	t.Password = v
	return t
}

// SetV__Password updates property "Password".
func (t AWS_MediaPackage_Channel_IngestEndpoint) SetV__Password(v string) AWS_MediaPackage_Channel_IngestEndpoint {
	t.Password = cfz.V(v)
	return t
}

// Set__Url updates property "Url".
func (t AWS_MediaPackage_Channel_IngestEndpoint) Set__Url(v cfz.Expression[string]) AWS_MediaPackage_Channel_IngestEndpoint {
	t.Url = v
	return t
}

// SetV__Url updates property "Url".
func (t AWS_MediaPackage_Channel_IngestEndpoint) SetV__Url(v string) AWS_MediaPackage_Channel_IngestEndpoint {
	t.Url = cfz.V(v)
	return t
}

// Set__Username updates property "Username".
func (t AWS_MediaPackage_Channel_IngestEndpoint) Set__Username(v cfz.Expression[string]) AWS_MediaPackage_Channel_IngestEndpoint {
	t.Username = v
	return t
}

// SetV__Username updates property "Username".
func (t AWS_MediaPackage_Channel_IngestEndpoint) SetV__Username(v string) AWS_MediaPackage_Channel_IngestEndpoint {
	t.Username = cfz.V(v)
	return t
}
