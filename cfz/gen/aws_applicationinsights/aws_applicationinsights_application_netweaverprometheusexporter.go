// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_applicationinsights

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__Type is the CloudFormation type for AWS::ApplicationInsights::Application.NetWeaverPrometheusExporter.
	AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__Type = "AWS::ApplicationInsights::Application.NetWeaverPrometheusExporter"
)

var (
	// AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__PropertiesMap reports all the CloudFormation properties for AWS::ApplicationInsights::Application.NetWeaverPrometheusExporter.
	AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__PropertiesMap = struct {
		InstanceNumbers string
		PrometheusPort  string
		SAPSID          string
	}{
		InstanceNumbers: "InstanceNumbers",
		PrometheusPort:  "PrometheusPort",
		SAPSID:          "SAPSID",
	}

	// AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__PropertiesSlice reports all the CloudFormation properties for AWS::ApplicationInsights::Application.NetWeaverPrometheusExporter.
	AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__PropertiesSlice = []string{
		AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__PropertiesMap.InstanceNumbers,
		AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__PropertiesMap.PrometheusPort,
		AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__PropertiesMap.SAPSID,
	}
)

// AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter is a binding for AWS::ApplicationInsights::Application.NetWeaverPrometheusExporter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html
type AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter struct {
	// InstanceNumbers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html#cfn-applicationinsights-application-netweaverprometheusexporter-instancenumbers
	InstanceNumbers cfz.ExpressionSlice[string] `json:"InstanceNumbers,omitempty"`

	// PrometheusPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html#cfn-applicationinsights-application-netweaverprometheusexporter-prometheusport
	PrometheusPort cfz.Expression[string] `json:"PrometheusPort,omitempty"`

	// SAPSID is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-netweaverprometheusexporter.html#cfn-applicationinsights-application-netweaverprometheusexporter-sapsid
	SAPSID cfz.Expression[string] `json:"SAPSID,omitempty"`
}

// New__AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter initializes a new AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter.
func New__AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter() AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter {
	return AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter{}
}

// GetType returns the CloudFormation type.
func (AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter) GetType() string {
	return AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter__Type
}

// Set__InstanceNumbers updates property "InstanceNumbers".
func (t AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter) Set__InstanceNumbers(v cfz.ExpressionSlice[string]) AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter {
	t.InstanceNumbers = v
	return t
}

// SetS__InstanceNumbers updates property "InstanceNumbers".
func (t AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter) SetS__InstanceNumbers(v ...cfz.Expression[string]) AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter {
	t.InstanceNumbers = cfz.S(v...)
	return t
}

// SetSV__InstanceNumbers updates property "InstanceNumbers".
func (t AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter) SetSV__InstanceNumbers(v ...string) AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter {
	t.InstanceNumbers = cfz.SV(v...)
	return t
}

// Set__PrometheusPort updates property "PrometheusPort".
func (t AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter) Set__PrometheusPort(v cfz.Expression[string]) AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter {
	t.PrometheusPort = v
	return t
}

// SetV__PrometheusPort updates property "PrometheusPort".
func (t AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter) SetV__PrometheusPort(v string) AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter {
	t.PrometheusPort = cfz.V(v)
	return t
}

// Set__SAPSID updates property "SAPSID".
func (t AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter) Set__SAPSID(v cfz.Expression[string]) AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter {
	t.SAPSID = v
	return t
}

// SetV__SAPSID updates property "SAPSID".
func (t AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter) SetV__SAPSID(v string) AWS_ApplicationInsights_Application_NetWeaverPrometheusExporter {
	t.SAPSID = cfz.V(v)
	return t
}
