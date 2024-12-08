// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_signer

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Signer_SigningProfile)(nil)
	_ cfz.Resource                   = (*AWS_Signer_SigningProfile)(nil)
)

const (
	// AWS_Signer_SigningProfile__Type is the CloudFormation type for AWS::Signer::SigningProfile.
	AWS_Signer_SigningProfile__Type = "AWS::Signer::SigningProfile"
)

var (
	// AWS_Signer_SigningProfile__AttributesMap reports all the CloudFormation attributes for AWS::Signer::SigningProfile.
	AWS_Signer_SigningProfile__AttributesMap = struct {
		Arn               string
		ProfileName       string
		ProfileVersion    string
		ProfileVersionArn string
	}{
		Arn:               "Arn",
		ProfileName:       "ProfileName",
		ProfileVersion:    "ProfileVersion",
		ProfileVersionArn: "ProfileVersionArn",
	}

	// AWS_Signer_SigningProfile__AttributesSlice reports all the CloudFormation attributes for AWS::Signer::SigningProfile.
	AWS_Signer_SigningProfile__AttributesSlice = []string{
		AWS_Signer_SigningProfile__AttributesMap.Arn,
		AWS_Signer_SigningProfile__AttributesMap.ProfileName,
		AWS_Signer_SigningProfile__AttributesMap.ProfileVersion,
		AWS_Signer_SigningProfile__AttributesMap.ProfileVersionArn,
	}
)

var (
	// AWS_Signer_SigningProfile__PropertiesMap reports all the CloudFormation properties for AWS::Signer::SigningProfile.
	AWS_Signer_SigningProfile__PropertiesMap = struct {
		PlatformId              string
		SignatureValidityPeriod string
		Tags                    string
	}{
		PlatformId:              "PlatformId",
		SignatureValidityPeriod: "SignatureValidityPeriod",
		Tags:                    "Tags",
	}

	// AWS_Signer_SigningProfile__PropertiesSlice reports all the CloudFormation properties for AWS::Signer::SigningProfile.
	AWS_Signer_SigningProfile__PropertiesSlice = []string{
		AWS_Signer_SigningProfile__PropertiesMap.PlatformId,
		AWS_Signer_SigningProfile__PropertiesMap.SignatureValidityPeriod,
		AWS_Signer_SigningProfile__PropertiesMap.Tags,
	}
)

// AWS_Signer_SigningProfile is a binding for AWS::Signer::SigningProfile.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-signingprofile.html
type AWS_Signer_SigningProfile struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// PlatformId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-signingprofile.html#cfn-signer-signingprofile-platformid
	PlatformId cfz.Expression[string] `json:"PlatformId,omitempty"`

	// SignatureValidityPeriod is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-signingprofile.html#cfn-signer-signingprofile-signaturevalidityperiod
	SignatureValidityPeriod cfz.Expression[AWS_Signer_SigningProfile_SignatureValidityPeriod] `json:"SignatureValidityPeriod,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-signingprofile.html#cfn-signer-signingprofile-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_Signer_SigningProfile initializes a new *AWS_Signer_SigningProfile.
func New__AWS_Signer_SigningProfile(logicalName string) *AWS_Signer_SigningProfile {
	return &AWS_Signer_SigningProfile{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Signer_SigningProfile) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Signer_SigningProfile) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Signer_SigningProfile) GetType() string {
	return AWS_Signer_SigningProfile__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Signer_SigningProfile) Set__LogicalName(v string) *AWS_Signer_SigningProfile {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Signer_SigningProfile) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Signer_SigningProfile {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Signer_SigningProfile) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Signer_SigningProfile {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Signer_SigningProfile) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Signer_SigningProfile {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Signer_SigningProfile) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Signer_SigningProfile {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Signer_SigningProfile) Set__RequestedOutputs(v []cfz.Output) *AWS_Signer_SigningProfile {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Signer_SigningProfile) Add__RequestedOutputs(v ...cfz.Output) *AWS_Signer_SigningProfile {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__PlatformId updates property "PlatformId".
func (t *AWS_Signer_SigningProfile) Set__PlatformId(v cfz.Expression[string]) *AWS_Signer_SigningProfile {
	t.PlatformId = v
	return t
}

// SetV__PlatformId updates property "PlatformId".
func (t *AWS_Signer_SigningProfile) SetV__PlatformId(v string) *AWS_Signer_SigningProfile {
	t.PlatformId = cfz.V(v)
	return t
}

// Set__SignatureValidityPeriod updates property "SignatureValidityPeriod".
func (t *AWS_Signer_SigningProfile) Set__SignatureValidityPeriod(v cfz.Expression[AWS_Signer_SigningProfile_SignatureValidityPeriod]) *AWS_Signer_SigningProfile {
	t.SignatureValidityPeriod = v
	return t
}

// SetV__SignatureValidityPeriod updates property "SignatureValidityPeriod".
func (t *AWS_Signer_SigningProfile) SetV__SignatureValidityPeriod(v AWS_Signer_SigningProfile_SignatureValidityPeriod) *AWS_Signer_SigningProfile {
	t.SignatureValidityPeriod = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Signer_SigningProfile) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Signer_SigningProfile {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Signer_SigningProfile) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Signer_SigningProfile {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Signer_SigningProfile) SetSV__Tags(v ...cfz.Tag) *AWS_Signer_SigningProfile {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Signer_SigningProfile) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Signer_SigningProfile) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Signer_SigningProfile__AttributesMap.Arn))
}

// GetAtt__ProfileName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ProfileName
func (t *AWS_Signer_SigningProfile) GetAtt__ProfileName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Signer_SigningProfile__AttributesMap.ProfileName))
}

// GetAtt__ProfileVersion returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ProfileVersion
func (t *AWS_Signer_SigningProfile) GetAtt__ProfileVersion() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Signer_SigningProfile__AttributesMap.ProfileVersion))
}

// GetAtt__ProfileVersionArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ProfileVersionArn
func (t *AWS_Signer_SigningProfile) GetAtt__ProfileVersionArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Signer_SigningProfile__AttributesMap.ProfileVersionArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Signer_SigningProfile) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Signer_SigningProfile) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ProfileName returns a conventionally configured output for an attribute of this resource.
// Attribute: ProfileName
func (t *AWS_Signer_SigningProfile) GetConventionalOutputAtt__ProfileName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttProfileName", t.GetAtt__ProfileName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ProfileVersion returns a conventionally configured output for an attribute of this resource.
// Attribute: ProfileVersion
func (t *AWS_Signer_SigningProfile) GetConventionalOutputAtt__ProfileVersion(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttProfileVersion", t.GetAtt__ProfileVersion())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ProfileVersionArn returns a conventionally configured output for an attribute of this resource.
// Attribute: ProfileVersionArn
func (t *AWS_Signer_SigningProfile) GetConventionalOutputAtt__ProfileVersionArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttProfileVersionArn", t.GetAtt__ProfileVersionArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Signer_SigningProfile) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Signer_SigningProfile

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

func (t *AWS_Signer_SigningProfile) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
