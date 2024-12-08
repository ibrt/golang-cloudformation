// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_fsx

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FSx_Volume_OriginSnapshot__Type is the CloudFormation type for AWS::FSx::Volume.OriginSnapshot.
	AWS_FSx_Volume_OriginSnapshot__Type = "AWS::FSx::Volume.OriginSnapshot"
)

var (
	// AWS_FSx_Volume_OriginSnapshot__PropertiesMap reports all the CloudFormation properties for AWS::FSx::Volume.OriginSnapshot.
	AWS_FSx_Volume_OriginSnapshot__PropertiesMap = struct {
		CopyStrategy string
		SnapshotARN  string
	}{
		CopyStrategy: "CopyStrategy",
		SnapshotARN:  "SnapshotARN",
	}

	// AWS_FSx_Volume_OriginSnapshot__PropertiesSlice reports all the CloudFormation properties for AWS::FSx::Volume.OriginSnapshot.
	AWS_FSx_Volume_OriginSnapshot__PropertiesSlice = []string{
		AWS_FSx_Volume_OriginSnapshot__PropertiesMap.CopyStrategy,
		AWS_FSx_Volume_OriginSnapshot__PropertiesMap.SnapshotARN,
	}
)

// AWS_FSx_Volume_OriginSnapshot is a binding for AWS::FSx::Volume.OriginSnapshot.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-openzfsconfiguration-originsnapshot.html
type AWS_FSx_Volume_OriginSnapshot struct {
	// CopyStrategy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-openzfsconfiguration-originsnapshot.html#cfn-fsx-volume-openzfsconfiguration-originsnapshot-copystrategy
	CopyStrategy cfz.Expression[string] `json:"CopyStrategy,omitempty"`

	// SnapshotARN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-openzfsconfiguration-originsnapshot.html#cfn-fsx-volume-openzfsconfiguration-originsnapshot-snapshotarn
	SnapshotARN cfz.Expression[string] `json:"SnapshotARN,omitempty"`
}

// New__AWS_FSx_Volume_OriginSnapshot initializes a new AWS_FSx_Volume_OriginSnapshot.
func New__AWS_FSx_Volume_OriginSnapshot() AWS_FSx_Volume_OriginSnapshot {
	return AWS_FSx_Volume_OriginSnapshot{}
}

// GetType returns the CloudFormation type.
func (AWS_FSx_Volume_OriginSnapshot) GetType() string {
	return AWS_FSx_Volume_OriginSnapshot__Type
}

// Set__CopyStrategy updates property "CopyStrategy".
func (t AWS_FSx_Volume_OriginSnapshot) Set__CopyStrategy(v cfz.Expression[string]) AWS_FSx_Volume_OriginSnapshot {
	t.CopyStrategy = v
	return t
}

// SetV__CopyStrategy updates property "CopyStrategy".
func (t AWS_FSx_Volume_OriginSnapshot) SetV__CopyStrategy(v string) AWS_FSx_Volume_OriginSnapshot {
	t.CopyStrategy = cfz.V(v)
	return t
}

// Set__SnapshotARN updates property "SnapshotARN".
func (t AWS_FSx_Volume_OriginSnapshot) Set__SnapshotARN(v cfz.Expression[string]) AWS_FSx_Volume_OriginSnapshot {
	t.SnapshotARN = v
	return t
}

// SetV__SnapshotARN updates property "SnapshotARN".
func (t AWS_FSx_Volume_OriginSnapshot) SetV__SnapshotARN(v string) AWS_FSx_Volume_OriginSnapshot {
	t.SnapshotARN = cfz.V(v)
	return t
}
