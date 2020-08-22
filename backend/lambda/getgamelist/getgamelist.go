package getgamelist

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	gameTypes "StocksAndBonds/backend/lambda/game"
)

var dynamoTable = "StocksAndBonds"

// GameListGetter gets game lists
type GameListGetter struct {
	Client *dynamodb.DynamoDB
}

// GetGameList returns a list of games that are available in the dynamo table
func (g GameListGetter) GetGameList() ([]gameTypes.Game, error) {
	games, err := g.Client.Scan(&dynamodb.ScanInput{TableName: &dynamoTable})
	if err != nil {
		return []gameTypes.Game{}, err
	}

	var gamelist []gameTypes.Game

	for _, game := range games.Items {
		var gameName string
		var gameState gameTypes.GameState
		var day int

		dynamodbattribute.Unmarshal(game["GameName"], &gameName)
		dynamodbattribute.Unmarshal(game["Day"], &day)
		dynamodbattribute.Unmarshal(game["Gamestate"], &gameState)

		gameOut := gameTypes.Game{
			GameName:  gameName,
			Day:       day,
			Gamestate: gameState,
		}

		gamelist = append(gamelist, gameOut)
	}

	return gamelist, nil
}
