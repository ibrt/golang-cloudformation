// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_LaunchTemplate_ElasticGpuSpecification__Type is the CloudFormation type for AWS::EC2::LaunchTemplate.ElasticGpuSpecification.
	AWS_EC2_LaunchTemplate_ElasticGpuSpecification__Type = "AWS::EC2::LaunchTemplate.ElasticGpuSpecification"
)

var (
	// AWS_EC2_LaunchTemplate_ElasticGpuSpecification__PropertiesMap reports all the CloudFormation properties for AWS::EC2::LaunchTemplate.ElasticGpuSpecification.
	AWS_EC2_LaunchTemplate_ElasticGpuSpecification__PropertiesMap = struct {
		Type string
	}{
		Type: "Type",
	}

	// AWS_EC2_LaunchTemplate_ElasticGpuSpecification__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::LaunchTemplate.ElasticGpuSpecification.
	AWS_EC2_LaunchTemplate_ElasticGpuSpecification__PropertiesSlice = []string{
		AWS_EC2_LaunchTemplate_ElasticGpuSpecification__PropertiesMap.Type,
	}
)

// AWS_EC2_LaunchTemplate_ElasticGpuSpecification is a binding for AWS::EC2::LaunchTemplate.ElasticGpuSpecification.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-elasticgpuspecification.html
type AWS_EC2_LaunchTemplate_ElasticGpuSpecification struct {
	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-elasticgpuspecification.html#cfn-ec2-launchtemplate-elasticgpuspecification-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_EC2_LaunchTemplate_ElasticGpuSpecification initializes a new AWS_EC2_LaunchTemplate_ElasticGpuSpecification.
func New__AWS_EC2_LaunchTemplate_ElasticGpuSpecification() AWS_EC2_LaunchTemplate_ElasticGpuSpecification {
	return AWS_EC2_LaunchTemplate_ElasticGpuSpecification{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_LaunchTemplate_ElasticGpuSpecification) GetType() string {
	return AWS_EC2_LaunchTemplate_ElasticGpuSpecification__Type
}

// Set__Type updates property "Type".
func (t AWS_EC2_LaunchTemplate_ElasticGpuSpecification) Set__Type(v cfz.Expression[string]) AWS_EC2_LaunchTemplate_ElasticGpuSpecification {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_EC2_LaunchTemplate_ElasticGpuSpecification) SetV__Type(v string) AWS_EC2_LaunchTemplate_ElasticGpuSpecification {
	t.Type = cfz.V(v)
	return t
}
