// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_frauddetector

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FraudDetector_Detector_Label__Type is the CloudFormation type for AWS::FraudDetector::Detector.Label.
	AWS_FraudDetector_Detector_Label__Type = "AWS::FraudDetector::Detector.Label"
)

var (
	// AWS_FraudDetector_Detector_Label__PropertiesMap reports all the CloudFormation properties for AWS::FraudDetector::Detector.Label.
	AWS_FraudDetector_Detector_Label__PropertiesMap = struct {
		Arn             string
		CreatedTime     string
		Description     string
		Inline          string
		LastUpdatedTime string
		Name            string
		Tags            string
	}{
		Arn:             "Arn",
		CreatedTime:     "CreatedTime",
		Description:     "Description",
		Inline:          "Inline",
		LastUpdatedTime: "LastUpdatedTime",
		Name:            "Name",
		Tags:            "Tags",
	}

	// AWS_FraudDetector_Detector_Label__PropertiesSlice reports all the CloudFormation properties for AWS::FraudDetector::Detector.Label.
	AWS_FraudDetector_Detector_Label__PropertiesSlice = []string{
		AWS_FraudDetector_Detector_Label__PropertiesMap.Arn,
		AWS_FraudDetector_Detector_Label__PropertiesMap.CreatedTime,
		AWS_FraudDetector_Detector_Label__PropertiesMap.Description,
		AWS_FraudDetector_Detector_Label__PropertiesMap.Inline,
		AWS_FraudDetector_Detector_Label__PropertiesMap.LastUpdatedTime,
		AWS_FraudDetector_Detector_Label__PropertiesMap.Name,
		AWS_FraudDetector_Detector_Label__PropertiesMap.Tags,
	}
)

// AWS_FraudDetector_Detector_Label is a binding for AWS::FraudDetector::Detector.Label.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-label.html
type AWS_FraudDetector_Detector_Label struct {
	// Arn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-label.html#cfn-frauddetector-detector-label-arn
	Arn cfz.Expression[string] `json:"Arn,omitempty"`

	// CreatedTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-label.html#cfn-frauddetector-detector-label-createdtime
	CreatedTime cfz.Expression[string] `json:"CreatedTime,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-label.html#cfn-frauddetector-detector-label-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Inline is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-label.html#cfn-frauddetector-detector-label-inline
	Inline cfz.Expression[bool] `json:"Inline,omitempty"`

	// LastUpdatedTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-label.html#cfn-frauddetector-detector-label-lastupdatedtime
	LastUpdatedTime cfz.Expression[string] `json:"LastUpdatedTime,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-label.html#cfn-frauddetector-detector-label-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-label.html#cfn-frauddetector-detector-label-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_FraudDetector_Detector_Label initializes a new AWS_FraudDetector_Detector_Label.
func New__AWS_FraudDetector_Detector_Label() AWS_FraudDetector_Detector_Label {
	return AWS_FraudDetector_Detector_Label{}
}

// GetType returns the CloudFormation type.
func (AWS_FraudDetector_Detector_Label) GetType() string {
	return AWS_FraudDetector_Detector_Label__Type
}

// Set__Arn updates property "Arn".
func (t AWS_FraudDetector_Detector_Label) Set__Arn(v cfz.Expression[string]) AWS_FraudDetector_Detector_Label {
	t.Arn = v
	return t
}

// SetV__Arn updates property "Arn".
func (t AWS_FraudDetector_Detector_Label) SetV__Arn(v string) AWS_FraudDetector_Detector_Label {
	t.Arn = cfz.V(v)
	return t
}

// Set__CreatedTime updates property "CreatedTime".
func (t AWS_FraudDetector_Detector_Label) Set__CreatedTime(v cfz.Expression[string]) AWS_FraudDetector_Detector_Label {
	t.CreatedTime = v
	return t
}

// SetV__CreatedTime updates property "CreatedTime".
func (t AWS_FraudDetector_Detector_Label) SetV__CreatedTime(v string) AWS_FraudDetector_Detector_Label {
	t.CreatedTime = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_FraudDetector_Detector_Label) Set__Description(v cfz.Expression[string]) AWS_FraudDetector_Detector_Label {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_FraudDetector_Detector_Label) SetV__Description(v string) AWS_FraudDetector_Detector_Label {
	t.Description = cfz.V(v)
	return t
}

// Set__Inline updates property "Inline".
func (t AWS_FraudDetector_Detector_Label) Set__Inline(v cfz.Expression[bool]) AWS_FraudDetector_Detector_Label {
	t.Inline = v
	return t
}

// SetV__Inline updates property "Inline".
func (t AWS_FraudDetector_Detector_Label) SetV__Inline(v bool) AWS_FraudDetector_Detector_Label {
	t.Inline = cfz.V(v)
	return t
}

// Set__LastUpdatedTime updates property "LastUpdatedTime".
func (t AWS_FraudDetector_Detector_Label) Set__LastUpdatedTime(v cfz.Expression[string]) AWS_FraudDetector_Detector_Label {
	t.LastUpdatedTime = v
	return t
}

// SetV__LastUpdatedTime updates property "LastUpdatedTime".
func (t AWS_FraudDetector_Detector_Label) SetV__LastUpdatedTime(v string) AWS_FraudDetector_Detector_Label {
	t.LastUpdatedTime = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_FraudDetector_Detector_Label) Set__Name(v cfz.Expression[string]) AWS_FraudDetector_Detector_Label {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_FraudDetector_Detector_Label) SetV__Name(v string) AWS_FraudDetector_Detector_Label {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t AWS_FraudDetector_Detector_Label) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) AWS_FraudDetector_Detector_Label {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t AWS_FraudDetector_Detector_Label) SetS__Tags(v ...cfz.Expression[cfz.Tag]) AWS_FraudDetector_Detector_Label {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t AWS_FraudDetector_Detector_Label) SetSV__Tags(v ...cfz.Tag) AWS_FraudDetector_Detector_Label {
	t.Tags = cfz.SV(v...)
	return t
}
