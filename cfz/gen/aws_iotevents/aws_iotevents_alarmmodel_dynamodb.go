// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotevents

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTEvents_AlarmModel_DynamoDB__Type is the CloudFormation type for AWS::IoTEvents::AlarmModel.DynamoDB.
	AWS_IoTEvents_AlarmModel_DynamoDB__Type = "AWS::IoTEvents::AlarmModel.DynamoDB"
)

var (
	// AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap reports all the CloudFormation properties for AWS::IoTEvents::AlarmModel.DynamoDB.
	AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap = struct {
		HashKeyField  string
		HashKeyType   string
		HashKeyValue  string
		Operation     string
		Payload       string
		PayloadField  string
		RangeKeyField string
		RangeKeyType  string
		RangeKeyValue string
		TableName     string
	}{
		HashKeyField:  "HashKeyField",
		HashKeyType:   "HashKeyType",
		HashKeyValue:  "HashKeyValue",
		Operation:     "Operation",
		Payload:       "Payload",
		PayloadField:  "PayloadField",
		RangeKeyField: "RangeKeyField",
		RangeKeyType:  "RangeKeyType",
		RangeKeyValue: "RangeKeyValue",
		TableName:     "TableName",
	}

	// AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesSlice reports all the CloudFormation properties for AWS::IoTEvents::AlarmModel.DynamoDB.
	AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesSlice = []string{
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.HashKeyField,
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.HashKeyType,
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.HashKeyValue,
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.Operation,
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.Payload,
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.PayloadField,
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.RangeKeyField,
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.RangeKeyType,
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.RangeKeyValue,
		AWS_IoTEvents_AlarmModel_DynamoDB__PropertiesMap.TableName,
	}
)

// AWS_IoTEvents_AlarmModel_DynamoDB is a binding for AWS::IoTEvents::AlarmModel.DynamoDB.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html
type AWS_IoTEvents_AlarmModel_DynamoDB struct {
	// HashKeyField is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-hashkeyfield
	HashKeyField cfz.Expression[string] `json:"HashKeyField,omitempty"`

	// HashKeyType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-hashkeytype
	HashKeyType cfz.Expression[string] `json:"HashKeyType,omitempty"`

	// HashKeyValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-hashkeyvalue
	HashKeyValue cfz.Expression[string] `json:"HashKeyValue,omitempty"`

	// Operation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-operation
	Operation cfz.Expression[string] `json:"Operation,omitempty"`

	// Payload is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-payload
	Payload cfz.Expression[AWS_IoTEvents_AlarmModel_Payload] `json:"Payload,omitempty"`

	// PayloadField is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-payloadfield
	PayloadField cfz.Expression[string] `json:"PayloadField,omitempty"`

	// RangeKeyField is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-rangekeyfield
	RangeKeyField cfz.Expression[string] `json:"RangeKeyField,omitempty"`

	// RangeKeyType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-rangekeytype
	RangeKeyType cfz.Expression[string] `json:"RangeKeyType,omitempty"`

	// RangeKeyValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-rangekeyvalue
	RangeKeyValue cfz.Expression[string] `json:"RangeKeyValue,omitempty"`

	// TableName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-dynamodb.html#cfn-iotevents-alarmmodel-dynamodb-tablename
	TableName cfz.Expression[string] `json:"TableName,omitempty"`
}

// New__AWS_IoTEvents_AlarmModel_DynamoDB initializes a new AWS_IoTEvents_AlarmModel_DynamoDB.
func New__AWS_IoTEvents_AlarmModel_DynamoDB() AWS_IoTEvents_AlarmModel_DynamoDB {
	return AWS_IoTEvents_AlarmModel_DynamoDB{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTEvents_AlarmModel_DynamoDB) GetType() string {
	return AWS_IoTEvents_AlarmModel_DynamoDB__Type
}

// Set__HashKeyField updates property "HashKeyField".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__HashKeyField(v cfz.Expression[string]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.HashKeyField = v
	return t
}

// SetV__HashKeyField updates property "HashKeyField".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__HashKeyField(v string) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.HashKeyField = cfz.V(v)
	return t
}

// Set__HashKeyType updates property "HashKeyType".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__HashKeyType(v cfz.Expression[string]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.HashKeyType = v
	return t
}

// SetV__HashKeyType updates property "HashKeyType".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__HashKeyType(v string) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.HashKeyType = cfz.V(v)
	return t
}

// Set__HashKeyValue updates property "HashKeyValue".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__HashKeyValue(v cfz.Expression[string]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.HashKeyValue = v
	return t
}

// SetV__HashKeyValue updates property "HashKeyValue".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__HashKeyValue(v string) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.HashKeyValue = cfz.V(v)
	return t
}

// Set__Operation updates property "Operation".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__Operation(v cfz.Expression[string]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.Operation = v
	return t
}

// SetV__Operation updates property "Operation".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__Operation(v string) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.Operation = cfz.V(v)
	return t
}

// Set__Payload updates property "Payload".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__Payload(v cfz.Expression[AWS_IoTEvents_AlarmModel_Payload]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.Payload = v
	return t
}

// SetV__Payload updates property "Payload".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__Payload(v AWS_IoTEvents_AlarmModel_Payload) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.Payload = cfz.V(v)
	return t
}

// Set__PayloadField updates property "PayloadField".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__PayloadField(v cfz.Expression[string]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.PayloadField = v
	return t
}

// SetV__PayloadField updates property "PayloadField".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__PayloadField(v string) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.PayloadField = cfz.V(v)
	return t
}

// Set__RangeKeyField updates property "RangeKeyField".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__RangeKeyField(v cfz.Expression[string]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.RangeKeyField = v
	return t
}

// SetV__RangeKeyField updates property "RangeKeyField".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__RangeKeyField(v string) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.RangeKeyField = cfz.V(v)
	return t
}

// Set__RangeKeyType updates property "RangeKeyType".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__RangeKeyType(v cfz.Expression[string]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.RangeKeyType = v
	return t
}

// SetV__RangeKeyType updates property "RangeKeyType".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__RangeKeyType(v string) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.RangeKeyType = cfz.V(v)
	return t
}

// Set__RangeKeyValue updates property "RangeKeyValue".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__RangeKeyValue(v cfz.Expression[string]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.RangeKeyValue = v
	return t
}

// SetV__RangeKeyValue updates property "RangeKeyValue".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__RangeKeyValue(v string) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.RangeKeyValue = cfz.V(v)
	return t
}

// Set__TableName updates property "TableName".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) Set__TableName(v cfz.Expression[string]) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.TableName = v
	return t
}

// SetV__TableName updates property "TableName".
func (t AWS_IoTEvents_AlarmModel_DynamoDB) SetV__TableName(v string) AWS_IoTEvents_AlarmModel_DynamoDB {
	t.TableName = cfz.V(v)
	return t
}
