package dice

import (
	"Users/chirag.soni/druva/snakesandladders/stats"

	"github.com/pkg/errors"
)

// MSSumObj ...
type MSSumObj struct {
	diceCount int
	diceStats *stats.DiceStats
}

// NewMSSumObj ...
func NewMSSumObj(diceCount int, diceStats *stats.DiceStats) DiceRoll {
	return &MSSumObj{
		diceCount: diceCount,
		diceStats: diceStats,
	}
}

// GetMovementCount for sum ms
func (sumObj *MSSumObj) GetMovementCount(diceValues []int) (int, error) {
	if len(diceValues) == 0 {
		return 0, errors.Errorf("dice values should not be")
	}

	if len(diceValues) != sumObj.diceCount {
		return 0, errors.Errorf("invalid dice count and values %d & %d", len(diceValues), sumObj.diceCount)
	}

	totalSum := 0
	for _, val := range diceValues {
		totalSum += val
	}

	sumObj.diceStats.TotalMovement++
	return totalSum, nil
}

// GetStats ...
func (sumObj *MSSumObj) GetStats() *stats.DiceStats {
	return sumObj.diceStats
}
