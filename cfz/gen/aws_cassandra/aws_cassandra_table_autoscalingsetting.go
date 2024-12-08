// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cassandra

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Cassandra_Table_AutoScalingSetting__Type is the CloudFormation type for AWS::Cassandra::Table.AutoScalingSetting.
	AWS_Cassandra_Table_AutoScalingSetting__Type = "AWS::Cassandra::Table.AutoScalingSetting"
)

var (
	// AWS_Cassandra_Table_AutoScalingSetting__PropertiesMap reports all the CloudFormation properties for AWS::Cassandra::Table.AutoScalingSetting.
	AWS_Cassandra_Table_AutoScalingSetting__PropertiesMap = struct {
		AutoScalingDisabled string
		MaximumUnits        string
		MinimumUnits        string
		ScalingPolicy       string
	}{
		AutoScalingDisabled: "AutoScalingDisabled",
		MaximumUnits:        "MaximumUnits",
		MinimumUnits:        "MinimumUnits",
		ScalingPolicy:       "ScalingPolicy",
	}

	// AWS_Cassandra_Table_AutoScalingSetting__PropertiesSlice reports all the CloudFormation properties for AWS::Cassandra::Table.AutoScalingSetting.
	AWS_Cassandra_Table_AutoScalingSetting__PropertiesSlice = []string{
		AWS_Cassandra_Table_AutoScalingSetting__PropertiesMap.AutoScalingDisabled,
		AWS_Cassandra_Table_AutoScalingSetting__PropertiesMap.MaximumUnits,
		AWS_Cassandra_Table_AutoScalingSetting__PropertiesMap.MinimumUnits,
		AWS_Cassandra_Table_AutoScalingSetting__PropertiesMap.ScalingPolicy,
	}
)

// AWS_Cassandra_Table_AutoScalingSetting is a binding for AWS::Cassandra::Table.AutoScalingSetting.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html
type AWS_Cassandra_Table_AutoScalingSetting struct {
	// AutoScalingDisabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html#cfn-cassandra-table-autoscalingsetting-autoscalingdisabled
	AutoScalingDisabled cfz.Expression[bool] `json:"AutoScalingDisabled,omitempty"`

	// MaximumUnits is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html#cfn-cassandra-table-autoscalingsetting-maximumunits
	MaximumUnits cfz.Expression[int32] `json:"MaximumUnits,omitempty"`

	// MinimumUnits is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html#cfn-cassandra-table-autoscalingsetting-minimumunits
	MinimumUnits cfz.Expression[int32] `json:"MinimumUnits,omitempty"`

	// ScalingPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html#cfn-cassandra-table-autoscalingsetting-scalingpolicy
	ScalingPolicy cfz.Expression[AWS_Cassandra_Table_ScalingPolicy] `json:"ScalingPolicy,omitempty"`
}

// New__AWS_Cassandra_Table_AutoScalingSetting initializes a new AWS_Cassandra_Table_AutoScalingSetting.
func New__AWS_Cassandra_Table_AutoScalingSetting() AWS_Cassandra_Table_AutoScalingSetting {
	return AWS_Cassandra_Table_AutoScalingSetting{}
}

// GetType returns the CloudFormation type.
func (AWS_Cassandra_Table_AutoScalingSetting) GetType() string {
	return AWS_Cassandra_Table_AutoScalingSetting__Type
}

// Set__AutoScalingDisabled updates property "AutoScalingDisabled".
func (t AWS_Cassandra_Table_AutoScalingSetting) Set__AutoScalingDisabled(v cfz.Expression[bool]) AWS_Cassandra_Table_AutoScalingSetting {
	t.AutoScalingDisabled = v
	return t
}

// SetV__AutoScalingDisabled updates property "AutoScalingDisabled".
func (t AWS_Cassandra_Table_AutoScalingSetting) SetV__AutoScalingDisabled(v bool) AWS_Cassandra_Table_AutoScalingSetting {
	t.AutoScalingDisabled = cfz.V(v)
	return t
}

// Set__MaximumUnits updates property "MaximumUnits".
func (t AWS_Cassandra_Table_AutoScalingSetting) Set__MaximumUnits(v cfz.Expression[int32]) AWS_Cassandra_Table_AutoScalingSetting {
	t.MaximumUnits = v
	return t
}

// SetV__MaximumUnits updates property "MaximumUnits".
func (t AWS_Cassandra_Table_AutoScalingSetting) SetV__MaximumUnits(v int32) AWS_Cassandra_Table_AutoScalingSetting {
	t.MaximumUnits = cfz.V(v)
	return t
}

// Set__MinimumUnits updates property "MinimumUnits".
func (t AWS_Cassandra_Table_AutoScalingSetting) Set__MinimumUnits(v cfz.Expression[int32]) AWS_Cassandra_Table_AutoScalingSetting {
	t.MinimumUnits = v
	return t
}

// SetV__MinimumUnits updates property "MinimumUnits".
func (t AWS_Cassandra_Table_AutoScalingSetting) SetV__MinimumUnits(v int32) AWS_Cassandra_Table_AutoScalingSetting {
	t.MinimumUnits = cfz.V(v)
	return t
}

// Set__ScalingPolicy updates property "ScalingPolicy".
func (t AWS_Cassandra_Table_AutoScalingSetting) Set__ScalingPolicy(v cfz.Expression[AWS_Cassandra_Table_ScalingPolicy]) AWS_Cassandra_Table_AutoScalingSetting {
	t.ScalingPolicy = v
	return t
}

// SetV__ScalingPolicy updates property "ScalingPolicy".
func (t AWS_Cassandra_Table_AutoScalingSetting) SetV__ScalingPolicy(v AWS_Cassandra_Table_ScalingPolicy) AWS_Cassandra_Table_AutoScalingSetting {
	t.ScalingPolicy = cfz.V(v)
	return t
}
