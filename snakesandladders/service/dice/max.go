package dice

import (
	"Users/chirag.soni/druva/snakesandladders/stats"

	"github.com/pkg/errors"
)

// MSMaxObj ...
type MSMaxObj struct {
	diceCount int
	diceStats *stats.DiceStats
}

// NewMSMaxObj ...
func NewMSMaxObj(diceCount int, diceStats *stats.DiceStats) DiceRoll {
	return &MSMaxObj{
		diceCount: diceCount,
		diceStats: diceStats,
	}
}

// GetMovementCount for sum ms
func (maxObj *MSMaxObj) GetMovementCount(diceValues []int) (int, error) {
	if len(diceValues) == 0 {
		return 0, errors.Errorf("dice values should not be")
	}

	if len(diceValues) != maxObj.diceCount {
		return 0, errors.Errorf("invalid dice count and values %d & %d", len(diceValues), maxObj.diceCount)
	}

	maxVal := 0
	for _, val := range diceValues {
		if val > maxVal {
			maxVal = val
		}
	}

	maxObj.diceStats.TotalMovement++
	return maxVal, nil
}

// GetStats ...
func (maxObj *MSMaxObj) GetStats() *stats.DiceStats {
	return maxObj.diceStats
}
