Notes about gamestate

DynamoDB key should be in format:

GameNAME:gamestate
GameNAME:PlayerName

Everything should take a lock and have a retry built in



curl -d '{"requestType":"creategame", "requester":"memer420", "gameState": {"gameName":"tester555"}}' -H "Content-Type: application/json" -X POST https://eyu6c6iiy3.execute-api.us-east-2.amazonaws.com/development/stocks

curl https://eyu6c6iiy3.execute-api.us-east-2.amazonaws.com/development/stocks