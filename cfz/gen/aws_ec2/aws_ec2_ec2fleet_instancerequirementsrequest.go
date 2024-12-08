// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_EC2Fleet_InstanceRequirementsRequest__Type is the CloudFormation type for AWS::EC2::EC2Fleet.InstanceRequirementsRequest.
	AWS_EC2_EC2Fleet_InstanceRequirementsRequest__Type = "AWS::EC2::EC2Fleet.InstanceRequirementsRequest"
)

var (
	// AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap reports all the CloudFormation properties for AWS::EC2::EC2Fleet.InstanceRequirementsRequest.
	AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap = struct {
		AcceleratorCount                               string
		AcceleratorManufacturers                       string
		AcceleratorNames                               string
		AcceleratorTotalMemoryMiB                      string
		AcceleratorTypes                               string
		AllowedInstanceTypes                           string
		BareMetal                                      string
		BaselineEbsBandwidthMbps                       string
		BurstablePerformance                           string
		CpuManufacturers                               string
		ExcludedInstanceTypes                          string
		InstanceGenerations                            string
		LocalStorage                                   string
		LocalStorageTypes                              string
		MaxSpotPriceAsPercentageOfOptimalOnDemandPrice string
		MemoryGiBPerVCpu                               string
		MemoryMiB                                      string
		NetworkBandwidthGbps                           string
		NetworkInterfaceCount                          string
		OnDemandMaxPricePercentageOverLowestPrice      string
		RequireHibernateSupport                        string
		SpotMaxPricePercentageOverLowestPrice          string
		TotalLocalStorageGB                            string
		VCpuCount                                      string
	}{
		AcceleratorCount:          "AcceleratorCount",
		AcceleratorManufacturers:  "AcceleratorManufacturers",
		AcceleratorNames:          "AcceleratorNames",
		AcceleratorTotalMemoryMiB: "AcceleratorTotalMemoryMiB",
		AcceleratorTypes:          "AcceleratorTypes",
		AllowedInstanceTypes:      "AllowedInstanceTypes",
		BareMetal:                 "BareMetal",
		BaselineEbsBandwidthMbps:  "BaselineEbsBandwidthMbps",
		BurstablePerformance:      "BurstablePerformance",
		CpuManufacturers:          "CpuManufacturers",
		ExcludedInstanceTypes:     "ExcludedInstanceTypes",
		InstanceGenerations:       "InstanceGenerations",
		LocalStorage:              "LocalStorage",
		LocalStorageTypes:         "LocalStorageTypes",
		MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: "MaxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		MemoryGiBPerVCpu:      "MemoryGiBPerVCpu",
		MemoryMiB:             "MemoryMiB",
		NetworkBandwidthGbps:  "NetworkBandwidthGbps",
		NetworkInterfaceCount: "NetworkInterfaceCount",
		OnDemandMaxPricePercentageOverLowestPrice: "OnDemandMaxPricePercentageOverLowestPrice",
		RequireHibernateSupport:                   "RequireHibernateSupport",
		SpotMaxPricePercentageOverLowestPrice:     "SpotMaxPricePercentageOverLowestPrice",
		TotalLocalStorageGB:                       "TotalLocalStorageGB",
		VCpuCount:                                 "VCpuCount",
	}

	// AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::EC2Fleet.InstanceRequirementsRequest.
	AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesSlice = []string{
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.AcceleratorCount,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.AcceleratorManufacturers,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.AcceleratorNames,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.AcceleratorTotalMemoryMiB,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.AcceleratorTypes,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.AllowedInstanceTypes,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.BareMetal,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.BaselineEbsBandwidthMbps,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.BurstablePerformance,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.CpuManufacturers,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.ExcludedInstanceTypes,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.InstanceGenerations,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.LocalStorage,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.LocalStorageTypes,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.MaxSpotPriceAsPercentageOfOptimalOnDemandPrice,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.MemoryGiBPerVCpu,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.MemoryMiB,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.NetworkBandwidthGbps,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.NetworkInterfaceCount,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.OnDemandMaxPricePercentageOverLowestPrice,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.RequireHibernateSupport,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.SpotMaxPricePercentageOverLowestPrice,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.TotalLocalStorageGB,
		AWS_EC2_EC2Fleet_InstanceRequirementsRequest__PropertiesMap.VCpuCount,
	}
)

