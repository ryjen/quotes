package storage

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const awsRegion = "ca-central-1"

type DynamoDB struct {
	*dynamodb.DynamoDB
}

func NewDynamoDB() *DynamoDB {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cfg := aws.NewConfig().WithRegion(awsRegion)

	return &DynamoDB{dynamodb.New(sess, cfg)}
}

// Scan retrieves the table items as json-like format from a DynamoDB database
func (db *DynamoDB) ScanItems(tableName string) (*dynamodb.ScanOutput, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := db.Scan(input)

	return result, err
}

func (db *DynamoDB) PageItems(tableName string, key string, count int) (*dynamodb.ScanOutput, error) {
	input := &dynamodb.ScanInput{
		ExclusiveStartKey: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(key)},
		},
		Limit:     aws.Int64(int64(count)),
		TableName: aws.String(tableName),
	}
	result, err := db.Scan(input)

	return result, err
}

// NewQuote creates a new item as json-like to be stored in DynamoDB database
func (db *DynamoDB) NewPutItem(tableName string, in interface{}) (*dynamodb.PutItemInput, error) {
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

func (db *DynamoDB) NewBatchPutItem(tableName string, in interface{}) (*dynamodb.BatchWriteItemInput, error) {
	items, err := dynamodbattribute.MarshalList(in)

	if err != nil {
		return nil, err
	}

	request := make([]*dynamodb.WriteRequest, len(items))

	for i, item := range items {

		request[i] = &dynamodb.WriteRequest{
			PutRequest:    &dynamodb.PutRequest{Item: item.M},
			DeleteRequest: nil,
		}
	}

	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			tableName: request,
		},
	}

	return input, nil
}

// NewQuote creates a new item as json-like to be stored in DynamoDB database
func (db *DynamoDB) NewUpdateItem(tableName string, id string, in interface{}) (*dynamodb.UpdateItemInput, error) {
	item, err := dynamodbattribute.MarshalMap(in)
	if err != nil {
		return nil, err
	}

	var expr []string
	names := make(map[string]*string)

	for key := range item {
		valueKey := fmt.Sprintf(":%s", key)
		nameKey := fmt.Sprintf("#%s", key)
		names[nameKey] = aws.String(key)
		expr = append(expr, fmt.Sprintf("%s = %s", nameKey, valueKey))
		attr := item[key]
		delete(item, key)
		item[valueKey] = attr
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: item,
		ExpressionAttributeNames:  names,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
		TableName:        aws.String(tableName),
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String(fmt.Sprintf("SET %s", strings.Join(expr, ", "))),
	}

	return input, nil
}

func (db *DynamoDB) NewDeleteItem(tableName string, id string) (*dynamodb.DeleteItemInput, error) {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
		TableName: aws.String(tableName),
	}

	return input, nil
}
