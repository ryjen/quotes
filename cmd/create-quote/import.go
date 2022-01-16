package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"micrantha.com/quotes/internal/responses"
	"micrantha.com/quotes/internal/storage"
)

func importQuotes(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	quotes := []storage.QuoteWithOptionalID{}

	err := json.Unmarshal([]byte(e.Body), &quotes)

	if err != nil {
		return responses.InvalidEntity(), err
	}

	ids, err := db.AddQuotes(ctx, quotes)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.SuccessJSON(ids), nil
}
