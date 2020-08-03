package getgamelist

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamoTable = "StocksAndBonds"

// GameListGetter gets game lists
type GameListGetter struct {
	Client *dynamodb.DynamoDB
}

// GetGameList returns a list of games that are available in the dynamo table
func (g GameListGetter) GetGameList() ([]string, error) {
	games, err := g.Client.Scan(&dynamodb.ScanInput{TableName: &dynamoTable})
	if err != nil {
		return []string{}, err
	}

	var gamelist []string

	for _, game := range games.Items {
		var gameName string
		dynamodbattribute.Unmarshal(game["GameName"], &gameName)
		gamelist = append(gamelist, gameName)
	}

	return gamelist, nil
}
