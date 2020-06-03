package updatestate

import (
	"StocksAndBonds/backend/lambda/creategame"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func UpdateState(request events.APIGatewayProxyRequest, dynamo *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {

	gameCreator := creategame.GameCreator{
		DynamoClient: *dynamo,
	}

	return gameCreator.CreateGame(request)
}
