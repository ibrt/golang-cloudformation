// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent__Type is the CloudFormation type for AWS::EC2::NetworkInsightsAnalysis.AnalysisComponent.
	AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent__Type = "AWS::EC2::NetworkInsightsAnalysis.AnalysisComponent"
)

var (
	// AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent__PropertiesMap reports all the CloudFormation properties for AWS::EC2::NetworkInsightsAnalysis.AnalysisComponent.
	AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent__PropertiesMap = struct {
		Arn string
		Id  string
	}{
		Arn: "Arn",
		Id:  "Id",
	}

	// AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::NetworkInsightsAnalysis.AnalysisComponent.
	AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent__PropertiesSlice = []string{
		AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent__PropertiesMap.Arn,
		AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent__PropertiesMap.Id,
	}
)

// AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent is a binding for AWS::EC2::NetworkInsightsAnalysis.AnalysisComponent.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysiscomponent.html
type AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent struct {
	// Arn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysiscomponent.html#cfn-ec2-networkinsightsanalysis-analysiscomponent-arn
	Arn cfz.Expression[string] `json:"Arn,omitempty"`

	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysiscomponent.html#cfn-ec2-networkinsightsanalysis-analysiscomponent-id
	Id cfz.Expression[string] `json:"Id,omitempty"`
}

// New__AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent initializes a new AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent.
func New__AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent() AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent {
	return AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent) GetType() string {
	return AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent__Type
}

// Set__Arn updates property "Arn".
func (t AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent) Set__Arn(v cfz.Expression[string]) AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent {
	t.Arn = v
	return t
}

// SetV__Arn updates property "Arn".
func (t AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent) SetV__Arn(v string) AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent {
	t.Arn = cfz.V(v)
	return t
}

// Set__Id updates property "Id".
func (t AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent) Set__Id(v cfz.Expression[string]) AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent) SetV__Id(v string) AWS_EC2_NetworkInsightsAnalysis_AnalysisComponent {
	t.Id = cfz.V(v)
	return t
}
