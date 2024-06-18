package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type data struct {
	val   string
	count int
}

type Board struct {
	rows    []*data
	cols    []*data
	diagnal []*data
}

func (b *Board) fillRow(i int, v string) bool {
	i -= 1
	//fmt.Printf("fillRow fill i: %d\n", i)
	if i >= len(b.rows) {
		return false
	}
	lastData := b.rows[i]
	if lastData == nil {
		b.rows[i] = &data{
			val:   v,
			count: 1,
		}
	} else if lastData.val == v {
		b.rows[i].count += 1
	} else {
		b.rows[i] = &data{
			val:   v,
			count: 1,
		}
	}
	return b.rows[i].count >= len(b.rows)
}

func (b *Board) fillCol(i int, v string) bool {
	i -= 1
	//fmt.Printf("fillCol fill i: %d\n", i)
	if i >= len(b.cols) {
		return false
	}
	lastData := b.cols[i]
	if lastData == nil {
		b.cols[i] = &data{
			val:   v,
			count: 1,
		}
	} else if lastData.val == v {
		b.cols[i].count += 1
	} else {
		b.cols[i] = &data{
			val:   v,
			count: 1,
		}
	}
	return b.cols[i].count >= len(b.rows)
}

func (b *Board) fillDiagnal(i int, j int, v string) bool {
	i -= 1
	j -= 1
	//fmt.Printf("Diagnal fill i: %d j: %d\n", i, j)
	if i != j && i != len(b.rows)-j-1 {
		return false
	} else if i == j {
		lastData := b.diagnal[0]
		if lastData == nil {
			b.diagnal[0] = &data{
				val:   v,
				count: 1,
			}
		} else if lastData.val == v {
			b.diagnal[0].count += 1
		} else {
			b.diagnal[0] = &data{
				val:   v,
				count: 1,
			}
		}
		return b.diagnal[0].count >= len(b.rows)
	} else {
		lastData := b.diagnal[1]
		if lastData == nil {
			b.diagnal[1] = &data{
				val:   v,
				count: 1,
			}
		} else if lastData.val == v {
			b.diagnal[1].count += 1
		} else {
			b.diagnal[1] = &data{
				val:   v,
				count: 1,
			}
		}
		return b.diagnal[1].count >= len(b.rows)
	}
}

func main() {
	/*
		Tic-tac-toe is a game played on an n × n board (typically n = 3) where two
		players take consecutive turns placing “O” and “X” marks onto the board cells.
		The game is won if n consecutive “O” or ‘X” marks are placed in a row, column,
		or diagonal. Create a data structure with O(n) space that accepts a sequence
		of moves, and reports in constant time whether the last move won the game.
	*/
	var n int = 3
	var board *Board = &Board{
		rows:    make([]*data, n),
		cols:    make([]*data, n),
		diagnal: make([]*data, 2),
	}

	for i := 1; i <= n*n; i++ {
		var player int
		if i%2 == 1 {
			player = 1
		} else {
			player = 2
		}
		fmt.Printf("Please make your move player %d :\n", player)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		inputs := strings.Split(s, " ")
		r, err := strconv.Atoi(inputs[0])
		if err != nil || r > n {
			fmt.Println("some error in getting your input, please try again")
			i--
			continue
		}
		c, err := strconv.Atoi(inputs[1])
		if err != nil || c > n {
			fmt.Println("some error in getting your input, please try again")
			i--
			continue
		}
		v := inputs[2]
		if (v != "x" && v != "o") || (v == "x" && player != 1) || (v == "o" && player != 2) {
			fmt.Println("some error in getting your input, please try again")
			i--
			continue
		}
		won := board.fillRow(r, v)
		if won {
			fmt.Printf("Player %d has won the game by filling row %d with %s\n", player, r, v)
			break
		}
		won = board.fillCol(c, v)
		if won {
			fmt.Printf("Player %d has won the game by filling col %d with %s\n", player, r, v)
			break
		}
		won = board.fillDiagnal(r, c, v)
		if won {
			fmt.Printf("Player %d has won the game by filling diagnal %d with %s\n", player, r, v)
			break
		}
	}
	fmt.Println("Game over.")

}
