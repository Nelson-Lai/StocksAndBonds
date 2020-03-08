package getstate

import "github.com/aws/aws-lambda-go/events"

// GetHandler handles get requests for game state
func GetHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "nothing yet",
	}, nil
}
