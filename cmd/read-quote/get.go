package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"micrantha.com/quotes/internal/responses"
)

func getQuote(ctx context.Context, id string) (events.APIGatewayProxyResponse, error) {
	quote, err := db.GetQuote(ctx, id)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.SuccessJSON(quote), nil
}
