// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelPackage_Explainability__Type is the CloudFormation type for AWS::SageMaker::ModelPackage.Explainability.
	AWS_SageMaker_ModelPackage_Explainability__Type = "AWS::SageMaker::ModelPackage.Explainability"
)

var (
	// AWS_SageMaker_ModelPackage_Explainability__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelPackage.Explainability.
	AWS_SageMaker_ModelPackage_Explainability__PropertiesMap = struct {
		Report string
	}{
		Report: "Report",
	}

	// AWS_SageMaker_ModelPackage_Explainability__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelPackage.Explainability.
	AWS_SageMaker_ModelPackage_Explainability__PropertiesSlice = []string{
		AWS_SageMaker_ModelPackage_Explainability__PropertiesMap.Report,
	}
)

// AWS_SageMaker_ModelPackage_Explainability is a binding for AWS::SageMaker::ModelPackage.Explainability.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-explainability.html
type AWS_SageMaker_ModelPackage_Explainability struct {
	// Report is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-explainability.html#cfn-sagemaker-modelpackage-explainability-report
	Report cfz.Expression[AWS_SageMaker_ModelPackage_MetricsSource] `json:"Report,omitempty"`
}

// New__AWS_SageMaker_ModelPackage_Explainability initializes a new AWS_SageMaker_ModelPackage_Explainability.
func New__AWS_SageMaker_ModelPackage_Explainability() AWS_SageMaker_ModelPackage_Explainability {
	return AWS_SageMaker_ModelPackage_Explainability{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelPackage_Explainability) GetType() string {
	return AWS_SageMaker_ModelPackage_Explainability__Type
}

// Set__Report updates property "Report".
func (t AWS_SageMaker_ModelPackage_Explainability) Set__Report(v cfz.Expression[AWS_SageMaker_ModelPackage_MetricsSource]) AWS_SageMaker_ModelPackage_Explainability {
	t.Report = v
	return t
}

// SetV__Report updates property "Report".
func (t AWS_SageMaker_ModelPackage_Explainability) SetV__Report(v AWS_SageMaker_ModelPackage_MetricsSource) AWS_SageMaker_ModelPackage_Explainability {
	t.Report = cfz.V(v)
	return t
}
