package updatestate

import (
	"StocksAndBonds/backend/lambda/creategame"
	"StocksAndBonds/backend/lambda/game"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type updateRequest struct {
	RequestType string    `json:"requestType"`
	GameName    string    `json:"gameName"`
	GameState   game.Game `json:"gameState"`
}

func UpdateState(request events.APIGatewayProxyRequest, dynamo *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {

	requestStruct := updateRequest{}

	var requestBody = []byte(request.Body)
	if request.IsBase64Encoded {
		decode, _ := base64.StdEncoding.DecodeString(request.Body)
		requestBody = decode
	}

	json.Unmarshal(requestBody, &requestStruct)

	gameCreator := creategame.GameCreator{
		DynamoClient: *dynamo,
	}

	if requestStruct.RequestType == "creategame" {

		return gameCreator.CreateGame(request)
	}
	return events.APIGatewayProxyResponse{
		Body:       "Unknown request type",
		StatusCode: http.StatusBadRequest,
	}, nil
}
