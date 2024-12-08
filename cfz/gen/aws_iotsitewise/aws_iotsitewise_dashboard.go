// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotsitewise

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_IoTSiteWise_Dashboard)(nil)
	_ cfz.Resource                   = (*AWS_IoTSiteWise_Dashboard)(nil)
)

const (
	// AWS_IoTSiteWise_Dashboard__Type is the CloudFormation type for AWS::IoTSiteWise::Dashboard.
	AWS_IoTSiteWise_Dashboard__Type = "AWS::IoTSiteWise::Dashboard"
)

var (
	// AWS_IoTSiteWise_Dashboard__AttributesMap reports all the CloudFormation attributes for AWS::IoTSiteWise::Dashboard.
	AWS_IoTSiteWise_Dashboard__AttributesMap = struct {
		DashboardArn string
		DashboardId  string
	}{
		DashboardArn: "DashboardArn",
		DashboardId:  "DashboardId",
	}

	// AWS_IoTSiteWise_Dashboard__AttributesSlice reports all the CloudFormation attributes for AWS::IoTSiteWise::Dashboard.
	AWS_IoTSiteWise_Dashboard__AttributesSlice = []string{
		AWS_IoTSiteWise_Dashboard__AttributesMap.DashboardArn,
		AWS_IoTSiteWise_Dashboard__AttributesMap.DashboardId,
	}
)

var (
	// AWS_IoTSiteWise_Dashboard__PropertiesMap reports all the CloudFormation properties for AWS::IoTSiteWise::Dashboard.
	AWS_IoTSiteWise_Dashboard__PropertiesMap = struct {
		DashboardDefinition  string
		DashboardDescription string
		DashboardName        string
		ProjectId            string
		Tags                 string
	}{
		DashboardDefinition:  "DashboardDefinition",
		DashboardDescription: "DashboardDescription",
		DashboardName:        "DashboardName",
		ProjectId:            "ProjectId",
		Tags:                 "Tags",
	}

	// AWS_IoTSiteWise_Dashboard__PropertiesSlice reports all the CloudFormation properties for AWS::IoTSiteWise::Dashboard.
	AWS_IoTSiteWise_Dashboard__PropertiesSlice = []string{
		AWS_IoTSiteWise_Dashboard__PropertiesMap.DashboardDefinition,
		AWS_IoTSiteWise_Dashboard__PropertiesMap.DashboardDescription,
		AWS_IoTSiteWise_Dashboard__PropertiesMap.DashboardName,
		AWS_IoTSiteWise_Dashboard__PropertiesMap.ProjectId,
		AWS_IoTSiteWise_Dashboard__PropertiesMap.Tags,
	}
)

