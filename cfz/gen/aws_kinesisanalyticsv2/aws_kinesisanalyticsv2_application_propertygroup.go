// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisanalyticsv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisAnalyticsV2_Application_PropertyGroup__Type is the CloudFormation type for AWS::KinesisAnalyticsV2::Application.PropertyGroup.
	AWS_KinesisAnalyticsV2_Application_PropertyGroup__Type = "AWS::KinesisAnalyticsV2::Application.PropertyGroup"
)

var (
	// AWS_KinesisAnalyticsV2_Application_PropertyGroup__PropertiesMap reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.PropertyGroup.
	AWS_KinesisAnalyticsV2_Application_PropertyGroup__PropertiesMap = struct {
		PropertyGroupId string
		PropertyMap     string
	}{
		PropertyGroupId: "PropertyGroupId",
		PropertyMap:     "PropertyMap",
	}

	// AWS_KinesisAnalyticsV2_Application_PropertyGroup__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.PropertyGroup.
	AWS_KinesisAnalyticsV2_Application_PropertyGroup__PropertiesSlice = []string{
		AWS_KinesisAnalyticsV2_Application_PropertyGroup__PropertiesMap.PropertyGroupId,
		AWS_KinesisAnalyticsV2_Application_PropertyGroup__PropertiesMap.PropertyMap,
	}
)

// AWS_KinesisAnalyticsV2_Application_PropertyGroup is a binding for AWS::KinesisAnalyticsV2::Application.PropertyGroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html
type AWS_KinesisAnalyticsV2_Application_PropertyGroup struct {
	// PropertyGroupId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html#cfn-kinesisanalyticsv2-application-propertygroup-propertygroupid
	PropertyGroupId cfz.Expression[string] `json:"PropertyGroupId,omitempty"`

	// PropertyMap is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html#cfn-kinesisanalyticsv2-application-propertygroup-propertymap
	PropertyMap cfz.ExpressionMap[string] `json:"PropertyMap,omitempty"`
}

// New__AWS_KinesisAnalyticsV2_Application_PropertyGroup initializes a new AWS_KinesisAnalyticsV2_Application_PropertyGroup.
func New__AWS_KinesisAnalyticsV2_Application_PropertyGroup() AWS_KinesisAnalyticsV2_Application_PropertyGroup {
	return AWS_KinesisAnalyticsV2_Application_PropertyGroup{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisAnalyticsV2_Application_PropertyGroup) GetType() string {
	return AWS_KinesisAnalyticsV2_Application_PropertyGroup__Type
}

// Set__PropertyGroupId updates property "PropertyGroupId".
func (t AWS_KinesisAnalyticsV2_Application_PropertyGroup) Set__PropertyGroupId(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_Application_PropertyGroup {
	t.PropertyGroupId = v
	return t
}

// SetV__PropertyGroupId updates property "PropertyGroupId".
func (t AWS_KinesisAnalyticsV2_Application_PropertyGroup) SetV__PropertyGroupId(v string) AWS_KinesisAnalyticsV2_Application_PropertyGroup {
	t.PropertyGroupId = cfz.V(v)
	return t
}

// Set__PropertyMap updates property "PropertyMap".
func (t AWS_KinesisAnalyticsV2_Application_PropertyGroup) Set__PropertyMap(v cfz.ExpressionMap[string]) AWS_KinesisAnalyticsV2_Application_PropertyGroup {
	t.PropertyMap = v
	return t
}

// SetM__PropertyMap updates property "PropertyMap".
func (t AWS_KinesisAnalyticsV2_Application_PropertyGroup) SetM__PropertyMap(v ...map[string]cfz.Expression[string]) AWS_KinesisAnalyticsV2_Application_PropertyGroup {
	t.PropertyMap = cfz.M(v...)
	return t
}

// SetMV__PropertyMap updates property "PropertyMap".
func (t AWS_KinesisAnalyticsV2_Application_PropertyGroup) SetMV__PropertyMap(v ...map[string]string) AWS_KinesisAnalyticsV2_Application_PropertyGroup {
	t.PropertyMap = cfz.MV(v...)
	return t
}
