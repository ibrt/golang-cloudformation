// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_LaunchTemplate_LaunchTemplateData__Type is the CloudFormation type for AWS::EC2::LaunchTemplate.LaunchTemplateData.
	AWS_EC2_LaunchTemplate_LaunchTemplateData__Type = "AWS::EC2::LaunchTemplate.LaunchTemplateData"
)

var (
	// AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap reports all the CloudFormation properties for AWS::EC2::LaunchTemplate.LaunchTemplateData.
	AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap = struct {
		BlockDeviceMappings               string
		CapacityReservationSpecification  string
		CpuOptions                        string
		CreditSpecification               string
		DisableApiStop                    string
		DisableApiTermination             string
		EbsOptimized                      string
		ElasticGpuSpecifications          string
		ElasticInferenceAccelerators      string
		EnclaveOptions                    string
		HibernationOptions                string
		IamInstanceProfile                string
		ImageId                           string
		InstanceInitiatedShutdownBehavior string
		InstanceMarketOptions             string
		InstanceRequirements              string
		InstanceType                      string
		KernelId                          string
		KeyName                           string
		LicenseSpecifications             string
		MaintenanceOptions                string
		MetadataOptions                   string
		Monitoring                        string
		NetworkInterfaces                 string
		Placement                         string
		PrivateDnsNameOptions             string
		RamDiskId                         string
		SecurityGroupIds                  string
		SecurityGroups                    string
		TagSpecifications                 string
		UserData                          string
	}{
		BlockDeviceMappings:               "BlockDeviceMappings",
		CapacityReservationSpecification:  "CapacityReservationSpecification",
		CpuOptions:                        "CpuOptions",
		CreditSpecification:               "CreditSpecification",
		DisableApiStop:                    "DisableApiStop",
		DisableApiTermination:             "DisableApiTermination",
		EbsOptimized:                      "EbsOptimized",
		ElasticGpuSpecifications:          "ElasticGpuSpecifications",
		ElasticInferenceAccelerators:      "ElasticInferenceAccelerators",
		EnclaveOptions:                    "EnclaveOptions",
		HibernationOptions:                "HibernationOptions",
		IamInstanceProfile:                "IamInstanceProfile",
		ImageId:                           "ImageId",
		InstanceInitiatedShutdownBehavior: "InstanceInitiatedShutdownBehavior",
		InstanceMarketOptions:             "InstanceMarketOptions",
		InstanceRequirements:              "InstanceRequirements",
		InstanceType:                      "InstanceType",
		KernelId:                          "KernelId",
		KeyName:                           "KeyName",
		LicenseSpecifications:             "LicenseSpecifications",
		MaintenanceOptions:                "MaintenanceOptions",
		MetadataOptions:                   "MetadataOptions",
		Monitoring:                        "Monitoring",
		NetworkInterfaces:                 "NetworkInterfaces",
		Placement:                         "Placement",
		PrivateDnsNameOptions:             "PrivateDnsNameOptions",
		RamDiskId:                         "RamDiskId",
		SecurityGroupIds:                  "SecurityGroupIds",
		SecurityGroups:                    "SecurityGroups",
		TagSpecifications:                 "TagSpecifications",
		UserData:                          "UserData",
	}

	// AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::LaunchTemplate.LaunchTemplateData.
	AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesSlice = []string{
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.BlockDeviceMappings,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.CapacityReservationSpecification,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.CpuOptions,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.CreditSpecification,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.DisableApiStop,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.DisableApiTermination,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.EbsOptimized,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.ElasticGpuSpecifications,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.ElasticInferenceAccelerators,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.EnclaveOptions,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.HibernationOptions,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.IamInstanceProfile,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.ImageId,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.InstanceInitiatedShutdownBehavior,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.InstanceMarketOptions,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.InstanceRequirements,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.InstanceType,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.KernelId,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.KeyName,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.LicenseSpecifications,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.MaintenanceOptions,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.MetadataOptions,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.Monitoring,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.NetworkInterfaces,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.Placement,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.PrivateDnsNameOptions,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.RamDiskId,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.SecurityGroupIds,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.SecurityGroups,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.TagSpecifications,
		AWS_EC2_LaunchTemplate_LaunchTemplateData__PropertiesMap.UserData,
	}
)

