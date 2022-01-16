package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"micrantha.com/quotes/internal/storage"
)

var db storage.Storage

func createQuote(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	isImport := strings.HasSuffix(e.Path, "/import")

	if isImport {
		return importQuotes(ctx, e)
	}

	return addQuote(ctx, e)
}

func main() {
	db = storage.NewDynamoDB()
	lambda.Start(createQuote)
}
