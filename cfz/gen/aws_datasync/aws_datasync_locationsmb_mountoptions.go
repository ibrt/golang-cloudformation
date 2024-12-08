// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datasync

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataSync_LocationSMB_MountOptions__Type is the CloudFormation type for AWS::DataSync::LocationSMB.MountOptions.
	AWS_DataSync_LocationSMB_MountOptions__Type = "AWS::DataSync::LocationSMB.MountOptions"
)

var (
	// AWS_DataSync_LocationSMB_MountOptions__PropertiesMap reports all the CloudFormation properties for AWS::DataSync::LocationSMB.MountOptions.
	AWS_DataSync_LocationSMB_MountOptions__PropertiesMap = struct {
		Version string
	}{
		Version: "Version",
	}

	// AWS_DataSync_LocationSMB_MountOptions__PropertiesSlice reports all the CloudFormation properties for AWS::DataSync::LocationSMB.MountOptions.
	AWS_DataSync_LocationSMB_MountOptions__PropertiesSlice = []string{
		AWS_DataSync_LocationSMB_MountOptions__PropertiesMap.Version,
	}
)

// AWS_DataSync_LocationSMB_MountOptions is a binding for AWS::DataSync::LocationSMB.MountOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationsmb-mountoptions.html
type AWS_DataSync_LocationSMB_MountOptions struct {
	// Version is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationsmb-mountoptions.html#cfn-datasync-locationsmb-mountoptions-version
	Version cfz.Expression[string] `json:"Version,omitempty"`
}

// New__AWS_DataSync_LocationSMB_MountOptions initializes a new AWS_DataSync_LocationSMB_MountOptions.
func New__AWS_DataSync_LocationSMB_MountOptions() AWS_DataSync_LocationSMB_MountOptions {
	return AWS_DataSync_LocationSMB_MountOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_DataSync_LocationSMB_MountOptions) GetType() string {
	return AWS_DataSync_LocationSMB_MountOptions__Type
}

// Set__Version updates property "Version".
func (t AWS_DataSync_LocationSMB_MountOptions) Set__Version(v cfz.Expression[string]) AWS_DataSync_LocationSMB_MountOptions {
	t.Version = v
	return t
}

// SetV__Version updates property "Version".
func (t AWS_DataSync_LocationSMB_MountOptions) SetV__Version(v string) AWS_DataSync_LocationSMB_MountOptions {
	t.Version = cfz.V(v)
	return t
}
