# go-lambda

Demo Project for Go using AWS Lambda

https://www.softkraft.co/aws-lambda-in-golang/

## How to Deploy AWS Lambda on AWS

1. Create Lambda function:
   1. In **Console**, search for **Lambda**.
   2. In Lambda page, choose **Create function**.
   3. In next page:
      1. Function name: *helloLambda*, for example.
      2. Runtime: your language, for example: *Go*.
      3. Execution role: Create a new role from aws policy templates
      4. Role Name: *helloLambda-executor* - and choose Simple microservice permission
2. Creat API:
   1. In **Console** search for **API Gateway**
   2. In section of **REST Api**, click **Build** button.
   3. In the next page, choose:
      1. Protocol: *REST*
      2. Create new API: *New API*
      3. API name: *helloLambdaAPI*
   4. After click **Create API** button, on the next screen:
      1. Click select box **Action** and select **Create API** option.
      2. Select:
         1. Integration type: *Lambda Function*
         2. Use Lambda Proxy integration: check
         3. Lambda Function: *helloLambda*
         4. Use Default Timeout: check

## With [[Go]]

Demo project: https://github.com/ducminhgd/go-lambda

Simple steps:
1. Make sure that you install the dependencies `go get -v all`.
2. Build project `GOOS=linux go build -o build/main cmd/main.go`.
3. Zip binary to upload to AWS `zip -jrm build/main.zip build/main`
4. If you want to handle request and response in your code, then please make sure that the field **Use Lambda Proxy integration** in **Integration Request** checked.

Sample code

```go
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

```