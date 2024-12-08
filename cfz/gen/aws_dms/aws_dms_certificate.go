// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_dms

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_DMS_Certificate)(nil)
	_ cfz.Resource                   = (*AWS_DMS_Certificate)(nil)
)

const (
	// AWS_DMS_Certificate__Type is the CloudFormation type for AWS::DMS::Certificate.
	AWS_DMS_Certificate__Type = "AWS::DMS::Certificate"
)

var (
	// AWS_DMS_Certificate__PropertiesMap reports all the CloudFormation properties for AWS::DMS::Certificate.
	AWS_DMS_Certificate__PropertiesMap = struct {
		CertificateIdentifier string
		CertificatePem        string
		CertificateWallet     string
	}{
		CertificateIdentifier: "CertificateIdentifier",
		CertificatePem:        "CertificatePem",
		CertificateWallet:     "CertificateWallet",
	}

	// AWS_DMS_Certificate__PropertiesSlice reports all the CloudFormation properties for AWS::DMS::Certificate.
	AWS_DMS_Certificate__PropertiesSlice = []string{
		AWS_DMS_Certificate__PropertiesMap.CertificateIdentifier,
		AWS_DMS_Certificate__PropertiesMap.CertificatePem,
		AWS_DMS_Certificate__PropertiesMap.CertificateWallet,
	}
)

// AWS_DMS_Certificate is a binding for AWS::DMS::Certificate.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html
type AWS_DMS_Certificate struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CertificateIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificateidentifier
	CertificateIdentifier cfz.Expression[string] `json:"CertificateIdentifier,omitempty"`

	// CertificatePem is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificatepem
	CertificatePem cfz.Expression[string] `json:"CertificatePem,omitempty"`

	// CertificateWallet is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificatewallet
	CertificateWallet cfz.Expression[string] `json:"CertificateWallet,omitempty"`
}

// New__AWS_DMS_Certificate initializes a new *AWS_DMS_Certificate.
func New__AWS_DMS_Certificate(logicalName string) *AWS_DMS_Certificate {
	return &AWS_DMS_Certificate{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_DMS_Certificate) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_DMS_Certificate) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_DMS_Certificate) GetType() string {
	return AWS_DMS_Certificate__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_DMS_Certificate) Set__LogicalName(v string) *AWS_DMS_Certificate {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_DMS_Certificate) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_DMS_Certificate {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_DMS_Certificate) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_DMS_Certificate {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_DMS_Certificate) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_DMS_Certificate {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_DMS_Certificate) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_DMS_Certificate {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_DMS_Certificate) Set__RequestedOutputs(v []cfz.Output) *AWS_DMS_Certificate {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_DMS_Certificate) Add__RequestedOutputs(v ...cfz.Output) *AWS_DMS_Certificate {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CertificateIdentifier updates property "CertificateIdentifier".
func (t *AWS_DMS_Certificate) Set__CertificateIdentifier(v cfz.Expression[string]) *AWS_DMS_Certificate {
	t.CertificateIdentifier = v
	return t
}

// SetV__CertificateIdentifier updates property "CertificateIdentifier".
func (t *AWS_DMS_Certificate) SetV__CertificateIdentifier(v string) *AWS_DMS_Certificate {
	t.CertificateIdentifier = cfz.V(v)
	return t
}

// Set__CertificatePem updates property "CertificatePem".
func (t *AWS_DMS_Certificate) Set__CertificatePem(v cfz.Expression[string]) *AWS_DMS_Certificate {
	t.CertificatePem = v
	return t
}

// SetV__CertificatePem updates property "CertificatePem".
func (t *AWS_DMS_Certificate) SetV__CertificatePem(v string) *AWS_DMS_Certificate {
	t.CertificatePem = cfz.V(v)
	return t
}

// Set__CertificateWallet updates property "CertificateWallet".
func (t *AWS_DMS_Certificate) Set__CertificateWallet(v cfz.Expression[string]) *AWS_DMS_Certificate {
	t.CertificateWallet = v
	return t
}

// SetV__CertificateWallet updates property "CertificateWallet".
func (t *AWS_DMS_Certificate) SetV__CertificateWallet(v string) *AWS_DMS_Certificate {
	t.CertificateWallet = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_DMS_Certificate) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_DMS_Certificate) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_DMS_Certificate) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_DMS_Certificate

	return json.Marshal(struct {
		Type                string                          `json:"Type"`
		DependsOn           []string                        `json:"DependsOn,omitempty"`
		DeletionPolicy      cfz.ResourceDeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Properties          *Properties                     `json:"Properties,omitempty"`
	}{
		Type:       t.GetType(),
		DependsOn:  t.getDependsOn(),
		Properties: (*Properties)(t),
	})
}

func (t *AWS_DMS_Certificate) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
