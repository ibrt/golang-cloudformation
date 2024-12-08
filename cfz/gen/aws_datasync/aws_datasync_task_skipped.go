// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datasync

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataSync_Task_Skipped__Type is the CloudFormation type for AWS::DataSync::Task.Skipped.
	AWS_DataSync_Task_Skipped__Type = "AWS::DataSync::Task.Skipped"
)

var (
	// AWS_DataSync_Task_Skipped__PropertiesMap reports all the CloudFormation properties for AWS::DataSync::Task.Skipped.
	AWS_DataSync_Task_Skipped__PropertiesMap = struct {
		ReportLevel string
	}{
		ReportLevel: "ReportLevel",
	}

	// AWS_DataSync_Task_Skipped__PropertiesSlice reports all the CloudFormation properties for AWS::DataSync::Task.Skipped.
	AWS_DataSync_Task_Skipped__PropertiesSlice = []string{
		AWS_DataSync_Task_Skipped__PropertiesMap.ReportLevel,
	}
)

// AWS_DataSync_Task_Skipped is a binding for AWS::DataSync::Task.Skipped.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-skipped.html
type AWS_DataSync_Task_Skipped struct {
	// ReportLevel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-skipped.html#cfn-datasync-task-skipped-reportlevel
	ReportLevel cfz.Expression[string] `json:"ReportLevel,omitempty"`
}

// New__AWS_DataSync_Task_Skipped initializes a new AWS_DataSync_Task_Skipped.
func New__AWS_DataSync_Task_Skipped() AWS_DataSync_Task_Skipped {
	return AWS_DataSync_Task_Skipped{}
}

// GetType returns the CloudFormation type.
func (AWS_DataSync_Task_Skipped) GetType() string {
	return AWS_DataSync_Task_Skipped__Type
}

// Set__ReportLevel updates property "ReportLevel".
func (t AWS_DataSync_Task_Skipped) Set__ReportLevel(v cfz.Expression[string]) AWS_DataSync_Task_Skipped {
	t.ReportLevel = v
	return t
}

// SetV__ReportLevel updates property "ReportLevel".
func (t AWS_DataSync_Task_Skipped) SetV__ReportLevel(v string) AWS_DataSync_Task_Skipped {
	t.ReportLevel = cfz.V(v)
	return t
}
