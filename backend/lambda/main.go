package main

import (
	"net/http"

	"StocksAndBonds/backend/lambda/getstate"
	"StocksAndBonds/backend/lambda/updatestate"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	lambda.Start(Mux)
}

// Mux handles requests from API Gateway
func Mux(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoClient := dynamodb.New(sess)
	switch request.HTTPMethod {
	case "GET":
		return getstate.GetState(request, dynamoClient)
	case "POST":
		return updatestate.UpdateState(request, dynamoClient)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusCreated,
			Body:       "wut",
		}, nil
	}

}
