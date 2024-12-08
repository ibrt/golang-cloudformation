// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_rds

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_RDS_DBInstance_DBInstanceRole__Type is the CloudFormation type for AWS::RDS::DBInstance.DBInstanceRole.
	AWS_RDS_DBInstance_DBInstanceRole__Type = "AWS::RDS::DBInstance.DBInstanceRole"
)

var (
	// AWS_RDS_DBInstance_DBInstanceRole__PropertiesMap reports all the CloudFormation properties for AWS::RDS::DBInstance.DBInstanceRole.
	AWS_RDS_DBInstance_DBInstanceRole__PropertiesMap = struct {
		FeatureName string
		RoleArn     string
	}{
		FeatureName: "FeatureName",
		RoleArn:     "RoleArn",
	}

	// AWS_RDS_DBInstance_DBInstanceRole__PropertiesSlice reports all the CloudFormation properties for AWS::RDS::DBInstance.DBInstanceRole.
	AWS_RDS_DBInstance_DBInstanceRole__PropertiesSlice = []string{
		AWS_RDS_DBInstance_DBInstanceRole__PropertiesMap.FeatureName,
		AWS_RDS_DBInstance_DBInstanceRole__PropertiesMap.RoleArn,
	}
)

// AWS_RDS_DBInstance_DBInstanceRole is a binding for AWS::RDS::DBInstance.DBInstanceRole.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-dbinstancerole.html
type AWS_RDS_DBInstance_DBInstanceRole struct {
	// FeatureName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-dbinstancerole.html#cfn-rds-dbinstance-dbinstancerole-featurename
	FeatureName cfz.Expression[string] `json:"FeatureName,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-dbinstancerole.html#cfn-rds-dbinstance-dbinstancerole-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`
}

// New__AWS_RDS_DBInstance_DBInstanceRole initializes a new AWS_RDS_DBInstance_DBInstanceRole.
func New__AWS_RDS_DBInstance_DBInstanceRole() AWS_RDS_DBInstance_DBInstanceRole {
	return AWS_RDS_DBInstance_DBInstanceRole{}
}

// GetType returns the CloudFormation type.
func (AWS_RDS_DBInstance_DBInstanceRole) GetType() string {
	return AWS_RDS_DBInstance_DBInstanceRole__Type
}

// Set__FeatureName updates property "FeatureName".
func (t AWS_RDS_DBInstance_DBInstanceRole) Set__FeatureName(v cfz.Expression[string]) AWS_RDS_DBInstance_DBInstanceRole {
	t.FeatureName = v
	return t
}

// SetV__FeatureName updates property "FeatureName".
func (t AWS_RDS_DBInstance_DBInstanceRole) SetV__FeatureName(v string) AWS_RDS_DBInstance_DBInstanceRole {
	t.FeatureName = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t AWS_RDS_DBInstance_DBInstanceRole) Set__RoleArn(v cfz.Expression[string]) AWS_RDS_DBInstance_DBInstanceRole {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t AWS_RDS_DBInstance_DBInstanceRole) SetV__RoleArn(v string) AWS_RDS_DBInstance_DBInstanceRole {
	t.RoleArn = cfz.V(v)
	return t
}
