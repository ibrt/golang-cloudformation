// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_FeatureGroup_DataCatalogConfig__Type is the CloudFormation type for AWS::SageMaker::FeatureGroup.DataCatalogConfig.
	AWS_SageMaker_FeatureGroup_DataCatalogConfig__Type = "AWS::SageMaker::FeatureGroup.DataCatalogConfig"
)

var (
	// AWS_SageMaker_FeatureGroup_DataCatalogConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::FeatureGroup.DataCatalogConfig.
	AWS_SageMaker_FeatureGroup_DataCatalogConfig__PropertiesMap = struct {
		Catalog   string
		Database  string
		TableName string
	}{
		Catalog:   "Catalog",
		Database:  "Database",
		TableName: "TableName",
	}

	// AWS_SageMaker_FeatureGroup_DataCatalogConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::FeatureGroup.DataCatalogConfig.
	AWS_SageMaker_FeatureGroup_DataCatalogConfig__PropertiesSlice = []string{
		AWS_SageMaker_FeatureGroup_DataCatalogConfig__PropertiesMap.Catalog,
		AWS_SageMaker_FeatureGroup_DataCatalogConfig__PropertiesMap.Database,
		AWS_SageMaker_FeatureGroup_DataCatalogConfig__PropertiesMap.TableName,
	}
)

// AWS_SageMaker_FeatureGroup_DataCatalogConfig is a binding for AWS::SageMaker::FeatureGroup.DataCatalogConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-datacatalogconfig.html
type AWS_SageMaker_FeatureGroup_DataCatalogConfig struct {
	// Catalog is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-datacatalogconfig.html#cfn-sagemaker-featuregroup-datacatalogconfig-catalog
	Catalog cfz.Expression[string] `json:"Catalog,omitempty"`

	// Database is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-datacatalogconfig.html#cfn-sagemaker-featuregroup-datacatalogconfig-database
	Database cfz.Expression[string] `json:"Database,omitempty"`

	// TableName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-datacatalogconfig.html#cfn-sagemaker-featuregroup-datacatalogconfig-tablename
	TableName cfz.Expression[string] `json:"TableName,omitempty"`
}

// New__AWS_SageMaker_FeatureGroup_DataCatalogConfig initializes a new AWS_SageMaker_FeatureGroup_DataCatalogConfig.
func New__AWS_SageMaker_FeatureGroup_DataCatalogConfig() AWS_SageMaker_FeatureGroup_DataCatalogConfig {
	return AWS_SageMaker_FeatureGroup_DataCatalogConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_FeatureGroup_DataCatalogConfig) GetType() string {
	return AWS_SageMaker_FeatureGroup_DataCatalogConfig__Type
}

// Set__Catalog updates property "Catalog".
func (t AWS_SageMaker_FeatureGroup_DataCatalogConfig) Set__Catalog(v cfz.Expression[string]) AWS_SageMaker_FeatureGroup_DataCatalogConfig {
	t.Catalog = v
	return t
}

// SetV__Catalog updates property "Catalog".
func (t AWS_SageMaker_FeatureGroup_DataCatalogConfig) SetV__Catalog(v string) AWS_SageMaker_FeatureGroup_DataCatalogConfig {
	t.Catalog = cfz.V(v)
	return t
}

// Set__Database updates property "Database".
func (t AWS_SageMaker_FeatureGroup_DataCatalogConfig) Set__Database(v cfz.Expression[string]) AWS_SageMaker_FeatureGroup_DataCatalogConfig {
	t.Database = v
	return t
}

// SetV__Database updates property "Database".
func (t AWS_SageMaker_FeatureGroup_DataCatalogConfig) SetV__Database(v string) AWS_SageMaker_FeatureGroup_DataCatalogConfig {
	t.Database = cfz.V(v)
	return t
}

// Set__TableName updates property "TableName".
func (t AWS_SageMaker_FeatureGroup_DataCatalogConfig) Set__TableName(v cfz.Expression[string]) AWS_SageMaker_FeatureGroup_DataCatalogConfig {
	t.TableName = v
	return t
}

// SetV__TableName updates property "TableName".
func (t AWS_SageMaker_FeatureGroup_DataCatalogConfig) SetV__TableName(v string) AWS_SageMaker_FeatureGroup_DataCatalogConfig {
	t.TableName = cfz.V(v)
	return t
}
