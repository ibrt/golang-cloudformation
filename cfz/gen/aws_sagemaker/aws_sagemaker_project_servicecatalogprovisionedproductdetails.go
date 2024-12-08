// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails__Type is the CloudFormation type for AWS::SageMaker::Project.ServiceCatalogProvisionedProductDetails.
	AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails__Type = "AWS::SageMaker::Project.ServiceCatalogProvisionedProductDetails"
)

var (
	// AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::Project.ServiceCatalogProvisionedProductDetails.
	AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails__PropertiesMap = struct {
		ProvisionedProductId            string
		ProvisionedProductStatusMessage string
	}{
		ProvisionedProductId:            "ProvisionedProductId",
		ProvisionedProductStatusMessage: "ProvisionedProductStatusMessage",
	}

	// AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::Project.ServiceCatalogProvisionedProductDetails.
	AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails__PropertiesSlice = []string{
		AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails__PropertiesMap.ProvisionedProductId,
		AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails__PropertiesMap.ProvisionedProductStatusMessage,
	}
)

// AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails is a binding for AWS::SageMaker::Project.ServiceCatalogProvisionedProductDetails.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisionedproductdetails.html
type AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails struct {
	// ProvisionedProductId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisionedproductdetails.html#cfn-sagemaker-project-servicecatalogprovisionedproductdetails-provisionedproductid
	ProvisionedProductId cfz.Expression[string] `json:"ProvisionedProductId,omitempty"`

	// ProvisionedProductStatusMessage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisionedproductdetails.html#cfn-sagemaker-project-servicecatalogprovisionedproductdetails-provisionedproductstatusmessage
	ProvisionedProductStatusMessage cfz.Expression[string] `json:"ProvisionedProductStatusMessage,omitempty"`
}

// New__AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails initializes a new AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails.
func New__AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails() AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails {
	return AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails) GetType() string {
	return AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails__Type
}

// Set__ProvisionedProductId updates property "ProvisionedProductId".
func (t AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails) Set__ProvisionedProductId(v cfz.Expression[string]) AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails {
	t.ProvisionedProductId = v
	return t
}

// SetV__ProvisionedProductId updates property "ProvisionedProductId".
func (t AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails) SetV__ProvisionedProductId(v string) AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails {
	t.ProvisionedProductId = cfz.V(v)
	return t
}

// Set__ProvisionedProductStatusMessage updates property "ProvisionedProductStatusMessage".
func (t AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails) Set__ProvisionedProductStatusMessage(v cfz.Expression[string]) AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails {
	t.ProvisionedProductStatusMessage = v
	return t
}

// SetV__ProvisionedProductStatusMessage updates property "ProvisionedProductStatusMessage".
func (t AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails) SetV__ProvisionedProductStatusMessage(v string) AWS_SageMaker_Project_ServiceCatalogProvisionedProductDetails {
	t.ProvisionedProductStatusMessage = cfz.V(v)
	return t
}
