package main

import (
	"fmt"

	"github.com/hacel/adventofcode/internal/scan"
)

const WIDTH = 12

func main() {
	var n0, n1 [WIDTH]int
	for line := range scan.B("3/input.txt") {
		for i, c := range line {
			if c == '0' {
				n0[i]++
			} else if c == '1' {
				n1[i]++
			}
		}
	}

	var gammaRate int
	for i, n := range n1 {
		if n >= n0[i] {
			gammaRate = gammaRate | 1<<(WIDTH-1-i)
		}
	}
	epsilonRate := (^gammaRate) & 0xfff

	fmt.Printf("Gamma rate: %d\n", gammaRate)
	fmt.Printf("Epsilon rate: %d\n", epsilonRate)
	fmt.Printf("Power consumption: %d\n", gammaRate*epsilonRate)
}
