package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.QueryStringParameters["name"]
	fmt.Printf(">> %+v\n", request)
	if name == "" {
		return events.APIGatewayProxyResponse{Body: "Please provide a name in the 'name' query string parameter", StatusCode: 400}, nil
	}
	return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Hello, %s!", name), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
