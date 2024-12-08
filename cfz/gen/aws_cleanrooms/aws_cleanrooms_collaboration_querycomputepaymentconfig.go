// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cleanrooms

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CleanRooms_Collaboration_QueryComputePaymentConfig__Type is the CloudFormation type for AWS::CleanRooms::Collaboration.QueryComputePaymentConfig.
	AWS_CleanRooms_Collaboration_QueryComputePaymentConfig__Type = "AWS::CleanRooms::Collaboration.QueryComputePaymentConfig"
)

var (
	// AWS_CleanRooms_Collaboration_QueryComputePaymentConfig__PropertiesMap reports all the CloudFormation properties for AWS::CleanRooms::Collaboration.QueryComputePaymentConfig.
	AWS_CleanRooms_Collaboration_QueryComputePaymentConfig__PropertiesMap = struct {
		IsResponsible string
	}{
		IsResponsible: "IsResponsible",
	}

	// AWS_CleanRooms_Collaboration_QueryComputePaymentConfig__PropertiesSlice reports all the CloudFormation properties for AWS::CleanRooms::Collaboration.QueryComputePaymentConfig.
	AWS_CleanRooms_Collaboration_QueryComputePaymentConfig__PropertiesSlice = []string{
		AWS_CleanRooms_Collaboration_QueryComputePaymentConfig__PropertiesMap.IsResponsible,
	}
)

// AWS_CleanRooms_Collaboration_QueryComputePaymentConfig is a binding for AWS::CleanRooms::Collaboration.QueryComputePaymentConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-querycomputepaymentconfig.html
type AWS_CleanRooms_Collaboration_QueryComputePaymentConfig struct {
	// IsResponsible is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-querycomputepaymentconfig.html#cfn-cleanrooms-collaboration-querycomputepaymentconfig-isresponsible
	IsResponsible cfz.Expression[bool] `json:"IsResponsible,omitempty"`
}

// New__AWS_CleanRooms_Collaboration_QueryComputePaymentConfig initializes a new AWS_CleanRooms_Collaboration_QueryComputePaymentConfig.
func New__AWS_CleanRooms_Collaboration_QueryComputePaymentConfig() AWS_CleanRooms_Collaboration_QueryComputePaymentConfig {
	return AWS_CleanRooms_Collaboration_QueryComputePaymentConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_CleanRooms_Collaboration_QueryComputePaymentConfig) GetType() string {
	return AWS_CleanRooms_Collaboration_QueryComputePaymentConfig__Type
}

// Set__IsResponsible updates property "IsResponsible".
func (t AWS_CleanRooms_Collaboration_QueryComputePaymentConfig) Set__IsResponsible(v cfz.Expression[bool]) AWS_CleanRooms_Collaboration_QueryComputePaymentConfig {
	t.IsResponsible = v
	return t
}

// SetV__IsResponsible updates property "IsResponsible".
func (t AWS_CleanRooms_Collaboration_QueryComputePaymentConfig) SetV__IsResponsible(v bool) AWS_CleanRooms_Collaboration_QueryComputePaymentConfig {
	t.IsResponsible = cfz.V(v)
	return t
}