// AWS_EC2_LaunchTemplate_LaunchTemplateData is a binding for AWS::EC2::LaunchTemplate.LaunchTemplateData.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html
type AWS_EC2_LaunchTemplate_LaunchTemplateData struct {
	// BlockDeviceMappings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-blockdevicemappings
	BlockDeviceMappings cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_BlockDeviceMapping] `json:"BlockDeviceMappings,omitempty"`

	// CapacityReservationSpecification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-capacityreservationspecification
	CapacityReservationSpecification cfz.Expression[AWS_EC2_LaunchTemplate_CapacityReservationSpecification] `json:"CapacityReservationSpecification,omitempty"`

	// CpuOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-cpuoptions
	CpuOptions cfz.Expression[AWS_EC2_LaunchTemplate_CpuOptions] `json:"CpuOptions,omitempty"`

	// CreditSpecification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-creditspecification
	CreditSpecification cfz.Expression[AWS_EC2_LaunchTemplate_CreditSpecification] `json:"CreditSpecification,omitempty"`

	// DisableApiStop is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-disableapistop
	DisableApiStop cfz.Expression[bool] `json:"DisableApiStop,omitempty"`

	// DisableApiTermination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-disableapitermination
	DisableApiTermination cfz.Expression[bool] `json:"DisableApiTermination,omitempty"`

	// EbsOptimized is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-ebsoptimized
	EbsOptimized cfz.Expression[bool] `json:"EbsOptimized,omitempty"`

	// ElasticGpuSpecifications is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-elasticgpuspecifications
	ElasticGpuSpecifications cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_ElasticGpuSpecification] `json:"ElasticGpuSpecifications,omitempty"`

	// ElasticInferenceAccelerators is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-elasticinferenceaccelerators
	ElasticInferenceAccelerators cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_LaunchTemplateElasticInferenceAccelerator] `json:"ElasticInferenceAccelerators,omitempty"`

	// EnclaveOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-enclaveoptions
	EnclaveOptions cfz.Expression[AWS_EC2_LaunchTemplate_EnclaveOptions] `json:"EnclaveOptions,omitempty"`

	// HibernationOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-hibernationoptions
	HibernationOptions cfz.Expression[AWS_EC2_LaunchTemplate_HibernationOptions] `json:"HibernationOptions,omitempty"`

	// IamInstanceProfile is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-iaminstanceprofile
	IamInstanceProfile cfz.Expression[AWS_EC2_LaunchTemplate_IamInstanceProfile] `json:"IamInstanceProfile,omitempty"`

	// ImageId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-imageid
	ImageId cfz.Expression[string] `json:"ImageId,omitempty"`

	// InstanceInitiatedShutdownBehavior is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-instanceinitiatedshutdownbehavior
	InstanceInitiatedShutdownBehavior cfz.Expression[string] `json:"InstanceInitiatedShutdownBehavior,omitempty"`

	// InstanceMarketOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-instancemarketoptions
	InstanceMarketOptions cfz.Expression[AWS_EC2_LaunchTemplate_InstanceMarketOptions] `json:"InstanceMarketOptions,omitempty"`

	// InstanceRequirements is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-instancerequirements
	InstanceRequirements cfz.Expression[AWS_EC2_LaunchTemplate_InstanceRequirements] `json:"InstanceRequirements,omitempty"`

	// InstanceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-instancetype
	InstanceType cfz.Expression[string] `json:"InstanceType,omitempty"`

	// KernelId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-kernelid
	KernelId cfz.Expression[string] `json:"KernelId,omitempty"`

	// KeyName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-keyname
	KeyName cfz.Expression[string] `json:"KeyName,omitempty"`

	// LicenseSpecifications is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-licensespecifications
	LicenseSpecifications cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_LicenseSpecification] `json:"LicenseSpecifications,omitempty"`

	// MaintenanceOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-maintenanceoptions
	MaintenanceOptions cfz.Expression[AWS_EC2_LaunchTemplate_MaintenanceOptions] `json:"MaintenanceOptions,omitempty"`

	// MetadataOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions
	MetadataOptions cfz.Expression[AWS_EC2_LaunchTemplate_MetadataOptions] `json:"MetadataOptions,omitempty"`

	// Monitoring is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-monitoring
	Monitoring cfz.Expression[AWS_EC2_LaunchTemplate_Monitoring] `json:"Monitoring,omitempty"`

	// NetworkInterfaces is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-networkinterfaces
	NetworkInterfaces cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_NetworkInterface] `json:"NetworkInterfaces,omitempty"`

	// Placement is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-placement
	Placement cfz.Expression[AWS_EC2_LaunchTemplate_Placement] `json:"Placement,omitempty"`

	// PrivateDnsNameOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-privatednsnameoptions
	PrivateDnsNameOptions cfz.Expression[AWS_EC2_LaunchTemplate_PrivateDnsNameOptions] `json:"PrivateDnsNameOptions,omitempty"`

	// RamDiskId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-ramdiskid
	RamDiskId cfz.Expression[string] `json:"RamDiskId,omitempty"`

	// SecurityGroupIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-securitygroupids
	SecurityGroupIds cfz.ExpressionSlice[string] `json:"SecurityGroupIds,omitempty"`

	// SecurityGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-securitygroups
	SecurityGroups cfz.ExpressionSlice[string] `json:"SecurityGroups,omitempty"`

	// TagSpecifications is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-tagspecifications
	TagSpecifications cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_TagSpecification] `json:"TagSpecifications,omitempty"`

	// UserData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-userdata
	UserData cfz.Expression[string] `json:"UserData,omitempty"`
}

