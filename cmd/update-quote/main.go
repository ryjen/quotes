package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"micrantha.com/quotes/internal/responses"
	"micrantha.com/quotes/internal/storage"
)

var db storage.Storage

func updateQuote(ctx context.Context, quote *storage.Quote) (events.APIGatewayProxyResponse, error) {

	err := db.SaveQuote(ctx, quote)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.Success(), err
}

func main() {
	db = storage.NewDynamoDB()
	lambda.Start(updateQuote)
}
