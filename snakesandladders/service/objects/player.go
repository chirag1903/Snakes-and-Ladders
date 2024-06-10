package objects

import (
	"Users/chirag.soni/druva/snakesandladders/stats"
	"fmt"
)

// PlayerInfo ...
type PlayerInfo struct {
	Name     string `yaml:"name"`
	Position int    `yaml:"position"` //starting positon
	Wait     int    `yaml:"wait"`
}

// PlayerObjectInfo ...
type PlayerObjectInfo struct {
	players        []*PlayerInfo
	gameStats      *stats.GameStats
	playerStats    *stats.PlayerStats
	totalBoardSize int
	objectName     string
}

// previous pos player pos is set to 1
func (playerObj *PlayerObjectInfo) updatePrevPlayerPos(playerInfo []*PlayerInfo, pos int, playerName string) {
	replaceablePlayerName := ""
	for _, player := range playerInfo {
		if playerName != player.Name && player.Position == pos {
			replaceablePlayerName = player.Name
			player.Position = 1
			break
		}
	}

	fmt.Printf("Player: %s, moved to pos 1 due to collison\n", replaceablePlayerName)
}

// GetNewBoardPosition for players
func (playerObj *PlayerObjectInfo) GetNewBoardPosition(pos, origPos int, playerName string) (int, error) {
	playerObj.gameStats.Objectname = playerObj.objectName
	playerObj.gameStats.Position = pos
	playerObj.gameStats.PlayerName = playerName
	playerObj.gameStats.Collided = true

	playerObj.updatePrevPlayerPos(playerObj.players, pos, playerName)
	playerObj.playerStats.Name = playerName
	playerObj.playerStats.RollValue = pos - origPos
	playerObj.playerStats.StartPos = origPos
	playerObj.playerStats.EndPos = pos
	return pos, nil
}

// LogStats for players
func (playerObj *PlayerObjectInfo) LogStats() *stats.GameStats {
	fmt.Printf("%s rolled a %d and collided with player at %d and moved from %d to %d\n",
		playerObj.playerStats.Name, playerObj.playerStats.RollValue, playerObj.gameStats.Position,
		playerObj.playerStats.StartPos, playerObj.playerStats.EndPos)

	return playerObj.gameStats
}

// NewPlayerObj ...
func NewPlayerObj(playerInfo []*PlayerInfo, totalBoardSize int, objName string) BoardObjects {
	return &PlayerObjectInfo{
		players:        playerInfo,
		gameStats:      &stats.GameStats{},
		playerStats:    &stats.PlayerStats{},
		totalBoardSize: totalBoardSize,
		objectName:     objName,
	}
}
