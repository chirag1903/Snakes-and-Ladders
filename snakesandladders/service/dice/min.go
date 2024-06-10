package dice

import (
	"Users/chirag.soni/druva/snakesandladders/stats"

	"github.com/pkg/errors"
)

// MSMinObj ...
type MSMinObj struct {
	diceCount int
	diceStats *stats.DiceStats
}

// NewMSMinObj ...
func NewMSMinObj(diceCount int, diceStats *stats.DiceStats) DiceRoll {
	return &MSMinObj{
		diceCount: diceCount,
		diceStats: diceStats,
	}
}

// GetMovementCount for sum ms
func (minObj *MSMinObj) GetMovementCount(diceValues []int) (int, error) {
	if len(diceValues) == 0 {
		return 0, errors.Errorf("dice values should not be")
	}

	if len(diceValues) != minObj.diceCount {
		return 0, errors.Errorf("invalid dice count and values %d & %d", len(diceValues), minObj.diceCount)
	}

	minVal := 6
	for _, val := range diceValues {
		if val < minVal {
			minVal = val
		}
	}

	minObj.diceStats.TotalMovement++
	return minVal, nil
}
// GetStats ...
func (minObj *MSMinObj) GetStats() *stats.DiceStats {
	return minObj.diceStats
}
