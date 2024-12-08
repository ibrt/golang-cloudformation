// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_opsworks

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_OpsWorks_Layer_Recipes__Type is the CloudFormation type for AWS::OpsWorks::Layer.Recipes.
	AWS_OpsWorks_Layer_Recipes__Type = "AWS::OpsWorks::Layer.Recipes"
)

var (
	// AWS_OpsWorks_Layer_Recipes__PropertiesMap reports all the CloudFormation properties for AWS::OpsWorks::Layer.Recipes.
	AWS_OpsWorks_Layer_Recipes__PropertiesMap = struct {
		Configure string
		Deploy    string
		Setup     string
		Shutdown  string
		Undeploy  string
	}{
		Configure: "Configure",
		Deploy:    "Deploy",
		Setup:     "Setup",
		Shutdown:  "Shutdown",
		Undeploy:  "Undeploy",
	}

	// AWS_OpsWorks_Layer_Recipes__PropertiesSlice reports all the CloudFormation properties for AWS::OpsWorks::Layer.Recipes.
	AWS_OpsWorks_Layer_Recipes__PropertiesSlice = []string{
		AWS_OpsWorks_Layer_Recipes__PropertiesMap.Configure,
		AWS_OpsWorks_Layer_Recipes__PropertiesMap.Deploy,
		AWS_OpsWorks_Layer_Recipes__PropertiesMap.Setup,
		AWS_OpsWorks_Layer_Recipes__PropertiesMap.Shutdown,
		AWS_OpsWorks_Layer_Recipes__PropertiesMap.Undeploy,
	}
)

// AWS_OpsWorks_Layer_Recipes is a binding for AWS::OpsWorks::Layer.Recipes.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html
type AWS_OpsWorks_Layer_Recipes struct {
	// Configure is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-configure
	Configure cfz.ExpressionSlice[string] `json:"Configure,omitempty"`

	// Deploy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-deploy
	Deploy cfz.ExpressionSlice[string] `json:"Deploy,omitempty"`

	// Setup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-setup
	Setup cfz.ExpressionSlice[string] `json:"Setup,omitempty"`

	// Shutdown is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-shutdown
	Shutdown cfz.ExpressionSlice[string] `json:"Shutdown,omitempty"`

	// Undeploy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-undeploy
	Undeploy cfz.ExpressionSlice[string] `json:"Undeploy,omitempty"`
}

// New__AWS_OpsWorks_Layer_Recipes initializes a new AWS_OpsWorks_Layer_Recipes.
func New__AWS_OpsWorks_Layer_Recipes() AWS_OpsWorks_Layer_Recipes {
	return AWS_OpsWorks_Layer_Recipes{}
}

// GetType returns the CloudFormation type.
func (AWS_OpsWorks_Layer_Recipes) GetType() string {
	return AWS_OpsWorks_Layer_Recipes__Type
}

// Set__Configure updates property "Configure".
func (t AWS_OpsWorks_Layer_Recipes) Set__Configure(v cfz.ExpressionSlice[string]) AWS_OpsWorks_Layer_Recipes {
	t.Configure = v
	return t
}

// SetS__Configure updates property "Configure".
func (t AWS_OpsWorks_Layer_Recipes) SetS__Configure(v ...cfz.Expression[string]) AWS_OpsWorks_Layer_Recipes {
	t.Configure = cfz.S(v...)
	return t
}

// SetSV__Configure updates property "Configure".
func (t AWS_OpsWorks_Layer_Recipes) SetSV__Configure(v ...string) AWS_OpsWorks_Layer_Recipes {
	t.Configure = cfz.SV(v...)
	return t
}

// Set__Deploy updates property "Deploy".
func (t AWS_OpsWorks_Layer_Recipes) Set__Deploy(v cfz.ExpressionSlice[string]) AWS_OpsWorks_Layer_Recipes {
	t.Deploy = v
	return t
}

// SetS__Deploy updates property "Deploy".
func (t AWS_OpsWorks_Layer_Recipes) SetS__Deploy(v ...cfz.Expression[string]) AWS_OpsWorks_Layer_Recipes {
	t.Deploy = cfz.S(v...)
	return t
}

// SetSV__Deploy updates property "Deploy".
func (t AWS_OpsWorks_Layer_Recipes) SetSV__Deploy(v ...string) AWS_OpsWorks_Layer_Recipes {
	t.Deploy = cfz.SV(v...)
	return t
}

// Set__Setup updates property "Setup".
func (t AWS_OpsWorks_Layer_Recipes) Set__Setup(v cfz.ExpressionSlice[string]) AWS_OpsWorks_Layer_Recipes {
	t.Setup = v
	return t
}

// SetS__Setup updates property "Setup".
func (t AWS_OpsWorks_Layer_Recipes) SetS__Setup(v ...cfz.Expression[string]) AWS_OpsWorks_Layer_Recipes {
	t.Setup = cfz.S(v...)
	return t
}

// SetSV__Setup updates property "Setup".
func (t AWS_OpsWorks_Layer_Recipes) SetSV__Setup(v ...string) AWS_OpsWorks_Layer_Recipes {
	t.Setup = cfz.SV(v...)
	return t
}

// Set__Shutdown updates property "Shutdown".
func (t AWS_OpsWorks_Layer_Recipes) Set__Shutdown(v cfz.ExpressionSlice[string]) AWS_OpsWorks_Layer_Recipes {
	t.Shutdown = v
	return t
}

// SetS__Shutdown updates property "Shutdown".
func (t AWS_OpsWorks_Layer_Recipes) SetS__Shutdown(v ...cfz.Expression[string]) AWS_OpsWorks_Layer_Recipes {
	t.Shutdown = cfz.S(v...)
	return t
}

// SetSV__Shutdown updates property "Shutdown".
func (t AWS_OpsWorks_Layer_Recipes) SetSV__Shutdown(v ...string) AWS_OpsWorks_Layer_Recipes {
	t.Shutdown = cfz.SV(v...)
	return t
}

// Set__Undeploy updates property "Undeploy".
func (t AWS_OpsWorks_Layer_Recipes) Set__Undeploy(v cfz.ExpressionSlice[string]) AWS_OpsWorks_Layer_Recipes {
	t.Undeploy = v
	return t
}

// SetS__Undeploy updates property "Undeploy".
func (t AWS_OpsWorks_Layer_Recipes) SetS__Undeploy(v ...cfz.Expression[string]) AWS_OpsWorks_Layer_Recipes {
	t.Undeploy = cfz.S(v...)
	return t
}

// SetSV__Undeploy updates property "Undeploy".
func (t AWS_OpsWorks_Layer_Recipes) SetSV__Undeploy(v ...string) AWS_OpsWorks_Layer_Recipes {
	t.Undeploy = cfz.SV(v...)
	return t
}