// AWS_EC2_EC2Fleet_InstanceRequirementsRequest is a binding for AWS::EC2::EC2Fleet.InstanceRequirementsRequest.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html
type AWS_EC2_EC2Fleet_InstanceRequirementsRequest struct {
	// AcceleratorCount is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-acceleratorcount
	AcceleratorCount cfz.Expression[AWS_EC2_EC2Fleet_AcceleratorCountRequest] `json:"AcceleratorCount,omitempty"`

	// AcceleratorManufacturers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-acceleratormanufacturers
	AcceleratorManufacturers cfz.ExpressionSlice[string] `json:"AcceleratorManufacturers,omitempty"`

	// AcceleratorNames is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-acceleratornames
	AcceleratorNames cfz.ExpressionSlice[string] `json:"AcceleratorNames,omitempty"`

	// AcceleratorTotalMemoryMiB is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-acceleratortotalmemorymib
	AcceleratorTotalMemoryMiB cfz.Expression[AWS_EC2_EC2Fleet_AcceleratorTotalMemoryMiBRequest] `json:"AcceleratorTotalMemoryMiB,omitempty"`

	// AcceleratorTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-acceleratortypes
	AcceleratorTypes cfz.ExpressionSlice[string] `json:"AcceleratorTypes,omitempty"`

	// AllowedInstanceTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-allowedinstancetypes
	AllowedInstanceTypes cfz.ExpressionSlice[string] `json:"AllowedInstanceTypes,omitempty"`

	// BareMetal is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-baremetal
	BareMetal cfz.Expression[string] `json:"BareMetal,omitempty"`

	// BaselineEbsBandwidthMbps is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-baselineebsbandwidthmbps
	BaselineEbsBandwidthMbps cfz.Expression[AWS_EC2_EC2Fleet_BaselineEbsBandwidthMbpsRequest] `json:"BaselineEbsBandwidthMbps,omitempty"`

	// BurstablePerformance is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-burstableperformance
	BurstablePerformance cfz.Expression[string] `json:"BurstablePerformance,omitempty"`

	// CpuManufacturers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-cpumanufacturers
	CpuManufacturers cfz.ExpressionSlice[string] `json:"CpuManufacturers,omitempty"`

	// ExcludedInstanceTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-excludedinstancetypes
	ExcludedInstanceTypes cfz.ExpressionSlice[string] `json:"ExcludedInstanceTypes,omitempty"`

	// InstanceGenerations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-instancegenerations
	InstanceGenerations cfz.ExpressionSlice[string] `json:"InstanceGenerations,omitempty"`

	// LocalStorage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-localstorage
	LocalStorage cfz.Expression[string] `json:"LocalStorage,omitempty"`

	// LocalStorageTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-localstoragetypes
	LocalStorageTypes cfz.ExpressionSlice[string] `json:"LocalStorageTypes,omitempty"`

	// MaxSpotPriceAsPercentageOfOptimalOnDemandPrice is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-maxspotpriceaspercentageofoptimalondemandprice
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice cfz.Expression[int32] `json:"MaxSpotPriceAsPercentageOfOptimalOnDemandPrice,omitempty"`

	// MemoryGiBPerVCpu is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-memorygibpervcpu
	MemoryGiBPerVCpu cfz.Expression[AWS_EC2_EC2Fleet_MemoryGiBPerVCpuRequest] `json:"MemoryGiBPerVCpu,omitempty"`

	// MemoryMiB is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-memorymib
	MemoryMiB cfz.Expression[AWS_EC2_EC2Fleet_MemoryMiBRequest] `json:"MemoryMiB,omitempty"`

	// NetworkBandwidthGbps is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-networkbandwidthgbps
	NetworkBandwidthGbps cfz.Expression[AWS_EC2_EC2Fleet_NetworkBandwidthGbpsRequest] `json:"NetworkBandwidthGbps,omitempty"`

	// NetworkInterfaceCount is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-networkinterfacecount
	NetworkInterfaceCount cfz.Expression[AWS_EC2_EC2Fleet_NetworkInterfaceCountRequest] `json:"NetworkInterfaceCount,omitempty"`

	// OnDemandMaxPricePercentageOverLowestPrice is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-ondemandmaxpricepercentageoverlowestprice
	OnDemandMaxPricePercentageOverLowestPrice cfz.Expression[int32] `json:"OnDemandMaxPricePercentageOverLowestPrice,omitempty"`

	// RequireHibernateSupport is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-requirehibernatesupport
	RequireHibernateSupport cfz.Expression[bool] `json:"RequireHibernateSupport,omitempty"`

	// SpotMaxPricePercentageOverLowestPrice is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-spotmaxpricepercentageoverlowestprice
	SpotMaxPricePercentageOverLowestPrice cfz.Expression[int32] `json:"SpotMaxPricePercentageOverLowestPrice,omitempty"`

	// TotalLocalStorageGB is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-totallocalstoragegb
	TotalLocalStorageGB cfz.Expression[AWS_EC2_EC2Fleet_TotalLocalStorageGBRequest] `json:"TotalLocalStorageGB,omitempty"`

	// VCpuCount is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancerequirementsrequest.html#cfn-ec2-ec2fleet-instancerequirementsrequest-vcpucount
	VCpuCount cfz.Expression[AWS_EC2_EC2Fleet_VCpuCountRangeRequest] `json:"VCpuCount,omitempty"`
}

