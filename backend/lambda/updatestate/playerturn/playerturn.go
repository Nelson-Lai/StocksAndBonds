package playerturn

import (
	"StocksAndBonds/backend/lambda/game"
	"net/http"

	"StocksAndBonds/backend/lambda/locks"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	tableName = "StocksAndBonds"
)

func updatePlayerState(request game.UpdateRequest, dynamoClient dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	player := request.Requester
	newPortfolio := request.Game.Gamestate.PlayerState[player]
	valid, err := IsNewStateValid(newPortfolio)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	if !valid {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Turn invalid",
		}, nil
	}

	err = locks.Lock()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	defer locks.Unlock()

	gameMarshall, err := dynamodbattribute.MarshalMap(request.Game)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Marshall error oh no",
		}, err
	}

	dynamoClient.PutItem(&dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      gameMarshall,
	},
	)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func getGameState() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, nil
}
