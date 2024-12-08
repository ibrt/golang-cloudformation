// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_backupgateway

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_BackupGateway_Hypervisor)(nil)
	_ cfz.Resource                   = (*AWS_BackupGateway_Hypervisor)(nil)
)

const (
	// AWS_BackupGateway_Hypervisor__Type is the CloudFormation type for AWS::BackupGateway::Hypervisor.
	AWS_BackupGateway_Hypervisor__Type = "AWS::BackupGateway::Hypervisor"
)

var (
	// AWS_BackupGateway_Hypervisor__AttributesMap reports all the CloudFormation attributes for AWS::BackupGateway::Hypervisor.
	AWS_BackupGateway_Hypervisor__AttributesMap = struct {
		HypervisorArn string
	}{
		HypervisorArn: "HypervisorArn",
	}

	// AWS_BackupGateway_Hypervisor__AttributesSlice reports all the CloudFormation attributes for AWS::BackupGateway::Hypervisor.
	AWS_BackupGateway_Hypervisor__AttributesSlice = []string{
		AWS_BackupGateway_Hypervisor__AttributesMap.HypervisorArn,
	}
)

var (
	// AWS_BackupGateway_Hypervisor__PropertiesMap reports all the CloudFormation properties for AWS::BackupGateway::Hypervisor.
	AWS_BackupGateway_Hypervisor__PropertiesMap = struct {
		Host        string
		KmsKeyArn   string
		LogGroupArn string
		Name        string
		Password    string
		Tags        string
		Username    string
	}{
		Host:        "Host",
		KmsKeyArn:   "KmsKeyArn",
		LogGroupArn: "LogGroupArn",
		Name:        "Name",
		Password:    "Password",
		Tags:        "Tags",
		Username:    "Username",
	}

	// AWS_BackupGateway_Hypervisor__PropertiesSlice reports all the CloudFormation properties for AWS::BackupGateway::Hypervisor.
	AWS_BackupGateway_Hypervisor__PropertiesSlice = []string{
		AWS_BackupGateway_Hypervisor__PropertiesMap.Host,
		AWS_BackupGateway_Hypervisor__PropertiesMap.KmsKeyArn,
		AWS_BackupGateway_Hypervisor__PropertiesMap.LogGroupArn,
		AWS_BackupGateway_Hypervisor__PropertiesMap.Name,
		AWS_BackupGateway_Hypervisor__PropertiesMap.Password,
		AWS_BackupGateway_Hypervisor__PropertiesMap.Tags,
		AWS_BackupGateway_Hypervisor__PropertiesMap.Username,
	}
)

// AWS_BackupGateway_Hypervisor is a binding for AWS::BackupGateway::Hypervisor.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html
type AWS_BackupGateway_Hypervisor struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Host is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-host
	Host cfz.Expression[string] `json:"Host,omitempty"`

	// KmsKeyArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-kmskeyarn
	KmsKeyArn cfz.Expression[string] `json:"KmsKeyArn,omitempty"`

	// LogGroupArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-loggrouparn
	LogGroupArn cfz.Expression[string] `json:"LogGroupArn,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Password is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-password
	Password cfz.Expression[string] `json:"Password,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// Username is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupgateway-hypervisor.html#cfn-backupgateway-hypervisor-username
	Username cfz.Expression[string] `json:"Username,omitempty"`
}

