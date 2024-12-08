// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datazone

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_DataZone_UserProfile)(nil)
	_ cfz.Resource                   = (*AWS_DataZone_UserProfile)(nil)
)

const (
	// AWS_DataZone_UserProfile__Type is the CloudFormation type for AWS::DataZone::UserProfile.
	AWS_DataZone_UserProfile__Type = "AWS::DataZone::UserProfile"
)

var (
	// AWS_DataZone_UserProfile__AttributesMap reports all the CloudFormation attributes for AWS::DataZone::UserProfile.
	AWS_DataZone_UserProfile__AttributesMap = struct {
		Details               string
		Details_Iam           string
		Details_Iam_Arn       string
		Details_Sso           string
		Details_Sso_FirstName string
		Details_Sso_LastName  string
		Details_Sso_Username  string
		DomainId              string
		Id                    string
		Type                  string
	}{
		Details:               "Details",
		Details_Iam:           "Details.Iam",
		Details_Iam_Arn:       "Details.Iam.Arn",
		Details_Sso:           "Details.Sso",
		Details_Sso_FirstName: "Details.Sso.FirstName",
		Details_Sso_LastName:  "Details.Sso.LastName",
		Details_Sso_Username:  "Details.Sso.Username",
		DomainId:              "DomainId",
		Id:                    "Id",
		Type:                  "Type",
	}

	// AWS_DataZone_UserProfile__AttributesSlice reports all the CloudFormation attributes for AWS::DataZone::UserProfile.
	AWS_DataZone_UserProfile__AttributesSlice = []string{
		AWS_DataZone_UserProfile__AttributesMap.Details,
		AWS_DataZone_UserProfile__AttributesMap.Details_Iam,
		AWS_DataZone_UserProfile__AttributesMap.Details_Iam_Arn,
		AWS_DataZone_UserProfile__AttributesMap.Details_Sso,
		AWS_DataZone_UserProfile__AttributesMap.Details_Sso_FirstName,
		AWS_DataZone_UserProfile__AttributesMap.Details_Sso_LastName,
		AWS_DataZone_UserProfile__AttributesMap.Details_Sso_Username,
		AWS_DataZone_UserProfile__AttributesMap.DomainId,
		AWS_DataZone_UserProfile__AttributesMap.Id,
		AWS_DataZone_UserProfile__AttributesMap.Type,
	}
)

var (
	// AWS_DataZone_UserProfile__PropertiesMap reports all the CloudFormation properties for AWS::DataZone::UserProfile.
	AWS_DataZone_UserProfile__PropertiesMap = struct {
		DomainIdentifier string
		Status           string
		UserIdentifier   string
		UserType         string
	}{
		DomainIdentifier: "DomainIdentifier",
		Status:           "Status",
		UserIdentifier:   "UserIdentifier",
		UserType:         "UserType",
	}

	// AWS_DataZone_UserProfile__PropertiesSlice reports all the CloudFormation properties for AWS::DataZone::UserProfile.
	AWS_DataZone_UserProfile__PropertiesSlice = []string{
		AWS_DataZone_UserProfile__PropertiesMap.DomainIdentifier,
		AWS_DataZone_UserProfile__PropertiesMap.Status,
		AWS_DataZone_UserProfile__PropertiesMap.UserIdentifier,
		AWS_DataZone_UserProfile__PropertiesMap.UserType,
	}
)

// AWS_DataZone_UserProfile is a binding for AWS::DataZone::UserProfile.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html
type AWS_DataZone_UserProfile struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DomainIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-domainidentifier
	DomainIdentifier cfz.Expression[string] `json:"DomainIdentifier,omitempty"`

	// Status is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-status
	Status cfz.Expression[string] `json:"Status,omitempty"`

	// UserIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-useridentifier
	UserIdentifier cfz.Expression[string] `json:"UserIdentifier,omitempty"`

	// UserType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-userprofile.html#cfn-datazone-userprofile-usertype
	UserType cfz.Expression[string] `json:"UserType,omitempty"`
}

