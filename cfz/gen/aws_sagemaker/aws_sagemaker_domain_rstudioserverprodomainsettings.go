// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_Domain_RStudioServerProDomainSettings__Type is the CloudFormation type for AWS::SageMaker::Domain.RStudioServerProDomainSettings.
	AWS_SageMaker_Domain_RStudioServerProDomainSettings__Type = "AWS::SageMaker::Domain.RStudioServerProDomainSettings"
)

var (
	// AWS_SageMaker_Domain_RStudioServerProDomainSettings__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::Domain.RStudioServerProDomainSettings.
	AWS_SageMaker_Domain_RStudioServerProDomainSettings__PropertiesMap = struct {
		DefaultResourceSpec      string
		DomainExecutionRoleArn   string
		RStudioConnectUrl        string
		RStudioPackageManagerUrl string
	}{
		DefaultResourceSpec:      "DefaultResourceSpec",
		DomainExecutionRoleArn:   "DomainExecutionRoleArn",
		RStudioConnectUrl:        "RStudioConnectUrl",
		RStudioPackageManagerUrl: "RStudioPackageManagerUrl",
	}

	// AWS_SageMaker_Domain_RStudioServerProDomainSettings__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::Domain.RStudioServerProDomainSettings.
	AWS_SageMaker_Domain_RStudioServerProDomainSettings__PropertiesSlice = []string{
		AWS_SageMaker_Domain_RStudioServerProDomainSettings__PropertiesMap.DefaultResourceSpec,
		AWS_SageMaker_Domain_RStudioServerProDomainSettings__PropertiesMap.DomainExecutionRoleArn,
		AWS_SageMaker_Domain_RStudioServerProDomainSettings__PropertiesMap.RStudioConnectUrl,
		AWS_SageMaker_Domain_RStudioServerProDomainSettings__PropertiesMap.RStudioPackageManagerUrl,
	}
)

// AWS_SageMaker_Domain_RStudioServerProDomainSettings is a binding for AWS::SageMaker::Domain.RStudioServerProDomainSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-rstudioserverprodomainsettings.html
type AWS_SageMaker_Domain_RStudioServerProDomainSettings struct {
	// DefaultResourceSpec is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-rstudioserverprodomainsettings.html#cfn-sagemaker-domain-rstudioserverprodomainsettings-defaultresourcespec
	DefaultResourceSpec cfz.Expression[AWS_SageMaker_Domain_ResourceSpec] `json:"DefaultResourceSpec,omitempty"`

	// DomainExecutionRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-rstudioserverprodomainsettings.html#cfn-sagemaker-domain-rstudioserverprodomainsettings-domainexecutionrolearn
	DomainExecutionRoleArn cfz.Expression[string] `json:"DomainExecutionRoleArn,omitempty"`

	// RStudioConnectUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-rstudioserverprodomainsettings.html#cfn-sagemaker-domain-rstudioserverprodomainsettings-rstudioconnecturl
	RStudioConnectUrl cfz.Expression[string] `json:"RStudioConnectUrl,omitempty"`

	// RStudioPackageManagerUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-rstudioserverprodomainsettings.html#cfn-sagemaker-domain-rstudioserverprodomainsettings-rstudiopackagemanagerurl
	RStudioPackageManagerUrl cfz.Expression[string] `json:"RStudioPackageManagerUrl,omitempty"`
}

// New__AWS_SageMaker_Domain_RStudioServerProDomainSettings initializes a new AWS_SageMaker_Domain_RStudioServerProDomainSettings.
func New__AWS_SageMaker_Domain_RStudioServerProDomainSettings() AWS_SageMaker_Domain_RStudioServerProDomainSettings {
	return AWS_SageMaker_Domain_RStudioServerProDomainSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_Domain_RStudioServerProDomainSettings) GetType() string {
	return AWS_SageMaker_Domain_RStudioServerProDomainSettings__Type
}

// Set__DefaultResourceSpec updates property "DefaultResourceSpec".
func (t AWS_SageMaker_Domain_RStudioServerProDomainSettings) Set__DefaultResourceSpec(v cfz.Expression[AWS_SageMaker_Domain_ResourceSpec]) AWS_SageMaker_Domain_RStudioServerProDomainSettings {
	t.DefaultResourceSpec = v
	return t
}

// SetV__DefaultResourceSpec updates property "DefaultResourceSpec".
func (t AWS_SageMaker_Domain_RStudioServerProDomainSettings) SetV__DefaultResourceSpec(v AWS_SageMaker_Domain_ResourceSpec) AWS_SageMaker_Domain_RStudioServerProDomainSettings {
	t.DefaultResourceSpec = cfz.V(v)
	return t
}

// Set__DomainExecutionRoleArn updates property "DomainExecutionRoleArn".
func (t AWS_SageMaker_Domain_RStudioServerProDomainSettings) Set__DomainExecutionRoleArn(v cfz.Expression[string]) AWS_SageMaker_Domain_RStudioServerProDomainSettings {
	t.DomainExecutionRoleArn = v
	return t
}

// SetV__DomainExecutionRoleArn updates property "DomainExecutionRoleArn".
func (t AWS_SageMaker_Domain_RStudioServerProDomainSettings) SetV__DomainExecutionRoleArn(v string) AWS_SageMaker_Domain_RStudioServerProDomainSettings {
	t.DomainExecutionRoleArn = cfz.V(v)
	return t
}

// Set__RStudioConnectUrl updates property "RStudioConnectUrl".
func (t AWS_SageMaker_Domain_RStudioServerProDomainSettings) Set__RStudioConnectUrl(v cfz.Expression[string]) AWS_SageMaker_Domain_RStudioServerProDomainSettings {
	t.RStudioConnectUrl = v
	return t
}

// SetV__RStudioConnectUrl updates property "RStudioConnectUrl".
func (t AWS_SageMaker_Domain_RStudioServerProDomainSettings) SetV__RStudioConnectUrl(v string) AWS_SageMaker_Domain_RStudioServerProDomainSettings {
	t.RStudioConnectUrl = cfz.V(v)
	return t
}

// Set__RStudioPackageManagerUrl updates property "RStudioPackageManagerUrl".
func (t AWS_SageMaker_Domain_RStudioServerProDomainSettings) Set__RStudioPackageManagerUrl(v cfz.Expression[string]) AWS_SageMaker_Domain_RStudioServerProDomainSettings {
	t.RStudioPackageManagerUrl = v
	return t
}

// SetV__RStudioPackageManagerUrl updates property "RStudioPackageManagerUrl".
func (t AWS_SageMaker_Domain_RStudioServerProDomainSettings) SetV__RStudioPackageManagerUrl(v string) AWS_SageMaker_Domain_RStudioServerProDomainSettings {
	t.RStudioPackageManagerUrl = cfz.V(v)
	return t
}
