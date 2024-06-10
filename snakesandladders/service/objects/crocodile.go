package objects

import (
	"Users/chirag.soni/druva/snakesandladders/stats"
	"fmt"

	"github.com/pkg/errors"
)

// CrocodileInfo ...
type CrocodileInfo struct {
	CrocPos int `yaml:"pos"`
}

// CrocodileObjectInfo ...
type CrocodileObjectInfo struct {
	crocodiles     []*CrocodileInfo
	gameStats      *stats.GameStats
	playerStats    *stats.PlayerStats
	totalBoardSize int
	objectName     string
}

// GetNewBoardPosition for snakes
func (sObj *CrocodileObjectInfo) GetNewBoardPosition(pos, origPos int, playerName string) (int, error) {
	sObj.gameStats.Objectname = sObj.objectName
	sObj.gameStats.Position = pos
	sObj.gameStats.PlayerName = playerName
	sObj.gameStats.Collided = true

	for _, crocodile := range sObj.crocodiles {
		if pos == crocodile.CrocPos {
			sObj.playerStats.Name = playerName
			sObj.playerStats.RollValue = pos - origPos
			sObj.playerStats.StartPos = origPos
			endPos := pos
			if pos-5 >= 1 {
				endPos = pos - 5
				sObj.gameStats.RepeatStep = true
			}

			sObj.playerStats.EndPos = endPos
			return endPos, nil
		}
	}

	return 0, errors.Errorf("crocodile pos expected for this positon %d", pos)
}

// LogStats for crocodiles
func (sObj *CrocodileObjectInfo) LogStats() *stats.GameStats {
	fmt.Printf("%s rolled a %d and bitten by crocodile at %d and moved from %d to %d\n",
		sObj.playerStats.Name, sObj.playerStats.RollValue, sObj.gameStats.Position,
		sObj.gameStats.Position, sObj.playerStats.EndPos)

	return sObj.gameStats
}

// NewCrocodileObj ...
func NewCrocodileObj(crocodileInfo []*CrocodileInfo, totalBoardSize int, objName string) BoardObjects {
	return &CrocodileObjectInfo{
		crocodiles:     crocodileInfo,
		gameStats:      &stats.GameStats{},
		playerStats:    &stats.PlayerStats{},
		totalBoardSize: totalBoardSize,
		objectName:     objName,
	}
}