// New__AWS_EC2_EC2Fleet_InstanceRequirementsRequest initializes a new AWS_EC2_EC2Fleet_InstanceRequirementsRequest.
func New__AWS_EC2_EC2Fleet_InstanceRequirementsRequest() AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	return AWS_EC2_EC2Fleet_InstanceRequirementsRequest{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_EC2Fleet_InstanceRequirementsRequest) GetType() string {
	return AWS_EC2_EC2Fleet_InstanceRequirementsRequest__Type
}

// Set__AcceleratorCount updates property "AcceleratorCount".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__AcceleratorCount(v cfz.Expression[AWS_EC2_EC2Fleet_AcceleratorCountRequest]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorCount = v
	return t
}

// SetV__AcceleratorCount updates property "AcceleratorCount".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__AcceleratorCount(v AWS_EC2_EC2Fleet_AcceleratorCountRequest) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorCount = cfz.V(v)
	return t
}

// Set__AcceleratorManufacturers updates property "AcceleratorManufacturers".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__AcceleratorManufacturers(v cfz.ExpressionSlice[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorManufacturers = v
	return t
}

// SetS__AcceleratorManufacturers updates property "AcceleratorManufacturers".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetS__AcceleratorManufacturers(v ...cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorManufacturers = cfz.S(v...)
	return t
}

// SetSV__AcceleratorManufacturers updates property "AcceleratorManufacturers".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetSV__AcceleratorManufacturers(v ...string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorManufacturers = cfz.SV(v...)
	return t
}

// Set__AcceleratorNames updates property "AcceleratorNames".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__AcceleratorNames(v cfz.ExpressionSlice[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorNames = v
	return t
}

// SetS__AcceleratorNames updates property "AcceleratorNames".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetS__AcceleratorNames(v ...cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorNames = cfz.S(v...)
	return t
}

// SetSV__AcceleratorNames updates property "AcceleratorNames".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetSV__AcceleratorNames(v ...string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorNames = cfz.SV(v...)
	return t
}

// Set__AcceleratorTotalMemoryMiB updates property "AcceleratorTotalMemoryMiB".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__AcceleratorTotalMemoryMiB(v cfz.Expression[AWS_EC2_EC2Fleet_AcceleratorTotalMemoryMiBRequest]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorTotalMemoryMiB = v
	return t
}

// SetV__AcceleratorTotalMemoryMiB updates property "AcceleratorTotalMemoryMiB".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__AcceleratorTotalMemoryMiB(v AWS_EC2_EC2Fleet_AcceleratorTotalMemoryMiBRequest) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorTotalMemoryMiB = cfz.V(v)
	return t
}

// Set__AcceleratorTypes updates property "AcceleratorTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__AcceleratorTypes(v cfz.ExpressionSlice[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorTypes = v
	return t
}

// SetS__AcceleratorTypes updates property "AcceleratorTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetS__AcceleratorTypes(v ...cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorTypes = cfz.S(v...)
	return t
}

// SetSV__AcceleratorTypes updates property "AcceleratorTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetSV__AcceleratorTypes(v ...string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AcceleratorTypes = cfz.SV(v...)
	return t
}

// Set__AllowedInstanceTypes updates property "AllowedInstanceTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__AllowedInstanceTypes(v cfz.ExpressionSlice[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AllowedInstanceTypes = v
	return t
}

// SetS__AllowedInstanceTypes updates property "AllowedInstanceTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetS__AllowedInstanceTypes(v ...cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AllowedInstanceTypes = cfz.S(v...)
	return t
}

// SetSV__AllowedInstanceTypes updates property "AllowedInstanceTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetSV__AllowedInstanceTypes(v ...string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.AllowedInstanceTypes = cfz.SV(v...)
	return t
}

// Set__BareMetal updates property "BareMetal".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__BareMetal(v cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.BareMetal = v
	return t
}

// SetV__BareMetal updates property "BareMetal".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__BareMetal(v string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.BareMetal = cfz.V(v)
	return t
}

