// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emr

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EMR_Cluster_HadoopJarStepConfig__Type is the CloudFormation type for AWS::EMR::Cluster.HadoopJarStepConfig.
	AWS_EMR_Cluster_HadoopJarStepConfig__Type = "AWS::EMR::Cluster.HadoopJarStepConfig"
)

var (
	// AWS_EMR_Cluster_HadoopJarStepConfig__PropertiesMap reports all the CloudFormation properties for AWS::EMR::Cluster.HadoopJarStepConfig.
	AWS_EMR_Cluster_HadoopJarStepConfig__PropertiesMap = struct {
		Args           string
		Jar            string
		MainClass      string
		StepProperties string
	}{
		Args:           "Args",
		Jar:            "Jar",
		MainClass:      "MainClass",
		StepProperties: "StepProperties",
	}

	// AWS_EMR_Cluster_HadoopJarStepConfig__PropertiesSlice reports all the CloudFormation properties for AWS::EMR::Cluster.HadoopJarStepConfig.
	AWS_EMR_Cluster_HadoopJarStepConfig__PropertiesSlice = []string{
		AWS_EMR_Cluster_HadoopJarStepConfig__PropertiesMap.Args,
		AWS_EMR_Cluster_HadoopJarStepConfig__PropertiesMap.Jar,
		AWS_EMR_Cluster_HadoopJarStepConfig__PropertiesMap.MainClass,
		AWS_EMR_Cluster_HadoopJarStepConfig__PropertiesMap.StepProperties,
	}
)

// AWS_EMR_Cluster_HadoopJarStepConfig is a binding for AWS::EMR::Cluster.HadoopJarStepConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-hadoopjarstepconfig.html
type AWS_EMR_Cluster_HadoopJarStepConfig struct {
	// Args is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-hadoopjarstepconfig.html#cfn-elasticmapreduce-cluster-hadoopjarstepconfig-args
	Args cfz.ExpressionSlice[string] `json:"Args,omitempty"`

	// Jar is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-hadoopjarstepconfig.html#cfn-elasticmapreduce-cluster-hadoopjarstepconfig-jar
	Jar cfz.Expression[string] `json:"Jar,omitempty"`

	// MainClass is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-hadoopjarstepconfig.html#cfn-elasticmapreduce-cluster-hadoopjarstepconfig-mainclass
	MainClass cfz.Expression[string] `json:"MainClass,omitempty"`

	// StepProperties is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-hadoopjarstepconfig.html#cfn-elasticmapreduce-cluster-hadoopjarstepconfig-stepproperties
	StepProperties cfz.ExpressionSlice[AWS_EMR_Cluster_KeyValue] `json:"StepProperties,omitempty"`
}

// New__AWS_EMR_Cluster_HadoopJarStepConfig initializes a new AWS_EMR_Cluster_HadoopJarStepConfig.
func New__AWS_EMR_Cluster_HadoopJarStepConfig() AWS_EMR_Cluster_HadoopJarStepConfig {
	return AWS_EMR_Cluster_HadoopJarStepConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_EMR_Cluster_HadoopJarStepConfig) GetType() string {
	return AWS_EMR_Cluster_HadoopJarStepConfig__Type
}

// Set__Args updates property "Args".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) Set__Args(v cfz.ExpressionSlice[string]) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.Args = v
	return t
}

// SetS__Args updates property "Args".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) SetS__Args(v ...cfz.Expression[string]) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.Args = cfz.S(v...)
	return t
}

// SetSV__Args updates property "Args".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) SetSV__Args(v ...string) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.Args = cfz.SV(v...)
	return t
}

// Set__Jar updates property "Jar".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) Set__Jar(v cfz.Expression[string]) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.Jar = v
	return t
}

// SetV__Jar updates property "Jar".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) SetV__Jar(v string) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.Jar = cfz.V(v)
	return t
}

// Set__MainClass updates property "MainClass".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) Set__MainClass(v cfz.Expression[string]) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.MainClass = v
	return t
}

// SetV__MainClass updates property "MainClass".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) SetV__MainClass(v string) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.MainClass = cfz.V(v)
	return t
}

// Set__StepProperties updates property "StepProperties".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) Set__StepProperties(v cfz.ExpressionSlice[AWS_EMR_Cluster_KeyValue]) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.StepProperties = v
	return t
}

// SetS__StepProperties updates property "StepProperties".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) SetS__StepProperties(v ...cfz.Expression[AWS_EMR_Cluster_KeyValue]) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.StepProperties = cfz.S(v...)
	return t
}

// SetSV__StepProperties updates property "StepProperties".
func (t AWS_EMR_Cluster_HadoopJarStepConfig) SetSV__StepProperties(v ...AWS_EMR_Cluster_KeyValue) AWS_EMR_Cluster_HadoopJarStepConfig {
	t.StepProperties = cfz.SV(v...)
	return t
}
