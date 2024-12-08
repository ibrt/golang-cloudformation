// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_quicksight

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QuickSight_DataSource_StarburstParameters__Type is the CloudFormation type for AWS::QuickSight::DataSource.StarburstParameters.
	AWS_QuickSight_DataSource_StarburstParameters__Type = "AWS::QuickSight::DataSource.StarburstParameters"
)

var (
	// AWS_QuickSight_DataSource_StarburstParameters__PropertiesMap reports all the CloudFormation properties for AWS::QuickSight::DataSource.StarburstParameters.
	AWS_QuickSight_DataSource_StarburstParameters__PropertiesMap = struct {
		AuthenticationType        string
		Catalog                   string
		DatabaseAccessControlRole string
		Host                      string
		OAuthParameters           string
		Port                      string
		ProductType               string
	}{
		AuthenticationType:        "AuthenticationType",
		Catalog:                   "Catalog",
		DatabaseAccessControlRole: "DatabaseAccessControlRole",
		Host:                      "Host",
		OAuthParameters:           "OAuthParameters",
		Port:                      "Port",
		ProductType:               "ProductType",
	}

	// AWS_QuickSight_DataSource_StarburstParameters__PropertiesSlice reports all the CloudFormation properties for AWS::QuickSight::DataSource.StarburstParameters.
	AWS_QuickSight_DataSource_StarburstParameters__PropertiesSlice = []string{
		AWS_QuickSight_DataSource_StarburstParameters__PropertiesMap.AuthenticationType,
		AWS_QuickSight_DataSource_StarburstParameters__PropertiesMap.Catalog,
		AWS_QuickSight_DataSource_StarburstParameters__PropertiesMap.DatabaseAccessControlRole,
		AWS_QuickSight_DataSource_StarburstParameters__PropertiesMap.Host,
		AWS_QuickSight_DataSource_StarburstParameters__PropertiesMap.OAuthParameters,
		AWS_QuickSight_DataSource_StarburstParameters__PropertiesMap.Port,
		AWS_QuickSight_DataSource_StarburstParameters__PropertiesMap.ProductType,
	}
)

// AWS_QuickSight_DataSource_StarburstParameters is a binding for AWS::QuickSight::DataSource.StarburstParameters.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html
type AWS_QuickSight_DataSource_StarburstParameters struct {
	// AuthenticationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-authenticationtype
	AuthenticationType cfz.Expression[string] `json:"AuthenticationType,omitempty"`

	// Catalog is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-catalog
	Catalog cfz.Expression[string] `json:"Catalog,omitempty"`

	// DatabaseAccessControlRole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-databaseaccesscontrolrole
	DatabaseAccessControlRole cfz.Expression[string] `json:"DatabaseAccessControlRole,omitempty"`

	// Host is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-host
	Host cfz.Expression[string] `json:"Host,omitempty"`

	// OAuthParameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-oauthparameters
	OAuthParameters cfz.Expression[AWS_QuickSight_DataSource_OAuthParameters] `json:"OAuthParameters,omitempty"`

	// Port is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-port
	Port cfz.Expression[float64] `json:"Port,omitempty"`

	// ProductType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-producttype
	ProductType cfz.Expression[string] `json:"ProductType,omitempty"`
}

// New__AWS_QuickSight_DataSource_StarburstParameters initializes a new AWS_QuickSight_DataSource_StarburstParameters.
func New__AWS_QuickSight_DataSource_StarburstParameters() AWS_QuickSight_DataSource_StarburstParameters {
	return AWS_QuickSight_DataSource_StarburstParameters{}
}

// GetType returns the CloudFormation type.
func (AWS_QuickSight_DataSource_StarburstParameters) GetType() string {
	return AWS_QuickSight_DataSource_StarburstParameters__Type
}

