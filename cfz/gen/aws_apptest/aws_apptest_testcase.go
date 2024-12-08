// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apptest

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_AppTest_TestCase)(nil)
	_ cfz.Resource                   = (*AWS_AppTest_TestCase)(nil)
)

const (
	// AWS_AppTest_TestCase__Type is the CloudFormation type for AWS::AppTest::TestCase.
	AWS_AppTest_TestCase__Type = "AWS::AppTest::TestCase"
)

var (
	// AWS_AppTest_TestCase__AttributesMap reports all the CloudFormation attributes for AWS::AppTest::TestCase.
	AWS_AppTest_TestCase__AttributesMap = struct {
		CreationTime          string
		LastUpdateTime        string
		LatestVersion         string
		LatestVersion_Status  string
		LatestVersion_Version string
		Status                string
		TestCaseArn           string
		TestCaseId            string
		TestCaseVersion       string
	}{
		CreationTime:          "CreationTime",
		LastUpdateTime:        "LastUpdateTime",
		LatestVersion:         "LatestVersion",
		LatestVersion_Status:  "LatestVersion.Status",
		LatestVersion_Version: "LatestVersion.Version",
		Status:                "Status",
		TestCaseArn:           "TestCaseArn",
		TestCaseId:            "TestCaseId",
		TestCaseVersion:       "TestCaseVersion",
	}

	// AWS_AppTest_TestCase__AttributesSlice reports all the CloudFormation attributes for AWS::AppTest::TestCase.
	AWS_AppTest_TestCase__AttributesSlice = []string{
		AWS_AppTest_TestCase__AttributesMap.CreationTime,
		AWS_AppTest_TestCase__AttributesMap.LastUpdateTime,
		AWS_AppTest_TestCase__AttributesMap.LatestVersion,
		AWS_AppTest_TestCase__AttributesMap.LatestVersion_Status,
		AWS_AppTest_TestCase__AttributesMap.LatestVersion_Version,
		AWS_AppTest_TestCase__AttributesMap.Status,
		AWS_AppTest_TestCase__AttributesMap.TestCaseArn,
		AWS_AppTest_TestCase__AttributesMap.TestCaseId,
		AWS_AppTest_TestCase__AttributesMap.TestCaseVersion,
	}
)

var (
	// AWS_AppTest_TestCase__PropertiesMap reports all the CloudFormation properties for AWS::AppTest::TestCase.
	AWS_AppTest_TestCase__PropertiesMap = struct {
		Description string
		Name        string
		Steps       string
		Tags        string
	}{
		Description: "Description",
		Name:        "Name",
		Steps:       "Steps",
		Tags:        "Tags",
	}

	// AWS_AppTest_TestCase__PropertiesSlice reports all the CloudFormation properties for AWS::AppTest::TestCase.
	AWS_AppTest_TestCase__PropertiesSlice = []string{
		AWS_AppTest_TestCase__PropertiesMap.Description,
		AWS_AppTest_TestCase__PropertiesMap.Name,
		AWS_AppTest_TestCase__PropertiesMap.Steps,
		AWS_AppTest_TestCase__PropertiesMap.Tags,
	}
)

