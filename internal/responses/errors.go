package responses

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func FailureJSON(code int, message string) events.APIGatewayProxyResponse {

	obj := map[string]string{
		"message": message,
	}

	response, err := json.Marshal(obj)

	if err != nil {
		return ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}
}

func ServerError(err error) events.APIGatewayProxyResponse {
	return FailureJSON(500, err.Error())
}

func InvalidRequest() events.APIGatewayProxyResponse {
	return FailureJSON(400, "invalid request")
}

func InvalidEntity() events.APIGatewayProxyResponse {
	return FailureJSON(422, "invalid input entity")
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
