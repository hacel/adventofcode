package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hacel/adventofcode/internal/scan"
)

const (
	HEIGHT = 5
	WIDTH  = 5
)

type Number struct {
	n      int
	marked bool
}

type Board struct {
	grid [HEIGHT][WIDTH]*Number
	won  bool
}

func (b *Board) Mark(n int) {
	for _, row := range b.grid {
		for _, num := range row {
			if num.n == n {
				num.marked = true
			}
		}
	}
}

func (b *Board) SumUnmarked() int {
	sum := 0
	for _, row := range b.grid {
		for _, num := range row {
			if !num.marked {
				sum += num.n
			}
		}
	}
	return sum
}

func (b *Board) Won() bool {
	if b.won {
		return b.won
	}

	vertical := [HEIGHT]int{}
	horizontal := [WIDTH]int{}
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if b.grid[y][x].marked {
				horizontal[x]++
				if horizontal[x] == WIDTH {
					b.won = true
					return true
				}
				vertical[y]++
				if vertical[y] == WIDTH {
					b.won = true
					return true
				}
			}
		}
	}
	return false
}

func main() {
	lines := scan.S("4/input.txt")
	var numbers []int
	for _, s := range strings.Split(<-lines, ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	<-lines

	// Initialize boards
	boards := []*Board{}
	board := &Board{}
	y := 0
	for line := range lines {
		if line == "" {
			boards = append(boards, board)
			board = &Board{}
			y = 0
			continue
		}
		for x, s := range strings.Fields(line) {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			board.grid[y][x] = &Number{n: n}
		}
		y++
	}
	boards = append(boards, board)

	wonBoards := []*Board{}
	for _, n := range numbers {
		for i, b := range boards {
			if b.Won() {
				continue
			}
			b.Mark(n)
			if b.Won() {
				wonBoards = append(wonBoards, b)
			}
			if len(wonBoards) == len(boards) {
				b := wonBoards[len(wonBoards)-1]
				fmt.Printf("Board #%d has won last\n", i+1)
				fmt.Printf("Last called number: %d\n", n)
				fmt.Printf("Sum of unmarked numbers: %d\n", b.SumUnmarked())
				fmt.Printf("%d * %d = %d\n", b.SumUnmarked(), n, b.SumUnmarked()*n)
				return
			}
		}
	}
}
