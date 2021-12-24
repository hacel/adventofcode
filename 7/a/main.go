package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/hacel/adventofcode/internal/scan"
)

func Median(d []int) int {
	sort.Ints(d)
	return d[(len(d)+1)/2]
}

func main() {
	var crabs []int
	for s := range scan.S("7/input.txt") {
		for _, v := range strings.Split(s, ",") {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			crabs = append(crabs, n)
		}
	}

	cheapestPosition := Median(crabs)
	cost := 0
	for _, c := range crabs {
		cost += int(math.Abs(float64(cheapestPosition - c)))
	}

	fmt.Printf("Best position: %d\n", cheapestPosition)
	fmt.Printf("Cost: %d\n", cost)
}
