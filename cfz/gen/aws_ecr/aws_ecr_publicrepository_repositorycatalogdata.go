// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ecr

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ECR_PublicRepository_RepositoryCatalogData__Type is the CloudFormation type for AWS::ECR::PublicRepository.RepositoryCatalogData.
	AWS_ECR_PublicRepository_RepositoryCatalogData__Type = "AWS::ECR::PublicRepository.RepositoryCatalogData"
)

var (
	// AWS_ECR_PublicRepository_RepositoryCatalogData__PropertiesMap reports all the CloudFormation properties for AWS::ECR::PublicRepository.RepositoryCatalogData.
	AWS_ECR_PublicRepository_RepositoryCatalogData__PropertiesMap = struct {
		AboutText             string
		Architectures         string
		OperatingSystems      string
		RepositoryDescription string
		UsageText             string
	}{
		AboutText:             "AboutText",
		Architectures:         "Architectures",
		OperatingSystems:      "OperatingSystems",
		RepositoryDescription: "RepositoryDescription",
		UsageText:             "UsageText",
	}

	// AWS_ECR_PublicRepository_RepositoryCatalogData__PropertiesSlice reports all the CloudFormation properties for AWS::ECR::PublicRepository.RepositoryCatalogData.
	AWS_ECR_PublicRepository_RepositoryCatalogData__PropertiesSlice = []string{
		AWS_ECR_PublicRepository_RepositoryCatalogData__PropertiesMap.AboutText,
		AWS_ECR_PublicRepository_RepositoryCatalogData__PropertiesMap.Architectures,
		AWS_ECR_PublicRepository_RepositoryCatalogData__PropertiesMap.OperatingSystems,
		AWS_ECR_PublicRepository_RepositoryCatalogData__PropertiesMap.RepositoryDescription,
		AWS_ECR_PublicRepository_RepositoryCatalogData__PropertiesMap.UsageText,
	}
)

// AWS_ECR_PublicRepository_RepositoryCatalogData is a binding for AWS::ECR::PublicRepository.RepositoryCatalogData.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html
type AWS_ECR_PublicRepository_RepositoryCatalogData struct {
	// AboutText is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-abouttext
	AboutText cfz.Expression[string] `json:"AboutText,omitempty"`

	// Architectures is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-architectures
	Architectures cfz.ExpressionSlice[string] `json:"Architectures,omitempty"`

	// OperatingSystems is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-operatingsystems
	OperatingSystems cfz.ExpressionSlice[string] `json:"OperatingSystems,omitempty"`

	// RepositoryDescription is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-repositorydescription
	RepositoryDescription cfz.Expression[string] `json:"RepositoryDescription,omitempty"`

	// UsageText is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-usagetext
	UsageText cfz.Expression[string] `json:"UsageText,omitempty"`
}

// New__AWS_ECR_PublicRepository_RepositoryCatalogData initializes a new AWS_ECR_PublicRepository_RepositoryCatalogData.
func New__AWS_ECR_PublicRepository_RepositoryCatalogData() AWS_ECR_PublicRepository_RepositoryCatalogData {
	return AWS_ECR_PublicRepository_RepositoryCatalogData{}
}

// GetType returns the CloudFormation type.
func (AWS_ECR_PublicRepository_RepositoryCatalogData) GetType() string {
	return AWS_ECR_PublicRepository_RepositoryCatalogData__Type
}

// Set__AboutText updates property "AboutText".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) Set__AboutText(v cfz.Expression[string]) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.AboutText = v
	return t
}

// SetV__AboutText updates property "AboutText".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) SetV__AboutText(v string) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.AboutText = cfz.V(v)
	return t
}

// Set__Architectures updates property "Architectures".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) Set__Architectures(v cfz.ExpressionSlice[string]) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.Architectures = v
	return t
}

// SetS__Architectures updates property "Architectures".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) SetS__Architectures(v ...cfz.Expression[string]) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.Architectures = cfz.S(v...)
	return t
}

// SetSV__Architectures updates property "Architectures".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) SetSV__Architectures(v ...string) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.Architectures = cfz.SV(v...)
	return t
}

// Set__OperatingSystems updates property "OperatingSystems".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) Set__OperatingSystems(v cfz.ExpressionSlice[string]) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.OperatingSystems = v
	return t
}

// SetS__OperatingSystems updates property "OperatingSystems".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) SetS__OperatingSystems(v ...cfz.Expression[string]) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.OperatingSystems = cfz.S(v...)
	return t
}

// SetSV__OperatingSystems updates property "OperatingSystems".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) SetSV__OperatingSystems(v ...string) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.OperatingSystems = cfz.SV(v...)
	return t
}

// Set__RepositoryDescription updates property "RepositoryDescription".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) Set__RepositoryDescription(v cfz.Expression[string]) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.RepositoryDescription = v
	return t
}

// SetV__RepositoryDescription updates property "RepositoryDescription".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) SetV__RepositoryDescription(v string) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.RepositoryDescription = cfz.V(v)
	return t
}

// Set__UsageText updates property "UsageText".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) Set__UsageText(v cfz.Expression[string]) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.UsageText = v
	return t
}

// SetV__UsageText updates property "UsageText".
func (t AWS_ECR_PublicRepository_RepositoryCatalogData) SetV__UsageText(v string) AWS_ECR_PublicRepository_RepositoryCatalogData {
	t.UsageText = cfz.V(v)
	return t
}