// New__AWS_EC2_LaunchTemplate_LaunchTemplateData initializes a new AWS_EC2_LaunchTemplate_LaunchTemplateData.
func New__AWS_EC2_LaunchTemplate_LaunchTemplateData() AWS_EC2_LaunchTemplate_LaunchTemplateData {
	return AWS_EC2_LaunchTemplate_LaunchTemplateData{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_LaunchTemplate_LaunchTemplateData) GetType() string {
	return AWS_EC2_LaunchTemplate_LaunchTemplateData__Type
}

// Set__BlockDeviceMappings updates property "BlockDeviceMappings".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__BlockDeviceMappings(v cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_BlockDeviceMapping]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.BlockDeviceMappings = v
	return t
}

// SetS__BlockDeviceMappings updates property "BlockDeviceMappings".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetS__BlockDeviceMappings(v ...cfz.Expression[AWS_EC2_LaunchTemplate_BlockDeviceMapping]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.BlockDeviceMappings = cfz.S(v...)
	return t
}

// SetSV__BlockDeviceMappings updates property "BlockDeviceMappings".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetSV__BlockDeviceMappings(v ...AWS_EC2_LaunchTemplate_BlockDeviceMapping) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.BlockDeviceMappings = cfz.SV(v...)
	return t
}

// Set__CapacityReservationSpecification updates property "CapacityReservationSpecification".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__CapacityReservationSpecification(v cfz.Expression[AWS_EC2_LaunchTemplate_CapacityReservationSpecification]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.CapacityReservationSpecification = v
	return t
}

// SetV__CapacityReservationSpecification updates property "CapacityReservationSpecification".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__CapacityReservationSpecification(v AWS_EC2_LaunchTemplate_CapacityReservationSpecification) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.CapacityReservationSpecification = cfz.V(v)
	return t
}

// Set__CpuOptions updates property "CpuOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__CpuOptions(v cfz.Expression[AWS_EC2_LaunchTemplate_CpuOptions]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.CpuOptions = v
	return t
}

// SetV__CpuOptions updates property "CpuOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__CpuOptions(v AWS_EC2_LaunchTemplate_CpuOptions) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.CpuOptions = cfz.V(v)
	return t
}

// Set__CreditSpecification updates property "CreditSpecification".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__CreditSpecification(v cfz.Expression[AWS_EC2_LaunchTemplate_CreditSpecification]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.CreditSpecification = v
	return t
}

// SetV__CreditSpecification updates property "CreditSpecification".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__CreditSpecification(v AWS_EC2_LaunchTemplate_CreditSpecification) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.CreditSpecification = cfz.V(v)
	return t
}

