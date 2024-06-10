package initconfig

import (
	"os"

	"Users/chirag.soni/druva/snakesandladders/service/objects"

	"gopkg.in/yaml.v2"
)

// ManualInputInfo ...
type ManualInputInfo struct {
	PlayerName string
	RollValues []int
}

// Config ...
type Config struct {
	Board            *BoardParams             `yaml:"board"`
	EntityCount      *EntityCount             `yaml:"count"`
	Snakes           []*objects.SnakeInfo     `yaml:"snakes"`
	Ladders          []*objects.LadderInfo    `yaml:"ladders"`
	Players          []*objects.PlayerInfo    `yaml:"players"`
	Crocodiles       []*objects.CrocodileInfo `yaml:"crocodiles"`
	Mines            []*objects.MineInfo      `yaml:"mines"`
	DiceCount        int                      `yaml:"diceCount"`
	MovementStrategy string                   `yaml:"ms"`
}

// BoardParams ...
type BoardParams struct {
	Size int `yaml:"size"`
}

// EntityCount ...
type EntityCount struct {
	Snakes  int `yaml:"snakes"`
	Ladders int `yaml:"ladders"`
	Players int `yaml:"players"`
}

// PrepareConfigFromYml ...
func PrepareConfigFromYml() (*Config, error) {
	ymlData, err := os.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(ymlData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
