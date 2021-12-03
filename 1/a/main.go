package main

import (
	"fmt"
	"math"

	"github.com/hacel/adventofcode/internal/scan"
)

func main() {
	inc := 0
	prev := math.MaxInt64
	for n := range scan.I("1/input.txt") {
		if n > prev {
			inc++
		}
		prev = n
	}
	fmt.Println(inc)
}
