package getstate

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetState(ctx context.Context, request events.APIGatewayProxyRequest, dynamoClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
