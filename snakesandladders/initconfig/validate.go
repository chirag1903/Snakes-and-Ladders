package initconfig

import (
	"github.com/pkg/errors"
)

// ValidateInputParams ..
func (c *Config) ValidateInputParams() error {
	//validate input count and positions
	err := c.validateCount()
	if err != nil {
		return err
	}

	snakeMap := make(map[int]int)
	ladderMap := make(map[int]int)

	// check snakes position
	err = c.validateSnakePos(snakeMap)
	if err != nil {
		return err
	}

	// check ladders position
	err = c.validateLadderPos(ladderMap)
	if err != nil {
		return err
	}

	// check snakes tail and ladder bottom
	// check ladder top and snake head
	err = c.validateSnakeAndLadderPos(snakeMap, ladderMap)
	if err != nil {
		return err
	}

	// check player position is not on snake head or ladder bottom
	// err = c.validatePlayerPosition(snakeMap, ladderMap)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (c *Config) validateCount() error {
	if len(c.Snakes) != c.EntityCount.Snakes {
		return errors.Errorf("Snakes count and positions doesnot match: %d and %d", len(c.Snakes), c.EntityCount.Snakes)
	}

	if len(c.Ladders) != c.EntityCount.Ladders {
		return errors.Errorf("Ladders count and positions doesnot match: %d and %d", len(c.Ladders), c.EntityCount.Ladders)
	}

	if len(c.Players) != c.EntityCount.Players {
		return errors.Errorf("Players count and positions doesnot match: %d and %d", len(c.Players), c.EntityCount.Players)
	}

	return nil
}

func (c *Config) validateSnakePos(snakeMap map[int]int) error {
	totalBoardSize := c.Board.Size * c.Board.Size

	for _, snakeInfo := range c.Snakes {
		if snakeInfo.Head < snakeInfo.Tail+1 {
			return errors.Errorf("Snake head should be after tail position: %d and %d", snakeInfo.Head, snakeInfo.Tail)
		}

		if snakeInfo.Head < 1 || snakeInfo.Tail < 1 {
			return errors.Errorf("Snake head or tail position should not be less than 1: %d and %d", snakeInfo.Head, snakeInfo.Tail)
		}

		if snakeInfo.Head > totalBoardSize || snakeInfo.Tail > totalBoardSize {
			return errors.Errorf("Snake head or tail position should not be more than %d: %d and %d", totalBoardSize, snakeInfo.Head, snakeInfo.Tail)
		}

		_, ok := snakeMap[snakeInfo.Head]
		if ok {
			return errors.Errorf("snake head cannot be on same position %d", snakeInfo.Head)
		}
		snakeMap[snakeInfo.Head] = snakeInfo.Tail
	}

	return nil
}

func (c *Config) validateLadderPos(ladderMap map[int]int) error {
	totalBoardSize := c.Board.Size * c.Board.Size

	for _, ladderInfo := range c.Ladders {
		if ladderInfo.Top < ladderInfo.Bottom+1 {
			return errors.Errorf("Ladder top should be after bottom position: %d and %d", ladderInfo.Top, ladderInfo.Bottom)
		}

		if ((ladderInfo.Top - 1) / c.Board.Size) == ((ladderInfo.Bottom - 1) / c.Board.Size) {
			return errors.Errorf("Ladder top and bottom should be in different rows: %d and %d", ladderInfo.Top, ladderInfo.Bottom)
		}

		if ladderInfo.Top < 1 || ladderInfo.Bottom < 1 {
			return errors.Errorf("Ladder top or bottom position should not be less than 1: %d and %d", ladderInfo.Top, ladderInfo.Bottom)
		}

		if ladderInfo.Top > totalBoardSize || ladderInfo.Bottom > totalBoardSize {
			return errors.Errorf("Ladder top or bottom position should not be more than %d: %d and %d", totalBoardSize, ladderInfo.Top, ladderInfo.Bottom)
		}
		_, ok := ladderMap[ladderInfo.Bottom]
		if ok {
			return errors.Errorf("Ladder bottom cannot be on same position %d", ladderInfo.Bottom)
		}
		ladderMap[ladderInfo.Bottom] = ladderInfo.Top
	}

	return nil
}

func (c *Config) validateSnakeAndLadderPos(snakeMap, ladderMap map[int]int) error {
	for head, tail := range snakeMap {
		if _, ok := ladderMap[head]; ok {
			return errors.Errorf("snakes tail and ladder bottom position cannot be same %d", tail)
		}

		if top, ok := ladderMap[tail]; ok {
			if top == head {
				return errors.Errorf("snakes head and ladder top position cannot be same %d", top)
			}
		}
	}

	return nil
}

func (c *Config) validatePlayerPosition(snakeMap, ladderMap map[int]int) error {
	for _, playerInfo := range c.Players {
		_, ok := snakeMap[playerInfo.Position]
		if ok {
			return errors.Errorf("player position cannot be on snake head %d", playerInfo.Position)
		}

		_, ok = ladderMap[playerInfo.Position]
		if ok {
			return errors.Errorf("player position cannot be on ladder bottom %d", playerInfo.Position)
		}
	}

	return nil
}
