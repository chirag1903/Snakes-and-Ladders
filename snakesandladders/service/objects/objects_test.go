package objects

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoardObjectsInfo(t *testing.T) {

	tests := []struct {
		name       string
		obj        *BoardObjectsInfo
		pos        int
		playerName string
		want       bool
		wantOrgPos int
		wantNewPos int
		wantWait   int
		wantErr    bool
	}{
		{
			name: "success for players",
			obj: &BoardObjectsInfo{
				players: []*PlayerInfo{
					{
						Name:     "ABC",
						Position: 10,
					},
					{
						Name:     "DEF",
						Position: 10,
					},
				},
				totalBoardSize: 100,
			},
			pos:        10,
			playerName: "ABC",
			want:       true,
			wantOrgPos: 10,
			wantNewPos: 10,
		},
		{
			name: "success for snakes",
			obj: &BoardObjectsInfo{
				players: []*PlayerInfo{
					{
						Name:     "ABC",
						Position: 15,
					},
					{
						Name:     "DEF",
						Position: 10,
					},
				},
				snakes: []*SnakeInfo{
					{
						Head: 15,
						Tail: 5,
					},
				},
				totalBoardSize: 100,
			},
			pos:        15,
			playerName: "ABC",
			want:       true,
			wantOrgPos: 15,
			wantNewPos: 5,
		},
		{
			name: "success for ladders",
			obj: &BoardObjectsInfo{
				players: []*PlayerInfo{
					{
						Name:     "ABC",
						Position: 15,
					},
					{
						Name:     "DEF",
						Position: 10,
					},
				},
				ladders: []*LadderInfo{
					{
						Bottom: 15,
						Top:    21,
					},
				},
				totalBoardSize: 100,
			},
			pos:        15,
			playerName: "ABC",
			want:       true,
			wantOrgPos: 15,
			wantNewPos: 21,
		},
		{
			name: "success for crocodiles",
			obj: &BoardObjectsInfo{
				players: []*PlayerInfo{
					{
						Name:     "ABC",
						Position: 15,
					},
					{
						Name:     "DEF",
						Position: 10,
					},
				},
				crocodiles: []*CrocodileInfo{
					{
						CrocPos: 15,
					},
				},
				totalBoardSize: 100,
			},
			pos:        15,
			playerName: "ABC",
			want:       true,
			wantOrgPos: 15,
			wantNewPos: 10,
		},
		{
			name: "success for mines",
			obj: &BoardObjectsInfo{
				players: []*PlayerInfo{
					{
						Name:     "ABC",
						Position: 15,
					},
					{
						Name:     "DEF",
						Position: 10,
					},
				},
				mines: []*MineInfo{
					{
						MinePos: 15,
					},
				},
				totalBoardSize: 100,
			},
			pos:        15,
			playerName: "ABC",
			want:       true,
			wantOrgPos: 15,
			wantNewPos: 15,
			wantWait:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boardObjs := NewBoardObjects(tt.obj.snakes, tt.obj.ladders, tt.obj.players, tt.obj.crocodiles, tt.obj.mines, tt.obj.totalBoardSize)

			isPresent, boardObj, origPos, err := boardObjs.GetBoardObject(tt.pos, tt.playerName)
			if (err != nil) != tt.wantErr {
				t.Fatalf("BoardObjectsInfo.GetBoardObject() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if isPresent != tt.want {
					t.Fatalf("BoardObjectsInfo.GetBoardObject() got = %v, want %v", isPresent, tt.want)
				}

				if origPos != tt.wantOrgPos {
					t.Fatalf("BoardObjectsInfo.GetBoardObject() origPos = %v, want %v", origPos, tt.wantOrgPos)
				}
				newPos, err := boardObj.GetNewBoardPosition(tt.pos, origPos, tt.playerName)
				require.NoError(t, err)

				stats := boardObj.LogStats()
				if newPos != tt.wantNewPos {
					t.Fatalf("BoardObjectsInfo.GetBoardObject() newPos = %v, want %v", newPos, tt.wantNewPos)
				}

				if stats.HoldTurn != tt.wantWait {
					t.Fatalf("BoardObjectsInfo.GetBoardObject() holdturn = %v, want %v", stats.HoldTurn, tt.wantWait)
				}
			}
		})
	}
}
