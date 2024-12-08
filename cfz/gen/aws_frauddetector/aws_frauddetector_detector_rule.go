// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_frauddetector

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FraudDetector_Detector_Rule__Type is the CloudFormation type for AWS::FraudDetector::Detector.Rule.
	AWS_FraudDetector_Detector_Rule__Type = "AWS::FraudDetector::Detector.Rule"
)

var (
	// AWS_FraudDetector_Detector_Rule__PropertiesMap reports all the CloudFormation properties for AWS::FraudDetector::Detector.Rule.
	AWS_FraudDetector_Detector_Rule__PropertiesMap = struct {
		Arn             string
		CreatedTime     string
		Description     string
		DetectorId      string
		Expression      string
		Language        string
		LastUpdatedTime string
		Outcomes        string
		RuleId          string
		RuleVersion     string
		Tags            string
	}{
		Arn:             "Arn",
		CreatedTime:     "CreatedTime",
		Description:     "Description",
		DetectorId:      "DetectorId",
		Expression:      "Expression",
		Language:        "Language",
		LastUpdatedTime: "LastUpdatedTime",
		Outcomes:        "Outcomes",
		RuleId:          "RuleId",
		RuleVersion:     "RuleVersion",
		Tags:            "Tags",
	}

	// AWS_FraudDetector_Detector_Rule__PropertiesSlice reports all the CloudFormation properties for AWS::FraudDetector::Detector.Rule.
	AWS_FraudDetector_Detector_Rule__PropertiesSlice = []string{
		AWS_FraudDetector_Detector_Rule__PropertiesMap.Arn,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.CreatedTime,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.Description,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.DetectorId,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.Expression,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.Language,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.LastUpdatedTime,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.Outcomes,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.RuleId,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.RuleVersion,
		AWS_FraudDetector_Detector_Rule__PropertiesMap.Tags,
	}
)

