// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_qbusiness

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__Type is the CloudFormation type for AWS::QBusiness::DataSource.InlineDocumentEnrichmentConfiguration.
	AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__Type = "AWS::QBusiness::DataSource.InlineDocumentEnrichmentConfiguration"
)

var (
	// AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::QBusiness::DataSource.InlineDocumentEnrichmentConfiguration.
	AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__PropertiesMap = struct {
		Condition               string
		DocumentContentOperator string
		Target                  string
	}{
		Condition:               "Condition",
		DocumentContentOperator: "DocumentContentOperator",
		Target:                  "Target",
	}

	// AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::QBusiness::DataSource.InlineDocumentEnrichmentConfiguration.
	AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__PropertiesSlice = []string{
		AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__PropertiesMap.Condition,
		AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__PropertiesMap.DocumentContentOperator,
		AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__PropertiesMap.Target,
	}
)

// AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration is a binding for AWS::QBusiness::DataSource.InlineDocumentEnrichmentConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html
type AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration struct {
	// Condition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html#cfn-qbusiness-datasource-inlinedocumentenrichmentconfiguration-condition
	Condition cfz.Expression[AWS_QBusiness_DataSource_DocumentAttributeCondition] `json:"Condition,omitempty"`

	// DocumentContentOperator is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html#cfn-qbusiness-datasource-inlinedocumentenrichmentconfiguration-documentcontentoperator
	DocumentContentOperator cfz.Expression[string] `json:"DocumentContentOperator,omitempty"`

	// Target is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html#cfn-qbusiness-datasource-inlinedocumentenrichmentconfiguration-target
	Target cfz.Expression[AWS_QBusiness_DataSource_DocumentAttributeTarget] `json:"Target,omitempty"`
}

// New__AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration initializes a new AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration.
func New__AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration() AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration {
	return AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration) GetType() string {
	return AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration__Type
}

// Set__Condition updates property "Condition".
func (t AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration) Set__Condition(v cfz.Expression[AWS_QBusiness_DataSource_DocumentAttributeCondition]) AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration {
	t.Condition = v
	return t
}

// SetV__Condition updates property "Condition".
func (t AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration) SetV__Condition(v AWS_QBusiness_DataSource_DocumentAttributeCondition) AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration {
	t.Condition = cfz.V(v)
	return t
}

// Set__DocumentContentOperator updates property "DocumentContentOperator".
func (t AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration) Set__DocumentContentOperator(v cfz.Expression[string]) AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration {
	t.DocumentContentOperator = v
	return t
}

// SetV__DocumentContentOperator updates property "DocumentContentOperator".
func (t AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration) SetV__DocumentContentOperator(v string) AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration {
	t.DocumentContentOperator = cfz.V(v)
	return t
}

// Set__Target updates property "Target".
func (t AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration) Set__Target(v cfz.Expression[AWS_QBusiness_DataSource_DocumentAttributeTarget]) AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration {
	t.Target = v
	return t
}

// SetV__Target updates property "Target".
func (t AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration) SetV__Target(v AWS_QBusiness_DataSource_DocumentAttributeTarget) AWS_QBusiness_DataSource_InlineDocumentEnrichmentConfiguration {
	t.Target = cfz.V(v)
	return t
}
