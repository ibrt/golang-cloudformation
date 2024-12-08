// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lightsail

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lightsail_Distribution_CacheBehaviorPerPath__Type is the CloudFormation type for AWS::Lightsail::Distribution.CacheBehaviorPerPath.
	AWS_Lightsail_Distribution_CacheBehaviorPerPath__Type = "AWS::Lightsail::Distribution.CacheBehaviorPerPath"
)

var (
	// AWS_Lightsail_Distribution_CacheBehaviorPerPath__PropertiesMap reports all the CloudFormation properties for AWS::Lightsail::Distribution.CacheBehaviorPerPath.
	AWS_Lightsail_Distribution_CacheBehaviorPerPath__PropertiesMap = struct {
		Behavior string
		Path     string
	}{
		Behavior: "Behavior",
		Path:     "Path",
	}

	// AWS_Lightsail_Distribution_CacheBehaviorPerPath__PropertiesSlice reports all the CloudFormation properties for AWS::Lightsail::Distribution.CacheBehaviorPerPath.
	AWS_Lightsail_Distribution_CacheBehaviorPerPath__PropertiesSlice = []string{
		AWS_Lightsail_Distribution_CacheBehaviorPerPath__PropertiesMap.Behavior,
		AWS_Lightsail_Distribution_CacheBehaviorPerPath__PropertiesMap.Path,
	}
)

// AWS_Lightsail_Distribution_CacheBehaviorPerPath is a binding for AWS::Lightsail::Distribution.CacheBehaviorPerPath.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-cachebehaviorperpath.html
type AWS_Lightsail_Distribution_CacheBehaviorPerPath struct {
	// Behavior is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-cachebehaviorperpath.html#cfn-lightsail-distribution-cachebehaviorperpath-behavior
	Behavior cfz.Expression[string] `json:"Behavior,omitempty"`

	// Path is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-cachebehaviorperpath.html#cfn-lightsail-distribution-cachebehaviorperpath-path
	Path cfz.Expression[string] `json:"Path,omitempty"`
}

// New__AWS_Lightsail_Distribution_CacheBehaviorPerPath initializes a new AWS_Lightsail_Distribution_CacheBehaviorPerPath.
func New__AWS_Lightsail_Distribution_CacheBehaviorPerPath() AWS_Lightsail_Distribution_CacheBehaviorPerPath {
	return AWS_Lightsail_Distribution_CacheBehaviorPerPath{}
}

// GetType returns the CloudFormation type.
func (AWS_Lightsail_Distribution_CacheBehaviorPerPath) GetType() string {
	return AWS_Lightsail_Distribution_CacheBehaviorPerPath__Type
}

// Set__Behavior updates property "Behavior".
func (t AWS_Lightsail_Distribution_CacheBehaviorPerPath) Set__Behavior(v cfz.Expression[string]) AWS_Lightsail_Distribution_CacheBehaviorPerPath {
	t.Behavior = v
	return t
}

// SetV__Behavior updates property "Behavior".
func (t AWS_Lightsail_Distribution_CacheBehaviorPerPath) SetV__Behavior(v string) AWS_Lightsail_Distribution_CacheBehaviorPerPath {
	t.Behavior = cfz.V(v)
	return t
}

// Set__Path updates property "Path".
func (t AWS_Lightsail_Distribution_CacheBehaviorPerPath) Set__Path(v cfz.Expression[string]) AWS_Lightsail_Distribution_CacheBehaviorPerPath {
	t.Path = v
	return t
}

// SetV__Path updates property "Path".
func (t AWS_Lightsail_Distribution_CacheBehaviorPerPath) SetV__Path(v string) AWS_Lightsail_Distribution_CacheBehaviorPerPath {
	t.Path = cfz.V(v)
	return t
}
