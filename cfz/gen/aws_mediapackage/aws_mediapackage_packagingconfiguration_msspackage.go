// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediapackage

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaPackage_PackagingConfiguration_MssPackage__Type is the CloudFormation type for AWS::MediaPackage::PackagingConfiguration.MssPackage.
	AWS_MediaPackage_PackagingConfiguration_MssPackage__Type = "AWS::MediaPackage::PackagingConfiguration.MssPackage"
)

var (
	// AWS_MediaPackage_PackagingConfiguration_MssPackage__PropertiesMap reports all the CloudFormation properties for AWS::MediaPackage::PackagingConfiguration.MssPackage.
	AWS_MediaPackage_PackagingConfiguration_MssPackage__PropertiesMap = struct {
		Encryption             string
		MssManifests           string
		SegmentDurationSeconds string
	}{
		Encryption:             "Encryption",
		MssManifests:           "MssManifests",
		SegmentDurationSeconds: "SegmentDurationSeconds",
	}

	// AWS_MediaPackage_PackagingConfiguration_MssPackage__PropertiesSlice reports all the CloudFormation properties for AWS::MediaPackage::PackagingConfiguration.MssPackage.
	AWS_MediaPackage_PackagingConfiguration_MssPackage__PropertiesSlice = []string{
		AWS_MediaPackage_PackagingConfiguration_MssPackage__PropertiesMap.Encryption,
		AWS_MediaPackage_PackagingConfiguration_MssPackage__PropertiesMap.MssManifests,
		AWS_MediaPackage_PackagingConfiguration_MssPackage__PropertiesMap.SegmentDurationSeconds,
	}
)

// AWS_MediaPackage_PackagingConfiguration_MssPackage is a binding for AWS::MediaPackage::PackagingConfiguration.MssPackage.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-msspackage.html
type AWS_MediaPackage_PackagingConfiguration_MssPackage struct {
	// Encryption is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-msspackage.html#cfn-mediapackage-packagingconfiguration-msspackage-encryption
	Encryption cfz.Expression[AWS_MediaPackage_PackagingConfiguration_MssEncryption] `json:"Encryption,omitempty"`

	// MssManifests is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-msspackage.html#cfn-mediapackage-packagingconfiguration-msspackage-mssmanifests
	MssManifests cfz.ExpressionSlice[AWS_MediaPackage_PackagingConfiguration_MssManifest] `json:"MssManifests,omitempty"`

	// SegmentDurationSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-msspackage.html#cfn-mediapackage-packagingconfiguration-msspackage-segmentdurationseconds
	SegmentDurationSeconds cfz.Expression[int32] `json:"SegmentDurationSeconds,omitempty"`
}

// New__AWS_MediaPackage_PackagingConfiguration_MssPackage initializes a new AWS_MediaPackage_PackagingConfiguration_MssPackage.
func New__AWS_MediaPackage_PackagingConfiguration_MssPackage() AWS_MediaPackage_PackagingConfiguration_MssPackage {
	return AWS_MediaPackage_PackagingConfiguration_MssPackage{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaPackage_PackagingConfiguration_MssPackage) GetType() string {
	return AWS_MediaPackage_PackagingConfiguration_MssPackage__Type
}

// Set__Encryption updates property "Encryption".
func (t AWS_MediaPackage_PackagingConfiguration_MssPackage) Set__Encryption(v cfz.Expression[AWS_MediaPackage_PackagingConfiguration_MssEncryption]) AWS_MediaPackage_PackagingConfiguration_MssPackage {
	t.Encryption = v
	return t
}

// SetV__Encryption updates property "Encryption".
func (t AWS_MediaPackage_PackagingConfiguration_MssPackage) SetV__Encryption(v AWS_MediaPackage_PackagingConfiguration_MssEncryption) AWS_MediaPackage_PackagingConfiguration_MssPackage {
	t.Encryption = cfz.V(v)
	return t
}

// Set__MssManifests updates property "MssManifests".
func (t AWS_MediaPackage_PackagingConfiguration_MssPackage) Set__MssManifests(v cfz.ExpressionSlice[AWS_MediaPackage_PackagingConfiguration_MssManifest]) AWS_MediaPackage_PackagingConfiguration_MssPackage {
	t.MssManifests = v
	return t
}

// SetS__MssManifests updates property "MssManifests".
func (t AWS_MediaPackage_PackagingConfiguration_MssPackage) SetS__MssManifests(v ...cfz.Expression[AWS_MediaPackage_PackagingConfiguration_MssManifest]) AWS_MediaPackage_PackagingConfiguration_MssPackage {
	t.MssManifests = cfz.S(v...)
	return t
}

// SetSV__MssManifests updates property "MssManifests".
func (t AWS_MediaPackage_PackagingConfiguration_MssPackage) SetSV__MssManifests(v ...AWS_MediaPackage_PackagingConfiguration_MssManifest) AWS_MediaPackage_PackagingConfiguration_MssPackage {
	t.MssManifests = cfz.SV(v...)
	return t
}

// Set__SegmentDurationSeconds updates property "SegmentDurationSeconds".
func (t AWS_MediaPackage_PackagingConfiguration_MssPackage) Set__SegmentDurationSeconds(v cfz.Expression[int32]) AWS_MediaPackage_PackagingConfiguration_MssPackage {
	t.SegmentDurationSeconds = v
	return t
}

// SetV__SegmentDurationSeconds updates property "SegmentDurationSeconds".
func (t AWS_MediaPackage_PackagingConfiguration_MssPackage) SetV__SegmentDurationSeconds(v int32) AWS_MediaPackage_PackagingConfiguration_MssPackage {
	t.SegmentDurationSeconds = cfz.V(v)
	return t
}
