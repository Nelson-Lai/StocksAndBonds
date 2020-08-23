package creategame

import (
	"errors"
	"fmt"
	"net/http"

	"StocksAndBonds/backend/lambda/game"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	tableName = "StocksAndBonds"
)

type GameCreator struct {
	DynamoClient dynamodb.DynamoDB
}

func (gc GameCreator) CreateGame(request game.UpdateRequest) (events.APIGatewayProxyResponse, error) {
	gameName := request.Game.GameName

	gameExists, err := gc.checkGameExistence(gameName)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}
	if gameExists == true {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusFound,
			Body:       fmt.Sprintf("Game of name %s already exists", gameName),
		}, fmt.Errorf("Game of name %s already exists", gameName)
	}

	game := game.NewGame()
	game.GameName = gameName
	game.Players = 1
	game.Day = 1
	game.PlayerList = append(game.PlayerList, request.Requester)

	gameMarshall, err := dynamodbattribute.MarshalMap(game)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Marshall error oh no",
		}, err
	}

	gc.DynamoClient.PutItem(&dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      gameMarshall,
	},
	)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Game Created",
	}, nil
}

func (gc GameCreator) checkGameExistence(gamename string) (bool, error) {
	game, err := gc.DynamoClient.GetItem(
		&dynamodb.GetItemInput{
			TableName: &tableName,
			Key: map[string]*dynamodb.AttributeValue{
				"GameName": {
					S: &gamename,
				},
			},
		},
	)
	if err != nil {
		return false, err
	}
	if len(game.Item) == 0 {
		return false, nil
	}
	return true, nil
}

func stringSliceToPointerSlice(input []string) []*string {
	output := []*string{}
	for _, entry := range input {
		output = append(output, &entry)
	}
	return output
}

func (gc GameCreator) joinGame(game game.Game, player string) error {
	if player == "" {
		return errors.New("player name cannot be empty")
	}
	return nil
}
