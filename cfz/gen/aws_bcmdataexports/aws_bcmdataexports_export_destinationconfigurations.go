// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bcmdataexports

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_BCMDataExports_Export_DestinationConfigurations__Type is the CloudFormation type for AWS::BCMDataExports::Export.DestinationConfigurations.
	AWS_BCMDataExports_Export_DestinationConfigurations__Type = "AWS::BCMDataExports::Export.DestinationConfigurations"
)

var (
	// AWS_BCMDataExports_Export_DestinationConfigurations__PropertiesMap reports all the CloudFormation properties for AWS::BCMDataExports::Export.DestinationConfigurations.
	AWS_BCMDataExports_Export_DestinationConfigurations__PropertiesMap = struct {
		S3Destination string
	}{
		S3Destination: "S3Destination",
	}

	// AWS_BCMDataExports_Export_DestinationConfigurations__PropertiesSlice reports all the CloudFormation properties for AWS::BCMDataExports::Export.DestinationConfigurations.
	AWS_BCMDataExports_Export_DestinationConfigurations__PropertiesSlice = []string{
		AWS_BCMDataExports_Export_DestinationConfigurations__PropertiesMap.S3Destination,
	}
)

// AWS_BCMDataExports_Export_DestinationConfigurations is a binding for AWS::BCMDataExports::Export.DestinationConfigurations.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-destinationconfigurations.html
type AWS_BCMDataExports_Export_DestinationConfigurations struct {
	// S3Destination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-destinationconfigurations.html#cfn-bcmdataexports-export-destinationconfigurations-s3destination
	S3Destination cfz.Expression[AWS_BCMDataExports_Export_S3Destination] `json:"S3Destination,omitempty"`
}

// New__AWS_BCMDataExports_Export_DestinationConfigurations initializes a new AWS_BCMDataExports_Export_DestinationConfigurations.
func New__AWS_BCMDataExports_Export_DestinationConfigurations() AWS_BCMDataExports_Export_DestinationConfigurations {
	return AWS_BCMDataExports_Export_DestinationConfigurations{}
}

// GetType returns the CloudFormation type.
func (AWS_BCMDataExports_Export_DestinationConfigurations) GetType() string {
	return AWS_BCMDataExports_Export_DestinationConfigurations__Type
}

// Set__S3Destination updates property "S3Destination".
func (t AWS_BCMDataExports_Export_DestinationConfigurations) Set__S3Destination(v cfz.Expression[AWS_BCMDataExports_Export_S3Destination]) AWS_BCMDataExports_Export_DestinationConfigurations {
	t.S3Destination = v
	return t
}

// SetV__S3Destination updates property "S3Destination".
func (t AWS_BCMDataExports_Export_DestinationConfigurations) SetV__S3Destination(v AWS_BCMDataExports_Export_S3Destination) AWS_BCMDataExports_Export_DestinationConfigurations {
	t.S3Destination = cfz.V(v)
	return t
}
