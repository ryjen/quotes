package storage

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (db *DynamoDB) ListQuotes(ctx context.Context) ([]*Quote, error) {
	data, err := db.ScanQuotes()

	if err != nil {
		return nil, err
	}

	var quotes []*Quote

	err = dynamodbattribute.UnmarshalListOfMaps(data.Items, &quotes)
	if err != nil {
		return quotes, err
	}

	return quotes, nil
}

func (db *DynamoDB) SaveQuote(ctx context.Context, baseQuote BaseQuote) (string, error) {

	quote := &Quote{baseQuote, newId()}

	data, err := db.NewItem(quote)

	if err != nil {
		return "", err
	}

	_, err = db.PutItem(data)

	if err != nil {
		return "", err
	}

	return quote.ID, nil
}

func (db *DynamoDB) GetQuote(ctx context.Context, id string) (*Quote, error) {
	item, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: &id,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if item.Item == nil {
		return nil, errors.New("could not find item with id")
	}

	var quote *Quote

	err = dynamodbattribute.UnmarshalMap(item.Item, &quote)

	return quote, err
}
