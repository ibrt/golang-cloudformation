// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_databrew

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataBrew_Dataset_Input__Type is the CloudFormation type for AWS::DataBrew::Dataset.Input.
	AWS_DataBrew_Dataset_Input__Type = "AWS::DataBrew::Dataset.Input"
)

var (
	// AWS_DataBrew_Dataset_Input__PropertiesMap reports all the CloudFormation properties for AWS::DataBrew::Dataset.Input.
	AWS_DataBrew_Dataset_Input__PropertiesMap = struct {
		DataCatalogInputDefinition string
		DatabaseInputDefinition    string
		Metadata                   string
		S3InputDefinition          string
	}{
		DataCatalogInputDefinition: "DataCatalogInputDefinition",
		DatabaseInputDefinition:    "DatabaseInputDefinition",
		Metadata:                   "Metadata",
		S3InputDefinition:          "S3InputDefinition",
	}

	// AWS_DataBrew_Dataset_Input__PropertiesSlice reports all the CloudFormation properties for AWS::DataBrew::Dataset.Input.
	AWS_DataBrew_Dataset_Input__PropertiesSlice = []string{
		AWS_DataBrew_Dataset_Input__PropertiesMap.DataCatalogInputDefinition,
		AWS_DataBrew_Dataset_Input__PropertiesMap.DatabaseInputDefinition,
		AWS_DataBrew_Dataset_Input__PropertiesMap.Metadata,
		AWS_DataBrew_Dataset_Input__PropertiesMap.S3InputDefinition,
	}
)

// AWS_DataBrew_Dataset_Input is a binding for AWS::DataBrew::Dataset.Input.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html
type AWS_DataBrew_Dataset_Input struct {
	// DataCatalogInputDefinition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-datacataloginputdefinition
	DataCatalogInputDefinition cfz.Expression[AWS_DataBrew_Dataset_DataCatalogInputDefinition] `json:"DataCatalogInputDefinition,omitempty"`

	// DatabaseInputDefinition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-databaseinputdefinition
	DatabaseInputDefinition cfz.Expression[AWS_DataBrew_Dataset_DatabaseInputDefinition] `json:"DatabaseInputDefinition,omitempty"`

	// Metadata is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-metadata
	Metadata cfz.Expression[AWS_DataBrew_Dataset_Metadata] `json:"Metadata,omitempty"`

	// S3InputDefinition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-s3inputdefinition
	S3InputDefinition cfz.Expression[AWS_DataBrew_Dataset_S3Location] `json:"S3InputDefinition,omitempty"`
}

// New__AWS_DataBrew_Dataset_Input initializes a new AWS_DataBrew_Dataset_Input.
func New__AWS_DataBrew_Dataset_Input() AWS_DataBrew_Dataset_Input {
	return AWS_DataBrew_Dataset_Input{}
}

// GetType returns the CloudFormation type.
func (AWS_DataBrew_Dataset_Input) GetType() string {
	return AWS_DataBrew_Dataset_Input__Type
}

// Set__DataCatalogInputDefinition updates property "DataCatalogInputDefinition".
func (t AWS_DataBrew_Dataset_Input) Set__DataCatalogInputDefinition(v cfz.Expression[AWS_DataBrew_Dataset_DataCatalogInputDefinition]) AWS_DataBrew_Dataset_Input {
	t.DataCatalogInputDefinition = v
	return t
}

// SetV__DataCatalogInputDefinition updates property "DataCatalogInputDefinition".
func (t AWS_DataBrew_Dataset_Input) SetV__DataCatalogInputDefinition(v AWS_DataBrew_Dataset_DataCatalogInputDefinition) AWS_DataBrew_Dataset_Input {
	t.DataCatalogInputDefinition = cfz.V(v)
	return t
}

// Set__DatabaseInputDefinition updates property "DatabaseInputDefinition".
func (t AWS_DataBrew_Dataset_Input) Set__DatabaseInputDefinition(v cfz.Expression[AWS_DataBrew_Dataset_DatabaseInputDefinition]) AWS_DataBrew_Dataset_Input {
	t.DatabaseInputDefinition = v
	return t
}

// SetV__DatabaseInputDefinition updates property "DatabaseInputDefinition".
func (t AWS_DataBrew_Dataset_Input) SetV__DatabaseInputDefinition(v AWS_DataBrew_Dataset_DatabaseInputDefinition) AWS_DataBrew_Dataset_Input {
	t.DatabaseInputDefinition = cfz.V(v)
	return t
}

// Set__Metadata updates property "Metadata".
func (t AWS_DataBrew_Dataset_Input) Set__Metadata(v cfz.Expression[AWS_DataBrew_Dataset_Metadata]) AWS_DataBrew_Dataset_Input {
	t.Metadata = v
	return t
}

// SetV__Metadata updates property "Metadata".
func (t AWS_DataBrew_Dataset_Input) SetV__Metadata(v AWS_DataBrew_Dataset_Metadata) AWS_DataBrew_Dataset_Input {
	t.Metadata = cfz.V(v)
	return t
}

// Set__S3InputDefinition updates property "S3InputDefinition".
func (t AWS_DataBrew_Dataset_Input) Set__S3InputDefinition(v cfz.Expression[AWS_DataBrew_Dataset_S3Location]) AWS_DataBrew_Dataset_Input {
	t.S3InputDefinition = v
	return t
}

// SetV__S3InputDefinition updates property "S3InputDefinition".
func (t AWS_DataBrew_Dataset_Input) SetV__S3InputDefinition(v AWS_DataBrew_Dataset_S3Location) AWS_DataBrew_Dataset_Input {
	t.S3InputDefinition = cfz.V(v)
	return t
}
