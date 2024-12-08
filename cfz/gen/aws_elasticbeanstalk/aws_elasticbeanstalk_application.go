// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_elasticbeanstalk

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ElasticBeanstalk_Application)(nil)
	_ cfz.Resource                   = (*AWS_ElasticBeanstalk_Application)(nil)
)

const (
	// AWS_ElasticBeanstalk_Application__Type is the CloudFormation type for AWS::ElasticBeanstalk::Application.
	AWS_ElasticBeanstalk_Application__Type = "AWS::ElasticBeanstalk::Application"
)

var (
	// AWS_ElasticBeanstalk_Application__PropertiesMap reports all the CloudFormation properties for AWS::ElasticBeanstalk::Application.
	AWS_ElasticBeanstalk_Application__PropertiesMap = struct {
		ApplicationName         string
		Description             string
		ResourceLifecycleConfig string
	}{
		ApplicationName:         "ApplicationName",
		Description:             "Description",
		ResourceLifecycleConfig: "ResourceLifecycleConfig",
	}

	// AWS_ElasticBeanstalk_Application__PropertiesSlice reports all the CloudFormation properties for AWS::ElasticBeanstalk::Application.
	AWS_ElasticBeanstalk_Application__PropertiesSlice = []string{
		AWS_ElasticBeanstalk_Application__PropertiesMap.ApplicationName,
		AWS_ElasticBeanstalk_Application__PropertiesMap.Description,
		AWS_ElasticBeanstalk_Application__PropertiesMap.ResourceLifecycleConfig,
	}
)

// AWS_ElasticBeanstalk_Application is a binding for AWS::ElasticBeanstalk::Application.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-application.html
type AWS_ElasticBeanstalk_Application struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ApplicationName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-application.html#cfn-elasticbeanstalk-application-applicationname
	ApplicationName cfz.Expression[string] `json:"ApplicationName,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-application.html#cfn-elasticbeanstalk-application-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// ResourceLifecycleConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-application.html#cfn-elasticbeanstalk-application-resourcelifecycleconfig
	ResourceLifecycleConfig cfz.Expression[AWS_ElasticBeanstalk_Application_ApplicationResourceLifecycleConfig] `json:"ResourceLifecycleConfig,omitempty"`
}

// New__AWS_ElasticBeanstalk_Application initializes a new *AWS_ElasticBeanstalk_Application.
func New__AWS_ElasticBeanstalk_Application(logicalName string) *AWS_ElasticBeanstalk_Application {
	return &AWS_ElasticBeanstalk_Application{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ElasticBeanstalk_Application) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ElasticBeanstalk_Application) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ElasticBeanstalk_Application) GetType() string {
	return AWS_ElasticBeanstalk_Application__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ElasticBeanstalk_Application) Set__LogicalName(v string) *AWS_ElasticBeanstalk_Application {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ElasticBeanstalk_Application) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ElasticBeanstalk_Application {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ElasticBeanstalk_Application) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ElasticBeanstalk_Application {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ElasticBeanstalk_Application) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ElasticBeanstalk_Application {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ElasticBeanstalk_Application) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ElasticBeanstalk_Application {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ElasticBeanstalk_Application) Set__RequestedOutputs(v []cfz.Output) *AWS_ElasticBeanstalk_Application {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ElasticBeanstalk_Application) Add__RequestedOutputs(v ...cfz.Output) *AWS_ElasticBeanstalk_Application {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ApplicationName updates property "ApplicationName".
func (t *AWS_ElasticBeanstalk_Application) Set__ApplicationName(v cfz.Expression[string]) *AWS_ElasticBeanstalk_Application {
	t.ApplicationName = v
	return t
}

// SetV__ApplicationName updates property "ApplicationName".
func (t *AWS_ElasticBeanstalk_Application) SetV__ApplicationName(v string) *AWS_ElasticBeanstalk_Application {
	t.ApplicationName = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_ElasticBeanstalk_Application) Set__Description(v cfz.Expression[string]) *AWS_ElasticBeanstalk_Application {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_ElasticBeanstalk_Application) SetV__Description(v string) *AWS_ElasticBeanstalk_Application {
	t.Description = cfz.V(v)
	return t
}

// Set__ResourceLifecycleConfig updates property "ResourceLifecycleConfig".
func (t *AWS_ElasticBeanstalk_Application) Set__ResourceLifecycleConfig(v cfz.Expression[AWS_ElasticBeanstalk_Application_ApplicationResourceLifecycleConfig]) *AWS_ElasticBeanstalk_Application {
	t.ResourceLifecycleConfig = v
	return t
}

// SetV__ResourceLifecycleConfig updates property "ResourceLifecycleConfig".
func (t *AWS_ElasticBeanstalk_Application) SetV__ResourceLifecycleConfig(v AWS_ElasticBeanstalk_Application_ApplicationResourceLifecycleConfig) *AWS_ElasticBeanstalk_Application {
	t.ResourceLifecycleConfig = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ElasticBeanstalk_Application) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ElasticBeanstalk_Application) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ElasticBeanstalk_Application) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ElasticBeanstalk_Application

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

func (t *AWS_ElasticBeanstalk_Application) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
