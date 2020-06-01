package getgamelist

import (
	"StocksAndBonds/rpc/api/gamelist"
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type GameListGetter struct {
	client   *dynamodb.DynamoDB
	gamelist []string
}

// GetGameList returns a list of games that are available in the dynamo table
func (g GameListGetter) GetGameList(ctx context.Context, req gamelist.GameListRequest) (gamelist.GameListResponse, error) {
	games, err := g.client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return gamelist.GameListResponse{}, err
	}
	gamesList := make([]string, len(games.TableNames))
	for idx, gamePointer := range games.TableNames {
		gamesList[idx] = *gamePointer
	}
	return gamelist.GameListResponse{
		Games: gamesList,
	}, nil
}
