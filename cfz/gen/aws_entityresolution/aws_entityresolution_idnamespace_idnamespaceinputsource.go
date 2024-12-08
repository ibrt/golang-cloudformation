// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_entityresolution

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EntityResolution_IdNamespace_IdNamespaceInputSource__Type is the CloudFormation type for AWS::EntityResolution::IdNamespace.IdNamespaceInputSource.
	AWS_EntityResolution_IdNamespace_IdNamespaceInputSource__Type = "AWS::EntityResolution::IdNamespace.IdNamespaceInputSource"
)

var (
	// AWS_EntityResolution_IdNamespace_IdNamespaceInputSource__PropertiesMap reports all the CloudFormation properties for AWS::EntityResolution::IdNamespace.IdNamespaceInputSource.
	AWS_EntityResolution_IdNamespace_IdNamespaceInputSource__PropertiesMap = struct {
		InputSourceARN string
		SchemaName     string
	}{
		InputSourceARN: "InputSourceARN",
		SchemaName:     "SchemaName",
	}

	// AWS_EntityResolution_IdNamespace_IdNamespaceInputSource__PropertiesSlice reports all the CloudFormation properties for AWS::EntityResolution::IdNamespace.IdNamespaceInputSource.
	AWS_EntityResolution_IdNamespace_IdNamespaceInputSource__PropertiesSlice = []string{
		AWS_EntityResolution_IdNamespace_IdNamespaceInputSource__PropertiesMap.InputSourceARN,
		AWS_EntityResolution_IdNamespace_IdNamespaceInputSource__PropertiesMap.SchemaName,
	}
)

// AWS_EntityResolution_IdNamespace_IdNamespaceInputSource is a binding for AWS::EntityResolution::IdNamespace.IdNamespaceInputSource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html
type AWS_EntityResolution_IdNamespace_IdNamespaceInputSource struct {
	// InputSourceARN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html#cfn-entityresolution-idnamespace-idnamespaceinputsource-inputsourcearn
	InputSourceARN cfz.Expression[string] `json:"InputSourceARN,omitempty"`

	// SchemaName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html#cfn-entityresolution-idnamespace-idnamespaceinputsource-schemaname
	SchemaName cfz.Expression[string] `json:"SchemaName,omitempty"`
}

// New__AWS_EntityResolution_IdNamespace_IdNamespaceInputSource initializes a new AWS_EntityResolution_IdNamespace_IdNamespaceInputSource.
func New__AWS_EntityResolution_IdNamespace_IdNamespaceInputSource() AWS_EntityResolution_IdNamespace_IdNamespaceInputSource {
	return AWS_EntityResolution_IdNamespace_IdNamespaceInputSource{}
}

// GetType returns the CloudFormation type.
func (AWS_EntityResolution_IdNamespace_IdNamespaceInputSource) GetType() string {
	return AWS_EntityResolution_IdNamespace_IdNamespaceInputSource__Type
}

// Set__InputSourceARN updates property "InputSourceARN".
func (t AWS_EntityResolution_IdNamespace_IdNamespaceInputSource) Set__InputSourceARN(v cfz.Expression[string]) AWS_EntityResolution_IdNamespace_IdNamespaceInputSource {
	t.InputSourceARN = v
	return t
}

// SetV__InputSourceARN updates property "InputSourceARN".
func (t AWS_EntityResolution_IdNamespace_IdNamespaceInputSource) SetV__InputSourceARN(v string) AWS_EntityResolution_IdNamespace_IdNamespaceInputSource {
	t.InputSourceARN = cfz.V(v)
	return t
}

// Set__SchemaName updates property "SchemaName".
func (t AWS_EntityResolution_IdNamespace_IdNamespaceInputSource) Set__SchemaName(v cfz.Expression[string]) AWS_EntityResolution_IdNamespace_IdNamespaceInputSource {
	t.SchemaName = v
	return t
}

// SetV__SchemaName updates property "SchemaName".
func (t AWS_EntityResolution_IdNamespace_IdNamespaceInputSource) SetV__SchemaName(v string) AWS_EntityResolution_IdNamespace_IdNamespaceInputSource {
	t.SchemaName = cfz.V(v)
	return t
}
