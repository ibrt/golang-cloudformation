// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_databrew

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataBrew_Project_Sample__Type is the CloudFormation type for AWS::DataBrew::Project.Sample.
	AWS_DataBrew_Project_Sample__Type = "AWS::DataBrew::Project.Sample"
)

var (
	// AWS_DataBrew_Project_Sample__PropertiesMap reports all the CloudFormation properties for AWS::DataBrew::Project.Sample.
	AWS_DataBrew_Project_Sample__PropertiesMap = struct {
		Size string
		Type string
	}{
		Size: "Size",
		Type: "Type",
	}

	// AWS_DataBrew_Project_Sample__PropertiesSlice reports all the CloudFormation properties for AWS::DataBrew::Project.Sample.
	AWS_DataBrew_Project_Sample__PropertiesSlice = []string{
		AWS_DataBrew_Project_Sample__PropertiesMap.Size,
		AWS_DataBrew_Project_Sample__PropertiesMap.Type,
	}
)

// AWS_DataBrew_Project_Sample is a binding for AWS::DataBrew::Project.Sample.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-project-sample.html
type AWS_DataBrew_Project_Sample struct {
	// Size is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-project-sample.html#cfn-databrew-project-sample-size
	Size cfz.Expression[int32] `json:"Size,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-project-sample.html#cfn-databrew-project-sample-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_DataBrew_Project_Sample initializes a new AWS_DataBrew_Project_Sample.
func New__AWS_DataBrew_Project_Sample() AWS_DataBrew_Project_Sample {
	return AWS_DataBrew_Project_Sample{}
}

// GetType returns the CloudFormation type.
func (AWS_DataBrew_Project_Sample) GetType() string {
	return AWS_DataBrew_Project_Sample__Type
}

// Set__Size updates property "Size".
func (t AWS_DataBrew_Project_Sample) Set__Size(v cfz.Expression[int32]) AWS_DataBrew_Project_Sample {
	t.Size = v
	return t
}

// SetV__Size updates property "Size".
func (t AWS_DataBrew_Project_Sample) SetV__Size(v int32) AWS_DataBrew_Project_Sample {
	t.Size = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_DataBrew_Project_Sample) Set__Type(v cfz.Expression[string]) AWS_DataBrew_Project_Sample {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_DataBrew_Project_Sample) SetV__Type(v string) AWS_DataBrew_Project_Sample {
	t.Type = cfz.V(v)
	return t
}
