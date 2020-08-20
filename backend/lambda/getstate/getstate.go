package getstate

import (
	"StocksAndBonds/backend/lambda/getgamelist"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetGamelist should be the only response to a "GET" request and returns the active game list. It is the first thing the client does.
func GetGamelist(request events.APIGatewayProxyRequest, dynamoClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	gameFetcher := getgamelist.GameListGetter{
		Client: dynamoClient,
	}

	gamelist, err := gameFetcher.GetGameList()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 420}, err
	}

	bodyJSON, _ := json.Marshal(gamelist)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(bodyJSON),
	}, nil
}