// New__AWS_DataZone_UserProfile initializes a new *AWS_DataZone_UserProfile.
func New__AWS_DataZone_UserProfile(logicalName string) *AWS_DataZone_UserProfile {
	return &AWS_DataZone_UserProfile{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_DataZone_UserProfile) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_DataZone_UserProfile) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_DataZone_UserProfile) GetType() string {
	return AWS_DataZone_UserProfile__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_DataZone_UserProfile) Set__LogicalName(v string) *AWS_DataZone_UserProfile {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_DataZone_UserProfile) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_DataZone_UserProfile {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_DataZone_UserProfile) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_DataZone_UserProfile {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_DataZone_UserProfile) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_DataZone_UserProfile {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_DataZone_UserProfile) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_DataZone_UserProfile {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_DataZone_UserProfile) Set__RequestedOutputs(v []cfz.Output) *AWS_DataZone_UserProfile {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_DataZone_UserProfile) Add__RequestedOutputs(v ...cfz.Output) *AWS_DataZone_UserProfile {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DomainIdentifier updates property "DomainIdentifier".
func (t *AWS_DataZone_UserProfile) Set__DomainIdentifier(v cfz.Expression[string]) *AWS_DataZone_UserProfile {
	t.DomainIdentifier = v
	return t
}

// SetV__DomainIdentifier updates property "DomainIdentifier".
func (t *AWS_DataZone_UserProfile) SetV__DomainIdentifier(v string) *AWS_DataZone_UserProfile {
	t.DomainIdentifier = cfz.V(v)
	return t
}

// Set__Status updates property "Status".
func (t *AWS_DataZone_UserProfile) Set__Status(v cfz.Expression[string]) *AWS_DataZone_UserProfile {
	t.Status = v
	return t
}

// SetV__Status updates property "Status".
func (t *AWS_DataZone_UserProfile) SetV__Status(v string) *AWS_DataZone_UserProfile {
	t.Status = cfz.V(v)
	return t
}

// Set__UserIdentifier updates property "UserIdentifier".
func (t *AWS_DataZone_UserProfile) Set__UserIdentifier(v cfz.Expression[string]) *AWS_DataZone_UserProfile {
	t.UserIdentifier = v
	return t
}

// SetV__UserIdentifier updates property "UserIdentifier".
func (t *AWS_DataZone_UserProfile) SetV__UserIdentifier(v string) *AWS_DataZone_UserProfile {
	t.UserIdentifier = cfz.V(v)
	return t
}

// Set__UserType updates property "UserType".
func (t *AWS_DataZone_UserProfile) Set__UserType(v cfz.Expression[string]) *AWS_DataZone_UserProfile {
	t.UserType = v
	return t
}

// SetV__UserType updates property "UserType".
func (t *AWS_DataZone_UserProfile) SetV__UserType(v string) *AWS_DataZone_UserProfile {
	t.UserType = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_DataZone_UserProfile) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Details returns a $cfz.Expression[AWS_DataZone_UserProfile_UserProfileDetails] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Details
func (t *AWS_DataZone_UserProfile) GetAtt__Details() cfz.Expression[AWS_DataZone_UserProfile_UserProfileDetails] {
	return cfz.GetAtt[AWS_DataZone_UserProfile_UserProfileDetails](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.Details))
}

// GetAtt__Details_Iam returns a $cfz.Expression[AWS_DataZone_UserProfile_IamUserProfileDetails] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Details.Iam
func (t *AWS_DataZone_UserProfile) GetAtt__Details_Iam() cfz.Expression[AWS_DataZone_UserProfile_IamUserProfileDetails] {
	return cfz.GetAtt[AWS_DataZone_UserProfile_IamUserProfileDetails](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.Details_Iam))
}

// GetAtt__Details_Iam_Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Details.Iam.Arn
func (t *AWS_DataZone_UserProfile) GetAtt__Details_Iam_Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.Details_Iam_Arn))
}

// GetAtt__Details_Sso returns a $cfz.Expression[AWS_DataZone_UserProfile_SsoUserProfileDetails] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Details.Sso
func (t *AWS_DataZone_UserProfile) GetAtt__Details_Sso() cfz.Expression[AWS_DataZone_UserProfile_SsoUserProfileDetails] {
	return cfz.GetAtt[AWS_DataZone_UserProfile_SsoUserProfileDetails](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.Details_Sso))
}

// GetAtt__Details_Sso_FirstName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Details.Sso.FirstName
func (t *AWS_DataZone_UserProfile) GetAtt__Details_Sso_FirstName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.Details_Sso_FirstName))
}

// GetAtt__Details_Sso_LastName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Details.Sso.LastName
func (t *AWS_DataZone_UserProfile) GetAtt__Details_Sso_LastName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.Details_Sso_LastName))
}

// GetAtt__Details_Sso_Username returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Details.Sso.Username
func (t *AWS_DataZone_UserProfile) GetAtt__Details_Sso_Username() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.Details_Sso_Username))
}

// GetAtt__DomainId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DomainId
func (t *AWS_DataZone_UserProfile) GetAtt__DomainId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.DomainId))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_DataZone_UserProfile) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.Id))
}

// GetAtt__Type returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Type
func (t *AWS_DataZone_UserProfile) GetAtt__Type() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_DataZone_UserProfile__AttributesMap.Type))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_DataZone_UserProfile) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Details returns a conventionally configured output for an attribute of this resource.
// Attribute: Details
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__Details(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDetails", t.GetAtt__Details())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Details_Iam returns a conventionally configured output for an attribute of this resource.
// Attribute: Details.Iam
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__Details_Iam(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDetailsIam", t.GetAtt__Details_Iam())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Details_Iam_Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Details.Iam.Arn
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__Details_Iam_Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDetailsIamArn", t.GetAtt__Details_Iam_Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Details_Sso returns a conventionally configured output for an attribute of this resource.
// Attribute: Details.Sso
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__Details_Sso(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDetailsSso", t.GetAtt__Details_Sso())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Details_Sso_FirstName returns a conventionally configured output for an attribute of this resource.
// Attribute: Details.Sso.FirstName
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__Details_Sso_FirstName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDetailsSsoFirstName", t.GetAtt__Details_Sso_FirstName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Details_Sso_LastName returns a conventionally configured output for an attribute of this resource.
// Attribute: Details.Sso.LastName
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__Details_Sso_LastName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDetailsSsoLastName", t.GetAtt__Details_Sso_LastName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Details_Sso_Username returns a conventionally configured output for an attribute of this resource.
// Attribute: Details.Sso.Username
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__Details_Sso_Username(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDetailsSsoUsername", t.GetAtt__Details_Sso_Username())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DomainId returns a conventionally configured output for an attribute of this resource.
// Attribute: DomainId
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__DomainId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDomainId", t.GetAtt__DomainId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Type returns a conventionally configured output for an attribute of this resource.
// Attribute: Type
func (t *AWS_DataZone_UserProfile) GetConventionalOutputAtt__Type(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttType", t.GetAtt__Type())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_DataZone_UserProfile) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_DataZone_UserProfile

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

func (t *AWS_DataZone_UserProfile) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
