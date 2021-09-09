package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

func loadOpponent() *uci.Engine {
	if _, err := exec.LookPath("stockfish"); err != nil {
		return nil
	}

	e, err := uci.New("stockfish") // you must have stockfish installed and on $PATH
	if err != nil {
		panic(err)
	}

	if err := e.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}
	return e
}

func playResponse(u *ui) {
	var m *chess.Move
	if u.eng != nil {
		cmdPos := uci.CmdPosition{Position: u.game.Position()}
		cmdGo := uci.CmdGo{MoveTime: time.Millisecond}
		if err := u.eng.Run(cmdPos, cmdGo); err != nil {
			panic(err)
		}

		m = u.eng.SearchResults().BestMove
	} else {
		m = AI(u.game)
	}
	if m == nil {
		return // somehow end of game and we didn't notice?
	}

	off := squareToOffset(m.S1())
	cell := u.grid.objects[off].(*fyne.Container)

	u.over.Move(cell.Position())
	move(m, u.game, false, u)
	fmt.Println(u.game.FEN())
}

//func randomResponse(game *chess.Game) *chess.Move {
//	valid := game.ValidMoves()
//	if len(valid) == 0 {
//		return nil
//	}
//
//	return valid[rand.Intn(len(valid))]
//}

func AI(game *chess.Game) *chess.Move {
	test := game.FEN()
	fmt.Println("pre move:")
	fenReturn(test)
	valid := game.ValidMoves()

	if len(valid) == 0 {
		return nil
	}
	testa := valid[rand.Intn(len(valid))]

	return testa
}

func fenReturn(fennotation string) []string {
	rows := strings.Split(fennotation, "/")

	lol := map[string]string{"1": " ", "2": "  ", "3": "   ", "4": "    ", "5": "     ", "6": "      ", "7": "       ", "8": "        "}
	for _, row := range rows {
		for index, element := range lol {
			res := strings.Contains(row, index)
			if res {
				done := strings.Replace(row, index, element, -1)
				fmt.Println(done)
			}
		}
	}
	return nil
}
