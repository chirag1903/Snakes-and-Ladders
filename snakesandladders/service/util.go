package service

import (
	"Users/chirag.soni/druva/snakesandladders/initconfig"
	"Users/chirag.soni/druva/snakesandladders/service/dice"
	"Users/chirag.soni/druva/snakesandladders/service/objects"
)

// Services ...
type Services struct {
	DiceInfo     *dice.DiceInfo
	BoardObjects *objects.BoardObjectsInfo
}

// InitService ...
func InitService(config *initconfig.Config) *Services {
	totalBoardSize := config.Board.Size * config.Board.Size
	diceInfo := dice.NewDiceRoll(config.MovementStrategy, config.DiceCount)
	boardObjects := objects.NewBoardObjects(config.Snakes, config.Ladders, config.Players, config.Crocodiles, config.Mines, totalBoardSize)

	return &Services{
		DiceInfo:     diceInfo,
		BoardObjects: boardObjects,
	}

}