// AWS_IoTSiteWise_Dashboard is a binding for AWS::IoTSiteWise::Dashboard.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html
type AWS_IoTSiteWise_Dashboard struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DashboardDefinition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-dashboarddefinition
	DashboardDefinition cfz.Expression[string] `json:"DashboardDefinition,omitempty"`

	// DashboardDescription is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-dashboarddescription
	DashboardDescription cfz.Expression[string] `json:"DashboardDescription,omitempty"`

	// DashboardName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-dashboardname
	DashboardName cfz.Expression[string] `json:"DashboardName,omitempty"`

	// ProjectId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-projectid
	ProjectId cfz.Expression[string] `json:"ProjectId,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dashboard.html#cfn-iotsitewise-dashboard-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_IoTSiteWise_Dashboard initializes a new *AWS_IoTSiteWise_Dashboard.
func New__AWS_IoTSiteWise_Dashboard(logicalName string) *AWS_IoTSiteWise_Dashboard {
	return &AWS_IoTSiteWise_Dashboard{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_IoTSiteWise_Dashboard) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_IoTSiteWise_Dashboard) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_IoTSiteWise_Dashboard) GetType() string {
	return AWS_IoTSiteWise_Dashboard__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_IoTSiteWise_Dashboard) Set__LogicalName(v string) *AWS_IoTSiteWise_Dashboard {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_IoTSiteWise_Dashboard) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_IoTSiteWise_Dashboard {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_IoTSiteWise_Dashboard) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_IoTSiteWise_Dashboard {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_IoTSiteWise_Dashboard) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_IoTSiteWise_Dashboard {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_IoTSiteWise_Dashboard) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_IoTSiteWise_Dashboard {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_IoTSiteWise_Dashboard) Set__RequestedOutputs(v []cfz.Output) *AWS_IoTSiteWise_Dashboard {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_IoTSiteWise_Dashboard) Add__RequestedOutputs(v ...cfz.Output) *AWS_IoTSiteWise_Dashboard {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DashboardDefinition updates property "DashboardDefinition".
func (t *AWS_IoTSiteWise_Dashboard) Set__DashboardDefinition(v cfz.Expression[string]) *AWS_IoTSiteWise_Dashboard {
	t.DashboardDefinition = v
	return t
}

// SetV__DashboardDefinition updates property "DashboardDefinition".
func (t *AWS_IoTSiteWise_Dashboard) SetV__DashboardDefinition(v string) *AWS_IoTSiteWise_Dashboard {
	t.DashboardDefinition = cfz.V(v)
	return t
}

// Set__DashboardDescription updates property "DashboardDescription".
func (t *AWS_IoTSiteWise_Dashboard) Set__DashboardDescription(v cfz.Expression[string]) *AWS_IoTSiteWise_Dashboard {
	t.DashboardDescription = v
	return t
}

// SetV__DashboardDescription updates property "DashboardDescription".
func (t *AWS_IoTSiteWise_Dashboard) SetV__DashboardDescription(v string) *AWS_IoTSiteWise_Dashboard {
	t.DashboardDescription = cfz.V(v)
	return t
}

// Set__DashboardName updates property "DashboardName".
func (t *AWS_IoTSiteWise_Dashboard) Set__DashboardName(v cfz.Expression[string]) *AWS_IoTSiteWise_Dashboard {
	t.DashboardName = v
	return t
}

// SetV__DashboardName updates property "DashboardName".
func (t *AWS_IoTSiteWise_Dashboard) SetV__DashboardName(v string) *AWS_IoTSiteWise_Dashboard {
	t.DashboardName = cfz.V(v)
	return t
}

// Set__ProjectId updates property "ProjectId".
func (t *AWS_IoTSiteWise_Dashboard) Set__ProjectId(v cfz.Expression[string]) *AWS_IoTSiteWise_Dashboard {
	t.ProjectId = v
	return t
}

// SetV__ProjectId updates property "ProjectId".
func (t *AWS_IoTSiteWise_Dashboard) SetV__ProjectId(v string) *AWS_IoTSiteWise_Dashboard {
	t.ProjectId = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_IoTSiteWise_Dashboard) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_IoTSiteWise_Dashboard {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_IoTSiteWise_Dashboard) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_IoTSiteWise_Dashboard {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_IoTSiteWise_Dashboard) SetSV__Tags(v ...cfz.Tag) *AWS_IoTSiteWise_Dashboard {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_IoTSiteWise_Dashboard) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__DashboardArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DashboardArn
func (t *AWS_IoTSiteWise_Dashboard) GetAtt__DashboardArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IoTSiteWise_Dashboard__AttributesMap.DashboardArn))
}

// GetAtt__DashboardId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DashboardId
func (t *AWS_IoTSiteWise_Dashboard) GetAtt__DashboardId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IoTSiteWise_Dashboard__AttributesMap.DashboardId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_IoTSiteWise_Dashboard) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DashboardArn returns a conventionally configured output for an attribute of this resource.
// Attribute: DashboardArn
func (t *AWS_IoTSiteWise_Dashboard) GetConventionalOutputAtt__DashboardArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDashboardArn", t.GetAtt__DashboardArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DashboardId returns a conventionally configured output for an attribute of this resource.
// Attribute: DashboardId
func (t *AWS_IoTSiteWise_Dashboard) GetConventionalOutputAtt__DashboardId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDashboardId", t.GetAtt__DashboardId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_IoTSiteWise_Dashboard) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_IoTSiteWise_Dashboard

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

func (t *AWS_IoTSiteWise_Dashboard) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