// AWS_AppTest_TestCase is a binding for AWS::AppTest::TestCase.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html
type AWS_AppTest_TestCase struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html#cfn-apptest-testcase-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html#cfn-apptest-testcase-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Steps is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html#cfn-apptest-testcase-steps
	Steps cfz.ExpressionSlice[AWS_AppTest_TestCase_Step] `json:"Steps,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apptest-testcase.html#cfn-apptest-testcase-tags
	Tags cfz.ExpressionMap[string] `json:"Tags,omitempty"`
}

// New__AWS_AppTest_TestCase initializes a new *AWS_AppTest_TestCase.
func New__AWS_AppTest_TestCase(logicalName string) *AWS_AppTest_TestCase {
	return &AWS_AppTest_TestCase{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_AppTest_TestCase) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_AppTest_TestCase) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_AppTest_TestCase) GetType() string {
	return AWS_AppTest_TestCase__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_AppTest_TestCase) Set__LogicalName(v string) *AWS_AppTest_TestCase {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_AppTest_TestCase) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_AppTest_TestCase {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_AppTest_TestCase) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_AppTest_TestCase {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_AppTest_TestCase) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_AppTest_TestCase {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_AppTest_TestCase) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_AppTest_TestCase {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_AppTest_TestCase) Set__RequestedOutputs(v []cfz.Output) *AWS_AppTest_TestCase {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_AppTest_TestCase) Add__RequestedOutputs(v ...cfz.Output) *AWS_AppTest_TestCase {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_AppTest_TestCase) Set__Description(v cfz.Expression[string]) *AWS_AppTest_TestCase {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_AppTest_TestCase) SetV__Description(v string) *AWS_AppTest_TestCase {
	t.Description = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_AppTest_TestCase) Set__Name(v cfz.Expression[string]) *AWS_AppTest_TestCase {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_AppTest_TestCase) SetV__Name(v string) *AWS_AppTest_TestCase {
	t.Name = cfz.V(v)
	return t
}

// Set__Steps updates property "Steps".
func (t *AWS_AppTest_TestCase) Set__Steps(v cfz.ExpressionSlice[AWS_AppTest_TestCase_Step]) *AWS_AppTest_TestCase {
	t.Steps = v
	return t
}

// SetS__Steps updates property "Steps".
func (t *AWS_AppTest_TestCase) SetS__Steps(v ...cfz.Expression[AWS_AppTest_TestCase_Step]) *AWS_AppTest_TestCase {
	t.Steps = cfz.S(v...)
	return t
}

// SetSV__Steps updates property "Steps".
func (t *AWS_AppTest_TestCase) SetSV__Steps(v ...AWS_AppTest_TestCase_Step) *AWS_AppTest_TestCase {
	t.Steps = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_AppTest_TestCase) Set__Tags(v cfz.ExpressionMap[string]) *AWS_AppTest_TestCase {
	t.Tags = v
	return t
}

// SetM__Tags updates property "Tags".
func (t *AWS_AppTest_TestCase) SetM__Tags(v ...map[string]cfz.Expression[string]) *AWS_AppTest_TestCase {
	t.Tags = cfz.M(v...)
	return t
}

// SetMV__Tags updates property "Tags".
func (t *AWS_AppTest_TestCase) SetMV__Tags(v ...map[string]string) *AWS_AppTest_TestCase {
	t.Tags = cfz.MV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_AppTest_TestCase) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CreationTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreationTime
func (t *AWS_AppTest_TestCase) GetAtt__CreationTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppTest_TestCase__AttributesMap.CreationTime))
}

// GetAtt__LastUpdateTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LastUpdateTime
func (t *AWS_AppTest_TestCase) GetAtt__LastUpdateTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppTest_TestCase__AttributesMap.LastUpdateTime))
}

// GetAtt__LatestVersion returns a $cfz.Expression[AWS_AppTest_TestCase_TestCaseLatestVersion] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LatestVersion
func (t *AWS_AppTest_TestCase) GetAtt__LatestVersion() cfz.Expression[AWS_AppTest_TestCase_TestCaseLatestVersion] {
	return cfz.GetAtt[AWS_AppTest_TestCase_TestCaseLatestVersion](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppTest_TestCase__AttributesMap.LatestVersion))
}

// GetAtt__LatestVersion_Status returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LatestVersion.Status
func (t *AWS_AppTest_TestCase) GetAtt__LatestVersion_Status() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppTest_TestCase__AttributesMap.LatestVersion_Status))
}

// GetAtt__LatestVersion_Version returns a $cfz.Expression[float64] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LatestVersion.Version
func (t *AWS_AppTest_TestCase) GetAtt__LatestVersion_Version() cfz.Expression[float64] {
	return cfz.GetAtt[float64](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppTest_TestCase__AttributesMap.LatestVersion_Version))
}

// GetAtt__Status returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Status
func (t *AWS_AppTest_TestCase) GetAtt__Status() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppTest_TestCase__AttributesMap.Status))
}

// GetAtt__TestCaseArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TestCaseArn
func (t *AWS_AppTest_TestCase) GetAtt__TestCaseArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppTest_TestCase__AttributesMap.TestCaseArn))
}

// GetAtt__TestCaseId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TestCaseId
func (t *AWS_AppTest_TestCase) GetAtt__TestCaseId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppTest_TestCase__AttributesMap.TestCaseId))
}

// GetAtt__TestCaseVersion returns a $cfz.Expression[float64] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TestCaseVersion
func (t *AWS_AppTest_TestCase) GetAtt__TestCaseVersion() cfz.Expression[float64] {
	return cfz.GetAtt[float64](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppTest_TestCase__AttributesMap.TestCaseVersion))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_AppTest_TestCase) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreationTime returns a conventionally configured output for an attribute of this resource.
// Attribute: CreationTime
func (t *AWS_AppTest_TestCase) GetConventionalOutputAtt__CreationTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreationTime", t.GetAtt__CreationTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LastUpdateTime returns a conventionally configured output for an attribute of this resource.
// Attribute: LastUpdateTime
func (t *AWS_AppTest_TestCase) GetConventionalOutputAtt__LastUpdateTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLastUpdateTime", t.GetAtt__LastUpdateTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LatestVersion returns a conventionally configured output for an attribute of this resource.
// Attribute: LatestVersion
func (t *AWS_AppTest_TestCase) GetConventionalOutputAtt__LatestVersion(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLatestVersion", t.GetAtt__LatestVersion())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LatestVersion_Status returns a conventionally configured output for an attribute of this resource.
// Attribute: LatestVersion.Status
func (t *AWS_AppTest_TestCase) GetConventionalOutputAtt__LatestVersion_Status(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLatestVersionStatus", t.GetAtt__LatestVersion_Status())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LatestVersion_Version returns a conventionally configured output for an attribute of this resource.
// Attribute: LatestVersion.Version
func (t *AWS_AppTest_TestCase) GetConventionalOutputAtt__LatestVersion_Version(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLatestVersionVersion", t.GetAtt__LatestVersion_Version())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Status returns a conventionally configured output for an attribute of this resource.
// Attribute: Status
func (t *AWS_AppTest_TestCase) GetConventionalOutputAtt__Status(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatus", t.GetAtt__Status())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TestCaseArn returns a conventionally configured output for an attribute of this resource.
// Attribute: TestCaseArn
func (t *AWS_AppTest_TestCase) GetConventionalOutputAtt__TestCaseArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTestCaseArn", t.GetAtt__TestCaseArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TestCaseId returns a conventionally configured output for an attribute of this resource.
// Attribute: TestCaseId
func (t *AWS_AppTest_TestCase) GetConventionalOutputAtt__TestCaseId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTestCaseId", t.GetAtt__TestCaseId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TestCaseVersion returns a conventionally configured output for an attribute of this resource.
// Attribute: TestCaseVersion
func (t *AWS_AppTest_TestCase) GetConventionalOutputAtt__TestCaseVersion(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTestCaseVersion", t.GetAtt__TestCaseVersion())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_AppTest_TestCase) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_AppTest_TestCase

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

func (t *AWS_AppTest_TestCase) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
