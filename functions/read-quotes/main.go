package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"micrantha.com/quotes/internal/responses"
	"micrantha.com/quotes/internal/storage"
)

var db storage.Storage

func readQuotes(ctx context.Context, e events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {

	quotes, err := db.ListQuotes(context.Background())

	if err != nil {
		return responses.ServerError(err)
	}

	return responses.SuccessJSON(quotes)
}

func main() {
	db = storage.NewDynamoDB()
	lambda.Start(readQuotes)
}
