package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"micrantha.com/quotes/internal/responses"
	"micrantha.com/quotes/internal/storage"
)

var db storage.Storage

func deleteQuote(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id := e.PathParameters["id"]

	if len(id) == 0 {
		return responses.InvalidRequest(), nil
	}

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