// New__AWS_BackupGateway_Hypervisor initializes a new *AWS_BackupGateway_Hypervisor.
func New__AWS_BackupGateway_Hypervisor(logicalName string) *AWS_BackupGateway_Hypervisor {
	return &AWS_BackupGateway_Hypervisor{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_BackupGateway_Hypervisor) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_BackupGateway_Hypervisor) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_BackupGateway_Hypervisor) GetType() string {
	return AWS_BackupGateway_Hypervisor__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_BackupGateway_Hypervisor) Set__LogicalName(v string) *AWS_BackupGateway_Hypervisor {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_BackupGateway_Hypervisor) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_BackupGateway_Hypervisor {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_BackupGateway_Hypervisor) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_BackupGateway_Hypervisor {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_BackupGateway_Hypervisor) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_BackupGateway_Hypervisor {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_BackupGateway_Hypervisor) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_BackupGateway_Hypervisor {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_BackupGateway_Hypervisor) Set__RequestedOutputs(v []cfz.Output) *AWS_BackupGateway_Hypervisor {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_BackupGateway_Hypervisor) Add__RequestedOutputs(v ...cfz.Output) *AWS_BackupGateway_Hypervisor {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Host updates property "Host".
func (t *AWS_BackupGateway_Hypervisor) Set__Host(v cfz.Expression[string]) *AWS_BackupGateway_Hypervisor {
	t.Host = v
	return t
}

// SetV__Host updates property "Host".
func (t *AWS_BackupGateway_Hypervisor) SetV__Host(v string) *AWS_BackupGateway_Hypervisor {
	t.Host = cfz.V(v)
	return t
}

// Set__KmsKeyArn updates property "KmsKeyArn".
func (t *AWS_BackupGateway_Hypervisor) Set__KmsKeyArn(v cfz.Expression[string]) *AWS_BackupGateway_Hypervisor {
	t.KmsKeyArn = v
	return t
}

// SetV__KmsKeyArn updates property "KmsKeyArn".
func (t *AWS_BackupGateway_Hypervisor) SetV__KmsKeyArn(v string) *AWS_BackupGateway_Hypervisor {
	t.KmsKeyArn = cfz.V(v)
	return t
}

// Set__LogGroupArn updates property "LogGroupArn".
func (t *AWS_BackupGateway_Hypervisor) Set__LogGroupArn(v cfz.Expression[string]) *AWS_BackupGateway_Hypervisor {
	t.LogGroupArn = v
	return t
}

// SetV__LogGroupArn updates property "LogGroupArn".
func (t *AWS_BackupGateway_Hypervisor) SetV__LogGroupArn(v string) *AWS_BackupGateway_Hypervisor {
	t.LogGroupArn = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_BackupGateway_Hypervisor) Set__Name(v cfz.Expression[string]) *AWS_BackupGateway_Hypervisor {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_BackupGateway_Hypervisor) SetV__Name(v string) *AWS_BackupGateway_Hypervisor {
	t.Name = cfz.V(v)
	return t
}

// Set__Password updates property "Password".
func (t *AWS_BackupGateway_Hypervisor) Set__Password(v cfz.Expression[string]) *AWS_BackupGateway_Hypervisor {
	t.Password = v
	return t
}

// SetV__Password updates property "Password".
func (t *AWS_BackupGateway_Hypervisor) SetV__Password(v string) *AWS_BackupGateway_Hypervisor {
	t.Password = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_BackupGateway_Hypervisor) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_BackupGateway_Hypervisor {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_BackupGateway_Hypervisor) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_BackupGateway_Hypervisor {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_BackupGateway_Hypervisor) SetSV__Tags(v ...cfz.Tag) *AWS_BackupGateway_Hypervisor {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__Username updates property "Username".
func (t *AWS_BackupGateway_Hypervisor) Set__Username(v cfz.Expression[string]) *AWS_BackupGateway_Hypervisor {
	t.Username = v
	return t
}

// SetV__Username updates property "Username".
func (t *AWS_BackupGateway_Hypervisor) SetV__Username(v string) *AWS_BackupGateway_Hypervisor {
	t.Username = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_BackupGateway_Hypervisor) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__HypervisorArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: HypervisorArn
func (t *AWS_BackupGateway_Hypervisor) GetAtt__HypervisorArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_BackupGateway_Hypervisor__AttributesMap.HypervisorArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_BackupGateway_Hypervisor) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__HypervisorArn returns a conventionally configured output for an attribute of this resource.
// Attribute: HypervisorArn
func (t *AWS_BackupGateway_Hypervisor) GetConventionalOutputAtt__HypervisorArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttHypervisorArn", t.GetAtt__HypervisorArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_BackupGateway_Hypervisor) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_BackupGateway_Hypervisor

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

func (t *AWS_BackupGateway_Hypervisor) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
