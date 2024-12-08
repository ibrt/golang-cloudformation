// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_databrew

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataBrew_Dataset_Metadata__Type is the CloudFormation type for AWS::DataBrew::Dataset.Metadata.
	AWS_DataBrew_Dataset_Metadata__Type = "AWS::DataBrew::Dataset.Metadata"
)

var (
	// AWS_DataBrew_Dataset_Metadata__PropertiesMap reports all the CloudFormation properties for AWS::DataBrew::Dataset.Metadata.
	AWS_DataBrew_Dataset_Metadata__PropertiesMap = struct {
		SourceArn string
	}{
		SourceArn: "SourceArn",
	}

	// AWS_DataBrew_Dataset_Metadata__PropertiesSlice reports all the CloudFormation properties for AWS::DataBrew::Dataset.Metadata.
	AWS_DataBrew_Dataset_Metadata__PropertiesSlice = []string{
		AWS_DataBrew_Dataset_Metadata__PropertiesMap.SourceArn,
	}
)

// AWS_DataBrew_Dataset_Metadata is a binding for AWS::DataBrew::Dataset.Metadata.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-metadata.html
type AWS_DataBrew_Dataset_Metadata struct {
	// SourceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-metadata.html#cfn-databrew-dataset-metadata-sourcearn
	SourceArn cfz.Expression[string] `json:"SourceArn,omitempty"`
}

// New__AWS_DataBrew_Dataset_Metadata initializes a new AWS_DataBrew_Dataset_Metadata.
func New__AWS_DataBrew_Dataset_Metadata() AWS_DataBrew_Dataset_Metadata {
	return AWS_DataBrew_Dataset_Metadata{}
}

// GetType returns the CloudFormation type.
func (AWS_DataBrew_Dataset_Metadata) GetType() string {
	return AWS_DataBrew_Dataset_Metadata__Type
}

// Set__SourceArn updates property "SourceArn".
func (t AWS_DataBrew_Dataset_Metadata) Set__SourceArn(v cfz.Expression[string]) AWS_DataBrew_Dataset_Metadata {
	t.SourceArn = v
	return t
}

// SetV__SourceArn updates property "SourceArn".
func (t AWS_DataBrew_Dataset_Metadata) SetV__SourceArn(v string) AWS_DataBrew_Dataset_Metadata {
	t.SourceArn = cfz.V(v)
	return t
}
