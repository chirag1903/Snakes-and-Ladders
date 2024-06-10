package dice

import (
	"math/rand"

	"Users/chirag.soni/druva/snakesandladders/stats"

	"github.com/pkg/errors"
)

// DiceInfo ...
type DiceInfo struct {
	movementStrategy string
	diceCount        int
	diceStats        *stats.DiceStats
}

// NewDiceRoll ...
func NewDiceRoll(ms string, diceCount int) *DiceInfo {
	diceStats := &stats.DiceStats{
		MovementStrategy: ms,
		RollNumbers:      make([]int, 0),
	}
	return &DiceInfo{
		movementStrategy: ms,
		diceCount:        diceCount,
		diceStats:        diceStats,
	}
}

// DiceRoll interface implemented
type DiceRoll interface {
	GetMovementCount(diceValues []int) (int, error)
	GetStats() *stats.DiceStats
}

// GetDiceMovementCount returns the positon to be incremented for a dice roll
func (dInfo *DiceInfo) GetDiceMovementCount() (DiceRoll, error) {
	var diceObj DiceRoll
	switch dInfo.movementStrategy {
	case "SUM":
		diceObj = NewMSSumObj(dInfo.diceCount, dInfo.diceStats)
	case "MAX":
		diceObj = NewMSMaxObj(dInfo.diceCount, dInfo.diceStats)
	case "MIN":
		diceObj = NewMSMinObj(dInfo.diceCount, dInfo.diceStats)
	default:
		return nil, errors.Errorf("invalid movement strategy %s", dInfo.movementStrategy)
	}

	return diceObj, nil
}

// RollDice ...
func (dInfo *DiceInfo) RollDice() int {
	randNum := 1 + rand.Intn(6)
	dInfo.diceStats.RollNumbers = append(dInfo.diceStats.RollNumbers, randNum)
	return randNum
}

// GetStats ...
func (dInfo *DiceInfo) GetStats() *stats.DiceStats {
	return dInfo.diceStats
}
