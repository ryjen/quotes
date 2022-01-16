package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"micrantha.com/quotes/internal/responses"
)

func listQuotes(ctx context.Context) (events.APIGatewayProxyResponse, error) {

	quotes, err := db.ListQuotes(ctx)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.SuccessJSON(quotes), nil
}
