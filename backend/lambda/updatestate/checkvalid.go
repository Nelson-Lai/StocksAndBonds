package updatestate

import (
	"StocksAndBonds/backend/lambda/game"
	"fmt"
)

// This file contains all the logic necessary to prevent players from
// taking bad turns or making illegal moves

func isNewStateValid(playerPortolio game.Portfolio) (bool, error) {
	if playerPortolio.Cash < 0 {
		return false, nil
	}

	for _, company := range game.CompanyList {
		numShares, found := playerPortolio.Portfolio[company]
		if !found {
			return false, fmt.Errorf("Could not find company %s in map", company)
		}
		if numShares[len(numShares)-1] < 0 {
			return false, nil
		}
	}
	return true, nil
}
