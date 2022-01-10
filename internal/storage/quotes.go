package storage

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "quotes"

func (db *DynamoDB) ListQuotes(ctx context.Context) ([]*Quote, error) {
	data, err := db.ScanItems(tableName)

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

func (db *DynamoDB) AddQuote(ctx context.Context, quote NewQuote) (string, error) {

	var id string

	if quote.ID == nil {
		id = newId()
		quote.ID = &id
	} else {
		id = *quote.ID
	}

	data, err := db.NewPutItem(tableName, quote)

	if err != nil {
		return "", err
	}

	_, err = db.PutItem(data)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (db *DynamoDB) SaveQuote(ctx context.Context, quote *Quote) error {

	data, err := db.NewUpdateItem(tableName, quote.ID, quote)

	if err != nil {
		return err
	}

	_, err = db.UpdateItem(data)

	if err != nil {
		return err
	}

	return nil
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

func (db *DynamoDB) RemoveQuote(ctx context.Context, id string) error {

	item, err := db.NewDeleteItem(tableName, id)

	if err != nil {
		return err
	}

	_, err = db.DeleteItem(item)

	return err
}