// Set__DisableApiStop updates property "DisableApiStop".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__DisableApiStop(v cfz.Expression[bool]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.DisableApiStop = v
	return t
}

// SetV__DisableApiStop updates property "DisableApiStop".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__DisableApiStop(v bool) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.DisableApiStop = cfz.V(v)
	return t
}

// Set__DisableApiTermination updates property "DisableApiTermination".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__DisableApiTermination(v cfz.Expression[bool]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.DisableApiTermination = v
	return t
}

// SetV__DisableApiTermination updates property "DisableApiTermination".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__DisableApiTermination(v bool) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.DisableApiTermination = cfz.V(v)
	return t
}

// Set__EbsOptimized updates property "EbsOptimized".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__EbsOptimized(v cfz.Expression[bool]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.EbsOptimized = v
	return t
}

// SetV__EbsOptimized updates property "EbsOptimized".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__EbsOptimized(v bool) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.EbsOptimized = cfz.V(v)
	return t
}

// Set__ElasticGpuSpecifications updates property "ElasticGpuSpecifications".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__ElasticGpuSpecifications(v cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_ElasticGpuSpecification]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.ElasticGpuSpecifications = v
	return t
}

// SetS__ElasticGpuSpecifications updates property "ElasticGpuSpecifications".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetS__ElasticGpuSpecifications(v ...cfz.Expression[AWS_EC2_LaunchTemplate_ElasticGpuSpecification]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.ElasticGpuSpecifications = cfz.S(v...)
	return t
}

// SetSV__ElasticGpuSpecifications updates property "ElasticGpuSpecifications".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetSV__ElasticGpuSpecifications(v ...AWS_EC2_LaunchTemplate_ElasticGpuSpecification) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.ElasticGpuSpecifications = cfz.SV(v...)
	return t
}

// Set__ElasticInferenceAccelerators updates property "ElasticInferenceAccelerators".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__ElasticInferenceAccelerators(v cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_LaunchTemplateElasticInferenceAccelerator]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.ElasticInferenceAccelerators = v
	return t
}

// SetS__ElasticInferenceAccelerators updates property "ElasticInferenceAccelerators".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetS__ElasticInferenceAccelerators(v ...cfz.Expression[AWS_EC2_LaunchTemplate_LaunchTemplateElasticInferenceAccelerator]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.ElasticInferenceAccelerators = cfz.S(v...)
	return t
}

// SetSV__ElasticInferenceAccelerators updates property "ElasticInferenceAccelerators".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetSV__ElasticInferenceAccelerators(v ...AWS_EC2_LaunchTemplate_LaunchTemplateElasticInferenceAccelerator) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.ElasticInferenceAccelerators = cfz.SV(v...)
	return t
}

// Set__EnclaveOptions updates property "EnclaveOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__EnclaveOptions(v cfz.Expression[AWS_EC2_LaunchTemplate_EnclaveOptions]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.EnclaveOptions = v
	return t
}

// SetV__EnclaveOptions updates property "EnclaveOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__EnclaveOptions(v AWS_EC2_LaunchTemplate_EnclaveOptions) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.EnclaveOptions = cfz.V(v)
	return t
}

// Set__HibernationOptions updates property "HibernationOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__HibernationOptions(v cfz.Expression[AWS_EC2_LaunchTemplate_HibernationOptions]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.HibernationOptions = v
	return t
}

// SetV__HibernationOptions updates property "HibernationOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__HibernationOptions(v AWS_EC2_LaunchTemplate_HibernationOptions) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.HibernationOptions = cfz.V(v)
	return t
}

// Set__IamInstanceProfile updates property "IamInstanceProfile".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__IamInstanceProfile(v cfz.Expression[AWS_EC2_LaunchTemplate_IamInstanceProfile]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.IamInstanceProfile = v
	return t
}

// SetV__IamInstanceProfile updates property "IamInstanceProfile".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__IamInstanceProfile(v AWS_EC2_LaunchTemplate_IamInstanceProfile) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.IamInstanceProfile = cfz.V(v)
	return t
}

// Set__ImageId updates property "ImageId".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__ImageId(v cfz.Expression[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.ImageId = v
	return t
}

