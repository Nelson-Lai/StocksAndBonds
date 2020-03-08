package main

import (
	"fmt"
	"strings"

	"main/getstate"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	dynamo := dynamodb.New(sess)
	fmt.Println(dynamo.ListTables(&dynamodb.ListTablesInput{})) // This works if you have aws configured your keys correctly
}

func main() {
	// Theres nothing here yet!
	fmt.Println("lol from main")
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

// Mux represents the main transmuxxer
func Mux(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" {
		return getstate.
	}
	if request.HTTPMethod == "POST" {
		if strings.Contains(request.Path, "transaction") {
			return events.APIGatewayProxyResponse{
				StatusCode: 200,
			}, nil
		}

	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func transactionHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
