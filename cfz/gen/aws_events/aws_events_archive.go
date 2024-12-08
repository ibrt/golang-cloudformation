// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_events

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Events_Archive)(nil)
	_ cfz.Resource                   = (*AWS_Events_Archive)(nil)
)

const (
	// AWS_Events_Archive__Type is the CloudFormation type for AWS::Events::Archive.
	AWS_Events_Archive__Type = "AWS::Events::Archive"
)

var (
	// AWS_Events_Archive__AttributesMap reports all the CloudFormation attributes for AWS::Events::Archive.
	AWS_Events_Archive__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Events_Archive__AttributesSlice reports all the CloudFormation attributes for AWS::Events::Archive.
	AWS_Events_Archive__AttributesSlice = []string{
		AWS_Events_Archive__AttributesMap.Arn,
	}
)

var (
	// AWS_Events_Archive__PropertiesMap reports all the CloudFormation properties for AWS::Events::Archive.
	AWS_Events_Archive__PropertiesMap = struct {
		ArchiveName   string
		Description   string
		EventPattern  string
		RetentionDays string
		SourceArn     string
	}{
		ArchiveName:   "ArchiveName",
		Description:   "Description",
		EventPattern:  "EventPattern",
		RetentionDays: "RetentionDays",
		SourceArn:     "SourceArn",
	}

	// AWS_Events_Archive__PropertiesSlice reports all the CloudFormation properties for AWS::Events::Archive.
	AWS_Events_Archive__PropertiesSlice = []string{
		AWS_Events_Archive__PropertiesMap.ArchiveName,
		AWS_Events_Archive__PropertiesMap.Description,
		AWS_Events_Archive__PropertiesMap.EventPattern,
		AWS_Events_Archive__PropertiesMap.RetentionDays,
		AWS_Events_Archive__PropertiesMap.SourceArn,
	}
)

// AWS_Events_Archive is a binding for AWS::Events::Archive.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html
type AWS_Events_Archive struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ArchiveName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-archivename
	ArchiveName cfz.Expression[string] `json:"ArchiveName,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// EventPattern is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-eventpattern
	EventPattern cfz.Expression[json.RawMessage] `json:"EventPattern,omitempty"`

	// RetentionDays is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-retentiondays
	RetentionDays cfz.Expression[int32] `json:"RetentionDays,omitempty"`

	// SourceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-sourcearn
	SourceArn cfz.Expression[string] `json:"SourceArn,omitempty"`
}

// New__AWS_Events_Archive initializes a new *AWS_Events_Archive.
func New__AWS_Events_Archive(logicalName string) *AWS_Events_Archive {
	return &AWS_Events_Archive{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Events_Archive) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Events_Archive) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Events_Archive) GetType() string {
	return AWS_Events_Archive__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Events_Archive) Set__LogicalName(v string) *AWS_Events_Archive {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Events_Archive) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Events_Archive {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Events_Archive) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Events_Archive {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Events_Archive) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Events_Archive {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Events_Archive) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Events_Archive {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Events_Archive) Set__RequestedOutputs(v []cfz.Output) *AWS_Events_Archive {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Events_Archive) Add__RequestedOutputs(v ...cfz.Output) *AWS_Events_Archive {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ArchiveName updates property "ArchiveName".
func (t *AWS_Events_Archive) Set__ArchiveName(v cfz.Expression[string]) *AWS_Events_Archive {
	t.ArchiveName = v
	return t
}

// SetV__ArchiveName updates property "ArchiveName".
func (t *AWS_Events_Archive) SetV__ArchiveName(v string) *AWS_Events_Archive {
	t.ArchiveName = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Events_Archive) Set__Description(v cfz.Expression[string]) *AWS_Events_Archive {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Events_Archive) SetV__Description(v string) *AWS_Events_Archive {
	t.Description = cfz.V(v)
	return t
}

// Set__EventPattern updates property "EventPattern".
func (t *AWS_Events_Archive) Set__EventPattern(v cfz.Expression[json.RawMessage]) *AWS_Events_Archive {
	t.EventPattern = v
	return t
}

// SetV__EventPattern updates property "EventPattern".
func (t *AWS_Events_Archive) SetV__EventPattern(v json.RawMessage) *AWS_Events_Archive {
	t.EventPattern = cfz.V(v)
	return t
}

// Set__RetentionDays updates property "RetentionDays".
func (t *AWS_Events_Archive) Set__RetentionDays(v cfz.Expression[int32]) *AWS_Events_Archive {
	t.RetentionDays = v
	return t
}

// SetV__RetentionDays updates property "RetentionDays".
func (t *AWS_Events_Archive) SetV__RetentionDays(v int32) *AWS_Events_Archive {
	t.RetentionDays = cfz.V(v)
	return t
}

// Set__SourceArn updates property "SourceArn".
func (t *AWS_Events_Archive) Set__SourceArn(v cfz.Expression[string]) *AWS_Events_Archive {
	t.SourceArn = v
	return t
}

// SetV__SourceArn updates property "SourceArn".
func (t *AWS_Events_Archive) SetV__SourceArn(v string) *AWS_Events_Archive {
	t.SourceArn = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Events_Archive) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Events_Archive) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Events_Archive__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Events_Archive) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Events_Archive) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Events_Archive) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Events_Archive

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

func (t *AWS_Events_Archive) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
