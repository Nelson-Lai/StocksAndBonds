package getgamelist

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamoTable = "StocksAndBonds"

// GameListGetter gets game lists
type GameListGetter struct {
	Client   *dynamodb.DynamoDB
	Gamelist []string
}

// GetGameList returns a list of games that are available in the dynamo table
func (g GameListGetter) GetGameList() error {
	games, err := g.Client.Scan(&dynamodb.ScanInput{TableName: &dynamoTable})
	if err != nil {
		return err
	}

	for _, game := range games.Items {
		var gameName string
		dynamodbattribute.Unmarshal(game["GameName"], &gameName)
		g.Gamelist = append(g.Gamelist, gameName)
	}
	return nil
}
