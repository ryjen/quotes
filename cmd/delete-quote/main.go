package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"micrantha.com/quotes/internal/responses"
	"micrantha.com/quotes/internal/storage"
)

var db storage.Storage

func deleteQuote(ctx context.Context, id string) (events.APIGatewayProxyResponse, error) {

	err := db.RemoveQuote(ctx, id)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.Success(), nil
}

func main() {
	db = storage.NewDynamoDB()
	lambda.Start(deleteQuote)
}
