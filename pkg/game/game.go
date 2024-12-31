package game

import (
	"fmt"
	"os"
)

type Game struct {
	State   [3][3]string
	Current string
}

func New() *Game {
	return &Game{State: [3][3]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}, Current: "X"}
}

// Пофиксить строку с столбцом. Непонятно что где
func (g *Game) StartGame() {
	var stroka, stolb int
	for true {
		fmt.Println("---------------------")
		g.FmtGame()
		fmt.Println("Выберете куда сходить: x y")
		fmt.Print("x :")
		fmt.Fscan(os.Stdin, &stroka)
		fmt.Print("y :")
		fmt.Fscan(os.Stdin, &stolb)
		if (stroka < 1 || stroka > 3) || (stolb < 1 || stolb > 3) {
			fmt.Println("Неправильный диапазон")
			continue
		}
		if g.State[stroka-1][stolb-1] != "_" {
			fmt.Printf("Ай, так нельзя, там уже %s\n", g.State[stroka-1][stolb-1])
			continue
		}
		g.State[stroka-1][stolb-1] = g.Current
		if g.Current == "X" {
			g.Current = "O"
		} else {
			g.Current = "X"
		}
	}
}

// func (g *Game) isLegal(x, y int) bool {

// }

// func (g *Game) isWin() bool {
// 	var n int = 3
// 	// i - Это строка
// 	// j - Это столбец
// 	for i := 0; i < n; i++ {

// 	}
// }

func (g *Game) FmtGame() {
	fmt.Println("Текущее состояние:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(g.State[i][j] + " ")
		}
		fmt.Println()
	}
	fmt.Printf("Текущий ход за игроком: %s\n", g.Current)
}