// SetV__ImageId updates property "ImageId".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__ImageId(v string) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.ImageId = cfz.V(v)
	return t
}

// Set__InstanceInitiatedShutdownBehavior updates property "InstanceInitiatedShutdownBehavior".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__InstanceInitiatedShutdownBehavior(v cfz.Expression[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.InstanceInitiatedShutdownBehavior = v
	return t
}

// SetV__InstanceInitiatedShutdownBehavior updates property "InstanceInitiatedShutdownBehavior".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__InstanceInitiatedShutdownBehavior(v string) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.InstanceInitiatedShutdownBehavior = cfz.V(v)
	return t
}

// Set__InstanceMarketOptions updates property "InstanceMarketOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__InstanceMarketOptions(v cfz.Expression[AWS_EC2_LaunchTemplate_InstanceMarketOptions]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.InstanceMarketOptions = v
	return t
}

// SetV__InstanceMarketOptions updates property "InstanceMarketOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__InstanceMarketOptions(v AWS_EC2_LaunchTemplate_InstanceMarketOptions) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.InstanceMarketOptions = cfz.V(v)
	return t
}

// Set__InstanceRequirements updates property "InstanceRequirements".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__InstanceRequirements(v cfz.Expression[AWS_EC2_LaunchTemplate_InstanceRequirements]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.InstanceRequirements = v
	return t
}

// SetV__InstanceRequirements updates property "InstanceRequirements".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__InstanceRequirements(v AWS_EC2_LaunchTemplate_InstanceRequirements) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.InstanceRequirements = cfz.V(v)
	return t
}

// Set__InstanceType updates property "InstanceType".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__InstanceType(v cfz.Expression[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.InstanceType = v
	return t
}

// SetV__InstanceType updates property "InstanceType".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__InstanceType(v string) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.InstanceType = cfz.V(v)
	return t
}

// Set__KernelId updates property "KernelId".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__KernelId(v cfz.Expression[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.KernelId = v
	return t
}

// SetV__KernelId updates property "KernelId".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__KernelId(v string) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.KernelId = cfz.V(v)
	return t
}

// Set__KeyName updates property "KeyName".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__KeyName(v cfz.Expression[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.KeyName = v
	return t
}

// SetV__KeyName updates property "KeyName".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__KeyName(v string) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.KeyName = cfz.V(v)
	return t
}

// Set__LicenseSpecifications updates property "LicenseSpecifications".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__LicenseSpecifications(v cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_LicenseSpecification]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.LicenseSpecifications = v
	return t
}

// SetS__LicenseSpecifications updates property "LicenseSpecifications".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetS__LicenseSpecifications(v ...cfz.Expression[AWS_EC2_LaunchTemplate_LicenseSpecification]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.LicenseSpecifications = cfz.S(v...)
	return t
}

// SetSV__LicenseSpecifications updates property "LicenseSpecifications".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetSV__LicenseSpecifications(v ...AWS_EC2_LaunchTemplate_LicenseSpecification) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.LicenseSpecifications = cfz.SV(v...)
	return t
}

// Set__MaintenanceOptions updates property "MaintenanceOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__MaintenanceOptions(v cfz.Expression[AWS_EC2_LaunchTemplate_MaintenanceOptions]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.MaintenanceOptions = v
	return t
}

// SetV__MaintenanceOptions updates property "MaintenanceOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__MaintenanceOptions(v AWS_EC2_LaunchTemplate_MaintenanceOptions) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.MaintenanceOptions = cfz.V(v)
	return t
}

// Set__MetadataOptions updates property "MetadataOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__MetadataOptions(v cfz.Expression[AWS_EC2_LaunchTemplate_MetadataOptions]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.MetadataOptions = v
	return t
}

// SetV__MetadataOptions updates property "MetadataOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__MetadataOptions(v AWS_EC2_LaunchTemplate_MetadataOptions) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.MetadataOptions = cfz.V(v)
	return t
}

// Set__Monitoring updates property "Monitoring".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__Monitoring(v cfz.Expression[AWS_EC2_LaunchTemplate_Monitoring]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.Monitoring = v
	return t
}

