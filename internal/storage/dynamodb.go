package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const awsRegion = "ca-central-1"
const tableName = "quotes"

type DynamoDB struct {
	*dynamodb.DynamoDB
}

func NewDynamoDB() *DynamoDB {

	sess := session.Must(session.NewSession())

	cfg := aws.NewConfig().WithRegion(awsRegion)

	return &DynamoDB{dynamodb.New(sess, cfg)}
}

// Scan retrieves the table items as json-like format from a DynamoDB database
func (db *DynamoDB) ScanQuotes() (*dynamodb.ScanOutput, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := db.Scan(input)

	return result, err
}

// NewItem creates a new item as json-like to be stored in DynamoDB database
func (db *DynamoDB) NewItem(in interface{}) (*dynamodb.PutItemInput, error) {
	item, err := dynamodbattribute.MarshalMap(in)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	return input, nil
}
