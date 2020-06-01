package creategame

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type GameCreator struct {
	dynamoClient dynamodb.DynamoDB
}

func (gc GameCreator) CreateGame(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	gameName := request.Body
	gameExists, err := gc.checkGameExistence(gameName)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	if gameExists == true {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("Game of name %s already exists", gameName)
	}

	gameKey := "Game"
	GameTag := dynamodb.Tag{
		Key:   &gameKey,
		Value: &gameName,
	}

	gc.dynamoClient.CreateTable(&dynamodb.CreateTableInput{
		TableName: &gameName,
		Tags:      []*dynamodb.Tag{&GameTag},
	})

	return events.APIGatewayProxyResponse{
		Body: "Game is creating",
	}, nil
}

func (gc GameCreator) checkGameExistence(gamename string) (bool, error) {
	tables, err := gc.dynamoClient.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return false, err
	}
	for _, tableName := range tables.TableNames {
		if gamename == *tableName {
			return true, nil
		}
	}
	return false, nil

}
