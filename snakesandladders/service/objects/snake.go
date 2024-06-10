package objects

import (
	"Users/chirag.soni/druva/snakesandladders/stats"
	"fmt"

	"github.com/pkg/errors"
)

// SnakeInfo ...
type SnakeInfo struct {
	Head int `yaml:"head"`
	Tail int `yaml:"tail"`
}

// SnakeObjectInfo ...
type SnakeObjectInfo struct {
	snakes         []*SnakeInfo
	gameStats      *stats.GameStats
	playerStats    *stats.PlayerStats
	totalBoardSize int
	objectName     string
}

// GetNewBoardPosition for snakes
func (sObj *SnakeObjectInfo) GetNewBoardPosition(pos, origPos int, playerName string) (int, error) {
	sObj.gameStats.Objectname = sObj.objectName
	sObj.gameStats.Position = pos
	sObj.gameStats.PlayerName = playerName
	sObj.gameStats.Collided = true

	for _, snake := range sObj.snakes {
		if pos == snake.Head {
			sObj.playerStats.Name = playerName
			sObj.playerStats.RollValue = pos - origPos
			sObj.playerStats.StartPos = origPos
			sObj.playerStats.EndPos = snake.Tail
			return snake.Tail, nil
		}
	}

	return 0, errors.Errorf("snake head expected for this positon %d", pos)
}

// LogStats for snakes
func (sObj *SnakeObjectInfo) LogStats() *stats.GameStats {
	fmt.Printf("%s rolled a %d and bitten by snake at %d and moved from %d to %d\n",
		sObj.playerStats.Name, sObj.playerStats.RollValue, sObj.gameStats.Position,
		sObj.gameStats.Position, sObj.playerStats.EndPos)

	return sObj.gameStats
}

// NewSnakeObj ...
func NewSnakeObj(snakeInfo []*SnakeInfo, totalBoardSize int, objName string) BoardObjects {
	return &SnakeObjectInfo{
		snakes:         snakeInfo,
		gameStats:      &stats.GameStats{},
		playerStats:    &stats.PlayerStats{},
		totalBoardSize: totalBoardSize,
		objectName:     objName,
	}
}
