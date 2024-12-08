// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_NetworkInsightsAccessScopeAnalysis)(nil)
	_ cfz.Resource                   = (*AWS_EC2_NetworkInsightsAccessScopeAnalysis)(nil)
)

const (
	// AWS_EC2_NetworkInsightsAccessScopeAnalysis__Type is the CloudFormation type for AWS::EC2::NetworkInsightsAccessScopeAnalysis.
	AWS_EC2_NetworkInsightsAccessScopeAnalysis__Type = "AWS::EC2::NetworkInsightsAccessScopeAnalysis"
)

var (
	// AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap reports all the CloudFormation attributes for AWS::EC2::NetworkInsightsAccessScopeAnalysis.
	AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap = struct {
		AnalyzedEniCount                      string
		EndDate                               string
		FindingsFound                         string
		NetworkInsightsAccessScopeAnalysisArn string
		NetworkInsightsAccessScopeAnalysisId  string
		StartDate                             string
		Status                                string
		StatusMessage                         string
	}{
		AnalyzedEniCount:                      "AnalyzedEniCount",
		EndDate:                               "EndDate",
		FindingsFound:                         "FindingsFound",
		NetworkInsightsAccessScopeAnalysisArn: "NetworkInsightsAccessScopeAnalysisArn",
		NetworkInsightsAccessScopeAnalysisId:  "NetworkInsightsAccessScopeAnalysisId",
		StartDate:                             "StartDate",
		Status:                                "Status",
		StatusMessage:                         "StatusMessage",
	}

	// AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesSlice reports all the CloudFormation attributes for AWS::EC2::NetworkInsightsAccessScopeAnalysis.
	AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesSlice = []string{
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.AnalyzedEniCount,
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.EndDate,
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.FindingsFound,
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.NetworkInsightsAccessScopeAnalysisArn,
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.NetworkInsightsAccessScopeAnalysisId,
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.StartDate,
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.Status,
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.StatusMessage,
	}
)

var (
	// AWS_EC2_NetworkInsightsAccessScopeAnalysis__PropertiesMap reports all the CloudFormation properties for AWS::EC2::NetworkInsightsAccessScopeAnalysis.
	AWS_EC2_NetworkInsightsAccessScopeAnalysis__PropertiesMap = struct {
		NetworkInsightsAccessScopeId string
		Tags                         string
	}{
		NetworkInsightsAccessScopeId: "NetworkInsightsAccessScopeId",
		Tags:                         "Tags",
	}

	// AWS_EC2_NetworkInsightsAccessScopeAnalysis__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::NetworkInsightsAccessScopeAnalysis.
	AWS_EC2_NetworkInsightsAccessScopeAnalysis__PropertiesSlice = []string{
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__PropertiesMap.NetworkInsightsAccessScopeId,
		AWS_EC2_NetworkInsightsAccessScopeAnalysis__PropertiesMap.Tags,
	}
)

