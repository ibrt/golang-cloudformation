// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_gamelift

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GameLift_Build_StorageLocation__Type is the CloudFormation type for AWS::GameLift::Build.StorageLocation.
	AWS_GameLift_Build_StorageLocation__Type = "AWS::GameLift::Build.StorageLocation"
)

var (
	// AWS_GameLift_Build_StorageLocation__PropertiesMap reports all the CloudFormation properties for AWS::GameLift::Build.StorageLocation.
	AWS_GameLift_Build_StorageLocation__PropertiesMap = struct {
		Bucket        string
		Key           string
		ObjectVersion string
		RoleArn       string
	}{
		Bucket:        "Bucket",
		Key:           "Key",
		ObjectVersion: "ObjectVersion",
		RoleArn:       "RoleArn",
	}

	// AWS_GameLift_Build_StorageLocation__PropertiesSlice reports all the CloudFormation properties for AWS::GameLift::Build.StorageLocation.
	AWS_GameLift_Build_StorageLocation__PropertiesSlice = []string{
		AWS_GameLift_Build_StorageLocation__PropertiesMap.Bucket,
		AWS_GameLift_Build_StorageLocation__PropertiesMap.Key,
		AWS_GameLift_Build_StorageLocation__PropertiesMap.ObjectVersion,
		AWS_GameLift_Build_StorageLocation__PropertiesMap.RoleArn,
	}
)

// AWS_GameLift_Build_StorageLocation is a binding for AWS::GameLift::Build.StorageLocation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html
type AWS_GameLift_Build_StorageLocation struct {
	// Bucket is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storagelocation-bucket
	Bucket cfz.Expression[string] `json:"Bucket,omitempty"`

	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storagelocation-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// ObjectVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storagelocation-objectversion
	ObjectVersion cfz.Expression[string] `json:"ObjectVersion,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storagelocation-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`
}

// New__AWS_GameLift_Build_StorageLocation initializes a new AWS_GameLift_Build_StorageLocation.
func New__AWS_GameLift_Build_StorageLocation() AWS_GameLift_Build_StorageLocation {
	return AWS_GameLift_Build_StorageLocation{}
}

// GetType returns the CloudFormation type.
func (AWS_GameLift_Build_StorageLocation) GetType() string {
	return AWS_GameLift_Build_StorageLocation__Type
}

// Set__Bucket updates property "Bucket".
func (t AWS_GameLift_Build_StorageLocation) Set__Bucket(v cfz.Expression[string]) AWS_GameLift_Build_StorageLocation {
	t.Bucket = v
	return t
}

// SetV__Bucket updates property "Bucket".
func (t AWS_GameLift_Build_StorageLocation) SetV__Bucket(v string) AWS_GameLift_Build_StorageLocation {
	t.Bucket = cfz.V(v)
	return t
}

// Set__Key updates property "Key".
func (t AWS_GameLift_Build_StorageLocation) Set__Key(v cfz.Expression[string]) AWS_GameLift_Build_StorageLocation {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_GameLift_Build_StorageLocation) SetV__Key(v string) AWS_GameLift_Build_StorageLocation {
	t.Key = cfz.V(v)
	return t
}

// Set__ObjectVersion updates property "ObjectVersion".
func (t AWS_GameLift_Build_StorageLocation) Set__ObjectVersion(v cfz.Expression[string]) AWS_GameLift_Build_StorageLocation {
	t.ObjectVersion = v
	return t
}

// SetV__ObjectVersion updates property "ObjectVersion".
func (t AWS_GameLift_Build_StorageLocation) SetV__ObjectVersion(v string) AWS_GameLift_Build_StorageLocation {
	t.ObjectVersion = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t AWS_GameLift_Build_StorageLocation) Set__RoleArn(v cfz.Expression[string]) AWS_GameLift_Build_StorageLocation {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t AWS_GameLift_Build_StorageLocation) SetV__RoleArn(v string) AWS_GameLift_Build_StorageLocation {
	t.RoleArn = cfz.V(v)
	return t
}
