package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Theres nothing here yet!
	fmt.Println("lol")
	lambda.Start(Handler)
}

// Handler handles requests from API Gateway stop bitching at me
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO HANDLER
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Kraz sux",
	}, nil
}