// Set__BaselineEbsBandwidthMbps updates property "BaselineEbsBandwidthMbps".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__BaselineEbsBandwidthMbps(v cfz.Expression[AWS_EC2_EC2Fleet_BaselineEbsBandwidthMbpsRequest]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.BaselineEbsBandwidthMbps = v
	return t
}

// SetV__BaselineEbsBandwidthMbps updates property "BaselineEbsBandwidthMbps".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__BaselineEbsBandwidthMbps(v AWS_EC2_EC2Fleet_BaselineEbsBandwidthMbpsRequest) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.BaselineEbsBandwidthMbps = cfz.V(v)
	return t
}

// Set__BurstablePerformance updates property "BurstablePerformance".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__BurstablePerformance(v cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.BurstablePerformance = v
	return t
}

// SetV__BurstablePerformance updates property "BurstablePerformance".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__BurstablePerformance(v string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.BurstablePerformance = cfz.V(v)
	return t
}

// Set__CpuManufacturers updates property "CpuManufacturers".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__CpuManufacturers(v cfz.ExpressionSlice[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.CpuManufacturers = v
	return t
}

// SetS__CpuManufacturers updates property "CpuManufacturers".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetS__CpuManufacturers(v ...cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.CpuManufacturers = cfz.S(v...)
	return t
}

// SetSV__CpuManufacturers updates property "CpuManufacturers".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetSV__CpuManufacturers(v ...string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.CpuManufacturers = cfz.SV(v...)
	return t
}

// Set__ExcludedInstanceTypes updates property "ExcludedInstanceTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__ExcludedInstanceTypes(v cfz.ExpressionSlice[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.ExcludedInstanceTypes = v
	return t
}

// SetS__ExcludedInstanceTypes updates property "ExcludedInstanceTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetS__ExcludedInstanceTypes(v ...cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.ExcludedInstanceTypes = cfz.S(v...)
	return t
}

// SetSV__ExcludedInstanceTypes updates property "ExcludedInstanceTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetSV__ExcludedInstanceTypes(v ...string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.ExcludedInstanceTypes = cfz.SV(v...)
	return t
}

// Set__InstanceGenerations updates property "InstanceGenerations".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__InstanceGenerations(v cfz.ExpressionSlice[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.InstanceGenerations = v
	return t
}

// SetS__InstanceGenerations updates property "InstanceGenerations".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetS__InstanceGenerations(v ...cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.InstanceGenerations = cfz.S(v...)
	return t
}

// SetSV__InstanceGenerations updates property "InstanceGenerations".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetSV__InstanceGenerations(v ...string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.InstanceGenerations = cfz.SV(v...)
	return t
}

// Set__LocalStorage updates property "LocalStorage".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__LocalStorage(v cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.LocalStorage = v
	return t
}

// SetV__LocalStorage updates property "LocalStorage".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__LocalStorage(v string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.LocalStorage = cfz.V(v)
	return t
}

// Set__LocalStorageTypes updates property "LocalStorageTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__LocalStorageTypes(v cfz.ExpressionSlice[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.LocalStorageTypes = v
	return t
}

// SetS__LocalStorageTypes updates property "LocalStorageTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetS__LocalStorageTypes(v ...cfz.Expression[string]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.LocalStorageTypes = cfz.S(v...)
	return t
}

// SetSV__LocalStorageTypes updates property "LocalStorageTypes".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetSV__LocalStorageTypes(v ...string) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.LocalStorageTypes = cfz.SV(v...)
	return t
}

// Set__MaxSpotPriceAsPercentageOfOptimalOnDemandPrice updates property "MaxSpotPriceAsPercentageOfOptimalOnDemandPrice".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__MaxSpotPriceAsPercentageOfOptimalOnDemandPrice(v cfz.Expression[int32]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.MaxSpotPriceAsPercentageOfOptimalOnDemandPrice = v
	return t
}

// SetV__MaxSpotPriceAsPercentageOfOptimalOnDemandPrice updates property "MaxSpotPriceAsPercentageOfOptimalOnDemandPrice".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__MaxSpotPriceAsPercentageOfOptimalOnDemandPrice(v int32) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.MaxSpotPriceAsPercentageOfOptimalOnDemandPrice = cfz.V(v)
	return t
}

// Set__MemoryGiBPerVCpu updates property "MemoryGiBPerVCpu".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__MemoryGiBPerVCpu(v cfz.Expression[AWS_EC2_EC2Fleet_MemoryGiBPerVCpuRequest]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.MemoryGiBPerVCpu = v
	return t
}

