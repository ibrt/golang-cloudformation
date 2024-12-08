// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_redshiftserverless

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_RedshiftServerless_Workgroup_Workgroup__Type is the CloudFormation type for AWS::RedshiftServerless::Workgroup.Workgroup.
	AWS_RedshiftServerless_Workgroup_Workgroup__Type = "AWS::RedshiftServerless::Workgroup.Workgroup"
)

var (
	// AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap reports all the CloudFormation properties for AWS::RedshiftServerless::Workgroup.Workgroup.
	AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap = struct {
		BaseCapacity       string
		ConfigParameters   string
		CreationDate       string
		Endpoint           string
		EnhancedVpcRouting string
		MaxCapacity        string
		NamespaceName      string
		PubliclyAccessible string
		SecurityGroupIds   string
		Status             string
		SubnetIds          string
		WorkgroupArn       string
		WorkgroupId        string
		WorkgroupName      string
	}{
		BaseCapacity:       "BaseCapacity",
		ConfigParameters:   "ConfigParameters",
		CreationDate:       "CreationDate",
		Endpoint:           "Endpoint",
		EnhancedVpcRouting: "EnhancedVpcRouting",
		MaxCapacity:        "MaxCapacity",
		NamespaceName:      "NamespaceName",
		PubliclyAccessible: "PubliclyAccessible",
		SecurityGroupIds:   "SecurityGroupIds",
		Status:             "Status",
		SubnetIds:          "SubnetIds",
		WorkgroupArn:       "WorkgroupArn",
		WorkgroupId:        "WorkgroupId",
		WorkgroupName:      "WorkgroupName",
	}

	// AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesSlice reports all the CloudFormation properties for AWS::RedshiftServerless::Workgroup.Workgroup.
	AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesSlice = []string{
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.BaseCapacity,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.ConfigParameters,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.CreationDate,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.Endpoint,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.EnhancedVpcRouting,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.MaxCapacity,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.NamespaceName,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.PubliclyAccessible,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.SecurityGroupIds,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.Status,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.SubnetIds,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.WorkgroupArn,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.WorkgroupId,
		AWS_RedshiftServerless_Workgroup_Workgroup__PropertiesMap.WorkgroupName,
	}
)

// AWS_RedshiftServerless_Workgroup_Workgroup is a binding for AWS::RedshiftServerless::Workgroup.Workgroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html
type AWS_RedshiftServerless_Workgroup_Workgroup struct {
	// BaseCapacity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-basecapacity
	BaseCapacity cfz.Expression[int32] `json:"BaseCapacity,omitempty"`

	// ConfigParameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-configparameters
	ConfigParameters cfz.ExpressionSlice[AWS_RedshiftServerless_Workgroup_ConfigParameter] `json:"ConfigParameters,omitempty"`

	// CreationDate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-creationdate
	CreationDate cfz.Expression[string] `json:"CreationDate,omitempty"`

	// Endpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-endpoint
	Endpoint cfz.Expression[AWS_RedshiftServerless_Workgroup_Endpoint] `json:"Endpoint,omitempty"`

	// EnhancedVpcRouting is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-enhancedvpcrouting
	EnhancedVpcRouting cfz.Expression[bool] `json:"EnhancedVpcRouting,omitempty"`

	// MaxCapacity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-maxcapacity
	MaxCapacity cfz.Expression[int32] `json:"MaxCapacity,omitempty"`

	// NamespaceName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-namespacename
	NamespaceName cfz.Expression[string] `json:"NamespaceName,omitempty"`

	// PubliclyAccessible is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-publiclyaccessible
	PubliclyAccessible cfz.Expression[bool] `json:"PubliclyAccessible,omitempty"`

	// SecurityGroupIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-securitygroupids
	SecurityGroupIds cfz.ExpressionSlice[string] `json:"SecurityGroupIds,omitempty"`

	// Status is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-status
	Status cfz.Expression[string] `json:"Status,omitempty"`

	// SubnetIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-subnetids
	SubnetIds cfz.ExpressionSlice[string] `json:"SubnetIds,omitempty"`

	// WorkgroupArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-workgrouparn
	WorkgroupArn cfz.Expression[string] `json:"WorkgroupArn,omitempty"`

	// WorkgroupId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-workgroupid
	WorkgroupId cfz.Expression[string] `json:"WorkgroupId,omitempty"`

	// WorkgroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-workgroupname
	WorkgroupName cfz.Expression[string] `json:"WorkgroupName,omitempty"`
}

// New__AWS_RedshiftServerless_Workgroup_Workgroup initializes a new AWS_RedshiftServerless_Workgroup_Workgroup.
func New__AWS_RedshiftServerless_Workgroup_Workgroup() AWS_RedshiftServerless_Workgroup_Workgroup {
	return AWS_RedshiftServerless_Workgroup_Workgroup{}
}

// GetType returns the CloudFormation type.
func (AWS_RedshiftServerless_Workgroup_Workgroup) GetType() string {
	return AWS_RedshiftServerless_Workgroup_Workgroup__Type
}

// Set__BaseCapacity updates property "BaseCapacity".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__BaseCapacity(v cfz.Expression[int32]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.BaseCapacity = v
	return t
}

