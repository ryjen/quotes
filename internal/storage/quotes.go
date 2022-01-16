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

func (db *DynamoDB) PageQuotes(ctx context.Context, key string, count int) (*PagedQuotes, error) {
	data, err := db.PageItems(tableName, key, count)

	if err != nil {
		return nil, err
	}

	lastKey, hasLastKey := data.LastEvaluatedKey["id"]

	var lastID *string

	if hasLastKey {
		lastID = lastKey.S
	}

	var quotes []*Quote

	err = dynamodbattribute.UnmarshalListOfMaps(data.Items, &quotes)

	return &PagedQuotes{quotes, lastID}, err

}

func (db *DynamoDB) AddQuote(ctx context.Context, quote QuoteWithOptionalID) (string, error) {

	id := quote.GetOrGenerateID()

	item := &Quote{quote.QuoteWithoutID, id}

	data, err := db.NewPutItem(tableName, item)

	if err != nil {
		return "", err
	}

	_, err = db.PutItem(data)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (db *DynamoDB) AddQuotes(ctx context.Context, quotes []QuoteWithOptionalID) ([]string, error) {

	var items []*Quote
	var ids []string

	for _, q := range quotes {
		item := &Quote{q.QuoteWithoutID, q.GetOrGenerateID()}
		items = append(items, item)
		ids = append(ids, item.ID)
	}

	data, err := db.NewBatchPutItem(tableName, items)

	if err != nil {
		return nil, err
	}

	_, err = db.BatchWriteItem(data)

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (db *DynamoDB) SaveQuote(ctx context.Context, quote *Quote) error {

	data, err := db.NewUpdateItem(tableName, quote.ID, quote.QuoteWithoutID)

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
