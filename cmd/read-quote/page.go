package main

import (
	"context"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"micrantha.com/quotes/internal/responses"
)

func pageQuotes(ctx context.Context, key string, pageSize string) (events.APIGatewayProxyResponse, error) {

	count := 20

	if len(pageSize) > 0 {
		count, _ = strconv.Atoi(pageSize)
	}

	quotes, err := db.PageQuotes(ctx, key, count)

	if err != nil {
		return responses.ServerError(err), err
	}

	return responses.SuccessJSON(quotes), nil
}
