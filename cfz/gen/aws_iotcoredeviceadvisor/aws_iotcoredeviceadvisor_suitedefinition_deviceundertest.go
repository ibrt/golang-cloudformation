// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotcoredeviceadvisor

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest__Type is the CloudFormation type for AWS::IoTCoreDeviceAdvisor::SuiteDefinition.DeviceUnderTest.
	AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest__Type = "AWS::IoTCoreDeviceAdvisor::SuiteDefinition.DeviceUnderTest"
)

var (
	// AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest__PropertiesMap reports all the CloudFormation properties for AWS::IoTCoreDeviceAdvisor::SuiteDefinition.DeviceUnderTest.
	AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest__PropertiesMap = struct {
		CertificateArn string
		ThingArn       string
	}{
		CertificateArn: "CertificateArn",
		ThingArn:       "ThingArn",
	}

	// AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest__PropertiesSlice reports all the CloudFormation properties for AWS::IoTCoreDeviceAdvisor::SuiteDefinition.DeviceUnderTest.
	AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest__PropertiesSlice = []string{
		AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest__PropertiesMap.CertificateArn,
		AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest__PropertiesMap.ThingArn,
	}
)

// AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest is a binding for AWS::IoTCoreDeviceAdvisor::SuiteDefinition.DeviceUnderTest.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-deviceundertest.html
type AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest struct {
	// CertificateArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-deviceundertest.html#cfn-iotcoredeviceadvisor-suitedefinition-deviceundertest-certificatearn
	CertificateArn cfz.Expression[string] `json:"CertificateArn,omitempty"`

	// ThingArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-deviceundertest.html#cfn-iotcoredeviceadvisor-suitedefinition-deviceundertest-thingarn
	ThingArn cfz.Expression[string] `json:"ThingArn,omitempty"`
}

// New__AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest initializes a new AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest.
func New__AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest() AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest {
	return AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest) GetType() string {
	return AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest__Type
}

// Set__CertificateArn updates property "CertificateArn".
func (t AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest) Set__CertificateArn(v cfz.Expression[string]) AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest {
	t.CertificateArn = v
	return t
}

// SetV__CertificateArn updates property "CertificateArn".
func (t AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest) SetV__CertificateArn(v string) AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest {
	t.CertificateArn = cfz.V(v)
	return t
}

// Set__ThingArn updates property "ThingArn".
func (t AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest) Set__ThingArn(v cfz.Expression[string]) AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest {
	t.ThingArn = v
	return t
}

// SetV__ThingArn updates property "ThingArn".
func (t AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest) SetV__ThingArn(v string) AWS_IoTCoreDeviceAdvisor_SuiteDefinition_DeviceUnderTest {
	t.ThingArn = cfz.V(v)
	return t
}
