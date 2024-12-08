// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_eks

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EKS_Nodegroup_LaunchTemplateSpecification__Type is the CloudFormation type for AWS::EKS::Nodegroup.LaunchTemplateSpecification.
	AWS_EKS_Nodegroup_LaunchTemplateSpecification__Type = "AWS::EKS::Nodegroup.LaunchTemplateSpecification"
)

var (
	// AWS_EKS_Nodegroup_LaunchTemplateSpecification__PropertiesMap reports all the CloudFormation properties for AWS::EKS::Nodegroup.LaunchTemplateSpecification.
	AWS_EKS_Nodegroup_LaunchTemplateSpecification__PropertiesMap = struct {
		Id      string
		Name    string
		Version string
	}{
		Id:      "Id",
		Name:    "Name",
		Version: "Version",
	}

	// AWS_EKS_Nodegroup_LaunchTemplateSpecification__PropertiesSlice reports all the CloudFormation properties for AWS::EKS::Nodegroup.LaunchTemplateSpecification.
	AWS_EKS_Nodegroup_LaunchTemplateSpecification__PropertiesSlice = []string{
		AWS_EKS_Nodegroup_LaunchTemplateSpecification__PropertiesMap.Id,
		AWS_EKS_Nodegroup_LaunchTemplateSpecification__PropertiesMap.Name,
		AWS_EKS_Nodegroup_LaunchTemplateSpecification__PropertiesMap.Version,
	}
)

// AWS_EKS_Nodegroup_LaunchTemplateSpecification is a binding for AWS::EKS::Nodegroup.LaunchTemplateSpecification.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-launchtemplatespecification.html
type AWS_EKS_Nodegroup_LaunchTemplateSpecification struct {
	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-launchtemplatespecification.html#cfn-eks-nodegroup-launchtemplatespecification-id
	Id cfz.Expression[string] `json:"Id,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-launchtemplatespecification.html#cfn-eks-nodegroup-launchtemplatespecification-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Version is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-launchtemplatespecification.html#cfn-eks-nodegroup-launchtemplatespecification-version
	Version cfz.Expression[string] `json:"Version,omitempty"`
}

// New__AWS_EKS_Nodegroup_LaunchTemplateSpecification initializes a new AWS_EKS_Nodegroup_LaunchTemplateSpecification.
func New__AWS_EKS_Nodegroup_LaunchTemplateSpecification() AWS_EKS_Nodegroup_LaunchTemplateSpecification {
	return AWS_EKS_Nodegroup_LaunchTemplateSpecification{}
}

// GetType returns the CloudFormation type.
func (AWS_EKS_Nodegroup_LaunchTemplateSpecification) GetType() string {
	return AWS_EKS_Nodegroup_LaunchTemplateSpecification__Type
}

// Set__Id updates property "Id".
func (t AWS_EKS_Nodegroup_LaunchTemplateSpecification) Set__Id(v cfz.Expression[string]) AWS_EKS_Nodegroup_LaunchTemplateSpecification {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t AWS_EKS_Nodegroup_LaunchTemplateSpecification) SetV__Id(v string) AWS_EKS_Nodegroup_LaunchTemplateSpecification {
	t.Id = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_EKS_Nodegroup_LaunchTemplateSpecification) Set__Name(v cfz.Expression[string]) AWS_EKS_Nodegroup_LaunchTemplateSpecification {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_EKS_Nodegroup_LaunchTemplateSpecification) SetV__Name(v string) AWS_EKS_Nodegroup_LaunchTemplateSpecification {
	t.Name = cfz.V(v)
	return t
}

// Set__Version updates property "Version".
func (t AWS_EKS_Nodegroup_LaunchTemplateSpecification) Set__Version(v cfz.Expression[string]) AWS_EKS_Nodegroup_LaunchTemplateSpecification {
	t.Version = v
	return t
}

// SetV__Version updates property "Version".
func (t AWS_EKS_Nodegroup_LaunchTemplateSpecification) SetV__Version(v string) AWS_EKS_Nodegroup_LaunchTemplateSpecification {
	t.Version = cfz.V(v)
	return t
}
