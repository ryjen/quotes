package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"micrantha.com/quotes/internal/responses"
	"micrantha.com/quotes/internal/storage"
)

func addQuote(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var quote storage.QuoteWithOptionalID

	err := json.Unmarshal([]byte(e.Body), &quote)

	if err != nil {
		return responses.InvalidEntity(), nil
	}

	id, err := db.AddQuote(ctx, quote)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.SuccessID(id), err
}
