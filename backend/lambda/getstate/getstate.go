package getstate

import (
	"StocksAndBonds/backend/lambda/getgamelist"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetState should be the only response to a "GET" request and returns the active game list
func GetState(request events.APIGatewayProxyRequest, dynamoClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	gameFetcher := getgamelist.GameListGetter{
		Client:   dynamoClient,
		Gamelist: []string{},
	}

	gamelist, err := gameFetcher.GetGameList()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 420}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       strings.Join(gamelist.Games, ","),
	}, nil
}