// SetV__Monitoring updates property "Monitoring".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__Monitoring(v AWS_EC2_LaunchTemplate_Monitoring) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.Monitoring = cfz.V(v)
	return t
}

// Set__NetworkInterfaces updates property "NetworkInterfaces".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__NetworkInterfaces(v cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_NetworkInterface]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.NetworkInterfaces = v
	return t
}

// SetS__NetworkInterfaces updates property "NetworkInterfaces".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetS__NetworkInterfaces(v ...cfz.Expression[AWS_EC2_LaunchTemplate_NetworkInterface]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.NetworkInterfaces = cfz.S(v...)
	return t
}

// SetSV__NetworkInterfaces updates property "NetworkInterfaces".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetSV__NetworkInterfaces(v ...AWS_EC2_LaunchTemplate_NetworkInterface) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.NetworkInterfaces = cfz.SV(v...)
	return t
}

// Set__Placement updates property "Placement".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__Placement(v cfz.Expression[AWS_EC2_LaunchTemplate_Placement]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.Placement = v
	return t
}

// SetV__Placement updates property "Placement".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__Placement(v AWS_EC2_LaunchTemplate_Placement) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.Placement = cfz.V(v)
	return t
}

// Set__PrivateDnsNameOptions updates property "PrivateDnsNameOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__PrivateDnsNameOptions(v cfz.Expression[AWS_EC2_LaunchTemplate_PrivateDnsNameOptions]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.PrivateDnsNameOptions = v
	return t
}

// SetV__PrivateDnsNameOptions updates property "PrivateDnsNameOptions".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__PrivateDnsNameOptions(v AWS_EC2_LaunchTemplate_PrivateDnsNameOptions) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.PrivateDnsNameOptions = cfz.V(v)
	return t
}

// Set__RamDiskId updates property "RamDiskId".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__RamDiskId(v cfz.Expression[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.RamDiskId = v
	return t
}

// SetV__RamDiskId updates property "RamDiskId".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__RamDiskId(v string) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.RamDiskId = cfz.V(v)
	return t
}

// Set__SecurityGroupIds updates property "SecurityGroupIds".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__SecurityGroupIds(v cfz.ExpressionSlice[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.SecurityGroupIds = v
	return t
}

// SetS__SecurityGroupIds updates property "SecurityGroupIds".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetS__SecurityGroupIds(v ...cfz.Expression[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.SecurityGroupIds = cfz.S(v...)
	return t
}

// SetSV__SecurityGroupIds updates property "SecurityGroupIds".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetSV__SecurityGroupIds(v ...string) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.SecurityGroupIds = cfz.SV(v...)
	return t
}

// Set__SecurityGroups updates property "SecurityGroups".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__SecurityGroups(v cfz.ExpressionSlice[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.SecurityGroups = v
	return t
}

// SetS__SecurityGroups updates property "SecurityGroups".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetS__SecurityGroups(v ...cfz.Expression[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.SecurityGroups = cfz.S(v...)
	return t
}

// SetSV__SecurityGroups updates property "SecurityGroups".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetSV__SecurityGroups(v ...string) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.SecurityGroups = cfz.SV(v...)
	return t
}

// Set__TagSpecifications updates property "TagSpecifications".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__TagSpecifications(v cfz.ExpressionSlice[AWS_EC2_LaunchTemplate_TagSpecification]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.TagSpecifications = v
	return t
}

// SetS__TagSpecifications updates property "TagSpecifications".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetS__TagSpecifications(v ...cfz.Expression[AWS_EC2_LaunchTemplate_TagSpecification]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.TagSpecifications = cfz.S(v...)
	return t
}

// SetSV__TagSpecifications updates property "TagSpecifications".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetSV__TagSpecifications(v ...AWS_EC2_LaunchTemplate_TagSpecification) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.TagSpecifications = cfz.SV(v...)
	return t
}

// Set__UserData updates property "UserData".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) Set__UserData(v cfz.Expression[string]) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.UserData = v
	return t
}

// SetV__UserData updates property "UserData".
func (t AWS_EC2_LaunchTemplate_LaunchTemplateData) SetV__UserData(v string) AWS_EC2_LaunchTemplate_LaunchTemplateData {
	t.UserData = cfz.V(v)
	return t
}
