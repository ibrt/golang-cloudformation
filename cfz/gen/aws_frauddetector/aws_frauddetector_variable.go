// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_frauddetector

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_FraudDetector_Variable)(nil)
	_ cfz.Resource                   = (*AWS_FraudDetector_Variable)(nil)
)

const (
	// AWS_FraudDetector_Variable__Type is the CloudFormation type for AWS::FraudDetector::Variable.
	AWS_FraudDetector_Variable__Type = "AWS::FraudDetector::Variable"
)

var (
	// AWS_FraudDetector_Variable__AttributesMap reports all the CloudFormation attributes for AWS::FraudDetector::Variable.
	AWS_FraudDetector_Variable__AttributesMap = struct {
		Arn             string
		CreatedTime     string
		LastUpdatedTime string
	}{
		Arn:             "Arn",
		CreatedTime:     "CreatedTime",
		LastUpdatedTime: "LastUpdatedTime",
	}

	// AWS_FraudDetector_Variable__AttributesSlice reports all the CloudFormation attributes for AWS::FraudDetector::Variable.
	AWS_FraudDetector_Variable__AttributesSlice = []string{
		AWS_FraudDetector_Variable__AttributesMap.Arn,
		AWS_FraudDetector_Variable__AttributesMap.CreatedTime,
		AWS_FraudDetector_Variable__AttributesMap.LastUpdatedTime,
	}
)

var (
	// AWS_FraudDetector_Variable__PropertiesMap reports all the CloudFormation properties for AWS::FraudDetector::Variable.
	AWS_FraudDetector_Variable__PropertiesMap = struct {
		DataSource   string
		DataType     string
		DefaultValue string
		Description  string
		Name         string
		Tags         string
		VariableType string
	}{
		DataSource:   "DataSource",
		DataType:     "DataType",
		DefaultValue: "DefaultValue",
		Description:  "Description",
		Name:         "Name",
		Tags:         "Tags",
		VariableType: "VariableType",
	}

	// AWS_FraudDetector_Variable__PropertiesSlice reports all the CloudFormation properties for AWS::FraudDetector::Variable.
	AWS_FraudDetector_Variable__PropertiesSlice = []string{
		AWS_FraudDetector_Variable__PropertiesMap.DataSource,
		AWS_FraudDetector_Variable__PropertiesMap.DataType,
		AWS_FraudDetector_Variable__PropertiesMap.DefaultValue,
		AWS_FraudDetector_Variable__PropertiesMap.Description,
		AWS_FraudDetector_Variable__PropertiesMap.Name,
		AWS_FraudDetector_Variable__PropertiesMap.Tags,
		AWS_FraudDetector_Variable__PropertiesMap.VariableType,
	}
)