// AWS_FraudDetector_Detector_Rule is a binding for AWS::FraudDetector::Detector.Rule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html
type AWS_FraudDetector_Detector_Rule struct {
	// Arn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-arn
	Arn cfz.Expression[string] `json:"Arn,omitempty"`

	// CreatedTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-createdtime
	CreatedTime cfz.Expression[string] `json:"CreatedTime,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// DetectorId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-detectorid
	DetectorId cfz.Expression[string] `json:"DetectorId,omitempty"`

	// Expression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-expression
	Expression cfz.Expression[string] `json:"Expression,omitempty"`

	// Language is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-language
	Language cfz.Expression[string] `json:"Language,omitempty"`

	// LastUpdatedTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-lastupdatedtime
	LastUpdatedTime cfz.Expression[string] `json:"LastUpdatedTime,omitempty"`

	// Outcomes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-outcomes
	Outcomes cfz.ExpressionSlice[AWS_FraudDetector_Detector_Outcome] `json:"Outcomes,omitempty"`

	// RuleId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-ruleid
	RuleId cfz.Expression[string] `json:"RuleId,omitempty"`

	// RuleVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-ruleversion
	RuleVersion cfz.Expression[string] `json:"RuleVersion,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-rule.html#cfn-frauddetector-detector-rule-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_FraudDetector_Detector_Rule initializes a new AWS_FraudDetector_Detector_Rule.
func New__AWS_FraudDetector_Detector_Rule() AWS_FraudDetector_Detector_Rule {
	return AWS_FraudDetector_Detector_Rule{}
}

// GetType returns the CloudFormation type.
func (AWS_FraudDetector_Detector_Rule) GetType() string {
	return AWS_FraudDetector_Detector_Rule__Type
}

// Set__Arn updates property "Arn".
func (t AWS_FraudDetector_Detector_Rule) Set__Arn(v cfz.Expression[string]) AWS_FraudDetector_Detector_Rule {
	t.Arn = v
	return t
}

// SetV__Arn updates property "Arn".
func (t AWS_FraudDetector_Detector_Rule) SetV__Arn(v string) AWS_FraudDetector_Detector_Rule {
	t.Arn = cfz.V(v)
	return t
}

// Set__CreatedTime updates property "CreatedTime".
func (t AWS_FraudDetector_Detector_Rule) Set__CreatedTime(v cfz.Expression[string]) AWS_FraudDetector_Detector_Rule {
	t.CreatedTime = v
	return t
}

// SetV__CreatedTime updates property "CreatedTime".
func (t AWS_FraudDetector_Detector_Rule) SetV__CreatedTime(v string) AWS_FraudDetector_Detector_Rule {
	t.CreatedTime = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_FraudDetector_Detector_Rule) Set__Description(v cfz.Expression[string]) AWS_FraudDetector_Detector_Rule {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_FraudDetector_Detector_Rule) SetV__Description(v string) AWS_FraudDetector_Detector_Rule {
	t.Description = cfz.V(v)
	return t
}

// Set__DetectorId updates property "DetectorId".
func (t AWS_FraudDetector_Detector_Rule) Set__DetectorId(v cfz.Expression[string]) AWS_FraudDetector_Detector_Rule {
	t.DetectorId = v
	return t
}

// SetV__DetectorId updates property "DetectorId".
func (t AWS_FraudDetector_Detector_Rule) SetV__DetectorId(v string) AWS_FraudDetector_Detector_Rule {
	t.DetectorId = cfz.V(v)
	return t
}

// Set__Expression updates property "Expression".
func (t AWS_FraudDetector_Detector_Rule) Set__Expression(v cfz.Expression[string]) AWS_FraudDetector_Detector_Rule {
	t.Expression = v
	return t
}

// SetV__Expression updates property "Expression".
func (t AWS_FraudDetector_Detector_Rule) SetV__Expression(v string) AWS_FraudDetector_Detector_Rule {
	t.Expression = cfz.V(v)
	return t
}

// Set__Language updates property "Language".
func (t AWS_FraudDetector_Detector_Rule) Set__Language(v cfz.Expression[string]) AWS_FraudDetector_Detector_Rule {
	t.Language = v
	return t
}

// SetV__Language updates property "Language".
func (t AWS_FraudDetector_Detector_Rule) SetV__Language(v string) AWS_FraudDetector_Detector_Rule {
	t.Language = cfz.V(v)
	return t
}

// Set__LastUpdatedTime updates property "LastUpdatedTime".
func (t AWS_FraudDetector_Detector_Rule) Set__LastUpdatedTime(v cfz.Expression[string]) AWS_FraudDetector_Detector_Rule {
	t.LastUpdatedTime = v
	return t
}

// SetV__LastUpdatedTime updates property "LastUpdatedTime".
func (t AWS_FraudDetector_Detector_Rule) SetV__LastUpdatedTime(v string) AWS_FraudDetector_Detector_Rule {
	t.LastUpdatedTime = cfz.V(v)
	return t
}

// Set__Outcomes updates property "Outcomes".
func (t AWS_FraudDetector_Detector_Rule) Set__Outcomes(v cfz.ExpressionSlice[AWS_FraudDetector_Detector_Outcome]) AWS_FraudDetector_Detector_Rule {
	t.Outcomes = v
	return t
}

// SetS__Outcomes updates property "Outcomes".
func (t AWS_FraudDetector_Detector_Rule) SetS__Outcomes(v ...cfz.Expression[AWS_FraudDetector_Detector_Outcome]) AWS_FraudDetector_Detector_Rule {
	t.Outcomes = cfz.S(v...)
	return t
}

// SetSV__Outcomes updates property "Outcomes".
func (t AWS_FraudDetector_Detector_Rule) SetSV__Outcomes(v ...AWS_FraudDetector_Detector_Outcome) AWS_FraudDetector_Detector_Rule {
	t.Outcomes = cfz.SV(v...)
	return t
}

// Set__RuleId updates property "RuleId".
func (t AWS_FraudDetector_Detector_Rule) Set__RuleId(v cfz.Expression[string]) AWS_FraudDetector_Detector_Rule {
	t.RuleId = v
	return t
}

// SetV__RuleId updates property "RuleId".
func (t AWS_FraudDetector_Detector_Rule) SetV__RuleId(v string) AWS_FraudDetector_Detector_Rule {
	t.RuleId = cfz.V(v)
	return t
}

// Set__RuleVersion updates property "RuleVersion".
func (t AWS_FraudDetector_Detector_Rule) Set__RuleVersion(v cfz.Expression[string]) AWS_FraudDetector_Detector_Rule {
	t.RuleVersion = v
	return t
}

// SetV__RuleVersion updates property "RuleVersion".
func (t AWS_FraudDetector_Detector_Rule) SetV__RuleVersion(v string) AWS_FraudDetector_Detector_Rule {
	t.RuleVersion = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t AWS_FraudDetector_Detector_Rule) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) AWS_FraudDetector_Detector_Rule {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t AWS_FraudDetector_Detector_Rule) SetS__Tags(v ...cfz.Expression[cfz.Tag]) AWS_FraudDetector_Detector_Rule {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t AWS_FraudDetector_Detector_Rule) SetSV__Tags(v ...cfz.Tag) AWS_FraudDetector_Detector_Rule {
	t.Tags = cfz.SV(v...)
	return t
}
