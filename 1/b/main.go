package main

import (
	"fmt"

	"github.com/hacel/adventofcode/internal/scan"
)

func main() {
	w := map[int]int{}
	inc := 0
	i := 0
	for n := range scan.I("1/input.txt") {
		w[i] += n
		if i > 0 {
			w[i-1] += n
		}
		if i > 1 {
			w[i-2] += n
		}
		if i > 2 {
			if w[i-2] > w[i-3] {
				inc++
			}
		}
		i++
	}
	fmt.Println(inc)
}
