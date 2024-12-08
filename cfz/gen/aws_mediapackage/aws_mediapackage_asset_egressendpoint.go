// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediapackage

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaPackage_Asset_EgressEndpoint__Type is the CloudFormation type for AWS::MediaPackage::Asset.EgressEndpoint.
	AWS_MediaPackage_Asset_EgressEndpoint__Type = "AWS::MediaPackage::Asset.EgressEndpoint"
)

var (
	// AWS_MediaPackage_Asset_EgressEndpoint__PropertiesMap reports all the CloudFormation properties for AWS::MediaPackage::Asset.EgressEndpoint.
	AWS_MediaPackage_Asset_EgressEndpoint__PropertiesMap = struct {
		PackagingConfigurationId string
		Url                      string
	}{
		PackagingConfigurationId: "PackagingConfigurationId",
		Url:                      "Url",
	}

	// AWS_MediaPackage_Asset_EgressEndpoint__PropertiesSlice reports all the CloudFormation properties for AWS::MediaPackage::Asset.EgressEndpoint.
	AWS_MediaPackage_Asset_EgressEndpoint__PropertiesSlice = []string{
		AWS_MediaPackage_Asset_EgressEndpoint__PropertiesMap.PackagingConfigurationId,
		AWS_MediaPackage_Asset_EgressEndpoint__PropertiesMap.Url,
	}
)

// AWS_MediaPackage_Asset_EgressEndpoint is a binding for AWS::MediaPackage::Asset.EgressEndpoint.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-asset-egressendpoint.html
type AWS_MediaPackage_Asset_EgressEndpoint struct {
	// PackagingConfigurationId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-asset-egressendpoint.html#cfn-mediapackage-asset-egressendpoint-packagingconfigurationid
	PackagingConfigurationId cfz.Expression[string] `json:"PackagingConfigurationId,omitempty"`

	// Url is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-asset-egressendpoint.html#cfn-mediapackage-asset-egressendpoint-url
	Url cfz.Expression[string] `json:"Url,omitempty"`
}

// New__AWS_MediaPackage_Asset_EgressEndpoint initializes a new AWS_MediaPackage_Asset_EgressEndpoint.
func New__AWS_MediaPackage_Asset_EgressEndpoint() AWS_MediaPackage_Asset_EgressEndpoint {
	return AWS_MediaPackage_Asset_EgressEndpoint{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaPackage_Asset_EgressEndpoint) GetType() string {
	return AWS_MediaPackage_Asset_EgressEndpoint__Type
}

// Set__PackagingConfigurationId updates property "PackagingConfigurationId".
func (t AWS_MediaPackage_Asset_EgressEndpoint) Set__PackagingConfigurationId(v cfz.Expression[string]) AWS_MediaPackage_Asset_EgressEndpoint {
	t.PackagingConfigurationId = v
	return t
}

// SetV__PackagingConfigurationId updates property "PackagingConfigurationId".
func (t AWS_MediaPackage_Asset_EgressEndpoint) SetV__PackagingConfigurationId(v string) AWS_MediaPackage_Asset_EgressEndpoint {
	t.PackagingConfigurationId = cfz.V(v)
	return t
}

// Set__Url updates property "Url".
func (t AWS_MediaPackage_Asset_EgressEndpoint) Set__Url(v cfz.Expression[string]) AWS_MediaPackage_Asset_EgressEndpoint {
	t.Url = v
	return t
}

// SetV__Url updates property "Url".
func (t AWS_MediaPackage_Asset_EgressEndpoint) SetV__Url(v string) AWS_MediaPackage_Asset_EgressEndpoint {
	t.Url = cfz.V(v)
	return t
}
