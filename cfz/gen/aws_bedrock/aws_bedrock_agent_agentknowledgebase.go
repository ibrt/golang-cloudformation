// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Agent_AgentKnowledgeBase__Type is the CloudFormation type for AWS::Bedrock::Agent.AgentKnowledgeBase.
	AWS_Bedrock_Agent_AgentKnowledgeBase__Type = "AWS::Bedrock::Agent.AgentKnowledgeBase"
)

var (
	// AWS_Bedrock_Agent_AgentKnowledgeBase__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Agent.AgentKnowledgeBase.
	AWS_Bedrock_Agent_AgentKnowledgeBase__PropertiesMap = struct {
		Description        string
		KnowledgeBaseId    string
		KnowledgeBaseState string
	}{
		Description:        "Description",
		KnowledgeBaseId:    "KnowledgeBaseId",
		KnowledgeBaseState: "KnowledgeBaseState",
	}

	// AWS_Bedrock_Agent_AgentKnowledgeBase__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Agent.AgentKnowledgeBase.
	AWS_Bedrock_Agent_AgentKnowledgeBase__PropertiesSlice = []string{
		AWS_Bedrock_Agent_AgentKnowledgeBase__PropertiesMap.Description,
		AWS_Bedrock_Agent_AgentKnowledgeBase__PropertiesMap.KnowledgeBaseId,
		AWS_Bedrock_Agent_AgentKnowledgeBase__PropertiesMap.KnowledgeBaseState,
	}
)

// AWS_Bedrock_Agent_AgentKnowledgeBase is a binding for AWS::Bedrock::Agent.AgentKnowledgeBase.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentknowledgebase.html
type AWS_Bedrock_Agent_AgentKnowledgeBase struct {
	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentknowledgebase.html#cfn-bedrock-agent-agentknowledgebase-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// KnowledgeBaseId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentknowledgebase.html#cfn-bedrock-agent-agentknowledgebase-knowledgebaseid
	KnowledgeBaseId cfz.Expression[string] `json:"KnowledgeBaseId,omitempty"`

	// KnowledgeBaseState is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentknowledgebase.html#cfn-bedrock-agent-agentknowledgebase-knowledgebasestate
	KnowledgeBaseState cfz.Expression[string] `json:"KnowledgeBaseState,omitempty"`
}

// New__AWS_Bedrock_Agent_AgentKnowledgeBase initializes a new AWS_Bedrock_Agent_AgentKnowledgeBase.
func New__AWS_Bedrock_Agent_AgentKnowledgeBase() AWS_Bedrock_Agent_AgentKnowledgeBase {
	return AWS_Bedrock_Agent_AgentKnowledgeBase{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Agent_AgentKnowledgeBase) GetType() string {
	return AWS_Bedrock_Agent_AgentKnowledgeBase__Type
}

// Set__Description updates property "Description".
func (t AWS_Bedrock_Agent_AgentKnowledgeBase) Set__Description(v cfz.Expression[string]) AWS_Bedrock_Agent_AgentKnowledgeBase {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_Bedrock_Agent_AgentKnowledgeBase) SetV__Description(v string) AWS_Bedrock_Agent_AgentKnowledgeBase {
	t.Description = cfz.V(v)
	return t
}

// Set__KnowledgeBaseId updates property "KnowledgeBaseId".
func (t AWS_Bedrock_Agent_AgentKnowledgeBase) Set__KnowledgeBaseId(v cfz.Expression[string]) AWS_Bedrock_Agent_AgentKnowledgeBase {
	t.KnowledgeBaseId = v
	return t
}

// SetV__KnowledgeBaseId updates property "KnowledgeBaseId".
func (t AWS_Bedrock_Agent_AgentKnowledgeBase) SetV__KnowledgeBaseId(v string) AWS_Bedrock_Agent_AgentKnowledgeBase {
	t.KnowledgeBaseId = cfz.V(v)
	return t
}

// Set__KnowledgeBaseState updates property "KnowledgeBaseState".
func (t AWS_Bedrock_Agent_AgentKnowledgeBase) Set__KnowledgeBaseState(v cfz.Expression[string]) AWS_Bedrock_Agent_AgentKnowledgeBase {
	t.KnowledgeBaseState = v
	return t
}

// SetV__KnowledgeBaseState updates property "KnowledgeBaseState".
func (t AWS_Bedrock_Agent_AgentKnowledgeBase) SetV__KnowledgeBaseState(v string) AWS_Bedrock_Agent_AgentKnowledgeBase {
	t.KnowledgeBaseState = cfz.V(v)
	return t
}
