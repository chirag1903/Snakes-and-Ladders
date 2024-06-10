package objects

import (
	"Users/chirag.soni/druva/snakesandladders/stats"

	"github.com/pkg/errors"
)

// BoardObjectsInfo ...
type BoardObjectsInfo struct {
	snakes         []*SnakeInfo
	ladders        []*LadderInfo
	players        []*PlayerInfo
	crocodiles     []*CrocodileInfo
	mines          []*MineInfo
	gameStats      *stats.GameStats
	playerStats    *stats.PlayerStats
	totalBoardSize int
}

// NewBoardObjects ...
func NewBoardObjects(snakeInfo []*SnakeInfo, ladderInfo []*LadderInfo, playerInfo []*PlayerInfo, crocodiles []*CrocodileInfo, mines []*MineInfo, totalBoardSize int) *BoardObjectsInfo {
	return &BoardObjectsInfo{
		snakes:         snakeInfo,
		ladders:        ladderInfo,
		players:        playerInfo,
		crocodiles:     crocodiles,
		mines:          mines,
		gameStats:      &stats.GameStats{},
		playerStats:    &stats.PlayerStats{},
		totalBoardSize: totalBoardSize,
	}
}

// BoardObjects interface implemented
type BoardObjects interface {
	GetNewBoardPosition(pos, originalPos int, playerName string) (int, error)
	LogStats() *stats.GameStats
}

// GetBoardObject returns board object if present on that position
func (obj *BoardObjectsInfo) GetBoardObject(pos int, playerName string) (bool, BoardObjects, int, error) {
	if pos > obj.totalBoardSize {
		return false, nil, 0, errors.Errorf("invalid position input to fetch board object %d", pos)
	}

	isPresent := false
	objName := ""
	var boardObject BoardObjects
	originalPos := 0
	for _, player := range obj.players {
		if player.Name != playerName {
			if player.Position == pos {
				isPresent = true
				objName = Player
				boardObject = NewPlayerObj(obj.players, obj.totalBoardSize, objName)
			}
		} else {
			originalPos = player.Position
		}
	}

	// use map for snake and ladder
	for _, snake := range obj.snakes {
		if snake.Head == pos {
			isPresent = true
			if objName != "" {
				return false, nil, 0, errors.Errorf("invalid object present at snake head %d", pos)
			}
			objName = Snake
			boardObject = NewSnakeObj(obj.snakes, obj.totalBoardSize, objName)
		}
	}

	for _, ladder := range obj.ladders {
		if ladder.Bottom == pos {
			isPresent = true
			if objName != "" {
				return false, nil, 0, errors.Errorf("invalid object present at ladder bottom %d", pos)
			}
			objName = Ladder
			boardObject = NewLadderObj(obj.ladders, obj.totalBoardSize, objName)
		}
	}

	for _, crocodile := range obj.crocodiles {
		if crocodile.CrocPos == pos {
			isPresent = true
			if objName != "" {
				return false, nil, 0, errors.Errorf("invalid object present at crocodile pos %d", pos)
			}
			objName = Crocodile
			boardObject = NewCrocodileObj(obj.crocodiles, obj.totalBoardSize, objName)
		}
	}

	for _, mine := range obj.mines {
		if mine.MinePos == pos {
			isPresent = true
			if objName != "" {
				return false, nil, 0, errors.Errorf("invalid object present at mine pos %d", pos)
			}
			objName = Mine
			boardObject = NewMineObj(obj.mines, obj.totalBoardSize, objName)
		}
	}

	obj.playerStats.Name = playerName
	obj.playerStats.RollValue = pos - originalPos
	obj.playerStats.StartPos = originalPos
	obj.playerStats.EndPos = pos
	return isPresent, boardObject, originalPos, nil
}

// GetPlayerStats ...
func (obj *BoardObjectsInfo) GetPlayerStats() *stats.PlayerStats {
	return obj.playerStats
}