// AWS_EC2_NetworkInsightsAccessScopeAnalysis is a binding for AWS::EC2::NetworkInsightsAccessScopeAnalysis.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightsaccessscopeanalysis.html
type AWS_EC2_NetworkInsightsAccessScopeAnalysis struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// NetworkInsightsAccessScopeId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightsaccessscopeanalysis.html#cfn-ec2-networkinsightsaccessscopeanalysis-networkinsightsaccessscopeid
	NetworkInsightsAccessScopeId cfz.Expression[string] `json:"NetworkInsightsAccessScopeId,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightsaccessscopeanalysis.html#cfn-ec2-networkinsightsaccessscopeanalysis-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_EC2_NetworkInsightsAccessScopeAnalysis initializes a new *AWS_EC2_NetworkInsightsAccessScopeAnalysis.
func New__AWS_EC2_NetworkInsightsAccessScopeAnalysis(logicalName string) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	return &AWS_EC2_NetworkInsightsAccessScopeAnalysis{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetType() string {
	return AWS_EC2_NetworkInsightsAccessScopeAnalysis__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Set__LogicalName(v string) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__NetworkInsightsAccessScopeId updates property "NetworkInsightsAccessScopeId".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Set__NetworkInsightsAccessScopeId(v cfz.Expression[string]) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.NetworkInsightsAccessScopeId = v
	return t
}

// SetV__NetworkInsightsAccessScopeId updates property "NetworkInsightsAccessScopeId".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) SetV__NetworkInsightsAccessScopeId(v string) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.NetworkInsightsAccessScopeId = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) SetSV__Tags(v ...cfz.Tag) *AWS_EC2_NetworkInsightsAccessScopeAnalysis {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__AnalyzedEniCount returns a $cfz.Expression[int32] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AnalyzedEniCount
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetAtt__AnalyzedEniCount() cfz.Expression[int32] {
	return cfz.GetAtt[int32](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.AnalyzedEniCount))
}

// GetAtt__EndDate returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: EndDate
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetAtt__EndDate() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.EndDate))
}

// GetAtt__FindingsFound returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: FindingsFound
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetAtt__FindingsFound() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.FindingsFound))
}

// GetAtt__NetworkInsightsAccessScopeAnalysisArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: NetworkInsightsAccessScopeAnalysisArn
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetAtt__NetworkInsightsAccessScopeAnalysisArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.NetworkInsightsAccessScopeAnalysisArn))
}

// GetAtt__NetworkInsightsAccessScopeAnalysisId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: NetworkInsightsAccessScopeAnalysisId
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetAtt__NetworkInsightsAccessScopeAnalysisId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.NetworkInsightsAccessScopeAnalysisId))
}

// GetAtt__StartDate returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: StartDate
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetAtt__StartDate() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.StartDate))
}

// GetAtt__Status returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Status
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetAtt__Status() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.Status))
}

// GetAtt__StatusMessage returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: StatusMessage
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetAtt__StatusMessage() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_NetworkInsightsAccessScopeAnalysis__AttributesMap.StatusMessage))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AnalyzedEniCount returns a conventionally configured output for an attribute of this resource.
// Attribute: AnalyzedEniCount
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetConventionalOutputAtt__AnalyzedEniCount(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAnalyzedEniCount", t.GetAtt__AnalyzedEniCount())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__EndDate returns a conventionally configured output for an attribute of this resource.
// Attribute: EndDate
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetConventionalOutputAtt__EndDate(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttEndDate", t.GetAtt__EndDate())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__FindingsFound returns a conventionally configured output for an attribute of this resource.
// Attribute: FindingsFound
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetConventionalOutputAtt__FindingsFound(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFindingsFound", t.GetAtt__FindingsFound())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__NetworkInsightsAccessScopeAnalysisArn returns a conventionally configured output for an attribute of this resource.
// Attribute: NetworkInsightsAccessScopeAnalysisArn
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetConventionalOutputAtt__NetworkInsightsAccessScopeAnalysisArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttNetworkInsightsAccessScopeAnalysisArn", t.GetAtt__NetworkInsightsAccessScopeAnalysisArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__NetworkInsightsAccessScopeAnalysisId returns a conventionally configured output for an attribute of this resource.
// Attribute: NetworkInsightsAccessScopeAnalysisId
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetConventionalOutputAtt__NetworkInsightsAccessScopeAnalysisId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttNetworkInsightsAccessScopeAnalysisId", t.GetAtt__NetworkInsightsAccessScopeAnalysisId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__StartDate returns a conventionally configured output for an attribute of this resource.
// Attribute: StartDate
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetConventionalOutputAtt__StartDate(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStartDate", t.GetAtt__StartDate())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Status returns a conventionally configured output for an attribute of this resource.
// Attribute: Status
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetConventionalOutputAtt__Status(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatus", t.GetAtt__Status())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__StatusMessage returns a conventionally configured output for an attribute of this resource.
// Attribute: StatusMessage
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) GetConventionalOutputAtt__StatusMessage(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatusMessage", t.GetAtt__StatusMessage())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_NetworkInsightsAccessScopeAnalysis

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

func (t *AWS_EC2_NetworkInsightsAccessScopeAnalysis) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
