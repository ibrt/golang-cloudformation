// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_batch

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Batch_JobDefinition_EksVolume__Type is the CloudFormation type for AWS::Batch::JobDefinition.EksVolume.
	AWS_Batch_JobDefinition_EksVolume__Type = "AWS::Batch::JobDefinition.EksVolume"
)

var (
	// AWS_Batch_JobDefinition_EksVolume__PropertiesMap reports all the CloudFormation properties for AWS::Batch::JobDefinition.EksVolume.
	AWS_Batch_JobDefinition_EksVolume__PropertiesMap = struct {
		EmptyDir string
		HostPath string
		Name     string
		Secret   string
	}{
		EmptyDir: "EmptyDir",
		HostPath: "HostPath",
		Name:     "Name",
		Secret:   "Secret",
	}

	// AWS_Batch_JobDefinition_EksVolume__PropertiesSlice reports all the CloudFormation properties for AWS::Batch::JobDefinition.EksVolume.
	AWS_Batch_JobDefinition_EksVolume__PropertiesSlice = []string{
		AWS_Batch_JobDefinition_EksVolume__PropertiesMap.EmptyDir,
		AWS_Batch_JobDefinition_EksVolume__PropertiesMap.HostPath,
		AWS_Batch_JobDefinition_EksVolume__PropertiesMap.Name,
		AWS_Batch_JobDefinition_EksVolume__PropertiesMap.Secret,
	}
)

// AWS_Batch_JobDefinition_EksVolume is a binding for AWS::Batch::JobDefinition.EksVolume.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html
type AWS_Batch_JobDefinition_EksVolume struct {
	// EmptyDir is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html#cfn-batch-jobdefinition-eksvolume-emptydir
	EmptyDir cfz.Expression[AWS_Batch_JobDefinition_EksEmptyDir] `json:"EmptyDir,omitempty"`

	// HostPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html#cfn-batch-jobdefinition-eksvolume-hostpath
	HostPath cfz.Expression[AWS_Batch_JobDefinition_EksHostPath] `json:"HostPath,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html#cfn-batch-jobdefinition-eksvolume-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Secret is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html#cfn-batch-jobdefinition-eksvolume-secret
	Secret cfz.Expression[AWS_Batch_JobDefinition_EksSecret] `json:"Secret,omitempty"`
}

// New__AWS_Batch_JobDefinition_EksVolume initializes a new AWS_Batch_JobDefinition_EksVolume.
func New__AWS_Batch_JobDefinition_EksVolume() AWS_Batch_JobDefinition_EksVolume {
	return AWS_Batch_JobDefinition_EksVolume{}
}

// GetType returns the CloudFormation type.
func (AWS_Batch_JobDefinition_EksVolume) GetType() string {
	return AWS_Batch_JobDefinition_EksVolume__Type
}

// Set__EmptyDir updates property "EmptyDir".
func (t AWS_Batch_JobDefinition_EksVolume) Set__EmptyDir(v cfz.Expression[AWS_Batch_JobDefinition_EksEmptyDir]) AWS_Batch_JobDefinition_EksVolume {
	t.EmptyDir = v
	return t
}

// SetV__EmptyDir updates property "EmptyDir".
func (t AWS_Batch_JobDefinition_EksVolume) SetV__EmptyDir(v AWS_Batch_JobDefinition_EksEmptyDir) AWS_Batch_JobDefinition_EksVolume {
	t.EmptyDir = cfz.V(v)
	return t
}

// Set__HostPath updates property "HostPath".
func (t AWS_Batch_JobDefinition_EksVolume) Set__HostPath(v cfz.Expression[AWS_Batch_JobDefinition_EksHostPath]) AWS_Batch_JobDefinition_EksVolume {
	t.HostPath = v
	return t
}

// SetV__HostPath updates property "HostPath".
func (t AWS_Batch_JobDefinition_EksVolume) SetV__HostPath(v AWS_Batch_JobDefinition_EksHostPath) AWS_Batch_JobDefinition_EksVolume {
	t.HostPath = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_Batch_JobDefinition_EksVolume) Set__Name(v cfz.Expression[string]) AWS_Batch_JobDefinition_EksVolume {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_Batch_JobDefinition_EksVolume) SetV__Name(v string) AWS_Batch_JobDefinition_EksVolume {
	t.Name = cfz.V(v)
	return t
}

// Set__Secret updates property "Secret".
func (t AWS_Batch_JobDefinition_EksVolume) Set__Secret(v cfz.Expression[AWS_Batch_JobDefinition_EksSecret]) AWS_Batch_JobDefinition_EksVolume {
	t.Secret = v
	return t
}

// SetV__Secret updates property "Secret".
func (t AWS_Batch_JobDefinition_EksVolume) SetV__Secret(v AWS_Batch_JobDefinition_EksSecret) AWS_Batch_JobDefinition_EksVolume {
	t.Secret = cfz.V(v)
	return t
}
