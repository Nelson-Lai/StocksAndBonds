package getgamelist

import (
	"StocksAndBonds/rpc/api/gamelist"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamoTable = "StocksAndBonds"

type GameListGetter struct {
	Client   *dynamodb.DynamoDB
	Gamelist []string
}

// GetGameList returns a list of games that are available in the dynamo table
func (g GameListGetter) GetGameList() (gamelist.GameListResponse, error) {
	games, err := g.Client.Scan(&dynamodb.ScanInput{TableName: &dynamoTable})
	if err != nil {
		return gamelist.GameListResponse{}, err
	}

	gameList := gamelist.GameListResponse{Games: []string{}}

	for _, game := range games.Items {
		var gameName string
		dynamodbattribute.Unmarshal(game["GameName"], &gameName)
		gameList.Games = append(gameList.Games, gameName)
	}
	return gameList, nil
}
