package game

import (
	"os"
	"strconv"
	"testing"
)

func genPlayers(count int) []Player {
	var players []Player
	for i := 0; i < count; i++ {
		players = append(players, NewPlayer(i, "player"+strconv.Itoa(i)))
	}
	return players
}

func randomKill(game *Game) *Player {
	goodMan := GetGoodManLeft(game)
	return goodMan[0]
}

func randomCheck(game *Game) *Player {
	man := GetManLeft(game)
	return man[0]
}

func randomPoison(game *Game) *Player {
	man := GetManLeft(game)
	return man[0]
}

func TestGameFlow(t *testing.T) {
	game := New()

	players := genPlayers(game.PlayersCount)
	for _, p := range players {
		game.JoinPlayer(p)
	}

	go StartGame(&game)

	for {
		select {
		case s := <-game.Lifecycle:
			if s == WaitingPlayerAction {
				game.PlayerActions.WerewolfKill <- randomKill(&game)
				game.PlayerActions.SeerCheck <- randomCheck(&game)
				if game.RoundNumber == 1 {
					game.PlayerActions.UsePoison <- randomPoison(&game)
				} else {
					game.PlayerActions.UsePoison <- nil
				}
				if game.RoundNumber == 2 {
					game.PlayerActions.UseAntidote <- true
				} else {
					game.PlayerActions.UseAntidote <- false
				}
			}
			if s == Over {
				os.Exit(0)
			}
		}
	}
}