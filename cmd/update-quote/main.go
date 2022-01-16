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

func updateQuote(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id := e.PathParameters["id"]

	if len(id) == 0 {
		return responses.InvalidRequest(), nil
	}

	var quote storage.Quote

	err := json.Unmarshal([]byte(e.Body), &quote)

	if err != nil {
		return responses.InvalidEntity(), err
	}

	quote.ID = id

	err = db.SaveQuote(ctx, &quote)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.Success(), err
}

func main() {
	db = storage.NewDynamoDB()
	lambda.Start(updateQuote)
}
