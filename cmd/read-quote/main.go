package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"micrantha.com/quotes/internal/storage"
)

var db storage.Storage

func readQuote(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id := e.PathParameters["id"]

	if len(id) > 0 {
		return getQuote(ctx, id)
	}

	key, isPaged := e.QueryStringParameters["key"]

	if isPaged {
		return pageQuotes(ctx, key, e.QueryStringParameters["count"])
	}

	return listQuotes(ctx)
}

func main() {
	db = storage.NewDynamoDB()
	lambda.Start(readQuote)
}
