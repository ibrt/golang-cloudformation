// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3outposts

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_S3Outposts_AccessPoint)(nil)
	_ cfz.Resource                   = (*AWS_S3Outposts_AccessPoint)(nil)
)

const (
	// AWS_S3Outposts_AccessPoint__Type is the CloudFormation type for AWS::S3Outposts::AccessPoint.
	AWS_S3Outposts_AccessPoint__Type = "AWS::S3Outposts::AccessPoint"
)

var (
	// AWS_S3Outposts_AccessPoint__AttributesMap reports all the CloudFormation attributes for AWS::S3Outposts::AccessPoint.
	AWS_S3Outposts_AccessPoint__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_S3Outposts_AccessPoint__AttributesSlice reports all the CloudFormation attributes for AWS::S3Outposts::AccessPoint.
	AWS_S3Outposts_AccessPoint__AttributesSlice = []string{
		AWS_S3Outposts_AccessPoint__AttributesMap.Arn,
	}
)

var (
	// AWS_S3Outposts_AccessPoint__PropertiesMap reports all the CloudFormation properties for AWS::S3Outposts::AccessPoint.
	AWS_S3Outposts_AccessPoint__PropertiesMap = struct {
		Bucket           string
		Name             string
		Policy           string
		VpcConfiguration string
	}{
		Bucket:           "Bucket",
		Name:             "Name",
		Policy:           "Policy",
		VpcConfiguration: "VpcConfiguration",
	}

	// AWS_S3Outposts_AccessPoint__PropertiesSlice reports all the CloudFormation properties for AWS::S3Outposts::AccessPoint.
	AWS_S3Outposts_AccessPoint__PropertiesSlice = []string{
		AWS_S3Outposts_AccessPoint__PropertiesMap.Bucket,
		AWS_S3Outposts_AccessPoint__PropertiesMap.Name,
		AWS_S3Outposts_AccessPoint__PropertiesMap.Policy,
		AWS_S3Outposts_AccessPoint__PropertiesMap.VpcConfiguration,
	}
)

// AWS_S3Outposts_AccessPoint is a binding for AWS::S3Outposts::AccessPoint.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html
type AWS_S3Outposts_AccessPoint struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Bucket is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html#cfn-s3outposts-accesspoint-bucket
	Bucket cfz.Expression[string] `json:"Bucket,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html#cfn-s3outposts-accesspoint-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Policy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html#cfn-s3outposts-accesspoint-policy
	Policy cfz.Expression[json.RawMessage] `json:"Policy,omitempty"`

	// VpcConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3outposts-accesspoint.html#cfn-s3outposts-accesspoint-vpcconfiguration
	VpcConfiguration cfz.Expression[AWS_S3Outposts_AccessPoint_VpcConfiguration] `json:"VpcConfiguration,omitempty"`
}

// New__AWS_S3Outposts_AccessPoint initializes a new *AWS_S3Outposts_AccessPoint.
func New__AWS_S3Outposts_AccessPoint(logicalName string) *AWS_S3Outposts_AccessPoint {
	return &AWS_S3Outposts_AccessPoint{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_S3Outposts_AccessPoint) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_S3Outposts_AccessPoint) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_S3Outposts_AccessPoint) GetType() string {
	return AWS_S3Outposts_AccessPoint__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_S3Outposts_AccessPoint) Set__LogicalName(v string) *AWS_S3Outposts_AccessPoint {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_S3Outposts_AccessPoint) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_S3Outposts_AccessPoint {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_S3Outposts_AccessPoint) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_S3Outposts_AccessPoint {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_S3Outposts_AccessPoint) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_S3Outposts_AccessPoint {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_S3Outposts_AccessPoint) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_S3Outposts_AccessPoint {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_S3Outposts_AccessPoint) Set__RequestedOutputs(v []cfz.Output) *AWS_S3Outposts_AccessPoint {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_S3Outposts_AccessPoint) Add__RequestedOutputs(v ...cfz.Output) *AWS_S3Outposts_AccessPoint {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Bucket updates property "Bucket".
func (t *AWS_S3Outposts_AccessPoint) Set__Bucket(v cfz.Expression[string]) *AWS_S3Outposts_AccessPoint {
	t.Bucket = v
	return t
}

// SetV__Bucket updates property "Bucket".
func (t *AWS_S3Outposts_AccessPoint) SetV__Bucket(v string) *AWS_S3Outposts_AccessPoint {
	t.Bucket = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_S3Outposts_AccessPoint) Set__Name(v cfz.Expression[string]) *AWS_S3Outposts_AccessPoint {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_S3Outposts_AccessPoint) SetV__Name(v string) *AWS_S3Outposts_AccessPoint {
	t.Name = cfz.V(v)
	return t
}

// Set__Policy updates property "Policy".
func (t *AWS_S3Outposts_AccessPoint) Set__Policy(v cfz.Expression[json.RawMessage]) *AWS_S3Outposts_AccessPoint {
	t.Policy = v
	return t
}

// SetV__Policy updates property "Policy".
func (t *AWS_S3Outposts_AccessPoint) SetV__Policy(v json.RawMessage) *AWS_S3Outposts_AccessPoint {
	t.Policy = cfz.V(v)
	return t
}

// Set__VpcConfiguration updates property "VpcConfiguration".
func (t *AWS_S3Outposts_AccessPoint) Set__VpcConfiguration(v cfz.Expression[AWS_S3Outposts_AccessPoint_VpcConfiguration]) *AWS_S3Outposts_AccessPoint {
	t.VpcConfiguration = v
	return t
}

// SetV__VpcConfiguration updates property "VpcConfiguration".
func (t *AWS_S3Outposts_AccessPoint) SetV__VpcConfiguration(v AWS_S3Outposts_AccessPoint_VpcConfiguration) *AWS_S3Outposts_AccessPoint {
	t.VpcConfiguration = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_S3Outposts_AccessPoint) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_S3Outposts_AccessPoint) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_S3Outposts_AccessPoint__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_S3Outposts_AccessPoint) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_S3Outposts_AccessPoint) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_S3Outposts_AccessPoint) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_S3Outposts_AccessPoint

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

func (t *AWS_S3Outposts_AccessPoint) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
