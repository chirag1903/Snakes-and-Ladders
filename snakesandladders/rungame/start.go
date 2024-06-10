package rungame

import (
	"Users/chirag.soni/druva/snakesandladders/initconfig"
	"Users/chirag.soni/druva/snakesandladders/service"
	"fmt"

	"github.com/pkg/errors"
)

var totalTurnLimit = 10000

// StartGame ...
func StartGame(config *initconfig.Config, svcs *service.Services, manualInputInfoMap []*initconfig.ManualInputInfo) error {
	totalBoardSize := config.Board.Size * config.Board.Size
	totalTurns := 0
	totalManualInputCount := len(manualInputInfoMap)

	for {
		if totalTurns == totalTurnLimit {
			return errors.Errorf("total turn limit exceeded")

		} else if totalManualInputCount != 0 && totalManualInputCount == totalTurns {
			return errors.Errorf("total manual input exhausted")
		}

		for _, player := range config.Players {
			diceVals := make([]int, 0)

			if totalManualInputCount != 0 {
				if manualInputInfoMap[totalTurns].PlayerName != player.Name {
					continue
				}
				diceVals = manualInputInfoMap[totalTurns].RollValues
			}

			totalTurns++
			movementCount, err := getMovementCount(config.DiceCount, svcs, diceVals)
			if err != nil {
				return err
			}

			// assuming player will be rolling dice and will wait later
			if player.Wait != 0 {
				player.Wait--
				continue
			}

			newPos := player.Position + movementCount
			if newPos > totalBoardSize {
				continue
			}

			newPos, wait, err := getPositonIfObjectsPresent(newPos, player.Name, svcs)
			if err != nil {
				return err
			}

			player.Position = newPos
			player.Wait = wait
			if player.Position == totalBoardSize {
				fmt.Printf("Player %s Wins in totalturns %d\n", player.Name, totalTurns)
				return nil
			}
		}
	}

}

func getMovementCount(diceCount int, svcs *service.Services, diceVals []int) (int, error) {
	if len(diceVals) == 0 {
		for j := 0; j < diceCount; j++ {
			// roll dice
			diceVal := svcs.DiceInfo.RollDice()
			diceVals = append(diceVals, diceVal)
		}
	}

	// get ms obj
	msObj, err := svcs.DiceInfo.GetDiceMovementCount()
	if err != nil {
		return 0, err
	}

	movementCount, err := msObj.GetMovementCount(diceVals)
	if err != nil {
		return 0, err
	}

	diceStats := msObj.GetStats()
	fmt.Printf("Movement Strategy:%s, TotalMovementCount: %d, dice rolls", diceStats.MovementStrategy, diceStats.TotalMovement)
	fmt.Println(diceStats.RollNumbers)

	// reset dice rolled number
	diceStats.RollNumbers = make([]int, 0)

	return movementCount, nil
}

func getPositonIfObjectsPresent(newPos int, playerName string, svcs *service.Services) (int, int, error) {
	// fetch object at the position
	isPresent, boardObj, origPos, err := svcs.BoardObjects.GetBoardObject(newPos, playerName)
	if err != nil {
		return 0, 0, nil
	}

	updatedPos := newPos
	if isPresent {
		// update position as per object
		updatedPos, err = boardObj.GetNewBoardPosition(newPos, origPos, playerName)
		if err != nil {
			return 0, 0, err
		}

		gameStats := boardObj.LogStats()

		// case where a player can have multiple steps
		if gameStats.RepeatStep {
			updatedPos, _, err = getPositonIfObjectsPresent(updatedPos, playerName, svcs)
			if err != nil {
				return 0, 0, err
			}
		}

		if gameStats.HoldTurn != 0 {
			return updatedPos, gameStats.HoldTurn, nil
		}

	} else {
		playerStatus := svcs.BoardObjects.GetPlayerStats()
		fmt.Printf("%s rolled a %d and moved from %d to %d\n", playerStatus.Name, playerStatus.RollValue, playerStatus.StartPos, playerStatus.EndPos)
	}

	return updatedPos, 0, nil
}
