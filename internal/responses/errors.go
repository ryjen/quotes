package responses

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func ServerError(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       err.Error(),
	}
}

func InvalidError() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       "invalid input",
	}
}

func SuccessJSON(obj interface{}) events.APIGatewayProxyResponse {
	response, err := json.Marshal(obj)

	if err != nil {
		return ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}
}

func Success() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func SuccessID(id string) events.APIGatewayProxyResponse {
	return SuccessJSON(map[string]string{
		"id": id,
	})
}
