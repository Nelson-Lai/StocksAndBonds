package updatestate

import (
	"StocksAndBonds/backend/lambda/game"
	"StocksAndBonds/backend/lambda/updatestate/creategame"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func UpdateState(request events.APIGatewayProxyRequest, dynamo *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {

	requestStruct := game.UpdateRequest{}

	var requestBody = []byte(request.Body)
	if request.IsBase64Encoded {
		decoded, _ := base64.StdEncoding.DecodeString(request.Body)
		requestBody = decoded
	}

	json.Unmarshal(requestBody, &requestStruct)

	gameCreator := creategame.GameCreator{
		DynamoClient: *dynamo,
	}

	if requestStruct.RequestType == "creategame" {

		return gameCreator.CreateGame(requestStruct)
	}
	return events.APIGatewayProxyResponse{
		Body:       "Unknown request type",
		StatusCode: http.StatusBadRequest,
	}, nil
}
