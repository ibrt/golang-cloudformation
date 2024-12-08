// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_opensearchservice

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_OpenSearchService_Domain_ServiceSoftwareOptions__Type is the CloudFormation type for AWS::OpenSearchService::Domain.ServiceSoftwareOptions.
	AWS_OpenSearchService_Domain_ServiceSoftwareOptions__Type = "AWS::OpenSearchService::Domain.ServiceSoftwareOptions"
)

var (
	// AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap reports all the CloudFormation properties for AWS::OpenSearchService::Domain.ServiceSoftwareOptions.
	AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap = struct {
		AutomatedUpdateDate string
		Cancellable         string
		CurrentVersion      string
		Description         string
		NewVersion          string
		OptionalDeployment  string
		UpdateAvailable     string
		UpdateStatus        string
	}{
		AutomatedUpdateDate: "AutomatedUpdateDate",
		Cancellable:         "Cancellable",
		CurrentVersion:      "CurrentVersion",
		Description:         "Description",
		NewVersion:          "NewVersion",
		OptionalDeployment:  "OptionalDeployment",
		UpdateAvailable:     "UpdateAvailable",
		UpdateStatus:        "UpdateStatus",
	}

	// AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesSlice reports all the CloudFormation properties for AWS::OpenSearchService::Domain.ServiceSoftwareOptions.
	AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesSlice = []string{
		AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap.AutomatedUpdateDate,
		AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap.Cancellable,
		AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap.CurrentVersion,
		AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap.Description,
		AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap.NewVersion,
		AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap.OptionalDeployment,
		AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap.UpdateAvailable,
		AWS_OpenSearchService_Domain_ServiceSoftwareOptions__PropertiesMap.UpdateStatus,
	}
)

// AWS_OpenSearchService_Domain_ServiceSoftwareOptions is a binding for AWS::OpenSearchService::Domain.ServiceSoftwareOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html
type AWS_OpenSearchService_Domain_ServiceSoftwareOptions struct {
	// AutomatedUpdateDate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-automatedupdatedate
	AutomatedUpdateDate cfz.Expression[string] `json:"AutomatedUpdateDate,omitempty"`

	// Cancellable is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-cancellable
	Cancellable cfz.Expression[bool] `json:"Cancellable,omitempty"`

	// CurrentVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-currentversion
	CurrentVersion cfz.Expression[string] `json:"CurrentVersion,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// NewVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-newversion
	NewVersion cfz.Expression[string] `json:"NewVersion,omitempty"`

	// OptionalDeployment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-optionaldeployment
	OptionalDeployment cfz.Expression[bool] `json:"OptionalDeployment,omitempty"`

	// UpdateAvailable is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-updateavailable
	UpdateAvailable cfz.Expression[bool] `json:"UpdateAvailable,omitempty"`

	// UpdateStatus is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-servicesoftwareoptions.html#cfn-opensearchservice-domain-servicesoftwareoptions-updatestatus
	UpdateStatus cfz.Expression[string] `json:"UpdateStatus,omitempty"`
}

// New__AWS_OpenSearchService_Domain_ServiceSoftwareOptions initializes a new AWS_OpenSearchService_Domain_ServiceSoftwareOptions.
func New__AWS_OpenSearchService_Domain_ServiceSoftwareOptions() AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	return AWS_OpenSearchService_Domain_ServiceSoftwareOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_OpenSearchService_Domain_ServiceSoftwareOptions) GetType() string {
	return AWS_OpenSearchService_Domain_ServiceSoftwareOptions__Type
}

// Set__AutomatedUpdateDate updates property "AutomatedUpdateDate".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) Set__AutomatedUpdateDate(v cfz.Expression[string]) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.AutomatedUpdateDate = v
	return t
}

// SetV__AutomatedUpdateDate updates property "AutomatedUpdateDate".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) SetV__AutomatedUpdateDate(v string) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.AutomatedUpdateDate = cfz.V(v)
	return t
}

// Set__Cancellable updates property "Cancellable".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) Set__Cancellable(v cfz.Expression[bool]) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.Cancellable = v
	return t
}

// SetV__Cancellable updates property "Cancellable".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) SetV__Cancellable(v bool) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.Cancellable = cfz.V(v)
	return t
}

// Set__CurrentVersion updates property "CurrentVersion".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) Set__CurrentVersion(v cfz.Expression[string]) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.CurrentVersion = v
	return t
}

// SetV__CurrentVersion updates property "CurrentVersion".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) SetV__CurrentVersion(v string) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.CurrentVersion = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) Set__Description(v cfz.Expression[string]) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) SetV__Description(v string) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.Description = cfz.V(v)
	return t
}

// Set__NewVersion updates property "NewVersion".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) Set__NewVersion(v cfz.Expression[string]) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.NewVersion = v
	return t
}

// SetV__NewVersion updates property "NewVersion".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) SetV__NewVersion(v string) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.NewVersion = cfz.V(v)
	return t
}

// Set__OptionalDeployment updates property "OptionalDeployment".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) Set__OptionalDeployment(v cfz.Expression[bool]) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.OptionalDeployment = v
	return t
}

// SetV__OptionalDeployment updates property "OptionalDeployment".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) SetV__OptionalDeployment(v bool) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.OptionalDeployment = cfz.V(v)
	return t
}

// Set__UpdateAvailable updates property "UpdateAvailable".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) Set__UpdateAvailable(v cfz.Expression[bool]) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.UpdateAvailable = v
	return t
}

// SetV__UpdateAvailable updates property "UpdateAvailable".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) SetV__UpdateAvailable(v bool) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.UpdateAvailable = cfz.V(v)
	return t
}

// Set__UpdateStatus updates property "UpdateStatus".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) Set__UpdateStatus(v cfz.Expression[string]) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.UpdateStatus = v
	return t
}

// SetV__UpdateStatus updates property "UpdateStatus".
func (t AWS_OpenSearchService_Domain_ServiceSoftwareOptions) SetV__UpdateStatus(v string) AWS_OpenSearchService_Domain_ServiceSoftwareOptions {
	t.UpdateStatus = cfz.V(v)
	return t
}
