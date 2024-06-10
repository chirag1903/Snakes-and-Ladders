package stats

// PlayerStats for player name, roll value, start and end pos
type PlayerStats struct {
	Name      string
	RollValue int
	StartPos  int
	EndPos    int
}

// GameStats for ladder, snake bite and player collison
// Collided with object at pos
type GameStats struct {
	Objectname string
	Position   int
	PlayerName string
	Collided   bool
	RepeatStep bool
	HoldTurn   int
}

// DiceStats for movement strategy, number
type DiceStats struct {
	MovementStrategy string
	TotalMovement    int
	RollNumbers      []int
}

// WinnerStatus for playername, simulation count, playerstats
type WinnerStatus struct {
	PlayerName  string
	TotalTurns  int
	PlayerStats *PlayerStats
}