// SetV__BaseCapacity updates property "BaseCapacity".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__BaseCapacity(v int32) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.BaseCapacity = cfz.V(v)
	return t
}

// Set__ConfigParameters updates property "ConfigParameters".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__ConfigParameters(v cfz.ExpressionSlice[AWS_RedshiftServerless_Workgroup_ConfigParameter]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.ConfigParameters = v
	return t
}

// SetS__ConfigParameters updates property "ConfigParameters".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetS__ConfigParameters(v ...cfz.Expression[AWS_RedshiftServerless_Workgroup_ConfigParameter]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.ConfigParameters = cfz.S(v...)
	return t
}

// SetSV__ConfigParameters updates property "ConfigParameters".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetSV__ConfigParameters(v ...AWS_RedshiftServerless_Workgroup_ConfigParameter) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.ConfigParameters = cfz.SV(v...)
	return t
}

// Set__CreationDate updates property "CreationDate".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__CreationDate(v cfz.Expression[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.CreationDate = v
	return t
}

// SetV__CreationDate updates property "CreationDate".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__CreationDate(v string) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.CreationDate = cfz.V(v)
	return t
}

// Set__Endpoint updates property "Endpoint".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__Endpoint(v cfz.Expression[AWS_RedshiftServerless_Workgroup_Endpoint]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.Endpoint = v
	return t
}

// SetV__Endpoint updates property "Endpoint".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__Endpoint(v AWS_RedshiftServerless_Workgroup_Endpoint) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.Endpoint = cfz.V(v)
	return t
}

// Set__EnhancedVpcRouting updates property "EnhancedVpcRouting".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__EnhancedVpcRouting(v cfz.Expression[bool]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.EnhancedVpcRouting = v
	return t
}

// SetV__EnhancedVpcRouting updates property "EnhancedVpcRouting".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__EnhancedVpcRouting(v bool) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.EnhancedVpcRouting = cfz.V(v)
	return t
}

// Set__MaxCapacity updates property "MaxCapacity".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__MaxCapacity(v cfz.Expression[int32]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.MaxCapacity = v
	return t
}

// SetV__MaxCapacity updates property "MaxCapacity".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__MaxCapacity(v int32) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.MaxCapacity = cfz.V(v)
	return t
}

// Set__NamespaceName updates property "NamespaceName".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__NamespaceName(v cfz.Expression[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.NamespaceName = v
	return t
}

// SetV__NamespaceName updates property "NamespaceName".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__NamespaceName(v string) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.NamespaceName = cfz.V(v)
	return t
}

// Set__PubliclyAccessible updates property "PubliclyAccessible".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__PubliclyAccessible(v cfz.Expression[bool]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.PubliclyAccessible = v
	return t
}

// SetV__PubliclyAccessible updates property "PubliclyAccessible".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__PubliclyAccessible(v bool) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.PubliclyAccessible = cfz.V(v)
	return t
}

// Set__SecurityGroupIds updates property "SecurityGroupIds".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__SecurityGroupIds(v cfz.ExpressionSlice[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.SecurityGroupIds = v
	return t
}

// SetS__SecurityGroupIds updates property "SecurityGroupIds".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetS__SecurityGroupIds(v ...cfz.Expression[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.SecurityGroupIds = cfz.S(v...)
	return t
}

// SetSV__SecurityGroupIds updates property "SecurityGroupIds".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetSV__SecurityGroupIds(v ...string) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.SecurityGroupIds = cfz.SV(v...)
	return t
}

// Set__Status updates property "Status".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__Status(v cfz.Expression[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.Status = v
	return t
}

// SetV__Status updates property "Status".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__Status(v string) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.Status = cfz.V(v)
	return t
}

// Set__SubnetIds updates property "SubnetIds".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__SubnetIds(v cfz.ExpressionSlice[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.SubnetIds = v
	return t
}

// SetS__SubnetIds updates property "SubnetIds".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetS__SubnetIds(v ...cfz.Expression[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.SubnetIds = cfz.S(v...)
	return t
}

// SetSV__SubnetIds updates property "SubnetIds".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetSV__SubnetIds(v ...string) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.SubnetIds = cfz.SV(v...)
	return t
}

// Set__WorkgroupArn updates property "WorkgroupArn".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__WorkgroupArn(v cfz.Expression[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.WorkgroupArn = v
	return t
}

// SetV__WorkgroupArn updates property "WorkgroupArn".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__WorkgroupArn(v string) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.WorkgroupArn = cfz.V(v)
	return t
}

// Set__WorkgroupId updates property "WorkgroupId".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__WorkgroupId(v cfz.Expression[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.WorkgroupId = v
	return t
}

// SetV__WorkgroupId updates property "WorkgroupId".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__WorkgroupId(v string) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.WorkgroupId = cfz.V(v)
	return t
}

// Set__WorkgroupName updates property "WorkgroupName".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) Set__WorkgroupName(v cfz.Expression[string]) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.WorkgroupName = v
	return t
}

// SetV__WorkgroupName updates property "WorkgroupName".
func (t AWS_RedshiftServerless_Workgroup_Workgroup) SetV__WorkgroupName(v string) AWS_RedshiftServerless_Workgroup_Workgroup {
	t.WorkgroupName = cfz.V(v)
	return t
}
