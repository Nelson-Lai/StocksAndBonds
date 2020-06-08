package game

// Game contains all the information needed to play a game of stocsk and bonds
type Game struct {
	GameName   string
	Players    int
	Day        int
	PlayerList []string
	Gamestate  GameState
}

// GameState records historic prices per company
type GameState struct {
	Prices      map[string][]int
	PlayerState map[string]Portfolio
}

// Portfolio represents the assets each player has in each company
type Portfolio struct {
	Portfolio map[string][]int
	Cash      int
}

var CompanyList = []string{
	"Central City",
	"Growth Corporation",
	"Metro Properties",
	"Pioneer Mutual",
	"Shady Brooks",
	"Stryker Drilling",
	"Tri-City Transport",
	"United Auto",
	"Uranium Enterprises",
	"Valley Power",
}

// NewGame returns a game struct with all the boring shit filled in
func NewGame() Game {

	newPrices := make(map[string][]int)
	for _, company := range CompanyList {
		newPrices[company] = []int{100}
	}

	gameState := GameState{
		Prices: newPrices,
	}

	return Game{
		GameName:   "",
		Players:    0,
		PlayerList: []string{},
		Gamestate:  gameState,
	}
}

// NewPortfolio returns a portfolio object with no stocks and 10000 cash
func NewPortfolio() Portfolio {
	newPortfolio := make(map[string][]int)
	for _, company := range CompanyList {
		newPortfolio[company] = []int{0}
	}

	return Portfolio{
		Portfolio: newPortfolio,
		Cash:      10000,
	}
}