// SetV__MemoryGiBPerVCpu updates property "MemoryGiBPerVCpu".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__MemoryGiBPerVCpu(v AWS_EC2_EC2Fleet_MemoryGiBPerVCpuRequest) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.MemoryGiBPerVCpu = cfz.V(v)
	return t
}

// Set__MemoryMiB updates property "MemoryMiB".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__MemoryMiB(v cfz.Expression[AWS_EC2_EC2Fleet_MemoryMiBRequest]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.MemoryMiB = v
	return t
}

// SetV__MemoryMiB updates property "MemoryMiB".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__MemoryMiB(v AWS_EC2_EC2Fleet_MemoryMiBRequest) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.MemoryMiB = cfz.V(v)
	return t
}

// Set__NetworkBandwidthGbps updates property "NetworkBandwidthGbps".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__NetworkBandwidthGbps(v cfz.Expression[AWS_EC2_EC2Fleet_NetworkBandwidthGbpsRequest]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.NetworkBandwidthGbps = v
	return t
}

// SetV__NetworkBandwidthGbps updates property "NetworkBandwidthGbps".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__NetworkBandwidthGbps(v AWS_EC2_EC2Fleet_NetworkBandwidthGbpsRequest) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.NetworkBandwidthGbps = cfz.V(v)
	return t
}

// Set__NetworkInterfaceCount updates property "NetworkInterfaceCount".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__NetworkInterfaceCount(v cfz.Expression[AWS_EC2_EC2Fleet_NetworkInterfaceCountRequest]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.NetworkInterfaceCount = v
	return t
}

// SetV__NetworkInterfaceCount updates property "NetworkInterfaceCount".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__NetworkInterfaceCount(v AWS_EC2_EC2Fleet_NetworkInterfaceCountRequest) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.NetworkInterfaceCount = cfz.V(v)
	return t
}

// Set__OnDemandMaxPricePercentageOverLowestPrice updates property "OnDemandMaxPricePercentageOverLowestPrice".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__OnDemandMaxPricePercentageOverLowestPrice(v cfz.Expression[int32]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.OnDemandMaxPricePercentageOverLowestPrice = v
	return t
}

// SetV__OnDemandMaxPricePercentageOverLowestPrice updates property "OnDemandMaxPricePercentageOverLowestPrice".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__OnDemandMaxPricePercentageOverLowestPrice(v int32) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.OnDemandMaxPricePercentageOverLowestPrice = cfz.V(v)
	return t
}

// Set__RequireHibernateSupport updates property "RequireHibernateSupport".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__RequireHibernateSupport(v cfz.Expression[bool]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.RequireHibernateSupport = v
	return t
}

// SetV__RequireHibernateSupport updates property "RequireHibernateSupport".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__RequireHibernateSupport(v bool) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.RequireHibernateSupport = cfz.V(v)
	return t
}

// Set__SpotMaxPricePercentageOverLowestPrice updates property "SpotMaxPricePercentageOverLowestPrice".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__SpotMaxPricePercentageOverLowestPrice(v cfz.Expression[int32]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.SpotMaxPricePercentageOverLowestPrice = v
	return t
}

// SetV__SpotMaxPricePercentageOverLowestPrice updates property "SpotMaxPricePercentageOverLowestPrice".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__SpotMaxPricePercentageOverLowestPrice(v int32) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.SpotMaxPricePercentageOverLowestPrice = cfz.V(v)
	return t
}

// Set__TotalLocalStorageGB updates property "TotalLocalStorageGB".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__TotalLocalStorageGB(v cfz.Expression[AWS_EC2_EC2Fleet_TotalLocalStorageGBRequest]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.TotalLocalStorageGB = v
	return t
}

// SetV__TotalLocalStorageGB updates property "TotalLocalStorageGB".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__TotalLocalStorageGB(v AWS_EC2_EC2Fleet_TotalLocalStorageGBRequest) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.TotalLocalStorageGB = cfz.V(v)
	return t
}

// Set__VCpuCount updates property "VCpuCount".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) Set__VCpuCount(v cfz.Expression[AWS_EC2_EC2Fleet_VCpuCountRangeRequest]) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.VCpuCount = v
	return t
}

// SetV__VCpuCount updates property "VCpuCount".
func (t AWS_EC2_EC2Fleet_InstanceRequirementsRequest) SetV__VCpuCount(v AWS_EC2_EC2Fleet_VCpuCountRangeRequest) AWS_EC2_EC2Fleet_InstanceRequirementsRequest {
	t.VCpuCount = cfz.V(v)
	return t
}