// AWS_FraudDetector_Variable is a binding for AWS::FraudDetector::Variable.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-variable.html
type AWS_FraudDetector_Variable struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DataSource is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-variable.html#cfn-frauddetector-variable-datasource
	DataSource cfz.Expression[string] `json:"DataSource,omitempty"`

	// DataType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-variable.html#cfn-frauddetector-variable-datatype
	DataType cfz.Expression[string] `json:"DataType,omitempty"`

	// DefaultValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-variable.html#cfn-frauddetector-variable-defaultvalue
	DefaultValue cfz.Expression[string] `json:"DefaultValue,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-variable.html#cfn-frauddetector-variable-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-variable.html#cfn-frauddetector-variable-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-variable.html#cfn-frauddetector-variable-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// VariableType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-variable.html#cfn-frauddetector-variable-variabletype
	VariableType cfz.Expression[string] `json:"VariableType,omitempty"`
}

// New__AWS_FraudDetector_Variable initializes a new *AWS_FraudDetector_Variable.
func New__AWS_FraudDetector_Variable(logicalName string) *AWS_FraudDetector_Variable {
	return &AWS_FraudDetector_Variable{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_FraudDetector_Variable) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_FraudDetector_Variable) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_FraudDetector_Variable) GetType() string {
	return AWS_FraudDetector_Variable__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_FraudDetector_Variable) Set__LogicalName(v string) *AWS_FraudDetector_Variable {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_FraudDetector_Variable) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_FraudDetector_Variable {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_FraudDetector_Variable) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_FraudDetector_Variable {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_FraudDetector_Variable) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_FraudDetector_Variable {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_FraudDetector_Variable) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_FraudDetector_Variable {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_FraudDetector_Variable) Set__RequestedOutputs(v []cfz.Output) *AWS_FraudDetector_Variable {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_FraudDetector_Variable) Add__RequestedOutputs(v ...cfz.Output) *AWS_FraudDetector_Variable {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DataSource updates property "DataSource".
func (t *AWS_FraudDetector_Variable) Set__DataSource(v cfz.Expression[string]) *AWS_FraudDetector_Variable {
	t.DataSource = v
	return t
}

// SetV__DataSource updates property "DataSource".
func (t *AWS_FraudDetector_Variable) SetV__DataSource(v string) *AWS_FraudDetector_Variable {
	t.DataSource = cfz.V(v)
	return t
}

// Set__DataType updates property "DataType".
func (t *AWS_FraudDetector_Variable) Set__DataType(v cfz.Expression[string]) *AWS_FraudDetector_Variable {
	t.DataType = v
	return t
}

// SetV__DataType updates property "DataType".
func (t *AWS_FraudDetector_Variable) SetV__DataType(v string) *AWS_FraudDetector_Variable {
	t.DataType = cfz.V(v)
	return t
}

// Set__DefaultValue updates property "DefaultValue".
func (t *AWS_FraudDetector_Variable) Set__DefaultValue(v cfz.Expression[string]) *AWS_FraudDetector_Variable {
	t.DefaultValue = v
	return t
}

// SetV__DefaultValue updates property "DefaultValue".
func (t *AWS_FraudDetector_Variable) SetV__DefaultValue(v string) *AWS_FraudDetector_Variable {
	t.DefaultValue = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_FraudDetector_Variable) Set__Description(v cfz.Expression[string]) *AWS_FraudDetector_Variable {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_FraudDetector_Variable) SetV__Description(v string) *AWS_FraudDetector_Variable {
	t.Description = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_FraudDetector_Variable) Set__Name(v cfz.Expression[string]) *AWS_FraudDetector_Variable {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_FraudDetector_Variable) SetV__Name(v string) *AWS_FraudDetector_Variable {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_FraudDetector_Variable) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_FraudDetector_Variable {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_FraudDetector_Variable) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_FraudDetector_Variable {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_FraudDetector_Variable) SetSV__Tags(v ...cfz.Tag) *AWS_FraudDetector_Variable {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__VariableType updates property "VariableType".
func (t *AWS_FraudDetector_Variable) Set__VariableType(v cfz.Expression[string]) *AWS_FraudDetector_Variable {
	t.VariableType = v
	return t
}

// SetV__VariableType updates property "VariableType".
func (t *AWS_FraudDetector_Variable) SetV__VariableType(v string) *AWS_FraudDetector_Variable {
	t.VariableType = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_FraudDetector_Variable) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_FraudDetector_Variable) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_FraudDetector_Variable__AttributesMap.Arn))
}

// GetAtt__CreatedTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreatedTime
func (t *AWS_FraudDetector_Variable) GetAtt__CreatedTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_FraudDetector_Variable__AttributesMap.CreatedTime))
}

// GetAtt__LastUpdatedTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LastUpdatedTime
func (t *AWS_FraudDetector_Variable) GetAtt__LastUpdatedTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_FraudDetector_Variable__AttributesMap.LastUpdatedTime))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_FraudDetector_Variable) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_FraudDetector_Variable) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreatedTime returns a conventionally configured output for an attribute of this resource.
// Attribute: CreatedTime
func (t *AWS_FraudDetector_Variable) GetConventionalOutputAtt__CreatedTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreatedTime", t.GetAtt__CreatedTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LastUpdatedTime returns a conventionally configured output for an attribute of this resource.
// Attribute: LastUpdatedTime
func (t *AWS_FraudDetector_Variable) GetConventionalOutputAtt__LastUpdatedTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLastUpdatedTime", t.GetAtt__LastUpdatedTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_FraudDetector_Variable) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_FraudDetector_Variable

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

func (t *AWS_FraudDetector_Variable) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
