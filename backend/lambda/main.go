package main

import (
	"fmt"

	"StocksAndBonds/backend/lambda/getstate"
	"StocksAndBonds/backend/lambda/updatestate"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	dynamoClient := dynamodb.New(sess)
	fmt.Println(dynamoClient.ListTables(&dynamodb.ListTablesInput{})) // This works if you have aws configured your keys correctly

	lambda.Start(Mux)
}

// Mux handles requests from API Gateway
func Mux(request events.APIGatewayProxyRequest, dynamoClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" {
		return getstate.GetState(request, dynamoClient)
	}
	if request.HTTPMethod == "GET" {
		return updatestate.UpdateState(request, dynamoClient)
	}
	return events.APIGatewayProxyResponse{}, nil
}
