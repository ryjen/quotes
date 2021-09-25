package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"micrantha.com/quotes/internal/responses"
	"micrantha.com/quotes/internal/storage"
)

var db storage.Storage

func readQuotes(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	quotes, err := db.ListQuotes(ctx)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.SuccessJSON(quotes), nil
}

func main() {
	db = storage.NewDynamoDB()
	lambda.Start(readQuotes)
}
