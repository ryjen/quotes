package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"micrantha.com/quotes/internal/responses"
	"micrantha.com/quotes/internal/storage"
)

var db storage.Storage

func createQuote(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var quote storage.BaseQuote
	err := json.Unmarshal([]byte(e.Body), &quote)

	if err != nil {
		return responses.InvalidError(), err
	}

	id, err := db.SaveQuote(context.Background(), quote)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.SuccessID(id), err
}

func main() {
	db = storage.NewMemoryDB()
	lambda.Start(createQuote)
}
