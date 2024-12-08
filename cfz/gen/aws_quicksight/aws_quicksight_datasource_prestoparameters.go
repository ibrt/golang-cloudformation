// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_DataSource_PrestoParameters__Type is the CloudFormation type for AWS::QuickSight::DataSource.PrestoParameters.
	AWS_QuickSight_DataSource_PrestoParameters__Type = "AWS::QuickSight::DataSource.PrestoParameters"
)

var (
	// AWS_QuickSight_DataSource_PrestoParameters__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::DataSource.PrestoParameters.
	AWS_QuickSight_DataSource_PrestoParameters__PropertiesMap = struct {
		Catalog string
		Host    string
		Port    string
	}{
		Catalog: "Catalog",
		Host:    "Host",
		Port:    "Port",
	}

	// AWS_QuickSight_DataSource_PrestoParameters__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::DataSource.PrestoParameters.
	AWS_QuickSight_DataSource_PrestoParameters__PropertiesSlice = []string{
		AWS_QuickSight_DataSource_PrestoParameters__PropertiesMap.Catalog,
		AWS_QuickSight_DataSource_PrestoParameters__PropertiesMap.Host,
		AWS_QuickSight_DataSource_PrestoParameters__PropertiesMap.Port,
	}
)

// AWS_QuickSight_DataSource_PrestoParameters is a binding for AWS::QuickSight::DataSource.PrestoParameters.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-prestoparameters.html
type AWS_QuickSight_DataSource_PrestoParameters struct {
	// Catalog is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-prestoparameters.html#cfn-quicksight-datasource-prestoparameters-catalog
	Catalog cfz.Expression[string] `json:"Catalog,omitempty"`

	// Host is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-prestoparameters.html#cfn-quicksight-datasource-prestoparameters-host
	Host cfz.Expression[string] `json:"Host,omitempty"`

	// Port is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-prestoparameters.html#cfn-quicksight-datasource-prestoparameters-port
	Port cfz.Expression[float64] `json:"Port,omitempty"`
}

// New__AWS_QuickSight_DataSource_PrestoParameters initializes a new AWS_QuickSight_DataSource_PrestoParameters.
func New__AWS_QuickSight_DataSource_PrestoParameters() AWS_QuickSight_DataSource_PrestoParameters {
	return AWS_QuickSight_DataSource_PrestoParameters{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_DataSource_PrestoParameters) GetType() string {
	return AWS_QuickSight_DataSource_PrestoParameters__Type
}

// Set__Catalog updates property "Catalog".
func (t AWS_QuickSight_DataSource_PrestoParameters) Set__Catalog(v cfz.Expression[string]) AWS_QuickSight_DataSource_PrestoParameters {
	t.Catalog = v
	return t
}

// SetV__Catalog updates property "Catalog".
func (t AWS_QuickSight_DataSource_PrestoParameters) SetV__Catalog(v string) AWS_QuickSight_DataSource_PrestoParameters {
	t.Catalog = cfz.V(v)
	return t
}

// Set__Host updates property "Host".
func (t AWS_QuickSight_DataSource_PrestoParameters) Set__Host(v cfz.Expression[string]) AWS_QuickSight_DataSource_PrestoParameters {
	t.Host = v
	return t
}

// SetV__Host updates property "Host".
func (t AWS_QuickSight_DataSource_PrestoParameters) SetV__Host(v string) AWS_QuickSight_DataSource_PrestoParameters {
	t.Host = cfz.V(v)
	return t
}

// Set__Port updates property "Port".
func (t AWS_QuickSight_DataSource_PrestoParameters) Set__Port(v cfz.Expression[float64]) AWS_QuickSight_DataSource_PrestoParameters {
	t.Port = v
	return t
}

// SetV__Port updates property "Port".
func (t AWS_QuickSight_DataSource_PrestoParameters) SetV__Port(v float64) AWS_QuickSight_DataSource_PrestoParameters {
	t.Port = cfz.V(v)
	return t
}
