package objects

import (
	"Users/chirag.soni/druva/snakesandladders/stats"
	"fmt"

	"github.com/pkg/errors"
)

// MineInfo ...
type MineInfo struct {
	MinePos int `yaml:"pos"`
}

// MineObjectInfo ...
type MineObjectInfo struct {
	mines          []*MineInfo
	gameStats      *stats.GameStats
	playerStats    *stats.PlayerStats
	totalBoardSize int
	objectName     string
}

// GetNewBoardPosition for snakes
func (mineObj *MineObjectInfo) GetNewBoardPosition(pos, origPos int, playerName string) (int, error) {
	mineObj.gameStats.Objectname = mineObj.objectName
	mineObj.gameStats.Position = pos
	mineObj.gameStats.PlayerName = playerName
	mineObj.gameStats.Collided = true

	for _, mine := range mineObj.mines {
		if pos == mine.MinePos {
			mineObj.playerStats.Name = playerName
			mineObj.playerStats.RollValue = pos - origPos
			mineObj.playerStats.StartPos = origPos
			mineObj.playerStats.EndPos = pos

			// if mine is encountered hold turn for 2 times
			mineObj.gameStats.HoldTurn = 2
			return pos, nil
		}
	}

	return 0, errors.Errorf("mine pos expected for this positon %d", pos)
}

// LogStats for mines
func (mineObj *MineObjectInfo) LogStats() *stats.GameStats {
	fmt.Printf("%s rolled a %d and hit by mine at %d and moved from %d to %d\n",
		mineObj.playerStats.Name, mineObj.playerStats.RollValue, mineObj.gameStats.Position,
		mineObj.gameStats.Position, mineObj.playerStats.EndPos)

	return mineObj.gameStats
}

// NewMineObj ...
func NewMineObj(mineInfo []*MineInfo, totalBoardSize int, objName string) BoardObjects {
	return &MineObjectInfo{
		mines:          mineInfo,
		gameStats:      &stats.GameStats{},
		playerStats:    &stats.PlayerStats{},
		totalBoardSize: totalBoardSize,
		objectName:     objName,
	}
}