// Set__AuthenticationType updates property "AuthenticationType".
func (t AWS_QuickSight_DataSource_StarburstParameters) Set__AuthenticationType(v cfz.Expression[string]) AWS_QuickSight_DataSource_StarburstParameters {
	t.AuthenticationType = v
	return t
}

// SetV__AuthenticationType updates property "AuthenticationType".
func (t AWS_QuickSight_DataSource_StarburstParameters) SetV__AuthenticationType(v string) AWS_QuickSight_DataSource_StarburstParameters {
	t.AuthenticationType = cfz.V(v)
	return t
}

// Set__Catalog updates property "Catalog".
func (t AWS_QuickSight_DataSource_StarburstParameters) Set__Catalog(v cfz.Expression[string]) AWS_QuickSight_DataSource_StarburstParameters {
	t.Catalog = v
	return t
}

// SetV__Catalog updates property "Catalog".
func (t AWS_QuickSight_DataSource_StarburstParameters) SetV__Catalog(v string) AWS_QuickSight_DataSource_StarburstParameters {
	t.Catalog = cfz.V(v)
	return t
}

// Set__DatabaseAccessControlRole updates property "DatabaseAccessControlRole".
func (t AWS_QuickSight_DataSource_StarburstParameters) Set__DatabaseAccessControlRole(v cfz.Expression[string]) AWS_QuickSight_DataSource_StarburstParameters {
	t.DatabaseAccessControlRole = v
	return t
}

// SetV__DatabaseAccessControlRole updates property "DatabaseAccessControlRole".
func (t AWS_QuickSight_DataSource_StarburstParameters) SetV__DatabaseAccessControlRole(v string) AWS_QuickSight_DataSource_StarburstParameters {
	t.DatabaseAccessControlRole = cfz.V(v)
	return t
}

// Set__Host updates property "Host".
func (t AWS_QuickSight_DataSource_StarburstParameters) Set__Host(v cfz.Expression[string]) AWS_QuickSight_DataSource_StarburstParameters {
	t.Host = v
	return t
}

// SetV__Host updates property "Host".
func (t AWS_QuickSight_DataSource_StarburstParameters) SetV__Host(v string) AWS_QuickSight_DataSource_StarburstParameters {
	t.Host = cfz.V(v)
	return t
}

// Set__OAuthParameters updates property "OAuthParameters".
func (t AWS_QuickSight_DataSource_StarburstParameters) Set__OAuthParameters(v cfz.Expression[AWS_QuickSight_DataSource_OAuthParameters]) AWS_QuickSight_DataSource_StarburstParameters {
	t.OAuthParameters = v
	return t
}

// SetV__OAuthParameters updates property "OAuthParameters".
func (t AWS_QuickSight_DataSource_StarburstParameters) SetV__OAuthParameters(v AWS_QuickSight_DataSource_OAuthParameters) AWS_QuickSight_DataSource_StarburstParameters {
	t.OAuthParameters = cfz.V(v)
	return t
}

// Set__Port updates property "Port".
func (t AWS_QuickSight_DataSource_StarburstParameters) Set__Port(v cfz.Expression[float64]) AWS_QuickSight_DataSource_StarburstParameters {
	t.Port = v
	return t
}

// SetV__Port updates property "Port".
func (t AWS_QuickSight_DataSource_StarburstParameters) SetV__Port(v float64) AWS_QuickSight_DataSource_StarburstParameters {
	t.Port = cfz.V(v)
	return t
}

// Set__ProductType updates property "ProductType".
func (t AWS_QuickSight_DataSource_StarburstParameters) Set__ProductType(v cfz.Expression[string]) AWS_QuickSight_DataSource_StarburstParameters {
	t.ProductType = v
	return t
}

// SetV__ProductType updates property "ProductType".
func (t AWS_QuickSight_DataSource_StarburstParameters) SetV__ProductType(v string) AWS_QuickSight_DataSource_StarburstParameters {
	t.ProductType = cfz.V(v)
	return t
}
