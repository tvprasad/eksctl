package cloudformation

import (
	"encoding/json"
)

// AWSDynamoDBTable_KeySchema AWS CloudFormation Resource (AWS::DynamoDB::Table.KeySchema)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html
type AWSDynamoDBTable_KeySchema struct {

	// AttributeName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-attributename
	AttributeName *Value `json:"AttributeName,omitempty"`

	// KeyType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-keytype
	KeyType *Value `json:"KeyType,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_KeySchema) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.KeySchema"
}

func (r *AWSDynamoDBTable_KeySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
