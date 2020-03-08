package creategame

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateGame(ctx context.Context, request events.APIGatewayProxyRequest, dynamoClient dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	gameName := request.Body
	dynamoClient.
}
