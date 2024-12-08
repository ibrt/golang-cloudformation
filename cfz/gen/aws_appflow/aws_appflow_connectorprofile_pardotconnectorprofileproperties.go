// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appflow

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__Type is the CloudFormation type for AWS::AppFlow::ConnectorProfile.PardotConnectorProfileProperties.
	AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__Type = "AWS::AppFlow::ConnectorProfile.PardotConnectorProfileProperties"
)

var (
	// AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__PropertiesMap reports all the CloudFormation properties for AWS::AppFlow::ConnectorProfile.PardotConnectorProfileProperties.
	AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__PropertiesMap = struct {
		BusinessUnitId       string
		InstanceUrl          string
		IsSandboxEnvironment string
	}{
		BusinessUnitId:       "BusinessUnitId",
		InstanceUrl:          "InstanceUrl",
		IsSandboxEnvironment: "IsSandboxEnvironment",
	}

	// AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__PropertiesSlice reports all the CloudFormation properties for AWS::AppFlow::ConnectorProfile.PardotConnectorProfileProperties.
	AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__PropertiesSlice = []string{
		AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__PropertiesMap.BusinessUnitId,
		AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__PropertiesMap.InstanceUrl,
		AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__PropertiesMap.IsSandboxEnvironment,
	}
)

// AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties is a binding for AWS::AppFlow::ConnectorProfile.PardotConnectorProfileProperties.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html
type AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties struct {
	// BusinessUnitId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html#cfn-appflow-connectorprofile-pardotconnectorprofileproperties-businessunitid
	BusinessUnitId cfz.Expression[string] `json:"BusinessUnitId,omitempty"`

	// InstanceUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html#cfn-appflow-connectorprofile-pardotconnectorprofileproperties-instanceurl
	InstanceUrl cfz.Expression[string] `json:"InstanceUrl,omitempty"`

	// IsSandboxEnvironment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html#cfn-appflow-connectorprofile-pardotconnectorprofileproperties-issandboxenvironment
	IsSandboxEnvironment cfz.Expression[bool] `json:"IsSandboxEnvironment,omitempty"`
}

// New__AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties initializes a new AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties.
func New__AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties() AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties {
	return AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties{}
}

// GetType returns the CloudFormation type.
func (AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties) GetType() string {
	return AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties__Type
}

// Set__BusinessUnitId updates property "BusinessUnitId".
func (t AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties) Set__BusinessUnitId(v cfz.Expression[string]) AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties {
	t.BusinessUnitId = v
	return t
}

// SetV__BusinessUnitId updates property "BusinessUnitId".
func (t AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties) SetV__BusinessUnitId(v string) AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties {
	t.BusinessUnitId = cfz.V(v)
	return t
}

// Set__InstanceUrl updates property "InstanceUrl".
func (t AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties) Set__InstanceUrl(v cfz.Expression[string]) AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties {
	t.InstanceUrl = v
	return t
}

// SetV__InstanceUrl updates property "InstanceUrl".
func (t AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties) SetV__InstanceUrl(v string) AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties {
	t.InstanceUrl = cfz.V(v)
	return t
}

// Set__IsSandboxEnvironment updates property "IsSandboxEnvironment".
func (t AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties) Set__IsSandboxEnvironment(v cfz.Expression[bool]) AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties {
	t.IsSandboxEnvironment = v
	return t
}

// SetV__IsSandboxEnvironment updates property "IsSandboxEnvironment".
func (t AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties) SetV__IsSandboxEnvironment(v bool) AWS_AppFlow_ConnectorProfile_PardotConnectorProfileProperties {
	t.IsSandboxEnvironment = cfz.V(v)
	return t
}
