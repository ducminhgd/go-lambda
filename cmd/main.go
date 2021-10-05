package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	name, found := req.QueryStringParameters["name"]
	if !found {
		name = "Anonymous"
	}
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "text/plain"}}
	resp.StatusCode = 200
	resp.Body = string(fmt.Sprintf("Hello, %s! Welcome to AWS Lambda", name))
	return &resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
