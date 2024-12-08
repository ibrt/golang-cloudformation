// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cleanroomsml

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CleanRoomsML_TrainingDataset_Dataset__Type is the CloudFormation type for AWS::CleanRoomsML::TrainingDataset.Dataset.
	AWS_CleanRoomsML_TrainingDataset_Dataset__Type = "AWS::CleanRoomsML::TrainingDataset.Dataset"
)

var (
	// AWS_CleanRoomsML_TrainingDataset_Dataset__PropertiesMap reports all the CloudFormation properties for AWS::CleanRoomsML::TrainingDataset.Dataset.
	AWS_CleanRoomsML_TrainingDataset_Dataset__PropertiesMap = struct {
		InputConfig string
		Type        string
	}{
		InputConfig: "InputConfig",
		Type:        "Type",
	}

	// AWS_CleanRoomsML_TrainingDataset_Dataset__PropertiesSlice reports all the CloudFormation properties for AWS::CleanRoomsML::TrainingDataset.Dataset.
	AWS_CleanRoomsML_TrainingDataset_Dataset__PropertiesSlice = []string{
		AWS_CleanRoomsML_TrainingDataset_Dataset__PropertiesMap.InputConfig,
		AWS_CleanRoomsML_TrainingDataset_Dataset__PropertiesMap.Type,
	}
)

// AWS_CleanRoomsML_TrainingDataset_Dataset is a binding for AWS::CleanRoomsML::TrainingDataset.Dataset.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-dataset.html
type AWS_CleanRoomsML_TrainingDataset_Dataset struct {
	// InputConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-dataset.html#cfn-cleanroomsml-trainingdataset-dataset-inputconfig
	InputConfig cfz.Expression[AWS_CleanRoomsML_TrainingDataset_DatasetInputConfig] `json:"InputConfig,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-dataset.html#cfn-cleanroomsml-trainingdataset-dataset-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_CleanRoomsML_TrainingDataset_Dataset initializes a new AWS_CleanRoomsML_TrainingDataset_Dataset.
func New__AWS_CleanRoomsML_TrainingDataset_Dataset() AWS_CleanRoomsML_TrainingDataset_Dataset {
	return AWS_CleanRoomsML_TrainingDataset_Dataset{}
}

// GetType returns the CloudFormation type.
func (AWS_CleanRoomsML_TrainingDataset_Dataset) GetType() string {
	return AWS_CleanRoomsML_TrainingDataset_Dataset__Type
}

// Set__InputConfig updates property "InputConfig".
func (t AWS_CleanRoomsML_TrainingDataset_Dataset) Set__InputConfig(v cfz.Expression[AWS_CleanRoomsML_TrainingDataset_DatasetInputConfig]) AWS_CleanRoomsML_TrainingDataset_Dataset {
	t.InputConfig = v
	return t
}

// SetV__InputConfig updates property "InputConfig".
func (t AWS_CleanRoomsML_TrainingDataset_Dataset) SetV__InputConfig(v AWS_CleanRoomsML_TrainingDataset_DatasetInputConfig) AWS_CleanRoomsML_TrainingDataset_Dataset {
	t.InputConfig = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_CleanRoomsML_TrainingDataset_Dataset) Set__Type(v cfz.Expression[string]) AWS_CleanRoomsML_TrainingDataset_Dataset {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_CleanRoomsML_TrainingDataset_Dataset) SetV__Type(v string) AWS_CleanRoomsML_TrainingDataset_Dataset {
	t.Type = cfz.V(v)
	return t
}
