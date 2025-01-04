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

func (g *Game) StartGame() {
	var y, x int
	var winCount = 0
	g.FmtRule()
	for {
		fmt.Println("---------------------")
		g.FmtGame()
		fmt.Println("Выберете куда сходить: y x")
		fmt.Print("y :")
		fmt.Fscan(os.Stdin, &y)
		fmt.Print("x :")
		fmt.Fscan(os.Stdin, &x)
		if (y < 1 || y > 3) || (x < 1 || x > 3) {
			fmt.Println("Неправильный диапазон")
			continue
		}
		if g.State[y-1][x-1] != "_" {
			fmt.Printf("Ай, так нельзя, там уже %s\n", g.State[y-1][x-1])
			continue
		}

		g.State[y-1][x-1] = g.Current

		if g.isWin() {
			g.FmtGame()
			fmt.Printf("Победил %s", g.Current)
			break
		}
		// Наверное можно switch-case
		if g.Current == "X" {
			g.Current = "O"
		} else {
			g.Current = "X"
		}
		winCount++
		// Проверка на ничью
		if winCount == 9 {
			fmt.Println("Ничья!")
			g.FmtGame()
		}
	}
}

// func (g *Game) isLegal(y, x int) bool {

// }

func (g *Game) isWin() bool {
	var n int = 3
	// var winCount = 0
	// i - Это строка y
	// j - Это столбец x

	// Проверка по строкам
	for i := 0; i < n; i++ {
		if g.State[i][1] == g.State[i][0] && g.State[i][1] == g.State[i][2] && g.State[i][1] != "_" {
			fmt.Printf("Проверка по строкам %d", i)
			return true
		}
	}

	// Проверка по столбцам
	for i := 0; i < n; i++ {
		if g.State[1][i] == g.State[0][i] && g.State[1][i] == g.State[2][i] && g.State[1][i] != "_" {
			fmt.Printf("Проверка по столбцам %d", i)
			return true
		}
	}

	// Проверка по горизонтали
	if ((g.State[1][1] == g.State[0][0] && g.State[1][1] == g.State[2][2]) || (g.State[1][1] == g.State[2][0] && g.State[1][1] == g.State[0][2])) && g.State[1][1] != "_" {
		fmt.Printf("Проверка по горизонтали ")
		return true
	}

	// Проверку на ничью было бы неплохо сюда
	return false

}

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

func (g *Game) FmtRule() {
	fmt.Println(`
	Правила: 
	1) разрешённый диапазон значений: 1-3 (включительно)
	2) Стандартные правила крестики-нолики
	3) Первым ходит X
	
	`)
}
