// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_aps

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_APS_Scraper_Source__Type is the CloudFormation type for AWS::APS::Scraper.Source.
	AWS_APS_Scraper_Source__Type = "AWS::APS::Scraper.Source"
)

var (
	// AWS_APS_Scraper_Source__PropertiesMap reports all the CloudFormation properties for AWS::APS::Scraper.Source.
	AWS_APS_Scraper_Source__PropertiesMap = struct {
		EksConfiguration string
	}{
		EksConfiguration: "EksConfiguration",
	}

	// AWS_APS_Scraper_Source__PropertiesSlice reports all the CloudFormation properties for AWS::APS::Scraper.Source.
	AWS_APS_Scraper_Source__PropertiesSlice = []string{
		AWS_APS_Scraper_Source__PropertiesMap.EksConfiguration,
	}
)

// AWS_APS_Scraper_Source is a binding for AWS::APS::Scraper.Source.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-source.html
type AWS_APS_Scraper_Source struct {
	// EksConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-source.html#cfn-aps-scraper-source-eksconfiguration
	EksConfiguration cfz.Expression[AWS_APS_Scraper_EksConfiguration] `json:"EksConfiguration,omitempty"`
}

// New__AWS_APS_Scraper_Source initializes a new AWS_APS_Scraper_Source.
func New__AWS_APS_Scraper_Source() AWS_APS_Scraper_Source {
	return AWS_APS_Scraper_Source{}
}

// GetType returns the CloudFormation type.
func (AWS_APS_Scraper_Source) GetType() string {
	return AWS_APS_Scraper_Source__Type
}

// Set__EksConfiguration updates property "EksConfiguration".
func (t AWS_APS_Scraper_Source) Set__EksConfiguration(v cfz.Expression[AWS_APS_Scraper_EksConfiguration]) AWS_APS_Scraper_Source {
	t.EksConfiguration = v
	return t
}

// SetV__EksConfiguration updates property "EksConfiguration".
func (t AWS_APS_Scraper_Source) SetV__EksConfiguration(v AWS_APS_Scraper_EksConfiguration) AWS_APS_Scraper_Source {
	t.EksConfiguration = cfz.V(v)
	return t
}
