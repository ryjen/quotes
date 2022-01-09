package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"micrantha.com/quotes/internal/responses"
	"micrantha.com/quotes/internal/storage"
)

var db storage.Storage

func createQuote(ctx context.Context, quote storage.NewQuote) (events.APIGatewayProxyResponse, error) {

	id, err := db.SaveQuote(ctx, quote)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.SuccessID(id), err
}

func main() {
	db = storage.NewDynamoDB()
	lambda.Start(createQuote)
}
