package dice

import (
	"Users/chirag.soni/druva/snakesandladders/stats"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiceInfo(t *testing.T) {
	tests := []struct {
		name       string
		dInfo      *DiceInfo
		diceValues []int
		want       int
		wantErr    bool
	}{
		{
			name: "success for sum",
			dInfo: &DiceInfo{
				movementStrategy: "SUM",
				diceCount:        2,
				diceStats:        &stats.DiceStats{},
			},
			diceValues: []int{3, 4},
			want:       7,
		},
		{
			name: "failure: invalid dice count",
			dInfo: &DiceInfo{
				movementStrategy: "SUM",
				diceCount:        3,
				diceStats:        &stats.DiceStats{},
			},
			diceValues: []int{3, 4},
			wantErr:    true,
		},
		{
			name: "failure: invalid dice values",
			dInfo: &DiceInfo{
				movementStrategy: "SUM",
				diceCount:        3,
				diceStats:        &stats.DiceStats{},
			},
			wantErr: true,
		},
		{
			name: "success for max",
			dInfo: &DiceInfo{
				movementStrategy: "MAX",
				diceCount:        2,
				diceStats:        &stats.DiceStats{},
			},
			diceValues: []int{3, 4},
			want:       4,
		},
		{
			name: "success for min",
			dInfo: &DiceInfo{
				movementStrategy: "MIN",
				diceCount:        2,
				diceStats:        &stats.DiceStats{},
			},
			diceValues: []int{3, 4},
			want:       3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diceObj, err := tt.dInfo.GetDiceMovementCount()
			require.NoError(t, err)

			got, err := diceObj.GetMovementCount(tt.diceValues)
			if (err != nil) != tt.wantErr {
				t.Fatalf("DiceInfo.GetDiceMovementCount() error = %v, wantErr %v", err, tt.wantErr)
			}

			require.Equal(t, tt.want, got, "DiceInfo.GetDiceMovementCount() error")
		})
	}
}
