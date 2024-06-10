package objects

import (
	"Users/chirag.soni/druva/snakesandladders/stats"
	"fmt"

	"github.com/pkg/errors"
)

// LadderInfo ...
type LadderInfo struct {
	Bottom int `yaml:"bottom"`
	Top    int `yaml:"top"`
}

// LadderObjectInfo ...
type LadderObjectInfo struct {
	ladders        []*LadderInfo
	gameStats      *stats.GameStats
	playerStats    *stats.PlayerStats
	totalBoardSize int
	objectName     string
}

// GetNewBoardPosition ladder
func (ladderObj *LadderObjectInfo) GetNewBoardPosition(pos, origPos int, playerName string) (int, error) {
	ladderObj.gameStats.Objectname = ladderObj.objectName
	ladderObj.gameStats.Position = pos
	ladderObj.gameStats.PlayerName = playerName
	ladderObj.gameStats.Collided = true

	for _, ladder := range ladderObj.ladders {
		if pos == ladder.Bottom {
			ladderObj.playerStats.Name = playerName
			ladderObj.playerStats.RollValue = pos - origPos
			ladderObj.playerStats.StartPos = origPos
			ladderObj.playerStats.EndPos = ladder.Top
			return ladder.Top, nil
		}
	}

	return 0, errors.Errorf("ladder bottom expected for this positon %d", pos)
}

// LogStats for ladder
func (ladderObj *LadderObjectInfo) LogStats() *stats.GameStats {
	fmt.Printf("%s rolled a %d and climbed the ladder at %d and moved from %d to %d\n",
		ladderObj.playerStats.Name, ladderObj.playerStats.RollValue, ladderObj.gameStats.Position,
		ladderObj.gameStats.Position, ladderObj.playerStats.EndPos)

	return ladderObj.gameStats
}

// NewLadderObj ...
func NewLadderObj(ladderInfo []*LadderInfo, totalBoardSize int, objName string) BoardObjects {
	return &LadderObjectInfo{
		ladders:        ladderInfo,
		gameStats:      &stats.GameStats{},
		playerStats:    &stats.PlayerStats{},
		totalBoardSize: totalBoardSize,
		objectName:     objName,
	}
}
